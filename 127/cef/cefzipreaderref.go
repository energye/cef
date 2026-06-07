//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/127/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefZipReader Parent: ICefBaseRefCounted
type ICefZipReader interface {
	ICefBaseRefCounted
	// MoveToFirstFile
	//  Moves the cursor to the first file in the archive. Returns true (1) if the
	//  cursor position was set successfully.
	MoveToFirstFile() bool // function
	// MoveToNextFile
	//  Moves the cursor to the next file in the archive. Returns true (1) if the
	//  cursor position was set successfully.
	MoveToNextFile() bool // function
	// MoveToFile
	//  Moves the cursor to the specified file in the archive. If |caseSensitive|
	//  is true (1) then the search will be case sensitive. Returns true (1) if
	//  the cursor position was set successfully.
	MoveToFile(fileName string, caseSensitive bool) bool // function
	// Close
	//  Closes the archive. This should be called directly to ensure that cleanup
	//  occurs on the correct thread.
	Close() bool // function
	// GetFileName
	//  Returns the name of the file.
	GetFileName() string // function
	// GetFileSize
	//  Returns the uncompressed size of the file.
	GetFileSize() int64 // function
	// GetFileLastModified
	//  Returns the last modified timestamp for the file.
	GetFileLastModified() int64 // function
	// OpenFile
	//  Opens the file for reading of uncompressed data. A read password may
	//  optionally be specified.
	OpenFile(password string) bool // function
	// CloseFile
	//  Closes the file.
	CloseFile() bool // function
	// ReadFile
	//  Read uncompressed file contents into the specified buffer. Returns < 0 if
	//  an error occurred, 0 if at the end of file, or the number of bytes read.
	ReadFile(buffer uintptr, bufferSize cefTypes.NativeUInt) int32 // function
	// Tell
	//  Returns the current offset in the uncompressed file contents.
	Tell() int64 // function
	// Eof
	//  Returns true (1) if at end of the file contents.
	Eof() bool // function
}

// ICefZipReaderRef Parent: ICefZipReader ICefBaseRefCountedRef
type ICefZipReaderRef interface {
	ICefZipReader
	ICefBaseRefCountedRef
	AsIntfZipReader() uintptr
}

type TCefZipReaderRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefZipReaderRef) MoveToFirstFile() bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) MoveToNextFile() bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) MoveToFile(fileName string, caseSensitive bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(3, m.Instance(), api.PasStr(fileName), api.PasBool(caseSensitive))
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) Close() bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) GetFileName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefZipReaderRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefZipReaderRef) GetFileSize() (result int64) {
	if !m.IsValid() {
		return
	}
	cefZipReaderRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefZipReaderRef) GetFileLastModified() (result int64) {
	if !m.IsValid() {
		return
	}
	cefZipReaderRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefZipReaderRef) OpenFile(password string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(8, m.Instance(), api.PasStr(password))
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) CloseFile() bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) ReadFile(buffer uintptr, bufferSize cefTypes.NativeUInt) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefZipReaderRefAPI().SysCallN(10, m.Instance(), uintptr(buffer), uintptr(bufferSize))
	return int32(r)
}

func (m *TCefZipReaderRef) Tell() (result int64) {
	if !m.IsValid() {
		return
	}
	cefZipReaderRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefZipReaderRef) Eof() bool {
	if !m.IsValid() {
		return false
	}
	r := cefZipReaderRefAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCefZipReaderRef) AsIntfZipReader() uintptr {
	return m.GetIntfPointer(0)
}

// ZipReaderRef  is static instance
var ZipReaderRef _ZipReaderRefClass

// _ZipReaderRefClass is class type defined by TCefZipReaderRef
type _ZipReaderRefClass uintptr

func (_ZipReaderRefClass) UnWrap(data uintptr) (result ICefZipReader) {
	var resultPtr uintptr
	cefZipReaderRefAPI().SysCallN(13, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefZipReaderRef(resultPtr)
	return
}

func (_ZipReaderRefClass) New(stream ICefStreamReader) (result ICefZipReader) {
	var resultPtr uintptr
	cefZipReaderRefAPI().SysCallN(14, base.GetObjectUintptr(stream), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefZipReaderRef(resultPtr)
	return
}

// NewZipReaderRef class constructor
func NewZipReaderRef(data uintptr) ICefZipReaderRef {
	var zipReaderPtr uintptr // ICefZipReader
	r := cefZipReaderRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&zipReaderPtr)))
	ret := AsCefZipReaderRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, zipReaderPtr)
	}
	return ret
}

var (
	cefZipReaderRefOnce   base.Once
	cefZipReaderRefImport *imports.Imports = nil
)

func cefZipReaderRefAPI() *imports.Imports {
	cefZipReaderRefOnce.Do(func() {
		cefZipReaderRefImport = api.NewDefaultImports()
		cefZipReaderRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefZipReaderRef_Create", 0), // constructor NewZipReaderRef
			/* 1 */ imports.NewTable("TCefZipReaderRef_MoveToFirstFile", 0), // function MoveToFirstFile
			/* 2 */ imports.NewTable("TCefZipReaderRef_MoveToNextFile", 0), // function MoveToNextFile
			/* 3 */ imports.NewTable("TCefZipReaderRef_MoveToFile", 0), // function MoveToFile
			/* 4 */ imports.NewTable("TCefZipReaderRef_Close", 0), // function Close
			/* 5 */ imports.NewTable("TCefZipReaderRef_GetFileName", 0), // function GetFileName
			/* 6 */ imports.NewTable("TCefZipReaderRef_GetFileSize", 0), // function GetFileSize
			/* 7 */ imports.NewTable("TCefZipReaderRef_GetFileLastModified", 0), // function GetFileLastModified
			/* 8 */ imports.NewTable("TCefZipReaderRef_OpenFile", 0), // function OpenFile
			/* 9 */ imports.NewTable("TCefZipReaderRef_CloseFile", 0), // function CloseFile
			/* 10 */ imports.NewTable("TCefZipReaderRef_ReadFile", 0), // function ReadFile
			/* 11 */ imports.NewTable("TCefZipReaderRef_Tell", 0), // function Tell
			/* 12 */ imports.NewTable("TCefZipReaderRef_Eof", 0), // function Eof
			/* 13 */ imports.NewTable("TCefZipReaderRef_UnWrap", 0), // static function UnWrap
			/* 14 */ imports.NewTable("TCefZipReaderRef_New", 0), // static function New
		}
	})
	return cefZipReaderRefImport
}
