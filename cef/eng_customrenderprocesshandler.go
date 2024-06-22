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

// ICefCustomRenderProcessHandler Parent: ICefRenderProcessHandler
type ICefCustomRenderProcessHandler interface {
	ICefRenderProcessHandler
}

// TCefCustomRenderProcessHandler Parent: TCefRenderProcessHandler
type TCefCustomRenderProcessHandler struct {
	TCefRenderProcessHandler
}

func NewCefCustomRenderProcessHandler(aCefApp ICefApplicationCore) ICefCustomRenderProcessHandler {
	r1 := customRenderProcessHandlerImportAPI().SysCallN(0, GetObjectUintptr(aCefApp))
	return AsCefCustomRenderProcessHandler(r1)
}

var (
	customRenderProcessHandlerImport       *imports.Imports = nil
	customRenderProcessHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomRenderProcessHandler_Create", 0),
	}
)

func customRenderProcessHandlerImportAPI() *imports.Imports {
	if customRenderProcessHandlerImport == nil {
		customRenderProcessHandlerImport = NewDefaultImports()
		customRenderProcessHandlerImport.SetImportTable(customRenderProcessHandlerImportTables)
		customRenderProcessHandlerImportTables = nil
	}
	return customRenderProcessHandlerImport
}
