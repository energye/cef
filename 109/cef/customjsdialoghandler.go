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

// ICustomJsDialogHandler Parent: ICefJsDialogHandlerOwn
type ICustomJsDialogHandler interface {
	ICefJsDialogHandlerOwn
	AsIntfJsDialogHandler() uintptr
}

type TCustomJsDialogHandler struct {
	TCefJsDialogHandlerOwn
}

func (m *TCustomJsDialogHandler) AsIntfJsDialogHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomJsDialogHandler class constructor
func NewCustomJsDialogHandler(events IChromiumCore) ICustomJsDialogHandler {
	var jsDialogHandlerPtr uintptr // ICefJsDialogHandler
	r := customJsDialogHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&jsDialogHandlerPtr)))
	ret := AsCustomJsDialogHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, jsDialogHandlerPtr)
	}
	return ret
}

var (
	customJsDialogHandlerOnce   base.Once
	customJsDialogHandlerImport *imports.Imports = nil
)

func customJsDialogHandlerAPI() *imports.Imports {
	customJsDialogHandlerOnce.Do(func() {
		customJsDialogHandlerImport = api.NewDefaultImports()
		customJsDialogHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomJsDialogHandler_Create", 0), // constructor NewCustomJsDialogHandler
		}
	})
	return customJsDialogHandlerImport
}
