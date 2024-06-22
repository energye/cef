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

// ICEFButtonComponent Parent: ICEFViewComponent
type ICEFButtonComponent interface {
	ICEFViewComponent
	// AsLabelButton
	//  Returns this Button as a LabelButton or NULL if this is not a LabelButton.
	AsLabelButton() ICefLabelButton // property
	// State
	//  Returns the current display state of the Button.
	State() TCefButtonState // property
	// SetState Set State
	SetState(AValue TCefButtonState) // property
	// SetInkDropEnabled
	//  Sets the Button will use an ink drop effect for displaying state changes.
	SetInkDropEnabled(enabled bool) // procedure
	// SetTooltipText
	//  Sets the tooltip text that will be displayed when the user hovers the
	//  mouse cursor over the Button.
	SetTooltipText(tooltiptext string) // procedure
	// SetAccessibleName
	//  Sets the accessible name that will be exposed to assistive technology
	SetAccessibleName(name string) // procedure
	// SetOnButtonPressed
	//  Called when |button| is pressed.
	SetOnButtonPressed(fn TOnButtonPressed) // property event
	// SetOnButtonStateChanged
	//  Called when the state of |button| changes.
	SetOnButtonStateChanged(fn TOnButtonStateChanged) // property event
}

// TCEFButtonComponent Parent: TCEFViewComponent
type TCEFButtonComponent struct {
	TCEFViewComponent
	buttonPressedPtr      uintptr
	buttonStateChangedPtr uintptr
}

func NewCEFButtonComponent(aOwner IComponent) ICEFButtonComponent {
	r1 := buttonComponentImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsCEFButtonComponent(r1)
}

func (m *TCEFButtonComponent) AsLabelButton() ICefLabelButton {
	var resultCefLabelButton uintptr
	buttonComponentImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

func (m *TCEFButtonComponent) State() TCefButtonState {
	r1 := buttonComponentImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TCefButtonState(r1)
}

func (m *TCEFButtonComponent) SetState(AValue TCefButtonState) {
	buttonComponentImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFButtonComponent) SetInkDropEnabled(enabled bool) {
	buttonComponentImportAPI().SysCallN(3, m.Instance(), PascalBool(enabled))
}

func (m *TCEFButtonComponent) SetTooltipText(tooltiptext string) {
	buttonComponentImportAPI().SysCallN(6, m.Instance(), PascalStr(tooltiptext))
}

func (m *TCEFButtonComponent) SetAccessibleName(name string) {
	buttonComponentImportAPI().SysCallN(2, m.Instance(), PascalStr(name))
}

func (m *TCEFButtonComponent) SetOnButtonPressed(fn TOnButtonPressed) {
	if m.buttonPressedPtr != 0 {
		RemoveEventElement(m.buttonPressedPtr)
	}
	m.buttonPressedPtr = MakeEventDataPtr(fn)
	buttonComponentImportAPI().SysCallN(4, m.Instance(), m.buttonPressedPtr)
}

func (m *TCEFButtonComponent) SetOnButtonStateChanged(fn TOnButtonStateChanged) {
	if m.buttonStateChangedPtr != 0 {
		RemoveEventElement(m.buttonStateChangedPtr)
	}
	m.buttonStateChangedPtr = MakeEventDataPtr(fn)
	buttonComponentImportAPI().SysCallN(5, m.Instance(), m.buttonStateChangedPtr)
}

var (
	buttonComponentImport       *imports.Imports = nil
	buttonComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFButtonComponent_AsLabelButton", 0),
		/*1*/ imports.NewTable("CEFButtonComponent_Create", 0),
		/*2*/ imports.NewTable("CEFButtonComponent_SetAccessibleName", 0),
		/*3*/ imports.NewTable("CEFButtonComponent_SetInkDropEnabled", 0),
		/*4*/ imports.NewTable("CEFButtonComponent_SetOnButtonPressed", 0),
		/*5*/ imports.NewTable("CEFButtonComponent_SetOnButtonStateChanged", 0),
		/*6*/ imports.NewTable("CEFButtonComponent_SetTooltipText", 0),
		/*7*/ imports.NewTable("CEFButtonComponent_State", 0),
	}
)

func buttonComponentImportAPI() *imports.Imports {
	if buttonComponentImport == nil {
		buttonComponentImport = NewDefaultImports()
		buttonComponentImport.SetImportTable(buttonComponentImportTables)
		buttonComponentImportTables = nil
	}
	return buttonComponentImport
}
