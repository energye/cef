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

// ICEFViewComponent Parent: IComponent
type ICEFViewComponent interface {
	IComponent
	// Initialized
	//  Returns true when the control is fully initialized.
	Initialized() bool // property
	// AsView
	//  Returns this control as a View.
	AsView() ICefView // property
	// AsBrowserView
	//  Returns this View as a BrowserView or NULL if this is not a BrowserView.
	AsBrowserView() ICefBrowserView // property
	// AsButton
	//  Returns this View as a Button or NULL if this is not a Button.
	AsButton() ICefButton // property
	// AsPanel
	//  Returns this View as a Panel or NULL if this is not a Panel.
	AsPanel() ICefPanel // property
	// AsScrollView
	//  Returns this View as a ScrollView or NULL if this is not a ScrollView.
	AsScrollView() ICefScrollView // property
	// AsTextfield
	//  Returns this View as a Textfield or NULL if this is not a Textfield.
	AsTextfield() ICefTextfield // property
	// ViewForID
	//  Recursively descends the view tree starting at this View, and returns the
	//  first child that it encounters with the given ID. Returns NULL if no
	//  matching child view is found.
	ViewForID(id int32) ICefView // property
	// Valid
	//  Returns true(1) if this View is valid.
	Valid() bool // property
	// Attached
	//  Returns true(1) if this View is currently attached to another View. A
	//  View can only be attached to one View at a time.
	Attached() bool // property
	// Delegate
	//  Returns the delegate associated with this View, if any.
	Delegate() ICefViewDelegate // property
	// Window
	//  Returns the top-level Window hosting this View, if any.
	Window() ICefWindow // property
	// ParentView
	//  Returns the View that contains this View, if any.
	ParentView() ICefView // property
	// BoundsInScreen
	//  Returns the bounds(size and position) of this View in DIP screen
	//  coordinates.
	BoundsInScreen() (resultCefRect TCefRect) // property
	// PreferredSize
	//  Returns the size this View would like to be if enough space is available.
	//  Size is in parent coordinates, or DIP screen coordinates if there is no
	//  parent.
	PreferredSize() (resultCefSize TCefSize) // property
	// MinimumSize
	//  Returns the minimum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	MinimumSize() (resultCefSize TCefSize) // property
	// MaximumSize
	//  Returns the maximum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	MaximumSize() (resultCefSize TCefSize) // property
	// Visible
	//  Returns whether this View is visible. A view may be visible but still not
	//  drawn in a Window if any parent views are hidden. If this View is a Window
	//  then a return value of true(1) indicates that this Window is currently
	//  visible to the user on-screen. If this View is not a Window then call
	//  is_drawn() to determine whether this View and all parent views are visible
	//  and will be drawn.
	Visible() bool // property
	// SetVisible Set Visible
	SetVisible(AValue bool) // property
	// Drawn
	//  Returns whether this View is visible and drawn in a Window. A view is
	//  drawn if it and all parent views are visible. If this View is a Window
	//  then calling this function is equivalent to calling is_visible().
	//  Otherwise, to determine if the containing Window is visible to the user
	//  on-screen call is_visible() on the Window.
	Drawn() bool // property
	// Enabled
	//  Get or set whether this View is enabled. A disabled View does not receive
	//  keyboard or mouse inputs. If |enabled| differs from the current value the
	//  View will be repainted. Also, clears focus if the focused View is
	//  disabled.
	Enabled() bool // property
	// SetEnabled Set Enabled
	SetEnabled(AValue bool) // property
	// Focusable
	//  Gets and sets whether this View is capable of taking focus. It will clear focus if
	//  the focused View is set to be non-focusable. This is false(0) by default
	//  so that a View used as a container does not get the focus.
	Focusable() bool // property
	// SetFocusable Set Focusable
	SetFocusable(AValue bool) // property
	// AccessibilityFocusable
	//  Return whether this View is focusable when the user requires full keyboard
	//  access, even though it may not be normally focusable.
	AccessibilityFocusable() bool // property
	// BackgroundColor
	//  Returns the background color for this View.
	BackgroundColor() TCefColor // property
	// SetBackgroundColor Set BackgroundColor
	SetBackgroundColor(AValue TCefColor) // property
	// ID
	//  Gets or sets the ID for this View. ID should be unique within the subtree that you
	//  intend to search for it. 0 is the default ID for views.
	ID() int32 // property
	// SetID Set ID
	SetID(AValue int32) // property
	// GroupID
	//  Returns the group id of this View, or -1 if not set.
	GroupID() int32 // property
	// SetGroupID Set GroupID
	SetGroupID(AValue int32) // property
	// Bounds
	//  Returns the bounds(size and position) of this View in parent coordinates,
	//  or DIP screen coordinates if there is no parent.
	Bounds() (resultCefRect TCefRect) // property
	// SetBounds Set Bounds
	SetBounds(AValue *TCefRect) // property
	// Size
	//  Returns the size of this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	Size() (resultCefSize TCefSize) // property
	// SetSize Set Size
	SetSize(AValue *TCefSize) // property
	// Position
	//  Returns the position of this View. Position is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	Position() (resultCefPoint TCefPoint) // property
	// SetPosition Set Position
	SetPosition(AValue *TCefPoint) // property
	// Insets
	//  Returns the insets for this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	Insets() (resultCefInsets TCefInsets) // property
	// SetInsets Set Insets
	SetInsets(AValue *TCefInsets) // property
	// TypeString
	//  Returns the type of this View as a string. Used primarily for testing
	//  purposes.
	TypeString() string // property
	// HeightForWidth
	//  Returns the height necessary to display this View with the provided width.
	HeightForWidth(width int32) int32 // property
	// ToStringEx
	//  Returns a string representation of this View which includes the type and
	//  various type-specific identifying attributes. If |include_children| is
	//  true(1) any child Views will also be included. Used primarily for testing
	//  purposes.
	ToStringEx(includechildren bool) string // function
	// IsSame
	//  Returns true(1) if this View is the same as |that| View.
	IsSame(that ICefView) bool // function
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
	// SizeToPreferredSize
	//  Size this View to its preferred size. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	SizeToPreferredSize() // procedure
	// InvalidateLayout
	//  Indicate that this View and all parent Views require a re-layout. This
	//  ensures the next call to layout() will propagate to this View even if the
	//  bounds of parent Views do not change.
	InvalidateLayout() // procedure
	// RequestFocus
	//  Request keyboard focus. If this View is focusable it will become the
	//  focused View.
	RequestFocus() // procedure
	// SetOnGetPreferredSize
	//  Return the preferred size for |view|. The Layout will use this information
	//  to determine the display size.
	SetOnGetPreferredSize(fn TOnGetPreferredSize) // property event
	// SetOnGetMinimumSize
	//  Return the minimum size for |view|.
	SetOnGetMinimumSize(fn TOnGetMinimumSize) // property event
	// SetOnGetMaximumSize
	//  Return the maximum size for |view|.
	SetOnGetMaximumSize(fn TOnGetMaximumSize) // property event
	// SetOnGetHeightForWidth
	//  Return the height necessary to display |view| with the provided |width|.
	//  If not specified the result of get_preferred_size().height will be used by
	//  default. Override if |view|'s preferred height depends upon the width(for
	//  example, with Labels).
	SetOnGetHeightForWidth(fn TOnGetHeightForWidth) // property event
	// SetOnParentViewChanged
	//  Called when the parent of |view| has changed. If |view| is being added to
	//  |parent| then |added| will be true(1). If |view| is being removed from
	//  |parent| then |added| will be false(0). If |view| is being reparented the
	//  remove notification will be sent before the add notification. Do not
	//  modify the view hierarchy in this callback.
	SetOnParentViewChanged(fn TOnParentViewChanged) // property event
	// SetOnChildViewChanged
	//  Called when a child of |view| has changed. If |child| is being added to
	//  |view| then |added| will be true(1). If |child| is being removed from
	//  |view| then |added| will be false(0). If |child| is being reparented the
	//  remove notification will be sent to the old parent before the add
	//  notification is sent to the new parent. Do not modify the view hierarchy
	//  in this callback.
	SetOnChildViewChanged(fn TOnChildViewChanged) // property event
	// SetOnWindowChanged
	//  Called when |view| is added or removed from the ICefWindow.
	SetOnWindowChanged(fn TOnWindowChanged) // property event
	// SetOnLayoutChanged
	//  Called when the layout of |view| has changed.
	SetOnLayoutChanged(fn TOnLayoutChanged) // property event
	// SetOnFocus
	//  Called when |view| gains focus.
	SetOnFocus(fn TOnFocus) // property event
	// SetOnBlur
	//  Called when |view| loses focus.
	SetOnBlur(fn TOnBlur) // property event
}

