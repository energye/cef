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

// ICustomTextfieldDelegate Parent: ICefTextfieldDelegate
//
//	This class handles all the ICefTextfieldDelegate and ICefViewDelegate methods which call the ICefTextfieldDelegateEvents methods.
//	ICefTextfieldDelegateEvents will be implemented by the control receiving the ICefTextfieldDelegate events.
type ICustomTextfieldDelegate interface {
	ICefTextfieldDelegate
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTextfieldDelegate // procedure
}

// TCustomTextfieldDelegate Parent: TCefTextfieldDelegate
//
//	This class handles all the ICefTextfieldDelegate and ICefViewDelegate methods which call the ICefTextfieldDelegateEvents methods.
//	ICefTextfieldDelegateEvents will be implemented by the control receiving the ICefTextfieldDelegate events.
type TCustomTextfieldDelegate struct {
	TCefTextfieldDelegate
}

func NewCustomTextfieldDelegate(events ICefTextfieldDelegateEvents) ICustomTextfieldDelegate {
	r1 := customTextfieldDelegateImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomTextfieldDelegate(r1)
}

func (m *TCustomTextfieldDelegate) AsInterface() ICefTextfieldDelegate {
	var resultCefTextfieldDelegate uintptr
	customTextfieldDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTextfieldDelegate)))
	return AsCefTextfieldDelegate(resultCefTextfieldDelegate)
}

var (
	customTextfieldDelegateImport       *imports.Imports = nil
	customTextfieldDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTextfieldDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("CustomTextfieldDelegate_Create", 0),
	}
)

func customTextfieldDelegateImportAPI() *imports.Imports {
	if customTextfieldDelegateImport == nil {
		customTextfieldDelegateImport = NewDefaultImports()
		customTextfieldDelegateImport.SetImportTable(customTextfieldDelegateImportTables)
		customTextfieldDelegateImportTables = nil
	}
	return customTextfieldDelegateImport
}
