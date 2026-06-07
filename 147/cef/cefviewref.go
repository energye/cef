//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/147/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefView Parent: ICefBaseRefCounted
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
	//  true (1) any child Views will also be included. Used primarily for testing
	//  purposes.
	ToStringEx(includeChildren bool) string // function
	// IsValid
	//  Returns true (1) if this View is valid.
	IsValid() bool // function
	// IsAttached
	//  Returns true (1) if this View is currently attached to another View. A
	//  View can only be attached to one View at a time.
	IsAttached() bool // function
	// IsSame
	//  Returns true (1) if this View is the same as |that| View.
	IsSame(that ICefView) bool // function
	// GetDelegate
	//  Returns the delegate associated with this View, if any.
	GetDelegate() IEngViewDelegate // function
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
	//  Returns the bounds (size and position) of this View in parent coordinates,
	//  or DIP screen coordinates if there is no parent.
	GetBounds() TCefRect // function
	// GetBoundsInScreen
	//  Returns the bounds (size and position) of this View in DIP screen
	//  coordinates.
	GetBoundsInScreen() TCefRect // function
	// GetSize
	//  Returns the size of this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetSize() TCefSize // function
	// GetPosition
	//  Returns the position of this View. Position is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetPosition() TCefPoint // function
	// GetInsets
	//  Returns the insets for this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetInsets() TCefInsets // function
	// GetPreferredSize
	//  Returns the size this View would like to be if enough space is available.
	//  Size is in parent coordinates, or DIP screen coordinates if there is no
	//  parent.
	GetPreferredSize() TCefSize // function
	// GetMinimumSize
	//  Returns the minimum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMinimumSize() TCefSize // function
	// GetMaximumSize
	//  Returns the maximum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMaximumSize() TCefSize // function
	// GetHeightForWidth
	//  Returns the height necessary to display this View with the provided width.
	GetHeightForWidth(width int32) int32 // function
	// IsVisible
	//  Returns whether this View is visible. A view may be visible but still not
	//  drawn in a Window if any parent views are hidden. If this View is a Window
	//  then a return value of true (1) indicates that this Window is currently
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
	//  Returns true (1) if this View is focusable, enabled and drawn.
	IsFocusable() bool // function
	// IsAccessibilityFocusable
	//  Return whether this View is focusable when the user requires full keyboard
	//  access, even though it may not be normally focusable.
	IsAccessibilityFocusable() bool // function
	// HasFocus
	//  Returns true (1) if this View has focus in the context of the containing
	//  Window. Check both this function and ICefWindow.IsActive to determine
	//  global keyboard focus.
	HasFocus() bool // function
	// GetBackgroundColor
	//  Returns the background color for this View. If the background color is
	//  unset then the current `GetThemeColor(CEF_ColorPrimaryBackground)` value
	//  will be returned. If this View belongs to an overlay (created with
	//  ICefWindow.AddOverlayView), and the background color is unset, then a
	//  value of transparent (0) will be returned.
	GetBackgroundColor() cefTypes.TCefColor // function
	// GetThemeColor
	//  Returns the current theme color associated with |color_id|, or the
	//  placeholder color (red) if unset. See cef_color_ids.h for standard ID
	//  values. Standard colors can be overridden and custom colors can be added
	//  using ICefWindow.SetThemeColor.
	GetThemeColor(colorId int32) cefTypes.TCefColor // function
	// ConvertPointToScreen
	//  Convert |point| from this View's coordinate system to DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true (1) if the conversion is successful or false (0) otherwise.
	//  Use ICefDisplay.ConvertPointToPixels() after calling this function
	//  if further conversion to display-specific pixel coordinates is desired.
	ConvertPointToScreen(point *TCefPoint) bool // function
	// ConvertPointFromScreen
	//  Convert |point| to this View's coordinate system from DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true (1) if the conversion is successful or false (0) otherwise.
	//  Use ICefDisplay.ConvertPointFromPixels() before calling this
	//  function if conversion from display-specific pixel coordinates is
	//  necessary.
	ConvertPointFromScreen(point *TCefPoint) bool // function
	// ConvertPointToWindow
	//  Convert |point| from this View's coordinate system to that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	//  (1) if the conversion is successful or false (0) otherwise.
	ConvertPointToWindow(point *TCefPoint) bool // function
	// ConvertPointFromWindow
	//  Convert |point| to this View's coordinate system from that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	//  (1) if the conversion is successful or false (0) otherwise.
	ConvertPointFromWindow(point *TCefPoint) bool // function
	// ConvertPointToView
	//  Convert |point| from this View's coordinate system to that of |view|.
	//  |view| needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true (1) if the conversion is successful or false (0)
	//  otherwise.
	ConvertPointToView(view ICefView, point *TCefPoint) bool // function
	// ConvertPointFromView
	//  Convert |point| to this View's coordinate system from that |view|. |view|
	//  needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true (1) if the conversion is successful or false (0)
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
	SetGroupID(groupId int32) // procedure
	// SetBounds
	//  Sets the bounds (size and position) of this View. |bounds| is in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetBounds(bounds TCefRect) // procedure
	// SetSize
	//  Sets the size of this View without changing the position. |size| in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetSize(size TCefSize) // procedure
	// SetPosition
	//  Sets the position of this View without changing the size. |position| is in
	//  parent coordinates, or DIP screen coordinates if there is no parent.
	SetPosition(position TCefPoint) // procedure
	// SetInsets
	//  Sets the insets for this View. |insets| is in parent coordinates, or DIP
	//  screen coordinates if there is no parent.
	SetInsets(insets TCefInsets) // procedure
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
	//  the focused View is set to be non-focusable. This is false (0) by default
	//  so that a View used as a container does not get the focus.
	SetFocusable(focusable bool) // procedure
	// RequestFocus
	//  Request focus for this View in the context of the containing Window. If
	//  this View is focusable it will become the focused View. Any focus changes
	//  while a Window is not active may be applied after that Window next becomes
	//  active.
	RequestFocus() // procedure
	// SetBackgroundColor
	//  Sets the background color for this View. The background color will be
	//  automatically reset when ICefViewDelegate.OnThemeChanged is called.
	SetBackgroundColor(color cefTypes.TCefColor) // procedure
}

