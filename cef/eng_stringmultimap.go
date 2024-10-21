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

// ICefStringMultimap Parent: IObject
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type ICefStringMultimap interface {
	IObject
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringMultimap     // procedure
	GetHandle() TCefStringMultimapHandle // function
	// GetSize
	//  Return the number of elements in the string multimap.
	GetSize() NativeUInt // function
	// FindCount
	//  Return the number of values with the specified key.
	FindCount(key string) NativeUInt // function
	// GetEnumerate
	//  Return the value_index-th value with the specified key.
	GetEnumerate(key string, valueIndex NativeUInt) string // function
	// GetKey
	//  Return the key at the specified zero-based string multimap index.
	GetKey(index NativeUInt) string // function
	// GetValue
	//  Return the value at the specified zero-based string multimap index.
	GetValue(index NativeUInt) string // function
	// Append
	//  Append a new key/value pair at the end of the string multimap.
	Append(key, value string) bool // function
	// Clear
	//  Clear the string multimap.
	Clear() // procedure
}

// TCefStringMultimap Parent: TObject
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type TCefStringMultimap struct {
	TObject
}

func NewCefStringMultimap() ICefStringMultimap {
	r1 := stringMultimapImportAPI().SysCallN(3)
	return AsCefStringMultimap(r1)
}

func (m *TCefStringMultimap) AsInterface() ICefStringMultimap {
	var resultCefStringMultimap uintptr
	stringMultimapImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefStringMultimap)))
	return AsCefStringMultimap(resultCefStringMultimap)
}

func (m *TCefStringMultimap) GetHandle() TCefStringMultimapHandle {
	r1 := stringMultimapImportAPI().SysCallN(6, m.Instance())
	return TCefStringMultimapHandle(r1)
}

func (m *TCefStringMultimap) GetSize() NativeUInt {
	r1 := stringMultimapImportAPI().SysCallN(8, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefStringMultimap) FindCount(key string) NativeUInt {
	r1 := stringMultimapImportAPI().SysCallN(4, m.Instance(), PascalStr(key))
	return NativeUInt(r1)
}

func (m *TCefStringMultimap) GetEnumerate(key string, valueIndex NativeUInt) string {
	r1 := stringMultimapImportAPI().SysCallN(5, m.Instance(), PascalStr(key), uintptr(valueIndex))
	return GoStr(r1)
}

func (m *TCefStringMultimap) GetKey(index NativeUInt) (r string) {
	value := NewTString()
	stringMultimapImportAPI().SysCallN(7, m.Instance(), uintptr(index), value.Instance())
	r = value.Value()
	value.Free()
	return
}

func (m *TCefStringMultimap) GetValue(index NativeUInt) string {
	r1 := stringMultimapImportAPI().SysCallN(9, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringMultimap) Append(key, value string) bool {
	r1 := stringMultimapImportAPI().SysCallN(0, m.Instance(), PascalStr(key), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefStringMultimap) Clear() {
	stringMultimapImportAPI().SysCallN(2, m.Instance())
}

var (
	stringMultimapImport       *imports.Imports = nil
	stringMultimapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMultimap_Append", 0),
		/*1*/ imports.NewTable("CefStringMultimap_AsInterface", 0),
		/*2*/ imports.NewTable("CefStringMultimap_Clear", 0),
		/*3*/ imports.NewTable("CefStringMultimap_Create", 0),
		/*4*/ imports.NewTable("CefStringMultimap_FindCount", 0),
		/*5*/ imports.NewTable("CefStringMultimap_GetEnumerate", 0),
		/*6*/ imports.NewTable("CefStringMultimap_GetHandle", 0),
		/*7*/ imports.NewTable("CefStringMultimap_GetKey", 0),
		/*8*/ imports.NewTable("CefStringMultimap_GetSize", 0),
		/*9*/ imports.NewTable("CefStringMultimap_GetValue", 0),
	}
)

func stringMultimapImportAPI() *imports.Imports {
	if stringMultimapImport == nil {
		stringMultimapImport = NewDefaultImports()
		stringMultimapImport.SetImportTable(stringMultimapImportTables)
		stringMultimapImportTables = nil
	}
	return stringMultimapImport
}
