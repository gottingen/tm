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

package keyvisual

import (
	"github.com/pingcap/tidb-dashboard/pkg/keyvisual/region"

	"github.com/gottingen/tm/pkg/dashboard/keyvisual/input"
	"github.com/gottingen/tm/server"
)

// GenCustomDataProvider generates a custom DataProvider for the dashboard `keyvisual` package.
func GenCustomDataProvider(srv *server.Server) *region.DataProvider {
	return &region.DataProvider{
		PeriodicGetter: input.NewCorePeriodicGetter(srv),
	}
}
