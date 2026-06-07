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

// ICustomViewDelegate Parent: ICefViewDelegateOwn
type ICustomViewDelegate interface {
	ICefViewDelegateOwn
	AsIntfViewDelegate() uintptr
}

type TCustomViewDelegate struct {
	TCefViewDelegateOwn
}

func (m *TCustomViewDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomViewDelegate class constructor
func NewCustomViewDelegate(events ICEFViewComponent) ICustomViewDelegate {
	var viewDelegatePtr uintptr // ICefViewDelegate
	r := customViewDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomViewDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewDelegatePtr)
	}
	return ret
}

var (
	customViewDelegateOnce   base.Once
	customViewDelegateImport *imports.Imports = nil
)

func customViewDelegateAPI() *imports.Imports {
	customViewDelegateOnce.Do(func() {
		customViewDelegateImport = api.NewDefaultImports()
		customViewDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomViewDelegate_Create", 0), // constructor NewCustomViewDelegate
		}
	})
	return customViewDelegateImport
}
