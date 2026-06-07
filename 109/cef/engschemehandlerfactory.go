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

	cefTypes "github.com/energye/cef/109/types"
)

// IEngSchemeHandlerFactory Parent: ICefSchemeHandlerFactoryOwn
type IEngSchemeHandlerFactory interface {
	ICefSchemeHandlerFactoryOwn
	SetOnSchemeFactoryNew(fn TOnSchemeFactoryNewEvent) // property event
	AsIntfSchemeHandlerFactory() uintptr
}

type TEngSchemeHandlerFactory struct {
	TCefSchemeHandlerFactoryOwn
}

func (m *TEngSchemeHandlerFactory) SetOnSchemeFactoryNew(fn TOnSchemeFactoryNewEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSchemeFactoryNewEvent(fn)
	base.SetEvent(m, 1, engSchemeHandlerFactoryAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngSchemeHandlerFactory) AsIntfSchemeHandlerFactory() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngSchemeHandlerFactory class constructor
func NewEngSchemeHandlerFactory(class cefTypes.TCefResourceHandlerClass) IEngSchemeHandlerFactory {
	var schemeHandlerFactoryPtr uintptr // ICefSchemeHandlerFactory
	r := engSchemeHandlerFactoryAPI().SysCallN(0, uintptr(class), uintptr(base.UnsafePointer(&schemeHandlerFactoryPtr)))
	ret := AsEngSchemeHandlerFactory(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, schemeHandlerFactoryPtr)
	}
	return ret
}

var (
	engSchemeHandlerFactoryOnce   base.Once
	engSchemeHandlerFactoryImport *imports.Imports = nil
)

func engSchemeHandlerFactoryAPI() *imports.Imports {
	engSchemeHandlerFactoryOnce.Do(func() {
		engSchemeHandlerFactoryImport = api.NewDefaultImports()
		engSchemeHandlerFactoryImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngSchemeHandlerFactory_Create", 0), // constructor NewEngSchemeHandlerFactory
			/* 1 */ imports.NewTable("TEngSchemeHandlerFactory_OnSchemeFactoryNew", 0), // event OnSchemeFactoryNew
		}
	})
	return engSchemeHandlerFactoryImport
}
