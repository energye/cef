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

// ICefEventCompletionCallback Parent: ICefCompletionCallbackOwn
type ICefEventCompletionCallback interface {
	ICefCompletionCallbackOwn
	AsIntfCompletionCallback() uintptr
}

type TCefEventCompletionCallback struct {
	TCefCompletionCallbackOwn
}

func (m *TCefEventCompletionCallback) AsIntfCompletionCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewEventCompletionCallback class constructor
func NewEventCompletionCallback(event ICefWaitableEvent) ICefEventCompletionCallback {
	var completionCallbackPtr uintptr // ICefCompletionCallback
	r := cefEventCompletionCallbackAPI().SysCallN(0, base.GetObjectUintptr(event), uintptr(base.UnsafePointer(&completionCallbackPtr)))
	ret := AsCefEventCompletionCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, completionCallbackPtr)
	}
	return ret
}

var (
	cefEventCompletionCallbackOnce   base.Once
	cefEventCompletionCallbackImport *imports.Imports = nil
)

func cefEventCompletionCallbackAPI() *imports.Imports {
	cefEventCompletionCallbackOnce.Do(func() {
		cefEventCompletionCallbackImport = api.NewDefaultImports()
		cefEventCompletionCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefEventCompletionCallback_Create", 0), // constructor NewEventCompletionCallback
		}
	})
	return cefEventCompletionCallbackImport
}
