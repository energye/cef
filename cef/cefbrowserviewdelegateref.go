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

// ICefBrowserViewDelegate Parent: ICefViewDelegate
type ICefBrowserViewDelegate interface {
	ICefViewDelegate
}

// ICefBrowserViewDelegateRef Parent: ICefBrowserViewDelegate ICefViewDelegateRef
type ICefBrowserViewDelegateRef interface {
	ICefBrowserViewDelegate
	ICefViewDelegateRef
	AsIntfBrowserViewDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TCefBrowserViewDelegateRef struct {
	TCefViewDelegateRef
}

func (m *TCefBrowserViewDelegateRef) AsIntfBrowserViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefBrowserViewDelegateRef) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// BrowserViewDelegateRef  is static instance
var BrowserViewDelegateRef _BrowserViewDelegateRefClass

// _BrowserViewDelegateRefClass is class type defined by TCefBrowserViewDelegateRef
type _BrowserViewDelegateRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefBrowserViewDelegate instance using a PCefBrowserViewDelegate data pointer.
func (_BrowserViewDelegateRefClass) UnWrapWithPointer(data uintptr) (result IEngBrowserViewDelegate) {
	var resultPtr uintptr
	cefBrowserViewDelegateRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngBrowserViewDelegate(resultPtr)
	return
}

// NewBrowserViewDelegateRef class constructor
func NewBrowserViewDelegateRef(data uintptr) ICefBrowserViewDelegateRef {
	var browserViewDelegatePtr uintptr // ICefBrowserViewDelegate
	var viewDelegatePtr uintptr        // ICefViewDelegate
	r := cefBrowserViewDelegateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&browserViewDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsCefBrowserViewDelegateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	cefBrowserViewDelegateRefOnce   base.Once
	cefBrowserViewDelegateRefImport *imports.Imports = nil
)

func cefBrowserViewDelegateRefAPI() *imports.Imports {
	cefBrowserViewDelegateRefOnce.Do(func() {
		cefBrowserViewDelegateRefImport = api.NewDefaultImports()
		cefBrowserViewDelegateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserViewDelegateRef_Create", 0), // constructor NewBrowserViewDelegateRef
			/* 1 */ imports.NewTable("TCefBrowserViewDelegateRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
		}
	})
	return cefBrowserViewDelegateRefImport
}
