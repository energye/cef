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

// ICustomMenuButtonDelegate Parent: ICefMenuButtonDelegate
//
//	This class handles all the ICefMenuButtonDelegate methods which call the ICefMenuButtonDelegateEvents methods.
//	ICefMenuButtonDelegateEvents will be implemented by the control receiving the ICefMenuButtonDelegate events.
type ICustomMenuButtonDelegate interface {
	ICefMenuButtonDelegate
}

// TCustomMenuButtonDelegate Parent: TCefMenuButtonDelegate
//
//	This class handles all the ICefMenuButtonDelegate methods which call the ICefMenuButtonDelegateEvents methods.
//	ICefMenuButtonDelegateEvents will be implemented by the control receiving the ICefMenuButtonDelegate events.
type TCustomMenuButtonDelegate struct {
	TCefMenuButtonDelegate
}

func NewCustomMenuButtonDelegate(events ICefMenuButtonDelegateEvents) ICustomMenuButtonDelegate {
	r1 := customMenuButtonDelegateImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomMenuButtonDelegate(r1)
}

var (
	customMenuButtonDelegateImport       *imports.Imports = nil
	customMenuButtonDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMenuButtonDelegate_Create", 0),
	}
)

func customMenuButtonDelegateImportAPI() *imports.Imports {
	if customMenuButtonDelegateImport == nil {
		customMenuButtonDelegateImport = NewDefaultImports()
		customMenuButtonDelegateImport.SetImportTable(customMenuButtonDelegateImportTables)
		customMenuButtonDelegateImportTables = nil
	}
	return customMenuButtonDelegateImport
}
