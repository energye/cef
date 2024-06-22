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
}

// TCefEnableFocusTask Parent: TCefTask
type TCefEnableFocusTask struct {
	TCefTask
}

func NewCefEnableFocusTask(aEvents IChromiumEvents) ICefEnableFocusTask {
	r1 := enableFocusTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefEnableFocusTask(r1)
}

var (
	enableFocusTaskImport       *imports.Imports = nil
	enableFocusTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefEnableFocusTask_Create", 0),
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
