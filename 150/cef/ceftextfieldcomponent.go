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

	cefTypes "github.com/energye/cef/150/types"
)

// ICefTextfieldDelegateEvents Parent: ICefViewDelegateEvents
type ICefTextfieldDelegateEvents interface {
	ICefViewDelegateEvents
}

// ICEFTextfieldComponent Parent: ICefTextfieldDelegateEvents ICEFViewComponent
type ICEFTextfieldComponent interface {
	ICefTextfieldDelegateEvents
	ICEFViewComponent
	// IsCommandEnabled
	//  Returns true (1) if the action associated with the specified command id is
	//  enabled. See additional comments on execute_command().
	IsCommandEnabled(commandId cefTypes.TCefTextFieldCommands) bool // function
	// CreateTextField
	//  Create a new Textfield.
	CreateTextField() // procedure
	// AppendText
	//  Appends |text| to the previously-existing text.
	AppendText(text string) // procedure
	// InsertOrReplaceText
	//  Inserts |text| at the current cursor position replacing any selected text.
	InsertOrReplaceText(text string) // procedure
	// SelectAll
	//  Selects all text. If |reversed| is true (1) the range will end at the
	//  logical beginning of the text; this generally shows the leading portion of
	//  text that overflows its display area.
	SelectAll(reversed bool) // procedure
	// ClearSelection
	//  Clears the text selection and sets the caret to the end.
	ClearSelection() // procedure
	// SetFontList
	//  Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
	//  where:
	//  <code>
	//  - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names (case-sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//  </code>
	//  Here are examples of valid font description strings:
	//  <code>
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
	//  </code>
	SetFontList(fontList string) // procedure
	// ApplyTextColor
	//  Applies |color| to the specified |range| without changing the default
	//  color. If |range| is NULL the color will be set on the complete text
	//  contents.
	ApplyTextColor(color cefTypes.TCefColor, range_ TCefRange) // procedure
	// ApplyTextStyle
	//  Applies |style| to the specified |range| without changing the default
	//  style. If |add| is true (1) the style will be added, otherwise the style
	//  will be removed. If |range| is NULL the style will be set on the complete
	//  text contents.
	ApplyTextStyle(style cefTypes.TCefTextStyle, add bool, range_ TCefRange) // procedure
	// ExecuteCommand
	//  Performs the action associated with the specified command id.
	ExecuteCommand(commandId cefTypes.TCefTextFieldCommands) // procedure
	// ClearEditHistory
	//  Clears Edit history.
	ClearEditHistory() // procedure
	// SetAccessibleName
	//  Set the accessible name that will be exposed to assistive technology (AT).
	SetAccessibleName(name string) // procedure
	// PasswordInput
	//  Returns true (1) if the text will be displayed as asterisks.
	PasswordInput() bool         // property PasswordInput Getter
	SetPasswordInput(value bool) // property PasswordInput Setter
	// ReadOnly
	//  Returns true (1) if the text is read-only.
	ReadOnly() bool         // property ReadOnly Getter
	SetReadOnly(value bool) // property ReadOnly Setter
	// Text
	//  Returns the currently displayed text.
	Text() string         // property Text Getter
	SetText(value string) // property Text Setter
	// SelectedText
	//  Returns the currently selected text.
	SelectedText() string // property SelectedText Getter
	// SelectedRange
	//  Returns the selected logical text range.
	SelectedRange() TCefRange         // property SelectedRange Getter
	SetSelectedRange(value TCefRange) // property SelectedRange Setter
	// CursorPosition
	//  Returns the current cursor position.
	CursorPosition() cefTypes.NativeUInt // property CursorPosition Getter
	// PlaceholderText
	//  Returns the placeholder text that will be displayed when the Textfield is
	//  NULL.
	PlaceholderText() string         // property PlaceholderText Getter
	SetPlaceholderText(value string) // property PlaceholderText Setter
	// HasSelection
	//  Returns true (1) if there is any selected text.
	HasSelection() bool                                  // property HasSelection Getter
	SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEventEvent) // property event
	SetOnAfterUserAction(fn TOnAfterUserActionEvent)     // property event
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

