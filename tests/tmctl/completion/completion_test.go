// Copyright 2020 TiKV Project Authors.
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

package completion_test

import (
	"testing"

	"github.com/gottingen/tm/tests/tmctl"
	pdctlCmd "github.com/gottingen/tm/tools/tm-ctl/tmctl"
	"github.com/stretchr/testify/require"
)

func TestCompletion(t *testing.T) {
	re := require.New(t)
	cmd := pdctlCmd.GetRootCmd()

	// completion command
	args := []string{"completion", "bash"}
	_, err := tmctl.ExecuteCommand(cmd, args...)
	re.NoError(err)

	// completion command
	args = []string{"completion", "zsh"}
	_, err = tmctl.ExecuteCommand(cmd, args...)
	re.NoError(err)
}
