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
type ICefValue interface {
	ICefBaseRefCounted
	IsValid() bool                                // function
	IsOwned() bool                                // function
	IsReadOnly() bool                             // function
	IsSame(that ICefValue) bool                   // function
	IsEqual(that ICefValue) bool                  // function
	Copy() ICefValue                              // function
	GetType() TCefValueType                       // function
	GetBool() bool                                // function
	GetInt() int32                                // function
	GetDouble() (resultFloat64 float64)           // function
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

// TCefValue Parent: TCefBaseRefCounted
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
