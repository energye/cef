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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefDownloadHandler // procedure
}

// TCustomDownloadHandler Parent: TCefDownloadHandler
type TCustomDownloadHandler struct {
	TCefDownloadHandler
}

func NewCustomDownloadHandler(events IChromiumEvents) ICustomDownloadHandler {
	r1 := customDownloadHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomDownloadHandler(r1)
}

func (m *TCustomDownloadHandler) AsInterface() ICefDownloadHandler {
	var resultCefDownloadHandler uintptr
	customDownloadHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefDownloadHandler)))
	return AsCefDownloadHandler(resultCefDownloadHandler)
}

var (
	customDownloadHandlerImport       *imports.Imports = nil
	customDownloadHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDownloadHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomDownloadHandler_Create", 0),
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
