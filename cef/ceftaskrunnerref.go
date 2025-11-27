//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefTaskRunner Parent: ICefBaseRefCountedRef
type ICefTaskRunner interface {
	ICefBaseRefCountedRef
	// IsSame
	//  Returns true (1) if this object is pointing to the same task runner as
	//  |that| object.
	IsSame(that ICefTaskRunner) bool // function
	// BelongsToCurrentThread
	//  Returns true (1) if this task runner belongs to the current thread.
	BelongsToCurrentThread() bool // function
	// BelongsToThread
	//  Returns true (1) if this task runner is for the specified CEF thread.
	BelongsToThread(threadId cefTypes.TCefThreadId) bool // function
	// PostTask
	//  Post a task for execution on the thread associated with this task runner.
	//  Execution will occur asynchronously.
	PostTask(task IEngTask) bool // function
	// PostDelayedTask
	//  Post a task for delayed execution on the thread associated with this task
	//  runner. Execution will occur asynchronously. Delayed tasks are not
	//  supported on V8 WebWorker threads and will be executed without the
	//  specified delay.
	PostDelayedTask(task IEngTask, delayMs int64) bool // function
}

// ICefTaskRunnerRef Parent: ICefTaskRunner
type ICefTaskRunnerRef interface {
	ICefTaskRunner
	AsIntfTaskRunner() uintptr
}

type TCefTaskRunnerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefTaskRunnerRef) IsSame(that ICefTaskRunner) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskRunnerRefAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefTaskRunnerRef) BelongsToCurrentThread() bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskRunnerRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefTaskRunnerRef) BelongsToThread(threadId cefTypes.TCefThreadId) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskRunnerRefAPI().SysCallN(3, m.Instance(), uintptr(threadId))
	return api.GoBool(r)
}

func (m *TCefTaskRunnerRef) PostTask(task IEngTask) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskRunnerRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(task))
	return api.GoBool(r)
}

func (m *TCefTaskRunnerRef) PostDelayedTask(task IEngTask, delayMs int64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskRunnerRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(task), uintptr(base.UnsafePointer(&delayMs)))
	return api.GoBool(r)
}

func (m *TCefTaskRunnerRef) AsIntfTaskRunner() uintptr {
	return m.GetIntfPointer(0)
}

// TaskRunnerRef  is static instance
var TaskRunnerRef _TaskRunnerRefClass

// _TaskRunnerRefClass is class type defined by TCefTaskRunnerRef
type _TaskRunnerRefClass uintptr

// UnWrap
//
//	Returns a ICefTaskRunner instance using a PCefTaskRunner data pointer.
func (_TaskRunnerRefClass) UnWrap(data uintptr) (result ICefTaskRunner) {
	var resultPtr uintptr
	cefTaskRunnerRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

// GetForCurrentThread
//
//	Returns the task runner for the current thread. Only CEF threads will have
//	task runners. An NULL reference will be returned if this function is called
//	on an invalid thread.
func (_TaskRunnerRefClass) GetForCurrentThread() (result ICefTaskRunner) {
	var resultPtr uintptr
	cefTaskRunnerRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

// GetForThread
//
//	Returns the task runner for the specified CEF thread.
func (_TaskRunnerRefClass) GetForThread(threadId cefTypes.TCefThreadId) (result ICefTaskRunner) {
	var resultPtr uintptr
	cefTaskRunnerRefAPI().SysCallN(8, uintptr(threadId), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

// NewTaskRunnerRef class constructor
func NewTaskRunnerRef(data uintptr) ICefTaskRunnerRef {
	var taskRunnerPtr uintptr // ICefTaskRunner
	r := cefTaskRunnerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&taskRunnerPtr)))
	ret := AsCefTaskRunnerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskRunnerPtr)
	}
	return ret
}

var (
	cefTaskRunnerRefOnce   base.Once
	cefTaskRunnerRefImport *imports.Imports = nil
)

func cefTaskRunnerRefAPI() *imports.Imports {
	cefTaskRunnerRefOnce.Do(func() {
		cefTaskRunnerRefImport = api.NewDefaultImports()
		cefTaskRunnerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTaskRunnerRef_Create", 0), // constructor NewTaskRunnerRef
			/* 1 */ imports.NewTable("TCefTaskRunnerRef_IsSame", 0), // function IsSame
			/* 2 */ imports.NewTable("TCefTaskRunnerRef_BelongsToCurrentThread", 0), // function BelongsToCurrentThread
			/* 3 */ imports.NewTable("TCefTaskRunnerRef_BelongsToThread", 0), // function BelongsToThread
			/* 4 */ imports.NewTable("TCefTaskRunnerRef_PostTask", 0), // function PostTask
			/* 5 */ imports.NewTable("TCefTaskRunnerRef_PostDelayedTask", 0), // function PostDelayedTask
			/* 6 */ imports.NewTable("TCefTaskRunnerRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefTaskRunnerRef_GetForCurrentThread", 0), // static function GetForCurrentThread
			/* 8 */ imports.NewTable("TCefTaskRunnerRef_GetForThread", 0), // static function GetForThread
		}
	})
	return cefTaskRunnerRefImport
}
