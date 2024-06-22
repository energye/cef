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

// ICefToggleAudioMutedTask Parent: ICefTask
type ICefToggleAudioMutedTask interface {
	ICefTask
}

// TCefToggleAudioMutedTask Parent: TCefTask
type TCefToggleAudioMutedTask struct {
	TCefTask
}

func NewCefToggleAudioMutedTask(aEvents IChromiumEvents) ICefToggleAudioMutedTask {
	r1 := oggleAudioMutedTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefToggleAudioMutedTask(r1)
}

var (
	oggleAudioMutedTaskImport       *imports.Imports = nil
	oggleAudioMutedTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefToggleAudioMutedTask_Create", 0),
	}
)

func oggleAudioMutedTaskImportAPI() *imports.Imports {
	if oggleAudioMutedTaskImport == nil {
		oggleAudioMutedTaskImport = NewDefaultImports()
		oggleAudioMutedTaskImport.SetImportTable(oggleAudioMutedTaskImportTables)
		oggleAudioMutedTaskImportTables = nil
	}
	return oggleAudioMutedTaskImport
}
