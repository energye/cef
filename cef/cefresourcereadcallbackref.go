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

// ICefResourceReadCallback Parent: ICefBaseRefCounted
type ICefResourceReadCallback interface {
	ICefBaseRefCounted
	Cont(bytesRead int64) // procedure
}

// ICefResourceReadCallbackRef Parent: ICefResourceReadCallback ICefBaseRefCountedRef
type ICefResourceReadCallbackRef interface {
	ICefResourceReadCallback
	ICefBaseRefCountedRef
	AsIntfResourceReadCallback() uintptr
}

type TCefResourceReadCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResourceReadCallbackRef) Cont(bytesRead int64) {
	if !m.IsValid() {
		return
	}
	cefResourceReadCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&bytesRead)))
}

func (m *TCefResourceReadCallbackRef) AsIntfResourceReadCallback() uintptr {
	return m.GetIntfPointer(0)
}

// ResourceReadCallbackRef  is static instance
var ResourceReadCallbackRef _ResourceReadCallbackRefClass

// _ResourceReadCallbackRefClass is class type defined by TCefResourceReadCallbackRef
type _ResourceReadCallbackRefClass uintptr

func (_ResourceReadCallbackRefClass) UnWrap(data uintptr) (result ICefResourceReadCallback) {
	var resultPtr uintptr
	cefResourceReadCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResourceReadCallbackRef(resultPtr)
	return
}

// NewResourceReadCallbackRef class constructor
func NewResourceReadCallbackRef(data uintptr) ICefResourceReadCallbackRef {
	var resourceReadCallbackPtr uintptr // ICefResourceReadCallback
	r := cefResourceReadCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&resourceReadCallbackPtr)))
	ret := AsCefResourceReadCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceReadCallbackPtr)
	}
	return ret
}

var (
	cefResourceReadCallbackRefOnce   base.Once
	cefResourceReadCallbackRefImport *imports.Imports = nil
)

func cefResourceReadCallbackRefAPI() *imports.Imports {
	cefResourceReadCallbackRefOnce.Do(func() {
		cefResourceReadCallbackRefImport = api.NewDefaultImports()
		cefResourceReadCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResourceReadCallbackRef_Create", 0), // constructor NewResourceReadCallbackRef
			/* 1 */ imports.NewTable("TCefResourceReadCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefResourceReadCallbackRef_Cont", 0), // procedure Cont
		}
	})
	return cefResourceReadCallbackRefImport
}
