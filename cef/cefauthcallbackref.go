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

// ICefAuthCallback Parent: ICefBaseRefCountedRef
type ICefAuthCallback interface {
	ICefBaseRefCountedRef
	// Cont
	//  Continue the authentication request.
	Cont(username string, password string) // procedure
	// Cancel
	//  Cancel the authentication request.
	Cancel() // procedure
}

// ICefAuthCallbackRef Parent: ICefAuthCallback
type ICefAuthCallbackRef interface {
	ICefAuthCallback
	AsIntfAuthCallback() uintptr
}

type TCefAuthCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefAuthCallbackRef) Cont(username string, password string) {
	if !m.IsValid() {
		return
	}
	cefAuthCallbackRefAPI().SysCallN(2, m.Instance(), api.PasStr(username), api.PasStr(password))
}

func (m *TCefAuthCallbackRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefAuthCallbackRefAPI().SysCallN(3, m.Instance())
}

func (m *TCefAuthCallbackRef) AsIntfAuthCallback() uintptr {
	return m.GetIntfPointer(0)
}

// AuthCallbackRef  is static instance
var AuthCallbackRef _AuthCallbackRefClass

// _AuthCallbackRefClass is class type defined by TCefAuthCallbackRef
type _AuthCallbackRefClass uintptr

func (_AuthCallbackRefClass) UnWrap(data uintptr) (result ICefAuthCallback) {
	var resultPtr uintptr
	cefAuthCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefAuthCallbackRef(resultPtr)
	return
}

// NewAuthCallbackRef class constructor
func NewAuthCallbackRef(data uintptr) ICefAuthCallbackRef {
	var authCallbackPtr uintptr // ICefAuthCallback
	r := cefAuthCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&authCallbackPtr)))
	ret := AsCefAuthCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, authCallbackPtr)
	}
	return ret
}

var (
	cefAuthCallbackRefOnce   base.Once
	cefAuthCallbackRefImport *imports.Imports = nil
)

func cefAuthCallbackRefAPI() *imports.Imports {
	cefAuthCallbackRefOnce.Do(func() {
		cefAuthCallbackRefImport = api.NewDefaultImports()
		cefAuthCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefAuthCallbackRef_Create", 0), // constructor NewAuthCallbackRef
			/* 1 */ imports.NewTable("TCefAuthCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefAuthCallbackRef_Cont", 0), // procedure Cont
			/* 3 */ imports.NewTable("TCefAuthCallbackRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefAuthCallbackRefImport
}
