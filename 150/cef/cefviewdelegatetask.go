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

// ICefViewDelegateTask Parent: ICefManagedTask
type ICefViewDelegateTask interface {
	ICefManagedTask
	AsIntfTask() uintptr
}

type TCefViewDelegateTask struct {
	TCefManagedTask
}

func (m *TCefViewDelegateTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewViewDelegateTask class constructor
func NewViewDelegateTask(events ICEFViewComponent) ICefViewDelegateTask {
	var taskPtr uintptr // ICefTask
	r := cefViewDelegateTaskAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefViewDelegateTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefViewDelegateTaskOnce   base.Once
	cefViewDelegateTaskImport *imports.Imports = nil
)

func cefViewDelegateTaskAPI() *imports.Imports {
	cefViewDelegateTaskOnce.Do(func() {
		cefViewDelegateTaskImport = api.NewDefaultImports()
		cefViewDelegateTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefViewDelegateTask_Create", 0), // constructor NewViewDelegateTask
		}
	})
	return cefViewDelegateTaskImport
}
