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

// ICustomPermissionHandler Parent: ICefPermissionHandlerOwn
type ICustomPermissionHandler interface {
	ICefPermissionHandlerOwn
	AsIntfPermissionHandler() uintptr
}

type TCustomPermissionHandler struct {
	TCefPermissionHandlerOwn
}

func (m *TCustomPermissionHandler) AsIntfPermissionHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomPermissionHandler class constructor
func NewCustomPermissionHandler(events IChromiumCore) ICustomPermissionHandler {
	var permissionHandlerPtr uintptr // ICefPermissionHandler
	r := customPermissionHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&permissionHandlerPtr)))
	ret := AsCustomPermissionHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, permissionHandlerPtr)
	}
	return ret
}

var (
	customPermissionHandlerOnce   base.Once
	customPermissionHandlerImport *imports.Imports = nil
)

func customPermissionHandlerAPI() *imports.Imports {
	customPermissionHandlerOnce.Do(func() {
		customPermissionHandlerImport = api.NewDefaultImports()
		customPermissionHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPermissionHandler_Create", 0), // constructor NewCustomPermissionHandler
		}
	})
	return customPermissionHandlerImport
}
