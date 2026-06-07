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

// ICustomRenderLoadHandler Parent: ICefLoadHandlerOwn
type ICustomRenderLoadHandler interface {
	ICefLoadHandlerOwn
	AsIntfLoadHandler() uintptr
}

type TCustomRenderLoadHandler struct {
	TCefLoadHandlerOwn
}

func (m *TCustomRenderLoadHandler) AsIntfLoadHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomRenderLoadHandler class constructor
func NewCustomRenderLoadHandler(cefApp ICefApplicationCore) ICustomRenderLoadHandler {
	var loadHandlerPtr uintptr // ICefLoadHandler
	r := customRenderLoadHandlerAPI().SysCallN(0, base.GetObjectUintptr(cefApp), uintptr(base.UnsafePointer(&loadHandlerPtr)))
	ret := AsCustomRenderLoadHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, loadHandlerPtr)
	}
	return ret
}

var (
	customRenderLoadHandlerOnce   base.Once
	customRenderLoadHandlerImport *imports.Imports = nil
)

func customRenderLoadHandlerAPI() *imports.Imports {
	customRenderLoadHandlerOnce.Do(func() {
		customRenderLoadHandlerImport = api.NewDefaultImports()
		customRenderLoadHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRenderLoadHandler_Create", 0), // constructor NewCustomRenderLoadHandler
		}
	})
	return customRenderLoadHandlerImport
}