// TCEFViewComponent Parent: TComponent
type TCEFViewComponent struct {
	TComponent
	getPreferredSizePtr  uintptr
	getMinimumSizePtr    uintptr
	getMaximumSizePtr    uintptr
	getHeightForWidthPtr uintptr
	parentViewChangedPtr uintptr
	childViewChangedPtr  uintptr
	windowChangedPtr     uintptr
	layoutChangedPtr     uintptr
	focusPtr             uintptr
	blurPtr              uintptr
}

func NewCEFViewComponent(aOwner IComponent) ICEFViewComponent {
	r1 := viewComponentImportAPI().SysCallN(17, GetObjectUintptr(aOwner))
	return AsCEFViewComponent(r1)
}

func (m *TCEFViewComponent) Initialized() bool {
	r1 := viewComponentImportAPI().SysCallN(25, m.Instance())
	return GoBool(r1)
}

func (m *TCEFViewComponent) AsView() ICefView {
	var resultCefView uintptr
	viewComponentImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFViewComponent) AsBrowserView() ICefBrowserView {
	var resultCefBrowserView uintptr
	viewComponentImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

func (m *TCEFViewComponent) AsButton() ICefButton {
	var resultCefButton uintptr
	viewComponentImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefButton)))
	return AsCefButton(resultCefButton)
}

