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

// ICefPrintJobCallback Parent: ICefBaseRefCounted
type ICefPrintJobCallback interface {
	ICefBaseRefCounted
	Cont() // procedure
}

// ICefPrintJobCallbackRef Parent: ICefPrintJobCallback ICefBaseRefCountedRef
type ICefPrintJobCallbackRef interface {
	ICefPrintJobCallback
	ICefBaseRefCountedRef
	AsIntfPrintJobCallback() uintptr
}

type TCefPrintJobCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPrintJobCallbackRef) Cont() {
	if !m.IsValid() {
		return
	}
	cefPrintJobCallbackRefAPI().SysCallN(2, m.Instance())
}

func (m *TCefPrintJobCallbackRef) AsIntfPrintJobCallback() uintptr {
	return m.GetIntfPointer(0)
}

// PrintJobCallbackRef  is static instance
var PrintJobCallbackRef _PrintJobCallbackRefClass

// _PrintJobCallbackRefClass is class type defined by TCefPrintJobCallbackRef
type _PrintJobCallbackRefClass uintptr

func (_PrintJobCallbackRefClass) UnWrap(data uintptr) (result ICefPrintJobCallback) {
	var resultPtr uintptr
	cefPrintJobCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPrintJobCallbackRef(resultPtr)
	return
}

// NewPrintJobCallbackRef class constructor
func NewPrintJobCallbackRef(data uintptr) ICefPrintJobCallbackRef {
	var printJobCallbackPtr uintptr // ICefPrintJobCallback
	r := cefPrintJobCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&printJobCallbackPtr)))
	ret := AsCefPrintJobCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printJobCallbackPtr)
	}
	return ret
}

var (
	cefPrintJobCallbackRefOnce   base.Once
	cefPrintJobCallbackRefImport *imports.Imports = nil
)

func cefPrintJobCallbackRefAPI() *imports.Imports {
	cefPrintJobCallbackRefOnce.Do(func() {
		cefPrintJobCallbackRefImport = api.NewDefaultImports()
		cefPrintJobCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPrintJobCallbackRef_Create", 0), // constructor NewPrintJobCallbackRef
			/* 1 */ imports.NewTable("TCefPrintJobCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefPrintJobCallbackRef_cont", 0), // procedure Cont
		}
	})
	return cefPrintJobCallbackRefImport
}
