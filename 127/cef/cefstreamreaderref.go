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

	cefTypes "github.com/energye/cef/127/types"
)

// ICefStreamReader Parent: ICefBaseRefCounted
type ICefStreamReader interface {
	ICefBaseRefCounted
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

// ICefStreamReaderRef Parent: ICefStreamReader ICefBaseRefCountedRef
type ICefStreamReaderRef interface {
	ICefStreamReader
	ICefBaseRefCountedRef
	AsIntfStreamReader() uintptr
}

type TCefStreamReaderRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefStreamReaderRef) Read(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefStreamReaderRefAPI().SysCallN(1, m.Instance(), uintptr(ptr), uintptr(size), uintptr(N))
	return cefTypes.NativeUInt(r)
}

func (m *TCefStreamReaderRef) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefStreamReaderRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(whence))
	return int32(r)
}

func (m *TCefStreamReaderRef) Tell() (result int64) {
	if !m.IsValid() {
		return
	}
	cefStreamReaderRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefStreamReaderRef) Eof() bool {
	if !m.IsValid() {
		return false
	}
	r := cefStreamReaderRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefStreamReaderRef) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r := cefStreamReaderRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefStreamReaderRef) AsIntfStreamReader() uintptr {
	return m.GetIntfPointer(0)
}

// StreamReaderRef  is static instance
var StreamReaderRef _StreamReaderRefClass

// _StreamReaderRefClass is class type defined by TCefStreamReaderRef
type _StreamReaderRefClass uintptr

func (_StreamReaderRefClass) UnWrap(data uintptr) (result ICefStreamReader) {
	var resultPtr uintptr
	cefStreamReaderRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamReaderRef(resultPtr)
	return
}

func (_StreamReaderRefClass) CreateForFile(filename string) (result ICefStreamReader) {
	var resultPtr uintptr
	cefStreamReaderRefAPI().SysCallN(7, api.PasStr(filename), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamReaderRef(resultPtr)
	return
}

func (_StreamReaderRefClass) CreateForCustomStream(stream ICefCustomStreamReader) (result ICefStreamReader) {
	var resultPtr uintptr
	cefStreamReaderRefAPI().SysCallN(8, base.GetObjectUintptr(stream), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamReaderRef(resultPtr)
	return
}

func (_StreamReaderRefClass) CreateForStream(stream lcl.IStream, owned bool) (result ICefStreamReader) {
	var resultPtr uintptr
	cefStreamReaderRefAPI().SysCallN(9, base.GetObjectUintptr(stream), api.PasBool(owned), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamReaderRef(resultPtr)
	return
}

func (_StreamReaderRefClass) CreateForData(data uintptr, size cefTypes.NativeUInt) (result ICefStreamReader) {
	var resultPtr uintptr
	cefStreamReaderRefAPI().SysCallN(10, uintptr(data), uintptr(size), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamReaderRef(resultPtr)
	return
}

// NewStreamReaderRef class constructor
func NewStreamReaderRef(data uintptr) ICefStreamReaderRef {
	var streamReaderPtr uintptr // ICefStreamReader
	r := cefStreamReaderRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&streamReaderPtr)))
	ret := AsCefStreamReaderRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, streamReaderPtr)
	}
	return ret
}

var (
	cefStreamReaderRefOnce   base.Once
	cefStreamReaderRefImport *imports.Imports = nil
)

func cefStreamReaderRefAPI() *imports.Imports {
	cefStreamReaderRefOnce.Do(func() {
		cefStreamReaderRefImport = api.NewDefaultImports()
		cefStreamReaderRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStreamReaderRef_Create", 0), // constructor NewStreamReaderRef
			/* 1 */ imports.NewTable("TCefStreamReaderRef_Read", 0), // function Read
			/* 2 */ imports.NewTable("TCefStreamReaderRef_Seek", 0), // function Seek
			/* 3 */ imports.NewTable("TCefStreamReaderRef_Tell", 0), // function Tell
			/* 4 */ imports.NewTable("TCefStreamReaderRef_Eof", 0), // function Eof
			/* 5 */ imports.NewTable("TCefStreamReaderRef_MayBlock", 0), // function MayBlock
			/* 6 */ imports.NewTable("TCefStreamReaderRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefStreamReaderRef_CreateForFile", 0), // static function CreateForFile
			/* 8 */ imports.NewTable("TCefStreamReaderRef_CreateForCustomStream", 0), // static function CreateForCustomStream
			/* 9 */ imports.NewTable("TCefStreamReaderRef_CreateForStream", 0), // static function CreateForStream
			/* 10 */ imports.NewTable("TCefStreamReaderRef_CreateForData", 0), // static function CreateForData
		}
	})
	return cefStreamReaderRefImport
}
