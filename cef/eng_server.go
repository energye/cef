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

// ICEFServer Parent: ICefBaseRefCounted
type ICEFServer interface {
	ICefBaseRefCounted
	GetTaskRunner() ICefTaskRunner                                                                                               // function
	IsRunning() bool                                                                                                             // function
	GetAddress() string                                                                                                          // function
	HasConnection() bool                                                                                                         // function
	IsValidConnection(connectionid int32) bool                                                                                   // function
	Shutdown()                                                                                                                   // procedure
	SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt)                               // procedure
	SendHttp404response(connectionid int32)                                                                                      // procedure
	SendHttp500response(connectionid int32, errormessage string)                                                                 // procedure
	SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) // procedure
	SendRawData(connectionid int32, data uintptr, datasize NativeUInt)                                                           // procedure
	CloseConnection(connectionid int32)                                                                                          // procedure
	SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt)                                                  // procedure
}

// TCEFServer Parent: TCefBaseRefCounted
type TCEFServer struct {
	TCefBaseRefCounted
}

// ServerRef -> ICEFServer
var ServerRef server

// server TCEFServer Ref
type server uintptr

func (m *server) UnWrap(data uintptr) ICEFServer {
	var resultCEFServer uintptr
	serverImportAPI().SysCallN(13, uintptr(data), uintptr(unsafePointer(&resultCEFServer)))
	return AsCEFServer(resultCEFServer)
}

func (m *TCEFServer) GetTaskRunner() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	serverImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCEFServer) IsRunning() bool {
	r1 := serverImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServer) GetAddress() string {
	r1 := serverImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCEFServer) HasConnection() bool {
	r1 := serverImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServer) IsValidConnection(connectionid int32) bool {
	r1 := serverImportAPI().SysCallN(5, m.Instance(), uintptr(connectionid))
	return GoBool(r1)
}

func (m *TCEFServer) Shutdown() {
	serverImportAPI().SysCallN(12, m.Instance())
}

func (m *TCEFServer) SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) {
	serverImportAPI().SysCallN(6, m.Instance(), uintptr(connectionid), PascalStr(contenttype), uintptr(data), uintptr(datasize))
}

func (m *TCEFServer) SendHttp404response(connectionid int32) {
	serverImportAPI().SysCallN(7, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServer) SendHttp500response(connectionid int32, errormessage string) {
	serverImportAPI().SysCallN(8, m.Instance(), uintptr(connectionid), PascalStr(errormessage))
}

func (m *TCEFServer) SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) {
	serverImportAPI().SysCallN(9, m.Instance(), uintptr(connectionid), uintptr(responsecode), PascalStr(contenttype), uintptr(unsafePointer(&contentlength)), GetObjectUintptr(extraheaders))
}

func (m *TCEFServer) SendRawData(connectionid int32, data uintptr, datasize NativeUInt) {
	serverImportAPI().SysCallN(10, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServer) CloseConnection(connectionid int32) {
	serverImportAPI().SysCallN(0, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServer) SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) {
	serverImportAPI().SysCallN(11, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

var (
	serverImport       *imports.Imports = nil
	serverImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFServer_CloseConnection", 0),
		/*1*/ imports.NewTable("CEFServer_GetAddress", 0),
		/*2*/ imports.NewTable("CEFServer_GetTaskRunner", 0),
		/*3*/ imports.NewTable("CEFServer_HasConnection", 0),
		/*4*/ imports.NewTable("CEFServer_IsRunning", 0),
		/*5*/ imports.NewTable("CEFServer_IsValidConnection", 0),
		/*6*/ imports.NewTable("CEFServer_SendHttp200response", 0),
		/*7*/ imports.NewTable("CEFServer_SendHttp404response", 0),
		/*8*/ imports.NewTable("CEFServer_SendHttp500response", 0),
		/*9*/ imports.NewTable("CEFServer_SendHttpResponse", 0),
		/*10*/ imports.NewTable("CEFServer_SendRawData", 0),
		/*11*/ imports.NewTable("CEFServer_SendWebSocketMessage", 0),
		/*12*/ imports.NewTable("CEFServer_Shutdown", 0),
		/*13*/ imports.NewTable("CEFServer_UnWrap", 0),
	}
)

func serverImportAPI() *imports.Imports {
	if serverImport == nil {
		serverImport = NewDefaultImports()
		serverImport.SetImportTable(serverImportTables)
		serverImportTables = nil
	}
	return serverImport
}
