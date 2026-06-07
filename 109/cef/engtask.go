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

// IEngTask Parent: ICefTaskOwn
type IEngTask interface {
	ICefTaskOwn
	SetOnTaskExecute(fn TOnTaskExecuteEvent) // property event
	AsIntfTask() uintptr
}

type TEngTask struct {
	TCefTaskOwn
}

func (m *TEngTask) SetOnTaskExecute(fn TOnTaskExecuteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTaskExecuteEvent(fn)
	base.SetEvent(m, 1, engTaskAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngTask class constructor
func NewEngTask() IEngTask {
	var taskPtr uintptr // ICefTask
	r := engTaskAPI().SysCallN(0, uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsEngTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	engTaskOnce   base.Once
	engTaskImport *imports.Imports = nil
)

func engTaskAPI() *imports.Imports {
	engTaskOnce.Do(func() {
		engTaskImport = api.NewDefaultImports()
		engTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngTask_Create", 0), // constructor NewEngTask
			/* 1 */ imports.NewTable("TEngTask_OnTaskExecute", 0), // event OnTaskExecute
		}
	})
	return engTaskImport
}
