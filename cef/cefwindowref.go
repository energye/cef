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

// ICefWindow Parent: ICefPanel
type ICefWindow interface {
	ICefPanel
	// IsClosed
	//  Returns true (1) if the Window has been closed.
	IsClosed() bool // function
	// IsActive
	//  Returns whether the Window is the currently active Window.
	IsActive() bool // function
	// IsAlwaysOnTop
	//  Returns whether the Window has been set to be on top of other Windows in
	//  the Windowing system.
	IsAlwaysOnTop() bool // function
	// IsMaximized
	//  Returns true (1) if the Window is maximized.
	IsMaximized() bool // function
	// IsMinimized
	//  Returns true (1) if the Window is minimized.
	IsMinimized() bool // function
	// IsFullscreen
	//  Returns true (1) if the Window is fullscreen.
	IsFullscreen() bool // function
	// GetTitle
	//  Get the Window title.
	GetTitle() string // function
	// GetWindowIcon
	//  Get the Window icon.
	GetWindowIcon() ICefImage // function
	// GetWindowAppIcon
	//  Get the Window App icon.
	GetWindowAppIcon() ICefImage // function
	// AddOverlayView
	//  Add a View that will be overlayed on the Window contents with absolute
	//  positioning and high z-order. Positioning is controlled by |docking_mode|
	//  as described below. Setting |can_activate| to true (1) will allow the
	//  overlay view to receive input focus. The returned cef_overlay_controller_t
	//  object is used to control the overlay. Overlays are hidden by default.
	//
	//  With CEF_DOCKING_MODE_CUSTOM:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned in the top-left corner.
	//  2. Optionally change the overlay position and/or size by calling
	//  CefOverlayController methods.
	//  3. Call ICefOverlayController.SetVisible(true) to show the overlay.
	//  4. The overlay will be automatically re-sized if |view|'s layout
	//  changes. Optionally change the overlay position and/or size when
	//  OnLayoutChanged is called on the Window's delegate to indicate a
	//  change in Window bounds.
	//
	//  With other docking modes:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned based on |docking_mode|.
	//  2. Call ICefOverlayController.SetVisible(true) to show the overlay.
	//  3. The overlay will be automatically re-sized if |view|'s layout changes
	//  and re-positioned as appropriate when the Window resizes.
	//
	//  Overlays created by this function will receive a higher z-order then any
	//  child Views added previously. It is therefore recommended to call this
	//  function last after all other child Views have been added so that the
	//  overlay displays as the top-most child of the Window.
	AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode, canActivate bool) ICefOverlayController // function
	// GetDisplay
	//  Returns the Display that most closely intersects the bounds of this
	//  Window. May return NULL if this Window is not currently displayed.
	GetDisplay() ICefDisplay // function
	// GetClientAreaBoundsInScreen
	//  Returns the bounds (size and position) of this Window's client area.
	//  Position is in screen coordinates.
	GetClientAreaBoundsInScreen() TCefRect // function
	// GetWindowHandle
	//  Retrieve the platform window handle for this Window.
	GetWindowHandle() cefTypes.TCefWindowHandle // function
	// GetRuntimeStyle
	//  Returns the runtime style for this Window (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	GetRuntimeStyle() cefTypes.TCefRuntimeStyle // function
	// Show
	//  Show the Window.
	Show() // procedure
	// ShowAsBrowserModalDialog
	//  Show the Window as a browser modal dialog relative to |browser_view|. A
	//  parent Window must be returned via
	//  ICefWindowDelegate.GetParentWindow() and |browser_view| must belong
	//  to that parent Window. While this Window is visible, |browser_view| will
	//  be disabled while other controls in the parent Window remain enabled.
	//  Navigating or destroying the |browser_view| will close this Window
	//  automatically. Alternately, use show() and return true (1) from
	//  ICefWindowDelegate.IsWindowModalDialog() for a window modal dialog
	//  where all controls in the parent Window are disabled.
	ShowAsBrowserModalDialog(browserView ICefBrowserView) // procedure
	// Hide
	//  Hide the Window.
	Hide() // procedure
	// CenterWindow
	//  Sizes the Window to |size| and centers it in the current display.
	CenterWindow(size TCefSize) // procedure
	// Close
	//  Close the Window.
	Close() // procedure
	// Activate
	//  Activate the Window, assuming it already exists and is visible.
	Activate() // procedure
	// Deactivate
	//  Deactivate the Window, making the next Window in the Z order the active
	//  Window.
	Deactivate() // procedure
	// BringToTop
	//  Bring this Window to the top of other Windows in the Windowing system.
	BringToTop() // procedure
	// SetAlwaysOnTop
	//  Set the Window to be on top of other Windows in the Windowing system.
	SetAlwaysOnTop(onTop bool) // procedure
	// Maximize
	//  Maximize the Window.
	Maximize() // procedure
	// Minimize
	//  Minimize the Window.
	Minimize() // procedure
	// Restore
	//  Restore the Window.
	Restore() // procedure
	// SetFullscreen
	//  Set fullscreen Window state. The
	//  ICefWindowDelegate.OnWindowFullscreenTransition function will be
	//  called during the fullscreen transition for notification purposes.
	SetFullscreen(fullscreen bool) // procedure
	// SetTitle
	//  Set the Window title.
	SetTitle(title string) // procedure
	// SetWindowIcon
	//  Set the Window icon. This should be a 16x16 icon suitable for use in the
	//  Windows's title bar.
	SetWindowIcon(image ICefImage) // procedure
	// SetWindowAppIcon
	//  Set the Window App icon. This should be a larger icon for use in the host
	//  environment app switching UI. On Windows, this is the ICON_BIG used in
	//  Alt-Tab list and Windows taskbar. The Window icon will be used by default
	//  if no Window App icon is specified.
	SetWindowAppIcon(image ICefImage) // procedure
	// ShowMenu
	//  Show a menu with contents |menu_model|. |screen_point| specifies the menu
	//  position in screen coordinates. |anchor_position| specifies how the menu
	//  will be anchored relative to |screen_point|.
	ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) // procedure
	// CancelMenu
	//  Cancel the menu that is currently showing, if any.
	CancelMenu() // procedure
	// SetDraggableRegions
	//  Set the regions where mouse events will be intercepted by this Window to
	//  support drag operations. Call this function with an NULL vector to clear
	//  the draggable regions. The draggable region bounds should be in window
	//  coordinates.
	SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray) // procedure
	// SendKeyPress
	//  Simulate a key press. |key_code| is the VKEY_* value from Chromium's
	//  ui/events/keycodes/keyboard_codes.h header (VK_* values on Windows).
	//  |event_flags| is some combination of EVENTFLAG_SHIFT_DOWN,
	//  EVENTFLAG_CONTROL_DOWN and/or EVENTFLAG_ALT_DOWN. This function is exposed
	//  primarily for testing purposes.
	SendKeyPress(keyCode int32, eventFlags uint32) // procedure
	// SendMouseMove
	//  Simulate a mouse move. The mouse cursor will be moved to the specified
	//  (screen_x, screen_y) position. This function is exposed primarily for
	//  testing purposes.
	SendMouseMove(screenX int32, screenY int32) // procedure
	// SendMouseEvents
	//  Simulate mouse down and/or mouse up events. |button| is the mouse button
	//  type. If |mouse_down| is true (1) a mouse down event will be sent. If
	//  |mouse_up| is true (1) a mouse up event will be sent. If both are true (1)
	//  a mouse down event will be sent followed by a mouse up event (equivalent
	//  to clicking the mouse button). The events will be sent using the current
	//  cursor position so make sure to call send_mouse_move() first to position
	//  the mouse. This function is exposed primarily for testing purposes.
	SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool) // procedure
	// SetAccelerator
	//  Set the keyboard accelerator for the specified |command_id|. |key_code|
	//  can be any virtual key or character value. Required modifier keys are
	//  specified by |shift_pressed|, |ctrl_pressed| and/or |alt_pressed|.
	//  ICefWindowDelegate.OnAccelerator will be called if the keyboard
	//  combination is triggered while this window has focus.
	//  The |high_priority| value will be considered if a child ICefBrowserView
	//  has focus when the keyboard combination is triggered. If |high_priority|
	//  is true (1) then the key event will not be forwarded to the web content
	//  (`keydown` event handler) or ICefKeyboardHandler first. If
	//  |high_priority| is false (0) then the behavior will depend on the
	//  ICefBrowserView.SetPreferAccelerators configuration.
	SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool, highPriority bool) // procedure
	// RemoveAccelerator
	//  Remove the keyboard accelerator for the specified |command_id|.
	RemoveAccelerator(commandId int32) // procedure
	// RemoveAllAccelerators
	//  Remove all keyboard accelerators.
	RemoveAllAccelerators() // procedure
	// SetThemeColor
	//  Override a standard theme color or add a custom color associated with
	//  |color_id|. See cef_color_ids.h for standard ID values. Recommended usage
	//  is as follows:
	//  <code>
	//  1. Customize the default native/OS theme by calling SetThemeColor before
	//  showing the first Window. When done setting colors call
	//  ICefWindow.ThemeChanged to trigger ICefViewDelegate.OnThemeChanged
	//  notifications.
	//  2. Customize the current native/OS or Chrome theme after it changes by
	//  calling SetThemeColor from the ICefWindowDelegate.OnThemeColorsChanged
	//  callback. ICefViewDelegate.OnThemeChanged notifications will then be
	//  triggered automatically.
	//  </code>
	//  The configured color will be available immediately via
	//  ICefView.GetThemeColor and will be applied to each View in this
	//  Window's component hierarchy when ICefViewDelegate.OnThemeChanged is
	//  called. See OnThemeColorsChanged documentation for additional details.
	//  Clients wishing to add custom colors should use |color_id| values >=
	//  CEF_ChromeColorsEnd.
	SetThemeColor(colorId int32, color cefTypes.TCefColor) // procedure
	// ThemeChanged
	//  Trigger ICefViewDelegate.OnThemeChanged callbacks for each View in
	//  this Window's component hierarchy. Unlike a native/OS or Chrome theme
	//  change this function does not reset theme colors to standard values and
	//  does not result in a call to ICefWindowDelegate.OnThemeColorsChanged.
	//  Do not call this function from ICefWindowDelegate.OnThemeColorsChanged
	//  or ICefViewDelegate.OnThemeChanged.
	ThemeChanged() // procedure
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

