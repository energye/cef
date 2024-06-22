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

// ICefPanel Parent: ICefView
//
//	A Panel is a container in the views hierarchy that can contain other Views
//	as children. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_capi.h">CEF source file: /include/capi/views/cef_panel_capi.h (cef_panel_t)</a>
type ICefPanel interface {
	ICefView
	// GetAsWindow
	//  Returns this Panel as a Window or NULL if this is not a Window.
	GetAsWindow() ICefWindow // function
	// SetToFillLayout
	//  Set this Panel's Layout to FillLayout and return the FillLayout object.
	SetToFillLayout() ICefFillLayout // function
	// SetToBoxLayout
	//  Set this Panel's Layout to BoxLayout and return the BoxLayout object.
	SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout // function
	// GetLayout
	//  Get the Layout.
	GetLayout() ICefLayout // function
	// GetChildViewCount
	//  Returns the number of child Views.
	GetChildViewCount() NativeUInt // function
	// GetChildViewAt
	//  Returns the child View at the specified |index|.
	GetChildViewAt(index int32) ICefView // function
	// Layout
	//  Lay out the child Views(set their bounds based on sizing heuristics
	//  specific to the current Layout).
	Layout() // procedure
	// AddChildView
	//  Add a child View.
	AddChildView(view ICefView) // procedure
	// AddChildViewAt
	//  Add a child View at the specified |index|. If |index| matches the result
	//  of GetChildCount() then the View will be added at the end.
	AddChildViewAt(view ICefView, index int32) // procedure
	// ReorderChildView
	//  Move the child View to the specified |index|. A negative value for |index|
	//  will move the View to the end.
	ReorderChildView(view ICefView, index int32) // procedure
	// RemoveChildView
	//  Remove a child View. The View can then be added to another Panel.
	RemoveChildView(view ICefView) // procedure
	// RemoveAllChildViews
	//  Remove all child Views. The removed Views will be deleted if the client
	//  holds no references to them.
	RemoveAllChildViews() // procedure
}

// TCefPanel Parent: TCefView
//
//	A Panel is a container in the views hierarchy that can contain other Views
//	as children. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_capi.h">CEF source file: /include/capi/views/cef_panel_capi.h (cef_panel_t)</a>
type TCefPanel struct {
	TCefView
}

// PanelRef -> ICefPanel
var PanelRef panel

// panel TCefPanel Ref
type panel uintptr

// UnWrap
//
//	Returns a ICefPanel instance using a PCefPanel data pointer.
func (m *panel) UnWrap(data uintptr) ICefPanel {
	var resultCefPanel uintptr
	panelImportAPI().SysCallN(13, uintptr(data), uintptr(unsafePointer(&resultCefPanel)))
	return AsCefPanel(resultCefPanel)
}

// CreatePanel
//
//	Create a new Panel.
func (m *panel) CreatePanel(delegate ICefPanelDelegate) ICefPanel {
	var resultCefPanel uintptr
	panelImportAPI().SysCallN(2, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefPanel)))
	return AsCefPanel(resultCefPanel)
}

func (m *TCefPanel) GetAsWindow() ICefWindow {
	var resultCefWindow uintptr
	panelImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefPanel) SetToFillLayout() ICefFillLayout {
	var resultCefFillLayout uintptr
	panelImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}

func (m *TCefPanel) SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout {
	var resultCefBoxLayout uintptr
	panelImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(settings)), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCefPanel) GetLayout() ICefLayout {
	var resultCefLayout uintptr
	panelImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefLayout)))
	return AsCefLayout(resultCefLayout)
}

func (m *TCefPanel) GetChildViewCount() NativeUInt {
	r1 := panelImportAPI().SysCallN(5, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPanel) GetChildViewAt(index int32) ICefView {
	var resultCefView uintptr
	panelImportAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefPanel) Layout() {
	panelImportAPI().SysCallN(7, m.Instance())
}

func (m *TCefPanel) AddChildView(view ICefView) {
	panelImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(view))
}

func (m *TCefPanel) AddChildViewAt(view ICefView, index int32) {
	panelImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCefPanel) ReorderChildView(view ICefView, index int32) {
	panelImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCefPanel) RemoveChildView(view ICefView) {
	panelImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(view))
}

func (m *TCefPanel) RemoveAllChildViews() {
	panelImportAPI().SysCallN(8, m.Instance())
}

var (
	panelImport       *imports.Imports = nil
	panelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPanel_AddChildView", 0),
		/*1*/ imports.NewTable("CefPanel_AddChildViewAt", 0),
		/*2*/ imports.NewTable("CefPanel_CreatePanel", 0),
		/*3*/ imports.NewTable("CefPanel_GetAsWindow", 0),
		/*4*/ imports.NewTable("CefPanel_GetChildViewAt", 0),
		/*5*/ imports.NewTable("CefPanel_GetChildViewCount", 0),
		/*6*/ imports.NewTable("CefPanel_GetLayout", 0),
		/*7*/ imports.NewTable("CefPanel_Layout", 0),
		/*8*/ imports.NewTable("CefPanel_RemoveAllChildViews", 0),
		/*9*/ imports.NewTable("CefPanel_RemoveChildView", 0),
		/*10*/ imports.NewTable("CefPanel_ReorderChildView", 0),
		/*11*/ imports.NewTable("CefPanel_SetToBoxLayout", 0),
		/*12*/ imports.NewTable("CefPanel_SetToFillLayout", 0),
		/*13*/ imports.NewTable("CefPanel_UnWrap", 0),
	}
)

func panelImportAPI() *imports.Imports {
	if panelImport == nil {
		panelImport = NewDefaultImports()
		panelImport.SetImportTable(panelImportTables)
		panelImportTables = nil
	}
	return panelImport
}
