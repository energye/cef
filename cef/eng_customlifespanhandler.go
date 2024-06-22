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

// ICustomLifeSpanHandler Parent: ICefLifeSpanHandler
type ICustomLifeSpanHandler interface {
	ICefLifeSpanHandler
}

// TCustomLifeSpanHandler Parent: TCefLifeSpanHandler
type TCustomLifeSpanHandler struct {
	TCefLifeSpanHandler
}

func NewCustomLifeSpanHandler(events IChromiumEvents) ICustomLifeSpanHandler {
	r1 := customLifeSpanHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomLifeSpanHandler(r1)
}

var (
	customLifeSpanHandlerImport       *imports.Imports = nil
	customLifeSpanHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomLifeSpanHandler_Create", 0),
	}
)

func customLifeSpanHandlerImportAPI() *imports.Imports {
	if customLifeSpanHandlerImport == nil {
		customLifeSpanHandlerImport = NewDefaultImports()
		customLifeSpanHandlerImport.SetImportTable(customLifeSpanHandlerImportTables)
		customLifeSpanHandlerImportTables = nil
	}
	return customLifeSpanHandlerImport
}
