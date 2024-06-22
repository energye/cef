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
type ICefFileDialogCallback interface {
	ICefBaseRefCounted
	Cont(filePaths IStrings) // procedure
	Cancel()                 // procedure
}

// TCefFileDialogCallback Parent: TCefBaseRefCounted
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
