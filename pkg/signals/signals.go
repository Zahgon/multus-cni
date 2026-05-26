// Copyright (c) 2024 Multus Authors
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

// Package signals provides handling for os signals.
package signals

import (
	"context"
)

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler registers for SIGTERM and SIGINT. A context is returned
// which is canceled on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() context.Context {
	_ = "STUB: not implemented"
	return *
	// panics when called twice
	new(context.Context)
}

// second signal. Exit directly.
