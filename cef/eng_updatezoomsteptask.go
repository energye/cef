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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefUpdateZoomStepTask Parent: TCefTask
type TCefUpdateZoomStepTask struct {
	TCefTask
}

func NewCefUpdateZoomStepTask(aEvents IChromiumEvents, aInc bool) ICefUpdateZoomStepTask {
	r1 := updateZoomStepTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), PascalBool(aInc))
	return AsCefUpdateZoomStepTask(r1)
}

func (m *TCefUpdateZoomStepTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	updateZoomStepTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	updateZoomStepTaskImport       *imports.Imports = nil
	updateZoomStepTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefUpdateZoomStepTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefUpdateZoomStepTask_Create", 0),
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
