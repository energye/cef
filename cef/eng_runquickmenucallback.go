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

// ICefRunQuickMenuCallback Parent: ICefBaseRefCounted
type ICefRunQuickMenuCallback interface {
	ICefBaseRefCounted
	Cont(commandid int32, eventflags TCefEventFlags) // procedure
	Cancel()                                         // procedure
}

// TCefRunQuickMenuCallback Parent: TCefBaseRefCounted
type TCefRunQuickMenuCallback struct {
	TCefBaseRefCounted
}

// RunQuickMenuCallbackRef -> ICefRunQuickMenuCallback
var RunQuickMenuCallbackRef runQuickMenuCallback

// runQuickMenuCallback TCefRunQuickMenuCallback Ref
type runQuickMenuCallback uintptr

func (m *runQuickMenuCallback) UnWrap(data uintptr) ICefRunQuickMenuCallback {
	var resultCefRunQuickMenuCallback uintptr
	runQuickMenuCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefRunQuickMenuCallback)))
	return AsCefRunQuickMenuCallback(resultCefRunQuickMenuCallback)
}

func (m *TCefRunQuickMenuCallback) Cont(commandid int32, eventflags TCefEventFlags) {
	runQuickMenuCallbackImportAPI().SysCallN(1, m.Instance(), uintptr(commandid), uintptr(eventflags))
}

func (m *TCefRunQuickMenuCallback) Cancel() {
	runQuickMenuCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	runQuickMenuCallbackImport       *imports.Imports = nil
	runQuickMenuCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefRunQuickMenuCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefRunQuickMenuCallback_Cont", 0),
		/*2*/ imports.NewTable("CefRunQuickMenuCallback_UnWrap", 0),
	}
)

func runQuickMenuCallbackImportAPI() *imports.Imports {
	if runQuickMenuCallbackImport == nil {
		runQuickMenuCallbackImport = NewDefaultImports()
		runQuickMenuCallbackImport.SetImportTable(runQuickMenuCallbackImportTables)
		runQuickMenuCallbackImportTables = nil
	}
	return runQuickMenuCallbackImport
}
