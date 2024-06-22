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

// ICefBrowser Parent: ICefBaseRefCounted
type ICefBrowser interface {
	ICefBaseRefCounted
	IsValid() bool                                                                                     // function
	GetHost() ICefBrowserHost                                                                          // function
	CanGoBack() bool                                                                                   // function
	CanGoForward() bool                                                                                // function
	IsLoading() bool                                                                                   // function
	GetIdentifier() int32                                                                              // function
	IsSame(that ICefBrowser) bool                                                                      // function
	IsPopup() bool                                                                                     // function
	HasDocument() bool                                                                                 // function
	GetMainFrame() ICefFrame                                                                           // function
	GetFocusedFrame() ICefFrame                                                                        // function
	GetFrameByident(identifier int64) ICefFrame                                                        // function
	GetFrame(name string) ICefFrame                                                                    // function
	GetFrameCount() NativeUInt                                                                         // function
	GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool // function
	GetFrameNames(aFrameNames *IStrings) bool                                                          // function
	GoBack()                                                                                           // procedure
	GoForward()                                                                                        // procedure
	Reload()                                                                                           // procedure
	ReloadIgnoreCache()                                                                                // procedure
	StopLoad()                                                                                         // procedure
}

// TCefBrowser Parent: TCefBaseRefCounted
type TCefBrowser struct {
	TCefBaseRefCounted
}

// BrowserRef -> ICefBrowser
var BrowserRef browser

// browser TCefBrowser Ref
type browser uintptr

func (m *browser) UnWrap(data uintptr) ICefBrowser {
	var resultCefBrowser uintptr
	browserImportAPI().SysCallN(21, uintptr(data), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefBrowser) IsValid() bool {
	r1 := browserImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetHost() ICefBrowserHost {
	var resultCefBrowserHost uintptr
	browserImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefBrowserHost)))
	return AsCefBrowserHost(resultCefBrowserHost)
}

func (m *TCefBrowser) CanGoBack() bool {
	r1 := browserImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) CanGoForward() bool {
	r1 := browserImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) IsLoading() bool {
	r1 := browserImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetIdentifier() int32 {
	r1 := browserImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCefBrowser) IsSame(that ICefBrowser) bool {
	r1 := browserImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBrowser) IsPopup() bool {
	r1 := browserImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) HasDocument() bool {
	r1 := browserImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetMainFrame() ICefFrame {
	var resultCefFrame uintptr
	browserImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFocusedFrame() ICefFrame {
	var resultCefFrame uintptr
	browserImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrameByident(identifier int64) ICefFrame {
	var resultCefFrame uintptr
	browserImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&identifier)), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrame(name string) ICefFrame {
	var resultCefFrame uintptr
	browserImportAPI().SysCallN(3, m.Instance(), PascalStr(name), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrameCount() NativeUInt {
	r1 := browserImportAPI().SysCallN(5, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefBrowser) GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := browserImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*aFrameCount = NativeUInt(result0)
	*aFrameIdentifierArray = FrameIdentifierArrayRef.New(int(*aFrameCount), result1)
	return GoBool(r1)
}

func (m *TCefBrowser) GetFrameNames(aFrameNames *IStrings) bool {
	var result0 uintptr
	r1 := browserImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&result0)))
	*aFrameNames = AsStrings(result0)
	return GoBool(r1)
}

func (m *TCefBrowser) GoBack() {
	browserImportAPI().SysCallN(11, m.Instance())
}

func (m *TCefBrowser) GoForward() {
	browserImportAPI().SysCallN(12, m.Instance())
}

func (m *TCefBrowser) Reload() {
	browserImportAPI().SysCallN(18, m.Instance())
}

func (m *TCefBrowser) ReloadIgnoreCache() {
	browserImportAPI().SysCallN(19, m.Instance())
}

func (m *TCefBrowser) StopLoad() {
	browserImportAPI().SysCallN(20, m.Instance())
}

var (
	browserImport       *imports.Imports = nil
	browserImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBrowser_CanGoBack", 0),
		/*1*/ imports.NewTable("CefBrowser_CanGoForward", 0),
		/*2*/ imports.NewTable("CefBrowser_GetFocusedFrame", 0),
		/*3*/ imports.NewTable("CefBrowser_GetFrame", 0),
		/*4*/ imports.NewTable("CefBrowser_GetFrameByident", 0),
		/*5*/ imports.NewTable("CefBrowser_GetFrameCount", 0),
		/*6*/ imports.NewTable("CefBrowser_GetFrameIdentifiers", 0),
		/*7*/ imports.NewTable("CefBrowser_GetFrameNames", 0),
		/*8*/ imports.NewTable("CefBrowser_GetHost", 0),
		/*9*/ imports.NewTable("CefBrowser_GetIdentifier", 0),
		/*10*/ imports.NewTable("CefBrowser_GetMainFrame", 0),
		/*11*/ imports.NewTable("CefBrowser_GoBack", 0),
		/*12*/ imports.NewTable("CefBrowser_GoForward", 0),
		/*13*/ imports.NewTable("CefBrowser_HasDocument", 0),
		/*14*/ imports.NewTable("CefBrowser_IsLoading", 0),
		/*15*/ imports.NewTable("CefBrowser_IsPopup", 0),
		/*16*/ imports.NewTable("CefBrowser_IsSame", 0),
		/*17*/ imports.NewTable("CefBrowser_IsValid", 0),
		/*18*/ imports.NewTable("CefBrowser_Reload", 0),
		/*19*/ imports.NewTable("CefBrowser_ReloadIgnoreCache", 0),
		/*20*/ imports.NewTable("CefBrowser_StopLoad", 0),
		/*21*/ imports.NewTable("CefBrowser_UnWrap", 0),
	}
)

func browserImportAPI() *imports.Imports {
	if browserImport == nil {
		browserImport = NewDefaultImports()
		browserImport.SetImportTable(browserImportTables)
		browserImportTables = nil
	}
	return browserImport
}
