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

// ICustomDialogHandler Parent: ICefDialogHandlerOwn
type ICustomDialogHandler interface {
	ICefDialogHandlerOwn
	AsIntfDialogHandler() uintptr
}

type TCustomDialogHandler struct {
	TCefDialogHandlerOwn
}

func (m *TCustomDialogHandler) AsIntfDialogHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDialogHandler class constructor
func NewCustomDialogHandler(events IChromiumCore) ICustomDialogHandler {
	var dialogHandlerPtr uintptr // ICefDialogHandler
	r := customDialogHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&dialogHandlerPtr)))
	ret := AsCustomDialogHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dialogHandlerPtr)
	}
	return ret
}

var (
	customDialogHandlerOnce   base.Once
	customDialogHandlerImport *imports.Imports = nil
)

func customDialogHandlerAPI() *imports.Imports {
	customDialogHandlerOnce.Do(func() {
		customDialogHandlerImport = api.NewDefaultImports()
		customDialogHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDialogHandler_Create", 0), // constructor NewCustomDialogHandler
		}
	})
	return customDialogHandlerImport
}
