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

// ICefCustomStringList Parent: IObject
//
//	CEF string maps are a set of key/value string pairs.
type ICefCustomStringList interface {
	IObject
}

// TCefCustomStringList Parent: TObject
//
//	CEF string maps are a set of key/value string pairs.
type TCefCustomStringList struct {
	TObject
}

func NewCefCustomStringList() ICefCustomStringList {
	r1 := customStringListImportAPI().SysCallN(0)
	return AsCefCustomStringList(r1)
}

var (
	customStringListImport       *imports.Imports = nil
	customStringListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomStringList_Create", 0),
	}
)

func customStringListImportAPI() *imports.Imports {
	if customStringListImport == nil {
		customStringListImport = NewDefaultImports()
		customStringListImport.SetImportTable(customStringListImportTables)
		customStringListImportTables = nil
	}
	return customStringListImport
}
