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

// ICefWindowDelegateOwn Parent: ICefPanelDelegateOwn
type ICefWindowDelegateOwn interface {
	ICefPanelDelegateOwn
	AsIntfWindowDelegate() uintptr
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefWindowDelegateOwn struct {
	TCefPanelDelegateOwn
}

func (m *TCefWindowDelegateOwn) AsIntfWindowDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefWindowDelegateOwn) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefWindowDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewWindowDelegateOwn class constructor
func NewWindowDelegateOwn() ICefWindowDelegateOwn {
	var windowDelegatePtr uintptr // ICefWindowDelegate
	var panelDelegatePtr uintptr  // ICefPanelDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := cefWindowDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&windowDelegatePtr)), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefWindowDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowDelegatePtr)
		intf.SetIntfPointer(1, panelDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	cefWindowDelegateOwnOnce   base.Once
	cefWindowDelegateOwnImport *imports.Imports = nil
)

func cefWindowDelegateOwnAPI() *imports.Imports {
	cefWindowDelegateOwnOnce.Do(func() {
		cefWindowDelegateOwnImport = api.NewDefaultImports()
		cefWindowDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefWindowDelegateOwn_Create", 0), // constructor NewWindowDelegateOwn
		}
	})
	return cefWindowDelegateOwnImport
}
