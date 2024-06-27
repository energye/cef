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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefMenuButtonDelegate // procedure
}

// TCustomMenuButtonDelegate Parent: TCefMenuButtonDelegate
//
//	This class handles all the ICefMenuButtonDelegate methods which call the ICefMenuButtonDelegateEvents methods.
//	ICefMenuButtonDelegateEvents will be implemented by the control receiving the ICefMenuButtonDelegate events.
type TCustomMenuButtonDelegate struct {
	TCefMenuButtonDelegate
}

func NewCustomMenuButtonDelegate(events ICefMenuButtonDelegateEvents) ICustomMenuButtonDelegate {
	r1 := customMenuButtonDelegateImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomMenuButtonDelegate(r1)
}

func (m *TCustomMenuButtonDelegate) AsInterface() ICefMenuButtonDelegate {
	var resultCefMenuButtonDelegate uintptr
	customMenuButtonDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMenuButtonDelegate)))
	return AsCefMenuButtonDelegate(resultCefMenuButtonDelegate)
}

var (
	customMenuButtonDelegateImport       *imports.Imports = nil
	customMenuButtonDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMenuButtonDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("CustomMenuButtonDelegate_Create", 0),
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
