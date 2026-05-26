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
	"io"

	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/version"
)

// ChrootExec implements invoke.Exec to execute CNI with chroot
type ChrootExec struct {
	Stderr    io.Writer
	chrootDir string
	version.PluginDecoder
}

var _ invoke.Exec = &ChrootExec{}

// ExecPlugin executes CNI plugin with given environment/stdin data.
func (e *ChrootExec) ExecPlugin(ctx context.Context, pluginPath string, stdinData []byte, environ []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// execute delegate CNI with host filesystem context.

// Retry the command on "text file busy" errors

// Command succeeded

// If the plugin is currently about to be written, then we wait a
// second and try it again

// All other errors except than the busy text file

// Copy stderr to caller's buffer in case plugin printed to both
// stdout and stderr for some reason. Ignore failures as stderr is
// only informational.

func (e *ChrootExec) pluginErr(err error, stdout, stderr []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// FindInPath try to find CNI plugin based on given path
func (e *ChrootExec) FindInPath(plugin string, paths []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
