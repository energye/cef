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
}

// TCefCustomPDFPrintCallBack Parent: TCefPdfPrintCallback
type TCefCustomPDFPrintCallBack struct {
	TCefPdfPrintCallback
}

func NewCefCustomPDFPrintCallBack(aEvents IChromiumEvents) ICefCustomPDFPrintCallBack {
	r1 := customPDFPrintCallBackImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefCustomPDFPrintCallBack(r1)
}

var (
	customPDFPrintCallBackImport       *imports.Imports = nil
	customPDFPrintCallBackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomPDFPrintCallBack_Create", 0),
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
