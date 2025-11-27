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

// IEngFrameHandler Parent: ICefFrameHandlerOwn
type IEngFrameHandler interface {
	ICefFrameHandlerOwn
	SetOnFrameFrameCreated(fn TOnFrameFrameCreatedEvent)         // property event
	SetOnFrameFrameAttached(fn TOnFrameFrameAttachedEvent)       // property event
	SetOnFrameFrameDetached(fn TOnFrameFrameDetachedEvent)       // property event
	SetOnFrameMainFrameChanged(fn TOnFrameMainFrameChangedEvent) // property event
	AsIntfFrameHandler() uintptr
}

type TEngFrameHandler struct {
	TCefFrameHandlerOwn
}

func (m *TEngFrameHandler) SetOnFrameFrameCreated(fn TOnFrameFrameCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameFrameCreatedEvent(fn)
	base.SetEvent(m, 1, engFrameHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFrameHandler) SetOnFrameFrameAttached(fn TOnFrameFrameAttachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameFrameAttachedEvent(fn)
	base.SetEvent(m, 2, engFrameHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFrameHandler) SetOnFrameFrameDetached(fn TOnFrameFrameDetachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameFrameDetachedEvent(fn)
	base.SetEvent(m, 3, engFrameHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFrameHandler) SetOnFrameMainFrameChanged(fn TOnFrameMainFrameChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameMainFrameChangedEvent(fn)
	base.SetEvent(m, 4, engFrameHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFrameHandler) AsIntfFrameHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngFrameHandler class constructor
func NewEngFrameHandler() IEngFrameHandler {
	var frameHandlerPtr uintptr // ICefFrameHandler
	r := engFrameHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&frameHandlerPtr)))
	ret := AsEngFrameHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, frameHandlerPtr)
	}
	return ret
}

var (
	engFrameHandlerOnce   base.Once
	engFrameHandlerImport *imports.Imports = nil
)

func engFrameHandlerAPI() *imports.Imports {
	engFrameHandlerOnce.Do(func() {
		engFrameHandlerImport = api.NewDefaultImports()
		engFrameHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngFrameHandler_Create", 0), // constructor NewEngFrameHandler
			/* 1 */ imports.NewTable("TEngFrameHandler_OnFrameFrameCreated", 0), // event OnFrameFrameCreated
			/* 2 */ imports.NewTable("TEngFrameHandler_OnFrameFrameAttached", 0), // event OnFrameFrameAttached
			/* 3 */ imports.NewTable("TEngFrameHandler_OnFrameFrameDetached", 0), // event OnFrameFrameDetached
			/* 4 */ imports.NewTable("TEngFrameHandler_OnFrameMainFrameChanged", 0), // event OnFrameMainFrameChanged
		}
	})
	return engFrameHandlerImport
}
