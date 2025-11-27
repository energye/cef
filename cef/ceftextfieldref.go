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

// ICefTextfield Parent: ICefViewRef
type ICefTextfield interface {
	ICefViewRef
	// IsPasswordInput
	//  Returns true (1) if the text will be displayed as asterisks.
	IsPasswordInput() bool // function
	// IsReadOnly
	//  Returns true (1) if the text is read-only.
	IsReadOnly() bool // function
	// GetText
	//  Returns the currently displayed text.
	GetText() string // function
	// HasSelection
	//  Returns true (1) if there is any selected text.
	HasSelection() bool // function
	// GetSelectedText
	//  Returns the currently selected text.
	GetSelectedText() string // function
	// GetSelectedRange
	//  Returns the selected logical text range.
	GetSelectedRange() TCefRange // function
	// GetCursorPosition
	//  Returns the current cursor position.
	GetCursorPosition() cefTypes.NativeUInt // function
	// GetTextColor
	//  Returns the text color.
	GetTextColor() cefTypes.TCefColor // function
	// GetSelectionTextColor
	//  Returns the selection text color.
	GetSelectionTextColor() cefTypes.TCefColor // function
	// GetSelectionBackgroundColor
	//  Returns the selection background color.
	GetSelectionBackgroundColor() cefTypes.TCefColor // function
	// IsCommandEnabled
	//  Returns true (1) if the action associated with the specified command id is
	//  enabled. See additional comments on execute_command().
	IsCommandEnabled(commandId cefTypes.TCefTextFieldCommands) bool // function
	// GetPlaceholderText
	//  Returns the placeholder text that will be displayed when the Textfield is
	//  NULL.
	GetPlaceholderText() string // function
	// SetPasswordInput
	//  Sets whether the text will be displayed as asterisks.
	SetPasswordInput(passwordInput bool) // procedure
	// SetReadOnly
	//  Sets whether the text will read-only.
	SetReadOnly(readOnly bool) // procedure
	// SetText
	//  Sets the contents to |text|. The cursor will be moved to end of the text
	//  if the current position is outside of the text range.
	SetText(text string) // procedure
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
	// SelectRange
	//  Selects the specified logical text range.
	SelectRange(range_ TCefRange) // procedure
	// SetTextColor
	//  Sets the text color.
	SetTextColor(color cefTypes.TCefColor) // procedure
	// SetSelectionTextColor
	//  Sets the selection text color.
	SetSelectionTextColor(color cefTypes.TCefColor) // procedure
	// SetSelectionBackgroundColor
	//  Sets the selection background color.
	SetSelectionBackgroundColor(color cefTypes.TCefColor) // procedure
	// SetFontList
	//  Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
	//  where:
	//  - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names (case-sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//
	//  Here are examples of valid font description strings:
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
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
	// SetPlaceholderText
	//  Sets the placeholder text that will be displayed when the Textfield is
	//  NULL.
	SetPlaceholderText(text string) // procedure
	// SetPlaceholderTextColor
	//  Sets the placeholder text color.
	SetPlaceholderTextColor(color cefTypes.TCefColor) // procedure
	// SetAccessibleName
	//  Set the accessible name that will be exposed to assistive technology (AT).
	SetAccessibleName(name string) // procedure
}

// ICefTextfieldRef Parent: ICefTextfield
type ICefTextfieldRef interface {
	ICefTextfield
	AsIntfTextfield() uintptr
	AsIntfView() uintptr
}

type TCefTextfieldRef struct {
	TCefViewRef
}

func (m *TCefTextfieldRef) IsPasswordInput() bool {
	if !m.IsValid() {
		return false
	}
	r := cefTextfieldRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefTextfieldRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefTextfieldRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefTextfieldRef) GetText() string {
	if !m.IsValid() {
		return ""
	}
	r := cefTextfieldRefAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCefTextfieldRef) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := cefTextfieldRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefTextfieldRef) GetSelectedText() string {
	if !m.IsValid() {
		return ""
	}
	r := cefTextfieldRefAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCefTextfieldRef) GetSelectedRange() (result TCefRange) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefTextfieldRef) GetCursorPosition() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefTextfieldRefAPI().SysCallN(7, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefTextfieldRef) GetTextColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefTextfieldRefAPI().SysCallN(8, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefTextfieldRef) GetSelectionTextColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefTextfieldRefAPI().SysCallN(9, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefTextfieldRef) GetSelectionBackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefTextfieldRefAPI().SysCallN(10, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefTextfieldRef) IsCommandEnabled(commandId cefTypes.TCefTextFieldCommands) bool {
	if !m.IsValid() {
		return false
	}
	r := cefTextfieldRefAPI().SysCallN(11, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefTextfieldRef) GetPlaceholderText() string {
	if !m.IsValid() {
		return ""
	}
	r := cefTextfieldRefAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TCefTextfieldRef) SetPasswordInput(passwordInput bool) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(15, m.Instance(), api.PasBool(passwordInput))
}

