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

// ICefBinaryValue Parent: ICefBaseRefCounted
type ICefBinaryValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. This object may become invalid
	//  if the underlying data is owned by another object (e.g. list or
	//  dictionary) and that other object is then modified or destroyed. Do not
	//  call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if this object is currently owned by another object.
	IsOwned() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying
	//  data.
	IsSame(that ICefBinaryValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent
	//  underlying value but are not necessarily the same object.
	IsEqual(that ICefBinaryValue) bool // function
	// Copy
	//  Returns a copy of this object. The data in this object will also be
	//  copied.
	Copy() ICefBinaryValue // function
	// GetRawData
	//  Returns a pointer to the beginning of the memory block. The returned
	//  pointer is valid as long as the ICefBinaryValue is alive.
	GetRawData() uintptr // function
	// GetSize
	//  Returns the data size.
	GetSize() cefTypes.NativeUInt // function
	// GetData
	//  Read up to |buffer_size| number of bytes into |buffer|. Reading begins at
	//  the specified byte |data_offset|. Returns the number of bytes read.
	GetData(buffer uintptr, bufferSize cefTypes.NativeUInt, dataOffset cefTypes.NativeUInt) cefTypes.NativeUInt // function
}

// ICefBinaryValueRef Parent: ICefBinaryValue ICefBaseRefCountedRef
type ICefBinaryValueRef interface {
	ICefBinaryValue
	ICefBaseRefCountedRef
	AsIntfBinaryValue() uintptr
}

type TCefBinaryValueRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefBinaryValueRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefBinaryValueRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBinaryValueRef) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBinaryValueRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBinaryValueRef) IsSame(that ICefBinaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBinaryValueRefAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefBinaryValueRef) IsEqual(that ICefBinaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBinaryValueRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefBinaryValueRef) Copy() (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBinaryValueRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefBinaryValueRef) GetRawData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefBinaryValueRefAPI().SysCallN(6, m.Instance())
	return uintptr(r)
}

func (m *TCefBinaryValueRef) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBinaryValueRefAPI().SysCallN(7, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefBinaryValueRef) GetData(buffer uintptr, bufferSize cefTypes.NativeUInt, dataOffset cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBinaryValueRefAPI().SysCallN(8, m.Instance(), uintptr(buffer), uintptr(bufferSize), uintptr(dataOffset))
	return cefTypes.NativeUInt(r)
}

func (m *TCefBinaryValueRef) AsIntfBinaryValue() uintptr {
	return m.GetIntfPointer(0)
}

// BinaryValueRef  is static instance
var BinaryValueRef _BinaryValueRefClass

// _BinaryValueRefClass is class type defined by TCefBinaryValueRef
type _BinaryValueRefClass uintptr

func (_BinaryValueRefClass) UnWrap(data uintptr) (result ICefBinaryValue) {
	var resultPtr uintptr
	cefBinaryValueRefAPI().SysCallN(9, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (_BinaryValueRefClass) New(data uintptr, dataSize cefTypes.NativeUInt) (result ICefBinaryValue) {
	var resultPtr uintptr
	cefBinaryValueRefAPI().SysCallN(10, uintptr(data), uintptr(dataSize), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

// NewBinaryValueRef class constructor
func NewBinaryValueRef(data uintptr) ICefBinaryValueRef {
	var binaryValuePtr uintptr // ICefBinaryValue
	r := cefBinaryValueRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&binaryValuePtr)))
	ret := AsCefBinaryValueRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, binaryValuePtr)
	}
	return ret
}

var (
	cefBinaryValueRefOnce   base.Once
	cefBinaryValueRefImport *imports.Imports = nil
)

func cefBinaryValueRefAPI() *imports.Imports {
	cefBinaryValueRefOnce.Do(func() {
		cefBinaryValueRefImport = api.NewDefaultImports()
		cefBinaryValueRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBinaryValueRef_Create", 0), // constructor NewBinaryValueRef
			/* 1 */ imports.NewTable("TCefBinaryValueRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefBinaryValueRef_IsOwned", 0), // function IsOwned
			/* 3 */ imports.NewTable("TCefBinaryValueRef_IsSame", 0), // function IsSame
			/* 4 */ imports.NewTable("TCefBinaryValueRef_IsEqual", 0), // function IsEqual
			/* 5 */ imports.NewTable("TCefBinaryValueRef_Copy", 0), // function Copy
			/* 6 */ imports.NewTable("TCefBinaryValueRef_GetRawData", 0), // function GetRawData
			/* 7 */ imports.NewTable("TCefBinaryValueRef_GetSize", 0), // function GetSize
			/* 8 */ imports.NewTable("TCefBinaryValueRef_GetData", 0), // function GetData
			/* 9 */ imports.NewTable("TCefBinaryValueRef_UnWrap", 0), // static function UnWrap
			/* 10 */ imports.NewTable("TCefBinaryValueRef_New", 0), // static function New
		}
	})
	return cefBinaryValueRefImport
}
