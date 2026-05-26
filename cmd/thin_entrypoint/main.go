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

// This is a entrypoint for thin (stand-alone) images.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"

	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/cmdutils"
	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/signals"
)

// Options stores command line options
type Options struct {
	CNIBinDir                string
	CNIConfDir               string
	CNIVersion               string
	MultusConfFile           string
	MultusBinFile            string // may be hidden or remove?
	MultusCNIConfDir         string
	SkipMultusBinaryCopy     bool
	MultusKubeConfigFileHost string
	MultusMasterCNIFileName  string
	NamespaceIsolation       bool
	GlobalNamespaces         string
	MultusAutoconfigDir      string
	MultusLogToStderr        bool
	MultusLogLevel           string
	MultusLogFile            string
	OverrideNetworkName      bool
	CleanupConfigOnExit      bool
	RenameConfFile           bool
	ReadinessIndicatorFile   string
	AdditionalBinDir         string
	ForceCNIVersion          bool
	SkipTLSVerify            bool
	SkipMultusConfWatch      bool
}

const (
	serviceAccountTokenFile = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	serviceAccountCAFile    = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func (o *Options) addFlags() {
	_ = "STUB: not implemented"
	// suppress error message for help
	return
}

func (o *Options) verifyFileExists() error {
	_ = "STUB: not implemented"
	// CNIConfDir
	return nil
}

// CNIBinDir

// MultusBinFile

// MultusConfFile

const kubeConfigTemplate = `# Kubeconfig file for Multus CNI plugin.
apiVersion: v1
kind: Config
clusters:
- name: local
  cluster:
    server: {{ .KubeConfigHost }}
    {{ .KubeServerTLS }}
users:
- name: multus
  user:
    token: "{{ .KubeServiceAccountToken }}"
contexts:
- name: multus-context
  context:
    cluster: local
    user: multus
current-context: multus-context
`

func getFileAndHash(filepath string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (o *Options) createKubeConfig(prevCAHash, prevSATokenHash []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// don't log "recreating" on first function execution

// create multus.d directory

// create multus cni conf directory

// get Kubernetes service protocol/host/port

// check tlsConfig

// create tlsConfig by service account CA file

// create kubeconfig by template and replace it by atomic

// generate kubeconfig from template

// replace file with tempfile

const multusConflistTemplate = `{
    "cniVersion": "{{ .CNIVersion }}",
    "name": "{{ .MasterPluginNetworkName }}",
    "plugins": [ {
        "type": "multus",{{
            .NestedCapabilities
        }}{{
            .NamespaceIsolationConfig
        }}{{
            .GlobalNamespacesConfig
        }}{{
            .LogToStderrConfig
        }}{{
            .LogLevelConfig
        }}{{
            .LogFileConfig
        }}{{
            .AdditionalBinDirConfig
        }}{{
            .MultusCNIConfDirConfig
        }}{{
            .ReadinessIndicatorFileConfig
        }}
        "kubeconfig": "{{ .MultusKubeConfigFileHost }}",
        "delegates": [
            {{ .MasterPluginJSON }}
        ]
    }]
}
`

const multusConfTemplate = `{
        "cniVersion": "{{ .CNIVersion }}",
        "name": "{{ .MasterPluginNetworkName }}",
        "type": "multus",{{
            .NestedCapabilities
        }}{{
            .NamespaceIsolationConfig
        }}{{
            .GlobalNamespacesConfig
        }}{{
            .LogToStderrConfig
        }}{{
            .LogLevelConfig
        }}{{
            .LogFileConfig
        }}{{
            .AdditionalBinDirConfig
        }}{{
            .MultusCNIConfDirConfig
        }}{{
            .ReadinessIndicatorFileConfig
        }}
        "kubeconfig": "{{ .MultusKubeConfigFileHost }}",
        "delegates": [
                {{ .MasterPluginJSON }}
        ]
}
`

func (o *Options) getMasterConfigPath() (string, error) {
	_ = "STUB: not implemented"
	// Master config file is specified
	return "", nil
}

// Pick the alphabetically first config file from MultusAutoconfigDir

// No config file found

func (o *Options) createMultusConfig(prevMasterConfigFileHash []byte) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// don't log "recreating" on first function execution

// check CNIVersion

// check OverrideNetworkName (if true, get master plugin name, otherwise 'multus-cni-network'

// check capabilities (from master conf, top and 'plugins')

// check NamespaceIsolation

// check GlobalNamespaces

// check MultusLogToStderr

// check MultusLogLevel (debug/error/panic/verbose) and reject others

// no logLevel config, skipped

// check MultusLogFile

// check AdditionalBinDir

// check MultusCNIConfDir

// check ReadinessIndicatorFile

// fill .MasterPluginJSON

// generate multus config

// use conflist template if cniVersionConfig >= "1.0.0"

// be fixed?

//masterConfigPath

func main() {
	opt := Options{}
	opt.addFlags()
	helpFlag := pflag.BoolP("help", "h", false, "show help message and quit")

	pflag.Parse()
	if *helpFlag {
		pflag.PrintDefaults()
		os.Exit(1)
	}

	err := opt.verifyFileExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// copy multus binary
	if !opt.SkipMultusBinaryCopy {
		// Copy
		if err = cmdutils.CopyFileAtomic(opt.MultusBinFile, opt.CNIBinDir, "_multus", "multus"); err != nil {
			fmt.Fprintf(os.Stderr, "failed at multus copy: %v\n", err)
			return
		}
	}

	var masterConfigHash, caHash, saTokenHash []byte
	var masterConfigFilePath string
	// copy user specified multus conf to CNI conf directory
	if opt.MultusConfFile != "auto" {
		caHash, saTokenHash, err = opt.createKubeConfig(nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create multus kubeconfig: %v\n", err)
			return
		}
		confFileName := filepath.Base(opt.MultusConfFile)
		tempConfFileName := fmt.Sprintf("%s.temp", confFileName)
		if err = cmdutils.CopyFileAtomic(opt.MultusConfFile, opt.CNIConfDir, tempConfFileName, confFileName); err != nil {
			fmt.Fprintf(os.Stderr, "failed at copy multus conf file: %v\n", err)
			return
		}
		fmt.Printf("multus config file %s is copied.\n", opt.MultusConfFile)
	} else { // auto generate multus config
		caHash, saTokenHash, err = opt.createKubeConfig(nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create multus kubeconfig: %v\n", err)
			return
		}
		fmt.Printf("kubeconfig file is created.\n")
		masterConfigFilePath, masterConfigHash, err = opt.createMultusConfig(nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create multus config: %v\n", err)
			return
		}
		fmt.Printf("multus config file is created.\n")
	}

	ctx := signals.SetupSignalHandler()

	if opt.CleanupConfigOnExit {
		defer cleanupMultusConf(&opt)
	}

	watchChanges := opt.CleanupConfigOnExit && opt.MultusConfFile == "auto" && !opt.SkipMultusConfWatch
	if watchChanges {
		fmt.Printf("Entering watch loop...\n")
		masterConfigExists := true

	outer:
		for range time.Tick(1 * time.Second) {
			select {
			case <-ctx.Done():
				// signal received break from loop
				break outer
			default:
				// Check kubeconfig and update if different (i.e. service account updated)
				caHash, saTokenHash, err = opt.createKubeConfig(caHash, saTokenHash)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to update multus kubeconfig: %v\n", err)
					return
				}

				// TODO: should we watch master CNI config (by fsnotify? https://github.com/fsnotify/fsnotify)
				_, err = os.Stat(masterConfigFilePath)

				// if masterConfigFilePath is no longer exists
				if os.IsNotExist(err) {
					if masterConfigExists {
						fmt.Printf("Master plugin @ %q has been deleted. waiting for its restoration...\n", masterConfigFilePath)
					}
					masterConfigExists = false
					continue
				}

				if !masterConfigExists {
					fmt.Printf("Master plugin @ %q was restored. Regenerating given configuration.\n", masterConfigFilePath)
					masterConfigExists = true
				}

				masterConfigFilePath, masterConfigHash, err = opt.createMultusConfig(masterConfigHash)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to create multus config: %v\n", err)
					return
				}
			}
		}
	} else {
		// wait until signal received
		<-ctx.Done()
	}
}

func cleanupMultusConf(opt *Options) {
	_ = "STUB: not implemented"
	// try remove multus conf
	return
}
