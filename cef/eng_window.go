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

// ICefWindow Parent: ICefPanel
//
//	A Window is a top-level Window/widget in the Views hierarchy. By default it
//	will have a non-client area with title bar, icon and buttons that supports
//	moving and resizing. All size and position values are in density independent
//	pixels (DIP) unless otherwise indicated. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
type ICefWindow interface {
	ICefPanel
	// IsClosed
	//  Returns true(1) if the Window has been closed.
	IsClosed() bool // function
	// IsActive
	//  Returns whether the Window is the currently active Window.
	IsActive() bool // function
	// IsAlwaysOnTop
	//  Returns whether the Window has been set to be on top of other Windows in
	//  the Windowing system.
	IsAlwaysOnTop() bool // function
	// IsMaximized
	//  Returns true(1) if the Window is maximized.
	IsMaximized() bool // function
	// IsMinimized
	//  Returns true(1) if the Window is minimized.
	IsMinimized() bool // function
	// IsFullscreen
	//  Returns true(1) if the Window is fullscreen.
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
	//  as described below. The returned cef_overlay_controller_t object is used
	//  to control the overlay. Overlays are hidden by default.
	//  With CEF_DOCKING_MODE_CUSTOM:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned in the top-left corner.
	//  2. Optionally change the overlay position and/or size by calling
	//  CefOverlayController methods.
	//  3. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  4. The overlay will be automatically re-sized if |view|'s layout
	//  changes. Optionally change the overlay position and/or size when
	//  OnLayoutChanged is called on the Window's delegate to indicate a
	//  change in Window bounds.
	//  With other docking modes:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned based on |docking_mode|.
	//  2. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  3. The overlay will be automatically re-sized if |view|'s layout changes
	//  and re-positioned as appropriate when the Window resizes.
	//  Overlays created by this function will receive a higher z-order then any
	//  child Views added previously. It is therefore recommended to call this
	//  function last after all other child Views have been added so that the
	//  overlay displays as the top-most child of the Window.
	AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController // function
	// GetDisplay
	//  Returns the Display that most closely intersects the bounds of this
	//  Window. May return NULL if this Window is not currently displayed.
	GetDisplay() ICefDisplay // function
	// GetClientAreaBoundsInScreen
	//  Returns the bounds(size and position) of this Window's client area.
	//  Position is in screen coordinates.
	GetClientAreaBoundsInScreen() (resultCefRect TCefRect) // function
	// GetWindowHandle
	//  Retrieve the platform window handle for this Window.
	GetWindowHandle() TCefWindowHandle // function
	// Show
	//  Show the Window.
	Show() // procedure
	// ShowAsBrowserModalDialog
	//  Show the Window as a browser modal dialog relative to |browser_view|. A
	//  parent Window must be returned via
	//  cef_window_delegate_t::get_parent_window() and |browser_view| must belong
	//  to that parent Window. While this Window is visible, |browser_view| will
	//  be disabled while other controls in the parent Window remain enabled.
	//  Navigating or destroying the |browser_view| will close this Window
	//  automatically. Alternately, use show() and return true(1) from
	//  cef_window_delegate_t::is_window_modal_dialog() for a window modal dialog
	//  where all controls in the parent Window are disabled.
	ShowAsBrowserModalDialog(browserview ICefBrowserView) // procedure
	// Hide
	//  Hide the Window.
	Hide() // procedure
	// CenterWindow
	//  Sizes the Window to |size| and centers it in the current display.
	CenterWindow(size *TCefSize) // procedure
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
	SetAlwaysOnTop(ontop bool) // procedure
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
	ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) // procedure
	// CancelMenu
	//  Cancel the menu that is currently showing, if any.
	CancelMenu() // procedure
	// SetDraggableRegions
	//  Set the regions where mouse events will be intercepted by this Window to
	//  support drag operations. Call this function with an NULL vector to clear
	//  the draggable regions. The draggable region bounds should be in window
	//  coordinates.
	SetDraggableRegions(regionsCount NativeUInt, regions TCefDraggableRegionArray) // procedure
	// SendKeyPress
	//  Simulate a key press. |key_code| is the VKEY_* value from Chromium's
	//  ui/events/keycodes/keyboard_codes.h header(VK_* values on Windows).
	//  |event_flags| is some combination of EVENTFLAG_SHIFT_DOWN,
	//  EVENTFLAG_CONTROL_DOWN and/or EVENTFLAG_ALT_DOWN. This function is exposed
	//  primarily for testing purposes.
	SendKeyPress(keycode int32, eventflags uint32) // procedure
	// SendMouseMove
	//  Simulate a mouse move. The mouse cursor will be moved to the specified
	//  (screen_x, screen_y) position. This function is exposed primarily for
	//  testing purposes.
	SendMouseMove(screenx, screeny int32) // procedure
	// SendMouseEvents
	//  Simulate mouse down and/or mouse up events. |button| is the mouse button
	//  type. If |mouse_down| is true(1) a mouse down event will be sent. If
	//  |mouse_up| is true(1) a mouse up event will be sent. If both are true(1)
	//  a mouse down event will be sent followed by a mouse up event(equivalent
	//  to clicking the mouse button). The events will be sent using the current
	//  cursor position so make sure to call send_mouse_move() first to position
	//  the mouse. This function is exposed primarily for testing purposes.
	SendMouseEvents(button TCefMouseButtonType, mousedown, mouseup bool) // procedure
	// SetAccelerator
	//  Set the keyboard accelerator for the specified |command_id|. |key_code|
	//  can be any virtual key or character value.
	//  cef_window_delegate_t::OnAccelerator will be called if the keyboard
	//  combination is triggered while this window has focus.
	SetAccelerator(commandid, keycode int32, shiftpressed, ctrlpressed, altpressed bool) // procedure
	// RemoveAccelerator
	//  Remove the keyboard accelerator for the specified |command_id|.
	RemoveAccelerator(commandid int32) // procedure
	// RemoveAllAccelerators
	//  Remove all keyboard accelerators.
	RemoveAllAccelerators() // procedure
}

