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

package watch_test

import (
	"context"
	"testing"
	"time"

	"github.com/gottingen/tm/pkg/utils/testutil"
	"github.com/gottingen/tm/server/config"
	"github.com/gottingen/tm/tests"
	"github.com/pingcap/failpoint"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, testutil.LeakOptions...)
}

func TestWatcher(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cluster, err := tests.NewTestCluster(ctx, 1, func(conf *config.Config, serverName string) { conf.AutoCompactionRetention = "1s" })
	defer cluster.Destroy()
	re.NoError(err)

	err = cluster.RunInitialServers()
	re.NoError(err)
	cluster.WaitLeader()
	tm1 := cluster.GetServer(cluster.GetLeader())
	re.NotNil(tm1)

	tm2, err := cluster.Join(ctx)
	re.NoError(err)
	err = tm2.Run()
	re.NoError(err)
	cluster.WaitLeader()

	time.Sleep(5 * time.Second)
	tm3, err := cluster.Join(ctx)
	re.NoError(err)
	re.NoError(failpoint.Enable("github.com/gottingen/tm/server/delayWatcher", `pause`))
	err = tm3.Run()
	re.NoError(err)
	time.Sleep(200 * time.Millisecond)
	re.Equal(tm1.GetConfig().Name, tm3.GetLeader().GetName())
	err = tm1.Stop()
	re.NoError(err)
	cluster.WaitLeader()
	re.Equal(tm2.GetConfig().Name, tm2.GetLeader().GetName())
	re.NoError(failpoint.Disable("github.com/gottingen/tm/server/delayWatcher"))
	testutil.Eventually(re, func() bool {
		return tm3.GetLeader().GetName() == tm2.GetConfig().Name
	})
}

func TestWatcherCompacted(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cluster, err := tests.NewTestCluster(ctx, 1, func(conf *config.Config, serverName string) { conf.AutoCompactionRetention = "1s" })
	defer cluster.Destroy()
	re.NoError(err)

	err = cluster.RunInitialServers()
	re.NoError(err)
	cluster.WaitLeader()
	tm1 := cluster.GetServer(cluster.GetLeader())
	re.NotNil(tm1)
	client := tm1.GetEtcdClient()
	_, err = client.Put(context.Background(), "test", "v")
	re.NoError(err)
	// wait compaction
	time.Sleep(2 * time.Second)
	tm2, err := cluster.Join(ctx)
	re.NoError(err)
	err = tm2.Run()
	re.NoError(err)
	testutil.Eventually(re, func() bool {
		return tm2.GetLeader().GetName() == tm1.GetConfig().Name
	})
}
