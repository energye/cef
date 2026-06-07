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

// ICefWindowDelegate Parent: ICefPanelDelegate
type ICefWindowDelegate interface {
	ICefPanelDelegate
}

// ICefWindowDelegateRef Parent: ICefWindowDelegate ICefPanelDelegateRef
type ICefWindowDelegateRef interface {
	ICefWindowDelegate
	ICefPanelDelegateRef
	AsIntfWindowDelegate() uintptr
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefWindowDelegateRef struct {
	TCefPanelDelegateRef
}

func (m *TCefWindowDelegateRef) AsIntfWindowDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefWindowDelegateRef) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefWindowDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// WindowDelegateRef  is static instance
var WindowDelegateRef _WindowDelegateRefClass

// _WindowDelegateRefClass is class type defined by TCefWindowDelegateRef
type _WindowDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefWindowDelegate instance using a PCefWindowDelegate data pointer.
func (_WindowDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngWindowDelegate) {
	var resultPtr uintptr
	cefWindowDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngWindowDelegate(resultPtr)
	return
}

// NewWindowDelegateRef class constructor
func NewWindowDelegateRef(data uintptr) ICefWindowDelegateRef {
	var windowDelegatePtr uintptr // ICefWindowDelegate
	var panelDelegatePtr uintptr  // ICefPanelDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := cefWindowDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&windowDelegatePtr)), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefWindowDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowDelegatePtr)
		intf.SetIntfPointer(1, panelDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	cefWindowDelegateRefOnce   base.Once
	cefWindowDelegateRefImport *imports.Imports = nil
)

func cefWindowDelegateRefAPI() *imports.Imports {
	cefWindowDelegateRefOnce.Do(func() {
		cefWindowDelegateRefImport = api.NewDefaultImports()
		cefWindowDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefWindowDelegateRef_Create", 0), // constructor NewWindowDelegateRef
			/* 1 */ imports.NewTable("TCefWindowDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefWindowDelegateRefImport
}
