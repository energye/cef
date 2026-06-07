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

// ICefStringMultimapOwn Parent: ICefCustomStringMultimap
type ICefStringMultimapOwn interface {
	ICefCustomStringMultimap
	AsIntfStringMultimap() uintptr
}

type TCefStringMultimapOwn struct {
	TCefCustomStringMultimap
}

func (m *TCefStringMultimapOwn) GetHandle() cefTypes.TCefStringMultimap {
	if !m.IsValid() {
		return 0
	}
	r := cefStringMultimapOwnAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefStringMultimap(r)
}

func (m *TCefStringMultimapOwn) GetSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefStringMultimapOwnAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefStringMultimapOwn) FindCount(key string) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefStringMultimapOwnAPI().SysCallN(3, m.Instance(), api.PasStr(key))
	return cefTypes.NativeUInt(r)
}

func (m *TCefStringMultimapOwn) GetEnumerate(key string, valueIndex cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefStringMultimapOwnAPI().SysCallN(4, m.Instance(), api.PasStr(key), uintptr(valueIndex), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefStringMultimapOwn) GetKey(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefStringMultimapOwnAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefStringMultimapOwn) GetValue(index cefTypes.NativeUInt) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefStringMultimapOwnAPI().SysCallN(6, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefStringMultimapOwn) Append(key string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefStringMultimapOwnAPI().SysCallN(7, m.Instance(), api.PasStr(key), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefStringMultimapOwn) Clear() {
	if !m.IsValid() {
		return
	}
	cefStringMultimapOwnAPI().SysCallN(8, m.Instance())
}

func (m *TCefStringMultimapOwn) AsIntfStringMultimap() uintptr {
	return m.GetIntfPointer(0)
}

// NewStringMultimapOwn class constructor
func NewStringMultimapOwn() ICefStringMultimapOwn {
	var stringMultimapPtr uintptr // ICefStringMultimap
	r := cefStringMultimapOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringMultimapPtr)))
	ret := AsCefStringMultimapOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringMultimapPtr)
	}
	return ret
}

var (
	cefStringMultimapOwnOnce   base.Once
	cefStringMultimapOwnImport *imports.Imports = nil
)

func cefStringMultimapOwnAPI() *imports.Imports {
	cefStringMultimapOwnOnce.Do(func() {
		cefStringMultimapOwnImport = api.NewDefaultImports()
		cefStringMultimapOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStringMultimapOwn_Create", 0), // constructor NewStringMultimapOwn
			/* 1 */ imports.NewTable("TCefStringMultimapOwn_GetHandle", 0), // function GetHandle
			/* 2 */ imports.NewTable("TCefStringMultimapOwn_GetSize", 0), // function GetSize
			/* 3 */ imports.NewTable("TCefStringMultimapOwn_FindCount", 0), // function FindCount
			/* 4 */ imports.NewTable("TCefStringMultimapOwn_GetEnumerate", 0), // function GetEnumerate
			/* 5 */ imports.NewTable("TCefStringMultimapOwn_GetKey", 0), // function GetKey
			/* 6 */ imports.NewTable("TCefStringMultimapOwn_GetValue", 0), // function GetValue
			/* 7 */ imports.NewTable("TCefStringMultimapOwn_Append", 0), // function Append
			/* 8 */ imports.NewTable("TCefStringMultimapOwn_Clear", 0), // procedure Clear
		}
	})
	return cefStringMultimapOwnImport
}
