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

// ICefMenuModel Parent: ICefBaseRefCounted
type ICefMenuModel interface {
	ICefBaseRefCounted
	IsSubMenu() bool                                                                                                // function
	Clear() bool                                                                                                    // function
	GetCount() NativeUInt                                                                                           // function
	AddSeparator() bool                                                                                             // function
	AddItem(commandId int32, text string) bool                                                                      // function
	AddCheckItem(commandId int32, text string) bool                                                                 // function
	AddRadioItem(commandId int32, text string, groupId int32) bool                                                  // function
	AddSubMenu(commandId int32, text string) ICefMenuModel                                                          // function
	InsertSeparatorAt(index NativeUInt) bool                                                                        // function
	InsertItemAt(index NativeUInt, commandId int32, text string) bool                                               // function
	InsertCheckItemAt(index NativeUInt, commandId int32, text string) bool                                          // function
	InsertRadioItemAt(index NativeUInt, commandId int32, text string, groupId int32) bool                           // function
	InsertSubMenuAt(index NativeUInt, commandId int32, text string) ICefMenuModel                                   // function
	Remove(commandId int32) bool                                                                                    // function
	RemoveAt(index NativeUInt) bool                                                                                 // function
	GetIndexOf(commandId int32) int32                                                                               // function
	GetCommandIdAt(index NativeUInt) int32                                                                          // function
	SetCommandIdAt(index NativeUInt, commandId int32) bool                                                          // function
	GetLabel(commandId int32) string                                                                                // function
	GetLabelAt(index NativeUInt) string                                                                             // function
	SetLabel(commandId int32, text string) bool                                                                     // function
	SetLabelAt(index NativeUInt, text string) bool                                                                  // function
	GetType(commandId int32) TCefMenuItemType                                                                       // function
	GetTypeAt(index NativeUInt) TCefMenuItemType                                                                    // function
	GetGroupId(commandId int32) int32                                                                               // function
	GetGroupIdAt(index NativeUInt) int32                                                                            // function
	SetGroupId(commandId, groupId int32) bool                                                                       // function
	SetGroupIdAt(index NativeUInt, groupId int32) bool                                                              // function
	GetSubMenu(commandId int32) ICefMenuModel                                                                       // function
	GetSubMenuAt(index NativeUInt) ICefMenuModel                                                                    // function
	IsVisible(commandId int32) bool                                                                                 // function
	IsVisibleAt(index NativeUInt) bool                                                                              // function
	SetVisible(commandId int32, visible bool) bool                                                                  // function
	SetVisibleAt(index NativeUInt, visible bool) bool                                                               // function
	IsEnabled(commandId int32) bool                                                                                 // function
	IsEnabledAt(index NativeUInt) bool                                                                              // function
	SetEnabled(commandId int32, enabled bool) bool                                                                  // function
	SetEnabledAt(index NativeUInt, enabled bool) bool                                                               // function
	IsChecked(commandId int32) bool                                                                                 // function
	IsCheckedAt(index NativeUInt) bool                                                                              // function
	SetChecked(commandId int32, checked bool) bool                                                                  // function
	SetCheckedAt(index NativeUInt, checked bool) bool                                                               // function
	HasAccelerator(commandId int32) bool                                                                            // function
	HasAcceleratorAt(index NativeUInt) bool                                                                         // function
	SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool                       // function
	SetAcceleratorAt(index NativeUInt, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool              // function
	RemoveAccelerator(commandId int32) bool                                                                         // function
	RemoveAcceleratorAt(index NativeUInt) bool                                                                      // function
	GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool    // function
	GetAcceleratorAt(index NativeUInt, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool // function
	SetColor(commandId int32, colorType TCefMenuColorType, color TCefColor) bool                                    // function
	SetColorAt(index int32, colorType TCefMenuColorType, color TCefColor) bool                                      // function
	GetColor(commandId int32, colorType TCefMenuColorType, outColor *TCefColor) bool                                // function
	GetColorAt(index int32, colorType TCefMenuColorType, outColor *TCefColor) bool                                  // function
	SetFontList(commandId int32, fontList string) bool                                                              // function
	SetFontListAt(index int32, fontList string) bool                                                                // function
}

// TCefMenuModel Parent: TCefBaseRefCounted
type TCefMenuModel struct {
	TCefBaseRefCounted
}

// MenuModelRef -> ICefMenuModel
var MenuModelRef menuModel

// menuModel TCefMenuModel Ref
type menuModel uintptr

func (m *menuModel) UnWrap(data uintptr) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(57, uintptr(data), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *menuModel) New(delegate ICefMenuModelDelegate) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(35, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) IsSubMenu() bool {
	r1 := menuModelImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) Clear() bool {
	r1 := menuModelImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) GetCount() NativeUInt {
	r1 := menuModelImportAPI().SysCallN(11, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefMenuModel) AddSeparator() bool {
	r1 := menuModelImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) AddItem(commandId int32, text string) bool {
	r1 := menuModelImportAPI().SysCallN(1, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddCheckItem(commandId int32, text string) bool {
	r1 := menuModelImportAPI().SysCallN(0, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddRadioItem(commandId int32, text string, groupId int32) bool {
	r1 := menuModelImportAPI().SysCallN(2, m.Instance(), uintptr(commandId), PascalStr(text), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddSubMenu(commandId int32, text string) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(4, m.Instance(), uintptr(commandId), PascalStr(text), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) InsertSeparatorAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(26, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertItemAt(index NativeUInt, commandId int32, text string) bool {
	r1 := menuModelImportAPI().SysCallN(24, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertCheckItemAt(index NativeUInt, commandId int32, text string) bool {
	r1 := menuModelImportAPI().SysCallN(23, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertRadioItemAt(index NativeUInt, commandId int32, text string, groupId int32) bool {
	r1 := menuModelImportAPI().SysCallN(25, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertSubMenuAt(index NativeUInt, commandId int32, text string) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(27, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) Remove(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(36, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(39, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetIndexOf(commandId int32) int32 {
	r1 := menuModelImportAPI().SysCallN(14, m.Instance(), uintptr(commandId))
	return int32(r1)
}

func (m *TCefMenuModel) GetCommandIdAt(index NativeUInt) int32 {
	r1 := menuModelImportAPI().SysCallN(10, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefMenuModel) SetCommandIdAt(index NativeUInt, commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(46, m.Instance(), uintptr(index), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetLabel(commandId int32) string {
	r1 := menuModelImportAPI().SysCallN(15, m.Instance(), uintptr(commandId))
	return GoStr(r1)
}

func (m *TCefMenuModel) GetLabelAt(index NativeUInt) string {
	r1 := menuModelImportAPI().SysCallN(16, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefMenuModel) SetLabel(commandId int32, text string) bool {
	r1 := menuModelImportAPI().SysCallN(53, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetLabelAt(index NativeUInt, text string) bool {
	r1 := menuModelImportAPI().SysCallN(54, m.Instance(), uintptr(index), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetType(commandId int32) TCefMenuItemType {
	r1 := menuModelImportAPI().SysCallN(19, m.Instance(), uintptr(commandId))
	return TCefMenuItemType(r1)
}

func (m *TCefMenuModel) GetTypeAt(index NativeUInt) TCefMenuItemType {
	r1 := menuModelImportAPI().SysCallN(20, m.Instance(), uintptr(index))
	return TCefMenuItemType(r1)
}

func (m *TCefMenuModel) GetGroupId(commandId int32) int32 {
	r1 := menuModelImportAPI().SysCallN(12, m.Instance(), uintptr(commandId))
	return int32(r1)
}

func (m *TCefMenuModel) GetGroupIdAt(index NativeUInt) int32 {
	r1 := menuModelImportAPI().SysCallN(13, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefMenuModel) SetGroupId(commandId, groupId int32) bool {
	r1 := menuModelImportAPI().SysCallN(51, m.Instance(), uintptr(commandId), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetGroupIdAt(index NativeUInt, groupId int32) bool {
	r1 := menuModelImportAPI().SysCallN(52, m.Instance(), uintptr(index), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetSubMenu(commandId int32) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(17, m.Instance(), uintptr(commandId), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) GetSubMenuAt(index NativeUInt) ICefMenuModel {
	var resultCefMenuModel uintptr
	menuModelImportAPI().SysCallN(18, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) IsVisible(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(33, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsVisibleAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(34, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetVisible(commandId int32, visible bool) bool {
	r1 := menuModelImportAPI().SysCallN(55, m.Instance(), uintptr(commandId), PascalBool(visible))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetVisibleAt(index NativeUInt, visible bool) bool {
	r1 := menuModelImportAPI().SysCallN(56, m.Instance(), uintptr(index), PascalBool(visible))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsEnabled(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(30, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsEnabledAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(31, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetEnabled(commandId int32, enabled bool) bool {
	r1 := menuModelImportAPI().SysCallN(47, m.Instance(), uintptr(commandId), PascalBool(enabled))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetEnabledAt(index NativeUInt, enabled bool) bool {
	r1 := menuModelImportAPI().SysCallN(48, m.Instance(), uintptr(index), PascalBool(enabled))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsChecked(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(28, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsCheckedAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(29, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetChecked(commandId int32, checked bool) bool {
	r1 := menuModelImportAPI().SysCallN(42, m.Instance(), uintptr(commandId), PascalBool(checked))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetCheckedAt(index NativeUInt, checked bool) bool {
	r1 := menuModelImportAPI().SysCallN(43, m.Instance(), uintptr(index), PascalBool(checked))
	return GoBool(r1)
}

func (m *TCefMenuModel) HasAccelerator(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(21, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) HasAcceleratorAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(22, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1 := menuModelImportAPI().SysCallN(40, m.Instance(), uintptr(commandId), uintptr(keyCode), PascalBool(shiftPressed), PascalBool(ctrlPressed), PascalBool(altPressed))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetAcceleratorAt(index NativeUInt, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1 := menuModelImportAPI().SysCallN(41, m.Instance(), uintptr(index), uintptr(keyCode), PascalBool(shiftPressed), PascalBool(ctrlPressed), PascalBool(altPressed))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAccelerator(commandId int32) bool {
	r1 := menuModelImportAPI().SysCallN(37, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAcceleratorAt(index NativeUInt) bool {
	r1 := menuModelImportAPI().SysCallN(38, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := menuModelImportAPI().SysCallN(6, m.Instance(), uintptr(commandId), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*outKeyCode = int32(result1)
	*outShiftPressed = GoBool(result2)
	*outCtrlPressed = GoBool(result3)
	*outLtPressed = GoBool(result4)
	return GoBool(r1)
}

func (m *TCefMenuModel) GetAcceleratorAt(index NativeUInt, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := menuModelImportAPI().SysCallN(7, m.Instance(), uintptr(index), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*outKeyCode = int32(result1)
	*outShiftPressed = GoBool(result2)
	*outCtrlPressed = GoBool(result3)
	*outLtPressed = GoBool(result4)
	return GoBool(r1)
}

func (m *TCefMenuModel) SetColor(commandId int32, colorType TCefMenuColorType, color TCefColor) bool {
	r1 := menuModelImportAPI().SysCallN(44, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(color))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetColorAt(index int32, colorType TCefMenuColorType, color TCefColor) bool {
	r1 := menuModelImportAPI().SysCallN(45, m.Instance(), uintptr(index), uintptr(colorType), uintptr(color))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetColor(commandId int32, colorType TCefMenuColorType, outColor *TCefColor) bool {
	var result2 uintptr
	r1 := menuModelImportAPI().SysCallN(8, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(unsafePointer(&result2)))
	*outColor = TCefColor(result2)
	return GoBool(r1)
}

func (m *TCefMenuModel) GetColorAt(index int32, colorType TCefMenuColorType, outColor *TCefColor) bool {
	var result2 uintptr
	r1 := menuModelImportAPI().SysCallN(9, m.Instance(), uintptr(index), uintptr(colorType), uintptr(unsafePointer(&result2)))
	*outColor = TCefColor(result2)
	return GoBool(r1)
}

func (m *TCefMenuModel) SetFontList(commandId int32, fontList string) bool {
	r1 := menuModelImportAPI().SysCallN(49, m.Instance(), uintptr(commandId), PascalStr(fontList))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetFontListAt(index int32, fontList string) bool {
	r1 := menuModelImportAPI().SysCallN(50, m.Instance(), uintptr(index), PascalStr(fontList))
	return GoBool(r1)
}

var (
	menuModelImport       *imports.Imports = nil
	menuModelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMenuModel_AddCheckItem", 0),
		/*1*/ imports.NewTable("CefMenuModel_AddItem", 0),
		/*2*/ imports.NewTable("CefMenuModel_AddRadioItem", 0),
		/*3*/ imports.NewTable("CefMenuModel_AddSeparator", 0),
		/*4*/ imports.NewTable("CefMenuModel_AddSubMenu", 0),
		/*5*/ imports.NewTable("CefMenuModel_Clear", 0),
		/*6*/ imports.NewTable("CefMenuModel_GetAccelerator", 0),
		/*7*/ imports.NewTable("CefMenuModel_GetAcceleratorAt", 0),
		/*8*/ imports.NewTable("CefMenuModel_GetColor", 0),
		/*9*/ imports.NewTable("CefMenuModel_GetColorAt", 0),
		/*10*/ imports.NewTable("CefMenuModel_GetCommandIdAt", 0),
		/*11*/ imports.NewTable("CefMenuModel_GetCount", 0),
		/*12*/ imports.NewTable("CefMenuModel_GetGroupId", 0),
		/*13*/ imports.NewTable("CefMenuModel_GetGroupIdAt", 0),
		/*14*/ imports.NewTable("CefMenuModel_GetIndexOf", 0),
		/*15*/ imports.NewTable("CefMenuModel_GetLabel", 0),
		/*16*/ imports.NewTable("CefMenuModel_GetLabelAt", 0),
		/*17*/ imports.NewTable("CefMenuModel_GetSubMenu", 0),
		/*18*/ imports.NewTable("CefMenuModel_GetSubMenuAt", 0),
		/*19*/ imports.NewTable("CefMenuModel_GetType", 0),
		/*20*/ imports.NewTable("CefMenuModel_GetTypeAt", 0),
		/*21*/ imports.NewTable("CefMenuModel_HasAccelerator", 0),
		/*22*/ imports.NewTable("CefMenuModel_HasAcceleratorAt", 0),
		/*23*/ imports.NewTable("CefMenuModel_InsertCheckItemAt", 0),
		/*24*/ imports.NewTable("CefMenuModel_InsertItemAt", 0),
		/*25*/ imports.NewTable("CefMenuModel_InsertRadioItemAt", 0),
		/*26*/ imports.NewTable("CefMenuModel_InsertSeparatorAt", 0),
		/*27*/ imports.NewTable("CefMenuModel_InsertSubMenuAt", 0),
		/*28*/ imports.NewTable("CefMenuModel_IsChecked", 0),
		/*29*/ imports.NewTable("CefMenuModel_IsCheckedAt", 0),
		/*30*/ imports.NewTable("CefMenuModel_IsEnabled", 0),
		/*31*/ imports.NewTable("CefMenuModel_IsEnabledAt", 0),
		/*32*/ imports.NewTable("CefMenuModel_IsSubMenu", 0),
		/*33*/ imports.NewTable("CefMenuModel_IsVisible", 0),
		/*34*/ imports.NewTable("CefMenuModel_IsVisibleAt", 0),
		/*35*/ imports.NewTable("CefMenuModel_New", 0),
		/*36*/ imports.NewTable("CefMenuModel_Remove", 0),
		/*37*/ imports.NewTable("CefMenuModel_RemoveAccelerator", 0),
		/*38*/ imports.NewTable("CefMenuModel_RemoveAcceleratorAt", 0),
		/*39*/ imports.NewTable("CefMenuModel_RemoveAt", 0),
		/*40*/ imports.NewTable("CefMenuModel_SetAccelerator", 0),
		/*41*/ imports.NewTable("CefMenuModel_SetAcceleratorAt", 0),
		/*42*/ imports.NewTable("CefMenuModel_SetChecked", 0),
		/*43*/ imports.NewTable("CefMenuModel_SetCheckedAt", 0),
		/*44*/ imports.NewTable("CefMenuModel_SetColor", 0),
		/*45*/ imports.NewTable("CefMenuModel_SetColorAt", 0),
		/*46*/ imports.NewTable("CefMenuModel_SetCommandIdAt", 0),
		/*47*/ imports.NewTable("CefMenuModel_SetEnabled", 0),
		/*48*/ imports.NewTable("CefMenuModel_SetEnabledAt", 0),
		/*49*/ imports.NewTable("CefMenuModel_SetFontList", 0),
		/*50*/ imports.NewTable("CefMenuModel_SetFontListAt", 0),
		/*51*/ imports.NewTable("CefMenuModel_SetGroupId", 0),
		/*52*/ imports.NewTable("CefMenuModel_SetGroupIdAt", 0),
		/*53*/ imports.NewTable("CefMenuModel_SetLabel", 0),
		/*54*/ imports.NewTable("CefMenuModel_SetLabelAt", 0),
		/*55*/ imports.NewTable("CefMenuModel_SetVisible", 0),
		/*56*/ imports.NewTable("CefMenuModel_SetVisibleAt", 0),
		/*57*/ imports.NewTable("CefMenuModel_UnWrap", 0),
	}
)

func menuModelImportAPI() *imports.Imports {
	if menuModelImport == nil {
		menuModelImport = NewDefaultImports()
		menuModelImport.SetImportTable(menuModelImportTables)
		menuModelImportTables = nil
	}
	return menuModelImport
}
