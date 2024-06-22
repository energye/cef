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

// ICefCustomDeleteCookiesCallback Parent: ICefDeleteCookiesCallback
type ICefCustomDeleteCookiesCallback interface {
	ICefDeleteCookiesCallback
}

// TCefCustomDeleteCookiesCallback Parent: TCefDeleteCookiesCallback
type TCefCustomDeleteCookiesCallback struct {
	TCefDeleteCookiesCallback
}

func NewCefCustomDeleteCookiesCallback(aEvents IChromiumEvents) ICefCustomDeleteCookiesCallback {
	r1 := customDeleteCookiesCallbackImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefCustomDeleteCookiesCallback(r1)
}

var (
	customDeleteCookiesCallbackImport       *imports.Imports = nil
	customDeleteCookiesCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomDeleteCookiesCallback_Create", 0),
	}
)

func customDeleteCookiesCallbackImportAPI() *imports.Imports {
	if customDeleteCookiesCallbackImport == nil {
		customDeleteCookiesCallbackImport = NewDefaultImports()
		customDeleteCookiesCallbackImport.SetImportTable(customDeleteCookiesCallbackImportTables)
		customDeleteCookiesCallbackImportTables = nil
	}
	return customDeleteCookiesCallbackImport
}
