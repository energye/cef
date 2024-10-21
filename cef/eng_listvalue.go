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

// ICefListValue Parent: ICefBaseRefCounted
//
//	Interface representing a list value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_list_value_t))</a>
type ICefListValue interface {
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
	IsSame(that ICefListValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefListValue) bool // function
	// Copy
	//  Returns a writable copy of this object.
	Copy() ICefListValue // function
	// SetSize
	//  Sets the number of values. If the number of values is expanded all new value slots will default to type null. Returns true (1) on success.
	SetSize(size NativeUInt) bool // function
	// GetSize
	//  Returns the number of values.
	GetSize() NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// Remove
	//  Removes the value at the specified index.
	Remove(index NativeUInt) bool // function
	// GetType
	//  Returns the value type at the specified index.
	GetType(index NativeUInt) TCefValueType // function
	// GetValue
	//  Returns the value at the specified index. For simple types the returned value will copy existing data and modifications to the value will not modify this object. For complex types (binary, dictionary and list) the returned value will reference existing data and modifications to the value will modify this object.
	GetValue(index NativeUInt) ICefValue // function
	// GetBool
	//  Returns the value at the specified index as type bool.
	GetBool(index NativeUInt) bool // function
	// GetInt
	//  Returns the value at the specified index as type int.
	GetInt(index NativeUInt) int32 // function
	// GetDouble
	//  Returns the value at the specified index as type double.
	GetDouble(index NativeUInt) (resultFloat64 float64) // function
	// GetString
	//  Returns the value at the specified index as type string.
	GetString(index NativeUInt) string // function
	// GetBinary
	//  Returns the value at the specified index as type binary. The returned value will reference existing data.
	GetBinary(index NativeUInt) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified index as type dictionary. The returned value will reference existing data and modifications to the value will modify this object.
	GetDictionary(index NativeUInt) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified index as type list. The returned value will reference existing data and modifications to the value will modify this object.
	GetList(index NativeUInt) ICefListValue // function
	// SetValue
	//  Sets the value at the specified index. Returns true (1) if the value was set successfully. If |value| represents simple data then the underlying data will be copied and modifications to |value| will not modify this object. If |value| represents complex data (binary, dictionary or list) then the underlying data will be referenced and modifications to |value| will modify this object.
	SetValue(index NativeUInt, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified index as type null. Returns true (1) if the value was set successfully.
	SetNull(index NativeUInt) bool // function
	// SetBool
	//  Sets the value at the specified index as type bool. Returns true (1) if the value was set successfully.
	SetBool(index NativeUInt, value bool) bool // function
	// SetInt
	//  Sets the value at the specified index as type int. Returns true (1) if the value was set successfully.
	SetInt(index NativeUInt, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified index as type double. Returns true (1) if the value was set successfully.
	SetDouble(index NativeUInt, value float64) bool // function
	// SetString
	//  Sets the value at the specified index as type string. Returns true (1) if the value was set successfully.
	SetString(index NativeUInt, value string) bool // function
	// SetBinary
	//  Sets the value at the specified index as type binary. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetBinary(index NativeUInt, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified index as type dict. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetDictionary(index NativeUInt, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified index as type list. Returns true (1) if the value was set successfully. If |value| is currently owned by another object then the value will be copied and the |value| reference will not change. Otherwise, ownership will be transferred to this object and the |value| reference will be invalidated.
	SetList(index NativeUInt, value ICefListValue) bool // function
}

// TCefListValue Parent: TCefBaseRefCounted
//
//	Interface representing a list value. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_list_value_t))</a>
type TCefListValue struct {
	TCefBaseRefCounted
}

// ListValueRef -> ICefListValue
var ListValueRef listValue

// listValue TCefListValue Ref
type listValue uintptr

func (m *listValue) UnWrap(data uintptr) ICefListValue {
	var resultCefListValue uintptr
	listValueImportAPI().SysCallN(29, uintptr(data), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *listValue) New() ICefListValue {
	var resultCefListValue uintptr
	listValueImportAPI().SysCallN(17, uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) IsValid() bool {
	r1 := listValueImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsOwned() bool {
	r1 := listValueImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsReadOnly() bool {
	r1 := listValueImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) IsSame(that ICefListValue) bool {
	r1 := listValueImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefListValue) IsEqual(that ICefListValue) bool {
	r1 := listValueImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefListValue) Copy() ICefListValue {
	var resultCefListValue uintptr
	listValueImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) SetSize(size NativeUInt) bool {
	r1 := listValueImportAPI().SysCallN(26, m.Instance(), uintptr(size))
	return GoBool(r1)
}

func (m *TCefListValue) GetSize() NativeUInt {
	r1 := listValueImportAPI().SysCallN(8, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefListValue) Clear() bool {
	r1 := listValueImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCefListValue) Remove(index NativeUInt) bool {
	r1 := listValueImportAPI().SysCallN(18, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) GetType(index NativeUInt) TCefValueType {
	r1 := listValueImportAPI().SysCallN(10, m.Instance(), uintptr(index))
	return TCefValueType(r1)
}

func (m *TCefListValue) GetValue(index NativeUInt) ICefValue {
	var resultCefValue uintptr
	listValueImportAPI().SysCallN(11, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefListValue) GetBool(index NativeUInt) bool {
	r1 := listValueImportAPI().SysCallN(3, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) GetInt(index NativeUInt) int32 {
	r1 := listValueImportAPI().SysCallN(6, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefListValue) GetDouble(index NativeUInt) (resultFloat64 float64) {
	listValueImportAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefListValue) GetString(index NativeUInt) string {
	value := NewTString()
	defer value.Free()
	listValueImportAPI().SysCallN(9, m.Instance(), uintptr(index), value.Instance())
	return value.Value()
}

func (m *TCefListValue) GetBinary(index NativeUInt) ICefBinaryValue {
	var resultCefBinaryValue uintptr
	listValueImportAPI().SysCallN(2, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefListValue) GetDictionary(index NativeUInt) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	listValueImportAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefListValue) GetList(index NativeUInt) ICefListValue {
	var resultCefListValue uintptr
	listValueImportAPI().SysCallN(7, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefListValue) SetValue(index NativeUInt, value ICefValue) bool {
	r1 := listValueImportAPI().SysCallN(28, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetNull(index NativeUInt) bool {
	r1 := listValueImportAPI().SysCallN(25, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefListValue) SetBool(index NativeUInt, value bool) bool {
	r1 := listValueImportAPI().SysCallN(20, m.Instance(), uintptr(index), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetInt(index NativeUInt, value int32) bool {
	r1 := listValueImportAPI().SysCallN(23, m.Instance(), uintptr(index), uintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetDouble(index NativeUInt, value float64) bool {
	r1 := listValueImportAPI().SysCallN(22, m.Instance(), uintptr(index), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefListValue) SetString(index NativeUInt, value string) bool {
	r1 := listValueImportAPI().SysCallN(27, m.Instance(), uintptr(index), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetBinary(index NativeUInt, value ICefBinaryValue) bool {
	r1 := listValueImportAPI().SysCallN(19, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetDictionary(index NativeUInt, value ICefDictionaryValue) bool {
	r1 := listValueImportAPI().SysCallN(21, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefListValue) SetList(index NativeUInt, value ICefListValue) bool {
	r1 := listValueImportAPI().SysCallN(24, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

var (
	listValueImport       *imports.Imports = nil
	listValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefListValue_Clear", 0),
		/*1*/ imports.NewTable("CefListValue_Copy", 0),
		/*2*/ imports.NewTable("CefListValue_GetBinary", 0),
		/*3*/ imports.NewTable("CefListValue_GetBool", 0),
		/*4*/ imports.NewTable("CefListValue_GetDictionary", 0),
		/*5*/ imports.NewTable("CefListValue_GetDouble", 0),
		/*6*/ imports.NewTable("CefListValue_GetInt", 0),
		/*7*/ imports.NewTable("CefListValue_GetList", 0),
		/*8*/ imports.NewTable("CefListValue_GetSize", 0),
		/*9*/ imports.NewTable("CefListValue_GetString", 0),
		/*10*/ imports.NewTable("CefListValue_GetType", 0),
		/*11*/ imports.NewTable("CefListValue_GetValue", 0),
		/*12*/ imports.NewTable("CefListValue_IsEqual", 0),
		/*13*/ imports.NewTable("CefListValue_IsOwned", 0),
		/*14*/ imports.NewTable("CefListValue_IsReadOnly", 0),
		/*15*/ imports.NewTable("CefListValue_IsSame", 0),
		/*16*/ imports.NewTable("CefListValue_IsValid", 0),
		/*17*/ imports.NewTable("CefListValue_New", 0),
		/*18*/ imports.NewTable("CefListValue_Remove", 0),
		/*19*/ imports.NewTable("CefListValue_SetBinary", 0),
		/*20*/ imports.NewTable("CefListValue_SetBool", 0),
		/*21*/ imports.NewTable("CefListValue_SetDictionary", 0),
		/*22*/ imports.NewTable("CefListValue_SetDouble", 0),
		/*23*/ imports.NewTable("CefListValue_SetInt", 0),
		/*24*/ imports.NewTable("CefListValue_SetList", 0),
		/*25*/ imports.NewTable("CefListValue_SetNull", 0),
		/*26*/ imports.NewTable("CefListValue_SetSize", 0),
		/*27*/ imports.NewTable("CefListValue_SetString", 0),
		/*28*/ imports.NewTable("CefListValue_SetValue", 0),
		/*29*/ imports.NewTable("CefListValue_UnWrap", 0),
	}
)

func listValueImportAPI() *imports.Imports {
	if listValueImport == nil {
		listValueImport = NewDefaultImports()
		listValueImport.SetImportTable(listValueImportTables)
		listValueImportTables = nil
	}
	return listValueImport
}
