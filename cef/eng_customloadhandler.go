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
}

// TCustomLoadHandler Parent: TCefLoadHandler
type TCustomLoadHandler struct {
	TCefLoadHandler
}

func NewCustomLoadHandler(events IChromiumEvents) ICustomLoadHandler {
	r1 := customLoadHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomLoadHandler(r1)
}

var (
	customLoadHandlerImport       *imports.Imports = nil
	customLoadHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomLoadHandler_Create", 0),
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
