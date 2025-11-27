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

// ICefBrowserViewDelegateOwn Parent: ICefViewDelegateOwn
type ICefBrowserViewDelegateOwn interface {
	ICefViewDelegateOwn
	AsIntfBrowserViewDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefBrowserViewDelegateOwn struct {
	TCefViewDelegateOwn
}

func (m *TCefBrowserViewDelegateOwn) AsIntfBrowserViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefBrowserViewDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewBrowserViewDelegateOwn class constructor
func NewBrowserViewDelegateOwn() ICefBrowserViewDelegateOwn {
	var browserViewDelegatePtr uintptr // ICefBrowserViewDelegate
	var viewDelegatePtr uintptr        // ICefViewDelegate
	r := cefBrowserViewDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&browserViewDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefBrowserViewDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefBrowserViewDelegateOwnOnce   base.Once
	cefBrowserViewDelegateOwnImport *imports.Imports = nil
)

func cefBrowserViewDelegateOwnAPI() *imports.Imports {
	cefBrowserViewDelegateOwnOnce.Do(func() {
		cefBrowserViewDelegateOwnImport = api.NewDefaultImports()
		cefBrowserViewDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserViewDelegateOwn_Create", 0), // constructor NewBrowserViewDelegateOwn
		}
	})
	return cefBrowserViewDelegateOwnImport
}
