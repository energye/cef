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

// ICustomDevToolsMessageObserver Parent: ICefDevToolsMessageObserver
type ICustomDevToolsMessageObserver interface {
	ICefDevToolsMessageObserver
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDevToolsMessageObserver // procedure
}

// TCustomDevToolsMessageObserver Parent: TCefDevToolsMessageObserver
type TCustomDevToolsMessageObserver struct {
	TCefDevToolsMessageObserver
}

func NewCustomDevToolsMessageObserver(events IChromiumEvents) ICustomDevToolsMessageObserver {
	r1 := customDevToolsMessageObserverImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomDevToolsMessageObserver(r1)
}

func (m *TCustomDevToolsMessageObserver) AsInterface() ICefDevToolsMessageObserver {
	var resultCefDevToolsMessageObserver uintptr
	customDevToolsMessageObserverImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDevToolsMessageObserver)))
	return AsCefDevToolsMessageObserver(resultCefDevToolsMessageObserver)
}

var (
	customDevToolsMessageObserverImport       *imports.Imports = nil
	customDevToolsMessageObserverImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDevToolsMessageObserver_AsInterface", 0),
		/*1*/ imports.NewTable("CustomDevToolsMessageObserver_Create", 0),
	}
)

func customDevToolsMessageObserverImportAPI() *imports.Imports {
	if customDevToolsMessageObserverImport == nil {
		customDevToolsMessageObserverImport = NewDefaultImports()
		customDevToolsMessageObserverImport.SetImportTable(customDevToolsMessageObserverImportTables)
		customDevToolsMessageObserverImportTables = nil
	}
	return customDevToolsMessageObserverImport
}
