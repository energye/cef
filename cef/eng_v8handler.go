//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IV8Handler Parent: ICefV8Handler
//
//	Interface that should be implemented to handle V8 function calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 function.
//	<a cref="uCEFTypes|TCefv8Handler">Implements TCefv8Handler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t)</a>
type IV8Handler interface {
	ICefV8Handler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefV8Handler // procedure
	// SetOnExecute
	//  Handle execution of the function identified by |name|. |object| is the
	//  receiver('this' object) of the function. |arguments| is the list of
	//  arguments passed to the function. If execution succeeds set |retval| to
	//  the function return value. If execution fails set |exception| to the
	//  exception that will be thrown. Return true(1) if execution was handled.
	SetOnExecute(fn TOnV8HandlerExecute) // property event
}

// TV8Handler Parent: TCefV8Handler
//
//	Interface that should be implemented to handle V8 function calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 function.
//	<a cref="uCEFTypes|TCefv8Handler">Implements TCefv8Handler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t)</a>
type TV8Handler struct {
	TCefV8Handler
	executePtr uintptr
}

func NewV8Handler() IV8Handler {
	r1 := v8HandlerImportAPI().SysCallN(1)
	return AsV8Handler(r1)
}

func (m *TV8Handler) AsInterface() ICefV8Handler {
	var resultCefV8Handler uintptr
	v8HandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefV8Handler)))
	return AsCefV8Handler(resultCefV8Handler)
}

func (m *TV8Handler) SetOnExecute(fn TOnV8HandlerExecute) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	v8HandlerImportAPI().SysCallN(2, m.Instance(), m.executePtr)
}

var (
	v8HandlerImport       *imports.Imports = nil
	v8HandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("V8Handler_AsInterface", 0),
		/*1*/ imports.NewTable("V8Handler_Create", 0),
		/*2*/ imports.NewTable("V8Handler_SetOnExecute", 0),
	}
)

func v8HandlerImportAPI() *imports.Imports {
	if v8HandlerImport == nil {
		v8HandlerImport = NewDefaultImports()
		v8HandlerImport.SetImportTable(v8HandlerImportTables)
		v8HandlerImportTables = nil
	}
	return v8HandlerImport
}
