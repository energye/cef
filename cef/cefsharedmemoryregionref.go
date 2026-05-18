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

// ICefSharedMemoryRegion Parent: ICefBaseRefCounted
type ICefSharedMemoryRegion interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the mapping is valid.
	IsValid() bool // function
	// Size
	//  Returns the size of the mapping in bytes. Returns 0 for invalid instances.
	Size() cefTypes.NativeUInt // function
	// Memory
	//  Returns the pointer to the memory. Returns nullptr for invalid instances.
	//  The returned pointer is only valid for the life span of this object.
	Memory() uintptr // function
}

// ICefSharedMemoryRegionRef Parent: ICefSharedMemoryRegion ICefBaseRefCountedRef
type ICefSharedMemoryRegionRef interface {
	ICefSharedMemoryRegion
	ICefBaseRefCountedRef
	AsIntfSharedMemoryRegion() uintptr
}

type TCefSharedMemoryRegionRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefSharedMemoryRegionRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefSharedMemoryRegionRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefSharedMemoryRegionRef) Size() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefSharedMemoryRegionRefAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefSharedMemoryRegionRef) Memory() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefSharedMemoryRegionRefAPI().SysCallN(3, m.Instance())
	return uintptr(r)
}

func (m *TCefSharedMemoryRegionRef) AsIntfSharedMemoryRegion() uintptr {
	return m.GetIntfPointer(0)
}

// SharedMemoryRegionRef  is static instance
var SharedMemoryRegionRef _SharedMemoryRegionRefClass

// _SharedMemoryRegionRefClass is class type defined by TCefSharedMemoryRegionRef
type _SharedMemoryRegionRefClass uintptr

func (_SharedMemoryRegionRefClass) UnWrap(data uintptr) (result ICefSharedMemoryRegion) {
	var resultPtr uintptr
	cefSharedMemoryRegionRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSharedMemoryRegionRef(resultPtr)
	return
}

// NewSharedMemoryRegionRef class constructor
func NewSharedMemoryRegionRef(data uintptr) ICefSharedMemoryRegionRef {
	var sharedMemoryRegionPtr uintptr // ICefSharedMemoryRegion
	r := cefSharedMemoryRegionRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&sharedMemoryRegionPtr)))
	ret := AsCefSharedMemoryRegionRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, sharedMemoryRegionPtr)
	}
	return ret
}

var (
	cefSharedMemoryRegionRefOnce   base.Once
	cefSharedMemoryRegionRefImport *imports.Imports = nil
)

func cefSharedMemoryRegionRefAPI() *imports.Imports {
	cefSharedMemoryRegionRefOnce.Do(func() {
		cefSharedMemoryRegionRefImport = api.NewDefaultImports()
		cefSharedMemoryRegionRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSharedMemoryRegionRef_Create", 0), // constructor NewSharedMemoryRegionRef
			/* 1 */ imports.NewTable("TCefSharedMemoryRegionRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefSharedMemoryRegionRef_Size", 0), // function Size
			/* 3 */ imports.NewTable("TCefSharedMemoryRegionRef_Memory", 0), // function Memory
			/* 4 */ imports.NewTable("TCefSharedMemoryRegionRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefSharedMemoryRegionRefImport
}
