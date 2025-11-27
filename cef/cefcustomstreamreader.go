//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/types"
)

// ICefCustomStreamReader0 Parent: ICefBaseRefCountedOwn
type ICefCustomStreamReader0 interface {
	ICefBaseRefCountedOwn
	// Read
	//  Read raw binary data.
	Read(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt // function
	// Seek
	//  Seek to the specified offset position. |whence| may be any one of
	//  SEEK_CUR, SEEK_END or SEEK_SET. Return zero on success and non-zero on
	//  failure.
	Seek(offset int64, whence int32) int32 // function
	// Tell
	//  Return the current offset position.
	Tell() int64 // function
	// Eof
	//  Return non-zero if at end of file.
	Eof() bool // function
	// MayBlock
	//  Return true (1) if this handler performs work like accessing the file
	//  system which may block. Used as a hint for determining the thread to
	//  access the handler from.
	MayBlock() bool // function
}

// ICefCustomStreamReader Parent: ICefCustomStreamReader0
type ICefCustomStreamReader interface {
	ICefCustomStreamReader0
	AsIntfCustomStreamReader() uintptr
}

type TCefCustomStreamReader struct {
	TCefBaseRefCountedOwn
}

func (m *TCefCustomStreamReader) Read(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStreamReaderAPI().SysCallN(2, m.Instance(), uintptr(ptr), uintptr(size), uintptr(N))
	return cefTypes.NativeUInt(r)
}

func (m *TCefCustomStreamReader) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStreamReaderAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(whence))
	return int32(r)
}

func (m *TCefCustomStreamReader) Tell() (result int64) {
	if !m.IsValid() {
		return
	}
	cefCustomStreamReaderAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefCustomStreamReader) Eof() bool {
	if !m.IsValid() {
		return false
	}
	r := cefCustomStreamReaderAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCustomStreamReader) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r := cefCustomStreamReaderAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefCustomStreamReader) AsIntfCustomStreamReader() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomStreamReader class constructor
func NewCustomStreamReader(stream lcl.IStream, owned bool) ICefCustomStreamReader {
	var customStreamReaderPtr uintptr // ICefCustomStreamReader
	r := cefCustomStreamReaderAPI().SysCallN(0, base.GetObjectUintptr(stream), api.PasBool(owned), uintptr(base.UnsafePointer(&customStreamReaderPtr)))
	ret := AsCefCustomStreamReader(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customStreamReaderPtr)
	}
	return ret
}

// NewCustomStreamReaderWithString class constructor
func NewCustomStreamReaderWithString(filename string) ICefCustomStreamReader {
	var customStreamReaderPtr uintptr // ICefCustomStreamReader
	r := cefCustomStreamReaderAPI().SysCallN(1, api.PasStr(filename), uintptr(base.UnsafePointer(&customStreamReaderPtr)))
	ret := AsCefCustomStreamReader(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customStreamReaderPtr)
	}
	return ret
}

var (
	cefCustomStreamReaderOnce   base.Once
	cefCustomStreamReaderImport *imports.Imports = nil
)

func cefCustomStreamReaderAPI() *imports.Imports {
	cefCustomStreamReaderOnce.Do(func() {
		cefCustomStreamReaderImport = api.NewDefaultImports()
		cefCustomStreamReaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomStreamReader_Create", 0), // constructor NewCustomStreamReader
			/* 1 */ imports.NewTable("TCefCustomStreamReader_CreateWithString", 0), // constructor NewCustomStreamReaderWithString
			/* 2 */ imports.NewTable("TCefCustomStreamReader_Read", 0), // function Read
			/* 3 */ imports.NewTable("TCefCustomStreamReader_Seek", 0), // function Seek
			/* 4 */ imports.NewTable("TCefCustomStreamReader_Tell", 0), // function Tell
			/* 5 */ imports.NewTable("TCefCustomStreamReader_Eof", 0), // function Eof
			/* 6 */ imports.NewTable("TCefCustomStreamReader_MayBlock", 0), // function MayBlock
		}
	})
	return cefCustomStreamReaderImport
}
