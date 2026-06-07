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

// IEngCompletionCallback Parent: ICefCompletionCallbackOwn
type IEngCompletionCallback interface {
	ICefCompletionCallbackOwn
	SetOnCompletionCallbackComplete(fn TOnCompletionCallbackCompleteEvent) // property event
	AsIntfCompletionCallback() uintptr
}

type TEngCompletionCallback struct {
	TCefCompletionCallbackOwn
}

func (m *TEngCompletionCallback) SetOnCompletionCallbackComplete(fn TOnCompletionCallbackCompleteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompletionCallbackCompleteEvent(fn)
	base.SetEvent(m, 1, engCompletionCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngCompletionCallback) AsIntfCompletionCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngCompletionCallback class constructor
func NewEngCompletionCallback() IEngCompletionCallback {
	var completionCallbackPtr uintptr // ICefCompletionCallback
	r := engCompletionCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&completionCallbackPtr)))
	ret := AsEngCompletionCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, completionCallbackPtr)
	}
	return ret
}

var (
	engCompletionCallbackOnce   base.Once
	engCompletionCallbackImport *imports.Imports = nil
)

func engCompletionCallbackAPI() *imports.Imports {
	engCompletionCallbackOnce.Do(func() {
		engCompletionCallbackImport = api.NewDefaultImports()
		engCompletionCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngCompletionCallback_Create", 0), // constructor NewEngCompletionCallback
			/* 1 */ imports.NewTable("TEngCompletionCallback_OnCompletionCallbackComplete", 0), // event OnCompletionCallbackComplete
		}
	})
	return engCompletionCallbackImport
}
