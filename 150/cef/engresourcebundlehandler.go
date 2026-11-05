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

// IEngResourceBundleHandler Parent: ICefResourceBundleHandlerOwn
type IEngResourceBundleHandler interface {
	ICefResourceBundleHandlerOwn
	SetOnResourceBundleGetLocalizedString(fn TOnResourceBundleGetLocalizedStringEvent)           // property event
	SetOnResourceBundleGetDataResource(fn TOnResourceBundleGetDataResourceEvent)                 // property event
	SetOnResourceBundleGetDataResourceForScale(fn TOnResourceBundleGetDataResourceForScaleEvent) // property event
	AsIntfResourceBundleHandler() uintptr
}

type TEngResourceBundleHandler struct {
	TCefResourceBundleHandlerOwn
}

func (m *TEngResourceBundleHandler) SetOnResourceBundleGetLocalizedString(fn TOnResourceBundleGetLocalizedStringEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceBundleGetLocalizedStringEvent(fn)
	base.SetEvent(m, 1, engResourceBundleHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceBundleHandler) SetOnResourceBundleGetDataResource(fn TOnResourceBundleGetDataResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceBundleGetDataResourceEvent(fn)
	base.SetEvent(m, 2, engResourceBundleHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceBundleHandler) SetOnResourceBundleGetDataResourceForScale(fn TOnResourceBundleGetDataResourceForScaleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceBundleGetDataResourceForScaleEvent(fn)
	base.SetEvent(m, 3, engResourceBundleHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngResourceBundleHandler) AsIntfResourceBundleHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngResourceBundleHandler class constructor
func NewEngResourceBundleHandler() IEngResourceBundleHandler {
	var resourceBundleHandlerPtr uintptr // ICefResourceBundleHandler
	r := engResourceBundleHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&resourceBundleHandlerPtr)))
	ret := AsEngResourceBundleHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceBundleHandlerPtr)
	}
	return ret
}

var (
	engResourceBundleHandlerOnce   base.Once
	engResourceBundleHandlerImport *imports.Imports = nil
)

func engResourceBundleHandlerAPI() *imports.Imports {
	engResourceBundleHandlerOnce.Do(func() {
		engResourceBundleHandlerImport = api.NewDefaultImports()
		engResourceBundleHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngResourceBundleHandler_Create", 0), // constructor NewEngResourceBundleHandler
			/* 1 */ imports.NewTable("TEngResourceBundleHandler_OnResourceBundleGetLocalizedString", 0), // event OnResourceBundleGetLocalizedString
			/* 2 */ imports.NewTable("TEngResourceBundleHandler_OnResourceBundleGetDataResource", 0), // event OnResourceBundleGetDataResource
			/* 3 */ imports.NewTable("TEngResourceBundleHandler_OnResourceBundleGetDataResourceForScale", 0), // event OnResourceBundleGetDataResourceForScale
		}
	})
	return engResourceBundleHandlerImport
}
