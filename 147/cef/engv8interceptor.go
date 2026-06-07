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

// IEngV8Interceptor Parent: ICefV8InterceptorOwn
type IEngV8Interceptor interface {
	ICefV8InterceptorOwn
	SetOnV8InterceptorGetByName(fn TOnV8InterceptorGetByNameEvent)   // property event
	SetOnV8InterceptorGetByIndex(fn TOnV8InterceptorGetByIndexEvent) // property event
	SetOnV8InterceptorSetByName(fn TOnV8InterceptorSetByNameEvent)   // property event
	SetOnV8InterceptorSetByIndex(fn TOnV8InterceptorSetByIndexEvent) // property event
	AsIntfV8Interceptor() uintptr
}

type TEngV8Interceptor struct {
	TCefV8InterceptorOwn
}

func (m *TEngV8Interceptor) SetOnV8InterceptorGetByName(fn TOnV8InterceptorGetByNameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8InterceptorGetByNameEvent(fn)
	base.SetEvent(m, 1, engV8InterceptorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Interceptor) SetOnV8InterceptorGetByIndex(fn TOnV8InterceptorGetByIndexEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8InterceptorGetByIndexEvent(fn)
	base.SetEvent(m, 2, engV8InterceptorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Interceptor) SetOnV8InterceptorSetByName(fn TOnV8InterceptorSetByNameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8InterceptorSetByNameEvent(fn)
	base.SetEvent(m, 3, engV8InterceptorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Interceptor) SetOnV8InterceptorSetByIndex(fn TOnV8InterceptorSetByIndexEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8InterceptorSetByIndexEvent(fn)
	base.SetEvent(m, 4, engV8InterceptorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8Interceptor) AsIntfV8Interceptor() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngV8Interceptor class constructor
func NewEngV8Interceptor() IEngV8Interceptor {
	var v8InterceptorPtr uintptr // ICefV8Interceptor
	r := engV8InterceptorAPI().SysCallN(0, uintptr(base.UnsafePointer(&v8InterceptorPtr)))
	ret := AsEngV8Interceptor(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8InterceptorPtr)
	}
	return ret
}

var (
	engV8InterceptorOnce   base.Once
	engV8InterceptorImport *imports.Imports = nil
)

func engV8InterceptorAPI() *imports.Imports {
	engV8InterceptorOnce.Do(func() {
		engV8InterceptorImport = api.NewDefaultImports()
		engV8InterceptorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngV8Interceptor_Create", 0), // constructor NewEngV8Interceptor
			/* 1 */ imports.NewTable("TEngV8Interceptor_OnV8InterceptorGetByName", 0), // event OnV8InterceptorGetByName
			/* 2 */ imports.NewTable("TEngV8Interceptor_OnV8InterceptorGetByIndex", 0), // event OnV8InterceptorGetByIndex
			/* 3 */ imports.NewTable("TEngV8Interceptor_OnV8InterceptorSetByName", 0), // event OnV8InterceptorSetByName
			/* 4 */ imports.NewTable("TEngV8Interceptor_OnV8InterceptorSetByIndex", 0), // event OnV8InterceptorSetByIndex
		}
	})
	return engV8InterceptorImport
}
