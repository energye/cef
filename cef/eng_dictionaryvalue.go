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
type ICefDictionaryValue interface {
	ICefBaseRefCounted
	IsValid() bool                                            // function
	IsOwned() bool                                            // function
	IsReadOnly() bool                                         // function
	IsSame(that ICefDictionaryValue) bool                     // function
	IsEqual(that ICefDictionaryValue) bool                    // function
	Copy(excludeEmptyChildren bool) ICefDictionaryValue       // function
	GetSize() NativeUInt                                      // function
	Clear() bool                                              // function
	HasKey(key string) bool                                   // function
	GetKeys(keys IStrings) bool                               // function
	Remove(key string) bool                                   // function
	GetType(key string) TCefValueType                         // function
	GetValue(key string) ICefValue                            // function
	GetBool(key string) bool                                  // function
	GetInt(key string) int32                                  // function
	GetDouble(key string) (resultFloat64 float64)             // function
	GetString(key string) string                              // function
	GetBinary(key string) ICefBinaryValue                     // function
	GetDictionary(key string) ICefDictionaryValue             // function
	GetList(key string) ICefListValue                         // function
	SetValue(key string, value ICefValue) bool                // function
	SetNull(key string) bool                                  // function
	SetBool(key string, value bool) bool                      // function
	SetInt(key string, value int32) bool                      // function
	SetDouble(key string, value float64) bool                 // function
	SetString(key, value string) bool                         // function
	SetBinary(key string, value ICefBinaryValue) bool         // function
	SetDictionary(key string, value ICefDictionaryValue) bool // function
	SetList(key string, value ICefListValue) bool             // function
}

// TCefDictionaryValue Parent: TCefBaseRefCounted
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
