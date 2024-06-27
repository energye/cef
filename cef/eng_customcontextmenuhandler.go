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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefContextMenuHandler // procedure
}

// TCustomContextMenuHandler Parent: TCefContextMenuHandler
type TCustomContextMenuHandler struct {
	TCefContextMenuHandler
}

func NewCustomContextMenuHandler(events IChromiumEvents) ICustomContextMenuHandler {
	r1 := customContextMenuHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomContextMenuHandler(r1)
}

func (m *TCustomContextMenuHandler) AsInterface() ICefContextMenuHandler {
	var resultCefContextMenuHandler uintptr
	customContextMenuHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefContextMenuHandler)))
	return AsCefContextMenuHandler(resultCefContextMenuHandler)
}

var (
	customContextMenuHandlerImport       *imports.Imports = nil
	customContextMenuHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomContextMenuHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomContextMenuHandler_Create", 0),
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
