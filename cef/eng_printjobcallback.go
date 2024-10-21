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

// ICefPrintJobCallback Parent: ICefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print job requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_job_callback_t))</a>
type ICefPrintJobCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Indicate completion of the print job.
	Cont() // procedure
}

// TCefPrintJobCallback Parent: TCefBaseRefCounted
//
//	Callback interface for asynchronous continuation of print job requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_job_callback_t))</a>
type TCefPrintJobCallback struct {
	TCefBaseRefCounted
}

// PrintJobCallbackRef -> ICefPrintJobCallback
var PrintJobCallbackRef printJobCallback

// printJobCallback TCefPrintJobCallback Ref
type printJobCallback uintptr

func (m *printJobCallback) UnWrap(data uintptr) ICefPrintJobCallback {
	var resultCefPrintJobCallback uintptr
	printJobCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefPrintJobCallback)))
	return AsCefPrintJobCallback(resultCefPrintJobCallback)
}

func (m *TCefPrintJobCallback) Cont() {
	printJobCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	printJobCallbackImport       *imports.Imports = nil
	printJobCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPrintJobCallback_Cont", 0),
		/*1*/ imports.NewTable("CefPrintJobCallback_UnWrap", 0),
	}
)

func printJobCallbackImportAPI() *imports.Imports {
	if printJobCallbackImport == nil {
		printJobCallbackImport = NewDefaultImports()
		printJobCallbackImport.SetImportTable(printJobCallbackImportTables)
		printJobCallbackImportTables = nil
	}
	return printJobCallbackImport
}
