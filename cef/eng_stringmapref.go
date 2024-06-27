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

// ICefStringMapRef Parent: ICefStringMap
type ICefStringMapRef interface {
	ICefStringMap
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringMap // procedure
}

// TCefStringMapRef Parent: TCefStringMap
type TCefStringMapRef struct {
	TCefStringMap
}

func NewCefStringMapRef(aHandle TCefStringMapHandle) ICefStringMapRef {
	r1 := stringMapRefImportAPI().SysCallN(1, uintptr(aHandle))
	return AsCefStringMapRef(r1)
}

func (m *TCefStringMapRef) AsInterface() ICefStringMap {
	var resultCefStringMap uintptr
	stringMapRefImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringMap)))
	return AsCefStringMap(resultCefStringMap)
}

var (
	stringMapRefImport       *imports.Imports = nil
	stringMapRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMapRef_AsInterface", 0),
		/*1*/ imports.NewTable("CefStringMapRef_Create", 0),
	}
)

func stringMapRefImportAPI() *imports.Imports {
	if stringMapRefImport == nil {
		stringMapRefImport = NewDefaultImports()
		stringMapRefImport.SetImportTable(stringMapRefImportTables)
		stringMapRefImportTables = nil
	}
	return stringMapRefImport
}
