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

// ICefMenuButtonDelegateOwn Parent: ICefButtonDelegateOwn
type ICefMenuButtonDelegateOwn interface {
	ICefButtonDelegateOwn
	AsIntfMenuButtonDelegate() uintptr
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefMenuButtonDelegateOwn struct {
	TCefButtonDelegateOwn
}

func (m *TCefMenuButtonDelegateOwn) AsIntfMenuButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefMenuButtonDelegateOwn) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefMenuButtonDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewMenuButtonDelegateOwn class constructor
func NewMenuButtonDelegateOwn() ICefMenuButtonDelegateOwn {
	var menuButtonDelegatePtr uintptr // ICefMenuButtonDelegate
	var buttonDelegatePtr uintptr     // ICefButtonDelegate
	var viewDelegatePtr uintptr       // ICefViewDelegate
	r := cefMenuButtonDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&menuButtonDelegatePtr)), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefMenuButtonDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, menuButtonDelegatePtr)
		intf.SetIntfPointer(1, buttonDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	cefMenuButtonDelegateOwnOnce   base.Once
	cefMenuButtonDelegateOwnImport *imports.Imports = nil
)

func cefMenuButtonDelegateOwnAPI() *imports.Imports {
	cefMenuButtonDelegateOwnOnce.Do(func() {
		cefMenuButtonDelegateOwnImport = api.NewDefaultImports()
		cefMenuButtonDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMenuButtonDelegateOwn_Create", 0), // constructor NewMenuButtonDelegateOwn
		}
	})
	return cefMenuButtonDelegateOwnImport
}
