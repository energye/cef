//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICefEnableFocusTask Parent: ICefTask
type ICefEnableFocusTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefEnableFocusTask Parent: TCefTask
type TCefEnableFocusTask struct {
	TCefTask
}

func NewCefEnableFocusTask(aEvents IChromiumEvents) ICefEnableFocusTask {
	r1 := enableFocusTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefEnableFocusTask(r1)
}

func (m *TCefEnableFocusTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	enableFocusTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	enableFocusTaskImport       *imports.Imports = nil
	enableFocusTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefEnableFocusTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefEnableFocusTask_Create", 0),
	}
)

func enableFocusTaskImportAPI() *imports.Imports {
	if enableFocusTaskImport == nil {
		enableFocusTaskImport = NewDefaultImports()
		enableFocusTaskImport.SetImportTable(enableFocusTaskImportTables)
		enableFocusTaskImportTables = nil
	}
	return enableFocusTaskImport
}
