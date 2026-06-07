//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// :predefine:

package cef

import (
	"unsafe"

	cefTypes "github.com/energye/cef/147/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// WindowInfoAsChild  Set to the specified window as a child window
func (_MiscFuncClass) WindowInfoAsChild(windowInfo *TCefWindowInfo, windowParent ICEFWinControl, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	miscFunctionsImportAPI().SysCallN(0, uintptr(unsafe.Pointer(windowInfoPtr)), windowParent.Instance(), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

// WindowInfoAsPopUp  Set to as a pop-up window
func (_MiscFuncClass) WindowInfoAsPopUp(windowInfo *TCefWindowInfo, windowParent ICEFWinControl, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	miscFunctionsImportAPI().SysCallN(1, uintptr(unsafe.Pointer(windowInfoPtr)), windowParent.Instance(), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

// WindowInfoAsWindowless  Set to as no window
func (_MiscFuncClass) WindowInfoAsWindowless(windowInfo *TCefWindowInfo, windowParent ICEFWinControl, windowName string, exStyle types.DWORD) {
	windowInfoPtr := windowInfo.ToPas()
	miscFunctionsImportAPI().SysCallN(2, uintptr(unsafe.Pointer(windowInfoPtr)), windowParent.Instance(), api.PasStr(windowName), uintptr(exStyle))
	*windowInfo = windowInfoPtr.ToGo()
}

func (_MiscFuncClass) CefCursorToWindowsCursor(cefCursor cefTypes.TCefCursorType) (result types.TCursor) {
	switch cefCursor {
	case cefTypes.CT_POINTER:
		result = types.CrArrow
	case cefTypes.CT_CROSS:
		result = types.CrCross
	case cefTypes.CT_HAND:
		result = types.CrHandPoint
	case cefTypes.CT_IBEAM:
		result = types.CrIBeam
	case cefTypes.CT_WAIT:
		result = types.CrHourGlass
	case cefTypes.CT_HELP:
		result = types.CrHelp
	case cefTypes.CT_EASTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_NORTHEASTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_NORTHWESTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_SOUTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_SOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_SOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_WESTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHSOUTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_EASTWESTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHEASTSOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_NORTHWESTSOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_COLUMNRESIZE:
		result = types.CrHSplit
	case cefTypes.CT_ROWRESIZE:
		result = types.CrVSplit
	case cefTypes.CT_MOVE:
		result = types.CrSizeAll
	case cefTypes.CT_PROGRESS:
		result = types.CrAppStart
	case cefTypes.CT_NONE:
		result = types.CrNone
	case cefTypes.CT_NODROP, cefTypes.CT_NOTALLOWED:
		result = types.CrNo
	case cefTypes.CT_GRAB, cefTypes.CT_GRABBING:
		result = types.CrDrag
	default:
		result = types.CrDefault
	}
	return
}

var (
	miscFunctionsImportOnce   base.Once
	miscFunctionsImportTables *imports.Imports = nil
)

func miscFunctionsImportAPI() *imports.Imports {
	miscFunctionsImportOnce.Do(func() {
		miscFunctionsImportTables = api.NewDefaultImports()
		miscFunctionsImportTables.Table = []*imports.Table{
			/* 0 */ imports.NewTable("Misc_WindowInfoAsChild", 0),
			/* 1 */ imports.NewTable("Misc_WindowInfoAsPopUp", 0),
			/* 2 */ imports.NewTable("Misc_WindowInfoAsWindowless", 0),
		}
	})
	return miscFunctionsImportTables
}

func CursorToWindowsCursor(cefCursor cefTypes.TCefCursorType) (result types.TCursor) {
	switch cefCursor {
	case cefTypes.CT_POINTER:
		result = types.CrArrow
	case cefTypes.CT_CROSS:
		result = types.CrCross
	case cefTypes.CT_HAND:
		result = types.CrHandPoint
	case cefTypes.CT_IBEAM:
		result = types.CrIBeam
	case cefTypes.CT_WAIT:
		result = types.CrHourGlass
	case cefTypes.CT_HELP:
		result = types.CrHelp
	case cefTypes.CT_EASTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_NORTHEASTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_NORTHWESTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_SOUTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_SOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_SOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_WESTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHSOUTHRESIZE:
		result = types.CrSizeNS
	case cefTypes.CT_EASTWESTRESIZE:
		result = types.CrSizeWE
	case cefTypes.CT_NORTHEASTSOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case cefTypes.CT_NORTHWESTSOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case cefTypes.CT_COLUMNRESIZE:
		result = types.CrHSplit
	case cefTypes.CT_ROWRESIZE:
		result = types.CrVSplit
	case cefTypes.CT_MOVE:
		result = types.CrSizeAll
	case cefTypes.CT_PROGRESS:
		result = types.CrAppStart
	case cefTypes.CT_NONE:
		result = types.CrNone
	case cefTypes.CT_NODROP, cefTypes.CT_NOTALLOWED:
		result = types.CrNo
	case cefTypes.CT_GRAB, cefTypes.CT_GRABBING:
		result = types.CrDrag
	default:
		result = types.CrDefault
	}
	return
}

func ColorGetA(color cefTypes.TCefColor) uint8 {
	return uint8(color>>24) & 0xFF
}

func ColorGetR(color cefTypes.TCefColor) uint8 {
	return uint8(color>>16) & 0xFF
}

func ColorGetG(color cefTypes.TCefColor) uint8 {
	return uint8(color>>8) & 0xFF
}

func ColorGetB(color cefTypes.TCefColor) uint8 {
	return uint8(color) & 0xFF
}

func ColorSetARGB(a, r, g, b byte) cefTypes.TCefColor {
	return cefTypes.TCefColor((uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | (uint32(b)))
}
