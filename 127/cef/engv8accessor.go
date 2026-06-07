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

// IEngV8Accessor Parent: ICefV8AccessorOwn
type IEngV8Accessor interface {
	ICefV8AccessorOwn
	SetOnV8AccessorGet(fn TOnV8AccessorGetEvent)   // property event
	SetOnV8AccessorSet_(fn TOnV8AccessorSet_Event) // property event
	AsIntfV8Accessor() uintptr
}

type TEngV8Accessor struct {
	TCefV8AccessorOwn
}

func (m *TEngV8Accessor) SetOnV8AccessorGet(fn TOnV8AccessorGetEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8AccessorGetEvent(fn)
	base.SetEvent(m, 1, engV8AccessorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Accessor) SetOnV8AccessorSet_(fn TOnV8AccessorSet_Event) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8AccessorSet_Event(fn)
	base.SetEvent(m, 2, engV8AccessorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Accessor) AsIntfV8Accessor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngV8Accessor class constructor
func NewEngV8Accessor() IEngV8Accessor {
	var v8AccessorPtr uintptr // ICefV8Accessor
	r := engV8AccessorAPI().SysCallN(0, uintptr(base.UnsafePointer(&v8AccessorPtr)))
	ret := AsEngV8Accessor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8AccessorPtr)
	}
	return ret
}

var (
	engV8AccessorOnce   base.Once
	engV8AccessorImport *imports.Imports = nil
)

func engV8AccessorAPI() *imports.Imports {
	engV8AccessorOnce.Do(func() {
		engV8AccessorImport = api.NewDefaultImports()
		engV8AccessorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngV8Accessor_Create", 0), // constructor NewEngV8Accessor
			/* 1 */ imports.NewTable("TEngV8Accessor_OnV8AccessorGet", 0), // event OnV8AccessorGet
			/* 2 */ imports.NewTable("TEngV8Accessor_OnV8AccessorSet_", 0), // event OnV8AccessorSet_
		}
	})
	return engV8AccessorImport
}
