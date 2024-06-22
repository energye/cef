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
}

// TCustomJsDialogHandler Parent: TCefJsDialogHandler
type TCustomJsDialogHandler struct {
	TCefJsDialogHandler
}

func NewCustomJsDialogHandler(events IChromiumEvents) ICustomJsDialogHandler {
	r1 := customJsDialogHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomJsDialogHandler(r1)
}

var (
	customJsDialogHandlerImport       *imports.Imports = nil
	customJsDialogHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomJsDialogHandler_Create", 0),
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
