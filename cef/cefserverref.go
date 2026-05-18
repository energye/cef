//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefServer Parent: ICefBaseRefCounted
type ICefServer interface {
	ICefBaseRefCounted
	// GetTaskRunner
	//  Returns the task runner for the dedicated server thread.
	GetTaskRunner() ICefTaskRunner // function
	// IsRunning
	//  Returns true (1) if the server is currently running and accepting incoming
	//  connections. See ICefServerHandler.OnServerCreated documentation for a
	//  description of server lifespan. This function must be called on the
	//  dedicated server thread.
	IsRunning() bool // function
	// GetAddress
	//  Returns the server address including the port number.
	GetAddress() string // function
	// HasConnection
	//  Returns true (1) if the server currently has a connection. This function
	//  must be called on the dedicated server thread.
	HasConnection() bool // function
	// IsValidConnection
	//  Returns true (1) if |connection_id| represents a valid connection. This
	//  function must be called on the dedicated server thread.
	IsValidConnection(connectionId int32) bool // function
	// Shutdown
	//  Stop the server and shut down the dedicated server thread. See
	//  ICefServerHandler.OnServerCreated documentation for a description of
	//  server lifespan.
	Shutdown() // procedure
	// SendHttp200response
	//  Send an HTTP 200 "OK" response to the connection identified by
	//  |connection_id|. |content_type| is the response content type (e.g.
	//  "text/html"), |data| is the response content, and |data_size| is the size
	//  of |data| in bytes. The contents of |data| will be copied. The connection
	//  will be closed automatically after the response is sent.
	SendHttp200response(connectionId int32, contentType string, data uintptr, dataSize cefTypes.NativeUInt) // procedure
	// SendHttp404response
	//  Send an HTTP 404 "Not Found" response to the connection identified by
	//  |connection_id|. The connection will be closed automatically after the
	//  response is sent.
	SendHttp404response(connectionId int32) // procedure
	// SendHttp500response
	//  Send an HTTP 500 "Internal Server Error" response to the connection
	//  identified by |connection_id|. |error_message| is the associated error
	//  message. The connection will be closed automatically after the response is
	//  sent.
	SendHttp500response(connectionId int32, errorMessage string) // procedure
	// SendHttpResponse
	//  Send a custom HTTP response to the connection identified by
	//  |connection_id|. |response_code| is the HTTP response code sent in the
	//  status line (e.g. 200), |content_type| is the response content type sent
	//  as the "Content-Type" header (e.g. "text/html"), |content_length| is the
	//  expected content length, and |extra_headers| is the map of extra response
	//  headers. If |content_length| is >= 0 then the "Content-Length" header will
	//  be sent. If |content_length| is 0 then no content is expected and the
	//  connection will be closed automatically after the response is sent. If
	//  |content_length| is < 0 then no "Content-Length" header will be sent and
	//  the client will continue reading until the connection is closed. Use the
	//  SendRawData function to send the content, if applicable, and call
	//  CloseConnection after all content has been sent.
	SendHttpResponse(connectionId int32, responseCode int32, contentType string, contentLength int64, extraHeaders ICefStringMultimap) // procedure
	// SendRawData
	//  Send raw data directly to the connection identified by |connection_id|.
	//  |data| is the raw data and |data_size| is the size of |data| in bytes. The
	//  contents of |data| will be copied. No validation of |data| is performed
	//  internally so the client should be careful to send the amount indicated by
	//  the "Content-Length" header, if specified. See SendHttpResponse
	//  documentation for intended usage.
	SendRawData(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) // procedure
	// CloseConnection
	//  Close the connection identified by |connection_id|. See SendHttpResponse
	//  documentation for intended usage.
	CloseConnection(connectionId int32) // procedure
	// SendWebSocketMessage
	//  Send a WebSocket message to the connection identified by |connection_id|.
	//  |data| is the response content and |data_size| is the size of |data| in
	//  bytes. The contents of |data| will be copied. See
	//  ICefServerHandler.OnWebSocketRequest documentation for intended usage.
	SendWebSocketMessage(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) // procedure
}

// ICEFServerRef Parent: ICefServer ICefBaseRefCountedRef
type ICEFServerRef interface {
	ICefServer
	ICefBaseRefCountedRef
	AsIntfServer() uintptr
}

type TCEFServerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCEFServerRef) GetTaskRunner() (result ICefTaskRunner) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFServerRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTaskRunnerRef(resultPtr)
	return
}

