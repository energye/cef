//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef misc functions

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/types"
	"unsafe"
)

// WindowInfoAsChild  Set to the specified window as a child window
func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	miscFunctionsImportAPI().Proc(misc_WindowInfoAsChild).Call(windowInfo, windowHandle, api.PascalStr(windowName))
}

// WindowInfoAsPopUp  Set to as a pop-up window
func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	miscFunctionsImportAPI().Proc(misc_WindowInfoAsPopUp).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

// WindowInfoAsWindowless  Set to as no window
func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	miscFunctionsImportAPI().Proc(misc_WindowInfoAsWindowless).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

// RegisterExtension registers the JS extension
//
// Insert custom JS code into the current browser
// Used in the WebKitInitialized callback function
//
// Parameter:
//
//	name: name of the root object
//	code: js code
//	handler: handler, which is called back according to the local function name
func RegisterExtension(name, code string, handler ICefV8Handler) {
	miscFunctionsImportAPI().Proc(misc_CefRegisterExtension).Call(api.PascalStr(name), api.PascalStr(code), handler.Instance())
}

func CheckSubprocessPath(subprocessPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CheckSubprocessPath).Call(api.PascalStr(subprocessPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckLocales(localesDirPath, localesRequired string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CheckLocales).Call(api.PascalStr(localesDirPath), uintptr(unsafe.Pointer(&missingFiles)), api.PascalStr(localesRequired))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckResources(resourcesDirPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CheckResources).Call(api.PascalStr(resourcesDirPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckDLLs(frameworkDirPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CheckDLLs).Call(api.PascalStr(frameworkDirPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func RegisterSchemeHandlerFactory(schemeName, domainName string, handler TCefResourceHandlerClass) bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefRegisterSchemeHandlerFactory).Call(api.PascalStr(schemeName), api.PascalStr(domainName), uintptr(handler))
	return api.GoBool(r1)
}

func ClearSchemeHandlerFactories() bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefClearSchemeHandlerFactories).Call()
	return api.GoBool(r1)
}

func GetMimeType(extension string) string {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefGetMimeType).Call(api.PascalStr(extension))
	return api.GoStr(r1)
}

func DeviceToLogicalInt32(value int32, deviceScaleFactor float64) int32 {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_DeviceToLogicalInt32).Call(uintptr(value), uintptr(unsafe.Pointer(&deviceScaleFactor)))
	return int32(r1)
}

func DeviceToLogicalFloat32(value float32, deviceScaleFactor float64) (result float32) {
	miscFunctionsImportAPI().Proc(misc_DeviceToLogicalFloat32).Call(uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&deviceScaleFactor)), uintptr(unsafe.Pointer(&result)))
	return
}

