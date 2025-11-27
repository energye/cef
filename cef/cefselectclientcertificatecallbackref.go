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

// ICefSelectClientCertificateCallback Parent: ICefBaseRefCountedRef
type ICefSelectClientCertificateCallback interface {
	ICefBaseRefCountedRef
	// Select
	//  Chooses the specified certificate for client certificate authentication.
	//  NULL value means that no client certificate should be used.
	Select(cert ICefX509Certificate) // procedure
}

// ICefSelectClientCertificateCallbackRef Parent: ICefSelectClientCertificateCallback
type ICefSelectClientCertificateCallbackRef interface {
	ICefSelectClientCertificateCallback
	AsIntfSelectClientCertificateCallback() uintptr
}

type TCefSelectClientCertificateCallbackRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefSelectClientCertificateCallbackRef) Select(cert ICefX509Certificate) {
	if !m.IsValid() {
		return
	}
	cefSelectClientCertificateCallbackRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(cert))
}

func (m *TCefSelectClientCertificateCallbackRef) AsIntfSelectClientCertificateCallback() uintptr {
	return m.GetIntfPointer(0)
}

// SelectClientCertificateCallbackRef  is static instance
var SelectClientCertificateCallbackRef _SelectClientCertificateCallbackRefClass

// _SelectClientCertificateCallbackRefClass is class type defined by TCefSelectClientCertificateCallbackRef
type _SelectClientCertificateCallbackRefClass uintptr

func (_SelectClientCertificateCallbackRefClass) UnWrap(data uintptr) (result ICefSelectClientCertificateCallback) {
	var resultPtr uintptr
	cefSelectClientCertificateCallbackRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSelectClientCertificateCallbackRef(resultPtr)
	return
}

// NewSelectClientCertificateCallbackRef class constructor
func NewSelectClientCertificateCallbackRef(data uintptr) ICefSelectClientCertificateCallbackRef {
	var selectClientCertificateCallbackPtr uintptr // ICefSelectClientCertificateCallback
	r := cefSelectClientCertificateCallbackRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&selectClientCertificateCallbackPtr)))
	ret := AsCefSelectClientCertificateCallbackRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, selectClientCertificateCallbackPtr)
	}
	return ret
}

var (
	cefSelectClientCertificateCallbackRefOnce   base.Once
	cefSelectClientCertificateCallbackRefImport *imports.Imports = nil
)

func cefSelectClientCertificateCallbackRefAPI() *imports.Imports {
	cefSelectClientCertificateCallbackRefOnce.Do(func() {
		cefSelectClientCertificateCallbackRefImport = api.NewDefaultImports()
		cefSelectClientCertificateCallbackRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSelectClientCertificateCallbackRef_Create", 0), // constructor NewSelectClientCertificateCallbackRef
			/* 1 */ imports.NewTable("TCefSelectClientCertificateCallbackRef_UnWrap", 0), // static function UnWrap
			/* 2 */ imports.NewTable("TCefSelectClientCertificateCallbackRef_Select", 0), // procedure Select
		}
	})
	return cefSelectClientCertificateCallbackRefImport
}
