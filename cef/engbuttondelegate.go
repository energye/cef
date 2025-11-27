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

// IEngButtonDelegate Parent: ICefButtonDelegateOwn
type IEngButtonDelegate interface {
	ICefButtonDelegateOwn
	SetOnButtonButtonPressed(fn TOnButtonButtonPressedEvent)           // property event
	SetOnButtonButtonStateChanged(fn TOnButtonButtonStateChangedEvent) // property event
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngButtonDelegate struct {
	TCefButtonDelegateOwn
}

func (m *TEngButtonDelegate) SetOnButtonButtonPressed(fn TOnButtonButtonPressedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnButtonButtonPressedEvent(fn)
	base.SetEvent(m, 1, engButtonDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngButtonDelegate) SetOnButtonButtonStateChanged(fn TOnButtonButtonStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnButtonButtonStateChangedEvent(fn)
	base.SetEvent(m, 2, engButtonDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngButtonDelegate) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngButtonDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewEngButtonDelegate class constructor
func NewEngButtonDelegate() IEngButtonDelegate {
	var buttonDelegatePtr uintptr // ICefButtonDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := engButtonDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngButtonDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	engButtonDelegateOnce   base.Once
	engButtonDelegateImport *imports.Imports = nil
)

func engButtonDelegateAPI() *imports.Imports {
	engButtonDelegateOnce.Do(func() {
		engButtonDelegateImport = api.NewDefaultImports()
		engButtonDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngButtonDelegate_Create", 0), // constructor NewEngButtonDelegate
			/* 1 */ imports.NewTable("TEngButtonDelegate_OnButtonButtonPressed", 0), // event OnButtonButtonPressed
			/* 2 */ imports.NewTable("TEngButtonDelegate_OnButtonButtonStateChanged", 0), // event OnButtonButtonStateChanged
		}
	})
	return engButtonDelegateImport
}
