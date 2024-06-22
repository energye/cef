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

// ICefOverlayController Parent: ICefBaseRefCounted
//
//	Controller for an overlay that contains a contents View added via
//	ICefWindow.AddOverlayView. Methods exposed by this controller should be
//	called in preference to functions of the same name exposed by the contents
//	View unless otherwise indicated. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_overlay_controller_capi.h">CEF source file: /include/capi/views/cef_overlay_controller_capi.h (cef_overlay_controller_t)</a>
type ICefOverlayController interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true(1) if this object is valid.
	IsValid() bool // function
	// IsSame
	//  Returns true(1) if this object is the same as |that| object.
	IsSame(that ICefOverlayController) bool // function
	// GetContentsView
	//  Returns the contents View for this overlay.
	GetContentsView() ICefView // function
	// GetWindow
	//  Returns the top-level Window hosting this overlay. Use this function
	//  instead of calling get_window() on the contents View.
	GetWindow() ICefWindow // function
	// GetDockingMode
	//  Returns the docking mode for this overlay.
	GetDockingMode() TCefDockingMode // function
	// GetBounds
	//  Returns the bounds(size and position) of this overlay in parent
	//  coordinates.
	GetBounds() (resultCefRect TCefRect) // function
	// GetBoundsInScreen
	//  Returns the bounds(size and position) of this overlay in DIP screen
	//  coordinates.
	GetBoundsInScreen() (resultCefRect TCefRect) // function
	// GetSize
	//  Returns the size of this overlay in parent coordinates.
	GetSize() (resultCefSize TCefSize) // function
	// GetPosition
	//  Returns the position of this overlay in parent coordinates.
	GetPosition() (resultCefPoint TCefPoint) // function
	// GetInsets
	//  Returns the insets for this overlay in parent coordinates.
	GetInsets() (resultCefInsets TCefInsets) // function
	// IsVisible
	//  Returns whether this overlay is visible. A View may be visible but still
	//  not drawn in a Window if any parent Views are hidden. Call is_drawn() to
	//  determine whether this overlay and all parent Views are visible and will
	//  be drawn.
	IsVisible() bool // function
	// IsDrawn
	//  Returns whether this overlay is visible and drawn in a Window. A View is
	//  drawn if it and all parent Views are visible. To determine if the
	//  containing Window is visible to the user on-screen call is_visible() on
	//  the Window.
	IsDrawn() bool // function
	// DestroyOverlay
	//  Destroy this overlay.
	DestroyOverlay() // procedure
	// SetBounds
	//  Sets the bounds(size and position) of this overlay. This will set the
	//  bounds of the contents View to match and trigger a re-layout if necessary.
	//  |bounds| is in parent coordinates and any insets configured on this
	//  overlay will be ignored. Use this function only for overlays created with
	//  a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking modes
	//  modify the insets of this overlay and/or layout of the contents View and
	//  call size_to_preferred_size() instead to calculate the new size and re-
	//  position the overlay if necessary.
	SetBounds(bounds *TCefRect) // procedure
	// SetSize
	//  Sets the size of this overlay without changing the position. This will set
	//  the size of the contents View to match and trigger a re-layout if
	//  necessary. |size| is in parent coordinates and any insets configured on
	//  this overlay will be ignored. Use this function only for overlays created
	//  with a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking
	//  modes modify the insets of this overlay and/or layout of the contents View
	//  and call size_to_preferred_size() instead to calculate the new size and
	//  re-position the overlay if necessary.
	SetSize(size *TCefSize) // procedure
	// SetPosition
	//  Sets the position of this overlay without changing the size. |position| is
	//  in parent coordinates and any insets configured on this overlay will be
	//  ignored. Use this function only for overlays created with a docking mode
	//  value of CEF_DOCKING_MODE_CUSTOM. With other docking modes modify the
	//  insets of this overlay and/or layout of the contents View and call
	//  size_to_preferred_size() instead to calculate the new size and re-position
	//  the overlay if necessary.
	SetPosition(position *TCefPoint) // procedure
	// SetInsets
	//  Sets the insets for this overlay. |insets| is in parent coordinates. Use
	//  this function only for overlays created with a docking mode value other
	//  than CEF_DOCKING_MODE_CUSTOM.
	SetInsets(insets *TCefInsets) // procedure
	// SizeToPreferredSize
	//  Size this overlay to its preferred size and trigger a re-layout if
	//  necessary. The position of overlays created with a docking mode value of
	//  CEF_DOCKING_MODE_CUSTOM will not be modified by calling this function.
	//  With other docking modes this function may re-position the overlay if
	//  necessary to accommodate the new size and any insets configured on the
	//  contents View.
	SizeToPreferredSize() // procedure
	// SetVisible
	//  Sets whether this overlay is visible. Overlays are hidden by default. If
	//  this overlay is hidden then it and any child Views will not be drawn and,
	//  if any of those Views currently have focus, then focus will also be
	//  cleared. Painting is scheduled as needed.
	SetVisible(visible bool) // procedure
}

