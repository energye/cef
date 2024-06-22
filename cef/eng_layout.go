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

// ICefLayout Parent: ICefBaseRefCounted
//
//	A Layout handles the sizing of the children of a Panel according to
//	implementation-specific heuristics. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_layout_capi.h">CEF source file: /include/capi/views/cef_layout_capi.h (cef_layout_t)</a>
type ICefLayout interface {
	ICefBaseRefCounted
	// AsBoxLayout
	//  Returns this Layout as a BoxLayout or NULL if this is not a BoxLayout.
	AsBoxLayout() ICefBoxLayout // function
	// AsFillLayout
	//  Returns this Layout as a FillLayout or NULL if this is not a FillLayout.
	AsFillLayout() ICefFillLayout // function
	// IsValid
	//  Returns true(1) if this Layout is valid.
	IsValid() bool // function
}

// TCefLayout Parent: TCefBaseRefCounted
//
//	A Layout handles the sizing of the children of a Panel according to
//	implementation-specific heuristics. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_layout_capi.h">CEF source file: /include/capi/views/cef_layout_capi.h (cef_layout_t)</a>
type TCefLayout struct {
	TCefBaseRefCounted
}

// LayoutRef -> ICefLayout
var LayoutRef layout

// layout TCefLayout Ref
type layout uintptr

// UnWrap
//
//	Returns a ICefLayout instance using a PCefLayout data pointer.
func (m *layout) UnWrap(data uintptr) ICefLayout {
	var resultCefLayout uintptr
	layoutImportAPI().SysCallN(3, uintptr(data), uintptr(unsafePointer(&resultCefLayout)))
	return AsCefLayout(resultCefLayout)
}

func (m *TCefLayout) AsBoxLayout() ICefBoxLayout {
	var resultCefBoxLayout uintptr
	layoutImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCefLayout) AsFillLayout() ICefFillLayout {
	var resultCefFillLayout uintptr
	layoutImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}

func (m *TCefLayout) IsValid() bool {
	r1 := layoutImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

var (
	layoutImport       *imports.Imports = nil
	layoutImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefLayout_AsBoxLayout", 0),
		/*1*/ imports.NewTable("CefLayout_AsFillLayout", 0),
		/*2*/ imports.NewTable("CefLayout_IsValid", 0),
		/*3*/ imports.NewTable("CefLayout_UnWrap", 0),
	}
)

func layoutImportAPI() *imports.Imports {
	if layoutImport == nil {
		layoutImport = NewDefaultImports()
		layoutImport.SetImportTable(layoutImportTables)
		layoutImportTables = nil
	}
	return layoutImport
}
