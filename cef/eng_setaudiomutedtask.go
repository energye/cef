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

// ICefSetAudioMutedTask Parent: ICefTask
type ICefSetAudioMutedTask interface {
	ICefTask
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefSetAudioMutedTask Parent: TCefTask
type TCefSetAudioMutedTask struct {
	TCefTask
}

func NewCefSetAudioMutedTask(aEvents IChromiumEvents, aValue bool) ICefSetAudioMutedTask {
	r1 := setAudioMutedTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents), PascalBool(aValue))
	return AsCefSetAudioMutedTask(r1)
}

func (m *TCefSetAudioMutedTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	setAudioMutedTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	setAudioMutedTaskImport       *imports.Imports = nil
	setAudioMutedTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefSetAudioMutedTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefSetAudioMutedTask_Create", 0),
	}
)

func setAudioMutedTaskImportAPI() *imports.Imports {
	if setAudioMutedTaskImport == nil {
		setAudioMutedTaskImport = NewDefaultImports()
		setAudioMutedTaskImport.SetImportTable(setAudioMutedTaskImportTables)
		setAudioMutedTaskImportTables = nil
	}
	return setAudioMutedTaskImport
}
