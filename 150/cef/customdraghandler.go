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

// ICustomDragHandler Parent: ICefDragHandlerOwn
type ICustomDragHandler interface {
	ICefDragHandlerOwn
	AsIntfDragHandler() uintptr
}

type TCustomDragHandler struct {
	TCefDragHandlerOwn
}

func (m *TCustomDragHandler) AsIntfDragHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDragHandler class constructor
func NewCustomDragHandler(events IChromiumCore) ICustomDragHandler {
	var dragHandlerPtr uintptr // ICefDragHandler
	r := customDragHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&dragHandlerPtr)))
	ret := AsCustomDragHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dragHandlerPtr)
	}
	return ret
}

var (
	customDragHandlerOnce   base.Once
	customDragHandlerImport *imports.Imports = nil
)

func customDragHandlerAPI() *imports.Imports {
	customDragHandlerOnce.Do(func() {
		customDragHandlerImport = api.NewDefaultImports()
		customDragHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDragHandler_Create", 0), // constructor NewCustomDragHandler
		}
	})
	return customDragHandlerImport
}
