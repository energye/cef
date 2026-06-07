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

	cefTypes "github.com/energye/cef/types"
)

// ICefWindowDelegateEvents Parent: ICefPanelDelegateEvents
type ICefWindowDelegateEvents interface {
	ICefPanelDelegateEvents
}

// ICEFWindowComponent Parent: ICefWindowDelegateEvents ICEFPanelComponent
type ICEFWindowComponent interface {
	ICefWindowDelegateEvents
	ICEFPanelComponent
	AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode) ICefOverlayController                // function
	CreateTopLevelWindow()                                                                                   // procedure
	Show()                                                                                                   // procedure
	Hide()                                                                                                   // procedure
	CenterWindow(size TCefSize)                                                                              // procedure
	Close()                                                                                                  // procedure
	Activate()                                                                                               // procedure
	Deactivate()                                                                                             // procedure
	BringToTop()                                                                                             // procedure
	Maximize()                                                                                               // procedure
	Minimize()                                                                                               // procedure
	Restore()                                                                                                // procedure
	ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) // procedure
	CancelMenu()                                                                                             // procedure
	SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray)                  // procedure
	SendKeyPress(keyCode int32, eventFlags uint32)                                                           // procedure
	SendMouseMove(screenX int32, screenY int32)                                                              // procedure
	SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool)                       // procedure
	SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool)     // procedure
	RemoveAccelerator(commandId int32)                                                                       // procedure
	RemoveAllAccelerators()                                                                                  // procedure
	Title() string                                                                                           // property Title Getter
	SetTitle(value string)                                                                                   // property Title Setter
	WindowIcon() ICefImage                                                                                   // property WindowIcon Getter
	SetWindowIcon(value ICefImage)                                                                           // property WindowIcon Setter
	WindowAppIcon() ICefImage                                                                                // property WindowAppIcon Getter
	SetWindowAppIcon(value ICefImage)                                                                        // property WindowAppIcon Setter
	Display() ICefDisplay                                                                                    // property Display Getter
	ClientAreaBoundsInScreen() TCefRect                                                                      // property ClientAreaBoundsInScreen Getter
	WindowHandle() cefTypes.TCefWindowHandle                                                                 // property WindowHandle Getter
	IsClosed() bool                                                                                          // property IsClosed Getter
	IsActive() bool                                                                                          // property IsActive Getter
	IsAlwaysOnTop() bool                                                                                     // property IsAlwaysOnTop Getter
	SetIsAlwaysOnTop(value bool)                                                                             // property IsAlwaysOnTop Setter
	IsFullscreen() bool                                                                                      // property IsFullscreen Getter
	SetIsFullscreen(value bool)                                                                              // property IsFullscreen Setter
	IsMaximized() bool                                                                                       // property IsMaximized Getter
	IsMinimized() bool                                                                                       // property IsMinimized Getter
	SetOnWindowCreated(fn TOnWindowCreatedEvent)                                                             // property event
	SetOnWindowClosing(fn TOnWindowClosingEvent)                                                             // property event
	SetOnWindowDestroyed(fn TOnWindowDestroyedEvent)                                                         // property event
	SetOnWindowActivationChanged(fn TOnWindowActivationChangedEvent)                                         // property event
	SetOnWindowBoundsChanged(fn TOnWindowBoundsChangedEvent)                                                 // property event
	SetOnGetParentWindow(fn TOnGetParentWindowEvent)                                                         // property event
	SetOnGetInitialBounds(fn TOnGetInitialBoundsEvent)                                                       // property event
	SetOnGetInitialShowState(fn TOnGetInitialShowStateEvent)                                                 // property event
	SetOnIsFrameless(fn TOnIsFramelessEvent)                                                                 // property event
	SetOnCanResize(fn TOnCanResizeEvent)                                                                     // property event
	SetOnCanMaximize(fn TOnCanMaximizeEvent)                                                                 // property event
	SetOnCanMinimize(fn TOnCanMinimizeEvent)                                                                 // property event
	SetOnCanClose(fn TOnCanCloseEvent)                                                                       // property event
	SetOnAccelerator(fn TOnAcceleratorEvent)                                                                 // property event
	SetOnKeyEvent(fn TOnWindowKeyEventEvent)                                                                 // property event
	AsIntfWindowDelegateEvents() uintptr
	AsIntfPanelDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFWindowComponent struct {
	TCEFPanelComponent
}

