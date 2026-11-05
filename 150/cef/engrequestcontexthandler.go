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

// IEngRequestContextHandler Parent: ICefRequestContextHandlerOwn
type IEngRequestContextHandler interface {
	ICefRequestContextHandlerOwn
	SetOnRequestContextRequestContextInitialized(fn TOnRequestContextRequestContextInitializedEvent) // property event
	SetOnRequestContextGetResourceRequestHandler(fn TOnRequestContextGetResourceRequestHandlerEvent) // property event
	AsIntfRequestContextHandler() uintptr
}

type TEngRequestContextHandler struct {
	TCefRequestContextHandlerOwn
}

func (m *TEngRequestContextHandler) SetOnRequestContextRequestContextInitialized(fn TOnRequestContextRequestContextInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestContextRequestContextInitializedEvent(fn)
	base.SetEvent(m, 1, engRequestContextHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestContextHandler) SetOnRequestContextGetResourceRequestHandler(fn TOnRequestContextGetResourceRequestHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestContextGetResourceRequestHandlerEvent(fn)
	base.SetEvent(m, 2, engRequestContextHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRequestContextHandler) AsIntfRequestContextHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngRequestContextHandler class constructor
func NewEngRequestContextHandler() IEngRequestContextHandler {
	var requestContextHandlerPtr uintptr // ICefRequestContextHandler
	r := engRequestContextHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&requestContextHandlerPtr)))
	ret := AsEngRequestContextHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, requestContextHandlerPtr)
	}
	return ret
}

var (
	engRequestContextHandlerOnce   base.Once
	engRequestContextHandlerImport *imports.Imports = nil
)

func engRequestContextHandlerAPI() *imports.Imports {
	engRequestContextHandlerOnce.Do(func() {
		engRequestContextHandlerImport = api.NewDefaultImports()
		engRequestContextHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngRequestContextHandler_Create", 0), // constructor NewEngRequestContextHandler
			/* 1 */ imports.NewTable("TEngRequestContextHandler_OnRequestContextRequestContextInitialized", 0), // event OnRequestContextRequestContextInitialized
			/* 2 */ imports.NewTable("TEngRequestContextHandler_OnRequestContextGetResourceRequestHandler", 0), // event OnRequestContextGetResourceRequestHandler
		}
	})
	return engRequestContextHandlerImport
}
