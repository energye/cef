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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/types"
)

// ICefX509Certificate Parent: ICefBaseRefCounted
type ICefX509Certificate interface {
	ICefBaseRefCounted
	// GetSubject
	//  Returns the subject of the X.509 certificate. For HTTPS server
	//  certificates this represents the web server. The common name of the
	//  subject should match the host name of the web server.
	GetSubject() ICefX509CertPrincipal // function
	// GetIssuer
	//  Returns the issuer of the X.509 certificate.
	GetIssuer() ICefX509CertPrincipal // function
	// GetSerialNumber
	//  Returns the DER encoded serial number for the X.509 certificate. The value
	//  possibly includes a leading 00 byte.
	GetSerialNumber() ICefBinaryValue // function
	// GetValidStart
	//  Returns the date before which the X.509 certificate is invalid.
	//  CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidStart() int64 // function
	// GetValidExpiry
	//  Returns the date after which the X.509 certificate is invalid.
	//  CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidExpiry() int64 // function
	// GetValidStartAsDateTime
	//  Returns the date before which the X.509 certificate is invalid.
	//  CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidStartAsDateTime() types.TDateTime // function
	// GetValidExpiryAsDateTime
	//  Returns the date after which the X.509 certificate is invalid.
	//  CefBaseTime.GetTimeT() will return 0 if no date was specified.
	GetValidExpiryAsDateTime() types.TDateTime // function
	// GetDerEncoded
	//  Returns the DER encoded data for the X.509 certificate.
	GetDerEncoded() ICefBinaryValue // function
	// GetPemEncoded
	//  Returns the PEM encoded data for the X.509 certificate.
	GetPemEncoded() ICefBinaryValue // function
	// GetIssuerChainSize
	//  Returns the number of certificates in the issuer chain. If 0, the
	//  certificate is self-signed.
	GetIssuerChainSize() cefTypes.NativeUInt // function
	// GetDEREncodedIssuerChain
	//  Returns the DER encoded data for the certificate issuer chain. If we
	//  failed to encode a certificate in the chain it is still present in the
	//  array but is an NULL string.
	GetDEREncodedIssuerChain(chainCount cefTypes.NativeUInt, chain *ICefBinaryValueArray) // procedure
	// GetPEMEncodedIssuerChain
	//  Returns the PEM encoded data for the certificate issuer chain. If we
	//  failed to encode a certificate in the chain it is still present in the
	//  array but is an NULL string.
	GetPEMEncodedIssuerChain(chainCount cefTypes.NativeUInt, chain *ICefBinaryValueArray) // procedure
}

// ICEFX509CertificateRef Parent: ICefX509Certificate ICefBaseRefCountedRef
type ICEFX509CertificateRef interface {
	ICefX509Certificate
	ICefBaseRefCountedRef
	AsIntfX509Certificate() uintptr
}

type TCEFX509CertificateRef struct {
	TCefBaseRefCountedRef
}

func (m *TCEFX509CertificateRef) GetSubject() (result ICefX509CertPrincipal) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefX509CertPrincipalRef(resultPtr)
	return
}

func (m *TCEFX509CertificateRef) GetIssuer() (result ICefX509CertPrincipal) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefX509CertPrincipalRef(resultPtr)
	return
}

func (m *TCEFX509CertificateRef) GetSerialNumber() (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCEFX509CertificateRef) GetValidStart() (result int64) {
	if !m.IsValid() {
		return
	}
	cEFX509CertificateRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFX509CertificateRef) GetValidExpiry() (result int64) {
	if !m.IsValid() {
		return
	}
	cEFX509CertificateRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFX509CertificateRef) GetValidStartAsDateTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cEFX509CertificateRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFX509CertificateRef) GetValidExpiryAsDateTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cEFX509CertificateRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFX509CertificateRef) GetDerEncoded() (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCEFX509CertificateRef) GetPemEncoded() (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCEFX509CertificateRef) GetIssuerChainSize() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cEFX509CertificateRefAPI().SysCallN(10, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCEFX509CertificateRef) GetDEREncodedIssuerChain(chainCount cefTypes.NativeUInt, chain *ICefBinaryValueArray) {
	if !m.IsValid() {
		return
	}
	var chainPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(12, m.Instance(), uintptr(chainCount), uintptr(base.UnsafePointer(&chainPtr)))
	*chain = NewCefBinaryValueArray(int(chainCount), chainPtr)
}

