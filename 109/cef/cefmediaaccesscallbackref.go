//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefMediaAccessCallback Parent: ICefBaseRefCounted
type ICefMediaAccessCallback interface {
	ICefBaseRefCounted
	Cont(allowedPermissions cefTypes.TCefMediaAccessPermissionTypes) // procedure
	Cancel()                                                         // procedure
}

// ICefMediaAccessCallbackRef Parent: ICefMediaAccessCallback ICefBaseRefCountedRef
type ICefMediaAccessCallbackRef interface {
	ICefMediaAccessCallback
	ICefBaseRefCountedRef
	AsIntfMediaAccessCallback() uintptr
}

type TCefMediaAccessCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMediaAccessCallbackRef) Cont(allowedPermissions cefTypes.TCefMediaAccessPermissionTypes) {
	if !m.IsValid() {
		return
	}
	cefMediaAccessCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(allowedPermissions))
}

func (m *TCefMediaAccessCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefMediaAccessCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefMediaAccessCallbackRef) AsIntfMediaAccessCallback() uintptr {
	return m.GetIntfPointer(0)
}

// MediaAccessCallbackRef  is static instance
var MediaAccessCallbackRef _MediaAccessCallbackRefClass

// _MediaAccessCallbackRefClass is class type defined by TCefMediaAccessCallbackRef
type _MediaAccessCallbackRefClass uintptr

func (_MediaAccessCallbackRefClass) UnWrap(data uintptr) (result ICefMediaAccessCallback) {
	var resultPtr uintptr
	cefMediaAccessCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaAccessCallbackRef(resultPtr)
	return
}

// NewMediaAccessCallbackRef class constructor
func NewMediaAccessCallbackRef(data uintptr) ICefMediaAccessCallbackRef {
	var mediaAccessCallbackPtr uintptr // ICefMediaAccessCallback
	r := cefMediaAccessCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&mediaAccessCallbackPtr)))
	ret := AsCefMediaAccessCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaAccessCallbackPtr)
	}
	return ret
}

var (
	cefMediaAccessCallbackRefOnce   base.Once
	cefMediaAccessCallbackRefImport *imports.Imports = nil
)

func cefMediaAccessCallbackRefAPI() *imports.Imports {
	cefMediaAccessCallbackRefOnce.Do(func() {
		cefMediaAccessCallbackRefImport = api.NewDefaultImports()
		cefMediaAccessCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMediaAccessCallbackRef_Create", 0), // constructor NewMediaAccessCallbackRef
			/* 1 */ imports.NewTable("TCefMediaAccessCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefMediaAccessCallbackRef_cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefMediaAccessCallbackRef_cancel", 0), // procedure Cancel
		}
	})
	return cefMediaAccessCallbackRefImport
}
