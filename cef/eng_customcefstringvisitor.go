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

// ICustomCefStringVisitor Parent: ICefStringVisitor
type ICustomCefStringVisitor interface {
	ICefStringVisitor
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringVisitor // procedure
}

// TCustomCefStringVisitor Parent: TCefStringVisitor
type TCustomCefStringVisitor struct {
	TCefStringVisitor
}

func NewCustomCefStringVisitor(aEvents IChromiumEvents) ICustomCefStringVisitor {
	r1 := customCefStringVisitorImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCustomCefStringVisitor(r1)
}

func (m *TCustomCefStringVisitor) AsInterface() ICefStringVisitor {
	var resultCefStringVisitor uintptr
	customCefStringVisitorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringVisitor)))
	return AsCefStringVisitor(resultCefStringVisitor)
}

var (
	customCefStringVisitorImport       *imports.Imports = nil
	customCefStringVisitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCefStringVisitor_AsInterface", 0),
		/*1*/ imports.NewTable("CustomCefStringVisitor_Create", 0),
	}
)

func customCefStringVisitorImportAPI() *imports.Imports {
	if customCefStringVisitorImport == nil {
		customCefStringVisitorImport = NewDefaultImports()
		customCefStringVisitorImport.SetImportTable(customCefStringVisitorImportTables)
		customCefStringVisitorImportTables = nil
	}
	return customCefStringVisitorImport
}
