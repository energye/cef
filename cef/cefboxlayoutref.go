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

// ICefBoxLayout Parent: ICefLayoutRef
type ICefBoxLayout interface {
	ICefLayoutRef
	// SetFlexForView
	//  Set the flex weight for the given |view|. Using the preferred size as the
	//  basis, free space along the main axis is distributed to views in the ratio
	//  of their flex weights. Similarly, if the views will overflow the parent,
	//  space is subtracted in these ratios. A flex of 0 means this view is not
	//  resized. Flex values must not be negative.
	SetFlexForView(view ICefView, flex int32) // procedure
	// ClearFlexForView
	//  Clears the flex for the given |view|, causing it to use the default flex
	//  specified via TCefBoxLayoutSettings.default_flex.
	ClearFlexForView(view ICefView) // procedure
}

// ICefBoxLayoutRef Parent: ICefBoxLayout
type ICefBoxLayoutRef interface {
	ICefBoxLayout
	AsIntfBoxLayout() uintptr
	AsIntfLayout() uintptr
}

type TCefBoxLayoutRef struct {
	TCefLayoutRef
}

func (m *TCefBoxLayoutRef) SetFlexForView(view ICefView, flex int32) {
	if !m.IsValid() {
		return
	}
	cefBoxLayoutRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(view), uintptr(flex))
}

func (m *TCefBoxLayoutRef) ClearFlexForView(view ICefView) {
	if !m.IsValid() {
		return
	}
	cefBoxLayoutRefAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(view))
}

func (m *TCefBoxLayoutRef) AsIntfBoxLayout() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefBoxLayoutRef) AsIntfLayout() uintptr {
	return m.GetIntfPointer(1)
}

// BoxLayoutRef  is static instance
var BoxLayoutRef _BoxLayoutRefClass

// _BoxLayoutRefClass is class type defined by TCefBoxLayoutRef
type _BoxLayoutRefClass uintptr

// UnWrapWithPointer
//
//	Returns a ICefBoxLayout instance using a PCefBoxLayout data pointer.
func (_BoxLayoutRefClass) UnWrapWithPointer(data uintptr) (result ICefBoxLayout) {
	var resultPtr uintptr
	cefBoxLayoutRefAPI().SysCallN(1, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBoxLayoutRef(resultPtr)
	return
}

// NewBoxLayoutRef class constructor
func NewBoxLayoutRef(data uintptr) ICefBoxLayoutRef {
	var boxLayoutPtr uintptr // ICefBoxLayout
	var layoutPtr uintptr    // ICefLayout
	r := cefBoxLayoutRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&boxLayoutPtr)), uintptr(base.UnsafePointer(&layoutPtr)))
	ret := AsCefBoxLayoutRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, boxLayoutPtr)
		intf.SetIntfPointer(1, layoutPtr)
	}
	return ret
}

var (
	cefBoxLayoutRefOnce   base.Once
	cefBoxLayoutRefImport *imports.Imports = nil
)

func cefBoxLayoutRefAPI() *imports.Imports {
	cefBoxLayoutRefOnce.Do(func() {
		cefBoxLayoutRefImport = api.NewDefaultImports()
		cefBoxLayoutRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBoxLayoutRef_Create", 0), // constructor NewBoxLayoutRef
			/* 1 */ imports.NewTable("TCefBoxLayoutRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 2 */ imports.NewTable("TCefBoxLayoutRef_SetFlexForView", 0), // procedure SetFlexForView
			/* 3 */ imports.NewTable("TCefBoxLayoutRef_ClearFlexForView", 0), // procedure ClearFlexForView
		}
	})
	return cefBoxLayoutRefImport
}
