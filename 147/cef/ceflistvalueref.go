//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/147/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefListValue Parent: ICefBaseRefCounted
type ICefListValue interface {
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
	IsSame(that ICefListValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent
	//  underlying value but are not necessarily the same object.
	IsEqual(that ICefListValue) bool // function
	// Copy
	//  Returns a writable copy of this object.
	Copy() ICefListValue // function
	// SetSize
	//  Sets the number of values. If the number of values is expanded all new
	//  value slots will default to type null. Returns true (1) on success.
	SetSize(size cefTypes.NativeUInt) bool // function
	// GetSize
	//  Returns the number of values.
	GetSize() cefTypes.NativeUInt // function
	// Clear
	//  Removes all values. Returns true (1) on success.
	Clear() bool // function
	// Remove
	//  Removes the value at the specified index.
	Remove(index cefTypes.NativeUInt) bool // function
	// GetType
	//  Returns the value type at the specified index.
	GetType(index cefTypes.NativeUInt) cefTypes.TCefValueType // function
	// GetValue
	//  Returns the value at the specified index. For simple types the returned
	//  value will copy existing data and modifications to the value will not
	//  modify this object. For complex types (binary, dictionary and list) the
	//  returned value will reference existing data and modifications to the value
	//  will modify this object.
	GetValue(index cefTypes.NativeUInt) ICefValue // function
	// GetBool
	//  Returns the value at the specified index as type bool.
	GetBool(index cefTypes.NativeUInt) bool // function
	// GetInt
	//  Returns the value at the specified index as type int.
	GetInt(index cefTypes.NativeUInt) int32 // function
	// GetDouble
	//  Returns the value at the specified index as type double.
	GetDouble(index cefTypes.NativeUInt) float64 // function
	// GetString
	//  Returns the value at the specified index as type string.
	GetString(index cefTypes.NativeUInt) string // function
	// GetBinary
	//  Returns the value at the specified index as type binary. The returned
	//  value will reference existing data.
	GetBinary(index cefTypes.NativeUInt) ICefBinaryValue // function
	// GetDictionary
	//  Returns the value at the specified index as type dictionary. The returned
	//  value will reference existing data and modifications to the value will
	//  modify this object.
	GetDictionary(index cefTypes.NativeUInt) ICefDictionaryValue // function
	// GetList
	//  Returns the value at the specified index as type list. The returned value
	//  will reference existing data and modifications to the value will modify
	//  this object.
	GetList(index cefTypes.NativeUInt) ICefListValue // function
	// SetValue
	//  Sets the value at the specified index. Returns true (1) if the value was
	//  set successfully. If |value| represents simple data then the underlying
	//  data will be copied and modifications to |value| will not modify this
	//  object. If |value| represents complex data (binary, dictionary or list)
	//  then the underlying data will be referenced and modifications to |value|
	//  will modify this object.
	SetValue(index cefTypes.NativeUInt, value ICefValue) bool // function
	// SetNull
	//  Sets the value at the specified index as type null. Returns true (1) if
	//  the value was set successfully.
	SetNull(index cefTypes.NativeUInt) bool // function
	// SetBool
	//  Sets the value at the specified index as type bool. Returns true (1) if
	//  the value was set successfully.
	SetBool(index cefTypes.NativeUInt, value bool) bool // function
	// SetInt
	//  Sets the value at the specified index as type int. Returns true (1) if the
	//  value was set successfully.
	SetInt(index cefTypes.NativeUInt, value int32) bool // function
	// SetDouble
	//  Sets the value at the specified index as type double. Returns true (1) if
	//  the value was set successfully.
	SetDouble(index cefTypes.NativeUInt, value float64) bool // function
	// SetString
	//  Sets the value at the specified index as type string. Returns true (1) if
	//  the value was set successfully.
	SetString(index cefTypes.NativeUInt, value string) bool // function
	// SetBinary
	//  Sets the value at the specified index as type binary. Returns true (1) if
	//  the value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetBinary(index cefTypes.NativeUInt, value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the value at the specified index as type dict. Returns true (1) if
	//  the value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetDictionary(index cefTypes.NativeUInt, value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the value at the specified index as type list. Returns true (1) if
	//  the value was set successfully. If |value| is currently owned by another
	//  object then the value will be copied and the |value| reference will not
	//  change. Otherwise, ownership will be transferred to this object and the
	//  |value| reference will be invalidated.
	SetList(index cefTypes.NativeUInt, value ICefListValue) bool // function
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
