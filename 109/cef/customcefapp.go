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

// ICustomCefApp Parent: ICefAppOwn
type ICustomCefApp interface {
	ICefAppOwn
	AsIntfApp() uintptr
}

type TCustomCefApp struct {
	TCefAppOwn
}

func (m *TCustomCefApp) AsIntfApp() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCefApp class constructor
func NewCustomCefApp(cefApp ICefApplicationCore) ICustomCefApp {
	var appPtr uintptr // ICefApp
	r := customCefAppAPI().SysCallN(0, base.GetObjectUintptr(cefApp), uintptr(base.UnsafePointer(&appPtr)))
	ret := AsCustomCefApp(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, appPtr)
	}
	return ret
}

var (
	customCefAppOnce   base.Once
	customCefAppImport *imports.Imports = nil
)

func customCefAppAPI() *imports.Imports {
	customCefAppOnce.Do(func() {
		customCefAppImport = api.NewDefaultImports()
		customCefAppImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCefApp_Create", 0), // constructor NewCustomCefApp
		}
	})
	return customCefAppImport
}
