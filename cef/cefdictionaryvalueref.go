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

// ICefDictionaryValue Parent: ICefBaseRefCounted
type ICefDictionaryValue interface {
	ICefBaseRefCounted
	IsValid() bool                                            // function
	IsOwned() bool                                            // function
	IsReadOnly() bool                                         // function
	IsSame(that ICefDictionaryValue) bool                     // function
	IsEqual(that ICefDictionaryValue) bool                    // function
	Copy(excludeEmptyChildren bool) ICefDictionaryValue       // function
	GetSize() cefTypes.NativeUInt                             // function
	Clear() bool                                              // function
	HasKey(key string) bool                                   // function
	GetKeys(keys lcl.IStrings) bool                           // function
	Remove(key string) bool                                   // function
	GetType(key string) cefTypes.TCefValueType                // function
	GetValue(key string) ICefValue                            // function
	GetBool(key string) bool                                  // function
	GetInt(key string) int32                                  // function
	GetDouble(key string) float64                             // function
	GetString(key string) string                              // function
	GetBinary(key string) ICefBinaryValue                     // function
	GetDictionary(key string) ICefDictionaryValue             // function
	GetList(key string) ICefListValue                         // function
	SetValue(key string, value ICefValue) bool                // function
	SetNull(key string) bool                                  // function
	SetBool(key string, value bool) bool                      // function
	SetInt(key string, value int32) bool                      // function
	SetDouble(key string, value float64) bool                 // function
	SetString(key string, value string) bool                  // function
	SetBinary(key string, value ICefBinaryValue) bool         // function
	SetDictionary(key string, value ICefDictionaryValue) bool // function
	SetList(key string, value ICefListValue) bool             // function
}

// ICefDictionaryValueRef Parent: ICefDictionaryValue ICefBaseRefCountedRef
type ICefDictionaryValueRef interface {
	ICefDictionaryValue
	ICefBaseRefCountedRef
	AsIntfDictionaryValue() uintptr
}

type TCefDictionaryValueRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDictionaryValueRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) IsSame(that ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) IsEqual(that ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) Copy(excludeEmptyChildren bool) (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(6, m.Instance(), api.PasBool(excludeEmptyChildren), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefDictionaryValueRef) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefDictionaryValueRefAPI().SysCallN(7, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefDictionaryValueRef) Clear() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) HasKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(9, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) GetKeys(keys lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(keys))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) Remove(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(11, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) GetType(key string) cefTypes.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r := cefDictionaryValueRefAPI().SysCallN(12, m.Instance(), api.PasStr(key))
	return cefTypes.TCefValueType(r)
}

