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

// ICefStringListRef Parent: ICefCustomStringList
type ICefStringListRef interface {
	ICefCustomStringList
}

// TCefStringListRef Parent: TCefCustomStringList
type TCefStringListRef struct {
	TCefCustomStringList
}

func NewCefStringListRef(aHandle TCefStringList) ICefStringListRef {
	r1 := stringListRefImportAPI().SysCallN(0, uintptr(aHandle))
	return AsCefStringListRef(r1)
}

var (
	stringListRefImport       *imports.Imports = nil
	stringListRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringListRef_Create", 0),
	}
)

func stringListRefImportAPI() *imports.Imports {
	if stringListRefImport == nil {
		stringListRefImport = NewDefaultImports()
		stringListRefImport.SetImportTable(stringListRefImportTables)
		stringListRefImportTables = nil
	}
	return stringListRefImport
}
