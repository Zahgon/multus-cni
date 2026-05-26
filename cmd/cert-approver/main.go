// Copyright (c) 2023 Network Plumbing Working Group
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

// This is Kubernetes controller which approves CSR submitted by multus.
// This command is required only if multus runs with per-node certificate.
package main

// Note: cert-approver should be simple, just approve multus' CSR, hence
// this go code should not have any dependencies from pkg/, if possible,
// to keep its code simplicity.
import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	certificatesv1 "k8s.io/api/certificates/v1"
	"k8s.io/klog/v2"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

// CertController object
type CertController struct {
	clientset          kubernetes.Interface
	queue              workqueue.RateLimitingInterface
	informer           cache.SharedIndexInformer
	broadcaster        record.EventBroadcaster
	recorder           record.EventRecorder
	commonNamePrefixes string
}

const (
	maxDuration                = time.Hour * 24 * 365
	resyncPeriod time.Duration = time.Second * 3600 // resync every one hour, default is 10 hour
	maxRetries                 = 5
)

var (
	// ControllerName provides controller name
	ControllerName = "csr-approver"
	// NamePrefix specifies which name in certification request should be target to approve
	NamePrefix = "system:multus"
	// Organization specifies which org in certification request should be target to approve
	Organization = []string{"system:multus"}
	// Groups specifies which group in certification request should be target to approve
	Groups = sets.New[string]("system:nodes", "system:multus", "system:authenticated")
	// UserPrefixes specifies which name prefix in certification request should be target to approve
	UserPrefixes = sets.New[string]("system:node", NamePrefix)
	// Usages specifies which usage in certification request should be target to approve
	Usages = sets.New[certificatesv1.KeyUsage](
		certificatesv1.UsageDigitalSignature,
		certificatesv1.UsageClientAuth)
)

// NewCertController creates certcontroller
func NewCertController() (*CertController, error) { _ = "STUB: not implemented"; return nil, nil }

// setup Kubernetes API client

// Run starts controller
func (c *CertController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// HasSynced is required for the cache.Controller interface.
func (c *CertController) HasSynced() bool { _ = "STUB: not implemented"; return false }

// LastSyncResourceVersion is required for the cache.Controller interface.
func (c *CertController) LastSyncResourceVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CertController) runWorker() { _ = "STUB: not implemented"; return }

// continue looping

func (c *CertController) processNextItem() bool {
	_ = "STUB: not implemented"
	// Wait until there is a new item in the working queue
	return false
}

// Tell the queue that we are done with processing this key. This unblocks the key for other workers
// This allows safe parallel processing because two pods with the same key are never processed in
// parallel.

// Invoke the method containing the business logic

// Handle the error if something went wrong during the execution of the business logic

// handleErr checks if an error happened and makes sure we will retry later.
func (c *CertController) handleErr(err error, key interface{}) {
	_ = "STUB: not implemented"

	// Forget about the #AddRateLimited history of the key on every successful synchronization.
	// This ensures that future processing of updates for this key is not delayed because of
	// an outdated error history.
	return
}

// This controller retries 5 times if something goes wrong. After that, it stops trying.

// Re-enqueue the key rate limited. Based on the rate limiter on the
// queue and the re-enqueue history, the key will be processed later again.

// Report to an external entity that, even after several retries, we could not successfully process this key

func (c *CertController) processItem(key string) error { _ = "STUB: not implemented"; return nil }

// CSR specific functions

func (c *CertController) filterCSR(csr *certificatesv1.CertificateSigningRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CertController) approveCSR(ctx context.Context, csr *certificatesv1.CertificateSigningRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CertController) denyCSR(ctx context.Context, csr *certificatesv1.CertificateSigningRequest, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func isApprovedOrDenied(status *certificatesv1.CertificateSigningRequestStatus) bool {
	_ = "STUB: not implemented"
	return false
}

func main() {
	klog.Infof("starting cert-approver")

	// Start watching for pod creations
	certController, err := NewCertController()
	if err != nil {
		klog.Fatal(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	go certController.Run(stopCh)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-sigterm
}
