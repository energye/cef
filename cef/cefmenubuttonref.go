//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefMenuButton Parent: ICefLabelButton
type ICefMenuButton interface {
	ICefLabelButton
	ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) // procedure
	TriggerMenu()                                                                                            // procedure
}

// ICefMenuButtonRef Parent: ICefMenuButton ICefLabelButtonRef
type ICefMenuButtonRef interface {
	ICefMenuButton
	ICefLabelButtonRef
	AsIntfMenuButton() uintptr
	AsIntfLabelButton() uintptr
	AsIntfButton() uintptr
	AsIntfView() uintptr
}

type TCefMenuButtonRef struct {
	TCefLabelButtonRef
}

func (m *TCefMenuButtonRef) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cefMenuButtonRefAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCefMenuButtonRef) TriggerMenu() {
	if !m.IsValid() {
		return
	}
	cefMenuButtonRefAPI().SysCallN(4, m.Instance())
}

func (m *TCefMenuButtonRef) AsIntfMenuButton() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefMenuButtonRef) AsIntfLabelButton() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefMenuButtonRef) AsIntfButton() uintptr {
	return m.GetIntfPointer(2)
}
func (m *TCefMenuButtonRef) AsIntfView() uintptr {
	return m.GetIntfPointer(3)
}

// MenuButtonRef  is static instance
var MenuButtonRef _MenuButtonRefClass

// _MenuButtonRefClass is class type defined by TCefMenuButtonRef
type _MenuButtonRefClass uintptr

func (_MenuButtonRefClass) UnWrapWithPointer(data uintptr) (result ICefMenuButton) {
	var resultPtr uintptr
	cefMenuButtonRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuButtonRef(resultPtr)
	return
}

func (_MenuButtonRefClass) CreateMenuButton(delegate IEngMenuButtonDelegate, text string) (result ICefMenuButton) {
	var resultPtr uintptr
	cefMenuButtonRefAPI().SysCallN(2, base.GetObjectUintptr(delegate), api.PasStr(text), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuButtonRef(resultPtr)
	return
}

// NewMenuButtonRef class constructor
func NewMenuButtonRef(data uintptr) ICefMenuButtonRef {
	var menuButtonPtr uintptr  // ICefMenuButton
	var labelButtonPtr uintptr // ICefLabelButton
	var buttonPtr uintptr      // ICefButton
	var viewPtr uintptr        // ICefView
	r := cefMenuButtonRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&menuButtonPtr)), uintptr(base.UnsafePointer(&labelButtonPtr)), uintptr(base.UnsafePointer(&buttonPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefMenuButtonRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(4)
		intf.SetIntfPointer(0, menuButtonPtr)
		intf.SetIntfPointer(1, labelButtonPtr)
		intf.SetIntfPointer(2, buttonPtr)
		intf.SetIntfPointer(3, viewPtr)
	}
	return ret
}

var (
	cefMenuButtonRefOnce   base.Once
	cefMenuButtonRefImport *imports.Imports = nil
)

func cefMenuButtonRefAPI() *imports.Imports {
	cefMenuButtonRefOnce.Do(func() {
		cefMenuButtonRefImport = api.NewDefaultImports()
		cefMenuButtonRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMenuButtonRef_Create", 0), // constructor NewMenuButtonRef
			/* 1 */ imports.NewTable("TCefMenuButtonRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 2 */ imports.NewTable("TCefMenuButtonRef_CreateMenuButton", 0), // static function CreateMenuButton
			/* 3 */ imports.NewTable("TCefMenuButtonRef_ShowMenu", 0), // procedure ShowMenu
			/* 4 */ imports.NewTable("TCefMenuButtonRef_TriggerMenu", 0), // procedure TriggerMenu
		}
	})
	return cefMenuButtonRefImport
}
