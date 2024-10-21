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

// ICefSharedMemoryRegion Parent: ICefBaseRefCounted
//
//	Interface that wraps platform-dependent share memory region mapping.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_memory_region_capi.h">CEF source file: /include/capi/cef_shared_memory_region_capi.h (cef_shared_memory_region_t))</a>
type ICefSharedMemoryRegion interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the mapping is valid.
	IsValid() bool // function
	// Size
	//  Returns the size of the mapping in bytes. Returns 0 for invalid instances.
	Size() NativeUInt // function
	// Memory
	//  Returns the pointer to the memory. Returns nullptr for invalid instances. The returned pointer is only valid for the life span of this object.
	Memory() uintptr // function
}

// TCefSharedMemoryRegion Parent: TCefBaseRefCounted
//
//	Interface that wraps platform-dependent share memory region mapping.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_memory_region_capi.h">CEF source file: /include/capi/cef_shared_memory_region_capi.h (cef_shared_memory_region_t))</a>
type TCefSharedMemoryRegion struct {
	TCefBaseRefCounted
}

// SharedMemoryRegionRef -> ICefSharedMemoryRegion
var SharedMemoryRegionRef sharedMemoryRegion

// sharedMemoryRegion TCefSharedMemoryRegion Ref
type sharedMemoryRegion uintptr

func (m *sharedMemoryRegion) UnWrap(data uintptr) ICefSharedMemoryRegion {
	var resultCefSharedMemoryRegion uintptr
	sharedMemoryRegionImportAPI().SysCallN(3, uintptr(data), uintptr(unsafePointer(&resultCefSharedMemoryRegion)))
	return AsCefSharedMemoryRegion(resultCefSharedMemoryRegion)
}

func (m *TCefSharedMemoryRegion) IsValid() bool {
	r1 := sharedMemoryRegionImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCefSharedMemoryRegion) Size() NativeUInt {
	r1 := sharedMemoryRegionImportAPI().SysCallN(2, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefSharedMemoryRegion) Memory() uintptr {
	r1 := sharedMemoryRegionImportAPI().SysCallN(1, m.Instance())
	return uintptr(r1)
}

var (
	sharedMemoryRegionImport       *imports.Imports = nil
	sharedMemoryRegionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSharedMemoryRegion_IsValid", 0),
		/*1*/ imports.NewTable("CefSharedMemoryRegion_Memory", 0),
		/*2*/ imports.NewTable("CefSharedMemoryRegion_Size", 0),
		/*3*/ imports.NewTable("CefSharedMemoryRegion_UnWrap", 0),
	}
)

func sharedMemoryRegionImportAPI() *imports.Imports {
	if sharedMemoryRegionImport == nil {
		sharedMemoryRegionImport = NewDefaultImports()
		sharedMemoryRegionImport.SetImportTable(sharedMemoryRegionImportTables)
		sharedMemoryRegionImportTables = nil
	}
	return sharedMemoryRegionImport
}
