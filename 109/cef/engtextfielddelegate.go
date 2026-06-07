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

// IEngTextfieldDelegate Parent: ICefTextfieldDelegateOwn
type IEngTextfieldDelegate interface {
	ICefTextfieldDelegateOwn
	SetOnTextfieldKey(fn TOnTextfieldKeyEvent)                         // property event
	SetOnTextfieldAfterUserAction(fn TOnTextfieldAfterUserActionEvent) // property event
	AsIntfTextfieldDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngTextfieldDelegate struct {
	TCefTextfieldDelegateOwn
}

func (m *TEngTextfieldDelegate) SetOnTextfieldKey(fn TOnTextfieldKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextfieldKeyEvent(fn)
	base.SetEvent(m, 1, engTextfieldDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngTextfieldDelegate) SetOnTextfieldAfterUserAction(fn TOnTextfieldAfterUserActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextfieldAfterUserActionEvent(fn)
	base.SetEvent(m, 2, engTextfieldDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngTextfieldDelegate) AsIntfTextfieldDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngTextfieldDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewEngTextfieldDelegate class constructor
func NewEngTextfieldDelegate() IEngTextfieldDelegate {
	var textfieldDelegatePtr uintptr // ICefTextfieldDelegate
	var viewDelegatePtr uintptr      // ICefViewDelegate
	r := engTextfieldDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&textfieldDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngTextfieldDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	engTextfieldDelegateOnce   base.Once
	engTextfieldDelegateImport *imports.Imports = nil
)

func engTextfieldDelegateAPI() *imports.Imports {
	engTextfieldDelegateOnce.Do(func() {
		engTextfieldDelegateImport = api.NewDefaultImports()
		engTextfieldDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngTextfieldDelegate_Create", 0), // constructor NewEngTextfieldDelegate
			/* 1 */ imports.NewTable("TEngTextfieldDelegate_OnTextfieldKey", 0), // event OnTextfieldKey
			/* 2 */ imports.NewTable("TEngTextfieldDelegate_OnTextfieldAfterUserAction", 0), // event OnTextfieldAfterUserAction
		}
	})
	return engTextfieldDelegateImport
}
