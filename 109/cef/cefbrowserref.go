//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/109/types"
)

// ICefBrowser Parent: ICefBaseRefCounted
type ICefBrowser interface {
	ICefBaseRefCounted
	IsValid() bool                                                                                            // function
	GetHost() ICefBrowserHost                                                                                 // function
	CanGoBack() bool                                                                                          // function
	CanGoForward() bool                                                                                       // function
	IsLoading() bool                                                                                          // function
	GetIdentifier() int32                                                                                     // function
	IsSame(that ICefBrowser) bool                                                                             // function
	IsPopup() bool                                                                                            // function
	HasDocument() bool                                                                                        // function
	GetMainFrame() ICefFrame                                                                                  // function
	GetFocusedFrame() ICefFrame                                                                               // function
	GetFrameByident(identifier int64) ICefFrame                                                               // function
	GetFrame(name string) ICefFrame                                                                           // function
	GetFrameCount() cefTypes.NativeUInt                                                                       // function
	GetFrameIdentifiers(frameCount *cefTypes.NativeUInt, frameIdentifierArray *ICefFrameIdentifierArray) bool // function
	GetFrameNames(frameNames *lcl.IStrings) bool                                                              // function
	GoBack()                                                                                                  // procedure
	GoForward()                                                                                               // procedure
	Reload()                                                                                                  // procedure
	ReloadIgnoreCache()                                                                                       // procedure
	StopLoad()                                                                                                // procedure
}

// ICefBrowserRef Parent: ICefBrowser ICefBaseRefCountedRef
type ICefBrowserRef interface {
	ICefBrowser
	ICefBaseRefCountedRef
	AsIntfBrowser() uintptr
}

type TCefBrowserRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefBrowserRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) GetHost() (result ICefBrowserHost) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserHostRef(resultPtr)
	return
}

func (m *TCefBrowserRef) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) GetIdentifier() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefBrowserRef) IsSame(that ICefBrowser) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefBrowserRef) IsPopup() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) HasDocument() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserRefAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserRef) GetMainFrame() (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefBrowserRef) GetFocusedFrame() (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefBrowserRef) GetFrameByident(identifier int64) (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&identifier)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefBrowserRef) GetFrame(name string) (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(13, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefBrowserRef) GetFrameCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserRefAPI().SysCallN(14, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefBrowserRef) GetFrameIdentifiers(frameCount *cefTypes.NativeUInt, frameIdentifierArray *ICefFrameIdentifierArray) bool {
	if !m.IsValid() {
		return false
	}
	frameCountPtr := uintptr(*frameCount)
	var frameIdentifierArrayPtr uintptr
	var frameIdentifierArrayCountPtr uintptr
	r := cefBrowserRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&frameCountPtr)), uintptr(base.UnsafePointer(&frameIdentifierArrayPtr)), uintptr(base.UnsafePointer(&frameIdentifierArrayCountPtr)))
	*frameCount = cefTypes.NativeUInt(frameCountPtr)
	*frameIdentifierArray = NewCefFrameIdentifierArray(int(frameIdentifierArrayCountPtr), frameIdentifierArrayPtr)
	return api.GoBool(r)
}

func (m *TCefBrowserRef) GetFrameNames(frameNames *lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	frameNamesPtr := base.GetObjectUintptr(*frameNames)
	r := cefBrowserRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&frameNamesPtr)))
	*frameNames = lcl.AsStrings(frameNamesPtr)
	return api.GoBool(r)
}

func (m *TCefBrowserRef) GoBack() {
	if !m.IsValid() {
		return
	}
	cefBrowserRefAPI().SysCallN(18, m.Instance())
}

func (m *TCefBrowserRef) GoForward() {
	if !m.IsValid() {
		return
	}
	cefBrowserRefAPI().SysCallN(19, m.Instance())
}

func (m *TCefBrowserRef) Reload() {
	if !m.IsValid() {
		return
	}
	cefBrowserRefAPI().SysCallN(20, m.Instance())
}

