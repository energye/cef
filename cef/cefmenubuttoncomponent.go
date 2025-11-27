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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/types"
)

// ICEFMenuButtonComponent Parent: ICEFLabelButtonComponent
type ICEFMenuButtonComponent interface {
	ICEFLabelButtonComponent
	// CreateMenuButton
	//  Create a new MenuButton.
	CreateMenuButton(text string) // procedure
	// ShowMenu
	//  Show a menu with contents |menu_model|. |screen_point| specifies the menu
	//  position in screen coordinates. |anchor_position| specifies how the menu
	//  will be anchored relative to |screen_point|. This function should be
	//  called from ICefMenuButtonDelegate.OnMenuButtonPressed().
	ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) // procedure
	// TriggerMenu
	//  Show the menu for this button. Results in a call to
	//  ICefMenuButtonDelegate.OnMenuButtonPressed().
	TriggerMenu()                                        // procedure
	SetOnMenuButtonPressed(fn TOnMenuButtonPressedEvent) // property event
	AsIntfMenuButtonDelegateEvents() uintptr
	AsIntfButtonDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFMenuButtonComponent struct {
	TCEFLabelButtonComponent
}

func (m *TCEFMenuButtonComponent) CreateMenuButton(text string) {
	if !m.IsValid() {
		return
	}
	cEFMenuButtonComponentAPI().SysCallN(1, m.Instance(), api.PasStr(text))
}

func (m *TCEFMenuButtonComponent) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cEFMenuButtonComponentAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCEFMenuButtonComponent) TriggerMenu() {
	if !m.IsValid() {
		return
	}
	cEFMenuButtonComponentAPI().SysCallN(3, m.Instance())
}

func (m *TCEFMenuButtonComponent) SetOnMenuButtonPressed(fn TOnMenuButtonPressedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMenuButtonPressedEvent(fn)
	base.SetEvent(m, 4, cEFMenuButtonComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFMenuButtonComponent) AsIntfMenuButtonDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFMenuButtonComponent) AsIntfButtonDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCEFMenuButtonComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(2)
}

// NewMenuButtonComponent class constructor
func NewMenuButtonComponent(owner lcl.IComponent) ICEFMenuButtonComponent {
	var menuButtonDelegateEventsPtr uintptr // ICefMenuButtonDelegateEvents
	var buttonDelegateEventsPtr uintptr     // ICefButtonDelegateEvents
	var viewDelegateEventsPtr uintptr       // ICefViewDelegateEvents
	r := cEFMenuButtonComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&menuButtonDelegateEventsPtr)), uintptr(base.UnsafePointer(&buttonDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFMenuButtonComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, menuButtonDelegateEventsPtr)
		intf.SetIntfPointer(1, buttonDelegateEventsPtr)
		intf.SetIntfPointer(2, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFMenuButtonComponentOnce   base.Once
	cEFMenuButtonComponentImport *imports.Imports = nil
)

func cEFMenuButtonComponentAPI() *imports.Imports {
	cEFMenuButtonComponentOnce.Do(func() {
		cEFMenuButtonComponentImport = api.NewDefaultImports()
		cEFMenuButtonComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFMenuButtonComponent_Create", 0), // constructor NewMenuButtonComponent
			/* 1 */ imports.NewTable("TCEFMenuButtonComponent_CreateMenuButton", 0), // procedure CreateMenuButton
			/* 2 */ imports.NewTable("TCEFMenuButtonComponent_ShowMenu", 0), // procedure ShowMenu
			/* 3 */ imports.NewTable("TCEFMenuButtonComponent_TriggerMenu", 0), // procedure TriggerMenu
			/* 4 */ imports.NewTable("TCEFMenuButtonComponent_OnMenuButtonPressed", 0), // event OnMenuButtonPressed
		}
	})
	return cEFMenuButtonComponentImport
}
