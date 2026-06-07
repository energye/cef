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

// ICustomCommandHandler Parent: ICefCommandHandlerOwn
type ICustomCommandHandler interface {
	ICefCommandHandlerOwn
	AsIntfCommandHandler() uintptr
}

type TCustomCommandHandler struct {
	TCefCommandHandlerOwn
}

func (m *TCustomCommandHandler) AsIntfCommandHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCommandHandler class constructor
func NewCustomCommandHandler(events IChromiumCore) ICustomCommandHandler {
	var commandHandlerPtr uintptr // ICefCommandHandler
	r := customCommandHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&commandHandlerPtr)))
	ret := AsCustomCommandHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, commandHandlerPtr)
	}
	return ret
}

var (
	customCommandHandlerOnce   base.Once
	customCommandHandlerImport *imports.Imports = nil
)

func customCommandHandlerAPI() *imports.Imports {
	customCommandHandlerOnce.Do(func() {
		customCommandHandlerImport = api.NewDefaultImports()
		customCommandHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCommandHandler_Create", 0), // constructor NewCustomCommandHandler
		}
	})
	return customCommandHandlerImport
}
