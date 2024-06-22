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

// ICEFScrollViewComponent Parent: ICEFViewComponent
type ICEFScrollViewComponent interface {
	ICEFViewComponent
	// ContentView
	//  Get and set the content View. The content View must have a specified size(e.g.
	//  via ICefView.SetBounds or ICefViewDelegate.GetPreferredSize).
	ContentView() ICefView // property
	// SetContentView Set ContentView
	SetContentView(AValue ICefView) // property
	// VisibleContentRect
	//  Returns the visible region of the content View.
	VisibleContentRect() (resultCefRect TCefRect) // property
	// HorizontalScrollbarHeight
	//  Returns the height of the horizontal scrollbar.
	HorizontalScrollbarHeight() int32 // property
	// VerticalScrollbarWidth
	//  Returns the width of the vertical scrollbar.
	VerticalScrollbarWidth() int32 // property
	// HasHorizontalScrollbar
	//  Returns true(1) if the horizontal scrollbar is currently showing.
	HasHorizontalScrollbar() bool // property
	// HasVerticalScrollbar
	//  Returns true(1) if the vertical scrollbar is currently showing.
	HasVerticalScrollbar() bool // property
	// CreateScrollView
	//  Create a new ScrollView.
	CreateScrollView() // procedure
}

// TCEFScrollViewComponent Parent: TCEFViewComponent
type TCEFScrollViewComponent struct {
	TCEFViewComponent
}

func NewCEFScrollViewComponent(aOwner IComponent) ICEFScrollViewComponent {
	r1 := scrollViewComponentImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsCEFScrollViewComponent(r1)
}

func (m *TCEFScrollViewComponent) ContentView() ICefView {
	var resultCefView uintptr
	scrollViewComponentImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCEFScrollViewComponent) SetContentView(AValue ICefView) {
	scrollViewComponentImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFScrollViewComponent) VisibleContentRect() (resultCefRect TCefRect) {
	scrollViewComponentImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCEFScrollViewComponent) HorizontalScrollbarHeight() int32 {
	r1 := scrollViewComponentImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TCEFScrollViewComponent) VerticalScrollbarWidth() int32 {
	r1 := scrollViewComponentImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TCEFScrollViewComponent) HasHorizontalScrollbar() bool {
	r1 := scrollViewComponentImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCEFScrollViewComponent) HasVerticalScrollbar() bool {
	r1 := scrollViewComponentImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCEFScrollViewComponent) CreateScrollView() {
	scrollViewComponentImportAPI().SysCallN(2, m.Instance())
}

var (
	scrollViewComponentImport       *imports.Imports = nil
	scrollViewComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFScrollViewComponent_ContentView", 0),
		/*1*/ imports.NewTable("CEFScrollViewComponent_Create", 0),
		/*2*/ imports.NewTable("CEFScrollViewComponent_CreateScrollView", 0),
		/*3*/ imports.NewTable("CEFScrollViewComponent_HasHorizontalScrollbar", 0),
		/*4*/ imports.NewTable("CEFScrollViewComponent_HasVerticalScrollbar", 0),
		/*5*/ imports.NewTable("CEFScrollViewComponent_HorizontalScrollbarHeight", 0),
		/*6*/ imports.NewTable("CEFScrollViewComponent_VerticalScrollbarWidth", 0),
		/*7*/ imports.NewTable("CEFScrollViewComponent_VisibleContentRect", 0),
	}
)

func scrollViewComponentImportAPI() *imports.Imports {
	if scrollViewComponentImport == nil {
		scrollViewComponentImport = NewDefaultImports()
		scrollViewComponentImport.SetImportTable(scrollViewComponentImportTables)
		scrollViewComponentImportTables = nil
	}
	return scrollViewComponentImport
}
