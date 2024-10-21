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

// ICefBaseRefCounted Parent: IObject
//
//	All ref-counted framework interfaces must inherit from this interface.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_base_capi.h">CEF source file: /include/capi/cef_base_capi.h (cef_base_ref_counted_t))</a>
type ICefBaseRefCounted interface {
	IObject
	FreeAndNil()
	// HasOneRef
	//  Returns true (1) if the current reference count is 1.
	HasOneRef() bool // function
	// HasAtLeastOneRef
	//  Returns true (1) if the current reference count is at least 1.
	HasAtLeastOneRef() bool // function
	// SameAs
	//  Compares the aData pointer with the FData field if the current instance.
	SameAs(aData uintptr) bool // function
	// SameAs1
	//  Compares the aData pointer with the FData field if the current instance.
	SameAs1(aBaseRefCounted ICefBaseRefCounted) bool // function
	// Wrap
	//  Called to increment the reference count for the object. Should be called for every new copy of a pointer to a given object.
	Wrap() uintptr // function
	// DestroyOtherRefs
	//  Releases all other instances.
	DestroyOtherRefs() // procedure
}

// TCefBaseRefCounted Parent: TObject
//
//	All ref-counted framework interfaces must inherit from this interface.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_base_capi.h">CEF source file: /include/capi/cef_base_capi.h (cef_base_ref_counted_t))</a>
type TCefBaseRefCounted struct {
	TObject
}

func NewCefBaseRefCounted(data uintptr) ICefBaseRefCounted {
	r1 := baseRefCountedImportAPI().SysCallN(0, uintptr(data))
	return AsCefBaseRefCounted(r1)
}

// BaseRefCountedRef -> ICefBaseRefCounted
var BaseRefCountedRef baseRefCounted

// baseRefCounted TCefBaseRefCounted Ref
type baseRefCounted uintptr

func (m *baseRefCounted) UnWrap(data uintptr) ICefBaseRefCounted {
	var resultCefBaseRefCounted uintptr
	baseRefCountedImportAPI().SysCallN(6, uintptr(data), uintptr(unsafePointer(&resultCefBaseRefCounted)))
	return AsCefBaseRefCounted(resultCefBaseRefCounted)
}

func (m *TCefBaseRefCounted) HasOneRef() bool {
	r1 := baseRefCountedImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) HasAtLeastOneRef() bool {
	r1 := baseRefCountedImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) SameAs(aData uintptr) bool {
	r1 := baseRefCountedImportAPI().SysCallN(4, m.Instance(), uintptr(aData))
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) SameAs1(aBaseRefCounted ICefBaseRefCounted) bool {
	r1 := baseRefCountedImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(aBaseRefCounted))
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) Wrap() uintptr {
	r1 := baseRefCountedImportAPI().SysCallN(7, m.Instance())
	return uintptr(r1)
}

func (m *TCefBaseRefCounted) DestroyOtherRefs() {
	baseRefCountedImportAPI().SysCallN(1, m.Instance())
}

var (
	baseRefCountedImport       *imports.Imports = nil
	baseRefCountedImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBaseRefCounted_Create", 0),
		/*1*/ imports.NewTable("CefBaseRefCounted_DestroyOtherRefs", 0),
		/*2*/ imports.NewTable("CefBaseRefCounted_HasAtLeastOneRef", 0),
		/*3*/ imports.NewTable("CefBaseRefCounted_HasOneRef", 0),
		/*4*/ imports.NewTable("CefBaseRefCounted_SameAs", 0),
		/*5*/ imports.NewTable("CefBaseRefCounted_SameAs1", 0),
		/*6*/ imports.NewTable("CefBaseRefCounted_UnWrap", 0),
		/*7*/ imports.NewTable("CefBaseRefCounted_Wrap", 0),
	}
)

func baseRefCountedImportAPI() *imports.Imports {
	if baseRefCountedImport == nil {
		baseRefCountedImport = NewDefaultImports()
		baseRefCountedImport.SetImportTable(baseRefCountedImportTables)
		baseRefCountedImportTables = nil
	}
	return baseRefCountedImport
}
