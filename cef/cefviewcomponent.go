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
	ToStringEx(includeChildren bool) string                    // function
	IsSame(that ICefView) bool                                 // function
	ConvertPointToScreen(point *TCefPoint) bool                // function
	ConvertPointFromScreen(point *TCefPoint) bool              // function
	ConvertPointToWindow(point *TCefPoint) bool                // function
	ConvertPointFromWindow(point *TCefPoint) bool              // function
	ConvertPointToView(view ICefView, point *TCefPoint) bool   // function
	ConvertPointFromView(view ICefView, point *TCefPoint) bool // function
	SizeToPreferredSize()                                      // procedure
	InvalidateLayout()                                         // procedure
	RequestFocus()                                             // procedure
	Initialized() bool                                         // property Initialized Getter
	AsView() ICefView                                          // property AsView Getter
	AsBrowserView() ICefBrowserView                            // property AsBrowserView Getter
	AsButton() ICefButton                                      // property AsButton Getter
	AsPanel() ICefPanel                                        // property AsPanel Getter
	AsScrollView() ICefScrollView                              // property AsScrollView Getter
	AsTextfield() ICefTextfield                                // property AsTextfield Getter
	ViewForID(id int32) ICefView                               // property ViewForID Getter
	Valid() bool                                               // property Valid Getter
	Attached() bool                                            // property Attached Getter
	Delegate() IEngViewDelegate                                // property Delegate Getter
	Window() ICefWindow                                        // property Window Getter
	ParentView() ICefView                                      // property ParentView Getter
	BoundsInScreen() TCefRect                                  // property BoundsInScreen Getter
	PreferredSize() TCefSize                                   // property PreferredSize Getter
	MinimumSize() TCefSize                                     // property MinimumSize Getter
	MaximumSize() TCefSize                                     // property MaximumSize Getter
	Visible() bool                                             // property Visible Getter
	SetVisible(value bool)                                     // property Visible Setter
	Drawn() bool                                               // property Drawn Getter
	Enabled() bool                                             // property Enabled Getter
	SetEnabled(value bool)                                     // property Enabled Setter
	Focusable() bool                                           // property Focusable Getter
	SetFocusable(value bool)                                   // property Focusable Setter
	AccessibilityFocusable() bool                              // property AccessibilityFocusable Getter
	BackgroundColor() cefTypes.TCefColor                       // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor)               // property BackgroundColor Setter
	ID() int32                                                 // property ID Getter
	SetID(value int32)                                         // property ID Setter
	GroupID() int32                                            // property GroupID Getter
	SetGroupID(value int32)                                    // property GroupID Setter
	Bounds() TCefRect                                          // property Bounds Getter
	SetBounds(value TCefRect)                                  // property Bounds Setter
	Size() TCefSize                                            // property Size Getter
	SetSize(value TCefSize)                                    // property Size Setter
	Position() TCefPoint                                       // property Position Getter
	SetPosition(value TCefPoint)                               // property Position Setter
	TypeString() string                                        // property TypeString Getter
	HeightForWidth(width int32) int32                          // property HeightForWidth Getter
	SetOnGetPreferredSize(fn TOnGetPreferredSizeEvent)         // property event
	SetOnGetMinimumSize(fn TOnGetMinimumSizeEvent)             // property event
	SetOnGetMaximumSize(fn TOnGetMaximumSizeEvent)             // property event
	SetOnGetHeightForWidth(fn TOnGetHeightForWidthEvent)       // property event
	SetOnParentViewChanged(fn TOnParentViewChangedEvent)       // property event
	SetOnChildViewChanged(fn TOnChildViewChangedEvent)         // property event
	SetOnWindowChanged(fn TOnWindowChangedEvent)               // property event
	SetOnLayoutChanged(fn TOnLayoutChangedEvent)               // property event
	SetOnFocus(fn TOnFocusEvent)                               // property event
	SetOnBlur(fn TOnBlurEvent)                                 // property event
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

func (m *TCEFViewComponent) ConvertPointToScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromScreen(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointToWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromWindow(point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(view), uintptr(base.UnsafePointer(point)))
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(9, m.Instance())
}

func (m *TCEFViewComponent) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(10, m.Instance())
}

func (m *TCEFViewComponent) RequestFocus() {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(11, m.Instance())
}

