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

// IEngResolveCallback Parent: ICefResolveCallbackOwn
type IEngResolveCallback interface {
	ICefResolveCallbackOwn
	SetOnResolveCallbackResolveCompleted(fn TOnResolveCallbackResolveCompletedEvent) // property event
	AsIntfResolveCallback() uintptr
}

type TEngResolveCallback struct {
	TCefResolveCallbackOwn
}

func (m *TEngResolveCallback) SetOnResolveCallbackResolveCompleted(fn TOnResolveCallbackResolveCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResolveCallbackResolveCompletedEvent(fn)
	base.SetEvent(m, 1, engResolveCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResolveCallback) AsIntfResolveCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngResolveCallback class constructor
func NewEngResolveCallback() IEngResolveCallback {
	var resolveCallbackPtr uintptr // ICefResolveCallback
	r := engResolveCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&resolveCallbackPtr)))
	ret := AsEngResolveCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resolveCallbackPtr)
	}
	return ret
}

var (
	engResolveCallbackOnce   base.Once
	engResolveCallbackImport *imports.Imports = nil
)

func engResolveCallbackAPI() *imports.Imports {
	engResolveCallbackOnce.Do(func() {
		engResolveCallbackImport = api.NewDefaultImports()
		engResolveCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngResolveCallback_Create", 0), // constructor NewEngResolveCallback
			/* 1 */ imports.NewTable("TEngResolveCallback_OnResolveCallbackResolveCompleted", 0), // event OnResolveCallbackResolveCompleted
		}
	})
	return engResolveCallbackImport
}
