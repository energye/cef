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

// ICefTaskRunner Parent: ICefBaseRefCounted
type ICefTaskRunner interface {
	ICefBaseRefCounted
	IsSame(that ICefTaskRunner) bool                   // function
	BelongsToCurrentThread() bool                      // function
	BelongsToThread(threadId TCefThreadId) bool        // function
	PostTask(task ICefTask) bool                       // function
	PostDelayedTask(task ICefTask, delayMs int64) bool // function
}

// TCefTaskRunner Parent: TCefBaseRefCounted
type TCefTaskRunner struct {
	TCefBaseRefCounted
}

// TaskRunnerRef -> ICefTaskRunner
var TaskRunnerRef taskRunner

// taskRunner TCefTaskRunner Ref
type taskRunner uintptr

func (m *taskRunner) UnWrap(data uintptr) ICefTaskRunner {
	var resultCefTaskRunner uintptr
	askRunnerImportAPI().SysCallN(7, uintptr(data), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *taskRunner) GetForCurrentThread() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	askRunnerImportAPI().SysCallN(2, uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *taskRunner) GetForThread(threadId TCefThreadId) ICefTaskRunner {
	var resultCefTaskRunner uintptr
	askRunnerImportAPI().SysCallN(3, uintptr(threadId), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCefTaskRunner) IsSame(that ICefTaskRunner) bool {
	r1 := askRunnerImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefTaskRunner) BelongsToCurrentThread() bool {
	r1 := askRunnerImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCefTaskRunner) BelongsToThread(threadId TCefThreadId) bool {
	r1 := askRunnerImportAPI().SysCallN(1, m.Instance(), uintptr(threadId))
	return GoBool(r1)
}

func (m *TCefTaskRunner) PostTask(task ICefTask) bool {
	r1 := askRunnerImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(task))
	return GoBool(r1)
}

func (m *TCefTaskRunner) PostDelayedTask(task ICefTask, delayMs int64) bool {
	r1 := askRunnerImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(task), uintptr(unsafePointer(&delayMs)))
	return GoBool(r1)
}

var (
	askRunnerImport       *imports.Imports = nil
	askRunnerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefTaskRunner_BelongsToCurrentThread", 0),
		/*1*/ imports.NewTable("CefTaskRunner_BelongsToThread", 0),
		/*2*/ imports.NewTable("CefTaskRunner_GetForCurrentThread", 0),
		/*3*/ imports.NewTable("CefTaskRunner_GetForThread", 0),
		/*4*/ imports.NewTable("CefTaskRunner_IsSame", 0),
		/*5*/ imports.NewTable("CefTaskRunner_PostDelayedTask", 0),
		/*6*/ imports.NewTable("CefTaskRunner_PostTask", 0),
		/*7*/ imports.NewTable("CefTaskRunner_UnWrap", 0),
	}
)

func askRunnerImportAPI() *imports.Imports {
	if askRunnerImport == nil {
		askRunnerImport = NewDefaultImports()
		askRunnerImport.SetImportTable(askRunnerImportTables)
		askRunnerImportTables = nil
	}
	return askRunnerImport
}
