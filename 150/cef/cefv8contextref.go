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
	// GetTaskRunner
	//  Returns the task runner associated with this context. V8 handles can only
	//  be accessed from the thread on which they are created. This function can
	//  be called on any render process thread.
	GetTaskRunner() ICefTaskRunner // function
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed
	//  on the current thread. Do not call any other functions if this function
	//  returns false (0).
	IsValid() bool // function
	// GetBrowser
	//  Returns the browser for this context. This function will return an NULL
	//  reference for WebWorker contexts.
	GetBrowser() ICefBrowser // function
	// GetFrame
	//  Returns the frame for this context. This function will return an NULL
	//  reference for WebWorker contexts.
	GetFrame() ICefFrame // function
	// GetGlobal
	//  Returns the global object for this context. The context must be entered
	//  before calling this function.
	GetGlobal() ICefv8Value // function
	// Enter
	//  Enter this context. A context must be explicitly entered before creating a
	//  V8 Object, Array, Function or Date asynchronously. exit() must be called
	//  the same number of times as enter() before releasing this context. V8
	//  objects belong to the context in which they are created. Returns true (1)
	//  if the scope was entered successfully.
	Enter() bool // function
	// Exit
	//  Exit this context. Call this function only after calling enter(). Returns
	//  true (1) if the scope was exited successfully.
	Exit() bool // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that|
	//  object.
	IsSame(that ICefv8Context) bool // function
	// Eval
	//  Execute a string of JavaScript code in this V8 context. The |script_url|
	//  parameter is the URL where the script in question can be found, if any.
	//  The |start_line| parameter is the base line number to use for error
	//  reporting. On success |retval| will be set to the return value, if any,
	//  and the function will return true (1). On failure |exception| will be set
	//  to the exception, if any, and the function will return false (0).
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
