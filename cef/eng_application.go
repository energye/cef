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

// ICefApplication Parent: ICefApplicationCore
//
//	Main class used to simplify the CEF initialization and destruction.
type ICefApplication interface {
	ICefApplicationCore
	DestroyApplicationObject() bool          // property
	SetDestroyApplicationObject(AValue bool) // property
	DestroyAppWindows() bool                 // property
	SetDestroyAppWindows(AValue bool)        // property
}

// TCefApplication Parent: TCefApplicationCore
//
//	Main class used to simplify the CEF initialization and destruction.
type TCefApplication struct {
	TCefApplicationCore
}

func NewCefApplication() ICefApplication {
	r1 := applicationImportAPI().SysCallN(0)
	return AsCefApplication(r1)
}

func (m *TCefApplication) DestroyApplicationObject() bool {
	r1 := applicationImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplication) SetDestroyApplicationObject(AValue bool) {
	applicationImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplication) DestroyAppWindows() bool {
	r1 := applicationImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplication) SetDestroyAppWindows(AValue bool) {
	applicationImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

var (
	applicationImport       *imports.Imports = nil
	applicationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefApplication_Create", 0),
		/*1*/ imports.NewTable("CefApplication_DestroyAppWindows", 0),
		/*2*/ imports.NewTable("CefApplication_DestroyApplicationObject", 0),
	}
)

func applicationImportAPI() *imports.Imports {
	if applicationImport == nil {
		applicationImport = NewDefaultImports()
		applicationImport.SetImportTable(applicationImportTables)
		applicationImportTables = nil
	}
	return applicationImport
}
