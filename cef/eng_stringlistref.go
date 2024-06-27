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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringList // procedure
}

// TCefStringListRef Parent: TCefCustomStringList
type TCefStringListRef struct {
	TCefCustomStringList
}

func NewCefStringListRef(aHandle TCefStringListPtr) ICefStringListRef {
	r1 := stringListRefImportAPI().SysCallN(1, uintptr(aHandle))
	return AsCefStringListRef(r1)
}

func (m *TCefStringListRef) AsInterface() ICefStringList {
	var resultCefStringList uintptr
	stringListRefImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringList)))
	return AsCefStringList(resultCefStringList)
}

var (
	stringListRefImport       *imports.Imports = nil
	stringListRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringListRef_AsInterface", 0),
		/*1*/ imports.NewTable("CefStringListRef_Create", 0),
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
