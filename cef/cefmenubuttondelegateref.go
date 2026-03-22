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

// ICefMenuButtonDelegate Parent: ICefButtonDelegate
type ICefMenuButtonDelegate interface {
	ICefButtonDelegate
}

// ICefMenuButtonDelegateRef Parent: ICefMenuButtonDelegate ICefButtonDelegateRef
type ICefMenuButtonDelegateRef interface {
	ICefMenuButtonDelegate
	ICefButtonDelegateRef
	AsIntfMenuButtonDelegate() uintptr
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefMenuButtonDelegateRef struct {
	TCefButtonDelegateRef
}

func (m *TCefMenuButtonDelegateRef) AsIntfMenuButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefMenuButtonDelegateRef) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefMenuButtonDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// MenuButtonDelegateRef  is static instance
var MenuButtonDelegateRef _MenuButtonDelegateRefClass

// _MenuButtonDelegateRefClass is class type defined by TCefMenuButtonDelegateRef
type _MenuButtonDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefMenuButtonDelegate instance using a PCefMenuButtonDelegate data pointer.
func (_MenuButtonDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngMenuButtonDelegate) {
	var resultPtr uintptr
	cefMenuButtonDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngMenuButtonDelegate(resultPtr)
	return
}

// NewMenuButtonDelegateRef class constructor
func NewMenuButtonDelegateRef(data uintptr) ICefMenuButtonDelegateRef {
	var menuButtonDelegatePtr uintptr // ICefMenuButtonDelegate
	var buttonDelegatePtr uintptr     // ICefButtonDelegate
	var viewDelegatePtr uintptr       // ICefViewDelegate
	r := cefMenuButtonDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&menuButtonDelegatePtr)), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefMenuButtonDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, menuButtonDelegatePtr)
		intf.SetIntfPointer(1, buttonDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	cefMenuButtonDelegateRefOnce   base.Once
	cefMenuButtonDelegateRefImport *imports.Imports = nil
)

func cefMenuButtonDelegateRefAPI() *imports.Imports {
	cefMenuButtonDelegateRefOnce.Do(func() {
		cefMenuButtonDelegateRefImport = api.NewDefaultImports()
		cefMenuButtonDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMenuButtonDelegateRef_Create", 0), // constructor NewMenuButtonDelegateRef
			/* 1 */ imports.NewTable("TCefMenuButtonDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefMenuButtonDelegateRefImport
}
