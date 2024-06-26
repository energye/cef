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

// IV8Interceptor Parent: ICefV8Interceptor
//
//	Interface that should be implemented to handle V8 interceptor calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 interceptor. Interceptor's named property handlers (with first argument
//	of type CefString) are called when object is indexed by string. Indexed
//	property handlers (with first argument of type int) are called when object
//	is indexed by integer.
//	<a cref="uCEFTypes|TCefV8Interceptor">Implements TCefV8Interceptor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8interceptor_t)</a>
type IV8Interceptor interface {
	ICefV8Interceptor
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefV8Interceptor // procedure
	// SetOnGetByName
	//  Handle retrieval of the interceptor value identified by |name|. |object|
	//  is the receiver('this' object) of the interceptor. If retrieval succeeds,
	//  set |retval| to the return value. If the requested value does not exist,
	//  don't set either |retval| or |exception|. If retrieval fails, set
	//  |exception| to the exception that will be thrown. If the property has an
	//  associated accessor, it will be called only if you don't set |retval|.
	//  Return true(1) if interceptor retrieval was handled, false(0) otherwise.
	SetOnGetByName(fn TOnV8InterceptorGetByName) // property event
	// SetOnGetByIndex
	//  Handle retrieval of the interceptor value identified by |index|. |object|
	//  is the receiver('this' object) of the interceptor. If retrieval succeeds,
	//  set |retval| to the return value. If the requested value does not exist,
	//  don't set either |retval| or |exception|. If retrieval fails, set
	//  |exception| to the exception that will be thrown. Return true(1) if
	//  interceptor retrieval was handled, false(0) otherwise.
	SetOnGetByIndex(fn TOnV8InterceptorGetByIndex) // property event
	// SetOnSetByName
	//  Handle assignment of the interceptor value identified by |name|. |object|
	//  is the receiver('this' object) of the interceptor. |value| is the new
	//  value being assigned to the interceptor. If assignment fails, set
	//  |exception| to the exception that will be thrown. This setter will always
	//  be called, even when the property has an associated accessor. Return true
	//  (1) if interceptor assignment was handled, false(0) otherwise.
	SetOnSetByName(fn TOnV8InterceptorSetByName) // property event
	// SetOnSetByIndex
	//  Handle assignment of the interceptor value identified by |index|. |object|
	//  is the receiver('this' object) of the interceptor. |value| is the new
	//  value being assigned to the interceptor. If assignment fails, set
	//  |exception| to the exception that will be thrown. Return true(1) if
	//  interceptor assignment was handled, false(0) otherwise.
	SetOnSetByIndex(fn TOnV8InterceptorSetByIndex) // property event
}

// TV8Interceptor Parent: TCefV8Interceptor
//
//	Interface that should be implemented to handle V8 interceptor calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 interceptor. Interceptor's named property handlers (with first argument
//	of type CefString) are called when object is indexed by string. Indexed
//	property handlers (with first argument of type int) are called when object
//	is indexed by integer.
//	<a cref="uCEFTypes|TCefV8Interceptor">Implements TCefV8Interceptor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8interceptor_t)</a>
type TV8Interceptor struct {
	TCefV8Interceptor
	getByNamePtr  uintptr
	getByIndexPtr uintptr
	setByNamePtr  uintptr
	setByIndexPtr uintptr
}

func NewV8Interceptor() IV8Interceptor {
	r1 := v8InterceptorImportAPI().SysCallN(1)
	return AsV8Interceptor(r1)
}

func (m *TV8Interceptor) AsInterface() ICefV8Interceptor {
	var resultCefV8Interceptor uintptr
	v8InterceptorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefV8Interceptor)))
	return AsCefV8Interceptor(resultCefV8Interceptor)
}

func (m *TV8Interceptor) SetOnGetByName(fn TOnV8InterceptorGetByName) {
	if m.getByNamePtr != 0 {
		RemoveEventElement(m.getByNamePtr)
	}
	m.getByNamePtr = MakeEventDataPtr(fn)
	v8InterceptorImportAPI().SysCallN(3, m.Instance(), m.getByNamePtr)
}

func (m *TV8Interceptor) SetOnGetByIndex(fn TOnV8InterceptorGetByIndex) {
	if m.getByIndexPtr != 0 {
		RemoveEventElement(m.getByIndexPtr)
	}
	m.getByIndexPtr = MakeEventDataPtr(fn)
	v8InterceptorImportAPI().SysCallN(2, m.Instance(), m.getByIndexPtr)
}

func (m *TV8Interceptor) SetOnSetByName(fn TOnV8InterceptorSetByName) {
	if m.setByNamePtr != 0 {
		RemoveEventElement(m.setByNamePtr)
	}
	m.setByNamePtr = MakeEventDataPtr(fn)
	v8InterceptorImportAPI().SysCallN(5, m.Instance(), m.setByNamePtr)
}

func (m *TV8Interceptor) SetOnSetByIndex(fn TOnV8InterceptorSetByIndex) {
	if m.setByIndexPtr != 0 {
		RemoveEventElement(m.setByIndexPtr)
	}
	m.setByIndexPtr = MakeEventDataPtr(fn)
	v8InterceptorImportAPI().SysCallN(4, m.Instance(), m.setByIndexPtr)
}

var (
	v8InterceptorImport       *imports.Imports = nil
	v8InterceptorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("V8Interceptor_AsInterface", 0),
		/*1*/ imports.NewTable("V8Interceptor_Create", 0),
		/*2*/ imports.NewTable("V8Interceptor_SetOnGetByIndex", 0),
		/*3*/ imports.NewTable("V8Interceptor_SetOnGetByName", 0),
		/*4*/ imports.NewTable("V8Interceptor_SetOnSetByIndex", 0),
		/*5*/ imports.NewTable("V8Interceptor_SetOnSetByName", 0),
	}
)

func v8InterceptorImportAPI() *imports.Imports {
	if v8InterceptorImport == nil {
		v8InterceptorImport = NewDefaultImports()
		v8InterceptorImport.SetImportTable(v8InterceptorImportTables)
		v8InterceptorImportTables = nil
	}
	return v8InterceptorImport
}
