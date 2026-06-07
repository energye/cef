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
)

// ICefScrollView Parent: ICefView
type ICefScrollView interface {
	ICefView
	GetContentView() ICefView            // function
	GetVisibleContentRect() TCefRect     // function
	HasHorizontalScrollbar() bool        // function
	GetHorizontalScrollbarHeight() int32 // function
	HasVerticalScrollbar() bool          // function
	GetVerticalScrollbarWidth() int32    // function
	SetContentView(view ICefView)        // procedure
}

// ICefScrollViewRef Parent: ICefScrollView ICefViewRef
type ICefScrollViewRef interface {
	ICefScrollView
	ICefViewRef
	AsIntfScrollView() uintptr
	AsIntfView() uintptr
}

type TCefScrollViewRef struct {
	TCefViewRef
}

func (m *TCefScrollViewRef) GetContentView() (result ICefView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefScrollViewRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefViewRef(resultPtr)
	return
}

func (m *TCefScrollViewRef) GetVisibleContentRect() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefScrollViewRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefScrollViewRef) HasHorizontalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r := cefScrollViewRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefScrollViewRef) GetHorizontalScrollbarHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefScrollViewRefAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCefScrollViewRef) HasVerticalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r := cefScrollViewRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefScrollViewRef) GetVerticalScrollbarWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefScrollViewRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefScrollViewRef) SetContentView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cefScrollViewRefAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCefScrollViewRef) AsIntfScrollView() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefScrollViewRef) AsIntfView() uintptr {
	return m.GetIntfPointer(1)
}

// ScrollViewRef  is static instance
var ScrollViewRef _ScrollViewRefClass

// _ScrollViewRefClass is class type defined by TCefScrollViewRef
type _ScrollViewRefClass uintptr

func (_ScrollViewRefClass) UnWrapWithPointer(data uintptr) (result ICefScrollView) {
	var resultPtr uintptr
	cefScrollViewRefAPI().SysCallN(7, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefScrollViewRef(resultPtr)
	return
}

func (_ScrollViewRefClass) CreateScrollView(delegate IEngViewDelegate) (result ICefScrollView) {
	var resultPtr uintptr
	cefScrollViewRefAPI().SysCallN(8, base.GetObjectUintptr(delegate), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefScrollViewRef(resultPtr)
	return
}

// NewScrollViewRef class constructor
func NewScrollViewRef(data uintptr) ICefScrollViewRef {
	var scrollViewPtr uintptr // ICefScrollView
	var viewPtr uintptr       // ICefView
	r := cefScrollViewRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&scrollViewPtr)), uintptr(base.UnsafePointer(&viewPtr)))
	ret := AsCefScrollViewRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, scrollViewPtr)
		intf.SetIntfPointer(1, viewPtr)
	}
	return ret
}

var (
	cefScrollViewRefOnce   base.Once
	cefScrollViewRefImport *imports.Imports = nil
)

func cefScrollViewRefAPI() *imports.Imports {
	cefScrollViewRefOnce.Do(func() {
		cefScrollViewRefImport = api.NewDefaultImports()
		cefScrollViewRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefScrollViewRef_Create", 0), // constructor NewScrollViewRef
			/* 1 */ imports.NewTable("TCefScrollViewRef_GetContentView", 0), // function GetContentView
			/* 2 */ imports.NewTable("TCefScrollViewRef_GetVisibleContentRect", 0), // function GetVisibleContentRect
			/* 3 */ imports.NewTable("TCefScrollViewRef_HasHorizontalScrollbar", 0), // function HasHorizontalScrollbar
			/* 4 */ imports.NewTable("TCefScrollViewRef_GetHorizontalScrollbarHeight", 0), // function GetHorizontalScrollbarHeight
			/* 5 */ imports.NewTable("TCefScrollViewRef_HasVerticalScrollbar", 0), // function HasVerticalScrollbar
			/* 6 */ imports.NewTable("TCefScrollViewRef_GetVerticalScrollbarWidth", 0), // function GetVerticalScrollbarWidth
			/* 7 */ imports.NewTable("TCefScrollViewRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 8 */ imports.NewTable("TCefScrollViewRef_CreateScrollView", 0), // static function CreateScrollView
			/* 9 */ imports.NewTable("TCefScrollViewRef_SetContentView", 0), // procedure SetContentView
		}
	})
	return cefScrollViewRefImport
}