func (m *TCefDictionaryValueRef) GetValue(key string) (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(13, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TCefDictionaryValueRef) GetBool(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(14, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) GetInt(key string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDictionaryValueRefAPI().SysCallN(15, m.Instance(), api.PasStr(key))
	return int32(r)
}

func (m *TCefDictionaryValueRef) GetDouble(key string) (result float64) {
	if !m.IsValid() {
		return
	}
	cefDictionaryValueRefAPI().SysCallN(16, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDictionaryValueRef) GetString(key string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDictionaryValueRefAPI().SysCallN(17, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDictionaryValueRef) GetBinary(key string) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(18, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefDictionaryValueRef) GetDictionary(key string) (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(19, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefDictionaryValueRef) GetList(key string) (result ICefListValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(20, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (m *TCefDictionaryValueRef) SetValue(key string, value ICefValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(21, m.Instance(), api.PasStr(key), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetNull(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(22, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetBool(key string, value bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(23, m.Instance(), api.PasStr(key), api.PasBool(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetInt(key string, value int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(24, m.Instance(), api.PasStr(key), uintptr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetDouble(key string, value float64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(25, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&value)))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetString(key string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(26, m.Instance(), api.PasStr(key), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetBinary(key string, value ICefBinaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(27, m.Instance(), api.PasStr(key), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetDictionary(key string, value ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(28, m.Instance(), api.PasStr(key), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) SetList(key string, value ICefListValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDictionaryValueRefAPI().SysCallN(29, m.Instance(), api.PasStr(key), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefDictionaryValueRef) AsIntfDictionaryValue() uintptr {
	return m.GetIntfPointer(0)
}

// DictionaryValueRef  is static instance
var DictionaryValueRef _DictionaryValueRefClass

// _DictionaryValueRefClass is class type defined by TCefDictionaryValueRef
type _DictionaryValueRefClass uintptr

func (_DictionaryValueRefClass) UnWrap(data uintptr) (result ICefDictionaryValue) {
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(30, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (_DictionaryValueRefClass) New() (result ICefDictionaryValue) {
	var resultPtr uintptr
	cefDictionaryValueRefAPI().SysCallN(31, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

// NewDictionaryValueRef class constructor
func NewDictionaryValueRef(data uintptr) ICefDictionaryValueRef {
	var dictionaryValuePtr uintptr // ICefDictionaryValue
	r := cefDictionaryValueRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&dictionaryValuePtr)))
	ret := AsCefDictionaryValueRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dictionaryValuePtr)
	}
	return ret
}

var (
	cefDictionaryValueRefOnce   base.Once
	cefDictionaryValueRefImport *imports.Imports = nil
)

func cefDictionaryValueRefAPI() *imports.Imports {
	cefDictionaryValueRefOnce.Do(func() {
		cefDictionaryValueRefImport = api.NewDefaultImports()
		cefDictionaryValueRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDictionaryValueRef_Create", 0), // constructor NewDictionaryValueRef
			/* 1 */ imports.NewTable("TCefDictionaryValueRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefDictionaryValueRef_isOwned", 0), // function IsOwned
			/* 3 */ imports.NewTable("TCefDictionaryValueRef_IsReadOnly", 0), // function IsReadOnly
			/* 4 */ imports.NewTable("TCefDictionaryValueRef_IsSame", 0), // function IsSame
			/* 5 */ imports.NewTable("TCefDictionaryValueRef_IsEqual", 0), // function IsEqual
			/* 6 */ imports.NewTable("TCefDictionaryValueRef_Copy", 0), // function Copy
			/* 7 */ imports.NewTable("TCefDictionaryValueRef_GetSize", 0), // function GetSize
			/* 8 */ imports.NewTable("TCefDictionaryValueRef_Clear", 0), // function Clear
			/* 9 */ imports.NewTable("TCefDictionaryValueRef_HasKey", 0), // function HasKey
			/* 10 */ imports.NewTable("TCefDictionaryValueRef_GetKeys", 0), // function GetKeys
			/* 11 */ imports.NewTable("TCefDictionaryValueRef_Remove", 0), // function Remove
			/* 12 */ imports.NewTable("TCefDictionaryValueRef_GetType", 0), // function GetType
			/* 13 */ imports.NewTable("TCefDictionaryValueRef_GetValue", 0), // function GetValue
			/* 14 */ imports.NewTable("TCefDictionaryValueRef_GetBool", 0), // function GetBool
			/* 15 */ imports.NewTable("TCefDictionaryValueRef_GetInt", 0), // function GetInt
			/* 16 */ imports.NewTable("TCefDictionaryValueRef_GetDouble", 0), // function GetDouble
			/* 17 */ imports.NewTable("TCefDictionaryValueRef_GetString", 0), // function GetString
			/* 18 */ imports.NewTable("TCefDictionaryValueRef_GetBinary", 0), // function GetBinary
			/* 19 */ imports.NewTable("TCefDictionaryValueRef_GetDictionary", 0), // function GetDictionary
			/* 20 */ imports.NewTable("TCefDictionaryValueRef_GetList", 0), // function GetList
			/* 21 */ imports.NewTable("TCefDictionaryValueRef_SetValue", 0), // function SetValue
			/* 22 */ imports.NewTable("TCefDictionaryValueRef_SetNull", 0), // function SetNull
			/* 23 */ imports.NewTable("TCefDictionaryValueRef_SetBool", 0), // function SetBool
			/* 24 */ imports.NewTable("TCefDictionaryValueRef_SetInt", 0), // function SetInt
			/* 25 */ imports.NewTable("TCefDictionaryValueRef_SetDouble", 0), // function SetDouble
			/* 26 */ imports.NewTable("TCefDictionaryValueRef_SetString", 0), // function SetString
			/* 27 */ imports.NewTable("TCefDictionaryValueRef_SetBinary", 0), // function SetBinary
			/* 28 */ imports.NewTable("TCefDictionaryValueRef_SetDictionary", 0), // function SetDictionary
			/* 29 */ imports.NewTable("TCefDictionaryValueRef_SetList", 0), // function SetList
			/* 30 */ imports.NewTable("TCefDictionaryValueRef_UnWrap", 0), // static function UnWrap
			/* 31 */ imports.NewTable("TCefDictionaryValueRef_New", 0), // static function New
		}
	})
	return cefDictionaryValueRefImport
}
