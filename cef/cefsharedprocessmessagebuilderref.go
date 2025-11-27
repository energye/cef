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

// ICefSharedProcessMessageBuilder Parent: ICefBaseRefCountedRef
type ICefSharedProcessMessageBuilder interface {
	ICefBaseRefCountedRef
	// IsValid
	//  Returns true (1) if the builder is valid.
	IsValid() bool // function
	// Size
	//  Returns the size of the shared memory region in bytes. Returns 0 for
	//  invalid instances.
	Size() cefTypes.NativeUInt // function
	// Memory
	//  Returns the pointer to the writable memory. Returns nullptr for invalid
	//  instances. The returned pointer is only valid for the life span of this
	//  object.
	Memory() uintptr // function
	// Build
	//  Creates a new ICefProcessMessage from the data provided to the builder.
	//  Returns nullptr for invalid instances. Invalidates the builder instance.
	Build() ICefProcessMessage // function
}

// ICefSharedProcessMessageBuilderRef Parent: ICefSharedProcessMessageBuilder
type ICefSharedProcessMessageBuilderRef interface {
	ICefSharedProcessMessageBuilder
	AsIntfSharedProcessMessageBuilder() uintptr
}

type TCefSharedProcessMessageBuilderRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefSharedProcessMessageBuilderRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefSharedProcessMessageBuilderRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefSharedProcessMessageBuilderRef) Size() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefSharedProcessMessageBuilderRefAPI().SysCallN(2, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefSharedProcessMessageBuilderRef) Memory() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefSharedProcessMessageBuilderRefAPI().SysCallN(3, m.Instance())
	return uintptr(r)
}

func (m *TCefSharedProcessMessageBuilderRef) Build() (result ICefProcessMessage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefSharedProcessMessageBuilderRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefProcessMessageRef(resultPtr)
	return
}

func (m *TCefSharedProcessMessageBuilderRef) AsIntfSharedProcessMessageBuilder() uintptr {
	return m.GetIntfPointer(0)
}

// SharedProcessMessageBuilderRef  is static instance
var SharedProcessMessageBuilderRef _SharedProcessMessageBuilderRefClass

// _SharedProcessMessageBuilderRefClass is class type defined by TCefSharedProcessMessageBuilderRef
type _SharedProcessMessageBuilderRefClass uintptr

func (_SharedProcessMessageBuilderRefClass) UnWrap(data uintptr) (result ICefSharedProcessMessageBuilder) {
	var resultPtr uintptr
	cefSharedProcessMessageBuilderRefAPI().SysCallN(5, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSharedProcessMessageBuilderRef(resultPtr)
	return
}

func (_SharedProcessMessageBuilderRefClass) CreateBuilder(name string, byteSize cefTypes.NativeUInt) (result ICefSharedProcessMessageBuilder) {
	var resultPtr uintptr
	cefSharedProcessMessageBuilderRefAPI().SysCallN(6, api.PasStr(name), uintptr(byteSize), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSharedProcessMessageBuilderRef(resultPtr)
	return
}

// NewSharedProcessMessageBuilderRef class constructor
func NewSharedProcessMessageBuilderRef(data uintptr) ICefSharedProcessMessageBuilderRef {
	var sharedProcessMessageBuilderPtr uintptr // ICefSharedProcessMessageBuilder
	r := cefSharedProcessMessageBuilderRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&sharedProcessMessageBuilderPtr)))
	ret := AsCefSharedProcessMessageBuilderRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, sharedProcessMessageBuilderPtr)
	}
	return ret
}

var (
	cefSharedProcessMessageBuilderRefOnce   base.Once
	cefSharedProcessMessageBuilderRefImport *imports.Imports = nil
)

func cefSharedProcessMessageBuilderRefAPI() *imports.Imports {
	cefSharedProcessMessageBuilderRefOnce.Do(func() {
		cefSharedProcessMessageBuilderRefImport = api.NewDefaultImports()
		cefSharedProcessMessageBuilderRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_Create", 0), // constructor NewSharedProcessMessageBuilderRef
			/* 1 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_Size", 0), // function Size
			/* 3 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_Memory", 0), // function Memory
			/* 4 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_Build", 0), // function Build
			/* 5 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_UnWrap", 0), // static function UnWrap
			/* 6 */ imports.NewTable("TCefSharedProcessMessageBuilderRef_CreateBuilder", 0), // static function CreateBuilder
		}
	})
	return cefSharedProcessMessageBuilderRefImport
}