func (m *TCefWindowRef) GetTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := cefWindowRefAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
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

func (m *TCefWindowRef) AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode, canActivate bool) (result ICefOverlayController) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(view), uintptr(dockingMode), api.PasBool(canActivate), uintptr(base.UnsafePointer(&resultPtr)))
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

func (m *TCefWindowRef) GetRuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := cefWindowRefAPI().SysCallN(14, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TCefWindowRef) Show() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(17, m.Instance())
}

func (m *TCefWindowRef) ShowAsBrowserModalDialog(browserView ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(18, m.Instance(), base.GetObjectUintptr(browserView))
}

func (m *TCefWindowRef) Hide() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(19, m.Instance())
}

func (m *TCefWindowRef) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefWindowRef) Close() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(21, m.Instance())
}

func (m *TCefWindowRef) Activate() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(22, m.Instance())
}

func (m *TCefWindowRef) Deactivate() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(23, m.Instance())
}

func (m *TCefWindowRef) BringToTop() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(24, m.Instance())
}

func (m *TCefWindowRef) SetAlwaysOnTop(onTop bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(25, m.Instance(), api.PasBool(onTop))
}

func (m *TCefWindowRef) Maximize() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(26, m.Instance())
}

func (m *TCefWindowRef) Minimize() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(27, m.Instance())
}

