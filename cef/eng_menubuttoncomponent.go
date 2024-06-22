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

// ICEFMenuButtonComponent Parent: ICEFLabelButtonComponent
type ICEFMenuButtonComponent interface {
	ICEFLabelButtonComponent
	// CreateMenuButton
	//  Create a new MenuButton.
	CreateMenuButton(aText string) // procedure
	// ShowMenu
	//  Show a menu with contents |menu_model|. |screen_point| specifies the menu
	//  position in screen coordinates. |anchor_position| specifies how the menu
	//  will be anchored relative to |screen_point|. This function should be
	//  called from ICefMenuButtonDelegate.OnMenuButtonPressed().
	ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) // procedure
	// TriggerMenu
	//  Show the menu for this button. Results in a call to
	//  ICefMenuButtonDelegate.OnMenuButtonPressed().
	TriggerMenu() // procedure
	// SetOnMenuButtonPressed
	//  Called when |button| is pressed. Call ICefMenuButton.ShowMenu() to
	//  show a popup menu at |screen_point|. When showing a custom popup such as a
	//  window keep a reference to |button_pressed_lock| until the popup is hidden
	//  to maintain the pressed button state.
	SetOnMenuButtonPressed(fn TOnMenuButtonPressed) // property event
}

// TCEFMenuButtonComponent Parent: TCEFLabelButtonComponent
type TCEFMenuButtonComponent struct {
	TCEFLabelButtonComponent
	menuButtonPressedPtr uintptr
}

func NewCEFMenuButtonComponent(aOwner IComponent) ICEFMenuButtonComponent {
	r1 := menuButtonComponentImportAPI().SysCallN(0, GetObjectUintptr(aOwner))
	return AsCEFMenuButtonComponent(r1)
}

func (m *TCEFMenuButtonComponent) CreateMenuButton(aText string) {
	menuButtonComponentImportAPI().SysCallN(1, m.Instance(), PascalStr(aText))
}

func (m *TCEFMenuButtonComponent) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	menuButtonComponentImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCEFMenuButtonComponent) TriggerMenu() {
	menuButtonComponentImportAPI().SysCallN(4, m.Instance())
}

func (m *TCEFMenuButtonComponent) SetOnMenuButtonPressed(fn TOnMenuButtonPressed) {
	if m.menuButtonPressedPtr != 0 {
		RemoveEventElement(m.menuButtonPressedPtr)
	}
	m.menuButtonPressedPtr = MakeEventDataPtr(fn)
	menuButtonComponentImportAPI().SysCallN(2, m.Instance(), m.menuButtonPressedPtr)
}

var (
	menuButtonComponentImport       *imports.Imports = nil
	menuButtonComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFMenuButtonComponent_Create", 0),
		/*1*/ imports.NewTable("CEFMenuButtonComponent_CreateMenuButton", 0),
		/*2*/ imports.NewTable("CEFMenuButtonComponent_SetOnMenuButtonPressed", 0),
		/*3*/ imports.NewTable("CEFMenuButtonComponent_ShowMenu", 0),
		/*4*/ imports.NewTable("CEFMenuButtonComponent_TriggerMenu", 0),
	}
)

func menuButtonComponentImportAPI() *imports.Imports {
	if menuButtonComponentImport == nil {
		menuButtonComponentImport = NewDefaultImports()
		menuButtonComponentImport.SetImportTable(menuButtonComponentImportTables)
		menuButtonComponentImportTables = nil
	}
	return menuButtonComponentImport
}
