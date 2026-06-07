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

// ICustomResponseFilter Parent: ICefResponseFilterOwn
type ICustomResponseFilter interface {
	ICefResponseFilterOwn
	SetOnFilter(fn TOnFilterEvent)         // property event
	SetOnInitFilter(fn TOnInitFilterEvent) // property event
	AsIntfResponseFilter() uintptr
}

type TCustomResponseFilter struct {
	TCefResponseFilterOwn
}

func (m *TCustomResponseFilter) SetOnFilter(fn TOnFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFilterEvent(fn)
	base.SetEvent(m, 1, customResponseFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomResponseFilter) SetOnInitFilter(fn TOnInitFilterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnInitFilterEvent(fn)
	base.SetEvent(m, 2, customResponseFilterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomResponseFilter) AsIntfResponseFilter() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomResponseFilter class constructor
func NewCustomResponseFilter() ICustomResponseFilter {
	var responseFilterPtr uintptr // ICefResponseFilter
	r := customResponseFilterAPI().SysCallN(0, uintptr(base.UnsafePointer(&responseFilterPtr)))
	ret := AsCustomResponseFilter(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, responseFilterPtr)
	}
	return ret
}

var (
	customResponseFilterOnce   base.Once
	customResponseFilterImport *imports.Imports = nil
)

func customResponseFilterAPI() *imports.Imports {
	customResponseFilterOnce.Do(func() {
		customResponseFilterImport = api.NewDefaultImports()
		customResponseFilterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomResponseFilter_Create", 0), // constructor NewCustomResponseFilter
			/* 1 */ imports.NewTable("TCustomResponseFilter_OnFilter", 0), // event OnFilter
			/* 2 */ imports.NewTable("TCustomResponseFilter_OnInitFilter", 0), // event OnInitFilter
		}
	})
	return customResponseFilterImport
}
