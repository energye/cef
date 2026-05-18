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

// ICustomFocusHandler Parent: ICefFocusHandlerOwn
type ICustomFocusHandler interface {
	ICefFocusHandlerOwn
	AsIntfFocusHandler() uintptr
}

type TCustomFocusHandler struct {
	TCefFocusHandlerOwn
}

func (m *TCustomFocusHandler) AsIntfFocusHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomFocusHandler class constructor
func NewCustomFocusHandler(events IChromiumCore) ICustomFocusHandler {
	var focusHandlerPtr uintptr // ICefFocusHandler
	r := customFocusHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&focusHandlerPtr)))
	ret := AsCustomFocusHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, focusHandlerPtr)
	}
	return ret
}

var (
	customFocusHandlerOnce   base.Once
	customFocusHandlerImport *imports.Imports = nil
)

func customFocusHandlerAPI() *imports.Imports {
	customFocusHandlerOnce.Do(func() {
		customFocusHandlerImport = api.NewDefaultImports()
		customFocusHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFocusHandler_Create", 0), // constructor NewCustomFocusHandler
		}
	})
	return customFocusHandlerImport
}
