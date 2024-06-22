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

// ICEFLabelButtonComponent Parent: ICEFButtonComponent
type ICEFLabelButtonComponent interface {
	ICEFButtonComponent
	// Text
	//  Gets and sets the text shown on the LabelButton. By default |text| will also be
	//  used as the accessible name.
	Text() string // property
	// SetText Set Text
	SetText(AValue string) // property
	// Image
	//  Returns the image shown for |button_state|. If no image exists for that
	//  state then the image for CEF_BUTTON_STATE_NORMAL will be returned.
	Image(buttonstate TCefButtonState) ICefImage // property
	// SetImage Set Image
	SetImage(buttonstate TCefButtonState, AValue ICefImage) // property
	// AsMenuButton
	//  Returns this LabelButton as a MenuButton or NULL if this is not a
	//  MenuButton.
	AsMenuButton() ICefMenuButton // property
	// CreateLabelButton
	//  Create a new LabelButton. |aText| will be shown on the LabelButton and used as the default
	//  accessible name.
	CreateLabelButton(aText string) // procedure
	// SetTextColor
	//  Sets the text color shown for the specified button |for_state| to |color|.
	SetTextColor(forstate TCefButtonState, color TCefColor) // procedure
	// SetEnabledTextColors
	//  Sets the text colors shown for the non-disabled states to |color|.
	SetEnabledTextColors(color TCefColor) // procedure
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
	// SetHorizontalAlignment
	//  Sets the horizontal alignment; reversed in RTL. Default is
	//  CEF_HORIZONTAL_ALIGNMENT_CENTER.
	SetHorizontalAlignment(alignment TCefHorizontalAlignment) // procedure
	// SetMinimumSizeForCefSize
	//  Reset the minimum size of this LabelButton to |size|.
	SetMinimumSizeForCefSize(size *TCefSize) // procedure
	// SetMaximumSizeForCefSize
	//  Reset the maximum size of this LabelButton to |size|.
	SetMaximumSizeForCefSize(size *TCefSize) // procedure
}

// TCEFLabelButtonComponent Parent: TCEFButtonComponent
type TCEFLabelButtonComponent struct {
	TCEFButtonComponent
}

func NewCEFLabelButtonComponent(aOwner IComponent) ICEFLabelButtonComponent {
	r1 := labelButtonComponentImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsCEFLabelButtonComponent(r1)
}

func (m *TCEFLabelButtonComponent) Text() string {
	r1 := labelButtonComponentImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFLabelButtonComponent) SetText(AValue string) {
	labelButtonComponentImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFLabelButtonComponent) Image(buttonstate TCefButtonState) ICefImage {
	var resultCefImage uintptr
	labelButtonComponentImportAPI().SysCallN(3, 0, m.Instance(), uintptr(buttonstate), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCEFLabelButtonComponent) SetImage(buttonstate TCefButtonState, AValue ICefImage) {
	labelButtonComponentImportAPI().SysCallN(3, 1, m.Instance(), uintptr(buttonstate), GetObjectUintptr(AValue), uintptr(buttonstate), GetObjectUintptr(AValue))
}

func (m *TCEFLabelButtonComponent) AsMenuButton() ICefMenuButton {
	var resultCefMenuButton uintptr
	labelButtonComponentImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMenuButton)))
	return AsCefMenuButton(resultCefMenuButton)
}

func (m *TCEFLabelButtonComponent) CreateLabelButton(aText string) {
	labelButtonComponentImportAPI().SysCallN(2, m.Instance(), PascalStr(aText))
}

func (m *TCEFLabelButtonComponent) SetTextColor(forstate TCefButtonState, color TCefColor) {
	labelButtonComponentImportAPI().SysCallN(9, m.Instance(), uintptr(forstate), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetEnabledTextColors(color TCefColor) {
	labelButtonComponentImportAPI().SysCallN(4, m.Instance(), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetFontList(fontlist string) {
	labelButtonComponentImportAPI().SysCallN(5, m.Instance(), PascalStr(fontlist))
}

func (m *TCEFLabelButtonComponent) SetHorizontalAlignment(alignment TCefHorizontalAlignment) {
	labelButtonComponentImportAPI().SysCallN(6, m.Instance(), uintptr(alignment))
}

func (m *TCEFLabelButtonComponent) SetMinimumSizeForCefSize(size *TCefSize) {
	labelButtonComponentImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCEFLabelButtonComponent) SetMaximumSizeForCefSize(size *TCefSize) {
	labelButtonComponentImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(size)))
}

var (
	labelButtonComponentImport       *imports.Imports = nil
	labelButtonComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFLabelButtonComponent_AsMenuButton", 0),
		/*1*/ imports.NewTable("CEFLabelButtonComponent_Create", 0),
		/*2*/ imports.NewTable("CEFLabelButtonComponent_CreateLabelButton", 0),
		/*3*/ imports.NewTable("CEFLabelButtonComponent_Image", 0),
		/*4*/ imports.NewTable("CEFLabelButtonComponent_SetEnabledTextColors", 0),
		/*5*/ imports.NewTable("CEFLabelButtonComponent_SetFontList", 0),
		/*6*/ imports.NewTable("CEFLabelButtonComponent_SetHorizontalAlignment", 0),
		/*7*/ imports.NewTable("CEFLabelButtonComponent_SetMaximumSizeForCefSize", 0),
		/*8*/ imports.NewTable("CEFLabelButtonComponent_SetMinimumSizeForCefSize", 0),
		/*9*/ imports.NewTable("CEFLabelButtonComponent_SetTextColor", 0),
		/*10*/ imports.NewTable("CEFLabelButtonComponent_Text", 0),
	}
)

func labelButtonComponentImportAPI() *imports.Imports {
	if labelButtonComponentImport == nil {
		labelButtonComponentImport = NewDefaultImports()
		labelButtonComponentImport.SetImportTable(labelButtonComponentImportTables)
		labelButtonComponentImportTables = nil
	}
	return labelButtonComponentImport
}
