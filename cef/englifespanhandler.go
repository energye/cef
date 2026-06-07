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

// IEngLifeSpanHandler Parent: ICefLifeSpanHandlerOwn
type IEngLifeSpanHandler interface {
	ICefLifeSpanHandlerOwn
	SetOnLifeSpanBeforePopup(fn TOnLifeSpanBeforePopupEvent)                 // property event
	SetOnLifeSpanDoClose(fn TOnLifeSpanDoCloseEvent)                         // property event
	SetOnLifeSpanBeforePopupAborted(fn TOnLifeSpanBeforePopupAbortedEvent)   // property event
	SetOnLifeSpanBeforeDevToolsPopup(fn TOnLifeSpanBeforeDevToolsPopupEvent) // property event
	SetOnLifeSpanAfterCreated(fn TOnLifeSpanAfterCreatedEvent)               // property event
	SetOnLifeSpanBeforeClose(fn TOnLifeSpanBeforeCloseEvent)                 // property event
	AsIntfLifeSpanHandler() uintptr
}

type TEngLifeSpanHandler struct {
	TCefLifeSpanHandlerOwn
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanBeforePopup(fn TOnLifeSpanBeforePopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanBeforePopupEvent(fn)
	base.SetEvent(m, 1, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanDoClose(fn TOnLifeSpanDoCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanDoCloseEvent(fn)
	base.SetEvent(m, 2, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanBeforePopupAborted(fn TOnLifeSpanBeforePopupAbortedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanBeforePopupAbortedEvent(fn)
	base.SetEvent(m, 3, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanBeforeDevToolsPopup(fn TOnLifeSpanBeforeDevToolsPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanBeforeDevToolsPopupEvent(fn)
	base.SetEvent(m, 4, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanAfterCreated(fn TOnLifeSpanAfterCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanAfterCreatedEvent(fn)
	base.SetEvent(m, 5, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) SetOnLifeSpanBeforeClose(fn TOnLifeSpanBeforeCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLifeSpanBeforeCloseEvent(fn)
	base.SetEvent(m, 6, engLifeSpanHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLifeSpanHandler) AsIntfLifeSpanHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngLifeSpanHandler class constructor
func NewEngLifeSpanHandler() IEngLifeSpanHandler {
	var lifeSpanHandlerPtr uintptr // ICefLifeSpanHandler
	r := engLifeSpanHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&lifeSpanHandlerPtr)))
	ret := AsEngLifeSpanHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, lifeSpanHandlerPtr)
	}
	return ret
}

var (
	engLifeSpanHandlerOnce   base.Once
	engLifeSpanHandlerImport *imports.Imports = nil
)

func engLifeSpanHandlerAPI() *imports.Imports {
	engLifeSpanHandlerOnce.Do(func() {
		engLifeSpanHandlerImport = api.NewDefaultImports()
		engLifeSpanHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngLifeSpanHandler_Create", 0), // constructor NewEngLifeSpanHandler
			/* 1 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanBeforePopup", 0), // event OnLifeSpanBeforePopup
			/* 2 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanDoClose", 0), // event OnLifeSpanDoClose
			/* 3 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanBeforePopupAborted", 0), // event OnLifeSpanBeforePopupAborted
			/* 4 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanBeforeDevToolsPopup", 0), // event OnLifeSpanBeforeDevToolsPopup
			/* 5 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanAfterCreated", 0), // event OnLifeSpanAfterCreated
			/* 6 */ imports.NewTable("TEngLifeSpanHandler_OnLifeSpanBeforeClose", 0), // event OnLifeSpanBeforeClose
		}
	})
	return engLifeSpanHandlerImport
}
