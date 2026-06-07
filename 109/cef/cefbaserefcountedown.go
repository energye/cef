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

// ICefBaseRefCountedOwn Parent: ICefBaseRefCounted IInterfacedObject
type ICefBaseRefCountedOwn interface {
	ICefBaseRefCounted
	IInterfacedObject
}

type TCefBaseRefCountedOwn struct {
	TInterfacedObject
}

func (m *TCefBaseRefCountedOwn) HasOneRef() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedOwnAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedOwn) HasAtLeastOneRef() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedOwnAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedOwn) SameAs(data uintptr) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBaseRefCountedOwnAPI().SysCallN(3, m.Instance(), uintptr(data))
	return api.GoBool(r)
}

func (m *TCefBaseRefCountedOwn) Wrap() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefBaseRefCountedOwnAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TCefBaseRefCountedOwn) DestroyOtherRefs() {
	if !m.IsValid() {
		return
	}
	cefBaseRefCountedOwnAPI().SysCallN(5, m.Instance())
}

// NewBaseRefCountedOwnData class constructor
func NewBaseRefCountedOwnData(size uint32, owned bool) ICefBaseRefCountedOwn {
	r := cefBaseRefCountedOwnAPI().SysCallN(0, uintptr(size), api.PasBool(owned))
	return AsCefBaseRefCountedOwn(r)
}

var (
	cefBaseRefCountedOwnOnce   base.Once
	cefBaseRefCountedOwnImport *imports.Imports = nil
)

func cefBaseRefCountedOwnAPI() *imports.Imports {
	cefBaseRefCountedOwnOnce.Do(func() {
		cefBaseRefCountedOwnImport = api.NewDefaultImports()
		cefBaseRefCountedOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBaseRefCountedOwn_CreateData", 0), // constructor NewBaseRefCountedOwnData
			/* 1 */ imports.NewTable("TCefBaseRefCountedOwn_HasOneRef", 0), // function HasOneRef
			/* 2 */ imports.NewTable("TCefBaseRefCountedOwn_HasAtLeastOneRef", 0), // function HasAtLeastOneRef
			/* 3 */ imports.NewTable("TCefBaseRefCountedOwn_SameAs", 0), // function SameAs
			/* 4 */ imports.NewTable("TCefBaseRefCountedOwn_Wrap", 0), // function Wrap
			/* 5 */ imports.NewTable("TCefBaseRefCountedOwn_DestroyOtherRefs", 0), // procedure DestroyOtherRefs
		}
	})
	return cefBaseRefCountedOwnImport
}
