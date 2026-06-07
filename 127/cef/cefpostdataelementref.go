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

// ICefPostDataElement Parent: ICefBaseRefCounted
type ICefPostDataElement interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetType
	//  Return the type of this post data element.
	GetType() cefTypes.TCefPostDataElementType // function
	// GetFile
	//  Return the file name.
	GetFile() string // function
	// GetBytesCount
	//  Return the number of bytes.
	GetBytesCount() cefTypes.NativeUInt // function
	// GetBytes
	//  Read up to |size| bytes into |bytes| and return the number of bytes
	//  actually read.
	GetBytes(size cefTypes.NativeUInt, bytes uintptr) cefTypes.NativeUInt // function
	// SetToEmpty
	//  Remove all contents from the post data element.
	SetToEmpty() // procedure
	// SetToFile
	//  The post data element will represent a file.
	SetToFile(fileName string) // procedure
	// SetToBytes
	//  The post data element will represent bytes. The bytes passed in will be
	//  copied.
	SetToBytes(size cefTypes.NativeUInt, bytes uintptr) // procedure
}

// ICefPostDataElementRef Parent: ICefPostDataElement ICefBaseRefCountedRef
type ICefPostDataElementRef interface {
	ICefPostDataElement
	ICefBaseRefCountedRef
	AsIntfPostDataElement() uintptr
}

type TCefPostDataElementRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPostDataElementRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPostDataElementRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPostDataElementRef) GetType() cefTypes.TCefPostDataElementType {
	if !m.IsValid() {
		return 0
	}
	r := cefPostDataElementRefAPI().SysCallN(2, m.Instance())
	return cefTypes.TCefPostDataElementType(r)
}

func (m *TCefPostDataElementRef) GetFile() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefPostDataElementRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefPostDataElementRef) GetBytesCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefPostDataElementRefAPI().SysCallN(4, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefPostDataElementRef) GetBytes(size cefTypes.NativeUInt, bytes uintptr) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefPostDataElementRefAPI().SysCallN(5, m.Instance(), uintptr(size), uintptr(bytes))
	return cefTypes.NativeUInt(r)
}

func (m *TCefPostDataElementRef) SetToEmpty() {
	if !m.IsValid() {
		return
	}
	cefPostDataElementRefAPI().SysCallN(8, m.Instance())
}

func (m *TCefPostDataElementRef) SetToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	cefPostDataElementRefAPI().SysCallN(9, m.Instance(), api.PasStr(fileName))
}

func (m *TCefPostDataElementRef) SetToBytes(size cefTypes.NativeUInt, bytes uintptr) {
	if !m.IsValid() {
		return
	}
	cefPostDataElementRefAPI().SysCallN(10, m.Instance(), uintptr(size), uintptr(bytes))
}

func (m *TCefPostDataElementRef) AsIntfPostDataElement() uintptr {
	return m.GetIntfPointer(0)
}

// PostDataElementRef  is static instance
var PostDataElementRef _PostDataElementRefClass

// _PostDataElementRefClass is class type defined by TCefPostDataElementRef
type _PostDataElementRefClass uintptr

func (_PostDataElementRefClass) UnWrap(data uintptr) (result ICefPostDataElement) {
	var resultPtr uintptr
	cefPostDataElementRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPostDataElementRef(resultPtr)
	return
}

func (_PostDataElementRefClass) New() (result ICefPostDataElement) {
	var resultPtr uintptr
	cefPostDataElementRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPostDataElementRef(resultPtr)
	return
}

// NewPostDataElementRef class constructor
func NewPostDataElementRef(data uintptr) ICefPostDataElementRef {
	var postDataElementPtr uintptr // ICefPostDataElement
	r := cefPostDataElementRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&postDataElementPtr)))
	ret := AsCefPostDataElementRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, postDataElementPtr)
	}
	return ret
}

var (
	cefPostDataElementRefOnce   base.Once
	cefPostDataElementRefImport *imports.Imports = nil
)

func cefPostDataElementRefAPI() *imports.Imports {
	cefPostDataElementRefOnce.Do(func() {
		cefPostDataElementRefImport = api.NewDefaultImports()
		cefPostDataElementRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPostDataElementRef_Create", 0), // constructor NewPostDataElementRef
			/* 1 */ imports.NewTable("TCefPostDataElementRef_IsReadOnly", 0), // function IsReadOnly
			/* 2 */ imports.NewTable("TCefPostDataElementRef_GetType", 0), // function GetType
			/* 3 */ imports.NewTable("TCefPostDataElementRef_GetFile", 0), // function GetFile
			/* 4 */ imports.NewTable("TCefPostDataElementRef_GetBytesCount", 0), // function GetBytesCount
			/* 5 */ imports.NewTable("TCefPostDataElementRef_GetBytes", 0), // function GetBytes
			/* 6 */ imports.NewTable("TCefPostDataElementRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefPostDataElementRef_New", 0), // static function New
			/* 8 */ imports.NewTable("TCefPostDataElementRef_SetToEmpty", 0), // procedure SetToEmpty
			/* 9 */ imports.NewTable("TCefPostDataElementRef_SetToFile", 0), // procedure SetToFile
			/* 10 */ imports.NewTable("TCefPostDataElementRef_SetToBytes", 0), // procedure SetToBytes
		}
	})
	return cefPostDataElementRefImport
}
