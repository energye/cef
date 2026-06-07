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

	cefTypes "github.com/energye/cef/147/types"
)

// ICEFLabelButtonComponent Parent: ICEFButtonComponent
type ICEFLabelButtonComponent interface {
	ICEFButtonComponent
	// CreateLabelButton
	//  Create a new LabelButton. |aText| will be shown on the LabelButton and used as the default
	//  accessible name.
	CreateLabelButton(text string) // procedure
	// SetTextColor
	//  Sets the text color shown for the specified button |for_state| to |color|.
	SetTextColor(forState cefTypes.TCefButtonState, color cefTypes.TCefColor) // procedure
	// SetEnabledTextColors
	//  Sets the text colors shown for the non-disabled states to |color|.
	SetEnabledTextColors(color cefTypes.TCefColor) // procedure
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
	// SetHorizontalAlignment
	//  Sets the horizontal alignment; reversed in RTL. Default is
	//  CEF_HORIZONTAL_ALIGNMENT_CENTER.
	SetHorizontalAlignment(alignment cefTypes.TCefHorizontalAlignment) // procedure
	// SetMinimumSize
	//  Reset the minimum size of this LabelButton to |size|.
	SetMinimumSize(size TCefSize) // procedure
	// SetMaximumSize
	//  Reset the maximum size of this LabelButton to |size|.
	SetMaximumSize(size TCefSize) // procedure
	// Text
	//  Gets and sets the text shown on the LabelButton. By default |text| will also be
	//  used as the accessible name.
	Text() string         // property Text Getter
	SetText(value string) // property Text Setter
	// Image
	//  Returns the image shown for |button_state|. If no image exists for that
	//  state then the image for CEF_BUTTON_STATE_NORMAL will be returned.
	Image(buttonState cefTypes.TCefButtonState) ICefImage           // property Image Getter
	SetImage(buttonState cefTypes.TCefButtonState, value ICefImage) // property Image Setter
	// AsMenuButton
	//  Returns this LabelButton as a MenuButton or NULL if this is not a
	//  MenuButton.
	AsMenuButton() ICefMenuButton // property AsMenuButton Getter
	AsIntfButtonDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFLabelButtonComponent struct {
	TCEFButtonComponent
}

func (m *TCEFLabelButtonComponent) CreateLabelButton(text string) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(1, m.Instance(), api.PasStr(text))
}

func (m *TCEFLabelButtonComponent) SetTextColor(forState cefTypes.TCefButtonState, color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(2, m.Instance(), uintptr(forState), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetEnabledTextColors(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(3, m.Instance(), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(4, m.Instance(), api.PasStr(fontList))
}

func (m *TCEFLabelButtonComponent) SetHorizontalAlignment(alignment cefTypes.TCefHorizontalAlignment) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(5, m.Instance(), uintptr(alignment))
}

func (m *TCEFLabelButtonComponent) SetMinimumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCEFLabelButtonComponent) SetMaximumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCEFLabelButtonComponent) Text() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFLabelButtonComponentAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFLabelButtonComponent) SetText(value string) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFLabelButtonComponent) Image(buttonState cefTypes.TCefButtonState) (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFLabelButtonComponentAPI().SysCallN(9, 0, m.Instance(), uintptr(buttonState), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCEFLabelButtonComponent) SetImage(buttonState cefTypes.TCefButtonState, value ICefImage) {
	if !m.IsValid() {
		return
	}
	cEFLabelButtonComponentAPI().SysCallN(9, 1, m.Instance(), uintptr(buttonState), base.GetObjectUintptr(value))
}

func (m *TCEFLabelButtonComponent) AsMenuButton() (result ICefMenuButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFLabelButtonComponentAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuButtonRef(resultPtr)
	return
}

func (m *TCEFLabelButtonComponent) AsIntfButtonDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFLabelButtonComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}

// NewLabelButtonComponent class constructor
func NewLabelButtonComponent(owner lcl.IComponent) ICEFLabelButtonComponent {
	var buttonDelegateEventsPtr uintptr // ICefButtonDelegateEvents
	var viewDelegateEventsPtr uintptr   // ICefViewDelegateEvents
	r := cEFLabelButtonComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&buttonDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFLabelButtonComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, buttonDelegateEventsPtr)
		intf.SetIntfPointer(1, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFLabelButtonComponentOnce   base.Once
	cEFLabelButtonComponentImport *imports.Imports = nil
)

func cEFLabelButtonComponentAPI() *imports.Imports {
	cEFLabelButtonComponentOnce.Do(func() {
		cEFLabelButtonComponentImport = api.NewDefaultImports()
		cEFLabelButtonComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFLabelButtonComponent_Create", 0), // constructor NewLabelButtonComponent
			/* 1 */ imports.NewTable("TCEFLabelButtonComponent_CreateLabelButton", 0), // procedure CreateLabelButton
			/* 2 */ imports.NewTable("TCEFLabelButtonComponent_SetTextColor", 0), // procedure SetTextColor
			/* 3 */ imports.NewTable("TCEFLabelButtonComponent_SetEnabledTextColors", 0), // procedure SetEnabledTextColors
			/* 4 */ imports.NewTable("TCEFLabelButtonComponent_SetFontList", 0), // procedure SetFontList
			/* 5 */ imports.NewTable("TCEFLabelButtonComponent_SetHorizontalAlignment", 0), // procedure SetHorizontalAlignment
			/* 6 */ imports.NewTable("TCEFLabelButtonComponent_SetMinimumSize", 0), // procedure SetMinimumSize
			/* 7 */ imports.NewTable("TCEFLabelButtonComponent_SetMaximumSize", 0), // procedure SetMaximumSize
			/* 8 */ imports.NewTable("TCEFLabelButtonComponent_Text", 0), // property Text
			/* 9 */ imports.NewTable("TCEFLabelButtonComponent_Image", 0), // property Image
			/* 10 */ imports.NewTable("TCEFLabelButtonComponent_AsMenuButton", 0), // property AsMenuButton
		}
	})
	return cEFLabelButtonComponentImport
}
