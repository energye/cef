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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/types"
)

// ICefDisplay Parent: ICefBaseRefCounted
type ICefDisplay interface {
	ICefBaseRefCounted
	// GetID
	//  Returns the unique identifier for this Display.
	GetID() int64 // function
	// GetDeviceScaleFactor
	//  Returns this Display's device pixel scale factor. This specifies how much
	//  the UI should be scaled when the actual output has more pixels than
	//  standard displays (which is around 100~120dpi). The potential return
	//  values differ by platform.
	GetDeviceScaleFactor() float32 // function
	// GetBounds
	//  Returns this Display's bounds in DIP screen coordinates. This is the full
	//  size of the display.
	GetBounds() TCefRect // function
	// GetWorkArea
	//  Returns this Display's work area in DIP screen coordinates. This excludes
	//  areas of the display that are occupied with window manager toolbars, etc.
	GetWorkArea() TCefRect // function
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

// ICefDisplayRef Parent: ICefDisplay ICefBaseRefCountedRef
type ICefDisplayRef interface {
	ICefDisplay
	ICefBaseRefCountedRef
	AsIntfDisplay() uintptr
}

type TCefDisplayRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDisplayRef) GetID() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDisplayRef) GetDeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDisplayRef) GetBounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDisplayRef) GetWorkArea() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDisplayRef) GetRotation() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDisplayRefAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCefDisplayRef) ConvertPointToPixels(point *TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(point)))
}

func (m *TCefDisplayRef) ConvertPointFromPixels(point *TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefDisplayRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(point)))
}

func (m *TCefDisplayRef) AsIntfDisplay() uintptr {
	return m.GetIntfPointer(0)
}

// DisplayRef  is static instance
var DisplayRef _DisplayRefClass

// _DisplayRefClass is class type defined by TCefDisplayRef
type _DisplayRefClass uintptr

// UnWrap
//
//	Returns a ICefDisplay instance using a PCefDisplay data pointer.
func (_DisplayRefClass) UnWrap(data uintptr) (result ICefDisplay) {
	var resultPtr uintptr
	cefDisplayRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

// Primary
//
//	Returns the primary Display.
func (_DisplayRefClass) Primary() (result ICefDisplay) {
	var resultPtr uintptr
	cefDisplayRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

// NearestPoint
//
//	Returns the Display nearest |point|. Set |input_pixel_coords| to true (1) if
//	|point| is in pixel screen coordinates instead of DIP screen coordinates.
func (_DisplayRefClass) NearestPoint(point TCefPoint, inputPixelCoords bool) (result ICefDisplay) {
	var resultPtr uintptr
	cefDisplayRefAPI().SysCallN(8, uintptr(base.UnsafePointer(&point)), api.PasBool(inputPixelCoords), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

// MatchingBounds
//
//	Returns the Display that most closely intersects |bounds|. Set
//	|input_pixel_coords| to true (1) if |bounds| is in pixel screen coordinates
//	instead of DIP screen coordinates.
func (_DisplayRefClass) MatchingBounds(bounds TCefRect, inputPixelCoords bool) (result ICefDisplay) {
	var resultPtr uintptr
	cefDisplayRefAPI().SysCallN(9, uintptr(base.UnsafePointer(&bounds)), api.PasBool(inputPixelCoords), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDisplayRef(resultPtr)
	return
}

// GetCount
//
//	Returns the total number of Displays. Mirrored displays are excluded; this
//	function is intended to return the number of distinct, usable displays.
func (_DisplayRefClass) GetCount() cefTypes.NativeUInt {
	r := cefDisplayRefAPI().SysCallN(10)
	return cefTypes.NativeUInt(r)
}

// GetAlls
//
//	Returns all Displays. Mirrored displays are excluded; this function is
//	intended to return distinct, usable displays.
func (_DisplayRefClass) GetAlls(displayArray *ICefDisplayArray) bool {
	var displayArrayPtr uintptr
	var displayArrayCountPtr uintptr
	r := cefDisplayRefAPI().SysCallN(11, uintptr(base.UnsafePointer(&displayArrayPtr)), uintptr(base.UnsafePointer(&displayArrayCountPtr)))
	*displayArray = NewCefDisplayArray(int(displayArrayCountPtr), displayArrayPtr)
	return api.GoBool(r)
}

// ScreenPointToPixels
//
//	Convert |point| from DIP screen coordinates to pixel screen coordinates.
//	This function is only used on Windows.
func (_DisplayRefClass) ScreenPointToPixels(screenPoint types.TPoint) (result types.TPoint) {
	cefDisplayRefAPI().SysCallN(12, uintptr(base.UnsafePointer(&screenPoint)), uintptr(base.UnsafePointer(&result)))
	return
}

// ScreenPointFromPixels
//
//	Convert |point| from pixel screen coordinates to DIP screen coordinates.
//	This function is only used on Windows.
func (_DisplayRefClass) ScreenPointFromPixels(pixelsPoint types.TPoint) (result types.TPoint) {
	cefDisplayRefAPI().SysCallN(13, uintptr(base.UnsafePointer(&pixelsPoint)), uintptr(base.UnsafePointer(&result)))
	return
}

// ScreenRectToPixels
//
//	Convert |rect| from DIP screen coordinates to pixel screen coordinates. This
//	function is only used on Windows.
func (_DisplayRefClass) ScreenRectToPixels(screenRect types.TRect) (result types.TRect) {
	cefDisplayRefAPI().SysCallN(14, uintptr(base.UnsafePointer(&screenRect)), uintptr(base.UnsafePointer(&result)))
	return
}

// ScreenRectFromPixels
//
//	Convert |rect| from pixel screen coordinates to DIP screen coordinates. This
//	function is only used on Windows.
func (_DisplayRefClass) ScreenRectFromPixels(pixelsRect types.TRect) (result types.TRect) {
	cefDisplayRefAPI().SysCallN(15, uintptr(base.UnsafePointer(&pixelsRect)), uintptr(base.UnsafePointer(&result)))
	return
}

// NewDisplayRef class constructor
func NewDisplayRef(data uintptr) ICefDisplayRef {
	var displayPtr uintptr // ICefDisplay
	r := cefDisplayRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&displayPtr)))
	ret := AsCefDisplayRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, displayPtr)
	}
	return ret
}

