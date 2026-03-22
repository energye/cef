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

// ICefUnresponsiveProcessCallback Parent: ICefBaseRefCounted
type ICefUnresponsiveProcessCallback interface {
	ICefBaseRefCounted
	// Wait
	//  Reset the timeout for the unresponsive process.
	Wait() // procedure
	// Terminate
	//  Terminate the unresponsive process.
	Terminate() // procedure
}

// ICefUnresponsiveProcessCallbackRef Parent: ICefUnresponsiveProcessCallback ICefBaseRefCountedRef
type ICefUnresponsiveProcessCallbackRef interface {
	ICefUnresponsiveProcessCallback
	ICefBaseRefCountedRef
	AsIntfUnresponsiveProcessCallback() uintptr
}

type TCefUnresponsiveProcessCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefUnresponsiveProcessCallbackRef) Wait() {
	if !m.IsValid() {
		return
	}
	cefUnresponsiveProcessCallbackRefAPI().SysCallN(2, m.Instance())
}

func (m *TCefUnresponsiveProcessCallbackRef) Terminate() {
	if !m.IsValid() {
		return
	}
	cefUnresponsiveProcessCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefUnresponsiveProcessCallbackRef) AsIntfUnresponsiveProcessCallback() uintptr {
	return m.GetIntfPointer(0)
}

// UnresponsiveProcessCallbackRef  is static instance
var UnresponsiveProcessCallbackRef _UnresponsiveProcessCallbackRefClass

// _UnresponsiveProcessCallbackRefClass is class type defined by TCefUnresponsiveProcessCallbackRef
type _UnresponsiveProcessCallbackRefClass uintptr

func (_UnresponsiveProcessCallbackRefClass) UnWrap(data uintptr) (result ICefUnresponsiveProcessCallback) {
	var resultPtr uintptr
	cefUnresponsiveProcessCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUnresponsiveProcessCallbackRef(resultPtr)
	return
}

// NewUnresponsiveProcessCallbackRef class constructor
func NewUnresponsiveProcessCallbackRef(data uintptr) ICefUnresponsiveProcessCallbackRef {
	var unresponsiveProcessCallbackPtr uintptr // ICefUnresponsiveProcessCallback
	r := cefUnresponsiveProcessCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&unresponsiveProcessCallbackPtr)))
	ret := AsCefUnresponsiveProcessCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, unresponsiveProcessCallbackPtr)
	}
	return ret
}

var (
	cefUnresponsiveProcessCallbackRefOnce   base.Once
	cefUnresponsiveProcessCallbackRefImport *imports.Imports = nil
)

func cefUnresponsiveProcessCallbackRefAPI() *imports.Imports {
	cefUnresponsiveProcessCallbackRefOnce.Do(func() {
		cefUnresponsiveProcessCallbackRefImport = api.NewDefaultImports()
		cefUnresponsiveProcessCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefUnresponsiveProcessCallbackRef_Create", 0), // constructor NewUnresponsiveProcessCallbackRef
			/* 1 */ imports.NewTable("TCefUnresponsiveProcessCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefUnresponsiveProcessCallbackRef_Wait", 0), // procedure Wait
			/* 3 */ imports.NewTable("TCefUnresponsiveProcessCallbackRef_Terminate", 0), // procedure Terminate
		}
	})
	return cefUnresponsiveProcessCallbackRefImport
}
