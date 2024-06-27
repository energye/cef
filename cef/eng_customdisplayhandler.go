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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDisplayHandler // procedure
}

// TCustomDisplayHandler Parent: TCefDisplayHandler
type TCustomDisplayHandler struct {
	TCefDisplayHandler
}

func NewCustomDisplayHandler(events IChromiumEvents) ICustomDisplayHandler {
	r1 := customDisplayHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomDisplayHandler(r1)
}

func (m *TCustomDisplayHandler) AsInterface() ICefDisplayHandler {
	var resultCefDisplayHandler uintptr
	customDisplayHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDisplayHandler)))
	return AsCefDisplayHandler(resultCefDisplayHandler)
}

var (
	customDisplayHandlerImport       *imports.Imports = nil
	customDisplayHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDisplayHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomDisplayHandler_Create", 0),
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
