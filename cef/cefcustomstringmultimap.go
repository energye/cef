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

// ICefStringMultimap Parent: IObject
type ICefStringMultimap interface {
	IObject

	GetHandle() cefTypes.TCefStringMultimap                         // function
	GetSize() cefTypes.NativeUInt                                   // function
	FindCount(key string) cefTypes.NativeUInt                       // function
	GetEnumerate(key string, valueIndex cefTypes.NativeUInt) string // function
	GetKey(index cefTypes.NativeUInt) string                        // function
	GetValue(index cefTypes.NativeUInt) string                      // function
	Append(key string, value string) bool                           // function
	Clear()                                                         // procedure
}

// ICefCustomStringMultimap Parent: ICefStringMultimap IInterfacedObject
type ICefCustomStringMultimap interface {
	ICefStringMultimap
	IInterfacedObject
	AsIntfStringMultimap() uintptr
}

type TCefCustomStringMultimap struct {
	TInterfacedObject
}

func (m *TCefCustomStringMultimap) GetHandle() cefTypes.TCefStringMultimap {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringMultimapAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefStringMultimap(r)
}

func (m *TCefCustomStringMultimap) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringMultimapAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefCustomStringMultimap) FindCount(key string) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefCustomStringMultimapAPI().SysCallN(3, m.Instance(), api.PasStr(key))
	return cefTypes.NativeUInt(r)
}

func (m *TCefCustomStringMultimap) GetEnumerate(key string, valueIndex cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMultimapAPI().SysCallN(4, m.Instance(), api.PasStr(key), uintptr(valueIndex), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMultimap) GetKey(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMultimapAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMultimap) GetValue(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefCustomStringMultimapAPI().SysCallN(6, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefCustomStringMultimap) Append(key string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefCustomStringMultimapAPI().SysCallN(7, m.Instance(), api.PasStr(key), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefCustomStringMultimap) Clear() {
	if !m.IsValid() {
		return
	}
	cefCustomStringMultimapAPI().SysCallN(8, m.Instance())
}

func (m *TCefCustomStringMultimap) AsIntfStringMultimap() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomStringMultimap class constructor
func NewCustomStringMultimap() ICefCustomStringMultimap {
	var stringMultimapPtr uintptr // ICefStringMultimap
	r := cefCustomStringMultimapAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringMultimapPtr)))
	ret := AsCefCustomStringMultimap(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringMultimapPtr)
	}
	return ret
}

var (
	cefCustomStringMultimapOnce   base.Once
	cefCustomStringMultimapImport *imports.Imports = nil
)

func cefCustomStringMultimapAPI() *imports.Imports {
	cefCustomStringMultimapOnce.Do(func() {
		cefCustomStringMultimapImport = api.NewDefaultImports()
		cefCustomStringMultimapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomStringMultimap_Create", 0), // constructor NewCustomStringMultimap
			/* 1 */ imports.NewTable("TCefCustomStringMultimap_GetHandle", 0), // function GetHandle
			/* 2 */ imports.NewTable("TCefCustomStringMultimap_GetSize", 0), // function GetSize
			/* 3 */ imports.NewTable("TCefCustomStringMultimap_FindCount", 0), // function FindCount
			/* 4 */ imports.NewTable("TCefCustomStringMultimap_GetEnumerate", 0), // function GetEnumerate
			/* 5 */ imports.NewTable("TCefCustomStringMultimap_GetKey", 0), // function GetKey
			/* 6 */ imports.NewTable("TCefCustomStringMultimap_GetValue", 0), // function GetValue
			/* 7 */ imports.NewTable("TCefCustomStringMultimap_Append", 0), // function Append
			/* 8 */ imports.NewTable("TCefCustomStringMultimap_Clear", 0), // procedure Clear
		}
	})
	return cefCustomStringMultimapImport
}
