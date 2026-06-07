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

// ICefCustomResolveCallback Parent: ICefResolveCallbackOwn
type ICefCustomResolveCallback interface {
	ICefResolveCallbackOwn
	AsIntfResolveCallback() uintptr
}

type TCefCustomResolveCallback struct {
	TCefResolveCallbackOwn
}

func (m *TCefCustomResolveCallback) AsIntfResolveCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomResolveCallback class constructor
func NewCustomResolveCallback(events IChromiumCore) ICefCustomResolveCallback {
	var resolveCallbackPtr uintptr // ICefResolveCallback
	r := cefCustomResolveCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&resolveCallbackPtr)))
	ret := AsCefCustomResolveCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resolveCallbackPtr)
	}
	return ret
}

var (
	cefCustomResolveCallbackOnce   base.Once
	cefCustomResolveCallbackImport *imports.Imports = nil
)

func cefCustomResolveCallbackAPI() *imports.Imports {
	cefCustomResolveCallbackOnce.Do(func() {
		cefCustomResolveCallbackImport = api.NewDefaultImports()
		cefCustomResolveCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomResolveCallback_Create", 0), // constructor NewCustomResolveCallback
		}
	})
	return cefCustomResolveCallbackImport
}