// TCefOverlayController Parent: TCefBaseRefCounted
//
//	Controller for an overlay that contains a contents View added via
//	ICefWindow.AddOverlayView. Methods exposed by this controller should be
//	called in preference to functions of the same name exposed by the contents
//	View unless otherwise indicated. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_overlay_controller_capi.h">CEF source file: /include/capi/views/cef_overlay_controller_capi.h (cef_overlay_controller_t)</a>
type TCefOverlayController struct {
	TCefBaseRefCounted
}

// OverlayControllerRef -> ICefOverlayController
var OverlayControllerRef overlayController

// overlayController TCefOverlayController Ref
type overlayController uintptr

// UnWrap
//
//	Returns a ICefOverlayController instance using a PCefOverlayController data pointer.
func (m *overlayController) UnWrap(data uintptr) ICefOverlayController {
	var resultCefOverlayController uintptr
	overlayControllerImportAPI().SysCallN(19, uintptr(data), uintptr(unsafePointer(&resultCefOverlayController)))
	return AsCefOverlayController(resultCefOverlayController)
}

func (m *TCefOverlayController) IsValid() bool {
	r1 := overlayControllerImportAPI().SysCallN(11, m.Instance())
	return GoBool(r1)
}

func (m *TCefOverlayController) IsSame(that ICefOverlayController) bool {
	r1 := overlayControllerImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefOverlayController) GetContentsView() ICefView {
	var resultCefView uintptr
	overlayControllerImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefOverlayController) GetWindow() ICefWindow {
	var resultCefWindow uintptr
	overlayControllerImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefOverlayController) GetDockingMode() TCefDockingMode {
	r1 := overlayControllerImportAPI().SysCallN(4, m.Instance())
	return TCefDockingMode(r1)
}

func (m *TCefOverlayController) GetBounds() (resultCefRect TCefRect) {
	overlayControllerImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefOverlayController) GetBoundsInScreen() (resultCefRect TCefRect) {
	overlayControllerImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefOverlayController) GetSize() (resultCefSize TCefSize) {
	overlayControllerImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefOverlayController) GetPosition() (resultCefPoint TCefPoint) {
	overlayControllerImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCefOverlayController) GetInsets() (resultCefInsets TCefInsets) {
	overlayControllerImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefInsets)))
	return
}

func (m *TCefOverlayController) IsVisible() bool {
	r1 := overlayControllerImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TCefOverlayController) IsDrawn() bool {
	r1 := overlayControllerImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCefOverlayController) DestroyOverlay() {
	overlayControllerImportAPI().SysCallN(0, m.Instance())
}

func (m *TCefOverlayController) SetBounds(bounds *TCefRect) {
	overlayControllerImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(bounds)))
}

func (m *TCefOverlayController) SetSize(size *TCefSize) {
	overlayControllerImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefOverlayController) SetPosition(position *TCefPoint) {
	overlayControllerImportAPI().SysCallN(15, m.Instance(), uintptr(unsafePointer(position)))
}

func (m *TCefOverlayController) SetInsets(insets *TCefInsets) {
	overlayControllerImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(insets)))
}

func (m *TCefOverlayController) SizeToPreferredSize() {
	overlayControllerImportAPI().SysCallN(18, m.Instance())
}

func (m *TCefOverlayController) SetVisible(visible bool) {
	overlayControllerImportAPI().SysCallN(17, m.Instance(), PascalBool(visible))
}

var (
	overlayControllerImport       *imports.Imports = nil
	overlayControllerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefOverlayController_DestroyOverlay", 0),
		/*1*/ imports.NewTable("CefOverlayController_GetBounds", 0),
		/*2*/ imports.NewTable("CefOverlayController_GetBoundsInScreen", 0),
		/*3*/ imports.NewTable("CefOverlayController_GetContentsView", 0),
		/*4*/ imports.NewTable("CefOverlayController_GetDockingMode", 0),
		/*5*/ imports.NewTable("CefOverlayController_GetInsets", 0),
		/*6*/ imports.NewTable("CefOverlayController_GetPosition", 0),
		/*7*/ imports.NewTable("CefOverlayController_GetSize", 0),
		/*8*/ imports.NewTable("CefOverlayController_GetWindow", 0),
		/*9*/ imports.NewTable("CefOverlayController_IsDrawn", 0),
		/*10*/ imports.NewTable("CefOverlayController_IsSame", 0),
		/*11*/ imports.NewTable("CefOverlayController_IsValid", 0),
		/*12*/ imports.NewTable("CefOverlayController_IsVisible", 0),
		/*13*/ imports.NewTable("CefOverlayController_SetBounds", 0),
		/*14*/ imports.NewTable("CefOverlayController_SetInsets", 0),
		/*15*/ imports.NewTable("CefOverlayController_SetPosition", 0),
		/*16*/ imports.NewTable("CefOverlayController_SetSize", 0),
		/*17*/ imports.NewTable("CefOverlayController_SetVisible", 0),
		/*18*/ imports.NewTable("CefOverlayController_SizeToPreferredSize", 0),
		/*19*/ imports.NewTable("CefOverlayController_UnWrap", 0),
	}
)

func overlayControllerImportAPI() *imports.Imports {
	if overlayControllerImport == nil {
		overlayControllerImport = NewDefaultImports()
		overlayControllerImport.SetImportTable(overlayControllerImportTables)
		overlayControllerImportTables = nil
	}
	return overlayControllerImport
}
