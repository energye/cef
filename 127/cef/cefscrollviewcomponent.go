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
)

// ICEFScrollViewComponent Parent: ICEFViewComponent
type ICEFScrollViewComponent interface {
	ICEFViewComponent
	// CreateScrollView
	//  Create a new ScrollView.
	CreateScrollView() // procedure
	// ContentView
	//  Get and set the content View. The content View must have a specified size (e.g.
	//  via ICefView.SetBounds or ICefViewDelegate.GetPreferredSize).
	ContentView() ICefView         // property ContentView Getter
	SetContentView(value ICefView) // property ContentView Setter
	// VisibleContentRect
	//  Returns the visible region of the content View.
	VisibleContentRect() TCefRect // property VisibleContentRect Getter
	// HorizontalScrollbarHeight
	//  Returns the height of the horizontal scrollbar.
	HorizontalScrollbarHeight() int32 // property HorizontalScrollbarHeight Getter
	// VerticalScrollbarWidth
	//  Returns the width of the vertical scrollbar.
	VerticalScrollbarWidth() int32 // property VerticalScrollbarWidth Getter
	// HasHorizontalScrollbar
	//  Returns true (1) if the horizontal scrollbar is currently showing.
	HasHorizontalScrollbar() bool // property HasHorizontalScrollbar Getter
	// HasVerticalScrollbar
	//  Returns true (1) if the vertical scrollbar is currently showing.
	HasVerticalScrollbar() bool // property HasVerticalScrollbar Getter
	AsIntfViewDelegateEvents() uintptr
}

type TCEFScrollViewComponent struct {
	TCEFViewComponent
}

func (m *TCEFScrollViewComponent) CreateScrollView() {
	if !m.IsValid() {
		return
	}
	cEFScrollViewComponentAPI().SysCallN(1, m.Instance())
}

func (m *TCEFScrollViewComponent) ContentView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFScrollViewComponentAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCEFScrollViewComponent) SetContentView(value ICefView) {
	if !m.IsValid() {
		return
	}
	cEFScrollViewComponentAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCEFScrollViewComponent) VisibleContentRect() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cEFScrollViewComponentAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCEFScrollViewComponent) HorizontalScrollbarHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFScrollViewComponentAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCEFScrollViewComponent) VerticalScrollbarWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFScrollViewComponentAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCEFScrollViewComponent) HasHorizontalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFScrollViewComponentAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFScrollViewComponent) HasVerticalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFScrollViewComponentAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCEFScrollViewComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewScrollViewComponent class constructor
func NewScrollViewComponent(owner lcl.IComponent) ICEFScrollViewComponent {
	var viewDelegateEventsPtr uintptr // ICefViewDelegateEvents
	r := cEFScrollViewComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFScrollViewComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFScrollViewComponentOnce   base.Once
	cEFScrollViewComponentImport *imports.Imports = nil
)

func cEFScrollViewComponentAPI() *imports.Imports {
	cEFScrollViewComponentOnce.Do(func() {
		cEFScrollViewComponentImport = api.NewDefaultImports()
		cEFScrollViewComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFScrollViewComponent_Create", 0), // constructor NewScrollViewComponent
			/* 1 */ imports.NewTable("TCEFScrollViewComponent_CreateScrollView", 0), // procedure CreateScrollView
			/* 2 */ imports.NewTable("TCEFScrollViewComponent_ContentView", 0), // property ContentView
			/* 3 */ imports.NewTable("TCEFScrollViewComponent_VisibleContentRect", 0), // property VisibleContentRect
			/* 4 */ imports.NewTable("TCEFScrollViewComponent_HorizontalScrollbarHeight", 0), // property HorizontalScrollbarHeight
			/* 5 */ imports.NewTable("TCEFScrollViewComponent_VerticalScrollbarWidth", 0), // property VerticalScrollbarWidth
			/* 6 */ imports.NewTable("TCEFScrollViewComponent_HasHorizontalScrollbar", 0), // property HasHorizontalScrollbar
			/* 7 */ imports.NewTable("TCEFScrollViewComponent_HasVerticalScrollbar", 0), // property HasVerticalScrollbar
		}
	})
	return cEFScrollViewComponentImport
}
