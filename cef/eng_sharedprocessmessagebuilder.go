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
type ICefSharedProcessMessageBuilder interface {
	ICefBaseRefCounted
	IsValid() bool             // function
	Size() NativeUInt          // function
	Memory() uintptr           // function
	Build() ICefProcessMessage // function
}

// TCefSharedProcessMessageBuilder Parent: TCefBaseRefCounted
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
