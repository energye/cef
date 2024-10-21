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

// ICefProcessMessage Parent: ICefBaseRefCounted
//
//	Interface representing a message. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_process_message_capi.h">CEF source file: /include/capi/cef_process_message_capi.h (cef_process_message_t))</a>
type ICefProcessMessage interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// Copy
	//  Returns a writable copy of this object. Returns nullptr when message contains a shared memory region.
	Copy() ICefProcessMessage // function
	// GetName
	//  Returns the message name.
	GetName() string // function
	// GetArgumentList
	//  Returns the list of arguments. Returns nullptr when message contains a shared memory region.
	GetArgumentList() ICefListValue // function
	// GetSharedMemoryRegion
	//  Returns the shared memory region. Returns nullptr when message contains an argument list.
	GetSharedMemoryRegion() ICefSharedMemoryRegion // function
}

// TCefProcessMessage Parent: TCefBaseRefCounted
//
//	Interface representing a message. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_process_message_capi.h">CEF source file: /include/capi/cef_process_message_capi.h (cef_process_message_t))</a>
type TCefProcessMessage struct {
	TCefBaseRefCounted
}

// ProcessMessageRef -> ICefProcessMessage
var ProcessMessageRef processMessage

// processMessage TCefProcessMessage Ref
type processMessage uintptr

func (m *processMessage) UnWrap(data uintptr) ICefProcessMessage {
	var resultCefProcessMessage uintptr
	processMessageImportAPI().SysCallN(7, uintptr(data), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *processMessage) New(name string) ICefProcessMessage {
	var resultCefProcessMessage uintptr
	processMessageImportAPI().SysCallN(6, PascalStr(name), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *TCefProcessMessage) IsValid() bool {
	r1 := processMessageImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCefProcessMessage) IsReadOnly() bool {
	r1 := processMessageImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefProcessMessage) Copy() ICefProcessMessage {
	var resultCefProcessMessage uintptr
	processMessageImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *TCefProcessMessage) GetName() string {
	r1 := processMessageImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefProcessMessage) GetArgumentList() ICefListValue {
	var resultCefListValue uintptr
	processMessageImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefProcessMessage) GetSharedMemoryRegion() ICefSharedMemoryRegion {
	var resultCefSharedMemoryRegion uintptr
	processMessageImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefSharedMemoryRegion)))
	return AsCefSharedMemoryRegion(resultCefSharedMemoryRegion)
}

var (
	processMessageImport       *imports.Imports = nil
	processMessageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefProcessMessage_Copy", 0),
		/*1*/ imports.NewTable("CefProcessMessage_GetArgumentList", 0),
		/*2*/ imports.NewTable("CefProcessMessage_GetName", 0),
		/*3*/ imports.NewTable("CefProcessMessage_GetSharedMemoryRegion", 0),
		/*4*/ imports.NewTable("CefProcessMessage_IsReadOnly", 0),
		/*5*/ imports.NewTable("CefProcessMessage_IsValid", 0),
		/*6*/ imports.NewTable("CefProcessMessage_New", 0),
		/*7*/ imports.NewTable("CefProcessMessage_UnWrap", 0),
	}
)

func processMessageImportAPI() *imports.Imports {
	if processMessageImport == nil {
		processMessageImport = NewDefaultImports()
		processMessageImport.SetImportTable(processMessageImportTables)
		processMessageImportTables = nil
	}
	return processMessageImport
}
