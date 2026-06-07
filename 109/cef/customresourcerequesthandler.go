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

// ICustomResourceRequestHandler Parent: ICefResourceRequestHandlerOwn
type ICustomResourceRequestHandler interface {
	ICefResourceRequestHandlerOwn
	RemoveReferences() // procedure
	AsIntfResourceRequestHandler() uintptr
}

type TCustomResourceRequestHandler struct {
	TCefResourceRequestHandlerOwn
}

func (m *TCustomResourceRequestHandler) RemoveReferences() {
	if !m.IsValid() {
		return
	}
	customResourceRequestHandlerAPI().SysCallN(1, m.Instance())
}

func (m *TCustomResourceRequestHandler) AsIntfResourceRequestHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomResourceRequestHandler class constructor
func NewCustomResourceRequestHandler(events IChromiumCore) ICustomResourceRequestHandler {
	var resourceRequestHandlerPtr uintptr // ICefResourceRequestHandler
	r := customResourceRequestHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&resourceRequestHandlerPtr)))
	ret := AsCustomResourceRequestHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceRequestHandlerPtr)
	}
	return ret
}

var (
	customResourceRequestHandlerOnce   base.Once
	customResourceRequestHandlerImport *imports.Imports = nil
)

func customResourceRequestHandlerAPI() *imports.Imports {
	customResourceRequestHandlerOnce.Do(func() {
		customResourceRequestHandlerImport = api.NewDefaultImports()
		customResourceRequestHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomResourceRequestHandler_Create", 0), // constructor NewCustomResourceRequestHandler
			/* 1 */ imports.NewTable("TCustomResourceRequestHandler_RemoveReferences", 0), // procedure RemoveReferences
		}
	})
	return customResourceRequestHandlerImport
}
