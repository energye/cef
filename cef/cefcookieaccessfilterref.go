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

// ICefCookieAccessFilterRef Parent: ICefBaseRefCountedRef
type ICefCookieAccessFilterRef interface {
	ICefBaseRefCountedRef
	AsIntfCookieAccessFilter() uintptr
}

type TCefCookieAccessFilterRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefCookieAccessFilterRef) AsIntfCookieAccessFilter() uintptr {
	return m.GetIntfPointer(0)
}

// CookieAccessFilterRef  is static instance
var CookieAccessFilterRef _CookieAccessFilterRefClass

// _CookieAccessFilterRefClass is class type defined by TCefCookieAccessFilterRef
type _CookieAccessFilterRefClass uintptr

func (_CookieAccessFilterRefClass) UnWrap(data uintptr) (result IEngCookieAccessFilter) {
	var resultPtr uintptr
	cefCookieAccessFilterRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngCookieAccessFilter(resultPtr)
	return
}

// NewCookieAccessFilterRef class constructor
func NewCookieAccessFilterRef(data uintptr) ICefCookieAccessFilterRef {
	var cookieAccessFilterPtr uintptr // ICefCookieAccessFilter
	r := cefCookieAccessFilterRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&cookieAccessFilterPtr)))
	ret := AsCefCookieAccessFilterRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieAccessFilterPtr)
	}
	return ret
}

var (
	cefCookieAccessFilterRefOnce   base.Once
	cefCookieAccessFilterRefImport *imports.Imports = nil
)

func cefCookieAccessFilterRefAPI() *imports.Imports {
	cefCookieAccessFilterRefOnce.Do(func() {
		cefCookieAccessFilterRefImport = api.NewDefaultImports()
		cefCookieAccessFilterRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCookieAccessFilterRef_Create", 0), // constructor NewCookieAccessFilterRef
			/* 1 */ imports.NewTable("TCefCookieAccessFilterRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefCookieAccessFilterRefImport
}
