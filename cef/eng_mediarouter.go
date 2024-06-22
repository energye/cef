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
type ICefMediaRouter interface {
	ICefBaseRefCounted
	AddObserver(observer ICefMediaObserver) ICefRegistration                                       // function
	GetSource(urn string) ICefMediaSource                                                          // function
	NotifyCurrentSinks()                                                                           // procedure
	CreateRoute(source ICefMediaSource, sink ICefMediaSink, callback ICefMediaRouteCreateCallback) // procedure
	NotifyCurrentRoutes()                                                                          // procedure
}

// TCefMediaRouter Parent: TCefBaseRefCounted
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
