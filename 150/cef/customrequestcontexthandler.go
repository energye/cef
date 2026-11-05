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

// ICustomRequestContextHandler Parent: ICefRequestContextHandlerOwn
type ICustomRequestContextHandler interface {
	ICefRequestContextHandlerOwn
	RemoveReferences() // procedure
	AsIntfRequestContextHandler() uintptr
}

type TCustomRequestContextHandler struct {
	TCefRequestContextHandlerOwn
}

func (m *TCustomRequestContextHandler) RemoveReferences() {
	if !m.IsValid() {
		return
	}
	customRequestContextHandlerAPI().SysCallN(1, m.Instance())
}

func (m *TCustomRequestContextHandler) AsIntfRequestContextHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomRequestContextHandler class constructor
func NewCustomRequestContextHandler(events IChromiumCore) ICustomRequestContextHandler {
	var requestContextHandlerPtr uintptr // ICefRequestContextHandler
	r := customRequestContextHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&requestContextHandlerPtr)))
	ret := AsCustomRequestContextHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, requestContextHandlerPtr)
	}
	return ret
}

var (
	customRequestContextHandlerOnce   base.Once
	customRequestContextHandlerImport *imports.Imports = nil
)

func customRequestContextHandlerAPI() *imports.Imports {
	customRequestContextHandlerOnce.Do(func() {
		customRequestContextHandlerImport = api.NewDefaultImports()
		customRequestContextHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRequestContextHandler_Create", 0), // constructor NewCustomRequestContextHandler
			/* 1 */ imports.NewTable("TCustomRequestContextHandler_RemoveReferences", 0), // procedure RemoveReferences
		}
	})
	return customRequestContextHandlerImport
}
