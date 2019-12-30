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

// Package modulecontext package used to maintain the various
// contexts used across
// our application.
package modulecontext

import (
	"context"
	"crud-server/internal/logger"
	"os"
	"path/filepath"
)

// CtxStruct is a wrapper Context used to track the different modules in the app.
// application modules will use this context for it tracking and management.
// We dont create context for every go routine/thread as its over kill.
// There are many short lived go routines can be present in a application, and
// it doesnt need its own context.
// In an application, we keep a hierarchy of context. any
// new module must be inherited from the main module context.
// Any new go routine can inherit a new context from its parent module context
// if needed. Always any module context should be a child of main context.
type CtxStruct struct {
	// go context used to cancel.
	signalCtx  context.Context
	ctxName    string
	CancelFunc context.CancelFunc
	Log        logger.Logging
}

// CreateCancelCtx function to create a context with cancel handler.
// Any new module/go routine must create a cancel context from its parent
// and use it for tracking the cancel signal from the parent module/thread.
func CreateCancelCtx(name string, parentCtx *context.Context) *CtxStruct {
	var newCtx *CtxStruct
	newCtx = new(CtxStruct)
	if parentCtx == nil {
		// Must be the main context.
		newCtx.signalCtx, newCtx.CancelFunc =
			context.WithCancel(context.Background())
	} else {
		newCtx.signalCtx, newCtx.CancelFunc =
			context.WithCancel(*parentCtx)
	}
	newCtx.ctxName = name
	return newCtx
}

// InitContextLogger function to initialize the logger in the specific handler.
// By default there are no logger present in the context.
func (ctx *CtxStruct) InitContextLogger(loglevel logger.LogLeveltype,
	filename string) {
	//Create the log directory first before
	filename, _ = filepath.Abs(filename)
	logDir := filepath.Dir(filename)
	os.Mkdir(logDir, 0700)
	ctx.Log.LogInit(loglevel, filename)
}
