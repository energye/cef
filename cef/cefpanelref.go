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

// ICefPanel Parent: ICefView
type ICefPanel interface {
	ICefView
	GetAsWindow() ICefWindow                                     // function
	SetToFillLayout() ICefFillLayout                             // function
	SetToBoxLayout(settings TCefBoxLayoutSettings) ICefBoxLayout // function
	GetLayout() ICefLayout                                       // function
	GetChildViewCount() cefTypes.NativeUInt                      // function
	GetChildViewAt(index int32) ICefView                         // function
	Layout()                                                     // procedure
	AddChildView(view ICefView)                                  // procedure
	AddChildViewAt(view ICefView, index int32)                   // procedure
	ReorderChildView(view ICefView, index int32)                 // procedure
	RemoveChildView(view ICefView)                               // procedure
	RemoveAllChildViews()                                        // procedure
}

// ICefPanelRef Parent: ICefPanel ICefViewRef
type ICefPanelRef interface {
	ICefPanel
	ICefViewRef
	AsIntfPanel() uintptr
	AsIntfView() uintptr
}

type TCefPanelRef struct {
	TCefViewRef
}

func (m *TCefPanelRef) GetAsWindow() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCefPanelRef) SetToFillLayout() (result ICefFillLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFillLayoutRef(resultPtr)
	return
}

func (m *TCefPanelRef) SetToBoxLayout(settings TCefBoxLayoutSettings) (result ICefBoxLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&settings)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBoxLayoutRef(resultPtr)
	return
}

func (m *TCefPanelRef) GetLayout() (result ICefLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLayoutRef(resultPtr)
	return
}

func (m *TCefPanelRef) GetChildViewCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefPanelRefAPI().SysCallN(5, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefPanelRef) GetChildViewAt(index int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(6, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefPanelRef) Layout() {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(9, m.Instance())
}

func (m *TCefPanelRef) AddChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCefPanelRef) AddChildViewAt(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCefPanelRef) ReorderChildView(view ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(view), uintptr(index))
}

func (m *TCefPanelRef) RemoveChildView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCefPanelRef) RemoveAllChildViews() {
	if !m.IsValid() {
		return
	}
	cefPanelRefAPI().SysCallN(14, m.Instance())
}

func (m *TCefPanelRef) AsIntfPanel() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefPanelRef) AsIntfView() uintptr {
	return m.GetIntfPointer(1)
}

// PanelRef  is static instance
var PanelRef _PanelRefClass

// _PanelRefClass is class type defined by TCefPanelRef
type _PanelRefClass uintptr

func (_PanelRefClass) UnWrapWithPointer(data uintptr) (result ICefPanel) {
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(7, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPanelRef(resultPtr)
	return
}

func (_PanelRefClass) CreatePanel(delegate IEngPanelDelegate) (result ICefPanel) {
	var resultPtr uintptr
	cefPanelRefAPI().SysCallN(8, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPanelRef(resultPtr)
	return
}

// NewPanelRef class constructor
func NewPanelRef(data uintptr) ICefPanelRef {
	var panelPtr uintptr // ICefPanel
	var viewPtr uintptr  // ICefView
	r := cefPanelRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&panelPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefPanelRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelPtr)
		intf.SetIntfPointer(1, viewPtr)
	}
	return ret
}

var (
	cefPanelRefOnce   base.Once
	cefPanelRefImport *imports.Imports = nil
)

func cefPanelRefAPI() *imports.Imports {
	cefPanelRefOnce.Do(func() {
		cefPanelRefImport = api.NewDefaultImports()
		cefPanelRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPanelRef_Create", 0), // constructor NewPanelRef
			/* 1 */ imports.NewTable("TCefPanelRef_GetAsWindow", 0), // function GetAsWindow
			/* 2 */ imports.NewTable("TCefPanelRef_SetToFillLayout", 0), // function SetToFillLayout
			/* 3 */ imports.NewTable("TCefPanelRef_SetToBoxLayout", 0), // function SetToBoxLayout
			/* 4 */ imports.NewTable("TCefPanelRef_GetLayout", 0), // function GetLayout
			/* 5 */ imports.NewTable("TCefPanelRef_GetChildViewCount", 0), // function GetChildViewCount
			/* 6 */ imports.NewTable("TCefPanelRef_GetChildViewAt", 0), // function GetChildViewAt
			/* 7 */ imports.NewTable("TCefPanelRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 8 */ imports.NewTable("TCefPanelRef_CreatePanel", 0), // static function CreatePanel
			/* 9 */ imports.NewTable("TCefPanelRef_Layout", 0), // procedure Layout
			/* 10 */ imports.NewTable("TCefPanelRef_AddChildView", 0), // procedure AddChildView
			/* 11 */ imports.NewTable("TCefPanelRef_AddChildViewAt", 0), // procedure AddChildViewAt
			/* 12 */ imports.NewTable("TCefPanelRef_ReorderChildView", 0), // procedure ReorderChildView
			/* 13 */ imports.NewTable("TCefPanelRef_RemoveChildView", 0), // procedure RemoveChildView
			/* 14 */ imports.NewTable("TCefPanelRef_RemoveAllChildViews", 0), // procedure RemoveAllChildViews
		}
	})
	return cefPanelRefImport
}
