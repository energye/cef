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

// ICefChromiumTask Parent: ICefManagedTask
type ICefChromiumTask interface {
	ICefManagedTask
	AsIntfTask() uintptr
}

type TCefChromiumTask struct {
	TCefManagedTask
}

func (m *TCefChromiumTask) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// NewChromiumTask class constructor
func NewChromiumTask(events IChromiumCore) ICefChromiumTask {
	var taskPtr uintptr // ICefTask
	r := cefChromiumTaskAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefChromiumTask(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefChromiumTaskOnce   base.Once
	cefChromiumTaskImport *imports.Imports = nil
)

func cefChromiumTaskAPI() *imports.Imports {
	cefChromiumTaskOnce.Do(func() {
		cefChromiumTaskImport = api.NewDefaultImports()
		cefChromiumTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefChromiumTask_Create", 0), // constructor NewChromiumTask
		}
	})
	return cefChromiumTaskImport
}
