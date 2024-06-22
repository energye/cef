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

// ICefRequestContextHandler Parent: ICefBaseRefCounted
type ICefRequestContextHandler interface {
	ICefBaseRefCounted
	OnRequestContextInitialized(requestcontext ICefRequestContext)                                                                                                                                                                  // procedure
	GetResourceRequestHandler(browser ICefBrowser, frame ICefFrame, request ICefRequest, isnavigation, isdownload bool, requestinitiator string, disabledefaulthandling *bool, aResourceRequestHandler *ICefResourceRequestHandler) // procedure
	RemoveReferences()                                                                                                                                                                                                              // procedure
}

// TCefRequestContextHandler Parent: TCefBaseRefCounted
type TCefRequestContextHandler struct {
	TCefBaseRefCounted
}

// RequestContextHandlerRef -> ICefRequestContextHandler
var RequestContextHandlerRef requestContextHandler

// requestContextHandler TCefRequestContextHandler Ref
type requestContextHandler uintptr

func (m *requestContextHandler) UnWrap(data uintptr) ICefRequestContextHandler {
	var resultCefRequestContextHandler uintptr
	requestContextHandlerImportAPI().SysCallN(3, uintptr(data), uintptr(unsafePointer(&resultCefRequestContextHandler)))
	return AsCefRequestContextHandler(resultCefRequestContextHandler)
}

func (m *TCefRequestContextHandler) OnRequestContextInitialized(requestcontext ICefRequestContext) {
	requestContextHandlerImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(requestcontext))
}

func (m *TCefRequestContextHandler) GetResourceRequestHandler(browser ICefBrowser, frame ICefFrame, request ICefRequest, isnavigation, isdownload bool, requestinitiator string, disabledefaulthandling *bool, aResourceRequestHandler *ICefResourceRequestHandler) {
	var result5 uintptr
	var result6 uintptr
	requestContextHandlerImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(browser), GetObjectUintptr(frame), GetObjectUintptr(request), PascalBool(isnavigation), PascalBool(isdownload), PascalStr(requestinitiator), uintptr(unsafePointer(&result5)), uintptr(unsafePointer(&result6)))
	*disabledefaulthandling = GoBool(result5)
	*aResourceRequestHandler = AsCefResourceRequestHandler(result6)
}

func (m *TCefRequestContextHandler) RemoveReferences() {
	requestContextHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	requestContextHandlerImport       *imports.Imports = nil
	requestContextHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefRequestContextHandler_GetResourceRequestHandler", 0),
		/*1*/ imports.NewTable("CefRequestContextHandler_OnRequestContextInitialized", 0),
		/*2*/ imports.NewTable("CefRequestContextHandler_RemoveReferences", 0),
		/*3*/ imports.NewTable("CefRequestContextHandler_UnWrap", 0),
	}
)

func requestContextHandlerImportAPI() *imports.Imports {
	if requestContextHandlerImport == nil {
		requestContextHandlerImport = NewDefaultImports()
		requestContextHandlerImport.SetImportTable(requestContextHandlerImportTables)
		requestContextHandlerImportTables = nil
	}
	return requestContextHandlerImport
}
