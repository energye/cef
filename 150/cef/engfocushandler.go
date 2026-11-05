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

// IEngFocusHandler Parent: ICefFocusHandlerOwn
type IEngFocusHandler interface {
	ICefFocusHandlerOwn
	SetOnFocusSetFocus(fn TOnFocusSetFocusEvent)   // property event
	SetOnFocusTakeFocus(fn TOnFocusTakeFocusEvent) // property event
	SetOnFocusGotFocus(fn TOnFocusGotFocusEvent)   // property event
	AsIntfFocusHandler() uintptr
}

type TEngFocusHandler struct {
	TCefFocusHandlerOwn
}

func (m *TEngFocusHandler) SetOnFocusSetFocus(fn TOnFocusSetFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusSetFocusEvent(fn)
	base.SetEvent(m, 1, engFocusHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFocusHandler) SetOnFocusTakeFocus(fn TOnFocusTakeFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusTakeFocusEvent(fn)
	base.SetEvent(m, 2, engFocusHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFocusHandler) SetOnFocusGotFocus(fn TOnFocusGotFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusGotFocusEvent(fn)
	base.SetEvent(m, 3, engFocusHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFocusHandler) AsIntfFocusHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngFocusHandler class constructor
func NewEngFocusHandler() IEngFocusHandler {
	var focusHandlerPtr uintptr // ICefFocusHandler
	r := engFocusHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&focusHandlerPtr)))
	ret := AsEngFocusHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, focusHandlerPtr)
	}
	return ret
}

var (
	engFocusHandlerOnce   base.Once
	engFocusHandlerImport *imports.Imports = nil
)

func engFocusHandlerAPI() *imports.Imports {
	engFocusHandlerOnce.Do(func() {
		engFocusHandlerImport = api.NewDefaultImports()
		engFocusHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngFocusHandler_Create", 0), // constructor NewEngFocusHandler
			/* 1 */ imports.NewTable("TEngFocusHandler_OnFocusSetFocus", 0), // event OnFocusSetFocus
			/* 2 */ imports.NewTable("TEngFocusHandler_OnFocusTakeFocus", 0), // event OnFocusTakeFocus
			/* 3 */ imports.NewTable("TEngFocusHandler_OnFocusGotFocus", 0), // event OnFocusGotFocus
		}
	})
	return engFocusHandlerImport
}
