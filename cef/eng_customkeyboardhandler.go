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

// ICustomKeyboardHandler Parent: ICefKeyboardHandler
type ICustomKeyboardHandler interface {
	ICefKeyboardHandler
}

// TCustomKeyboardHandler Parent: TCefKeyboardHandler
type TCustomKeyboardHandler struct {
	TCefKeyboardHandler
}

func NewCustomKeyboardHandler(events IChromiumEvents) ICustomKeyboardHandler {
	r1 := customKeyboardHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomKeyboardHandler(r1)
}

var (
	customKeyboardHandlerImport       *imports.Imports = nil
	customKeyboardHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomKeyboardHandler_Create", 0),
	}
)

func customKeyboardHandlerImportAPI() *imports.Imports {
	if customKeyboardHandlerImport == nil {
		customKeyboardHandlerImport = NewDefaultImports()
		customKeyboardHandlerImport.SetImportTable(customKeyboardHandlerImportTables)
		customKeyboardHandlerImportTables = nil
	}
	return customKeyboardHandlerImport
}
