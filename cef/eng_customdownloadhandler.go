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

// ICustomDownloadHandler Parent: ICefDownloadHandler
type ICustomDownloadHandler interface {
	ICefDownloadHandler
}

// TCustomDownloadHandler Parent: TCefDownloadHandler
type TCustomDownloadHandler struct {
	TCefDownloadHandler
}

func NewCustomDownloadHandler(events IChromiumEvents) ICustomDownloadHandler {
	r1 := customDownloadHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomDownloadHandler(r1)
}

var (
	customDownloadHandlerImport       *imports.Imports = nil
	customDownloadHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDownloadHandler_Create", 0),
	}
)

func customDownloadHandlerImportAPI() *imports.Imports {
	if customDownloadHandlerImport == nil {
		customDownloadHandlerImport = NewDefaultImports()
		customDownloadHandlerImport.SetImportTable(customDownloadHandlerImportTables)
		customDownloadHandlerImportTables = nil
	}
	return customDownloadHandlerImport
}
