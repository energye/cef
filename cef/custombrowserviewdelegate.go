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

// ICustomBrowserViewDelegate Parent: ICefBrowserViewDelegateOwn
type ICustomBrowserViewDelegate interface {
	ICefBrowserViewDelegateOwn
	AsIntfBrowserViewDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomBrowserViewDelegate struct {
	TCefBrowserViewDelegateOwn
}

func (m *TCustomBrowserViewDelegate) AsIntfBrowserViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomBrowserViewDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewCustomBrowserViewDelegate class constructor
func NewCustomBrowserViewDelegate(events ICEFBrowserViewComponent) ICustomBrowserViewDelegate {
	var browserViewDelegatePtr uintptr // ICefBrowserViewDelegate
	var viewDelegatePtr uintptr        // ICefViewDelegate
	r := customBrowserViewDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&browserViewDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomBrowserViewDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	customBrowserViewDelegateOnce   base.Once
	customBrowserViewDelegateImport *imports.Imports = nil
)

func customBrowserViewDelegateAPI() *imports.Imports {
	customBrowserViewDelegateOnce.Do(func() {
		customBrowserViewDelegateImport = api.NewDefaultImports()
		customBrowserViewDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomBrowserViewDelegate_Create", 0), // constructor NewCustomBrowserViewDelegate
		}
	})
	return customBrowserViewDelegateImport
}
