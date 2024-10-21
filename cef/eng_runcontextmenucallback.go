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
//
//	Callback interface used for continuation of custom context menu display.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_context_menu_callback_t))</a>
type ICefRunContextMenuCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Complete context menu display by selecting the specified |command_id| and |event_flags|.
	Cont(commandId int32, eventFlags TCefEventFlags) // procedure
	// Cancel
	//  Cancel context menu display.
	Cancel() // procedure
}

// TCefRunContextMenuCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for continuation of custom context menu display.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_context_menu_callback_t))</a>
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
