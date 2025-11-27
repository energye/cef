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

// ICEFWindowComponent Parent: ICEFPanelComponent
type ICEFWindowComponent interface {
	ICEFPanelComponent
	// AddOverlayView
	//  Add a View that will be overlayed on the Window contents with absolute
	//  positioning and high z-order. Positioning is controlled by |docking_mode|
	//  as described below. Setting |can_activate| to true (1) will allow the
	//  overlay view to receive input focus. The returned cef_overlay_controller_t
	//  object is used to control the overlay. Overlays are hidden by default.
	//  With CEF_DOCKING_MODE_CUSTOM:
	//  <code>
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned in the top-left corner.
	//  2. Optionally change the overlay position and/or size by calling
	//  CefOverlayController methods.
	//  3. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  4. The overlay will be automatically re-sized if |view|'s layout
	//  changes. Optionally change the overlay position and/or size when
	//  OnLayoutChanged is called on the Window's delegate to indicate a
	//  change in Window bounds.</code>
	//  With other docking modes:
	//  <code>
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned based on |docking_mode|.
	//  2. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  3. The overlay will be automatically re-sized if |view|'s layout changes
	//  and re-positioned as appropriate when the Window resizes.</code>
	//  Overlays created by this function will receive a higher z-order then any
	//  child Views added previously. It is therefore recommended to call this
	//  function last after all other child Views have been added so that the
	//  overlay displays as the top-most child of the Window.
	AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode, canActivate bool) ICefOverlayController // function
	// CreateTopLevelWindow
	//  Create a new Window.
	CreateTopLevelWindow() // procedure
	// Show
	//  Show the Window.
	Show() // procedure
	// ShowAsBrowserModalDialog
	//  Show the Window as a browser modal dialog relative to |browser_view|. A
	//  parent Window must be returned via
	//  ICefWindowDelegate.OnGetParentWindow and |browser_view| must belong
	//  to that parent Window. While this Window is visible, |browser_view| will
	//  be disabled while other controls in the parent Window remain enabled.
	//  Navigating or destroying the |browser_view| will close this Window
	//  automatically. Alternately, use show() and return true (1) from
	//  ICefWindowDelegate.OnIsWindowModalDialog for a window modal dialog
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
	// Maximize
	//  Maximize the Window.
	Maximize() // procedure
	// Minimize
	//  Minimize the Window.
	Minimize() // procedure
	// Restore
	//  Restore the Window.
	Restore() // procedure
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
	// Title
	//  Get the Window title.
	Title() string         // property Title Getter
	SetTitle(value string) // property Title Setter
	// WindowIcon
	//  Get the Window icon.
	WindowIcon() ICefImage         // property WindowIcon Getter
	SetWindowIcon(value ICefImage) // property WindowIcon Setter
	// WindowAppIcon
	//  Get or set the Window App icon. This should be a larger icon for use in the host
	//  environment app switching UI. On Windows, this is the ICON_BIG used in
	//  Alt-Tab list and Windows taskbar. The Window icon will be used by default
	//  if no Window App icon is specified.
	WindowAppIcon() ICefImage         // property WindowAppIcon Getter
	SetWindowAppIcon(value ICefImage) // property WindowAppIcon Setter
	// Display
	//  Returns the Display that most closely intersects the bounds of this
	//  Window. May return NULL if this Window is not currently displayed.
	Display() ICefDisplay // property Display Getter
	// ClientAreaBoundsInScreen
	//  Returns the bounds (size and position) of this Window's client area.
	//  Position is in screen coordinates.
	ClientAreaBoundsInScreen() TCefRect // property ClientAreaBoundsInScreen Getter
	// WindowHandle
	//  Retrieve the platform window handle for this Window.
	WindowHandle() cefTypes.TCefWindowHandle // property WindowHandle Getter
	// IsClosed
	//  Returns true (1) if the Window has been closed.
	IsClosed() bool // property IsClosed Getter
	// IsActive
	//  Returns whether the Window is the currently active Window.
	IsActive() bool // property IsActive Getter
	// IsAlwaysOnTop
	//  Returns whether the Window has been set to be on top of other Windows in
	//  the Windowing system.
	IsAlwaysOnTop() bool         // property IsAlwaysOnTop Getter
	SetIsAlwaysOnTop(value bool) // property IsAlwaysOnTop Setter
	// IsFullscreen
	//  Returns true (1) if the Window is fullscreen.
	IsFullscreen() bool         // property IsFullscreen Getter
	SetIsFullscreen(value bool) // property IsFullscreen Setter
	// IsMaximized
	//  Returns true (1) if the Window is maximized.
	IsMaximized() bool // property IsMaximized Getter
	// IsMinimized
	//  Returns true (1) if the Window is minimized.
	IsMinimized() bool // property IsMinimized Getter
	// RuntimeStyle
	//  Returns the runtime style for this Window (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	RuntimeStyle() cefTypes.TCefRuntimeStyle                               // property RuntimeStyle Getter
	SetOnWindowCreated(fn TOnWindowCreatedEvent)                           // property event
	SetOnWindowClosing(fn TOnWindowClosingEvent)                           // property event
	SetOnWindowDestroyed(fn TOnWindowDestroyedEvent)                       // property event
	SetOnWindowActivationChanged(fn TOnWindowActivationChangedEvent)       // property event
	SetOnWindowBoundsChanged(fn TOnWindowBoundsChangedEvent)               // property event
	SetOnWindowFullscreenTransition(fn TOnWindowFullscreenTransitionEvent) // property event
	SetOnGetParentWindow(fn TOnGetParentWindowEvent)                       // property event
	SetOnIsWindowModalDialog(fn TOnIsWindowModalDialogEvent)               // property event
	SetOnGetInitialBounds(fn TOnGetInitialBoundsEvent)                     // property event
	SetOnGetInitialShowState(fn TOnGetInitialShowStateEvent)               // property event
	SetOnIsFrameless(fn TOnIsFramelessEvent)                               // property event
	SetOnWithStandardWindowButtons(fn TOnWithStandardWindowButtonsEvent)   // property event
	SetOnGetTitlebarHeight(fn TOnGetTitlebarHeightEvent)                   // property event
	SetOnAcceptsFirstMouse(fn TOnAcceptsFirstMouseEvent)                   // property event
	SetOnCanResize(fn TOnCanResizeEvent)                                   // property event
	SetOnCanMaximize(fn TOnCanMaximizeEvent)                               // property event
	SetOnCanMinimize(fn TOnCanMinimizeEvent)                               // property event
	SetOnCanClose(fn TOnCanCloseEvent)                                     // property event
	SetOnAccelerator(fn TOnAcceleratorEvent)                               // property event
	SetOnKeyEvent(fn TOnWindowKeyEventEvent)                               // property event
	SetOnThemeColorsChanged(fn TOnThemeColorsChangedEvent)                 // property event
	SetOnGetWindowRuntimeStyle(fn TOnGetWindowRuntimeStyleEvent)           // property event
	AsIntfWindowDelegateEvents() uintptr
	AsIntfPanelDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFWindowComponent struct {
	TCEFPanelComponent
}

