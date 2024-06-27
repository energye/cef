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

// ICefStringList Parent: IObject
//
//	Custom interface used to handle all the CEF functions related to CefStringList.
type ICefStringList interface {
	IObject
	GetHandle() TCefStringListPtr     // function
	GetSize() NativeUInt              // function
	GetValue(index NativeUInt) string // function
	Copy() TCefStringListPtr          // function
	Append(value string)              // procedure
	Clear()                           // procedure
	CopyToStrings(aStrings IStrings)  // procedure
	AddStrings(aStrings IStrings)     // procedure
}

// TCefStringList Parent: TObject
//
//	Custom interface used to handle all the CEF functions related to CefStringList.
type TCefStringList struct {
	TObject
}

func (m *TCefStringList) GetHandle() TCefStringListPtr {
	r1 := stringListImportAPI().SysCallN(5, m.Instance())
	return TCefStringListPtr(r1)
}

func (m *TCefStringList) GetSize() NativeUInt {
	r1 := stringListImportAPI().SysCallN(6, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefStringList) GetValue(index NativeUInt) string {
	r1 := stringListImportAPI().SysCallN(7, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefStringList) Copy() TCefStringListPtr {
	r1 := stringListImportAPI().SysCallN(3, m.Instance())
	return TCefStringListPtr(r1)
}

func (m *TCefStringList) Append(value string) {
	stringListImportAPI().SysCallN(1, m.Instance(), PascalStr(value))
}

func (m *TCefStringList) Clear() {
	stringListImportAPI().SysCallN(2, m.Instance())
}

func (m *TCefStringList) CopyToStrings(aStrings IStrings) {
	stringListImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(aStrings))
}

func (m *TCefStringList) AddStrings(aStrings IStrings) {
	stringListImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(aStrings))
}

var (
	stringListImport       *imports.Imports = nil
	stringListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ICefStringList_AddStrings", 0),
		/*1*/ imports.NewTable("ICefStringList_Append", 0),
		/*2*/ imports.NewTable("ICefStringList_Clear", 0),
		/*3*/ imports.NewTable("ICefStringList_Copy", 0),
		/*4*/ imports.NewTable("ICefStringList_CopyToStrings", 0),
		/*5*/ imports.NewTable("ICefStringList_GetHandle", 0),
		/*6*/ imports.NewTable("ICefStringList_GetSize", 0),
		/*7*/ imports.NewTable("ICefStringList_GetValue", 0),
	}
)

func stringListImportAPI() *imports.Imports {
	if stringListImport == nil {
		stringListImport = NewDefaultImports()
		stringListImport.SetImportTable(stringListImportTables)
		stringListImportTables = nil
	}
	return stringListImport
}
