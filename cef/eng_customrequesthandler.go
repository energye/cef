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

// ICustomRequestHandler Parent: ICefRequestHandler
type ICustomRequestHandler interface {
	ICefRequestHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefRequestHandler // procedure
	RemoveReferences()               // procedure
}

// TCustomRequestHandler Parent: TCefRequestHandler
type TCustomRequestHandler struct {
	TCefRequestHandler
}

func NewCustomRequestHandler(events IChromiumEvents) ICustomRequestHandler {
	r1 := customRequestHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomRequestHandler(r1)
}

func (m *TCustomRequestHandler) AsInterface() ICefRequestHandler {
	var resultCefRequestHandler uintptr
	customRequestHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefRequestHandler)))
	return AsCefRequestHandler(resultCefRequestHandler)
}

func (m *TCustomRequestHandler) RemoveReferences() {
	customRequestHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	customRequestHandlerImport       *imports.Imports = nil
	customRequestHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRequestHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomRequestHandler_Create", 0),
		/*2*/ imports.NewTable("CustomRequestHandler_RemoveReferences", 0),
	}
)

func customRequestHandlerImportAPI() *imports.Imports {
	if customRequestHandlerImport == nil {
		customRequestHandlerImport = NewDefaultImports()
		customRequestHandlerImport.SetImportTable(customRequestHandlerImportTables)
		customRequestHandlerImportTables = nil
	}
	return customRequestHandlerImport
}
