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

// ICustomFindHandler Parent: ICefFindHandlerOwn
type ICustomFindHandler interface {
	ICefFindHandlerOwn
	AsIntfFindHandler() uintptr
}

type TCustomFindHandler struct {
	TCefFindHandlerOwn
}

func (m *TCustomFindHandler) AsIntfFindHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomFindHandler class constructor
func NewCustomFindHandler(events IChromiumCore) ICustomFindHandler {
	var findHandlerPtr uintptr // ICefFindHandler
	r := customFindHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&findHandlerPtr)))
	ret := AsCustomFindHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, findHandlerPtr)
	}
	return ret
}

var (
	customFindHandlerOnce   base.Once
	customFindHandlerImport *imports.Imports = nil
)

func customFindHandlerAPI() *imports.Imports {
	customFindHandlerOnce.Do(func() {
		customFindHandlerImport = api.NewDefaultImports()
		customFindHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFindHandler_Create", 0), // constructor NewCustomFindHandler
		}
	})
	return customFindHandlerImport
}
