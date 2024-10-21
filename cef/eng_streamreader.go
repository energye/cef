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

// ICefStreamReader Parent: ICefBaseRefCounted
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))</a>
type ICefStreamReader interface {
	ICefBaseRefCounted
	// Read
	//  Read raw binary data.
	Read(ptr uintptr, size, n NativeUInt) NativeUInt // function
	// Seek
	//  Seek to the specified offset position. |whence| may be any one of SEEK_CUR, SEEK_END or SEEK_SET. Return zero on success and non-zero on failure.
	Seek(offset int64, whence int32) int32 // function
	// Tell
	//  Return the current offset position.
	Tell() (resultInt64 int64) // function
	// Eof
	//  Return non-zero if at end of file.
	Eof() bool // function
	// MayBlock
	//  Return true (1) if this handler performs work like accessing the file system which may block. Used as a hint for determining the thread to access the handler from.
	MayBlock() bool // function
}

// TCefStreamReader Parent: TCefBaseRefCounted
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))</a>
type TCefStreamReader struct {
	TCefBaseRefCounted
}

// StreamReaderRef -> ICefStreamReader
var StreamReaderRef streamReader

// streamReader TCefStreamReader Ref
type streamReader uintptr

func (m *streamReader) UnWrap(data uintptr) ICefStreamReader {
	var resultCefStreamReader uintptr
	streamReaderImportAPI().SysCallN(9, uintptr(data), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForFile(filename string) ICefStreamReader {
	var resultCefStreamReader uintptr
	streamReaderImportAPI().SysCallN(2, PascalStr(filename), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForCustomStream(stream ICefCustomStreamReader) ICefStreamReader {
	var resultCefStreamReader uintptr
	streamReaderImportAPI().SysCallN(0, GetObjectUintptr(stream), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForStream(stream IStream, owned bool) ICefStreamReader {
	var resultCefStreamReader uintptr
	streamReaderImportAPI().SysCallN(3, GetObjectUintptr(stream), PascalBool(owned), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForData(data uintptr, size NativeUInt) ICefStreamReader {
	var resultCefStreamReader uintptr
	streamReaderImportAPI().SysCallN(1, uintptr(data), uintptr(size), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *TCefStreamReader) Read(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := streamReaderImportAPI().SysCallN(6, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefStreamReader) Seek(offset int64, whence int32) int32 {
	r1 := streamReaderImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefStreamReader) Tell() (resultInt64 int64) {
	streamReaderImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefStreamReader) Eof() bool {
	r1 := streamReaderImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefStreamReader) MayBlock() bool {
	r1 := streamReaderImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

var (
	streamReaderImport       *imports.Imports = nil
	streamReaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStreamReader_CreateForCustomStream", 0),
		/*1*/ imports.NewTable("CefStreamReader_CreateForData", 0),
		/*2*/ imports.NewTable("CefStreamReader_CreateForFile", 0),
		/*3*/ imports.NewTable("CefStreamReader_CreateForStream", 0),
		/*4*/ imports.NewTable("CefStreamReader_Eof", 0),
		/*5*/ imports.NewTable("CefStreamReader_MayBlock", 0),
		/*6*/ imports.NewTable("CefStreamReader_Read", 0),
		/*7*/ imports.NewTable("CefStreamReader_Seek", 0),
		/*8*/ imports.NewTable("CefStreamReader_Tell", 0),
		/*9*/ imports.NewTable("CefStreamReader_UnWrap", 0),
	}
)

func streamReaderImportAPI() *imports.Imports {
	if streamReaderImport == nil {
		streamReaderImport = NewDefaultImports()
		streamReaderImport.SetImportTable(streamReaderImportTables)
		streamReaderImportTables = nil
	}
	return streamReaderImport
}
