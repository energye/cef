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
}

// TCustomExtensionHandler Parent: TCefExtensionHandler
type TCustomExtensionHandler struct {
	TCefExtensionHandler
}

func NewCustomExtensionHandler(events IChromiumEvents) ICustomExtensionHandler {
	r1 := customExtensionHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomExtensionHandler(r1)
}

var (
	customExtensionHandlerImport       *imports.Imports = nil
	customExtensionHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomExtensionHandler_Create", 0),
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