// ICefViewRef Parent: ICefView ICefBaseRefCountedRef
type ICefViewRef interface {
	ICefView
	ICefBaseRefCountedRef
	AsIntfView() uintptr
}

type TCefViewRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefViewRef) AsBrowserView() (result ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

func (m *TCefViewRef) AsButton() (result ICefButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefButtonRef(resultPtr)
	return
}

func (m *TCefViewRef) AsPanel() (result ICefPanel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPanelRef(resultPtr)
	return
}

func (m *TCefViewRef) AsScrollView() (result ICefScrollView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefScrollViewRef(resultPtr)
	return
}

func (m *TCefViewRef) AsTextfield() (result ICefTextfield) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTextfieldRef(resultPtr)
	return
}

func (m *TCefViewRef) GetTypeString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefViewRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefViewRef) ToStringEx(includeChildren bool) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefViewRefAPI().SysCallN(7, m.Instance(), api.PasBool(includeChildren), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefViewRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsAttached() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsSame(that ICefView) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefViewRef) GetDelegate() (result IEngViewDelegate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngViewDelegate(resultPtr)
	return
}

func (m *TCefViewRef) GetWindow() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCefViewRef) GetID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TCefViewRef) GetGroupID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(14, m.Instance())
	return int32(r)
}

func (m *TCefViewRef) GetParentView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefViewRef) GetViewForID(id int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(16, m.Instance(), uintptr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefViewRef) GetBounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetPosition() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetInsets() (result TCefInsets) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetPreferredSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetMinimumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetMaximumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefViewRef) GetHeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(25, m.Instance(), uintptr(width))
	return int32(r)
}

func (m *TCefViewRef) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(26, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsDrawn() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(27, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) IsAccessibilityFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) HasFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

func (m *TCefViewRef) GetBackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(32, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefViewRef) GetThemeColor(colorId int32) cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(33, m.Instance(), uintptr(colorId))
	return cefTypes.TCefColor(r)
}

