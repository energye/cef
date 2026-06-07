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

// IEngAudioHandler Parent: ICefAudioHandlerOwn
type IEngAudioHandler interface {
	ICefAudioHandlerOwn
	SetOnAudioGetAudioParameters(fn TOnAudioGetAudioParametersEvent) // property event
	SetOnAudioAudioStreamStarted(fn TOnAudioAudioStreamStartedEvent) // property event
	SetOnAudioAudioStreamPacket(fn TOnAudioAudioStreamPacketEvent)   // property event
	SetOnAudioAudioStreamStopped(fn TOnAudioAudioStreamStoppedEvent) // property event
	SetOnAudioAudioStreamError(fn TOnAudioAudioStreamErrorEvent)     // property event
	AsIntfAudioHandler() uintptr
}

type TEngAudioHandler struct {
	TCefAudioHandlerOwn
}

func (m *TEngAudioHandler) SetOnAudioGetAudioParameters(fn TOnAudioGetAudioParametersEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioGetAudioParametersEvent(fn)
	base.SetEvent(m, 1, engAudioHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAudioHandler) SetOnAudioAudioStreamStarted(fn TOnAudioAudioStreamStartedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioAudioStreamStartedEvent(fn)
	base.SetEvent(m, 2, engAudioHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAudioHandler) SetOnAudioAudioStreamPacket(fn TOnAudioAudioStreamPacketEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioAudioStreamPacketEvent(fn)
	base.SetEvent(m, 3, engAudioHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAudioHandler) SetOnAudioAudioStreamStopped(fn TOnAudioAudioStreamStoppedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioAudioStreamStoppedEvent(fn)
	base.SetEvent(m, 4, engAudioHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAudioHandler) SetOnAudioAudioStreamError(fn TOnAudioAudioStreamErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioAudioStreamErrorEvent(fn)
	base.SetEvent(m, 5, engAudioHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngAudioHandler) AsIntfAudioHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngAudioHandler class constructor
func NewEngAudioHandler() IEngAudioHandler {
	var audioHandlerPtr uintptr // ICefAudioHandler
	r := engAudioHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&audioHandlerPtr)))
	ret := AsEngAudioHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, audioHandlerPtr)
	}
	return ret
}

var (
	engAudioHandlerOnce   base.Once
	engAudioHandlerImport *imports.Imports = nil
)

func engAudioHandlerAPI() *imports.Imports {
	engAudioHandlerOnce.Do(func() {
		engAudioHandlerImport = api.NewDefaultImports()
		engAudioHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngAudioHandler_Create", 0), // constructor NewEngAudioHandler
			/* 1 */ imports.NewTable("TEngAudioHandler_OnAudioGetAudioParameters", 0), // event OnAudioGetAudioParameters
			/* 2 */ imports.NewTable("TEngAudioHandler_OnAudioAudioStreamStarted", 0), // event OnAudioAudioStreamStarted
			/* 3 */ imports.NewTable("TEngAudioHandler_OnAudioAudioStreamPacket", 0), // event OnAudioAudioStreamPacket
			/* 4 */ imports.NewTable("TEngAudioHandler_OnAudioAudioStreamStopped", 0), // event OnAudioAudioStreamStopped
			/* 5 */ imports.NewTable("TEngAudioHandler_OnAudioAudioStreamError", 0), // event OnAudioAudioStreamError
		}
	})
	return engAudioHandlerImport
}
