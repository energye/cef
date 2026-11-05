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

// ICustomTextfieldDelegate Parent: ICefTextfieldDelegateOwn
type ICustomTextfieldDelegate interface {
	ICefTextfieldDelegateOwn
	AsIntfTextfieldDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCustomTextfieldDelegate struct {
	TCefTextfieldDelegateOwn
}

func (m *TCustomTextfieldDelegate) AsIntfTextfieldDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCustomTextfieldDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewCustomTextfieldDelegate class constructor
func NewCustomTextfieldDelegate(events ICEFTextfieldComponent) ICustomTextfieldDelegate {
	var textfieldDelegatePtr uintptr // ICefTextfieldDelegate
	var viewDelegatePtr uintptr      // ICefViewDelegate
	r := customTextfieldDelegateAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&textfieldDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCustomTextfieldDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	customTextfieldDelegateOnce   base.Once
	customTextfieldDelegateImport *imports.Imports = nil
)

func customTextfieldDelegateAPI() *imports.Imports {
	customTextfieldDelegateOnce.Do(func() {
		customTextfieldDelegateImport = api.NewDefaultImports()
		customTextfieldDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTextfieldDelegate_Create", 0), // constructor NewCustomTextfieldDelegate
		}
	})
	return customTextfieldDelegateImport
}
