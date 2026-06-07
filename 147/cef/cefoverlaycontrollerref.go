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

// ICefOverlayController Parent: ICefBaseRefCounted
type ICefOverlayController interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid.
	IsValid() bool // function
	// IsSame
	//  Returns true (1) if this object is the same as |that| object.
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
	GetDockingMode() cefTypes.TCefDockingMode // function
	// GetBounds
	//  Returns the bounds (size and position) of this overlay in parent
	//  coordinates.
	GetBounds() TCefRect // function
	// GetBoundsInScreen
	//  Returns the bounds (size and position) of this overlay in DIP screen
	//  coordinates.
	GetBoundsInScreen() TCefRect // function
	// GetSize
	//  Returns the size of this overlay in parent coordinates.
	GetSize() TCefSize // function
	// GetPosition
	//  Returns the position of this overlay in parent coordinates.
	GetPosition() TCefPoint // function
	// GetInsets
	//  Returns the insets for this overlay in parent coordinates.
	GetInsets() TCefInsets // function
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
	//  Sets the bounds (size and position) of this overlay. This will set the
	//  bounds of the contents View to match and trigger a re-layout if necessary.
	//  |bounds| is in parent coordinates and any insets configured on this
	//  overlay will be ignored. Use this function only for overlays created with
	//  a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking modes
	//  modify the insets of this overlay and/or layout of the contents View and
	//  call size_to_preferred_size() instead to calculate the new size and re-
	//  position the overlay if necessary.
	SetBounds(bounds TCefRect) // procedure
	// SetSize
	//  Sets the size of this overlay without changing the position. This will set
	//  the size of the contents View to match and trigger a re-layout if
	//  necessary. |size| is in parent coordinates and any insets configured on
	//  this overlay will be ignored. Use this function only for overlays created
	//  with a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking
	//  modes modify the insets of this overlay and/or layout of the contents View
	//  and call size_to_preferred_size() instead to calculate the new size and
	//  re-position the overlay if necessary.
	SetSize(size TCefSize) // procedure
	// SetPosition
	//  Sets the position of this overlay without changing the size. |position| is
	//  in parent coordinates and any insets configured on this overlay will be
	//  ignored. Use this function only for overlays created with a docking mode
	//  value of CEF_DOCKING_MODE_CUSTOM. With other docking modes modify the
	//  insets of this overlay and/or layout of the contents View and call
	//  size_to_preferred_size() instead to calculate the new size and re-position
	//  the overlay if necessary.
	SetPosition(position TCefPoint) // procedure
	// SetInsets
	//  Sets the insets for this overlay. |insets| is in parent coordinates. Use
	//  this function only for overlays created with a docking mode value other
	//  than CEF_DOCKING_MODE_CUSTOM.
	SetInsets(insets TCefInsets) // procedure
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

// ICefOverlayControllerRef Parent: ICefOverlayController ICefBaseRefCountedRef
type ICefOverlayControllerRef interface {
	ICefOverlayController
	ICefBaseRefCountedRef
	AsIntfOverlayController() uintptr
}

type TCefOverlayControllerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefOverlayControllerRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefOverlayControllerRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefOverlayControllerRef) IsSame(that ICefOverlayController) bool {
	if !m.IsValid() {
		return false
	}
	r := cefOverlayControllerRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefOverlayControllerRef) GetContentsView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefOverlayControllerRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefOverlayControllerRef) GetWindow() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefOverlayControllerRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCefOverlayControllerRef) GetDockingMode() cefTypes.TCefDockingMode {
	if !m.IsValid() {
		return 0
	}
	r := cefOverlayControllerRefAPI().SysCallN(5, m.Instance())
	return cefTypes.TCefDockingMode(r)
}

func (m *TCefOverlayControllerRef) GetBounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefOverlayControllerRef) GetBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefOverlayControllerRef) GetSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefOverlayControllerRef) GetPosition() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefOverlayControllerRef) GetInsets() (result TCefInsets) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefOverlayControllerRef) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := cefOverlayControllerRefAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCefOverlayControllerRef) IsDrawn() bool {
	if !m.IsValid() {
		return false
	}
	r := cefOverlayControllerRefAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCefOverlayControllerRef) DestroyOverlay() {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(14, m.Instance())
}

