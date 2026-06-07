//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefListValue Parent: ICefBaseRefCounted
type ICefListValue interface {
	ICefBaseRefCounted
	IsValid() bool                                                           // function
	IsOwned() bool                                                           // function
	IsReadOnly() bool                                                        // function
	IsSame(that ICefListValue) bool                                          // function
	IsEqual(that ICefListValue) bool                                         // function
	Copy() ICefListValue                                                     // function
	SetSize(size cefTypes.NativeUInt) bool                                   // function
	GetSize() cefTypes.NativeUInt                                            // function
	Clear() bool                                                             // function
	Remove(index cefTypes.NativeUInt) bool                                   // function
	GetType(index cefTypes.NativeUInt) cefTypes.TCefValueType                // function
	GetValue(index cefTypes.NativeUInt) ICefValue                            // function
	GetBool(index cefTypes.NativeUInt) bool                                  // function
	GetInt(index cefTypes.NativeUInt) int32                                  // function
	GetDouble(index cefTypes.NativeUInt) float64                             // function
	GetString(index cefTypes.NativeUInt) string                              // function
	GetBinary(index cefTypes.NativeUInt) ICefBinaryValue                     // function
	GetDictionary(index cefTypes.NativeUInt) ICefDictionaryValue             // function
	GetList(index cefTypes.NativeUInt) ICefListValue                         // function
	SetValue(index cefTypes.NativeUInt, value ICefValue) bool                // function
	SetNull(index cefTypes.NativeUInt) bool                                  // function
	SetBool(index cefTypes.NativeUInt, value bool) bool                      // function
	SetInt(index cefTypes.NativeUInt, value int32) bool                      // function
	SetDouble(index cefTypes.NativeUInt, value float64) bool                 // function
	SetString(index cefTypes.NativeUInt, value string) bool                  // function
	SetBinary(index cefTypes.NativeUInt, value ICefBinaryValue) bool         // function
	SetDictionary(index cefTypes.NativeUInt, value ICefDictionaryValue) bool // function
	SetList(index cefTypes.NativeUInt, value ICefListValue) bool             // function
}

// ICefListValueRef Parent: ICefListValue ICefBaseRefCountedRef
type ICefListValueRef interface {
	ICefListValue
	ICefBaseRefCountedRef
	AsIntfListValue() uintptr
}

type TCefListValueRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefListValueRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefListValueRef) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefListValueRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefListValueRef) IsSame(that ICefListValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefListValueRef) IsEqual(that ICefListValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefListValueRef) Copy() (result ICefListValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (m *TCefListValueRef) SetSize(size cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(7, m.Instance(), uintptr(size))
	return api.GoBool(r)
}

func (m *TCefListValueRef) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefListValueRefAPI().SysCallN(8, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefListValueRef) Clear() bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCefListValueRef) Remove(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(10, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefListValueRef) GetType(index cefTypes.NativeUInt) cefTypes.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r := cefListValueRefAPI().SysCallN(11, m.Instance(), uintptr(index))
	return cefTypes.TCefValueType(r)
}

