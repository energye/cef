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

// IEngResourceRequestHandler Parent: ICefResourceRequestHandlerOwn
type IEngResourceRequestHandler interface {
	ICefResourceRequestHandlerOwn
	SetOnResourceRequestBeforeResourceLoad(fn TOnResourceRequestBeforeResourceLoadEvent)               // property event
	SetOnResourceRequestResourceResponse(fn TOnResourceRequestResourceResponseEvent)                   // property event
	SetOnResourceRequestGetCookieAccessFilter(fn TOnResourceRequestGetCookieAccessFilterEvent)         // property event
	SetOnResourceRequestGetResourceHandler(fn TOnResourceRequestGetResourceHandlerEvent)               // property event
	SetOnResourceRequestResourceRedirect(fn TOnResourceRequestResourceRedirectEvent)                   // property event
	SetOnResourceRequestGetResourceResponseFilter(fn TOnResourceRequestGetResourceResponseFilterEvent) // property event
	SetOnResourceRequestResourceLoadComplete(fn TOnResourceRequestResourceLoadCompleteEvent)           // property event
	SetOnResourceRequestProtocolExecution(fn TOnResourceRequestProtocolExecutionEvent)                 // property event
	AsIntfResourceRequestHandler() uintptr
}

type TEngResourceRequestHandler struct {
	TCefResourceRequestHandlerOwn
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestBeforeResourceLoad(fn TOnResourceRequestBeforeResourceLoadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestBeforeResourceLoadEvent(fn)
	base.SetEvent(m, 1, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestResourceResponse(fn TOnResourceRequestResourceResponseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestResourceResponseEvent(fn)
	base.SetEvent(m, 2, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestGetCookieAccessFilter(fn TOnResourceRequestGetCookieAccessFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestGetCookieAccessFilterEvent(fn)
	base.SetEvent(m, 3, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestGetResourceHandler(fn TOnResourceRequestGetResourceHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestGetResourceHandlerEvent(fn)
	base.SetEvent(m, 4, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestResourceRedirect(fn TOnResourceRequestResourceRedirectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestResourceRedirectEvent(fn)
	base.SetEvent(m, 5, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestGetResourceResponseFilter(fn TOnResourceRequestGetResourceResponseFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestGetResourceResponseFilterEvent(fn)
	base.SetEvent(m, 6, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestResourceLoadComplete(fn TOnResourceRequestResourceLoadCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestResourceLoadCompleteEvent(fn)
	base.SetEvent(m, 7, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) SetOnResourceRequestProtocolExecution(fn TOnResourceRequestProtocolExecutionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRequestProtocolExecutionEvent(fn)
	base.SetEvent(m, 8, engResourceRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceRequestHandler) AsIntfResourceRequestHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngResourceRequestHandler class constructor
func NewEngResourceRequestHandler() IEngResourceRequestHandler {
	var resourceRequestHandlerPtr uintptr // ICefResourceRequestHandler
	r := engResourceRequestHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&resourceRequestHandlerPtr)))
	ret := AsEngResourceRequestHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceRequestHandlerPtr)
	}
	return ret
}

var (
	engResourceRequestHandlerOnce   base.Once
	engResourceRequestHandlerImport *imports.Imports = nil
)

func engResourceRequestHandlerAPI() *imports.Imports {
	engResourceRequestHandlerOnce.Do(func() {
		engResourceRequestHandlerImport = api.NewDefaultImports()
		engResourceRequestHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngResourceRequestHandler_Create", 0), // constructor NewEngResourceRequestHandler
			/* 1 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestBeforeResourceLoad", 0), // event OnResourceRequestBeforeResourceLoad
			/* 2 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestResourceResponse", 0), // event OnResourceRequestResourceResponse
			/* 3 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestGetCookieAccessFilter", 0), // event OnResourceRequestGetCookieAccessFilter
			/* 4 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestGetResourceHandler", 0), // event OnResourceRequestGetResourceHandler
			/* 5 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestResourceRedirect", 0), // event OnResourceRequestResourceRedirect
			/* 6 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestGetResourceResponseFilter", 0), // event OnResourceRequestGetResourceResponseFilter
			/* 7 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestResourceLoadComplete", 0), // event OnResourceRequestResourceLoadComplete
			/* 8 */ imports.NewTable("TEngResourceRequestHandler_OnResourceRequestProtocolExecution", 0), // event OnResourceRequestProtocolExecution
		}
	})
	return engResourceRequestHandlerImport
}
