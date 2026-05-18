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

// IEngRunFileDialogCallback Parent: ICefRunFileDialogCallbackOwn
type IEngRunFileDialogCallback interface {
	ICefRunFileDialogCallbackOwn
	SetOnRunFileDialogCallbackFileDialogDismissed(fn TOnRunFileDialogCallbackFileDialogDismissedEvent) // property event
	AsIntfRunFileDialogCallback() uintptr
}

type TEngRunFileDialogCallback struct {
	TCefRunFileDialogCallbackOwn
}

func (m *TEngRunFileDialogCallback) SetOnRunFileDialogCallbackFileDialogDismissed(fn TOnRunFileDialogCallbackFileDialogDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRunFileDialogCallbackFileDialogDismissedEvent(fn)
	base.SetEvent(m, 1, engRunFileDialogCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRunFileDialogCallback) AsIntfRunFileDialogCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngRunFileDialogCallback class constructor
func NewEngRunFileDialogCallback() IEngRunFileDialogCallback {
	var runFileDialogCallbackPtr uintptr // ICefRunFileDialogCallback
	r := engRunFileDialogCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&runFileDialogCallbackPtr)))
	ret := AsEngRunFileDialogCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, runFileDialogCallbackPtr)
	}
	return ret
}

var (
	engRunFileDialogCallbackOnce   base.Once
	engRunFileDialogCallbackImport *imports.Imports = nil
)

func engRunFileDialogCallbackAPI() *imports.Imports {
	engRunFileDialogCallbackOnce.Do(func() {
		engRunFileDialogCallbackImport = api.NewDefaultImports()
		engRunFileDialogCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngRunFileDialogCallback_Create", 0), // constructor NewEngRunFileDialogCallback
			/* 1 */ imports.NewTable("TEngRunFileDialogCallback_OnRunFileDialogCallbackFileDialogDismissed", 0), // event OnRunFileDialogCallbackFileDialogDismissed
		}
	})
	return engRunFileDialogCallbackImport
}
