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

// ICefv8Context Parent: ICefBaseRefCounted
type ICefv8Context interface {
	ICefBaseRefCounted
	GetTaskRunner() ICefTaskRunner                                                                             // function
	IsValid() bool                                                                                             // function
	GetBrowser() ICefBrowser                                                                                   // function
	GetFrame() ICefFrame                                                                                       // function
	GetGlobal() ICefv8Value                                                                                    // function
	Enter() bool                                                                                               // function
	Exit() bool                                                                                                // function
	IsSame(that ICefv8Context) bool                                                                            // function
	Eval(code string, scriptUrl string, startLine int32, retval *ICefv8Value, exception *ICefV8Exception) bool // function
}

// ICefv8ContextRef Parent: ICefv8Context ICefBaseRefCountedRef
type ICefv8ContextRef interface {
	ICefv8Context
	ICefBaseRefCountedRef
	AsIntfV8Context() uintptr
}

type TCefv8ContextRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefv8ContextRef) GetTaskRunner() (result ICefTaskRunner) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

func (m *TCefv8ContextRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefv8ContextRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ContextRef) GetBrowser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCefv8ContextRef) GetFrame() (result ICefFrame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefFrameRef(resultPtr)
	return
}

func (m *TCefv8ContextRef) GetGlobal() (result ICefv8Value) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ValueRef(resultPtr)
	return
}

func (m *TCefv8ContextRef) Enter() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ContextRefAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ContextRef) Exit() bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ContextRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefv8ContextRef) IsSame(that ICefv8Context) bool {
	if !m.IsValid() {
		return false
	}
	r := cefv8ContextRefAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefv8ContextRef) Eval(code string, scriptUrl string, startLine int32, retval *ICefv8Value, exception *ICefV8Exception) bool {
	if !m.IsValid() {
		return false
	}
	retvalPtr := base.GetObjectUintptr(*retval)
	exceptionPtr := base.GetObjectUintptr(*exception)
	r := cefv8ContextRefAPI().SysCallN(9, m.Instance(), api.PasStr(code), api.PasStr(scriptUrl), uintptr(startLine), uintptr(base.UnsafePointer(&retvalPtr)), uintptr(base.UnsafePointer(&exceptionPtr)))
	*retval = AsCefv8ValueRef(retvalPtr)
	*exception = AsCefV8ExceptionRef(exceptionPtr)
	return api.GoBool(r)
}

func (m *TCefv8ContextRef) AsIntfV8Context() uintptr {
	return m.GetIntfPointer(0)
}

// V8ContextRef  is static instance
var V8ContextRef _V8ContextRefClass

// _V8ContextRefClass is class type defined by TCefv8ContextRef
type _V8ContextRefClass uintptr

func (_V8ContextRefClass) UnWrap(data uintptr) (result ICefv8Context) {
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(10, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ContextRef(resultPtr)
	return
}

func (_V8ContextRefClass) Current() (result ICefv8Context) {
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(11, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ContextRef(resultPtr)
	return
}

func (_V8ContextRefClass) Entered() (result ICefv8Context) {
	var resultPtr uintptr
	cefv8ContextRefAPI().SysCallN(12, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefv8ContextRef(resultPtr)
	return
}

// NewV8ContextRef class constructor
func NewV8ContextRef(data uintptr) ICefv8ContextRef {
	var v8ContextPtr uintptr // ICefv8Context
	r := cefv8ContextRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&v8ContextPtr)))
	ret := AsCefv8ContextRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8ContextPtr)
	}
	return ret
}

var (
	cefv8ContextRefOnce   base.Once
	cefv8ContextRefImport *imports.Imports = nil
)

func cefv8ContextRefAPI() *imports.Imports {
	cefv8ContextRefOnce.Do(func() {
		cefv8ContextRefImport = api.NewDefaultImports()
		cefv8ContextRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefv8ContextRef_Create", 0), // constructor NewV8ContextRef
			/* 1 */ imports.NewTable("TCefv8ContextRef_GetTaskRunner", 0), // function GetTaskRunner
			/* 2 */ imports.NewTable("TCefv8ContextRef_IsValid", 0), // function IsValid
			/* 3 */ imports.NewTable("TCefv8ContextRef_GetBrowser", 0), // function GetBrowser
			/* 4 */ imports.NewTable("TCefv8ContextRef_GetFrame", 0), // function GetFrame
			/* 5 */ imports.NewTable("TCefv8ContextRef_GetGlobal", 0), // function GetGlobal
			/* 6 */ imports.NewTable("TCefv8ContextRef_Enter", 0), // function Enter
			/* 7 */ imports.NewTable("TCefv8ContextRef_Exit", 0), // function Exit
			/* 8 */ imports.NewTable("TCefv8ContextRef_IsSame", 0), // function IsSame
			/* 9 */ imports.NewTable("TCefv8ContextRef_Eval", 0), // function Eval
			/* 10 */ imports.NewTable("TCefv8ContextRef_UnWrap", 0), // static function UnWrap
			/* 11 */ imports.NewTable("TCefv8ContextRef_Current", 0), // static function Current
			/* 12 */ imports.NewTable("TCefv8ContextRef_Entered", 0), // static function Entered
		}
	})
	return cefv8ContextRefImport
}
