// Copyright 2019 Sugesh Chandran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package configs is a module used across the application for managing
// application configuration.
// Update the application configuration in this file when starting the
// application. No runtime configuration is supported with this module
package configs

import (
	"crud-server/internal/logger"
)

var (
	//LogLevel Default log level for the application
	LogLevel = logger.Trace
	//LogPath is the path to store different logs from the modules
	LogPath = "/tmp/log"
)
