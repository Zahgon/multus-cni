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
	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/logging"
)

const (
	configListCapabilityKey   = "plugins"
	multusPluginName          = "multus-shim"
	singleConfigCapabilityKey = "capabilities"
)

// Option mutates the `conf` object
type Option func(conf *MultusConf) error

// MultusConf holds the multus configuration
type MultusConf struct {
	BinDir                   string              `json:"binDir,omitempty"`
	Capabilities             map[string]bool     `json:"capabilities,omitempty"`
	CNIVersion               string              `json:"cniVersion"`
	LogFile                  string              `json:"logFile,omitempty"`
	LogLevel                 string              `json:"logLevel,omitempty"`
	LogToStderr              bool                `json:"logToStderr,omitempty"`
	LogOptions               *logging.LogOptions `json:"logOptions,omitempty"`
	Name                     string              `json:"name"`
	ClusterNetwork           string              `json:"clusterNetwork,omitempty"`
	NamespaceIsolation       bool                `json:"namespaceIsolation,omitempty"`
	RawNonIsolatedNamespaces string              `json:"globalNamespaces,omitempty"`
	ReadinessIndicatorFile   string              `json:"readinessindicatorfile,omitempty"`
	Type                     string              `json:"type"`
	CniDir                   string              `json:"cniDir,omitempty"`
	CniConfigDir             string              `json:"cniConfigDir,omitempty"`
	AuxiliaryCNIChainName    string              `json:"auxiliaryCNIChainName,omitempty"`
	DaemonSocketDir          string              `json:"daemonSocketDir,omitempty"`
	MultusConfigFile         string              `json:"multusConfigFile,omitempty"`
	MultusMasterCni          string              `json:"multusMasterCNI,omitempty"`
	MultusAutoconfigDir      string              `json:"multusAutoconfigDir,omitempty"`
	ForceCNIVersion          bool                `json:"forceCNIVersion,omitempty"`
	OverrideNetworkName      bool                `json:"overrideNetworkName,omitempty"`
}

// ParseMultusConfig parses multus config from configPath and create MultusConf.
func ParseMultusConfig(configPath string) (*MultusConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// change name

// CheckVersionCompatibility checks compatibilty of the
// top level cni version with the delegate cni version.
// Since version 0.4.0, CHECK was introduced, which
// causes incompatibility.
func CheckVersionCompatibility(mc *MultusConf, delegate interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate generates the multus configuration from whatever state is currently
// held
func (mc *MultusConf) Generate() (string, error) {
	_ = "STUB: not implemented"
	// before marshal, flush variables which is not required for multus-shim config
	return "", nil
}

// Readiness indicator file existence is already handled by the
// ConfigManager via an fsnotify watch, so CmdAdd/CmdDel don't need to.

func (mc *MultusConf) setCapabilities(cniData interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func extractCapabilities(capabilitiesInterface interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func findMasterPlugin(cniConfigDirPath string, remainingTries int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
