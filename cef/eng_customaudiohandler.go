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
}

// TCustomAudioHandler Parent: TCefAudioHandler
type TCustomAudioHandler struct {
	TCefAudioHandler
}

func NewCustomAudioHandler(events IChromiumEvents) ICustomAudioHandler {
	r1 := customAudioHandlerImportAPI().SysCallN(0, GetObjectUintptr(events))
	return AsCustomAudioHandler(r1)
}

var (
	customAudioHandlerImport       *imports.Imports = nil
	customAudioHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomAudioHandler_Create", 0),
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
