//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefCustomRenderProcessHandler Parent: ICefRenderProcessHandlerOwn
type ICefCustomRenderProcessHandler interface {
	ICefRenderProcessHandlerOwn
	AsIntfRenderProcessHandler() uintptr
}

type TCefCustomRenderProcessHandler struct {
	TCefRenderProcessHandlerOwn
}

func (m *TCefCustomRenderProcessHandler) AsIntfRenderProcessHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomRenderProcessHandler class constructor
func NewCustomRenderProcessHandler(cefApp ICefApplicationCore) ICefCustomRenderProcessHandler {
	var renderProcessHandlerPtr uintptr // ICefRenderProcessHandler
	r := cefCustomRenderProcessHandlerAPI().SysCallN(0, base.GetObjectUintptr(cefApp), uintptr(base.UnsafePointer(&renderProcessHandlerPtr)))
	ret := AsCefCustomRenderProcessHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, renderProcessHandlerPtr)
	}
	return ret
}

var (
	cefCustomRenderProcessHandlerOnce   base.Once
	cefCustomRenderProcessHandlerImport *imports.Imports = nil
)

func cefCustomRenderProcessHandlerAPI() *imports.Imports {
	cefCustomRenderProcessHandlerOnce.Do(func() {
		cefCustomRenderProcessHandlerImport = api.NewDefaultImports()
		cefCustomRenderProcessHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomRenderProcessHandler_Create", 0), // constructor NewCustomRenderProcessHandler
		}
	})
	return cefCustomRenderProcessHandlerImport
}
