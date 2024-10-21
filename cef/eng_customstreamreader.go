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

// ICefCustomStreamReader Parent: ICefBaseRefCountedOwn
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))</a>
type ICefCustomStreamReader interface {
	ICefBaseRefCountedOwn
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefCustomStreamReader // procedure
}

// TCefCustomStreamReader Parent: TCefBaseRefCountedOwn
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))</a>
type TCefCustomStreamReader struct {
	TCefBaseRefCountedOwn
}

func NewCefCustomStreamReader(stream IStream, owned bool) ICefCustomStreamReader {
	r1 := customStreamReaderImportAPI().SysCallN(1, GetObjectUintptr(stream), PascalBool(owned))
	return AsCefCustomStreamReader(r1)
}

func NewCefCustomStreamReader1(filename string) ICefCustomStreamReader {
	r1 := customStreamReaderImportAPI().SysCallN(2, PascalStr(filename))
	return AsCefCustomStreamReader(r1)
}

func (m *TCefCustomStreamReader) AsInterface() ICefCustomStreamReader {
	var resultCefCustomStreamReader uintptr
	customStreamReaderImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefCustomStreamReader)))
	return AsCefCustomStreamReader(resultCefCustomStreamReader)
}

var (
	customStreamReaderImport       *imports.Imports = nil
	customStreamReaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomStreamReader_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomStreamReader_Create", 0),
		/*2*/ imports.NewTable("CefCustomStreamReader_Create1", 0),
	}
)

func customStreamReaderImportAPI() *imports.Imports {
	if customStreamReaderImport == nil {
		customStreamReaderImport = NewDefaultImports()
		customStreamReaderImport.SetImportTable(customStreamReaderImportTables)
		customStreamReaderImportTables = nil
	}
	return customStreamReaderImport
}
