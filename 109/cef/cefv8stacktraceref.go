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

// ICefV8StackTrace Parent: ICefBaseRefCounted
type ICefV8StackTrace interface {
	ICefBaseRefCounted
	IsValid() bool                         // function
	GetFrameCount() int32                  // function
	GetFrame(index int32) ICefV8StackFrame // function
}

// ICefV8StackTraceRef Parent: ICefV8StackTrace ICefBaseRefCountedRef
type ICefV8StackTraceRef interface {
	ICefV8StackTrace
	ICefBaseRefCountedRef
	AsIntfV8StackTrace() uintptr
}

type TCefV8StackTraceRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefV8StackTraceRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefV8StackTraceRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefV8StackTraceRef) GetFrameCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8StackTraceRefAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCefV8StackTraceRef) GetFrame(index int32) (result ICefV8StackFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefV8StackTraceRefAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8StackFrameRef(resultPtr)
	return
}

func (m *TCefV8StackTraceRef) AsIntfV8StackTrace() uintptr {
	return m.GetIntfPointer(0)
}

// V8StackTraceRef  is static instance
var V8StackTraceRef _V8StackTraceRefClass

// _V8StackTraceRefClass is class type defined by TCefV8StackTraceRef
type _V8StackTraceRefClass uintptr

func (_V8StackTraceRefClass) UnWrap(data uintptr) (result ICefV8StackTrace) {
	var resultPtr uintptr
	cefV8StackTraceRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8StackTraceRef(resultPtr)
	return
}

func (_V8StackTraceRefClass) Current(frameLimit int32) (result ICefV8StackTrace) {
	var resultPtr uintptr
	cefV8StackTraceRefAPI().SysCallN(5, uintptr(frameLimit), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8StackTraceRef(resultPtr)
	return
}

// NewV8StackTraceRef class constructor
func NewV8StackTraceRef(data uintptr) ICefV8StackTraceRef {
	var v8StackTracePtr uintptr // ICefV8StackTrace
	r := cefV8StackTraceRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8StackTracePtr)))
	ret := AsCefV8StackTraceRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8StackTracePtr)
	}
	return ret
}

var (
	cefV8StackTraceRefOnce   base.Once
	cefV8StackTraceRefImport *imports.Imports = nil
)

func cefV8StackTraceRefAPI() *imports.Imports {
	cefV8StackTraceRefOnce.Do(func() {
		cefV8StackTraceRefImport = api.NewDefaultImports()
		cefV8StackTraceRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefV8StackTraceRef_Create", 0), // constructor NewV8StackTraceRef
			/* 1 */ imports.NewTable("TCefV8StackTraceRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefV8StackTraceRef_GetFrameCount", 0), // function GetFrameCount
			/* 3 */ imports.NewTable("TCefV8StackTraceRef_GetFrame", 0), // function GetFrame
			/* 4 */ imports.NewTable("TCefV8StackTraceRef_UnWrap", 0), // static function UnWrap
			/* 5 */ imports.NewTable("TCefV8StackTraceRef_Current", 0), // static function Current
		}
	})
	return cefV8StackTraceRefImport
}
