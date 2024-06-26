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

// ICefLabelButton Parent: ICefButton
//
//	LabelButton is a button with optional text and/or icon. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_label_button_capi.h">CEF source file: /include/capi/views/cef_label_button_capi.h (cef_label_button_t)</a>
type ICefLabelButton interface {
	ICefButton
	// AsMenuButton
	//  Returns this LabelButton as a MenuButton or NULL if this is not a
	//  MenuButton.
	AsMenuButton() ICefMenuButton // function
	// GetText
	//  Returns the text shown on the LabelButton.
	GetText() string // function
	// GetImage
	//  Returns the image shown for |button_state|. If no image exists for that
	//  state then the image for CEF_BUTTON_STATE_NORMAL will be returned.
	GetImage(buttonstate TCefButtonState) ICefImage // function
	// SetText
	//  Sets the text shown on the LabelButton. By default |text| will also be
	//  used as the accessible name.
	SetText(text string) // procedure
	// SetImage
	//  Sets the image shown for |button_state|. When this Button is drawn if no
	//  image exists for the current state then the image for
	//  CEF_BUTTON_STATE_NORMAL, if any, will be shown.
	SetImage(buttonstate TCefButtonState, image ICefImage) // procedure
	// SetTextColor
	//  Sets the text color shown for the specified button |for_state| to |color|.
	SetTextColor(forstate TCefButtonState, color TCefColor) // procedure
	// SetEnabledTextColors
	//  Sets the text colors shown for the non-disabled states to |color|.
	SetEnabledTextColors(color TCefColor) // procedure
	// SetFontList
	//  Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
	//  where: - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names(case-
	//  sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//  Here are examples of valid font description strings: - "Arial, Helvetica,
	//  Bold Italic 14px" - "Arial, 14px"
	SetFontList(fontlist string) // procedure
	// SetHorizontalAlignment
	//  Sets the horizontal alignment; reversed in RTL. Default is
	//  CEF_HORIZONTAL_ALIGNMENT_CENTER.
	SetHorizontalAlignment(alignment TCefHorizontalAlignment) // procedure
	// SetMinimumSize
	//  Reset the minimum size of this LabelButton to |size|.
	SetMinimumSize(size *TCefSize) // procedure
	// SetMaximumSize
	//  Reset the maximum size of this LabelButton to |size|.
	SetMaximumSize(size *TCefSize) // procedure
}

// TCefLabelButton Parent: TCefButton
//
//	LabelButton is a button with optional text and/or icon. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_label_button_capi.h">CEF source file: /include/capi/views/cef_label_button_capi.h (cef_label_button_t)</a>
type TCefLabelButton struct {
	TCefButton
}

// LabelButtonRef -> ICefLabelButton
var LabelButtonRef labelButton

// labelButton TCefLabelButton Ref
type labelButton uintptr

// UnWrap
//
//	Returns a ICefLabelButton instance using a PCefLabelButton data pointer.
func (m *labelButton) UnWrap(data uintptr) ICefLabelButton {
	var resultCefLabelButton uintptr
	labelButtonImportAPI().SysCallN(12, uintptr(data), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

// CreateLabelButton
//
//	Create a new LabelButton. A |delegate| must be provided to handle the button
//	click. |text| will be shown on the LabelButton and used as the default
//	accessible name.
func (m *labelButton) CreateLabelButton(delegate ICefButtonDelegate, text string) ICefLabelButton {
	var resultCefLabelButton uintptr
	labelButtonImportAPI().SysCallN(1, GetObjectUintptr(delegate), PascalStr(text), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

func (m *TCefLabelButton) AsMenuButton() ICefMenuButton {
	var resultCefMenuButton uintptr
	labelButtonImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMenuButton)))
	return AsCefMenuButton(resultCefMenuButton)
}

func (m *TCefLabelButton) GetText() string {
	r1 := labelButtonImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefLabelButton) GetImage(buttonstate TCefButtonState) ICefImage {
	var resultCefImage uintptr
	labelButtonImportAPI().SysCallN(2, m.Instance(), uintptr(buttonstate), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefLabelButton) SetText(text string) {
	labelButtonImportAPI().SysCallN(10, m.Instance(), PascalStr(text))
}

func (m *TCefLabelButton) SetImage(buttonstate TCefButtonState, image ICefImage) {
	labelButtonImportAPI().SysCallN(7, m.Instance(), uintptr(buttonstate), GetObjectUintptr(image))
}

func (m *TCefLabelButton) SetTextColor(forstate TCefButtonState, color TCefColor) {
	labelButtonImportAPI().SysCallN(11, m.Instance(), uintptr(forstate), uintptr(color))
}

func (m *TCefLabelButton) SetEnabledTextColors(color TCefColor) {
	labelButtonImportAPI().SysCallN(4, m.Instance(), uintptr(color))
}

func (m *TCefLabelButton) SetFontList(fontlist string) {
	labelButtonImportAPI().SysCallN(5, m.Instance(), PascalStr(fontlist))
}

func (m *TCefLabelButton) SetHorizontalAlignment(alignment TCefHorizontalAlignment) {
	labelButtonImportAPI().SysCallN(6, m.Instance(), uintptr(alignment))
}

func (m *TCefLabelButton) SetMinimumSize(size *TCefSize) {
	labelButtonImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefLabelButton) SetMaximumSize(size *TCefSize) {
	labelButtonImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(size)))
}

var (
	labelButtonImport       *imports.Imports = nil
	labelButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefLabelButton_AsMenuButton", 0),
		/*1*/ imports.NewTable("CefLabelButton_CreateLabelButton", 0),
		/*2*/ imports.NewTable("CefLabelButton_GetImage", 0),
		/*3*/ imports.NewTable("CefLabelButton_GetText", 0),
		/*4*/ imports.NewTable("CefLabelButton_SetEnabledTextColors", 0),
		/*5*/ imports.NewTable("CefLabelButton_SetFontList", 0),
		/*6*/ imports.NewTable("CefLabelButton_SetHorizontalAlignment", 0),
		/*7*/ imports.NewTable("CefLabelButton_SetImage", 0),
		/*8*/ imports.NewTable("CefLabelButton_SetMaximumSize", 0),
		/*9*/ imports.NewTable("CefLabelButton_SetMinimumSize", 0),
		/*10*/ imports.NewTable("CefLabelButton_SetText", 0),
		/*11*/ imports.NewTable("CefLabelButton_SetTextColor", 0),
		/*12*/ imports.NewTable("CefLabelButton_UnWrap", 0),
	}
)

func labelButtonImportAPI() *imports.Imports {
	if labelButtonImport == nil {
		labelButtonImport = NewDefaultImports()
		labelButtonImport.SetImportTable(labelButtonImportTables)
		labelButtonImportTables = nil
	}
	return labelButtonImport
}
