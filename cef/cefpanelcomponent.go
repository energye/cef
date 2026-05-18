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

// ICefPanelDelegateEvents Parent: ICefViewDelegateEvents
type ICefPanelDelegateEvents interface {
	ICefViewDelegateEvents
}

// ICEFPanelComponent Parent: ICefPanelDelegateEvents ICEFViewComponent
type ICEFPanelComponent interface {
	ICefPanelDelegateEvents
	ICEFViewComponent
	// SetToFillLayout
	//  Set this Panel's Layout to FillLayout and return the FillLayout object.
	SetToFillLayout() ICefFillLayout // function
	// SetToBoxLayout
	//  Set this Panel's Layout to BoxLayout and return the BoxLayout object.
	SetToBoxLayout(settings TCefBoxLayoutSettings) ICefBoxLayout // function
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
	//  Lay out the child Views (set their bounds based on sizing heuristics
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
	// AsWindow
	//  Returns this Panel as a Window or NULL if this is not a Window.
	AsWindow() ICefWindow // property AsWindow Getter
	// ChildViewCount
	//  Returns the number of child Views.
	ChildViewCount() cefTypes.NativeUInt // property ChildViewCount Getter
	AsIntfPanelDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFPanelComponent struct {
	TCEFViewComponent
}

func (m *TCEFPanelComponent) SetToFillLayout() (result ICefFillLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFillLayoutRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) SetToBoxLayout(settings TCefBoxLayoutSettings) (result ICefBoxLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&settings)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBoxLayoutRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) GetLayout() (result ICefLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLayoutRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) GetChildViewAt(index int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) CreatePanel() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(5, m.Instance())
}

func (m *TCEFPanelComponent) Layout() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(6, m.Instance())
}

func (m *TCEFPanelComponent) AddChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) AddChildViewAt(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) ReorderChildView(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) RemoveChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) RemoveAllChildViews() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFPanelComponent) AsWindow() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) ChildViewCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cEFPanelComponentAPI().SysCallN(13, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCEFPanelComponent) AsIntfPanelDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFPanelComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}

// NewPanelComponent class constructor
func NewPanelComponent(owner lcl.IComponent) ICEFPanelComponent {
	var panelDelegateEventsPtr uintptr // ICefPanelDelegateEvents
	var viewDelegateEventsPtr uintptr  // ICefViewDelegateEvents
	r := cEFPanelComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&panelDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFPanelComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelDelegateEventsPtr)
		intf.SetIntfPointer(1, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFPanelComponentOnce   base.Once
	cEFPanelComponentImport *imports.Imports = nil
)

func cEFPanelComponentAPI() *imports.Imports {
	cEFPanelComponentOnce.Do(func() {
		cEFPanelComponentImport = api.NewDefaultImports()
		cEFPanelComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFPanelComponent_Create", 0), // constructor NewPanelComponent
			/* 1 */ imports.NewTable("TCEFPanelComponent_SetToFillLayout", 0), // function SetToFillLayout
			/* 2 */ imports.NewTable("TCEFPanelComponent_SetToBoxLayout", 0), // function SetToBoxLayout
			/* 3 */ imports.NewTable("TCEFPanelComponent_GetLayout", 0), // function GetLayout
			/* 4 */ imports.NewTable("TCEFPanelComponent_GetChildViewAt", 0), // function GetChildViewAt
			/* 5 */ imports.NewTable("TCEFPanelComponent_CreatePanel", 0), // procedure CreatePanel
			/* 6 */ imports.NewTable("TCEFPanelComponent_Layout", 0), // procedure Layout
			/* 7 */ imports.NewTable("TCEFPanelComponent_AddChildView", 0), // procedure AddChildView
			/* 8 */ imports.NewTable("TCEFPanelComponent_AddChildViewAt", 0), // procedure AddChildViewAt
			/* 9 */ imports.NewTable("TCEFPanelComponent_ReorderChildView", 0), // procedure ReorderChildView
			/* 10 */ imports.NewTable("TCEFPanelComponent_RemoveChildView", 0), // procedure RemoveChildView
			/* 11 */ imports.NewTable("TCEFPanelComponent_RemoveAllChildViews", 0), // procedure RemoveAllChildViews
			/* 12 */ imports.NewTable("TCEFPanelComponent_AsWindow", 0), // property AsWindow
			/* 13 */ imports.NewTable("TCEFPanelComponent_ChildViewCount", 0), // property ChildViewCount
		}
	})
	return cEFPanelComponentImport
}
