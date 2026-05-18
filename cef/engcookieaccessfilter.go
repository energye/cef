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

// IEngCookieAccessFilter Parent: ICefCookieAccessFilterOwn
type IEngCookieAccessFilter interface {
	ICefCookieAccessFilterOwn
	SetOnCookieAccessFilterCanSendCookie(fn TOnCookieAccessFilterCanSendCookieEvent) // property event
	SetOnCookieAccessFilterCanSaveCookie(fn TOnCookieAccessFilterCanSaveCookieEvent) // property event
	AsIntfCookieAccessFilter() uintptr
}

type TEngCookieAccessFilter struct {
	TCefCookieAccessFilterOwn
}

func (m *TEngCookieAccessFilter) SetOnCookieAccessFilterCanSendCookie(fn TOnCookieAccessFilterCanSendCookieEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieAccessFilterCanSendCookieEvent(fn)
	base.SetEvent(m, 1, engCookieAccessFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCookieAccessFilter) SetOnCookieAccessFilterCanSaveCookie(fn TOnCookieAccessFilterCanSaveCookieEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieAccessFilterCanSaveCookieEvent(fn)
	base.SetEvent(m, 2, engCookieAccessFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCookieAccessFilter) AsIntfCookieAccessFilter() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngCookieAccessFilter class constructor
func NewEngCookieAccessFilter() IEngCookieAccessFilter {
	var cookieAccessFilterPtr uintptr // ICefCookieAccessFilter
	r := engCookieAccessFilterAPI().SysCallN(0, uintptr(base.UnsafePointer(&cookieAccessFilterPtr)))
	ret := AsEngCookieAccessFilter(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieAccessFilterPtr)
	}
	return ret
}

var (
	engCookieAccessFilterOnce   base.Once
	engCookieAccessFilterImport *imports.Imports = nil
)

func engCookieAccessFilterAPI() *imports.Imports {
	engCookieAccessFilterOnce.Do(func() {
		engCookieAccessFilterImport = api.NewDefaultImports()
		engCookieAccessFilterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngCookieAccessFilter_Create", 0), // constructor NewEngCookieAccessFilter
			/* 1 */ imports.NewTable("TEngCookieAccessFilter_OnCookieAccessFilterCanSendCookie", 0), // event OnCookieAccessFilterCanSendCookie
			/* 2 */ imports.NewTable("TEngCookieAccessFilter_OnCookieAccessFilterCanSaveCookie", 0), // event OnCookieAccessFilterCanSaveCookie
		}
	})
	return engCookieAccessFilterImport
}
