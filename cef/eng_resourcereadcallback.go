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

// ICefResourceReadCallback Parent: ICefBaseRefCounted
//
//	Callback for asynchronous continuation of ICefResourceHandler.read.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_read_callback_t))</a>
type ICefResourceReadCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Callback for asynchronous continuation of read(). If |bytes_read| == 0 the response will be considered complete. If |bytes_read| > 0 then read() will be called again until the request is complete (based on either the result or the expected content length). If |bytes_read| < 0 then the request will fail and the |bytes_read| value will be treated as the error code.
	Cont(bytesread int64) // procedure
}

// TCefResourceReadCallback Parent: TCefBaseRefCounted
//
//	Callback for asynchronous continuation of ICefResourceHandler.read.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_read_callback_t))</a>
type TCefResourceReadCallback struct {
	TCefBaseRefCounted
}

// ResourceReadCallbackRef -> ICefResourceReadCallback
var ResourceReadCallbackRef resourceReadCallback

// resourceReadCallback TCefResourceReadCallback Ref
type resourceReadCallback uintptr

func (m *resourceReadCallback) UnWrap(data uintptr) ICefResourceReadCallback {
	var resultCefResourceReadCallback uintptr
	resourceReadCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefResourceReadCallback)))
	return AsCefResourceReadCallback(resultCefResourceReadCallback)
}

func (m *TCefResourceReadCallback) Cont(bytesread int64) {
	resourceReadCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&bytesread)))
}

var (
	resourceReadCallbackImport       *imports.Imports = nil
	resourceReadCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefResourceReadCallback_Cont", 0),
		/*1*/ imports.NewTable("CefResourceReadCallback_UnWrap", 0),
	}
)

func resourceReadCallbackImportAPI() *imports.Imports {
	if resourceReadCallbackImport == nil {
		resourceReadCallbackImport = NewDefaultImports()
		resourceReadCallbackImport.SetImportTable(resourceReadCallbackImportTables)
		resourceReadCallbackImportTables = nil
	}
	return resourceReadCallbackImport
}
