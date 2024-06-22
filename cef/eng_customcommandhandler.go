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

// ICustomCommandHandler Parent: ICefCommandHandler
type ICustomCommandHandler interface {
	ICefCommandHandler
}

// TCustomCommandHandler Parent: TCefCommandHandler
type TCustomCommandHandler struct {
	TCefCommandHandler
}

func NewCustomCommandHandler(events IChromiumEvents) ICustomCommandHandler {
	r1 := customCommandHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomCommandHandler(r1)
}

var (
	customCommandHandlerImport       *imports.Imports = nil
	customCommandHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCommandHandler_Create", 0),
	}
)

func customCommandHandlerImportAPI() *imports.Imports {
	if customCommandHandlerImport == nil {
		customCommandHandlerImport = NewDefaultImports()
		customCommandHandlerImport.SetImportTable(customCommandHandlerImportTables)
		customCommandHandlerImportTables = nil
	}
	return customCommandHandlerImport
}
