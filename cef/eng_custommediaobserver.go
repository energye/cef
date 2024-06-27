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

// ICustomMediaObserver Parent: ICefMediaObserver
type ICustomMediaObserver interface {
	ICefMediaObserver
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefMediaObserver // procedure
}

// TCustomMediaObserver Parent: TCefMediaObserver
type TCustomMediaObserver struct {
	TCefMediaObserver
}

func NewCustomMediaObserver(events IChromiumEvents) ICustomMediaObserver {
	r1 := customMediaObserverImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomMediaObserver(r1)
}

func (m *TCustomMediaObserver) AsInterface() ICefMediaObserver {
	var resultCefMediaObserver uintptr
	customMediaObserverImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMediaObserver)))
	return AsCefMediaObserver(resultCefMediaObserver)
}

var (
	customMediaObserverImport       *imports.Imports = nil
	customMediaObserverImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMediaObserver_AsInterface", 0),
		/*1*/ imports.NewTable("CustomMediaObserver_Create", 0),
	}
)

func customMediaObserverImportAPI() *imports.Imports {
	if customMediaObserverImport == nil {
		customMediaObserverImport = NewDefaultImports()
		customMediaObserverImport.SetImportTable(customMediaObserverImportTables)
		customMediaObserverImportTables = nil
	}
	return customMediaObserverImport
}
