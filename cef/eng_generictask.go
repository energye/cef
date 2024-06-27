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

// ICefGenericTask Parent: ICefTask
type ICefGenericTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefGenericTask Parent: TCefTask
type TCefGenericTask struct {
	TCefTask
}

func NewCefGenericTask(aEvents IChromiumEvents, aTaskID uint32) ICefGenericTask {
	r1 := genericTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), uintptr(aTaskID))
	return AsCefGenericTask(r1)
}

func (m *TCefGenericTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	genericTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	genericTaskImport       *imports.Imports = nil
	genericTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefGenericTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefGenericTask_Create", 0),
	}
)

func genericTaskImportAPI() *imports.Imports {
	if genericTaskImport == nil {
		genericTaskImport = NewDefaultImports()
		genericTaskImport.SetImportTable(genericTaskImportTables)
		genericTaskImportTables = nil
	}
	return genericTaskImport
}