func (m *TCEFViewComponent) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) AsView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsBrowserView() (result ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsButton() (result ICefButton) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefButtonRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsPanel() (result ICefPanel) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPanelRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsScrollView() (result ICefScrollView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefScrollViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) AsTextfield() (result ICefTextfield) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefTextfieldRef(resultPtr)
	return
}

func (m *TCEFViewComponent) ViewForID(id int32) (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(19, m.Instance(), uintptr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) Valid() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Attached() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Delegate() (result IEngViewDelegate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngViewDelegate(resultPtr)
	return
}

func (m *TCEFViewComponent) Window() (result ICefWindow) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefWindowRef(resultPtr)
	return
}

func (m *TCEFViewComponent) ParentView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFViewComponentAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFViewComponent) BoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) PreferredSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(26, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) MinimumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(27, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) MaximumSize() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(28, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) Drawn() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(31, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(31, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) Focusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) SetFocusable(value bool) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

func (m *TCEFViewComponent) AccessibilityFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFViewComponentAPI().SysCallN(33, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFViewComponent) BackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(34, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCEFViewComponent) SetBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) ID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(35, 0, m.Instance())
	return int32(r)
}

func (m *TCEFViewComponent) SetID(value int32) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) GroupID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(36, 0, m.Instance())
	return int32(r)
}

