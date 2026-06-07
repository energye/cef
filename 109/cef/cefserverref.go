//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefServer Parent: ICefBaseRefCounted
type ICefServer interface {
	ICefBaseRefCounted
	GetTaskRunner() ICefTaskRunner                                                                                                     // function
	IsRunning() bool                                                                                                                   // function
	GetAddress() string                                                                                                                // function
	HasConnection() bool                                                                                                               // function
	IsValidConnection(connectionId int32) bool                                                                                         // function
	Shutdown()                                                                                                                         // procedure
	SendHttp200response(connectionId int32, contentType string, data uintptr, dataSize cefTypes.NativeUInt)                            // procedure
	SendHttp404response(connectionId int32)                                                                                            // procedure
	SendHttp500response(connectionId int32, errorMessage string)                                                                       // procedure
	SendHttpResponse(connectionId int32, responseCode int32, contentType string, contentLength int64, extraHeaders ICefStringMultimap) // procedure
	SendRawData(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt)                                                        // procedure
	CloseConnection(connectionId int32)                                                                                                // procedure
	SendWebSocketMessage(connectionId int32, data uintptr, dataSize cefTypes.NativeUInt)                                               // procedure
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
