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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringList // procedure
}

// TCefCustomStringList Parent: TObject
//
//	CEF string maps are a set of key/value string pairs.
type TCefCustomStringList struct {
	TObject
}

func NewCefCustomStringList() ICefCustomStringList {
	r1 := customStringListImportAPI().SysCallN(1)
	return AsCefCustomStringList(r1)
}

func (m *TCefCustomStringList) AsInterface() ICefStringList {
	var resultCefStringList uintptr
	customStringListImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringList)))
	return AsCefStringList(resultCefStringList)
}

var (
	customStringListImport       *imports.Imports = nil
	customStringListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomStringList_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomStringList_Create", 0),
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
