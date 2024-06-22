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

// ICustomServerHandler Parent: ICEFServerHandler
type ICustomServerHandler interface {
	ICEFServerHandler
}

// TCustomServerHandler Parent: TCEFServerHandler
type TCustomServerHandler struct {
	TCEFServerHandler
}

func NewCustomServerHandler(events IServerEvents) ICustomServerHandler {
	r1 := customServerHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomServerHandler(r1)
}

var (
	customServerHandlerImport       *imports.Imports = nil
	customServerHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomServerHandler_Create", 0),
	}
)

func customServerHandlerImportAPI() *imports.Imports {
	if customServerHandlerImport == nil {
		customServerHandlerImport = NewDefaultImports()
		customServerHandlerImport.SetImportTable(customServerHandlerImportTables)
		customServerHandlerImportTables = nil
	}
	return customServerHandlerImport
}