func (m *TCEFX509CertificateRef) GetPEMEncodedIssuerChain(chainCount cefTypes.NativeUInt, chain *ICefBinaryValueArray) {
	if !m.IsValid() {
		return
	}
	var chainPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(13, m.Instance(), uintptr(chainCount), uintptr(base.UnsafePointer(&chainPtr)))
	*chain = NewCefBinaryValueArray(int(chainCount), chainPtr)
}

func (m *TCEFX509CertificateRef) AsIntfX509Certificate() uintptr {
	return m.GetIntfPointer(0)
}

// X509CertificateRef  is static instance
var X509CertificateRef _X509CertificateRefClass

// _X509CertificateRefClass is class type defined by TCEFX509CertificateRef
type _X509CertificateRefClass uintptr

func (_X509CertificateRefClass) UnWrap(data uintptr) (result ICefX509Certificate) {
	var resultPtr uintptr
	cEFX509CertificateRefAPI().SysCallN(11, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCEFX509CertificateRef(resultPtr)
	return
}

// NewX509CertificateRef class constructor
func NewX509CertificateRef(data uintptr) ICEFX509CertificateRef {
	var x509CertificatePtr uintptr // ICefX509Certificate
	r := cEFX509CertificateRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&x509CertificatePtr)))
	ret := AsCEFX509CertificateRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, x509CertificatePtr)
	}
	return ret
}

var (
	cEFX509CertificateRefOnce   base.Once
	cEFX509CertificateRefImport *imports.Imports = nil
)

func cEFX509CertificateRefAPI() *imports.Imports {
	cEFX509CertificateRefOnce.Do(func() {
		cEFX509CertificateRefImport = api.NewDefaultImports()
		cEFX509CertificateRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFX509CertificateRef_Create", 0), // constructor NewX509CertificateRef
			/* 1 */ imports.NewTable("TCEFX509CertificateRef_GetSubject", 0), // function GetSubject
			/* 2 */ imports.NewTable("TCEFX509CertificateRef_GetIssuer", 0), // function GetIssuer
			/* 3 */ imports.NewTable("TCEFX509CertificateRef_GetSerialNumber", 0), // function GetSerialNumber
			/* 4 */ imports.NewTable("TCEFX509CertificateRef_GetValidStart", 0), // function GetValidStart
			/* 5 */ imports.NewTable("TCEFX509CertificateRef_GetValidExpiry", 0), // function GetValidExpiry
			/* 6 */ imports.NewTable("TCEFX509CertificateRef_GetValidStartAsDateTime", 0), // function GetValidStartAsDateTime
			/* 7 */ imports.NewTable("TCEFX509CertificateRef_GetValidExpiryAsDateTime", 0), // function GetValidExpiryAsDateTime
			/* 8 */ imports.NewTable("TCEFX509CertificateRef_GetDerEncoded", 0), // function GetDerEncoded
			/* 9 */ imports.NewTable("TCEFX509CertificateRef_GetPemEncoded", 0), // function GetPemEncoded
			/* 10 */ imports.NewTable("TCEFX509CertificateRef_GetIssuerChainSize", 0), // function GetIssuerChainSize
			/* 11 */ imports.NewTable("TCEFX509CertificateRef_UnWrap", 0), // static function UnWrap
			/* 12 */ imports.NewTable("TCEFX509CertificateRef_GetDEREncodedIssuerChain", 0), // procedure GetDEREncodedIssuerChain
			/* 13 */ imports.NewTable("TCEFX509CertificateRef_GetPEMEncodedIssuerChain", 0), // procedure GetPEMEncodedIssuerChain
		}
	})
	return cEFX509CertificateRefImport
}
