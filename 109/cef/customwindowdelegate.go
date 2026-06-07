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

// ICustomWindowDelegate Parent: ICefWindowDelegateOwn
type ICustomWindowDelegate interface {
	ICefWindowDelegateOwn
	AsIntfWindowDelegate() uintptr
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomWindowDelegate struct {
	TCefWindowDelegateOwn
}

func (m *TCustomWindowDelegate) AsIntfWindowDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomWindowDelegate) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCustomWindowDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewCustomWindowDelegate class constructor
func NewCustomWindowDelegate(events ICEFWindowComponent) ICustomWindowDelegate {
	var windowDelegatePtr uintptr // ICefWindowDelegate
	var panelDelegatePtr uintptr  // ICefPanelDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := customWindowDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&windowDelegatePtr)), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomWindowDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowDelegatePtr)
		intf.SetIntfPointer(1, panelDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	customWindowDelegateOnce   base.Once
	customWindowDelegateImport *imports.Imports = nil
)

func customWindowDelegateAPI() *imports.Imports {
	customWindowDelegateOnce.Do(func() {
		customWindowDelegateImport = api.NewDefaultImports()
		customWindowDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomWindowDelegate_Create", 0), // constructor NewCustomWindowDelegate
		}
	})
	return customWindowDelegateImport
}
