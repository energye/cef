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

// ICustomFrameHandler Parent: ICefFrameHandler
type ICustomFrameHandler interface {
	ICefFrameHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefFrameHandler // procedure
}

// TCustomFrameHandler Parent: TCefFrameHandler
type TCustomFrameHandler struct {
	TCefFrameHandler
}

func NewCustomFrameHandler(events IChromiumEvents) ICustomFrameHandler {
	r1 := customFrameHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomFrameHandler(r1)
}

func (m *TCustomFrameHandler) AsInterface() ICefFrameHandler {
	var resultCefFrameHandler uintptr
	customFrameHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefFrameHandler)))
	return AsCefFrameHandler(resultCefFrameHandler)
}

var (
	customFrameHandlerImport       *imports.Imports = nil
	customFrameHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFrameHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomFrameHandler_Create", 0),
	}
)

func customFrameHandlerImportAPI() *imports.Imports {
	if customFrameHandlerImport == nil {
		customFrameHandlerImport = NewDefaultImports()
		customFrameHandlerImport.SetImportTable(customFrameHandlerImportTables)
		customFrameHandlerImportTables = nil
	}
	return customFrameHandlerImport
}
