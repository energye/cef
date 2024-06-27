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

// ICustomResourceRequestHandler Parent: ICefResourceRequestHandler
type ICustomResourceRequestHandler interface {
	ICefResourceRequestHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefResourceRequestHandler // procedure
	RemoveReferences()                       // procedure
}

// TCustomResourceRequestHandler Parent: TCefResourceRequestHandler
type TCustomResourceRequestHandler struct {
	TCefResourceRequestHandler
}

func NewCustomResourceRequestHandler(events IChromiumEvents) ICustomResourceRequestHandler {
	r1 := customResourceRequestHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomResourceRequestHandler(r1)
}

func (m *TCustomResourceRequestHandler) AsInterface() ICefResourceRequestHandler {
	var resultCefResourceRequestHandler uintptr
	customResourceRequestHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefResourceRequestHandler)))
	return AsCefResourceRequestHandler(resultCefResourceRequestHandler)
}

func (m *TCustomResourceRequestHandler) RemoveReferences() {
	customResourceRequestHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	customResourceRequestHandlerImport       *imports.Imports = nil
	customResourceRequestHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomResourceRequestHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomResourceRequestHandler_Create", 0),
		/*2*/ imports.NewTable("CustomResourceRequestHandler_RemoveReferences", 0),
	}
)

func customResourceRequestHandlerImportAPI() *imports.Imports {
	if customResourceRequestHandlerImport == nil {
		customResourceRequestHandlerImport = NewDefaultImports()
		customResourceRequestHandlerImport.SetImportTable(customResourceRequestHandlerImportTables)
		customResourceRequestHandlerImportTables = nil
	}
	return customResourceRequestHandlerImport
}
