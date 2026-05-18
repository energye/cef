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

// ICefValue Parent: ICefBaseRefCounted
type ICefValue interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying data is valid. This will always be true
	//  (1) for simple types. For complex types (binary, dictionary and list) the
	//  underlying data may become invalid if owned by another object (e.g. list
	//  or dictionary) and that other object is then modified or destroyed. This
	//  value object can be re-used by calling Set*() even if the underlying data
	//  is invalid.
	IsValid() bool // function
	// IsOwned
	//  Returns true (1) if the underlying data is owned by another object.
	IsOwned() bool // function
	// IsReadOnly
	//  Returns true (1) if the underlying data is read-only. Some APIs may expose
	//  read-only objects.
	IsReadOnly() bool // function
	// IsSame
	//  Returns true (1) if this object and |that| object have the same underlying
	//  data. If true (1) modifications to this object will also affect |that|
	//  object and vice-versa.
	IsSame(that ICefValue) bool // function
	// IsEqual
	//  Returns true (1) if this object and |that| object have an equivalent
	//  underlying value but are not necessarily the same object.
	IsEqual(that ICefValue) bool // function
	// Copy
	//  Returns a copy of this object. The underlying data will also be copied.
	Copy() ICefValue // function
	// GetType
	//  Returns the underlying value type.
	GetType() cefTypes.TCefValueType // function
	// GetBool
	//  Returns the underlying value as type bool.
	GetBool() bool // function
	// GetInt
	//  Returns the underlying value as type int.
	GetInt() int32 // function
	// GetDouble
	//  Returns the underlying value as type double.
	GetDouble() float64 // function
	// GetString
	//  Returns the underlying value as type string.
	GetString() string // function
	// GetBinary
	//  Returns the underlying value as type binary. The returned reference may
	//  become invalid if the value is owned by another object or if ownership is
	//  transferred to another object in the future. To maintain a reference to
	//  the value after assigning ownership to a dictionary or list pass this
	//  object to the set_value() function instead of passing the returned
	//  reference to set_binary().
	GetBinary() ICefBinaryValue // function
	// GetDictionary
	//  Returns the underlying value as type dictionary. The returned reference
	//  may become invalid if the value is owned by another object or if ownership
	//  is transferred to another object in the future. To maintain a reference to
	//  the value after assigning ownership to a dictionary or list pass this
	//  object to the set_value() function instead of passing the returned
	//  reference to set_dictionary().
	GetDictionary() ICefDictionaryValue // function
	// GetList
	//  Returns the underlying value as type list. The returned reference may
	//  become invalid if the value is owned by another object or if ownership is
	//  transferred to another object in the future. To maintain a reference to
	//  the value after assigning ownership to a dictionary or list pass this
	//  object to the set_value() function instead of passing the returned
	//  reference to set_list().
	GetList() ICefListValue // function
	// SetNull
	//  Sets the underlying value as type null. Returns true (1) if the value was
	//  set successfully.
	SetNull() bool // function
	// SetBool
	//  Sets the underlying value as type bool. Returns true (1) if the value was
	//  set successfully.
	SetBool(value bool) bool // function
	// SetInt
	//  Sets the underlying value as type int. Returns true (1) if the value was
	//  set successfully.
	SetInt(value int32) bool // function
	// SetDouble
	//  Sets the underlying value as type double. Returns true (1) if the value
	//  was set successfully.
	SetDouble(value float64) bool // function
	// SetString
	//  Sets the underlying value as type string. Returns true (1) if the value
	//  was set successfully.
	SetString(value string) bool // function
	// SetBinary
	//  Sets the underlying value as type binary. Returns true (1) if the value
	//  was set successfully. This object keeps a reference to |value| and
	//  ownership of the underlying data remains unchanged.
	SetBinary(value ICefBinaryValue) bool // function
	// SetDictionary
	//  Sets the underlying value as type dict. Returns true (1) if the value was
	//  set successfully. This object keeps a reference to |value| and ownership
	//  of the underlying data remains unchanged.
	SetDictionary(value ICefDictionaryValue) bool // function
	// SetList
	//  Sets the underlying value as type list. Returns true (1) if the value was
	//  set successfully. This object keeps a reference to |value| and ownership
	//  of the underlying data remains unchanged.
	SetList(value ICefListValue) bool // function
}

// ICefValueRef Parent: ICefValue ICefBaseRefCountedRef
type ICefValueRef interface {
	ICefValue
	ICefBaseRefCountedRef
	AsIntfValue() uintptr
}

type TCefValueRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefValueRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefValueRef) IsOwned() bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefValueRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefValueRef) IsSame(that ICefValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefValueRef) IsEqual(that ICefValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefValueRef) Copy() (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TCefValueRef) GetType() cefTypes.TCefValueType {
	if !m.IsValid() {
		return 0
	}
	r := cefValueRefAPI().SysCallN(7, m.Instance())
	return cefTypes.TCefValueType(r)
}

func (m *TCefValueRef) GetBool() bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefValueRef) GetInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefValueRefAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TCefValueRef) GetDouble() (result float64) {
	if !m.IsValid() {
		return
	}
	cefValueRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefValueRef) GetString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefValueRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefValueRef) GetBinary() (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefValueRef) GetDictionary() (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefValueRef) GetList() (result ICefListValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (m *TCefValueRef) SetNull() bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TCefValueRef) SetBool(value bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(16, m.Instance(), api.PasBool(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetInt(value int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(17, m.Instance(), uintptr(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetDouble(value float64) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&value)))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetString(value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(19, m.Instance(), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetBinary(value ICefBinaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetDictionary(value ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) SetList(value ICefListValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefValueRefAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCefValueRef) AsIntfValue() uintptr {
	return m.GetIntfPointer(0)
}

// ValueRef  is static instance
var ValueRef _ValueRefClass

// _ValueRefClass is class type defined by TCefValueRef
type _ValueRefClass uintptr

func (_ValueRefClass) UnWrap(data uintptr) (result ICefValue) {
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(23, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (_ValueRefClass) New() (result ICefValue) {
	var resultPtr uintptr
	cefValueRefAPI().SysCallN(24, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

// NewValueRef class constructor
func NewValueRef(data uintptr) ICefValueRef {
	var valuePtr uintptr // ICefValue
	r := cefValueRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&valuePtr)))
	ret := AsCefValueRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, valuePtr)
	}
	return ret
}

var (
	cefValueRefOnce   base.Once
	cefValueRefImport *imports.Imports = nil
)

func cefValueRefAPI() *imports.Imports {
	cefValueRefOnce.Do(func() {
		cefValueRefImport = api.NewDefaultImports()
		cefValueRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefValueRef_Create", 0), // constructor NewValueRef
			/* 1 */ imports.NewTable("TCefValueRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefValueRef_IsOwned", 0), // function IsOwned
			/* 3 */ imports.NewTable("TCefValueRef_IsReadOnly", 0), // function IsReadOnly
			/* 4 */ imports.NewTable("TCefValueRef_IsSame", 0), // function IsSame
			/* 5 */ imports.NewTable("TCefValueRef_IsEqual", 0), // function IsEqual
			/* 6 */ imports.NewTable("TCefValueRef_Copy", 0), // function Copy
			/* 7 */ imports.NewTable("TCefValueRef_GetType", 0), // function GetType
			/* 8 */ imports.NewTable("TCefValueRef_GetBool", 0), // function GetBool
			/* 9 */ imports.NewTable("TCefValueRef_GetInt", 0), // function GetInt
			/* 10 */ imports.NewTable("TCefValueRef_GetDouble", 0), // function GetDouble
			/* 11 */ imports.NewTable("TCefValueRef_GetString", 0), // function GetString
			/* 12 */ imports.NewTable("TCefValueRef_GetBinary", 0), // function GetBinary
			/* 13 */ imports.NewTable("TCefValueRef_GetDictionary", 0), // function GetDictionary
			/* 14 */ imports.NewTable("TCefValueRef_GetList", 0), // function GetList
			/* 15 */ imports.NewTable("TCefValueRef_SetNull", 0), // function SetNull
			/* 16 */ imports.NewTable("TCefValueRef_SetBool", 0), // function SetBool
			/* 17 */ imports.NewTable("TCefValueRef_SetInt", 0), // function SetInt
			/* 18 */ imports.NewTable("TCefValueRef_SetDouble", 0), // function SetDouble
			/* 19 */ imports.NewTable("TCefValueRef_SetString", 0), // function SetString
			/* 20 */ imports.NewTable("TCefValueRef_SetBinary", 0), // function SetBinary
			/* 21 */ imports.NewTable("TCefValueRef_SetDictionary", 0), // function SetDictionary
			/* 22 */ imports.NewTable("TCefValueRef_SetList", 0), // function SetList
			/* 23 */ imports.NewTable("TCefValueRef_UnWrap", 0), // static function UnWrap
			/* 24 */ imports.NewTable("TCefValueRef_New", 0), // static function New
		}
	})
	return cefValueRefImport
}
