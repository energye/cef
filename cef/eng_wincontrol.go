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

// ICEFWinControl Parent: IWinControl
//
//	Custom TWinControl used by CEF browsers.
type ICEFWinControl interface {
	IWinControl
	// ChildWindowHandle
	//  Handle of the first child window created by the browser.
	ChildWindowHandle() THandle   // property
	DragKind() TDragKind          // property
	SetDragKind(AValue TDragKind) // property
	DragCursor() TCursor          // property
	SetDragCursor(AValue TCursor) // property
	DragMode() TDragMode          // property
	SetDragMode(AValue TDragMode) // property
	// TakeSnapshot
	//  Take a snapshot of the browser contents into aBitmap. This function only works in Windows without hardware acceleration.
	TakeSnapshot(aBitmap *IBitmap) bool // function
	// DestroyChildWindow
	//  Destroy the child windows created by the browser.
	DestroyChildWindow() bool // function
	// CreateHandle
	//  Exposes the CreateHandle procedure to create the Handle at any time.
	CreateHandle() // procedure
	// InvalidateChildren
	//  Invalidate the child windows created by the browser.
	InvalidateChildren() // procedure
	// UpdateSize
	//  Updates the size of the child windows created by the browser.
	UpdateSize()                  // procedure
	SetOnDragDrop(fn TDragDrop)   // property event
	SetOnDragOver(fn TDragOver)   // property event
	SetOnStartDrag(fn TStartDrag) // property event
	SetOnEndDrag(fn TEndDrag)     // property event
}

// TCEFWinControl Parent: TWinControl
//
//	Custom TWinControl used by CEF browsers.
type TCEFWinControl struct {
	TWinControl
	dragDropPtr  uintptr
	dragOverPtr  uintptr
	startDragPtr uintptr
	endDragPtr   uintptr
}

func NewCEFWinControl(theOwner IComponent) ICEFWinControl {
	r1 := winControlImportAPI().SysCallN(1, GetObjectUintptr(theOwner))
	return AsCEFWinControl(r1)
}

func (m *TCEFWinControl) ChildWindowHandle() THandle {
	r1 := winControlImportAPI().SysCallN(0, m.Instance())
	return THandle(r1)
}

func (m *TCEFWinControl) DragKind() TDragKind {
	r1 := winControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCEFWinControl) SetDragKind(AValue TDragKind) {
	winControlImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) DragCursor() TCursor {
	r1 := winControlImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCEFWinControl) SetDragCursor(AValue TCursor) {
	winControlImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) DragMode() TDragMode {
	r1 := winControlImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCEFWinControl) SetDragMode(AValue TDragMode) {
	winControlImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWinControl) TakeSnapshot(aBitmap *IBitmap) bool {
	var result0 uintptr
	r1 := winControlImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&result0)))
	*aBitmap = AsBitmap(result0)
	return GoBool(r1)
}

func (m *TCEFWinControl) DestroyChildWindow() bool {
	r1 := winControlImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCEFWinControl) CreateHandle() {
	winControlImportAPI().SysCallN(2, m.Instance())
}

func (m *TCEFWinControl) InvalidateChildren() {
	winControlImportAPI().SysCallN(7, m.Instance())
}

func (m *TCEFWinControl) UpdateSize() {
	winControlImportAPI().SysCallN(14, m.Instance())
}

func (m *TCEFWinControl) SetOnDragDrop(fn TDragDrop) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TCEFWinControl) SetOnDragOver(fn TDragOver) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TCEFWinControl) SetOnStartDrag(fn TStartDrag) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(11, m.Instance(), m.startDragPtr)
}

func (m *TCEFWinControl) SetOnEndDrag(fn TEndDrag) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

var (
	winControlImport       *imports.Imports = nil
	winControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFWinControl_ChildWindowHandle", 0),
		/*1*/ imports.NewTable("CEFWinControl_Create", 0),
		/*2*/ imports.NewTable("CEFWinControl_CreateHandle", 0),
		/*3*/ imports.NewTable("CEFWinControl_DestroyChildWindow", 0),
		/*4*/ imports.NewTable("CEFWinControl_DragCursor", 0),
		/*5*/ imports.NewTable("CEFWinControl_DragKind", 0),
		/*6*/ imports.NewTable("CEFWinControl_DragMode", 0),
		/*7*/ imports.NewTable("CEFWinControl_InvalidateChildren", 0),
		/*8*/ imports.NewTable("CEFWinControl_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("CEFWinControl_SetOnDragOver", 0),
		/*10*/ imports.NewTable("CEFWinControl_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("CEFWinControl_SetOnStartDrag", 0),
		/*12*/ imports.NewTable("CEFWinControl_TakeSnapshot", 0),
		/*13*/ imports.NewTable("CEFWinControl_Touch", 0),
		/*14*/ imports.NewTable("CEFWinControl_UpdateSize", 0),
	}
)

func winControlImportAPI() *imports.Imports {
	if winControlImport == nil {
		winControlImport = NewDefaultImports()
		winControlImport.SetImportTable(winControlImportTables)
		winControlImportTables = nil
	}
	return winControlImport
}
