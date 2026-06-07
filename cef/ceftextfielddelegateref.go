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

// ICefTextfieldDelegate Parent: ICefViewDelegate
type ICefTextfieldDelegate interface {
	ICefViewDelegate
}

// ICefTextfieldDelegateRef Parent: ICefTextfieldDelegate ICefViewDelegateRef
type ICefTextfieldDelegateRef interface {
	ICefTextfieldDelegate
	ICefViewDelegateRef
	AsIntfTextfieldDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefTextfieldDelegateRef struct {
	TCefViewDelegateRef
}

func (m *TCefTextfieldDelegateRef) AsIntfTextfieldDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefTextfieldDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// TextfieldDelegateRef  is static instance
var TextfieldDelegateRef _TextfieldDelegateRefClass

// _TextfieldDelegateRefClass is class type defined by TCefTextfieldDelegateRef
type _TextfieldDelegateRefClass uintptr

func (_TextfieldDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngTextfieldDelegate) {
	var resultPtr uintptr
	cefTextfieldDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngTextfieldDelegate(resultPtr)
	return
}

// NewTextfieldDelegateRef class constructor
func NewTextfieldDelegateRef(data uintptr) ICefTextfieldDelegateRef {
	var textfieldDelegatePtr uintptr // ICefTextfieldDelegate
	var viewDelegatePtr uintptr      // ICefViewDelegate
	r := cefTextfieldDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&textfieldDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefTextfieldDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefTextfieldDelegateRefOnce   base.Once
	cefTextfieldDelegateRefImport *imports.Imports = nil
)

func cefTextfieldDelegateRefAPI() *imports.Imports {
	cefTextfieldDelegateRefOnce.Do(func() {
		cefTextfieldDelegateRefImport = api.NewDefaultImports()
		cefTextfieldDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTextfieldDelegateRef_Create", 0), // constructor NewTextfieldDelegateRef
			/* 1 */ imports.NewTable("TCefTextfieldDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefTextfieldDelegateRefImport
}
