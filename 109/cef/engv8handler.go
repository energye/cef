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

// IEngV8Handler Parent: ICefv8HandlerOwn
type IEngV8Handler interface {
	ICefv8HandlerOwn
	SetOnV8Execute(fn TOnV8ExecuteEvent) // property event
	AsIntfV8Handler() uintptr
}

type TEngV8Handler struct {
	TCefv8HandlerOwn
}

func (m *TEngV8Handler) SetOnV8Execute(fn TOnV8ExecuteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8ExecuteEvent(fn)
	base.SetEvent(m, 1, engV8HandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Handler) AsIntfV8Handler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngV8Handler class constructor
func NewEngV8Handler() IEngV8Handler {
	var v8HandlerPtr uintptr // ICefv8Handler
	r := engV8HandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&v8HandlerPtr)))
	ret := AsEngV8Handler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8HandlerPtr)
	}
	return ret
}

var (
	engV8HandlerOnce   base.Once
	engV8HandlerImport *imports.Imports = nil
)

func engV8HandlerAPI() *imports.Imports {
	engV8HandlerOnce.Do(func() {
		engV8HandlerImport = api.NewDefaultImports()
		engV8HandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngV8Handler_Create", 0), // constructor NewEngV8Handler
			/* 1 */ imports.NewTable("TEngV8Handler_OnV8Execute", 0), // event OnV8Execute
		}
	})
	return engV8HandlerImport
}
