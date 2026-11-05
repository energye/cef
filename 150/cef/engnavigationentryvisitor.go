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

// IEngNavigationEntryVisitor Parent: ICefNavigationEntryVisitorOwn
type IEngNavigationEntryVisitor interface {
	ICefNavigationEntryVisitorOwn
	SetOnNavigationEntryVisitorVisit(fn TOnNavigationEntryVisitorVisitEvent) // property event
	AsIntfNavigationEntryVisitor() uintptr
}

type TEngNavigationEntryVisitor struct {
	TCefNavigationEntryVisitorOwn
}

func (m *TEngNavigationEntryVisitor) SetOnNavigationEntryVisitorVisit(fn TOnNavigationEntryVisitorVisitEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationEntryVisitorVisitEvent(fn)
	base.SetEvent(m, 1, engNavigationEntryVisitorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngNavigationEntryVisitor) AsIntfNavigationEntryVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngNavigationEntryVisitor class constructor
func NewEngNavigationEntryVisitor() IEngNavigationEntryVisitor {
	var navigationEntryVisitorPtr uintptr // ICefNavigationEntryVisitor
	r := engNavigationEntryVisitorAPI().SysCallN(0, uintptr(base.UnsafePointer(&navigationEntryVisitorPtr)))
	ret := AsEngNavigationEntryVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, navigationEntryVisitorPtr)
	}
	return ret
}

var (
	engNavigationEntryVisitorOnce   base.Once
	engNavigationEntryVisitorImport *imports.Imports = nil
)

func engNavigationEntryVisitorAPI() *imports.Imports {
	engNavigationEntryVisitorOnce.Do(func() {
		engNavigationEntryVisitorImport = api.NewDefaultImports()
		engNavigationEntryVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngNavigationEntryVisitor_Create", 0), // constructor NewEngNavigationEntryVisitor
			/* 1 */ imports.NewTable("TEngNavigationEntryVisitor_OnNavigationEntryVisitorVisit", 0), // event OnNavigationEntryVisitorVisit
		}
	})
	return engNavigationEntryVisitorImport
}