func (m *TCEFWindowComponent) AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode) (result ICefOverlayController) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(view), uintptr(dockingMode), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefOverlayControllerRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(2, m.Instance())
}

func (m *TCEFWindowComponent) Show() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(3, m.Instance())
}

func (m *TCEFWindowComponent) Hide() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(4, m.Instance())
}

func (m *TCEFWindowComponent) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCEFWindowComponent) Close() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(6, m.Instance())
}

func (m *TCEFWindowComponent) Activate() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(7, m.Instance())
}

func (m *TCEFWindowComponent) Deactivate() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(8, m.Instance())
}

func (m *TCEFWindowComponent) BringToTop() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(9, m.Instance())
}

func (m *TCEFWindowComponent) Maximize() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(10, m.Instance())
}

func (m *TCEFWindowComponent) Minimize() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFWindowComponent) Restore() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(12, m.Instance())
}

func (m *TCEFWindowComponent) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCEFWindowComponent) CancelMenu() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(14, m.Instance())
}

func (m *TCEFWindowComponent) SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(15, m.Instance(), uintptr(regionsCount), regions.Instance())
}

func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(16, m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCEFWindowComponent) SendMouseMove(screenX int32, screenY int32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(17, m.Instance(), uintptr(screenX), uintptr(screenY))
}

func (m *TCEFWindowComponent) SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(18, m.Instance(), uintptr(button), api.PasBool(mouseDown), api.PasBool(mouseUp))
}

func (m *TCEFWindowComponent) SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(19, m.Instance(), uintptr(commandId), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed))
}

func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(20, m.Instance(), uintptr(commandId))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(21, m.Instance())
}