func (m *TCefListValueRef) GetValue(index cefTypes.NativeUInt) (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(12, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TCefListValueRef) GetBool(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(13, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefListValueRef) GetInt(index cefTypes.NativeUInt) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefListValueRefAPI().SysCallN(14, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TCefListValueRef) GetDouble(index cefTypes.NativeUInt) (result float64) {
	if !m.IsValid() {
		return
	}
	cefListValueRefAPI().SysCallN(15, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefListValueRef) GetString(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefListValueRefAPI().SysCallN(16, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefListValueRef) GetBinary(index cefTypes.NativeUInt) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(17, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefListValueRef) GetDictionary(index cefTypes.NativeUInt) (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(18, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefListValueRef) GetList(index cefTypes.NativeUInt) (result ICefListValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(19, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (m *TCefListValueRef) SetValue(index cefTypes.NativeUInt, value ICefValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(20, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetNull(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(21, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetBool(index cefTypes.NativeUInt, value bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(22, m.Instance(), uintptr(index), api.PasBool(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetInt(index cefTypes.NativeUInt, value int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(23, m.Instance(), uintptr(index), uintptr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetDouble(index cefTypes.NativeUInt, value float64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(24, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetString(index cefTypes.NativeUInt, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(25, m.Instance(), uintptr(index), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetBinary(index cefTypes.NativeUInt, value ICefBinaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(26, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetDictionary(index cefTypes.NativeUInt, value ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(27, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) SetList(index cefTypes.NativeUInt, value ICefListValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefListValueRefAPI().SysCallN(28, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefListValueRef) AsIntfListValue() uintptr {
	return m.GetIntfPointer(0)
}

// ListValueRef  is static instance
var ListValueRef _ListValueRefClass

// _ListValueRefClass is class type defined by TCefListValueRef
type _ListValueRefClass uintptr

func (_ListValueRefClass) UnWrap(data uintptr) (result ICefListValue) {
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(29, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (_ListValueRefClass) New() (result ICefListValue) {
	var resultPtr uintptr
	cefListValueRefAPI().SysCallN(30, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

// NewListValueRef class constructor
func NewListValueRef(data uintptr) ICefListValueRef {
	var listValuePtr uintptr // ICefListValue
	r := cefListValueRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&listValuePtr)))
	ret := AsCefListValueRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, listValuePtr)
	}
	return ret
}

var (
	cefListValueRefOnce   base.Once
	cefListValueRefImport *imports.Imports = nil
)

func cefListValueRefAPI() *imports.Imports {
	cefListValueRefOnce.Do(func() {
		cefListValueRefImport = api.NewDefaultImports()
		cefListValueRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefListValueRef_Create", 0), // constructor NewListValueRef
			/* 1 */ imports.NewTable("TCefListValueRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefListValueRef_IsOwned", 0), // function IsOwned
			/* 3 */ imports.NewTable("TCefListValueRef_IsReadOnly", 0), // function IsReadOnly
			/* 4 */ imports.NewTable("TCefListValueRef_IsSame", 0), // function IsSame
			/* 5 */ imports.NewTable("TCefListValueRef_IsEqual", 0), // function IsEqual
			/* 6 */ imports.NewTable("TCefListValueRef_Copy", 0), // function Copy
			/* 7 */ imports.NewTable("TCefListValueRef_SetSize", 0), // function SetSize
			/* 8 */ imports.NewTable("TCefListValueRef_GetSize", 0), // function GetSize
			/* 9 */ imports.NewTable("TCefListValueRef_Clear", 0), // function Clear
			/* 10 */ imports.NewTable("TCefListValueRef_Remove", 0), // function Remove
			/* 11 */ imports.NewTable("TCefListValueRef_GetType", 0), // function GetType
			/* 12 */ imports.NewTable("TCefListValueRef_GetValue", 0), // function GetValue
			/* 13 */ imports.NewTable("TCefListValueRef_GetBool", 0), // function GetBool
			/* 14 */ imports.NewTable("TCefListValueRef_GetInt", 0), // function GetInt
			/* 15 */ imports.NewTable("TCefListValueRef_GetDouble", 0), // function GetDouble
			/* 16 */ imports.NewTable("TCefListValueRef_GetString", 0), // function GetString
			/* 17 */ imports.NewTable("TCefListValueRef_GetBinary", 0), // function GetBinary
			/* 18 */ imports.NewTable("TCefListValueRef_GetDictionary", 0), // function GetDictionary
			/* 19 */ imports.NewTable("TCefListValueRef_GetList", 0), // function GetList
			/* 20 */ imports.NewTable("TCefListValueRef_SetValue", 0), // function SetValue
			/* 21 */ imports.NewTable("TCefListValueRef_SetNull", 0), // function SetNull
			/* 22 */ imports.NewTable("TCefListValueRef_SetBool", 0), // function SetBool
			/* 23 */ imports.NewTable("TCefListValueRef_SetInt", 0), // function SetInt
			/* 24 */ imports.NewTable("TCefListValueRef_SetDouble", 0), // function SetDouble
			/* 25 */ imports.NewTable("TCefListValueRef_SetString", 0), // function SetString
			/* 26 */ imports.NewTable("TCefListValueRef_SetBinary", 0), // function SetBinary
			/* 27 */ imports.NewTable("TCefListValueRef_SetDictionary", 0), // function SetDictionary
			/* 28 */ imports.NewTable("TCefListValueRef_SetList", 0), // function SetList
			/* 29 */ imports.NewTable("TCefListValueRef_UnWrap", 0), // static function UnWrap
			/* 30 */ imports.NewTable("TCefListValueRef_New", 0), // static function New
		}
	})
	return cefListValueRefImport
}
