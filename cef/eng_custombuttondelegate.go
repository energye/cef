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

// ICustomButtonDelegate Parent: ICefButtonDelegate
//
//	This class handles all the ICefButtonDelegate methods which call the ICefButtonDelegateEvents methods.
//	ICefButtonDelegateEvents will be implemented by the control receiving the ICefButtonDelegate events.
type ICustomButtonDelegate interface {
	ICefButtonDelegate
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefButtonDelegate // procedure
}

// TCustomButtonDelegate Parent: TCefButtonDelegate
//
//	This class handles all the ICefButtonDelegate methods which call the ICefButtonDelegateEvents methods.
//	ICefButtonDelegateEvents will be implemented by the control receiving the ICefButtonDelegate events.
type TCustomButtonDelegate struct {
	TCefButtonDelegate
}

func NewCustomButtonDelegate(events ICefButtonDelegateEvents) ICustomButtonDelegate {
	r1 := customButtonDelegateImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomButtonDelegate(r1)
}

func (m *TCustomButtonDelegate) AsInterface() ICefButtonDelegate {
	var resultCefButtonDelegate uintptr
	customButtonDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefButtonDelegate)))
	return AsCefButtonDelegate(resultCefButtonDelegate)
}

var (
	customButtonDelegateImport       *imports.Imports = nil
	customButtonDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomButtonDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("CustomButtonDelegate_Create", 0),
	}
)

func customButtonDelegateImportAPI() *imports.Imports {
	if customButtonDelegateImport == nil {
		customButtonDelegateImport = NewDefaultImports()
		customButtonDelegateImport.SetImportTable(customButtonDelegateImportTables)
		customButtonDelegateImportTables = nil
	}
	return customButtonDelegateImport
}
