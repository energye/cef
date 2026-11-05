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

// IEngPanelDelegate Parent: ICefPanelDelegateOwn
type IEngPanelDelegate interface {
	ICefPanelDelegateOwn
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngPanelDelegate struct {
	TCefPanelDelegateOwn
}

func (m *TEngPanelDelegate) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngPanelDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewEngPanelDelegate class constructor
func NewEngPanelDelegate() IEngPanelDelegate {
	var panelDelegatePtr uintptr // ICefPanelDelegate
	var viewDelegatePtr uintptr  // ICefViewDelegate
	r := engPanelDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngPanelDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	engPanelDelegateOnce   base.Once
	engPanelDelegateImport *imports.Imports = nil
)

func engPanelDelegateAPI() *imports.Imports {
	engPanelDelegateOnce.Do(func() {
		engPanelDelegateImport = api.NewDefaultImports()
		engPanelDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngPanelDelegate_Create", 0), // constructor NewEngPanelDelegate
		}
	})
	return engPanelDelegateImport
}
