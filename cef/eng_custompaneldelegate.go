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

// ICustomPanelDelegate Parent: ICefPanelDelegate
//
//	This class handles all the ICefPanelDelegate methods which call the ICefPanelDelegateEvents methods.
//	ICefPanelDelegateEvents will be implemented by the control receiving the ICefPanelDelegate events.
type ICustomPanelDelegate interface {
	ICefPanelDelegate
}

// TCustomPanelDelegate Parent: TCefPanelDelegate
//
//	This class handles all the ICefPanelDelegate methods which call the ICefPanelDelegateEvents methods.
//	ICefPanelDelegateEvents will be implemented by the control receiving the ICefPanelDelegate events.
type TCustomPanelDelegate struct {
	TCefPanelDelegate
}

func NewCustomPanelDelegate(events ICefPanelDelegateEvents) ICustomPanelDelegate {
	r1 := customPanelDelegateImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomPanelDelegate(r1)
}

var (
	customPanelDelegateImport       *imports.Imports = nil
	customPanelDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPanelDelegate_Create", 0),
	}
)

func customPanelDelegateImportAPI() *imports.Imports {
	if customPanelDelegateImport == nil {
		customPanelDelegateImport = NewDefaultImports()
		customPanelDelegateImport.SetImportTable(customPanelDelegateImportTables)
		customPanelDelegateImportTables = nil
	}
	return customPanelDelegateImport
}
