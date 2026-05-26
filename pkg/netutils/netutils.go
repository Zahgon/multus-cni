// Copyright (c) 2019 Intel Corporation
// Copyright (c) 2021 Multus Authors
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

package netutils

import (
	"net"

	"github.com/containernetworking/cni/libcni"
)

// DeleteDefaultGW removes the default gateway from marked interfaces.
func DeleteDefaultGW(netnsPath string, ifName string) error { _ = "STUB: not implemented"; return nil }

// SetDefaultGW adds a default gateway on a specific interface
func SetDefaultGW(netnsPath string, ifName string, gateways []net.IP) error {
	_ = "STUB: not implemented"
	// This ensures we're acting within the net namespace for the pod.
	return nil
}

// Do this within the net namespace.

// Pick up the link info as we need the index.

// Cycle through all the desired gateways.

// Create a new route (note: dst is nil by default)

// Perform the creation of the default route....

// DeleteDefaultGWCache updates libcni cache to remove default gateway routes in result
func DeleteDefaultGWCache(cacheDir string, rt *libcni.RuntimeConf, netName string, _ string, ipv4, ipv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteDefaultGWCacheBytes(cacheFile []byte, ipv4, ipv6 bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try to get result

func deleteDefaultGWResultRoutes(routes []interface{}, dstGW string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteDefaultGWResult(result map[string]interface{}, ipv4, ipv6 bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	// try to get cniVersion from result
	return nil, nil
}

// fallback to processing result for old cni version(0.1.0/0.2.0)

// fallback to processing result for old cni version(0.1.0/0.2.0)

// No route in result, hence we do nothing

// delete IPv4 default routes

func deleteDefaultGWResult020(result map[string]interface{}, ipv4, ipv6 bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddDefaultGWCache updates libcni cache to add default gateway result
func AddDefaultGWCache(cacheDir string, rt *libcni.RuntimeConf, netName string, _ string, gw []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func addDefaultGWCacheBytes(cacheFile []byte, gw []net.IP) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try to get result

func addDefaultGWResult(result map[string]interface{}, gw []net.IP) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	// try to get cniVersion from result
	return nil, nil
}

// fallback to processing result for old cni version(0.1.0/0.2.0)

// fallback to processing result for old cni version(0.1.0/0.2.0)

func isSupportedGatewayResultVersion(cniVersion string) bool {
	_ = "STUB: not implemented"
	return false
}

func addDefaultGWResult020(result map[string]interface{}, gw []net.IP) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
