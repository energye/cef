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

// IEngCommandHandler Parent: ICefCommandHandlerOwn
type IEngCommandHandler interface {
	ICefCommandHandlerOwn
	SetOnCommandChromeCommand(fn TOnCommandChromeCommandEvent)                                 // property event
	SetOnCommandIsChromeAppMenuItemVisible(fn TOnCommandIsChromeAppMenuItemVisibleEvent)       // property event
	SetOnCommandIsChromeAppMenuItemEnabled(fn TOnCommandIsChromeAppMenuItemEnabledEvent)       // property event
	SetOnCommandIsChromePageActionIconVisible(fn TOnCommandIsChromePageActionIconVisibleEvent) // property event
	SetOnCommandIsChromeToolbarButtonVisible(fn TOnCommandIsChromeToolbarButtonVisibleEvent)   // property event
	AsIntfCommandHandler() uintptr
}

type TEngCommandHandler struct {
	TCefCommandHandlerOwn
}

func (m *TEngCommandHandler) SetOnCommandChromeCommand(fn TOnCommandChromeCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCommandChromeCommandEvent(fn)
	base.SetEvent(m, 1, engCommandHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCommandHandler) SetOnCommandIsChromeAppMenuItemVisible(fn TOnCommandIsChromeAppMenuItemVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCommandIsChromeAppMenuItemVisibleEvent(fn)
	base.SetEvent(m, 2, engCommandHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCommandHandler) SetOnCommandIsChromeAppMenuItemEnabled(fn TOnCommandIsChromeAppMenuItemEnabledEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCommandIsChromeAppMenuItemEnabledEvent(fn)
	base.SetEvent(m, 3, engCommandHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCommandHandler) SetOnCommandIsChromePageActionIconVisible(fn TOnCommandIsChromePageActionIconVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCommandIsChromePageActionIconVisibleEvent(fn)
	base.SetEvent(m, 4, engCommandHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCommandHandler) SetOnCommandIsChromeToolbarButtonVisible(fn TOnCommandIsChromeToolbarButtonVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCommandIsChromeToolbarButtonVisibleEvent(fn)
	base.SetEvent(m, 5, engCommandHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCommandHandler) AsIntfCommandHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngCommandHandler class constructor
func NewEngCommandHandler() IEngCommandHandler {
	var commandHandlerPtr uintptr // ICefCommandHandler
	r := engCommandHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&commandHandlerPtr)))
	ret := AsEngCommandHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, commandHandlerPtr)
	}
	return ret
}

var (
	engCommandHandlerOnce   base.Once
	engCommandHandlerImport *imports.Imports = nil
)

func engCommandHandlerAPI() *imports.Imports {
	engCommandHandlerOnce.Do(func() {
		engCommandHandlerImport = api.NewDefaultImports()
		engCommandHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngCommandHandler_Create", 0), // constructor NewEngCommandHandler
			/* 1 */ imports.NewTable("TEngCommandHandler_OnCommandChromeCommand", 0), // event OnCommandChromeCommand
			/* 2 */ imports.NewTable("TEngCommandHandler_OnCommandIsChromeAppMenuItemVisible", 0), // event OnCommandIsChromeAppMenuItemVisible
			/* 3 */ imports.NewTable("TEngCommandHandler_OnCommandIsChromeAppMenuItemEnabled", 0), // event OnCommandIsChromeAppMenuItemEnabled
			/* 4 */ imports.NewTable("TEngCommandHandler_OnCommandIsChromePageActionIconVisible", 0), // event OnCommandIsChromePageActionIconVisible
			/* 5 */ imports.NewTable("TEngCommandHandler_OnCommandIsChromeToolbarButtonVisible", 0), // event OnCommandIsChromeToolbarButtonVisible
		}
	})
	return engCommandHandlerImport
}
