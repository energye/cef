//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefPermissionPromptCallback Parent: ICefBaseRefCounted
type ICefPermissionPromptCallback interface {
	ICefBaseRefCounted
	Cont(result cefTypes.TCefPermissionRequestResult) // procedure
}

// ICefPermissionPromptCallbackRef Parent: ICefPermissionPromptCallback ICefBaseRefCountedRef
type ICefPermissionPromptCallbackRef interface {
	ICefPermissionPromptCallback
	ICefBaseRefCountedRef
	AsIntfPermissionPromptCallback() uintptr
}

type TCefPermissionPromptCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPermissionPromptCallbackRef) Cont(result cefTypes.TCefPermissionRequestResult) {
	if !m.IsValid() {
		return
	}
	cefPermissionPromptCallbackRefAPI().SysCallN(2, m.Instance(), uintptr(result))
}

func (m *TCefPermissionPromptCallbackRef) AsIntfPermissionPromptCallback() uintptr {
	return m.GetIntfPointer(0)
}

// PermissionPromptCallbackRef  is static instance
var PermissionPromptCallbackRef _PermissionPromptCallbackRefClass

// _PermissionPromptCallbackRefClass is class type defined by TCefPermissionPromptCallbackRef
type _PermissionPromptCallbackRefClass uintptr

func (_PermissionPromptCallbackRefClass) UnWrap(data uintptr) (result ICefPermissionPromptCallback) {
	var resultPtr uintptr
	cefPermissionPromptCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPermissionPromptCallbackRef(resultPtr)
	return
}

// NewPermissionPromptCallbackRef class constructor
func NewPermissionPromptCallbackRef(data uintptr) ICefPermissionPromptCallbackRef {
	var permissionPromptCallbackPtr uintptr // ICefPermissionPromptCallback
	r := cefPermissionPromptCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&permissionPromptCallbackPtr)))
	ret := AsCefPermissionPromptCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, permissionPromptCallbackPtr)
	}
	return ret
}

var (
	cefPermissionPromptCallbackRefOnce   base.Once
	cefPermissionPromptCallbackRefImport *imports.Imports = nil
)

func cefPermissionPromptCallbackRefAPI() *imports.Imports {
	cefPermissionPromptCallbackRefOnce.Do(func() {
		cefPermissionPromptCallbackRefImport = api.NewDefaultImports()
		cefPermissionPromptCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPermissionPromptCallbackRef_Create", 0), // constructor NewPermissionPromptCallbackRef
			/* 1 */ imports.NewTable("TCefPermissionPromptCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefPermissionPromptCallbackRef_cont", 0), // procedure Cont
		}
	})
	return cefPermissionPromptCallbackRefImport
}
