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

// ICustomFocusHandler Parent: ICefFocusHandler
type ICustomFocusHandler interface {
	ICefFocusHandler
}

// TCustomFocusHandler Parent: TCefFocusHandler
type TCustomFocusHandler struct {
	TCefFocusHandler
}

func NewCustomFocusHandler(events IChromiumEvents) ICustomFocusHandler {
	r1 := customFocusHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomFocusHandler(r1)
}

var (
	customFocusHandlerImport       *imports.Imports = nil
	customFocusHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFocusHandler_Create", 0),
	}
)

func customFocusHandlerImportAPI() *imports.Imports {
	if customFocusHandlerImport == nil {
		customFocusHandlerImport = NewDefaultImports()
		customFocusHandlerImport.SetImportTable(customFocusHandlerImportTables)
		customFocusHandlerImportTables = nil
	}
	return customFocusHandlerImport
}
