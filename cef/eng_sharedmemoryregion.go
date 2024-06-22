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
type ICefSharedMemoryRegion interface {
	ICefBaseRefCounted
	IsValid() bool    // function
	Size() NativeUInt // function
	Memory() uintptr  // function
}

// TCefSharedMemoryRegion Parent: TCefBaseRefCounted
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
