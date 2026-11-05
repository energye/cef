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

// IEngPrintHandler Parent: ICefPrintHandlerOwn
type IEngPrintHandler interface {
	ICefPrintHandlerOwn
	SetOnPrintPrintStart(fn TOnPrintPrintStartEvent)           // property event
	SetOnPrintPrintSettings(fn TOnPrintPrintSettingsEvent)     // property event
	SetOnPrintPrintDialog(fn TOnPrintPrintDialogEvent)         // property event
	SetOnPrintPrintJob(fn TOnPrintPrintJobEvent)               // property event
	SetOnPrintPrintReset(fn TOnPrintPrintResetEvent)           // property event
	SetOnPrintGetPDFPaperSize(fn TOnPrintGetPDFPaperSizeEvent) // property event
	AsIntfPrintHandler() uintptr
}

type TEngPrintHandler struct {
	TCefPrintHandlerOwn
}

func (m *TEngPrintHandler) SetOnPrintPrintStart(fn TOnPrintPrintStartEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintPrintStartEvent(fn)
	base.SetEvent(m, 1, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) SetOnPrintPrintSettings(fn TOnPrintPrintSettingsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintPrintSettingsEvent(fn)
	base.SetEvent(m, 2, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) SetOnPrintPrintDialog(fn TOnPrintPrintDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintPrintDialogEvent(fn)
	base.SetEvent(m, 3, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) SetOnPrintPrintJob(fn TOnPrintPrintJobEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintPrintJobEvent(fn)
	base.SetEvent(m, 4, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) SetOnPrintPrintReset(fn TOnPrintPrintResetEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintPrintResetEvent(fn)
	base.SetEvent(m, 5, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) SetOnPrintGetPDFPaperSize(fn TOnPrintGetPDFPaperSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintGetPDFPaperSizeEvent(fn)
	base.SetEvent(m, 6, engPrintHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPrintHandler) AsIntfPrintHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngPrintHandler class constructor
func NewEngPrintHandler() IEngPrintHandler {
	var printHandlerPtr uintptr // ICefPrintHandler
	r := engPrintHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&printHandlerPtr)))
	ret := AsEngPrintHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printHandlerPtr)
	}
	return ret
}

var (
	engPrintHandlerOnce   base.Once
	engPrintHandlerImport *imports.Imports = nil
)

func engPrintHandlerAPI() *imports.Imports {
	engPrintHandlerOnce.Do(func() {
		engPrintHandlerImport = api.NewDefaultImports()
		engPrintHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngPrintHandler_Create", 0), // constructor NewEngPrintHandler
			/* 1 */ imports.NewTable("TEngPrintHandler_OnPrintPrintStart", 0), // event OnPrintPrintStart
			/* 2 */ imports.NewTable("TEngPrintHandler_OnPrintPrintSettings", 0), // event OnPrintPrintSettings
			/* 3 */ imports.NewTable("TEngPrintHandler_OnPrintPrintDialog", 0), // event OnPrintPrintDialog
			/* 4 */ imports.NewTable("TEngPrintHandler_OnPrintPrintJob", 0), // event OnPrintPrintJob
			/* 5 */ imports.NewTable("TEngPrintHandler_OnPrintPrintReset", 0), // event OnPrintPrintReset
			/* 6 */ imports.NewTable("TEngPrintHandler_OnPrintGetPDFPaperSize", 0), // event OnPrintGetPDFPaperSize
		}
	})
	return engPrintHandlerImport
}
