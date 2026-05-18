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

// ICefViewDelegateEvents Parent: IObject
type ICefViewDelegateEvents interface {
	IObject
}

// ICEFViewComponent Parent: ICefViewDelegateEvents IComponent
type ICEFViewComponent interface {
	ICefViewDelegateEvents
	IComponent
	// ToStringEx
	//  Returns a string representation of this View which includes the type and
	//  various type-specific identifying attributes. If |include_children| is
	//  true (1) any child Views will also be included. Used primarily for testing
	//  purposes.
	ToStringEx(includeChildren bool) string // function
	// IsSame
	//  Returns true (1) if this View is the same as |that| View.
	IsSame(that ICefView) bool // function
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
	// Initialized
	//  Returns true when the control is fully initialized.
	Initialized() bool // property Initialized Getter
	// AsView
	//  Returns this control as a View.
	AsView() ICefView // property AsView Getter
	// AsBrowserView
	//  Returns this View as a BrowserView or NULL if this is not a BrowserView.
	AsBrowserView() ICefBrowserView // property AsBrowserView Getter
	// AsButton
	//  Returns this View as a Button or NULL if this is not a Button.
	AsButton() ICefButton // property AsButton Getter
	// AsPanel
	//  Returns this View as a Panel or NULL if this is not a Panel.
	AsPanel() ICefPanel // property AsPanel Getter
	// AsScrollView
	//  Returns this View as a ScrollView or NULL if this is not a ScrollView.
	AsScrollView() ICefScrollView // property AsScrollView Getter
	// AsTextfield
	//  Returns this View as a Textfield or NULL if this is not a Textfield.
	AsTextfield() ICefTextfield // property AsTextfield Getter
	// ViewForID
	//  Recursively descends the view tree starting at this View, and returns the
	//  first child that it encounters with the given ID. Returns NULL if no
	//  matching child view is found.
	ViewForID(id int32) ICefView // property ViewForID Getter
	// Valid
	//  Returns true (1) if this View is valid.
	Valid() bool // property Valid Getter
	// Attached
	//  Returns true (1) if this View is currently attached to another View. A
	//  View can only be attached to one View at a time.
	Attached() bool // property Attached Getter
	// Delegate
	//  Returns the delegate associated with this View, if any.
	Delegate() IEngViewDelegate // property Delegate Getter
	// Window
	//  Returns the top-level Window hosting this View, if any.
	Window() ICefWindow // property Window Getter
	// ParentView
	//  Returns the View that contains this View, if any.
	ParentView() ICefView // property ParentView Getter
	// BoundsInScreen
	//  Returns the bounds (size and position) of this View in DIP screen
	//  coordinates.
	BoundsInScreen() TCefRect // property BoundsInScreen Getter
	// PreferredSize
	//  Returns the size this View would like to be if enough space is available.
	//  Size is in parent coordinates, or DIP screen coordinates if there is no
	//  parent.
	PreferredSize() TCefSize // property PreferredSize Getter
	// MinimumSize
	//  Returns the minimum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	MinimumSize() TCefSize // property MinimumSize Getter
	// MaximumSize
	//  Returns the maximum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	MaximumSize() TCefSize // property MaximumSize Getter
	// Visible
	//  Returns whether this View is visible. A view may be visible but still not
	//  drawn in a Window if any parent views are hidden. If this View is a Window
	//  then a return value of true (1) indicates that this Window is currently
	//  visible to the user on-screen. If this View is not a Window then call
	//  is_drawn() to determine whether this View and all parent views are visible
	//  and will be drawn.
	Visible() bool         // property Visible Getter
	SetVisible(value bool) // property Visible Setter
	// Drawn
	//  Returns whether this View is visible and drawn in a Window. A view is
	//  drawn if it and all parent views are visible. If this View is a Window
	//  then calling this function is equivalent to calling is_visible().
	//  Otherwise, to determine if the containing Window is visible to the user
	//  on-screen call is_visible() on the Window.
	Drawn() bool // property Drawn Getter
	// Enabled
	//  Get or set whether this View is enabled. A disabled View does not receive
	//  keyboard or mouse inputs. If |enabled| differs from the current value the
	//  View will be repainted. Also, clears focus if the focused View is
	//  disabled.
	Enabled() bool         // property Enabled Getter
	SetEnabled(value bool) // property Enabled Setter
	// Focusable
	//  Gets and sets whether this View is capable of taking focus. It will clear focus if
	//  the focused View is set to be non-focusable. This is false (0) by default
	//  so that a View used as a container does not get the focus.
	Focusable() bool         // property Focusable Getter
	SetFocusable(value bool) // property Focusable Setter
	// AccessibilityFocusable
	//  Return whether this View is focusable when the user requires full keyboard
	//  access, even though it may not be normally focusable.
	AccessibilityFocusable() bool // property AccessibilityFocusable Getter
	// BackgroundColor
	//  Returns the background color for this View. If the background color is
	//  unset then the current `GetThemeColor(CEF_ColorPrimaryBackground)` value
	//  will be returned. If this View belongs to an overlay (created with
	//  ICefWindow.AddOverlayView), and the background color is unset, then a
	//  value of transparent (0) will be returned.
	BackgroundColor() cefTypes.TCefColor         // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor) // property BackgroundColor Setter
	// ID
	//  Gets or sets the ID for this View. ID should be unique within the subtree that you
	//  intend to search for it. 0 is the default ID for views.
	ID() int32         // property ID Getter
	SetID(value int32) // property ID Setter
	// GroupID
	//  Returns the group id of this View, or -1 if not set.
	GroupID() int32         // property GroupID Getter
	SetGroupID(value int32) // property GroupID Setter
	// Bounds
	//  Returns the bounds (size and position) of this View in parent coordinates,
	//  or DIP screen coordinates if there is no parent.
	Bounds() TCefRect         // property Bounds Getter
	SetBounds(value TCefRect) // property Bounds Setter
	// Size
	//  Returns the size of this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	Size() TCefSize         // property Size Getter
	SetSize(value TCefSize) // property Size Setter
	// Position
	//  Returns the position of this View. Position is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	Position() TCefPoint         // property Position Getter
	SetPosition(value TCefPoint) // property Position Setter
	// Insets
	//  Returns the insets for this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	Insets() TCefInsets         // property Insets Getter
	SetInsets(value TCefInsets) // property Insets Setter
	// TypeString
	//  Returns the type of this View as a string. Used primarily for testing
	//  purposes.
	TypeString() string // property TypeString Getter
	// HeightForWidth
	//  Returns the height necessary to display this View with the provided width.
	HeightForWidth(width int32) int32                    // property HeightForWidth Getter
	SetOnGetPreferredSize(fn TOnGetPreferredSizeEvent)   // property event
	SetOnGetMinimumSize(fn TOnGetMinimumSizeEvent)       // property event
	SetOnGetMaximumSize(fn TOnGetMaximumSizeEvent)       // property event
	SetOnGetHeightForWidth(fn TOnGetHeightForWidthEvent) // property event
	SetOnParentViewChanged(fn TOnParentViewChangedEvent) // property event
	SetOnChildViewChanged(fn TOnChildViewChangedEvent)   // property event
	SetOnWindowChanged(fn TOnWindowChangedEvent)         // property event
	SetOnLayoutChanged(fn TOnLayoutChangedEvent)         // property event
	SetOnFocus(fn TOnFocusEvent)                         // property event
	SetOnBlur(fn TOnBlurEvent)                           // property event
	SetOnThemeChanged(fn TOnThemeChangedEvent)           // property event
	AsIntfViewDelegateEvents() uintptr
}