func (m *TCefTextfieldRef) SetReadOnly(readOnly bool) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(16, m.Instance(), api.PasBool(readOnly))
}

func (m *TCefTextfieldRef) SetText(text string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(17, m.Instance(), api.PasStr(text))
}

func (m *TCefTextfieldRef) AppendText(text string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(18, m.Instance(), api.PasStr(text))
}

func (m *TCefTextfieldRef) InsertOrReplaceText(text string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(19, m.Instance(), api.PasStr(text))
}

func (m *TCefTextfieldRef) SelectAll(reversed bool) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(20, m.Instance(), api.PasBool(reversed))
}

func (m *TCefTextfieldRef) ClearSelection() {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(21, m.Instance())
}

func (m *TCefTextfieldRef) SelectRange(range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCefTextfieldRef) SetTextColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(23, m.Instance(), uintptr(color))
}

func (m *TCefTextfieldRef) SetSelectionTextColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(24, m.Instance(), uintptr(color))
}

func (m *TCefTextfieldRef) SetSelectionBackgroundColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(25, m.Instance(), uintptr(color))
}

func (m *TCefTextfieldRef) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(26, m.Instance(), api.PasStr(fontList))
}

func (m *TCefTextfieldRef) ApplyTextColor(color cefTypes.TCefColor, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(27, m.Instance(), uintptr(color), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCefTextfieldRef) ApplyTextStyle(style cefTypes.TCefTextStyle, add bool, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(28, m.Instance(), uintptr(style), api.PasBool(add), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCefTextfieldRef) ExecuteCommand(commandId cefTypes.TCefTextFieldCommands) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(29, m.Instance(), uintptr(commandId))
}

func (m *TCefTextfieldRef) ClearEditHistory() {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(30, m.Instance())
}

func (m *TCefTextfieldRef) SetPlaceholderText(text string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(31, m.Instance(), api.PasStr(text))
}

func (m *TCefTextfieldRef) SetPlaceholderTextColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(32, m.Instance(), uintptr(color))
}

func (m *TCefTextfieldRef) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	cefTextfieldRefAPI().SysCallN(33, m.Instance(), api.PasStr(name))
}

func (m *TCefTextfieldRef) AsIntfTextfield() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefTextfieldRef) AsIntfView() uintptr {
	return m.GetIntfPointer(1)
}

// TextfieldRef  is static instance
var TextfieldRef _TextfieldRefClass

// _TextfieldRefClass is class type defined by TCefTextfieldRef
type _TextfieldRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefTextfield instance using a PCefTextfield data pointer.
func (_TextfieldRefClass) UnWrapWithPointer(data uintptr) (result ICefTextfield) {
	var resultPtr uintptr
	cefTextfieldRefAPI().SysCallN(13, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTextfieldRef(resultPtr)
	return
}

// CreateTextField
//
//	Create a new Textfield.
func (_TextfieldRefClass) CreateTextField(delegate IEngTextfieldDelegate) (result ICefTextfield) {
	var resultPtr uintptr
	cefTextfieldRefAPI().SysCallN(14, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTextfieldRef(resultPtr)
	return
}

// NewTextfieldRef class constructor
func NewTextfieldRef(data uintptr) ICefTextfieldRef {
	var textfieldPtr uintptr // ICefTextfield
	var viewPtr uintptr      // ICefView
	r := cefTextfieldRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&textfieldPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefTextfieldRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, textfieldPtr)
		intf.SetIntfPointer(1, viewPtr)
	}
	return ret
}

var (
	cefTextfieldRefOnce   base.Once
	cefTextfieldRefImport *imports.Imports = nil
)

