//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefWindow Parent: ICefPanel
type ICefWindow interface {
	ICefPanel
	IsClosed() bool                                                                                          // function
	IsActive() bool                                                                                          // function
	IsAlwaysOnTop() bool                                                                                     // function
	IsMaximized() bool                                                                                       // function
	IsMinimized() bool                                                                                       // function
	IsFullscreen() bool                                                                                      // function
	GetTitle() string                                                                                        // function
	GetWindowIcon() ICefImage                                                                                // function
	GetWindowAppIcon() ICefImage                                                                             // function
	AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode) ICefOverlayController                // function
	GetDisplay() ICefDisplay                                                                                 // function
	GetClientAreaBoundsInScreen() TCefRect                                                                   // function
	GetWindowHandle() cefTypes.TCefWindowHandle                                                              // function
	Show()                                                                                                   // procedure
	Hide()                                                                                                   // procedure
	CenterWindow(size TCefSize)                                                                              // procedure
	Close()                                                                                                  // procedure
	Activate()                                                                                               // procedure
	Deactivate()                                                                                             // procedure
	BringToTop()                                                                                             // procedure
	SetAlwaysOnTop(onTop bool)                                                                               // procedure
	Maximize()                                                                                               // procedure
	Minimize()                                                                                               // procedure
	Restore()                                                                                                // procedure
	SetFullscreen(fullscreen bool)                                                                           // procedure
	SetTitle(title string)                                                                                   // procedure
	SetWindowIcon(image ICefImage)                                                                           // procedure
	SetWindowAppIcon(image ICefImage)                                                                        // procedure
	ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) // procedure
	CancelMenu()                                                                                             // procedure
	SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray)                  // procedure
	SendKeyPress(keyCode int32, eventFlags uint32)                                                           // procedure
	SendMouseMove(screenX int32, screenY int32)                                                              // procedure
	SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool)                       // procedure
	SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool)     // procedure
	RemoveAccelerator(commandId int32)                                                                       // procedure
	RemoveAllAccelerators()                                                                                  // procedure
}

// ICefWindowRef Parent: ICefWindow ICefPanelRef
type ICefWindowRef interface {
	ICefWindow
	ICefPanelRef
	AsIntfWindow() uintptr
	AsIntfPanel() uintptr
	AsIntfView() uintptr
}

type TCefWindowRef struct {
	TCefPanelRef
}

func (m *TCefWindowRef) IsClosed() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) IsAlwaysOnTop() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) IsMaximized() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) IsMinimized() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := cefWindowRefAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefWindowRef) GetTitle() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefWindowRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefWindowRef) GetWindowIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCefWindowRef) GetWindowAppIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCefWindowRef) AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode) (result ICefOverlayController) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(view), uintptr(dockingMode), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefOverlayControllerRef(resultPtr)
	return
}

func (m *TCefWindowRef) GetDisplay() (result ICefDisplay) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

func (m *TCefWindowRef) GetClientAreaBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefWindowRef) GetWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cefWindowRefAPI().SysCallN(13, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCefWindowRef) Show() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(16, m.Instance())
}

func (m *TCefWindowRef) Hide() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(17, m.Instance())
}

func (m *TCefWindowRef) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefWindowRef) Close() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(19, m.Instance())
}

func (m *TCefWindowRef) Activate() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(20, m.Instance())
}

func (m *TCefWindowRef) Deactivate() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(21, m.Instance())
}

func (m *TCefWindowRef) BringToTop() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(22, m.Instance())
}

func (m *TCefWindowRef) SetAlwaysOnTop(onTop bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(23, m.Instance(), api.PasBool(onTop))
}

func (m *TCefWindowRef) Maximize() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(24, m.Instance())
}

func (m *TCefWindowRef) Minimize() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(25, m.Instance())
}

func (m *TCefWindowRef) Restore() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(26, m.Instance())
}

func (m *TCefWindowRef) SetFullscreen(fullscreen bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(27, m.Instance(), api.PasBool(fullscreen))
}

func (m *TCefWindowRef) SetTitle(title string) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(28, m.Instance(), api.PasStr(title))
}

func (m *TCefWindowRef) SetWindowIcon(image ICefImage) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(29, m.Instance(), base.GetObjectUintptr(image))
}

func (m *TCefWindowRef) SetWindowAppIcon(image ICefImage) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(30, m.Instance(), base.GetObjectUintptr(image))
}

func (m *TCefWindowRef) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCefWindowRef) CancelMenu() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(32, m.Instance())
}

func (m *TCefWindowRef) SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(33, m.Instance(), uintptr(regionsCount), regions.Instance())
}

func (m *TCefWindowRef) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(34, m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCefWindowRef) SendMouseMove(screenX int32, screenY int32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(35, m.Instance(), uintptr(screenX), uintptr(screenY))
}

func (m *TCefWindowRef) SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(36, m.Instance(), uintptr(button), api.PasBool(mouseDown), api.PasBool(mouseUp))
}

func (m *TCefWindowRef) SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(37, m.Instance(), uintptr(commandId), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed))
}

func (m *TCefWindowRef) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(38, m.Instance(), uintptr(commandId))
}

func (m *TCefWindowRef) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(39, m.Instance())
}

func (m *TCefWindowRef) AsIntfWindow() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefWindowRef) AsIntfPanel() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCefWindowRef) AsIntfView() uintptr {
	return m.GetIntfPointer(2)
}

// WindowRef  is static instance
var WindowRef _WindowRefClass

// _WindowRefClass is class type defined by TCefWindowRef
type _WindowRefClass uintptr

