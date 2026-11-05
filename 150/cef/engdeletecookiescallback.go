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

// IEngDeleteCookiesCallback Parent: ICefDeleteCookiesCallbackOwn
type IEngDeleteCookiesCallback interface {
	ICefDeleteCookiesCallbackOwn
	SetOnDeleteCookiesCallbackComplete(fn TOnDeleteCookiesCallbackCompleteEvent) // property event
	AsIntfDeleteCookiesCallback() uintptr
}

type TEngDeleteCookiesCallback struct {
	TCefDeleteCookiesCallbackOwn
}

func (m *TEngDeleteCookiesCallback) SetOnDeleteCookiesCallbackComplete(fn TOnDeleteCookiesCallbackCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDeleteCookiesCallbackCompleteEvent(fn)
	base.SetEvent(m, 1, engDeleteCookiesCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDeleteCookiesCallback) AsIntfDeleteCookiesCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDeleteCookiesCallback class constructor
func NewEngDeleteCookiesCallback() IEngDeleteCookiesCallback {
	var deleteCookiesCallbackPtr uintptr // ICefDeleteCookiesCallback
	r := engDeleteCookiesCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&deleteCookiesCallbackPtr)))
	ret := AsEngDeleteCookiesCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, deleteCookiesCallbackPtr)
	}
	return ret
}

var (
	engDeleteCookiesCallbackOnce   base.Once
	engDeleteCookiesCallbackImport *imports.Imports = nil
)

func engDeleteCookiesCallbackAPI() *imports.Imports {
	engDeleteCookiesCallbackOnce.Do(func() {
		engDeleteCookiesCallbackImport = api.NewDefaultImports()
		engDeleteCookiesCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDeleteCookiesCallback_Create", 0), // constructor NewEngDeleteCookiesCallback
			/* 1 */ imports.NewTable("TEngDeleteCookiesCallback_OnDeleteCookiesCallbackComplete", 0), // event OnDeleteCookiesCallbackComplete
		}
	})
	return engDeleteCookiesCallbackImport
}
