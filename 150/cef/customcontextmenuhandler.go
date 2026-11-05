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

// ICustomContextMenuHandler Parent: ICefContextMenuHandlerOwn
type ICustomContextMenuHandler interface {
	ICefContextMenuHandlerOwn
	AsIntfContextMenuHandler() uintptr
}

type TCustomContextMenuHandler struct {
	TCefContextMenuHandlerOwn
}

func (m *TCustomContextMenuHandler) AsIntfContextMenuHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomContextMenuHandler class constructor
func NewCustomContextMenuHandler(events IChromiumCore) ICustomContextMenuHandler {
	var contextMenuHandlerPtr uintptr // ICefContextMenuHandler
	r := customContextMenuHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&contextMenuHandlerPtr)))
	ret := AsCustomContextMenuHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, contextMenuHandlerPtr)
	}
	return ret
}

var (
	customContextMenuHandlerOnce   base.Once
	customContextMenuHandlerImport *imports.Imports = nil
)

func customContextMenuHandlerAPI() *imports.Imports {
	customContextMenuHandlerOnce.Do(func() {
		customContextMenuHandlerImport = api.NewDefaultImports()
		customContextMenuHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomContextMenuHandler_Create", 0), // constructor NewCustomContextMenuHandler
		}
	})
	return customContextMenuHandlerImport
}
