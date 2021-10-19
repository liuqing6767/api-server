// Copyright (c) 2021 The BFE Authors.
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

package bfe_cluster

import (
	"net/http"

	"github.com/bfenetworks/api-server/lib/xreq"
	"github.com/bfenetworks/api-server/model/iauth"
	"github.com/bfenetworks/api-server/stateful/container"
)

// BFEClusterDetail Request Param
// AUTO GEN BY ctrl, MODIFY AS U NEED
type BFEClusterDetail struct {
	Name string `json:"name" uri:"name"`
	Pool string `json:"pool" uri:"pool"`
}

// ListRoute route
// AUTO GEN BY ctrl, MODIFY AS U NEED
var ListEndpoint = &xreq.Endpoint{
	Path:       "/bfe-clusters",
	Method:     http.MethodGet,
	Handler:    xreq.Convert(ListAction),
	Authorizer: iauth.FA(iauth.FeatureBFECluster, iauth.ActionReadAll),
}

func listActionProcess(req *http.Request) ([]*BFEClusterDetail, error) {
	list, err := container.BFEClusterManager.FetchBFEClusters(req.Context(), nil)
	if err != nil {
		return nil, err
	}

	rst := []*BFEClusterDetail{}
	for _, one := range list {
		rst = append(rst, &BFEClusterDetail{
			Name: one.Name,
			Pool: one.Pool,
		})
	}
	return rst, nil
}

var _ xreq.Handler = ListAction

// ListAction action
// AUTO GEN BY ctrl, MODIFY AS U NEED
func ListAction(req *http.Request) (interface{}, error) {
	return listActionProcess(req)
}
