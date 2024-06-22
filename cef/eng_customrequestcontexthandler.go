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

// ICustomRequestContextHandler Parent: ICefRequestContextHandler
type ICustomRequestContextHandler interface {
	ICefRequestContextHandler
	RemoveReferences() // procedure
}

// TCustomRequestContextHandler Parent: TCefRequestContextHandler
type TCustomRequestContextHandler struct {
	TCefRequestContextHandler
}

func NewCustomRequestContextHandler(events IChromiumEvents) ICustomRequestContextHandler {
	r1 := customRequestContextHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomRequestContextHandler(r1)
}

func (m *TCustomRequestContextHandler) RemoveReferences() {
	customRequestContextHandlerImportAPI().SysCallN(1, m.Instance())
}

var (
	customRequestContextHandlerImport       *imports.Imports = nil
	customRequestContextHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRequestContextHandler_Create", 0),
		/*1*/ imports.NewTable("CustomRequestContextHandler_RemoveReferences", 0),
	}
)

func customRequestContextHandlerImportAPI() *imports.Imports {
	if customRequestContextHandlerImport == nil {
		customRequestContextHandlerImport = NewDefaultImports()
		customRequestContextHandlerImport.SetImportTable(customRequestContextHandlerImportTables)
		customRequestContextHandlerImportTables = nil
	}
	return customRequestContextHandlerImport
}
