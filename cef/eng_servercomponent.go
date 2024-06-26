//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICEFServerComponent Parent: IComponent
//
//	The TCEFServerComponent class puts together all CEF server procedures, functions, properties and events in one place.
type ICEFServerComponent interface {
	IComponent
	// AsInterface
	//  Class instance to interface instance
	AsInterface() IServerEvents // procedure
	// Initialized
	//  Returns true when the server and the handler are initialized.
	Initialized() bool // property
	// IsRunning
	//  Returns true(1) if the server is currently running and accepting incoming
	//  connections. See ICefServerHandler.OnServerCreated documentation for a
	//  description of server lifespan. This function must be called on the
	//  dedicated server thread.
	IsRunning() bool // property
	// Address
	//  Returns the server address including the port number.
	Address() string // property
	// HasConnection
	//  Returns true(1) if the server currently has a connection. This function
	//  must be called on the dedicated server thread.
	HasConnection() bool // property
	// IsValidConnection
	//  Returns true(1) if |connection_id| represents a valid connection. This
	//  function must be called on the dedicated server thread.
	IsValidConnection(connectionid int32) bool // function
	// CreateServer
	//  Create a new server that binds to |address| and |port|. |address| must be a
	//  valid IPv4 or IPv6 address(e.g. 127.0.0.1 or::1) and |port| must be a port
	//  number outside of the reserved range(e.g. between 1025 and 65535 on most
	//  platforms). |backlog| is the maximum number of pending connections. A new
	//  thread will be created for each CreateServer call(the "dedicated server
	//  thread"). It is therefore recommended to use a different
	//  ICefServerHandler instance for each CreateServer call to avoid thread
	//  safety issues in the ICefServerHandler implementation. The
	//  ICefServerHandler.OnServerCreated function will be called on the
	//  dedicated server thread to report success or failure. See
	//  ICefServerHandler.OnServerCreated documentation for a description of
	//  server lifespan.
	CreateServer(address string, port uint16, backlog int32) // procedure
	// Shutdown
	//  Stop the server and shut down the dedicated server thread. See
	//  ICefServerHandler.OnServerCreated documentation for a description of
	//  server lifespan.
	Shutdown() // procedure
	// SendHttp200response
	//  Send an HTTP 200 "OK" response to the connection identified by
	//  |connection_id|. |content_type| is the response content type(e.g.
	//  "text/html"), |data| is the response content, and |data_size| is the size
	//  of |data| in bytes. The contents of |data| will be copied. The connection
	//  will be closed automatically after the response is sent.
	SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) // procedure
	// SendHttp404response
	//  Send an HTTP 404 "Not Found" response to the connection identified by
	//  |connection_id|. The connection will be closed automatically after the
	//  response is sent.
	SendHttp404response(connectionid int32) // procedure
	// SendHttp500response
	//  Send an HTTP 500 "Internal Server Error" response to the connection
	//  identified by |connection_id|. |error_message| is the associated error
	//  message. The connection will be closed automatically after the response is
	//  sent.
	SendHttp500response(connectionid int32, errormessage string) // procedure
	// SendHttpResponse
	//  Send a custom HTTP response to the connection identified by
	//  |connection_id|. |response_code| is the HTTP response code sent in the
	//  status line(e.g. 200), |content_type| is the response content type sent
	//  as the "Content-Type" header(e.g. "text/html"), |content_length| is the
	//  expected content length, and |extra_headers| is the map of extra response
	//  headers. If |content_length| is >= 0 then the "Content-Length" header will
	//  be sent. If |content_length| is 0 then no content is expected and the
	//  connection will be closed automatically after the response is sent. If
	//  |content_length| is < 0 then no "Content-Length" header will be sent and
	//  the client will continue reading until the connection is closed. Use the
	//  SendRawData function to send the content, if applicable, and call
	//  CloseConnection after all content has been sent.
	SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) // procedure
	// SendRawData
	//  Send raw data directly to the connection identified by |connection_id|.
	//  |data| is the raw data and |data_size| is the size of |data| in bytes. The
	//  contents of |data| will be copied. No validation of |data| is performed
	//  internally so the client should be careful to send the amount indicated by
	//  the "Content-Length" header, if specified. See SendHttpResponse
	//  documentation for intended usage.
	SendRawData(connectionid int32, data uintptr, datasize NativeUInt) // procedure
	// CloseConnection
	//  Close the connection identified by |connection_id|. See SendHttpResponse
	//  documentation for intended usage.
	CloseConnection(connectionid int32) // procedure
	// SendWebSocketMessage
	//  Send a WebSocket message to the connection identified by |connection_id|.
	//  |data| is the response content and |data_size| is the size of |data| in
	//  bytes. The contents of |data| will be copied. See
	//  ICefServerHandler.OnWebSocketRequest documentation for intended usage.
	SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) // procedure
	// SetOnServerCreated
	//  Called when |server| is created. If the server was started successfully
	//  then ICefServer.IsRunning will return true(1). The server will
	//  continue running until ICefServerShutdown is called, after which time
	//  OnServerDestroyed will be called. If the server failed to start then
	//  OnServerDestroyed will be called immediately after this function returns.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnServerCreated(fn TOnServerCreated) // property event
	// SetOnServerDestroyed
	//  Called when |server| is destroyed. The server thread will be stopped after
	//  this function returns. The client should release any references to
	//  |server| when this function is called. See OnServerCreated documentation
	//  for a description of server lifespan.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnServerDestroyed(fn TOnServerDestroyed) // property event
	// SetOnClientConnected
	//  Called when a client connects to |server|. |connection_id| uniquely
	//  identifies the connection. Each call to this function will have a matching
	//  call to OnClientDisconnected.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnClientConnected(fn TOnClientConnected) // property event
	// SetOnClientDisconnected
	//  Called when a client disconnects from |server|. |connection_id| uniquely
	//  identifies the connection. The client should release any data associated
	//  with |connection_id| when this function is called and |connection_id|
	//  should no longer be passed to ICefServer functions. Disconnects can
	//  originate from either the client or the server. For example, the server
	//  will disconnect automatically after a ICefServer.SendHttpXXXResponse
	//  function is called.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnClientDisconnected(fn TOnClientDisconnected) // property event
	// SetOnHttpRequest
	//  Called when |server| receives an HTTP request. |connection_id| uniquely
	//  identifies the connection, |client_address| is the requesting IPv4 or IPv6
	//  client address including port number, and |request| contains the request
	//  contents(URL, function, headers and optional POST data). Call
	//  ICefServer functions either synchronously or asynchronusly to send a
	//  response.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnHttpRequest(fn TOnHttpRequest) // property event
	// SetOnWebSocketRequest
	//  Called when |server| receives a WebSocket request. |connection_id|
	//  uniquely identifies the connection, |client_address| is the requesting
	//  IPv4 or IPv6 client address including port number, and |request| contains
	//  the request contents(URL, function, headers and optional POST data).
	//  Execute |callback| either synchronously or asynchronously to accept or
	//  decline the WebSocket connection. If the request is accepted then
	//  OnWebSocketConnected will be called after the WebSocket has connected and
	//  incoming messages will be delivered to the OnWebSocketMessage callback. If
	//  the request is declined then the client will be disconnected and
	//  OnClientDisconnected will be called. Call the
	//  ICefServer.SendWebSocketMessage function after receiving the
	//  OnWebSocketConnected callback to respond with WebSocket messages.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketRequest(fn TOnWebSocketRequest) // property event
	// SetOnWebSocketConnected
	//  Called after the client has accepted the WebSocket connection for |server|
	//  and |connection_id| via the OnWebSocketRequest callback. See
	//  OnWebSocketRequest documentation for intended usage.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketConnected(fn TOnWebSocketConnected) // property event
	// SetOnWebSocketMessage
	//  Called when |server| receives an WebSocket message. |connection_id|
	//  uniquely identifies the connection, |data| is the message content and
	//  |data_size| is the size of |data| in bytes. Do not keep a reference to
	//  |data| outside of this function. See OnWebSocketRequest documentation for
	//  intended usage.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketMessage(fn TOnWebSocketMessage) // property event
}

