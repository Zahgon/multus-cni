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

package server

import (
	"context"
	"net"
	"net/http"

	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/skel"
	cnitypes "github.com/containernetworking/cni/pkg/types"

	k8s "gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/k8sclient"
	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/server/api"
	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/types"

	netdefclient "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned"
	netdefinformer "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/informers/externalversions"

	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	fullReadWriteExecutePermissions    = 0777
	thickPluginSocketRunDirPermissions = 0700
)

// FilesystemPreRequirements ensures the target `rundir` features the correct
// permissions.
func FilesystemPreRequirements(rundir string) error { _ = "STUB: not implemented"; return nil }

func printCmdArgs(args *skel.CmdArgs) string { _ = "STUB: not implemented"; return "" }

// HandleCNIRequest is the CNI server handler function; it is invoked whenever
// a CNI request is processed.
// Note: k8sArgs may be nil for plugin-level commands (STATUS, GC) that have no pod context.
func (s *Server) HandleCNIRequest(cmd string, k8sArgs *types.K8sArgs, cniCmdArgs *skel.CmdArgs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HandleDelegateRequest is the CNI server handler function; it is invoked whenever
// a CNI request is processed as delegate CNI request.
func (s *Server) HandleDelegateRequest(cmd string, k8sArgs *types.K8sArgs, cniCmdArgs *skel.CmdArgs, interfaceAttributes *api.DelegateInterfaceAttributes) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetListener creates a listener to a unix socket located in `socketPath`
func GetListener(socketPath string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// Informer transform to trim object fields for memory efficiency.
func informerObjectTrim(obj interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newNetDefInformer(netdefClient netdefclient.Interface) (netdefinformer.SharedInformerFactory, cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return *new(netdefinformer.SharedInformerFactory), *new(cache.SharedIndexInformer)
}

func newPodInformer(kubeClient kubernetes.Interface, nodeName string) (internalinterfaces.SharedInformerFactory, cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return *new(internalinterfaces.SharedInformerFactory), *new(cache.SharedIndexInformer)
}

// Only watch for local pods

func isPerNodeCertEnabled(config *PerNodeCertificate) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NewCNIServer creates and returns a new Server object which will listen on a socket in the given path
func NewCNIServer(daemonConfig *ControllerNetConf, serverConfig []byte, ignoreReadinessIndicator bool) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCNIServer(rundir string, kubeClient *k8s.ClientInfo, exec invoke.Exec, servConfig []byte, ignoreReadinessIndicator bool) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// register metrics

// handle for '/cni'

// Empty response JSON means success with no body

// handle for '/delegate'

// Empty response JSON means success with no body

// handle for '/healthz'

// this handle for the rest of above

// Start starts the server and begins serving on the given listener
func (s *Server) Start(ctx context.Context, l net.Listener) { _ = "STUB: not implemented"; return }

// Give the initial sync some time to complete in large clusters, but
// don't wait forever

// Give the initial sync some time to complete in large clusters, but
// don't wait forever

func (s *Server) writeCNIErrorResponse(w http.ResponseWriter, err error) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) wrapCNIRequestError(cmdArgs *skel.CmdArgs, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Prefix error with request information for easier debugging.

func (s *Server) handleCNIRequest(r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// STATUS and GC are plugin-level commands with no pod context,
// so they don't have K8S_POD_NAME/K8S_POD_NAMESPACE in CNI_ARGS.

func (s *Server) handleDelegateRequest(r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func overrideCNIConfigWithServerConfig(cniConf []byte, overrideConf []byte, ignoreReadinessIndicator bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy each key of the override config into the CNI config except for
// a few specific keys

func (s *Server) extractCniData(cniRequest *api.Request, overrideConf []byte) (string, *skel.CmdArgs, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// STATUS and GC are plugin-level commands with no pod context;
// they don't require CNI_CONTAINERID, CNI_NETNS, or CNI_ARGS.

func kubernetesRuntimeArgs(cniRequestEnvVariables map[string]string, kubeClient *k8s.ClientInfo) (*types.K8sArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gatherCNIArgs(env map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func podUID(kubeclient *k8s.ClientInfo, cniArgs map[string]string, podNamespace, podName string) (string, error) {
	_ = "STUB: not implemented"
	// UID may not be passed by all runtimes yet. Will be passed
	// by CRIO 1.20+ and containerd 1.5+ soon.
	// CRIO 1.20: https://github.com/cri-o/cri-o/pull/5029
	// CRIO 1.21: https://github.com/cri-o/cri-o/pull/5028
	// CRIO 1.22: https://github.com/cri-o/cri-o/pull/5026
	// containerd 1.6: https://github.com/containerd/containerd/pull/5640
	// containerd 1.5: https://github.com/containerd/containerd/pull/5643
	return "", nil
}

func (s *Server) cmdAdd(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) cmdDel(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) cmdCheck(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) cmdGC(cmdArgs *skel.CmdArgs, _ *types.K8sArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) cmdStatus(cmdArgs *skel.CmdArgs, _ *types.K8sArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func serializeResult(result cnitypes.Result) ([]byte, error) {
	_ = "STUB: not implemented"
	// cni result is converted to latest here and decoded to specific cni version at multus-shim
	return nil, nil
}

func (s *Server) cmdDelegateAdd(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs, multusConfig *types.NetConf, interfaceAttributes *api.DelegateInterfaceAttributes) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy deleate annotation into network selection element

func (s *Server) cmdDelegateCheck(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs, multusConfig *types.NetConf) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) cmdDelegateStatus(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs, multusConfig *types.NetConf) error {
	_ = "STUB: not implemented"
	return nil
}

// note: this function may send back error to the client. In cni spec, command DEL should NOT send any error
// because container deletion follows cni DEL command. But in delegateDel case, container is not removed by
// this delegateDel, hence we decide to send error message to the request sender.
func (s *Server) cmdDelegateDel(cmdArgs *skel.CmdArgs, k8sArgs *types.K8sArgs, multusConfig *types.NetConf) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadDaemonNetConf loads the configuration for the multus daemon
func LoadDaemonNetConf(config []byte) (*ControllerNetConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
