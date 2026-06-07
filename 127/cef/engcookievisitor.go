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

// IEngCookieVisitor Parent: ICefCookieVisitorOwn
type IEngCookieVisitor interface {
	ICefCookieVisitorOwn
	SetOnCookieVisitorVisit(fn TOnCookieVisitorVisitEvent) // property event
	AsIntfCookieVisitor() uintptr
}

type TEngCookieVisitor struct {
	TCefCookieVisitorOwn
}

func (m *TEngCookieVisitor) SetOnCookieVisitorVisit(fn TOnCookieVisitorVisitEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieVisitorVisitEvent(fn)
	base.SetEvent(m, 1, engCookieVisitorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCookieVisitor) AsIntfCookieVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngCookieVisitor class constructor
func NewEngCookieVisitor() IEngCookieVisitor {
	var cookieVisitorPtr uintptr // ICefCookieVisitor
	r := engCookieVisitorAPI().SysCallN(0, uintptr(base.UnsafePointer(&cookieVisitorPtr)))
	ret := AsEngCookieVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieVisitorPtr)
	}
	return ret
}

var (
	engCookieVisitorOnce   base.Once
	engCookieVisitorImport *imports.Imports = nil
)

func engCookieVisitorAPI() *imports.Imports {
	engCookieVisitorOnce.Do(func() {
		engCookieVisitorImport = api.NewDefaultImports()
		engCookieVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngCookieVisitor_Create", 0), // constructor NewEngCookieVisitor
			/* 1 */ imports.NewTable("TEngCookieVisitor_OnCookieVisitorVisit", 0), // event OnCookieVisitorVisit
		}
	})
	return engCookieVisitorImport
}
