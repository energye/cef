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

// ICefStringListOwn Parent: ICefCustomStringList
type ICefStringListOwn interface {
	ICefCustomStringList
	AsIntfStringList() uintptr
}

type TCefStringListOwn struct {
	TCefCustomStringList
}

func (m *TCefStringListOwn) AsIntfStringList() uintptr {
	return m.GetIntfPointer(0)
}

// NewStringListOwn class constructor
func NewStringListOwn() ICefStringListOwn {
	var stringListPtr uintptr // ICefStringList
	r := cefStringListOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringListPtr)))
	ret := AsCefStringListOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringListPtr)
	}
	return ret
}

var (
	cefStringListOwnOnce   base.Once
	cefStringListOwnImport *imports.Imports = nil
)

func cefStringListOwnAPI() *imports.Imports {
	cefStringListOwnOnce.Do(func() {
		cefStringListOwnImport = api.NewDefaultImports()
		cefStringListOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStringListOwn_Create", 0), // constructor NewStringListOwn
		}
	})
	return cefStringListOwnImport
}
