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

// ICefGetExtensionResourceCallback Parent: ICefBaseRefCounted
type ICefGetExtensionResourceCallback interface {
	ICefBaseRefCounted
	Cont(stream ICefStreamReader) // procedure
	Cancel()                      // procedure
}

// ICefGetExtensionResourceCallbackRef Parent: ICefGetExtensionResourceCallback ICefBaseRefCountedRef
type ICefGetExtensionResourceCallbackRef interface {
	ICefGetExtensionResourceCallback
	ICefBaseRefCountedRef
	AsIntfGetExtensionResourceCallback() uintptr
}

type TCefGetExtensionResourceCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefGetExtensionResourceCallbackRef) Cont(stream ICefStreamReader) {
	if !m.IsValid() {
		return
	}
	cefGetExtensionResourceCallbackRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCefGetExtensionResourceCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefGetExtensionResourceCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefGetExtensionResourceCallbackRef) AsIntfGetExtensionResourceCallback() uintptr {
	return m.GetIntfPointer(0)
}

// GetExtensionResourceCallbackRef  is static instance
var GetExtensionResourceCallbackRef _GetExtensionResourceCallbackRefClass

// _GetExtensionResourceCallbackRefClass is class type defined by TCefGetExtensionResourceCallbackRef
type _GetExtensionResourceCallbackRefClass uintptr

func (_GetExtensionResourceCallbackRefClass) UnWrap(data uintptr) (result ICefGetExtensionResourceCallback) {
	var resultPtr uintptr
	cefGetExtensionResourceCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefGetExtensionResourceCallbackRef(resultPtr)
	return
}

// NewGetExtensionResourceCallbackRef class constructor
func NewGetExtensionResourceCallbackRef(data uintptr) ICefGetExtensionResourceCallbackRef {
	var getExtensionResourceCallbackPtr uintptr // ICefGetExtensionResourceCallback
	r := cefGetExtensionResourceCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&getExtensionResourceCallbackPtr)))
	ret := AsCefGetExtensionResourceCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, getExtensionResourceCallbackPtr)
	}
	return ret
}

var (
	cefGetExtensionResourceCallbackRefOnce   base.Once
	cefGetExtensionResourceCallbackRefImport *imports.Imports = nil
)

func cefGetExtensionResourceCallbackRefAPI() *imports.Imports {
	cefGetExtensionResourceCallbackRefOnce.Do(func() {
		cefGetExtensionResourceCallbackRefImport = api.NewDefaultImports()
		cefGetExtensionResourceCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefGetExtensionResourceCallbackRef_Create", 0), // constructor NewGetExtensionResourceCallbackRef
			/* 1 */ imports.NewTable("TCefGetExtensionResourceCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefGetExtensionResourceCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefGetExtensionResourceCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefGetExtensionResourceCallbackRefImport
}
