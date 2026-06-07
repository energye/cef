//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/127/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefLabelButton Parent: ICefButton
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
	GetImage(buttonState cefTypes.TCefButtonState) ICefImage // function
	// SetText
	//  Sets the text shown on the LabelButton. By default |text| will also be
	//  used as the accessible name.
	SetText(text string) // procedure
	// SetImage
	//  Sets the image shown for |button_state|. When this Button is drawn if no
	//  image exists for the current state then the image for
	//  CEF_BUTTON_STATE_NORMAL, if any, will be shown.
	SetImage(buttonState cefTypes.TCefButtonState, image ICefImage) // procedure
	// SetTextColor
	//  Sets the text color shown for the specified button |for_state| to |color|.
	SetTextColor(forState cefTypes.TCefButtonState, color cefTypes.TCefColor) // procedure
	// SetEnabledTextColors
	//  Sets the text colors shown for the non-disabled states to |color|.
	SetEnabledTextColors(color cefTypes.TCefColor) // procedure
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
}

// ICefLabelButtonRef Parent: ICefLabelButton ICefButtonRef
type ICefLabelButtonRef interface {
	ICefLabelButton
	ICefButtonRef
	AsIntfLabelButton() uintptr
	AsIntfButton() uintptr
	AsIntfView() uintptr
}

type TCefLabelButtonRef struct {
	TCefButtonRef
}

func (m *TCefLabelButtonRef) AsMenuButton() (result ICefMenuButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefLabelButtonRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuButtonRef(resultPtr)
	return
}

func (m *TCefLabelButtonRef) GetText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefLabelButtonRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefLabelButtonRef) GetImage(buttonState cefTypes.TCefButtonState) (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefLabelButtonRefAPI().SysCallN(3, m.Instance(), uintptr(buttonState), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCefLabelButtonRef) SetText(text string) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(6, m.Instance(), api.PasStr(text))
}

func (m *TCefLabelButtonRef) SetImage(buttonState cefTypes.TCefButtonState, image ICefImage) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(7, m.Instance(), uintptr(buttonState), base.GetObjectUintptr(image))
}

func (m *TCefLabelButtonRef) SetTextColor(forState cefTypes.TCefButtonState, color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(8, m.Instance(), uintptr(forState), uintptr(color))
}

func (m *TCefLabelButtonRef) SetEnabledTextColors(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(9, m.Instance(), uintptr(color))
}

func (m *TCefLabelButtonRef) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(10, m.Instance(), api.PasStr(fontList))
}

func (m *TCefLabelButtonRef) SetHorizontalAlignment(alignment cefTypes.TCefHorizontalAlignment) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(11, m.Instance(), uintptr(alignment))
}

func (m *TCefLabelButtonRef) SetMinimumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefLabelButtonRef) SetMaximumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefLabelButtonRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefLabelButtonRef) AsIntfLabelButton() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefLabelButtonRef) AsIntfButton() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefLabelButtonRef) AsIntfView() uintptr {
	return m.GetIntfPointer(2)
}

// LabelButtonRef  is static instance
var LabelButtonRef _LabelButtonRefClass

// _LabelButtonRefClass is class type defined by TCefLabelButtonRef
type _LabelButtonRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefLabelButton instance using a PCefLabelButton data pointer.
func (_LabelButtonRefClass) UnWrapWithPointer(data uintptr) (result ICefLabelButton) {
	var resultPtr uintptr
	cefLabelButtonRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLabelButtonRef(resultPtr)
	return
}

// CreateLabelButton
//
//	Create a new LabelButton. A |delegate| must be provided to handle the button
//	click. |text| will be shown on the LabelButton and used as the default
//	accessible name.
func (_LabelButtonRefClass) CreateLabelButton(delegate IEngButtonDelegate, text string) (result ICefLabelButton) {
	var resultPtr uintptr
	cefLabelButtonRefAPI().SysCallN(5, base.GetObjectUintptr(delegate), api.PasStr(text), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLabelButtonRef(resultPtr)
	return
}

// NewLabelButtonRef class constructor
func NewLabelButtonRef(data uintptr) ICefLabelButtonRef {
	var labelButtonPtr uintptr // ICefLabelButton
	var buttonPtr uintptr      // ICefButton
	var viewPtr uintptr        // ICefView
	r := cefLabelButtonRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&labelButtonPtr)), uintptr(base.UnsafePointer(&buttonPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefLabelButtonRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, labelButtonPtr)
		intf.SetIntfPointer(1, buttonPtr)
		intf.SetIntfPointer(2, viewPtr)
	}
	return ret
}

var (
	cefLabelButtonRefOnce   base.Once
	cefLabelButtonRefImport *imports.Imports = nil
)

func cefLabelButtonRefAPI() *imports.Imports {
	cefLabelButtonRefOnce.Do(func() {
		cefLabelButtonRefImport = api.NewDefaultImports()
		cefLabelButtonRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefLabelButtonRef_Create", 0), // constructor NewLabelButtonRef
			/* 1 */ imports.NewTable("TCefLabelButtonRef_AsMenuButton", 0), // function AsMenuButton
			/* 2 */ imports.NewTable("TCefLabelButtonRef_GetText", 0), // function GetText
			/* 3 */ imports.NewTable("TCefLabelButtonRef_GetImage", 0), // function GetImage
			/* 4 */ imports.NewTable("TCefLabelButtonRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 5 */ imports.NewTable("TCefLabelButtonRef_CreateLabelButton", 0), // static function CreateLabelButton
			/* 6 */ imports.NewTable("TCefLabelButtonRef_SetText", 0), // procedure SetText
			/* 7 */ imports.NewTable("TCefLabelButtonRef_SetImage", 0), // procedure SetImage
			/* 8 */ imports.NewTable("TCefLabelButtonRef_SetTextColor", 0), // procedure SetTextColor
			/* 9 */ imports.NewTable("TCefLabelButtonRef_SetEnabledTextColors", 0), // procedure SetEnabledTextColors
			/* 10 */ imports.NewTable("TCefLabelButtonRef_SetFontList", 0), // procedure SetFontList
			/* 11 */ imports.NewTable("TCefLabelButtonRef_SetHorizontalAlignment", 0), // procedure SetHorizontalAlignment
			/* 12 */ imports.NewTable("TCefLabelButtonRef_SetMinimumSize", 0), // procedure SetMinimumSize
			/* 13 */ imports.NewTable("TCefLabelButtonRef_SetMaximumSize", 0), // procedure SetMaximumSize
		}
	})
	return cefLabelButtonRefImport
}
