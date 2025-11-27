//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build !windows && cgo
// +build !windows,cgo

package cef

//// #cgo darwin CFLAGS: -mmacosx-version-min=11.0 -DMACOSX_DEPLOYMENT_TARGET=11.0
// #cgo darwin CFLAGS: -mmacosx-version-min=11.0
// #cgo darwin LDFLAGS: -mmacosx-version-min=11.0
// #include <stdint.h>
//
// extern void* doCEFEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doCEFEventCallbackAddr() {
//    return &doCEFEventCallbackProc;
// }
//
// extern void* doCEFRemoveEventCallbackProc(uintptr_t ptr);
// static void* doCEFRemoveEventCallbackAddr() {
//    return &doCEFRemoveEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doCEFEventCallbackProc
func doCEFEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doCEFRemoveEventCallbackProc
func doCEFRemoveEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

var (
	eventCallback       = uintptr(C.doCEFEventCallbackAddr())
	removeEventCallback = uintptr(C.doCEFRemoveEventCallbackAddr())
)
