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
	"github.com/energye/lcl/lcl"
)

// ICefFileDialogCallback Parent: ICefBaseRefCounted
type ICefFileDialogCallback interface {
	ICefBaseRefCounted
	Cont(filePaths lcl.IStrings) // procedure
	Cancel()                     // procedure
}

// ICefFileDialogCallbackRef Parent: ICefFileDialogCallback ICefBaseRefCountedRef
type ICefFileDialogCallbackRef interface {
	ICefFileDialogCallback
	ICefBaseRefCountedRef
	AsIntfFileDialogCallback() uintptr
}

type TCefFileDialogCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefFileDialogCallbackRef) Cont(filePaths lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefFileDialogCallbackRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(filePaths))
}

func (m *TCefFileDialogCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefFileDialogCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefFileDialogCallbackRef) AsIntfFileDialogCallback() uintptr {
	return m.GetIntfPointer(0)
}

// FileDialogCallbackRef  is static instance
var FileDialogCallbackRef _FileDialogCallbackRefClass

// _FileDialogCallbackRefClass is class type defined by TCefFileDialogCallbackRef
type _FileDialogCallbackRefClass uintptr

func (_FileDialogCallbackRefClass) UnWrap(data uintptr) (result ICefFileDialogCallback) {
	var resultPtr uintptr
	cefFileDialogCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFileDialogCallbackRef(resultPtr)
	return
}

// NewFileDialogCallbackRef class constructor
func NewFileDialogCallbackRef(data uintptr) ICefFileDialogCallbackRef {
	var fileDialogCallbackPtr uintptr // ICefFileDialogCallback
	r := cefFileDialogCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&fileDialogCallbackPtr)))
	ret := AsCefFileDialogCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, fileDialogCallbackPtr)
	}
	return ret
}

var (
	cefFileDialogCallbackRefOnce   base.Once
	cefFileDialogCallbackRefImport *imports.Imports = nil
)

func cefFileDialogCallbackRefAPI() *imports.Imports {
	cefFileDialogCallbackRefOnce.Do(func() {
		cefFileDialogCallbackRefImport = api.NewDefaultImports()
		cefFileDialogCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefFileDialogCallbackRef_Create", 0), // constructor NewFileDialogCallbackRef
			/* 1 */ imports.NewTable("TCefFileDialogCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefFileDialogCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefFileDialogCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefFileDialogCallbackRefImport
}
