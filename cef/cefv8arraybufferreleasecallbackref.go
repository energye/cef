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

// ICefv8ArrayBufferReleaseCallback Parent: ICefBaseRefCounted
type ICefv8ArrayBufferReleaseCallback interface {
	ICefBaseRefCounted
}

// ICefv8ArrayBufferReleaseCallbackRef Parent: ICefv8ArrayBufferReleaseCallback ICefBaseRefCountedRef
type ICefv8ArrayBufferReleaseCallbackRef interface {
	ICefv8ArrayBufferReleaseCallback
	ICefBaseRefCountedRef
	AsIntfV8ArrayBufferReleaseCallback() uintptr
}

type TCefv8ArrayBufferReleaseCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefv8ArrayBufferReleaseCallbackRef) AsIntfV8ArrayBufferReleaseCallback() uintptr {
	return m.GetIntfPointer(0)
}

// V8ArrayBufferReleaseCallbackRef  is static instance
var V8ArrayBufferReleaseCallbackRef _V8ArrayBufferReleaseCallbackRefClass

// _V8ArrayBufferReleaseCallbackRefClass is class type defined by TCefv8ArrayBufferReleaseCallbackRef
type _V8ArrayBufferReleaseCallbackRefClass uintptr

func (_V8ArrayBufferReleaseCallbackRefClass) UnWrap(data uintptr) (result IEngV8ArrayBufferReleaseCallback) {
	var resultPtr uintptr
	cefv8ArrayBufferReleaseCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngV8ArrayBufferReleaseCallback(resultPtr)
	return
}

// NewV8ArrayBufferReleaseCallbackRef class constructor
func NewV8ArrayBufferReleaseCallbackRef(data uintptr) ICefv8ArrayBufferReleaseCallbackRef {
	var v8ArrayBufferReleaseCallbackPtr uintptr // ICefv8ArrayBufferReleaseCallback
	r := cefv8ArrayBufferReleaseCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8ArrayBufferReleaseCallbackPtr)))
	ret := AsCefv8ArrayBufferReleaseCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8ArrayBufferReleaseCallbackPtr)
	}
	return ret
}

var (
	cefv8ArrayBufferReleaseCallbackRefOnce   base.Once
	cefv8ArrayBufferReleaseCallbackRefImport *imports.Imports = nil
)

func cefv8ArrayBufferReleaseCallbackRefAPI() *imports.Imports {
	cefv8ArrayBufferReleaseCallbackRefOnce.Do(func() {
		cefv8ArrayBufferReleaseCallbackRefImport = api.NewDefaultImports()
		cefv8ArrayBufferReleaseCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefv8ArrayBufferReleaseCallbackRef_Create", 0), // constructor NewV8ArrayBufferReleaseCallbackRef
			/* 1 */ imports.NewTable("TCefv8ArrayBufferReleaseCallbackRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefv8ArrayBufferReleaseCallbackRefImport
}
