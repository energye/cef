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
	"github.com/energye/lcl/types"
)

// IChromium Parent: IChromiumCore
type IChromium interface {
	IChromiumCore
	CreateBrowserWithWControlStrRContextDValue(browserParent lcl.IWinControl, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue) bool // function
	SaveAsBitmapStream(stream lcl.IStream) bool                                                                                                                  // function
	TakeSnapshot(bitmap *lcl.IBitmap) bool                                                                                                                       // function
	InitializeDragAndDropWithWinControl(dropTargetCtrl lcl.IWinControl)                                                                                          // procedure
	ShowDevToolsWithPointWinControl(inspectElementAt types.TPoint, devTools lcl.IWinControl)                                                                     // procedure
	CloseDevToolsWithWinControl(devTools lcl.IWinControl)                                                                                                        // procedure
	MoveFormTo(X int32, Y int32)                                                                                                                                 // procedure
	MoveFormBy(X int32, Y int32)                                                                                                                                 // procedure
	ResizeFormWidthTo(X int32)                                                                                                                                   // procedure
	ResizeFormHeightTo(Y int32)                                                                                                                                  // procedure
	SetFormLeftTo(X int32)                                                                                                                                       // procedure
	SetFormTopTo(Y int32)                                                                                                                                        // procedure
	AsIntfChromiumEvents() uintptr
}

type TChromium struct {
	TChromiumCore
}

func (m *TChromium) CreateBrowserWithWControlStrRContextDValue(browserParent lcl.IWinControl, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserParent), api.PasStr(windowName), base.GetObjectUintptr(context), base.GetObjectUintptr(extraInfo))
	return api.GoBool(r)
}

func (m *TChromium) SaveAsBitmapStream(stream lcl.IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(stream))
	return api.GoBool(r)
}

func (m *TChromium) TakeSnapshot(bitmap *lcl.IBitmap) bool {
	if !m.IsValid() {
		return false
	}
	bitmapPtr := base.GetObjectUintptr(*bitmap)
	r := chromiumAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&bitmapPtr)))
	*bitmap = lcl.AsBitmap(bitmapPtr)
	return api.GoBool(r)
}

func (m *TChromium) InitializeDragAndDropWithWinControl(dropTargetCtrl lcl.IWinControl) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(dropTargetCtrl))
}

func (m *TChromium) ShowDevToolsWithPointWinControl(inspectElementAt types.TPoint, devTools lcl.IWinControl) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&inspectElementAt)), base.GetObjectUintptr(devTools))
}

func (m *TChromium) CloseDevToolsWithWinControl(devTools lcl.IWinControl) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(devTools))
}

func (m *TChromium) MoveFormTo(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(7, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TChromium) MoveFormBy(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(8, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TChromium) ResizeFormWidthTo(X int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(9, m.Instance(), uintptr(X))
}

func (m *TChromium) ResizeFormHeightTo(Y int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(10, m.Instance(), uintptr(Y))
}

func (m *TChromium) SetFormLeftTo(X int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(11, m.Instance(), uintptr(X))
}

func (m *TChromium) SetFormTopTo(Y int32) {
	if !m.IsValid() {
		return
	}
	chromiumAPI().SysCallN(12, m.Instance(), uintptr(Y))
}

func (m *TChromium) AsIntfChromiumEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewChromium class constructor
func NewChromium(owner lcl.IComponent) IChromium {
	var chromiumEventsPtr uintptr // IChromiumEvents
	r := chromiumAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&chromiumEventsPtr)))
	ret := AsChromium(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, chromiumEventsPtr)
	}
	return ret
}

var (
	chromiumOnce   base.Once
	chromiumImport *imports.Imports = nil
)

func chromiumAPI() *imports.Imports {
	chromiumOnce.Do(func() {
		chromiumImport = api.NewDefaultImports()
		chromiumImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChromium_Create", 0), // constructor NewChromium
			/* 1 */ imports.NewTable("TChromium_CreateBrowserWithWControlStrRContextDValue", 0), // function CreateBrowserWithWControlStrRContextDValue
			/* 2 */ imports.NewTable("TChromium_SaveAsBitmapStream", 0), // function SaveAsBitmapStream
			/* 3 */ imports.NewTable("TChromium_TakeSnapshot", 0), // function TakeSnapshot
			/* 4 */ imports.NewTable("TChromium_InitializeDragAndDropWithWinControl", 0), // procedure InitializeDragAndDropWithWinControl
			/* 5 */ imports.NewTable("TChromium_ShowDevToolsWithPointWinControl", 0), // procedure ShowDevToolsWithPointWinControl
			/* 6 */ imports.NewTable("TChromium_CloseDevToolsWithWinControl", 0), // procedure CloseDevToolsWithWinControl
			/* 7 */ imports.NewTable("TChromium_MoveFormTo", 0), // procedure MoveFormTo
			/* 8 */ imports.NewTable("TChromium_MoveFormBy", 0), // procedure MoveFormBy
			/* 9 */ imports.NewTable("TChromium_ResizeFormWidthTo", 0), // procedure ResizeFormWidthTo
			/* 10 */ imports.NewTable("TChromium_ResizeFormHeightTo", 0), // procedure ResizeFormHeightTo
			/* 11 */ imports.NewTable("TChromium_SetFormLeftTo", 0), // procedure SetFormLeftTo
			/* 12 */ imports.NewTable("TChromium_SetFormTopTo", 0), // procedure SetFormTopTo
		}
	})
	return chromiumImport
}
