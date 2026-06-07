//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/127/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefStringMap Parent: IObject
type ICefStringMap interface {
	IObject

	GetHandle() cefTypes.TCefStringMap         // function
	GetSize() cefTypes.NativeUInt              // function
	Find(key string) string                    // function
	GetKey(index cefTypes.NativeUInt) string   // function
	GetValue(index cefTypes.NativeUInt) string // function
	Append(key string, value string) bool      // function
	Clear()                                    // procedure
}

// ICefCustomStringMap Parent: ICefStringMap IInterfacedObject
type ICefCustomStringMap interface {
	ICefStringMap
	IInterfacedObject
	AsIntfStringMap() uintptr
}

type TCefCustomStringMap struct {
	TInterfacedObject
}

func (m *TCefCustomStringMap) GetHandle() cefTypes.TCefStringMap {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringMapAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefStringMap(r)
}

func (m *TCefCustomStringMap) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringMapAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefCustomStringMap) Find(key string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMapAPI().SysCallN(3, m.Instance(), api.PasStr(key), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMap) GetKey(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMap) GetValue(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMapAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMap) Append(key string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCustomStringMapAPI().SysCallN(6, m.Instance(), api.PasStr(key), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefCustomStringMap) Clear() {
	if !m.IsValid() {
		return
	}
	cefCustomStringMapAPI().SysCallN(7, m.Instance())
}

func (m *TCefCustomStringMap) AsIntfStringMap() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomStringMap class constructor
func NewCustomStringMap() ICefCustomStringMap {
	var stringMapPtr uintptr // ICefStringMap
	r := cefCustomStringMapAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringMapPtr)))
	ret := AsCefCustomStringMap(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringMapPtr)
	}
	return ret
}

var (
	cefCustomStringMapOnce   base.Once
	cefCustomStringMapImport *imports.Imports = nil
)

func cefCustomStringMapAPI() *imports.Imports {
	cefCustomStringMapOnce.Do(func() {
		cefCustomStringMapImport = api.NewDefaultImports()
		cefCustomStringMapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomStringMap_Create", 0), // constructor NewCustomStringMap
			/* 1 */ imports.NewTable("TCefCustomStringMap_GetHandle", 0), // function GetHandle
			/* 2 */ imports.NewTable("TCefCustomStringMap_GetSize", 0), // function GetSize
			/* 3 */ imports.NewTable("TCefCustomStringMap_Find", 0), // function Find
			/* 4 */ imports.NewTable("TCefCustomStringMap_GetKey", 0), // function GetKey
			/* 5 */ imports.NewTable("TCefCustomStringMap_GetValue", 0), // function GetValue
			/* 6 */ imports.NewTable("TCefCustomStringMap_Append", 0), // function Append
			/* 7 */ imports.NewTable("TCefCustomStringMap_Clear", 0), // procedure Clear
		}
	})
	return cefCustomStringMapImport
}
