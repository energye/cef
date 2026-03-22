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

// ICefV8StackFrame Parent: ICefBaseRefCounted
type ICefV8StackFrame interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed
	//  on the current thread. Do not call any other functions if this function
	//  returns false (0).
	IsValid() bool // function
	// GetScriptName
	//  Returns the name of the resource script that contains the function.
	GetScriptName() string // function
	// GetScriptNameOrSourceUrl
	//  Returns the name of the resource script that contains the function or the
	//  sourceURL value if the script name is undefined and its source ends with a
	//  "//@ sourceURL=..." string.
	GetScriptNameOrSourceUrl() string // function
	// GetFunctionName
	//  Returns the name of the function.
	GetFunctionName() string // function
	// GetLineNumber
	//  Returns the 1-based line number for the function call or 0 if unknown.
	GetLineNumber() int32 // function
	// GetColumn
	//  Returns the 1-based column offset on the line for the function call or 0
	//  if unknown.
	GetColumn() int32 // function
	// IsEval
	//  Returns true (1) if the function was compiled using eval().
	IsEval() bool // function
	// IsConstructor
	//  Returns true (1) if the function was called as a constructor via "new".
	IsConstructor() bool // function
}

// ICefV8StackFrameRef Parent: ICefV8StackFrame ICefBaseRefCountedRef
type ICefV8StackFrameRef interface {
	ICefV8StackFrame
	ICefBaseRefCountedRef
	AsIntfV8StackFrame() uintptr
}

type TCefV8StackFrameRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefV8StackFrameRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefV8StackFrameRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefV8StackFrameRef) GetScriptName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8StackFrameRefAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8StackFrameRef) GetScriptNameOrSourceUrl() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8StackFrameRefAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8StackFrameRef) GetFunctionName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8StackFrameRefAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8StackFrameRef) GetLineNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8StackFrameRefAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCefV8StackFrameRef) GetColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8StackFrameRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefV8StackFrameRef) IsEval() bool {
	if !m.IsValid() {
		return false
	}
	r := cefV8StackFrameRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefV8StackFrameRef) IsConstructor() bool {
	if !m.IsValid() {
		return false
	}
	r := cefV8StackFrameRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefV8StackFrameRef) AsIntfV8StackFrame() uintptr {
	return m.GetIntfPointer(0)
}

// V8StackFrameRef  is static instance
var V8StackFrameRef _V8StackFrameRefClass

// _V8StackFrameRefClass is class type defined by TCefV8StackFrameRef
type _V8StackFrameRefClass uintptr

func (_V8StackFrameRefClass) UnWrap(data uintptr) (result ICefV8StackFrame) {
	var resultPtr uintptr
	cefV8StackFrameRefAPI().SysCallN(9, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8StackFrameRef(resultPtr)
	return
}

// NewV8StackFrameRef class constructor
func NewV8StackFrameRef(data uintptr) ICefV8StackFrameRef {
	var v8StackFramePtr uintptr // ICefV8StackFrame
	r := cefV8StackFrameRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8StackFramePtr)))
	ret := AsCefV8StackFrameRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8StackFramePtr)
	}
	return ret
}

var (
	cefV8StackFrameRefOnce   base.Once
	cefV8StackFrameRefImport *imports.Imports = nil
)

func cefV8StackFrameRefAPI() *imports.Imports {
	cefV8StackFrameRefOnce.Do(func() {
		cefV8StackFrameRefImport = api.NewDefaultImports()
		cefV8StackFrameRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefV8StackFrameRef_Create", 0), // constructor NewV8StackFrameRef
			/* 1 */ imports.NewTable("TCefV8StackFrameRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefV8StackFrameRef_GetScriptName", 0), // function GetScriptName
			/* 3 */ imports.NewTable("TCefV8StackFrameRef_GetScriptNameOrSourceUrl", 0), // function GetScriptNameOrSourceUrl
			/* 4 */ imports.NewTable("TCefV8StackFrameRef_GetFunctionName", 0), // function GetFunctionName
			/* 5 */ imports.NewTable("TCefV8StackFrameRef_GetLineNumber", 0), // function GetLineNumber
			/* 6 */ imports.NewTable("TCefV8StackFrameRef_GetColumn", 0), // function GetColumn
			/* 7 */ imports.NewTable("TCefV8StackFrameRef_IsEval", 0), // function IsEval
			/* 8 */ imports.NewTable("TCefV8StackFrameRef_IsConstructor", 0), // function IsConstructor
			/* 9 */ imports.NewTable("TCefV8StackFrameRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefV8StackFrameRefImport
}
