//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/127/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefThread Parent: ICefBaseRefCounted
type ICefThread interface {
	ICefBaseRefCounted
	// GetTaskRunner
	//  Returns the ICefTaskRunner that will execute code on this thread's
	//  message loop. This function is safe to call from any thread.
	GetTaskRunner() ICefTaskRunner // function
	// GetPlatformThreadID
	//  Returns the platform thread ID. It will return the same value after stop()
	//  is called. This function is safe to call from any thread.
	GetPlatformThreadID() cefTypes.TCefPlatformThreadId // function
	// IsRunning
	//  Returns true (1) if the thread is currently running. This function must be
	//  called from the same thread that called cef_thread_create().
	IsRunning() bool // function
	// Stop
	//  Stop and join the thread. This function must be called from the same
	//  thread that called cef_thread_create(). Do not call this function if
	//  cef_thread_create() was called with a |stoppable| value of false (0).
	Stop() // procedure
}

// ICefThreadRef Parent: ICefThread ICefBaseRefCountedRef
type ICefThreadRef interface {
	ICefThread
	ICefBaseRefCountedRef
	AsIntfThread() uintptr
}

type TCefThreadRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefThreadRef) GetTaskRunner() (result ICefTaskRunner) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefThreadRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

func (m *TCefThreadRef) GetPlatformThreadID() cefTypes.TCefPlatformThreadId {
	if !m.IsValid() {
		return 0
	}
	r := cefThreadRefAPI().SysCallN(2, m.Instance())
	return cefTypes.TCefPlatformThreadId(r)
}

func (m *TCefThreadRef) IsRunning() bool {
	if !m.IsValid() {
		return false
	}
	r := cefThreadRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefThreadRef) Stop() {
	if !m.IsValid() {
		return
	}
	cefThreadRefAPI().SysCallN(6, m.Instance())
}

func (m *TCefThreadRef) AsIntfThread() uintptr {
	return m.GetIntfPointer(0)
}

// ThreadRef  is static instance
var ThreadRef _ThreadRefClass

// _ThreadRefClass is class type defined by TCefThreadRef
type _ThreadRefClass uintptr

func (_ThreadRefClass) UnWrap(data uintptr) (result ICefThread) {
	var resultPtr uintptr
	cefThreadRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefThreadRef(resultPtr)
	return
}

func (_ThreadRefClass) New(displayName string, priority cefTypes.TCefThreadPriority, messageLoopType cefTypes.TCefMessageLoopType, stoppable int32, comInitMode cefTypes.TCefCOMInitMode) (result ICefThread) {
	var resultPtr uintptr
	cefThreadRefAPI().SysCallN(5, api.PasStr(displayName), uintptr(priority), uintptr(messageLoopType), uintptr(stoppable), uintptr(comInitMode), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefThreadRef(resultPtr)
	return
}

// NewThreadRef class constructor
func NewThreadRef(data uintptr) ICefThreadRef {
	var threadPtr uintptr // ICefThread
	r := cefThreadRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&threadPtr)))
	ret := AsCefThreadRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, threadPtr)
	}
	return ret
}

var (
	cefThreadRefOnce   base.Once
	cefThreadRefImport *imports.Imports = nil
)

func cefThreadRefAPI() *imports.Imports {
	cefThreadRefOnce.Do(func() {
		cefThreadRefImport = api.NewDefaultImports()
		cefThreadRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefThreadRef_Create", 0), // constructor NewThreadRef
			/* 1 */ imports.NewTable("TCefThreadRef_GetTaskRunner", 0), // function GetTaskRunner
			/* 2 */ imports.NewTable("TCefThreadRef_GetPlatformThreadID", 0), // function GetPlatformThreadID
			/* 3 */ imports.NewTable("TCefThreadRef_IsRunning", 0), // function IsRunning
			/* 4 */ imports.NewTable("TCefThreadRef_UnWrap", 0), // static function UnWrap
			/* 5 */ imports.NewTable("TCefThreadRef_New", 0), // static function New
			/* 6 */ imports.NewTable("TCefThreadRef_Stop", 0), // procedure Stop
		}
	})
	return cefThreadRefImport
}
