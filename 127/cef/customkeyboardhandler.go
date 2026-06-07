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

// ICustomKeyboardHandler Parent: ICefKeyboardHandlerOwn
type ICustomKeyboardHandler interface {
	ICefKeyboardHandlerOwn
	AsIntfKeyboardHandler() uintptr
}

type TCustomKeyboardHandler struct {
	TCefKeyboardHandlerOwn
}

func (m *TCustomKeyboardHandler) AsIntfKeyboardHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomKeyboardHandler class constructor
func NewCustomKeyboardHandler(events IChromiumCore) ICustomKeyboardHandler {
	var keyboardHandlerPtr uintptr // ICefKeyboardHandler
	r := customKeyboardHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&keyboardHandlerPtr)))
	ret := AsCustomKeyboardHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, keyboardHandlerPtr)
	}
	return ret
}

var (
	customKeyboardHandlerOnce   base.Once
	customKeyboardHandlerImport *imports.Imports = nil
)

func customKeyboardHandlerAPI() *imports.Imports {
	customKeyboardHandlerOnce.Do(func() {
		customKeyboardHandlerImport = api.NewDefaultImports()
		customKeyboardHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomKeyboardHandler_Create", 0), // constructor NewCustomKeyboardHandler
		}
	})
	return customKeyboardHandlerImport
}
