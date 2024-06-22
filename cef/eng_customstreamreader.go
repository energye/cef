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
type ICefCustomStreamReader interface {
	ICefBaseRefCountedOwn
}

// TCefCustomStreamReader Parent: TCefBaseRefCountedOwn
type TCefCustomStreamReader struct {
	TCefBaseRefCountedOwn
}

func NewCefCustomStreamReader(stream IStream, owned bool) ICefCustomStreamReader {
	r1 := customStreamReaderImportAPI().SysCallN(0, GetObjectUintptr(stream), PascalBool(owned))
	return AsCefCustomStreamReader(r1)
}

func NewCefCustomStreamReader1(filename string) ICefCustomStreamReader {
	r1 := customStreamReaderImportAPI().SysCallN(1, PascalStr(filename))
	return AsCefCustomStreamReader(r1)
}

var (
	customStreamReaderImport       *imports.Imports = nil
	customStreamReaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomStreamReader_Create", 0),
		/*1*/ imports.NewTable("CefCustomStreamReader_Create1", 0),
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
