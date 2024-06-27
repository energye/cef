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
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefTask // procedure
}

// TCefURLRequestTask Parent: TCefTask
type TCefURLRequestTask struct {
	TCefTask
}

func NewCefURLRequestTask(aEvents ICEFUrlRequestClientEvents) ICefURLRequestTask {
	r1 := uRLRequestTaskImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefURLRequestTask(r1)
}

func (m *TCefURLRequestTask) AsInterface() ICefTask {
	var resultCefTask uintptr
	uRLRequestTaskImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefTask)))
	return AsCefTask(resultCefTask)
}

var (
	uRLRequestTaskImport       *imports.Imports = nil
	uRLRequestTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefURLRequestTask_AsInterface", 0),
		/*1*/ imports.NewTable("CefURLRequestTask_Create", 0),
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
