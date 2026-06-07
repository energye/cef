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

// IEngSetCookieCallback Parent: ICefSetCookieCallbackOwn
type IEngSetCookieCallback interface {
	ICefSetCookieCallbackOwn
	SetOnSetCookieCallbackComplete(fn TOnSetCookieCallbackCompleteEvent) // property event
	AsIntfSetCookieCallback() uintptr
}

type TEngSetCookieCallback struct {
	TCefSetCookieCallbackOwn
}

func (m *TEngSetCookieCallback) SetOnSetCookieCallbackComplete(fn TOnSetCookieCallbackCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSetCookieCallbackCompleteEvent(fn)
	base.SetEvent(m, 1, engSetCookieCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngSetCookieCallback) AsIntfSetCookieCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngSetCookieCallback class constructor
func NewEngSetCookieCallback() IEngSetCookieCallback {
	var setCookieCallbackPtr uintptr // ICefSetCookieCallback
	r := engSetCookieCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&setCookieCallbackPtr)))
	ret := AsEngSetCookieCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, setCookieCallbackPtr)
	}
	return ret
}

var (
	engSetCookieCallbackOnce   base.Once
	engSetCookieCallbackImport *imports.Imports = nil
)

func engSetCookieCallbackAPI() *imports.Imports {
	engSetCookieCallbackOnce.Do(func() {
		engSetCookieCallbackImport = api.NewDefaultImports()
		engSetCookieCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngSetCookieCallback_Create", 0), // constructor NewEngSetCookieCallback
			/* 1 */ imports.NewTable("TEngSetCookieCallback_OnSetCookieCallbackComplete", 0), // event OnSetCookieCallbackComplete
		}
	})
	return engSetCookieCallbackImport
}
