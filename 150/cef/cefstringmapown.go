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

// ICefStringMapOwn Parent: ICefCustomStringMap
type ICefStringMapOwn interface {
	ICefCustomStringMap
	AsIntfStringMap() uintptr
}

type TCefStringMapOwn struct {
	TCefCustomStringMap
}

func (m *TCefStringMapOwn) AsIntfStringMap() uintptr {
	return m.GetIntfPointer(0)
}

// NewStringMapOwn class constructor
func NewStringMapOwn() ICefStringMapOwn {
	var stringMapPtr uintptr // ICefStringMap
	r := cefStringMapOwnAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringMapPtr)))
	ret := AsCefStringMapOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringMapPtr)
	}
	return ret
}

var (
	cefStringMapOwnOnce   base.Once
	cefStringMapOwnImport *imports.Imports = nil
)

func cefStringMapOwnAPI() *imports.Imports {
	cefStringMapOwnOnce.Do(func() {
		cefStringMapOwnImport = api.NewDefaultImports()
		cefStringMapOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStringMapOwn_Create", 0), // constructor NewStringMapOwn
		}
	})
	return cefStringMapOwnImport
}
