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

// IEngRequestHandler Parent: ICefRequestHandlerOwn
type IEngRequestHandler interface {
	ICefRequestHandlerOwn
	SetOnRequestBeforeBrowse(fn TOnRequestBeforeBrowseEvent)                                 // property event
	SetOnRequestOpenUrlFromTab(fn TOnRequestOpenUrlFromTabEvent)                             // property event
	SetOnRequestGetAuthCredentials(fn TOnRequestGetAuthCredentialsEvent)                     // property event
	SetOnRequestCertificateError(fn TOnRequestCertificateErrorEvent)                         // property event
	SetOnRequestSelectClientCertificate(fn TOnRequestSelectClientCertificateEvent)           // property event
	SetOnRequestRenderProcessUnresponsive(fn TOnRequestRenderProcessUnresponsiveEvent)       // property event
	SetOnRequestGetResourceRequestHandler(fn TOnRequestGetResourceRequestHandlerEvent)       // property event
	SetOnRequestRenderViewReady(fn TOnRequestRenderViewReadyEvent)                           // property event
	SetOnRequestRenderProcessResponsive(fn TOnRequestRenderProcessResponsiveEvent)           // property event
	SetOnRequestRenderProcessTerminated(fn TOnRequestRenderProcessTerminatedEvent)           // property event
	SetOnRequestDocumentAvailableInMainFrame(fn TOnRequestDocumentAvailableInMainFrameEvent) // property event
	AsIntfRequestHandler() uintptr
}

type TEngRequestHandler struct {
	TCefRequestHandlerOwn
}

func (m *TEngRequestHandler) SetOnRequestBeforeBrowse(fn TOnRequestBeforeBrowseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestBeforeBrowseEvent(fn)
	base.SetEvent(m, 1, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestOpenUrlFromTab(fn TOnRequestOpenUrlFromTabEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestOpenUrlFromTabEvent(fn)
	base.SetEvent(m, 2, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestGetAuthCredentials(fn TOnRequestGetAuthCredentialsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestGetAuthCredentialsEvent(fn)
	base.SetEvent(m, 3, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestCertificateError(fn TOnRequestCertificateErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestCertificateErrorEvent(fn)
	base.SetEvent(m, 4, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestSelectClientCertificate(fn TOnRequestSelectClientCertificateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestSelectClientCertificateEvent(fn)
	base.SetEvent(m, 5, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestRenderProcessUnresponsive(fn TOnRequestRenderProcessUnresponsiveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestRenderProcessUnresponsiveEvent(fn)
	base.SetEvent(m, 6, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestGetResourceRequestHandler(fn TOnRequestGetResourceRequestHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestGetResourceRequestHandlerEvent(fn)
	base.SetEvent(m, 7, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestRenderViewReady(fn TOnRequestRenderViewReadyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestRenderViewReadyEvent(fn)
	base.SetEvent(m, 8, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestRenderProcessResponsive(fn TOnRequestRenderProcessResponsiveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestRenderProcessResponsiveEvent(fn)
	base.SetEvent(m, 9, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestRenderProcessTerminated(fn TOnRequestRenderProcessTerminatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestRenderProcessTerminatedEvent(fn)
	base.SetEvent(m, 10, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) SetOnRequestDocumentAvailableInMainFrame(fn TOnRequestDocumentAvailableInMainFrameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestDocumentAvailableInMainFrameEvent(fn)
	base.SetEvent(m, 11, engRequestHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestHandler) AsIntfRequestHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngRequestHandler class constructor
func NewEngRequestHandler() IEngRequestHandler {
	var requestHandlerPtr uintptr // ICefRequestHandler
	r := engRequestHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&requestHandlerPtr)))
	ret := AsEngRequestHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, requestHandlerPtr)
	}
	return ret
}

var (
	engRequestHandlerOnce   base.Once
	engRequestHandlerImport *imports.Imports = nil
)

func engRequestHandlerAPI() *imports.Imports {
	engRequestHandlerOnce.Do(func() {
		engRequestHandlerImport = api.NewDefaultImports()
		engRequestHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngRequestHandler_Create", 0), // constructor NewEngRequestHandler
			/* 1 */ imports.NewTable("TEngRequestHandler_OnRequestBeforeBrowse", 0), // event OnRequestBeforeBrowse
			/* 2 */ imports.NewTable("TEngRequestHandler_OnRequestOpenUrlFromTab", 0), // event OnRequestOpenUrlFromTab
			/* 3 */ imports.NewTable("TEngRequestHandler_OnRequestGetAuthCredentials", 0), // event OnRequestGetAuthCredentials
			/* 4 */ imports.NewTable("TEngRequestHandler_OnRequestCertificateError", 0), // event OnRequestCertificateError
			/* 5 */ imports.NewTable("TEngRequestHandler_OnRequestSelectClientCertificate", 0), // event OnRequestSelectClientCertificate
			/* 6 */ imports.NewTable("TEngRequestHandler_OnRequestRenderProcessUnresponsive", 0), // event OnRequestRenderProcessUnresponsive
			/* 7 */ imports.NewTable("TEngRequestHandler_OnRequestGetResourceRequestHandler", 0), // event OnRequestGetResourceRequestHandler
			/* 8 */ imports.NewTable("TEngRequestHandler_OnRequestRenderViewReady", 0), // event OnRequestRenderViewReady
			/* 9 */ imports.NewTable("TEngRequestHandler_OnRequestRenderProcessResponsive", 0), // event OnRequestRenderProcessResponsive
			/* 10 */ imports.NewTable("TEngRequestHandler_OnRequestRenderProcessTerminated", 0), // event OnRequestRenderProcessTerminated
			/* 11 */ imports.NewTable("TEngRequestHandler_OnRequestDocumentAvailableInMainFrame", 0), // event OnRequestDocumentAvailableInMainFrame
		}
	})
	return engRequestHandlerImport
}
