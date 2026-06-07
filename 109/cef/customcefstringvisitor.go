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

// ICustomCefStringVisitor Parent: ICefStringVisitorOwn
type ICustomCefStringVisitor interface {
	ICefStringVisitorOwn
	AsIntfStringVisitor() uintptr
}

type TCustomCefStringVisitor struct {
	TCefStringVisitorOwn
}

func (m *TCustomCefStringVisitor) AsIntfStringVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCefStringVisitor class constructor
func NewCustomCefStringVisitor(events IChromiumCore) ICustomCefStringVisitor {
	var stringVisitorPtr uintptr // ICefStringVisitor
	r := customCefStringVisitorAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&stringVisitorPtr)))
	ret := AsCustomCefStringVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringVisitorPtr)
	}
	return ret
}

var (
	customCefStringVisitorOnce   base.Once
	customCefStringVisitorImport *imports.Imports = nil
)

func customCefStringVisitorAPI() *imports.Imports {
	customCefStringVisitorOnce.Do(func() {
		customCefStringVisitorImport = api.NewDefaultImports()
		customCefStringVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCefStringVisitor_Create", 0), // constructor NewCustomCefStringVisitor
		}
	})
	return customCefStringVisitorImport
}
