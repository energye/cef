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

// ICefCustomDownloadImageCallback Parent: ICefDownloadImageCallback
type ICefCustomDownloadImageCallback interface {
	ICefDownloadImageCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDownloadImageCallback // procedure
}

// TCefCustomDownloadImageCallback Parent: TCefDownloadImageCallback
type TCefCustomDownloadImageCallback struct {
	TCefDownloadImageCallback
}

func NewCefCustomDownloadImageCallback(aEvents IChromiumEvents) ICefCustomDownloadImageCallback {
	r1 := customDownloadImageCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomDownloadImageCallback(r1)
}

func (m *TCefCustomDownloadImageCallback) AsInterface() ICefDownloadImageCallback {
	var resultCefDownloadImageCallback uintptr
	customDownloadImageCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDownloadImageCallback)))
	return AsCefDownloadImageCallback(resultCefDownloadImageCallback)
}

var (
	customDownloadImageCallbackImport       *imports.Imports = nil
	customDownloadImageCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomDownloadImageCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomDownloadImageCallback_Create", 0),
	}
)

func customDownloadImageCallbackImportAPI() *imports.Imports {
	if customDownloadImageCallbackImport == nil {
		customDownloadImageCallbackImport = NewDefaultImports()
		customDownloadImageCallbackImport.SetImportTable(customDownloadImageCallbackImportTables)
		customDownloadImageCallbackImportTables = nil
	}
	return customDownloadImageCallbackImport
}
