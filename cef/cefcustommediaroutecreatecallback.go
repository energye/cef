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

// ICefCustomMediaRouteCreateCallback Parent: ICefMediaRouteCreateCallbackOwn
type ICefCustomMediaRouteCreateCallback interface {
	ICefMediaRouteCreateCallbackOwn
	AsIntfMediaRouteCreateCallback() uintptr
}

type TCefCustomMediaRouteCreateCallback struct {
	TCefMediaRouteCreateCallbackOwn
}

func (m *TCefCustomMediaRouteCreateCallback) AsIntfMediaRouteCreateCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomMediaRouteCreateCallback class constructor
func NewCustomMediaRouteCreateCallback(events IChromiumCore) ICefCustomMediaRouteCreateCallback {
	var mediaRouteCreateCallbackPtr uintptr // ICefMediaRouteCreateCallback
	r := cefCustomMediaRouteCreateCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&mediaRouteCreateCallbackPtr)))
	ret := AsCefCustomMediaRouteCreateCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, mediaRouteCreateCallbackPtr)
	}
	return ret
}

var (
	cefCustomMediaRouteCreateCallbackOnce   base.Once
	cefCustomMediaRouteCreateCallbackImport *imports.Imports = nil
)

func cefCustomMediaRouteCreateCallbackAPI() *imports.Imports {
	cefCustomMediaRouteCreateCallbackOnce.Do(func() {
		cefCustomMediaRouteCreateCallbackImport = api.NewDefaultImports()
		cefCustomMediaRouteCreateCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomMediaRouteCreateCallback_Create", 0), // constructor NewCustomMediaRouteCreateCallback
		}
	})
	return cefCustomMediaRouteCreateCallbackImport
}
