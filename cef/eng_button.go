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

// ICefButton Parent: ICefView
//
//	A View representing a button. Depending on the specific type, the button
//	could be implemented by a native control or custom rendered. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_capi.h">CEF source file: /include/capi/views/cef_button_capi.h (cef_button_t)</a>
type ICefButton interface {
	ICefView
	// AsLabelButton
	//  Returns this Button as a LabelButton or NULL if this is not a LabelButton.
	AsLabelButton() ICefLabelButton // function
	// GetState
	//  Returns the current display state of the Button.
	GetState() TCefButtonState // function
	// SetState
	//  Sets the current display state of the Button.
	SetState(state TCefButtonState) // procedure
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
}

// TCefButton Parent: TCefView
//
//	A View representing a button. Depending on the specific type, the button
//	could be implemented by a native control or custom rendered. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_capi.h">CEF source file: /include/capi/views/cef_button_capi.h (cef_button_t)</a>
type TCefButton struct {
	TCefView
}

// ButtonRef -> ICefButton
var ButtonRef button

// button TCefButton Ref
type button uintptr

// UnWrap
//
//	Returns a ICefButton instance using a PCefButton data pointer.
func (m *button) UnWrap(data uintptr) ICefButton {
	var resultCefButton uintptr
	buttonImportAPI().SysCallN(6, uintptr(data), uintptr(unsafePointer(&resultCefButton)))
	return AsCefButton(resultCefButton)
}

func (m *TCefButton) AsLabelButton() ICefLabelButton {
	var resultCefLabelButton uintptr
	buttonImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

func (m *TCefButton) GetState() TCefButtonState {
	r1 := buttonImportAPI().SysCallN(1, m.Instance())
	return TCefButtonState(r1)
}

func (m *TCefButton) SetState(state TCefButtonState) {
	buttonImportAPI().SysCallN(4, m.Instance(), uintptr(state))
}

func (m *TCefButton) SetInkDropEnabled(enabled bool) {
	buttonImportAPI().SysCallN(3, m.Instance(), PascalBool(enabled))
}

func (m *TCefButton) SetTooltipText(tooltiptext string) {
	buttonImportAPI().SysCallN(5, m.Instance(), PascalStr(tooltiptext))
}

func (m *TCefButton) SetAccessibleName(name string) {
	buttonImportAPI().SysCallN(2, m.Instance(), PascalStr(name))
}

var (
	buttonImport       *imports.Imports = nil
	buttonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefButton_AsLabelButton", 0),
		/*1*/ imports.NewTable("CefButton_GetState", 0),
		/*2*/ imports.NewTable("CefButton_SetAccessibleName", 0),
		/*3*/ imports.NewTable("CefButton_SetInkDropEnabled", 0),
		/*4*/ imports.NewTable("CefButton_SetState", 0),
		/*5*/ imports.NewTable("CefButton_SetTooltipText", 0),
		/*6*/ imports.NewTable("CefButton_UnWrap", 0),
	}
)

func buttonImportAPI() *imports.Imports {
	if buttonImport == nil {
		buttonImport = NewDefaultImports()
		buttonImport.SetImportTable(buttonImportTables)
		buttonImportTables = nil
	}
	return buttonImport
}
