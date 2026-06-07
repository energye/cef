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

// IEngDownloadImageCallback Parent: ICefDownloadImageCallbackOwn
type IEngDownloadImageCallback interface {
	ICefDownloadImageCallbackOwn
	SetOnDownloadImageCallbackDownloadImageFinished(fn TOnDownloadImageCallbackDownloadImageFinishedEvent) // property event
	AsIntfDownloadImageCallback() uintptr
}

type TEngDownloadImageCallback struct {
	TCefDownloadImageCallbackOwn
}

func (m *TEngDownloadImageCallback) SetOnDownloadImageCallbackDownloadImageFinished(fn TOnDownloadImageCallbackDownloadImageFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadImageCallbackDownloadImageFinishedEvent(fn)
	base.SetEvent(m, 1, engDownloadImageCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDownloadImageCallback) AsIntfDownloadImageCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDownloadImageCallback class constructor
func NewEngDownloadImageCallback() IEngDownloadImageCallback {
	var downloadImageCallbackPtr uintptr // ICefDownloadImageCallback
	r := engDownloadImageCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&downloadImageCallbackPtr)))
	ret := AsEngDownloadImageCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadImageCallbackPtr)
	}
	return ret
}

var (
	engDownloadImageCallbackOnce   base.Once
	engDownloadImageCallbackImport *imports.Imports = nil
)

func engDownloadImageCallbackAPI() *imports.Imports {
	engDownloadImageCallbackOnce.Do(func() {
		engDownloadImageCallbackImport = api.NewDefaultImports()
		engDownloadImageCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDownloadImageCallback_Create", 0), // constructor NewEngDownloadImageCallback
			/* 1 */ imports.NewTable("TEngDownloadImageCallback_OnDownloadImageCallbackDownloadImageFinished", 0), // event OnDownloadImageCallbackDownloadImageFinished
		}
	})
	return engDownloadImageCallbackImport
}
