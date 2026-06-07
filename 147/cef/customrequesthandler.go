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

// ICustomRequestHandler Parent: ICefRequestHandlerOwn
type ICustomRequestHandler interface {
	ICefRequestHandlerOwn
	RemoveReferences() // procedure
	AsIntfRequestHandler() uintptr
}

type TCustomRequestHandler struct {
	TCefRequestHandlerOwn
}

func (m *TCustomRequestHandler) RemoveReferences() {
	if !m.IsValid() {
		return
	}
	customRequestHandlerAPI().SysCallN(1, m.Instance())
}

func (m *TCustomRequestHandler) AsIntfRequestHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomRequestHandler class constructor
func NewCustomRequestHandler(events IChromiumCore) ICustomRequestHandler {
	var requestHandlerPtr uintptr // ICefRequestHandler
	r := customRequestHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&requestHandlerPtr)))
	ret := AsCustomRequestHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, requestHandlerPtr)
	}
	return ret
}

var (
	customRequestHandlerOnce   base.Once
	customRequestHandlerImport *imports.Imports = nil
)

func customRequestHandlerAPI() *imports.Imports {
	customRequestHandlerOnce.Do(func() {
		customRequestHandlerImport = api.NewDefaultImports()
		customRequestHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRequestHandler_Create", 0), // constructor NewCustomRequestHandler
			/* 1 */ imports.NewTable("TCustomRequestHandler_RemoveReferences", 0), // procedure RemoveReferences
		}
	})
	return customRequestHandlerImport
}
