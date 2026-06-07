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

// ICustomRenderHandler Parent: ICefRenderHandlerOwn
type ICustomRenderHandler interface {
	ICefRenderHandlerOwn
	AsIntfRenderHandler() uintptr
}

type TCustomRenderHandler struct {
	TCefRenderHandlerOwn
}

func (m *TCustomRenderHandler) AsIntfRenderHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomRenderHandler class constructor
func NewCustomRenderHandler(events IChromiumCore) ICustomRenderHandler {
	var renderHandlerPtr uintptr // ICefRenderHandler
	r := customRenderHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&renderHandlerPtr)))
	ret := AsCustomRenderHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, renderHandlerPtr)
	}
	return ret
}

var (
	customRenderHandlerOnce   base.Once
	customRenderHandlerImport *imports.Imports = nil
)

func customRenderHandlerAPI() *imports.Imports {
	customRenderHandlerOnce.Do(func() {
		customRenderHandlerImport = api.NewDefaultImports()
		customRenderHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRenderHandler_Create", 0), // constructor NewCustomRenderHandler
		}
	})
	return customRenderHandlerImport
}
