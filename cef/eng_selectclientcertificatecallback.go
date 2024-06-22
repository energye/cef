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

// ICefSelectClientCertificateCallback Parent: ICefBaseRefCounted
type ICefSelectClientCertificateCallback interface {
	ICefBaseRefCounted
	Select(cert ICefX509Certificate) // procedure
}

// TCefSelectClientCertificateCallback Parent: TCefBaseRefCounted
type TCefSelectClientCertificateCallback struct {
	TCefBaseRefCounted
}

// SelectClientCertificateCallbackRef -> ICefSelectClientCertificateCallback
var SelectClientCertificateCallbackRef selectClientCertificateCallback

// selectClientCertificateCallback TCefSelectClientCertificateCallback Ref
type selectClientCertificateCallback uintptr

func (m *selectClientCertificateCallback) UnWrap(data uintptr) ICefSelectClientCertificateCallback {
	var resultCefSelectClientCertificateCallback uintptr
	selectClientCertificateCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefSelectClientCertificateCallback)))
	return AsCefSelectClientCertificateCallback(resultCefSelectClientCertificateCallback)
}

func (m *TCefSelectClientCertificateCallback) Select(cert ICefX509Certificate) {
	selectClientCertificateCallbackImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(cert))
}

var (
	selectClientCertificateCallbackImport       *imports.Imports = nil
	selectClientCertificateCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSelectClientCertificateCallback_Select", 0),
		/*1*/ imports.NewTable("CefSelectClientCertificateCallback_UnWrap", 0),
	}
)

func selectClientCertificateCallbackImportAPI() *imports.Imports {
	if selectClientCertificateCallbackImport == nil {
		selectClientCertificateCallbackImport = NewDefaultImports()
		selectClientCertificateCallbackImport.SetImportTable(selectClientCertificateCallbackImportTables)
		selectClientCertificateCallbackImportTables = nil
	}
	return selectClientCertificateCallbackImport
}
