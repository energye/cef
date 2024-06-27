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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefUpdateZoomPctTask Parent: TCefTask
type TCefUpdateZoomPctTask struct {
	TCefTask
}

func NewCefUpdateZoomPctTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomPctTask {
	r1 := updateZoomPctTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomPctTask(r1)
}

func (m *TCefUpdateZoomPctTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	updateZoomPctTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	updateZoomPctTaskImport       *imports.Imports = nil
	updateZoomPctTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdateZoomPctTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefUpdateZoomPctTask_Create", 0),
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
