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

// ICefCallback Parent: ICefBaseRefCounted
type ICefCallback interface {
	ICefBaseRefCounted
	Cont()   // procedure
	Cancel() // procedure
}

// TCefCallback Parent: TCefBaseRefCounted
type TCefCallback struct {
	TCefBaseRefCounted
}

// CallbackRef -> ICefCallback
var CallbackRef callback

// callback TCefCallback Ref
type callback uintptr

func (m *callback) UnWrap(data uintptr) ICefCallback {
	var resultCefCallback uintptr
	callbackImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefCallback)))
	return AsCefCallback(resultCefCallback)
}

func (m *TCefCallback) Cont() {
	callbackImportAPI().SysCallN(1, m.Instance())
}

func (m *TCefCallback) Cancel() {
	callbackImportAPI().SysCallN(0, m.Instance())
}

var (
	callbackImport       *imports.Imports = nil
	callbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCallback_Cancel", 0),
		/*1*/ imports.NewTable("CefCallback_Cont", 0),
		/*2*/ imports.NewTable("CefCallback_UnWrap", 0),
	}
)

func callbackImportAPI() *imports.Imports {
	if callbackImport == nil {
		callbackImport = NewDefaultImports()
		callbackImport.SetImportTable(callbackImportTables)
		callbackImportTables = nil
	}
	return callbackImport
}
