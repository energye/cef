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

// IEngContextMenuHandler Parent: ICefContextMenuHandlerOwn
type IEngContextMenuHandler interface {
	ICefContextMenuHandlerOwn
	SetOnContextMenuRunContextMenu(fn TOnContextMenuRunContextMenuEvent)             // property event
	SetOnContextMenuContextMenuCommand(fn TOnContextMenuContextMenuCommandEvent)     // property event
	SetOnContextMenuRunQuickMenu(fn TOnContextMenuRunQuickMenuEvent)                 // property event
	SetOnContextMenuQuickMenuCommand(fn TOnContextMenuQuickMenuCommandEvent)         // property event
	SetOnContextMenuBeforeContextMenu(fn TOnContextMenuBeforeContextMenuEvent)       // property event
	SetOnContextMenuContextMenuDismissed(fn TOnContextMenuContextMenuDismissedEvent) // property event
	SetOnContextMenuQuickMenuDismissed(fn TOnContextMenuQuickMenuDismissedEvent)     // property event
	AsIntfContextMenuHandler() uintptr
}

type TEngContextMenuHandler struct {
	TCefContextMenuHandlerOwn
}

func (m *TEngContextMenuHandler) SetOnContextMenuRunContextMenu(fn TOnContextMenuRunContextMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuRunContextMenuEvent(fn)
	base.SetEvent(m, 1, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuContextMenuCommand(fn TOnContextMenuContextMenuCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuContextMenuCommandEvent(fn)
	base.SetEvent(m, 2, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuRunQuickMenu(fn TOnContextMenuRunQuickMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuRunQuickMenuEvent(fn)
	base.SetEvent(m, 3, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuQuickMenuCommand(fn TOnContextMenuQuickMenuCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuQuickMenuCommandEvent(fn)
	base.SetEvent(m, 4, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuBeforeContextMenu(fn TOnContextMenuBeforeContextMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuBeforeContextMenuEvent(fn)
	base.SetEvent(m, 5, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuContextMenuDismissed(fn TOnContextMenuContextMenuDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuContextMenuDismissedEvent(fn)
	base.SetEvent(m, 6, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) SetOnContextMenuQuickMenuDismissed(fn TOnContextMenuQuickMenuDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuQuickMenuDismissedEvent(fn)
	base.SetEvent(m, 7, engContextMenuHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngContextMenuHandler) AsIntfContextMenuHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngContextMenuHandler class constructor
func NewEngContextMenuHandler() IEngContextMenuHandler {
	var contextMenuHandlerPtr uintptr // ICefContextMenuHandler
	r := engContextMenuHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&contextMenuHandlerPtr)))
	ret := AsEngContextMenuHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, contextMenuHandlerPtr)
	}
	return ret
}

var (
	engContextMenuHandlerOnce   base.Once
	engContextMenuHandlerImport *imports.Imports = nil
)

func engContextMenuHandlerAPI() *imports.Imports {
	engContextMenuHandlerOnce.Do(func() {
		engContextMenuHandlerImport = api.NewDefaultImports()
		engContextMenuHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngContextMenuHandler_Create", 0), // constructor NewEngContextMenuHandler
			/* 1 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuRunContextMenu", 0), // event OnContextMenuRunContextMenu
			/* 2 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuContextMenuCommand", 0), // event OnContextMenuContextMenuCommand
			/* 3 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuRunQuickMenu", 0), // event OnContextMenuRunQuickMenu
			/* 4 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuQuickMenuCommand", 0), // event OnContextMenuQuickMenuCommand
			/* 5 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuBeforeContextMenu", 0), // event OnContextMenuBeforeContextMenu
			/* 6 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuContextMenuDismissed", 0), // event OnContextMenuContextMenuDismissed
			/* 7 */ imports.NewTable("TEngContextMenuHandler_OnContextMenuQuickMenuDismissed", 0), // event OnContextMenuQuickMenuDismissed
		}
	})
	return engContextMenuHandlerImport
}
