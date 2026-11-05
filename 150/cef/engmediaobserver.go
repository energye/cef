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

// IEngMediaObserver Parent: ICefMediaObserverOwn
type IEngMediaObserver interface {
	ICefMediaObserverOwn
	SetOnMediaObserverSinks(fn TOnMediaObserverSinksEvent)                               // property event
	SetOnMediaObserverRoutes(fn TOnMediaObserverRoutesEvent)                             // property event
	SetOnMediaObserverRouteStateChanged(fn TOnMediaObserverRouteStateChangedEvent)       // property event
	SetOnMediaObserverRouteMessageReceived(fn TOnMediaObserverRouteMessageReceivedEvent) // property event
	AsIntfMediaObserver() uintptr
}

type TEngMediaObserver struct {
	TCefMediaObserverOwn
}

func (m *TEngMediaObserver) SetOnMediaObserverSinks(fn TOnMediaObserverSinksEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaObserverSinksEvent(fn)
	base.SetEvent(m, 1, engMediaObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaObserver) SetOnMediaObserverRoutes(fn TOnMediaObserverRoutesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaObserverRoutesEvent(fn)
	base.SetEvent(m, 2, engMediaObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaObserver) SetOnMediaObserverRouteStateChanged(fn TOnMediaObserverRouteStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaObserverRouteStateChangedEvent(fn)
	base.SetEvent(m, 3, engMediaObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaObserver) SetOnMediaObserverRouteMessageReceived(fn TOnMediaObserverRouteMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaObserverRouteMessageReceivedEvent(fn)
	base.SetEvent(m, 4, engMediaObserverAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaObserver) AsIntfMediaObserver() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngMediaObserver class constructor
func NewEngMediaObserver() IEngMediaObserver {
	var mediaObserverPtr uintptr // ICefMediaObserver
	r := engMediaObserverAPI().SysCallN(0, uintptr(base.UnsafePointer(&mediaObserverPtr)))
	ret := AsEngMediaObserver(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaObserverPtr)
	}
	return ret
}

var (
	engMediaObserverOnce   base.Once
	engMediaObserverImport *imports.Imports = nil
)

func engMediaObserverAPI() *imports.Imports {
	engMediaObserverOnce.Do(func() {
		engMediaObserverImport = api.NewDefaultImports()
		engMediaObserverImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMediaObserver_Create", 0), // constructor NewEngMediaObserver
			/* 1 */ imports.NewTable("TEngMediaObserver_OnMediaObserverSinks", 0), // event OnMediaObserverSinks
			/* 2 */ imports.NewTable("TEngMediaObserver_OnMediaObserverRoutes", 0), // event OnMediaObserverRoutes
			/* 3 */ imports.NewTable("TEngMediaObserver_OnMediaObserverRouteStateChanged", 0), // event OnMediaObserverRouteStateChanged
			/* 4 */ imports.NewTable("TEngMediaObserver_OnMediaObserverRouteMessageReceived", 0), // event OnMediaObserverRouteMessageReceived
		}
	})
	return engMediaObserverImport
}
