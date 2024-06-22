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

// ICefCustomCookieVisitor Parent: ICefCookieVisitor
type ICefCustomCookieVisitor interface {
	ICefCookieVisitor
}

// TCefCustomCookieVisitor Parent: TCefCookieVisitor
type TCefCustomCookieVisitor struct {
	TCefCookieVisitor
}

func NewCefCustomCookieVisitor(aEvents IChromiumEvents, aID int32) ICefCustomCookieVisitor {
	r1 := customCookieVisitorImportAPI().SysCallN(0, GetObjectUintptr(aEvents), uintptr(aID))
	return AsCefCustomCookieVisitor(r1)
}

var (
	customCookieVisitorImport       *imports.Imports = nil
	customCookieVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomCookieVisitor_Create", 0),
	}
)

func customCookieVisitorImportAPI() *imports.Imports {
	if customCookieVisitorImport == nil {
		customCookieVisitorImport = NewDefaultImports()
		customCookieVisitorImport.SetImportTable(customCookieVisitorImportTables)
		customCookieVisitorImportTables = nil
	}
	return customCookieVisitorImport
}
