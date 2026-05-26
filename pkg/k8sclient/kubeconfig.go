// Copyright (c) 2023 Multus Authors
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

package k8sclient

import (
	"time"

	certificatesv1 "k8s.io/api/certificates/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

const (
	certNamePrefix       = "multus-client"
	certCommonNamePrefix = "system:multus"
	certOrganization     = "system:multus"
)

var (
	certUsages = []certificatesv1.KeyUsage{certificatesv1.UsageDigitalSignature, certificatesv1.UsageClientAuth}
)

// getPerNodeKubeconfig creates new kubeConfig, based on bootstrap, with new certDir
func getPerNodeKubeconfig(bootstrap *rest.Config, certDir string) *rest.Config {
	_ = "STUB: not implemented"
	return nil
}

// Switch to the per-node client certificate instead of reusing bootstrap auth.

// Allow multus (especially in server mode) to make more concurrent requests
// to reduce client-side throttling

// Set the config timeout to one minute.

// PerNodeK8sClient creates/reload new multus kubeconfig per-node.
func PerNodeK8sClient(nodeName, bootstrapKubeconfigFile string, certDuration time.Duration, certDir string) (*ClientInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we have a valid certificate, use that to fetch CSRs.
// Otherwise, use the bootstrap credentials from bootstrapKubeconfig
// https://github.com/kubernetes/kubernetes/blob/068ee321bc7bfe1c2cefb87fb4d9e5deea84fbc8/cmd/kubelet/app/server.go#L953-L963

// When a current certificate exists, prefer it for CSR operations but fall back
// to bootstrap credentials if the stored per-node config is no longer trusted.

// the default value for CertCallbackRefreshDuration (5min) is too long for short-lived certs,
// set it to a more sensible value

// InClusterK8sClient returns the `k8s.ClientInfo` struct to use to connect to
// the k8s API.
func InClusterK8sClient() (*ClientInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// SetK8sClientInformers adds informer structure to ClientInfo to utilize in thick daemon
func (c *ClientInfo) SetK8sClientInformers(podInformer, netDefInformer cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return
}

// GetK8sClient gets client info from kubeconfig
func GetK8sClient(kubeconfig string, kubeClient *ClientInfo) (*ClientInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we get a valid kubeClient (eg from testcases) just return that
// one.

// Otherwise try to create a kubeClient from a given kubeConfig

// uses the current context in kubeconfig

// Try in-cluster config where multus might be running in a kubernetes pod

// No kubernetes config; assume we shouldn't talk to Kube at all

// Specify that we use gRPC

// Set the config timeout to one minute.

// Allow multus (especially in server mode) to make more concurrent requests
// to reduce client-side throttling

// newClientInfo returns a `ClientInfo` from a configuration created from an
// existing kubeconfig file.
func newClientInfo(config *rest.Config) (*ClientInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
