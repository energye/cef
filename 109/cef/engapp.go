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

// IEngApp Parent: ICefAppOwn
type IEngApp interface {
	ICefAppOwn
	SetOnAppBeforeCommandLineProcessing(fn TOnAppBeforeCommandLineProcessingEvent) // property event
	SetOnAppRegisterCustomSchemes(fn TOnAppRegisterCustomSchemesEvent)             // property event
	SetOnAppGetResourceBundleHandler(fn TOnAppGetResourceBundleHandlerEvent)       // property event
	SetOnAppGetBrowserProcessHandler(fn TOnAppGetBrowserProcessHandlerEvent)       // property event
	SetOnAppGetRenderProcessHandler(fn TOnAppGetRenderProcessHandlerEvent)         // property event
	AsIntfApp() uintptr
}

type TEngApp struct {
	TCefAppOwn
}

func (m *TEngApp) SetOnAppBeforeCommandLineProcessing(fn TOnAppBeforeCommandLineProcessingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAppBeforeCommandLineProcessingEvent(fn)
	base.SetEvent(m, 1, engAppAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngApp) SetOnAppRegisterCustomSchemes(fn TOnAppRegisterCustomSchemesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAppRegisterCustomSchemesEvent(fn)
	base.SetEvent(m, 2, engAppAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngApp) SetOnAppGetResourceBundleHandler(fn TOnAppGetResourceBundleHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAppGetResourceBundleHandlerEvent(fn)
	base.SetEvent(m, 3, engAppAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngApp) SetOnAppGetBrowserProcessHandler(fn TOnAppGetBrowserProcessHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAppGetBrowserProcessHandlerEvent(fn)
	base.SetEvent(m, 4, engAppAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngApp) SetOnAppGetRenderProcessHandler(fn TOnAppGetRenderProcessHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAppGetRenderProcessHandlerEvent(fn)
	base.SetEvent(m, 5, engAppAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngApp) AsIntfApp() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngApp class constructor
func NewEngApp() IEngApp {
	var appPtr uintptr // ICefApp
	r := engAppAPI().SysCallN(0, uintptr(base.UnsafePointer(&appPtr)))
	ret := AsEngApp(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, appPtr)
	}
	return ret
}

var (
	engAppOnce   base.Once
	engAppImport *imports.Imports = nil
)

func engAppAPI() *imports.Imports {
	engAppOnce.Do(func() {
		engAppImport = api.NewDefaultImports()
		engAppImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngApp_Create", 0), // constructor NewEngApp
			/* 1 */ imports.NewTable("TEngApp_OnAppBeforeCommandLineProcessing", 0), // event OnAppBeforeCommandLineProcessing
			/* 2 */ imports.NewTable("TEngApp_OnAppRegisterCustomSchemes", 0), // event OnAppRegisterCustomSchemes
			/* 3 */ imports.NewTable("TEngApp_OnAppGetResourceBundleHandler", 0), // event OnAppGetResourceBundleHandler
			/* 4 */ imports.NewTable("TEngApp_OnAppGetBrowserProcessHandler", 0), // event OnAppGetBrowserProcessHandler
			/* 5 */ imports.NewTable("TEngApp_OnAppGetRenderProcessHandler", 0), // event OnAppGetRenderProcessHandler
		}
	})
	return engAppImport
}
