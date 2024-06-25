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

// IRunFileDialogCallback Parent: ICefRunFileDialogCallback
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of
//	this interface will be called on the browser process UI thread.
//	<a cref="uCEFTypes|TCefRunFileDialogCallback">Implements TCefRunFileDialogCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t)</a>
type IRunFileDialogCallback interface {
	ICefRunFileDialogCallback
	AsInterface() ICefRunFileDialogCallback // function
	// SetOnRunFileDialogDismissed
	//  Called asynchronously after the file dialog is dismissed. |file_paths|
	//  will be a single value or a list of values depending on the dialog mode.
	//  If the selection was cancelled |file_paths| will be NULL.
	SetOnRunFileDialogDismissed(fn TOnRunFileDialogDismissed) // property event
}

// TRunFileDialogCallback Parent: TCefRunFileDialogCallback
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of
//	this interface will be called on the browser process UI thread.
//	<a cref="uCEFTypes|TCefRunFileDialogCallback">Implements TCefRunFileDialogCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t)</a>
type TRunFileDialogCallback struct {
	TCefRunFileDialogCallback
	runFileDialogDismissedPtr uintptr
}

func NewRunFileDialogCallback() IRunFileDialogCallback {
	r1 := runFileDialogCallbackImportAPI().SysCallN(1)
	return AsRunFileDialogCallback(r1)
}

func (m *TRunFileDialogCallback) AsInterface() ICefRunFileDialogCallback {
	var resultCefRunFileDialogCallback uintptr
	runFileDialogCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefRunFileDialogCallback)))
	return AsCefRunFileDialogCallback(resultCefRunFileDialogCallback)
}

func (m *TRunFileDialogCallback) SetOnRunFileDialogDismissed(fn TOnRunFileDialogDismissed) {
	if m.runFileDialogDismissedPtr != 0 {
		RemoveEventElement(m.runFileDialogDismissedPtr)
	}
	m.runFileDialogDismissedPtr = MakeEventDataPtr(fn)
	runFileDialogCallbackImportAPI().SysCallN(2, m.Instance(), m.runFileDialogDismissedPtr)
}

var (
	runFileDialogCallbackImport       *imports.Imports = nil
	runFileDialogCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RunFileDialogCallback_AsInterface", 0),
		/*1*/ imports.NewTable("RunFileDialogCallback_Create", 0),
		/*2*/ imports.NewTable("RunFileDialogCallback_SetOnRunFileDialogDismissed", 0),
	}
)

func runFileDialogCallbackImportAPI() *imports.Imports {
	if runFileDialogCallbackImport == nil {
		runFileDialogCallbackImport = NewDefaultImports()
		runFileDialogCallbackImport.SetImportTable(runFileDialogCallbackImportTables)
		runFileDialogCallbackImportTables = nil
	}
	return runFileDialogCallbackImport
}
