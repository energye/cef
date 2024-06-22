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

// ICefDisplay Parent: ICefBaseRefCounted
//
//	This class typically, but not always, corresponds to a physical display
//	connected to the system. A fake Display may exist on a headless system, or a
//	Display may correspond to a remote, virtual display. All size and position
//	values are in density independent pixel (DIP) coordinates unless otherwise
//	indicated. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_display_capi.h">CEF source file: /include/capi/views/cef_display_capi.h (cef_display_t)</a>
type ICefDisplay interface {
	ICefBaseRefCounted
	// GetID
	//  Returns the unique identifier for this Display.
	GetID() (resultInt64 int64) // function
	// GetDeviceScaleFactor
	//  Returns this Display's device pixel scale factor. This specifies how much
	//  the UI should be scaled when the actual output has more pixels than
	//  standard displays(which is around 100~120dpi). The potential return
	//  values differ by platform.
	GetDeviceScaleFactor() (resultFloat32 float32) // function
	// GetBounds
	//  Returns this Display's bounds in DIP screen coordinates. This is the full
	//  size of the display.
	GetBounds() (resultCefRect TCefRect) // function
	// GetWorkArea
	//  Returns this Display's work area in DIP screen coordinates. This excludes
	//  areas of the display that are occupied with window manager toolbars, etc.
	GetWorkArea() (resultCefRect TCefRect) // function
	// GetRotation
	//  Returns this Display's rotation in degrees.
	GetRotation() int32 // function
	// ConvertPointToPixels
	//  Convert |point| from DIP coordinates to pixel coordinates using this
	//  Display's device scale factor.
	ConvertPointToPixels(point *TCefPoint) // procedure
	// ConvertPointFromPixels
	//  Convert |point| from pixel coordinates to DIP coordinates using this
	//  Display's device scale factor.
	ConvertPointFromPixels(point *TCefPoint) // procedure
}

// TCefDisplay Parent: TCefBaseRefCounted
//
//	This class typically, but not always, corresponds to a physical display
//	connected to the system. A fake Display may exist on a headless system, or a
//	Display may correspond to a remote, virtual display. All size and position
//	values are in density independent pixel (DIP) coordinates unless otherwise
//	indicated. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_display_capi.h">CEF source file: /include/capi/views/cef_display_capi.h (cef_display_t)</a>
type TCefDisplay struct {
	TCefBaseRefCounted
}

// DisplayRef -> ICefDisplay
var DisplayRef display

// display TCefDisplay Ref
type display uintptr

