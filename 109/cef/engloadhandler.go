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

// IEngLoadHandler Parent: ICefLoadHandlerOwn
type IEngLoadHandler interface {
	ICefLoadHandlerOwn
	SetOnLoadLoadingStateChange(fn TOnLoadLoadingStateChangeEvent) // property event
	SetOnLoadLoadStart(fn TOnLoadLoadStartEvent)                   // property event
	SetOnLoadLoadEnd(fn TOnLoadLoadEndEvent)                       // property event
	SetOnLoadLoadError(fn TOnLoadLoadErrorEvent)                   // property event
	AsIntfLoadHandler() uintptr
}

type TEngLoadHandler struct {
	TCefLoadHandlerOwn
}

func (m *TEngLoadHandler) SetOnLoadLoadingStateChange(fn TOnLoadLoadingStateChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadLoadingStateChangeEvent(fn)
	base.SetEvent(m, 1, engLoadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLoadHandler) SetOnLoadLoadStart(fn TOnLoadLoadStartEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadLoadStartEvent(fn)
	base.SetEvent(m, 2, engLoadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLoadHandler) SetOnLoadLoadEnd(fn TOnLoadLoadEndEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadLoadEndEvent(fn)
	base.SetEvent(m, 3, engLoadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLoadHandler) SetOnLoadLoadError(fn TOnLoadLoadErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadLoadErrorEvent(fn)
	base.SetEvent(m, 4, engLoadHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngLoadHandler) AsIntfLoadHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngLoadHandler class constructor
func NewEngLoadHandler() IEngLoadHandler {
	var loadHandlerPtr uintptr // ICefLoadHandler
	r := engLoadHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&loadHandlerPtr)))
	ret := AsEngLoadHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, loadHandlerPtr)
	}
	return ret
}

var (
	engLoadHandlerOnce   base.Once
	engLoadHandlerImport *imports.Imports = nil
)

func engLoadHandlerAPI() *imports.Imports {
	engLoadHandlerOnce.Do(func() {
		engLoadHandlerImport = api.NewDefaultImports()
		engLoadHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngLoadHandler_Create", 0), // constructor NewEngLoadHandler
			/* 1 */ imports.NewTable("TEngLoadHandler_OnLoadLoadingStateChange", 0), // event OnLoadLoadingStateChange
			/* 2 */ imports.NewTable("TEngLoadHandler_OnLoadLoadStart", 0), // event OnLoadLoadStart
			/* 3 */ imports.NewTable("TEngLoadHandler_OnLoadLoadEnd", 0), // event OnLoadLoadEnd
			/* 4 */ imports.NewTable("TEngLoadHandler_OnLoadLoadError", 0), // event OnLoadLoadError
		}
	})
	return engLoadHandlerImport
}
