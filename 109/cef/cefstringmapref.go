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

	cefTypes "github.com/energye/cef/109/types"
)

// ICefStringMapRef Parent: ICefCustomStringMap
type ICefStringMapRef interface {
	ICefCustomStringMap
	AsIntfStringMap() uintptr
}

type TCefStringMapRef struct {
	TCefCustomStringMap
}

func (m *TCefStringMapRef) AsIntfStringMap() uintptr {
	return m.GetIntfPointer(0)
}

// NewStringMapRef class constructor
func NewStringMapRef(handle cefTypes.TCefStringMap) ICefStringMapRef {
	var stringMapPtr uintptr // ICefStringMap
	r := cefStringMapRefAPI().SysCallN(0, uintptr(handle), uintptr(base.UnsafePointer(&stringMapPtr)))
	ret := AsCefStringMapRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringMapPtr)
	}
	return ret
}

var (
	cefStringMapRefOnce   base.Once
	cefStringMapRefImport *imports.Imports = nil
)

func cefStringMapRefAPI() *imports.Imports {
	cefStringMapRefOnce.Do(func() {
		cefStringMapRefImport = api.NewDefaultImports()
		cefStringMapRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefStringMapRef_Create", 0), // constructor NewStringMapRef
		}
	})
	return cefStringMapRefImport
}
