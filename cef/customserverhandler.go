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

// ICustomServerHandler Parent: ICEFServerHandlerOwn
type ICustomServerHandler interface {
	ICEFServerHandlerOwn
	AsIntfServerHandler() uintptr
}

type TCustomServerHandler struct {
	TCEFServerHandlerOwn
}

func (m *TCustomServerHandler) AsIntfServerHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomServerHandler class constructor
func NewCustomServerHandler(events ICEFServerComponent) ICustomServerHandler {
	var serverHandlerPtr uintptr // ICefServerHandler
	r := customServerHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&serverHandlerPtr)))
	ret := AsCustomServerHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, serverHandlerPtr)
	}
	return ret
}

var (
	customServerHandlerOnce   base.Once
	customServerHandlerImport *imports.Imports = nil
)

func customServerHandlerAPI() *imports.Imports {
	customServerHandlerOnce.Do(func() {
		customServerHandlerImport = api.NewDefaultImports()
		customServerHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomServerHandler_Create", 0), // constructor NewCustomServerHandler
		}
	})
	return customServerHandlerImport
}
