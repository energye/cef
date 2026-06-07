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

// ICefBoxLayout Parent: ICefLayout
type ICefBoxLayout interface {
	ICefLayout
	SetFlexForView(view ICefView, flex int32) // procedure
	ClearFlexForView(view ICefView)           // procedure
}

// ICefBoxLayoutRef Parent: ICefBoxLayout ICefLayoutRef
type ICefBoxLayoutRef interface {
	ICefBoxLayout
	ICefLayoutRef
	AsIntfBoxLayout() uintptr
	AsIntfLayout() uintptr
}

type TCefBoxLayoutRef struct {
	TCefLayoutRef
}

func (m *TCefBoxLayoutRef) SetFlexForView(view ICefView, flex int32) {
	if !m.IsValid() {
		return
	}
	cefBoxLayoutRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(view), uintptr(flex))
}

func (m *TCefBoxLayoutRef) ClearFlexForView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cefBoxLayoutRefAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCefBoxLayoutRef) AsIntfBoxLayout() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefBoxLayoutRef) AsIntfLayout() uintptr {
	return m.GetIntfPointer(1)
}

// BoxLayoutRef  is static instance
var BoxLayoutRef _BoxLayoutRefClass

// _BoxLayoutRefClass is class type defined by TCefBoxLayoutRef
type _BoxLayoutRefClass uintptr

func (_BoxLayoutRefClass) UnWrapWithPointer(data uintptr) (result ICefBoxLayout) {
	var resultPtr uintptr
	cefBoxLayoutRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBoxLayoutRef(resultPtr)
	return
}

// NewBoxLayoutRef class constructor
func NewBoxLayoutRef(data uintptr) ICefBoxLayoutRef {
	var boxLayoutPtr uintptr // ICefBoxLayout
	var layoutPtr uintptr    // ICefLayout
	r := cefBoxLayoutRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&boxLayoutPtr)), uintptr(base.UnsafePointer(&layoutPtr)))
	ret := AsCefBoxLayoutRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, boxLayoutPtr)
		intf.SetIntfPointer(1, layoutPtr)
	}
	return ret
}

var (
	cefBoxLayoutRefOnce   base.Once
	cefBoxLayoutRefImport *imports.Imports = nil
)

func cefBoxLayoutRefAPI() *imports.Imports {
	cefBoxLayoutRefOnce.Do(func() {
		cefBoxLayoutRefImport = api.NewDefaultImports()
		cefBoxLayoutRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBoxLayoutRef_Create", 0), // constructor NewBoxLayoutRef
			/* 1 */ imports.NewTable("TCefBoxLayoutRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 2 */ imports.NewTable("TCefBoxLayoutRef_SetFlexForView", 0), // procedure SetFlexForView
			/* 3 */ imports.NewTable("TCefBoxLayoutRef_ClearFlexForView", 0), // procedure ClearFlexForView
		}
	})
	return cefBoxLayoutRefImport
}
