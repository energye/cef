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

// ICustomClientHandler Parent: ICefClient
type ICustomClientHandler interface {
	ICefClient
	RemoveReferences() // procedure
}

// TCustomClientHandler Parent: TCefClient
type TCustomClientHandler struct {
	TCefClient
}

func NewCustomClientHandler(events IChromiumEvents, aDevToolsClient bool) ICustomClientHandler {
	r1 := customClientHandlerImportAPI().SysCallN(0, GetObjectUintptr(events), PascalBool(aDevToolsClient))
	return AsCustomClientHandler(r1)
}

func (m *TCustomClientHandler) RemoveReferences() {
	customClientHandlerImportAPI().SysCallN(1, m.Instance())
}

var (
	customClientHandlerImport       *imports.Imports = nil
	customClientHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomClientHandler_Create", 0),
		/*1*/ imports.NewTable("CustomClientHandler_RemoveReferences", 0),
	}
)

func customClientHandlerImportAPI() *imports.Imports {
	if customClientHandlerImport == nil {
		customClientHandlerImport = NewDefaultImports()
		customClientHandlerImport.SetImportTable(customClientHandlerImportTables)
		customClientHandlerImportTables = nil
	}
	return customClientHandlerImport
}
