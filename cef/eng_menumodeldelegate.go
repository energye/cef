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

// IMenuModelDelegate Parent: ICefMenuModelDelegate
//
//	Implement this interface to handle menu model events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a cref="uCEFTypes|TCefMenuModelDelegate">Implements TCefMenuModelDelegate</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_delegate_capi.h">CEF source file: /include/capi/cef_menu_model_delegate_capi.h (cef_menu_model_delegate_t)</a>
type IMenuModelDelegate interface {
	ICefMenuModelDelegate
	AsInterface() ICefMenuModelDelegate // function
	// SetOnExecuteCommand
	//  Perform the action associated with the specified |command_id| and optional
	//  |event_flags|.
	SetOnExecuteCommand(fn TOnExecuteCommand) // property event
	// SetOnMouseOutsideMenu
	//  Called when the user moves the mouse outside the menu and over the owning
	//  window.
	SetOnMouseOutsideMenu(fn TOnMouseOutsideMenu) // property event
	// SetOnUnhandledOpenSubmenu
	//  Called on unhandled open submenu keyboard commands. |is_rtl| will be true
	//  (1) if the menu is displaying a right-to-left language.
	SetOnUnhandledOpenSubmenu(fn TOnUnhandledOpenSubmenu) // property event
	// SetOnUnhandledCloseSubmenu
	//  Called on unhandled close submenu keyboard commands. |is_rtl| will be true
	//  (1) if the menu is displaying a right-to-left language.
	SetOnUnhandledCloseSubmenu(fn TOnUnhandledCloseSubmenu) // property event
	// SetOnMenuWillShow
	//  The menu is about to show.
	SetOnMenuWillShow(fn TOnMenuWillShow) // property event
	// SetOnMenuClosed
	//  The menu has closed.
	SetOnMenuClosed(fn TOnMenuClosed) // property event
	// SetOnFormatLabel
	//  Optionally modify a menu item label. Return true(1) if |label| was
	//  modified.
	SetOnFormatLabel(fn TOnFormatLabel) // property event
}

// TMenuModelDelegate Parent: TCefMenuModelDelegate
//
//	Implement this interface to handle menu model events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a cref="uCEFTypes|TCefMenuModelDelegate">Implements TCefMenuModelDelegate</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_delegate_capi.h">CEF source file: /include/capi/cef_menu_model_delegate_capi.h (cef_menu_model_delegate_t)</a>
type TMenuModelDelegate struct {
	TCefMenuModelDelegate
	executeCommandPtr        uintptr
	mouseOutsideMenuPtr      uintptr
	unhandledOpenSubmenuPtr  uintptr
	unhandledCloseSubmenuPtr uintptr
	menuWillShowPtr          uintptr
	menuClosedPtr            uintptr
	formatLabelPtr           uintptr
}

func NewMenuModelDelegate() IMenuModelDelegate {
	r1 := menuModelDelegateImportAPI().SysCallN(1)
	return AsMenuModelDelegate(r1)
}

func (m *TMenuModelDelegate) AsInterface() ICefMenuModelDelegate {
	var resultCefMenuModelDelegate uintptr
	menuModelDelegateImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefMenuModelDelegate)))
	return AsCefMenuModelDelegate(resultCefMenuModelDelegate)
}

func (m *TMenuModelDelegate) SetOnExecuteCommand(fn TOnExecuteCommand) {
	if m.executeCommandPtr != 0 {
		RemoveEventElement(m.executeCommandPtr)
	}
	m.executeCommandPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(2, m.Instance(), m.executeCommandPtr)
}

func (m *TMenuModelDelegate) SetOnMouseOutsideMenu(fn TOnMouseOutsideMenu) {
	if m.mouseOutsideMenuPtr != 0 {
		RemoveEventElement(m.mouseOutsideMenuPtr)
	}
	m.mouseOutsideMenuPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(6, m.Instance(), m.mouseOutsideMenuPtr)
}

func (m *TMenuModelDelegate) SetOnUnhandledOpenSubmenu(fn TOnUnhandledOpenSubmenu) {
	if m.unhandledOpenSubmenuPtr != 0 {
		RemoveEventElement(m.unhandledOpenSubmenuPtr)
	}
	m.unhandledOpenSubmenuPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(8, m.Instance(), m.unhandledOpenSubmenuPtr)
}

func (m *TMenuModelDelegate) SetOnUnhandledCloseSubmenu(fn TOnUnhandledCloseSubmenu) {
	if m.unhandledCloseSubmenuPtr != 0 {
		RemoveEventElement(m.unhandledCloseSubmenuPtr)
	}
	m.unhandledCloseSubmenuPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(7, m.Instance(), m.unhandledCloseSubmenuPtr)
}

func (m *TMenuModelDelegate) SetOnMenuWillShow(fn TOnMenuWillShow) {
	if m.menuWillShowPtr != 0 {
		RemoveEventElement(m.menuWillShowPtr)
	}
	m.menuWillShowPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(5, m.Instance(), m.menuWillShowPtr)
}

func (m *TMenuModelDelegate) SetOnMenuClosed(fn TOnMenuClosed) {
	if m.menuClosedPtr != 0 {
		RemoveEventElement(m.menuClosedPtr)
	}
	m.menuClosedPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(4, m.Instance(), m.menuClosedPtr)
}

func (m *TMenuModelDelegate) SetOnFormatLabel(fn TOnFormatLabel) {
	if m.formatLabelPtr != 0 {
		RemoveEventElement(m.formatLabelPtr)
	}
	m.formatLabelPtr = MakeEventDataPtr(fn)
	menuModelDelegateImportAPI().SysCallN(3, m.Instance(), m.formatLabelPtr)
}

var (
	menuModelDelegateImport       *imports.Imports = nil
	menuModelDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MenuModelDelegate_AsInterface", 0),
		/*1*/ imports.NewTable("MenuModelDelegate_Create", 0),
		/*2*/ imports.NewTable("MenuModelDelegate_SetOnExecuteCommand", 0),
		/*3*/ imports.NewTable("MenuModelDelegate_SetOnFormatLabel", 0),
		/*4*/ imports.NewTable("MenuModelDelegate_SetOnMenuClosed", 0),
		/*5*/ imports.NewTable("MenuModelDelegate_SetOnMenuWillShow", 0),
		/*6*/ imports.NewTable("MenuModelDelegate_SetOnMouseOutsideMenu", 0),
		/*7*/ imports.NewTable("MenuModelDelegate_SetOnUnhandledCloseSubmenu", 0),
		/*8*/ imports.NewTable("MenuModelDelegate_SetOnUnhandledOpenSubmenu", 0),
	}
)

func menuModelDelegateImportAPI() *imports.Imports {
	if menuModelDelegateImport == nil {
		menuModelDelegateImport = NewDefaultImports()
		menuModelDelegateImport.SetImportTable(menuModelDelegateImportTables)
		menuModelDelegateImportTables = nil
	}
	return menuModelDelegateImport
}
