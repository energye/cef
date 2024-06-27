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

// ICustomBrowserViewDelegate Parent: ICefBrowserViewDelegate
//
//	This class handles all the ICefBrowserViewDelegate methods which call the ICefBrowserViewDelegateEvents methods.
//	ICefBrowserViewDelegateEvents will be implemented by the control receiving the ICefBrowserViewDelegate events.
type ICustomBrowserViewDelegate interface {
	ICefBrowserViewDelegate
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefBrowserViewDelegate // procedure
}

// TCustomBrowserViewDelegate Parent: TCefBrowserViewDelegate
//
//	This class handles all the ICefBrowserViewDelegate methods which call the ICefBrowserViewDelegateEvents methods.
//	ICefBrowserViewDelegateEvents will be implemented by the control receiving the ICefBrowserViewDelegate events.
type TCustomBrowserViewDelegate struct {
	TCefBrowserViewDelegate
}

func NewCustomBrowserViewDelegate(events ICefBrowserViewDelegateEvents) ICustomBrowserViewDelegate {
	r1 := customBrowserViewDelegateImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomBrowserViewDelegate(r1)
}

func (m *TCustomBrowserViewDelegate) AsInterface() ICefBrowserViewDelegate {
	var resultCefBrowserViewDelegate uintptr
	customBrowserViewDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefBrowserViewDelegate)))
	return AsCefBrowserViewDelegate(resultCefBrowserViewDelegate)
}

var (
	customBrowserViewDelegateImport       *imports.Imports = nil
	customBrowserViewDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomBrowserViewDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("CustomBrowserViewDelegate_Create", 0),
	}
)

func customBrowserViewDelegateImportAPI() *imports.Imports {
	if customBrowserViewDelegateImport == nil {
		customBrowserViewDelegateImport = NewDefaultImports()
		customBrowserViewDelegateImport.SetImportTable(customBrowserViewDelegateImportTables)
		customBrowserViewDelegateImportTables = nil
	}
	return customBrowserViewDelegateImport
}
