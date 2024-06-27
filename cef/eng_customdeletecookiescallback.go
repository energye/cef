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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDeleteCookiesCallback // procedure
}

// TCefCustomDeleteCookiesCallback Parent: TCefDeleteCookiesCallback
type TCefCustomDeleteCookiesCallback struct {
	TCefDeleteCookiesCallback
}

func NewCefCustomDeleteCookiesCallback(aEvents IChromiumEvents) ICefCustomDeleteCookiesCallback {
	r1 := customDeleteCookiesCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomDeleteCookiesCallback(r1)
}

func (m *TCefCustomDeleteCookiesCallback) AsInterface() ICefDeleteCookiesCallback {
	var resultCefDeleteCookiesCallback uintptr
	customDeleteCookiesCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDeleteCookiesCallback)))
	return AsCefDeleteCookiesCallback(resultCefDeleteCookiesCallback)
}

var (
	customDeleteCookiesCallbackImport       *imports.Imports = nil
	customDeleteCookiesCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomDeleteCookiesCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomDeleteCookiesCallback_Create", 0),
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
