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

// IEngBrowserViewDelegate Parent: ICefBrowserViewDelegateOwn
type IEngBrowserViewDelegate interface {
	ICefBrowserViewDelegateOwn
	SetOnBrowserViewGetChromeToolbarType(fn TOnBrowserViewGetChromeToolbarTypeEvent)                     // property event
	SetOnBrowserViewBrowserCreated(fn TOnBrowserViewBrowserCreatedEvent)                                 // property event
	SetOnBrowserViewBrowserDestroyed(fn TOnBrowserViewBrowserDestroyedEvent)                             // property event
	SetOnBrowserViewGetDelegateForPopupBrowserView(fn TOnBrowserViewGetDelegateForPopupBrowserViewEvent) // property event
	SetOnBrowserViewPopupBrowserViewCreated(fn TOnBrowserViewPopupBrowserViewCreatedEvent)               // property event
	AsIntfBrowserViewDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngBrowserViewDelegate struct {
	TCefBrowserViewDelegateOwn
}

func (m *TEngBrowserViewDelegate) SetOnBrowserViewGetChromeToolbarType(fn TOnBrowserViewGetChromeToolbarTypeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserViewGetChromeToolbarTypeEvent(fn)
	base.SetEvent(m, 1, engBrowserViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserViewDelegate) SetOnBrowserViewBrowserCreated(fn TOnBrowserViewBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserViewBrowserCreatedEvent(fn)
	base.SetEvent(m, 2, engBrowserViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserViewDelegate) SetOnBrowserViewBrowserDestroyed(fn TOnBrowserViewBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserViewBrowserDestroyedEvent(fn)
	base.SetEvent(m, 3, engBrowserViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserViewDelegate) SetOnBrowserViewGetDelegateForPopupBrowserView(fn TOnBrowserViewGetDelegateForPopupBrowserViewEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserViewGetDelegateForPopupBrowserViewEvent(fn)
	base.SetEvent(m, 4, engBrowserViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserViewDelegate) SetOnBrowserViewPopupBrowserViewCreated(fn TOnBrowserViewPopupBrowserViewCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserViewPopupBrowserViewCreatedEvent(fn)
	base.SetEvent(m, 5, engBrowserViewDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngBrowserViewDelegate) AsIntfBrowserViewDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngBrowserViewDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(1)
}

// NewEngBrowserViewDelegate class constructor
func NewEngBrowserViewDelegate() IEngBrowserViewDelegate {
	var browserViewDelegatePtr uintptr // ICefBrowserViewDelegate
	var viewDelegatePtr uintptr        // ICefViewDelegate
	r := engBrowserViewDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&browserViewDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngBrowserViewDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewDelegatePtr)
		intf.SetIntfPointer(1, viewDelegatePtr)
	}
	return ret
}

var (
	engBrowserViewDelegateOnce   base.Once
	engBrowserViewDelegateImport *imports.Imports = nil
)

func engBrowserViewDelegateAPI() *imports.Imports {
	engBrowserViewDelegateOnce.Do(func() {
		engBrowserViewDelegateImport = api.NewDefaultImports()
		engBrowserViewDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngBrowserViewDelegate_Create", 0), // constructor NewEngBrowserViewDelegate
			/* 1 */ imports.NewTable("TEngBrowserViewDelegate_OnBrowserViewGetChromeToolbarType", 0), // event OnBrowserViewGetChromeToolbarType
			/* 2 */ imports.NewTable("TEngBrowserViewDelegate_OnBrowserViewBrowserCreated", 0), // event OnBrowserViewBrowserCreated
			/* 3 */ imports.NewTable("TEngBrowserViewDelegate_OnBrowserViewBrowserDestroyed", 0), // event OnBrowserViewBrowserDestroyed
			/* 4 */ imports.NewTable("TEngBrowserViewDelegate_OnBrowserViewGetDelegateForPopupBrowserView", 0), // event OnBrowserViewGetDelegateForPopupBrowserView
			/* 5 */ imports.NewTable("TEngBrowserViewDelegate_OnBrowserViewPopupBrowserViewCreated", 0), // event OnBrowserViewPopupBrowserViewCreated
		}
	})
	return engBrowserViewDelegateImport
}
