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

// ICefCustomCompletionCallback Parent: ICefCompletionCallback
type ICefCustomCompletionCallback interface {
	ICefCompletionCallback
}

// TCefCustomCompletionCallback Parent: TCefCompletionCallback
type TCefCustomCompletionCallback struct {
	TCefCompletionCallback
}

func NewCefCustomCompletionCallback(aEvents IChromiumEvents) ICefCustomCompletionCallback {
	r1 := customCompletionCallbackImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefCustomCompletionCallback(r1)
}

var (
	customCompletionCallbackImport       *imports.Imports = nil
	customCompletionCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomCompletionCallback_Create", 0),
	}
)

func customCompletionCallbackImportAPI() *imports.Imports {
	if customCompletionCallbackImport == nil {
		customCompletionCallbackImport = NewDefaultImports()
		customCompletionCallbackImport.SetImportTable(customCompletionCallbackImportTables)
		customCompletionCallbackImportTables = nil
	}
	return customCompletionCallbackImport
}
