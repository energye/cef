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

// ICefCustomSetCookieCallback Parent: ICefSetCookieCallback
type ICefCustomSetCookieCallback interface {
	ICefSetCookieCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefSetCookieCallback // procedure
}

// TCefCustomSetCookieCallback Parent: TCefSetCookieCallback
type TCefCustomSetCookieCallback struct {
	TCefSetCookieCallback
}

func NewCefCustomSetCookieCallback(aEvents IChromiumEvents, aID int32) ICefCustomSetCookieCallback {
	r1 := customSetCookieCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomSetCookieCallback(r1)
}

func (m *TCefCustomSetCookieCallback) AsInterface() ICefSetCookieCallback {
	var resultCefSetCookieCallback uintptr
	customSetCookieCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefSetCookieCallback)))
	return AsCefSetCookieCallback(resultCefSetCookieCallback)
}

var (
	customSetCookieCallbackImport       *imports.Imports = nil
	customSetCookieCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomSetCookieCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomSetCookieCallback_Create", 0),
	}
)

func customSetCookieCallbackImportAPI() *imports.Imports {
	if customSetCookieCallbackImport == nil {
		customSetCookieCallbackImport = NewDefaultImports()
		customSetCookieCallbackImport.SetImportTable(customSetCookieCallbackImportTables)
		customSetCookieCallbackImportTables = nil
	}
	return customSetCookieCallbackImport
}
