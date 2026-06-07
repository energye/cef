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

// ICefUrlrequestClient Parent: ICefBaseRefCounted
type ICefUrlrequestClient interface {
	ICefBaseRefCounted
}

// ICefUrlrequestClientRef Parent: ICefUrlrequestClient ICefBaseRefCountedRef
type ICefUrlrequestClientRef interface {
	ICefUrlrequestClient
	ICefBaseRefCountedRef
	AsIntfUrlrequestClient() uintptr
}

type TCefUrlrequestClientRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefUrlrequestClientRef) AsIntfUrlrequestClient() uintptr {
	return m.GetIntfPointer(0)
}

// UrlrequestClientRef  is static instance
var UrlrequestClientRef _UrlrequestClientRefClass

// _UrlrequestClientRefClass is class type defined by TCefUrlrequestClientRef
type _UrlrequestClientRefClass uintptr

func (_UrlrequestClientRefClass) UnWrap(data uintptr) (result IEngUrlrequestClient) {
	var resultPtr uintptr
	cefUrlrequestClientRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngUrlrequestClient(resultPtr)
	return
}

// NewUrlrequestClientRef class constructor
func NewUrlrequestClientRef(data uintptr) ICefUrlrequestClientRef {
	var urlrequestClientPtr uintptr // ICefUrlrequestClient
	r := cefUrlrequestClientRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&urlrequestClientPtr)))
	ret := AsCefUrlrequestClientRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, urlrequestClientPtr)
	}
	return ret
}

var (
	cefUrlrequestClientRefOnce   base.Once
	cefUrlrequestClientRefImport *imports.Imports = nil
)

func cefUrlrequestClientRefAPI() *imports.Imports {
	cefUrlrequestClientRefOnce.Do(func() {
		cefUrlrequestClientRefImport = api.NewDefaultImports()
		cefUrlrequestClientRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefUrlrequestClientRef_Create", 0), // constructor NewUrlrequestClientRef
			/* 1 */ imports.NewTable("TCefUrlrequestClientRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefUrlrequestClientRefImport
}
