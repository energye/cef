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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/types"
)

// ICefTaskManager Parent: ICefBaseRefCounted
type ICefTaskManager interface {
	ICefBaseRefCounted
	// GetTasksCount
	//  Returns the number of tasks currently tracked by the task manager. Returns
	//  0 if the function was called from the incorrect thread.
	GetTasksCount() cefTypes.NativeUInt // function
	// GetTaskIdsList
	//  Gets the list of task IDs currently tracked by the task manager. Tasks
	//  that share the same process id will always be consecutive. The list will
	//  be sorted in a way that reflects the process tree: the browser process
	//  will be first, followed by the gpu process if it exists. Related processes
	//  (e.g., a subframe process and its parent) will be kept together if
	//  possible. Callers can expect this ordering to be stable when a process is
	//  added or removed. The task IDs are unique within the application lifespan.
	//  Returns false (0) if the function was called from the incorrect thread.
	GetTaskIdsList(taskIds *lcl.IInt64ArrayWrap) bool // function
	// GetTaskInfo
	//  Gets information about the task with |task_id|. Returns true (1) if the
	//  information about the task was successfully retrieved and false (0) if the
	//  |task_id| is invalid or the function was called from the incorrect thread.
	GetTaskInfo(taskId int64, info *TCustomTaskInfo) bool // function
	// KillTask
	//  Attempts to terminate a task with |task_id|. Returns false (0) if the
	//  |task_id| is invalid, the call is made from an incorrect thread, or if the
	//  task cannot be terminated.
	KillTask(taskId int64) bool // function
	// GetTaskIdForBrowserId
	//  Returns the task ID associated with the main task for |browser_id| (value
	//  from cef_browser_t::GetIdentifier). Returns -1 if |browser_id| is invalid,
	//  does not currently have an associated task, or the function was called
	//  from the incorrect thread.
	GetTaskIdForBrowserId(browserId int32) int64 // function
}

// ICefTaskManagerRef Parent: ICefTaskManager ICefBaseRefCountedRef
type ICefTaskManagerRef interface {
	ICefTaskManager
	ICefBaseRefCountedRef
	AsIntfTaskManager() uintptr
}

type TCefTaskManagerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefTaskManagerRef) GetTasksCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefTaskManagerRefAPI().SysCallN(1, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefTaskManagerRef) GetTaskIdsList(taskIds *lcl.IInt64ArrayWrap) bool {
	if !m.IsValid() {
		return false
	}
	taskIdsPtr := base.GetObjectUintptr(*taskIds)
	r := cefTaskManagerRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&taskIdsPtr)))
	*taskIds = lcl.AsInt64ArrayWrap(taskIdsPtr)
	return api.GoBool(r)
}

func (m *TCefTaskManagerRef) GetTaskInfo(taskId int64, info *TCustomTaskInfo) bool {
	if !m.IsValid() {
		return false
	}
	infoPtr := info.ToPas()
	r := cefTaskManagerRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&taskId)), uintptr(base.UnsafePointer(infoPtr)))
	*info = infoPtr.ToGo()
	return api.GoBool(r)
}

func (m *TCefTaskManagerRef) KillTask(taskId int64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTaskManagerRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&taskId)))
	return api.GoBool(r)
}

func (m *TCefTaskManagerRef) GetTaskIdForBrowserId(browserId int32) (result int64) {
	if !m.IsValid() {
		return
	}
	cefTaskManagerRefAPI().SysCallN(5, m.Instance(), uintptr(browserId), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefTaskManagerRef) AsIntfTaskManager() uintptr {
	return m.GetIntfPointer(0)
}

// TaskManagerRef  is static instance
var TaskManagerRef _TaskManagerRefClass

// _TaskManagerRefClass is class type defined by TCefTaskManagerRef
type _TaskManagerRefClass uintptr

func (_TaskManagerRefClass) UnWrap(data uintptr) (result ICefTaskManager) {
	var resultPtr uintptr
	cefTaskManagerRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskManagerRef(resultPtr)
	return
}

// New
//
//	Returns the global task manager.
//	This function may only be called on the CEF UI thread.
func (_TaskManagerRefClass) New() (result ICefTaskManager) {
	var resultPtr uintptr
	cefTaskManagerRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskManagerRef(resultPtr)
	return
}

// NewTaskManagerRef class constructor
func NewTaskManagerRef(data uintptr) ICefTaskManagerRef {
	var taskManagerPtr uintptr // ICefTaskManager
	r := cefTaskManagerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&taskManagerPtr)))
	ret := AsCefTaskManagerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskManagerPtr)
	}
	return ret
}

var (
	cefTaskManagerRefOnce   base.Once
	cefTaskManagerRefImport *imports.Imports = nil
)

func cefTaskManagerRefAPI() *imports.Imports {
	cefTaskManagerRefOnce.Do(func() {
		cefTaskManagerRefImport = api.NewDefaultImports()
		cefTaskManagerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTaskManagerRef_Create", 0), // constructor NewTaskManagerRef
			/* 1 */ imports.NewTable("TCefTaskManagerRef_GetTasksCount", 0), // function GetTasksCount
			/* 2 */ imports.NewTable("TCefTaskManagerRef_GetTaskIdsList", 0), // function GetTaskIdsList
			/* 3 */ imports.NewTable("TCefTaskManagerRef_GetTaskInfo", 0), // function GetTaskInfo
			/* 4 */ imports.NewTable("TCefTaskManagerRef_KillTask", 0), // function KillTask
			/* 5 */ imports.NewTable("TCefTaskManagerRef_GetTaskIdForBrowserId", 0), // function GetTaskIdForBrowserId
			/* 6 */ imports.NewTable("TCefTaskManagerRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefTaskManagerRef_New", 0), // static function New
		}
	})
	return cefTaskManagerRefImport
}
