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
type ICefSSLStatus interface {
	ICefBaseRefCounted
	IsSecureConnection() bool                // function
	GetCertStatus() TCefCertStatus           // function
	GetSSLVersion() TCefSSLVersion           // function
	GetContentStatus() TCefSSLContentStatus  // function
	GetX509Certificate() ICefX509Certificate // function
}

// TCefSSLStatus Parent: TCefBaseRefCounted
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
