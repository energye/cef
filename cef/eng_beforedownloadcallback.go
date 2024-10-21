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
//
//	Callback interface used to asynchronously continue a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_before_download_callback_t))</a>
type ICefBeforeDownloadCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Call to continue the download. Set |download_path| to the full file path for the download including the file name or leave blank to use the suggested name and the default temp directory. Set |show_dialog| to true (1) if you do wish to show the default "Save As" dialog.
	Cont(downloadPath string, showDialog bool) // procedure
}

// TCefBeforeDownloadCallback Parent: TCefBaseRefCounted
//
//	Callback interface used to asynchronously continue a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_before_download_callback_t))</a>
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
