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

// ICustomPrintHandler Parent: ICefPrintHandler
type ICustomPrintHandler interface {
	ICefPrintHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefPrintHandler // procedure
}

// TCustomPrintHandler Parent: TCefPrintHandler
type TCustomPrintHandler struct {
	TCefPrintHandler
}

func NewCustomPrintHandler(events IChromiumEvents) ICustomPrintHandler {
	r1 := customPrintHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomPrintHandler(r1)
}

func (m *TCustomPrintHandler) AsInterface() ICefPrintHandler {
	var resultCefPrintHandler uintptr
	customPrintHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefPrintHandler)))
	return AsCefPrintHandler(resultCefPrintHandler)
}

var (
	customPrintHandlerImport       *imports.Imports = nil
	customPrintHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPrintHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomPrintHandler_Create", 0),
	}
)

func customPrintHandlerImportAPI() *imports.Imports {
	if customPrintHandlerImport == nil {
		customPrintHandlerImport = NewDefaultImports()
		customPrintHandlerImport.SetImportTable(customPrintHandlerImportTables)
		customPrintHandlerImportTables = nil
	}
	return customPrintHandlerImport
}
