//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefButton Parent: ICefView
type ICefButton interface {
	ICefView
	AsLabelButton() ICefLabelButton          // function
	GetState() cefTypes.TCefButtonState      // function
	SetState(state cefTypes.TCefButtonState) // procedure
	SetInkDropEnabled(enabled bool)          // procedure
	SetTooltipText(tooltipText string)       // procedure
	SetAccessibleName(name string)           // procedure
}

// ICefButtonRef Parent: ICefButton ICefViewRef
type ICefButtonRef interface {
	ICefButton
	ICefViewRef
	AsIntfButton() uintptr
	AsIntfView() uintptr
}

type TCefButtonRef struct {
	TCefViewRef
}

func (m *TCefButtonRef) AsLabelButton() (result ICefLabelButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefButtonRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLabelButtonRef(resultPtr)
	return
}

func (m *TCefButtonRef) GetState() cefTypes.TCefButtonState {
	if !m.IsValid() {
		return 0
	}
	r := cefButtonRefAPI().SysCallN(2, m.Instance())
	return cefTypes.TCefButtonState(r)
}

func (m *TCefButtonRef) SetState(state cefTypes.TCefButtonState) {
	if !m.IsValid() {
		return
	}
	cefButtonRefAPI().SysCallN(4, m.Instance(), uintptr(state))
}

func (m *TCefButtonRef) SetInkDropEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	cefButtonRefAPI().SysCallN(5, m.Instance(), api.PasBool(enabled))
}

func (m *TCefButtonRef) SetTooltipText(tooltipText string) {
	if !m.IsValid() {
		return
	}
	cefButtonRefAPI().SysCallN(6, m.Instance(), api.PasStr(tooltipText))
}

func (m *TCefButtonRef) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	cefButtonRefAPI().SysCallN(7, m.Instance(), api.PasStr(name))
}

func (m *TCefButtonRef) AsIntfButton() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefButtonRef) AsIntfView() uintptr {
	return m.GetIntfPointer(1)
}

// ButtonRef  is static instance
var ButtonRef _ButtonRefClass

// _ButtonRefClass is class type defined by TCefButtonRef
type _ButtonRefClass uintptr

func (_ButtonRefClass) UnWrapWithPointer(data uintptr) (result ICefButton) {
	var resultPtr uintptr
	cefButtonRefAPI().SysCallN(3, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefButtonRef(resultPtr)
	return
}

// NewButtonRef class constructor
func NewButtonRef(data uintptr) ICefButtonRef {
	var buttonPtr uintptr // ICefButton
	var viewPtr uintptr   // ICefView
	r := cefButtonRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&buttonPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefButtonRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonPtr)
		intf.SetIntfPointer(1, viewPtr)
	}
	return ret
}

var (
	cefButtonRefOnce   base.Once
	cefButtonRefImport *imports.Imports = nil
)

func cefButtonRefAPI() *imports.Imports {
	cefButtonRefOnce.Do(func() {
		cefButtonRefImport = api.NewDefaultImports()
		cefButtonRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefButtonRef_Create", 0), // constructor NewButtonRef
			/* 1 */ imports.NewTable("TCefButtonRef_AsLabelButton", 0), // function AsLabelButton
			/* 2 */ imports.NewTable("TCefButtonRef_GetState", 0), // function GetState
			/* 3 */ imports.NewTable("TCefButtonRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 4 */ imports.NewTable("TCefButtonRef_SetState", 0), // procedure SetState
			/* 5 */ imports.NewTable("TCefButtonRef_SetInkDropEnabled", 0), // procedure SetInkDropEnabled
			/* 6 */ imports.NewTable("TCefButtonRef_SetTooltipText", 0), // procedure SetTooltipText
			/* 7 */ imports.NewTable("TCefButtonRef_SetAccessibleName", 0), // procedure SetAccessibleName
		}
	})
	return cefButtonRefImport
}
