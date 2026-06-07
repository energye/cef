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

// ICefViewDelegate Parent: ICefBaseRefCounted
type ICefViewDelegate interface {
	ICefBaseRefCounted
}

// ICefViewDelegateRef Parent: ICefViewDelegate ICefBaseRefCountedRef
type ICefViewDelegateRef interface {
	ICefViewDelegate
	ICefBaseRefCountedRef
	AsIntfViewDelegate() uintptr
}

type TCefViewDelegateRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefViewDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}

// ViewDelegateRef  is static instance
var ViewDelegateRef _ViewDelegateRefClass

// _ViewDelegateRefClass is class type defined by TCefViewDelegateRef
type _ViewDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefViewDelegate instance using a PCefViewDelegate data pointer.
func (_ViewDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngViewDelegate) {
	var resultPtr uintptr
	cefViewDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngViewDelegate(resultPtr)
	return
}

// NewViewDelegateRef class constructor
func NewViewDelegateRef(data uintptr) ICefViewDelegateRef {
	var viewDelegatePtr uintptr // ICefViewDelegate
	r := cefViewDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefViewDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewDelegatePtr)
	}
	return ret
}

var (
	cefViewDelegateRefOnce   base.Once
	cefViewDelegateRefImport *imports.Imports = nil
)

func cefViewDelegateRefAPI() *imports.Imports {
	cefViewDelegateRefOnce.Do(func() {
		cefViewDelegateRefImport = api.NewDefaultImports()
		cefViewDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefViewDelegateRef_Create", 0), // constructor NewViewDelegateRef
			/* 1 */ imports.NewTable("TCefViewDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefViewDelegateRefImport
}