func (m *TCEFTextfieldComponent) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(7, m.Instance(), api.PasStr(fontList))
}

func (m *TCEFTextfieldComponent) ApplyTextColor(color cefTypes.TCefColor, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(8, m.Instance(), uintptr(color), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCEFTextfieldComponent) ApplyTextStyle(style cefTypes.TCefTextStyle, add bool, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(9, m.Instance(), uintptr(style), api.PasBool(add), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCEFTextfieldComponent) ExecuteCommand(commandId cefTypes.TCefTextFieldCommands) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(10, m.Instance(), uintptr(commandId))
}

func (m *TCEFTextfieldComponent) ClearEditHistory() {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFTextfieldComponent) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(12, m.Instance(), api.PasStr(name))
}

func (m *TCEFTextfieldComponent) PasswordInput() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetPasswordInput(value bool) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFTextfieldComponent) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFTextfieldComponent) Text() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(15, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) SetText(value string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFTextfieldComponent) SelectedText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) SelectedRange() (result TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(17, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFTextfieldComponent) SetSelectedRange(value TCefRange) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(17, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFTextfieldComponent) CursorPosition() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cEFTextfieldComponentAPI().SysCallN(18, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCEFTextfieldComponent) PlaceholderText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFTextfieldComponentAPI().SysCallN(19, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFTextfieldComponent) SetPlaceholderText(value string) {
	if !m.IsValid() {
		return
	}
	cEFTextfieldComponentAPI().SysCallN(19, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFTextfieldComponent) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFTextfieldComponentAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFTextfieldComponent) SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextfieldKeyEventEvent(fn)
	base.SetEvent(m, 21, cEFTextfieldComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFTextfieldComponent) SetOnAfterUserAction(fn TOnAfterUserActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAfterUserActionEvent(fn)
	base.SetEvent(m, 22, cEFTextfieldComponentAPI(), api.MakeEventDataPtr(cb))
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
			/* 7 */ imports.NewTable("TCEFTextfieldComponent_SetFontList", 0), // procedure SetFontList
			/* 8 */ imports.NewTable("TCEFTextfieldComponent_ApplyTextColor", 0), // procedure ApplyTextColor
			/* 9 */ imports.NewTable("TCEFTextfieldComponent_ApplyTextStyle", 0), // procedure ApplyTextStyle
			/* 10 */ imports.NewTable("TCEFTextfieldComponent_ExecuteCommand", 0), // procedure ExecuteCommand
			/* 11 */ imports.NewTable("TCEFTextfieldComponent_ClearEditHistory", 0), // procedure ClearEditHistory
			/* 12 */ imports.NewTable("TCEFTextfieldComponent_SetAccessibleName", 0), // procedure SetAccessibleName
			/* 13 */ imports.NewTable("TCEFTextfieldComponent_PasswordInput", 0), // property PasswordInput
			/* 14 */ imports.NewTable("TCEFTextfieldComponent_ReadOnly", 0), // property ReadOnly
			/* 15 */ imports.NewTable("TCEFTextfieldComponent_Text", 0), // property Text
			/* 16 */ imports.NewTable("TCEFTextfieldComponent_SelectedText", 0), // property SelectedText
			/* 17 */ imports.NewTable("TCEFTextfieldComponent_SelectedRange", 0), // property SelectedRange
			/* 18 */ imports.NewTable("TCEFTextfieldComponent_CursorPosition", 0), // property CursorPosition
			/* 19 */ imports.NewTable("TCEFTextfieldComponent_PlaceholderText", 0), // property PlaceholderText
			/* 20 */ imports.NewTable("TCEFTextfieldComponent_HasSelection", 0), // property HasSelection
			/* 21 */ imports.NewTable("TCEFTextfieldComponent_OnTextfieldKeyEvent", 0), // event OnTextfieldKeyEvent
			/* 22 */ imports.NewTable("TCEFTextfieldComponent_OnAfterUserAction", 0), // event OnAfterUserAction
		}
	})
	return cEFTextfieldComponentImport
}
