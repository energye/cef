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

// ICefv8Handler Parent: ICefBaseRefCounted
type ICefv8Handler interface {
	ICefBaseRefCounted
}

// ICefv8HandlerRef Parent: ICefv8Handler ICefBaseRefCountedRef
type ICefv8HandlerRef interface {
	ICefv8Handler
	ICefBaseRefCountedRef
	AsIntfV8Handler() uintptr
}

type TCefv8HandlerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefv8HandlerRef) AsIntfV8Handler() uintptr {
	return m.GetIntfPointer(0)
}

// V8HandlerRef  is static instance
var V8HandlerRef _V8HandlerRefClass

// _V8HandlerRefClass is class type defined by TCefv8HandlerRef
type _V8HandlerRefClass uintptr

func (_V8HandlerRefClass) UnWrap(data uintptr) (result IEngV8Handler) {
	var resultPtr uintptr
	cefv8HandlerRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngV8Handler(resultPtr)
	return
}

// NewV8HandlerRef class constructor
func NewV8HandlerRef(data uintptr) ICefv8HandlerRef {
	var v8HandlerPtr uintptr // ICefv8Handler
	r := cefv8HandlerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8HandlerPtr)))
	ret := AsCefv8HandlerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8HandlerPtr)
	}
	return ret
}

var (
	cefv8HandlerRefOnce   base.Once
	cefv8HandlerRefImport *imports.Imports = nil
)

func cefv8HandlerRefAPI() *imports.Imports {
	cefv8HandlerRefOnce.Do(func() {
		cefv8HandlerRefImport = api.NewDefaultImports()
		cefv8HandlerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefv8HandlerRef_Create", 0), // constructor NewV8HandlerRef
			/* 1 */ imports.NewTable("TCefv8HandlerRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefv8HandlerRefImport
}
