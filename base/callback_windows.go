//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build windows

package base

import (
	"syscall"
)

var (
	eventCallback       = syscall.NewCallback(eventCallbackProc)
	removeEventCallback = syscall.NewCallback(removeEventCallbackProc)
)