func (m *TCEFWindowComponent) AddOverlayView(view ICefView, dockingMode cefTypes.TCefDockingMode, canActivate bool) (result ICefOverlayController) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(view), uintptr(dockingMode), api.PasBool(canActivate), uintptr(base.UnsafePointer(&resultPtr)))
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

func (m *TCEFWindowComponent) ShowAsBrowserModalDialog(browserView ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(browserView))
}

func (m *TCEFWindowComponent) Hide() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(5, m.Instance())
}

func (m *TCEFWindowComponent) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCEFWindowComponent) Close() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(7, m.Instance())
}

func (m *TCEFWindowComponent) Activate() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(8, m.Instance())
}

func (m *TCEFWindowComponent) Deactivate() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(9, m.Instance())
}

func (m *TCEFWindowComponent) BringToTop() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(10, m.Instance())
}

func (m *TCEFWindowComponent) Maximize() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFWindowComponent) Minimize() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(12, m.Instance())
}

func (m *TCEFWindowComponent) Restore() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(13, m.Instance())
}

func (m *TCEFWindowComponent) ShowMenu(menuModel ICefMenuModel, screenPoint TCefPoint, anchorPosition cefTypes.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(menuModel), uintptr(base.UnsafePointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCEFWindowComponent) CancelMenu() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(15, m.Instance())
}

