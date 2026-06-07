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
)

// ICefProcessMessage Parent: ICefBaseRefCounted
type ICefProcessMessage interface {
	ICefBaseRefCounted
	IsValid() bool                                 // function
	IsReadOnly() bool                              // function
	Copy() ICefProcessMessage                      // function
	GetName() string                               // function
	GetArgumentList() ICefListValue                // function
	GetSharedMemoryRegion() ICefSharedMemoryRegion // function
}

// ICefProcessMessageRef Parent: ICefProcessMessage ICefBaseRefCountedRef
type ICefProcessMessageRef interface {
	ICefProcessMessage
	ICefBaseRefCountedRef
	AsIntfProcessMessage() uintptr
}

type TCefProcessMessageRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefProcessMessageRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefProcessMessageRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefProcessMessageRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefProcessMessageRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefProcessMessageRef) Copy() (result ICefProcessMessage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefProcessMessageRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefProcessMessageRef(resultPtr)
	return
}

func (m *TCefProcessMessageRef) GetName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefProcessMessageRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefProcessMessageRef) GetArgumentList() (result ICefListValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefProcessMessageRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefListValueRef(resultPtr)
	return
}

func (m *TCefProcessMessageRef) GetSharedMemoryRegion() (result ICefSharedMemoryRegion) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefProcessMessageRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSharedMemoryRegionRef(resultPtr)
	return
}

func (m *TCefProcessMessageRef) AsIntfProcessMessage() uintptr {
	return m.GetIntfPointer(0)
}

// ProcessMessageRef  is static instance
var ProcessMessageRef _ProcessMessageRefClass

// _ProcessMessageRefClass is class type defined by TCefProcessMessageRef
type _ProcessMessageRefClass uintptr

func (_ProcessMessageRefClass) UnWrap(data uintptr) (result ICefProcessMessage) {
	var resultPtr uintptr
	cefProcessMessageRefAPI().SysCallN(7, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefProcessMessageRef(resultPtr)
	return
}

func (_ProcessMessageRefClass) New(name string) (result ICefProcessMessage) {
	var resultPtr uintptr
	cefProcessMessageRefAPI().SysCallN(8, api.PasStr(name), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefProcessMessageRef(resultPtr)
	return
}

// NewProcessMessageRef class constructor
func NewProcessMessageRef(data uintptr) ICefProcessMessageRef {
	var processMessagePtr uintptr // ICefProcessMessage
	r := cefProcessMessageRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&processMessagePtr)))
	ret := AsCefProcessMessageRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, processMessagePtr)
	}
	return ret
}

var (
	cefProcessMessageRefOnce   base.Once
	cefProcessMessageRefImport *imports.Imports = nil
)

func cefProcessMessageRefAPI() *imports.Imports {
	cefProcessMessageRefOnce.Do(func() {
		cefProcessMessageRefImport = api.NewDefaultImports()
		cefProcessMessageRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefProcessMessageRef_Create", 0), // constructor NewProcessMessageRef
			/* 1 */ imports.NewTable("TCefProcessMessageRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefProcessMessageRef_IsReadOnly", 0), // function IsReadOnly
			/* 3 */ imports.NewTable("TCefProcessMessageRef_Copy", 0), // function Copy
			/* 4 */ imports.NewTable("TCefProcessMessageRef_GetName", 0), // function GetName
			/* 5 */ imports.NewTable("TCefProcessMessageRef_GetArgumentList", 0), // function GetArgumentList
			/* 6 */ imports.NewTable("TCefProcessMessageRef_GetSharedMemoryRegion", 0), // function GetSharedMemoryRegion
			/* 7 */ imports.NewTable("TCefProcessMessageRef_UnWrap", 0), // static function UnWrap
			/* 8 */ imports.NewTable("TCefProcessMessageRef_New", 0), // static function New
		}
	})
	return cefProcessMessageRefImport
}
