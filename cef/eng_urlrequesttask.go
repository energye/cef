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

// ICefURLRequestTask Parent: ICefTask
type ICefURLRequestTask interface {
	ICefTask
}

// TCefURLRequestTask Parent: TCefTask
type TCefURLRequestTask struct {
	TCefTask
}

func NewCefURLRequestTask(aEvents ICEFUrlRequestClientEvents) ICefURLRequestTask {
	r1 := uRLRequestTaskImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCefURLRequestTask(r1)
}

var (
	uRLRequestTaskImport       *imports.Imports = nil
	uRLRequestTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefURLRequestTask_Create", 0),
	}
)

func uRLRequestTaskImportAPI() *imports.Imports {
	if uRLRequestTaskImport == nil {
		uRLRequestTaskImport = NewDefaultImports()
		uRLRequestTaskImport.SetImportTable(uRLRequestTaskImportTables)
		uRLRequestTaskImportTables = nil
	}
	return uRLRequestTaskImport
}