func (m *TCEFViewComponent) SetGroupID(value int32) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TCEFViewComponent) Bounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(37, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetBounds(value TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(37, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) Size() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(38, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetSize(value TCefSize) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(38, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) Position() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(39, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFViewComponent) SetPosition(value TCefPoint) {
	if !m.IsValid() {
		return
	}
	cEFViewComponentAPI().SysCallN(39, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCEFViewComponent) TypeString() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFViewComponentAPI().SysCallN(40, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFViewComponent) HeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFViewComponentAPI().SysCallN(41, m.Instance(), uintptr(width))
	return int32(r)
}

func (m *TCEFViewComponent) SetOnGetPreferredSize(fn TOnGetPreferredSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetPreferredSizeEvent(fn)
	base.SetEvent(m, 42, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetMinimumSize(fn TOnGetMinimumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetMinimumSizeEvent(fn)
	base.SetEvent(m, 43, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetMaximumSize(fn TOnGetMaximumSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetMaximumSizeEvent(fn)
	base.SetEvent(m, 44, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnGetHeightForWidth(fn TOnGetHeightForWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetHeightForWidthEvent(fn)
	base.SetEvent(m, 45, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnParentViewChanged(fn TOnParentViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnParentViewChangedEvent(fn)
	base.SetEvent(m, 46, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnChildViewChanged(fn TOnChildViewChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnChildViewChangedEvent(fn)
	base.SetEvent(m, 47, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnWindowChanged(fn TOnWindowChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowChangedEvent(fn)
	base.SetEvent(m, 48, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnLayoutChanged(fn TOnLayoutChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLayoutChangedEvent(fn)
	base.SetEvent(m, 49, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnFocus(fn TOnFocusEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusEvent(fn)
	base.SetEvent(m, 50, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFViewComponent) SetOnBlur(fn TOnBlurEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBlurEvent(fn)
	base.SetEvent(m, 51, cEFViewComponentAPI(), api.MakeEventDataPtr(cb))
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
			/* 3 */ imports.NewTable("TCEFViewComponent_ConvertPointToScreen", 0), // function ConvertPointToScreen
			/* 4 */ imports.NewTable("TCEFViewComponent_ConvertPointFromScreen", 0), // function ConvertPointFromScreen
			/* 5 */ imports.NewTable("TCEFViewComponent_ConvertPointToWindow", 0), // function ConvertPointToWindow
			/* 6 */ imports.NewTable("TCEFViewComponent_ConvertPointFromWindow", 0), // function ConvertPointFromWindow
			/* 7 */ imports.NewTable("TCEFViewComponent_ConvertPointToView", 0), // function ConvertPointToView
			/* 8 */ imports.NewTable("TCEFViewComponent_ConvertPointFromView", 0), // function ConvertPointFromView
			/* 9 */ imports.NewTable("TCEFViewComponent_SizeToPreferredSize", 0), // procedure SizeToPreferredSize
			/* 10 */ imports.NewTable("TCEFViewComponent_InvalidateLayout", 0), // procedure InvalidateLayout
			/* 11 */ imports.NewTable("TCEFViewComponent_RequestFocus", 0), // procedure RequestFocus
			/* 12 */ imports.NewTable("TCEFViewComponent_Initialized", 0), // property Initialized
			/* 13 */ imports.NewTable("TCEFViewComponent_AsView", 0), // property AsView
			/* 14 */ imports.NewTable("TCEFViewComponent_AsBrowserView", 0), // property AsBrowserView
			/* 15 */ imports.NewTable("TCEFViewComponent_AsButton", 0), // property AsButton
			/* 16 */ imports.NewTable("TCEFViewComponent_AsPanel", 0), // property AsPanel
			/* 17 */ imports.NewTable("TCEFViewComponent_AsScrollView", 0), // property AsScrollView
			/* 18 */ imports.NewTable("TCEFViewComponent_AsTextfield", 0), // property AsTextfield
			/* 19 */ imports.NewTable("TCEFViewComponent_ViewForID", 0), // property ViewForID
			/* 20 */ imports.NewTable("TCEFViewComponent_Valid", 0), // property Valid
			/* 21 */ imports.NewTable("TCEFViewComponent_Attached", 0), // property Attached
			/* 22 */ imports.NewTable("TCEFViewComponent_Delegate", 0), // property Delegate
			/* 23 */ imports.NewTable("TCEFViewComponent_Window", 0), // property Window
			/* 24 */ imports.NewTable("TCEFViewComponent_ParentView", 0), // property ParentView
			/* 25 */ imports.NewTable("TCEFViewComponent_BoundsInScreen", 0), // property BoundsInScreen
			/* 26 */ imports.NewTable("TCEFViewComponent_PreferredSize", 0), // property PreferredSize
			/* 27 */ imports.NewTable("TCEFViewComponent_MinimumSize", 0), // property MinimumSize
			/* 28 */ imports.NewTable("TCEFViewComponent_MaximumSize", 0), // property MaximumSize
			/* 29 */ imports.NewTable("TCEFViewComponent_Visible", 0), // property Visible
			/* 30 */ imports.NewTable("TCEFViewComponent_Drawn", 0), // property Drawn
			/* 31 */ imports.NewTable("TCEFViewComponent_Enabled", 0), // property Enabled
			/* 32 */ imports.NewTable("TCEFViewComponent_Focusable", 0), // property Focusable
			/* 33 */ imports.NewTable("TCEFViewComponent_AccessibilityFocusable", 0), // property AccessibilityFocusable
			/* 34 */ imports.NewTable("TCEFViewComponent_BackgroundColor", 0), // property BackgroundColor
			/* 35 */ imports.NewTable("TCEFViewComponent_ID", 0), // property ID
			/* 36 */ imports.NewTable("TCEFViewComponent_GroupID", 0), // property GroupID
			/* 37 */ imports.NewTable("TCEFViewComponent_Bounds", 0), // property Bounds
			/* 38 */ imports.NewTable("TCEFViewComponent_Size", 0), // property Size
			/* 39 */ imports.NewTable("TCEFViewComponent_Position", 0), // property Position
			/* 40 */ imports.NewTable("TCEFViewComponent_TypeString", 0), // property TypeString
			/* 41 */ imports.NewTable("TCEFViewComponent_HeightForWidth", 0), // property HeightForWidth
			/* 42 */ imports.NewTable("TCEFViewComponent_OnGetPreferredSize", 0), // event OnGetPreferredSize
			/* 43 */ imports.NewTable("TCEFViewComponent_OnGetMinimumSize", 0), // event OnGetMinimumSize
			/* 44 */ imports.NewTable("TCEFViewComponent_OnGetMaximumSize", 0), // event OnGetMaximumSize
			/* 45 */ imports.NewTable("TCEFViewComponent_OnGetHeightForWidth", 0), // event OnGetHeightForWidth
			/* 46 */ imports.NewTable("TCEFViewComponent_OnParentViewChanged", 0), // event OnParentViewChanged
			/* 47 */ imports.NewTable("TCEFViewComponent_OnChildViewChanged", 0), // event OnChildViewChanged
			/* 48 */ imports.NewTable("TCEFViewComponent_OnWindowChanged", 0), // event OnWindowChanged
			/* 49 */ imports.NewTable("TCEFViewComponent_OnLayoutChanged", 0), // event OnLayoutChanged
			/* 50 */ imports.NewTable("TCEFViewComponent_OnFocus", 0), // event OnFocus
			/* 51 */ imports.NewTable("TCEFViewComponent_OnBlur", 0), // event OnBlur
		}
	})
	return cEFViewComponentImport
}
