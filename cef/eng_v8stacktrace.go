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

// ICefV8StackTrace Parent: ICefBaseRefCounted
//
//	Interface representing a V8 stack trace handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_trace_t))</a>
type ICefV8StackTrace interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// GetFrameCount
	//  Returns the number of stack frames.
	GetFrameCount() int32 // function
	// GetFrame
	//  Returns the stack frame at the specified 0-based index.
	GetFrame(index int32) ICefV8StackFrame // function
}

// TCefV8StackTrace Parent: TCefBaseRefCounted
//
//	Interface representing a V8 stack trace handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8stack_trace_t))</a>
type TCefV8StackTrace struct {
	TCefBaseRefCounted
}

// V8StackTraceRef -> ICefV8StackTrace
var V8StackTraceRef v8StackTrace

// v8StackTrace TCefV8StackTrace Ref
type v8StackTrace uintptr

func (m *v8StackTrace) UnWrap(data uintptr) ICefV8StackTrace {
	var resultCefV8StackTrace uintptr
	v8StackTraceImportAPI().SysCallN(4, uintptr(data), uintptr(unsafePointer(&resultCefV8StackTrace)))
	return AsCefV8StackTrace(resultCefV8StackTrace)
}

func (m *v8StackTrace) Current(frameLimit int32) ICefV8StackTrace {
	var resultCefV8StackTrace uintptr
	v8StackTraceImportAPI().SysCallN(0, uintptr(frameLimit), uintptr(unsafePointer(&resultCefV8StackTrace)))
	return AsCefV8StackTrace(resultCefV8StackTrace)
}

func (m *TCefV8StackTrace) IsValid() bool {
	r1 := v8StackTraceImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCefV8StackTrace) GetFrameCount() int32 {
	r1 := v8StackTraceImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefV8StackTrace) GetFrame(index int32) ICefV8StackFrame {
	var resultCefV8StackFrame uintptr
	v8StackTraceImportAPI().SysCallN(1, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefV8StackFrame)))
	return AsCefV8StackFrame(resultCefV8StackFrame)
}

var (
	v8StackTraceImport       *imports.Imports = nil
	v8StackTraceImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefV8StackTrace_Current", 0),
		/*1*/ imports.NewTable("CefV8StackTrace_GetFrame", 0),
		/*2*/ imports.NewTable("CefV8StackTrace_GetFrameCount", 0),
		/*3*/ imports.NewTable("CefV8StackTrace_IsValid", 0),
		/*4*/ imports.NewTable("CefV8StackTrace_UnWrap", 0),
	}
)

func v8StackTraceImportAPI() *imports.Imports {
	if v8StackTraceImport == nil {
		v8StackTraceImport = NewDefaultImports()
		v8StackTraceImport.SetImportTable(v8StackTraceImportTables)
		v8StackTraceImportTables = nil
	}
	return v8StackTraceImport
}
