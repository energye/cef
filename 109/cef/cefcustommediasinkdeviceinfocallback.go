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

// ICefCustomMediaSinkDeviceInfoCallback Parent: ICefMediaSinkDeviceInfoCallbackOwn
type ICefCustomMediaSinkDeviceInfoCallback interface {
	ICefMediaSinkDeviceInfoCallbackOwn
	AsIntfMediaSinkDeviceInfoCallback() uintptr
}

type TCefCustomMediaSinkDeviceInfoCallback struct {
	TCefMediaSinkDeviceInfoCallbackOwn
}

func (m *TCefCustomMediaSinkDeviceInfoCallback) AsIntfMediaSinkDeviceInfoCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomMediaSinkDeviceInfoCallback class constructor
func NewCustomMediaSinkDeviceInfoCallback(events IChromiumCore) ICefCustomMediaSinkDeviceInfoCallback {
	var mediaSinkDeviceInfoCallbackPtr uintptr // ICefMediaSinkDeviceInfoCallback
	r := cefCustomMediaSinkDeviceInfoCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&mediaSinkDeviceInfoCallbackPtr)))
	ret := AsCefCustomMediaSinkDeviceInfoCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaSinkDeviceInfoCallbackPtr)
	}
	return ret
}

var (
	cefCustomMediaSinkDeviceInfoCallbackOnce   base.Once
	cefCustomMediaSinkDeviceInfoCallbackImport *imports.Imports = nil
)

func cefCustomMediaSinkDeviceInfoCallbackAPI() *imports.Imports {
	cefCustomMediaSinkDeviceInfoCallbackOnce.Do(func() {
		cefCustomMediaSinkDeviceInfoCallbackImport = api.NewDefaultImports()
		cefCustomMediaSinkDeviceInfoCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomMediaSinkDeviceInfoCallback_Create", 0), // constructor NewCustomMediaSinkDeviceInfoCallback
		}
	})
	return cefCustomMediaSinkDeviceInfoCallbackImport
}
