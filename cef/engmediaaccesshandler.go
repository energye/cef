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

// IEngMediaAccessHandler Parent: ICefMediaAccessHandlerOwn
type IEngMediaAccessHandler interface {
	ICefMediaAccessHandlerOwn
	SetOnMediaAccessRequestMediaAccessPermission(fn TOnMediaAccessRequestMediaAccessPermissionEvent) // property event
	AsIntfMediaAccessHandler() uintptr
}

type TEngMediaAccessHandler struct {
	TCefMediaAccessHandlerOwn
}

func (m *TEngMediaAccessHandler) SetOnMediaAccessRequestMediaAccessPermission(fn TOnMediaAccessRequestMediaAccessPermissionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaAccessRequestMediaAccessPermissionEvent(fn)
	base.SetEvent(m, 1, engMediaAccessHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaAccessHandler) AsIntfMediaAccessHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngMediaAccessHandler class constructor
func NewEngMediaAccessHandler() IEngMediaAccessHandler {
	var mediaAccessHandlerPtr uintptr // ICefMediaAccessHandler
	r := engMediaAccessHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&mediaAccessHandlerPtr)))
	ret := AsEngMediaAccessHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaAccessHandlerPtr)
	}
	return ret
}

var (
	engMediaAccessHandlerOnce   base.Once
	engMediaAccessHandlerImport *imports.Imports = nil
)

func engMediaAccessHandlerAPI() *imports.Imports {
	engMediaAccessHandlerOnce.Do(func() {
		engMediaAccessHandlerImport = api.NewDefaultImports()
		engMediaAccessHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMediaAccessHandler_Create", 0), // constructor NewEngMediaAccessHandler
			/* 1 */ imports.NewTable("TEngMediaAccessHandler_OnMediaAccessRequestMediaAccessPermission", 0), // event OnMediaAccessRequestMediaAccessPermission
		}
	})
	return engMediaAccessHandlerImport
}
