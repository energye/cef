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

// ICEFTextfieldComponent Parent: ICEFViewComponent
type ICEFTextfieldComponent interface {
	ICEFViewComponent
	// PasswordInput
	//  Returns true(1) if the text will be displayed as asterisks.
	PasswordInput() bool // property
	// SetPasswordInput Set PasswordInput
	SetPasswordInput(AValue bool) // property
	// ReadOnly
	//  Returns true(1) if the text is read-only.
	ReadOnly() bool // property
	// SetReadOnly Set ReadOnly
	SetReadOnly(AValue bool) // property
	// Text
	//  Returns the currently displayed text.
	Text() string // property
	// SetText Set Text
	SetText(AValue string) // property
	// SelectedText
	//  Returns the currently selected text.
	SelectedText() string // property
	// SelectedRange
	//  Returns the selected logical text range.
	SelectedRange() (resultCefRange TCefRange) // property
	// SetSelectedRange Set SelectedRange
	SetSelectedRange(AValue *TCefRange) // property
	// CursorPosition
	//  Returns the current cursor position.
	CursorPosition() NativeUInt // property
	// TextColor
	//  Returns the text color.
	TextColor() TCefColor // property
	// SetTextColor Set TextColor
	SetTextColor(AValue TCefColor) // property
	// SelectionTextColor
	//  Returns the selection text color.
	SelectionTextColor() TCefColor // property
	// SetSelectionTextColor Set SelectionTextColor
	SetSelectionTextColor(AValue TCefColor) // property
	// SelectionBackgroundColor
	//  Returns the selection background color.
	SelectionBackgroundColor() TCefColor // property
	// SetSelectionBackgroundColor Set SelectionBackgroundColor
	SetSelectionBackgroundColor(AValue TCefColor) // property
	// PlaceholderText
	//  Returns the placeholder text that will be displayed when the Textfield is
	//  NULL.
	PlaceholderText() string // property
	// SetPlaceholderText Set PlaceholderText
	SetPlaceholderText(AValue string) // property
	// HasSelection
	//  Returns true(1) if there is any selected text.
	HasSelection() bool // property
	// IsCommandEnabled
	//  Returns true(1) if the action associated with the specified command id is
	//  enabled. See additional comments on execute_command().
	IsCommandEnabled(commandid TCefTextFieldCommands) bool // function
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
	//  Selects all text. If |reversed| is true(1) the range will end at the
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
	//  - STYLES is an optional space-separated list of style names(case-sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//  </code>
	//  Here are examples of valid font description strings:
	//  <code>
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
	//  </code>
	SetFontList(fontlist string) // procedure
	// ApplyTextColor
	//  Applies |color| to the specified |range| without changing the default
	//  color. If |range| is NULL the color will be set on the complete text
	//  contents.
	ApplyTextColor(color TCefColor, range_ *TCefRange) // procedure
	// ApplyTextStyle
	//  Applies |style| to the specified |range| without changing the default
	//  style. If |add| is true(1) the style will be added, otherwise the style
	//  will be removed. If |range| is NULL the style will be set on the complete
	//  text contents.
	ApplyTextStyle(style TCefTextStyle, add bool, range_ *TCefRange) // procedure
	// ExecuteCommand
	//  Performs the action associated with the specified command id.
	ExecuteCommand(commandid TCefTextFieldCommands) // procedure
	// ClearEditHistory
	//  Clears Edit history.
	ClearEditHistory() // procedure
	// SetAccessibleName
	//  Set the accessible name that will be exposed to assistive technology(AT).
	SetAccessibleName(name string) // procedure
	// SetPlaceholderTextColor
	//  Sets the placeholder text color.
	SetPlaceholderTextColor(color TCefColor) // procedure
	// SetOnTextfieldKeyEvent
	//  Called when |textfield| recieves a keyboard event. |event| contains
	//  information about the keyboard event. Return true(1) if the keyboard
	//  event was handled or false(0) otherwise for default handling.
	SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEvent) // property event
	// SetOnAfterUserAction
	//  Called after performing a user action that may change |textfield|.
	SetOnAfterUserAction(fn TOnAfterUserAction) // property event
}

// TCEFTextfieldComponent Parent: TCEFViewComponent
type TCEFTextfieldComponent struct {
	TCEFViewComponent
	textfieldKeyEventPtr uintptr
	afterUserActionPtr   uintptr
}

