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

// ICefTextfieldDelegateEvents Parent: ICefViewDelegateEvents
type ICefTextfieldDelegateEvents interface {
	ICefViewDelegateEvents
}

// ICEFTextfieldComponent Parent: ICefTextfieldDelegateEvents ICEFViewComponent
type ICEFTextfieldComponent interface {
	ICefTextfieldDelegateEvents
	ICEFViewComponent
	IsCommandEnabled(commandId cefTypes.TCefTextFieldCommands) bool          // function
	CreateTextField()                                                        // procedure
	AppendText(text string)                                                  // procedure
	InsertOrReplaceText(text string)                                         // procedure
	SelectAll(reversed bool)                                                 // procedure
	ClearSelection()                                                         // procedure
	SelectRange(range_ TCefRange)                                            // procedure
	SetFontList(fontList string)                                             // procedure
	ApplyTextColor(color cefTypes.TCefColor, range_ TCefRange)               // procedure
	ApplyTextStyle(style cefTypes.TCefTextStyle, add bool, range_ TCefRange) // procedure
	ExecuteCommand(commandId cefTypes.TCefTextFieldCommands)                 // procedure
	ClearEditHistory()                                                       // procedure
	SetAccessibleName(name string)                                           // procedure
	SetPlaceholderTextColor(color cefTypes.TCefColor)                        // procedure
	PasswordInput() bool                                                     // property PasswordInput Getter
	SetPasswordInput(value bool)                                             // property PasswordInput Setter
	ReadOnly() bool                                                          // property ReadOnly Getter
	SetReadOnly(value bool)                                                  // property ReadOnly Setter
	Text() string                                                            // property Text Getter
	SetText(value string)                                                    // property Text Setter
	SelectedText() string                                                    // property SelectedText Getter
	TextColor() cefTypes.TCefColor                                           // property TextColor Getter
	SetTextColor(value cefTypes.TCefColor)                                   // property TextColor Setter
	SelectionTextColor() cefTypes.TCefColor                                  // property SelectionTextColor Getter
	SetSelectionTextColor(value cefTypes.TCefColor)                          // property SelectionTextColor Setter
	SelectionBackgroundColor() cefTypes.TCefColor                            // property SelectionBackgroundColor Getter
	SetSelectionBackgroundColor(value cefTypes.TCefColor)                    // property SelectionBackgroundColor Setter
	PlaceholderText() string                                                 // property PlaceholderText Getter
	SetPlaceholderText(value string)                                         // property PlaceholderText Setter
	HasSelection() bool                                                      // property HasSelection Getter
	SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEventEvent)                     // property event
	SetOnAfterUserAction(fn TOnAfterUserActionEvent)                         // property event
	AsIntfTextfieldDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFTextfieldComponent struct {
	TCEFViewComponent
}

func (m *TCEFTextfieldComponent) IsCommandEnabled(commandId cefTypes.TCefTextFieldCommands) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(1, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) CreateTextField() {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(2, m.Instance())
}

func (m *TCEFTextfieldComponent) AppendText(text string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(3, m.Instance(), api.PasStr(text))
}

func (m *TCEFTextfieldComponent) InsertOrReplaceText(text string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(4, m.Instance(), api.PasStr(text))
}

func (m *TCEFTextfieldComponent) SelectAll(reversed bool) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(5, m.Instance(), api.PasBool(reversed))
}

func (m *TCEFTextfieldComponent) ClearSelection() {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(6, m.Instance())
}

func (m *TCEFTextfieldComponent) SelectRange(range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCEFTextfieldComponent) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(8, m.Instance(), api.PasStr(fontList))
}

func (m *TCEFTextfieldComponent) ApplyTextColor(color cefTypes.TCefColor, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(9, m.Instance(), uintptr(color), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCEFTextfieldComponent) ApplyTextStyle(style cefTypes.TCefTextStyle, add bool, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(10, m.Instance(), uintptr(style), api.PasBool(add), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCEFTextfieldComponent) ExecuteCommand(commandId cefTypes.TCefTextFieldCommands) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(11, m.Instance(), uintptr(commandId))
}

func (m *TCEFTextfieldComponent) ClearEditHistory() {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(12, m.Instance())
}

func (m *TCEFTextfieldComponent) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(13, m.Instance(), api.PasStr(name))
}

func (m *TCEFTextfieldComponent) SetPlaceholderTextColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(14, m.Instance(), uintptr(color))
}

