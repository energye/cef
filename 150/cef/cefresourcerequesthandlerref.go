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

// ICefResourceRequestHandler Parent: ICefBaseRefCounted
type ICefResourceRequestHandler interface {
	ICefBaseRefCounted
}

// ICefResourceRequestHandlerRef Parent: ICefResourceRequestHandler ICefBaseRefCountedRef
type ICefResourceRequestHandlerRef interface {
	ICefResourceRequestHandler
	ICefBaseRefCountedRef
	AsIntfResourceRequestHandler() uintptr
}

type TCefResourceRequestHandlerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResourceRequestHandlerRef) AsIntfResourceRequestHandler() uintptr {
	return m.GetIntfPointer(0)
}

// ResourceRequestHandlerRef  is static instance
var ResourceRequestHandlerRef _ResourceRequestHandlerRefClass

// _ResourceRequestHandlerRefClass is class type defined by TCefResourceRequestHandlerRef
type _ResourceRequestHandlerRefClass uintptr

func (_ResourceRequestHandlerRefClass) UnWrap(data uintptr) (result IEngResourceRequestHandler) {
	var resultPtr uintptr
	cefResourceRequestHandlerRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngResourceRequestHandler(resultPtr)
	return
}

// NewResourceRequestHandlerRef class constructor
func NewResourceRequestHandlerRef(data uintptr) ICefResourceRequestHandlerRef {
	var resourceRequestHandlerPtr uintptr // ICefResourceRequestHandler
	r := cefResourceRequestHandlerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&resourceRequestHandlerPtr)))
	ret := AsCefResourceRequestHandlerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceRequestHandlerPtr)
	}
	return ret
}

var (
	cefResourceRequestHandlerRefOnce   base.Once
	cefResourceRequestHandlerRefImport *imports.Imports = nil
)

func cefResourceRequestHandlerRefAPI() *imports.Imports {
	cefResourceRequestHandlerRefOnce.Do(func() {
		cefResourceRequestHandlerRefImport = api.NewDefaultImports()
		cefResourceRequestHandlerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResourceRequestHandlerRef_Create", 0), // constructor NewResourceRequestHandlerRef
			/* 1 */ imports.NewTable("TCefResourceRequestHandlerRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefResourceRequestHandlerRefImport
}
