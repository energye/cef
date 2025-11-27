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

// ICefButtonDelegateOwn Parent: ICefViewDelegateOwn
type ICefButtonDelegateOwn interface {
	ICefViewDelegateOwn
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefButtonDelegateOwn struct {
	TCefViewDelegateOwn
}

func (m *TCefButtonDelegateOwn) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefButtonDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewButtonDelegateOwn class constructor
func NewButtonDelegateOwn() ICefButtonDelegateOwn {
	var buttonDelegatePtr uintptr // ICefButtonDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := cefButtonDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefButtonDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefButtonDelegateOwnOnce   base.Once
	cefButtonDelegateOwnImport *imports.Imports = nil
)

func cefButtonDelegateOwnAPI() *imports.Imports {
	cefButtonDelegateOwnOnce.Do(func() {
		cefButtonDelegateOwnImport = api.NewDefaultImports()
		cefButtonDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefButtonDelegateOwn_Create", 0), // constructor NewButtonDelegateOwn
		}
	})
	return cefButtonDelegateOwnImport
}
