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

// ICefCustomDeleteCookiesCallback Parent: ICefDeleteCookiesCallbackOwn
type ICefCustomDeleteCookiesCallback interface {
	ICefDeleteCookiesCallbackOwn
	AsIntfDeleteCookiesCallback() uintptr
}

type TCefCustomDeleteCookiesCallback struct {
	TCefDeleteCookiesCallbackOwn
}

func (m *TCefCustomDeleteCookiesCallback) AsIntfDeleteCookiesCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomDeleteCookiesCallback class constructor
func NewCustomDeleteCookiesCallback(events IChromiumCore) ICefCustomDeleteCookiesCallback {
	var deleteCookiesCallbackPtr uintptr // ICefDeleteCookiesCallback
	r := cefCustomDeleteCookiesCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&deleteCookiesCallbackPtr)))
	ret := AsCefCustomDeleteCookiesCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, deleteCookiesCallbackPtr)
	}
	return ret
}

var (
	cefCustomDeleteCookiesCallbackOnce   base.Once
	cefCustomDeleteCookiesCallbackImport *imports.Imports = nil
)

func cefCustomDeleteCookiesCallbackAPI() *imports.Imports {
	cefCustomDeleteCookiesCallbackOnce.Do(func() {
		cefCustomDeleteCookiesCallbackImport = api.NewDefaultImports()
		cefCustomDeleteCookiesCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomDeleteCookiesCallback_Create", 0), // constructor NewCustomDeleteCookiesCallback
		}
	})
	return cefCustomDeleteCookiesCallbackImport
}
