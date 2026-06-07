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

// ICustomLoadHandler Parent: ICefLoadHandlerOwn
type ICustomLoadHandler interface {
	ICefLoadHandlerOwn
	AsIntfLoadHandler() uintptr
}

type TCustomLoadHandler struct {
	TCefLoadHandlerOwn
}

func (m *TCustomLoadHandler) AsIntfLoadHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomLoadHandler class constructor
func NewCustomLoadHandler(events IChromiumCore) ICustomLoadHandler {
	var loadHandlerPtr uintptr // ICefLoadHandler
	r := customLoadHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&loadHandlerPtr)))
	ret := AsCustomLoadHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, loadHandlerPtr)
	}
	return ret
}

var (
	customLoadHandlerOnce   base.Once
	customLoadHandlerImport *imports.Imports = nil
)

func customLoadHandlerAPI() *imports.Imports {
	customLoadHandlerOnce.Do(func() {
		customLoadHandlerImport = api.NewDefaultImports()
		customLoadHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomLoadHandler_Create", 0), // constructor NewCustomLoadHandler
		}
	})
	return customLoadHandlerImport
}
