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

// ICefUpdateZoomPctTask Parent: ICefTask
type ICefUpdateZoomPctTask interface {
	ICefTask
}

// TCefUpdateZoomPctTask Parent: TCefTask
type TCefUpdateZoomPctTask struct {
	TCefTask
}

func NewCefUpdateZoomPctTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomPctTask {
	r1 := updateZoomPctTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomPctTask(r1)
}

var (
	updateZoomPctTaskImport       *imports.Imports = nil
	updateZoomPctTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdateZoomPctTask_Create", 0),
	}
)

func updateZoomPctTaskImportAPI() *imports.Imports {
	if updateZoomPctTaskImport == nil {
		updateZoomPctTaskImport = NewDefaultImports()
		updateZoomPctTaskImport.SetImportTable(updateZoomPctTaskImportTables)
		updateZoomPctTaskImportTables = nil
	}
	return updateZoomPctTaskImport
}
