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

// ICefSetZoomStepTask Parent: ICefTask
type ICefSetZoomStepTask interface {
	ICefTask
}

// TCefSetZoomStepTask Parent: TCefTask
type TCefSetZoomStepTask struct {
	TCefTask
}

func NewCefSetZoomStepTask(aEvents IChromiumEvents, aValue byte) ICefSetZoomStepTask {
	r1 := setZoomStepTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents), uintptr(aValue))
	return AsCefSetZoomStepTask(r1)
}

var (
	setZoomStepTaskImport       *imports.Imports = nil
	setZoomStepTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSetZoomStepTask_Create", 0),
	}
)

func setZoomStepTaskImportAPI() *imports.Imports {
	if setZoomStepTaskImport == nil {
		setZoomStepTaskImport = NewDefaultImports()
		setZoomStepTaskImport.SetImportTable(setZoomStepTaskImportTables)
		setZoomStepTaskImportTables = nil
	}
	return setZoomStepTaskImport
}
