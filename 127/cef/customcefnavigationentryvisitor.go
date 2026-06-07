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

// ICustomCefNavigationEntryVisitor Parent: ICefNavigationEntryVisitorOwn
type ICustomCefNavigationEntryVisitor interface {
	ICefNavigationEntryVisitorOwn
	AsIntfNavigationEntryVisitor() uintptr
}

type TCustomCefNavigationEntryVisitor struct {
	TCefNavigationEntryVisitorOwn
}

func (m *TCustomCefNavigationEntryVisitor) AsIntfNavigationEntryVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCefNavigationEntryVisitor class constructor
func NewCustomCefNavigationEntryVisitor(events IChromiumCore) ICustomCefNavigationEntryVisitor {
	var navigationEntryVisitorPtr uintptr // ICefNavigationEntryVisitor
	r := customCefNavigationEntryVisitorAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&navigationEntryVisitorPtr)))
	ret := AsCustomCefNavigationEntryVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, navigationEntryVisitorPtr)
	}
	return ret
}

var (
	customCefNavigationEntryVisitorOnce   base.Once
	customCefNavigationEntryVisitorImport *imports.Imports = nil
)

func customCefNavigationEntryVisitorAPI() *imports.Imports {
	customCefNavigationEntryVisitorOnce.Do(func() {
		customCefNavigationEntryVisitorImport = api.NewDefaultImports()
		customCefNavigationEntryVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCefNavigationEntryVisitor_Create", 0), // constructor NewCustomCefNavigationEntryVisitor
		}
	})
	return customCefNavigationEntryVisitorImport
}
