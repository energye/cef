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

// ICefBeforeDownloadCallback Parent: ICefBaseRefCounted
type ICefBeforeDownloadCallback interface {
	ICefBaseRefCounted
	Cont(downloadPath string, showDialog bool) // procedure
}

// TCefBeforeDownloadCallback Parent: TCefBaseRefCounted
type TCefBeforeDownloadCallback struct {
	TCefBaseRefCounted
}

// BeforeDownloadCallbackRef -> ICefBeforeDownloadCallback
var BeforeDownloadCallbackRef beforeDownloadCallback

// beforeDownloadCallback TCefBeforeDownloadCallback Ref
type beforeDownloadCallback uintptr

func (m *beforeDownloadCallback) UnWrap(data uintptr) ICefBeforeDownloadCallback {
	var resultCefBeforeDownloadCallback uintptr
	beforeDownloadCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefBeforeDownloadCallback)))
	return AsCefBeforeDownloadCallback(resultCefBeforeDownloadCallback)
}

func (m *TCefBeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	beforeDownloadCallbackImportAPI().SysCallN(0, m.Instance(), PascalStr(downloadPath), PascalBool(showDialog))
}

var (
	beforeDownloadCallbackImport       *imports.Imports = nil
	beforeDownloadCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBeforeDownloadCallback_Cont", 0),
		/*1*/ imports.NewTable("CefBeforeDownloadCallback_UnWrap", 0),
	}
)

func beforeDownloadCallbackImportAPI() *imports.Imports {
	if beforeDownloadCallbackImport == nil {
		beforeDownloadCallbackImport = NewDefaultImports()
		beforeDownloadCallbackImport.SetImportTable(beforeDownloadCallbackImportTables)
		beforeDownloadCallbackImportTables = nil
	}
	return beforeDownloadCallbackImport
}
