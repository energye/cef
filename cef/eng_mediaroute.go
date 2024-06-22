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

// ICefMediaRoute Parent: ICefBaseRefCounted
type ICefMediaRoute interface {
	ICefBaseRefCounted
	GetId() string                   // function
	GetSource() ICefMediaSource      // function
	GetSink() ICefMediaSink          // function
	SendRouteMessage(message string) // procedure
	Terminate()                      // procedure
}

// TCefMediaRoute Parent: TCefBaseRefCounted
type TCefMediaRoute struct {
	TCefBaseRefCounted
}

// MediaRouteRef -> ICefMediaRoute
var MediaRouteRef mediaRoute

// mediaRoute TCefMediaRoute Ref
type mediaRoute uintptr

func (m *mediaRoute) UnWrap(data uintptr) ICefMediaRoute {
	var resultCefMediaRoute uintptr
	mediaRouteImportAPI().SysCallN(5, uintptr(data), uintptr(unsafePointer(&resultCefMediaRoute)))
	return AsCefMediaRoute(resultCefMediaRoute)
}

func (m *TCefMediaRoute) GetId() string {
	r1 := mediaRouteImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaRoute) GetSource() ICefMediaSource {
	var resultCefMediaSource uintptr
	mediaRouteImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TCefMediaRoute) GetSink() ICefMediaSink {
	var resultCefMediaSink uintptr
	mediaRouteImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefMediaSink)))
	return AsCefMediaSink(resultCefMediaSink)
}

func (m *TCefMediaRoute) SendRouteMessage(message string) {
	mediaRouteImportAPI().SysCallN(3, m.Instance(), PascalStr(message))
}

func (m *TCefMediaRoute) Terminate() {
	mediaRouteImportAPI().SysCallN(4, m.Instance())
}

var (
	mediaRouteImport       *imports.Imports = nil
	mediaRouteImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMediaRoute_GetId", 0),
		/*1*/ imports.NewTable("CefMediaRoute_GetSink", 0),
		/*2*/ imports.NewTable("CefMediaRoute_GetSource", 0),
		/*3*/ imports.NewTable("CefMediaRoute_SendRouteMessage", 0),
		/*4*/ imports.NewTable("CefMediaRoute_Terminate", 0),
		/*5*/ imports.NewTable("CefMediaRoute_UnWrap", 0),
	}
)

func mediaRouteImportAPI() *imports.Imports {
	if mediaRouteImport == nil {
		mediaRouteImport = NewDefaultImports()
		mediaRouteImport.SetImportTable(mediaRouteImportTables)
		mediaRouteImportTables = nil
	}
	return mediaRouteImport
}
