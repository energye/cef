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

// ICustomLifeSpanHandler Parent: ICefLifeSpanHandlerOwn
type ICustomLifeSpanHandler interface {
	ICefLifeSpanHandlerOwn
	AsIntfLifeSpanHandler() uintptr
}

type TCustomLifeSpanHandler struct {
	TCefLifeSpanHandlerOwn
}

func (m *TCustomLifeSpanHandler) AsIntfLifeSpanHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomLifeSpanHandler class constructor
func NewCustomLifeSpanHandler(events IChromiumCore) ICustomLifeSpanHandler {
	var lifeSpanHandlerPtr uintptr // ICefLifeSpanHandler
	r := customLifeSpanHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&lifeSpanHandlerPtr)))
	ret := AsCustomLifeSpanHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, lifeSpanHandlerPtr)
	}
	return ret
}

var (
	customLifeSpanHandlerOnce   base.Once
	customLifeSpanHandlerImport *imports.Imports = nil
)

func customLifeSpanHandlerAPI() *imports.Imports {
	customLifeSpanHandlerOnce.Do(func() {
		customLifeSpanHandlerImport = api.NewDefaultImports()
		customLifeSpanHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomLifeSpanHandler_Create", 0), // constructor NewCustomLifeSpanHandler
		}
	})
	return customLifeSpanHandlerImport
}
