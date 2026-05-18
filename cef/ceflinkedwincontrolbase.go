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

// ICEFLinkedWinControlBase Parent: ICEFWinControl
type ICEFLinkedWinControlBase interface {
	ICEFWinControl
	// UseSetFocus
	//  Use TChromium.SetFocus when the component receives a WM_SETFOCUS message in Windows.
	UseSetFocus() bool         // property UseSetFocus Getter
	SetUseSetFocus(value bool) // property UseSetFocus Setter
}

type TCEFLinkedWinControlBase struct {
	TCEFWinControl
}

func (m *TCEFLinkedWinControlBase) UseSetFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFLinkedWinControlBaseAPI().SysCallN(0, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFLinkedWinControlBase) SetUseSetFocus(value bool) {
	if !m.IsValid() {
		return
	}
	cEFLinkedWinControlBaseAPI().SysCallN(0, 1, m.Instance(), api.PasBool(value))
}

var (
	cEFLinkedWinControlBaseOnce   base.Once
	cEFLinkedWinControlBaseImport *imports.Imports = nil
)

func cEFLinkedWinControlBaseAPI() *imports.Imports {
	cEFLinkedWinControlBaseOnce.Do(func() {
		cEFLinkedWinControlBaseImport = api.NewDefaultImports()
		cEFLinkedWinControlBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFLinkedWinControlBase_UseSetFocus", 0), // property UseSetFocus
		}
	})
	return cEFLinkedWinControlBaseImport
}
