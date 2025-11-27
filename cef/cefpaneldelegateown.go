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

// ICefPanelDelegateOwn Parent: ICefViewDelegateOwn
type ICefPanelDelegateOwn interface {
	ICefViewDelegateOwn
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefPanelDelegateOwn struct {
	TCefViewDelegateOwn
}

func (m *TCefPanelDelegateOwn) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefPanelDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewPanelDelegateOwn class constructor
func NewPanelDelegateOwn() ICefPanelDelegateOwn {
	var panelDelegatePtr uintptr // ICefPanelDelegate
	var viewDelegatePtr uintptr  // ICefViewDelegate
	r := cefPanelDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefPanelDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefPanelDelegateOwnOnce   base.Once
	cefPanelDelegateOwnImport *imports.Imports = nil
)

func cefPanelDelegateOwnAPI() *imports.Imports {
	cefPanelDelegateOwnOnce.Do(func() {
		cefPanelDelegateOwnImport = api.NewDefaultImports()
		cefPanelDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPanelDelegateOwn_Create", 0), // constructor NewPanelDelegateOwn
		}
	})
	return cefPanelDelegateOwnImport
}
