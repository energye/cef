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

// ICefUpdatePrefsTask Parent: ICefTask
type ICefUpdatePrefsTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefUpdatePrefsTask Parent: TCefTask
type TCefUpdatePrefsTask struct {
	TCefTask
}

func NewCefUpdatePrefsTask(aEvents IChromiumEvents) ICefUpdatePrefsTask {
	r1 := updatePrefsTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefUpdatePrefsTask(r1)
}

func (m *TCefUpdatePrefsTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	updatePrefsTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	updatePrefsTaskImport       *imports.Imports = nil
	updatePrefsTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdatePrefsTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefUpdatePrefsTask_Create", 0),
	}
)

func updatePrefsTaskImportAPI() *imports.Imports {
	if updatePrefsTaskImport == nil {
		updatePrefsTaskImport = NewDefaultImports()
		updatePrefsTaskImport.SetImportTable(updatePrefsTaskImportTables)
		updatePrefsTaskImportTables = nil
	}
	return updatePrefsTaskImport
}
