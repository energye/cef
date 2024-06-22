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
type ICefListValue interface {
	ICefBaseRefCounted
	IsValid() bool                                                  // function
	IsOwned() bool                                                  // function
	IsReadOnly() bool                                               // function
	IsSame(that ICefListValue) bool                                 // function
	IsEqual(that ICefListValue) bool                                // function
	Copy() ICefListValue                                            // function
	SetSize(size NativeUInt) bool                                   // function
	GetSize() NativeUInt                                            // function
	Clear() bool                                                    // function
	Remove(index NativeUInt) bool                                   // function
	GetType(index NativeUInt) TCefValueType                         // function
	GetValue(index NativeUInt) ICefValue                            // function
	GetBool(index NativeUInt) bool                                  // function
	GetInt(index NativeUInt) int32                                  // function
	GetDouble(index NativeUInt) (resultFloat64 float64)             // function
	GetString(index NativeUInt) string                              // function
	GetBinary(index NativeUInt) ICefBinaryValue                     // function
	GetDictionary(index NativeUInt) ICefDictionaryValue             // function
	GetList(index NativeUInt) ICefListValue                         // function
	SetValue(index NativeUInt, value ICefValue) bool                // function
	SetNull(index NativeUInt) bool                                  // function
	SetBool(index NativeUInt, value bool) bool                      // function
	SetInt(index NativeUInt, value int32) bool                      // function
	SetDouble(index NativeUInt, value float64) bool                 // function
	SetString(index NativeUInt, value string) bool                  // function
	SetBinary(index NativeUInt, value ICefBinaryValue) bool         // function
	SetDictionary(index NativeUInt, value ICefDictionaryValue) bool // function
	SetList(index NativeUInt, value ICefListValue) bool             // function
}

// TCefListValue Parent: TCefBaseRefCounted
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
	r1 := listValueImportAPI().SysCallN(9, m.Instance(), uintptr(index))
	return GoStr(r1)
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
