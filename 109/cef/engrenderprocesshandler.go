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

// IEngRenderProcessHandler Parent: ICefRenderProcessHandlerOwn
type IEngRenderProcessHandler interface {
	ICefRenderProcessHandlerOwn
	SetOnRenderProcessGetLoadHandler(fn TOnRenderProcessGetLoadHandlerEvent)                 // property event
	SetOnRenderProcessProcessMessageReceived(fn TOnRenderProcessProcessMessageReceivedEvent) // property event
	SetOnRenderProcessWebKitInitialized(fn TOnRenderProcessWebKitInitializedEvent)           // property event
	SetOnRenderProcessBrowserCreated(fn TOnRenderProcessBrowserCreatedEvent)                 // property event
	SetOnRenderProcessBrowserDestroyed(fn TOnRenderProcessBrowserDestroyedEvent)             // property event
	SetOnRenderProcessContextCreated(fn TOnRenderProcessContextCreatedEvent)                 // property event
	SetOnRenderProcessContextReleased(fn TOnRenderProcessContextReleasedEvent)               // property event
	SetOnRenderProcessUncaughtException(fn TOnRenderProcessUncaughtExceptionEvent)           // property event
	SetOnRenderProcessFocusedNodeChanged(fn TOnRenderProcessFocusedNodeChangedEvent)         // property event
	AsIntfRenderProcessHandler() uintptr
}

type TEngRenderProcessHandler struct {
	TCefRenderProcessHandlerOwn
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessGetLoadHandler(fn TOnRenderProcessGetLoadHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessGetLoadHandlerEvent(fn)
	base.SetEvent(m, 1, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessProcessMessageReceived(fn TOnRenderProcessProcessMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessProcessMessageReceivedEvent(fn)
	base.SetEvent(m, 2, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessWebKitInitialized(fn TOnRenderProcessWebKitInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessWebKitInitializedEvent(fn)
	base.SetEvent(m, 3, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessBrowserCreated(fn TOnRenderProcessBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessBrowserCreatedEvent(fn)
	base.SetEvent(m, 4, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessBrowserDestroyed(fn TOnRenderProcessBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessBrowserDestroyedEvent(fn)
	base.SetEvent(m, 5, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessContextCreated(fn TOnRenderProcessContextCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessContextCreatedEvent(fn)
	base.SetEvent(m, 6, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessContextReleased(fn TOnRenderProcessContextReleasedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessContextReleasedEvent(fn)
	base.SetEvent(m, 7, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessUncaughtException(fn TOnRenderProcessUncaughtExceptionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessUncaughtExceptionEvent(fn)
	base.SetEvent(m, 8, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) SetOnRenderProcessFocusedNodeChanged(fn TOnRenderProcessFocusedNodeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessFocusedNodeChangedEvent(fn)
	base.SetEvent(m, 9, engRenderProcessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderProcessHandler) AsIntfRenderProcessHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngRenderProcessHandler class constructor
func NewEngRenderProcessHandler() IEngRenderProcessHandler {
	var renderProcessHandlerPtr uintptr // ICefRenderProcessHandler
	r := engRenderProcessHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&renderProcessHandlerPtr)))
	ret := AsEngRenderProcessHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, renderProcessHandlerPtr)
	}
	return ret
}

var (
	engRenderProcessHandlerOnce   base.Once
	engRenderProcessHandlerImport *imports.Imports = nil
)

func engRenderProcessHandlerAPI() *imports.Imports {
	engRenderProcessHandlerOnce.Do(func() {
		engRenderProcessHandlerImport = api.NewDefaultImports()
		engRenderProcessHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngRenderProcessHandler_Create", 0), // constructor NewEngRenderProcessHandler
			/* 1 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessGetLoadHandler", 0), // event OnRenderProcessGetLoadHandler
			/* 2 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessProcessMessageReceived", 0), // event OnRenderProcessProcessMessageReceived
			/* 3 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessWebKitInitialized", 0), // event OnRenderProcessWebKitInitialized
			/* 4 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessBrowserCreated", 0), // event OnRenderProcessBrowserCreated
			/* 5 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessBrowserDestroyed", 0), // event OnRenderProcessBrowserDestroyed
			/* 6 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessContextCreated", 0), // event OnRenderProcessContextCreated
			/* 7 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessContextReleased", 0), // event OnRenderProcessContextReleased
			/* 8 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessUncaughtException", 0), // event OnRenderProcessUncaughtException
			/* 9 */ imports.NewTable("TEngRenderProcessHandler_OnRenderProcessFocusedNodeChanged", 0), // event OnRenderProcessFocusedNodeChanged
		}
	})
	return engRenderProcessHandlerImport
}
