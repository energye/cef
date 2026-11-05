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

// ICefBaseRefCounted Parent: IInterfacedObject
type ICefBaseRefCounted interface {
	IInterfacedObject

	// HasOneRef
	//  Returns true (1) if the current reference count is 1.
	HasOneRef() bool // function
	// HasAtLeastOneRef
	//  Returns true (1) if the current reference count is at least 1.
	HasAtLeastOneRef() bool // function
	// SameAs
	//  Compares the aData pointer with the FData field if the current instance.
	SameAs(data uintptr) bool // function
	// Wrap
	//  Called to increment the reference count for the object. Should be called
	//  for every new copy of a pointer to a given object.
	Wrap() uintptr // function
	// DestroyOtherRefs
	//  Releases all other instances.
	DestroyOtherRefs() // procedure
}

// ICefBaseRefCountedRef Parent: ICefBaseRefCounted IInterfacedObject
type ICefBaseRefCountedRef interface {
	ICefBaseRefCounted
	IInterfacedObject
}

type TCefBaseRefCountedRef struct {
	TInterfacedObject
}

func (m *TCefBaseRefCountedRef) HasOneRef() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedRef) HasAtLeastOneRef() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedRef) SameAs(data uintptr) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedRefAPI().SysCallN(3, m.Instance(), uintptr(data))
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedRef) Wrap() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefBaseRefCountedRefAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TCefBaseRefCountedRef) DestroyOtherRefs() {
	if !m.IsValid() {
		return
	}
	cefBaseRefCountedRefAPI().SysCallN(5, m.Instance())
}

// NewBaseRefCountedRef class constructor
func NewBaseRefCountedRef(data uintptr) ICefBaseRefCountedRef {
	r := cefBaseRefCountedRefAPI().SysCallN(0, uintptr(data))
	return AsCefBaseRefCountedRef(r)
}

var (
	cefBaseRefCountedRefOnce   base.Once
	cefBaseRefCountedRefImport *imports.Imports = nil
)

func cefBaseRefCountedRefAPI() *imports.Imports {
	cefBaseRefCountedRefOnce.Do(func() {
		cefBaseRefCountedRefImport = api.NewDefaultImports()
		cefBaseRefCountedRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBaseRefCountedRef_Create", 0), // constructor NewBaseRefCountedRef
			/* 1 */ imports.NewTable("TCefBaseRefCountedRef_HasOneRef", 0), // function HasOneRef
			/* 2 */ imports.NewTable("TCefBaseRefCountedRef_HasAtLeastOneRef", 0), // function HasAtLeastOneRef
			/* 3 */ imports.NewTable("TCefBaseRefCountedRef_SameAs", 0), // function SameAs
			/* 4 */ imports.NewTable("TCefBaseRefCountedRef_Wrap", 0), // function Wrap
			/* 5 */ imports.NewTable("TCefBaseRefCountedRef_DestroyOtherRefs", 0), // procedure DestroyOtherRefs
		}
	})
	return cefBaseRefCountedRefImport
}
