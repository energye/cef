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

// IEngV8ArrayBufferReleaseCallback Parent: ICefv8ArrayBufferReleaseCallbackOwn
type IEngV8ArrayBufferReleaseCallback interface {
	ICefv8ArrayBufferReleaseCallbackOwn
	SetOnV8ArrayBufferReleaseCallbackReleaseBuffer(fn TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent) // property event
	AsIntfV8ArrayBufferReleaseCallback() uintptr
}

type TEngV8ArrayBufferReleaseCallback struct {
	TCefv8ArrayBufferReleaseCallbackOwn
}

func (m *TEngV8ArrayBufferReleaseCallback) SetOnV8ArrayBufferReleaseCallbackReleaseBuffer(fn TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnV8ArrayBufferReleaseCallbackReleaseBufferEvent(fn)
	base.SetEvent(m, 1, engV8ArrayBufferReleaseCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngV8ArrayBufferReleaseCallback) AsIntfV8ArrayBufferReleaseCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngV8ArrayBufferReleaseCallback class constructor
func NewEngV8ArrayBufferReleaseCallback() IEngV8ArrayBufferReleaseCallback {
	var v8ArrayBufferReleaseCallbackPtr uintptr // ICefv8ArrayBufferReleaseCallback
	r := engV8ArrayBufferReleaseCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&v8ArrayBufferReleaseCallbackPtr)))
	ret := AsEngV8ArrayBufferReleaseCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, v8ArrayBufferReleaseCallbackPtr)
	}
	return ret
}

var (
	engV8ArrayBufferReleaseCallbackOnce   base.Once
	engV8ArrayBufferReleaseCallbackImport *imports.Imports = nil
)

func engV8ArrayBufferReleaseCallbackAPI() *imports.Imports {
	engV8ArrayBufferReleaseCallbackOnce.Do(func() {
		engV8ArrayBufferReleaseCallbackImport = api.NewDefaultImports()
		engV8ArrayBufferReleaseCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngV8ArrayBufferReleaseCallback_Create", 0), // constructor NewEngV8ArrayBufferReleaseCallback
			/* 1 */ imports.NewTable("TEngV8ArrayBufferReleaseCallback_OnV8ArrayBufferReleaseCallbackReleaseBuffer", 0), // event OnV8ArrayBufferReleaseCallbackReleaseBuffer
		}
	})
	return engV8ArrayBufferReleaseCallbackImport
}
