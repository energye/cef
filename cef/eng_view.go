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

// ICefView Parent: ICefBaseRefCounted
//
//	A View is a rectangle within the views View hierarchy. It is the base
//	interface for all Views. All size and position values are in density
//	independent pixels (DIP) unless otherwise indicated. Methods must be called
//	on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_capi.h">CEF source file: /include/capi/views/cef_view_capi.h (cef_view_t)</a>
type ICefView interface {
	ICefBaseRefCounted
	// AsBrowserView
	//  Returns this View as a BrowserView or NULL if this is not a BrowserView.
	AsBrowserView() ICefBrowserView // function
	// AsButton
	//  Returns this View as a Button or NULL if this is not a Button.
	AsButton() ICefButton // function
	// AsPanel
	//  Returns this View as a Panel or NULL if this is not a Panel.
	AsPanel() ICefPanel // function
	// AsScrollView
	//  Returns this View as a ScrollView or NULL if this is not a ScrollView.
	AsScrollView() ICefScrollView // function
	// AsTextfield
	//  Returns this View as a Textfield or NULL if this is not a Textfield.
	AsTextfield() ICefTextfield // function
	// GetTypeString
	//  Returns the type of this View as a string. Used primarily for testing
	//  purposes.
	GetTypeString() string // function
	// ToStringEx
	//  Returns a string representation of this View which includes the type and
	//  various type-specific identifying attributes. If |include_children| is
	//  true(1) any child Views will also be included. Used primarily for testing
	//  purposes.
	ToStringEx(includechildren bool) string // function
	// IsValid
	//  Returns true(1) if this View is valid.
	IsValid() bool // function
	// IsAttached
	//  Returns true(1) if this View is currently attached to another View. A
	//  View can only be attached to one View at a time.
	IsAttached() bool // function
	// IsSame
	//  Returns true(1) if this View is the same as |that| View.
	IsSame(that ICefView) bool // function
	// GetDelegate
	//  Returns the delegate associated with this View, if any.
	GetDelegate() ICefViewDelegate // function
	// GetWindow
	//  Returns the top-level Window hosting this View, if any.
	GetWindow() ICefWindow // function
	// GetID
	//  Returns the ID for this View.
	GetID() int32 // function
	// GetGroupID
	//  Returns the group id of this View, or -1 if not set.
	GetGroupID() int32 // function
	// GetParentView
	//  Returns the View that contains this View, if any.
	GetParentView() ICefView // function
	// GetViewForID
	//  Recursively descends the view tree starting at this View, and returns the
	//  first child that it encounters with the given ID. Returns NULL if no
	//  matching child view is found.
	GetViewForID(id int32) ICefView // function
	// GetBounds
	//  Returns the bounds(size and position) of this View in parent coordinates,
	//  or DIP screen coordinates if there is no parent.
	GetBounds() (resultCefRect TCefRect) // function
	// GetBoundsInScreen
	//  Returns the bounds(size and position) of this View in DIP screen
	//  coordinates.
	GetBoundsInScreen() (resultCefRect TCefRect) // function
	// GetSize
	//  Returns the size of this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetSize() (resultCefSize TCefSize) // function
	// GetPosition
	//  Returns the position of this View. Position is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetPosition() (resultCefPoint TCefPoint) // function
	// GetInsets
	//  Returns the insets for this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetInsets() (resultCefInsets TCefInsets) // function
	// GetPreferredSize
	//  Returns the size this View would like to be if enough space is available.
	//  Size is in parent coordinates, or DIP screen coordinates if there is no
	//  parent.
	GetPreferredSize() (resultCefSize TCefSize) // function
	// GetMinimumSize
	//  Returns the minimum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMinimumSize() (resultCefSize TCefSize) // function
	// GetMaximumSize
	//  Returns the maximum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMaximumSize() (resultCefSize TCefSize) // function
	// GetHeightForWidth
	//  Returns the height necessary to display this View with the provided width.
	GetHeightForWidth(width int32) int32 // function
	// IsVisible
	//  Returns whether this View is visible. A view may be visible but still not
	//  drawn in a Window if any parent views are hidden. If this View is a Window
	//  then a return value of true(1) indicates that this Window is currently
	//  visible to the user on-screen. If this View is not a Window then call
	//  is_drawn() to determine whether this View and all parent views are visible
	//  and will be drawn.
	IsVisible() bool // function
	// IsDrawn
	//  Returns whether this View is visible and drawn in a Window. A view is
	//  drawn if it and all parent views are visible. If this View is a Window
	//  then calling this function is equivalent to calling is_visible().
	//  Otherwise, to determine if the containing Window is visible to the user
	//  on-screen call is_visible() on the Window.
	IsDrawn() bool // function
	// IsEnabled
	//  Returns whether this View is enabled.
	IsEnabled() bool // function
	// IsFocusable
	//  Returns true(1) if this View is focusable, enabled and drawn.
	IsFocusable() bool // function
	// IsAccessibilityFocusable
	//  Return whether this View is focusable when the user requires full keyboard
	//  access, even though it may not be normally focusable.
	IsAccessibilityFocusable() bool // function
	// GetBackgroundColor
	//  Returns the background color for this View.
	GetBackgroundColor() TCefColor // function
	// ConvertPointToScreen
	//  Convert |point| from this View's coordinate system to DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true(1) if the conversion is successful or false(0) otherwise.
	//  Use ICefDisplay.ConvertPointToPixels() after calling this function
	//  if further conversion to display-specific pixel coordinates is desired.
	ConvertPointToScreen(point *TCefPoint) bool // function
	// ConvertPointFromScreen
	//  Convert |point| to this View's coordinate system from DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true(1) if the conversion is successful or false(0) otherwise.
	//  Use ICefDisplay.ConvertPointFromPixels() before calling this
	//  function if conversion from display-specific pixel coordinates is
	//  necessary.
	ConvertPointFromScreen(point *TCefPoint) bool // function
	// ConvertPointToWindow
	//  Convert |point| from this View's coordinate system to that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	//  (1) if the conversion is successful or false(0) otherwise.
	ConvertPointToWindow(point *TCefPoint) bool // function
	// ConvertPointFromWindow
	//  Convert |point| to this View's coordinate system from that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	//  (1) if the conversion is successful or false(0) otherwise.
	ConvertPointFromWindow(point *TCefPoint) bool // function
	// ConvertPointToView
	//  Convert |point| from this View's coordinate system to that of |view|.
	//  |view| needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true(1) if the conversion is successful or false(0)
	//  otherwise.
	ConvertPointToView(view ICefView, point *TCefPoint) bool // function
	// ConvertPointFromView
	//  Convert |point| to this View's coordinate system from that |view|. |view|
	//  needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true(1) if the conversion is successful or false(0)
	//  otherwise.
	ConvertPointFromView(view ICefView, point *TCefPoint) bool // function
	// SetID
	//  Sets the ID for this View. ID should be unique within the subtree that you
	//  intend to search for it. 0 is the default ID for views.
	SetID(id int32) // procedure
	// SetGroupID
	//  A group id is used to tag Views which are part of the same logical group.
	//  Focus can be moved between views with the same group using the arrow keys.
	//  The group id is immutable once it's set.
	SetGroupID(groupid int32) // procedure
	// SetBounds
	//  Sets the bounds(size and position) of this View. |bounds| is in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetBounds(bounds *TCefRect) // procedure
	// SetSize
	//  Sets the size of this View without changing the position. |size| in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetSize(size *TCefSize) // procedure
	// SetPosition
	//  Sets the position of this View without changing the size. |position| is in
	//  parent coordinates, or DIP screen coordinates if there is no parent.
	SetPosition(position *TCefPoint) // procedure
	// SetInsets
	//  Sets the insets for this View. |insets| is in parent coordinates, or DIP
	//  screen coordinates if there is no parent.
	SetInsets(insets *TCefInsets) // procedure
	// SizeToPreferredSize
	//  Size this View to its preferred size. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	SizeToPreferredSize() // procedure
	// InvalidateLayout
	//  Indicate that this View and all parent Views require a re-layout. This
	//  ensures the next call to layout() will propagate to this View even if the
	//  bounds of parent Views do not change.
	InvalidateLayout() // procedure
	// SetVisible
	//  Sets whether this View is visible. Windows are hidden by default and other
	//  views are visible by default. This View and any parent views must be set
	//  as visible for this View to be drawn in a Window. If this View is set as
	//  hidden then it and any child views will not be drawn and, if any of those
	//  views currently have focus, then focus will also be cleared. Painting is
	//  scheduled as needed. If this View is a Window then calling this function
	//  is equivalent to calling the Window show() and hide() functions.
	SetVisible(visible bool) // procedure
	// SetEnabled
	//  Set whether this View is enabled. A disabled View does not receive
	//  keyboard or mouse inputs. If |enabled| differs from the current value the
	//  View will be repainted. Also, clears focus if the focused View is
	//  disabled.
	SetEnabled(enabled bool) // procedure
	// SetFocusable
	//  Sets whether this View is capable of taking focus. It will clear focus if
	//  the focused View is set to be non-focusable. This is false(0) by default
	//  so that a View used as a container does not get the focus.
	SetFocusable(focusable bool) // procedure
	// RequestFocus
	//  Request keyboard focus. If this View is focusable it will become the
	//  focused View.
	RequestFocus() // procedure
	// SetBackgroundColor
	//  Sets the background color for this View.
	SetBackgroundColor(color TCefColor) // procedure
}

