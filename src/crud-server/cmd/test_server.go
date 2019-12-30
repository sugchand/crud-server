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

// Package main for the test-server. Application testing is done with
// this main(test_main.go).
package main

import (
	"crud-server/configs"
	"crud-server/internal/logger"
	"crud-server/internal/modulecontext"
)

func main() {
	var mainCtx *modulecontext.CtxStruct
	mainCtx = modulecontext.CreateCancelCtx("main", nil)
	mainCtx.InitContextLogger(logger.LogLeveltype(configs.LogLevel),
		configs.LogPath+"/main.log")
	mainCtx.Log.Trace("main started")
	mainCtx.CancelFunc()
}