func (_WindowRefClass) UnWrapWithPointer(data uintptr) (result ICefWindow) {
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(14, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (_WindowRefClass) CreateTopLevel(delegate IEngWindowDelegate) (result ICefWindow) {
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(15, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

// NewWindowRef class constructor
func NewWindowRef(data uintptr) ICefWindowRef {
	var windowPtr uintptr // ICefWindow
	var panelPtr uintptr  // ICefPanel
	var viewPtr uintptr   // ICefView
	r := cefWindowRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&windowPtr)), uintptr(base.UnsafePointer(&panelPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefWindowRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowPtr)
		intf.SetIntfPointer(1, panelPtr)
		intf.SetIntfPointer(2, viewPtr)
	}
	return ret
}

var (
	cefWindowRefOnce   base.Once
	cefWindowRefImport *imports.Imports = nil
)

func cefWindowRefAPI() *imports.Imports {
	cefWindowRefOnce.Do(func() {
		cefWindowRefImport = api.NewDefaultImports()
		cefWindowRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefWindowRef_Create", 0), // constructor NewWindowRef
			/* 1 */ imports.NewTable("TCefWindowRef_IsClosed", 0), // function IsClosed
			/* 2 */ imports.NewTable("TCefWindowRef_IsActive", 0), // function IsActive
			/* 3 */ imports.NewTable("TCefWindowRef_IsAlwaysOnTop", 0), // function IsAlwaysOnTop
			/* 4 */ imports.NewTable("TCefWindowRef_IsMaximized", 0), // function IsMaximized
			/* 5 */ imports.NewTable("TCefWindowRef_IsMinimized", 0), // function IsMinimized
			/* 6 */ imports.NewTable("TCefWindowRef_IsFullscreen", 0), // function IsFullscreen
			/* 7 */ imports.NewTable("TCefWindowRef_GetTitle", 0), // function GetTitle
			/* 8 */ imports.NewTable("TCefWindowRef_GetWindowIcon", 0), // function GetWindowIcon
			/* 9 */ imports.NewTable("TCefWindowRef_GetWindowAppIcon", 0), // function GetWindowAppIcon
			/* 10 */ imports.NewTable("TCefWindowRef_AddOverlayView", 0), // function AddOverlayView
			/* 11 */ imports.NewTable("TCefWindowRef_GetDisplay", 0), // function GetDisplay
			/* 12 */ imports.NewTable("TCefWindowRef_GetClientAreaBoundsInScreen", 0), // function GetClientAreaBoundsInScreen
			/* 13 */ imports.NewTable("TCefWindowRef_GetWindowHandle", 0), // function GetWindowHandle
			/* 14 */ imports.NewTable("TCefWindowRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 15 */ imports.NewTable("TCefWindowRef_CreateTopLevel", 0), // static function CreateTopLevel
			/* 16 */ imports.NewTable("TCefWindowRef_Show", 0), // procedure Show
			/* 17 */ imports.NewTable("TCefWindowRef_Hide", 0), // procedure Hide
			/* 18 */ imports.NewTable("TCefWindowRef_CenterWindow", 0), // procedure CenterWindow
			/* 19 */ imports.NewTable("TCefWindowRef_Close", 0), // procedure Close
			/* 20 */ imports.NewTable("TCefWindowRef_Activate", 0), // procedure Activate
			/* 21 */ imports.NewTable("TCefWindowRef_Deactivate", 0), // procedure Deactivate
			/* 22 */ imports.NewTable("TCefWindowRef_BringToTop", 0), // procedure BringToTop
			/* 23 */ imports.NewTable("TCefWindowRef_SetAlwaysOnTop", 0), // procedure SetAlwaysOnTop
			/* 24 */ imports.NewTable("TCefWindowRef_Maximize", 0), // procedure Maximize
			/* 25 */ imports.NewTable("TCefWindowRef_Minimize", 0), // procedure Minimize
			/* 26 */ imports.NewTable("TCefWindowRef_Restore", 0), // procedure Restore
			/* 27 */ imports.NewTable("TCefWindowRef_SetFullscreen", 0), // procedure SetFullscreen
			/* 28 */ imports.NewTable("TCefWindowRef_SetTitle", 0), // procedure SetTitle
			/* 29 */ imports.NewTable("TCefWindowRef_SetWindowIcon", 0), // procedure SetWindowIcon
			/* 30 */ imports.NewTable("TCefWindowRef_SetWindowAppIcon", 0), // procedure SetWindowAppIcon
			/* 31 */ imports.NewTable("TCefWindowRef_ShowMenu", 0), // procedure ShowMenu
			/* 32 */ imports.NewTable("TCefWindowRef_CancelMenu", 0), // procedure CancelMenu
			/* 33 */ imports.NewTable("TCefWindowRef_SetDraggableRegions", 0), // procedure SetDraggableRegions
			/* 34 */ imports.NewTable("TCefWindowRef_SendKeyPress", 0), // procedure SendKeyPress
			/* 35 */ imports.NewTable("TCefWindowRef_SendMouseMove", 0), // procedure SendMouseMove
			/* 36 */ imports.NewTable("TCefWindowRef_SendMouseEvents", 0), // procedure SendMouseEvents
			/* 37 */ imports.NewTable("TCefWindowRef_SetAccelerator", 0), // procedure SetAccelerator
			/* 38 */ imports.NewTable("TCefWindowRef_RemoveAccelerator", 0), // procedure RemoveAccelerator
			/* 39 */ imports.NewTable("TCefWindowRef_RemoveAllAccelerators", 0), // procedure RemoveAllAccelerators
		}
	})
	return cefWindowRefImport
}
