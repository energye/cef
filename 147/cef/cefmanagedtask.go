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

// ICefManagedTask Parent: ICefTaskOwn
type ICefManagedTask interface {
	ICefTaskOwn
	AsIntfTask() uintptr
}

type TCefManagedTask struct {
	TCefTaskOwn
}

func (m *TCefManagedTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewManagedTask class constructor
func NewManagedTask() ICefManagedTask {
	var taskPtr uintptr // ICefTask
	r := cefManagedTaskAPI().SysCallN(0, uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefManagedTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefManagedTaskOnce   base.Once
	cefManagedTaskImport *imports.Imports = nil
)

func cefManagedTaskAPI() *imports.Imports {
	cefManagedTaskOnce.Do(func() {
		cefManagedTaskImport = api.NewDefaultImports()
		cefManagedTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefManagedTask_Create", 0), // constructor NewManagedTask
		}
	})
	return cefManagedTaskImport
}
