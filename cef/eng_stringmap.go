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

// ICefStringMap Parent: IObject
//
//	CEF string maps are a set of key/value string pairs.
type ICefStringMap interface {
	IObject
	GetHandle() TCefStringMapHandle // function
	// GetSize
	//  Return the number of elements in the string map.
	GetSize() NativeUInt // function
	// Find
	//  Return the value assigned to the specified key.
	Find(key string) string // function
	// GetKey
	//  Return the key at the specified zero-based string map index.
	GetKey(index NativeUInt) string // function
	// GetValue
	//  Return the value at the specified zero-based string map index.
	GetValue(index NativeUInt) string // function
	// Append
	//  Append a new key/value pair at the end of the string map. If the key exists,
	//  overwrite the existing value with a new value w/o changing the pair order.
	Append(key, value string) bool // function
	// Clear
	//  Clear the string map.
	Clear() // procedure
}

// TCefStringMap Parent: TObject
//
//	CEF string maps are a set of key/value string pairs.
type TCefStringMap struct {
	TObject
}

func NewCefStringMap() ICefStringMap {
	r1 := stringMapImportAPI().SysCallN(2)
	return AsCefStringMap(r1)
}

func (m *TCefStringMap) GetHandle() TCefStringMapHandle {
	r1 := stringMapImportAPI().SysCallN(4, m.Instance())
	return TCefStringMapHandle(r1)
}

func (m *TCefStringMap) GetSize() NativeUInt {
	r1 := stringMapImportAPI().SysCallN(6, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefStringMap) Find(key string) string {
	r1 := stringMapImportAPI().SysCallN(3, m.Instance(), PascalStr(key))
	return GoStr(r1)
}

func (m *TCefStringMap) GetKey(index NativeUInt) string {
	r1 := stringMapImportAPI().SysCallN(5, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringMap) GetValue(index NativeUInt) string {
	r1 := stringMapImportAPI().SysCallN(7, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringMap) Append(key, value string) bool {
	r1 := stringMapImportAPI().SysCallN(0, m.Instance(), PascalStr(key), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefStringMap) Clear() {
	stringMapImportAPI().SysCallN(1, m.Instance())
}

var (
	stringMapImport       *imports.Imports = nil
	stringMapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMap_Append", 0),
		/*1*/ imports.NewTable("CefStringMap_Clear", 0),
		/*2*/ imports.NewTable("CefStringMap_Create", 0),
		/*3*/ imports.NewTable("CefStringMap_Find", 0),
		/*4*/ imports.NewTable("CefStringMap_GetHandle", 0),
		/*5*/ imports.NewTable("CefStringMap_GetKey", 0),
		/*6*/ imports.NewTable("CefStringMap_GetSize", 0),
		/*7*/ imports.NewTable("CefStringMap_GetValue", 0),
	}
)

func stringMapImportAPI() *imports.Imports {
	if stringMapImport == nil {
		stringMapImport = NewDefaultImports()
		stringMapImport.SetImportTable(stringMapImportTables)
		stringMapImportTables = nil
	}
	return stringMapImport
}
