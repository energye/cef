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

// ICefSslInfo Parent: ICefBaseRefCounted
//
//	Interface representing SSL information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_info_capi.h">CEF source file: /include/capi/cef_ssl_info_capi.h (cef_sslinfo_t))</a>
type ICefSslInfo interface {
	ICefBaseRefCounted
	// GetCertStatus
	//  Returns a bitmask containing any and all problems verifying the server certificate.
	GetCertStatus() TCefCertStatus // function
	// GetX509Certificate
	//  Returns the X.509 certificate.
	GetX509Certificate() ICefX509Certificate // function
}

// TCefSslInfo Parent: TCefBaseRefCounted
//
//	Interface representing SSL information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_ssl_info_capi.h">CEF source file: /include/capi/cef_ssl_info_capi.h (cef_sslinfo_t))</a>
type TCefSslInfo struct {
	TCefBaseRefCounted
}

// SslInfoRef -> ICefSslInfo
var SslInfoRef sslInfo

// sslInfo TCefSslInfo Ref
type sslInfo uintptr

func (m *sslInfo) UnWrap(data uintptr) ICefSslInfo {
	var resultCefSslInfo uintptr
	sslInfoImportAPI().SysCallN(2, uintptr(data), uintptr(unsafePointer(&resultCefSslInfo)))
	return AsCefSslInfo(resultCefSslInfo)
}

func (m *TCefSslInfo) GetCertStatus() TCefCertStatus {
	r1 := sslInfoImportAPI().SysCallN(0, m.Instance())
	return TCefCertStatus(r1)
}

func (m *TCefSslInfo) GetX509Certificate() ICefX509Certificate {
	var resultCefX509Certificate uintptr
	sslInfoImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefX509Certificate)))
	return AsCefX509Certificate(resultCefX509Certificate)
}

var (
	sslInfoImport       *imports.Imports = nil
	sslInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSslInfo_GetCertStatus", 0),
		/*1*/ imports.NewTable("CefSslInfo_GetX509Certificate", 0),
		/*2*/ imports.NewTable("CefSslInfo_UnWrap", 0),
	}
)

func sslInfoImportAPI() *imports.Imports {
	if sslInfoImport == nil {
		sslInfoImport = NewDefaultImports()
		sslInfoImport.SetImportTable(sslInfoImportTables)
		sslInfoImportTables = nil
	}
	return sslInfoImport
}