func (m *TCEFTextfieldComponent) PasswordInput() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetPasswordInput(value bool) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFTextfieldComponent) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFTextfieldComponent) Text() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(17, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) SetText(value string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(17, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFTextfieldComponent) SelectedText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) TextColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFTextfieldComponentAPI().SysCallN(19, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCEFTextfieldComponent) SetTextColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCEFTextfieldComponent) SelectionTextColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFTextfieldComponentAPI().SysCallN(20, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCEFTextfieldComponent) SetSelectionTextColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCEFTextfieldComponent) SelectionBackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFTextfieldComponentAPI().SysCallN(21, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCEFTextfieldComponent) SetSelectionBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCEFTextfieldComponent) PlaceholderText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(22, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) SetPlaceholderText(value string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(22, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFTextfieldComponent) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextfieldKeyEventEvent(fn)
	base.SetEvent(m, 24, cEFTextfieldComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFTextfieldComponent) SetOnAfterUserAction(fn TOnAfterUserActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAfterUserActionEvent(fn)
	base.SetEvent(m, 25, cEFTextfieldComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFTextfieldComponent) AsIntfTextfieldDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFTextfieldComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}

// NewTextfieldComponent class constructor
func NewTextfieldComponent(owner lcl.IComponent) ICEFTextfieldComponent {
	var textfieldDelegateEventsPtr uintptr // ICefTextfieldDelegateEvents
	var viewDelegateEventsPtr uintptr      // ICefViewDelegateEvents
	r := cEFTextfieldComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&textfieldDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFTextfieldComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldDelegateEventsPtr)
		intf.SetIntfPointer(1, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFTextfieldComponentOnce   base.Once
	cEFTextfieldComponentImport *imports.Imports = nil
)

func cEFTextfieldComponentAPI() *imports.Imports {
	cEFTextfieldComponentOnce.Do(func() {
		cEFTextfieldComponentImport = api.NewDefaultImports()
		cEFTextfieldComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFTextfieldComponent_Create", 0), // constructor NewTextfieldComponent
			/* 1 */ imports.NewTable("TCEFTextfieldComponent_IsCommandEnabled", 0), // function IsCommandEnabled
			/* 2 */ imports.NewTable("TCEFTextfieldComponent_CreateTextField", 0), // procedure CreateTextField
			/* 3 */ imports.NewTable("TCEFTextfieldComponent_AppendText", 0), // procedure AppendText
			/* 4 */ imports.NewTable("TCEFTextfieldComponent_InsertOrReplaceText", 0), // procedure InsertOrReplaceText
			/* 5 */ imports.NewTable("TCEFTextfieldComponent_SelectAll", 0), // procedure SelectAll
			/* 6 */ imports.NewTable("TCEFTextfieldComponent_ClearSelection", 0), // procedure ClearSelection
			/* 7 */ imports.NewTable("TCEFTextfieldComponent_SelectRange", 0), // procedure SelectRange
			/* 8 */ imports.NewTable("TCEFTextfieldComponent_SetFontList", 0), // procedure SetFontList
			/* 9 */ imports.NewTable("TCEFTextfieldComponent_ApplyTextColor", 0), // procedure ApplyTextColor
			/* 10 */ imports.NewTable("TCEFTextfieldComponent_ApplyTextStyle", 0), // procedure ApplyTextStyle
			/* 11 */ imports.NewTable("TCEFTextfieldComponent_ExecuteCommand", 0), // procedure ExecuteCommand
			/* 12 */ imports.NewTable("TCEFTextfieldComponent_ClearEditHistory", 0), // procedure ClearEditHistory
			/* 13 */ imports.NewTable("TCEFTextfieldComponent_SetAccessibleName", 0), // procedure SetAccessibleName
			/* 14 */ imports.NewTable("TCEFTextfieldComponent_SetPlaceholderTextColor", 0), // procedure SetPlaceholderTextColor
			/* 15 */ imports.NewTable("TCEFTextfieldComponent_PasswordInput", 0), // property PasswordInput
			/* 16 */ imports.NewTable("TCEFTextfieldComponent_ReadOnly", 0), // property ReadOnly
			/* 17 */ imports.NewTable("TCEFTextfieldComponent_Text", 0), // property Text
			/* 18 */ imports.NewTable("TCEFTextfieldComponent_SelectedText", 0), // property SelectedText
			/* 19 */ imports.NewTable("TCEFTextfieldComponent_TextColor", 0), // property TextColor
			/* 20 */ imports.NewTable("TCEFTextfieldComponent_SelectionTextColor", 0), // property SelectionTextColor
			/* 21 */ imports.NewTable("TCEFTextfieldComponent_SelectionBackgroundColor", 0), // property SelectionBackgroundColor
			/* 22 */ imports.NewTable("TCEFTextfieldComponent_PlaceholderText", 0), // property PlaceholderText
			/* 23 */ imports.NewTable("TCEFTextfieldComponent_HasSelection", 0), // property HasSelection
			/* 24 */ imports.NewTable("TCEFTextfieldComponent_OnTextfieldKeyEvent", 0), // event OnTextfieldKeyEvent
			/* 25 */ imports.NewTable("TCEFTextfieldComponent_OnAfterUserAction", 0), // event OnAfterUserAction
		}
	})
	return cEFTextfieldComponentImport
}
