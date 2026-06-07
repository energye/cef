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

// IEngDragHandler Parent: ICefDragHandlerOwn
type IEngDragHandler interface {
	ICefDragHandlerOwn
	SetOnDragDragEnter(fn TOnDragDragEnterEvent)                             // property event
	SetOnDragDraggableRegionsChanged(fn TOnDragDraggableRegionsChangedEvent) // property event
	AsIntfDragHandler() uintptr
}

type TEngDragHandler struct {
	TCefDragHandlerOwn
}

func (m *TEngDragHandler) SetOnDragDragEnter(fn TOnDragDragEnterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDragDragEnterEvent(fn)
	base.SetEvent(m, 1, engDragHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDragHandler) SetOnDragDraggableRegionsChanged(fn TOnDragDraggableRegionsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDragDraggableRegionsChangedEvent(fn)
	base.SetEvent(m, 2, engDragHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDragHandler) AsIntfDragHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDragHandler class constructor
func NewEngDragHandler() IEngDragHandler {
	var dragHandlerPtr uintptr // ICefDragHandler
	r := engDragHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&dragHandlerPtr)))
	ret := AsEngDragHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dragHandlerPtr)
	}
	return ret
}

var (
	engDragHandlerOnce   base.Once
	engDragHandlerImport *imports.Imports = nil
)

func engDragHandlerAPI() *imports.Imports {
	engDragHandlerOnce.Do(func() {
		engDragHandlerImport = api.NewDefaultImports()
		engDragHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDragHandler_Create", 0), // constructor NewEngDragHandler
			/* 1 */ imports.NewTable("TEngDragHandler_OnDragDragEnter", 0), // event OnDragDragEnter
			/* 2 */ imports.NewTable("TEngDragHandler_OnDragDraggableRegionsChanged", 0), // event OnDragDraggableRegionsChanged
		}
	})
	return engDragHandlerImport
}