func (m *TCEFServerRef) IsRunning() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFServerRef) GetAddress() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFServerRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFServerRef) HasConnection() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFServerRef) IsValidConnection(connectionId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFServerRefAPI().SysCallN(5, m.Instance(), uintptr(connectionId))
	return api.GoBool(r)
}

func (m *TCEFServerRef) Shutdown() {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(7, m.Instance())
}

func (m *TCEFServerRef) SendHttp200response(connectionId int32, contentType string, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(8, m.Instance(), uintptr(connectionId), api.PasStr(contentType), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerRef) SendHttp404response(connectionId int32) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(9, m.Instance(), uintptr(connectionId))
}

func (m *TCEFServerRef) SendHttp500response(connectionId int32, errorMessage string) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(10, m.Instance(), uintptr(connectionId), api.PasStr(errorMessage))
}

func (m *TCEFServerRef) SendHttpResponse(connectionId int32, responseCode int32, contentType string, contentLength int64, extraHeaders ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(11, m.Instance(), uintptr(connectionId), uintptr(responseCode), api.PasStr(contentType), uintptr(base.UnsafePointer(&contentLength)), base.GetObjectUintptr(extraHeaders))
}

func (m *TCEFServerRef) SendRawData(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(12, m.Instance(), uintptr(connectionId), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerRef) CloseConnection(connectionId int32) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(13, m.Instance(), uintptr(connectionId))
}

func (m *TCEFServerRef) SendWebSocketMessage(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt) {
	if !m.IsValid() {
		return
	}
	cEFServerRefAPI().SysCallN(14, m.Instance(), uintptr(connectionId), uintptr(data), uintptr(dataSize))
}

func (m *TCEFServerRef) AsIntfServer() uintptr {
	return m.GetIntfPointer(0)
}

// ServerRef  is static instance
var ServerRef _ServerRefClass

// _ServerRefClass is class type defined by TCEFServerRef
type _ServerRefClass uintptr

func (_ServerRefClass) UnWrap(data uintptr) (result ICefServer) {
	var resultPtr uintptr
	cEFServerRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCEFServerRef(resultPtr)
	return
}

// NewServerRef class constructor
func NewServerRef(data uintptr) ICEFServerRef {
	var serverPtr uintptr // ICefServer
	r := cEFServerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&serverPtr)))
	ret := AsCEFServerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, serverPtr)
	}
	return ret
}

var (
	cEFServerRefOnce   base.Once
	cEFServerRefImport *imports.Imports = nil
)

func cEFServerRefAPI() *imports.Imports {
	cEFServerRefOnce.Do(func() {
		cEFServerRefImport = api.NewDefaultImports()
		cEFServerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFServerRef_Create", 0), // constructor NewServerRef
			/* 1 */ imports.NewTable("TCEFServerRef_GetTaskRunner", 0), // function GetTaskRunner
			/* 2 */ imports.NewTable("TCEFServerRef_IsRunning", 0), // function IsRunning
			/* 3 */ imports.NewTable("TCEFServerRef_GetAddress", 0), // function GetAddress
			/* 4 */ imports.NewTable("TCEFServerRef_HasConnection", 0), // function HasConnection
			/* 5 */ imports.NewTable("TCEFServerRef_IsValidConnection", 0), // function IsValidConnection
			/* 6 */ imports.NewTable("TCEFServerRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCEFServerRef_Shutdown", 0), // procedure Shutdown
			/* 8 */ imports.NewTable("TCEFServerRef_SendHttp200response", 0), // procedure SendHttp200response
			/* 9 */ imports.NewTable("TCEFServerRef_SendHttp404response", 0), // procedure SendHttp404response
			/* 10 */ imports.NewTable("TCEFServerRef_SendHttp500response", 0), // procedure SendHttp500response
			/* 11 */ imports.NewTable("TCEFServerRef_SendHttpResponse", 0), // procedure SendHttpResponse
			/* 12 */ imports.NewTable("TCEFServerRef_SendRawData", 0), // procedure SendRawData
			/* 13 */ imports.NewTable("TCEFServerRef_CloseConnection", 0), // procedure CloseConnection
			/* 14 */ imports.NewTable("TCEFServerRef_SendWebSocketMessage", 0), // procedure SendWebSocketMessage
		}
	})
	return cEFServerRefImport
}
