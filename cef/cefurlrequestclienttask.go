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

// ICefURLRequestClientTask Parent: ICefManagedTask
type ICefURLRequestClientTask interface {
	ICefManagedTask
	AsIntfTask() uintptr
}

type TCefURLRequestClientTask struct {
	TCefManagedTask
}

func (m *TCefURLRequestClientTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewURLRequestClientTask class constructor
func NewURLRequestClientTask(events ICEFUrlRequestClientComponent) ICefURLRequestClientTask {
	var taskPtr uintptr // ICefTask
	r := cefURLRequestClientTaskAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefURLRequestClientTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefURLRequestClientTaskOnce   base.Once
	cefURLRequestClientTaskImport *imports.Imports = nil
)

func cefURLRequestClientTaskAPI() *imports.Imports {
	cefURLRequestClientTaskOnce.Do(func() {
		cefURLRequestClientTaskImport = api.NewDefaultImports()
		cefURLRequestClientTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefURLRequestClientTask_Create", 0), // constructor NewURLRequestClientTask
		}
	})
	return cefURLRequestClientTaskImport
}
