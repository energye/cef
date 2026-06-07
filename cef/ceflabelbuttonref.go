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

// ICefLabelButton Parent: ICefButton
type ICefLabelButton interface {
	ICefButton
	AsMenuButton() ICefMenuButton                                             // function
	GetText() string                                                          // function
	GetImage(buttonState cefTypes.TCefButtonState) ICefImage                  // function
	SetText(text string)                                                      // procedure
	SetImage(buttonState cefTypes.TCefButtonState, image ICefImage)           // procedure
	SetTextColor(forState cefTypes.TCefButtonState, color cefTypes.TCefColor) // procedure
	SetEnabledTextColors(color cefTypes.TCefColor)                            // procedure
	SetFontList(fontList string)                                              // procedure
	SetHorizontalAlignment(alignment cefTypes.TCefHorizontalAlignment)        // procedure
	SetMinimumSize(size TCefSize)                                             // procedure
	SetMaximumSize(size TCefSize)                                             // procedure
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

func (_LabelButtonRefClass) UnWrapWithPointer(data uintptr) (result ICefLabelButton) {
	var resultPtr uintptr
	cefLabelButtonRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefLabelButtonRef(resultPtr)
	return
}

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
