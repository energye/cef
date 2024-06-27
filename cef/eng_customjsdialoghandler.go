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

// ICustomJsDialogHandler Parent: ICefJsDialogHandler
type ICustomJsDialogHandler interface {
	ICefJsDialogHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefJsDialogHandler // procedure
}

// TCustomJsDialogHandler Parent: TCefJsDialogHandler
type TCustomJsDialogHandler struct {
	TCefJsDialogHandler
}

func NewCustomJsDialogHandler(events IChromiumEvents) ICustomJsDialogHandler {
	r1 := customJsDialogHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomJsDialogHandler(r1)
}

func (m *TCustomJsDialogHandler) AsInterface() ICefJsDialogHandler {
	var resultCefJsDialogHandler uintptr
	customJsDialogHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefJsDialogHandler)))
	return AsCefJsDialogHandler(resultCefJsDialogHandler)
}

var (
	customJsDialogHandlerImport       *imports.Imports = nil
	customJsDialogHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomJsDialogHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomJsDialogHandler_Create", 0),
	}
)

func customJsDialogHandlerImportAPI() *imports.Imports {
	if customJsDialogHandlerImport == nil {
		customJsDialogHandlerImport = NewDefaultImports()
		customJsDialogHandlerImport.SetImportTable(customJsDialogHandlerImportTables)
		customJsDialogHandlerImportTables = nil
	}
	return customJsDialogHandlerImport
}
