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

// ICefValue Parent: ICefBaseRefCounted
//
//	Interface that wraps other data value types. Complex types (binary, dictionary and list) will be referenced but not owned by this object. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_value_t))</a>
type ICefValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying data is valid. This will always be true (1) for simple types. For complex types (binary, dictionary and list) the underlying data may become invalid if owned by another object (e.g. list or dictionary) and that other object is then modified or destroyed. This value object can be re-used by calling Set*() even if the underlying data is invalid.
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if the underlying data is owned by another object.
	IsOwned() bool // function
	// IsReadOnly
	//  Returns true (1) if the underlying data is read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying data. If true (1) modifications to this object will also affect |that| object and vice-versa.
	IsSame(that ICefValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent underlying value but are not necessarily the same object.
	IsEqual(that ICefValue) bool // function
	// Copy
	//  Returns a copy of this object. The underlying data will also be copied.
	Copy() ICefValue // function
	// GetType
	//  Returns the underlying value type.
	GetType() TCefValueType // function
	// GetBool
	//  Returns the underlying value as type bool.
	GetBool() bool // function
	// GetInt
	//  Returns the underlying value as type int.
	GetInt() int32 // function
	// GetDouble
	//  Returns the underlying value as type double.
	GetDouble() (resultFloat64 float64) // function
	// GetString
	//  Returns the underlying value as type string.
	GetString() string // function
	// GetBinary
	//  Returns the underlying value as type binary. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_binary().
	GetBinary() ICefBinaryValue // function
	// GetDictionary
	//  Returns the underlying value as type dictionary. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_dictionary().
	GetDictionary() ICefDictionaryValue // function
	// GetList
	//  Returns the underlying value as type list. The returned reference may become invalid if the value is owned by another object or if ownership is transferred to another object in the future. To maintain a reference to the value after assigning ownership to a dictionary or list pass this object to the set_value() function instead of passing the returned reference to set_list().
	GetList() ICefListValue // function
	// SetNull
	//  Sets the underlying value as type null. Returns true (1) if the value was set successfully.
	SetNull() bool // function
	// SetBool
	//  Sets the underlying value as type bool. Returns true (1) if the value was set successfully.
	SetBool(value bool) bool // function
	// SetInt
	//  Sets the underlying value as type int. Returns true (1) if the value was set successfully.
	SetInt(value int32) bool // function
	// SetDouble
	//  Sets the underlying value as type double. Returns true (1) if the value was set successfully.
	SetDouble(value float64) bool // function
	// SetString
	//  Sets the underlying value as type string. Returns true (1) if the value was set successfully.
	SetString(value string) bool // function
	// SetBinary
	//  Sets the underlying value as type binary. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetBinary(value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the underlying value as type dict. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetDictionary(value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the underlying value as type list. Returns true (1) if the value was set successfully. This object keeps a reference to |value| and ownership of the underlying data remains unchanged.
	SetList(value ICefListValue) bool // function
}

// TCefValue Parent: TCefBaseRefCounted
//
//	Interface that wraps other data value types. Complex types (binary, dictionary and list) will be referenced but not owned by this object. Can be used on any process and thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_values_capi.h">CEF source file: /include/capi/cef_values_capi.h (cef_value_t))</a>
type TCefValue struct {
	TCefBaseRefCounted
}

// ValueRef -> ICefValue
var ValueRef value

// value TCefValue Ref
type value uintptr

func (m *value) UnWrap(data uintptr) ICefValue {
	var resultCefValue uintptr
	valueImportAPI().SysCallN(23, uintptr(data), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *value) New() ICefValue {
	var resultCefValue uintptr
	valueImportAPI().SysCallN(14, uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefValue) IsValid() bool {
	r1 := valueImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsOwned() bool {
	r1 := valueImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsReadOnly() bool {
	r1 := valueImportAPI().SysCallN(11, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) IsSame(that ICefValue) bool {
	r1 := valueImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefValue) IsEqual(that ICefValue) bool {
	r1 := valueImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefValue) Copy() ICefValue {
	var resultCefValue uintptr
	valueImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefValue) GetType() TCefValueType {
	r1 := valueImportAPI().SysCallN(8, m.Instance())
	return TCefValueType(r1)
}

func (m *TCefValue) GetBool() bool {
	r1 := valueImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) GetInt() int32 {
	r1 := valueImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TCefValue) GetDouble() (resultFloat64 float64) {
	valueImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefValue) GetString() string {
	r1 := valueImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefValue) GetBinary() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	valueImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefValue) GetDictionary() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	valueImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefValue) GetList() ICefListValue {
	var resultCefListValue uintptr
	valueImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefValue) SetNull() bool {
	r1 := valueImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TCefValue) SetBool(value bool) bool {
	r1 := valueImportAPI().SysCallN(16, m.Instance(), PascalBool(value))
	return GoBool(r1)
}

func (m *TCefValue) SetInt(value int32) bool {
	r1 := valueImportAPI().SysCallN(19, m.Instance(), uintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetDouble(value float64) bool {
	r1 := valueImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(&value)))
	return GoBool(r1)
}

func (m *TCefValue) SetString(value string) bool {
	r1 := valueImportAPI().SysCallN(22, m.Instance(), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetBinary(value ICefBinaryValue) bool {
	r1 := valueImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetDictionary(value ICefDictionaryValue) bool {
	r1 := valueImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefValue) SetList(value ICefListValue) bool {
	r1 := valueImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(value))
	return GoBool(r1)
}

var (
	valueImport       *imports.Imports = nil
	valueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefValue_Copy", 0),
		/*1*/ imports.NewTable("CefValue_GetBinary", 0),
		/*2*/ imports.NewTable("CefValue_GetBool", 0),
		/*3*/ imports.NewTable("CefValue_GetDictionary", 0),
		/*4*/ imports.NewTable("CefValue_GetDouble", 0),
		/*5*/ imports.NewTable("CefValue_GetInt", 0),
		/*6*/ imports.NewTable("CefValue_GetList", 0),
		/*7*/ imports.NewTable("CefValue_GetString", 0),
		/*8*/ imports.NewTable("CefValue_GetType", 0),
		/*9*/ imports.NewTable("CefValue_IsEqual", 0),
		/*10*/ imports.NewTable("CefValue_IsOwned", 0),
		/*11*/ imports.NewTable("CefValue_IsReadOnly", 0),
		/*12*/ imports.NewTable("CefValue_IsSame", 0),
		/*13*/ imports.NewTable("CefValue_IsValid", 0),
		/*14*/ imports.NewTable("CefValue_New", 0),
		/*15*/ imports.NewTable("CefValue_SetBinary", 0),
		/*16*/ imports.NewTable("CefValue_SetBool", 0),
		/*17*/ imports.NewTable("CefValue_SetDictionary", 0),
		/*18*/ imports.NewTable("CefValue_SetDouble", 0),
		/*19*/ imports.NewTable("CefValue_SetInt", 0),
		/*20*/ imports.NewTable("CefValue_SetList", 0),
		/*21*/ imports.NewTable("CefValue_SetNull", 0),
		/*22*/ imports.NewTable("CefValue_SetString", 0),
		/*23*/ imports.NewTable("CefValue_UnWrap", 0),
	}
)

func valueImportAPI() *imports.Imports {
	if valueImport == nil {
		valueImport = NewDefaultImports()
		valueImport.SetImportTable(valueImportTables)
		valueImportTables = nil
	}
	return valueImport
}
