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
}

// TCefCustomSetCookieCallback Parent: TCefSetCookieCallback
type TCefCustomSetCookieCallback struct {
	TCefSetCookieCallback
}

func NewCefCustomSetCookieCallback(aEvents IChromiumEvents, aID int32) ICefCustomSetCookieCallback {
	r1 := customSetCookieCallbackImportAPI().SysCallN(0, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomSetCookieCallback(r1)
}

var (
	customSetCookieCallbackImport       *imports.Imports = nil
	customSetCookieCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomSetCookieCallback_Create", 0),
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
