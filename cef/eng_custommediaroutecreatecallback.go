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

// ICefCustomMediaRouteCreateCallback Parent: ICefMediaRouteCreateCallback
type ICefCustomMediaRouteCreateCallback interface {
	ICefMediaRouteCreateCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefMediaRouteCreateCallback // procedure
}

// TCefCustomMediaRouteCreateCallback Parent: TCefMediaRouteCreateCallback
type TCefCustomMediaRouteCreateCallback struct {
	TCefMediaRouteCreateCallback
}

func NewCefCustomMediaRouteCreateCallback(aEvents IChromiumEvents) ICefCustomMediaRouteCreateCallback {
	r1 := customMediaRouteCreateCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomMediaRouteCreateCallback(r1)
}

func (m *TCefCustomMediaRouteCreateCallback) AsInterface() ICefMediaRouteCreateCallback {
	var resultCefMediaRouteCreateCallback uintptr
	customMediaRouteCreateCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMediaRouteCreateCallback)))
	return AsCefMediaRouteCreateCallback(resultCefMediaRouteCreateCallback)
}

var (
	customMediaRouteCreateCallbackImport       *imports.Imports = nil
	customMediaRouteCreateCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomMediaRouteCreateCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomMediaRouteCreateCallback_Create", 0),
	}
)

func customMediaRouteCreateCallbackImportAPI() *imports.Imports {
	if customMediaRouteCreateCallbackImport == nil {
		customMediaRouteCreateCallbackImport = NewDefaultImports()
		customMediaRouteCreateCallbackImport.SetImportTable(customMediaRouteCreateCallbackImportTables)
		customMediaRouteCreateCallbackImportTables = nil
	}
	return customMediaRouteCreateCallbackImport
}
