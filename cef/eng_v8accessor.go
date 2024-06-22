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

// IV8Accessor Parent: ICefV8Accessor
//
//	Interface that should be implemented to handle V8 accessor calls. Accessor
//	identifiers are registered by calling ICefV8value.SetValue(). The
//	functions of this interface will be called on the thread associated with the
//	V8 accessor.
//	<a cref="uCEFTypes|TCefV8Accessor">Implements TCefV8Accessor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8accessor_t)</a>
type IV8Accessor interface {
	ICefV8Accessor
	// SetOnGet
	//  Handle retrieval the accessor value identified by |name|. |object| is the
	//  receiver('this' object) of the accessor. If retrieval succeeds set
	//  |retval| to the return value. If retrieval fails set |exception| to the
	//  exception that will be thrown. Return true(1) if accessor retrieval was
	//  handled.
	SetOnGet(fn TOnV8AccessorGet) // property event
	// SetOnSet
	//  Handle assignment of the accessor value identified by |name|. |object| is
	//  the receiver('this' object) of the accessor. |value| is the new value
	//  being assigned to the accessor. If assignment fails set |exception| to the
	//  exception that will be thrown. Return true(1) if accessor assignment was
	//  handled.
	SetOnSet(fn TOnV8AccessorSet) // property event
}

// TV8Accessor Parent: TCefV8Accessor
//
//	Interface that should be implemented to handle V8 accessor calls. Accessor
//	identifiers are registered by calling ICefV8value.SetValue(). The
//	functions of this interface will be called on the thread associated with the
//	V8 accessor.
//	<a cref="uCEFTypes|TCefV8Accessor">Implements TCefV8Accessor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8accessor_t)</a>
type TV8Accessor struct {
	TCefV8Accessor
	getPtr uintptr
	setPtr uintptr
}

func NewV8Accessor() IV8Accessor {
	r1 := v8AccessorImportAPI().SysCallN(0)
	return AsV8Accessor(r1)
}

func (m *TV8Accessor) SetOnGet(fn TOnV8AccessorGet) {
	if m.getPtr != 0 {
		RemoveEventElement(m.getPtr)
	}
	m.getPtr = MakeEventDataPtr(fn)
	v8AccessorImportAPI().SysCallN(1, m.Instance(), m.getPtr)
}

func (m *TV8Accessor) SetOnSet(fn TOnV8AccessorSet) {
	if m.setPtr != 0 {
		RemoveEventElement(m.setPtr)
	}
	m.setPtr = MakeEventDataPtr(fn)
	v8AccessorImportAPI().SysCallN(2, m.Instance(), m.setPtr)
}

var (
	v8AccessorImport       *imports.Imports = nil
	v8AccessorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("V8Accessor_Create", 0),
		/*1*/ imports.NewTable("V8Accessor_SetOnGet", 0),
		/*2*/ imports.NewTable("V8Accessor_SetOnSet", 0),
	}
)

func v8AccessorImportAPI() *imports.Imports {
	if v8AccessorImport == nil {
		v8AccessorImport = NewDefaultImports()
		v8AccessorImport.SetImportTable(v8AccessorImportTables)
		v8AccessorImportTables = nil
	}
	return v8AccessorImport
}
