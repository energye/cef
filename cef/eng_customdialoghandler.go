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

// ICustomDialogHandler Parent: ICefDialogHandler
type ICustomDialogHandler interface {
	ICefDialogHandler
}

// TCustomDialogHandler Parent: TCefDialogHandler
type TCustomDialogHandler struct {
	TCefDialogHandler
}

func NewCustomDialogHandler(events IChromiumEvents) ICustomDialogHandler {
	r1 := customDialogHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomDialogHandler(r1)
}

var (
	customDialogHandlerImport       *imports.Imports = nil
	customDialogHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDialogHandler_Create", 0),
	}
)

func customDialogHandlerImportAPI() *imports.Imports {
	if customDialogHandlerImport == nil {
		customDialogHandlerImport = NewDefaultImports()
		customDialogHandlerImport.SetImportTable(customDialogHandlerImportTables)
		customDialogHandlerImportTables = nil
	}
	return customDialogHandlerImport
}
