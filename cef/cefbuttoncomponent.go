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

// ICefButtonDelegateEvents Parent: ICefViewDelegateEvents
type ICefButtonDelegateEvents interface {
	ICefViewDelegateEvents
}

// ICEFButtonComponent Parent: ICefButtonDelegateEvents ICEFViewComponent
type ICEFButtonComponent interface {
	ICefButtonDelegateEvents
	ICEFViewComponent
	// SetInkDropEnabled
	//  Sets the Button will use an ink drop effect for displaying state changes.
	SetInkDropEnabled(enabled bool) // procedure
	// SetTooltipText
	//  Sets the tooltip text that will be displayed when the user hovers the
	//  mouse cursor over the Button.
	SetTooltipText(tooltipText string) // procedure
	// SetAccessibleName
	//  Sets the accessible name that will be exposed to assistive technology
	//  (AT).
	SetAccessibleName(name string) // procedure
	// AsLabelButton
	//  Returns this Button as a LabelButton or NULL if this is not a LabelButton.
	AsLabelButton() ICefLabelButton // property AsLabelButton Getter
	// State
	//  Returns the current display state of the Button.
	State() cefTypes.TCefButtonState                       // property State Getter
	SetState(value cefTypes.TCefButtonState)               // property State Setter
	SetOnButtonPressed(fn TOnButtonPressedEvent)           // property event
	SetOnButtonStateChanged(fn TOnButtonStateChangedEvent) // property event
	AsIntfButtonDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFButtonComponent struct {
	TCEFViewComponent
}

func (m *TCEFButtonComponent) SetInkDropEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	cEFButtonComponentAPI().SysCallN(1, m.Instance(), api.PasBool(enabled))
}

func (m *TCEFButtonComponent) SetTooltipText(tooltipText string) {
	if !m.IsValid() {
		return
	}
	cEFButtonComponentAPI().SysCallN(2, m.Instance(), api.PasStr(tooltipText))
}

func (m *TCEFButtonComponent) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	cEFButtonComponentAPI().SysCallN(3, m.Instance(), api.PasStr(name))
}

func (m *TCEFButtonComponent) AsLabelButton() (result ICefLabelButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFButtonComponentAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLabelButtonRef(resultPtr)
	return
}

func (m *TCEFButtonComponent) State() cefTypes.TCefButtonState {
	if !m.IsValid() {
		return 0
	}
	r := cEFButtonComponentAPI().SysCallN(5, 0, m.Instance())
	return cefTypes.TCefButtonState(r)
}

func (m *TCEFButtonComponent) SetState(value cefTypes.TCefButtonState) {
	if !m.IsValid() {
		return
	}
	cEFButtonComponentAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCEFButtonComponent) SetOnButtonPressed(fn TOnButtonPressedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnButtonPressedEvent(fn)
	base.SetEvent(m, 6, cEFButtonComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFButtonComponent) SetOnButtonStateChanged(fn TOnButtonStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnButtonStateChangedEvent(fn)
	base.SetEvent(m, 7, cEFButtonComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFButtonComponent) AsIntfButtonDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFButtonComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}

// NewButtonComponent class constructor
func NewButtonComponent(owner lcl.IComponent) ICEFButtonComponent {
	var buttonDelegateEventsPtr uintptr // ICefButtonDelegateEvents
	var viewDelegateEventsPtr uintptr   // ICefViewDelegateEvents
	r := cEFButtonComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&buttonDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFButtonComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegateEventsPtr)
		intf.SetIntfPointer(1, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFButtonComponentOnce   base.Once
	cEFButtonComponentImport *imports.Imports = nil
)

func cEFButtonComponentAPI() *imports.Imports {
	cEFButtonComponentOnce.Do(func() {
		cEFButtonComponentImport = api.NewDefaultImports()
		cEFButtonComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFButtonComponent_Create", 0), // constructor NewButtonComponent
			/* 1 */ imports.NewTable("TCEFButtonComponent_SetInkDropEnabled", 0), // procedure SetInkDropEnabled
			/* 2 */ imports.NewTable("TCEFButtonComponent_SetTooltipText", 0), // procedure SetTooltipText
			/* 3 */ imports.NewTable("TCEFButtonComponent_SetAccessibleName", 0), // procedure SetAccessibleName
			/* 4 */ imports.NewTable("TCEFButtonComponent_AsLabelButton", 0), // property AsLabelButton
			/* 5 */ imports.NewTable("TCEFButtonComponent_State", 0), // property State
			/* 6 */ imports.NewTable("TCEFButtonComponent_OnButtonPressed", 0), // event OnButtonPressed
			/* 7 */ imports.NewTable("TCEFButtonComponent_OnButtonStateChanged", 0), // event OnButtonStateChanged
		}
	})
	return cEFButtonComponentImport
}
