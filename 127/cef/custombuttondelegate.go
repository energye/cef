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

// ICustomButtonDelegate Parent: ICefButtonDelegateOwn
type ICustomButtonDelegate interface {
	ICefButtonDelegateOwn
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomButtonDelegate struct {
	TCefButtonDelegateOwn
}

func (m *TCustomButtonDelegate) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomButtonDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewCustomButtonDelegate class constructor
func NewCustomButtonDelegate(events ICEFButtonComponent) ICustomButtonDelegate {
	var buttonDelegatePtr uintptr // ICefButtonDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := customButtonDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomButtonDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	customButtonDelegateOnce   base.Once
	customButtonDelegateImport *imports.Imports = nil
)

func customButtonDelegateAPI() *imports.Imports {
	customButtonDelegateOnce.Do(func() {
		customButtonDelegateImport = api.NewDefaultImports()
		customButtonDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomButtonDelegate_Create", 0), // constructor NewCustomButtonDelegate
		}
	})
	return customButtonDelegateImport
}