func (m *TCefBrowserRef) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	cefBrowserRefAPI().SysCallN(21, m.Instance())
}

func (m *TCefBrowserRef) StopLoad() {
	if !m.IsValid() {
		return
	}
	cefBrowserRefAPI().SysCallN(22, m.Instance())
}

func (m *TCefBrowserRef) AsIntfBrowser() uintptr {
	return m.GetIntfPointer(0)
}

// BrowserRef  is static instance
var BrowserRef _BrowserRefClass

// _BrowserRefClass is class type defined by TCefBrowserRef
type _BrowserRefClass uintptr

func (_BrowserRefClass) UnWrap(data uintptr) (result ICefBrowser) {
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(17, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

// NewBrowserRef class constructor
func NewBrowserRef(data uintptr) ICefBrowserRef {
	var browserPtr uintptr // ICefBrowser
	r := cefBrowserRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&browserPtr)))
	ret := AsCefBrowserRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserPtr)
	}
	return ret
}

var (
	cefBrowserRefOnce   base.Once
	cefBrowserRefImport *imports.Imports = nil
)

func cefBrowserRefAPI() *imports.Imports {
	cefBrowserRefOnce.Do(func() {
		cefBrowserRefImport = api.NewDefaultImports()
		cefBrowserRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserRef_Create", 0), // constructor NewBrowserRef
			/* 1 */ imports.NewTable("TCefBrowserRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefBrowserRef_GetHost", 0), // function GetHost
			/* 3 */ imports.NewTable("TCefBrowserRef_CanGoBack", 0), // function CanGoBack
			/* 4 */ imports.NewTable("TCefBrowserRef_CanGoForward", 0), // function CanGoForward
			/* 5 */ imports.NewTable("TCefBrowserRef_IsLoading", 0), // function IsLoading
			/* 6 */ imports.NewTable("TCefBrowserRef_GetIdentifier", 0), // function GetIdentifier
			/* 7 */ imports.NewTable("TCefBrowserRef_IsSame", 0), // function IsSame
			/* 8 */ imports.NewTable("TCefBrowserRef_IsPopup", 0), // function IsPopup
			/* 9 */ imports.NewTable("TCefBrowserRef_HasDocument", 0), // function HasDocument
			/* 10 */ imports.NewTable("TCefBrowserRef_GetMainFrame", 0), // function GetMainFrame
			/* 11 */ imports.NewTable("TCefBrowserRef_GetFocusedFrame", 0), // function GetFocusedFrame
			/* 12 */ imports.NewTable("TCefBrowserRef_GetFrameByident", 0), // function GetFrameByident
			/* 13 */ imports.NewTable("TCefBrowserRef_GetFrame", 0), // function GetFrame
			/* 14 */ imports.NewTable("TCefBrowserRef_GetFrameCount", 0), // function GetFrameCount
			/* 15 */ imports.NewTable("TCefBrowserRef_GetFrameIdentifiers", 0), // function GetFrameIdentifiers
			/* 16 */ imports.NewTable("TCefBrowserRef_GetFrameNames", 0), // function GetFrameNames
			/* 17 */ imports.NewTable("TCefBrowserRef_UnWrap", 0), // static function UnWrap
			/* 18 */ imports.NewTable("TCefBrowserRef_GoBack", 0), // procedure GoBack
			/* 19 */ imports.NewTable("TCefBrowserRef_GoForward", 0), // procedure GoForward
			/* 20 */ imports.NewTable("TCefBrowserRef_Reload", 0), // procedure Reload
			/* 21 */ imports.NewTable("TCefBrowserRef_ReloadIgnoreCache", 0), // procedure ReloadIgnoreCache
			/* 22 */ imports.NewTable("TCefBrowserRef_StopLoad", 0), // procedure StopLoad
		}
	})
	return cefBrowserRefImport
}
