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
}

// TCustomWindowDelegate Parent: TCefWindowDelegate
//
//	This class handles all the TCustomWindowDelegate methods which call the ICefWindowDelegateEvents methods.
//	ICefWindowDelegateEvents will be implemented by the control receiving the TCustomWindowDelegate events.
type TCustomWindowDelegate struct {
	TCefWindowDelegate
}

func NewCustomWindowDelegate(events ICefWindowDelegateEvents) ICustomWindowDelegate {
	r1 := customWindowDelegateImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomWindowDelegate(r1)
}

var (
	customWindowDelegateImport       *imports.Imports = nil
	customWindowDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomWindowDelegate_Create", 0),
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
