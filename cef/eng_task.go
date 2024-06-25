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

// ITask Parent: ICefTask
//
//	Implement this interface for asynchronous task execution. If the task is
//	posted successfully and if the associated message loop is still running then
//	the execute() function will be called on the target thread. If the task
//	fails to post then the task object may be destroyed on the source thread
//	instead of the target thread. For this reason be cautious when performing
//	work in the task object destructor.
//	<a cref="uCEFTypes|TCefTask">Implements TCefTask</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_t)</a>
type ITask interface {
	ICefTask
	AsInterface() ICefTask                                       // function
	OnlyPostTask(threadId TCefThreadId) bool                     // function
	OnlyPostDelayedTask(threadId TCefThreadId, delay int64) bool // function
	// SetOnExecute
	//  Method that will be executed on the target thread.
	SetOnExecute(fn TTaskExecute) // property event
}

// TTask Parent: TCefTask
//
//	Implement this interface for asynchronous task execution. If the task is
//	posted successfully and if the associated message loop is still running then
//	the execute() function will be called on the target thread. If the task
//	fails to post then the task object may be destroyed on the source thread
//	instead of the target thread. For this reason be cautious when performing
//	work in the task object destructor.
//	<a cref="uCEFTypes|TCefTask">Implements TCefTask</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_t)</a>
type TTask struct {
	TCefTask
	executePtr uintptr
}

func NewTask() ITask {
	r1 := askImportAPI().SysCallN(1)
	return AsTask(r1)
}

func (m *TTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	askImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

func (m *TTask) OnlyPostTask(threadId TCefThreadId) bool {
	r1 := askImportAPI().SysCallN(3, m.Instance(), uintptr(threadId))
	return GoBool(r1)
}

func (m *TTask) OnlyPostDelayedTask(threadId TCefThreadId, delay int64) bool {
	r1 := askImportAPI().SysCallN(2, m.Instance(), uintptr(threadId), uintptr(unsafePointer(&delay)))
	return GoBool(r1)
}

func (m *TTask) SetOnExecute(fn TTaskExecute) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	askImportAPI().SysCallN(4, m.Instance(), m.executePtr)
}

var (
	askImport       *imports.Imports = nil
	askImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Task_AsInterface", 0),
		/*1*/ imports.NewTable("Task_Create", 0),
		/*2*/ imports.NewTable("Task_OnlyPostDelayedTask", 0),
		/*3*/ imports.NewTable("Task_OnlyPostTask", 0),
		/*4*/ imports.NewTable("Task_SetOnExecute", 0),
	}
)

func askImportAPI() *imports.Imports {
	if askImport == nil {
		askImport = NewDefaultImports()
		askImport.SetImportTable(askImportTables)
		askImportTables = nil
	}
	return askImport
}
