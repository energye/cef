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

// ICustomAccessibilityHandler Parent: ICEFAccessibilityHandlerOwn
type ICustomAccessibilityHandler interface {
	ICEFAccessibilityHandlerOwn
	SetOnTreeChange(fn TOnAccessibilityEvent)     // property event
	SetOnLocationChange(fn TOnAccessibilityEvent) // property event
	AsIntfAccessibilityHandler() uintptr
}

type TCustomAccessibilityHandler struct {
	TCEFAccessibilityHandlerOwn
}

func (m *TCustomAccessibilityHandler) SetOnTreeChange(fn TOnAccessibilityEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAccessibilityEvent(fn)
	base.SetEvent(m, 1, customAccessibilityHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAccessibilityHandler) SetOnLocationChange(fn TOnAccessibilityEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAccessibilityEvent(fn)
	base.SetEvent(m, 2, customAccessibilityHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAccessibilityHandler) AsIntfAccessibilityHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomAccessibilityHandler class constructor
func NewCustomAccessibilityHandler() ICustomAccessibilityHandler {
	var accessibilityHandlerPtr uintptr // ICefAccessibilityHandler
	r := customAccessibilityHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&accessibilityHandlerPtr)))
	ret := AsCustomAccessibilityHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, accessibilityHandlerPtr)
	}
	return ret
}

var (
	customAccessibilityHandlerOnce   base.Once
	customAccessibilityHandlerImport *imports.Imports = nil
)

func customAccessibilityHandlerAPI() *imports.Imports {
	customAccessibilityHandlerOnce.Do(func() {
		customAccessibilityHandlerImport = api.NewDefaultImports()
		customAccessibilityHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomAccessibilityHandler_Create", 0), // constructor NewCustomAccessibilityHandler
			/* 1 */ imports.NewTable("TCustomAccessibilityHandler_OnTreeChange", 0), // event OnTreeChange
			/* 2 */ imports.NewTable("TCustomAccessibilityHandler_OnLocationChange", 0), // event OnLocationChange
		}
	})
	return customAccessibilityHandlerImport
}
