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

// ICefMenuModel Parent: ICefBaseRefCountedRef
type ICefMenuModel interface {
	ICefBaseRefCountedRef
	// IsSubMenu
	//  Returns true (1) if this menu is a submenu.
	IsSubMenu() bool // function
	// Clear
	//  Clears the menu. Returns true (1) on success.
	Clear() bool // function
	// GetCount
	//  Returns the number of items in this menu.
	GetCount() cefTypes.NativeUInt // function
	// AddSeparator
	//  Add a separator to the menu. Returns true (1) on success.
	AddSeparator() bool // function
	// AddItem
	//  Add an item to the menu. Returns true (1) on success.
	AddItem(commandId int32, text string) bool // function
	// AddCheckItem
	//  Add a check item to the menu. Returns true (1) on success.
	AddCheckItem(commandId int32, text string) bool // function
	// AddRadioItem
	//  Add a radio item to the menu. Only a single item with the specified
	//  |group_id| can be checked at a time. Returns true (1) on success.
	AddRadioItem(commandId int32, text string, groupId int32) bool // function
	// AddSubMenu
	//  Add a sub-menu to the menu. The new sub-menu is returned.
	AddSubMenu(commandId int32, text string) ICefMenuModel // function
	// InsertSeparatorAt
	//  Insert a separator in the menu at the specified |index|. Returns true (1)
	//  on success.
	InsertSeparatorAt(index cefTypes.NativeUInt) bool // function
	// InsertItemAt
	//  Insert an item in the menu at the specified |index|. Returns true (1) on
	//  success.
	InsertItemAt(index cefTypes.NativeUInt, commandId int32, text string) bool // function
	// InsertCheckItemAt
	//  Insert a check item in the menu at the specified |index|. Returns true (1)
	//  on success.
	InsertCheckItemAt(index cefTypes.NativeUInt, commandId int32, text string) bool // function
	// InsertRadioItemAt
	//  Insert a radio item in the menu at the specified |index|. Only a single
	//  item with the specified |group_id| can be checked at a time. Returns true
	//  (1) on success.
	InsertRadioItemAt(index cefTypes.NativeUInt, commandId int32, text string, groupId int32) bool // function
	// InsertSubMenuAt
	//  Insert a sub-menu in the menu at the specified |index|. The new sub-menu
	//  is returned.
	InsertSubMenuAt(index cefTypes.NativeUInt, commandId int32, text string) ICefMenuModel // function
	// Remove
	//  Removes the item with the specified |command_id|. Returns true (1) on
	//  success.
	Remove(commandId int32) bool // function
	// RemoveAt
	//  Removes the item at the specified |index|. Returns true (1) on success.
	RemoveAt(index cefTypes.NativeUInt) bool // function
	// GetIndexOf
	//  Returns the index associated with the specified |command_id| or -1 if not
	//  found due to the command id not existing in the menu.
	GetIndexOf(commandId int32) int32 // function
	// GetCommandIdAt
	//  Returns the command id at the specified |index| or -1 if not found due to
	//  invalid range or the index being a separator.
	GetCommandIdAt(index cefTypes.NativeUInt) int32 // function
	// SetCommandIdAt
	//  Sets the command id at the specified |index|. Returns true (1) on success.
	SetCommandIdAt(index cefTypes.NativeUInt, commandId int32) bool // function
	// GetLabel
	//  Returns the label for the specified |command_id| or NULL if not found.
	GetLabel(commandId int32) string // function
	// GetLabelAt
	//  Returns the label at the specified |index| or NULL if not found due to
	//  invalid range or the index being a separator.
	GetLabelAt(index cefTypes.NativeUInt) string // function
	// SetLabel
	//  Sets the label for the specified |command_id|. Returns true (1) on
	//  success.
	SetLabel(commandId int32, text string) bool // function
	// SetLabelAt
	//  Set the label at the specified |index|. Returns true (1) on success.
	SetLabelAt(index cefTypes.NativeUInt, text string) bool // function
	// GetType
	//  Returns the item type for the specified |command_id|.
	GetType(commandId int32) cefTypes.TCefMenuItemType // function
	// GetTypeAt
	//  Returns the item type at the specified |index|.
	GetTypeAt(index cefTypes.NativeUInt) cefTypes.TCefMenuItemType // function
	// GetGroupId
	//  Returns the group id for the specified |command_id| or -1 if invalid.
	GetGroupId(commandId int32) int32 // function
	// GetGroupIdAt
	//  Returns the group id at the specified |index| or -1 if invalid.
	GetGroupIdAt(index cefTypes.NativeUInt) int32 // function
	// SetGroupId
	//  Sets the group id for the specified |command_id|. Returns true (1) on
	//  success.
	SetGroupId(commandId int32, groupId int32) bool // function
	// SetGroupIdAt
	//  Sets the group id at the specified |index|. Returns true (1) on success.
	SetGroupIdAt(index cefTypes.NativeUInt, groupId int32) bool // function
	// GetSubMenu
	//  Returns the submenu for the specified |command_id| or NULL if invalid.
	GetSubMenu(commandId int32) ICefMenuModel // function
	// GetSubMenuAt
	//  Returns the submenu at the specified |index| or NULL if invalid.
	GetSubMenuAt(index cefTypes.NativeUInt) ICefMenuModel // function
	// IsVisible
	//  Returns true (1) if the specified |command_id| is visible.
	IsVisible(commandId int32) bool // function
	// IsVisibleAt
	//  Returns true (1) if the specified |index| is visible.
	IsVisibleAt(index cefTypes.NativeUInt) bool // function
	// SetVisible
	//  Change the visibility of the specified |command_id|. Returns true (1) on
	//  success.
	SetVisible(commandId int32, visible bool) bool // function
	// SetVisibleAt
	//  Change the visibility at the specified |index|. Returns true (1) on
	//  success.
	SetVisibleAt(index cefTypes.NativeUInt, visible bool) bool // function
	// IsEnabled
	//  Returns true (1) if the specified |command_id| is enabled.
	IsEnabled(commandId int32) bool // function
	// IsEnabledAt
	//  Returns true (1) if the specified |index| is enabled.
	IsEnabledAt(index cefTypes.NativeUInt) bool // function
	// SetEnabled
	//  Change the enabled status of the specified |command_id|. Returns true (1)
	//  on success.
	SetEnabled(commandId int32, enabled bool) bool // function
	// SetEnabledAt
	//  Change the enabled status at the specified |index|. Returns true (1) on
	//  success.
	SetEnabledAt(index cefTypes.NativeUInt, enabled bool) bool // function
	// IsChecked
	//  Returns true (1) if the specified |command_id| is checked. Only applies to
	//  check and radio items.
	IsChecked(commandId int32) bool // function
	// IsCheckedAt
	//  Returns true (1) if the specified |index| is checked. Only applies to
	//  check and radio items.
	IsCheckedAt(index cefTypes.NativeUInt) bool // function
	// SetChecked
	//  Check the specified |command_id|. Only applies to check and radio items.
	//  Returns true (1) on success.
	SetChecked(commandId int32, checked bool) bool // function
	// SetCheckedAt
	//  Check the specified |index|. Only applies to check and radio items.
	//  Returns true (1) on success.
	SetCheckedAt(index cefTypes.NativeUInt, checked bool) bool // function
	// HasAccelerator
	//  Returns true (1) if the specified |command_id| has a keyboard accelerator
	//  assigned.
	HasAccelerator(commandId int32) bool // function
	// HasAcceleratorAt
	//  Returns true (1) if the specified |index| has a keyboard accelerator
	//  assigned.
	HasAcceleratorAt(index cefTypes.NativeUInt) bool // function
	// SetAccelerator
	//  Set the keyboard accelerator for the specified |command_id|. |key_code|
	//  can be any virtual key or character value. Returns true (1) on success.
	SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) bool // function
	// SetAcceleratorAt
	//  Set the keyboard accelerator at the specified |index|. |key_code| can be
	//  any virtual key or character value. Returns true (1) on success.
	SetAcceleratorAt(index cefTypes.NativeUInt, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) bool // function
	// RemoveAccelerator
	//  Remove the keyboard accelerator for the specified |command_id|. Returns
	//  true (1) on success.
	RemoveAccelerator(commandId int32) bool // function
	// RemoveAcceleratorAt
	//  Remove the keyboard accelerator at the specified |index|. Returns true (1)
	//  on success.
	RemoveAcceleratorAt(index cefTypes.NativeUInt) bool // function
	// GetAccelerator
	//  Retrieves the keyboard accelerator for the specified |command_id|. Returns
	//  true (1) on success.
	GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed *bool, outCtrlPressed *bool, outAltPressed *bool) bool // function
	// GetAcceleratorAt
	//  Retrieves the keyboard accelerator for the specified |index|. Returns true
	//  (1) on success.
	GetAcceleratorAt(index cefTypes.NativeUInt, outKeyCode *int32, outShiftPressed *bool, outCtrlPressed *bool, outAltPressed *bool) bool // function
	// SetColor
	//  Set the explicit color for |command_id| and |color_type| to |color|.
	//  Specify a |color| value of 0 to remove the explicit color. If no explicit
	//  color or default color is set for |color_type| then the system color will
	//  be used. Returns true (1) on success.
	SetColor(commandId int32, colorType cefTypes.TCefMenuColorType, color cefTypes.TCefColor) bool // function
	// SetColorAt
	//  Set the explicit color for |command_id| and |index| to |color|. Specify a
	//  |color| value of 0 to remove the explicit color. Specify an |index| value
	//  of -1 to set the default color for items that do not have an explicit
	//  color set. If no explicit color or default color is set for |color_type|
	//  then the system color will be used. Returns true (1) on success.
	SetColorAt(index int32, colorType cefTypes.TCefMenuColorType, color cefTypes.TCefColor) bool // function
	// GetColor
	//  Returns in |color| the color that was explicitly set for |command_id| and
	//  |color_type|. If a color was not set then 0 will be returned in |color|.
	//  Returns true (1) on success.
	GetColor(commandId int32, colorType cefTypes.TCefMenuColorType, outColor *cefTypes.TCefColor) bool // function
	// GetColorAt
	//  Returns in |color| the color that was explicitly set for |command_id| and
	//  |color_type|. Specify an |index| value of -1 to return the default color
	//  in |color|. If a color was not set then 0 will be returned in |color|.
	//  Returns true (1) on success.
	GetColorAt(index int32, colorType cefTypes.TCefMenuColorType, outColor *cefTypes.TCefColor) bool // function
	// SetFontList
	//  Sets the font list for the specified |command_id|. If |font_list| is NULL
	//  the system font will be used. Returns true (1) on success. The format is
	//  "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where:
	//  - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names
	//  (case-sensitive "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//
	//  Here are examples of valid font description strings:
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
	SetFontList(commandId int32, fontList string) bool // function
	// SetFontListAt
	//  Sets the font list for the specified |index|. Specify an |index| value of
	//  -1 to set the default font. If |font_list| is NULL the system font will be
	//  used. Returns true (1) on success. The format is
	//  "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where:
	//  - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names
	//  (case-sensitive "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//
	//  Here are examples of valid font description strings:
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
	SetFontListAt(index int32, fontList string) bool // function
}

