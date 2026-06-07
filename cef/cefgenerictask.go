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

// ICefGenericTask Parent: ICefTaskOwn
type ICefGenericTask interface {
	ICefTaskOwn
	AsIntfTask() uintptr
}

type TCefGenericTask struct {
	TCefTaskOwn
}

func (m *TCefGenericTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewGenericTask class constructor
func NewGenericTask(events IChromiumCore, taskID uint32) ICefGenericTask {
	var taskPtr uintptr // ICefTask
	r := cefGenericTaskAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(taskID), uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefGenericTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefGenericTaskOnce   base.Once
	cefGenericTaskImport *imports.Imports = nil
)

func cefGenericTaskAPI() *imports.Imports {
	cefGenericTaskOnce.Do(func() {
		cefGenericTaskImport = api.NewDefaultImports()
		cefGenericTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefGenericTask_Create", 0), // constructor NewGenericTask
		}
	})
	return cefGenericTaskImport
}
