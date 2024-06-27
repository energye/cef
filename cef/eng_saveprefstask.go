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

// ICefSavePrefsTask Parent: ICefTask
type ICefSavePrefsTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefSavePrefsTask Parent: TCefTask
type TCefSavePrefsTask struct {
	TCefTask
}

func NewCefSavePrefsTask(aEvents IChromiumEvents) ICefSavePrefsTask {
	r1 := savePrefsTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefSavePrefsTask(r1)
}

func (m *TCefSavePrefsTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	savePrefsTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	savePrefsTaskImport       *imports.Imports = nil
	savePrefsTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSavePrefsTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefSavePrefsTask_Create", 0),
	}
)

func savePrefsTaskImportAPI() *imports.Imports {
	if savePrefsTaskImport == nil {
		savePrefsTaskImport = NewDefaultImports()
		savePrefsTaskImport.SetImportTable(savePrefsTaskImportTables)
		savePrefsTaskImportTables = nil
	}
	return savePrefsTaskImport
}
