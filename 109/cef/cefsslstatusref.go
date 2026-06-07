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

// ICefSSLStatus Parent: ICefBaseRefCounted
type ICefSSLStatus interface {
	ICefBaseRefCounted
	IsSecureConnection() bool                        // function
	GetCertStatus() cefTypes.TCefCertStatus          // function
	GetSSLVersion() cefTypes.TCefSSLVersion          // function
	GetContentStatus() cefTypes.TCefSSLContentStatus // function
	GetX509Certificate() ICefX509Certificate         // function
}

// ICefSSLStatusRef Parent: ICefSSLStatus ICefBaseRefCountedRef
type ICefSSLStatusRef interface {
	ICefSSLStatus
	ICefBaseRefCountedRef
	AsIntfSSLStatus() uintptr
}

type TCefSSLStatusRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefSSLStatusRef) IsSecureConnection() bool {
	if !m.IsValid() {
		return false
	}
	r := cefSSLStatusRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefSSLStatusRef) GetCertStatus() cefTypes.TCefCertStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefSSLStatusRefAPI().SysCallN(2, m.Instance())
	return cefTypes.TCefCertStatus(r)
}

func (m *TCefSSLStatusRef) GetSSLVersion() cefTypes.TCefSSLVersion {
	if !m.IsValid() {
		return 0
	}
	r := cefSSLStatusRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefSSLVersion(r)
}

func (m *TCefSSLStatusRef) GetContentStatus() cefTypes.TCefSSLContentStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefSSLStatusRefAPI().SysCallN(4, m.Instance())
	return cefTypes.TCefSSLContentStatus(r)
}

func (m *TCefSSLStatusRef) GetX509Certificate() (result ICefX509Certificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefSSLStatusRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCEFX509CertificateRef(resultPtr)
	return
}

func (m *TCefSSLStatusRef) AsIntfSSLStatus() uintptr {
	return m.GetIntfPointer(0)
}

// SSLStatusRef  is static instance
var SSLStatusRef _SSLStatusRefClass

// _SSLStatusRefClass is class type defined by TCefSSLStatusRef
type _SSLStatusRefClass uintptr

func (_SSLStatusRefClass) UnWrap(data uintptr) (result ICefSSLStatus) {
	var resultPtr uintptr
	cefSSLStatusRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSSLStatusRef(resultPtr)
	return
}

// NewSSLStatusRef class constructor
func NewSSLStatusRef(data uintptr) ICefSSLStatusRef {
	var sSLStatusPtr uintptr // ICefSSLStatus
	r := cefSSLStatusRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&sSLStatusPtr)))
	ret := AsCefSSLStatusRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, sSLStatusPtr)
	}
	return ret
}

var (
	cefSSLStatusRefOnce   base.Once
	cefSSLStatusRefImport *imports.Imports = nil
)

func cefSSLStatusRefAPI() *imports.Imports {
	cefSSLStatusRefOnce.Do(func() {
		cefSSLStatusRefImport = api.NewDefaultImports()
		cefSSLStatusRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSSLStatusRef_Create", 0), // constructor NewSSLStatusRef
			/* 1 */ imports.NewTable("TCefSSLStatusRef_IsSecureConnection", 0), // function IsSecureConnection
			/* 2 */ imports.NewTable("TCefSSLStatusRef_GetCertStatus", 0), // function GetCertStatus
			/* 3 */ imports.NewTable("TCefSSLStatusRef_GetSSLVersion", 0), // function GetSSLVersion
			/* 4 */ imports.NewTable("TCefSSLStatusRef_GetContentStatus", 0), // function GetContentStatus
			/* 5 */ imports.NewTable("TCefSSLStatusRef_GetX509Certificate", 0), // function GetX509Certificate
			/* 6 */ imports.NewTable("TCefSSLStatusRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefSSLStatusRefImport
}
