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

// IEngResourceHandler Parent: ICefResourceHandlerOwn
type IEngResourceHandler interface {
	ICefResourceHandlerOwn
	SetOnResourceOpen(fn TOnResourceOpenEvent)                             // property event
	SetOnResourceProcessRequest(fn TOnResourceProcessRequestEvent)         // property event
	SetOnResourceSkip(fn TOnResourceSkipEvent)                             // property event
	SetOnResourceRead(fn TOnResourceReadEvent)                             // property event
	SetOnResourceReadResponse(fn TOnResourceReadResponseEvent)             // property event
	SetOnResourceGetResponseHeaders(fn TOnResourceGetResponseHeadersEvent) // property event
	SetOnResourceCancel(fn TOnResourceCancelEvent)                         // property event
	AsIntfResourceHandler() uintptr
}

type TEngResourceHandler struct {
	TCefResourceHandlerOwn
}

func (m *TEngResourceHandler) SetOnResourceOpen(fn TOnResourceOpenEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceOpenEvent(fn)
	base.SetEvent(m, 1, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceProcessRequest(fn TOnResourceProcessRequestEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceProcessRequestEvent(fn)
	base.SetEvent(m, 2, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceSkip(fn TOnResourceSkipEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceSkipEvent(fn)
	base.SetEvent(m, 3, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceRead(fn TOnResourceReadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceReadEvent(fn)
	base.SetEvent(m, 4, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceReadResponse(fn TOnResourceReadResponseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceReadResponseEvent(fn)
	base.SetEvent(m, 5, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceGetResponseHeaders(fn TOnResourceGetResponseHeadersEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceGetResponseHeadersEvent(fn)
	base.SetEvent(m, 6, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) SetOnResourceCancel(fn TOnResourceCancelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceCancelEvent(fn)
	base.SetEvent(m, 7, engResourceHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceHandler) AsIntfResourceHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngResourceHandler class constructor
func NewEngResourceHandler(browser ICefBrowser, frame ICefFrame, schemeName string, request ICefRequest) IEngResourceHandler {
	var resourceHandlerPtr uintptr // ICefResourceHandler
	r := engResourceHandlerAPI().SysCallN(0, base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), api.PasStr(schemeName), base.GetObjectUintptr(request), uintptr(base.UnsafePointer(&resourceHandlerPtr)))
	ret := AsEngResourceHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceHandlerPtr)
	}
	return ret
}

var (
	engResourceHandlerOnce   base.Once
	engResourceHandlerImport *imports.Imports = nil
)

func engResourceHandlerAPI() *imports.Imports {
	engResourceHandlerOnce.Do(func() {
		engResourceHandlerImport = api.NewDefaultImports()
		engResourceHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngResourceHandler_Create", 0), // constructor NewEngResourceHandler
			/* 1 */ imports.NewTable("TEngResourceHandler_OnResourceOpen", 0), // event OnResourceOpen
			/* 2 */ imports.NewTable("TEngResourceHandler_OnResourceProcessRequest", 0), // event OnResourceProcessRequest
			/* 3 */ imports.NewTable("TEngResourceHandler_OnResourceSkip", 0), // event OnResourceSkip
			/* 4 */ imports.NewTable("TEngResourceHandler_OnResourceRead", 0), // event OnResourceRead
			/* 5 */ imports.NewTable("TEngResourceHandler_OnResourceReadResponse", 0), // event OnResourceReadResponse
			/* 6 */ imports.NewTable("TEngResourceHandler_OnResourceGetResponseHeaders", 0), // event OnResourceGetResponseHeaders
			/* 7 */ imports.NewTable("TEngResourceHandler_OnResourceCancel", 0), // event OnResourceCancel
		}
	})
	return engResourceHandlerImport
}
