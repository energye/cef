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

// ICefBytesWriteHandler Parent: ICefWriteHandler
type ICefBytesWriteHandler interface {
	ICefWriteHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefWriteHandler                    // procedure
	Write(ptr uintptr, size, n NativeUInt) NativeUInt // function
	Seek(offset int64, whence int32) int32            // function
	Tell() (resultInt64 int64)                        // function
	Flush() int32                                     // function
	MayBlock() bool                                   // function
	GetData() uintptr                                 // function
	GetDataSize() (resultInt64 int64)                 // function
}

// TCefBytesWriteHandler Parent: TCefWriteHandler
type TCefBytesWriteHandler struct {
	TCefWriteHandler
}

func NewCefBytesWriteHandler(aGrow NativeUInt) ICefBytesWriteHandler {
	r1 := bytesWriteHandlerImportAPI().SysCallN(1, uintptr(aGrow))
	return AsCefBytesWriteHandler(r1)
}

func (m *TCefBytesWriteHandler) AsInterface() ICefWriteHandler {
	var resultCefWriteHandler uintptr
	bytesWriteHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefWriteHandler)))
	return AsCefWriteHandler(resultCefWriteHandler)
}

func (m *TCefBytesWriteHandler) Write(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := bytesWriteHandlerImportAPI().SysCallN(8, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefBytesWriteHandler) Seek(offset int64, whence int32) int32 {
	r1 := bytesWriteHandlerImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefBytesWriteHandler) Tell() (resultInt64 int64) {
	bytesWriteHandlerImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefBytesWriteHandler) Flush() int32 {
	r1 := bytesWriteHandlerImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCefBytesWriteHandler) MayBlock() bool {
	r1 := bytesWriteHandlerImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCefBytesWriteHandler) GetData() uintptr {
	r1 := bytesWriteHandlerImportAPI().SysCallN(3, m.Instance())
	return uintptr(r1)
}

func (m *TCefBytesWriteHandler) GetDataSize() (resultInt64 int64) {
	bytesWriteHandlerImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

var (
	bytesWriteHandlerImport       *imports.Imports = nil
	bytesWriteHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBytesWriteHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CefBytesWriteHandler_Create", 0),
		/*2*/ imports.NewTable("CefBytesWriteHandler_Flush", 0),
		/*3*/ imports.NewTable("CefBytesWriteHandler_GetData", 0),
		/*4*/ imports.NewTable("CefBytesWriteHandler_GetDataSize", 0),
		/*5*/ imports.NewTable("CefBytesWriteHandler_MayBlock", 0),
		/*6*/ imports.NewTable("CefBytesWriteHandler_Seek", 0),
		/*7*/ imports.NewTable("CefBytesWriteHandler_Tell", 0),
		/*8*/ imports.NewTable("CefBytesWriteHandler_Write", 0),
	}
)

func bytesWriteHandlerImportAPI() *imports.Imports {
	if bytesWriteHandlerImport == nil {
		bytesWriteHandlerImport = NewDefaultImports()
		bytesWriteHandlerImport.SetImportTable(bytesWriteHandlerImportTables)
		bytesWriteHandlerImportTables = nil
	}
	return bytesWriteHandlerImport
}
