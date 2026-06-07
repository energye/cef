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

// IEngResponseFilter Parent: ICefResponseFilterOwn
type IEngResponseFilter interface {
	ICefResponseFilterOwn
	SetOnResponseFilterInitFilter(fn TOnResponseFilterInitFilterEvent) // property event
	SetOnResponseFilterFilter(fn TOnResponseFilterFilterEvent)         // property event
	AsIntfResponseFilter() uintptr
}

type TEngResponseFilter struct {
	TCefResponseFilterOwn
}

func (m *TEngResponseFilter) SetOnResponseFilterInitFilter(fn TOnResponseFilterInitFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResponseFilterInitFilterEvent(fn)
	base.SetEvent(m, 1, engResponseFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResponseFilter) SetOnResponseFilterFilter(fn TOnResponseFilterFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResponseFilterFilterEvent(fn)
	base.SetEvent(m, 2, engResponseFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResponseFilter) AsIntfResponseFilter() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngResponseFilter class constructor
func NewEngResponseFilter() IEngResponseFilter {
	var responseFilterPtr uintptr // ICefResponseFilter
	r := engResponseFilterAPI().SysCallN(0, uintptr(base.UnsafePointer(&responseFilterPtr)))
	ret := AsEngResponseFilter(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, responseFilterPtr)
	}
	return ret
}

var (
	engResponseFilterOnce   base.Once
	engResponseFilterImport *imports.Imports = nil
)

func engResponseFilterAPI() *imports.Imports {
	engResponseFilterOnce.Do(func() {
		engResponseFilterImport = api.NewDefaultImports()
		engResponseFilterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngResponseFilter_Create", 0), // constructor NewEngResponseFilter
			/* 1 */ imports.NewTable("TEngResponseFilter_OnResponseFilterInitFilter", 0), // event OnResponseFilterInitFilter
			/* 2 */ imports.NewTable("TEngResponseFilter_OnResponseFilterFilter", 0), // event OnResponseFilterFilter
		}
	})
	return engResponseFilterImport
}
