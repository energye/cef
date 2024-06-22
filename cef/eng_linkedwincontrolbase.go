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

// ICEFLinkedWinControlBase Is Abstract Class Parent: ICEFWinControl
//
//	TCEFLinkedWinControlBase is a custom TWinControl to host the child controls created by the web browser
//	to show the web contents and it's linked to the TChromium instance that handles that web browser.
//	TCEFLinkedWinControlBase is the parent class of TChromiumWindow, TBrowserWindow and TCEFLinkedWindowParent.
type ICEFLinkedWinControlBase interface {
	ICEFWinControl
	// UseSetFocus
	//  Use TChromium.SetFocus when the component receives a WM_SETFOCUS message in Windows.
	UseSetFocus() bool // property
	// SetUseSetFocus Set UseSetFocus
	SetUseSetFocus(AValue bool) // property
}

// TCEFLinkedWinControlBase Is Abstract Class Parent: TCEFWinControl
//
//	TCEFLinkedWinControlBase is a custom TWinControl to host the child controls created by the web browser
//	to show the web contents and it's linked to the TChromium instance that handles that web browser.
//	TCEFLinkedWinControlBase is the parent class of TChromiumWindow, TBrowserWindow and TCEFLinkedWindowParent.
type TCEFLinkedWinControlBase struct {
	TCEFWinControl
}

func (m *TCEFLinkedWinControlBase) UseSetFocus() bool {
	r1 := linkedWinControlBaseImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFLinkedWinControlBase) SetUseSetFocus(AValue bool) {
	linkedWinControlBaseImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

var (
	linkedWinControlBaseImport       *imports.Imports = nil
	linkedWinControlBaseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFLinkedWinControlBase_UseSetFocus", 0),
	}
)

func linkedWinControlBaseImportAPI() *imports.Imports {
	if linkedWinControlBaseImport == nil {
		linkedWinControlBaseImport = NewDefaultImports()
		linkedWinControlBaseImport.SetImportTable(linkedWinControlBaseImportTables)
		linkedWinControlBaseImportTables = nil
	}
	return linkedWinControlBaseImport
}
