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

package config

import (
	"context"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// MultusDefaultNetworkName holds the default name of the multus network
const (
	multusConfigFileName     = "00-multus.conf"
	MultusDefaultNetworkName = "multus-cni-network"
	UserRWPermission         = 0600
)

// Manager monitors the configuration of the primary CNI plugin, and
// regenerates multus configuration whenever it gets updated.
type Manager struct {
	cniConfigData              map[string]interface{}
	configWatcher              *fsnotify.Watcher
	multusConfig               *MultusConf
	multusConfigDir            string
	multusConfigFilePath       string
	readinessIndicatorFilePath string
	primaryCNIConfigPath       string
}

// NewManager returns a config manager object, configured to read the
// primary CNI configuration in `config.MultusAutoconfigDir`. If
// `config.MultusMasterCni` is empty, this constructor will auto-discover the
// primary CNI for which it will delegate.
func NewManager(config MultusConf) (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

// overrideCNIVersion overrides cniVersion in cniConfigFile, it should be used only in kind case
func overrideCNIVersion(cniConfigFile string, multusCNIVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func newManager(config MultusConf, defaultCNIPluginName string) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start generates an updated Multus config, writes it, and begins watching
// the config directory and readiness indicator files for changes
func (m *Manager) Start(ctx context.Context, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) loadPrimaryCNIConfigFromFile() error { _ = "STUB: not implemented"; return nil }

// overrideNetworkName overrides the name of the multus configuration with the
// name of the delegated primary CNI.
func (m *Manager) overrideNetworkName() error { _ = "STUB: not implemented"; return nil }

func (m *Manager) loadPrimaryCNIConfigurationData(primaryCNIConfigData interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateConfig generates a multus configuration from its current state
func (m *Manager) GenerateConfig() (string, error) { _ = "STUB: not implemented"; return "", nil }

// monitorPluginConfiguration monitors the configuration file pointed
// to by the primaryCNIPluginName attribute, and re-generates the multus
// configuration whenever the primary CNI config is updated.
func (m *Manager) monitorPluginConfiguration(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// if readinessIndicatorFile is removed, then restart multus

// PersistMultusConfig persists the provided configuration to the disc, with
// Read / Write permissions. The output file path is `<multus auto config dir>/00-multus.conf`
func (m *Manager) PersistMultusConfig(config string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *Manager) shouldRegenerateConfig(event fsnotify.Event) bool {
	_ = "STUB: not implemented"
	// first, check the readiness indicator file existence
	return false
}

// we're watching the DIR where the config sits, and the event
// does not concern the primary CNI config. Skip it.

func getPrimaryCNIPluginName(multusAutoconfigDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newWatcher(cniConfigDir string, readinessIndicatorDir string) (*fsnotify.Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close watcher on error

// if readinessIndicatorDir is different from cniConfigDir,

func primaryCNIData(masterCNIPluginPath string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
