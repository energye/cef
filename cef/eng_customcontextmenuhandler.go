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

// ICustomContextMenuHandler Parent: ICefContextMenuHandler
type ICustomContextMenuHandler interface {
	ICefContextMenuHandler
}

// TCustomContextMenuHandler Parent: TCefContextMenuHandler
type TCustomContextMenuHandler struct {
	TCefContextMenuHandler
}

func NewCustomContextMenuHandler(events IChromiumEvents) ICustomContextMenuHandler {
	r1 := customContextMenuHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomContextMenuHandler(r1)
}

var (
	customContextMenuHandlerImport       *imports.Imports = nil
	customContextMenuHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomContextMenuHandler_Create", 0),
	}
)

func customContextMenuHandlerImportAPI() *imports.Imports {
	if customContextMenuHandlerImport == nil {
		customContextMenuHandlerImport = NewDefaultImports()
		customContextMenuHandlerImport.SetImportTable(customContextMenuHandlerImportTables)
		customContextMenuHandlerImportTables = nil
	}
	return customContextMenuHandlerImport
}
