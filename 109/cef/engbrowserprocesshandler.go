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

// IEngBrowserProcessHandler Parent: ICefBrowserProcessHandlerOwn
type IEngBrowserProcessHandler interface {
	ICefBrowserProcessHandlerOwn
	SetOnBrowserProcessRegisterCustomPreferences(fn TOnBrowserProcessRegisterCustomPreferencesEvent) // property event
	SetOnBrowserProcessContextInitialized(fn TOnBrowserProcessContextInitializedEvent)               // property event
	SetOnBrowserProcessBeforeChildProcessLaunch(fn TOnBrowserProcessBeforeChildProcessLaunchEvent)   // property event
	SetOnBrowserProcessScheduleMessagePumpWork(fn TOnBrowserProcessScheduleMessagePumpWorkEvent)     // property event
	SetOnBrowserProcessGetDefaultClient(fn TOnBrowserProcessGetDefaultClientEvent)                   // property event
	AsIntfBrowserProcessHandler() uintptr
}

type TEngBrowserProcessHandler struct {
	TCefBrowserProcessHandlerOwn
}

func (m *TEngBrowserProcessHandler) SetOnBrowserProcessRegisterCustomPreferences(fn TOnBrowserProcessRegisterCustomPreferencesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessRegisterCustomPreferencesEvent(fn)
	base.SetEvent(m, 1, engBrowserProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserProcessHandler) SetOnBrowserProcessContextInitialized(fn TOnBrowserProcessContextInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessContextInitializedEvent(fn)
	base.SetEvent(m, 2, engBrowserProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserProcessHandler) SetOnBrowserProcessBeforeChildProcessLaunch(fn TOnBrowserProcessBeforeChildProcessLaunchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessBeforeChildProcessLaunchEvent(fn)
	base.SetEvent(m, 3, engBrowserProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserProcessHandler) SetOnBrowserProcessScheduleMessagePumpWork(fn TOnBrowserProcessScheduleMessagePumpWorkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessScheduleMessagePumpWorkEvent(fn)
	base.SetEvent(m, 4, engBrowserProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserProcessHandler) SetOnBrowserProcessGetDefaultClient(fn TOnBrowserProcessGetDefaultClientEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessGetDefaultClientEvent(fn)
	base.SetEvent(m, 5, engBrowserProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserProcessHandler) AsIntfBrowserProcessHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngBrowserProcessHandler class constructor
func NewEngBrowserProcessHandler() IEngBrowserProcessHandler {
	var browserProcessHandlerPtr uintptr // ICefBrowserProcessHandler
	r := engBrowserProcessHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&browserProcessHandlerPtr)))
	ret := AsEngBrowserProcessHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserProcessHandlerPtr)
	}
	return ret
}

var (
	engBrowserProcessHandlerOnce   base.Once
	engBrowserProcessHandlerImport *imports.Imports = nil
)

func engBrowserProcessHandlerAPI() *imports.Imports {
	engBrowserProcessHandlerOnce.Do(func() {
		engBrowserProcessHandlerImport = api.NewDefaultImports()
		engBrowserProcessHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngBrowserProcessHandler_Create", 0), // constructor NewEngBrowserProcessHandler
			/* 1 */ imports.NewTable("TEngBrowserProcessHandler_OnBrowserProcessRegisterCustomPreferences", 0), // event OnBrowserProcessRegisterCustomPreferences
			/* 2 */ imports.NewTable("TEngBrowserProcessHandler_OnBrowserProcessContextInitialized", 0), // event OnBrowserProcessContextInitialized
			/* 3 */ imports.NewTable("TEngBrowserProcessHandler_OnBrowserProcessBeforeChildProcessLaunch", 0), // event OnBrowserProcessBeforeChildProcessLaunch
			/* 4 */ imports.NewTable("TEngBrowserProcessHandler_OnBrowserProcessScheduleMessagePumpWork", 0), // event OnBrowserProcessScheduleMessagePumpWork
			/* 5 */ imports.NewTable("TEngBrowserProcessHandler_OnBrowserProcessGetDefaultClient", 0), // event OnBrowserProcessGetDefaultClient
		}
	})
	return engBrowserProcessHandlerImport
}
