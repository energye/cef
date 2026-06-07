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

// ICefResourceSkipCallback Parent: ICefBaseRefCounted
type ICefResourceSkipCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Callback for asynchronous continuation of skip(). If |bytes_skipped| > 0
	//  then either skip() will be called again until the requested number of
	//  bytes have been skipped or the request will proceed. If |bytes_skipped| <=
	//  0 the request will fail with ERR_REQUEST_RANGE_NOT_SATISFIABLE.
	Cont(bytesSkipped int64) // procedure
}

// ICefResourceSkipCallbackRef Parent: ICefResourceSkipCallback ICefBaseRefCountedRef
type ICefResourceSkipCallbackRef interface {
	ICefResourceSkipCallback
	ICefBaseRefCountedRef
	AsIntfResourceSkipCallback() uintptr
}

type TCefResourceSkipCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResourceSkipCallbackRef) Cont(bytesSkipped int64) {
	if !m.IsValid() {
		return
	}
	cefResourceSkipCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&bytesSkipped)))
}

func (m *TCefResourceSkipCallbackRef) AsIntfResourceSkipCallback() uintptr {
	return m.GetIntfPointer(0)
}

// ResourceSkipCallbackRef  is static instance
var ResourceSkipCallbackRef _ResourceSkipCallbackRefClass

// _ResourceSkipCallbackRefClass is class type defined by TCefResourceSkipCallbackRef
type _ResourceSkipCallbackRefClass uintptr

func (_ResourceSkipCallbackRefClass) UnWrap(data uintptr) (result ICefResourceSkipCallback) {
	var resultPtr uintptr
	cefResourceSkipCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResourceSkipCallbackRef(resultPtr)
	return
}

// NewResourceSkipCallbackRef class constructor
func NewResourceSkipCallbackRef(data uintptr) ICefResourceSkipCallbackRef {
	var resourceSkipCallbackPtr uintptr // ICefResourceSkipCallback
	r := cefResourceSkipCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&resourceSkipCallbackPtr)))
	ret := AsCefResourceSkipCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceSkipCallbackPtr)
	}
	return ret
}

var (
	cefResourceSkipCallbackRefOnce   base.Once
	cefResourceSkipCallbackRefImport *imports.Imports = nil
)

func cefResourceSkipCallbackRefAPI() *imports.Imports {
	cefResourceSkipCallbackRefOnce.Do(func() {
		cefResourceSkipCallbackRefImport = api.NewDefaultImports()
		cefResourceSkipCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResourceSkipCallbackRef_Create", 0), // constructor NewResourceSkipCallbackRef
			/* 1 */ imports.NewTable("TCefResourceSkipCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefResourceSkipCallbackRef_Cont", 0), // procedure Cont
		}
	})
	return cefResourceSkipCallbackRefImport
}
