//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefRunContextMenuCallback Parent: ICefBaseRefCounted
type ICefRunContextMenuCallback interface {
	ICefBaseRefCounted
	Cont(commandId int32, eventFlags cefTypes.TCefEventFlags) // procedure
	Cancel()                                                  // procedure
}

// ICefRunContextMenuCallbackRef Parent: ICefRunContextMenuCallback ICefBaseRefCountedRef
type ICefRunContextMenuCallbackRef interface {
	ICefRunContextMenuCallback
	ICefBaseRefCountedRef
	AsIntfRunContextMenuCallback() uintptr
}

type TCefRunContextMenuCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefRunContextMenuCallbackRef) Cont(commandId int32, eventFlags cefTypes.TCefEventFlags) {
	if !m.IsValid() {
		return
	}
	cefRunContextMenuCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(commandId), uintptr(eventFlags))
}

func (m *TCefRunContextMenuCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefRunContextMenuCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefRunContextMenuCallbackRef) AsIntfRunContextMenuCallback() uintptr {
	return m.GetIntfPointer(0)
}

// RunContextMenuCallbackRef  is static instance
var RunContextMenuCallbackRef _RunContextMenuCallbackRefClass

// _RunContextMenuCallbackRefClass is class type defined by TCefRunContextMenuCallbackRef
type _RunContextMenuCallbackRefClass uintptr

func (_RunContextMenuCallbackRefClass) UnWrap(data uintptr) (result ICefRunContextMenuCallback) {
	var resultPtr uintptr
	cefRunContextMenuCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRunContextMenuCallbackRef(resultPtr)
	return
}

// NewRunContextMenuCallbackRef class constructor
func NewRunContextMenuCallbackRef(data uintptr) ICefRunContextMenuCallbackRef {
	var runContextMenuCallbackPtr uintptr // ICefRunContextMenuCallback
	r := cefRunContextMenuCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&runContextMenuCallbackPtr)))
	ret := AsCefRunContextMenuCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, runContextMenuCallbackPtr)
	}
	return ret
}

var (
	cefRunContextMenuCallbackRefOnce   base.Once
	cefRunContextMenuCallbackRefImport *imports.Imports = nil
)

func cefRunContextMenuCallbackRefAPI() *imports.Imports {
	cefRunContextMenuCallbackRefOnce.Do(func() {
		cefRunContextMenuCallbackRefImport = api.NewDefaultImports()
		cefRunContextMenuCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefRunContextMenuCallbackRef_Create", 0), // constructor NewRunContextMenuCallbackRef
			/* 1 */ imports.NewTable("TCefRunContextMenuCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefRunContextMenuCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefRunContextMenuCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefRunContextMenuCallbackRefImport
}