func (m *TCefWindowRef) Restore() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(28, m.Instance())
}

func (m *TCefWindowRef) SetFullscreen(fullscreen bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(29, m.Instance(), api.PasBool(fullscreen))
}

func (m *TCefWindowRef) SetTitle(title string) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(30, m.Instance(), api.PasStr(title))
}

func (m *TCefWindowRef) SetWindowIcon(image ICefImage) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(image))
}

func (m *TCefWindowRef) SetWindowAppIcon(image ICefImage) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(32, m.Instance(), base.GetObjectUintptr(image))
}

func (m *TCefWindowRef) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCefWindowRef) CancelMenu() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(34, m.Instance())
}

func (m *TCefWindowRef) SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(35, m.Instance(), uintptr(regionsCount), regions.Instance())
}

func (m *TCefWindowRef) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(36, m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCefWindowRef) SendMouseMove(screenX int32, screenY int32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(37, m.Instance(), uintptr(screenX), uintptr(screenY))
}

func (m *TCefWindowRef) SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(38, m.Instance(), uintptr(button), api.PasBool(mouseDown), api.PasBool(mouseUp))
}

func (m *TCefWindowRef) SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool, highPriority bool) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(39, m.Instance(), uintptr(commandId), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed), api.PasBool(highPriority))
}

func (m *TCefWindowRef) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(40, m.Instance(), uintptr(commandId))
}

