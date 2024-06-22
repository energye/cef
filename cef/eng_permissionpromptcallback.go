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

// ICefPermissionPromptCallback Parent: ICefBaseRefCounted
type ICefPermissionPromptCallback interface {
	ICefBaseRefCounted
	Cont(result TCefPermissionRequestResult) // procedure
}

// TCefPermissionPromptCallback Parent: TCefBaseRefCounted
type TCefPermissionPromptCallback struct {
	TCefBaseRefCounted
}

// PermissionPromptCallbackRef -> ICefPermissionPromptCallback
var PermissionPromptCallbackRef permissionPromptCallback

// permissionPromptCallback TCefPermissionPromptCallback Ref
type permissionPromptCallback uintptr

func (m *permissionPromptCallback) UnWrap(data uintptr) ICefPermissionPromptCallback {
	var resultCefPermissionPromptCallback uintptr
	permissionPromptCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefPermissionPromptCallback)))
	return AsCefPermissionPromptCallback(resultCefPermissionPromptCallback)
}

func (m *TCefPermissionPromptCallback) Cont(result TCefPermissionRequestResult) {
	permissionPromptCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(result))
}

var (
	permissionPromptCallbackImport       *imports.Imports = nil
	permissionPromptCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPermissionPromptCallback_Cont", 0),
		/*1*/ imports.NewTable("CefPermissionPromptCallback_UnWrap", 0),
	}
)

func permissionPromptCallbackImportAPI() *imports.Imports {
	if permissionPromptCallbackImport == nil {
		permissionPromptCallbackImport = NewDefaultImports()
		permissionPromptCallbackImport.SetImportTable(permissionPromptCallbackImportTables)
		permissionPromptCallbackImportTables = nil
	}
	return permissionPromptCallbackImport
}
