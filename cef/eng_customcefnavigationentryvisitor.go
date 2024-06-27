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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefNavigationEntryVisitor // procedure
}

// TCustomCefNavigationEntryVisitor Parent: TCefNavigationEntryVisitor
type TCustomCefNavigationEntryVisitor struct {
	TCefNavigationEntryVisitor
}

func NewCustomCefNavigationEntryVisitor(aEvents IChromiumEvents) ICustomCefNavigationEntryVisitor {
	r1 := customCefNavigationEntryVisitorImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCustomCefNavigationEntryVisitor(r1)
}

func (m *TCustomCefNavigationEntryVisitor) AsInterface() ICefNavigationEntryVisitor {
	var resultCefNavigationEntryVisitor uintptr
	customCefNavigationEntryVisitorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefNavigationEntryVisitor)))
	return AsCefNavigationEntryVisitor(resultCefNavigationEntryVisitor)
}

var (
	customCefNavigationEntryVisitorImport       *imports.Imports = nil
	customCefNavigationEntryVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCefNavigationEntryVisitor_AsInterface", 0),
		/*1*/ imports.NewTable("CustomCefNavigationEntryVisitor_Create", 0),
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
