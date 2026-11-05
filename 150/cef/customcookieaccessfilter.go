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

// ICustomCookieAccessFilter Parent: ICefCookieAccessFilterOwn
type ICustomCookieAccessFilter interface {
	ICefCookieAccessFilterOwn
	RemoveReferences() // procedure
	AsIntfCookieAccessFilter() uintptr
}

type TCustomCookieAccessFilter struct {
	TCefCookieAccessFilterOwn
}

func (m *TCustomCookieAccessFilter) RemoveReferences() {
	if !m.IsValid() {
		return
	}
	customCookieAccessFilterAPI().SysCallN(1, m.Instance())
}

func (m *TCustomCookieAccessFilter) AsIntfCookieAccessFilter() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCookieAccessFilter class constructor
func NewCustomCookieAccessFilter(events IChromiumCore) ICustomCookieAccessFilter {
	var cookieAccessFilterPtr uintptr // ICefCookieAccessFilter
	r := customCookieAccessFilterAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&cookieAccessFilterPtr)))
	ret := AsCustomCookieAccessFilter(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, cookieAccessFilterPtr)
	}
	return ret
}

var (
	customCookieAccessFilterOnce   base.Once
	customCookieAccessFilterImport *imports.Imports = nil
)

func customCookieAccessFilterAPI() *imports.Imports {
	customCookieAccessFilterOnce.Do(func() {
		customCookieAccessFilterImport = api.NewDefaultImports()
		customCookieAccessFilterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCookieAccessFilter_Create", 0), // constructor NewCustomCookieAccessFilter
			/* 1 */ imports.NewTable("TCustomCookieAccessFilter_RemoveReferences", 0), // procedure RemoveReferences
		}
	})
	return customCookieAccessFilterImport
}
