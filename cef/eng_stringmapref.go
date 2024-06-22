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
}

// TCefStringMapRef Parent: TCefStringMap
type TCefStringMapRef struct {
	TCefStringMap
}

func NewCefStringMapRef(aHandle TCefStringMapHandle) ICefStringMapRef {
	r1 := stringMapRefImportAPI().SysCallN(0, uintptr(aHandle))
	return AsCefStringMapRef(r1)
}

var (
	stringMapRefImport       *imports.Imports = nil
	stringMapRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMapRef_Create", 0),
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
