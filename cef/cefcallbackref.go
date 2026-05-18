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

// ICefCallback Parent: ICefBaseRefCounted
type ICefCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue processing.
	Cont() // procedure
	// Cancel
	//  Cancel processing.
	Cancel() // procedure
}

// ICefCallbackRef Parent: ICefCallback ICefBaseRefCountedRef
type ICefCallbackRef interface {
	ICefCallback
	ICefBaseRefCountedRef
	AsIntfCallback() uintptr
}

type TCefCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefCallbackRef) Cont() {
	if !m.IsValid() {
		return
	}
	cefCallbackRefAPI().SysCallN(2, m.Instance())
}

func (m *TCefCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefCallbackRef) AsIntfCallback() uintptr {
	return m.GetIntfPointer(0)
}

// CallbackRef  is static instance
var CallbackRef _CallbackRefClass

// _CallbackRefClass is class type defined by TCefCallbackRef
type _CallbackRefClass uintptr

func (_CallbackRefClass) UnWrap(data uintptr) (result ICefCallback) {
	var resultPtr uintptr
	cefCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCallbackRef(resultPtr)
	return
}

// NewCallbackRef class constructor
func NewCallbackRef(data uintptr) ICefCallbackRef {
	var callbackPtr uintptr // ICefCallback
	r := cefCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&callbackPtr)))
	ret := AsCefCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, callbackPtr)
	}
	return ret
}

var (
	cefCallbackRefOnce   base.Once
	cefCallbackRefImport *imports.Imports = nil
)

func cefCallbackRefAPI() *imports.Imports {
	cefCallbackRefOnce.Do(func() {
		cefCallbackRefImport = api.NewDefaultImports()
		cefCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCallbackRef_Create", 0), // constructor NewCallbackRef
			/* 1 */ imports.NewTable("TCefCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefCallbackRefImport
}
