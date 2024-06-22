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

// ICefRunContextMenuCallback Parent: ICefBaseRefCounted
type ICefRunContextMenuCallback interface {
	ICefBaseRefCounted
	Cont(commandId int32, eventFlags TCefEventFlags) // procedure
	Cancel()                                         // procedure
}

// TCefRunContextMenuCallback Parent: TCefBaseRefCounted
type TCefRunContextMenuCallback struct {
	TCefBaseRefCounted
}

// RunContextMenuCallbackRef -> ICefRunContextMenuCallback
var RunContextMenuCallbackRef runContextMenuCallback

// runContextMenuCallback TCefRunContextMenuCallback Ref
type runContextMenuCallback uintptr

func (m *runContextMenuCallback) UnWrap(data uintptr) ICefRunContextMenuCallback {
	var resultCefRunContextMenuCallback uintptr
	runContextMenuCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefRunContextMenuCallback)))
	return AsCefRunContextMenuCallback(resultCefRunContextMenuCallback)
}

func (m *TCefRunContextMenuCallback) Cont(commandId int32, eventFlags TCefEventFlags) {
	runContextMenuCallbackImportAPI().SysCallN(1, m.Instance(), uintptr(commandId), uintptr(eventFlags))
}

func (m *TCefRunContextMenuCallback) Cancel() {
	runContextMenuCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	runContextMenuCallbackImport       *imports.Imports = nil
	runContextMenuCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefRunContextMenuCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefRunContextMenuCallback_Cont", 0),
		/*2*/ imports.NewTable("CefRunContextMenuCallback_UnWrap", 0),
	}
)

func runContextMenuCallbackImportAPI() *imports.Imports {
	if runContextMenuCallbackImport == nil {
		runContextMenuCallbackImport = NewDefaultImports()
		runContextMenuCallbackImport.SetImportTable(runContextMenuCallbackImportTables)
		runContextMenuCallbackImportTables = nil
	}
	return runContextMenuCallbackImport
}
