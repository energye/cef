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

// ICustomRequestHandler Parent: ICefRequestHandler
type ICustomRequestHandler interface {
	ICefRequestHandler
	RemoveReferences() // procedure
}

// TCustomRequestHandler Parent: TCefRequestHandler
type TCustomRequestHandler struct {
	TCefRequestHandler
}

func NewCustomRequestHandler(events IChromiumEvents) ICustomRequestHandler {
	r1 := customRequestHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomRequestHandler(r1)
}

func (m *TCustomRequestHandler) RemoveReferences() {
	customRequestHandlerImportAPI().SysCallN(1, m.Instance())
}

var (
	customRequestHandlerImport       *imports.Imports = nil
	customRequestHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRequestHandler_Create", 0),
		/*1*/ imports.NewTable("CustomRequestHandler_RemoveReferences", 0),
	}
)

func customRequestHandlerImportAPI() *imports.Imports {
	if customRequestHandlerImport == nil {
		customRequestHandlerImport = NewDefaultImports()
		customRequestHandlerImport.SetImportTable(customRequestHandlerImportTables)
		customRequestHandlerImportTables = nil
	}
	return customRequestHandlerImport
}
