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

// ICustomClientHandler Parent: ICefClientOwn
type ICustomClientHandler interface {
	ICefClientOwn
	RemoveReferences() // procedure
	AsIntfClient() uintptr
}

type TCustomClientHandler struct {
	TCefClientOwn
}

func (m *TCustomClientHandler) RemoveReferences() {
	if !m.IsValid() {
		return
	}
	customClientHandlerAPI().SysCallN(1, m.Instance())
}

func (m *TCustomClientHandler) AsIntfClient() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomClientHandler class constructor
func NewCustomClientHandler(events IChromiumCore, devToolsClient bool) ICustomClientHandler {
	var clientPtr uintptr // ICefClient
	r := customClientHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), api.PasBool(devToolsClient), uintptr(base.UnsafePointer(&clientPtr)))
	ret := AsCustomClientHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, clientPtr)
	}
	return ret
}

var (
	customClientHandlerOnce   base.Once
	customClientHandlerImport *imports.Imports = nil
)

func customClientHandlerAPI() *imports.Imports {
	customClientHandlerOnce.Do(func() {
		customClientHandlerImport = api.NewDefaultImports()
		customClientHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomClientHandler_Create", 0), // constructor NewCustomClientHandler
			/* 1 */ imports.NewTable("TCustomClientHandler_RemoveReferences", 0), // procedure RemoveReferences
		}
	})
	return customClientHandlerImport
}
