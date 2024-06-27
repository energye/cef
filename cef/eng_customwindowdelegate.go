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

// ICustomWindowDelegate Parent: ICefWindowDelegate
//
//	This class handles all the TCustomWindowDelegate methods which call the ICefWindowDelegateEvents methods.
//	ICefWindowDelegateEvents will be implemented by the control receiving the TCustomWindowDelegate events.
type ICustomWindowDelegate interface {
	ICefWindowDelegate
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefWindowDelegate // procedure
}

// TCustomWindowDelegate Parent: TCefWindowDelegate
//
//	This class handles all the TCustomWindowDelegate methods which call the ICefWindowDelegateEvents methods.
//	ICefWindowDelegateEvents will be implemented by the control receiving the TCustomWindowDelegate events.
type TCustomWindowDelegate struct {
	TCefWindowDelegate
}

func NewCustomWindowDelegate(events ICefWindowDelegateEvents) ICustomWindowDelegate {
	r1 := customWindowDelegateImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomWindowDelegate(r1)
}

func (m *TCustomWindowDelegate) AsInterface() ICefWindowDelegate {
	var resultCefWindowDelegate uintptr
	customWindowDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefWindowDelegate)))
	return AsCefWindowDelegate(resultCefWindowDelegate)
}

var (
	customWindowDelegateImport       *imports.Imports = nil
	customWindowDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomWindowDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("CustomWindowDelegate_Create", 0),
	}
)

func customWindowDelegateImportAPI() *imports.Imports {
	if customWindowDelegateImport == nil {
		customWindowDelegateImport = NewDefaultImports()
		customWindowDelegateImport.SetImportTable(customWindowDelegateImportTables)
		customWindowDelegateImportTables = nil
	}
	return customWindowDelegateImport
}
