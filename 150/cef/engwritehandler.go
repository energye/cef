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

// IEngWriteHandler Parent: ICefWriteHandlerOwn
type IEngWriteHandler interface {
	ICefWriteHandlerOwn
	SetOnWriteWrite(fn TOnWriteWriteEvent)       // property event
	SetOnWriteSeek(fn TOnWriteSeekEvent)         // property event
	SetOnWriteEll(fn TOnWriteEllEvent)           // property event
	SetOnWriteFlush(fn TOnWriteFlushEvent)       // property event
	SetOnWriteMayBlock(fn TOnWriteMayBlockEvent) // property event
	AsIntfWriteHandler() uintptr
}

type TEngWriteHandler struct {
	TCefWriteHandlerOwn
}

func (m *TEngWriteHandler) SetOnWriteWrite(fn TOnWriteWriteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWriteWriteEvent(fn)
	base.SetEvent(m, 1, engWriteHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWriteHandler) SetOnWriteSeek(fn TOnWriteSeekEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWriteSeekEvent(fn)
	base.SetEvent(m, 2, engWriteHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWriteHandler) SetOnWriteEll(fn TOnWriteEllEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWriteEllEvent(fn)
	base.SetEvent(m, 3, engWriteHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWriteHandler) SetOnWriteFlush(fn TOnWriteFlushEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWriteFlushEvent(fn)
	base.SetEvent(m, 4, engWriteHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWriteHandler) SetOnWriteMayBlock(fn TOnWriteMayBlockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWriteMayBlockEvent(fn)
	base.SetEvent(m, 5, engWriteHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWriteHandler) AsIntfWriteHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngWriteHandler class constructor
func NewEngWriteHandler() IEngWriteHandler {
	var writeHandlerPtr uintptr // ICefWriteHandler
	r := engWriteHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&writeHandlerPtr)))
	ret := AsEngWriteHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, writeHandlerPtr)
	}
	return ret
}

var (
	engWriteHandlerOnce   base.Once
	engWriteHandlerImport *imports.Imports = nil
)

func engWriteHandlerAPI() *imports.Imports {
	engWriteHandlerOnce.Do(func() {
		engWriteHandlerImport = api.NewDefaultImports()
		engWriteHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngWriteHandler_Create", 0), // constructor NewEngWriteHandler
			/* 1 */ imports.NewTable("TEngWriteHandler_OnWriteWrite", 0), // event OnWriteWrite
			/* 2 */ imports.NewTable("TEngWriteHandler_OnWriteSeek", 0), // event OnWriteSeek
			/* 3 */ imports.NewTable("TEngWriteHandler_OnWriteEll", 0), // event OnWriteEll
			/* 4 */ imports.NewTable("TEngWriteHandler_OnWriteFlush", 0), // event OnWriteFlush
			/* 5 */ imports.NewTable("TEngWriteHandler_OnWriteMayBlock", 0), // event OnWriteMayBlock
		}
	})
	return engWriteHandlerImport
}
