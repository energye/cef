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

// ICefFillLayout Parent: ICefLayoutRef
type ICefFillLayout interface {
	ICefLayoutRef
}

// ICefFillLayoutRef Parent: ICefFillLayout
type ICefFillLayoutRef interface {
	ICefFillLayout
	AsIntfFillLayout() uintptr
	AsIntfLayout() uintptr
}

type TCefFillLayoutRef struct {
	TCefLayoutRef
}

func (m *TCefFillLayoutRef) AsIntfFillLayout() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefFillLayoutRef) AsIntfLayout() uintptr {
	return m.GetIntfPointer(1)
}

// FillLayoutRef  is static instance
var FillLayoutRef _FillLayoutRefClass

// _FillLayoutRefClass is class type defined by TCefFillLayoutRef
type _FillLayoutRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefFillLayout instance using a PCefFillLayout data pointer.
func (_FillLayoutRefClass) UnWrapWithPointer(data uintptr) (result ICefFillLayout) {
	var resultPtr uintptr
	cefFillLayoutRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFillLayoutRef(resultPtr)
	return
}

// NewFillLayoutRef class constructor
func NewFillLayoutRef(data uintptr) ICefFillLayoutRef {
	var fillLayoutPtr uintptr // ICefFillLayout
	var layoutPtr uintptr     // ICefLayout
	r := cefFillLayoutRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&fillLayoutPtr)), uintptr(base.UnsafePointer(&layoutPtr)))
	ret := AsCefFillLayoutRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, fillLayoutPtr)
		intf.SetIntfPointer(1, layoutPtr)
	}
	return ret
}

var (
	cefFillLayoutRefOnce   base.Once
	cefFillLayoutRefImport *imports.Imports = nil
)

func cefFillLayoutRefAPI() *imports.Imports {
	cefFillLayoutRefOnce.Do(func() {
		cefFillLayoutRefImport = api.NewDefaultImports()
		cefFillLayoutRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefFillLayoutRef_Create", 0), // constructor NewFillLayoutRef
			/* 1 */ imports.NewTable("TCefFillLayoutRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefFillLayoutRefImport
}
