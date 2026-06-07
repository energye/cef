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

// IEngExtensionHandler Parent: ICefExtensionHandlerOwn
type IEngExtensionHandler interface {
	ICefExtensionHandlerOwn
	SetOnExtensionBeforeBackgroundBrowser(fn TOnExtensionBeforeBackgroundBrowserEvent) // property event
	SetOnExtensionBeforeBrowser(fn TOnExtensionBeforeBrowserEvent)                     // property event
	SetOnExtensionCanAccessBrowser(fn TOnExtensionCanAccessBrowserEvent)               // property event
	SetOnExtensionGetExtensionResource(fn TOnExtensionGetExtensionResourceEvent)       // property event
	SetOnExtensionExtensionLoadFailed(fn TOnExtensionExtensionLoadFailedEvent)         // property event
	SetOnExtensionExtensionLoaded(fn TOnExtensionExtensionLoadedEvent)                 // property event
	SetOnExtensionExtensionUnloaded(fn TOnExtensionExtensionUnloadedEvent)             // property event
	SetOnExtensionGetActiveBrowser(fn TOnExtensionGetActiveBrowserEvent)               // property event
	AsIntfExtensionHandler() uintptr
}

type TEngExtensionHandler struct {
	TCefExtensionHandlerOwn
}

func (m *TEngExtensionHandler) SetOnExtensionBeforeBackgroundBrowser(fn TOnExtensionBeforeBackgroundBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionBeforeBackgroundBrowserEvent(fn)
	base.SetEvent(m, 1, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionBeforeBrowser(fn TOnExtensionBeforeBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionBeforeBrowserEvent(fn)
	base.SetEvent(m, 2, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionCanAccessBrowser(fn TOnExtensionCanAccessBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionCanAccessBrowserEvent(fn)
	base.SetEvent(m, 3, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionGetExtensionResource(fn TOnExtensionGetExtensionResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionGetExtensionResourceEvent(fn)
	base.SetEvent(m, 4, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionExtensionLoadFailed(fn TOnExtensionExtensionLoadFailedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionExtensionLoadFailedEvent(fn)
	base.SetEvent(m, 5, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionExtensionLoaded(fn TOnExtensionExtensionLoadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionExtensionLoadedEvent(fn)
	base.SetEvent(m, 6, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionExtensionUnloaded(fn TOnExtensionExtensionUnloadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionExtensionUnloadedEvent(fn)
	base.SetEvent(m, 7, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) SetOnExtensionGetActiveBrowser(fn TOnExtensionGetActiveBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionGetActiveBrowserEvent(fn)
	base.SetEvent(m, 8, engExtensionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngExtensionHandler) AsIntfExtensionHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngExtensionHandler class constructor
func NewEngExtensionHandler() IEngExtensionHandler {
	var extensionHandlerPtr uintptr // ICefExtensionHandler
	r := engExtensionHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&extensionHandlerPtr)))
	ret := AsEngExtensionHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, extensionHandlerPtr)
	}
	return ret
}

var (
	engExtensionHandlerOnce   base.Once
	engExtensionHandlerImport *imports.Imports = nil
)

func engExtensionHandlerAPI() *imports.Imports {
	engExtensionHandlerOnce.Do(func() {
		engExtensionHandlerImport = api.NewDefaultImports()
		engExtensionHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngExtensionHandler_Create", 0), // constructor NewEngExtensionHandler
			/* 1 */ imports.NewTable("TEngExtensionHandler_OnExtensionBeforeBackgroundBrowser", 0), // event OnExtensionBeforeBackgroundBrowser
			/* 2 */ imports.NewTable("TEngExtensionHandler_OnExtensionBeforeBrowser", 0), // event OnExtensionBeforeBrowser
			/* 3 */ imports.NewTable("TEngExtensionHandler_OnExtensionCanAccessBrowser", 0), // event OnExtensionCanAccessBrowser
			/* 4 */ imports.NewTable("TEngExtensionHandler_OnExtensionGetExtensionResource", 0), // event OnExtensionGetExtensionResource
			/* 5 */ imports.NewTable("TEngExtensionHandler_OnExtensionExtensionLoadFailed", 0), // event OnExtensionExtensionLoadFailed
			/* 6 */ imports.NewTable("TEngExtensionHandler_OnExtensionExtensionLoaded", 0), // event OnExtensionExtensionLoaded
			/* 7 */ imports.NewTable("TEngExtensionHandler_OnExtensionExtensionUnloaded", 0), // event OnExtensionExtensionUnloaded
			/* 8 */ imports.NewTable("TEngExtensionHandler_OnExtensionGetActiveBrowser", 0), // event OnExtensionGetActiveBrowser
		}
	})
	return engExtensionHandlerImport
}
