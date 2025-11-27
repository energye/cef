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

// ICefPrintDialogCallback Parent: ICefBaseRefCountedRef
type ICefPrintDialogCallback interface {
	ICefBaseRefCountedRef
	// Cont
	//  Continue printing with the specified |settings|.
	Cont(settings ICefPrintSettings) // procedure
	// Cancel
	//  Cancel the printing.
	Cancel() // procedure
}

// ICefPrintDialogCallbackRef Parent: ICefPrintDialogCallback
type ICefPrintDialogCallbackRef interface {
	ICefPrintDialogCallback
	AsIntfPrintDialogCallback() uintptr
}

type TCefPrintDialogCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPrintDialogCallbackRef) Cont(settings ICefPrintSettings) {
	if !m.IsValid() {
		return
	}
	cefPrintDialogCallbackRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(settings))
}

func (m *TCefPrintDialogCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefPrintDialogCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefPrintDialogCallbackRef) AsIntfPrintDialogCallback() uintptr {
	return m.GetIntfPointer(0)
}

// PrintDialogCallbackRef  is static instance
var PrintDialogCallbackRef _PrintDialogCallbackRefClass

// _PrintDialogCallbackRefClass is class type defined by TCefPrintDialogCallbackRef
type _PrintDialogCallbackRefClass uintptr

func (_PrintDialogCallbackRefClass) UnWrap(data uintptr) (result ICefPrintDialogCallback) {
	var resultPtr uintptr
	cefPrintDialogCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPrintDialogCallbackRef(resultPtr)
	return
}

// NewPrintDialogCallbackRef class constructor
func NewPrintDialogCallbackRef(data uintptr) ICefPrintDialogCallbackRef {
	var printDialogCallbackPtr uintptr // ICefPrintDialogCallback
	r := cefPrintDialogCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&printDialogCallbackPtr)))
	ret := AsCefPrintDialogCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printDialogCallbackPtr)
	}
	return ret
}

var (
	cefPrintDialogCallbackRefOnce   base.Once
	cefPrintDialogCallbackRefImport *imports.Imports = nil
)

func cefPrintDialogCallbackRefAPI() *imports.Imports {
	cefPrintDialogCallbackRefOnce.Do(func() {
		cefPrintDialogCallbackRefImport = api.NewDefaultImports()
		cefPrintDialogCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPrintDialogCallbackRef_Create", 0), // constructor NewPrintDialogCallbackRef
			/* 1 */ imports.NewTable("TCefPrintDialogCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefPrintDialogCallbackRef_cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefPrintDialogCallbackRef_cancel", 0), // procedure Cancel
		}
	})
	return cefPrintDialogCallbackRefImport
}
