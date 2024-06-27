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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefSetZoomStepTask Parent: TCefTask
type TCefSetZoomStepTask struct {
	TCefTask
}

func NewCefSetZoomStepTask(aEvents IChromiumEvents, aValue byte) ICefSetZoomStepTask {
	r1 := setZoomStepTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(aValue))
	return AsCefSetZoomStepTask(r1)
}

func (m *TCefSetZoomStepTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	setZoomStepTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	setZoomStepTaskImport       *imports.Imports = nil
	setZoomStepTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSetZoomStepTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefSetZoomStepTask_Create", 0),
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
