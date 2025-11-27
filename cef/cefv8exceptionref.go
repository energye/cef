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

// ICefV8Exception Parent: ICefBaseRefCountedRef
type ICefV8Exception interface {
	ICefBaseRefCountedRef
	// GetMessage
	//  Returns the exception message.
	GetMessage() string // function
	// GetSourceLine
	//  Returns the line of source code that the exception occurred within.
	GetSourceLine() string // function
	// GetScriptResourceName
	//  Returns the resource name for the script from where the function causing
	//  the error originates.
	GetScriptResourceName() string // function
	// GetLineNumber
	//  Returns the 1-based number of the line where the error occurred or 0 if
	//  the line number is unknown.
	GetLineNumber() int32 // function
	// GetStartPosition
	//  Returns the index within the script of the first character where the error
	//  occurred.
	GetStartPosition() int32 // function
	// GetEndPosition
	//  Returns the index within the script of the last character where the error
	//  occurred.
	GetEndPosition() int32 // function
	// GetStartColumn
	//  Returns the index within the line of the first character where the error
	//  occurred.
	GetStartColumn() int32 // function
	// GetEndColumn
	//  Returns the index within the line of the last character where the error
	//  occurred.
	GetEndColumn() int32 // function
}

// ICefV8ExceptionRef Parent: ICefV8Exception
type ICefV8ExceptionRef interface {
	ICefV8Exception
	AsIntfV8Exception() uintptr
}

type TCefV8ExceptionRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefV8ExceptionRef) GetMessage() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8ExceptionRefAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8ExceptionRef) GetSourceLine() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8ExceptionRefAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8ExceptionRef) GetScriptResourceName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefV8ExceptionRefAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCefV8ExceptionRef) GetLineNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8ExceptionRefAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCefV8ExceptionRef) GetStartPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8ExceptionRefAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCefV8ExceptionRef) GetEndPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8ExceptionRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefV8ExceptionRef) GetStartColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8ExceptionRefAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TCefV8ExceptionRef) GetEndColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefV8ExceptionRefAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TCefV8ExceptionRef) AsIntfV8Exception() uintptr {
	return m.GetIntfPointer(0)
}

// V8ExceptionRef  is static instance
var V8ExceptionRef _V8ExceptionRefClass

// _V8ExceptionRefClass is class type defined by TCefV8ExceptionRef
type _V8ExceptionRefClass uintptr

func (_V8ExceptionRefClass) UnWrap(data uintptr) (result ICefV8Exception) {
	var resultPtr uintptr
	cefV8ExceptionRefAPI().SysCallN(9, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefV8ExceptionRef(resultPtr)
	return
}

// NewV8ExceptionRef class constructor
func NewV8ExceptionRef(data uintptr) ICefV8ExceptionRef {
	var v8ExceptionPtr uintptr // ICefV8Exception
	r := cefV8ExceptionRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8ExceptionPtr)))
	ret := AsCefV8ExceptionRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8ExceptionPtr)
	}
	return ret
}

var (
	cefV8ExceptionRefOnce   base.Once
	cefV8ExceptionRefImport *imports.Imports = nil
)

func cefV8ExceptionRefAPI() *imports.Imports {
	cefV8ExceptionRefOnce.Do(func() {
		cefV8ExceptionRefImport = api.NewDefaultImports()
		cefV8ExceptionRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefV8ExceptionRef_Create", 0), // constructor NewV8ExceptionRef
			/* 1 */ imports.NewTable("TCefV8ExceptionRef_GetMessage", 0), // function GetMessage
			/* 2 */ imports.NewTable("TCefV8ExceptionRef_GetSourceLine", 0), // function GetSourceLine
			/* 3 */ imports.NewTable("TCefV8ExceptionRef_GetScriptResourceName", 0), // function GetScriptResourceName
			/* 4 */ imports.NewTable("TCefV8ExceptionRef_GetLineNumber", 0), // function GetLineNumber
			/* 5 */ imports.NewTable("TCefV8ExceptionRef_GetStartPosition", 0), // function GetStartPosition
			/* 6 */ imports.NewTable("TCefV8ExceptionRef_GetEndPosition", 0), // function GetEndPosition
			/* 7 */ imports.NewTable("TCefV8ExceptionRef_GetStartColumn", 0), // function GetStartColumn
			/* 8 */ imports.NewTable("TCefV8ExceptionRef_GetEndColumn", 0), // function GetEndColumn
			/* 9 */ imports.NewTable("TCefV8ExceptionRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefV8ExceptionRefImport
}
