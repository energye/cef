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

// ICefResponseFilterRef Parent: ICefBaseRefCountedRef
type ICefResponseFilterRef interface {
	ICefBaseRefCountedRef
	AsIntfResponseFilter() uintptr
}

type TCefResponseFilterRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResponseFilterRef) AsIntfResponseFilter() uintptr {
	return m.GetIntfPointer(0)
}

// ResponseFilterRef  is static instance
var ResponseFilterRef _ResponseFilterRefClass

// _ResponseFilterRefClass is class type defined by TCefResponseFilterRef
type _ResponseFilterRefClass uintptr

func (_ResponseFilterRefClass) UnWrap(data uintptr) (result IEngResponseFilter) {
	var resultPtr uintptr
	cefResponseFilterRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngResponseFilter(resultPtr)
	return
}

// NewResponseFilterRef class constructor
func NewResponseFilterRef(data uintptr) ICefResponseFilterRef {
	var responseFilterPtr uintptr // ICefResponseFilter
	r := cefResponseFilterRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&responseFilterPtr)))
	ret := AsCefResponseFilterRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, responseFilterPtr)
	}
	return ret
}

var (
	cefResponseFilterRefOnce   base.Once
	cefResponseFilterRefImport *imports.Imports = nil
)

func cefResponseFilterRefAPI() *imports.Imports {
	cefResponseFilterRefOnce.Do(func() {
		cefResponseFilterRefImport = api.NewDefaultImports()
		cefResponseFilterRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResponseFilterRef_Create", 0), // constructor NewResponseFilterRef
			/* 1 */ imports.NewTable("TCefResponseFilterRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefResponseFilterRefImport
}