type TCEFViewComponent struct {
	TComponent
}

func (m *TCEFViewComponent) ToStringEx(includeChildren bool) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFViewComponentAPI().SysCallN(1, m.Instance(), api.PasBool(includeChildren), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFViewComponent) IsSame(that ICefView) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) GetThemeColor(colorId int32) cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(3, m.Instance(), uintptr(colorId))
	return cefTypes.TCefColor(r)
}

func (m *TCEFViewComponent) ConvertPointToScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointToWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(10, m.Instance())
}

func (m *TCEFViewComponent) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFViewComponent) RequestFocus() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(12, m.Instance())
}

func (m *TCEFViewComponent) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) AsView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsBrowserView() (result ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsButton() (result ICefButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefButtonRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsPanel() (result ICefPanel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPanelRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsScrollView() (result ICefScrollView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefScrollViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsTextfield() (result ICefTextfield) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTextfieldRef(resultPtr)
	return
}

func (m *TCEFViewComponent) ViewForID(id int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(20, m.Instance(), uintptr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) Valid() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Attached() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(22, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Delegate() (result IEngViewDelegate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngViewDelegate(resultPtr)
	return
}

func (m *TCEFViewComponent) Window() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCEFViewComponent) ParentView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) BoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(26, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) PreferredSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(27, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) MinimumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(28, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) MaximumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(30, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(30, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) Drawn() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) Focusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetFocusable(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) AccessibilityFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(34, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) BackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(35, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCEFViewComponent) SetBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) ID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(36, 0, m.Instance())
	return int32(r)
}

func (m *TCEFViewComponent) SetID(value int32) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) GroupID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(37, 0, m.Instance())
	return int32(r)
}

