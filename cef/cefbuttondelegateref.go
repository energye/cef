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

// ICefButtonDelegate Parent: ICefViewDelegate
type ICefButtonDelegate interface {
	ICefViewDelegate
}

// ICefButtonDelegateRef Parent: ICefButtonDelegate ICefViewDelegateRef
type ICefButtonDelegateRef interface {
	ICefButtonDelegate
	ICefViewDelegateRef
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefButtonDelegateRef struct {
	TCefViewDelegateRef
}

func (m *TCefButtonDelegateRef) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefButtonDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// ButtonDelegateRef  is static instance
var ButtonDelegateRef _ButtonDelegateRefClass

// _ButtonDelegateRefClass is class type defined by TCefButtonDelegateRef
type _ButtonDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefButtonDelegate instance using a PCefButtonDelegate data pointer.
func (_ButtonDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngButtonDelegate) {
	var resultPtr uintptr
	cefButtonDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngButtonDelegate(resultPtr)
	return
}

// NewButtonDelegateRef class constructor
func NewButtonDelegateRef(data uintptr) ICefButtonDelegateRef {
	var buttonDelegatePtr uintptr // ICefButtonDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := cefButtonDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefButtonDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefButtonDelegateRefOnce   base.Once
	cefButtonDelegateRefImport *imports.Imports = nil
)

func cefButtonDelegateRefAPI() *imports.Imports {
	cefButtonDelegateRefOnce.Do(func() {
		cefButtonDelegateRefImport = api.NewDefaultImports()
		cefButtonDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefButtonDelegateRef_Create", 0), // constructor NewButtonDelegateRef
			/* 1 */ imports.NewTable("TCefButtonDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefButtonDelegateRefImport
}
