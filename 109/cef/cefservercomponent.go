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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/109/types"
)

// IServerEvents Parent: IObject
type IServerEvents interface {
	IObject
}

// ICEFServerComponent Parent: IServerEvents IComponent
type ICEFServerComponent interface {
	IServerEvents
	IComponent
	IsValidConnection(connectionId int32) bool                                                                                         // function
	CreateServer(address string, port uint16, backlog int32)                                                                           // procedure
	Shutdown()                                                                                                                         // procedure
	SendHttp200response(connectionId int32, contentType string, data uintptr, dataSize cefTypes.NativeUInt)                            // procedure
	SendHttp404response(connectionId int32)                                                                                            // procedure
	SendHttp500response(connectionId int32, errorMessage string)                                                                       // procedure
	SendHttpResponse(connectionId int32, responseCode int32, contentType string, contentLength int64, extraHeaders ICefStringMultimap) // procedure
	SendRawData(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt)                                                        // procedure
	CloseConnection(connectionId int32)                                                                                                // procedure
	SendWebSocketMessage(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt)                                               // procedure
	Initialized() bool                                                                                                                 // property Initialized Getter
	IsRunning() bool                                                                                                                   // property IsRunning Getter
	Address() string                                                                                                                   // property Address Getter
	HasConnection() bool                                                                                                               // property HasConnection Getter
	SetOnServerCreated(fn TOnServerCreated)                                                                                            // property event
	SetOnServerDestroyed(fn TOnServerDestroyed)                                                                                        // property event
	SetOnClientConnected(fn TOnClientConnected)                                                                                        // property event
	SetOnClientDisconnected(fn TOnClientDisconnected)                                                                                  // property event
	SetOnHttpRequest(fn TOnHttpRequest)                                                                                                // property event
	SetOnWebSocketRequest(fn TOnWebSocketRequest)                                                                                      // property event
	SetOnWebSocketConnected(fn TOnWebSocketConnected)                                                                                  // property event
	SetOnWebSocketMessage(fn TOnWebSocketMessage)                                                                                      // property event
	AsIntfServerEvents() uintptr
}

type TCEFServerComponent struct {
	TComponent
}

func (m *TCEFServerComponent) IsValidConnection(connectionId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerComponentAPI().SysCallN(1, m.Instance(), uintptr(connectionId))
	return api.GoBool(r)
}

func (m *TCEFServerComponent) CreateServer(address string, port uint16, backlog int32) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(2, m.Instance(), api.PasStr(address), uintptr(port), uintptr(backlog))
}

func (m *TCEFServerComponent) Shutdown() {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(3, m.Instance())
}

func (m *TCEFServerComponent) SendHttp200response(connectionId int32, contentType string, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(4, m.Instance(), uintptr(connectionId), api.PasStr(contentType), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerComponent) SendHttp404response(connectionId int32) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(5, m.Instance(), uintptr(connectionId))
}

func (m *TCEFServerComponent) SendHttp500response(connectionId int32, errorMessage string) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(6, m.Instance(), uintptr(connectionId), api.PasStr(errorMessage))
}

func (m *TCEFServerComponent) SendHttpResponse(connectionId int32, responseCode int32, contentType string, contentLength int64, extraHeaders ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(7, m.Instance(), uintptr(connectionId), uintptr(responseCode), api.PasStr(contentType), uintptr(base.UnsafePointer(&contentLength)), base.GetObjectUintptr(extraHeaders))
}

func (m *TCEFServerComponent) SendRawData(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(8, m.Instance(), uintptr(connectionId), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerComponent) CloseConnection(connectionId int32) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(9, m.Instance(), uintptr(connectionId))
}

