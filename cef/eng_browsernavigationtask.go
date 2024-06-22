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

// ICefBrowserNavigationTask Parent: ICefTask
type ICefBrowserNavigationTask interface {
	ICefTask
}

// TCefBrowserNavigationTask Parent: TCefTask
type TCefBrowserNavigationTask struct {
	TCefTask
}

func NewCefBrowserNavigationTask(aEvents IChromiumEvents, aTask TCefBrowserNavigation) ICefBrowserNavigationTask {
	r1 := browserNavigationTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents), uintptr(aTask))
	return AsCefBrowserNavigationTask(r1)
}

var (
	browserNavigationTaskImport       *imports.Imports = nil
	browserNavigationTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefBrowserNavigationTask_Create", 0),
	}
)

func browserNavigationTaskImportAPI() *imports.Imports {
	if browserNavigationTaskImport == nil {
		browserNavigationTaskImport = NewDefaultImports()
		browserNavigationTaskImport.SetImportTable(browserNavigationTaskImportTables)
		browserNavigationTaskImportTables = nil
	}
	return browserNavigationTaskImport
}