func (m *TCefOverlayControllerRef) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TCefOverlayControllerRef) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefOverlayControllerRef) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&position)))
}

func (m *TCefOverlayControllerRef) SetInsets(insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&insets)))
}

func (m *TCefOverlayControllerRef) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(19, m.Instance())
}

func (m *TCefOverlayControllerRef) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	cefOverlayControllerRefAPI().SysCallN(20, m.Instance(), api.PasBool(visible))
}

func (m *TCefOverlayControllerRef) AsIntfOverlayController() uintptr {
	return m.GetIntfPointer(0)
}

// OverlayControllerRef  is static instance
var OverlayControllerRef _OverlayControllerRefClass

// _OverlayControllerRefClass is class type defined by TCefOverlayControllerRef
type _OverlayControllerRefClass uintptr

// UnWrap
//
//	Returns a ICefOverlayController instance using a PCefOverlayController data pointer.
func (_OverlayControllerRefClass) UnWrap(data uintptr) (result ICefOverlayController) {
	var resultPtr uintptr
	cefOverlayControllerRefAPI().SysCallN(13, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefOverlayControllerRef(resultPtr)
	return
}

// NewOverlayControllerRef class constructor
func NewOverlayControllerRef(data uintptr) ICefOverlayControllerRef {
	var overlayControllerPtr uintptr // ICefOverlayController
	r := cefOverlayControllerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&overlayControllerPtr)))
	ret := AsCefOverlayControllerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, overlayControllerPtr)
	}
	return ret
}

var (
	cefOverlayControllerRefOnce   base.Once
	cefOverlayControllerRefImport *imports.Imports = nil
)

func cefOverlayControllerRefAPI() *imports.Imports {
	cefOverlayControllerRefOnce.Do(func() {
		cefOverlayControllerRefImport = api.NewDefaultImports()
		cefOverlayControllerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefOverlayControllerRef_Create", 0), // constructor NewOverlayControllerRef
			/* 1 */ imports.NewTable("TCefOverlayControllerRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefOverlayControllerRef_IsSame", 0), // function IsSame
			/* 3 */ imports.NewTable("TCefOverlayControllerRef_GetContentsView", 0), // function GetContentsView
			/* 4 */ imports.NewTable("TCefOverlayControllerRef_GetWindow", 0), // function GetWindow
			/* 5 */ imports.NewTable("TCefOverlayControllerRef_GetDockingMode", 0), // function GetDockingMode
			/* 6 */ imports.NewTable("TCefOverlayControllerRef_GetBounds", 0), // function GetBounds
			/* 7 */ imports.NewTable("TCefOverlayControllerRef_GetBoundsInScreen", 0), // function GetBoundsInScreen
			/* 8 */ imports.NewTable("TCefOverlayControllerRef_GetSize", 0), // function GetSize
			/* 9 */ imports.NewTable("TCefOverlayControllerRef_GetPosition", 0), // function GetPosition
			/* 10 */ imports.NewTable("TCefOverlayControllerRef_GetInsets", 0), // function GetInsets
			/* 11 */ imports.NewTable("TCefOverlayControllerRef_IsVisible", 0), // function IsVisible
			/* 12 */ imports.NewTable("TCefOverlayControllerRef_IsDrawn", 0), // function IsDrawn
			/* 13 */ imports.NewTable("TCefOverlayControllerRef_UnWrap", 0), // static function UnWrap
			/* 14 */ imports.NewTable("TCefOverlayControllerRef_DestroyOverlay", 0), // procedure DestroyOverlay
			/* 15 */ imports.NewTable("TCefOverlayControllerRef_SetBounds", 0), // procedure SetBounds
			/* 16 */ imports.NewTable("TCefOverlayControllerRef_SetSize", 0), // procedure SetSize
			/* 17 */ imports.NewTable("TCefOverlayControllerRef_SetPosition", 0), // procedure SetPosition
			/* 18 */ imports.NewTable("TCefOverlayControllerRef_SetInsets", 0), // procedure SetInsets
			/* 19 */ imports.NewTable("TCefOverlayControllerRef_SizeToPreferredSize", 0), // procedure SizeToPreferredSize
			/* 20 */ imports.NewTable("TCefOverlayControllerRef_SetVisible", 0), // procedure SetVisible
		}
	})
	return cefOverlayControllerRefImport
}
