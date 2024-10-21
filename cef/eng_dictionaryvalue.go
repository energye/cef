//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICefDictionaryValue Parent: ICefBaseRefCounted
//
//	Interface representing a dictionary value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_dictionary_value_t))</a>
type ICefDictionaryValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. This object may become invalid if the underlying data is owned by another object (e.g. list or dictionary) and that other object is then modified or destroyed. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if this object is currently owned by another object.
	IsOwned() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying data. If true (1) modifications to this object will also affect |that| object and vice-versa.
	IsSame(that ICefDictionaryValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefDictionaryValue) bool // function
	// Copy
	//  Returns a writable copy of this object. If |exclude_NULL_children| is true (1) any NULL dictionaries or lists will be excluded from the copy.
	Copy(excludeEmptyChildren bool) ICefDictionaryValue // function
	// GetSize
	//  Returns the number of values.
	GetSize() NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// HasKey
	//  Returns true (1) if the current dictionary has a value for the given key.
	HasKey(key string) bool // function
	// GetKeys
	//  Reads all keys for this dictionary into the specified vector.
	GetKeys(keys IStrings) bool // function
	// Remove
	//  Removes the value at the specified key. Returns true (1) is the value was removed successfully.
	Remove(key string) bool // function
	// GetType
	//  Returns the value type for the specified key.
	GetType(key string) TCefValueType // function
	// GetValue
	//  Returns the value at the specified key. For simple types the returned value will copy existing data and modifications to the value will not modify this object. For complex types (binary, dictionary and list) the returned value will reference existing data and modifications to the value will modify this object.
	GetValue(key string) ICefValue // function
	// GetBool
	//  Returns the value at the specified key as type bool.
	GetBool(key string) bool // function
	// GetInt
	//  Returns the value at the specified key as type int.
	GetInt(key string) int32 // function
	// GetDouble
	//  Returns the value at the specified key as type double.
	GetDouble(key string) (resultFloat64 float64) // function
	// GetString
	//  Returns the value at the specified key as type string.
	GetString(key string) string // function
	// GetBinary
	//  Returns the value at the specified key as type binary. The returned value will reference existing data.
	GetBinary(key string) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified key as type dictionary. The returned value will reference existing data and modifications to the value will modify this object.
	GetDictionary(key string) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified key as type list. The returned value will reference existing data and modifications to the value will modify this object.
	GetList(key string) ICefListValue // function
	// SetValue
	//  Sets the value at the specified key. Returns true (1) if the value was set successfully. If |value| represents simple data then the underlying data will be copied and modifications to |value| will not modify this object. If |value| represents complex data (binary, dictionary or list) then the underlying data will be referenced and modifications to |value| will modify this object.
	SetValue(key string, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified key as type null. Returns true (1) if the value was set successfully.
	SetNull(key string) bool // function
	// SetBool
	//  Sets the value at the specified key as type bool. Returns true (1) if the value was set successfully.
	SetBool(key string, value bool) bool // function
	// SetInt
	//  Sets the value at the specified key as type int. Returns true (1) if the value was set successfully.
	SetInt(key string, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified key as type double. Returns true (1) if the value was set successfully.
	SetDouble(key string, value float64) bool // function
	// SetString
	//  Sets the value at the specified key as type string. Returns true (1) if the value was set successfully.
	SetString(key, value string) bool // function
	// SetBinary
	//  Sets the value at the specified key as type binary. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetBinary(key string, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified key as type dict. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetDictionary(key string, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified key as type list. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetList(key string, value ICefListValue) bool // function
}

// TCefDictionaryValue Parent: TCefBaseRefCounted
//
//	Interface representing a dictionary value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_dictionary_value_t))</a>
type TCefDictionaryValue struct {
	TCefBaseRefCounted
}

// DictionaryValueRef -> ICefDictionaryValue
var DictionaryValueRef dictionaryValue

// dictionaryValue TCefDictionaryValue Ref
type dictionaryValue uintptr

func (m *dictionaryValue) UnWrap(data uintptr) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	dictionaryValueImportAPI().SysCallN(30, uintptr(data), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *dictionaryValue) New() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	dictionaryValueImportAPI().SysCallN(19, uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) IsValid() bool {
	r1 := dictionaryValueImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsOwned() bool {
	r1 := dictionaryValueImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsReadOnly() bool {
	r1 := dictionaryValueImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsSame(that ICefDictionaryValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) IsEqual(that ICefDictionaryValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) Copy(excludeEmptyChildren bool) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	dictionaryValueImportAPI().SysCallN(1, m.Instance(), PascalBool(excludeEmptyChildren), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) GetSize() NativeUInt {
	r1 := dictionaryValueImportAPI().SysCallN(9, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefDictionaryValue) Clear() bool {
	r1 := dictionaryValueImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCefDictionaryValue) HasKey(key string) bool {
	r1 := dictionaryValueImportAPI().SysCallN(13, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetKeys(keys IStrings) bool {
	r1 := dictionaryValueImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(keys))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) Remove(key string) bool {
	r1 := dictionaryValueImportAPI().SysCallN(20, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetType(key string) TCefValueType {
	r1 := dictionaryValueImportAPI().SysCallN(11, m.Instance(), PascalStr(key))
	return TCefValueType(r1)
}

func (m *TCefDictionaryValue) GetValue(key string) ICefValue {
	var resultCefValue uintptr
	dictionaryValueImportAPI().SysCallN(12, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefDictionaryValue) GetBool(key string) bool {
	r1 := dictionaryValueImportAPI().SysCallN(3, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) GetInt(key string) int32 {
	r1 := dictionaryValueImportAPI().SysCallN(6, m.Instance(), PascalStr(key))
	return int32(r1)
}

func (m *TCefDictionaryValue) GetDouble(key string) (resultFloat64 float64) {
	dictionaryValueImportAPI().SysCallN(5, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefDictionaryValue) GetString(key string) string {
	r1 := dictionaryValueImportAPI().SysCallN(10, m.Instance(), PascalStr(key))
	return GoStr(r1)
}

func (m *TCefDictionaryValue) GetBinary(key string) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	dictionaryValueImportAPI().SysCallN(2, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefDictionaryValue) GetDictionary(key string) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	dictionaryValueImportAPI().SysCallN(4, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefDictionaryValue) GetList(key string) ICefListValue {
	var resultCefListValue uintptr
	dictionaryValueImportAPI().SysCallN(8, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefDictionaryValue) SetValue(key string, value ICefValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(29, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetNull(key string) bool {
	r1 := dictionaryValueImportAPI().SysCallN(27, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetBool(key string, value bool) bool {
	r1 := dictionaryValueImportAPI().SysCallN(22, m.Instance(), PascalStr(key), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetInt(key string, value int32) bool {
	r1 := dictionaryValueImportAPI().SysCallN(25, m.Instance(), PascalStr(key), uintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetDouble(key string, value float64) bool {
	r1 := dictionaryValueImportAPI().SysCallN(24, m.Instance(), PascalStr(key), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetString(key, value string) bool {
	r1 := dictionaryValueImportAPI().SysCallN(28, m.Instance(), PascalStr(key), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetBinary(key string, value ICefBinaryValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(21, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetDictionary(key string, value ICefDictionaryValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(23, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefDictionaryValue) SetList(key string, value ICefListValue) bool {
	r1 := dictionaryValueImportAPI().SysCallN(26, m.Instance(), PascalStr(key), GetObjectUintptr(value))
	return GoBool(r1)
}

var (
	dictionaryValueImport       *imports.Imports = nil
	dictionaryValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDictionaryValue_Clear", 0),
		/*1*/ imports.NewTable("CefDictionaryValue_Copy", 0),
		/*2*/ imports.NewTable("CefDictionaryValue_GetBinary", 0),
		/*3*/ imports.NewTable("CefDictionaryValue_GetBool", 0),
		/*4*/ imports.NewTable("CefDictionaryValue_GetDictionary", 0),
		/*5*/ imports.NewTable("CefDictionaryValue_GetDouble", 0),
		/*6*/ imports.NewTable("CefDictionaryValue_GetInt", 0),
		/*7*/ imports.NewTable("CefDictionaryValue_GetKeys", 0),
		/*8*/ imports.NewTable("CefDictionaryValue_GetList", 0),
		/*9*/ imports.NewTable("CefDictionaryValue_GetSize", 0),
		/*10*/ imports.NewTable("CefDictionaryValue_GetString", 0),
		/*11*/ imports.NewTable("CefDictionaryValue_GetType", 0),
		/*12*/ imports.NewTable("CefDictionaryValue_GetValue", 0),
		/*13*/ imports.NewTable("CefDictionaryValue_HasKey", 0),
		/*14*/ imports.NewTable("CefDictionaryValue_IsEqual", 0),
		/*15*/ imports.NewTable("CefDictionaryValue_IsOwned", 0),
		/*16*/ imports.NewTable("CefDictionaryValue_IsReadOnly", 0),
		/*17*/ imports.NewTable("CefDictionaryValue_IsSame", 0),
		/*18*/ imports.NewTable("CefDictionaryValue_IsValid", 0),
		/*19*/ imports.NewTable("CefDictionaryValue_New", 0),
		/*20*/ imports.NewTable("CefDictionaryValue_Remove", 0),
		/*21*/ imports.NewTable("CefDictionaryValue_SetBinary", 0),
		/*22*/ imports.NewTable("CefDictionaryValue_SetBool", 0),
		/*23*/ imports.NewTable("CefDictionaryValue_SetDictionary", 0),
		/*24*/ imports.NewTable("CefDictionaryValue_SetDouble", 0),
		/*25*/ imports.NewTable("CefDictionaryValue_SetInt", 0),
		/*26*/ imports.NewTable("CefDictionaryValue_SetList", 0),
		/*27*/ imports.NewTable("CefDictionaryValue_SetNull", 0),
		/*28*/ imports.NewTable("CefDictionaryValue_SetString", 0),
		/*29*/ imports.NewTable("CefDictionaryValue_SetValue", 0),
		/*30*/ imports.NewTable("CefDictionaryValue_UnWrap", 0),
	}
)

func dictionaryValueImportAPI() *imports.Imports {
	if dictionaryValueImport == nil {
		dictionaryValueImport = NewDefaultImports()
		dictionaryValueImport.SetImportTable(dictionaryValueImportTables)
		dictionaryValueImportTables = nil
	}
	return dictionaryValueImport
}