func (m *TCEFWindowComponent) SetDraggableRegions(regionsCount cefTypes.NativeUInt, regions ICefDraggableRegionArray) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(16, m.Instance(), uintptr(regionsCount), regions.Instance())
}

func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(17, m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCEFWindowComponent) SendMouseMove(screenX int32, screenY int32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(18, m.Instance(), uintptr(screenX), uintptr(screenY))
}

func (m *TCEFWindowComponent) SendMouseEvents(button cefTypes.TCefMouseButtonType, mouseDown bool, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(19, m.Instance(), uintptr(button), api.PasBool(mouseDown), api.PasBool(mouseUp))
}

func (m *TCEFWindowComponent) SetAccelerator(commandId int32, keyCode int32, shiftPressed bool, ctrlPressed bool, altPressed bool, highPriority bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(20, m.Instance(), uintptr(commandId), uintptr(keyCode), api.PasBool(shiftPressed), api.PasBool(ctrlPressed), api.PasBool(altPressed), api.PasBool(highPriority))
}

func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(21, m.Instance(), uintptr(commandId))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(22, m.Instance())
}

func (m *TCEFWindowComponent) SetThemeColor(colorId int32, color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(23, m.Instance(), uintptr(colorId), uintptr(color))
}

func (m *TCEFWindowComponent) ThemeChanged() {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(24, m.Instance())
}

