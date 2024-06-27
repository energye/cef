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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDragHandler // procedure
}

// TCustomDragHandler Parent: TCefDragHandler
type TCustomDragHandler struct {
	TCefDragHandler
}

func NewCustomDragHandler(events IChromiumEvents) ICustomDragHandler {
	r1 := customDragHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomDragHandler(r1)
}

func (m *TCustomDragHandler) AsInterface() ICefDragHandler {
	var resultCefDragHandler uintptr
	customDragHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDragHandler)))
	return AsCefDragHandler(resultCefDragHandler)
}

var (
	customDragHandlerImport       *imports.Imports = nil
	customDragHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDragHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomDragHandler_Create", 0),
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
