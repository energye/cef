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

// ICefV8Exception Parent: ICefBaseRefCounted
type ICefV8Exception interface {
	ICefBaseRefCounted
	GetMessage() string            // function
	GetSourceLine() string         // function
	GetScriptResourceName() string // function
	GetLineNumber() int32          // function
	GetStartPosition() int32       // function
	GetEndPosition() int32         // function
	GetStartColumn() int32         // function
	GetEndColumn() int32           // function
}

// TCefV8Exception Parent: TCefBaseRefCounted
type TCefV8Exception struct {
	TCefBaseRefCounted
}

// V8ExceptionRef -> ICefV8Exception
var V8ExceptionRef v8Exception

// v8Exception TCefV8Exception Ref
type v8Exception uintptr

func (m *v8Exception) UnWrap(data uintptr) ICefV8Exception {
	var resultCefV8Exception uintptr
	v8ExceptionImportAPI().SysCallN(8, uintptr(data), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefV8Exception) GetMessage() string {
	r1 := v8ExceptionImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetSourceLine() string {
	r1 := v8ExceptionImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetScriptResourceName() string {
	r1 := v8ExceptionImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCefV8Exception) GetLineNumber() int32 {
	r1 := v8ExceptionImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetStartPosition() int32 {
	r1 := v8ExceptionImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetEndPosition() int32 {
	r1 := v8ExceptionImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetStartColumn() int32 {
	r1 := v8ExceptionImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TCefV8Exception) GetEndColumn() int32 {
	r1 := v8ExceptionImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

var (
	v8ExceptionImport       *imports.Imports = nil
	v8ExceptionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8Exception_GetEndColumn", 0),
		/*1*/ imports.NewTable("CefV8Exception_GetEndPosition", 0),
		/*2*/ imports.NewTable("CefV8Exception_GetLineNumber", 0),
		/*3*/ imports.NewTable("CefV8Exception_GetMessage", 0),
		/*4*/ imports.NewTable("CefV8Exception_GetScriptResourceName", 0),
		/*5*/ imports.NewTable("CefV8Exception_GetSourceLine", 0),
		/*6*/ imports.NewTable("CefV8Exception_GetStartColumn", 0),
		/*7*/ imports.NewTable("CefV8Exception_GetStartPosition", 0),
		/*8*/ imports.NewTable("CefV8Exception_UnWrap", 0),
	}
)

func v8ExceptionImportAPI() *imports.Imports {
	if v8ExceptionImport == nil {
		v8ExceptionImport = NewDefaultImports()
		v8ExceptionImport.SetImportTable(v8ExceptionImportTables)
		v8ExceptionImportTables = nil
	}
	return v8ExceptionImport
}
