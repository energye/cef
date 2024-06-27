//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IV8ArrayBufferReleaseCallback Parent: ICefV8ArrayBufferReleaseCallback
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a cref="uCEFTypes|TCefv8ArrayBufferReleaseCallback">Implements TCefv8ArrayBufferReleaseCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t)</a>
type IV8ArrayBufferReleaseCallback interface {
	ICefV8ArrayBufferReleaseCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefV8ArrayBufferReleaseCallback // procedure
	// SetOnReleaseBuffer
	//  Called to release |buffer| when the ArrayBuffer JS object is garbage
	//  collected. |buffer| is the value that was passed to CreateArrayBuffer
	//  along with this object.
	SetOnReleaseBuffer(fn TOnV8ArrayBufferReleaseBuffer) // property event
}

// TV8ArrayBufferReleaseCallback Parent: TCefV8ArrayBufferReleaseCallback
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a cref="uCEFTypes|TCefv8ArrayBufferReleaseCallback">Implements TCefv8ArrayBufferReleaseCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t)</a>
type TV8ArrayBufferReleaseCallback struct {
	TCefV8ArrayBufferReleaseCallback
	releaseBufferPtr uintptr
}

func NewV8ArrayBufferReleaseCallback() IV8ArrayBufferReleaseCallback {
	r1 := v8ArrayBufferReleaseCallbackImportAPI().SysCallN(1)
	return AsV8ArrayBufferReleaseCallback(r1)
}

func (m *TV8ArrayBufferReleaseCallback) AsInterface() ICefV8ArrayBufferReleaseCallback {
	var resultCefV8ArrayBufferReleaseCallback uintptr
	v8ArrayBufferReleaseCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefV8ArrayBufferReleaseCallback)))
	return AsCefV8ArrayBufferReleaseCallback(resultCefV8ArrayBufferReleaseCallback)
}

func (m *TV8ArrayBufferReleaseCallback) SetOnReleaseBuffer(fn TOnV8ArrayBufferReleaseBuffer) {
	if m.releaseBufferPtr != 0 {
		RemoveEventElement(m.releaseBufferPtr)
	}
	m.releaseBufferPtr = MakeEventDataPtr(fn)
	v8ArrayBufferReleaseCallbackImportAPI().SysCallN(2, m.Instance(), m.releaseBufferPtr)
}

var (
	v8ArrayBufferReleaseCallbackImport       *imports.Imports = nil
	v8ArrayBufferReleaseCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("V8ArrayBufferReleaseCallback_AsInterface", 0),
		/*1*/ imports.NewTable("V8ArrayBufferReleaseCallback_Create", 0),
		/*2*/ imports.NewTable("V8ArrayBufferReleaseCallback_SetOnReleaseBuffer", 0),
	}
)

func v8ArrayBufferReleaseCallbackImportAPI() *imports.Imports {
	if v8ArrayBufferReleaseCallbackImport == nil {
		v8ArrayBufferReleaseCallbackImport = NewDefaultImports()
		v8ArrayBufferReleaseCallbackImport.SetImportTable(v8ArrayBufferReleaseCallbackImportTables)
		v8ArrayBufferReleaseCallbackImportTables = nil
	}
	return v8ArrayBufferReleaseCallbackImport
}
