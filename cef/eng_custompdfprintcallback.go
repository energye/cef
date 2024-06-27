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

// ICefCustomPDFPrintCallBack Parent: ICefPdfPrintCallback
type ICefCustomPDFPrintCallBack interface {
	ICefPdfPrintCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefPdfPrintCallback // procedure
}

// TCefCustomPDFPrintCallBack Parent: TCefPdfPrintCallback
type TCefCustomPDFPrintCallBack struct {
	TCefPdfPrintCallback
}

func NewCefCustomPDFPrintCallBack(aEvents IChromiumEvents) ICefCustomPDFPrintCallBack {
	r1 := customPDFPrintCallBackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomPDFPrintCallBack(r1)
}

func (m *TCefCustomPDFPrintCallBack) AsInterface() ICefPdfPrintCallback {
	var resultCefPdfPrintCallback uintptr
	customPDFPrintCallBackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefPdfPrintCallback)))
	return AsCefPdfPrintCallback(resultCefPdfPrintCallback)
}

var (
	customPDFPrintCallBackImport       *imports.Imports = nil
	customPDFPrintCallBackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomPDFPrintCallBack_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomPDFPrintCallBack_Create", 0),
	}
)

func customPDFPrintCallBackImportAPI() *imports.Imports {
	if customPDFPrintCallBackImport == nil {
		customPDFPrintCallBackImport = NewDefaultImports()
		customPDFPrintCallBackImport.SetImportTable(customPDFPrintCallBackImportTables)
		customPDFPrintCallBackImportTables = nil
	}
	return customPDFPrintCallBackImport
}
