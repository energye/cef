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

// ICefSetZoomLevelTask Parent: ICefTask
type ICefSetZoomLevelTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefSetZoomLevelTask Parent: TCefTask
type TCefSetZoomLevelTask struct {
	TCefTask
}

func NewCefSetZoomLevelTask(aEvents IChromiumEvents, aValue float64) ICefSetZoomLevelTask {
	r1 := setZoomLevelTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(unsafePointer(&aValue)))
	return AsCefSetZoomLevelTask(r1)
}

func (m *TCefSetZoomLevelTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	setZoomLevelTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	setZoomLevelTaskImport       *imports.Imports = nil
	setZoomLevelTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSetZoomLevelTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefSetZoomLevelTask_Create", 0),
	}
)

func setZoomLevelTaskImportAPI() *imports.Imports {
	if setZoomLevelTaskImport == nil {
		setZoomLevelTaskImport = NewDefaultImports()
		setZoomLevelTaskImport.SetImportTable(setZoomLevelTaskImportTables)
		setZoomLevelTaskImportTables = nil
	}
	return setZoomLevelTaskImport
}
