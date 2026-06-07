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

// ICefCustomBrowserProcessHandler Parent: ICefBrowserProcessHandlerOwn
type ICefCustomBrowserProcessHandler interface {
	ICefBrowserProcessHandlerOwn
	AsIntfBrowserProcessHandler() uintptr
}

type TCefCustomBrowserProcessHandler struct {
	TCefBrowserProcessHandlerOwn
}

func (m *TCefCustomBrowserProcessHandler) AsIntfBrowserProcessHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomBrowserProcessHandler class constructor
func NewCustomBrowserProcessHandler(cefApp ICefApplicationCore) ICefCustomBrowserProcessHandler {
	var browserProcessHandlerPtr uintptr // ICefBrowserProcessHandler
	r := cefCustomBrowserProcessHandlerAPI().SysCallN(0, base.GetObjectUintptr(cefApp), uintptr(base.UnsafePointer(&browserProcessHandlerPtr)))
	ret := AsCefCustomBrowserProcessHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserProcessHandlerPtr)
	}
	return ret
}

var (
	cefCustomBrowserProcessHandlerOnce   base.Once
	cefCustomBrowserProcessHandlerImport *imports.Imports = nil
)

func cefCustomBrowserProcessHandlerAPI() *imports.Imports {
	cefCustomBrowserProcessHandlerOnce.Do(func() {
		cefCustomBrowserProcessHandlerImport = api.NewDefaultImports()
		cefCustomBrowserProcessHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomBrowserProcessHandler_Create", 0), // constructor NewCustomBrowserProcessHandler
		}
	})
	return cefCustomBrowserProcessHandlerImport
}
