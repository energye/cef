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

// IEngEndTracingCallback Parent: ICefEndTracingCallbackOwn
type IEngEndTracingCallback interface {
	ICefEndTracingCallbackOwn
	SetOnEndTracingCallbackEndTracingComplete(fn TOnEndTracingCallbackEndTracingCompleteEvent) // property event
	AsIntfEndTracingCallback() uintptr
}

type TEngEndTracingCallback struct {
	TCefEndTracingCallbackOwn
}

func (m *TEngEndTracingCallback) SetOnEndTracingCallbackEndTracingComplete(fn TOnEndTracingCallbackEndTracingCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnEndTracingCallbackEndTracingCompleteEvent(fn)
	base.SetEvent(m, 1, engEndTracingCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngEndTracingCallback) AsIntfEndTracingCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngEndTracingCallback class constructor
func NewEngEndTracingCallback() IEngEndTracingCallback {
	var endTracingCallbackPtr uintptr // ICefEndTracingCallback
	r := engEndTracingCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&endTracingCallbackPtr)))
	ret := AsEngEndTracingCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, endTracingCallbackPtr)
	}
	return ret
}

var (
	engEndTracingCallbackOnce   base.Once
	engEndTracingCallbackImport *imports.Imports = nil
)

func engEndTracingCallbackAPI() *imports.Imports {
	engEndTracingCallbackOnce.Do(func() {
		engEndTracingCallbackImport = api.NewDefaultImports()
		engEndTracingCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngEndTracingCallback_Create", 0), // constructor NewEngEndTracingCallback
			/* 1 */ imports.NewTable("TEngEndTracingCallback_OnEndTracingCallbackEndTracingComplete", 0), // event OnEndTracingCallbackEndTracingComplete
		}
	})
	return engEndTracingCallbackImport
}
