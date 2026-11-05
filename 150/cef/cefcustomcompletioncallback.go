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

// ICefCustomCompletionCallback Parent: ICefCompletionCallbackOwn
type ICefCustomCompletionCallback interface {
	ICefCompletionCallbackOwn
	AsIntfCompletionCallback() uintptr
}

type TCefCustomCompletionCallback struct {
	TCefCompletionCallbackOwn
}

func (m *TCefCustomCompletionCallback) AsIntfCompletionCallback() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomCompletionCallback class constructor
func NewCustomCompletionCallback(events IChromiumCore) ICefCustomCompletionCallback {
	var completionCallbackPtr uintptr // ICefCompletionCallback
	r := cefCustomCompletionCallbackAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&completionCallbackPtr)))
	ret := AsCefCustomCompletionCallback(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, completionCallbackPtr)
	}
	return ret
}

var (
	cefCustomCompletionCallbackOnce   base.Once
	cefCustomCompletionCallbackImport *imports.Imports = nil
)

func cefCustomCompletionCallbackAPI() *imports.Imports {
	cefCustomCompletionCallbackOnce.Do(func() {
		cefCustomCompletionCallbackImport = api.NewDefaultImports()
		cefCustomCompletionCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefCustomCompletionCallback_Create", 0), // constructor NewCustomCompletionCallback
		}
	})
	return cefCustomCompletionCallbackImport
}
