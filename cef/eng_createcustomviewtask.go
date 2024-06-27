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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefCreateCustomViewTask Parent: TCefTask
type TCefCreateCustomViewTask struct {
	TCefTask
}

func NewCefCreateCustomViewTask(aEvents ICefViewDelegateEvents) ICefCreateCustomViewTask {
	r1 := createCustomViewTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCreateCustomViewTask(r1)
}

func (m *TCefCreateCustomViewTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	createCustomViewTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	createCustomViewTaskImport       *imports.Imports = nil
	createCustomViewTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCreateCustomViewTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefCreateCustomViewTask_Create", 0),
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
