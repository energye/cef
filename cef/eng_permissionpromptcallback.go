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
//
//	Callback interface used for asynchronous continuation of permission prompts.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_prompt_callback_t))</a>
type ICefPermissionPromptCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Complete the permissions request with the specified |result|.
	Cont(result TCefPermissionRequestResult) // procedure
}

// TCefPermissionPromptCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of permission prompts.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_prompt_callback_t))</a>
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