// ICefMenuModelRef Parent: ICefMenuModel
type ICefMenuModelRef interface {
	ICefMenuModel
	AsIntfMenuModel() uintptr
}

type TCefMenuModelRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefMenuModelRef) IsSubMenu() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) Clear() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(3, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefMenuModelRef) AddSeparator() bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) AddItem(commandId int32, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(5, m.Instance(), uintptr(commandId), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) AddCheckItem(commandId int32, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(6, m.Instance(), uintptr(commandId), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) AddRadioItem(commandId int32, text string, groupId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(7, m.Instance(), uintptr(commandId), api.PasStr(text), uintptr(groupId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) AddSubMenu(commandId int32, text string) (result ICefMenuModel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(8, m.Instance(), uintptr(commandId), api.PasStr(text), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

func (m *TCefMenuModelRef) InsertSeparatorAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(9, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) InsertItemAt(index cefTypes.NativeUInt, commandId int32, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(10, m.Instance(), uintptr(index), uintptr(commandId), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) InsertCheckItemAt(index cefTypes.NativeUInt, commandId int32, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(11, m.Instance(), uintptr(index), uintptr(commandId), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) InsertRadioItemAt(index cefTypes.NativeUInt, commandId int32, text string, groupId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(12, m.Instance(), uintptr(index), uintptr(commandId), api.PasStr(text), uintptr(groupId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) InsertSubMenuAt(index cefTypes.NativeUInt, commandId int32, text string) (result ICefMenuModel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(13, m.Instance(), uintptr(index), uintptr(commandId), api.PasStr(text), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

func (m *TCefMenuModelRef) Remove(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(14, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) RemoveAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(15, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetIndexOf(commandId int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(16, m.Instance(), uintptr(commandId))
	return int32(r)
}

func (m *TCefMenuModelRef) GetCommandIdAt(index cefTypes.NativeUInt) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(17, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TCefMenuModelRef) SetCommandIdAt(index cefTypes.NativeUInt, commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(18, m.Instance(), uintptr(index), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetLabel(commandId int32) string {
	if !m.IsValid() {
		return ""
	}
	r := cefMenuModelRefAPI().SysCallN(19, m.Instance(), uintptr(commandId))
	return api.GoStr(r)
}

func (m *TCefMenuModelRef) GetLabelAt(index cefTypes.NativeUInt) string {
	if !m.IsValid() {
		return ""
	}
	r := cefMenuModelRefAPI().SysCallN(20, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TCefMenuModelRef) SetLabel(commandId int32, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(21, m.Instance(), uintptr(commandId), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetLabelAt(index cefTypes.NativeUInt, text string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(22, m.Instance(), uintptr(index), api.PasStr(text))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetType(commandId int32) cefTypes.TCefMenuItemType {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(23, m.Instance(), uintptr(commandId))
	return cefTypes.TCefMenuItemType(r)
}

func (m *TCefMenuModelRef) GetTypeAt(index cefTypes.NativeUInt) cefTypes.TCefMenuItemType {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(24, m.Instance(), uintptr(index))
	return cefTypes.TCefMenuItemType(r)
}

func (m *TCefMenuModelRef) GetGroupId(commandId int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(25, m.Instance(), uintptr(commandId))
	return int32(r)
}

func (m *TCefMenuModelRef) GetGroupIdAt(index cefTypes.NativeUInt) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefMenuModelRefAPI().SysCallN(26, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TCefMenuModelRef) SetGroupId(commandId int32, groupId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(27, m.Instance(), uintptr(commandId), uintptr(groupId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetGroupIdAt(index cefTypes.NativeUInt, groupId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(28, m.Instance(), uintptr(index), uintptr(groupId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetSubMenu(commandId int32) (result ICefMenuModel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(29, m.Instance(), uintptr(commandId), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

func (m *TCefMenuModelRef) GetSubMenuAt(index cefTypes.NativeUInt) (result ICefMenuModel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(30, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

func (m *TCefMenuModelRef) IsVisible(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(31, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) IsVisibleAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(32, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetVisible(commandId int32, visible bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(33, m.Instance(), uintptr(commandId), api.PasBool(visible))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetVisibleAt(index cefTypes.NativeUInt, visible bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(34, m.Instance(), uintptr(index), api.PasBool(visible))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) IsEnabled(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(35, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) IsEnabledAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(36, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetEnabled(commandId int32, enabled bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(37, m.Instance(), uintptr(commandId), api.PasBool(enabled))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetEnabledAt(index cefTypes.NativeUInt, enabled bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(38, m.Instance(), uintptr(index), api.PasBool(enabled))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) IsChecked(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(39, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) IsCheckedAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(40, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetChecked(commandId int32, checked bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(41, m.Instance(), uintptr(commandId), api.PasBool(checked))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetCheckedAt(index cefTypes.NativeUInt, checked bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(42, m.Instance(), uintptr(index), api.PasBool(checked))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) HasAccelerator(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(43, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) HasAcceleratorAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(44, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(45, m.Instance(), uintptr(commandId), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetAcceleratorAt(index cefTypes.NativeUInt, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(46, m.Instance(), uintptr(index), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) RemoveAccelerator(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(47, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) RemoveAcceleratorAt(index cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(48, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed *bool, outCtrlPressed *bool, outAltPressed *bool) bool {
	if !m.IsValid() {
		return false
	}
	var keyCodePtr uintptr
	r := cefMenuModelRefAPI().SysCallN(49, m.Instance(), uintptr(commandId), uintptr(base.UnsafePointer(&keyCodePtr)), uintptr(base.UnsafePointer(outShiftPressed)), uintptr(base.UnsafePointer(outCtrlPressed)), uintptr(base.UnsafePointer(outAltPressed)))
	*outKeyCode = int32(keyCodePtr)
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetAcceleratorAt(index cefTypes.NativeUInt, outKeyCode *int32, outShiftPressed *bool, outCtrlPressed *bool, outAltPressed *bool) bool {
	if !m.IsValid() {
		return false
	}
	var keyCodePtr uintptr
	r := cefMenuModelRefAPI().SysCallN(50, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&keyCodePtr)), uintptr(base.UnsafePointer(outShiftPressed)), uintptr(base.UnsafePointer(outCtrlPressed)), uintptr(base.UnsafePointer(outAltPressed)))
	*outKeyCode = int32(keyCodePtr)
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetColor(commandId int32, colorType cefTypes.TCefMenuColorType, color cefTypes.TCefColor) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(51, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(color))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetColorAt(index int32, colorType cefTypes.TCefMenuColorType, color cefTypes.TCefColor) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(52, m.Instance(), uintptr(index), uintptr(colorType), uintptr(color))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetColor(commandId int32, colorType cefTypes.TCefMenuColorType, outColor *cefTypes.TCefColor) bool {
	if !m.IsValid() {
		return false
	}
	var colorPtr uintptr
	r := cefMenuModelRefAPI().SysCallN(53, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(base.UnsafePointer(&colorPtr)))
	*outColor = cefTypes.TCefColor(colorPtr)
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) GetColorAt(index int32, colorType cefTypes.TCefMenuColorType, outColor *cefTypes.TCefColor) bool {
	if !m.IsValid() {
		return false
	}
	var colorPtr uintptr
	r := cefMenuModelRefAPI().SysCallN(54, m.Instance(), uintptr(index), uintptr(colorType), uintptr(base.UnsafePointer(&colorPtr)))
	*outColor = cefTypes.TCefColor(colorPtr)
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetFontList(commandId int32, fontList string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(55, m.Instance(), uintptr(commandId), api.PasStr(fontList))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) SetFontListAt(index int32, fontList string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefMenuModelRefAPI().SysCallN(56, m.Instance(), uintptr(index), api.PasStr(fontList))
	return api.GoBool(r)
}

func (m *TCefMenuModelRef) AsIntfMenuModel() uintptr {
	return m.GetIntfPointer(0)
}

// MenuModelRef  is static instance
var MenuModelRef _MenuModelRefClass

// _MenuModelRefClass is class type defined by TCefMenuModelRef
type _MenuModelRefClass uintptr

func (_MenuModelRefClass) UnWrap(data uintptr) (result ICefMenuModel) {
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(57, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

func (_MenuModelRefClass) New(delegate IEngMenuModelDelegate) (result ICefMenuModel) {
	var resultPtr uintptr
	cefMenuModelRefAPI().SysCallN(58, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMenuModelRef(resultPtr)
	return
}

// NewMenuModelRef class constructor
func NewMenuModelRef(data uintptr) ICefMenuModelRef {
	var menuModelPtr uintptr // ICefMenuModel
	r := cefMenuModelRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&menuModelPtr)))
	ret := AsCefMenuModelRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, menuModelPtr)
	}
	return ret
}

var (
	cefMenuModelRefOnce   base.Once
	cefMenuModelRefImport *imports.Imports = nil
)

func cefMenuModelRefAPI() *imports.Imports {
	cefMenuModelRefOnce.Do(func() {
		cefMenuModelRefImport = api.NewDefaultImports()
		cefMenuModelRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefMenuModelRef_Create", 0), // constructor NewMenuModelRef
			/* 1 */ imports.NewTable("TCefMenuModelRef_IsSubMenu", 0), // function IsSubMenu
			/* 2 */ imports.NewTable("TCefMenuModelRef_Clear", 0), // function Clear
			/* 3 */ imports.NewTable("TCefMenuModelRef_GetCount", 0), // function GetCount
			/* 4 */ imports.NewTable("TCefMenuModelRef_AddSeparator", 0), // function AddSeparator
			/* 5 */ imports.NewTable("TCefMenuModelRef_AddItem", 0), // function AddItem
			/* 6 */ imports.NewTable("TCefMenuModelRef_AddCheckItem", 0), // function AddCheckItem
			/* 7 */ imports.NewTable("TCefMenuModelRef_AddRadioItem", 0), // function AddRadioItem
			/* 8 */ imports.NewTable("TCefMenuModelRef_AddSubMenu", 0), // function AddSubMenu
			/* 9 */ imports.NewTable("TCefMenuModelRef_InsertSeparatorAt", 0), // function InsertSeparatorAt
			/* 10 */ imports.NewTable("TCefMenuModelRef_InsertItemAt", 0), // function InsertItemAt
			/* 11 */ imports.NewTable("TCefMenuModelRef_InsertCheckItemAt", 0), // function InsertCheckItemAt
			/* 12 */ imports.NewTable("TCefMenuModelRef_InsertRadioItemAt", 0), // function InsertRadioItemAt
			/* 13 */ imports.NewTable("TCefMenuModelRef_InsertSubMenuAt", 0), // function InsertSubMenuAt
			/* 14 */ imports.NewTable("TCefMenuModelRef_Remove", 0), // function Remove
			/* 15 */ imports.NewTable("TCefMenuModelRef_RemoveAt", 0), // function RemoveAt
			/* 16 */ imports.NewTable("TCefMenuModelRef_GetIndexOf", 0), // function GetIndexOf
			/* 17 */ imports.NewTable("TCefMenuModelRef_GetCommandIdAt", 0), // function GetCommandIdAt
			/* 18 */ imports.NewTable("TCefMenuModelRef_SetCommandIdAt", 0), // function SetCommandIdAt
			/* 19 */ imports.NewTable("TCefMenuModelRef_GetLabel", 0), // function GetLabel
			/* 20 */ imports.NewTable("TCefMenuModelRef_GetLabelAt", 0), // function GetLabelAt
			/* 21 */ imports.NewTable("TCefMenuModelRef_SetLabel", 0), // function SetLabel
			/* 22 */ imports.NewTable("TCefMenuModelRef_SetLabelAt", 0), // function SetLabelAt
			/* 23 */ imports.NewTable("TCefMenuModelRef_GetType", 0), // function GetType
			/* 24 */ imports.NewTable("TCefMenuModelRef_GetTypeAt", 0), // function GetTypeAt
			/* 25 */ imports.NewTable("TCefMenuModelRef_GetGroupId", 0), // function GetGroupId
			/* 26 */ imports.NewTable("TCefMenuModelRef_GetGroupIdAt", 0), // function GetGroupIdAt
			/* 27 */ imports.NewTable("TCefMenuModelRef_SetGroupId", 0), // function SetGroupId
			/* 28 */ imports.NewTable("TCefMenuModelRef_SetGroupIdAt", 0), // function SetGroupIdAt
			/* 29 */ imports.NewTable("TCefMenuModelRef_GetSubMenu", 0), // function GetSubMenu
			/* 30 */ imports.NewTable("TCefMenuModelRef_GetSubMenuAt", 0), // function GetSubMenuAt
			/* 31 */ imports.NewTable("TCefMenuModelRef_IsVisible", 0), // function IsVisible
			/* 32 */ imports.NewTable("TCefMenuModelRef_isVisibleAt", 0), // function IsVisibleAt
			/* 33 */ imports.NewTable("TCefMenuModelRef_SetVisible", 0), // function SetVisible
			/* 34 */ imports.NewTable("TCefMenuModelRef_SetVisibleAt", 0), // function SetVisibleAt
			/* 35 */ imports.NewTable("TCefMenuModelRef_IsEnabled", 0), // function IsEnabled
			/* 36 */ imports.NewTable("TCefMenuModelRef_IsEnabledAt", 0), // function IsEnabledAt
			/* 37 */ imports.NewTable("TCefMenuModelRef_SetEnabled", 0), // function SetEnabled
			/* 38 */ imports.NewTable("TCefMenuModelRef_SetEnabledAt", 0), // function SetEnabledAt
			/* 39 */ imports.NewTable("TCefMenuModelRef_IsChecked", 0), // function IsChecked
			/* 40 */ imports.NewTable("TCefMenuModelRef_IsCheckedAt", 0), // function IsCheckedAt
			/* 41 */ imports.NewTable("TCefMenuModelRef_setChecked", 0), // function SetChecked
			/* 42 */ imports.NewTable("TCefMenuModelRef_setCheckedAt", 0), // function SetCheckedAt
			/* 43 */ imports.NewTable("TCefMenuModelRef_HasAccelerator", 0), // function HasAccelerator
			/* 44 */ imports.NewTable("TCefMenuModelRef_HasAcceleratorAt", 0), // function HasAcceleratorAt
			/* 45 */ imports.NewTable("TCefMenuModelRef_SetAccelerator", 0), // function SetAccelerator
			/* 46 */ imports.NewTable("TCefMenuModelRef_SetAcceleratorAt", 0), // function SetAcceleratorAt
			/* 47 */ imports.NewTable("TCefMenuModelRef_RemoveAccelerator", 0), // function RemoveAccelerator
			/* 48 */ imports.NewTable("TCefMenuModelRef_RemoveAcceleratorAt", 0), // function RemoveAcceleratorAt
			/* 49 */ imports.NewTable("TCefMenuModelRef_GetAccelerator", 0), // function GetAccelerator
			/* 50 */ imports.NewTable("TCefMenuModelRef_GetAcceleratorAt", 0), // function GetAcceleratorAt
			/* 51 */ imports.NewTable("TCefMenuModelRef_SetColor", 0), // function SetColor
			/* 52 */ imports.NewTable("TCefMenuModelRef_SetColorAt", 0), // function SetColorAt
			/* 53 */ imports.NewTable("TCefMenuModelRef_GetColor", 0), // function GetColor
			/* 54 */ imports.NewTable("TCefMenuModelRef_GetColorAt", 0), // function GetColorAt
			/* 55 */ imports.NewTable("TCefMenuModelRef_SetFontList", 0), // function SetFontList
			/* 56 */ imports.NewTable("TCefMenuModelRef_SetFontListAt", 0), // function SetFontListAt
			/* 57 */ imports.NewTable("TCefMenuModelRef_UnWrap", 0), // static function UnWrap
			/* 58 */ imports.NewTable("TCefMenuModelRef_New", 0), // static function New
		}
	})
	return cefMenuModelRefImport
}
