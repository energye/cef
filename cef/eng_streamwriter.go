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

// ICefStreamWriter Parent: ICefBaseRefCounted
type ICefStreamWriter interface {
	ICefBaseRefCounted
	Write(ptr uintptr, size, n NativeUInt) NativeUInt // function
	Seek(offset int64, whence int32) int32            // function
	Tell() (resultInt64 int64)                        // function
	Flush() int32                                     // function
	MayBlock() bool                                   // function
}

// TCefStreamWriter Parent: TCefBaseRefCounted
type TCefStreamWriter struct {
	TCefBaseRefCounted
}

// StreamWriterRef -> ICefStreamWriter
var StreamWriterRef streamWriter

// streamWriter TCefStreamWriter Ref
type streamWriter uintptr

func (m *streamWriter) UnWrap(data uintptr) ICefStreamWriter {
	var resultCefStreamWriter uintptr
	streamWriterImportAPI().SysCallN(5, uintptr(data), uintptr(unsafePointer(&resultCefStreamWriter)))
	return AsCefStreamWriter(resultCefStreamWriter)
}

func (m *streamWriter) CreateForFile(fileName string) ICefStreamWriter {
	var resultCefStreamWriter uintptr
	streamWriterImportAPI().SysCallN(0, PascalStr(fileName), uintptr(unsafePointer(&resultCefStreamWriter)))
	return AsCefStreamWriter(resultCefStreamWriter)
}

func (m *TCefStreamWriter) Write(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := streamWriterImportAPI().SysCallN(6, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefStreamWriter) Seek(offset int64, whence int32) int32 {
	r1 := streamWriterImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefStreamWriter) Tell() (resultInt64 int64) {
	streamWriterImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefStreamWriter) Flush() int32 {
	r1 := streamWriterImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCefStreamWriter) MayBlock() bool {
	r1 := streamWriterImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

var (
	streamWriterImport       *imports.Imports = nil
	streamWriterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStreamWriter_CreateForFile", 0),
		/*1*/ imports.NewTable("CefStreamWriter_Flush", 0),
		/*2*/ imports.NewTable("CefStreamWriter_MayBlock", 0),
		/*3*/ imports.NewTable("CefStreamWriter_Seek", 0),
		/*4*/ imports.NewTable("CefStreamWriter_Tell", 0),
		/*5*/ imports.NewTable("CefStreamWriter_UnWrap", 0),
		/*6*/ imports.NewTable("CefStreamWriter_Write", 0),
	}
)

func streamWriterImportAPI() *imports.Imports {
	if streamWriterImport == nil {
		streamWriterImport = NewDefaultImports()
		streamWriterImport.SetImportTable(streamWriterImportTables)
		streamWriterImportTables = nil
	}
	return streamWriterImport
}
