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

// ICefNavigationEntry Parent: ICefBaseRefCounted
type ICefNavigationEntry interface {
	ICefBaseRefCounted
	IsValid() bool                                 // function
	GetUrl() string                                // function
	GetDisplayUrl() string                         // function
	GetOriginalUrl() string                        // function
	GetTitle() string                              // function
	GetTransitionType() TCefTransitionType         // function
	HasPostData() bool                             // function
	GetCompletionTime() (resultDateTime TDateTime) // function
	GetHttpStatusCode() int32                      // function
	GetSSLStatus() ICefSSLStatus                   // function
}

// TCefNavigationEntry Parent: TCefBaseRefCounted
type TCefNavigationEntry struct {
	TCefBaseRefCounted
}

// NavigationEntryRef -> ICefNavigationEntry
var NavigationEntryRef navigationEntry

// navigationEntry TCefNavigationEntry Ref
type navigationEntry uintptr

func (m *navigationEntry) UnWrap(data uintptr) ICefNavigationEntry {
	var resultCefNavigationEntry uintptr
	navigationEntryImportAPI().SysCallN(10, uintptr(data), uintptr(unsafePointer(&resultCefNavigationEntry)))
	return AsCefNavigationEntry(resultCefNavigationEntry)
}

func (m *TCefNavigationEntry) IsValid() bool {
	r1 := navigationEntryImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCefNavigationEntry) GetUrl() string {
	r1 := navigationEntryImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetDisplayUrl() string {
	r1 := navigationEntryImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetOriginalUrl() string {
	r1 := navigationEntryImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetTitle() string {
	r1 := navigationEntryImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCefNavigationEntry) GetTransitionType() TCefTransitionType {
	r1 := navigationEntryImportAPI().SysCallN(6, m.Instance())
	return TCefTransitionType(r1)
}

func (m *TCefNavigationEntry) HasPostData() bool {
	r1 := navigationEntryImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCefNavigationEntry) GetCompletionTime() (resultDateTime TDateTime) {
	navigationEntryImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefNavigationEntry) GetHttpStatusCode() int32 {
	r1 := navigationEntryImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefNavigationEntry) GetSSLStatus() ICefSSLStatus {
	var resultCefSSLStatus uintptr
	navigationEntryImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefSSLStatus)))
	return AsCefSSLStatus(resultCefSSLStatus)
}

var (
	navigationEntryImport       *imports.Imports = nil
	navigationEntryImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefNavigationEntry_GetCompletionTime", 0),
		/*1*/ imports.NewTable("CefNavigationEntry_GetDisplayUrl", 0),
		/*2*/ imports.NewTable("CefNavigationEntry_GetHttpStatusCode", 0),
		/*3*/ imports.NewTable("CefNavigationEntry_GetOriginalUrl", 0),
		/*4*/ imports.NewTable("CefNavigationEntry_GetSSLStatus", 0),
		/*5*/ imports.NewTable("CefNavigationEntry_GetTitle", 0),
		/*6*/ imports.NewTable("CefNavigationEntry_GetTransitionType", 0),
		/*7*/ imports.NewTable("CefNavigationEntry_GetUrl", 0),
		/*8*/ imports.NewTable("CefNavigationEntry_HasPostData", 0),
		/*9*/ imports.NewTable("CefNavigationEntry_IsValid", 0),
		/*10*/ imports.NewTable("CefNavigationEntry_UnWrap", 0),
	}
)

func navigationEntryImportAPI() *imports.Imports {
	if navigationEntryImport == nil {
		navigationEntryImport = NewDefaultImports()
		navigationEntryImport.SetImportTable(navigationEntryImportTables)
		navigationEntryImportTables = nil
	}
	return navigationEntryImport
}
