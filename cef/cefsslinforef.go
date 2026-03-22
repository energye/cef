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

// ICefSslInfo Parent: ICefBaseRefCounted
type ICefSslInfo interface {
	ICefBaseRefCounted
	// GetCertStatus
	//  Returns a bitmask containing any and all problems verifying the server
	//  certificate.
	GetCertStatus() cefTypes.TCefCertStatus // function
	// GetX509Certificate
	//  Returns the X.509 certificate.
	GetX509Certificate() ICefX509Certificate // function
}

// ICefSslInfoRef Parent: ICefSslInfo ICefBaseRefCountedRef
type ICefSslInfoRef interface {
	ICefSslInfo
	ICefBaseRefCountedRef
	AsIntfSslInfo() uintptr
}

type TCefSslInfoRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefSslInfoRef) GetCertStatus() cefTypes.TCefCertStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefSslInfoRefAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefCertStatus(r)
}

func (m *TCefSslInfoRef) GetX509Certificate() (result ICefX509Certificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefSslInfoRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCEFX509CertificateRef(resultPtr)
	return
}

func (m *TCefSslInfoRef) AsIntfSslInfo() uintptr {
	return m.GetIntfPointer(0)
}

// SslInfoRef  is static instance
var SslInfoRef _SslInfoRefClass

// _SslInfoRefClass is class type defined by TCefSslInfoRef
type _SslInfoRefClass uintptr

func (_SslInfoRefClass) UnWrap(data uintptr) (result ICefSslInfo) {
	var resultPtr uintptr
	cefSslInfoRefAPI().SysCallN(3, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSslInfoRef(resultPtr)
	return
}

// NewSslInfoRef class constructor
func NewSslInfoRef(data uintptr) ICefSslInfoRef {
	var sslInfoPtr uintptr // ICefSslInfo
	r := cefSslInfoRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&sslInfoPtr)))
	ret := AsCefSslInfoRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, sslInfoPtr)
	}
	return ret
}

var (
	cefSslInfoRefOnce   base.Once
	cefSslInfoRefImport *imports.Imports = nil
)

func cefSslInfoRefAPI() *imports.Imports {
	cefSslInfoRefOnce.Do(func() {
		cefSslInfoRefImport = api.NewDefaultImports()
		cefSslInfoRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSslInfoRef_Create", 0), // constructor NewSslInfoRef
			/* 1 */ imports.NewTable("TCefSslInfoRef_GetCertStatus", 0), // function GetCertStatus
			/* 2 */ imports.NewTable("TCefSslInfoRef_GetX509Certificate", 0), // function GetX509Certificate
			/* 3 */ imports.NewTable("TCefSslInfoRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefSslInfoRefImport
}
