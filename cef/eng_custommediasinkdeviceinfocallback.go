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

// ICefCustomMediaSinkDeviceInfoCallback Parent: ICefMediaSinkDeviceInfoCallback
type ICefCustomMediaSinkDeviceInfoCallback interface {
	ICefMediaSinkDeviceInfoCallback
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefMediaSinkDeviceInfoCallback // procedure
}

// TCefCustomMediaSinkDeviceInfoCallback Parent: TCefMediaSinkDeviceInfoCallback
type TCefCustomMediaSinkDeviceInfoCallback struct {
	TCefMediaSinkDeviceInfoCallback
}

func NewCefCustomMediaSinkDeviceInfoCallback(aEvents IChromiumEvents) ICefCustomMediaSinkDeviceInfoCallback {
	r1 := customMediaSinkDeviceInfoCallbackImportAPI().SysCallN(1, GetObjectUintptr(aEvents))
	return AsCefCustomMediaSinkDeviceInfoCallback(r1)
}

func (m *TCefCustomMediaSinkDeviceInfoCallback) AsInterface() ICefMediaSinkDeviceInfoCallback {
	var resultCefMediaSinkDeviceInfoCallback uintptr
	customMediaSinkDeviceInfoCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMediaSinkDeviceInfoCallback)))
	return AsCefMediaSinkDeviceInfoCallback(resultCefMediaSinkDeviceInfoCallback)
}

var (
	customMediaSinkDeviceInfoCallbackImport       *imports.Imports = nil
	customMediaSinkDeviceInfoCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefCustomMediaSinkDeviceInfoCallback_AsInterface", 0),
		/*1*/ imports.NewTable("CefCustomMediaSinkDeviceInfoCallback_Create", 0),
	}
)

func customMediaSinkDeviceInfoCallbackImportAPI() *imports.Imports {
	if customMediaSinkDeviceInfoCallbackImport == nil {
		customMediaSinkDeviceInfoCallbackImport = NewDefaultImports()
		customMediaSinkDeviceInfoCallbackImport.SetImportTable(customMediaSinkDeviceInfoCallbackImportTables)
		customMediaSinkDeviceInfoCallbackImportTables = nil
	}
	return customMediaSinkDeviceInfoCallbackImport
}
