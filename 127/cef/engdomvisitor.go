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

// IEngDomVisitor Parent: ICefDomVisitorOwn
type IEngDomVisitor interface {
	ICefDomVisitorOwn
	SetOnDomVisitorVisit(fn TOnDomVisitorVisitEvent) // property event
	AsIntfDomVisitor() uintptr
}

type TEngDomVisitor struct {
	TCefDomVisitorOwn
}

func (m *TEngDomVisitor) SetOnDomVisitorVisit(fn TOnDomVisitorVisitEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDomVisitorVisitEvent(fn)
	base.SetEvent(m, 1, engDomVisitorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDomVisitor) AsIntfDomVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDomVisitor class constructor
func NewEngDomVisitor() IEngDomVisitor {
	var domVisitorPtr uintptr // ICefDomVisitor
	r := engDomVisitorAPI().SysCallN(0, uintptr(base.UnsafePointer(&domVisitorPtr)))
	ret := AsEngDomVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, domVisitorPtr)
	}
	return ret
}

var (
	engDomVisitorOnce   base.Once
	engDomVisitorImport *imports.Imports = nil
)

func engDomVisitorAPI() *imports.Imports {
	engDomVisitorOnce.Do(func() {
		engDomVisitorImport = api.NewDefaultImports()
		engDomVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDomVisitor_Create", 0), // constructor NewEngDomVisitor
			/* 1 */ imports.NewTable("TEngDomVisitor_OnDomVisitorVisit", 0), // event OnDomVisitorVisit
		}
	})
	return engDomVisitorImport
}
