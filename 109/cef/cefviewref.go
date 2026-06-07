//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefView Parent: ICefBaseRefCounted
type ICefView interface {
	ICefBaseRefCounted
	AsBrowserView() ICefBrowserView                            // function
	AsButton() ICefButton                                      // function
	AsPanel() ICefPanel                                        // function
	AsScrollView() ICefScrollView                              // function
	AsTextfield() ICefTextfield                                // function
	GetTypeString() string                                     // function
	ToStringEx(includeChildren bool) string                    // function
	IsValid() bool                                             // function
	IsAttached() bool                                          // function
	IsSame(that ICefView) bool                                 // function
	GetDelegate() IEngViewDelegate                             // function
	GetWindow() ICefWindow                                     // function
	GetID() int32                                              // function
	GetGroupID() int32                                         // function
	GetParentView() ICefView                                   // function
	GetViewForID(id int32) ICefView                            // function
	GetBounds() TCefRect                                       // function
	GetBoundsInScreen() TCefRect                               // function
	GetSize() TCefSize                                         // function
	GetPosition() TCefPoint                                    // function
	GetInsets() TCefInsets                                     // function
	GetPreferredSize() TCefSize                                // function
	GetMinimumSize() TCefSize                                  // function
	GetMaximumSize() TCefSize                                  // function
	GetHeightForWidth(width int32) int32                       // function
	IsVisible() bool                                           // function
	IsDrawn() bool                                             // function
	IsEnabled() bool                                           // function
	IsFocusable() bool                                         // function
	IsAccessibilityFocusable() bool                            // function
	GetBackgroundColor() cefTypes.TCefColor                    // function
	ConvertPointToScreen(point *TCefPoint) bool                // function
	ConvertPointFromScreen(point *TCefPoint) bool              // function
	ConvertPointToWindow(point *TCefPoint) bool                // function
	ConvertPointFromWindow(point *TCefPoint) bool              // function
	ConvertPointToView(view ICefView, point *TCefPoint) bool   // function
	ConvertPointFromView(view ICefView, point *TCefPoint) bool // function
	SetID(id int32)                                            // procedure
	SetGroupID(groupId int32)                                  // procedure
	SetBounds(bounds TCefRect)                                 // procedure
	SetSize(size TCefSize)                                     // procedure
	SetPosition(position TCefPoint)                            // procedure
	SetInsets(insets TCefInsets)                               // procedure
	SizeToPreferredSize()                                      // procedure
	InvalidateLayout()                                         // procedure
	SetVisible(visible bool)                                   // procedure
	SetEnabled(enabled bool)                                   // procedure
	SetFocusable(focusable bool)                               // procedure
	RequestFocus()                                             // procedure
	SetBackgroundColor(color cefTypes.TCefColor)               // procedure
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

func (m *TCefViewRef) GetBackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefViewRefAPI().SysCallN(31, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefViewRef) ConvertPointToScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(32, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(33, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointToWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(34, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(35, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(36, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cefViewRefAPI().SysCallN(37, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCefViewRef) SetID(id int32) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(39, m.Instance(), uintptr(id))
}

func (m *TCefViewRef) SetGroupID(groupId int32) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(40, m.Instance(), uintptr(groupId))
}

func (m *TCefViewRef) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(41, m.Instance(), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TCefViewRef) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(42, m.Instance(), uintptr(base.UnsafePointer(&size)))
}

func (m *TCefViewRef) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(43, m.Instance(), uintptr(base.UnsafePointer(&position)))
}

func (m *TCefViewRef) SetInsets(insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(44, m.Instance(), uintptr(base.UnsafePointer(&insets)))
}

func (m *TCefViewRef) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(45, m.Instance())
}

func (m *TCefViewRef) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(46, m.Instance())
}

func (m *TCefViewRef) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(47, m.Instance(), api.PasBool(visible))
}

func (m *TCefViewRef) SetEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(48, m.Instance(), api.PasBool(enabled))
}

func (m *TCefViewRef) SetFocusable(focusable bool) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(49, m.Instance(), api.PasBool(focusable))
}

func (m *TCefViewRef) RequestFocus() {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(50, m.Instance())
}

func (m *TCefViewRef) SetBackgroundColor(color cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefViewRefAPI().SysCallN(51, m.Instance(), uintptr(color))
}

func (m *TCefViewRef) AsIntfView() uintptr {
	return m.GetIntfPointer(0)
}

// ViewRef  is static instance
var ViewRef _ViewRefClass

// _ViewRefClass is class type defined by TCefViewRef
type _ViewRefClass uintptr

func (_ViewRefClass) UnWrapWithPointer(data uintptr) (result ICefView) {
	var resultPtr uintptr
	cefViewRefAPI().SysCallN(38, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
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
			/* 31 */ imports.NewTable("TCefViewRef_GetBackgroundColor", 0), // function GetBackgroundColor
			/* 32 */ imports.NewTable("TCefViewRef_ConvertPointToScreen", 0), // function ConvertPointToScreen
			/* 33 */ imports.NewTable("TCefViewRef_ConvertPointFromScreen", 0), // function ConvertPointFromScreen
			/* 34 */ imports.NewTable("TCefViewRef_ConvertPointToWindow", 0), // function ConvertPointToWindow
			/* 35 */ imports.NewTable("TCefViewRef_ConvertPointFromWindow", 0), // function ConvertPointFromWindow
			/* 36 */ imports.NewTable("TCefViewRef_ConvertPointToView", 0), // function ConvertPointToView
			/* 37 */ imports.NewTable("TCefViewRef_ConvertPointFromView", 0), // function ConvertPointFromView
			/* 38 */ imports.NewTable("TCefViewRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 39 */ imports.NewTable("TCefViewRef_SetID", 0), // procedure SetID
			/* 40 */ imports.NewTable("TCefViewRef_SetGroupID", 0), // procedure SetGroupID
			/* 41 */ imports.NewTable("TCefViewRef_SetBounds", 0), // procedure SetBounds
			/* 42 */ imports.NewTable("TCefViewRef_SetSize", 0), // procedure SetSize
			/* 43 */ imports.NewTable("TCefViewRef_SetPosition", 0), // procedure SetPosition
			/* 44 */ imports.NewTable("TCefViewRef_SetInsets", 0), // procedure SetInsets
			/* 45 */ imports.NewTable("TCefViewRef_SizeToPreferredSize", 0), // procedure SizeToPreferredSize
			/* 46 */ imports.NewTable("TCefViewRef_InvalidateLayout", 0), // procedure InvalidateLayout
			/* 47 */ imports.NewTable("TCefViewRef_SetVisible", 0), // procedure SetVisible
			/* 48 */ imports.NewTable("TCefViewRef_SetEnabled", 0), // procedure SetEnabled
			/* 49 */ imports.NewTable("TCefViewRef_SetFocusable", 0), // procedure SetFocusable
			/* 50 */ imports.NewTable("TCefViewRef_RequestFocus", 0), // procedure RequestFocus
			/* 51 */ imports.NewTable("TCefViewRef_SetBackgroundColor", 0), // procedure SetBackgroundColor
		}
	})
	return cefViewRefImport
}