var (
	cefDisplayRefOnce   base.Once
	cefDisplayRefImport *imports.Imports = nil
)

func cefDisplayRefAPI() *imports.Imports {
	cefDisplayRefOnce.Do(func() {
		cefDisplayRefImport = api.NewDefaultImports()
		cefDisplayRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDisplayRef_Create", 0), // constructor NewDisplayRef
			/* 1 */ imports.NewTable("TCefDisplayRef_GetID", 0), // function GetID
			/* 2 */ imports.NewTable("TCefDisplayRef_GetDeviceScaleFactor", 0), // function GetDeviceScaleFactor
			/* 3 */ imports.NewTable("TCefDisplayRef_GetBounds", 0), // function GetBounds
			/* 4 */ imports.NewTable("TCefDisplayRef_GetWorkArea", 0), // function GetWorkArea
			/* 5 */ imports.NewTable("TCefDisplayRef_GetRotation", 0), // function GetRotation
			/* 6 */ imports.NewTable("TCefDisplayRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefDisplayRef_Primary", 0), // static function Primary
			/* 8 */ imports.NewTable("TCefDisplayRef_NearestPoint", 0), // static function NearestPoint
			/* 9 */ imports.NewTable("TCefDisplayRef_MatchingBounds", 0), // static function MatchingBounds
			/* 10 */ imports.NewTable("TCefDisplayRef_GetCount", 0), // static function GetCount
			/* 11 */ imports.NewTable("TCefDisplayRef_GetAlls", 0), // static function GetAlls
			/* 12 */ imports.NewTable("TCefDisplayRef_ScreenPointToPixels", 0), // static function ScreenPointToPixels
			/* 13 */ imports.NewTable("TCefDisplayRef_ScreenPointFromPixels", 0), // static function ScreenPointFromPixels
			/* 14 */ imports.NewTable("TCefDisplayRef_ScreenRectToPixels", 0), // static function ScreenRectToPixels
			/* 15 */ imports.NewTable("TCefDisplayRef_ScreenRectFromPixels", 0), // static function ScreenRectFromPixels
			/* 16 */ imports.NewTable("TCefDisplayRef_ConvertPointToPixels", 0), // procedure ConvertPointToPixels
			/* 17 */ imports.NewTable("TCefDisplayRef_ConvertPointFromPixels", 0), // procedure ConvertPointFromPixels
		}
	})
	return cefDisplayRefImport
}
