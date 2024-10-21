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

// ICefFileDialogCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of file dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_file_dialog_callback_t))</a>
type ICefFileDialogCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the file selection. |file_paths| should be a single value or a list of values depending on the dialog mode. An NULL |file_paths| value is treated the same as calling cancel().
	Cont(filePaths IStrings) // procedure
	// Cancel
	//  Cancel the file selection.
	Cancel() // procedure
}

// TCefFileDialogCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of file dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_file_dialog_callback_t))</a>
type TCefFileDialogCallback struct {
	TCefBaseRefCounted
}

// FileDialogCallbackRef -> ICefFileDialogCallback
var FileDialogCallbackRef fileDialogCallback

// fileDialogCallback TCefFileDialogCallback Ref
type fileDialogCallback uintptr

func (m *fileDialogCallback) UnWrap(data uintptr) ICefFileDialogCallback {
	var resultCefFileDialogCallback uintptr
	fileDialogCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefFileDialogCallback)))
	return AsCefFileDialogCallback(resultCefFileDialogCallback)
}

func (m *TCefFileDialogCallback) Cont(filePaths IStrings) {
	fileDialogCallbackImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(filePaths))
}

func (m *TCefFileDialogCallback) Cancel() {
	fileDialogCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	fileDialogCallbackImport       *imports.Imports = nil
	fileDialogCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefFileDialogCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefFileDialogCallback_Cont", 0),
		/*2*/ imports.NewTable("CefFileDialogCallback_UnWrap", 0),
	}
)

func fileDialogCallbackImportAPI() *imports.Imports {
	if fileDialogCallbackImport == nil {
		fileDialogCallbackImport = NewDefaultImports()
		fileDialogCallbackImport.SetImportTable(fileDialogCallbackImportTables)
		fileDialogCallbackImportTables = nil
	}
	return fileDialogCallbackImport
}
