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
	IsValid() bool                                // function
	IsOwned() bool                                // function
	IsReadOnly() bool                             // function
	IsSame(that ICefValue) bool                   // function
	IsEqual(that ICefValue) bool                  // function
	Copy() ICefValue                              // function
	GetType() cefTypes.TCefValueType              // function
	GetBool() bool                                // function
	GetInt() int32                                // function
	GetDouble() float64                           // function
	GetString() string                            // function
	GetBinary() ICefBinaryValue                   // function
	GetDictionary() ICefDictionaryValue           // function
	GetList() ICefListValue                       // function
	SetNull() bool                                // function
	SetBool(value bool) bool                      // function
	SetInt(value int32) bool                      // function
	SetDouble(value float64) bool                 // function
	SetString(value string) bool                  // function
	SetBinary(value ICefBinaryValue) bool         // function
	SetDictionary(value ICefDictionaryValue) bool // function
	SetList(value ICefListValue) bool             // function
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
