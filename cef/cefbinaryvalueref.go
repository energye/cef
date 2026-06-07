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

// ICefBinaryValue Parent: ICefBaseRefCounted
type ICefBinaryValue interface {
	ICefBaseRefCounted
	IsValid() bool                                                                                              // function
	IsOwned() bool                                                                                              // function
	IsSame(that ICefBinaryValue) bool                                                                           // function
	IsEqual(that ICefBinaryValue) bool                                                                          // function
	Copy() ICefBinaryValue                                                                                      // function
	GetSize() cefTypes.NativeUInt                                                                               // function
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

func (m *TCefBinaryValueRef) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBinaryValueRefAPI().SysCallN(6, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefBinaryValueRef) GetData(buffer uintptr, bufferSize cefTypes.NativeUInt, dataOffset cefTypes.NativeUInt) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefBinaryValueRefAPI().SysCallN(7, m.Instance(), uintptr(buffer), uintptr(bufferSize), uintptr(dataOffset))
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
	cefBinaryValueRefAPI().SysCallN(8, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (_BinaryValueRefClass) New(data uintptr, dataSize cefTypes.NativeUInt) (result ICefBinaryValue) {
	var resultPtr uintptr
	cefBinaryValueRefAPI().SysCallN(9, uintptr(data), uintptr(dataSize), uintptr(base.UnsafePointer(&resultPtr)))
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
			/* 6 */ imports.NewTable("TCefBinaryValueRef_GetSize", 0), // function GetSize
			/* 7 */ imports.NewTable("TCefBinaryValueRef_GetData", 0), // function GetData
			/* 8 */ imports.NewTable("TCefBinaryValueRef_UnWrap", 0), // static function UnWrap
			/* 9 */ imports.NewTable("TCefBinaryValueRef_New", 0), // static function New
		}
	})
	return cefBinaryValueRefImport
}