// TCefWindow Parent: TCefPanel
//
//	A Window is a top-level Window/widget in the Views hierarchy. By default it
//	will have a non-client area with title bar, icon and buttons that supports
//	moving and resizing. All size and position values are in density independent
//	pixels (DIP) unless otherwise indicated. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
type TCefWindow struct {
	TCefPanel
}

// WindowRef -> ICefWindow
var WindowRef window

// window TCefWindow Ref
type window uintptr

// UnWrap
//
//	Returns a ICefWindow instance using a PCefWindow data pointer.
func (m *window) UnWrap(data uintptr) ICefWindow {
	var resultCefWindow uintptr
	windowImportAPI().SysCallN(39, uintptr(data), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

// CreateTopLevel
//
//	Create a new Window.
func (m *window) CreateTopLevel(delegate ICefWindowDelegate) ICefWindow {
	var resultCefWindow uintptr
	windowImportAPI().SysCallN(6, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefWindow) IsClosed() bool {
	r1 := windowImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsActive() bool {
	r1 := windowImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsAlwaysOnTop() bool {
	r1 := windowImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsMaximized() bool {
	r1 := windowImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsMinimized() bool {
	r1 := windowImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsFullscreen() bool {
	r1 := windowImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) GetTitle() string {
	r1 := windowImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TCefWindow) GetWindowIcon() ICefImage {
	var resultCefImage uintptr
	windowImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefWindow) GetWindowAppIcon() ICefImage {
	var resultCefImage uintptr
	windowImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefWindow) AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController {
	var resultCefOverlayController uintptr
	windowImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(view), uintptr(dockingmode), uintptr(unsafePointer(&resultCefOverlayController)))
	return AsCefOverlayController(resultCefOverlayController)
}

func (m *TCefWindow) GetDisplay() ICefDisplay {
	var resultCefDisplay uintptr
	windowImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

func (m *TCefWindow) GetClientAreaBoundsInScreen() (resultCefRect TCefRect) {
	windowImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefWindow) GetWindowHandle() TCefWindowHandle {
	r1 := windowImportAPI().SysCallN(12, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefWindow) Show() {
	windowImportAPI().SysCallN(36, m.Instance())
}

func (m *TCefWindow) ShowAsBrowserModalDialog(browserview ICefBrowserView) {
	windowImportAPI().SysCallN(37, m.Instance(), GetObjectUintptr(browserview))
}

func (m *TCefWindow) Hide() {
	windowImportAPI().SysCallN(14, m.Instance())
}

func (m *TCefWindow) CenterWindow(size *TCefSize) {
	windowImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefWindow) Close() {
	windowImportAPI().SysCallN(5, m.Instance())
}

func (m *TCefWindow) Activate() {
	windowImportAPI().SysCallN(0, m.Instance())
}

func (m *TCefWindow) Deactivate() {
	windowImportAPI().SysCallN(7, m.Instance())
}

func (m *TCefWindow) BringToTop() {
	windowImportAPI().SysCallN(2, m.Instance())
}

func (m *TCefWindow) SetAlwaysOnTop(ontop bool) {
	windowImportAPI().SysCallN(30, m.Instance(), PascalBool(ontop))
}

func (m *TCefWindow) Maximize() {
	windowImportAPI().SysCallN(21, m.Instance())
}

func (m *TCefWindow) Minimize() {
	windowImportAPI().SysCallN(22, m.Instance())
}

func (m *TCefWindow) Restore() {
	windowImportAPI().SysCallN(25, m.Instance())
}

func (m *TCefWindow) SetFullscreen(fullscreen bool) {
	windowImportAPI().SysCallN(32, m.Instance(), PascalBool(fullscreen))
}

func (m *TCefWindow) SetTitle(title string) {
	windowImportAPI().SysCallN(33, m.Instance(), PascalStr(title))
}

func (m *TCefWindow) SetWindowIcon(image ICefImage) {
	windowImportAPI().SysCallN(35, m.Instance(), GetObjectUintptr(image))
}

func (m *TCefWindow) SetWindowAppIcon(image ICefImage) {
	windowImportAPI().SysCallN(34, m.Instance(), GetObjectUintptr(image))
}

func (m *TCefWindow) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	windowImportAPI().SysCallN(38, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCefWindow) CancelMenu() {
	windowImportAPI().SysCallN(3, m.Instance())
}

func (m *TCefWindow) SetDraggableRegions(regionsCount NativeUInt, regions TCefDraggableRegionArray) {
	windowImportAPI().SysCallN(31, m.Instance(), uintptr(regionsCount), uintptr(unsafePointer(&regions[0])))
}

func (m *TCefWindow) SendKeyPress(keycode int32, eventflags uint32) {
	windowImportAPI().SysCallN(26, m.Instance(), uintptr(keycode), uintptr(eventflags))
}

func (m *TCefWindow) SendMouseMove(screenx, screeny int32) {
	windowImportAPI().SysCallN(28, m.Instance(), uintptr(screenx), uintptr(screeny))
}

func (m *TCefWindow) SendMouseEvents(button TCefMouseButtonType, mousedown, mouseup bool) {
	windowImportAPI().SysCallN(27, m.Instance(), uintptr(button), PascalBool(mousedown), PascalBool(mouseup))
}

func (m *TCefWindow) SetAccelerator(commandid, keycode int32, shiftpressed, ctrlpressed, altpressed bool) {
	windowImportAPI().SysCallN(29, m.Instance(), uintptr(commandid), uintptr(keycode), PascalBool(shiftpressed), PascalBool(ctrlpressed), PascalBool(altpressed))
}

func (m *TCefWindow) RemoveAccelerator(commandid int32) {
	windowImportAPI().SysCallN(23, m.Instance(), uintptr(commandid))
}

func (m *TCefWindow) RemoveAllAccelerators() {
	windowImportAPI().SysCallN(24, m.Instance())
}

var (
	windowImport       *imports.Imports = nil
	windowImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefWindow_Activate", 0),
		/*1*/ imports.NewTable("CefWindow_AddOverlayView", 0),
		/*2*/ imports.NewTable("CefWindow_BringToTop", 0),
		/*3*/ imports.NewTable("CefWindow_CancelMenu", 0),
		/*4*/ imports.NewTable("CefWindow_CenterWindow", 0),
		/*5*/ imports.NewTable("CefWindow_Close", 0),
		/*6*/ imports.NewTable("CefWindow_CreateTopLevel", 0),
		/*7*/ imports.NewTable("CefWindow_Deactivate", 0),
		/*8*/ imports.NewTable("CefWindow_GetClientAreaBoundsInScreen", 0),
		/*9*/ imports.NewTable("CefWindow_GetDisplay", 0),
		/*10*/ imports.NewTable("CefWindow_GetTitle", 0),
		/*11*/ imports.NewTable("CefWindow_GetWindowAppIcon", 0),
		/*12*/ imports.NewTable("CefWindow_GetWindowHandle", 0),
		/*13*/ imports.NewTable("CefWindow_GetWindowIcon", 0),
		/*14*/ imports.NewTable("CefWindow_Hide", 0),
		/*15*/ imports.NewTable("CefWindow_IsActive", 0),
		/*16*/ imports.NewTable("CefWindow_IsAlwaysOnTop", 0),
		/*17*/ imports.NewTable("CefWindow_IsClosed", 0),
		/*18*/ imports.NewTable("CefWindow_IsFullscreen", 0),
		/*19*/ imports.NewTable("CefWindow_IsMaximized", 0),
		/*20*/ imports.NewTable("CefWindow_IsMinimized", 0),
		/*21*/ imports.NewTable("CefWindow_Maximize", 0),
		/*22*/ imports.NewTable("CefWindow_Minimize", 0),
		/*23*/ imports.NewTable("CefWindow_RemoveAccelerator", 0),
		/*24*/ imports.NewTable("CefWindow_RemoveAllAccelerators", 0),
		/*25*/ imports.NewTable("CefWindow_Restore", 0),
		/*26*/ imports.NewTable("CefWindow_SendKeyPress", 0),
		/*27*/ imports.NewTable("CefWindow_SendMouseEvents", 0),
		/*28*/ imports.NewTable("CefWindow_SendMouseMove", 0),
		/*29*/ imports.NewTable("CefWindow_SetAccelerator", 0),
		/*30*/ imports.NewTable("CefWindow_SetAlwaysOnTop", 0),
		/*31*/ imports.NewTable("CefWindow_SetDraggableRegions", 0),
		/*32*/ imports.NewTable("CefWindow_SetFullscreen", 0),
		/*33*/ imports.NewTable("CefWindow_SetTitle", 0),
		/*34*/ imports.NewTable("CefWindow_SetWindowAppIcon", 0),
		/*35*/ imports.NewTable("CefWindow_SetWindowIcon", 0),
		/*36*/ imports.NewTable("CefWindow_Show", 0),
		/*37*/ imports.NewTable("CefWindow_ShowAsBrowserModalDialog", 0),
		/*38*/ imports.NewTable("CefWindow_ShowMenu", 0),
		/*39*/ imports.NewTable("CefWindow_UnWrap", 0),
	}
)

func windowImportAPI() *imports.Imports {
	if windowImport == nil {
		windowImport = NewDefaultImports()
		windowImport.SetImportTable(windowImportTables)
		windowImportTables = nil
	}
	return windowImport
}
