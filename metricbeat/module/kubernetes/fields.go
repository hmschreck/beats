// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