// TCEFServerComponent Parent: TComponent
//
//	The TCEFServerComponent class puts together all CEF server procedures, functions, properties and events in one place.
type TCEFServerComponent struct {
	TComponent
	serverCreatedPtr      uintptr
	serverDestroyedPtr    uintptr
	clientConnectedPtr    uintptr
	clientDisconnectedPtr uintptr
	httpRequestPtr        uintptr
	webSocketRequestPtr   uintptr
	webSocketConnectedPtr uintptr
	webSocketMessagePtr   uintptr
}

func NewCEFServerComponent(aOwner IComponent) ICEFServerComponent {
	r1 := serverComponentImportAPI().SysCallN(3, GetObjectUintptr(aOwner))
	return AsCEFServerComponent(r1)
}

func (m *TCEFServerComponent) AsInterface() IServerEvents {
	var resultServerEvents uintptr
	serverComponentImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultServerEvents)))
	return AsServerEvents(resultServerEvents)
}

func (m *TCEFServerComponent) Initialized() bool {
	r1 := serverComponentImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) IsRunning() bool {
	r1 := serverComponentImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) Address() string {
	r1 := serverComponentImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCEFServerComponent) HasConnection() bool {
	r1 := serverComponentImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) IsValidConnection(connectionid int32) bool {
	r1 := serverComponentImportAPI().SysCallN(8, m.Instance(), uintptr(connectionid))
	return GoBool(r1)
}

