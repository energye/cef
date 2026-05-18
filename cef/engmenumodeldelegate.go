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

// IEngMenuModelDelegate Parent: ICefMenuModelDelegateOwn
type IEngMenuModelDelegate interface {
	ICefMenuModelDelegateOwn
	SetOnMenuModelFormatLabel(fn TOnMenuModelFormatLabelEvent)                     // property event
	SetOnMenuModelExecuteCommand(fn TOnMenuModelExecuteCommandEvent)               // property event
	SetOnMenuModelMouseOutsideMenu(fn TOnMenuModelMouseOutsideMenuEvent)           // property event
	SetOnMenuModelUnhandledOpenSubmenu(fn TOnMenuModelUnhandledOpenSubmenuEvent)   // property event
	SetOnMenuModelUnhandledCloseSubmenu(fn TOnMenuModelUnhandledCloseSubmenuEvent) // property event
	SetOnMenuModelMenuWillShow(fn TOnMenuModelMenuWillShowEvent)                   // property event
	SetOnMenuModelMenuClosed(fn TOnMenuModelMenuClosedEvent)                       // property event
	AsIntfMenuModelDelegate() uintptr
}

type TEngMenuModelDelegate struct {
	TCefMenuModelDelegateOwn
}

func (m *TEngMenuModelDelegate) SetOnMenuModelFormatLabel(fn TOnMenuModelFormatLabelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelFormatLabelEvent(fn)
	base.SetEvent(m, 1, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelExecuteCommand(fn TOnMenuModelExecuteCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelExecuteCommandEvent(fn)
	base.SetEvent(m, 2, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelMouseOutsideMenu(fn TOnMenuModelMouseOutsideMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelMouseOutsideMenuEvent(fn)
	base.SetEvent(m, 3, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelUnhandledOpenSubmenu(fn TOnMenuModelUnhandledOpenSubmenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelUnhandledOpenSubmenuEvent(fn)
	base.SetEvent(m, 4, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelUnhandledCloseSubmenu(fn TOnMenuModelUnhandledCloseSubmenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelUnhandledCloseSubmenuEvent(fn)
	base.SetEvent(m, 5, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelMenuWillShow(fn TOnMenuModelMenuWillShowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelMenuWillShowEvent(fn)
	base.SetEvent(m, 6, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) SetOnMenuModelMenuClosed(fn TOnMenuModelMenuClosedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuModelMenuClosedEvent(fn)
	base.SetEvent(m, 7, engMenuModelDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMenuModelDelegate) AsIntfMenuModelDelegate() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngMenuModelDelegate class constructor
func NewEngMenuModelDelegate() IEngMenuModelDelegate {
	var menuModelDelegatePtr uintptr // ICefMenuModelDelegate
	r := engMenuModelDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&menuModelDelegatePtr)))
	ret := AsEngMenuModelDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, menuModelDelegatePtr)
	}
	return ret
}

var (
	engMenuModelDelegateOnce   base.Once
	engMenuModelDelegateImport *imports.Imports = nil
)

func engMenuModelDelegateAPI() *imports.Imports {
	engMenuModelDelegateOnce.Do(func() {
		engMenuModelDelegateImport = api.NewDefaultImports()
		engMenuModelDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMenuModelDelegate_Create", 0), // constructor NewEngMenuModelDelegate
			/* 1 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelFormatLabel", 0), // event OnMenuModelFormatLabel
			/* 2 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelExecuteCommand", 0), // event OnMenuModelExecuteCommand
			/* 3 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelMouseOutsideMenu", 0), // event OnMenuModelMouseOutsideMenu
			/* 4 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelUnhandledOpenSubmenu", 0), // event OnMenuModelUnhandledOpenSubmenu
			/* 5 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelUnhandledCloseSubmenu", 0), // event OnMenuModelUnhandledCloseSubmenu
			/* 6 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelMenuWillShow", 0), // event OnMenuModelMenuWillShow
			/* 7 */ imports.NewTable("TEngMenuModelDelegate_OnMenuModelMenuClosed", 0), // event OnMenuModelMenuClosed
		}
	})
	return engMenuModelDelegateImport
}
