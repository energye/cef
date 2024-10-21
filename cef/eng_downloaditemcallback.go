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
//
//	Callback interface used to asynchronously cancel a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_item_callback_t))</a>
type ICefDownloadItemCallback interface {
	ICefBaseRefCounted
	// Cancel
	//  Call to cancel the download.
	Cancel() // procedure
	// Pause
	//  Call to pause the download.
	Pause() // procedure
	// Resume
	//  Call to resume the download.
	Resume() // procedure
}

// TCefDownloadItemCallback Parent: TCefBaseRefCounted
//
//	Callback interface used to asynchronously cancel a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_item_callback_t))</a>
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