func (m *TCEFServerComponent) SendWebSocketMessage(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerComponentAPI().SysCallN(10, m.Instance(), uintptr(connectionId), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerComponent) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerComponentAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFServerComponent) IsRunning() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerComponentAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFServerComponent) Address() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFServerComponentAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFServerComponent) HasConnection() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerComponentAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFServerComponent) SetOnServerCreated(fn TOnServerCreated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerCreated(fn)
	base.SetEvent(m, 15, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnServerDestroyed(fn TOnServerDestroyed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerDestroyed(fn)
	base.SetEvent(m, 16, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnClientConnected(fn TOnClientConnected) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientConnected(fn)
	base.SetEvent(m, 17, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnClientDisconnected(fn TOnClientDisconnected) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientDisconnected(fn)
	base.SetEvent(m, 18, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnHttpRequest(fn TOnHttpRequest) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnHttpRequest(fn)
	base.SetEvent(m, 19, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnWebSocketRequest(fn TOnWebSocketRequest) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebSocketRequest(fn)
	base.SetEvent(m, 20, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnWebSocketConnected(fn TOnWebSocketConnected) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebSocketConnected(fn)
	base.SetEvent(m, 21, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) SetOnWebSocketMessage(fn TOnWebSocketMessage) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebSocketMessage(fn)
	base.SetEvent(m, 22, cEFServerComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFServerComponent) AsIntfServerEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewServerComponent class constructor
func NewServerComponent(owner lcl.IComponent) ICEFServerComponent {
	var serverEventsPtr uintptr // IServerEvents
	r := cEFServerComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&serverEventsPtr)))
	ret := AsCEFServerComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, serverEventsPtr)
	}
	return ret
}

var (
	cEFServerComponentOnce   base.Once
	cEFServerComponentImport *imports.Imports = nil
)

func cEFServerComponentAPI() *imports.Imports {
	cEFServerComponentOnce.Do(func() {
		cEFServerComponentImport = api.NewDefaultImports()
		cEFServerComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFServerComponent_Create", 0), // constructor NewServerComponent
			/* 1 */ imports.NewTable("TCEFServerComponent_IsValidConnection", 0), // function IsValidConnection
			/* 2 */ imports.NewTable("TCEFServerComponent_CreateServer", 0), // procedure CreateServer
			/* 3 */ imports.NewTable("TCEFServerComponent_Shutdown", 0), // procedure Shutdown
			/* 4 */ imports.NewTable("TCEFServerComponent_SendHttp200response", 0), // procedure SendHttp200response
			/* 5 */ imports.NewTable("TCEFServerComponent_SendHttp404response", 0), // procedure SendHttp404response
			/* 6 */ imports.NewTable("TCEFServerComponent_SendHttp500response", 0), // procedure SendHttp500response
			/* 7 */ imports.NewTable("TCEFServerComponent_SendHttpResponse", 0), // procedure SendHttpResponse
			/* 8 */ imports.NewTable("TCEFServerComponent_SendRawData", 0), // procedure SendRawData
			/* 9 */ imports.NewTable("TCEFServerComponent_CloseConnection", 0), // procedure CloseConnection
			/* 10 */ imports.NewTable("TCEFServerComponent_SendWebSocketMessage", 0), // procedure SendWebSocketMessage
			/* 11 */ imports.NewTable("TCEFServerComponent_Initialized", 0), // property Initialized
			/* 12 */ imports.NewTable("TCEFServerComponent_IsRunning", 0), // property IsRunning
			/* 13 */ imports.NewTable("TCEFServerComponent_Address", 0), // property Address
			/* 14 */ imports.NewTable("TCEFServerComponent_HasConnection", 0), // property HasConnection
			/* 15 */ imports.NewTable("TCEFServerComponent_OnServerCreated", 0), // event OnServerCreated
			/* 16 */ imports.NewTable("TCEFServerComponent_OnServerDestroyed", 0), // event OnServerDestroyed
			/* 17 */ imports.NewTable("TCEFServerComponent_OnClientConnected", 0), // event OnClientConnected
			/* 18 */ imports.NewTable("TCEFServerComponent_OnClientDisconnected", 0), // event OnClientDisconnected
			/* 19 */ imports.NewTable("TCEFServerComponent_OnHttpRequest", 0), // event OnHttpRequest
			/* 20 */ imports.NewTable("TCEFServerComponent_OnWebSocketRequest", 0), // event OnWebSocketRequest
			/* 21 */ imports.NewTable("TCEFServerComponent_OnWebSocketConnected", 0), // event OnWebSocketConnected
			/* 22 */ imports.NewTable("TCEFServerComponent_OnWebSocketMessage", 0), // event OnWebSocketMessage
		}
	})
	return cEFServerComponentImport
}
