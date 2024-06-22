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

// ICefCreateCustomViewTask Parent: ICefTask
type ICefCreateCustomViewTask interface {
	ICefTask
}

// TCefCreateCustomViewTask Parent: TCefTask
type TCefCreateCustomViewTask struct {
	TCefTask
}

func NewCefCreateCustomViewTask(aEvents ICefViewDelegateEvents) ICefCreateCustomViewTask {
	r1 := createCustomViewTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefCreateCustomViewTask(r1)
}

var (
	createCustomViewTaskImport       *imports.Imports = nil
	createCustomViewTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCreateCustomViewTask_Create", 0),
	}
)

func createCustomViewTaskImportAPI() *imports.Imports {
	if createCustomViewTaskImport == nil {
		createCustomViewTaskImport = NewDefaultImports()
		createCustomViewTaskImport.SetImportTable(createCustomViewTaskImportTables)
		createCustomViewTaskImportTables = nil
	}
	return createCustomViewTaskImport
}
