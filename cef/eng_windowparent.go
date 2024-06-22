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

// ICEFWindowParent Parent: ICEFWinControl
type ICEFWindowParent interface {
	ICEFWinControl
}

// TCEFWindowParent Parent: TCEFWinControl
type TCEFWindowParent struct {
	TCEFWinControl
}

func NewCEFWindowParent(theOwner IComponent) ICEFWindowParent {
	r1 := windowParentImportAPI().SysCallN(0, GetObjectUintptr(theOwner))
	return AsCEFWindowParent(r1)
}

var (
	windowParentImport       *imports.Imports = nil
	windowParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFWindowParent_Create", 0),
	}
)

func windowParentImportAPI() *imports.Imports {
	if windowParentImport == nil {
		windowParentImport = NewDefaultImports()
		windowParentImport.SetImportTable(windowParentImportTables)
		windowParentImportTables = nil
	}
	return windowParentImport
}
