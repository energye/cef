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
}

// TCefUpdatePrefsTask Parent: TCefTask
type TCefUpdatePrefsTask struct {
	TCefTask
}

func NewCefUpdatePrefsTask(aEvents IChromiumEvents) ICefUpdatePrefsTask {
	r1 := updatePrefsTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefUpdatePrefsTask(r1)
}

var (
	updatePrefsTaskImport       *imports.Imports = nil
	updatePrefsTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdatePrefsTask_Create", 0),
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
