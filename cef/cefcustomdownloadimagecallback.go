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

// ICefCustomDownloadImageCallback Parent: ICefDownloadImageCallbackOwn
type ICefCustomDownloadImageCallback interface {
	ICefDownloadImageCallbackOwn
	AsIntfDownloadImageCallback() uintptr
}

type TCefCustomDownloadImageCallback struct {
	TCefDownloadImageCallbackOwn
}

func (m *TCefCustomDownloadImageCallback) AsIntfDownloadImageCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDownloadImageCallback class constructor
func NewCustomDownloadImageCallback(events IChromiumCore) ICefCustomDownloadImageCallback {
	var downloadImageCallbackPtr uintptr // ICefDownloadImageCallback
	r := cefCustomDownloadImageCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&downloadImageCallbackPtr)))
	ret := AsCefCustomDownloadImageCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadImageCallbackPtr)
	}
	return ret
}

var (
	cefCustomDownloadImageCallbackOnce   base.Once
	cefCustomDownloadImageCallbackImport *imports.Imports = nil
)

func cefCustomDownloadImageCallbackAPI() *imports.Imports {
	cefCustomDownloadImageCallbackOnce.Do(func() {
		cefCustomDownloadImageCallbackImport = api.NewDefaultImports()
		cefCustomDownloadImageCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomDownloadImageCallback_Create", 0), // constructor NewCustomDownloadImageCallback
		}
	})
	return cefCustomDownloadImageCallbackImport
}
