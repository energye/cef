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

// ICefMediaRouter Parent: ICefBaseRefCounted
//
//	Supports discovery of and communication with media devices on the local network via the Cast and DIAL protocols. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_router_t))</a>
type ICefMediaRouter interface {
	ICefBaseRefCounted
	// AddObserver
	//  Add an observer for MediaRouter events. The observer will remain registered until the returned Registration object is destroyed.
	AddObserver(observer ICefMediaObserver) ICefRegistration // function
	// GetSource
	//  Returns a MediaSource object for the specified media source URN. Supported URN schemes include "cast:" and "dial:", and will be already known by the client application (e.g. "cast:<appId>?clientId=<clientId>").
	GetSource(urn string) ICefMediaSource // function
	// NotifyCurrentSinks
	//  Trigger an asynchronous call to ICefMediaObserver.OnSinks on all registered observers.
	NotifyCurrentSinks() // procedure
	// CreateRoute
	//  Create a new route between |source| and |sink|. Source and sink must be valid, compatible (as reported by ICefMediaSink.IsCompatibleWith), and a route between them must not already exist. |callback| will be executed on success or failure. If route creation succeeds it will also trigger an asynchronous call to ICefMediaObserver.OnRoutes on all registered observers.
	CreateRoute(source ICefMediaSource, sink ICefMediaSink, callback ICefMediaRouteCreateCallback) // procedure
	// NotifyCurrentRoutes
	//  Trigger an asynchronous call to ICefMediaObserver.OnRoutes on all registered observers.
	NotifyCurrentRoutes() // procedure
}

// TCefMediaRouter Parent: TCefBaseRefCounted
//
//	Supports discovery of and communication with media devices on the local network via the Cast and DIAL protocols. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_router_t))</a>
type TCefMediaRouter struct {
	TCefBaseRefCounted
}

// MediaRouterRef -> ICefMediaRouter
var MediaRouterRef mediaRouter

// mediaRouter TCefMediaRouter Ref
type mediaRouter uintptr

func (m *mediaRouter) UnWrap(data uintptr) ICefMediaRouter {
	var resultCefMediaRouter uintptr
	mediaRouterImportAPI().SysCallN(6, uintptr(data), uintptr(unsafePointer(&resultCefMediaRouter)))
	return AsCefMediaRouter(resultCefMediaRouter)
}

func (m *mediaRouter) Global() ICefMediaRouter {
	var resultCefMediaRouter uintptr
	mediaRouterImportAPI().SysCallN(3, uintptr(unsafePointer(&resultCefMediaRouter)))
	return AsCefMediaRouter(resultCefMediaRouter)
}

func (m *TCefMediaRouter) AddObserver(observer ICefMediaObserver) ICefRegistration {
	var resultCefRegistration uintptr
	mediaRouterImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(observer), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TCefMediaRouter) GetSource(urn string) ICefMediaSource {
	var resultCefMediaSource uintptr
	mediaRouterImportAPI().SysCallN(2, m.Instance(), PascalStr(urn), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TCefMediaRouter) NotifyCurrentSinks() {
	mediaRouterImportAPI().SysCallN(5, m.Instance())
}

func (m *TCefMediaRouter) CreateRoute(source ICefMediaSource, sink ICefMediaSink, callback ICefMediaRouteCreateCallback) {
	mediaRouterImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(source), GetObjectUintptr(sink), GetObjectUintptr(callback))
}

func (m *TCefMediaRouter) NotifyCurrentRoutes() {
	mediaRouterImportAPI().SysCallN(4, m.Instance())
}

var (
	mediaRouterImport       *imports.Imports = nil
	mediaRouterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMediaRouter_AddObserver", 0),
		/*1*/ imports.NewTable("CefMediaRouter_CreateRoute", 0),
		/*2*/ imports.NewTable("CefMediaRouter_GetSource", 0),
		/*3*/ imports.NewTable("CefMediaRouter_Global", 0),
		/*4*/ imports.NewTable("CefMediaRouter_NotifyCurrentRoutes", 0),
		/*5*/ imports.NewTable("CefMediaRouter_NotifyCurrentSinks", 0),
		/*6*/ imports.NewTable("CefMediaRouter_UnWrap", 0),
	}
)

func mediaRouterImportAPI() *imports.Imports {
	if mediaRouterImport == nil {
		mediaRouterImport = NewDefaultImports()
		mediaRouterImport.SetImportTable(mediaRouterImportTables)
		mediaRouterImportTables = nil
	}
	return mediaRouterImport
}
