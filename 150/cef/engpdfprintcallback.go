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

// IEngPdfPrintCallback Parent: ICefPdfPrintCallbackOwn
type IEngPdfPrintCallback interface {
	ICefPdfPrintCallbackOwn
	SetOnPdfPrintCallbackPdfPrintFinished(fn TOnPdfPrintCallbackPdfPrintFinishedEvent) // property event
	AsIntfPdfPrintCallback() uintptr
}

type TEngPdfPrintCallback struct {
	TCefPdfPrintCallbackOwn
}

func (m *TEngPdfPrintCallback) SetOnPdfPrintCallbackPdfPrintFinished(fn TOnPdfPrintCallbackPdfPrintFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPdfPrintCallbackPdfPrintFinishedEvent(fn)
	base.SetEvent(m, 1, engPdfPrintCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPdfPrintCallback) AsIntfPdfPrintCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngPdfPrintCallback class constructor
func NewEngPdfPrintCallback() IEngPdfPrintCallback {
	var pdfPrintCallbackPtr uintptr // ICefPdfPrintCallback
	r := engPdfPrintCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&pdfPrintCallbackPtr)))
	ret := AsEngPdfPrintCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, pdfPrintCallbackPtr)
	}
	return ret
}

var (
	engPdfPrintCallbackOnce   base.Once
	engPdfPrintCallbackImport *imports.Imports = nil
)

func engPdfPrintCallbackAPI() *imports.Imports {
	engPdfPrintCallbackOnce.Do(func() {
		engPdfPrintCallbackImport = api.NewDefaultImports()
		engPdfPrintCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngPdfPrintCallback_Create", 0), // constructor NewEngPdfPrintCallback
			/* 1 */ imports.NewTable("TEngPdfPrintCallback_OnPdfPrintCallbackPdfPrintFinished", 0), // event OnPdfPrintCallbackPdfPrintFinished
		}
	})
	return engPdfPrintCallbackImport
}
