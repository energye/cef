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

// IEngMediaSinkDeviceInfoCallback Parent: ICefMediaSinkDeviceInfoCallbackOwn
type IEngMediaSinkDeviceInfoCallback interface {
	ICefMediaSinkDeviceInfoCallbackOwn
	SetOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent) // property event
	AsIntfMediaSinkDeviceInfoCallback() uintptr
}

type TEngMediaSinkDeviceInfoCallback struct {
	TCefMediaSinkDeviceInfoCallbackOwn
}

func (m *TEngMediaSinkDeviceInfoCallback) SetOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent(fn)
	base.SetEvent(m, 1, engMediaSinkDeviceInfoCallbackAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngMediaSinkDeviceInfoCallback) AsIntfMediaSinkDeviceInfoCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngMediaSinkDeviceInfoCallback class constructor
func NewEngMediaSinkDeviceInfoCallback() IEngMediaSinkDeviceInfoCallback {
	var mediaSinkDeviceInfoCallbackPtr uintptr // ICefMediaSinkDeviceInfoCallback
	r := engMediaSinkDeviceInfoCallbackAPI().SysCallN(0, uintptr(base.UnsafePointer(&mediaSinkDeviceInfoCallbackPtr)))
	ret := AsEngMediaSinkDeviceInfoCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaSinkDeviceInfoCallbackPtr)
	}
	return ret
}

var (
	engMediaSinkDeviceInfoCallbackOnce   base.Once
	engMediaSinkDeviceInfoCallbackImport *imports.Imports = nil
)

func engMediaSinkDeviceInfoCallbackAPI() *imports.Imports {
	engMediaSinkDeviceInfoCallbackOnce.Do(func() {
		engMediaSinkDeviceInfoCallbackImport = api.NewDefaultImports()
		engMediaSinkDeviceInfoCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngMediaSinkDeviceInfoCallback_Create", 0), // constructor NewEngMediaSinkDeviceInfoCallback
			/* 1 */ imports.NewTable("TEngMediaSinkDeviceInfoCallback_OnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfo", 0), // event OnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfo
		}
	})
	return engMediaSinkDeviceInfoCallbackImport
}
