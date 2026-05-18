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

// IEngDownloadHandler Parent: ICefDownloadHandlerOwn
type IEngDownloadHandler interface {
	ICefDownloadHandlerOwn
	SetOnDownloadCanDownload(fn TOnDownloadCanDownloadEvent)         // property event
	SetOnDownloadBeforeDownload(fn TOnDownloadBeforeDownloadEvent)   // property event
	SetOnDownloadDownloadUpdated(fn TOnDownloadDownloadUpdatedEvent) // property event
	AsIntfDownloadHandler() uintptr
}

type TEngDownloadHandler struct {
	TCefDownloadHandlerOwn
}

func (m *TEngDownloadHandler) SetOnDownloadCanDownload(fn TOnDownloadCanDownloadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadCanDownloadEvent(fn)
	base.SetEvent(m, 1, engDownloadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDownloadHandler) SetOnDownloadBeforeDownload(fn TOnDownloadBeforeDownloadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadBeforeDownloadEvent(fn)
	base.SetEvent(m, 2, engDownloadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDownloadHandler) SetOnDownloadDownloadUpdated(fn TOnDownloadDownloadUpdatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadDownloadUpdatedEvent(fn)
	base.SetEvent(m, 3, engDownloadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDownloadHandler) AsIntfDownloadHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDownloadHandler class constructor
func NewEngDownloadHandler() IEngDownloadHandler {
	var downloadHandlerPtr uintptr // ICefDownloadHandler
	r := engDownloadHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&downloadHandlerPtr)))
	ret := AsEngDownloadHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadHandlerPtr)
	}
	return ret
}

var (
	engDownloadHandlerOnce   base.Once
	engDownloadHandlerImport *imports.Imports = nil
)

func engDownloadHandlerAPI() *imports.Imports {
	engDownloadHandlerOnce.Do(func() {
		engDownloadHandlerImport = api.NewDefaultImports()
		engDownloadHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDownloadHandler_Create", 0), // constructor NewEngDownloadHandler
			/* 1 */ imports.NewTable("TEngDownloadHandler_OnDownloadCanDownload", 0), // event OnDownloadCanDownload
			/* 2 */ imports.NewTable("TEngDownloadHandler_OnDownloadBeforeDownload", 0), // event OnDownloadBeforeDownload
			/* 3 */ imports.NewTable("TEngDownloadHandler_OnDownloadDownloadUpdated", 0), // event OnDownloadDownloadUpdated
		}
	})
	return engDownloadHandlerImport
}
