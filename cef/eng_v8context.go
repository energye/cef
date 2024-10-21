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

// ICefV8Context Parent: ICefBaseRefCounted
//
//	Interface representing a V8 context handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefV8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8context_t))</a>
type ICefV8Context interface {
	ICefBaseRefCounted
	// GetTaskRunner
	//  Returns the task runner associated with this context. V8 handles can only be accessed from the thread on which they are created. This function can be called on any render process thread.
	GetTaskRunner() ICefTaskRunner // function
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetBrowser
	//  Returns the browser for this context. This function will return an NULL reference for WebWorker contexts.
	GetBrowser() ICefBrowser // function
	// GetFrame
	//  Returns the frame for this context. This function will return an NULL reference for WebWorker contexts.
	GetFrame() ICefFrame // function
	// GetGlobal
	//  Returns the global object for this context. The context must be entered before calling this function.
	GetGlobal() ICefV8Value // function
	// Enter
	//  Enter this context. A context must be explicitly entered before creating a V8 Object, Array, Function or Date asynchronously. exit() must be called the same number of times as enter() before releasing this context. V8 objects belong to the context in which they are created. Returns true (1) if the scope was entered successfully.
	Enter() bool // function
	// Exit
	//  Exit this context. Call this function only after calling enter(). Returns true (1) if the scope was exited successfully.
	Exit() bool // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefV8Context) bool // function
	// Eval
	//  Execute a string of JavaScript code in this V8 context. The |script_url| parameter is the URL where the script in question can be found, if any. The |start_line| parameter is the base line number to use for error reporting. On success |retval| will be set to the return value, if any, and the function will return true (1). On failure |exception| will be set to the exception, if any, and the function will return false (0).
	Eval(code string, scripturl string, startline int32, retval *ICefV8Value, exception *ICefV8Exception) bool // function
}

// TCefV8Context Parent: TCefBaseRefCounted
//
//	Interface representing a V8 context handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefV8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8context_t))</a>
type TCefV8Context struct {
	TCefBaseRefCounted
}

// V8ContextRef -> ICefV8Context
var V8ContextRef v8Context

// v8Context TCefV8Context Ref
type v8Context uintptr

func (m *v8Context) UnWrap(data uintptr) ICefV8Context {
	var resultCefV8Context uintptr
	v8ContextImportAPI().SysCallN(11, uintptr(data), uintptr(unsafePointer(&resultCefV8Context)))
	return AsCefV8Context(resultCefV8Context)
}

func (m *v8Context) Current() ICefV8Context {
	var resultCefV8Context uintptr
	v8ContextImportAPI().SysCallN(0, uintptr(unsafePointer(&resultCefV8Context)))
	return AsCefV8Context(resultCefV8Context)
}

func (m *v8Context) Entered() ICefV8Context {
	var resultCefV8Context uintptr
	v8ContextImportAPI().SysCallN(2, uintptr(unsafePointer(&resultCefV8Context)))
	return AsCefV8Context(resultCefV8Context)
}

func (m *TCefV8Context) GetTaskRunner() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	v8ContextImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCefV8Context) IsValid() bool {
	r1 := v8ContextImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Context) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	v8ContextImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefV8Context) GetFrame() ICefFrame {
	var resultCefFrame uintptr
	v8ContextImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefV8Context) GetGlobal() ICefV8Value {
	var resultCefV8Value uintptr
	v8ContextImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCefV8Value)))
	return AsCefV8Value(resultCefV8Value)
}

func (m *TCefV8Context) Enter() bool {
	r1 := v8ContextImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Context) Exit() bool {
	r1 := v8ContextImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8Context) IsSame(that ICefV8Context) bool {
	r1 := v8ContextImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefV8Context) Eval(code string, scripturl string, startline int32, retval *ICefV8Value, exception *ICefV8Exception) bool {
	var result3 uintptr
	var result4 uintptr
	r1 := v8ContextImportAPI().SysCallN(3, m.Instance(), PascalStr(code), PascalStr(scripturl), uintptr(startline), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*retval = AsCefV8Value(result3)
	*exception = AsCefV8Exception(result4)
	return GoBool(r1)
}

var (
	v8ContextImport       *imports.Imports = nil
	v8ContextImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8Context_Current", 0),
		/*1*/ imports.NewTable("CefV8Context_Enter", 0),
		/*2*/ imports.NewTable("CefV8Context_Entered", 0),
		/*3*/ imports.NewTable("CefV8Context_Eval", 0),
		/*4*/ imports.NewTable("CefV8Context_Exit", 0),
		/*5*/ imports.NewTable("CefV8Context_GetBrowser", 0),
		/*6*/ imports.NewTable("CefV8Context_GetFrame", 0),
		/*7*/ imports.NewTable("CefV8Context_GetGlobal", 0),
		/*8*/ imports.NewTable("CefV8Context_GetTaskRunner", 0),
		/*9*/ imports.NewTable("CefV8Context_IsSame", 0),
		/*10*/ imports.NewTable("CefV8Context_IsValid", 0),
		/*11*/ imports.NewTable("CefV8Context_UnWrap", 0),
	}
)

func v8ContextImportAPI() *imports.Imports {
	if v8ContextImport == nil {
		v8ContextImport = NewDefaultImports()
		v8ContextImport.SetImportTable(v8ContextImportTables)
		v8ContextImportTables = nil
	}
	return v8ContextImport
}
