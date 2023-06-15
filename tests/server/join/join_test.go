// Copyright 2018 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package join_test

import (
	"context"
	"os"
	"path"
	"testing"
	"time"

	"github.com/gottingen/tm/pkg/utils/etcdutil"
	"github.com/gottingen/tm/server"
	"github.com/gottingen/tm/server/join"
	"github.com/gottingen/tm/tests"
	"github.com/stretchr/testify/require"
)

// TODO: enable it when we fix TestFailedAndDeletedTMJoinsPreviousCluster
// func TestMain(m *testing.M) {
// 	goleak.VerifyTestMain(m, testutil.LeakOptions...)
// }

func TestSimpleJoin(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cluster, err := tests.NewTestCluster(ctx, 1)
	defer cluster.Destroy()
	re.NoError(err)

	err = cluster.RunInitialServers()
	re.NoError(err)
	cluster.WaitLeader()

	tm1 := cluster.GetServer("tm1")
	client := tm1.GetEtcdClient()
	members, err := etcdutil.ListEtcdMembers(client)
	re.NoError(err)
	re.Len(members.Members, 1)

	// Join the second TM.
	tm2, err := cluster.Join(ctx)
	re.NoError(err)
	err = tm2.Run()
	re.NoError(err)
	_, err = os.Stat(path.Join(tm2.GetConfig().DataDir, "join"))
	re.False(os.IsNotExist(err))
	members, err = etcdutil.ListEtcdMembers(client)
	re.NoError(err)
	re.Len(members.Members, 2)
	re.Equal(tm1.GetClusterID(), tm2.GetClusterID())

	// Wait for all nodes becoming healthy.
	time.Sleep(time.Second * 5)

	// Join another TM.
	tm3, err := cluster.Join(ctx)
	re.NoError(err)
	err = tm3.Run()
	re.NoError(err)
	_, err = os.Stat(path.Join(tm3.GetConfig().DataDir, "join"))
	re.False(os.IsNotExist(err))
	members, err = etcdutil.ListEtcdMembers(client)
	re.NoError(err)
	re.Len(members.Members, 3)
	re.Equal(tm1.GetClusterID(), tm3.GetClusterID())
}

// A failed TM tries to join the previous cluster but it has been deleted
// during its downtime.
func TestFailedAndDeletedTMJoinsPreviousCluster(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.EtcdStartTimeout = 10 * time.Second
	cluster, err := tests.NewTestCluster(ctx, 3)
	defer cluster.Destroy()
	re.NoError(err)

	err = cluster.RunInitialServers()
	re.NoError(err)
	cluster.WaitLeader()
	// Wait for all nodes becoming healthy.
	time.Sleep(time.Second * 5)

	tm3 := cluster.GetServer("tm3")
	err = tm3.Stop()
	re.NoError(err)

	client := cluster.GetServer("tm1").GetEtcdClient()
	_, err = client.MemberRemove(context.TODO(), tm3.GetServerID())
	re.NoError(err)

	// The server should not successfully start.
	res := cluster.RunServer(tm3)
	re.Error(<-res)

	members, err := etcdutil.ListEtcdMembers(client)
	re.NoError(err)
	re.Len(members.Members, 2)
}

// A deleted TM joins the previous cluster.
func TestDeletedPDJoinsPreviousCluster(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.EtcdStartTimeout = 10 * time.Second
	cluster, err := tests.NewTestCluster(ctx, 3)
	defer cluster.Destroy()
	re.NoError(err)

	err = cluster.RunInitialServers()
	re.NoError(err)
	cluster.WaitLeader()
	// Wait for all nodes becoming healthy.
	time.Sleep(time.Second * 5)

	tm3 := cluster.GetServer("tm3")
	client := cluster.GetServer("tm1").GetEtcdClient()
	_, err = client.MemberRemove(context.TODO(), tm3.GetServerID())
	re.NoError(err)

	err = tm3.Stop()
	re.NoError(err)

	// The server should not successfully start.
	res := cluster.RunServer(tm3)
	re.Error(<-res)

	members, err := etcdutil.ListEtcdMembers(client)
	re.NoError(err)
	re.Len(members.Members, 2)
}

func TestFailedPDJoinsPreviousCluster(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cluster, err := tests.NewTestCluster(ctx, 1)
	defer cluster.Destroy()
	re.NoError(err)

	re.NoError(cluster.RunInitialServers())
	cluster.WaitLeader()

	// Join the second TM.
	tm2, err := cluster.Join(ctx)
	re.NoError(err)
	re.NoError(tm2.Run())
	re.NoError(tm2.Stop())
	re.NoError(tm2.Destroy())
	re.Error(join.PrepareJoinCluster(tm2.GetConfig()))
}
