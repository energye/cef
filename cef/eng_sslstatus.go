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

// ICefSSLStatus Parent: ICefBaseRefCounted
//
//	Interface representing the SSL information for a navigation entry.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_status_capi.h">CEF source file: /include/capi/cef_ssl_status_capi.h (cef_sslstatus_t))</a>
type ICefSSLStatus interface {
	ICefBaseRefCounted
	// IsSecureConnection
	//  Returns true (1) if the status is related to a secure SSL/TLS connection.
	IsSecureConnection() bool // function
	// GetCertStatus
	//  Returns a bitmask containing any and all problems verifying the server certificate.
	GetCertStatus() TCefCertStatus // function
	// GetSSLVersion
	//  Returns the SSL version used for the SSL connection.
	GetSSLVersion() TCefSSLVersion // function
	// GetContentStatus
	//  Returns a bitmask containing the page security content status.
	GetContentStatus() TCefSSLContentStatus // function
	// GetX509Certificate
	//  Returns the X.509 certificate.
	GetX509Certificate() ICefX509Certificate // function
}

// TCefSSLStatus Parent: TCefBaseRefCounted
//
//	Interface representing the SSL information for a navigation entry.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_status_capi.h">CEF source file: /include/capi/cef_ssl_status_capi.h (cef_sslstatus_t))</a>
type TCefSSLStatus struct {
	TCefBaseRefCounted
}

// SSLStatusRef -> ICefSSLStatus
var SSLStatusRef sSLStatus

// sSLStatus TCefSSLStatus Ref
type sSLStatus uintptr

func (m *sSLStatus) UnWrap(data uintptr) ICefSSLStatus {
	var resultCefSSLStatus uintptr
	sSLStatusImportAPI().SysCallN(5, uintptr(data), uintptr(unsafePointer(&resultCefSSLStatus)))
	return AsCefSSLStatus(resultCefSSLStatus)
}

func (m *TCefSSLStatus) IsSecureConnection() bool {
	r1 := sSLStatusImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefSSLStatus) GetCertStatus() TCefCertStatus {
	r1 := sSLStatusImportAPI().SysCallN(0, m.Instance())
	return TCefCertStatus(r1)
}

func (m *TCefSSLStatus) GetSSLVersion() TCefSSLVersion {
	r1 := sSLStatusImportAPI().SysCallN(2, m.Instance())
	return TCefSSLVersion(r1)
}

func (m *TCefSSLStatus) GetContentStatus() TCefSSLContentStatus {
	r1 := sSLStatusImportAPI().SysCallN(1, m.Instance())
	return TCefSSLContentStatus(r1)
}

func (m *TCefSSLStatus) GetX509Certificate() ICefX509Certificate {
	var resultCefX509Certificate uintptr
	sSLStatusImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefX509Certificate)))
	return AsCefX509Certificate(resultCefX509Certificate)
}

var (
	sSLStatusImport       *imports.Imports = nil
	sSLStatusImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSSLStatus_GetCertStatus", 0),
		/*1*/ imports.NewTable("CefSSLStatus_GetContentStatus", 0),
		/*2*/ imports.NewTable("CefSSLStatus_GetSSLVersion", 0),
		/*3*/ imports.NewTable("CefSSLStatus_GetX509Certificate", 0),
		/*4*/ imports.NewTable("CefSSLStatus_IsSecureConnection", 0),
		/*5*/ imports.NewTable("CefSSLStatus_UnWrap", 0),
	}
)

func sSLStatusImportAPI() *imports.Imports {
	if sSLStatusImport == nil {
		sSLStatusImport = NewDefaultImports()
		sSLStatusImport.SetImportTable(sSLStatusImportTables)
		sSLStatusImportTables = nil
	}
	return sSLStatusImport
}
