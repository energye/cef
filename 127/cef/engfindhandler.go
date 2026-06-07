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

// IEngFindHandler Parent: ICefFindHandlerOwn
type IEngFindHandler interface {
	ICefFindHandlerOwn
	SetOnFindFindResult(fn TOnFindFindResultEvent) // property event
	AsIntfFindHandler() uintptr
}

type TEngFindHandler struct {
	TCefFindHandlerOwn
}

func (m *TEngFindHandler) SetOnFindFindResult(fn TOnFindFindResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFindFindResultEvent(fn)
	base.SetEvent(m, 1, engFindHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngFindHandler) AsIntfFindHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngFindHandler class constructor
func NewEngFindHandler() IEngFindHandler {
	var findHandlerPtr uintptr // ICefFindHandler
	r := engFindHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&findHandlerPtr)))
	ret := AsEngFindHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, findHandlerPtr)
	}
	return ret
}

var (
	engFindHandlerOnce   base.Once
	engFindHandlerImport *imports.Imports = nil
)

func engFindHandlerAPI() *imports.Imports {
	engFindHandlerOnce.Do(func() {
		engFindHandlerImport = api.NewDefaultImports()
		engFindHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngFindHandler_Create", 0), // constructor NewEngFindHandler
			/* 1 */ imports.NewTable("TEngFindHandler_OnFindFindResult", 0), // event OnFindFindResult
		}
	})
	return engFindHandlerImport
}
