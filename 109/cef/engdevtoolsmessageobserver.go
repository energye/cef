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

// IEngDevToolsMessageObserver Parent: ICEFDevToolsMessageObserverOwn
type IEngDevToolsMessageObserver interface {
	ICEFDevToolsMessageObserverOwn
	SetOnDevToolsMessageObserverDevToolsMessage(fn TOnDevToolsMessageObserverDevToolsMessageEvent)             // property event
	SetOnDevToolsMessageObserverDevToolsMethodResult(fn TOnDevToolsMessageObserverDevToolsMethodResultEvent)   // property event
	SetOnDevToolsMessageObserverDevTools(fn TOnDevToolsMessageObserverDevToolsEvent)                           // property event
	SetOnDevToolsMessageObserverDevToolsAgentAttached(fn TOnDevToolsMessageObserverDevToolsAgentAttachedEvent) // property event
	SetOnDevToolsMessageObserverDevToolsAgentDetached(fn TOnDevToolsMessageObserverDevToolsAgentDetachedEvent) // property event
	AsIntfDevToolsMessageObserver() uintptr
}

type TEngDevToolsMessageObserver struct {
	TCEFDevToolsMessageObserverOwn
}

func (m *TEngDevToolsMessageObserver) SetOnDevToolsMessageObserverDevToolsMessage(fn TOnDevToolsMessageObserverDevToolsMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageObserverDevToolsMessageEvent(fn)
	base.SetEvent(m, 1, engDevToolsMessageObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDevToolsMessageObserver) SetOnDevToolsMessageObserverDevToolsMethodResult(fn TOnDevToolsMessageObserverDevToolsMethodResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageObserverDevToolsMethodResultEvent(fn)
	base.SetEvent(m, 2, engDevToolsMessageObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDevToolsMessageObserver) SetOnDevToolsMessageObserverDevTools(fn TOnDevToolsMessageObserverDevToolsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageObserverDevToolsEvent(fn)
	base.SetEvent(m, 3, engDevToolsMessageObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDevToolsMessageObserver) SetOnDevToolsMessageObserverDevToolsAgentAttached(fn TOnDevToolsMessageObserverDevToolsAgentAttachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageObserverDevToolsAgentAttachedEvent(fn)
	base.SetEvent(m, 4, engDevToolsMessageObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDevToolsMessageObserver) SetOnDevToolsMessageObserverDevToolsAgentDetached(fn TOnDevToolsMessageObserverDevToolsAgentDetachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageObserverDevToolsAgentDetachedEvent(fn)
	base.SetEvent(m, 5, engDevToolsMessageObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDevToolsMessageObserver) AsIntfDevToolsMessageObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDevToolsMessageObserver class constructor
func NewEngDevToolsMessageObserver() IEngDevToolsMessageObserver {
	var devToolsMessageObserverPtr uintptr // ICefDevToolsMessageObserver
	r := engDevToolsMessageObserverAPI().SysCallN(0, uintptr(base.UnsafePointer(&devToolsMessageObserverPtr)))
	ret := AsEngDevToolsMessageObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, devToolsMessageObserverPtr)
	}
	return ret
}

var (
	engDevToolsMessageObserverOnce   base.Once
	engDevToolsMessageObserverImport *imports.Imports = nil
)

func engDevToolsMessageObserverAPI() *imports.Imports {
	engDevToolsMessageObserverOnce.Do(func() {
		engDevToolsMessageObserverImport = api.NewDefaultImports()
		engDevToolsMessageObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDevToolsMessageObserver_Create", 0), // constructor NewEngDevToolsMessageObserver
			/* 1 */ imports.NewTable("TEngDevToolsMessageObserver_OnDevToolsMessageObserverDevToolsMessage", 0), // event OnDevToolsMessageObserverDevToolsMessage
			/* 2 */ imports.NewTable("TEngDevToolsMessageObserver_OnDevToolsMessageObserverDevToolsMethodResult", 0), // event OnDevToolsMessageObserverDevToolsMethodResult
			/* 3 */ imports.NewTable("TEngDevToolsMessageObserver_OnDevToolsMessageObserverDevTools", 0), // event OnDevToolsMessageObserverDevTools
			/* 4 */ imports.NewTable("TEngDevToolsMessageObserver_OnDevToolsMessageObserverDevToolsAgentAttached", 0), // event OnDevToolsMessageObserverDevToolsAgentAttached
			/* 5 */ imports.NewTable("TEngDevToolsMessageObserver_OnDevToolsMessageObserverDevToolsAgentDetached", 0), // event OnDevToolsMessageObserverDevToolsAgentDetached
		}
	})
	return engDevToolsMessageObserverImport
}
