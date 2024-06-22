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

// ICustomRenderHandler Parent: ICefRenderHandler
type ICustomRenderHandler interface {
	ICefRenderHandler
}

// TCustomRenderHandler Parent: TCefRenderHandler
type TCustomRenderHandler struct {
	TCefRenderHandler
}

func NewCustomRenderHandler(events IChromiumEvents) ICustomRenderHandler {
	r1 := customRenderHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomRenderHandler(r1)
}

var (
	customRenderHandlerImport       *imports.Imports = nil
	customRenderHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRenderHandler_Create", 0),
	}
)

func customRenderHandlerImportAPI() *imports.Imports {
	if customRenderHandlerImport == nil {
		customRenderHandlerImport = NewDefaultImports()
		customRenderHandlerImport.SetImportTable(customRenderHandlerImportTables)
		customRenderHandlerImportTables = nil
	}
	return customRenderHandlerImport
}
