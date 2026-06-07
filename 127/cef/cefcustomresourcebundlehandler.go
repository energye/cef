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

// ICefCustomResourceBundleHandler Parent: ICefResourceBundleHandlerOwn
type ICefCustomResourceBundleHandler interface {
	ICefResourceBundleHandlerOwn
	AsIntfResourceBundleHandler() uintptr
}

type TCefCustomResourceBundleHandler struct {
	TCefResourceBundleHandlerOwn
}

func (m *TCefCustomResourceBundleHandler) AsIntfResourceBundleHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomResourceBundleHandler class constructor
func NewCustomResourceBundleHandler(cefApp ICefApplicationCore) ICefCustomResourceBundleHandler {
	var resourceBundleHandlerPtr uintptr // ICefResourceBundleHandler
	r := cefCustomResourceBundleHandlerAPI().SysCallN(0, base.GetObjectUintptr(cefApp), uintptr(base.UnsafePointer(&resourceBundleHandlerPtr)))
	ret := AsCefCustomResourceBundleHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceBundleHandlerPtr)
	}
	return ret
}

var (
	cefCustomResourceBundleHandlerOnce   base.Once
	cefCustomResourceBundleHandlerImport *imports.Imports = nil
)

func cefCustomResourceBundleHandlerAPI() *imports.Imports {
	cefCustomResourceBundleHandlerOnce.Do(func() {
		cefCustomResourceBundleHandlerImport = api.NewDefaultImports()
		cefCustomResourceBundleHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomResourceBundleHandler_Create", 0), // constructor NewCustomResourceBundleHandler
		}
	})
	return cefCustomResourceBundleHandlerImport
}
