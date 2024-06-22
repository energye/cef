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
}

// TCefSavePrefsTask Parent: TCefTask
type TCefSavePrefsTask struct {
	TCefTask
}

func NewCefSavePrefsTask(aEvents IChromiumEvents) ICefSavePrefsTask {
	r1 := savePrefsTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefSavePrefsTask(r1)
}

var (
	savePrefsTaskImport       *imports.Imports = nil
	savePrefsTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSavePrefsTask_Create", 0),
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
