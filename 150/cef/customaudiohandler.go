//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICustomAudioHandler Parent: ICefAudioHandlerOwn
type ICustomAudioHandler interface {
	ICefAudioHandlerOwn
	AsIntfAudioHandler() uintptr
}

type TCustomAudioHandler struct {
	TCefAudioHandlerOwn
}

func (m *TCustomAudioHandler) AsIntfAudioHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomAudioHandler class constructor
func NewCustomAudioHandler(events IChromiumCore) ICustomAudioHandler {
	var audioHandlerPtr uintptr // ICefAudioHandler
	r := customAudioHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&audioHandlerPtr)))
	ret := AsCustomAudioHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, audioHandlerPtr)
	}
	return ret
}

var (
	customAudioHandlerOnce   base.Once
	customAudioHandlerImport *imports.Imports = nil
)

func customAudioHandlerAPI() *imports.Imports {
	customAudioHandlerOnce.Do(func() {
		customAudioHandlerImport = api.NewDefaultImports()
		customAudioHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomAudioHandler_Create", 0), // constructor NewCustomAudioHandler
		}
	})
	return customAudioHandlerImport
}
