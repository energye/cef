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

// ICefMediaAccessCallback Parent: ICefBaseRefCounted
type ICefMediaAccessCallback interface {
	ICefBaseRefCounted
	Cont(allowedpermissions TCefMediaAccessPermissionTypes) // procedure
	Cancel()                                                // procedure
}

// TCefMediaAccessCallback Parent: TCefBaseRefCounted
type TCefMediaAccessCallback struct {
	TCefBaseRefCounted
}

// MediaAccessCallbackRef -> ICefMediaAccessCallback
var MediaAccessCallbackRef mediaAccessCallback

// mediaAccessCallback TCefMediaAccessCallback Ref
type mediaAccessCallback uintptr

func (m *mediaAccessCallback) UnWrap(data uintptr) ICefMediaAccessCallback {
	var resultCefMediaAccessCallback uintptr
	mediaAccessCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefMediaAccessCallback)))
	return AsCefMediaAccessCallback(resultCefMediaAccessCallback)
}

func (m *TCefMediaAccessCallback) Cont(allowedpermissions TCefMediaAccessPermissionTypes) {
	mediaAccessCallbackImportAPI().SysCallN(1, m.Instance(), uintptr(allowedpermissions))
}

func (m *TCefMediaAccessCallback) Cancel() {
	mediaAccessCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	mediaAccessCallbackImport       *imports.Imports = nil
	mediaAccessCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMediaAccessCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefMediaAccessCallback_Cont", 0),
		/*2*/ imports.NewTable("CefMediaAccessCallback_UnWrap", 0),
	}
)

func mediaAccessCallbackImportAPI() *imports.Imports {
	if mediaAccessCallbackImport == nil {
		mediaAccessCallbackImport = NewDefaultImports()
		mediaAccessCallbackImport.SetImportTable(mediaAccessCallbackImportTables)
		mediaAccessCallbackImportTables = nil
	}
	return mediaAccessCallbackImport
}
