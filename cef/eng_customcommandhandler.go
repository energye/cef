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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefCommandHandler // procedure
}

// TCustomCommandHandler Parent: TCefCommandHandler
type TCustomCommandHandler struct {
	TCefCommandHandler
}

func NewCustomCommandHandler(events IChromiumEvents) ICustomCommandHandler {
	r1 := customCommandHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomCommandHandler(r1)
}

func (m *TCustomCommandHandler) AsInterface() ICefCommandHandler {
	var resultCefCommandHandler uintptr
	customCommandHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefCommandHandler)))
	return AsCefCommandHandler(resultCefCommandHandler)
}

var (
	customCommandHandlerImport       *imports.Imports = nil
	customCommandHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCommandHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomCommandHandler_Create", 0),
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
