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
	SetToFillLayout() ICefFillLayout                             // function
	SetToBoxLayout(settings TCefBoxLayoutSettings) ICefBoxLayout // function
	GetLayout() ICefLayout                                       // function
	GetChildViewCount() cefTypes.NativeUInt                      // function
	GetChildViewAt(index int32) ICefView                         // function
	CreatePanel()                                                // procedure
	Layout()                                                     // procedure
	AddChildView(view ICefView)                                  // procedure
	AddChildViewAt(view ICefView, index int32)                   // procedure
	ReorderChildView(view ICefView, index int32)                 // procedure
	RemoveChildView(view ICefView)                               // procedure
	RemoveAllChildViews()                                        // procedure
	AsWindow() ICefWindow                                        // property AsWindow Getter
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

func (m *TCEFPanelComponent) GetChildViewCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cEFPanelComponentAPI().SysCallN(4, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCEFPanelComponent) GetChildViewAt(index int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFPanelComponent) CreatePanel() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(6, m.Instance())
}

func (m *TCEFPanelComponent) Layout() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(7, m.Instance())
}

func (m *TCEFPanelComponent) AddChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) AddChildViewAt(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) ReorderChildView(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCEFPanelComponent) RemoveChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCEFPanelComponent) RemoveAllChildViews() {
	if !m.IsValid() {
		return
	}
	cEFPanelComponentAPI().SysCallN(12, m.Instance())
}

func (m *TCEFPanelComponent) AsWindow() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFPanelComponentAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
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
			/* 4 */ imports.NewTable("TCEFPanelComponent_GetChildViewCount", 0), // function GetChildViewCount
			/* 5 */ imports.NewTable("TCEFPanelComponent_GetChildViewAt", 0), // function GetChildViewAt
			/* 6 */ imports.NewTable("TCEFPanelComponent_CreatePanel", 0), // procedure CreatePanel
			/* 7 */ imports.NewTable("TCEFPanelComponent_Layout", 0), // procedure Layout
			/* 8 */ imports.NewTable("TCEFPanelComponent_AddChildView", 0), // procedure AddChildView
			/* 9 */ imports.NewTable("TCEFPanelComponent_AddChildViewAt", 0), // procedure AddChildViewAt
			/* 10 */ imports.NewTable("TCEFPanelComponent_ReorderChildView", 0), // procedure ReorderChildView
			/* 11 */ imports.NewTable("TCEFPanelComponent_RemoveChildView", 0), // procedure RemoveChildView
			/* 12 */ imports.NewTable("TCEFPanelComponent_RemoveAllChildViews", 0), // procedure RemoveAllChildViews
			/* 13 */ imports.NewTable("TCEFPanelComponent_AsWindow", 0), // property AsWindow
		}
	})
	return cEFPanelComponentImport
}
