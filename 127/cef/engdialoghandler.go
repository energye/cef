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

// IEngDialogHandler Parent: ICefDialogHandlerOwn
type IEngDialogHandler interface {
	ICefDialogHandlerOwn
	SetOnDialogFileDialog(fn TOnDialogFileDialogEvent) // property event
	AsIntfDialogHandler() uintptr
}

type TEngDialogHandler struct {
	TCefDialogHandlerOwn
}

func (m *TEngDialogHandler) SetOnDialogFileDialog(fn TOnDialogFileDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDialogFileDialogEvent(fn)
	base.SetEvent(m, 1, engDialogHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDialogHandler) AsIntfDialogHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDialogHandler class constructor
func NewEngDialogHandler() IEngDialogHandler {
	var dialogHandlerPtr uintptr // ICefDialogHandler
	r := engDialogHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&dialogHandlerPtr)))
	ret := AsEngDialogHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dialogHandlerPtr)
	}
	return ret
}

var (
	engDialogHandlerOnce   base.Once
	engDialogHandlerImport *imports.Imports = nil
)

func engDialogHandlerAPI() *imports.Imports {
	engDialogHandlerOnce.Do(func() {
		engDialogHandlerImport = api.NewDefaultImports()
		engDialogHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDialogHandler_Create", 0), // constructor NewEngDialogHandler
			/* 1 */ imports.NewTable("TEngDialogHandler_OnDialogFileDialog", 0), // event OnDialogFileDialog
		}
	})
	return engDialogHandlerImport
}
