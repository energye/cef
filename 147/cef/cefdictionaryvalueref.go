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

	cefTypes "github.com/energye/cef/147/types"
)

// ICefDictionaryValue Parent: ICefBaseRefCounted
type ICefDictionaryValue interface {
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
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may
	//  expose read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying
	//  data. If true (1) modifications to this object will also affect |that|
	//  object and vice-versa.
	IsSame(that ICefDictionaryValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent
	//  underlying value but are not necessarily the same object.
	IsEqual(that ICefDictionaryValue) bool // function
	// Copy
	//  Returns a writable copy of this object. If |exclude_NULL_children| is true
	//  (1) any NULL dictionaries or lists will be excluded from the copy.
	Copy(excludeEmptyChildren bool) ICefDictionaryValue // function
	// GetSize
	//  Returns the number of values.
	GetSize() cefTypes.NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// HasKey
	//  Returns true (1) if the current dictionary has a value for the given key.
	HasKey(key string) bool // function
	// GetKeys
	//  Reads all keys for this dictionary into the specified vector.
	GetKeys(keys lcl.IStrings) bool // function
	// Remove
	//  Removes the value at the specified key. Returns true (1) is the value was
	//  removed successfully.
	Remove(key string) bool // function
	// GetType
	//  Returns the value type for the specified key.
	GetType(key string) cefTypes.TCefValueType // function
	// GetValue
	//  Returns the value at the specified key. For simple types the returned
	//  value will copy existing data and modifications to the value will not
	//  modify this object. For complex types (binary, dictionary and list) the
	//  returned value will reference existing data and modifications to the value
	//  will modify this object.
	GetValue(key string) ICefValue // function
	// GetBool
	//  Returns the value at the specified key as type bool.
	GetBool(key string) bool // function
	// GetInt
	//  Returns the value at the specified key as type int.
	GetInt(key string) int32 // function
	// GetDouble
	//  Returns the value at the specified key as type double.
	GetDouble(key string) float64 // function
	// GetString
	//  Returns the value at the specified key as type string.
	GetString(key string) string // function
	// GetBinary
	//  Returns the value at the specified key as type binary. The returned value
	//  will reference existing data.
	GetBinary(key string) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified key as type dictionary. The returned
	//  value will reference existing data and modifications to the value will
	//  modify this object.
	GetDictionary(key string) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified key as type list. The returned value
	//  will reference existing data and modifications to the value will modify
	//  this object.
	GetList(key string) ICefListValue // function
	// SetValue
	//  Sets the value at the specified key. Returns true (1) if the value was set
	//  successfully. If |value| represents simple data then the underlying data
	//  will be copied and modifications to |value| will not modify this object.
	//  If |value| represents complex data (binary, dictionary or list) then the
	//  underlying data will be referenced and modifications to |value| will
	//  modify this object.
	SetValue(key string, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified key as type null. Returns true (1) if the
	//  value was set successfully.
	SetNull(key string) bool // function
	// SetBool
	//  Sets the value at the specified key as type bool. Returns true (1) if the
	//  value was set successfully.
	SetBool(key string, value bool) bool // function
	// SetInt
	//  Sets the value at the specified key as type int. Returns true (1) if the
	//  value was set successfully.
	SetInt(key string, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified key as type double. Returns true (1) if
	//  the value was set successfully.
	SetDouble(key string, value float64) bool // function
	// SetString
	//  Sets the value at the specified key as type string. Returns true (1) if
	//  the value was set successfully.
	SetString(key string, value string) bool // function
	// SetBinary
	//  Sets the value at the specified key as type binary. Returns true (1) if
	//  the value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetBinary(key string, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified key as type dict. Returns true (1) if the
	//  value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetDictionary(key string, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified key as type list. Returns true (1) if the
	//  value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetList(key string, value ICefListValue) bool // function
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
