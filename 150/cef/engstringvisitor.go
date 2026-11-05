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

// IEngStringVisitor Parent: ICefStringVisitorOwn
type IEngStringVisitor interface {
	ICefStringVisitorOwn
	SetOnStringVisitorVisit(fn TOnStringVisitorVisitEvent) // property event
	AsIntfStringVisitor() uintptr
}

type TEngStringVisitor struct {
	TCefStringVisitorOwn
}

func (m *TEngStringVisitor) SetOnStringVisitorVisit(fn TOnStringVisitorVisitEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStringVisitorVisitEvent(fn)
	base.SetEvent(m, 1, engStringVisitorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngStringVisitor) AsIntfStringVisitor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngStringVisitor class constructor
func NewEngStringVisitor() IEngStringVisitor {
	var stringVisitorPtr uintptr // ICefStringVisitor
	r := engStringVisitorAPI().SysCallN(0, uintptr(base.UnsafePointer(&stringVisitorPtr)))
	ret := AsEngStringVisitor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, stringVisitorPtr)
	}
	return ret
}

var (
	engStringVisitorOnce   base.Once
	engStringVisitorImport *imports.Imports = nil
)

func engStringVisitorAPI() *imports.Imports {
	engStringVisitorOnce.Do(func() {
		engStringVisitorImport = api.NewDefaultImports()
		engStringVisitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngStringVisitor_Create", 0), // constructor NewEngStringVisitor
			/* 1 */ imports.NewTable("TEngStringVisitor_OnStringVisitorVisit", 0), // event OnStringVisitorVisit
		}
	})
	return engStringVisitorImport
}