func cefTextfieldRefAPI() *imports.Imports {
	cefTextfieldRefOnce.Do(func() {
		cefTextfieldRefImport = api.NewDefaultImports()
		cefTextfieldRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefTextfieldRef_Create", 0), // constructor NewTextfieldRef
			/* 1 */ imports.NewTable("TCefTextfieldRef_IsPasswordInput", 0), // function IsPasswordInput
			/* 2 */ imports.NewTable("TCefTextfieldRef_IsReadOnly", 0), // function IsReadOnly
			/* 3 */ imports.NewTable("TCefTextfieldRef_GetText", 0), // function GetText
			/* 4 */ imports.NewTable("TCefTextfieldRef_HasSelection", 0), // function HasSelection
			/* 5 */ imports.NewTable("TCefTextfieldRef_GetSelectedText", 0), // function GetSelectedText
			/* 6 */ imports.NewTable("TCefTextfieldRef_GetSelectedRange", 0), // function GetSelectedRange
			/* 7 */ imports.NewTable("TCefTextfieldRef_GetCursorPosition", 0), // function GetCursorPosition
			/* 8 */ imports.NewTable("TCefTextfieldRef_GetTextColor", 0), // function GetTextColor
			/* 9 */ imports.NewTable("TCefTextfieldRef_GetSelectionTextColor", 0), // function GetSelectionTextColor
			/* 10 */ imports.NewTable("TCefTextfieldRef_GetSelectionBackgroundColor", 0), // function GetSelectionBackgroundColor
			/* 11 */ imports.NewTable("TCefTextfieldRef_IsCommandEnabled", 0), // function IsCommandEnabled
			/* 12 */ imports.NewTable("TCefTextfieldRef_GetPlaceholderText", 0), // function GetPlaceholderText
			/* 13 */ imports.NewTable("TCefTextfieldRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 14 */ imports.NewTable("TCefTextfieldRef_CreateTextField", 0), // static function CreateTextField
			/* 15 */ imports.NewTable("TCefTextfieldRef_SetPasswordInput", 0), // procedure SetPasswordInput
			/* 16 */ imports.NewTable("TCefTextfieldRef_SetReadOnly", 0), // procedure SetReadOnly
			/* 17 */ imports.NewTable("TCefTextfieldRef_SetText", 0), // procedure SetText
			/* 18 */ imports.NewTable("TCefTextfieldRef_AppendText", 0), // procedure AppendText
			/* 19 */ imports.NewTable("TCefTextfieldRef_InsertOrReplaceText", 0), // procedure InsertOrReplaceText
			/* 20 */ imports.NewTable("TCefTextfieldRef_SelectAll", 0), // procedure SelectAll
			/* 21 */ imports.NewTable("TCefTextfieldRef_ClearSelection", 0), // procedure ClearSelection
			/* 22 */ imports.NewTable("TCefTextfieldRef_SelectRange", 0), // procedure SelectRange
			/* 23 */ imports.NewTable("TCefTextfieldRef_SetTextColor", 0), // procedure SetTextColor
			/* 24 */ imports.NewTable("TCefTextfieldRef_SetSelectionTextColor", 0), // procedure SetSelectionTextColor
			/* 25 */ imports.NewTable("TCefTextfieldRef_SetSelectionBackgroundColor", 0), // procedure SetSelectionBackgroundColor
			/* 26 */ imports.NewTable("TCefTextfieldRef_SetFontList", 0), // procedure SetFontList
			/* 27 */ imports.NewTable("TCefTextfieldRef_ApplyTextColor", 0), // procedure ApplyTextColor
			/* 28 */ imports.NewTable("TCefTextfieldRef_ApplyTextStyle", 0), // procedure ApplyTextStyle
			/* 29 */ imports.NewTable("TCefTextfieldRef_ExecuteCommand", 0), // procedure ExecuteCommand
			/* 30 */ imports.NewTable("TCefTextfieldRef_ClearEditHistory", 0), // procedure ClearEditHistory
			/* 31 */ imports.NewTable("TCefTextfieldRef_SetPlaceholderText", 0), // procedure SetPlaceholderText
			/* 32 */ imports.NewTable("TCefTextfieldRef_SetPlaceholderTextColor", 0), // procedure SetPlaceholderTextColor
			/* 33 */ imports.NewTable("TCefTextfieldRef_SetAccessibleName", 0), // procedure SetAccessibleName
		}
	})
	return cefTextfieldRefImport
}