func (m *TCefWindowRef) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(41, m.Instance())
}

func (m *TCefWindowRef) SetThemeColor(colorId int32, color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(42, m.Instance(), uintptr(colorId), uintptr(color))
}

func (m *TCefWindowRef) ThemeChanged() {
	if !m.IsValid() {
		return
	}
	cefWindowRefAPI().SysCallN(43, m.Instance())
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

// UnWrapWithPointer
//
//	Returns a ICefWindow instance using a PCefWindow data pointer.
func (_WindowRefClass) UnWrapWithPointer(data uintptr) (result ICefWindow) {
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(15, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

// CreateTopLevel
//
//	Create a new Window.
func (_WindowRefClass) CreateTopLevel(delegate IEngWindowDelegate) (result ICefWindow) {
	var resultPtr uintptr
	cefWindowRefAPI().SysCallN(16, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
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
			/* 14 */ imports.NewTable("TCefWindowRef_GetRuntimeStyle", 0), // function GetRuntimeStyle
			/* 15 */ imports.NewTable("TCefWindowRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 16 */ imports.NewTable("TCefWindowRef_CreateTopLevel", 0), // static function CreateTopLevel
			/* 17 */ imports.NewTable("TCefWindowRef_Show", 0), // procedure Show
			/* 18 */ imports.NewTable("TCefWindowRef_ShowAsBrowserModalDialog", 0), // procedure ShowAsBrowserModalDialog
			/* 19 */ imports.NewTable("TCefWindowRef_Hide", 0), // procedure Hide
			/* 20 */ imports.NewTable("TCefWindowRef_CenterWindow", 0), // procedure CenterWindow
			/* 21 */ imports.NewTable("TCefWindowRef_Close", 0), // procedure Close
			/* 22 */ imports.NewTable("TCefWindowRef_Activate", 0), // procedure Activate
			/* 23 */ imports.NewTable("TCefWindowRef_Deactivate", 0), // procedure Deactivate
			/* 24 */ imports.NewTable("TCefWindowRef_BringToTop", 0), // procedure BringToTop
			/* 25 */ imports.NewTable("TCefWindowRef_SetAlwaysOnTop", 0), // procedure SetAlwaysOnTop
			/* 26 */ imports.NewTable("TCefWindowRef_Maximize", 0), // procedure Maximize
			/* 27 */ imports.NewTable("TCefWindowRef_Minimize", 0), // procedure Minimize
			/* 28 */ imports.NewTable("TCefWindowRef_Restore", 0), // procedure Restore
			/* 29 */ imports.NewTable("TCefWindowRef_SetFullscreen", 0), // procedure SetFullscreen
			/* 30 */ imports.NewTable("TCefWindowRef_SetTitle", 0), // procedure SetTitle
			/* 31 */ imports.NewTable("TCefWindowRef_SetWindowIcon", 0), // procedure SetWindowIcon
			/* 32 */ imports.NewTable("TCefWindowRef_SetWindowAppIcon", 0), // procedure SetWindowAppIcon
			/* 33 */ imports.NewTable("TCefWindowRef_ShowMenu", 0), // procedure ShowMenu
			/* 34 */ imports.NewTable("TCefWindowRef_CancelMenu", 0), // procedure CancelMenu
			/* 35 */ imports.NewTable("TCefWindowRef_SetDraggableRegions", 0), // procedure SetDraggableRegions
			/* 36 */ imports.NewTable("TCefWindowRef_SendKeyPress", 0), // procedure SendKeyPress
			/* 37 */ imports.NewTable("TCefWindowRef_SendMouseMove", 0), // procedure SendMouseMove
			/* 38 */ imports.NewTable("TCefWindowRef_SendMouseEvents", 0), // procedure SendMouseEvents
			/* 39 */ imports.NewTable("TCefWindowRef_SetAccelerator", 0), // procedure SetAccelerator
			/* 40 */ imports.NewTable("TCefWindowRef_RemoveAccelerator", 0), // procedure RemoveAccelerator
			/* 41 */ imports.NewTable("TCefWindowRef_RemoveAllAccelerators", 0), // procedure RemoveAllAccelerators
			/* 42 */ imports.NewTable("TCefWindowRef_SetThemeColor", 0), // procedure SetThemeColor
			/* 43 */ imports.NewTable("TCefWindowRef_ThemeChanged", 0), // procedure ThemeChanged
		}
	})
	return cefWindowRefImport
}
