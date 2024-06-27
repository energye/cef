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

// ICustomLoadHandler Parent: ICefLoadHandler
type ICustomLoadHandler interface {
	ICefLoadHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefLoadHandler // procedure
}

// TCustomLoadHandler Parent: TCefLoadHandler
type TCustomLoadHandler struct {
	TCefLoadHandler
}

func NewCustomLoadHandler(events IChromiumEvents) ICustomLoadHandler {
	r1 := customLoadHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomLoadHandler(r1)
}

func (m *TCustomLoadHandler) AsInterface() ICefLoadHandler {
	var resultCefLoadHandler uintptr
	customLoadHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefLoadHandler)))
	return AsCefLoadHandler(resultCefLoadHandler)
}

var (
	customLoadHandlerImport       *imports.Imports = nil
	customLoadHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomLoadHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomLoadHandler_Create", 0),
	}
)

func customLoadHandlerImportAPI() *imports.Imports {
	if customLoadHandlerImport == nil {
		customLoadHandlerImport = NewDefaultImports()
		customLoadHandlerImport.SetImportTable(customLoadHandlerImportTables)
		customLoadHandlerImportTables = nil
	}
	return customLoadHandlerImport
}
