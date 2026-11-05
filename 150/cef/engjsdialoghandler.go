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

// IEngJsDialogHandler Parent: ICefJsDialogHandlerOwn
type IEngJsDialogHandler interface {
	ICefJsDialogHandlerOwn
	SetOnJsDialogJsdialog(fn TOnJsDialogJsdialogEvent)                     // property event
	SetOnJsDialogBeforeUnloadDialog(fn TOnJsDialogBeforeUnloadDialogEvent) // property event
	SetOnJsDialogResetDialogState(fn TOnJsDialogResetDialogStateEvent)     // property event
	SetOnJsDialogDialogClosed(fn TOnJsDialogDialogClosedEvent)             // property event
	AsIntfJsDialogHandler() uintptr
}

type TEngJsDialogHandler struct {
	TCefJsDialogHandlerOwn
}

func (m *TEngJsDialogHandler) SetOnJsDialogJsdialog(fn TOnJsDialogJsdialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsDialogJsdialogEvent(fn)
	base.SetEvent(m, 1, engJsDialogHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngJsDialogHandler) SetOnJsDialogBeforeUnloadDialog(fn TOnJsDialogBeforeUnloadDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsDialogBeforeUnloadDialogEvent(fn)
	base.SetEvent(m, 2, engJsDialogHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngJsDialogHandler) SetOnJsDialogResetDialogState(fn TOnJsDialogResetDialogStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsDialogResetDialogStateEvent(fn)
	base.SetEvent(m, 3, engJsDialogHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngJsDialogHandler) SetOnJsDialogDialogClosed(fn TOnJsDialogDialogClosedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsDialogDialogClosedEvent(fn)
	base.SetEvent(m, 4, engJsDialogHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngJsDialogHandler) AsIntfJsDialogHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngJsDialogHandler class constructor
func NewEngJsDialogHandler() IEngJsDialogHandler {
	var jsDialogHandlerPtr uintptr // ICefJsDialogHandler
	r := engJsDialogHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&jsDialogHandlerPtr)))
	ret := AsEngJsDialogHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, jsDialogHandlerPtr)
	}
	return ret
}

var (
	engJsDialogHandlerOnce   base.Once
	engJsDialogHandlerImport *imports.Imports = nil
)

func engJsDialogHandlerAPI() *imports.Imports {
	engJsDialogHandlerOnce.Do(func() {
		engJsDialogHandlerImport = api.NewDefaultImports()
		engJsDialogHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngJsDialogHandler_Create", 0), // constructor NewEngJsDialogHandler
			/* 1 */ imports.NewTable("TEngJsDialogHandler_OnJsDialogJsdialog", 0), // event OnJsDialogJsdialog
			/* 2 */ imports.NewTable("TEngJsDialogHandler_OnJsDialogBeforeUnloadDialog", 0), // event OnJsDialogBeforeUnloadDialog
			/* 3 */ imports.NewTable("TEngJsDialogHandler_OnJsDialogResetDialogState", 0), // event OnJsDialogResetDialogState
			/* 4 */ imports.NewTable("TEngJsDialogHandler_OnJsDialogDialogClosed", 0), // event OnJsDialogDialogClosed
		}
	})
	return engJsDialogHandlerImport
}
