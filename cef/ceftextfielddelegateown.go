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

// ICefTextfieldDelegateOwn Parent: ICefViewDelegateOwn
type ICefTextfieldDelegateOwn interface {
	ICefViewDelegateOwn
	AsIntfTextfieldDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefTextfieldDelegateOwn struct {
	TCefViewDelegateOwn
}

func (m *TCefTextfieldDelegateOwn) AsIntfTextfieldDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefTextfieldDelegateOwn) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewTextfieldDelegateOwn class constructor
func NewTextfieldDelegateOwn() ICefTextfieldDelegateOwn {
	var textfieldDelegatePtr uintptr // ICefTextfieldDelegate
	var viewDelegatePtr uintptr      // ICefViewDelegate
	r := cefTextfieldDelegateOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&textfieldDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefTextfieldDelegateOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefTextfieldDelegateOwnOnce   base.Once
	cefTextfieldDelegateOwnImport *imports.Imports = nil
)

func cefTextfieldDelegateOwnAPI() *imports.Imports {
	cefTextfieldDelegateOwnOnce.Do(func() {
		cefTextfieldDelegateOwnImport = api.NewDefaultImports()
		cefTextfieldDelegateOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTextfieldDelegateOwn_Create", 0), // constructor NewTextfieldDelegateOwn
		}
	})
	return cefTextfieldDelegateOwnImport
}
