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

// IEngMediaRouteCreateCallback Parent: ICefMediaRouteCreateCallbackOwn
type IEngMediaRouteCreateCallback interface {
	ICefMediaRouteCreateCallbackOwn
	SetOnMediaRouteCreateCallbackMediaRouteCreateFinished(fn TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent) // property event
	AsIntfMediaRouteCreateCallback() uintptr
}

type TEngMediaRouteCreateCallback struct {
	TCefMediaRouteCreateCallbackOwn
}

func (m *TEngMediaRouteCreateCallback) SetOnMediaRouteCreateCallbackMediaRouteCreateFinished(fn TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent(fn)
	base.SetEvent(m, 1, engMediaRouteCreateCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaRouteCreateCallback) AsIntfMediaRouteCreateCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngMediaRouteCreateCallback class constructor
func NewEngMediaRouteCreateCallback() IEngMediaRouteCreateCallback {
	var mediaRouteCreateCallbackPtr uintptr // ICefMediaRouteCreateCallback
	r := engMediaRouteCreateCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&mediaRouteCreateCallbackPtr)))
	ret := AsEngMediaRouteCreateCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaRouteCreateCallbackPtr)
	}
	return ret
}

var (
	engMediaRouteCreateCallbackOnce   base.Once
	engMediaRouteCreateCallbackImport *imports.Imports = nil
)

func engMediaRouteCreateCallbackAPI() *imports.Imports {
	engMediaRouteCreateCallbackOnce.Do(func() {
		engMediaRouteCreateCallbackImport = api.NewDefaultImports()
		engMediaRouteCreateCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMediaRouteCreateCallback_Create", 0), // constructor NewEngMediaRouteCreateCallback
			/* 1 */ imports.NewTable("TEngMediaRouteCreateCallback_OnMediaRouteCreateCallbackMediaRouteCreateFinished", 0), // event OnMediaRouteCreateCallbackMediaRouteCreateFinished
		}
	})
	return engMediaRouteCreateCallbackImport
}