func (m *TCEFViewComponent) AsPanel() ICefPanel {
	var resultCefPanel uintptr
	viewComponentImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefPanel)))
	return AsCefPanel(resultCefPanel)
}

func (m *TCEFViewComponent) AsScrollView() ICefScrollView {
	var resultCefScrollView uintptr
	viewComponentImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

func (m *TCEFViewComponent) AsTextfield() ICefTextfield {
	var resultCefTextfield uintptr
	viewComponentImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefTextfield)))
	return AsCefTextfield(resultCefTextfield)
}

func (m *TCEFViewComponent) ViewForID(id int32) ICefView {
	var resultCefView uintptr
	viewComponentImportAPI().SysCallN(50, m.Instance(), uintptr(id), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFViewComponent) Valid() bool {
	r1 := viewComponentImportAPI().SysCallN(49, m.Instance())
	return GoBool(r1)
}

func (m *TCEFViewComponent) Attached() bool {
	r1 := viewComponentImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCEFViewComponent) Delegate() ICefViewDelegate {
	var resultCefViewDelegate uintptr
	viewComponentImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(&resultCefViewDelegate)))
	return AsCefViewDelegate(resultCefViewDelegate)
}

func (m *TCEFViewComponent) Window() ICefWindow {
	var resultCefWindow uintptr
	viewComponentImportAPI().SysCallN(52, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCEFViewComponent) ParentView() ICefView {
	var resultCefView uintptr
	viewComponentImportAPI().SysCallN(31, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFViewComponent) BoundsInScreen() (resultCefRect TCefRect) {
	viewComponentImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCEFViewComponent) PreferredSize() (resultCefSize TCefSize) {
	viewComponentImportAPI().SysCallN(33, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCEFViewComponent) MinimumSize() (resultCefSize TCefSize) {
	viewComponentImportAPI().SysCallN(30, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCEFViewComponent) MaximumSize() (resultCefSize TCefSize) {
	viewComponentImportAPI().SysCallN(29, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCEFViewComponent) Visible() bool {
	r1 := viewComponentImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFViewComponent) SetVisible(AValue bool) {
	viewComponentImportAPI().SysCallN(51, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFViewComponent) Drawn() bool {
	r1 := viewComponentImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCEFViewComponent) Enabled() bool {
	r1 := viewComponentImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFViewComponent) SetEnabled(AValue bool) {
	viewComponentImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFViewComponent) Focusable() bool {
	r1 := viewComponentImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFViewComponent) SetFocusable(AValue bool) {
	viewComponentImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFViewComponent) AccessibilityFocusable() bool {
	r1 := viewComponentImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCEFViewComponent) BackgroundColor() TCefColor {
	r1 := viewComponentImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFViewComponent) SetBackgroundColor(AValue TCefColor) {
	viewComponentImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFViewComponent) ID() int32 {
	r1 := viewComponentImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFViewComponent) SetID(AValue int32) {
	viewComponentImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFViewComponent) GroupID() int32 {
	r1 := viewComponentImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFViewComponent) SetGroupID(AValue int32) {
	viewComponentImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFViewComponent) Bounds() (resultCefRect TCefRect) {
	viewComponentImportAPI().SysCallN(9, 0, m.Instance(), uintptr(unsafePointer(&resultCefRect)), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCEFViewComponent) SetBounds(AValue *TCefRect) {
	viewComponentImportAPI().SysCallN(9, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFViewComponent) Size() (resultCefSize TCefSize) {
	viewComponentImportAPI().SysCallN(45, 0, m.Instance(), uintptr(unsafePointer(&resultCefSize)), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCEFViewComponent) SetSize(AValue *TCefSize) {
	viewComponentImportAPI().SysCallN(45, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFViewComponent) Position() (resultCefPoint TCefPoint) {
	viewComponentImportAPI().SysCallN(32, 0, m.Instance(), uintptr(unsafePointer(&resultCefPoint)), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCEFViewComponent) SetPosition(AValue *TCefPoint) {
	viewComponentImportAPI().SysCallN(32, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFViewComponent) Insets() (resultCefInsets TCefInsets) {
	viewComponentImportAPI().SysCallN(26, 0, m.Instance(), uintptr(unsafePointer(&resultCefInsets)), uintptr(unsafePointer(&resultCefInsets)))
	return
}

func (m *TCEFViewComponent) SetInsets(AValue *TCefInsets) {
	viewComponentImportAPI().SysCallN(26, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFViewComponent) TypeString() string {
	r1 := viewComponentImportAPI().SysCallN(48, m.Instance())
	return GoStr(r1)
}

func (m *TCEFViewComponent) HeightForWidth(width int32) int32 {
	r1 := viewComponentImportAPI().SysCallN(23, m.Instance(), uintptr(width))
	return int32(r1)
}

func (m *TCEFViewComponent) ToStringEx(includechildren bool) string {
	r1 := viewComponentImportAPI().SysCallN(47, m.Instance(), PascalBool(includechildren))
	return GoStr(r1)
}

func (m *TCEFViewComponent) IsSame(that ICefView) bool {
	r1 := viewComponentImportAPI().SysCallN(28, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointToScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewComponentImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewComponentImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointToWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewComponentImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := viewComponentImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := viewComponentImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := viewComponentImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCEFViewComponent) SizeToPreferredSize() {
	viewComponentImportAPI().SysCallN(46, m.Instance())
}

func (m *TCEFViewComponent) InvalidateLayout() {
	viewComponentImportAPI().SysCallN(27, m.Instance())
}

func (m *TCEFViewComponent) RequestFocus() {
	viewComponentImportAPI().SysCallN(34, m.Instance())
}

func (m *TCEFViewComponent) SetOnGetPreferredSize(fn TOnGetPreferredSize) {
	if m.getPreferredSizePtr != 0 {
		RemoveEventElement(m.getPreferredSizePtr)
	}
	m.getPreferredSizePtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(41, m.Instance(), m.getPreferredSizePtr)
}

func (m *TCEFViewComponent) SetOnGetMinimumSize(fn TOnGetMinimumSize) {
	if m.getMinimumSizePtr != 0 {
		RemoveEventElement(m.getMinimumSizePtr)
	}
	m.getMinimumSizePtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(40, m.Instance(), m.getMinimumSizePtr)
}

func (m *TCEFViewComponent) SetOnGetMaximumSize(fn TOnGetMaximumSize) {
	if m.getMaximumSizePtr != 0 {
		RemoveEventElement(m.getMaximumSizePtr)
	}
	m.getMaximumSizePtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(39, m.Instance(), m.getMaximumSizePtr)
}

func (m *TCEFViewComponent) SetOnGetHeightForWidth(fn TOnGetHeightForWidth) {
	if m.getHeightForWidthPtr != 0 {
		RemoveEventElement(m.getHeightForWidthPtr)
	}
	m.getHeightForWidthPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(38, m.Instance(), m.getHeightForWidthPtr)
}

func (m *TCEFViewComponent) SetOnParentViewChanged(fn TOnParentViewChanged) {
	if m.parentViewChangedPtr != 0 {
		RemoveEventElement(m.parentViewChangedPtr)
	}
	m.parentViewChangedPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(43, m.Instance(), m.parentViewChangedPtr)
}

func (m *TCEFViewComponent) SetOnChildViewChanged(fn TOnChildViewChanged) {
	if m.childViewChangedPtr != 0 {
		RemoveEventElement(m.childViewChangedPtr)
	}
	m.childViewChangedPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(36, m.Instance(), m.childViewChangedPtr)
}

func (m *TCEFViewComponent) SetOnWindowChanged(fn TOnWindowChanged) {
	if m.windowChangedPtr != 0 {
		RemoveEventElement(m.windowChangedPtr)
	}
	m.windowChangedPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(44, m.Instance(), m.windowChangedPtr)
}

func (m *TCEFViewComponent) SetOnLayoutChanged(fn TOnLayoutChanged) {
	if m.layoutChangedPtr != 0 {
		RemoveEventElement(m.layoutChangedPtr)
	}
	m.layoutChangedPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(42, m.Instance(), m.layoutChangedPtr)
}

func (m *TCEFViewComponent) SetOnFocus(fn TOnFocus) {
	if m.focusPtr != 0 {
		RemoveEventElement(m.focusPtr)
	}
	m.focusPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(37, m.Instance(), m.focusPtr)
}

func (m *TCEFViewComponent) SetOnBlur(fn TOnBlur) {
	if m.blurPtr != 0 {
		RemoveEventElement(m.blurPtr)
	}
	m.blurPtr = MakeEventDataPtr(fn)
	viewComponentImportAPI().SysCallN(35, m.Instance(), m.blurPtr)
}

var (
	viewComponentImport       *imports.Imports = nil
	viewComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFViewComponent_AccessibilityFocusable", 0),
		/*1*/ imports.NewTable("CEFViewComponent_AsBrowserView", 0),
		/*2*/ imports.NewTable("CEFViewComponent_AsButton", 0),
		/*3*/ imports.NewTable("CEFViewComponent_AsPanel", 0),
		/*4*/ imports.NewTable("CEFViewComponent_AsScrollView", 0),
		/*5*/ imports.NewTable("CEFViewComponent_AsTextfield", 0),
		/*6*/ imports.NewTable("CEFViewComponent_AsView", 0),
		/*7*/ imports.NewTable("CEFViewComponent_Attached", 0),
		/*8*/ imports.NewTable("CEFViewComponent_BackgroundColor", 0),
		/*9*/ imports.NewTable("CEFViewComponent_Bounds", 0),
		/*10*/ imports.NewTable("CEFViewComponent_BoundsInScreen", 0),
		/*11*/ imports.NewTable("CEFViewComponent_ConvertPointFromScreen", 0),
		/*12*/ imports.NewTable("CEFViewComponent_ConvertPointFromView", 0),
		/*13*/ imports.NewTable("CEFViewComponent_ConvertPointFromWindow", 0),
		/*14*/ imports.NewTable("CEFViewComponent_ConvertPointToScreen", 0),
		/*15*/ imports.NewTable("CEFViewComponent_ConvertPointToView", 0),
		/*16*/ imports.NewTable("CEFViewComponent_ConvertPointToWindow", 0),
		/*17*/ imports.NewTable("CEFViewComponent_Create", 0),
		/*18*/ imports.NewTable("CEFViewComponent_Delegate", 0),
		/*19*/ imports.NewTable("CEFViewComponent_Drawn", 0),
		/*20*/ imports.NewTable("CEFViewComponent_Enabled", 0),
		/*21*/ imports.NewTable("CEFViewComponent_Focusable", 0),
		/*22*/ imports.NewTable("CEFViewComponent_GroupID", 0),
		/*23*/ imports.NewTable("CEFViewComponent_HeightForWidth", 0),
		/*24*/ imports.NewTable("CEFViewComponent_ID", 0),
		/*25*/ imports.NewTable("CEFViewComponent_Initialized", 0),
		/*26*/ imports.NewTable("CEFViewComponent_Insets", 0),
		/*27*/ imports.NewTable("CEFViewComponent_InvalidateLayout", 0),
		/*28*/ imports.NewTable("CEFViewComponent_IsSame", 0),
		/*29*/ imports.NewTable("CEFViewComponent_MaximumSize", 0),
		/*30*/ imports.NewTable("CEFViewComponent_MinimumSize", 0),
		/*31*/ imports.NewTable("CEFViewComponent_ParentView", 0),
		/*32*/ imports.NewTable("CEFViewComponent_Position", 0),
		/*33*/ imports.NewTable("CEFViewComponent_PreferredSize", 0),
		/*34*/ imports.NewTable("CEFViewComponent_RequestFocus", 0),
		/*35*/ imports.NewTable("CEFViewComponent_SetOnBlur", 0),
		/*36*/ imports.NewTable("CEFViewComponent_SetOnChildViewChanged", 0),
		/*37*/ imports.NewTable("CEFViewComponent_SetOnFocus", 0),
		/*38*/ imports.NewTable("CEFViewComponent_SetOnGetHeightForWidth", 0),
		/*39*/ imports.NewTable("CEFViewComponent_SetOnGetMaximumSize", 0),
		/*40*/ imports.NewTable("CEFViewComponent_SetOnGetMinimumSize", 0),
		/*41*/ imports.NewTable("CEFViewComponent_SetOnGetPreferredSize", 0),
		/*42*/ imports.NewTable("CEFViewComponent_SetOnLayoutChanged", 0),
		/*43*/ imports.NewTable("CEFViewComponent_SetOnParentViewChanged", 0),
		/*44*/ imports.NewTable("CEFViewComponent_SetOnWindowChanged", 0),
		/*45*/ imports.NewTable("CEFViewComponent_Size", 0),
		/*46*/ imports.NewTable("CEFViewComponent_SizeToPreferredSize", 0),
		/*47*/ imports.NewTable("CEFViewComponent_ToStringEx", 0),
		/*48*/ imports.NewTable("CEFViewComponent_TypeString", 0),
		/*49*/ imports.NewTable("CEFViewComponent_Valid", 0),
		/*50*/ imports.NewTable("CEFViewComponent_ViewForID", 0),
		/*51*/ imports.NewTable("CEFViewComponent_Visible", 0),
		/*52*/ imports.NewTable("CEFViewComponent_Window", 0),
	}
)

func viewComponentImportAPI() *imports.Imports {
	if viewComponentImport == nil {
		viewComponentImport = NewDefaultImports()
		viewComponentImport.SetImportTable(viewComponentImportTables)
		viewComponentImportTables = nil
	}
	return viewComponentImport
}
