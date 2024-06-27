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

// ICefSetZoomPctTask Parent: ICefTask
type ICefSetZoomPctTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefSetZoomPctTask Parent: TCefTask
type TCefSetZoomPctTask struct {
	TCefTask
}

func NewCefSetZoomPctTask(aEvents IChromiumEvents, aValue float64) ICefSetZoomPctTask {
	r1 := setZoomPctTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(unsafePointer(&aValue)))
	return AsCefSetZoomPctTask(r1)
}

func (m *TCefSetZoomPctTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	setZoomPctTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	setZoomPctTaskImport       *imports.Imports = nil
	setZoomPctTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSetZoomPctTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefSetZoomPctTask_Create", 0),
	}
)

func setZoomPctTaskImportAPI() *imports.Imports {
	if setZoomPctTaskImport == nil {
		setZoomPctTaskImport = NewDefaultImports()
		setZoomPctTaskImport.SetImportTable(setZoomPctTaskImportTables)
		setZoomPctTaskImportTables = nil
	}
	return setZoomPctTaskImport
}
