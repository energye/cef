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

// ICustomPermissionHandler Parent: ICefPermissionHandler
type ICustomPermissionHandler interface {
	ICefPermissionHandler
}

// TCustomPermissionHandler Parent: TCefPermissionHandler
type TCustomPermissionHandler struct {
	TCefPermissionHandler
}

func NewCustomPermissionHandler(events IChromiumEvents) ICustomPermissionHandler {
	r1 := customPermissionHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomPermissionHandler(r1)
}

var (
	customPermissionHandlerImport       *imports.Imports = nil
	customPermissionHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPermissionHandler_Create", 0),
	}
)

func customPermissionHandlerImportAPI() *imports.Imports {
	if customPermissionHandlerImport == nil {
		customPermissionHandlerImport = NewDefaultImports()
		customPermissionHandlerImport.SetImportTable(customPermissionHandlerImportTables)
		customPermissionHandlerImportTables = nil
	}
	return customPermissionHandlerImport
}
