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

// ICustomFrameHandler Parent: ICefFrameHandlerOwn
type ICustomFrameHandler interface {
	ICefFrameHandlerOwn
	AsIntfFrameHandler() uintptr
}

type TCustomFrameHandler struct {
	TCefFrameHandlerOwn
}

func (m *TCustomFrameHandler) AsIntfFrameHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomFrameHandler class constructor
func NewCustomFrameHandler(events IChromiumCore) ICustomFrameHandler {
	var frameHandlerPtr uintptr // ICefFrameHandler
	r := customFrameHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&frameHandlerPtr)))
	ret := AsCustomFrameHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, frameHandlerPtr)
	}
	return ret
}

var (
	customFrameHandlerOnce   base.Once
	customFrameHandlerImport *imports.Imports = nil
)

func customFrameHandlerAPI() *imports.Imports {
	customFrameHandlerOnce.Do(func() {
		customFrameHandlerImport = api.NewDefaultImports()
		customFrameHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFrameHandler_Create", 0), // constructor NewCustomFrameHandler
		}
	})
	return customFrameHandlerImport
}
