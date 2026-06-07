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

// IEngKeyboardHandler Parent: ICefKeyboardHandlerOwn
type IEngKeyboardHandler interface {
	ICefKeyboardHandlerOwn
	SetOnKeyboardPreKey(fn TOnKeyboardPreKeyEvent) // property event
	SetOnKeyboardKey(fn TOnKeyboardKeyEvent)       // property event
	AsIntfKeyboardHandler() uintptr
}

type TEngKeyboardHandler struct {
	TCefKeyboardHandlerOwn
}

func (m *TEngKeyboardHandler) SetOnKeyboardPreKey(fn TOnKeyboardPreKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnKeyboardPreKeyEvent(fn)
	base.SetEvent(m, 1, engKeyboardHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngKeyboardHandler) SetOnKeyboardKey(fn TOnKeyboardKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnKeyboardKeyEvent(fn)
	base.SetEvent(m, 2, engKeyboardHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngKeyboardHandler) AsIntfKeyboardHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngKeyboardHandler class constructor
func NewEngKeyboardHandler() IEngKeyboardHandler {
	var keyboardHandlerPtr uintptr // ICefKeyboardHandler
	r := engKeyboardHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&keyboardHandlerPtr)))
	ret := AsEngKeyboardHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, keyboardHandlerPtr)
	}
	return ret
}

var (
	engKeyboardHandlerOnce   base.Once
	engKeyboardHandlerImport *imports.Imports = nil
)

func engKeyboardHandlerAPI() *imports.Imports {
	engKeyboardHandlerOnce.Do(func() {
		engKeyboardHandlerImport = api.NewDefaultImports()
		engKeyboardHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngKeyboardHandler_Create", 0), // constructor NewEngKeyboardHandler
			/* 1 */ imports.NewTable("TEngKeyboardHandler_OnKeyboardPreKey", 0), // event OnKeyboardPreKey
			/* 2 */ imports.NewTable("TEngKeyboardHandler_OnKeyboardKey", 0), // event OnKeyboardKey
		}
	})
	return engKeyboardHandlerImport
}