// TCefView Parent: TCefBaseRefCounted
//
//	A View is a rectangle within the views View hierarchy. It is the base
//	interface for all Views. All size and position values are in density
//	independent pixels (DIP) unless otherwise indicated. Methods must be called
//	on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_capi.h">CEF source file: /include/capi/views/cef_view_capi.h (cef_view_t)</a>
type TCefView struct {
	TCefBaseRefCounted
}

// ViewRef -> ICefView
var ViewRef view

// view TCefView Ref
type view uintptr

// UnWrap
//
//	Returns a ICefView instance using a PCefView data pointer.
func (m *view) UnWrap(data uintptr) ICefView {
	var resultCefView uintptr
	viewImportAPI().SysCallN(50, uintptr(data), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) AsBrowserView() ICefBrowserView {
	var resultCefBrowserView uintptr
	viewImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

func (m *TCefView) AsButton() ICefButton {
	var resultCefButton uintptr
	viewImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefButton)))
	return AsCefButton(resultCefButton)
}

func (m *TCefView) AsPanel() ICefPanel {
	var resultCefPanel uintptr
	viewImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefPanel)))
	return AsCefPanel(resultCefPanel)
}

func (m *TCefView) AsScrollView() ICefScrollView {
	var resultCefScrollView uintptr
	viewImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

func (m *TCefView) AsTextfield() ICefTextfield {
	var resultCefTextfield uintptr
	viewImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefTextfield)))
	return AsCefTextfield(resultCefTextfield)
}