func (m *TCEFWindowComponent) Title() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFWindowComponentAPI().SysCallN(22, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFWindowComponent) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(22, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFWindowComponent) WindowIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(23, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) SetWindowIcon(value ICefImage) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFWindowComponent) WindowAppIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(24, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) SetWindowAppIcon(value ICefImage) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(24, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFWindowComponent) Display() (result ICefDisplay) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(26, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFWindowComponent) WindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cEFWindowComponentAPI().SysCallN(27, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCEFWindowComponent) IsClosed() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(30, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) SetIsAlwaysOnTop(value bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(30, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFWindowComponent) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(31, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) SetIsFullscreen(value bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(31, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFWindowComponent) IsMaximized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(32, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsMinimized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(33, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) SetOnWindowCreated(fn TOnWindowCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCreatedEvent(fn)
	base.SetEvent(m, 34, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowClosing(fn TOnWindowClosingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowClosingEvent(fn)
	base.SetEvent(m, 35, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn TOnWindowDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowDestroyedEvent(fn)
	base.SetEvent(m, 36, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn TOnWindowActivationChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowActivationChangedEvent(fn)
	base.SetEvent(m, 37, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowBoundsChanged(fn TOnWindowBoundsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowBoundsChangedEvent(fn)
	base.SetEvent(m, 38, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetParentWindow(fn TOnGetParentWindowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetParentWindowEvent(fn)
	base.SetEvent(m, 39, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn TOnGetInitialBoundsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetInitialBoundsEvent(fn)
	base.SetEvent(m, 40, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn TOnGetInitialShowStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetInitialShowStateEvent(fn)
	base.SetEvent(m, 41, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnIsFrameless(fn TOnIsFramelessEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsFramelessEvent(fn)
	base.SetEvent(m, 42, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanResize(fn TOnCanResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanResizeEvent(fn)
	base.SetEvent(m, 43, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanMaximize(fn TOnCanMaximizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanMaximizeEvent(fn)
	base.SetEvent(m, 44, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanMinimize(fn TOnCanMinimizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanMinimizeEvent(fn)
	base.SetEvent(m, 45, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanClose(fn TOnCanCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanCloseEvent(fn)
	base.SetEvent(m, 46, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnAccelerator(fn TOnAcceleratorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceleratorEvent(fn)
	base.SetEvent(m, 47, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnKeyEvent(fn TOnWindowKeyEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowKeyEventEvent(fn)
	base.SetEvent(m, 48, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) AsIntfWindowDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFWindowComponent) AsIntfPanelDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCEFWindowComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(2)
}

// NewWindowComponent class constructor
func NewWindowComponent(owner lcl.IComponent) ICEFWindowComponent {
	var windowDelegateEventsPtr uintptr // ICefWindowDelegateEvents
	var panelDelegateEventsPtr uintptr  // ICefPanelDelegateEvents
	var viewDelegateEventsPtr uintptr   // ICefViewDelegateEvents
	r := cEFWindowComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&windowDelegateEventsPtr)), uintptr(base.UnsafePointer(&panelDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFWindowComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowDelegateEventsPtr)
		intf.SetIntfPointer(1, panelDelegateEventsPtr)
		intf.SetIntfPointer(2, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFWindowComponentOnce   base.Once
	cEFWindowComponentImport *imports.Imports = nil
)

func cEFWindowComponentAPI() *imports.Imports {
	cEFWindowComponentOnce.Do(func() {
		cEFWindowComponentImport = api.NewDefaultImports()
		cEFWindowComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFWindowComponent_Create", 0), // constructor NewWindowComponent
			/* 1 */ imports.NewTable("TCEFWindowComponent_AddOverlayView", 0), // function AddOverlayView
			/* 2 */ imports.NewTable("TCEFWindowComponent_CreateTopLevelWindow", 0), // procedure CreateTopLevelWindow
			/* 3 */ imports.NewTable("TCEFWindowComponent_Show", 0), // procedure Show
			/* 4 */ imports.NewTable("TCEFWindowComponent_Hide", 0), // procedure Hide
			/* 5 */ imports.NewTable("TCEFWindowComponent_CenterWindow", 0), // procedure CenterWindow
			/* 6 */ imports.NewTable("TCEFWindowComponent_Close", 0), // procedure Close
			/* 7 */ imports.NewTable("TCEFWindowComponent_Activate", 0), // procedure Activate
			/* 8 */ imports.NewTable("TCEFWindowComponent_Deactivate", 0), // procedure Deactivate
			/* 9 */ imports.NewTable("TCEFWindowComponent_BringToTop", 0), // procedure BringToTop
			/* 10 */ imports.NewTable("TCEFWindowComponent_Maximize", 0), // procedure Maximize
			/* 11 */ imports.NewTable("TCEFWindowComponent_Minimize", 0), // procedure Minimize
			/* 12 */ imports.NewTable("TCEFWindowComponent_Restore", 0), // procedure Restore
			/* 13 */ imports.NewTable("TCEFWindowComponent_ShowMenu", 0), // procedure ShowMenu
			/* 14 */ imports.NewTable("TCEFWindowComponent_CancelMenu", 0), // procedure CancelMenu
			/* 15 */ imports.NewTable("TCEFWindowComponent_SetDraggableRegions", 0), // procedure SetDraggableRegions
			/* 16 */ imports.NewTable("TCEFWindowComponent_SendKeyPress", 0), // procedure SendKeyPress
			/* 17 */ imports.NewTable("TCEFWindowComponent_SendMouseMove", 0), // procedure SendMouseMove
			/* 18 */ imports.NewTable("TCEFWindowComponent_SendMouseEvents", 0), // procedure SendMouseEvents
			/* 19 */ imports.NewTable("TCEFWindowComponent_SetAccelerator", 0), // procedure SetAccelerator
			/* 20 */ imports.NewTable("TCEFWindowComponent_RemoveAccelerator", 0), // procedure RemoveAccelerator
			/* 21 */ imports.NewTable("TCEFWindowComponent_RemoveAllAccelerators", 0), // procedure RemoveAllAccelerators
			/* 22 */ imports.NewTable("TCEFWindowComponent_Title", 0), // property Title
			/* 23 */ imports.NewTable("TCEFWindowComponent_WindowIcon", 0), // property WindowIcon
			/* 24 */ imports.NewTable("TCEFWindowComponent_WindowAppIcon", 0), // property WindowAppIcon
			/* 25 */ imports.NewTable("TCEFWindowComponent_Display", 0), // property Display
			/* 26 */ imports.NewTable("TCEFWindowComponent_ClientAreaBoundsInScreen", 0), // property ClientAreaBoundsInScreen
			/* 27 */ imports.NewTable("TCEFWindowComponent_WindowHandle", 0), // property WindowHandle
			/* 28 */ imports.NewTable("TCEFWindowComponent_IsClosed", 0), // property IsClosed
			/* 29 */ imports.NewTable("TCEFWindowComponent_IsActive", 0), // property IsActive
			/* 30 */ imports.NewTable("TCEFWindowComponent_IsAlwaysOnTop", 0), // property IsAlwaysOnTop
			/* 31 */ imports.NewTable("TCEFWindowComponent_IsFullscreen", 0), // property IsFullscreen
			/* 32 */ imports.NewTable("TCEFWindowComponent_IsMaximized", 0), // property IsMaximized
			/* 33 */ imports.NewTable("TCEFWindowComponent_IsMinimized", 0), // property IsMinimized
			/* 34 */ imports.NewTable("TCEFWindowComponent_OnWindowCreated", 0), // event OnWindowCreated
			/* 35 */ imports.NewTable("TCEFWindowComponent_OnWindowClosing", 0), // event OnWindowClosing
			/* 36 */ imports.NewTable("TCEFWindowComponent_OnWindowDestroyed", 0), // event OnWindowDestroyed
			/* 37 */ imports.NewTable("TCEFWindowComponent_OnWindowActivationChanged", 0), // event OnWindowActivationChanged
			/* 38 */ imports.NewTable("TCEFWindowComponent_OnWindowBoundsChanged", 0), // event OnWindowBoundsChanged
			/* 39 */ imports.NewTable("TCEFWindowComponent_OnGetParentWindow", 0), // event OnGetParentWindow
			/* 40 */ imports.NewTable("TCEFWindowComponent_OnGetInitialBounds", 0), // event OnGetInitialBounds
			/* 41 */ imports.NewTable("TCEFWindowComponent_OnGetInitialShowState", 0), // event OnGetInitialShowState
			/* 42 */ imports.NewTable("TCEFWindowComponent_OnIsFrameless", 0), // event OnIsFrameless
			/* 43 */ imports.NewTable("TCEFWindowComponent_OnCanResize", 0), // event OnCanResize
			/* 44 */ imports.NewTable("TCEFWindowComponent_OnCanMaximize", 0), // event OnCanMaximize
			/* 45 */ imports.NewTable("TCEFWindowComponent_OnCanMinimize", 0), // event OnCanMinimize
			/* 46 */ imports.NewTable("TCEFWindowComponent_OnCanClose", 0), // event OnCanClose
			/* 47 */ imports.NewTable("TCEFWindowComponent_OnAccelerator", 0), // event OnAccelerator
			/* 48 */ imports.NewTable("TCEFWindowComponent_OnKeyEvent", 0), // event OnKeyEvent
		}
	})
	return cEFWindowComponentImport
}
