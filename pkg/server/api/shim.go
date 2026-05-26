// Copyright (c) 2022 Multus Authors
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

package api

import (
	"github.com/containernetworking/cni/pkg/skel"
)

// ShimNetConf for the SHIM cni config file written in json
type ShimNetConf struct {
	// Note: This struct contains NetConf in pkg/types, but this struct is only used to parse
	// following fields, so we skip to include NetConf here. Other fields are directly send to
	// multus-daemon as a part of skel.CmdArgs, StdinData.
	// types.NetConf

	CNIVersion      string `json:"cniVersion,omitempty"`
	MultusSocketDir string `json:"daemonSocketDir"`
	LogFile         string `json:"logFile,omitempty"`
	LogLevel        string `json:"logLevel,omitempty"`
	LogToStderr     bool   `json:"logToStderr,omitempty"`
}

// readyCheckFunc defines a type for API readiness check functions
type readyCheckFunc func(string) error

// CmdAdd implements the CNI spec ADD command handler
func CmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// CmdCheck implements the CNI spec CHECK command handler
func CmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// CmdDel implements the CNI spec DEL command handler
func CmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// No error in DEL (as of CNI spec)

// CmdGC implements the CNI spec GC command handler
func CmdGC(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// CmdStatus implements the CNI spec STATUS command handler
func CmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func postRequest(args *skel.CmdArgs, readinessCheck readyCheckFunc) (*Response, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Execute the readiness check as necessary (e.g. don't wait on CNI DEL)

// Create and fill a Request with this Plugin's environment and stdin which
// contain the CNI variables and configuration
func newCNIRequest(args *skel.CmdArgs) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func shimConfig(cniConfig []byte) (*ShimNetConf, error) { _ = "STUB: not implemented"; return nil, nil }

// Logging
