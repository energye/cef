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

// ICefJsDialogCallback Parent: ICefBaseRefCountedRef
type ICefJsDialogCallback interface {
	ICefBaseRefCountedRef
	// Cont
	//  Continue the JS dialog request. Set |success| to true (1) if the OK button
	//  was pressed. The |user_input| value should be specified for prompt
	//  dialogs.
	Cont(success bool, userInput string) // procedure
}

// ICefJsDialogCallbackRef Parent: ICefJsDialogCallback
type ICefJsDialogCallbackRef interface {
	ICefJsDialogCallback
	AsIntfJsDialogCallback() uintptr
}

type TCefJsDialogCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefJsDialogCallbackRef) Cont(success bool, userInput string) {
	if !m.IsValid() {
		return
	}
	cefJsDialogCallbackRefAPI().SysCallN(2, m.Instance(), api.PasBool(success), api.PasStr(userInput))
}

func (m *TCefJsDialogCallbackRef) AsIntfJsDialogCallback() uintptr {
	return m.GetIntfPointer(0)
}

// JsDialogCallbackRef  is static instance
var JsDialogCallbackRef _JsDialogCallbackRefClass

// _JsDialogCallbackRefClass is class type defined by TCefJsDialogCallbackRef
type _JsDialogCallbackRefClass uintptr

func (_JsDialogCallbackRefClass) UnWrap(data uintptr) (result ICefJsDialogCallback) {
	var resultPtr uintptr
	cefJsDialogCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefJsDialogCallbackRef(resultPtr)
	return
}

// NewJsDialogCallbackRef class constructor
func NewJsDialogCallbackRef(data uintptr) ICefJsDialogCallbackRef {
	var jsDialogCallbackPtr uintptr // ICefJsDialogCallback
	r := cefJsDialogCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&jsDialogCallbackPtr)))
	ret := AsCefJsDialogCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, jsDialogCallbackPtr)
	}
	return ret
}

var (
	cefJsDialogCallbackRefOnce   base.Once
	cefJsDialogCallbackRefImport *imports.Imports = nil
)

func cefJsDialogCallbackRefAPI() *imports.Imports {
	cefJsDialogCallbackRefOnce.Do(func() {
		cefJsDialogCallbackRefImport = api.NewDefaultImports()
		cefJsDialogCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefJsDialogCallbackRef_Create", 0), // constructor NewJsDialogCallbackRef
			/* 1 */ imports.NewTable("TCefJsDialogCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefJsDialogCallbackRef_Cont", 0), // procedure Cont
		}
	})
	return cefJsDialogCallbackRefImport
}
