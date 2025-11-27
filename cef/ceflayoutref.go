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

// ICefLayout Parent: ICefBaseRefCountedRef
type ICefLayout interface {
	ICefBaseRefCountedRef
	// AsBoxLayout
	//  Returns this Layout as a BoxLayout or NULL if this is not a BoxLayout.
	AsBoxLayout() ICefBoxLayout // function
	// AsFillLayout
	//  Returns this Layout as a FillLayout or NULL if this is not a FillLayout.
	AsFillLayout() ICefFillLayout // function
	// IsValid
	//  Returns true (1) if this Layout is valid.
	IsValid() bool // function
}

// ICefLayoutRef Parent: ICefLayout
type ICefLayoutRef interface {
	ICefLayout
	AsIntfLayout() uintptr
}

type TCefLayoutRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefLayoutRef) AsBoxLayout() (result ICefBoxLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefLayoutRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBoxLayoutRef(resultPtr)
	return
}

func (m *TCefLayoutRef) AsFillLayout() (result ICefFillLayout) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefLayoutRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFillLayoutRef(resultPtr)
	return
}

func (m *TCefLayoutRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefLayoutRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefLayoutRef) AsIntfLayout() uintptr {
	return m.GetIntfPointer(0)
}

// LayoutRef  is static instance
var LayoutRef _LayoutRefClass

// _LayoutRefClass is class type defined by TCefLayoutRef
type _LayoutRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefLayout instance using a PCefLayout data pointer.
func (_LayoutRefClass) UnWrapWithPointer(data uintptr) (result ICefLayout) {
	var resultPtr uintptr
	cefLayoutRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLayoutRef(resultPtr)
	return
}

// NewLayoutRef class constructor
func NewLayoutRef(data uintptr) ICefLayoutRef {
	var layoutPtr uintptr // ICefLayout
	r := cefLayoutRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&layoutPtr)))
	ret := AsCefLayoutRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, layoutPtr)
	}
	return ret
}

var (
	cefLayoutRefOnce   base.Once
	cefLayoutRefImport *imports.Imports = nil
)

func cefLayoutRefAPI() *imports.Imports {
	cefLayoutRefOnce.Do(func() {
		cefLayoutRefImport = api.NewDefaultImports()
		cefLayoutRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefLayoutRef_Create", 0), // constructor NewLayoutRef
			/* 1 */ imports.NewTable("TCefLayoutRef_AsBoxLayout", 0), // function AsBoxLayout
			/* 2 */ imports.NewTable("TCefLayoutRef_AsFillLayout", 0), // function AsFillLayout
			/* 3 */ imports.NewTable("TCefLayoutRef_IsValid", 0), // function IsValid
			/* 4 */ imports.NewTable("TCefLayoutRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefLayoutRefImport
}
