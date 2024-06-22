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

// ICefDownloadItemCallback Parent: ICefBaseRefCounted
type ICefDownloadItemCallback interface {
	ICefBaseRefCounted
	Cancel() // procedure
	Pause()  // procedure
	Resume() // procedure
}

// TCefDownloadItemCallback Parent: TCefBaseRefCounted
type TCefDownloadItemCallback struct {
	TCefBaseRefCounted
}

// DownloadItemCallbackRef -> ICefDownloadItemCallback
var DownloadItemCallbackRef downloadItemCallback

// downloadItemCallback TCefDownloadItemCallback Ref
type downloadItemCallback uintptr

func (m *downloadItemCallback) UnWrap(data uintptr) ICefDownloadItemCallback {
	var resultCefDownloadItemCallback uintptr
	downloadItemCallbackImportAPI().SysCallN(3, uintptr(data), uintptr(unsafePointer(&resultCefDownloadItemCallback)))
	return AsCefDownloadItemCallback(resultCefDownloadItemCallback)
}

func (m *TCefDownloadItemCallback) Cancel() {
	downloadItemCallbackImportAPI().SysCallN(0, m.Instance())
}

func (m *TCefDownloadItemCallback) Pause() {
	downloadItemCallbackImportAPI().SysCallN(1, m.Instance())
}

func (m *TCefDownloadItemCallback) Resume() {
	downloadItemCallbackImportAPI().SysCallN(2, m.Instance())
}

var (
	downloadItemCallbackImport       *imports.Imports = nil
	downloadItemCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDownloadItemCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefDownloadItemCallback_Pause", 0),
		/*2*/ imports.NewTable("CefDownloadItemCallback_Resume", 0),
		/*3*/ imports.NewTable("CefDownloadItemCallback_UnWrap", 0),
	}
)

func downloadItemCallbackImportAPI() *imports.Imports {
	if downloadItemCallbackImport == nil {
		downloadItemCallbackImport = NewDefaultImports()
		downloadItemCallbackImport.SetImportTable(downloadItemCallbackImportTables)
		downloadItemCallbackImportTables = nil
	}
	return downloadItemCallbackImport
}
