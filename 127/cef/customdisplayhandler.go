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

// ICustomDisplayHandler Parent: ICefDisplayHandlerOwn
type ICustomDisplayHandler interface {
	ICefDisplayHandlerOwn
	AsIntfDisplayHandler() uintptr
}

type TCustomDisplayHandler struct {
	TCefDisplayHandlerOwn
}

func (m *TCustomDisplayHandler) AsIntfDisplayHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDisplayHandler class constructor
func NewCustomDisplayHandler(events IChromiumCore) ICustomDisplayHandler {
	var displayHandlerPtr uintptr // ICefDisplayHandler
	r := customDisplayHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&displayHandlerPtr)))
	ret := AsCustomDisplayHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, displayHandlerPtr)
	}
	return ret
}

var (
	customDisplayHandlerOnce   base.Once
	customDisplayHandlerImport *imports.Imports = nil
)

func customDisplayHandlerAPI() *imports.Imports {
	customDisplayHandlerOnce.Do(func() {
		customDisplayHandlerImport = api.NewDefaultImports()
		customDisplayHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDisplayHandler_Create", 0), // constructor NewCustomDisplayHandler
		}
	})
	return customDisplayHandlerImport
}
