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

// ICustomPrintHandler Parent: ICefPrintHandlerOwn
type ICustomPrintHandler interface {
	ICefPrintHandlerOwn
	AsIntfPrintHandler() uintptr
}

type TCustomPrintHandler struct {
	TCefPrintHandlerOwn
}

func (m *TCustomPrintHandler) AsIntfPrintHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomPrintHandler class constructor
func NewCustomPrintHandler(events IChromiumCore) ICustomPrintHandler {
	var printHandlerPtr uintptr // ICefPrintHandler
	r := customPrintHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&printHandlerPtr)))
	ret := AsCustomPrintHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printHandlerPtr)
	}
	return ret
}

var (
	customPrintHandlerOnce   base.Once
	customPrintHandlerImport *imports.Imports = nil
)

func customPrintHandlerAPI() *imports.Imports {
	customPrintHandlerOnce.Do(func() {
		customPrintHandlerImport = api.NewDefaultImports()
		customPrintHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPrintHandler_Create", 0), // constructor NewCustomPrintHandler
		}
	})
	return customPrintHandlerImport
}
