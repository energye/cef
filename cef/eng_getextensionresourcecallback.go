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

// ICefGetExtensionResourceCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of ICefExtensionHandler.GetExtensionResource.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_get_extension_resource_callback_t))</a>
type ICefGetExtensionResourceCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the request. Read the resource contents from |stream|.
	Cont(stream ICefStreamReader) // procedure
	// Cancel
	//  Cancel the request.
	Cancel() // procedure
}

// TCefGetExtensionResourceCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of ICefExtensionHandler.GetExtensionResource.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_get_extension_resource_callback_t))</a>
type TCefGetExtensionResourceCallback struct {
	TCefBaseRefCounted
}

// GetExtensionResourceCallbackRef -> ICefGetExtensionResourceCallback
var GetExtensionResourceCallbackRef getExtensionResourceCallback

// getExtensionResourceCallback TCefGetExtensionResourceCallback Ref
type getExtensionResourceCallback uintptr

func (m *getExtensionResourceCallback) UnWrap(data uintptr) ICefGetExtensionResourceCallback {
	var resultCefGetExtensionResourceCallback uintptr
	getExtensionResourceCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefGetExtensionResourceCallback)))
	return AsCefGetExtensionResourceCallback(resultCefGetExtensionResourceCallback)
}

func (m *TCefGetExtensionResourceCallback) Cont(stream ICefStreamReader) {
	getExtensionResourceCallbackImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(stream))
}

func (m *TCefGetExtensionResourceCallback) Cancel() {
	getExtensionResourceCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	getExtensionResourceCallbackImport       *imports.Imports = nil
	getExtensionResourceCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefGetExtensionResourceCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefGetExtensionResourceCallback_Cont", 0),
		/*2*/ imports.NewTable("CefGetExtensionResourceCallback_UnWrap", 0),
	}
)

func getExtensionResourceCallbackImportAPI() *imports.Imports {
	if getExtensionResourceCallbackImport == nil {
		getExtensionResourceCallbackImport = NewDefaultImports()
		getExtensionResourceCallbackImport.SetImportTable(getExtensionResourceCallbackImportTables)
		getExtensionResourceCallbackImportTables = nil
	}
	return getExtensionResourceCallbackImport
}
