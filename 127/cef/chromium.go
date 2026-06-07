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
	// CreateBrowserWithWControlStrRContextDValue
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowserWithWControlStrRContextDValue(browserParent lcl.IWinControl, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue) bool // function
	// SaveAsBitmapStream
	//  Copy the DC to a bitmap stream. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	SaveAsBitmapStream(stream lcl.IStream) bool // function
	// TakeSnapshot
	//  Copy the DC to a TBitmap. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	TakeSnapshot(bitmap *lcl.IBitmap) bool // function
	// InitializeDragAndDropWithWinControl
	//  Used with browsers in OSR mode to initialize drag and drop in Windows.
	InitializeDragAndDropWithWinControl(dropTargetCtrl lcl.IWinControl) // procedure
	// ShowDevToolsWithPointWinControl
	//  Open developer tools (DevTools) in its own browser. If inspectElementAt has a valid point
	//  with coordinates different than low(integer) then the element at the specified location
	//  will be inspected. If the DevTools browser is already open then it will be focused.
	ShowDevToolsWithPointWinControl(inspectElementAt types.TPoint, devTools lcl.IWinControl) // procedure
	// CloseDevToolsWithWinControl
	//  Close the developer tools.
	CloseDevToolsWithWinControl(devTools lcl.IWinControl) // procedure
	// MoveFormTo
	//  Move the parent form to the x and y coordinates.
	MoveFormTo(X int32, Y int32) // procedure
	// MoveFormBy
	//  Move the parent form adding x and y to the coordinates.
	MoveFormBy(X int32, Y int32) // procedure
	// ResizeFormWidthTo
	//  Add x to the parent form width.
	ResizeFormWidthTo(X int32) // procedure
	// ResizeFormHeightTo
	//  Add y to the parent form height.
	ResizeFormHeightTo(Y int32) // procedure
	// SetFormLeftTo
	//  Set the parent form left property to x.
	SetFormLeftTo(X int32) // procedure
	// SetFormTopTo
	//  Set the parent form top property to y.
	SetFormTopTo(Y int32) // procedure
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
