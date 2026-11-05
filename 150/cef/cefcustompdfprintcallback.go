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
)

// ICefCustomPDFPrintCallBack Parent: ICefPdfPrintCallbackOwn
type ICefCustomPDFPrintCallBack interface {
	ICefPdfPrintCallbackOwn
	AsIntfPdfPrintCallback() uintptr
}

type TCefCustomPDFPrintCallBack struct {
	TCefPdfPrintCallbackOwn
}

func (m *TCefCustomPDFPrintCallBack) AsIntfPdfPrintCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomPDFPrintCallBack class constructor
func NewCustomPDFPrintCallBack(events IChromiumCore) ICefCustomPDFPrintCallBack {
	var pdfPrintCallbackPtr uintptr // ICefPdfPrintCallback
	r := cefCustomPDFPrintCallBackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&pdfPrintCallbackPtr)))
	ret := AsCefCustomPDFPrintCallBack(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, pdfPrintCallbackPtr)
	}
	return ret
}

var (
	cefCustomPDFPrintCallBackOnce   base.Once
	cefCustomPDFPrintCallBackImport *imports.Imports = nil
)

func cefCustomPDFPrintCallBackAPI() *imports.Imports {
	cefCustomPDFPrintCallBackOnce.Do(func() {
		cefCustomPDFPrintCallBackImport = api.NewDefaultImports()
		cefCustomPDFPrintCallBackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomPDFPrintCallBack_Create", 0), // constructor NewCustomPDFPrintCallBack
		}
	})
	return cefCustomPDFPrintCallBackImport
}