func (m *TCefView) GetTypeString() string {
	r1 := viewImportAPI().SysCallN(25, m.Instance())
	return GoStr(r1)
}

func (m *TCefView) ToStringEx(includechildren bool) string {
	r1 := viewImportAPI().SysCallN(49, m.Instance(), PascalBool(includechildren))
	return GoStr(r1)
}

func (m *TCefView) IsValid() bool {
	r1 := viewImportAPI().SysCallN(35, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsAttached() bool {
	r1 := viewImportAPI().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsSame(that ICefView) bool {
	r1 := viewImportAPI().SysCallN(34, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefView) GetDelegate() ICefViewDelegate {
	var resultCefViewDelegate uintptr
	viewImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&resultCefViewDelegate)))
	return AsCefViewDelegate(resultCefViewDelegate)
}

func (m *TCefView) GetWindow() ICefWindow {
	var resultCefWindow uintptr
	viewImportAPI().SysCallN(27, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefView) GetID() int32 {
	r1 := viewImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TCefView) GetGroupID() int32 {
	r1 := viewImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TCefView) GetParentView() ICefView {
	var resultCefView uintptr
	viewImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) GetViewForID(id int32) ICefView {
	var resultCefView uintptr
	viewImportAPI().SysCallN(26, m.Instance(), uintptr(id), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) GetBounds() (resultCefRect TCefRect) {
	viewImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefView) GetBoundsInScreen() (resultCefRect TCefRect) {
	viewImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefView) GetSize() (resultCefSize TCefSize) {
	viewImportAPI().SysCallN(24, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetPosition() (resultCefPoint TCefPoint) {
	viewImportAPI().SysCallN(22, m.Instance(), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCefView) GetInsets() (resultCefInsets TCefInsets) {
	viewImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(&resultCefInsets)))
	return
}

func (m *TCefView) GetPreferredSize() (resultCefSize TCefSize) {
	viewImportAPI().SysCallN(23, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetMinimumSize() (resultCefSize TCefSize) {
	viewImportAPI().SysCallN(20, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetMaximumSize() (resultCefSize TCefSize) {
	viewImportAPI().SysCallN(19, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetHeightForWidth(width int32) int32 {
	r1 := viewImportAPI().SysCallN(16, m.Instance(), uintptr(width))
	return int32(r1)
}

func (m *TCefView) IsVisible() bool {
	r1 := viewImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsDrawn() bool {
	r1 := viewImportAPI().SysCallN(31, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsEnabled() bool {
	r1 := viewImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsFocusable() bool {
	r1 := viewImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsAccessibilityFocusable() bool {
	r1 := viewImportAPI().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) GetBackgroundColor() TCefColor {
	r1 := viewImportAPI().SysCallN(11, m.Instance())
	return TCefColor(r1)
}

func (m *TCefView) ConvertPointToScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointToWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := viewImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := viewImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCefView) SetID(id int32) {
	viewImportAPI().SysCallN(43, m.Instance(), uintptr(id))
}

func (m *TCefView) SetGroupID(groupid int32) {
	viewImportAPI().SysCallN(42, m.Instance(), uintptr(groupid))
}

func (m *TCefView) SetBounds(bounds *TCefRect) {
	viewImportAPI().SysCallN(39, m.Instance(), uintptr(unsafePointer(bounds)))
}

func (m *TCefView) SetSize(size *TCefSize) {
	viewImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefView) SetPosition(position *TCefPoint) {
	viewImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(position)))
}

func (m *TCefView) SetInsets(insets *TCefInsets) {
	viewImportAPI().SysCallN(44, m.Instance(), uintptr(unsafePointer(insets)))
}

func (m *TCefView) SizeToPreferredSize() {
	viewImportAPI().SysCallN(48, m.Instance())
}

func (m *TCefView) InvalidateLayout() {
	viewImportAPI().SysCallN(28, m.Instance())
}

func (m *TCefView) SetVisible(visible bool) {
	viewImportAPI().SysCallN(47, m.Instance(), PascalBool(visible))
}

func (m *TCefView) SetEnabled(enabled bool) {
	viewImportAPI().SysCallN(40, m.Instance(), PascalBool(enabled))
}

func (m *TCefView) SetFocusable(focusable bool) {
	viewImportAPI().SysCallN(41, m.Instance(), PascalBool(focusable))
}

func (m *TCefView) RequestFocus() {
	viewImportAPI().SysCallN(37, m.Instance())
}

func (m *TCefView) SetBackgroundColor(color TCefColor) {
	viewImportAPI().SysCallN(38, m.Instance(), uintptr(color))
}

var (
	viewImport       *imports.Imports = nil
	viewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefView_AsBrowserView", 0),
		/*1*/ imports.NewTable("CefView_AsButton", 0),
		/*2*/ imports.NewTable("CefView_AsPanel", 0),
		/*3*/ imports.NewTable("CefView_AsScrollView", 0),
		/*4*/ imports.NewTable("CefView_AsTextfield", 0),
		/*5*/ imports.NewTable("CefView_ConvertPointFromScreen", 0),
		/*6*/ imports.NewTable("CefView_ConvertPointFromView", 0),
		/*7*/ imports.NewTable("CefView_ConvertPointFromWindow", 0),
		/*8*/ imports.NewTable("CefView_ConvertPointToScreen", 0),
		/*9*/ imports.NewTable("CefView_ConvertPointToView", 0),
		/*10*/ imports.NewTable("CefView_ConvertPointToWindow", 0),
		/*11*/ imports.NewTable("CefView_GetBackgroundColor", 0),
		/*12*/ imports.NewTable("CefView_GetBounds", 0),
		/*13*/ imports.NewTable("CefView_GetBoundsInScreen", 0),
		/*14*/ imports.NewTable("CefView_GetDelegate", 0),
		/*15*/ imports.NewTable("CefView_GetGroupID", 0),
		/*16*/ imports.NewTable("CefView_GetHeightForWidth", 0),
		/*17*/ imports.NewTable("CefView_GetID", 0),
		/*18*/ imports.NewTable("CefView_GetInsets", 0),
		/*19*/ imports.NewTable("CefView_GetMaximumSize", 0),
		/*20*/ imports.NewTable("CefView_GetMinimumSize", 0),
		/*21*/ imports.NewTable("CefView_GetParentView", 0),
		/*22*/ imports.NewTable("CefView_GetPosition", 0),
		/*23*/ imports.NewTable("CefView_GetPreferredSize", 0),
		/*24*/ imports.NewTable("CefView_GetSize", 0),
		/*25*/ imports.NewTable("CefView_GetTypeString", 0),
		/*26*/ imports.NewTable("CefView_GetViewForID", 0),
		/*27*/ imports.NewTable("CefView_GetWindow", 0),
		/*28*/ imports.NewTable("CefView_InvalidateLayout", 0),
		/*29*/ imports.NewTable("CefView_IsAccessibilityFocusable", 0),
		/*30*/ imports.NewTable("CefView_IsAttached", 0),
		/*31*/ imports.NewTable("CefView_IsDrawn", 0),
		/*32*/ imports.NewTable("CefView_IsEnabled", 0),
		/*33*/ imports.NewTable("CefView_IsFocusable", 0),
		/*34*/ imports.NewTable("CefView_IsSame", 0),
		/*35*/ imports.NewTable("CefView_IsValid", 0),
		/*36*/ imports.NewTable("CefView_IsVisible", 0),
		/*37*/ imports.NewTable("CefView_RequestFocus", 0),
		/*38*/ imports.NewTable("CefView_SetBackgroundColor", 0),
		/*39*/ imports.NewTable("CefView_SetBounds", 0),
		/*40*/ imports.NewTable("CefView_SetEnabled", 0),
		/*41*/ imports.NewTable("CefView_SetFocusable", 0),
		/*42*/ imports.NewTable("CefView_SetGroupID", 0),
		/*43*/ imports.NewTable("CefView_SetID", 0),
		/*44*/ imports.NewTable("CefView_SetInsets", 0),
		/*45*/ imports.NewTable("CefView_SetPosition", 0),
		/*46*/ imports.NewTable("CefView_SetSize", 0),
		/*47*/ imports.NewTable("CefView_SetVisible", 0),
		/*48*/ imports.NewTable("CefView_SizeToPreferredSize", 0),
		/*49*/ imports.NewTable("CefView_ToStringEx", 0),
		/*50*/ imports.NewTable("CefView_UnWrap", 0),
	}
)

func viewImportAPI() *imports.Imports {
	if viewImport == nil {
		viewImport = NewDefaultImports()
		viewImport.SetImportTable(viewImportTables)
		viewImportTables = nil
	}
	return viewImport
}
