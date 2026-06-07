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

// ICefTask Parent: ICefBaseRefCounted
type ICefTask interface {
	ICefBaseRefCounted
}

// ICefTaskRef Parent: ICefTask ICefBaseRefCountedRef
type ICefTaskRef interface {
	ICefTask
	ICefBaseRefCountedRef
	AsIntfTask() uintptr
}

type TCefTaskRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefTaskRef) AsIntfTask() uintptr {
	return m.GetIntfPointer(0)
}

// TaskRef  is static instance
var TaskRef _TaskRefClass

// _TaskRefClass is class type defined by TCefTaskRef
type _TaskRefClass uintptr

func (_TaskRefClass) UnWrap(data uintptr) (result IEngTask) {
	var resultPtr uintptr
	cefTaskRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngTask(resultPtr)
	return
}

// NewTaskRef class constructor
func NewTaskRef(data uintptr) ICefTaskRef {
	var taskPtr uintptr // ICefTask
	r := cefTaskRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&taskPtr)))
	ret := AsCefTaskRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, taskPtr)
	}
	return ret
}

var (
	cefTaskRefOnce   base.Once
	cefTaskRefImport *imports.Imports = nil
)

func cefTaskRefAPI() *imports.Imports {
	cefTaskRefOnce.Do(func() {
		cefTaskRefImport = api.NewDefaultImports()
		cefTaskRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTaskRef_Create", 0), // constructor NewTaskRef
			/* 1 */ imports.NewTable("TCefTaskRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefTaskRefImport
}
