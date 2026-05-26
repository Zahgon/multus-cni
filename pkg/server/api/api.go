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
	"time"
)

const (
	// APIReadyPollDuration specifies duration for API readiness check polling
	APIReadyPollDuration = 100 * time.Millisecond
	// APIReadyPollTimeout specifies timeout for API readiness check polling
	APIReadyPollTimeout = 60000 * time.Millisecond

	// MultusCNIAPIEndpoint is an endpoint for multus CNI request (for multus-shim)
	MultusCNIAPIEndpoint = "/cni"
	// MultusDelegateAPIEndpoint is an endpoint for multus delegate request (for hotplug)
	MultusDelegateAPIEndpoint = "/delegate"
	defaultMultusRunDir       = "/run/multus/"

	// MultusHealthAPIEndpoint is an endpoint API clients can query to know if they can communicate w/ multus server
	MultusHealthAPIEndpoint = "/healthz"
)

// DoCNI sends a CNI request to the CNI server via JSON + HTTP over a root-owned unix socket,
// and returns the result
func DoCNI(url string, req interface{}, socketPath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAPIEndpoint returns endpoint URL for multus-daemon
func GetAPIEndpoint(endpoint string) string { _ = "STUB: not implemented"; return "" }

// CreateDelegateRequest creates Request for delegate API request
func CreateDelegateRequest(cniCommand, cniContainerID, cniNetNS, cniIFName, podNamespace, podName, podUID string, cniConfig []byte, interfaceAttributes *DelegateInterfaceAttributes) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WaitUntilAPIReady checks API readiness
func WaitUntilAPIReady(socketPath string) error { _ = "STUB: not implemented"; return nil }

// CheckAPIReadyNow checks API readiness once
func CheckAPIReadyNow(socketPath string) error { _ = "STUB: not implemented"; return nil }
