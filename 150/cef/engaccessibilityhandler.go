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

// IEngAccessibilityHandler Parent: ICEFAccessibilityHandlerOwn
type IEngAccessibilityHandler interface {
	ICEFAccessibilityHandlerOwn
	SetOnAccessibilityAccessibilityTreeChange(fn TOnAccessibilityAccessibilityTreeChangeEvent)         // property event
	SetOnAccessibilityAccessibilityLocationChange(fn TOnAccessibilityAccessibilityLocationChangeEvent) // property event
	AsIntfAccessibilityHandler() uintptr
}

type TEngAccessibilityHandler struct {
	TCEFAccessibilityHandlerOwn
}

func (m *TEngAccessibilityHandler) SetOnAccessibilityAccessibilityTreeChange(fn TOnAccessibilityAccessibilityTreeChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAccessibilityAccessibilityTreeChangeEvent(fn)
	base.SetEvent(m, 1, engAccessibilityHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAccessibilityHandler) SetOnAccessibilityAccessibilityLocationChange(fn TOnAccessibilityAccessibilityLocationChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAccessibilityAccessibilityLocationChangeEvent(fn)
	base.SetEvent(m, 2, engAccessibilityHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAccessibilityHandler) AsIntfAccessibilityHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngAccessibilityHandler class constructor
func NewEngAccessibilityHandler() IEngAccessibilityHandler {
	var accessibilityHandlerPtr uintptr // ICefAccessibilityHandler
	r := engAccessibilityHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&accessibilityHandlerPtr)))
	ret := AsEngAccessibilityHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, accessibilityHandlerPtr)
	}
	return ret
}

var (
	engAccessibilityHandlerOnce   base.Once
	engAccessibilityHandlerImport *imports.Imports = nil
)

func engAccessibilityHandlerAPI() *imports.Imports {
	engAccessibilityHandlerOnce.Do(func() {
		engAccessibilityHandlerImport = api.NewDefaultImports()
		engAccessibilityHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngAccessibilityHandler_Create", 0), // constructor NewEngAccessibilityHandler
			/* 1 */ imports.NewTable("TEngAccessibilityHandler_OnAccessibilityAccessibilityTreeChange", 0), // event OnAccessibilityAccessibilityTreeChange
			/* 2 */ imports.NewTable("TEngAccessibilityHandler_OnAccessibilityAccessibilityLocationChange", 0), // event OnAccessibilityAccessibilityLocationChange
		}
	})
	return engAccessibilityHandlerImport
}