func (m *TCEFServerComponent) CreateServer(address string, port uint16, backlog int32) {
	serverComponentImportAPI().SysCallN(4, m.Instance(), PascalStr(address), uintptr(port), uintptr(backlog))
}

func (m *TCEFServerComponent) Shutdown() {
	serverComponentImportAPI().SysCallN(23, m.Instance())
}

func (m *TCEFServerComponent) SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) {
	serverComponentImportAPI().SysCallN(9, m.Instance(), uintptr(connectionid), PascalStr(contenttype), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) SendHttp404response(connectionid int32) {
	serverComponentImportAPI().SysCallN(10, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServerComponent) SendHttp500response(connectionid int32, errormessage string) {
	serverComponentImportAPI().SysCallN(11, m.Instance(), uintptr(connectionid), PascalStr(errormessage))
}

func (m *TCEFServerComponent) SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) {
	serverComponentImportAPI().SysCallN(12, m.Instance(), uintptr(connectionid), uintptr(responsecode), PascalStr(contenttype), uintptr(unsafePointer(&contentlength)), GetObjectUintptr(extraheaders))
}

func (m *TCEFServerComponent) SendRawData(connectionid int32, data uintptr, datasize NativeUInt) {
	serverComponentImportAPI().SysCallN(13, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) CloseConnection(connectionid int32) {
	serverComponentImportAPI().SysCallN(2, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServerComponent) SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) {
	serverComponentImportAPI().SysCallN(14, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) SetOnServerCreated(fn TOnServerCreated) {
	if m.serverCreatedPtr != 0 {
		RemoveEventElement(m.serverCreatedPtr)
	}
	m.serverCreatedPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(18, m.Instance(), m.serverCreatedPtr)
}

func (m *TCEFServerComponent) SetOnServerDestroyed(fn TOnServerDestroyed) {
	if m.serverDestroyedPtr != 0 {
		RemoveEventElement(m.serverDestroyedPtr)
	}
	m.serverDestroyedPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(19, m.Instance(), m.serverDestroyedPtr)
}

func (m *TCEFServerComponent) SetOnClientConnected(fn TOnClientConnected) {
	if m.clientConnectedPtr != 0 {
		RemoveEventElement(m.clientConnectedPtr)
	}
	m.clientConnectedPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(15, m.Instance(), m.clientConnectedPtr)
}

func (m *TCEFServerComponent) SetOnClientDisconnected(fn TOnClientDisconnected) {
	if m.clientDisconnectedPtr != 0 {
		RemoveEventElement(m.clientDisconnectedPtr)
	}
	m.clientDisconnectedPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(16, m.Instance(), m.clientDisconnectedPtr)
}

func (m *TCEFServerComponent) SetOnHttpRequest(fn TOnHttpRequest) {
	if m.httpRequestPtr != 0 {
		RemoveEventElement(m.httpRequestPtr)
	}
	m.httpRequestPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(17, m.Instance(), m.httpRequestPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketRequest(fn TOnWebSocketRequest) {
	if m.webSocketRequestPtr != 0 {
		RemoveEventElement(m.webSocketRequestPtr)
	}
	m.webSocketRequestPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(22, m.Instance(), m.webSocketRequestPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketConnected(fn TOnWebSocketConnected) {
	if m.webSocketConnectedPtr != 0 {
		RemoveEventElement(m.webSocketConnectedPtr)
	}
	m.webSocketConnectedPtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(20, m.Instance(), m.webSocketConnectedPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketMessage(fn TOnWebSocketMessage) {
	if m.webSocketMessagePtr != 0 {
		RemoveEventElement(m.webSocketMessagePtr)
	}
	m.webSocketMessagePtr = MakeEventDataPtr(fn)
	serverComponentImportAPI().SysCallN(21, m.Instance(), m.webSocketMessagePtr)
}

var (
	serverComponentImport       *imports.Imports = nil
	serverComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFServerComponent_Address", 0),
		/*1*/ imports.NewTable("CEFServerComponent_AsInterface", 0),
		/*2*/ imports.NewTable("CEFServerComponent_CloseConnection", 0),
		/*3*/ imports.NewTable("CEFServerComponent_Create", 0),
		/*4*/ imports.NewTable("CEFServerComponent_CreateServer", 0),
		/*5*/ imports.NewTable("CEFServerComponent_HasConnection", 0),
		/*6*/ imports.NewTable("CEFServerComponent_Initialized", 0),
		/*7*/ imports.NewTable("CEFServerComponent_IsRunning", 0),
		/*8*/ imports.NewTable("CEFServerComponent_IsValidConnection", 0),
		/*9*/ imports.NewTable("CEFServerComponent_SendHttp200response", 0),
		/*10*/ imports.NewTable("CEFServerComponent_SendHttp404response", 0),
		/*11*/ imports.NewTable("CEFServerComponent_SendHttp500response", 0),
		/*12*/ imports.NewTable("CEFServerComponent_SendHttpResponse", 0),
		/*13*/ imports.NewTable("CEFServerComponent_SendRawData", 0),
		/*14*/ imports.NewTable("CEFServerComponent_SendWebSocketMessage", 0),
		/*15*/ imports.NewTable("CEFServerComponent_SetOnClientConnected", 0),
		/*16*/ imports.NewTable("CEFServerComponent_SetOnClientDisconnected", 0),
		/*17*/ imports.NewTable("CEFServerComponent_SetOnHttpRequest", 0),
		/*18*/ imports.NewTable("CEFServerComponent_SetOnServerCreated", 0),
		/*19*/ imports.NewTable("CEFServerComponent_SetOnServerDestroyed", 0),
		/*20*/ imports.NewTable("CEFServerComponent_SetOnWebSocketConnected", 0),
		/*21*/ imports.NewTable("CEFServerComponent_SetOnWebSocketMessage", 0),
		/*22*/ imports.NewTable("CEFServerComponent_SetOnWebSocketRequest", 0),
		/*23*/ imports.NewTable("CEFServerComponent_Shutdown", 0),
	}
)

func serverComponentImportAPI() *imports.Imports {
	if serverComponentImport == nil {
		serverComponentImport = NewDefaultImports()
		serverComponentImport.SetImportTable(serverComponentImportTables)
		serverComponentImportTables = nil
	}
	return serverComponentImport
}
