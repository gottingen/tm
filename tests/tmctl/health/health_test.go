// Copyright 2019 TiKV Project Authors.
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

package health_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gottingen/tm/server/api"
	"github.com/gottingen/tm/server/cluster"
	"github.com/gottingen/tm/tests"
	"github.com/gottingen/tm/tests/tmctl"
	pdctlCmd "github.com/gottingen/tm/tools/tm-ctl/tmctl"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	re := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tc, err := tests.NewTestCluster(ctx, 3)
	re.NoError(err)
	err = tc.RunInitialServers()
	re.NoError(err)
	tc.WaitLeader()
	leaderServer := tc.GetServer(tc.GetLeader())
	re.NoError(leaderServer.BootstrapCluster())
	pdAddr := tc.GetConfig().GetClientURL()
	cmd := pdctlCmd.GetRootCmd()
	defer tc.Destroy()

	client := tc.GetEtcdClient()
	members, err := cluster.GetMembers(client)
	re.NoError(err)
	healthMembers := cluster.CheckHealth(tc.GetHTTPClient(), members)
	healths := []api.Health{}
	for _, member := range members {
		h := api.Health{
			Name:       member.Name,
			MemberID:   member.MemberId,
			ClientUrls: member.ClientUrls,
			Health:     false,
		}
		if _, ok := healthMembers[member.GetMemberId()]; ok {
			h.Health = true
		}
		healths = append(healths, h)
	}

	// health command
	args := []string{"-u", pdAddr, "health"}
	output, err := tmctl.ExecuteCommand(cmd, args...)
	re.NoError(err)
	h := make([]api.Health, len(healths))
	re.NoError(json.Unmarshal(output, &h))
	re.Equal(healths, h)
}
