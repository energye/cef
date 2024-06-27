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

// ICustomAudioHandler Parent: ICefAudioHandler
type ICustomAudioHandler interface {
	ICefAudioHandler
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefAudioHandler // procedure
}

// TCustomAudioHandler Parent: TCefAudioHandler
type TCustomAudioHandler struct {
	TCefAudioHandler
}

func NewCustomAudioHandler(events IChromiumEvents) ICustomAudioHandler {
	r1 := customAudioHandlerImportAPI().SysCallN(1, GetObjectUintptr(events))
	return AsCustomAudioHandler(r1)
}

func (m *TCustomAudioHandler) AsInterface() ICefAudioHandler {
	var resultCefAudioHandler uintptr
	customAudioHandlerImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefAudioHandler)))
	return AsCefAudioHandler(resultCefAudioHandler)
}

var (
	customAudioHandlerImport       *imports.Imports = nil
	customAudioHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomAudioHandler_AsInterface", 0),
		/*1*/ imports.NewTable("CustomAudioHandler_Create", 0),
	}
)

func customAudioHandlerImportAPI() *imports.Imports {
	if customAudioHandlerImport == nil {
		customAudioHandlerImport = NewDefaultImports()
		customAudioHandlerImport.SetImportTable(customAudioHandlerImportTables)
		customAudioHandlerImportTables = nil
	}
	return customAudioHandlerImport
}
