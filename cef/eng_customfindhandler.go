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
}

// TCustomFindHandler Parent: TCefFindHandler
type TCustomFindHandler struct {
	TCefFindHandler
}

func NewCustomFindHandler(events IChromiumEvents) ICustomFindHandler {
	r1 := customFindHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomFindHandler(r1)
}

var (
	customFindHandlerImport       *imports.Imports = nil
	customFindHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFindHandler_Create", 0),
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
