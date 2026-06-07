//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefBrowserView Parent: ICefView
type ICefBrowserView interface {
	ICefView
	// GetBrowser
	//  Returns the ICefBrowser hosted by this BrowserView. Will return NULL if
	//  the browser has not yet been created or has already been destroyed.
	GetBrowser() ICefBrowser // function
	// GetChromeToolbar
	//  Returns the Chrome toolbar associated with this BrowserView. Only
	//  supported when using Chrome style. The ICefBrowserViewDelegate.GetChromeToolbarType
	//  function must return a value other than
	//  CEF_CTT_NONE and the toolbar will not be available until after this
	//  BrowserView is added to a ICefWindow and
	//  ICefViewDelegate.OnWindowChanged() has been called.
	GetChromeToolbar() ICefView // function
	// GetRuntimeStyle
	//  Returns the runtime style for this BrowserView (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	GetRuntimeStyle() cefTypes.TCefRuntimeStyle // function
	// SetPreferAccelerators
	//  Sets whether normal priority accelerators are first forwarded to the web
	//  content (`keydown` event handler) or ICefKeyboardHandler. Normal priority
	//  accelerators can be registered via ICefWindow.SetAccelerator (with
	//  |high_priority|=false) or internally for standard accelerators supported
	//  by Chrome style. If |prefer_accelerators| is true then the matching
	//  accelerator will be triggered immediately (calling
	//  ICefWindowDelegate.OnAccelerator or ICefCommandHandler.OnChromeCommand
	//  respectively) and the event will not be forwarded to the web content or
	//  ICefKeyboardHandler first. If |prefer_accelerators| is false then the
	//  matching accelerator will only be triggered if the event is not handled by
	//  web content (`keydown` event handler that calls `event.preventDefault()`)
	//  or by ICefKeyboardHandler. The default value is false.
	SetPreferAccelerators(preferAccelerators bool) // procedure
}

// ICefBrowserViewRef Parent: ICefBrowserView ICefViewRef
type ICefBrowserViewRef interface {
	ICefBrowserView
	ICefViewRef
	AsIntfBrowserView() uintptr
	AsIntfView() uintptr
}

type TCefBrowserViewRef struct {
	TCefViewRef
}

func (m *TCefBrowserViewRef) GetBrowser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserViewRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCefBrowserViewRef) GetChromeToolbar() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserViewRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefBrowserViewRef) GetRuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserViewRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TCefBrowserViewRef) SetPreferAccelerators(preferAccelerators bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserViewRefAPI().SysCallN(7, m.Instance(), api.PasBool(preferAccelerators))
}

func (m *TCefBrowserViewRef) AsIntfBrowserView() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefBrowserViewRef) AsIntfView() uintptr {
	return m.GetIntfPointer(1)
}

// BrowserViewRef  is static instance
var BrowserViewRef _BrowserViewRefClass

// _BrowserViewRefClass is class type defined by TCefBrowserViewRef
type _BrowserViewRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefBrowserView instance using a PCefBrowserView data pointer.
func (_BrowserViewRefClass) UnWrapWithPointer(data uintptr) (result ICefBrowserView) {
	var resultPtr uintptr
	cefBrowserViewRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

// CreateBrowserView
//
//	Create a new BrowserView. The underlying cef_browser_t will not be created
//	until this view is added to the views hierarchy. The optional |extra_info|
//	parameter provides an opportunity to specify extra information specific to
//	the created browser that will be passed to
//	cef_render_process_handler_t::on_browser_created() in the render process.
func (_BrowserViewRefClass) CreateBrowserView(client IEngClient, url string, settings TCefBrowserSettings, extraInfo ICefDictionaryValue, requestContext ICefRequestContext, delegate IEngBrowserViewDelegate) (result ICefBrowserView) {
	settingsPtr := settings.ToPas()
	var resultPtr uintptr
	cefBrowserViewRefAPI().SysCallN(5, base.GetObjectUintptr(client), api.PasStr(url), uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(extraInfo), base.GetObjectUintptr(requestContext), base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

// GetForBrowser
//
//	Returns the BrowserView associated with |browser|.
func (_BrowserViewRefClass) GetForBrowser(browser ICefBrowser) (result ICefBrowserView) {
	var resultPtr uintptr
	cefBrowserViewRefAPI().SysCallN(6, base.GetObjectUintptr(browser), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

// NewBrowserViewRef class constructor
func NewBrowserViewRef(data uintptr) ICefBrowserViewRef {
	var browserViewPtr uintptr // ICefBrowserView
	var viewPtr uintptr        // ICefView
	r := cefBrowserViewRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&browserViewPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefBrowserViewRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewPtr)
		intf.SetIntfPointer(1, viewPtr)
	}
	return ret
}

var (
	cefBrowserViewRefOnce   base.Once
	cefBrowserViewRefImport *imports.Imports = nil
)

func cefBrowserViewRefAPI() *imports.Imports {
	cefBrowserViewRefOnce.Do(func() {
		cefBrowserViewRefImport = api.NewDefaultImports()
		cefBrowserViewRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserViewRef_Create", 0), // constructor NewBrowserViewRef
			/* 1 */ imports.NewTable("TCefBrowserViewRef_GetBrowser", 0), // function GetBrowser
			/* 2 */ imports.NewTable("TCefBrowserViewRef_GetChromeToolbar", 0), // function GetChromeToolbar
			/* 3 */ imports.NewTable("TCefBrowserViewRef_GetRuntimeStyle", 0), // function GetRuntimeStyle
			/* 4 */ imports.NewTable("TCefBrowserViewRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 5 */ imports.NewTable("TCefBrowserViewRef_CreateBrowserView", 0), // static function CreateBrowserView
			/* 6 */ imports.NewTable("TCefBrowserViewRef_GetForBrowser", 0), // static function GetForBrowser
			/* 7 */ imports.NewTable("TCefBrowserViewRef_SetPreferAccelerators", 0), // procedure SetPreferAccelerators
		}
	})
	return cefBrowserViewRefImport
}
