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

// ICefResourceHandler Parent: ICefBaseRefCounted
type ICefResourceHandler interface {
	ICefBaseRefCounted
}

// ICefResourceHandlerRef Parent: ICefResourceHandler ICefBaseRefCountedRef
type ICefResourceHandlerRef interface {
	ICefResourceHandler
	ICefBaseRefCountedRef
	AsIntfResourceHandler() uintptr
}

type TCefResourceHandlerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResourceHandlerRef) AsIntfResourceHandler() uintptr {
	return m.GetIntfPointer(0)
}

// ResourceHandlerRef  is static instance
var ResourceHandlerRef _ResourceHandlerRefClass

// _ResourceHandlerRefClass is class type defined by TCefResourceHandlerRef
type _ResourceHandlerRefClass uintptr

func (_ResourceHandlerRefClass) UnWrap(data uintptr) (result IEngResourceHandler) {
	var resultPtr uintptr
	cefResourceHandlerRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngResourceHandler(resultPtr)
	return
}

// NewResourceHandlerRef class constructor
func NewResourceHandlerRef(data uintptr) ICefResourceHandlerRef {
	var resourceHandlerPtr uintptr // ICefResourceHandler
	r := cefResourceHandlerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&resourceHandlerPtr)))
	ret := AsCefResourceHandlerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceHandlerPtr)
	}
	return ret
}

var (
	cefResourceHandlerRefOnce   base.Once
	cefResourceHandlerRefImport *imports.Imports = nil
)

func cefResourceHandlerRefAPI() *imports.Imports {
	cefResourceHandlerRefOnce.Do(func() {
		cefResourceHandlerRefImport = api.NewDefaultImports()
		cefResourceHandlerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResourceHandlerRef_Create", 0), // constructor NewResourceHandlerRef
			/* 1 */ imports.NewTable("TCefResourceHandlerRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefResourceHandlerRefImport
}
