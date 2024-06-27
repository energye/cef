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

// IChromium Parent: IChromiumCore
//
//	VCL and LCL version of TChromiumCore that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IChromium interface {
	IChromiumCore
	AsChromiumEvents() IChromiumEvents
	// CreateBrowserByWinControl
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowserByWinControl(aBrowserParent IWinControl, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool // function
	// SaveAsBitmapStream
	//  Copy the DC to a bitmap stream. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	SaveAsBitmapStream(aStream IStream) bool // function
	// TakeSnapshot
	//  Copy the DC to a TBitmap. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	TakeSnapshot(aBitmap *IBitmap) bool // function
	// InitializeDragAndDropForIWinControl
	//  Used with browsers in OSR mode to initialize drag and drop in Windows.
	InitializeDragAndDropForIWinControl(aDropTargetCtrl IWinControl) // procedure
	// ShowDevToolsForIWinControl
	//  Open developer tools(DevTools) in its own browser. If inspectElementAt has a valid point
	//  with coordinates different than low(integer) then the element at the specified location
	//  will be inspected. If the DevTools browser is already open then it will be focused.
	ShowDevToolsForIWinControl(inspectElementAt *TPoint, aDevTools IWinControl) // procedure
	// CloseDevToolsForIWinControl
	//  Close the developer tools.
	CloseDevToolsForIWinControl(aDevTools IWinControl) // procedure
	// MoveFormTo
	//  Move the parent form to the x and y coordinates.
	MoveFormTo(x, y int32) // procedure
	// MoveFormBy
	//  Move the parent form adding x and y to the coordinates.
	MoveFormBy(x, y int32) // procedure
	// ResizeFormWidthTo
	//  Add x to the parent form width.
	ResizeFormWidthTo(x int32) // procedure
	// ResizeFormHeightTo
	//  Add y to the parent form height.
	ResizeFormHeightTo(y int32) // procedure
	// SetFormLeftTo
	//  Set the parent form left property to x.
	SetFormLeftTo(x int32) // procedure
	// SetFormTopTo
	//  Set the parent form top property to y.
	SetFormTopTo(y int32) // procedure
}

// TChromium Parent: TChromiumCore
//
//	VCL and LCL version of TChromiumCore that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TChromium struct {
	TChromiumCore
}

func NewChromium(aOwner IComponent) IChromium {
	r1 := chromiumImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsChromium(r1)
}

func (m *TChromium) CreateBrowserByWinControl(aBrowserParent IWinControl, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool {
	r1 := chromiumImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(aBrowserParent), PascalStr(aWindowName), GetObjectUintptr(aContext), GetObjectUintptr(aExtraInfo))
	return GoBool(r1)
}

func (m *TChromium) SaveAsBitmapStream(aStream IStream) bool {
	r1 := chromiumImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(aStream))
	return GoBool(r1)
}

func (m *TChromium) TakeSnapshot(aBitmap *IBitmap) bool {
	var result0 uintptr
	r1 := chromiumImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&result0)))
	*aBitmap = AsBitmap(result0)
	return GoBool(r1)
}

func (m *TChromium) InitializeDragAndDropForIWinControl(aDropTargetCtrl IWinControl) {
	chromiumImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(aDropTargetCtrl))
}

func (m *TChromium) ShowDevToolsForIWinControl(inspectElementAt *TPoint, aDevTools IWinControl) {
	chromiumImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(inspectElementAt)), GetObjectUintptr(aDevTools))
}

func (m *TChromium) CloseDevToolsForIWinControl(aDevTools IWinControl) {
	chromiumImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(aDevTools))
}

func (m *TChromium) MoveFormTo(x, y int32) {
	chromiumImportAPI().SysCallN(5, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TChromium) MoveFormBy(x, y int32) {
	chromiumImportAPI().SysCallN(4, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TChromium) ResizeFormWidthTo(x int32) {
	chromiumImportAPI().SysCallN(7, m.Instance(), uintptr(x))
}

func (m *TChromium) ResizeFormHeightTo(y int32) {
	chromiumImportAPI().SysCallN(6, m.Instance(), uintptr(y))
}

func (m *TChromium) SetFormLeftTo(x int32) {
	chromiumImportAPI().SysCallN(9, m.Instance(), uintptr(x))
}

func (m *TChromium) SetFormTopTo(y int32) {
	chromiumImportAPI().SysCallN(10, m.Instance(), uintptr(y))
}

var (
	chromiumImport       *imports.Imports = nil
	chromiumImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Chromium_CloseDevToolsForIWinControl", 0),
		/*1*/ imports.NewTable("Chromium_Create", 0),
		/*2*/ imports.NewTable("Chromium_CreateBrowserByWinControl", 0),
		/*3*/ imports.NewTable("Chromium_InitializeDragAndDropForIWinControl", 0),
		/*4*/ imports.NewTable("Chromium_MoveFormBy", 0),
		/*5*/ imports.NewTable("Chromium_MoveFormTo", 0),
		/*6*/ imports.NewTable("Chromium_ResizeFormHeightTo", 0),
		/*7*/ imports.NewTable("Chromium_ResizeFormWidthTo", 0),
		/*8*/ imports.NewTable("Chromium_SaveAsBitmapStream", 0),
		/*9*/ imports.NewTable("Chromium_SetFormLeftTo", 0),
		/*10*/ imports.NewTable("Chromium_SetFormTopTo", 0),
		/*11*/ imports.NewTable("Chromium_ShowDevToolsForIWinControl", 0),
		/*12*/ imports.NewTable("Chromium_TakeSnapshot", 0),
	}
)

func chromiumImportAPI() *imports.Imports {
	if chromiumImport == nil {
		chromiumImport = NewDefaultImports()
		chromiumImport.SetImportTable(chromiumImportTables)
		chromiumImportTables = nil
	}
	return chromiumImport
}
