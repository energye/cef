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
	"github.com/energye/lcl/types"
)

// ICEFWinControl Parent: IWinControl
type ICEFWinControl interface {
	IWinControl
	// TakeSnapshot
	//  Take a snapshot of the browser contents into aBitmap. This function only works in Windows without hardware acceleration.
	TakeSnapshot(bitmap *lcl.IBitmap) bool // function
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
	UpdateSize() // procedure
	// ChildWindowHandle
	//  Handle of the first child window created by the browser.
	ChildWindowHandle() types.THandle  // property ChildWindowHandle Getter
	DragKind() types.TDragKind         // property DragKind Getter
	SetDragKind(value types.TDragKind) // property DragKind Setter
	DragCursor() types.TCursor         // property DragCursor Getter
	SetDragCursor(value types.TCursor) // property DragCursor Setter
	DragMode() types.TDragMode         // property DragMode Getter
	SetDragMode(value types.TDragMode) // property DragMode Setter
	SetOnDragDrop(fn TDragDropEvent)   // property event
	SetOnDragOver(fn TDragOverEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent) // property event
	SetOnEndDrag(fn TEndDragEvent)     // property event
}

type TCEFWinControl struct {
	TWinControl
}

func (m *TCEFWinControl) TakeSnapshot(bitmap *lcl.IBitmap) bool {
	if !m.IsValid() {
		return false
	}
	bitmapPtr := base.GetObjectUintptr(*bitmap)
	r := cEFWinControlAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&bitmapPtr)))
	*bitmap = lcl.AsBitmap(bitmapPtr)
	return api.GoBool(r)
}

func (m *TCEFWinControl) DestroyChildWindow() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWinControlAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWinControl) CreateHandle() {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(4, m.Instance())
}

func (m *TCEFWinControl) InvalidateChildren() {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(5, m.Instance())
}

func (m *TCEFWinControl) UpdateSize() {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(6, m.Instance())
}

func (m *TCEFWinControl) ChildWindowHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := cEFWinControlAPI().SysCallN(7, m.Instance())
	return types.THandle(r)
}

func (m *TCEFWinControl) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := cEFWinControlAPI().SysCallN(8, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCEFWinControl) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWinControl) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := cEFWinControlAPI().SysCallN(9, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCEFWinControl) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWinControl) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := cEFWinControlAPI().SysCallN(10, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCEFWinControl) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	cEFWinControlAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCEFWinControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 11, cEFWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWinControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 12, cEFWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWinControl) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 13, cEFWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWinControl) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, cEFWinControlAPI(), api.MakeEventDataPtr(cb))
}

// NewWinControl class constructor
func NewWinControl(theOwner lcl.IComponent) ICEFWinControl {
	r := cEFWinControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCEFWinControl(r)
}

// NewWinControlParented class constructor
func NewWinControlParented(parentWindow types.HWND) ICEFWinControl {
	r := cEFWinControlAPI().SysCallN(1, uintptr(parentWindow))
	return AsCEFWinControl(r)
}

var (
	cEFWinControlOnce   base.Once
	cEFWinControlImport *imports.Imports = nil
)

func cEFWinControlAPI() *imports.Imports {
	cEFWinControlOnce.Do(func() {
		cEFWinControlImport = api.NewDefaultImports()
		cEFWinControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFWinControl_Create", 0), // constructor NewWinControl
			/* 1 */ imports.NewTable("TCEFWinControl_CreateParented", 0), // constructor NewWinControlParented
			/* 2 */ imports.NewTable("TCEFWinControl_TakeSnapshot", 0), // function TakeSnapshot
			/* 3 */ imports.NewTable("TCEFWinControl_DestroyChildWindow", 0), // function DestroyChildWindow
			/* 4 */ imports.NewTable("TCEFWinControl_CreateHandle", 0), // procedure CreateHandle
			/* 5 */ imports.NewTable("TCEFWinControl_InvalidateChildren", 0), // procedure InvalidateChildren
			/* 6 */ imports.NewTable("TCEFWinControl_UpdateSize", 0), // procedure UpdateSize
			/* 7 */ imports.NewTable("TCEFWinControl_ChildWindowHandle", 0), // property ChildWindowHandle
			/* 8 */ imports.NewTable("TCEFWinControl_DragKind", 0), // property DragKind
			/* 9 */ imports.NewTable("TCEFWinControl_DragCursor", 0), // property DragCursor
			/* 10 */ imports.NewTable("TCEFWinControl_DragMode", 0), // property DragMode
			/* 11 */ imports.NewTable("TCEFWinControl_OnDragDrop", 0), // event OnDragDrop
			/* 12 */ imports.NewTable("TCEFWinControl_OnDragOver", 0), // event OnDragOver
			/* 13 */ imports.NewTable("TCEFWinControl_OnStartDrag", 0), // event OnStartDrag
			/* 14 */ imports.NewTable("TCEFWinControl_OnEndDrag", 0), // event OnEndDrag
		}
	})
	return cEFWinControlImport
}
