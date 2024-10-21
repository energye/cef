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

// ICefPrintDialogCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_dialog_callback_t))</a>
type ICefPrintDialogCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue printing with the specified |settings|.
	Cont(settings ICefPrintSettings) // procedure
	// Cancel
	//  Cancel the printing.
	Cancel() // procedure
}

// TCefPrintDialogCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print dialog requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_dialog_callback_t))</a>
type TCefPrintDialogCallback struct {
	TCefBaseRefCounted
}

// PrintDialogCallbackRef -> ICefPrintDialogCallback
var PrintDialogCallbackRef printDialogCallback

// printDialogCallback TCefPrintDialogCallback Ref
type printDialogCallback uintptr

func (m *printDialogCallback) UnWrap(data uintptr) ICefPrintDialogCallback {
	var resultCefPrintDialogCallback uintptr
	printDialogCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefPrintDialogCallback)))
	return AsCefPrintDialogCallback(resultCefPrintDialogCallback)
}

func (m *TCefPrintDialogCallback) Cont(settings ICefPrintSettings) {
	printDialogCallbackImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(settings))
}

func (m *TCefPrintDialogCallback) Cancel() {
	printDialogCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	printDialogCallbackImport       *imports.Imports = nil
	printDialogCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPrintDialogCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefPrintDialogCallback_Cont", 0),
		/*2*/ imports.NewTable("CefPrintDialogCallback_UnWrap", 0),
	}
)

func printDialogCallbackImportAPI() *imports.Imports {
	if printDialogCallbackImport == nil {
		printDialogCallbackImport = NewDefaultImports()
		printDialogCallbackImport.SetImportTable(printDialogCallbackImportTables)
		printDialogCallbackImportTables = nil
	}
	return printDialogCallbackImport
}
