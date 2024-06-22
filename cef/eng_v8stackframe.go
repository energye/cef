//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICefV8StackFrame Parent: ICefBaseRefCounted
type ICefV8StackFrame interface {
	ICefBaseRefCounted
	IsValid() bool                    // function
	GetScriptName() string            // function
	GetScriptNameOrSourceUrl() string // function
	GetFunctionName() string          // function
	GetLineNumber() int32             // function
	GetColumn() int32                 // function
	IsEval() bool                     // function
	IsConstructor() bool              // function
}

// TCefV8StackFrame Parent: TCefBaseRefCounted
type TCefV8StackFrame struct {
	TCefBaseRefCounted
}

// V8StackFrameRef -> ICefV8StackFrame
var V8StackFrameRef v8StackFrame

// v8StackFrame TCefV8StackFrame Ref
type v8StackFrame uintptr

func (m *v8StackFrame) UnWrap(data uintptr) ICefV8StackFrame {
	var resultCefV8StackFrame uintptr
	v8StackFrameImportAPI().SysCallN(8, uintptr(data), uintptr(unsafePointer(&resultCefV8StackFrame)))
	return AsCefV8StackFrame(resultCefV8StackFrame)
}

func (m *TCefV8StackFrame) IsValid() bool {
	r1 := v8StackFrameImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackFrame) GetScriptName() string {
	r1 := v8StackFrameImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetScriptNameOrSourceUrl() string {
	r1 := v8StackFrameImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetFunctionName() string {
	r1 := v8StackFrameImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8StackFrame) GetLineNumber() int32 {
	r1 := v8StackFrameImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackFrame) GetColumn() int32 {
	r1 := v8StackFrameImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackFrame) IsEval() bool {
	r1 := v8StackFrameImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackFrame) IsConstructor() bool {
	r1 := v8StackFrameImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

var (
	v8StackFrameImport       *imports.Imports = nil
	v8StackFrameImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8StackFrame_GetColumn", 0),
		/*1*/ imports.NewTable("CefV8StackFrame_GetFunctionName", 0),
		/*2*/ imports.NewTable("CefV8StackFrame_GetLineNumber", 0),
		/*3*/ imports.NewTable("CefV8StackFrame_GetScriptName", 0),
		/*4*/ imports.NewTable("CefV8StackFrame_GetScriptNameOrSourceUrl", 0),
		/*5*/ imports.NewTable("CefV8StackFrame_IsConstructor", 0),
		/*6*/ imports.NewTable("CefV8StackFrame_IsEval", 0),
		/*7*/ imports.NewTable("CefV8StackFrame_IsValid", 0),
		/*8*/ imports.NewTable("CefV8StackFrame_UnWrap", 0),
	}
)

func v8StackFrameImportAPI() *imports.Imports {
	if v8StackFrameImport == nil {
		v8StackFrameImport = NewDefaultImports()
		v8StackFrameImport.SetImportTable(v8StackFrameImportTables)
		v8StackFrameImportTables = nil
	}
	return v8StackFrameImport
}
