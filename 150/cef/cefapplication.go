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

// ICefApplication Parent: ICefApplicationCore
type ICefApplication interface {
	ICefApplicationCore
	DestroyApplicationObject() bool         // property DestroyApplicationObject Getter
	SetDestroyApplicationObject(value bool) // property DestroyApplicationObject Setter
	DestroyAppWindows() bool                // property DestroyAppWindows Getter
	SetDestroyAppWindows(value bool)        // property DestroyAppWindows Setter
	AsIntfApplicationCoreEvents() uintptr
}

type TCefApplication struct {
	TCefApplicationCore
}

func (m *TCefApplication) DestroyApplicationObject() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplication) SetDestroyApplicationObject(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplication) DestroyAppWindows() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplication) SetDestroyAppWindows(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplication) AsIntfApplicationCoreEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewApplication class constructor
func NewApplication() ICefApplication {
	var applicationCoreEventsPtr uintptr // IApplicationCoreEvents
	r := cefApplicationAPI().SysCallN(0, uintptr(base.UnsafePointer(&applicationCoreEventsPtr)))
	ret := AsCefApplication(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, applicationCoreEventsPtr)
	}
	return ret
}

var (
	cefApplicationOnce   base.Once
	cefApplicationImport *imports.Imports = nil
)

func cefApplicationAPI() *imports.Imports {
	cefApplicationOnce.Do(func() {
		cefApplicationImport = api.NewDefaultImports()
		cefApplicationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefApplication_Create", 0), // constructor NewApplication
			/* 1 */ imports.NewTable("TCefApplication_DestroyApplicationObject", 0), // property DestroyApplicationObject
			/* 2 */ imports.NewTable("TCefApplication_DestroyAppWindows", 0), // property DestroyAppWindows
		}
	})
	return cefApplicationImport
}
