//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/147/types"
)

// ICefStringList Parent: IObject
type ICefStringList interface {
	IObject

	GetHandle() cefTypes.TCefStringList        // function
	GetSize() cefTypes.NativeUInt              // function
	GetValue(index cefTypes.NativeUInt) string // function
	Copy() cefTypes.TCefStringList             // function
	Append(value string)                       // procedure
	Clear()                                    // procedure
	CopyToStrings(strings lcl.IStrings)        // procedure
	AddStrings(strings lcl.IStrings)           // procedure
}

// ICefCustomStringList Parent: ICefStringList IInterfacedObject
type ICefCustomStringList interface {
	ICefStringList
	IInterfacedObject
	AsIntfStringList() uintptr
}

type TCefCustomStringList struct {
	TInterfacedObject
}

func (m *TCefCustomStringList) GetHandle() cefTypes.TCefStringList {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringListAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefStringList(r)
}

func (m *TCefCustomStringList) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringListAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefCustomStringList) GetValue(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringListAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringList) Copy() cefTypes.TCefStringList {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringListAPI().SysCallN(4, m.Instance())
	return cefTypes.TCefStringList(r)
}

func (m *TCefCustomStringList) Append(value string) {
	if !m.IsValid() {
		return
	}
	cefCustomStringListAPI().SysCallN(5, m.Instance(), api.PasStr(value))
}

func (m *TCefCustomStringList) Clear() {
	if !m.IsValid() {
		return
	}
	cefCustomStringListAPI().SysCallN(6, m.Instance())
}

func (m *TCefCustomStringList) CopyToStrings(strings lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefCustomStringListAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TCefCustomStringList) AddStrings(strings lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefCustomStringListAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TCefCustomStringList) AsIntfStringList() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomStringList class constructor
func NewCustomStringList() ICefCustomStringList {
	var stringListPtr uintptr // ICefStringList
	r := cefCustomStringListAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringListPtr)))
	ret := AsCefCustomStringList(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringListPtr)
	}
	return ret
}

var (
	cefCustomStringListOnce   base.Once
	cefCustomStringListImport *imports.Imports = nil
)

func cefCustomStringListAPI() *imports.Imports {
	cefCustomStringListOnce.Do(func() {
		cefCustomStringListImport = api.NewDefaultImports()
		cefCustomStringListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomStringList_Create", 0), // constructor NewCustomStringList
			/* 1 */ imports.NewTable("TCefCustomStringList_GetHandle", 0), // function GetHandle
			/* 2 */ imports.NewTable("TCefCustomStringList_GetSize", 0), // function GetSize
			/* 3 */ imports.NewTable("TCefCustomStringList_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TCefCustomStringList_Copy", 0), // function Copy
			/* 5 */ imports.NewTable("TCefCustomStringList_Append", 0), // procedure Append
			/* 6 */ imports.NewTable("TCefCustomStringList_Clear", 0), // procedure Clear
			/* 7 */ imports.NewTable("TCefCustomStringList_CopyToStrings", 0), // procedure CopyToStrings
			/* 8 */ imports.NewTable("TCefCustomStringList_AddStrings", 0), // procedure AddStrings
		}
	})
	return cefCustomStringListImport
}
