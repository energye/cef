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

// ICustomExtensionHandler Parent: ICefExtensionHandler
type ICustomExtensionHandler interface {
	ICefExtensionHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefExtensionHandler // procedure
}

// TCustomExtensionHandler Parent: TCefExtensionHandler
type TCustomExtensionHandler struct {
	TCefExtensionHandler
}

func NewCustomExtensionHandler(events IChromiumEvents) ICustomExtensionHandler {
	r1 := customExtensionHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomExtensionHandler(r1)
}

func (m *TCustomExtensionHandler) AsInterface() ICefExtensionHandler {
	var resultCefExtensionHandler uintptr
	customExtensionHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefExtensionHandler)))
	return AsCefExtensionHandler(resultCefExtensionHandler)
}

var (
	customExtensionHandlerImport       *imports.Imports = nil
	customExtensionHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomExtensionHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomExtensionHandler_Create", 0),
	}
)

func customExtensionHandlerImportAPI() *imports.Imports {
	if customExtensionHandlerImport == nil {
		customExtensionHandlerImport = NewDefaultImports()
		customExtensionHandlerImport.SetImportTable(customExtensionHandlerImportTables)
		customExtensionHandlerImportTables = nil
	}
	return customExtensionHandlerImport
}
