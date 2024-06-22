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
}

// TCefReadZoomTask Parent: TCefTask
type TCefReadZoomTask struct {
	TCefTask
}

func NewCefReadZoomTask(aEvents IChromiumEvents) ICefReadZoomTask {
	r1 := readZoomTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefReadZoomTask(r1)
}

var (
	readZoomTaskImport       *imports.Imports = nil
	readZoomTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefReadZoomTask_Create", 0),
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
