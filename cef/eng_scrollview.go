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

// ICefScrollView Parent: ICefView
//
//	A ScrollView will show horizontal and/or vertical scrollbars when necessary
//	based on the size of the attached content view. Methods must be called on
//	the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_scroll_view_capi.h">CEF source file: /include/capi/views/cef_scroll_view_capi.h (cef_scroll_view_t)</a>
type ICefScrollView interface {
	ICefView
	// GetContentView
	//  Returns the content View.
	GetContentView() ICefView // function
	// GetVisibleContentRect
	//  Returns the visible region of the content View.
	GetVisibleContentRect() (resultCefRect TCefRect) // function
	// HasHorizontalScrollbar
	//  Returns true(1) if the horizontal scrollbar is currently showing.
	HasHorizontalScrollbar() bool // function
	// GetHorizontalScrollbarHeight
	//  Returns the height of the horizontal scrollbar.
	GetHorizontalScrollbarHeight() int32 // function
	// HasVerticalScrollbar
	//  Returns true(1) if the vertical scrollbar is currently showing.
	HasVerticalScrollbar() bool // function
	// GetVerticalScrollbarWidth
	//  Returns the width of the vertical scrollbar.
	GetVerticalScrollbarWidth() int32 // function
	// SetContentView
	//  Set the content View. The content View must have a specified size(e.g.
	//  via ICefView.SetBounds or ICefViewDelegate.GetPreferredSize).
	SetContentView(view ICefView) // procedure
}

// TCefScrollView Parent: TCefView
//
//	A ScrollView will show horizontal and/or vertical scrollbars when necessary
//	based on the size of the attached content view. Methods must be called on
//	the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_scroll_view_capi.h">CEF source file: /include/capi/views/cef_scroll_view_capi.h (cef_scroll_view_t)</a>
type TCefScrollView struct {
	TCefView
}

// ScrollViewRef -> ICefScrollView
var ScrollViewRef scrollView

// scrollView TCefScrollView Ref
type scrollView uintptr

// UnWrap
//
//	Returns a ICefScrollView instance using a PCefScrollView data pointer.
func (m *scrollView) UnWrap(data uintptr) ICefScrollView {
	var resultCefScrollView uintptr
	scrollViewImportAPI().SysCallN(8, uintptr(data), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

// CreateScrollView
//
//	Create a new ScrollView.
func (m *scrollView) CreateScrollView(delegate ICefViewDelegate) ICefScrollView {
	var resultCefScrollView uintptr
	scrollViewImportAPI().SysCallN(0, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

func (m *TCefScrollView) GetContentView() ICefView {
	var resultCefView uintptr
	scrollViewImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefScrollView) GetVisibleContentRect() (resultCefRect TCefRect) {
	scrollViewImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefScrollView) HasHorizontalScrollbar() bool {
	r1 := scrollViewImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCefScrollView) GetHorizontalScrollbarHeight() int32 {
	r1 := scrollViewImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefScrollView) HasVerticalScrollbar() bool {
	r1 := scrollViewImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCefScrollView) GetVerticalScrollbarWidth() int32 {
	r1 := scrollViewImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TCefScrollView) SetContentView(view ICefView) {
	scrollViewImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(view))
}

var (
	scrollViewImport       *imports.Imports = nil
	scrollViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefScrollView_CreateScrollView", 0),
		/*1*/ imports.NewTable("CefScrollView_GetContentView", 0),
		/*2*/ imports.NewTable("CefScrollView_GetHorizontalScrollbarHeight", 0),
		/*3*/ imports.NewTable("CefScrollView_GetVerticalScrollbarWidth", 0),
		/*4*/ imports.NewTable("CefScrollView_GetVisibleContentRect", 0),
		/*5*/ imports.NewTable("CefScrollView_HasHorizontalScrollbar", 0),
		/*6*/ imports.NewTable("CefScrollView_HasVerticalScrollbar", 0),
		/*7*/ imports.NewTable("CefScrollView_SetContentView", 0),
		/*8*/ imports.NewTable("CefScrollView_UnWrap", 0),
	}
)

func scrollViewImportAPI() *imports.Imports {
	if scrollViewImport == nil {
		scrollViewImport = NewDefaultImports()
		scrollViewImport.SetImportTable(scrollViewImportTables)
		scrollViewImportTables = nil
	}
	return scrollViewImport
}
