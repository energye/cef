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

// IEngMenuButtonDelegate Parent: ICefMenuButtonDelegateOwn
type IEngMenuButtonDelegate interface {
	ICefMenuButtonDelegateOwn
	SetOnMenuButtonMenuButtonPressed(fn TOnMenuButtonMenuButtonPressedEvent) // property event
	AsIntfMenuButtonDelegate() uintptr
	AsIntfButtonDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngMenuButtonDelegate struct {
	TCefMenuButtonDelegateOwn
}

func (m *TEngMenuButtonDelegate) SetOnMenuButtonMenuButtonPressed(fn TOnMenuButtonMenuButtonPressedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuButtonMenuButtonPressedEvent(fn)
	base.SetEvent(m, 1, engMenuButtonDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuButtonDelegate) AsIntfMenuButtonDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngMenuButtonDelegate) AsIntfButtonDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TEngMenuButtonDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewEngMenuButtonDelegate class constructor
func NewEngMenuButtonDelegate() IEngMenuButtonDelegate {
	var menuButtonDelegatePtr uintptr // ICefMenuButtonDelegate
	var buttonDelegatePtr uintptr     // ICefButtonDelegate
	var viewDelegatePtr uintptr       // ICefViewDelegate
	r := engMenuButtonDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&menuButtonDelegatePtr)), uintptr(base.UnsafePointer(&buttonDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngMenuButtonDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, menuButtonDelegatePtr)
		intf.SetIntfPointer(1, buttonDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	engMenuButtonDelegateOnce   base.Once
	engMenuButtonDelegateImport *imports.Imports = nil
)

func engMenuButtonDelegateAPI() *imports.Imports {
	engMenuButtonDelegateOnce.Do(func() {
		engMenuButtonDelegateImport = api.NewDefaultImports()
		engMenuButtonDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMenuButtonDelegate_Create", 0), // constructor NewEngMenuButtonDelegate
			/* 1 */ imports.NewTable("TEngMenuButtonDelegate_OnMenuButtonMenuButtonPressed", 0), // event OnMenuButtonMenuButtonPressed
		}
	})
	return engMenuButtonDelegateImport
}