func (m *TCefViewRef) ConvertPointToScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(34, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(35, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointToWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(36, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(37, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(38, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) SetID(id int32) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(41, m.Instance(), uintptr(id))
}

func (m *TCefViewRef) SetGroupID(groupId int32) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(42, m.Instance(), uintptr(groupId))
}

func (m *TCefViewRef) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(43, m.Instance(), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TCefViewRef) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(44, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefViewRef) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(45, m.Instance(), uintptr(base.UnsafePointer(&position)))
}

func (m *TCefViewRef) SetInsets(insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(46, m.Instance(), uintptr(base.UnsafePointer(&insets)))
}

func (m *TCefViewRef) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(47, m.Instance())
}

func (m *TCefViewRef) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(48, m.Instance())
}

func (m *TCefViewRef) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(49, m.Instance(), api.PasBool(visible))
}

func (m *TCefViewRef) SetEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(50, m.Instance(), api.PasBool(enabled))
}

func (m *TCefViewRef) SetFocusable(focusable bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(51, m.Instance(), api.PasBool(focusable))
}

func (m *TCefViewRef) RequestFocus() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(52, m.Instance())
}

func (m *TCefViewRef) SetBackgroundColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(53, m.Instance(), uintptr(color))
}

func (m *TCefViewRef) AsIntfView() uintptr {
	return m.GetIntfPointer(0)
}

// ViewRef  is static instance
var ViewRef _ViewRefClass

// _ViewRefClass is class type defined by TCefViewRef
type _ViewRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefView instance using a PCefView data pointer.
func (_ViewRefClass) UnWrapWithPointer(data uintptr) (result ICefView) {
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(40, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

// NewViewRef class constructor
func NewViewRef(data uintptr) ICefViewRef {
	var viewPtr uintptr // ICefView
	r := cefViewRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefViewRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewPtr)
	}
	return ret
}

var (
	cefViewRefOnce   base.Once
	cefViewRefImport *imports.Imports = nil
)

