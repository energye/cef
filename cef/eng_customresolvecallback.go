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

// ICefCustomResolveCallback Parent: ICefResolveCallback
type ICefCustomResolveCallback interface {
	ICefResolveCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefResolveCallback // procedure
}

// TCefCustomResolveCallback Parent: TCefResolveCallback
type TCefCustomResolveCallback struct {
	TCefResolveCallback
}

func NewCefCustomResolveCallback(aEvents IChromiumEvents) ICefCustomResolveCallback {
	r1 := customResolveCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomResolveCallback(r1)
}

func (m *TCefCustomResolveCallback) AsInterface() ICefResolveCallback {
	var resultCefResolveCallback uintptr
	customResolveCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefResolveCallback)))
	return AsCefResolveCallback(resultCefResolveCallback)
}

var (
	customResolveCallbackImport       *imports.Imports = nil
	customResolveCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomResolveCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomResolveCallback_Create", 0),
	}
)

func customResolveCallbackImportAPI() *imports.Imports {
	if customResolveCallbackImport == nil {
		customResolveCallbackImport = NewDefaultImports()
		customResolveCallbackImport.SetImportTable(customResolveCallbackImportTables)
		customResolveCallbackImportTables = nil
	}
	return customResolveCallbackImport
}
