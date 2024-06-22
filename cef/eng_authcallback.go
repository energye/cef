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

// ICefAuthCallback Parent: ICefBaseRefCounted
type ICefAuthCallback interface {
	ICefBaseRefCounted
	Cont(username, password string) // procedure
	Cancel()                        // procedure
}

// TCefAuthCallback Parent: TCefBaseRefCounted
type TCefAuthCallback struct {
	TCefBaseRefCounted
}

// AuthCallbackRef -> ICefAuthCallback
var AuthCallbackRef authCallback

// authCallback TCefAuthCallback Ref
type authCallback uintptr

func (m *authCallback) UnWrap(data uintptr) ICefAuthCallback {
	var resultCefAuthCallback uintptr
	authCallbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefAuthCallback)))
	return AsCefAuthCallback(resultCefAuthCallback)
}

func (m *TCefAuthCallback) Cont(username, password string) {
	authCallbackImportAPI().SysCallN(1, m.Instance(), PascalStr(username), PascalStr(password))
}

func (m *TCefAuthCallback) Cancel() {
	authCallbackImportAPI().SysCallN(0, m.Instance())
}

var (
	authCallbackImport       *imports.Imports = nil
	authCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefAuthCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefAuthCallback_Cont", 0),
		/*2*/ imports.NewTable("CefAuthCallback_UnWrap", 0),
	}
)

func authCallbackImportAPI() *imports.Imports {
	if authCallbackImport == nil {
		authCallbackImport = NewDefaultImports()
		authCallbackImport.SetImportTable(authCallbackImportTables)
		authCallbackImportTables = nil
	}
	return authCallbackImport
}
