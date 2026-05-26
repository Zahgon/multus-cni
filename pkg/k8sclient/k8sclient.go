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

package k8sclient

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"

	"github.com/containernetworking/cni/libcni"
	"github.com/containernetworking/cni/pkg/skel"
	nettypes "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	netclient "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned"
	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/types"
)

const (
	resourceNameAnnot      = "k8s.v1.cni.cncf.io/resourceName"
	defaultNetAnnot        = "v1.multus-cni.io/default-network"
	networkAttachmentAnnot = "k8s.v1.cni.cncf.io/networks"
)

// NoK8sNetworkError indicates error, no network in kubernetes
type NoK8sNetworkError struct {
	message string
}

// ClientInfo contains information given from k8s client
type ClientInfo struct {
	Client           kubernetes.Interface
	NetClient        netclient.Interface
	EventBroadcaster record.EventBroadcaster
	EventRecorder    record.EventRecorder

	// multus-thick uses these informer
	PodInformer    cache.SharedIndexInformer
	NetDefInformer cache.SharedIndexInformer
}

// AddPod adds pod into kubernetes
func (c *ClientInfo) AddPod(pod *v1.Pod) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPod gets pod from kubernetes
func (c *ClientInfo) GetPod(namespace, name string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodContext gets pod from kubernetes with context
func (c *ClientInfo) GetPodContext(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodAPILiveQuery does a live API query for the pod, instead of using informers, for cases when a failure occurred, as to prevent a cache miss.
func (c *ClientInfo) GetPodAPILiveQuery(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePod deletes a pod from kubernetes
func (c *ClientInfo) DeletePod(namespace, name string) error { _ = "STUB: not implemented"; return nil }

// AddNetAttachDef adds net-attach-def into kubernetes
func (c *ClientInfo) AddNetAttachDef(netattach *nettypes.NetworkAttachmentDefinition) (*nettypes.NetworkAttachmentDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNetAttachDef get net-attach-def from kubernetes
func (c *ClientInfo) GetNetAttachDef(namespace, name string) (*nettypes.NetworkAttachmentDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Eventf puts event into kubernetes events
func (c *ClientInfo) Eventf(object runtime.Object, eventtype, reason, messageFmt string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (e *NoK8sNetworkError) Error() string {
	_ = "STUB: not implemented"

	// SetNetworkStatus sets network status into Pod annotation
	return ""
}

func SetNetworkStatus(client *ClientInfo, k8sArgs *types.K8sArgs, netStatus []nettypes.NetworkStatus, conf *types.NetConf) error {
	_ = "STUB: not implemented"
	return nil
}

// SetPodNetworkStatusAnnotation sets network status into Pod annotation
func SetPodNetworkStatusAnnotation(client *ClientInfo, podName string, podNamespace string, podUID string, netStatus []nettypes.NetworkStatus, conf *types.NetConf) error {
	_ = "STUB: not implemented"
	return nil
}

// No available kube client and no delegates, we can't do anything

func parsePodNetworkObjectName(podnetwork string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// Check and see if each item matches the specification for valid attachment name.
// "Valid attachment names must be comprised of units of the DNS-1123 label format"
// [a-z0-9]([-a-z0-9]*[a-z0-9])?
// It must start and end alphanumerically.

func parsePodNetworkAnnotation(podNetworks, defaultNamespace string) ([]*types.NetworkSelectionElement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Comma-delimited list of network attachment object names

// Remove leading and trailing whitespace.

// Parse network name (i.e. <namespace>/<network name>@<ifname>)

// validate MAC address

// validate GUID address

// validate IP address

// compatibility pre v3.2, will be removed in v4.0

func getKubernetesDelegate(client *ClientInfo, net *types.NetworkSelectionElement, confdir string, pod *v1.Pod, resourceMap map[string]*types.ResourceInfo) (*types.DelegateNetConf, map[string]*types.ResourceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get resourceName annotation from NetworkAttachmentDefinition

// ResourceName annotation is found; try to get device info from resourceMap

// increment Index for next delegate

// GetK8sArgs gets k8s related args from CNI args
func GetK8sArgs(args *skel.CmdArgs) (*types.K8sArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TryLoadPodDelegates attempts to load Kubernetes-defined delegates and add them to the Multus config.
// Returns the number of Kubernetes-defined delegates added or an error.
func TryLoadPodDelegates(pod *v1.Pod, conf *types.NetConf, clientInfo *ClientInfo, resourceMap map[string]*types.ResourceInfo) (int, *ClientInfo, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// No available kube client and no delegates, we can't do anything

// Check gatewayRequest is configured in delegates
// and mark its config if gateway filter is required

// GetPodNetwork gets net-attach-def annotation from pod
func GetPodNetwork(pod *v1.Pod) ([]*types.NetworkSelectionElement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNetworkDelegates returns delegatenetconf from net-attach-def annotation in pod
func GetNetworkDelegates(k8sclient *ClientInfo, pod *v1.Pod, networks []*types.NetworkSelectionElement, conf *types.NetConf, resourceMap map[string]*types.ResourceInfo) ([]*types.DelegateNetConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read all network objects referenced by 'networks'

// The pods namespace (stored as defaultNamespace, does not equal the annotation's target namespace in net.Namespace)
// In the case that this is a mismatch when namespaceisolation is enabled, this should be an error.

// We allow exceptions based on the specified list of non-isolated namespaces (and/or "default" namespace, by default)

func isValidNamespaceReference(targetns string, allowednamespaces []string) bool {
	_ = "STUB: not implemented"
	return false
}

// getNetDelegate loads delegate network for clusterNetwork/defaultNetworks
func getNetDelegate(client *ClientInfo, pod *v1.Pod, netname, confdir, namespace string, resourceMap map[string]*types.ResourceInfo) (*types.DelegateNetConf, map[string]*types.ResourceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// if netname is not directory or file, it must be net-attach-def name or CNI config name

// option1) search CRD object for the network

// option2) search CNI json config file, which has <netname> as CNI name, from confDir

// option3) search directory

// option4) if file path (absolute), then load it directly

// Or it's not a conflist...
// after libcni v1.2.3 there's no support support this old-school method with non-conflists.
// this method doesn't check if there's a 0 length plugins field, that is.

func loadSubdirectoryChain(bytes []byte, cniconfdir string) (*libcni.NetworkConfigList, error) {
	_ = "STUB: not implemented"
	// Load the network configuration from the byte array
	return nil, nil
}

// Check if plugins need to be loaded from files

// Let's validate that conf.Name
// From the CNI spec:
// > Must start with an alphanumeric character, optionally followed by any combination of one or more alphanumeric characters,
// > underscore, dot (.) or hyphen (-). Must not contain characters disallowed in file paths.

// LoadChainedDelegatesFromBytes loads a CNI configuration byte array and returns a DelegateNetConf with the chain added.
func LoadChainedDelegatesFromBytes(bytes []byte, cniconfdir string) *types.DelegateNetConf {
	_ = "STUB: not implemented"
	return nil
}

// Create and return a DelegateNetConf from the configuration list

// LoadChainedPluginsFromFile loads a CNI configuration file and returns the NetworkConfigList
func LoadChainedPluginsFromFile(filename string) (*libcni.NetworkConfigList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// stat the file to make sure it's a normal file.

// GetDefaultNetworks parses 'defaultNetwork' config, gets network json and put it into netconf.Delegates.
func GetDefaultNetworks(pod *v1.Pod, conf *types.NetConf, kubeClient *ClientInfo, resourceMap map[string]*types.ResourceInfo) (map[string]*types.ResourceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No available kube client and no delegates, we can't do anything

// Pod in kube-system namespace does not have default network for now.

// tryLoadK8sPodDefaultNetwork get pod default network from annotations
func tryLoadK8sPodDefaultNetwork(kubeClient *ClientInfo, pod *v1.Pod, conf *types.NetConf) (*types.DelegateNetConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The CRD object of default network should only be defined in multusNamespace

// ConfigSourceAnnotationKey specifies kubernetes annotation, defined in k8s.io/kubernetes/pkg/kubelet/types
const ConfigSourceAnnotationKey = "kubernetes.io/config.source"

// IsStaticPod returns true if the pod is static pod.
func IsStaticPod(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }
