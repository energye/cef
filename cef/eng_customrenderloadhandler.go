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

// ICustomRenderLoadHandler Parent: ICefLoadHandler
type ICustomRenderLoadHandler interface {
	ICefLoadHandler
}

// TCustomRenderLoadHandler Parent: TCefLoadHandler
type TCustomRenderLoadHandler struct {
	TCefLoadHandler
}

func NewCustomRenderLoadHandler(aCefApp ICefApplicationCore) ICustomRenderLoadHandler {
	r1 := customRenderLoadHandlerImportAPI().SysCallN(0, GetObjectUintptr(aCefApp))
	return AsCustomRenderLoadHandler(r1)
}

var (
	customRenderLoadHandlerImport       *imports.Imports = nil
	customRenderLoadHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRenderLoadHandler_Create", 0),
	}
)

func customRenderLoadHandlerImportAPI() *imports.Imports {
	if customRenderLoadHandlerImport == nil {
		customRenderLoadHandlerImport = NewDefaultImports()
		customRenderLoadHandlerImport.SetImportTable(customRenderLoadHandlerImportTables)
		customRenderLoadHandlerImportTables = nil
	}
	return customRenderLoadHandlerImport
}
