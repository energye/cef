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

// ICEFPanelComponent Parent: ICEFViewComponent
type ICEFPanelComponent interface {
	ICEFViewComponent
	// AsWindow
	//  Returns this Panel as a Window or NULL if this is not a Window.
	AsWindow() ICefWindow // property
	// ChildViewCount
	//  Returns the number of child Views.
	ChildViewCount() NativeUInt // property
	// SetToFillLayout
	//  Set this Panel's Layout to FillLayout and return the FillLayout object.
	SetToFillLayout() ICefFillLayout // function
	// SetToBoxLayout
	//  Set this Panel's Layout to BoxLayout and return the BoxLayout object.
	SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout // function
	// GetLayout
	//  Get the Layout.
	GetLayout() ICefLayout // function
	// GetChildViewAt
	//  Returns the child View at the specified |index|.
	GetChildViewAt(index int32) ICefView // function
	// CreatePanel
	//  Create a new Panel.
	CreatePanel() // procedure
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

// TCEFPanelComponent Parent: TCEFViewComponent
type TCEFPanelComponent struct {
	TCEFViewComponent
}

func NewCEFPanelComponent(aOwner IComponent) ICEFPanelComponent {
	r1 := panelComponentImportAPI().SysCallN(4, GetObjectUintptr(aOwner))
	return AsCEFPanelComponent(r1)
}

func (m *TCEFPanelComponent) AsWindow() ICefWindow {
	var resultCefWindow uintptr
	panelComponentImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCEFPanelComponent) ChildViewCount() NativeUInt {
	r1 := panelComponentImportAPI().SysCallN(3, m.Instance())
	return NativeUInt(r1)
}

func (m *TCEFPanelComponent) SetToFillLayout() ICefFillLayout {
	var resultCefFillLayout uintptr
	panelComponentImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefFillLayout)))
	return AsCefFillLayout(resultCefFillLayout)
}

func (m *TCEFPanelComponent) SetToBoxLayout(settings *TCefBoxLayoutSettings) ICefBoxLayout {
	var resultCefBoxLayout uintptr
	panelComponentImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(settings)), uintptr(unsafePointer(&resultCefBoxLayout)))
	return AsCefBoxLayout(resultCefBoxLayout)
}

func (m *TCEFPanelComponent) GetLayout() ICefLayout {
	var resultCefLayout uintptr
	panelComponentImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCefLayout)))
	return AsCefLayout(resultCefLayout)
}

func (m *TCEFPanelComponent) GetChildViewAt(index int32) ICefView {
	var resultCefView uintptr
	panelComponentImportAPI().SysCallN(6, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFPanelComponent) CreatePanel() {
	panelComponentImportAPI().SysCallN(5, m.Instance())
}

func (m *TCEFPanelComponent) Layout() {
	panelComponentImportAPI().SysCallN(8, m.Instance())
}

func (m *TCEFPanelComponent) AddChildView(view ICefView) {
	panelComponentImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) AddChildViewAt(view ICefView, index int32) {
	panelComponentImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) ReorderChildView(view ICefView, index int32) {
	panelComponentImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) RemoveChildView(view ICefView) {
	panelComponentImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) RemoveAllChildViews() {
	panelComponentImportAPI().SysCallN(9, m.Instance())
}

var (
	panelComponentImport       *imports.Imports = nil
	panelComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFPanelComponent_AddChildView", 0),
		/*1*/ imports.NewTable("CEFPanelComponent_AddChildViewAt", 0),
		/*2*/ imports.NewTable("CEFPanelComponent_AsWindow", 0),
		/*3*/ imports.NewTable("CEFPanelComponent_ChildViewCount", 0),
		/*4*/ imports.NewTable("CEFPanelComponent_Create", 0),
		/*5*/ imports.NewTable("CEFPanelComponent_CreatePanel", 0),
		/*6*/ imports.NewTable("CEFPanelComponent_GetChildViewAt", 0),
		/*7*/ imports.NewTable("CEFPanelComponent_GetLayout", 0),
		/*8*/ imports.NewTable("CEFPanelComponent_Layout", 0),
		/*9*/ imports.NewTable("CEFPanelComponent_RemoveAllChildViews", 0),
		/*10*/ imports.NewTable("CEFPanelComponent_RemoveChildView", 0),
		/*11*/ imports.NewTable("CEFPanelComponent_ReorderChildView", 0),
		/*12*/ imports.NewTable("CEFPanelComponent_SetToBoxLayout", 0),
		/*13*/ imports.NewTable("CEFPanelComponent_SetToFillLayout", 0),
	}
)

func panelComponentImportAPI() *imports.Imports {
	if panelComponentImport == nil {
		panelComponentImport = NewDefaultImports()
		panelComponentImport.SetImportTable(panelComponentImportTables)
		panelComponentImportTables = nil
	}
	return panelComponentImport
}
