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

// ICustomPanelDelegate Parent: ICefPanelDelegateOwn
type ICustomPanelDelegate interface {
	ICefPanelDelegateOwn
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomPanelDelegate struct {
	TCefPanelDelegateOwn
}

func (m *TCustomPanelDelegate) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomPanelDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewCustomPanelDelegate class constructor
func NewCustomPanelDelegate(events ICEFPanelComponent) ICustomPanelDelegate {
	var panelDelegatePtr uintptr // ICefPanelDelegate
	var viewDelegatePtr uintptr  // ICefViewDelegate
	r := customPanelDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomPanelDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, panelDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	customPanelDelegateOnce   base.Once
	customPanelDelegateImport *imports.Imports = nil
)

func customPanelDelegateAPI() *imports.Imports {
	customPanelDelegateOnce.Do(func() {
		customPanelDelegateImport = api.NewDefaultImports()
		customPanelDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPanelDelegate_Create", 0), // constructor NewCustomPanelDelegate
		}
	})
	return customPanelDelegateImport
}
