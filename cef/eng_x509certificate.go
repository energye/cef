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

// ICefX509Certificate Parent: ICefBaseRefCounted
type ICefX509Certificate interface {
	ICefBaseRefCounted
	GetDEREncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray)
	GetPEMEncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray)
	GetSubject() ICefX509CertPrincipal                    // function
	GetIssuer() ICefX509CertPrincipal                     // function
	GetSerialNumber() ICefBinaryValue                     // function
	GetValidStart() TCefBaseTime                          // function
	GetValidExpiry() TCefBaseTime                         // function
	GetValidStartAsDateTime() (resultDateTime TDateTime)  // function
	GetValidExpiryAsDateTime() (resultDateTime TDateTime) // function
	GetDerEncoded() ICefBinaryValue                       // function
	GetPemEncoded() ICefBinaryValue                       // function
	GetIssuerChainSize() NativeUInt                       // function
}

// TCefX509Certificate Parent: TCefBaseRefCounted
type TCefX509Certificate struct {
	TCefBaseRefCounted
}

// X509CertificateRef -> ICefX509Certificate
var X509CertificateRef x509Certificate

// x509Certificate TCefX509Certificate Ref
type x509Certificate uintptr

func (m *x509Certificate) UnWrap(data uintptr) ICefX509Certificate {
	var resultCefX509Certificate uintptr
	x509CertificateImportAPI().SysCallN(10, uintptr(data), uintptr(unsafePointer(&resultCefX509Certificate)))
	return AsCefX509Certificate(resultCefX509Certificate)
}

func (m *TCefX509Certificate) GetSubject() ICefX509CertPrincipal {
	var resultCefX509CertPrincipal uintptr
	x509CertificateImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefX509CertPrincipal)))
	return AsCefX509CertPrincipal(resultCefX509CertPrincipal)
}

func (m *TCefX509Certificate) GetIssuer() ICefX509CertPrincipal {
	var resultCefX509CertPrincipal uintptr
	x509CertificateImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefX509CertPrincipal)))
	return AsCefX509CertPrincipal(resultCefX509CertPrincipal)
}

func (m *TCefX509Certificate) GetSerialNumber() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	x509CertificateImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefX509Certificate) GetValidStart() TCefBaseTime {
	r1 := x509CertificateImportAPI().SysCallN(8, m.Instance())
	return TCefBaseTime(r1)
}

func (m *TCefX509Certificate) GetValidExpiry() TCefBaseTime {
	r1 := x509CertificateImportAPI().SysCallN(6, m.Instance())
	return TCefBaseTime(r1)
}

func (m *TCefX509Certificate) GetValidStartAsDateTime() (resultDateTime TDateTime) {
	x509CertificateImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefX509Certificate) GetValidExpiryAsDateTime() (resultDateTime TDateTime) {
	x509CertificateImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefX509Certificate) GetDerEncoded() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	x509CertificateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefX509Certificate) GetPemEncoded() ICefBinaryValue {
	var resultCefBinaryValue uintptr
	x509CertificateImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefBinaryValue)))
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefX509Certificate) GetIssuerChainSize() NativeUInt {
	r1 := x509CertificateImportAPI().SysCallN(2, m.Instance())
	return NativeUInt(r1)
}

var (
	x509CertificateImport       *imports.Imports = nil
	x509CertificateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefX509Certificate_GetDerEncoded", 0),
		/*1*/ imports.NewTable("CefX509Certificate_GetIssuer", 0),
		/*2*/ imports.NewTable("CefX509Certificate_GetIssuerChainSize", 0),
		/*3*/ imports.NewTable("CefX509Certificate_GetPemEncoded", 0),
		/*4*/ imports.NewTable("CefX509Certificate_GetSerialNumber", 0),
		/*5*/ imports.NewTable("CefX509Certificate_GetSubject", 0),
		/*6*/ imports.NewTable("CefX509Certificate_GetValidExpiry", 0),
		/*7*/ imports.NewTable("CefX509Certificate_GetValidExpiryAsDateTime", 0),
		/*8*/ imports.NewTable("CefX509Certificate_GetValidStart", 0),
		/*9*/ imports.NewTable("CefX509Certificate_GetValidStartAsDateTime", 0),
		/*10*/ imports.NewTable("CefX509Certificate_UnWrap", 0),
	}
)

func x509CertificateImportAPI() *imports.Imports {
	if x509CertificateImport == nil {
		x509CertificateImport = NewDefaultImports()
		x509CertificateImport.SetImportTable(x509CertificateImportTables)
		x509CertificateImportTables = nil
	}
	return x509CertificateImport
}
