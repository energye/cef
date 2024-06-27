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

// ICustomFindHandler Parent: ICefFindHandler
type ICustomFindHandler interface {
	ICefFindHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefFindHandler // procedure
}

// TCustomFindHandler Parent: TCefFindHandler
type TCustomFindHandler struct {
	TCefFindHandler
}

func NewCustomFindHandler(events IChromiumEvents) ICustomFindHandler {
	r1 := customFindHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomFindHandler(r1)
}

func (m *TCustomFindHandler) AsInterface() ICefFindHandler {
	var resultCefFindHandler uintptr
	customFindHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefFindHandler)))
	return AsCefFindHandler(resultCefFindHandler)
}

var (
	customFindHandlerImport       *imports.Imports = nil
	customFindHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFindHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomFindHandler_Create", 0),
	}
)

func customFindHandlerImportAPI() *imports.Imports {
	if customFindHandlerImport == nil {
		customFindHandlerImport = NewDefaultImports()
		customFindHandlerImport.SetImportTable(customFindHandlerImportTables)
		customFindHandlerImportTables = nil
	}
	return customFindHandlerImport
}