// UnWrap
//
//	Returns a ICefDisplay instance using a PCefDisplay data pointer.
func (m *display) UnWrap(data uintptr) ICefDisplay {
	var resultCefDisplay uintptr
	displayImportAPI().SysCallN(15, uintptr(data), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// Primary
//
//	Returns the primary Display.
func (m *display) Primary() ICefDisplay {
	var resultCefDisplay uintptr
	displayImportAPI().SysCallN(10, uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// NearestPoint
//
//	Returns the Display nearest |point|. Set |input_pixel_coords| to true(1) if
//	|point| is in pixel screen coordinates instead of DIP screen coordinates.
func (m *display) NearestPoint(point *TCefPoint, inputpixelcoords bool) ICefDisplay {
	var resultCefDisplay uintptr
	displayImportAPI().SysCallN(9, uintptr(unsafePointer(point)), PascalBool(inputpixelcoords), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// MatchingBounds
//
//	Returns the Display that most closely intersects |bounds|. Set
//	|input_pixel_coords| to true(1) if |bounds| is in pixel screen coordinates
//	instead of DIP screen coordinates.
func (m *display) MatchingBounds(bounds *TCefRect, inputpixelcoords bool) ICefDisplay {
	var resultCefDisplay uintptr
	displayImportAPI().SysCallN(8, uintptr(unsafePointer(bounds)), PascalBool(inputpixelcoords), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// GetCount
//
//	Returns the total number of Displays. Mirrored displays are excluded; this
//	function is intended to return the number of distinct, usable displays.
func (m *display) GetCount() NativeUInt {
	r1 := displayImportAPI().SysCallN(3)
	return NativeUInt(r1)
}

// ScreenPointToPixels
//
//	Convert |point| from DIP screen coordinates to pixel screen coordinates.
//	This function is only used on Windows.
func (m *display) ScreenPointToPixels(aScreenPoint *TPoint) (resultPoint TPoint) {
	displayImportAPI().SysCallN(12, uintptr(unsafePointer(aScreenPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

// ScreenPointFromPixels
//
//	Convert |point| from pixel screen coordinates to DIP screen coordinates.
//	This function is only used on Windows.
func (m *display) ScreenPointFromPixels(aPixelsPoint *TPoint) (resultPoint TPoint) {
	displayImportAPI().SysCallN(11, uintptr(unsafePointer(aPixelsPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

// ScreenRectToPixels
//
//	Convert |rect| from DIP screen coordinates to pixel screen coordinates. This
//	function is only used on Windows.
func (m *display) ScreenRectToPixels(aScreenRect *TRect) (resultRect TRect) {
	displayImportAPI().SysCallN(14, uintptr(unsafePointer(aScreenRect)), uintptr(unsafePointer(&resultRect)))
	return
}

// ScreenRectFromPixels
//
//	Convert |rect| from pixel screen coordinates to DIP screen coordinates. This
//	function is only used on Windows.
func (m *display) ScreenRectFromPixels(aPixelsRect *TRect) (resultRect TRect) {
	displayImportAPI().SysCallN(13, uintptr(unsafePointer(aPixelsRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCefDisplay) GetID() (resultInt64 int64) {
	displayImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDisplay) GetDeviceScaleFactor() (resultFloat32 float32) {
	displayImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TCefDisplay) GetBounds() (resultCefRect TCefRect) {
	displayImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDisplay) GetWorkArea() (resultCefRect TCefRect) {
	displayImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDisplay) GetRotation() int32 {
	r1 := displayImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TCefDisplay) ConvertPointToPixels(point *TCefPoint) {
	var result0 uintptr
	displayImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
}

func (m *TCefDisplay) ConvertPointFromPixels(point *TCefPoint) {
	var result0 uintptr
	displayImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
}

var (
	displayImport       *imports.Imports = nil
	displayImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDisplay_ConvertPointFromPixels", 0),
		/*1*/ imports.NewTable("CefDisplay_ConvertPointToPixels", 0),
		/*2*/ imports.NewTable("CefDisplay_GetBounds", 0),
		/*3*/ imports.NewTable("CefDisplay_GetCount", 0),
		/*4*/ imports.NewTable("CefDisplay_GetDeviceScaleFactor", 0),
		/*5*/ imports.NewTable("CefDisplay_GetID", 0),
		/*6*/ imports.NewTable("CefDisplay_GetRotation", 0),
		/*7*/ imports.NewTable("CefDisplay_GetWorkArea", 0),
		/*8*/ imports.NewTable("CefDisplay_MatchingBounds", 0),
		/*9*/ imports.NewTable("CefDisplay_NearestPoint", 0),
		/*10*/ imports.NewTable("CefDisplay_Primary", 0),
		/*11*/ imports.NewTable("CefDisplay_ScreenPointFromPixels", 0),
		/*12*/ imports.NewTable("CefDisplay_ScreenPointToPixels", 0),
		/*13*/ imports.NewTable("CefDisplay_ScreenRectFromPixels", 0),
		/*14*/ imports.NewTable("CefDisplay_ScreenRectToPixels", 0),
		/*15*/ imports.NewTable("CefDisplay_UnWrap", 0),
	}
)

func displayImportAPI() *imports.Imports {
	if displayImport == nil {
		displayImport = NewDefaultImports()
		displayImport.SetImportTable(displayImportTables)
		displayImportTables = nil
	}
	return displayImport
}
