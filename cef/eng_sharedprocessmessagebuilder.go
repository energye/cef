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

// ICefSharedProcessMessageBuilder Parent: ICefBaseRefCounted
//
//	Interface that builds a ICefProcessMessage containing a shared memory region. This interface is not thread-safe but may be used exclusively on a different thread from the one which constructed it.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_process_message_builder_capi.h">CEF source file: /include/capi/cef_shared_process_message_builder_capi.h (cef_shared_process_message_builder_t))</a>
type ICefSharedProcessMessageBuilder interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the builder is valid.
	IsValid() bool // function
	// Size
	//  Returns the size of the shared memory region in bytes. Returns 0 for invalid instances.
	Size() NativeUInt // function
	// Memory
	//  Returns the pointer to the writable memory. Returns nullptr for invalid instances. The returned pointer is only valid for the life span of this object.
	Memory() uintptr // function
	// Build
	//  Creates a new ICefProcessMessage from the data provided to the builder. Returns nullptr for invalid instances. Invalidates the builder instance.
	Build() ICefProcessMessage // function
}

// TCefSharedProcessMessageBuilder Parent: TCefBaseRefCounted
//
//	Interface that builds a ICefProcessMessage containing a shared memory region. This interface is not thread-safe but may be used exclusively on a different thread from the one which constructed it.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_process_message_builder_capi.h">CEF source file: /include/capi/cef_shared_process_message_builder_capi.h (cef_shared_process_message_builder_t))</a>
type TCefSharedProcessMessageBuilder struct {
	TCefBaseRefCounted
}

// SharedProcessMessageBuilderRef -> ICefSharedProcessMessageBuilder
var SharedProcessMessageBuilderRef sharedProcessMessageBuilder

// sharedProcessMessageBuilder TCefSharedProcessMessageBuilder Ref
type sharedProcessMessageBuilder uintptr

func (m *sharedProcessMessageBuilder) UnWrap(data uintptr) ICefSharedProcessMessageBuilder {
	var resultCefSharedProcessMessageBuilder uintptr
	sharedProcessMessageBuilderImportAPI().SysCallN(5, uintptr(data), uintptr(unsafePointer(&resultCefSharedProcessMessageBuilder)))
	return AsCefSharedProcessMessageBuilder(resultCefSharedProcessMessageBuilder)
}

func (m *sharedProcessMessageBuilder) CreateBuilder(name string, bytesize NativeUInt) ICefSharedProcessMessageBuilder {
	var resultCefSharedProcessMessageBuilder uintptr
	sharedProcessMessageBuilderImportAPI().SysCallN(1, PascalStr(name), uintptr(bytesize), uintptr(unsafePointer(&resultCefSharedProcessMessageBuilder)))
	return AsCefSharedProcessMessageBuilder(resultCefSharedProcessMessageBuilder)
}

func (m *TCefSharedProcessMessageBuilder) IsValid() bool {
	r1 := sharedProcessMessageBuilderImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCefSharedProcessMessageBuilder) Size() NativeUInt {
	r1 := sharedProcessMessageBuilderImportAPI().SysCallN(4, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefSharedProcessMessageBuilder) Memory() uintptr {
	r1 := sharedProcessMessageBuilderImportAPI().SysCallN(3, m.Instance())
	return uintptr(r1)
}

func (m *TCefSharedProcessMessageBuilder) Build() ICefProcessMessage {
	var resultCefProcessMessage uintptr
	sharedProcessMessageBuilderImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

var (
	sharedProcessMessageBuilderImport       *imports.Imports = nil
	sharedProcessMessageBuilderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSharedProcessMessageBuilder_Build", 0),
		/*1*/ imports.NewTable("CefSharedProcessMessageBuilder_CreateBuilder", 0),
		/*2*/ imports.NewTable("CefSharedProcessMessageBuilder_IsValid", 0),
		/*3*/ imports.NewTable("CefSharedProcessMessageBuilder_Memory", 0),
		/*4*/ imports.NewTable("CefSharedProcessMessageBuilder_Size", 0),
		/*5*/ imports.NewTable("CefSharedProcessMessageBuilder_UnWrap", 0),
	}
)

func sharedProcessMessageBuilderImportAPI() *imports.Imports {
	if sharedProcessMessageBuilderImport == nil {
		sharedProcessMessageBuilderImport = NewDefaultImports()
		sharedProcessMessageBuilderImport.SetImportTable(sharedProcessMessageBuilderImportTables)
		sharedProcessMessageBuilderImportTables = nil
	}
	return sharedProcessMessageBuilderImport
}
