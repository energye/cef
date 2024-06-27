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

// ICefReadZoomTask Parent: ICefTask
type ICefReadZoomTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefReadZoomTask Parent: TCefTask
type TCefReadZoomTask struct {
	TCefTask
}

func NewCefReadZoomTask(aEvents IChromiumEvents) ICefReadZoomTask {
	r1 := readZoomTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefReadZoomTask(r1)
}

func (m *TCefReadZoomTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	readZoomTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	readZoomTaskImport       *imports.Imports = nil
	readZoomTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefReadZoomTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefReadZoomTask_Create", 0),
	}
)

func readZoomTaskImportAPI() *imports.Imports {
	if readZoomTaskImport == nil {
		readZoomTaskImport = NewDefaultImports()
		readZoomTaskImport.SetImportTable(readZoomTaskImportTables)
		readZoomTaskImportTables = nil
	}
	return readZoomTaskImport
}
