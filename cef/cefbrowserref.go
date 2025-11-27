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

	cefTypes "github.com/energye/cef/types"
)

// ICefBrowser Parent: ICefBaseRefCountedRef
type ICefBrowser interface {
	ICefBaseRefCountedRef
	// IsValid
	//  True if this object is currently valid. This will return false (0) after
	//  ICefLifeSpanHandler.OnBeforeClose is called.
	IsValid() bool // function
	// GetHost
	//  Returns the browser host object. This function can only be called in the
	//  browser process.
	GetHost() ICefBrowserHost // function
	// CanGoBack
	//  Returns true (1) if the browser can navigate backwards.
	CanGoBack() bool // function
	// CanGoForward
	//  Returns true (1) if the browser can navigate forwards.
	CanGoForward() bool // function
	// IsLoading
	//  Returns true (1) if the browser is currently loading.
	IsLoading() bool // function
	// GetIdentifier
	//  Returns the globally unique identifier for this browser. This value is
	//  also used as the tabId for extension APIs.
	GetIdentifier() int32 // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that|
	//  object.
	IsSame(that ICefBrowser) bool // function
	// IsPopup
	//  Returns true (1) if the browser is a popup.
	IsPopup() bool // function
	// HasDocument
	//  Returns true (1) if a document has been loaded in the browser.
	HasDocument() bool // function
	// GetMainFrame
	//  Returns the main (top-level) frame for the browser. In the browser process
	//  this will return a valid object until after
	//  ICefLifeSpanHandler.OnBeforeClose is called. In the renderer process
	//  this will return NULL if the main frame is hosted in a different renderer
	//  process (e.g. for cross-origin sub-frames). The main frame object will
	//  change during cross-origin navigation or re-navigation after renderer
	//  process termination (due to crashes, etc).
	GetMainFrame() ICefFrame // function
	// GetFocusedFrame
	//  Returns the focused frame for the browser.
	GetFocusedFrame() ICefFrame // function
	// GetFrameByIdentifier
	//  Returns the frame with the specified identifier, or NULL if not found.
	GetFrameByIdentifier(identifier string) ICefFrame // function
	// GetFrameByName
	//  Returns the frame with the specified name, or NULL if not found.
	GetFrameByName(name string) ICefFrame // function
	// GetFrameCount
	//  Returns the number of frames that currently exist.
	GetFrameCount() cefTypes.NativeUInt // function
	// GetFrameIdentifiers
	//  Returns the identifiers of all existing frames.
	GetFrameIdentifiers(frameIdentifiers *lcl.IStrings) bool // function
	// GetFrameNames
	//  Returns the names of all existing frames.
	GetFrameNames(frameNames *lcl.IStrings) bool // function
	// GoBack
	//  Navigate backwards.
	GoBack() // procedure
	// GoForward
	//  Navigate forwards.
	GoForward() // procedure
	// Reload
	//  Reload the current page.
	Reload() // procedure
	// ReloadIgnoreCache
	//  Reload the current page ignoring any cached data.
	ReloadIgnoreCache() // procedure
	// StopLoad
	//  Stop loading the page.
	StopLoad() // procedure
}

// ICefBrowserRef Parent: ICefBrowser
type ICefBrowserRef interface {
	ICefBrowser
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

func (m *TCefBrowserRef) GetFrameByIdentifier(identifier string) (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserRefAPI().SysCallN(12, m.Instance(), api.PasStr(identifier), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefBrowserRef) GetFrameByName(name string) (result ICefFrame) {
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

func (m *TCefBrowserRef) GetFrameIdentifiers(frameIdentifiers *lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	frameIdentifiersPtr := base.GetObjectUintptr(*frameIdentifiers)
	r := cefBrowserRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&frameIdentifiersPtr)))
	*frameIdentifiers = lcl.AsStrings(frameIdentifiersPtr)
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
			/* 12 */ imports.NewTable("TCefBrowserRef_GetFrameByIdentifier", 0), // function GetFrameByIdentifier
			/* 13 */ imports.NewTable("TCefBrowserRef_GetFrameByName", 0), // function GetFrameByName
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
