// Copyright (c) 2018 Intel Corporation
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

package testing

import (
	"io"
	"net"

	netv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	v1 "k8s.io/api/core/v1"

	"github.com/containernetworking/cni/pkg/types"
	types020 "github.com/containernetworking/cni/pkg/types/020"
)

// NewFakeNetAttachDef returns net-attach-def for testing
func NewFakeNetAttachDef(namespace, name, config string) *netv1.NetworkAttachmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

// NewFakeNetAttachDefFile returns net-attach-def for testing with conf file
func NewFakeNetAttachDefFile(namespace, name, filePath, fileData string) *netv1.NetworkAttachmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

// NewFakeNetAttachDefAnnotation returns net-attach-def with resource annotation
func NewFakeNetAttachDefAnnotation(namespace, name, config string) *netv1.NetworkAttachmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

// NewFakePod creates fake Pod object
func NewFakePod(name string, netAnnotation string, defaultNetAnnotation string) *v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// EnsureCIDR parses/verify CIDR ip string and convert to net.IPNet
func EnsureCIDR(cidr string) *net.IPNet { _ = "STUB: not implemented"; return nil }

// Result is stub Result for testing
type Result struct {
	CNIVersion string             `json:"cniVersion,omitempty"`
	IP4        *types020.IPConfig `json:"ip4,omitempty"`
	IP6        *types020.IPConfig `json:"ip6,omitempty"`
	DNS        types.DNS          `json:"dns,omitempty"`
}

// Version returns current CNIVersion of the given Result
func (r *Result) Version() string { _ = "STUB: not implemented"; return "" }

// GetAsVersion returns a Result object given a version
func (r *Result) GetAsVersion(version string) (types.Result, error) {
	_ = "STUB: not implemented"
	return *new(types.Result), nil
}

// Print prints a Result's information to std out
func (r *Result) Print() error { _ = "STUB: not implemented"; return nil }

// PrintTo prints a Result's information to the provided writer
func (r *Result) PrintTo(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// String returns a formatted string in the form of "[IP4: $1,][ IP6: $2,] DNS: $3" where
// $1 represents the receiver's IPv4, $2 represents the receiver's IPv6 and $3 the
// receiver's DNS. If $1 or $2 are nil, they won't be present in the returned string.
func (r *Result) String() string { _ = "STUB: not implemented"; return "" }

// Int returns a pointer to an int
func Int(i int) *int {
	_ = "STUB: not implemented"

	// Bool returns a pointer to a bool.
	return nil
}

func Bool(b bool) *bool { _ = "STUB: not implemented"; return nil }