func (m *TCEFWindowComponent) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := cEFWindowComponentAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCEFWindowComponent) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFWindowComponent) WindowIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(26, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) SetWindowIcon(value ICefImage) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(26, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFWindowComponent) WindowAppIcon() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(27, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) SetWindowAppIcon(value ICefImage) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(27, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFWindowComponent) Display() (result ICefDisplay) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFWindowComponentAPI().SysCallN(28, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFWindowComponent) WindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cEFWindowComponentAPI().SysCallN(30, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCEFWindowComponent) IsClosed() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(32, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) SetIsAlwaysOnTop(value bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFWindowComponent) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(34, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) SetIsFullscreen(value bool) {
	if !m.IsValid() {
		return
	}
	cEFWindowComponentAPI().SysCallN(34, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFWindowComponent) IsMaximized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(35, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) IsMinimized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFWindowComponentAPI().SysCallN(36, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFWindowComponent) RuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := cEFWindowComponentAPI().SysCallN(37, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TCEFWindowComponent) SetOnWindowCreated(fn TOnWindowCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCreatedEvent(fn)
	base.SetEvent(m, 38, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowClosing(fn TOnWindowClosingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowClosingEvent(fn)
	base.SetEvent(m, 39, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn TOnWindowDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowDestroyedEvent(fn)
	base.SetEvent(m, 40, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn TOnWindowActivationChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowActivationChangedEvent(fn)
	base.SetEvent(m, 41, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowBoundsChanged(fn TOnWindowBoundsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowBoundsChangedEvent(fn)
	base.SetEvent(m, 42, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWindowFullscreenTransition(fn TOnWindowFullscreenTransitionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowFullscreenTransitionEvent(fn)
	base.SetEvent(m, 43, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetParentWindow(fn TOnGetParentWindowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetParentWindowEvent(fn)
	base.SetEvent(m, 44, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnIsWindowModalDialog(fn TOnIsWindowModalDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsWindowModalDialogEvent(fn)
	base.SetEvent(m, 45, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn TOnGetInitialBoundsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetInitialBoundsEvent(fn)
	base.SetEvent(m, 46, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn TOnGetInitialShowStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetInitialShowStateEvent(fn)
	base.SetEvent(m, 47, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnIsFrameless(fn TOnIsFramelessEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsFramelessEvent(fn)
	base.SetEvent(m, 48, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnWithStandardWindowButtons(fn TOnWithStandardWindowButtonsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWithStandardWindowButtonsEvent(fn)
	base.SetEvent(m, 49, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetTitlebarHeight(fn TOnGetTitlebarHeightEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetTitlebarHeightEvent(fn)
	base.SetEvent(m, 50, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnAcceptsFirstMouse(fn TOnAcceptsFirstMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceptsFirstMouseEvent(fn)
	base.SetEvent(m, 51, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanResize(fn TOnCanResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanResizeEvent(fn)
	base.SetEvent(m, 52, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanMaximize(fn TOnCanMaximizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanMaximizeEvent(fn)
	base.SetEvent(m, 53, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanMinimize(fn TOnCanMinimizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanMinimizeEvent(fn)
	base.SetEvent(m, 54, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnCanClose(fn TOnCanCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanCloseEvent(fn)
	base.SetEvent(m, 55, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnAccelerator(fn TOnAcceleratorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceleratorEvent(fn)
	base.SetEvent(m, 56, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnKeyEvent(fn TOnWindowKeyEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowKeyEventEvent(fn)
	base.SetEvent(m, 57, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnThemeColorsChanged(fn TOnThemeColorsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnThemeColorsChangedEvent(fn)
	base.SetEvent(m, 58, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFWindowComponent) SetOnGetWindowRuntimeStyle(fn TOnGetWindowRuntimeStyleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetWindowRuntimeStyleEvent(fn)
	base.SetEvent(m, 59, cEFWindowComponentAPI(), api.MakeEventDataPtr(cb))
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
			/* 4 */ imports.NewTable("TCEFWindowComponent_ShowAsBrowserModalDialog", 0), // procedure ShowAsBrowserModalDialog
			/* 5 */ imports.NewTable("TCEFWindowComponent_Hide", 0), // procedure Hide
			/* 6 */ imports.NewTable("TCEFWindowComponent_CenterWindow", 0), // procedure CenterWindow
			/* 7 */ imports.NewTable("TCEFWindowComponent_Close", 0), // procedure Close
			/* 8 */ imports.NewTable("TCEFWindowComponent_Activate", 0), // procedure Activate
			/* 9 */ imports.NewTable("TCEFWindowComponent_Deactivate", 0), // procedure Deactivate
			/* 10 */ imports.NewTable("TCEFWindowComponent_BringToTop", 0), // procedure BringToTop
			/* 11 */ imports.NewTable("TCEFWindowComponent_Maximize", 0), // procedure Maximize
			/* 12 */ imports.NewTable("TCEFWindowComponent_Minimize", 0), // procedure Minimize
			/* 13 */ imports.NewTable("TCEFWindowComponent_Restore", 0), // procedure Restore
			/* 14 */ imports.NewTable("TCEFWindowComponent_ShowMenu", 0), // procedure ShowMenu
			/* 15 */ imports.NewTable("TCEFWindowComponent_CancelMenu", 0), // procedure CancelMenu
			/* 16 */ imports.NewTable("TCEFWindowComponent_SetDraggableRegions", 0), // procedure SetDraggableRegions
			/* 17 */ imports.NewTable("TCEFWindowComponent_SendKeyPress", 0), // procedure SendKeyPress
			/* 18 */ imports.NewTable("TCEFWindowComponent_SendMouseMove", 0), // procedure SendMouseMove
			/* 19 */ imports.NewTable("TCEFWindowComponent_SendMouseEvents", 0), // procedure SendMouseEvents
			/* 20 */ imports.NewTable("TCEFWindowComponent_SetAccelerator", 0), // procedure SetAccelerator
			/* 21 */ imports.NewTable("TCEFWindowComponent_RemoveAccelerator", 0), // procedure RemoveAccelerator
			/* 22 */ imports.NewTable("TCEFWindowComponent_RemoveAllAccelerators", 0), // procedure RemoveAllAccelerators
			/* 23 */ imports.NewTable("TCEFWindowComponent_SetThemeColor", 0), // procedure SetThemeColor
			/* 24 */ imports.NewTable("TCEFWindowComponent_ThemeChanged", 0), // procedure ThemeChanged
			/* 25 */ imports.NewTable("TCEFWindowComponent_Title", 0), // property Title
			/* 26 */ imports.NewTable("TCEFWindowComponent_WindowIcon", 0), // property WindowIcon
			/* 27 */ imports.NewTable("TCEFWindowComponent_WindowAppIcon", 0), // property WindowAppIcon
			/* 28 */ imports.NewTable("TCEFWindowComponent_Display", 0), // property Display
			/* 29 */ imports.NewTable("TCEFWindowComponent_ClientAreaBoundsInScreen", 0), // property ClientAreaBoundsInScreen
			/* 30 */ imports.NewTable("TCEFWindowComponent_WindowHandle", 0), // property WindowHandle
			/* 31 */ imports.NewTable("TCEFWindowComponent_IsClosed", 0), // property IsClosed
			/* 32 */ imports.NewTable("TCEFWindowComponent_IsActive", 0), // property IsActive
			/* 33 */ imports.NewTable("TCEFWindowComponent_IsAlwaysOnTop", 0), // property IsAlwaysOnTop
			/* 34 */ imports.NewTable("TCEFWindowComponent_IsFullscreen", 0), // property IsFullscreen
			/* 35 */ imports.NewTable("TCEFWindowComponent_IsMaximized", 0), // property IsMaximized
			/* 36 */ imports.NewTable("TCEFWindowComponent_IsMinimized", 0), // property IsMinimized
			/* 37 */ imports.NewTable("TCEFWindowComponent_RuntimeStyle", 0), // property RuntimeStyle
			/* 38 */ imports.NewTable("TCEFWindowComponent_OnWindowCreated", 0), // event OnWindowCreated
			/* 39 */ imports.NewTable("TCEFWindowComponent_OnWindowClosing", 0), // event OnWindowClosing
			/* 40 */ imports.NewTable("TCEFWindowComponent_OnWindowDestroyed", 0), // event OnWindowDestroyed
			/* 41 */ imports.NewTable("TCEFWindowComponent_OnWindowActivationChanged", 0), // event OnWindowActivationChanged
			/* 42 */ imports.NewTable("TCEFWindowComponent_OnWindowBoundsChanged", 0), // event OnWindowBoundsChanged
			/* 43 */ imports.NewTable("TCEFWindowComponent_OnWindowFullscreenTransition", 0), // event OnWindowFullscreenTransition
			/* 44 */ imports.NewTable("TCEFWindowComponent_OnGetParentWindow", 0), // event OnGetParentWindow
			/* 45 */ imports.NewTable("TCEFWindowComponent_OnIsWindowModalDialog", 0), // event OnIsWindowModalDialog
			/* 46 */ imports.NewTable("TCEFWindowComponent_OnGetInitialBounds", 0), // event OnGetInitialBounds
			/* 47 */ imports.NewTable("TCEFWindowComponent_OnGetInitialShowState", 0), // event OnGetInitialShowState
			/* 48 */ imports.NewTable("TCEFWindowComponent_OnIsFrameless", 0), // event OnIsFrameless
			/* 49 */ imports.NewTable("TCEFWindowComponent_OnWithStandardWindowButtons", 0), // event OnWithStandardWindowButtons
			/* 50 */ imports.NewTable("TCEFWindowComponent_OnGetTitlebarHeight", 0), // event OnGetTitlebarHeight
			/* 51 */ imports.NewTable("TCEFWindowComponent_OnAcceptsFirstMouse", 0), // event OnAcceptsFirstMouse
			/* 52 */ imports.NewTable("TCEFWindowComponent_OnCanResize", 0), // event OnCanResize
			/* 53 */ imports.NewTable("TCEFWindowComponent_OnCanMaximize", 0), // event OnCanMaximize
			/* 54 */ imports.NewTable("TCEFWindowComponent_OnCanMinimize", 0), // event OnCanMinimize
			/* 55 */ imports.NewTable("TCEFWindowComponent_OnCanClose", 0), // event OnCanClose
			/* 56 */ imports.NewTable("TCEFWindowComponent_OnAccelerator", 0), // event OnAccelerator
			/* 57 */ imports.NewTable("TCEFWindowComponent_OnKeyEvent", 0), // event OnKeyEvent
			/* 58 */ imports.NewTable("TCEFWindowComponent_OnThemeColorsChanged", 0), // event OnThemeColorsChanged
			/* 59 */ imports.NewTable("TCEFWindowComponent_OnGetWindowRuntimeStyle", 0), // event OnGetWindowRuntimeStyle
		}
	})
	return cEFWindowComponentImport
}
