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

// ICefCustomSetCookieCallback Parent: ICefSetCookieCallbackOwn
type ICefCustomSetCookieCallback interface {
	ICefSetCookieCallbackOwn
	AsIntfSetCookieCallback() uintptr
}

type TCefCustomSetCookieCallback struct {
	TCefSetCookieCallbackOwn
}

func (m *TCefCustomSetCookieCallback) AsIntfSetCookieCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomSetCookieCallback class constructor
func NewCustomSetCookieCallback(events IChromiumCore, iD int32) ICefCustomSetCookieCallback {
	var setCookieCallbackPtr uintptr // ICefSetCookieCallback
	r := cefCustomSetCookieCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(iD), uintptr(base.UnsafePointer(&setCookieCallbackPtr)))
	ret := AsCefCustomSetCookieCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, setCookieCallbackPtr)
	}
	return ret
}

var (
	cefCustomSetCookieCallbackOnce   base.Once
	cefCustomSetCookieCallbackImport *imports.Imports = nil
)

func cefCustomSetCookieCallbackAPI() *imports.Imports {
	cefCustomSetCookieCallbackOnce.Do(func() {
		cefCustomSetCookieCallbackImport = api.NewDefaultImports()
		cefCustomSetCookieCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomSetCookieCallback_Create", 0), // constructor NewCustomSetCookieCallback
		}
	})
	return cefCustomSetCookieCallbackImport
}
