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
//
//	Interface representing a X.509 certificate.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_x509_certificate_capi.h">CEF source file: /include/capi/cef_x509_certificate_capi.h (cef_x509certificate_t))</a>
type ICefX509Certificate interface {
	ICefBaseRefCounted
	// GetDEREncodedIssuerChain
	//  Returns the DER encoded data for the certificate issuer chain. If we failed to encode a certificate in the chain it is still present in the array but is an NULL string.
	GetDEREncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray)
	// GetPEMEncodedIssuerChain
	//  Returns the PEM encoded data for the certificate issuer chain. If we failed to encode a certificate in the chain it is still present in the array but is an NULL string.
	GetPEMEncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray)
	// GetSubject
	//  Returns the subject of the X.509 certificate. For HTTPS server certificates this represents the web server. The common name of the subject should match the host name of the web server.
	GetSubject() ICefX509CertPrincipal // function
	// GetIssuer
	//  Returns the issuer of the X.509 certificate.
	GetIssuer() ICefX509CertPrincipal // function
	// GetSerialNumber
	//  Returns the DER encoded serial number for the X.509 certificate. The value possibly includes a leading 00 byte.
	GetSerialNumber() ICefBinaryValue // function
	// GetValidStart
	//  Returns the date before which the X.509 certificate is invalid. CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidStart() TCefBaseTime // function
	// GetValidExpiry
	//  Returns the date after which the X.509 certificate is invalid. CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidExpiry() TCefBaseTime // function
	// GetValidStartAsDateTime
	//  Returns the date before which the X.509 certificate is invalid. CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidStartAsDateTime() (resultDateTime TDateTime) // function
	// GetValidExpiryAsDateTime
	//  Returns the date after which the X.509 certificate is invalid. CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidExpiryAsDateTime() (resultDateTime TDateTime) // function
	// GetDerEncoded
	//  Returns the DER encoded data for the X.509 certificate.
	GetDerEncoded() ICefBinaryValue // function
	// GetPemEncoded
	//  Returns the PEM encoded data for the X.509 certificate.
	GetPemEncoded() ICefBinaryValue // function
	// GetIssuerChainSize
	//  Returns the number of certificates in the issuer chain. If 0, the certificate is self-signed.
	GetIssuerChainSize() NativeUInt // function
}

// TCefX509Certificate Parent: TCefBaseRefCounted
//
//	Interface representing a X.509 certificate.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_x509_certificate_capi.h">CEF source file: /include/capi/cef_x509_certificate_capi.h (cef_x509certificate_t))</a>
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
