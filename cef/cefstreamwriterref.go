//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefStreamWriter Parent: ICefBaseRefCountedRef
type ICefStreamWriter interface {
	ICefBaseRefCountedRef
	// Write
	//  Write raw binary data.
	Write(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt // function
	// Seek
	//  Seek to the specified offset position. |whence| may be any one of
	//  SEEK_CUR, SEEK_END or SEEK_SET. Returns zero on success and non-zero on
	//  failure.
	Seek(offset int64, whence int32) int32 // function
	// Tell
	//  Return the current offset position.
	Tell() int64 // function
	// Flush
	//  Flush the stream.
	Flush() int32 // function
	// MayBlock
	//  Returns true (1) if this writer performs work like accessing the file
	//  system which may block. Used as a hint for determining the thread to
	//  access the writer from.
	MayBlock() bool // function
}

// ICefStreamWriterRef Parent: ICefStreamWriter
type ICefStreamWriterRef interface {
	ICefStreamWriter
	AsIntfStreamWriter() uintptr
}

type TCefStreamWriterRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefStreamWriterRef) Write(ptr uintptr, size cefTypes.NativeUInt, N cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefStreamWriterRefAPI().SysCallN(1, m.Instance(), uintptr(ptr), uintptr(size), uintptr(N))
	return cefTypes.NativeUInt(r)
}

func (m *TCefStreamWriterRef) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefStreamWriterRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(whence))
	return int32(r)
}

func (m *TCefStreamWriterRef) Tell() (result int64) {
	if !m.IsValid() {
		return
	}
	cefStreamWriterRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefStreamWriterRef) Flush() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefStreamWriterRefAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCefStreamWriterRef) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r := cefStreamWriterRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefStreamWriterRef) AsIntfStreamWriter() uintptr {
	return m.GetIntfPointer(0)
}

// StreamWriterRef  is static instance
var StreamWriterRef _StreamWriterRefClass

// _StreamWriterRefClass is class type defined by TCefStreamWriterRef
type _StreamWriterRefClass uintptr

func (_StreamWriterRefClass) UnWrap(data uintptr) (result ICefStreamWriter) {
	var resultPtr uintptr
	cefStreamWriterRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamWriterRef(resultPtr)
	return
}

func (_StreamWriterRefClass) CreateForFile(fileName string) (result ICefStreamWriter) {
	var resultPtr uintptr
	cefStreamWriterRefAPI().SysCallN(7, api.PasStr(fileName), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefStreamWriterRef(resultPtr)
	return
}

// NewStreamWriterRef class constructor
func NewStreamWriterRef(data uintptr) ICefStreamWriterRef {
	var streamWriterPtr uintptr // ICefStreamWriter
	r := cefStreamWriterRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&streamWriterPtr)))
	ret := AsCefStreamWriterRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, streamWriterPtr)
	}
	return ret
}

var (
	cefStreamWriterRefOnce   base.Once
	cefStreamWriterRefImport *imports.Imports = nil
)

func cefStreamWriterRefAPI() *imports.Imports {
	cefStreamWriterRefOnce.Do(func() {
		cefStreamWriterRefImport = api.NewDefaultImports()
		cefStreamWriterRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStreamWriterRef_Create", 0), // constructor NewStreamWriterRef
			/* 1 */ imports.NewTable("TCefStreamWriterRef_write", 0), // function Write
			/* 2 */ imports.NewTable("TCefStreamWriterRef_Seek", 0), // function Seek
			/* 3 */ imports.NewTable("TCefStreamWriterRef_Tell", 0), // function Tell
			/* 4 */ imports.NewTable("TCefStreamWriterRef_Flush", 0), // function Flush
			/* 5 */ imports.NewTable("TCefStreamWriterRef_MayBlock", 0), // function MayBlock
			/* 6 */ imports.NewTable("TCefStreamWriterRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefStreamWriterRef_CreateForFile", 0), // static function CreateForFile
		}
	})
	return cefStreamWriterRefImport
}
