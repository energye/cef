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
}

// TCustomPrintHandler Parent: TCefPrintHandler
type TCustomPrintHandler struct {
	TCefPrintHandler
}

func NewCustomPrintHandler(events IChromiumEvents) ICustomPrintHandler {
	r1 := customPrintHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomPrintHandler(r1)
}

var (
	customPrintHandlerImport       *imports.Imports = nil
	customPrintHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPrintHandler_Create", 0),
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
