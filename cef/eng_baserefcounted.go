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
type ICefBaseRefCounted interface {
	IObject
	HasOneRef() bool                                 // function
	HasAtLeastOneRef() bool                          // function
	SameAs(aData uintptr) bool                       // function
	SameAs1(aBaseRefCounted ICefBaseRefCounted) bool // function
	Wrap() uintptr                                   // function
	DestroyOtherRefs()                               // procedure
}

// TCefBaseRefCounted Parent: TObject
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
