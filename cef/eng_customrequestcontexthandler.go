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

// ICustomRequestContextHandler Parent: ICefRequestContextHandler
type ICustomRequestContextHandler interface {
	ICefRequestContextHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefRequestContextHandler // procedure
	RemoveReferences()                      // procedure
}

// TCustomRequestContextHandler Parent: TCefRequestContextHandler
type TCustomRequestContextHandler struct {
	TCefRequestContextHandler
}

func NewCustomRequestContextHandler(events IChromiumEvents) ICustomRequestContextHandler {
	r1 := customRequestContextHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomRequestContextHandler(r1)
}

func (m *TCustomRequestContextHandler) AsInterface() ICefRequestContextHandler {
	var resultCefRequestContextHandler uintptr
	customRequestContextHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefRequestContextHandler)))
	return AsCefRequestContextHandler(resultCefRequestContextHandler)
}

func (m *TCustomRequestContextHandler) RemoveReferences() {
	customRequestContextHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	customRequestContextHandlerImport       *imports.Imports = nil
	customRequestContextHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRequestContextHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomRequestContextHandler_Create", 0),
		/*2*/ imports.NewTable("CustomRequestContextHandler_RemoveReferences", 0),
	}
)

func customRequestContextHandlerImportAPI() *imports.Imports {
	if customRequestContextHandlerImport == nil {
		customRequestContextHandlerImport = NewDefaultImports()
		customRequestContextHandlerImport.SetImportTable(customRequestContextHandlerImportTables)
		customRequestContextHandlerImportTables = nil
	}
	return customRequestContextHandlerImport
}