func DeviceToLogicalMouse(event *TCefMouseEvent, deviceScaleFactor float64) {
	miscFunctionsImportAPI().Proc(misc_DeviceToLogicalMouse).Call(uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func DeviceToLogicalTouch(event *TCefTouchEvent, deviceScaleFactor float64) {
	miscFunctionsImportAPI().Proc(misc_DeviceToLogicalTouch).Call(uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func DeviceToLogicalPoint(point *types.TPoint, deviceScaleFactor float64) {
	miscFunctionsImportAPI().Proc(misc_DeviceToLogicalPoint).Call(uintptr(unsafe.Pointer(point)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func LogicalToDeviceInt32(value int32, deviceScaleFactor float64) int32 {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_LogicalToDeviceInt32).Call(uintptr(value), uintptr(unsafe.Pointer(&deviceScaleFactor)))
	return int32(r1)
}

func LogicalToDeviceRect(rect *TCefRect, deviceScaleFactor float64) {
	miscFunctionsImportAPI().Proc(misc_LogicalToDeviceRect).Call(uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func InitializeWindowHandle() TCefWindowHandle {
	var result uintptr
	miscFunctionsImportAPI().Proc(misc_InitializeWindowHandle).Call(uintptr(unsafe.Pointer(&result)))
	return TCefWindowHandle(result)
}

func ValidCefWindowHandle(handle TCefWindowHandle) bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_ValidCefWindowHandle).Call(uintptr(handle))
	return api.GoBool(r1)
}

func GetScreenDPI() int32 {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_GetScreenDPI).Call()
	return int32(r1)
}

func GetDeviceScaleFactor() (result float32) {
	miscFunctionsImportAPI().Proc(misc_GetDeviceScaleFactor).Call(uintptr(unsafe.Pointer(&result)))
	return
}

func CefPostTask(threadId TCefThreadId, task ITask) bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefPostTask).Call(uintptr(threadId), task.Instance())
	return api.GoBool(r1)
}

func CefPostDelayedTask(threadId TCefThreadId, task ITask, delayMs int64) bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefPostDelayedTask).Call(uintptr(threadId), task.Instance(), uintptr(unsafe.Pointer(&delayMs)))
	return api.GoBool(r1)
}

func CefCurrentlyOn(threadId TCefThreadId) bool {
	r1, _, _ := miscFunctionsImportAPI().Proc(misc_CefCurrentlyOn).Call(uintptr(threadId))
	return api.GoBool(r1)
}

func CefCursorToWindowsCursor(cefCursor TCefCursorType) (result types.TCursor) {
	switch cefCursor {
	case CT_POINTER:
		result = types.CrArrow
	case CT_CROSS:
		result = types.CrCross
	case CT_HAND:
		result = types.CrHandPoint
	case CT_IBEAM:
		result = types.CrIBeam
	case CT_WAIT:
		result = types.CrHourGlass
	case CT_HELP:
		result = types.CrHelp
	case CT_EASTRESIZE:
		result = types.CrSizeWE
	case CT_NORTHRESIZE:
		result = types.CrSizeNS
	case CT_NORTHEASTRESIZE:
		result = types.CrSizeNESW
	case CT_NORTHWESTRESIZE:
		result = types.CrSizeNWSE
	case CT_SOUTHRESIZE:
		result = types.CrSizeNS
	case CT_SOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case CT_SOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case CT_WESTRESIZE:
		result = types.CrSizeWE
	case CT_NORTHSOUTHRESIZE:
		result = types.CrSizeNS
	case CT_EASTWESTRESIZE:
		result = types.CrSizeWE
	case CT_NORTHEASTSOUTHWESTRESIZE:
		result = types.CrSizeNESW
	case CT_NORTHWESTSOUTHEASTRESIZE:
		result = types.CrSizeNWSE
	case CT_COLUMNRESIZE:
		result = types.CrHSplit
	case CT_ROWRESIZE:
		result = types.CrVSplit
	case CT_MOVE:
		result = types.CrSizeAll
	case CT_PROGRESS:
		result = types.CrAppStart
	case CT_NONE:
		result = types.CrNone
	case CT_NODROP, CT_NOTALLOWED:
		result = types.CrNo
	case CT_GRAB, CT_GRABBING:
		result = types.CrDrag
	default:
		result = types.CrDefault
	}
	return
}

func CefColorGetA(color TCefColor) uint8 {
	return uint8(color>>24) & 0xFF
}

func CefColorGetR(color TCefColor) uint8 {
	return uint8(color>>16) & 0xFF
}

func CefColorGetG(color TCefColor) uint8 {
	return uint8(color>>8) & 0xFF
}

func CefColorGetB(color TCefColor) uint8 {
	return uint8(color) & 0xFF
}

func CefColorSetARGB(a, r, g, b byte) TCefColor {
	return TCefColor((uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | (uint32(b)))
}

const (
	misc_WindowInfoAsChild = iota
	misc_WindowInfoAsPopUp
	misc_WindowInfoAsWindowless
	misc_CefRegisterExtension
	misc_CheckSubprocessPath
	misc_CheckLocales
	misc_CheckResources
	misc_CheckDLLs
	misc_CefRegisterSchemeHandlerFactory
	misc_CefClearSchemeHandlerFactories
	misc_CefGetMimeType
	misc_DeviceToLogicalInt32
	misc_DeviceToLogicalFloat32
	misc_DeviceToLogicalMouse
	misc_DeviceToLogicalTouch
	misc_DeviceToLogicalPoint
	misc_LogicalToDeviceInt32
	misc_LogicalToDeviceRect
	misc_InitializeWindowHandle
	misc_ValidCefWindowHandle
	misc_GetScreenDPI
	misc_GetDeviceScaleFactor
	misc_CefPostTask
	misc_CefPostDelayedTask
	misc_CefCurrentlyOn
)

var (
	miscFunctionsImport       *imports.Imports = nil
	miscFunctionsImportTables                  = []*imports.Table{
		imports.NewTable("Misc_WindowInfoAsChild", 0),
		imports.NewTable("Misc_WindowInfoAsPopUp", 0),
		imports.NewTable("Misc_WindowInfoAsWindowless", 0),
		imports.NewTable("Misc_CefRegisterExtension", 0),
		imports.NewTable("Misc_CheckSubprocessPath", 0),
		imports.NewTable("Misc_CheckLocales", 0),
		imports.NewTable("Misc_CheckResources", 0),
		imports.NewTable("Misc_CheckDLLs", 0),
		imports.NewTable("Misc_CefRegisterSchemeHandlerFactory", 0),
		imports.NewTable("Misc_CefClearSchemeHandlerFactories", 0),
		imports.NewTable("Misc_CefGetMimeType", 0),
		imports.NewTable("Misc_DeviceToLogicalInt32", 0),
		imports.NewTable("Misc_DeviceToLogicalFloat32", 0),
		imports.NewTable("Misc_DeviceToLogicalMouse", 0),
		imports.NewTable("Misc_DeviceToLogicalTouch", 0),
		imports.NewTable("Misc_DeviceToLogicalPoint", 0),
		imports.NewTable("Misc_LogicalToDeviceInt32", 0),
		imports.NewTable("Misc_LogicalToDeviceRect", 0),
		imports.NewTable("Misc_InitializeWindowHandle", 0),
		imports.NewTable("Misc_ValidCefWindowHandle", 0),
		imports.NewTable("Misc_GetScreenDPI", 0),
		imports.NewTable("Misc_GetDeviceScaleFactor", 0),
		imports.NewTable("Misc_CefPostTask", 0),
		imports.NewTable("Misc_CefPostDelayedTask", 0),
		imports.NewTable("Misc_CefCurrentlyOn", 0),
	}
)

func miscFunctionsImportAPI() *imports.Imports {
	if miscFunctionsImport == nil {
		miscFunctionsImport = api.NewDefaultImports()
		miscFunctionsImport.SetImportTable(miscFunctionsImportTables)
		miscFunctionsImportTables = nil
	}
	return miscFunctionsImport
}
