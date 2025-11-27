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

// ICustomMenuButtonDelegate Parent: ICefMenuButtonDelegateOwn
type ICustomMenuButtonDelegate interface {
	ICefMenuButtonDelegateOwn
	AsIntfMenuButtonDelegate() uintptr
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomMenuButtonDelegate struct {
	TCefMenuButtonDelegateOwn
}

func (m *TCustomMenuButtonDelegate) AsIntfMenuButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomMenuButtonDelegate) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCustomMenuButtonDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewCustomMenuButtonDelegate class constructor
func NewCustomMenuButtonDelegate(events ICEFMenuButtonComponent) ICustomMenuButtonDelegate {
	var menuButtonDelegatePtr uintptr // ICefMenuButtonDelegate
	var buttonDelegatePtr uintptr     // ICefButtonDelegate
	var viewDelegatePtr uintptr       // ICefViewDelegate
	r := customMenuButtonDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&menuButtonDelegatePtr)), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomMenuButtonDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, menuButtonDelegatePtr)
		intf.SetIntfPointer(1, buttonDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	customMenuButtonDelegateOnce   base.Once
	customMenuButtonDelegateImport *imports.Imports = nil
)

func customMenuButtonDelegateAPI() *imports.Imports {
	customMenuButtonDelegateOnce.Do(func() {
		customMenuButtonDelegateImport = api.NewDefaultImports()
		customMenuButtonDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomMenuButtonDelegate_Create", 0), // constructor NewCustomMenuButtonDelegate
		}
	})
	return customMenuButtonDelegateImport
}
