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

// IEngPermissionHandler Parent: ICefPermissionHandlerOwn
type IEngPermissionHandler interface {
	ICefPermissionHandlerOwn
	SetOnPermissionRequestMediaAccessPermission(fn TOnPermissionRequestMediaAccessPermissionEvent) // property event
	SetOnPermissionShowPermissionPrompt(fn TOnPermissionShowPermissionPromptEvent)                 // property event
	SetOnPermissionDismissPermissionPrompt(fn TOnPermissionDismissPermissionPromptEvent)           // property event
	AsIntfPermissionHandler() uintptr
}

type TEngPermissionHandler struct {
	TCefPermissionHandlerOwn
}

func (m *TEngPermissionHandler) SetOnPermissionRequestMediaAccessPermission(fn TOnPermissionRequestMediaAccessPermissionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPermissionRequestMediaAccessPermissionEvent(fn)
	base.SetEvent(m, 1, engPermissionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPermissionHandler) SetOnPermissionShowPermissionPrompt(fn TOnPermissionShowPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPermissionShowPermissionPromptEvent(fn)
	base.SetEvent(m, 2, engPermissionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPermissionHandler) SetOnPermissionDismissPermissionPrompt(fn TOnPermissionDismissPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPermissionDismissPermissionPromptEvent(fn)
	base.SetEvent(m, 3, engPermissionHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngPermissionHandler) AsIntfPermissionHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngPermissionHandler class constructor
func NewEngPermissionHandler() IEngPermissionHandler {
	var permissionHandlerPtr uintptr // ICefPermissionHandler
	r := engPermissionHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&permissionHandlerPtr)))
	ret := AsEngPermissionHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, permissionHandlerPtr)
	}
	return ret
}

var (
	engPermissionHandlerOnce   base.Once
	engPermissionHandlerImport *imports.Imports = nil
)

func engPermissionHandlerAPI() *imports.Imports {
	engPermissionHandlerOnce.Do(func() {
		engPermissionHandlerImport = api.NewDefaultImports()
		engPermissionHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngPermissionHandler_Create", 0), // constructor NewEngPermissionHandler
			/* 1 */ imports.NewTable("TEngPermissionHandler_OnPermissionRequestMediaAccessPermission", 0), // event OnPermissionRequestMediaAccessPermission
			/* 2 */ imports.NewTable("TEngPermissionHandler_OnPermissionShowPermissionPrompt", 0), // event OnPermissionShowPermissionPrompt
			/* 3 */ imports.NewTable("TEngPermissionHandler_OnPermissionDismissPermissionPrompt", 0), // event OnPermissionDismissPermissionPrompt
		}
	})
	return engPermissionHandlerImport
}
