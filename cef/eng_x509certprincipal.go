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

// ICefX509CertPrincipal Parent: ICefBaseRefCounted
//
//	Interface representing the issuer or subject field of an X.509 certificate.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_x509_certificate_capi.h">CEF source file: /include/capi/cef_x509_certificate_capi.h (cef_x509cert_principal_t))</a>
type ICefX509CertPrincipal interface {
	ICefBaseRefCounted
	// GetDisplayName
	//  Returns a name that can be used to represent the issuer. It tries in this order: Common Name (CN), Organization Name (O) and Organizational Unit Name (OU) and returns the first non-NULL one found.
	GetDisplayName() string // function
	// GetCommonName
	//  Returns the common name.
	GetCommonName() string // function
	// GetLocalityName
	//  Returns the locality name.
	GetLocalityName() string // function
	// GetStateOrProvinceName
	//  Returns the state or province name.
	GetStateOrProvinceName() string // function
	// GetCountryName
	//  Returns the country name.
	GetCountryName() string // function
	// GetOrganizationNames
	//  Retrieve the list of organization names.
	GetOrganizationNames(names IStrings) // procedure
	// GetOrganizationUnitNames
	//  Retrieve the list of organization unit names.
	GetOrganizationUnitNames(names IStrings) // procedure
}

// TCefX509CertPrincipal Parent: TCefBaseRefCounted
//
//	Interface representing the issuer or subject field of an X.509 certificate.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_x509_certificate_capi.h">CEF source file: /include/capi/cef_x509_certificate_capi.h (cef_x509cert_principal_t))</a>
type TCefX509CertPrincipal struct {
	TCefBaseRefCounted
}

// X509CertPrincipalRef -> ICefX509CertPrincipal
var X509CertPrincipalRef x509CertPrincipal

// x509CertPrincipal TCefX509CertPrincipal Ref
type x509CertPrincipal uintptr

func (m *x509CertPrincipal) UnWrap(data uintptr) ICefX509CertPrincipal {
	var resultCefX509CertPrincipal uintptr
	x509CertPrincipalImportAPI().SysCallN(7, uintptr(data), uintptr(unsafePointer(&resultCefX509CertPrincipal)))
	return AsCefX509CertPrincipal(resultCefX509CertPrincipal)
}

func (m *TCefX509CertPrincipal) GetDisplayName() string {
	r1 := x509CertPrincipalImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefX509CertPrincipal) GetCommonName() string {
	r1 := x509CertPrincipalImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefX509CertPrincipal) GetLocalityName() string {
	r1 := x509CertPrincipalImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefX509CertPrincipal) GetStateOrProvinceName() string {
	r1 := x509CertPrincipalImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCefX509CertPrincipal) GetCountryName() string {
	r1 := x509CertPrincipalImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCefX509CertPrincipal) GetOrganizationNames(names IStrings) {
	x509CertPrincipalImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(names))
}

func (m *TCefX509CertPrincipal) GetOrganizationUnitNames(names IStrings) {
	x509CertPrincipalImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(names))
}

var (
	x509CertPrincipalImport       *imports.Imports = nil
	x509CertPrincipalImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefX509CertPrincipal_GetCommonName", 0),
		/*1*/ imports.NewTable("CefX509CertPrincipal_GetCountryName", 0),
		/*2*/ imports.NewTable("CefX509CertPrincipal_GetDisplayName", 0),
		/*3*/ imports.NewTable("CefX509CertPrincipal_GetLocalityName", 0),
		/*4*/ imports.NewTable("CefX509CertPrincipal_GetOrganizationNames", 0),
		/*5*/ imports.NewTable("CefX509CertPrincipal_GetOrganizationUnitNames", 0),
		/*6*/ imports.NewTable("CefX509CertPrincipal_GetStateOrProvinceName", 0),
		/*7*/ imports.NewTable("CefX509CertPrincipal_UnWrap", 0),
	}
)

func x509CertPrincipalImportAPI() *imports.Imports {
	if x509CertPrincipalImport == nil {
		x509CertPrincipalImport = NewDefaultImports()
		x509CertPrincipalImport.SetImportTable(x509CertPrincipalImportTables)
		x509CertPrincipalImportTables = nil
	}
	return x509CertPrincipalImport
}
