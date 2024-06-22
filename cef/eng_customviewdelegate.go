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

// ICustomViewDelegate Parent: ICefViewDelegate
//
//	This class handles all the ICefViewDelegate methods which call the ICefViewDelegateEvents methods.
//	ICefViewDelegateEvents will be implemented by the control receiving the ICefViewDelegate events.
type ICustomViewDelegate interface {
	ICefViewDelegate
}

// TCustomViewDelegate Parent: TCefViewDelegate
//
//	This class handles all the ICefViewDelegate methods which call the ICefViewDelegateEvents methods.
//	ICefViewDelegateEvents will be implemented by the control receiving the ICefViewDelegate events.
type TCustomViewDelegate struct {
	TCefViewDelegate
}

func NewCustomViewDelegate(events ICefViewDelegateEvents) ICustomViewDelegate {
	r1 := customViewDelegateImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomViewDelegate(r1)
}

var (
	customViewDelegateImport       *imports.Imports = nil
	customViewDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomViewDelegate_Create", 0),
	}
)

func customViewDelegateImportAPI() *imports.Imports {
	if customViewDelegateImport == nil {
		customViewDelegateImport = NewDefaultImports()
		customViewDelegateImport.SetImportTable(customViewDelegateImportTables)
		customViewDelegateImportTables = nil
	}
	return customViewDelegateImport
}