func NewCEFTextfieldComponent(aOwner IComponent) ICEFTextfieldComponent {
	r1 := extfieldComponentImportAPI().SysCallN(5, GetObjectUintptr(aOwner))
	return AsCEFTextfieldComponent(r1)
}

func (m *TCEFTextfieldComponent) PasswordInput() bool {
	r1 := extfieldComponentImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) SetPasswordInput(AValue bool) {
	extfieldComponentImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFTextfieldComponent) ReadOnly() bool {
	r1 := extfieldComponentImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) SetReadOnly(AValue bool) {
	extfieldComponentImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFTextfieldComponent) Text() string {
	r1 := extfieldComponentImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SetText(AValue string) {
	extfieldComponentImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFTextfieldComponent) SelectedText() string {
	r1 := extfieldComponentImportAPI().SysCallN(17, m.Instance())
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SelectedRange() (resultCefRange TCefRange) {
	extfieldComponentImportAPI().SysCallN(16, 0, m.Instance(), uintptr(unsafePointer(&resultCefRange)), uintptr(unsafePointer(&resultCefRange)))
	return
}

func (m *TCEFTextfieldComponent) SetSelectedRange(AValue *TCefRange) {
	extfieldComponentImportAPI().SysCallN(16, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFTextfieldComponent) CursorPosition() NativeUInt {
	r1 := extfieldComponentImportAPI().SysCallN(7, m.Instance())
	return NativeUInt(r1)
}

func (m *TCEFTextfieldComponent) TextColor() TCefColor {
	r1 := extfieldComponentImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetTextColor(AValue TCefColor) {
	extfieldComponentImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) SelectionTextColor() TCefColor {
	r1 := extfieldComponentImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetSelectionTextColor(AValue TCefColor) {
	extfieldComponentImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) SelectionBackgroundColor() TCefColor {
	r1 := extfieldComponentImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetSelectionBackgroundColor(AValue TCefColor) {
	extfieldComponentImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) PlaceholderText() string {
	r1 := extfieldComponentImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SetPlaceholderText(AValue string) {
	extfieldComponentImportAPI().SysCallN(13, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFTextfieldComponent) HasSelection() bool {
	r1 := extfieldComponentImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) IsCommandEnabled(commandid TCefTextFieldCommands) bool {
	r1 := extfieldComponentImportAPI().SysCallN(11, m.Instance(), uintptr(commandid))
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) CreateTextField() {
	extfieldComponentImportAPI().SysCallN(6, m.Instance())
}

func (m *TCEFTextfieldComponent) AppendText(text string) {
	extfieldComponentImportAPI().SysCallN(0, m.Instance(), PascalStr(text))
}

func (m *TCEFTextfieldComponent) InsertOrReplaceText(text string) {
	extfieldComponentImportAPI().SysCallN(10, m.Instance(), PascalStr(text))
}

func (m *TCEFTextfieldComponent) SelectAll(reversed bool) {
	extfieldComponentImportAPI().SysCallN(15, m.Instance(), PascalBool(reversed))
}

func (m *TCEFTextfieldComponent) ClearSelection() {
	extfieldComponentImportAPI().SysCallN(4, m.Instance())
}

func (m *TCEFTextfieldComponent) SetFontList(fontlist string) {
	extfieldComponentImportAPI().SysCallN(21, m.Instance(), PascalStr(fontlist))
}

func (m *TCEFTextfieldComponent) ApplyTextColor(color TCefColor, range_ *TCefRange) {
	extfieldComponentImportAPI().SysCallN(1, m.Instance(), uintptr(color), uintptr(unsafePointer(range_)))
}

func (m *TCEFTextfieldComponent) ApplyTextStyle(style TCefTextStyle, add bool, range_ *TCefRange) {
	extfieldComponentImportAPI().SysCallN(2, m.Instance(), uintptr(style), PascalBool(add), uintptr(unsafePointer(range_)))
}

func (m *TCEFTextfieldComponent) ExecuteCommand(commandid TCefTextFieldCommands) {
	extfieldComponentImportAPI().SysCallN(8, m.Instance(), uintptr(commandid))
}

func (m *TCEFTextfieldComponent) ClearEditHistory() {
	extfieldComponentImportAPI().SysCallN(3, m.Instance())
}

func (m *TCEFTextfieldComponent) SetAccessibleName(name string) {
	extfieldComponentImportAPI().SysCallN(20, m.Instance(), PascalStr(name))
}

func (m *TCEFTextfieldComponent) SetPlaceholderTextColor(color TCefColor) {
	extfieldComponentImportAPI().SysCallN(24, m.Instance(), uintptr(color))
}

func (m *TCEFTextfieldComponent) SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEvent) {
	if m.textfieldKeyEventPtr != 0 {
		RemoveEventElement(m.textfieldKeyEventPtr)
	}
	m.textfieldKeyEventPtr = MakeEventDataPtr(fn)
	extfieldComponentImportAPI().SysCallN(23, m.Instance(), m.textfieldKeyEventPtr)
}

func (m *TCEFTextfieldComponent) SetOnAfterUserAction(fn TOnAfterUserAction) {
	if m.afterUserActionPtr != 0 {
		RemoveEventElement(m.afterUserActionPtr)
	}
	m.afterUserActionPtr = MakeEventDataPtr(fn)
	extfieldComponentImportAPI().SysCallN(22, m.Instance(), m.afterUserActionPtr)
}

var (
	extfieldComponentImport       *imports.Imports = nil
	extfieldComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFTextfieldComponent_AppendText", 0),
		/*1*/ imports.NewTable("CEFTextfieldComponent_ApplyTextColor", 0),
		/*2*/ imports.NewTable("CEFTextfieldComponent_ApplyTextStyle", 0),
		/*3*/ imports.NewTable("CEFTextfieldComponent_ClearEditHistory", 0),
		/*4*/ imports.NewTable("CEFTextfieldComponent_ClearSelection", 0),
		/*5*/ imports.NewTable("CEFTextfieldComponent_Create", 0),
		/*6*/ imports.NewTable("CEFTextfieldComponent_CreateTextField", 0),
		/*7*/ imports.NewTable("CEFTextfieldComponent_CursorPosition", 0),
		/*8*/ imports.NewTable("CEFTextfieldComponent_ExecuteCommand", 0),
		/*9*/ imports.NewTable("CEFTextfieldComponent_HasSelection", 0),
		/*10*/ imports.NewTable("CEFTextfieldComponent_InsertOrReplaceText", 0),
		/*11*/ imports.NewTable("CEFTextfieldComponent_IsCommandEnabled", 0),
		/*12*/ imports.NewTable("CEFTextfieldComponent_PasswordInput", 0),
		/*13*/ imports.NewTable("CEFTextfieldComponent_PlaceholderText", 0),
		/*14*/ imports.NewTable("CEFTextfieldComponent_ReadOnly", 0),
		/*15*/ imports.NewTable("CEFTextfieldComponent_SelectAll", 0),
		/*16*/ imports.NewTable("CEFTextfieldComponent_SelectedRange", 0),
		/*17*/ imports.NewTable("CEFTextfieldComponent_SelectedText", 0),
		/*18*/ imports.NewTable("CEFTextfieldComponent_SelectionBackgroundColor", 0),
		/*19*/ imports.NewTable("CEFTextfieldComponent_SelectionTextColor", 0),
		/*20*/ imports.NewTable("CEFTextfieldComponent_SetAccessibleName", 0),
		/*21*/ imports.NewTable("CEFTextfieldComponent_SetFontList", 0),
		/*22*/ imports.NewTable("CEFTextfieldComponent_SetOnAfterUserAction", 0),
		/*23*/ imports.NewTable("CEFTextfieldComponent_SetOnTextfieldKeyEvent", 0),
		/*24*/ imports.NewTable("CEFTextfieldComponent_SetPlaceholderTextColor", 0),
		/*25*/ imports.NewTable("CEFTextfieldComponent_Text", 0),
		/*26*/ imports.NewTable("CEFTextfieldComponent_TextColor", 0),
	}
)

func extfieldComponentImportAPI() *imports.Imports {
	if extfieldComponentImport == nil {
		extfieldComponentImport = NewDefaultImports()
		extfieldComponentImport.SetImportTable(extfieldComponentImportTables)
		extfieldComponentImportTables = nil
	}
	return extfieldComponentImport
}
