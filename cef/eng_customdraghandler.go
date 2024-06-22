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

// ICustomDragHandler Parent: ICefDragHandler
type ICustomDragHandler interface {
	ICefDragHandler
}

// TCustomDragHandler Parent: TCefDragHandler
type TCustomDragHandler struct {
	TCefDragHandler
}

func NewCustomDragHandler(events IChromiumEvents) ICustomDragHandler {
	r1 := customDragHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomDragHandler(r1)
}

var (
	customDragHandlerImport       *imports.Imports = nil
	customDragHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDragHandler_Create", 0),
	}
)

func customDragHandlerImportAPI() *imports.Imports {
	if customDragHandlerImport == nil {
		customDragHandlerImport = NewDefaultImports()
		customDragHandlerImport.SetImportTable(customDragHandlerImportTables)
		customDragHandlerImportTables = nil
	}
	return customDragHandlerImport
}
