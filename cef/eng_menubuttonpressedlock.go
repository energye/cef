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

// ICefMenuButtonPressedLock Parent: ICefBaseRefCounted
//
//	MenuButton pressed lock is released when this object is destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_pressed_lock_t)</a>
type ICefMenuButtonPressedLock interface {
	ICefBaseRefCounted
}

// TCefMenuButtonPressedLock Parent: TCefBaseRefCounted
//
//	MenuButton pressed lock is released when this object is destroyed.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_pressed_lock_t)</a>
type TCefMenuButtonPressedLock struct {
	TCefBaseRefCounted
}

// MenuButtonPressedLockRef -> ICefMenuButtonPressedLock
var MenuButtonPressedLockRef menuButtonPressedLock

// menuButtonPressedLock TCefMenuButtonPressedLock Ref
type menuButtonPressedLock uintptr

// UnWrap
//
//	Returns a ICefMenuButtonPressedLock instance using a PCefMenuButtonPressedLock data pointer.
func (m *menuButtonPressedLock) UnWrap(data uintptr) ICefMenuButtonPressedLock {
	var resultCefMenuButtonPressedLock uintptr
	menuButtonPressedLockImportAPI().SysCallN(0, uintptr(data), uintptr(unsafePointer(&resultCefMenuButtonPressedLock)))
	return AsCefMenuButtonPressedLock(resultCefMenuButtonPressedLock)
}

var (
	menuButtonPressedLockImport       *imports.Imports = nil
	menuButtonPressedLockImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMenuButtonPressedLock_UnWrap", 0),
	}
)

func menuButtonPressedLockImportAPI() *imports.Imports {
	if menuButtonPressedLockImport == nil {
		menuButtonPressedLockImport = NewDefaultImports()
		menuButtonPressedLockImport.SetImportTable(menuButtonPressedLockImportTables)
		menuButtonPressedLockImportTables = nil
	}
	return menuButtonPressedLockImport
}
