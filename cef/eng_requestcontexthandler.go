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
//
//	Implement this interface to provide handler implementations. The handler instance will not be released until all objects related to the context have been destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_handler_capi.h">CEF source file: /include/capi/cef_request_context_handler_capi.h (cef_request_context_handler_t))</a>
type ICefRequestContextHandler interface {
	ICefBaseRefCounted
	// OnRequestContextInitialized
	//  Called on the browser process UI thread immediately after the request context has been initialized.
	OnRequestContextInitialized(requestcontext ICefRequestContext) // procedure
	// GetResourceRequestHandler
	//  Called on the browser process IO thread before a resource request is initiated. The |browser| and |frame| values represent the source of the request, and may be NULL for requests originating from service workers or ICefUrlRequest. |request| represents the request contents and cannot be modified in this callback. |is_navigation| will be true (1) if the resource request is a navigation. |is_download| will be true (1) if the resource request is a download. |request_initiator| is the origin (scheme + domain) of the page that initiated the request. Set |disable_default_handling| to true (1) to disable default handling of the request, in which case it will need to be handled via ICefResourceRequestHandler.GetResourceHandler or it will be canceled. To allow the resource load to proceed with default handling return NULL. To specify a handler for the resource return a ICefResourceRequestHandler object. This function will not be called if the client associated with |browser| returns a non-NULL value from ICefRequestHandler.GetResourceRequestHandler for the same request (identified by ICefRequest.GetIdentifier).
	GetResourceRequestHandler(browser ICefBrowser, frame ICefFrame, request ICefRequest, isnavigation, isdownload bool, requestinitiator string, disabledefaulthandling *bool, aResourceRequestHandler *ICefResourceRequestHandler) // procedure
	// RemoveReferences
	//  Custom procedure to clear all references.
	RemoveReferences() // procedure
}

// TCefRequestContextHandler Parent: TCefBaseRefCounted
//
//	Implement this interface to provide handler implementations. The handler instance will not be released until all objects related to the context have been destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_handler_capi.h">CEF source file: /include/capi/cef_request_context_handler_capi.h (cef_request_context_handler_t))</a>
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
