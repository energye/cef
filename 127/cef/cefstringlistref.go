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

	cefTypes "github.com/energye/cef/127/types"
)

// ICefStringListRef Parent: ICefCustomStringList
type ICefStringListRef interface {
	ICefCustomStringList
	AsIntfStringList() uintptr
}

type TCefStringListRef struct {
	TCefCustomStringList
}

func (m *TCefStringListRef) AsIntfStringList() uintptr {
	return m.GetIntfPointer(0)
}

// NewStringListRef class constructor
func NewStringListRef(handle cefTypes.TCefStringList) ICefStringListRef {
	var stringListPtr uintptr // ICefStringList
	r := cefStringListRefAPI().SysCallN(0, uintptr(handle), uintptr(base.UnsafePointer(&stringListPtr)))
	ret := AsCefStringListRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringListPtr)
	}
	return ret
}

var (
	cefStringListRefOnce   base.Once
	cefStringListRefImport *imports.Imports = nil
)

func cefStringListRefAPI() *imports.Imports {
	cefStringListRefOnce.Do(func() {
		cefStringListRefImport = api.NewDefaultImports()
		cefStringListRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStringListRef_Create", 0), // constructor NewStringListRef
		}
	})
	return cefStringListRefImport
}
