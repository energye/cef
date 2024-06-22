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

// ICustomCefNavigationEntryVisitor Parent: ICefNavigationEntryVisitor
type ICustomCefNavigationEntryVisitor interface {
	ICefNavigationEntryVisitor
}

// TCustomCefNavigationEntryVisitor Parent: TCefNavigationEntryVisitor
type TCustomCefNavigationEntryVisitor struct {
	TCefNavigationEntryVisitor
}

func NewCustomCefNavigationEntryVisitor(aEvents IChromiumEvents) ICustomCefNavigationEntryVisitor {
	r1 := customCefNavigationEntryVisitorImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCustomCefNavigationEntryVisitor(r1)
}

var (
	customCefNavigationEntryVisitorImport       *imports.Imports = nil
	customCefNavigationEntryVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCefNavigationEntryVisitor_Create", 0),
	}
)

func customCefNavigationEntryVisitorImportAPI() *imports.Imports {
	if customCefNavigationEntryVisitorImport == nil {
		customCefNavigationEntryVisitorImport = NewDefaultImports()
		customCefNavigationEntryVisitorImport.SetImportTable(customCefNavigationEntryVisitorImportTables)
		customCefNavigationEntryVisitorImportTables = nil
	}
	return customCefNavigationEntryVisitorImport
}