func cefViewRefAPI() *imports.Imports {
	cefViewRefOnce.Do(func() {
		cefViewRefImport = api.NewDefaultImports()
		cefViewRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefViewRef_Create", 0), // constructor NewViewRef
			/* 1 */ imports.NewTable("TCefViewRef_AsBrowserView", 0), // function AsBrowserView
			/* 2 */ imports.NewTable("TCefViewRef_AsButton", 0), // function AsButton
			/* 3 */ imports.NewTable("TCefViewRef_AsPanel", 0), // function AsPanel
			/* 4 */ imports.NewTable("TCefViewRef_AsScrollView", 0), // function AsScrollView
			/* 5 */ imports.NewTable("TCefViewRef_AsTextfield", 0), // function AsTextfield
			/* 6 */ imports.NewTable("TCefViewRef_GetTypeString", 0), // function GetTypeString
			/* 7 */ imports.NewTable("TCefViewRef_ToStringEx", 0), // function ToStringEx
			/* 8 */ imports.NewTable("TCefViewRef_IsValid", 0), // function IsValid
			/* 9 */ imports.NewTable("TCefViewRef_IsAttached", 0), // function IsAttached
			/* 10 */ imports.NewTable("TCefViewRef_IsSame", 0), // function IsSame
			/* 11 */ imports.NewTable("TCefViewRef_GetDelegate", 0), // function GetDelegate
			/* 12 */ imports.NewTable("TCefViewRef_GetWindow", 0), // function GetWindow
			/* 13 */ imports.NewTable("TCefViewRef_GetID", 0), // function GetID
			/* 14 */ imports.NewTable("TCefViewRef_GetGroupID", 0), // function GetGroupID
			/* 15 */ imports.NewTable("TCefViewRef_GetParentView", 0), // function GetParentView
			/* 16 */ imports.NewTable("TCefViewRef_GetViewForID", 0), // function GetViewForID
			/* 17 */ imports.NewTable("TCefViewRef_GetBounds", 0), // function GetBounds
			/* 18 */ imports.NewTable("TCefViewRef_GetBoundsInScreen", 0), // function GetBoundsInScreen
			/* 19 */ imports.NewTable("TCefViewRef_GetSize", 0), // function GetSize
			/* 20 */ imports.NewTable("TCefViewRef_GetPosition", 0), // function GetPosition
			/* 21 */ imports.NewTable("TCefViewRef_GetInsets", 0), // function GetInsets
			/* 22 */ imports.NewTable("TCefViewRef_GetPreferredSize", 0), // function GetPreferredSize
			/* 23 */ imports.NewTable("TCefViewRef_GetMinimumSize", 0), // function GetMinimumSize
			/* 24 */ imports.NewTable("TCefViewRef_GetMaximumSize", 0), // function GetMaximumSize
			/* 25 */ imports.NewTable("TCefViewRef_GetHeightForWidth", 0), // function GetHeightForWidth
			/* 26 */ imports.NewTable("TCefViewRef_IsVisible", 0), // function IsVisible
			/* 27 */ imports.NewTable("TCefViewRef_IsDrawn", 0), // function IsDrawn
			/* 28 */ imports.NewTable("TCefViewRef_IsEnabled", 0), // function IsEnabled
			/* 29 */ imports.NewTable("TCefViewRef_IsFocusable", 0), // function IsFocusable
			/* 30 */ imports.NewTable("TCefViewRef_IsAccessibilityFocusable", 0), // function IsAccessibilityFocusable
			/* 31 */ imports.NewTable("TCefViewRef_HasFocus", 0), // function HasFocus
			/* 32 */ imports.NewTable("TCefViewRef_GetBackgroundColor", 0), // function GetBackgroundColor
			/* 33 */ imports.NewTable("TCefViewRef_GetThemeColor", 0), // function GetThemeColor
			/* 34 */ imports.NewTable("TCefViewRef_ConvertPointToScreen", 0), // function ConvertPointToScreen
			/* 35 */ imports.NewTable("TCefViewRef_ConvertPointFromScreen", 0), // function ConvertPointFromScreen
			/* 36 */ imports.NewTable("TCefViewRef_ConvertPointToWindow", 0), // function ConvertPointToWindow
			/* 37 */ imports.NewTable("TCefViewRef_ConvertPointFromWindow", 0), // function ConvertPointFromWindow
			/* 38 */ imports.NewTable("TCefViewRef_ConvertPointToView", 0), // function ConvertPointToView
			/* 39 */ imports.NewTable("TCefViewRef_ConvertPointFromView", 0), // function ConvertPointFromView
			/* 40 */ imports.NewTable("TCefViewRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 41 */ imports.NewTable("TCefViewRef_SetID", 0), // procedure SetID
			/* 42 */ imports.NewTable("TCefViewRef_SetGroupID", 0), // procedure SetGroupID
			/* 43 */ imports.NewTable("TCefViewRef_SetBounds", 0), // procedure SetBounds
			/* 44 */ imports.NewTable("TCefViewRef_SetSize", 0), // procedure SetSize
			/* 45 */ imports.NewTable("TCefViewRef_SetPosition", 0), // procedure SetPosition
			/* 46 */ imports.NewTable("TCefViewRef_SetInsets", 0), // procedure SetInsets
			/* 47 */ imports.NewTable("TCefViewRef_SizeToPreferredSize", 0), // procedure SizeToPreferredSize
			/* 48 */ imports.NewTable("TCefViewRef_InvalidateLayout", 0), // procedure InvalidateLayout
			/* 49 */ imports.NewTable("TCefViewRef_SetVisible", 0), // procedure SetVisible
			/* 50 */ imports.NewTable("TCefViewRef_SetEnabled", 0), // procedure SetEnabled
			/* 51 */ imports.NewTable("TCefViewRef_SetFocusable", 0), // procedure SetFocusable
			/* 52 */ imports.NewTable("TCefViewRef_RequestFocus", 0), // procedure RequestFocus
			/* 53 */ imports.NewTable("TCefViewRef_SetBackgroundColor", 0), // procedure SetBackgroundColor
		}
	})
	return cefViewRefImport
}
