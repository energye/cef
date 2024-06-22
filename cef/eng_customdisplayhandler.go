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

// ICustomDisplayHandler Parent: ICefDisplayHandler
type ICustomDisplayHandler interface {
	ICefDisplayHandler
}

// TCustomDisplayHandler Parent: TCefDisplayHandler
type TCustomDisplayHandler struct {
	TCefDisplayHandler
}

func NewCustomDisplayHandler(events IChromiumEvents) ICustomDisplayHandler {
	r1 := customDisplayHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomDisplayHandler(r1)
}

var (
	customDisplayHandlerImport       *imports.Imports = nil
	customDisplayHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDisplayHandler_Create", 0),
	}
)

func customDisplayHandlerImportAPI() *imports.Imports {
	if customDisplayHandlerImport == nil {
		customDisplayHandlerImport = NewDefaultImports()
		customDisplayHandlerImport.SetImportTable(customDisplayHandlerImportTables)
		customDisplayHandlerImportTables = nil
	}
	return customDisplayHandlerImport
}
