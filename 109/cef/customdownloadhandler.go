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

// ICustomDownloadHandler Parent: ICefDownloadHandlerOwn
type ICustomDownloadHandler interface {
	ICefDownloadHandlerOwn
	AsIntfDownloadHandler() uintptr
}

type TCustomDownloadHandler struct {
	TCefDownloadHandlerOwn
}

func (m *TCustomDownloadHandler) AsIntfDownloadHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDownloadHandler class constructor
func NewCustomDownloadHandler(events IChromiumCore) ICustomDownloadHandler {
	var downloadHandlerPtr uintptr // ICefDownloadHandler
	r := customDownloadHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&downloadHandlerPtr)))
	ret := AsCustomDownloadHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadHandlerPtr)
	}
	return ret
}

var (
	customDownloadHandlerOnce   base.Once
	customDownloadHandlerImport *imports.Imports = nil
)

func customDownloadHandlerAPI() *imports.Imports {
	customDownloadHandlerOnce.Do(func() {
		customDownloadHandlerImport = api.NewDefaultImports()
		customDownloadHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDownloadHandler_Create", 0), // constructor NewCustomDownloadHandler
		}
	})
	return customDownloadHandlerImport
}
