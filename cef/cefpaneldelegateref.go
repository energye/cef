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
)

// ICefPanelDelegate Parent: ICefViewDelegate
type ICefPanelDelegate interface {
	ICefViewDelegate
}

// ICefPanelDelegateRef Parent: ICefPanelDelegate ICefViewDelegateRef
type ICefPanelDelegateRef interface {
	ICefPanelDelegate
	ICefViewDelegateRef
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefPanelDelegateRef struct {
	TCefViewDelegateRef
}

func (m *TCefPanelDelegateRef) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefPanelDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// PanelDelegateRef  is static instance
var PanelDelegateRef _PanelDelegateRefClass

// _PanelDelegateRefClass is class type defined by TCefPanelDelegateRef
type _PanelDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefPanelDelegate instance using a PCefPanelDelegate data pointer.
func (_PanelDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngPanelDelegate) {
	var resultPtr uintptr
	cefPanelDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngPanelDelegate(resultPtr)
	return
}

// NewPanelDelegateRef class constructor
func NewPanelDelegateRef(data uintptr) ICefPanelDelegateRef {
	var panelDelegatePtr uintptr // ICefPanelDelegate
	var viewDelegatePtr uintptr  // ICefViewDelegate
	r := cefPanelDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefPanelDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefPanelDelegateRefOnce   base.Once
	cefPanelDelegateRefImport *imports.Imports = nil
)

func cefPanelDelegateRefAPI() *imports.Imports {
	cefPanelDelegateRefOnce.Do(func() {
		cefPanelDelegateRefImport = api.NewDefaultImports()
		cefPanelDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPanelDelegateRef_Create", 0), // constructor NewPanelDelegateRef
			/* 1 */ imports.NewTable("TCefPanelDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefPanelDelegateRefImport
}
