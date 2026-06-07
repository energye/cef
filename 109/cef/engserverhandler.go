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

// IEngServerHandler Parent: ICEFServerHandlerOwn
type IEngServerHandler interface {
	ICEFServerHandlerOwn
	SetOnServerServerCreated(fn TOnServerServerCreatedEvent)           // property event
	SetOnServerServerDestroyed(fn TOnServerServerDestroyedEvent)       // property event
	SetOnServerClientConnected(fn TOnServerClientConnectedEvent)       // property event
	SetOnServerClientDisconnected(fn TOnServerClientDisconnectedEvent) // property event
	SetOnServerHttpRequest(fn TOnServerHttpRequestEvent)               // property event
	SetOnServerWebSocketRequest(fn TOnServerWebSocketRequestEvent)     // property event
	SetOnServerWebSocketConnected(fn TOnServerWebSocketConnectedEvent) // property event
	SetOnServerWebSocketMessage(fn TOnServerWebSocketMessageEvent)     // property event
	AsIntfServerHandler() uintptr
}

type TEngServerHandler struct {
	TCEFServerHandlerOwn
}

func (m *TEngServerHandler) SetOnServerServerCreated(fn TOnServerServerCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerServerCreatedEvent(fn)
	base.SetEvent(m, 1, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerServerDestroyed(fn TOnServerServerDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerServerDestroyedEvent(fn)
	base.SetEvent(m, 2, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerClientConnected(fn TOnServerClientConnectedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerClientConnectedEvent(fn)
	base.SetEvent(m, 3, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerClientDisconnected(fn TOnServerClientDisconnectedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerClientDisconnectedEvent(fn)
	base.SetEvent(m, 4, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerHttpRequest(fn TOnServerHttpRequestEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerHttpRequestEvent(fn)
	base.SetEvent(m, 5, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerWebSocketRequest(fn TOnServerWebSocketRequestEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerWebSocketRequestEvent(fn)
	base.SetEvent(m, 6, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerWebSocketConnected(fn TOnServerWebSocketConnectedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerWebSocketConnectedEvent(fn)
	base.SetEvent(m, 7, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) SetOnServerWebSocketMessage(fn TOnServerWebSocketMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerWebSocketMessageEvent(fn)
	base.SetEvent(m, 8, engServerHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngServerHandler) AsIntfServerHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngServerHandler class constructor
func NewEngServerHandler() IEngServerHandler {
	var serverHandlerPtr uintptr // ICefServerHandler
	r := engServerHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&serverHandlerPtr)))
	ret := AsEngServerHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, serverHandlerPtr)
	}
	return ret
}

var (
	engServerHandlerOnce   base.Once
	engServerHandlerImport *imports.Imports = nil
)

func engServerHandlerAPI() *imports.Imports {
	engServerHandlerOnce.Do(func() {
		engServerHandlerImport = api.NewDefaultImports()
		engServerHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngServerHandler_Create", 0), // constructor NewEngServerHandler
			/* 1 */ imports.NewTable("TEngServerHandler_OnServerServerCreated", 0), // event OnServerServerCreated
			/* 2 */ imports.NewTable("TEngServerHandler_OnServerServerDestroyed", 0), // event OnServerServerDestroyed
			/* 3 */ imports.NewTable("TEngServerHandler_OnServerClientConnected", 0), // event OnServerClientConnected
			/* 4 */ imports.NewTable("TEngServerHandler_OnServerClientDisconnected", 0), // event OnServerClientDisconnected
			/* 5 */ imports.NewTable("TEngServerHandler_OnServerHttpRequest", 0), // event OnServerHttpRequest
			/* 6 */ imports.NewTable("TEngServerHandler_OnServerWebSocketRequest", 0), // event OnServerWebSocketRequest
			/* 7 */ imports.NewTable("TEngServerHandler_OnServerWebSocketConnected", 0), // event OnServerWebSocketConnected
			/* 8 */ imports.NewTable("TEngServerHandler_OnServerWebSocketMessage", 0), // event OnServerWebSocketMessage
		}
	})
	return engServerHandlerImport
}