func (m *TCEFViewComponent) SetGroupID(value int32) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) Bounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(38, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetBounds(value TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(38, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) Size() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(39, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetSize(value TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(39, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) Position() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(40, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetPosition(value TCefPoint) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(40, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) Insets() (result TCefInsets) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(41, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetInsets(value TCefInsets) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(41, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) TypeString() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFViewComponentAPI().SysCallN(42, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFViewComponent) HeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(43, m.Instance(), uintptr(width))
	return int32(r)
}

func (m *TCEFViewComponent) SetOnGetPreferredSize(fn TOnGetPreferredSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetPreferredSizeEvent(fn)
	base.SetEvent(m, 44, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetMinimumSize(fn TOnGetMinimumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetMinimumSizeEvent(fn)
	base.SetEvent(m, 45, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetMaximumSize(fn TOnGetMaximumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetMaximumSizeEvent(fn)
	base.SetEvent(m, 46, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetHeightForWidth(fn TOnGetHeightForWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetHeightForWidthEvent(fn)
	base.SetEvent(m, 47, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnParentViewChanged(fn TOnParentViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnParentViewChangedEvent(fn)
	base.SetEvent(m, 48, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnChildViewChanged(fn TOnChildViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnChildViewChangedEvent(fn)
	base.SetEvent(m, 49, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnWindowChanged(fn TOnWindowChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowChangedEvent(fn)
	base.SetEvent(m, 50, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnLayoutChanged(fn TOnLayoutChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLayoutChangedEvent(fn)
	base.SetEvent(m, 51, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnFocus(fn TOnFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusEvent(fn)
	base.SetEvent(m, 52, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnBlur(fn TOnBlurEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBlurEvent(fn)
	base.SetEvent(m, 53, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnThemeChanged(fn TOnThemeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnThemeChangedEvent(fn)
	base.SetEvent(m, 54, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewViewComponent class constructor
func NewViewComponent(owner lcl.IComponent) ICEFViewComponent {
	var viewDelegateEventsPtr uintptr // ICefViewDelegateEvents
	r := cEFViewComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFViewComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFViewComponentOnce   base.Once
	cEFViewComponentImport *imports.Imports = nil
)

func cEFViewComponentAPI() *imports.Imports {
	cEFViewComponentOnce.Do(func() {
		cEFViewComponentImport = api.NewDefaultImports()
		cEFViewComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFViewComponent_Create", 0), // constructor NewViewComponent
			/* 1 */ imports.NewTable("TCEFViewComponent_ToStringEx", 0), // function ToStringEx
			/* 2 */ imports.NewTable("TCEFViewComponent_IsSame", 0), // function IsSame
			/* 3 */ imports.NewTable("TCEFViewComponent_GetThemeColor", 0), // function GetThemeColor
			/* 4 */ imports.NewTable("TCEFViewComponent_ConvertPointToScreen", 0), // function ConvertPointToScreen
			/* 5 */ imports.NewTable("TCEFViewComponent_ConvertPointFromScreen", 0), // function ConvertPointFromScreen
			/* 6 */ imports.NewTable("TCEFViewComponent_ConvertPointToWindow", 0), // function ConvertPointToWindow
			/* 7 */ imports.NewTable("TCEFViewComponent_ConvertPointFromWindow", 0), // function ConvertPointFromWindow
			/* 8 */ imports.NewTable("TCEFViewComponent_ConvertPointToView", 0), // function ConvertPointToView
			/* 9 */ imports.NewTable("TCEFViewComponent_ConvertPointFromView", 0), // function ConvertPointFromView
			/* 10 */ imports.NewTable("TCEFViewComponent_SizeToPreferredSize", 0), // procedure SizeToPreferredSize
			/* 11 */ imports.NewTable("TCEFViewComponent_InvalidateLayout", 0), // procedure InvalidateLayout
			/* 12 */ imports.NewTable("TCEFViewComponent_RequestFocus", 0), // procedure RequestFocus
			/* 13 */ imports.NewTable("TCEFViewComponent_Initialized", 0), // property Initialized
			/* 14 */ imports.NewTable("TCEFViewComponent_AsView", 0), // property AsView
			/* 15 */ imports.NewTable("TCEFViewComponent_AsBrowserView", 0), // property AsBrowserView
			/* 16 */ imports.NewTable("TCEFViewComponent_AsButton", 0), // property AsButton
			/* 17 */ imports.NewTable("TCEFViewComponent_AsPanel", 0), // property AsPanel
			/* 18 */ imports.NewTable("TCEFViewComponent_AsScrollView", 0), // property AsScrollView
			/* 19 */ imports.NewTable("TCEFViewComponent_AsTextfield", 0), // property AsTextfield
			/* 20 */ imports.NewTable("TCEFViewComponent_ViewForID", 0), // property ViewForID
			/* 21 */ imports.NewTable("TCEFViewComponent_Valid", 0), // property Valid
			/* 22 */ imports.NewTable("TCEFViewComponent_Attached", 0), // property Attached
			/* 23 */ imports.NewTable("TCEFViewComponent_Delegate", 0), // property Delegate
			/* 24 */ imports.NewTable("TCEFViewComponent_Window", 0), // property Window
			/* 25 */ imports.NewTable("TCEFViewComponent_ParentView", 0), // property ParentView
			/* 26 */ imports.NewTable("TCEFViewComponent_BoundsInScreen", 0), // property BoundsInScreen
			/* 27 */ imports.NewTable("TCEFViewComponent_PreferredSize", 0), // property PreferredSize
			/* 28 */ imports.NewTable("TCEFViewComponent_MinimumSize", 0), // property MinimumSize
			/* 29 */ imports.NewTable("TCEFViewComponent_MaximumSize", 0), // property MaximumSize
			/* 30 */ imports.NewTable("TCEFViewComponent_Visible", 0), // property Visible
			/* 31 */ imports.NewTable("TCEFViewComponent_Drawn", 0), // property Drawn
			/* 32 */ imports.NewTable("TCEFViewComponent_Enabled", 0), // property Enabled
			/* 33 */ imports.NewTable("TCEFViewComponent_Focusable", 0), // property Focusable
			/* 34 */ imports.NewTable("TCEFViewComponent_AccessibilityFocusable", 0), // property AccessibilityFocusable
			/* 35 */ imports.NewTable("TCEFViewComponent_BackgroundColor", 0), // property BackgroundColor
			/* 36 */ imports.NewTable("TCEFViewComponent_ID", 0), // property ID
			/* 37 */ imports.NewTable("TCEFViewComponent_GroupID", 0), // property GroupID
			/* 38 */ imports.NewTable("TCEFViewComponent_Bounds", 0), // property Bounds
			/* 39 */ imports.NewTable("TCEFViewComponent_Size", 0), // property Size
			/* 40 */ imports.NewTable("TCEFViewComponent_Position", 0), // property Position
			/* 41 */ imports.NewTable("TCEFViewComponent_Insets", 0), // property Insets
			/* 42 */ imports.NewTable("TCEFViewComponent_TypeString", 0), // property TypeString
			/* 43 */ imports.NewTable("TCEFViewComponent_HeightForWidth", 0), // property HeightForWidth
			/* 44 */ imports.NewTable("TCEFViewComponent_OnGetPreferredSize", 0), // event OnGetPreferredSize
			/* 45 */ imports.NewTable("TCEFViewComponent_OnGetMinimumSize", 0), // event OnGetMinimumSize
			/* 46 */ imports.NewTable("TCEFViewComponent_OnGetMaximumSize", 0), // event OnGetMaximumSize
			/* 47 */ imports.NewTable("TCEFViewComponent_OnGetHeightForWidth", 0), // event OnGetHeightForWidth
			/* 48 */ imports.NewTable("TCEFViewComponent_OnParentViewChanged", 0), // event OnParentViewChanged
			/* 49 */ imports.NewTable("TCEFViewComponent_OnChildViewChanged", 0), // event OnChildViewChanged
			/* 50 */ imports.NewTable("TCEFViewComponent_OnWindowChanged", 0), // event OnWindowChanged
			/* 51 */ imports.NewTable("TCEFViewComponent_OnLayoutChanged", 0), // event OnLayoutChanged
			/* 52 */ imports.NewTable("TCEFViewComponent_OnFocus", 0), // event OnFocus
			/* 53 */ imports.NewTable("TCEFViewComponent_OnBlur", 0), // event OnBlur
			/* 54 */ imports.NewTable("TCEFViewComponent_OnThemeChanged", 0), // event OnThemeChanged
		}
	})
	return cEFViewComponentImport
}
