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

// ICefUpdateZoomStepTask Parent: ICefTask
type ICefUpdateZoomStepTask interface {
	ICefTask
}

// TCefUpdateZoomStepTask Parent: TCefTask
type TCefUpdateZoomStepTask struct {
	TCefTask
}

func NewCefUpdateZoomStepTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomStepTask {
	r1 := updateZoomStepTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomStepTask(r1)
}

var (
	updateZoomStepTaskImport       *imports.Imports = nil
	updateZoomStepTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdateZoomStepTask_Create", 0),
	}
)

func updateZoomStepTaskImportAPI() *imports.Imports {
	if updateZoomStepTaskImport == nil {
		updateZoomStepTaskImport = NewDefaultImports()
		updateZoomStepTaskImport.SetImportTable(updateZoomStepTaskImportTables)
		updateZoomStepTaskImportTables = nil
	}
	return updateZoomStepTaskImport
}
