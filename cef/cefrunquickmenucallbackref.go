//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefRunQuickMenuCallback Parent: ICefBaseRefCounted
type ICefRunQuickMenuCallback interface {
	ICefBaseRefCounted
	Cont(commandId int32, eventFlags cefTypes.TCefEventFlags) // procedure
	Cancel()                                                  // procedure
}

// ICefRunQuickMenuCallbackRef Parent: ICefRunQuickMenuCallback ICefBaseRefCountedRef
type ICefRunQuickMenuCallbackRef interface {
	ICefRunQuickMenuCallback
	ICefBaseRefCountedRef
	AsIntfRunQuickMenuCallback() uintptr
}

type TCefRunQuickMenuCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefRunQuickMenuCallbackRef) Cont(commandId int32, eventFlags cefTypes.TCefEventFlags) {
	if !m.IsValid() {
		return
	}
	cefRunQuickMenuCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(commandId), uintptr(eventFlags))
}

func (m *TCefRunQuickMenuCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefRunQuickMenuCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefRunQuickMenuCallbackRef) AsIntfRunQuickMenuCallback() uintptr {
	return m.GetIntfPointer(0)
}

// RunQuickMenuCallbackRef  is static instance
var RunQuickMenuCallbackRef _RunQuickMenuCallbackRefClass

// _RunQuickMenuCallbackRefClass is class type defined by TCefRunQuickMenuCallbackRef
type _RunQuickMenuCallbackRefClass uintptr

func (_RunQuickMenuCallbackRefClass) UnWrap(data uintptr) (result ICefRunQuickMenuCallback) {
	var resultPtr uintptr
	cefRunQuickMenuCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRunQuickMenuCallbackRef(resultPtr)
	return
}

// NewRunQuickMenuCallbackRef class constructor
func NewRunQuickMenuCallbackRef(data uintptr) ICefRunQuickMenuCallbackRef {
	var runQuickMenuCallbackPtr uintptr // ICefRunQuickMenuCallback
	r := cefRunQuickMenuCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&runQuickMenuCallbackPtr)))
	ret := AsCefRunQuickMenuCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, runQuickMenuCallbackPtr)
	}
	return ret
}

var (
	cefRunQuickMenuCallbackRefOnce   base.Once
	cefRunQuickMenuCallbackRefImport *imports.Imports = nil
)

func cefRunQuickMenuCallbackRefAPI() *imports.Imports {
	cefRunQuickMenuCallbackRefOnce.Do(func() {
		cefRunQuickMenuCallbackRefImport = api.NewDefaultImports()
		cefRunQuickMenuCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefRunQuickMenuCallbackRef_Create", 0), // constructor NewRunQuickMenuCallbackRef
			/* 1 */ imports.NewTable("TCefRunQuickMenuCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefRunQuickMenuCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefRunQuickMenuCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefRunQuickMenuCallbackRefImport
}
