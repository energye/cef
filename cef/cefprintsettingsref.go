//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefPrintSettings Parent: ICefBaseRefCounted
type ICefPrintSettings interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions
	//  if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may
	//  expose read-only objects.
	IsReadOnly() bool // function
	// IsLandscape
	//  Returns true (1) if the orientation is landscape.
	IsLandscape() bool // function
	// GetDeviceName
	//  Get the device name.
	GetDeviceName() string // function
	// GetDpi
	//  Get the DPI (dots per inch).
	GetDpi() int32 // function
	// GetPageRangesCount
	//  Returns the number of page ranges that currently exist.
	GetPageRangesCount() cefTypes.NativeUInt // function
	// IsSelectionOnly
	//  Returns true (1) if only the selection will be printed.
	IsSelectionOnly() bool // function
	// WillCollate
	//  Returns true (1) if pages will be collated.
	WillCollate() bool // function
	// GetColorModel
	//  Get the color model.
	GetColorModel() cefTypes.TCefColorModel // function
	// GetCopies
	//  Get the number of copies.
	GetCopies() int32 // function
	// GetDuplexMode
	//  Get the duplex mode.
	GetDuplexMode() cefTypes.TCefDuplexMode // function
	// SetOrientation
	//  Set the page orientation.
	SetOrientation(landscape bool) // procedure
	// SetPrinterPrintableArea
	//  Set the printer printable area in device units. Some platforms already
	//  provide flipped area. Set |landscape_needs_flip| to false (0) on those
	//  platforms to avoid double flipping.
	SetPrinterPrintableArea(physicalSizeDeviceUnits TCefSize, printableAreaDeviceUnits TCefRect, landscapeNeedsFlip bool) // procedure
	// SetDeviceName
	//  Set the device name.
	SetDeviceName(name string) // procedure
	// SetDpi
	//  Set the DPI (dots per inch).
	SetDpi(dpi int32) // procedure
	// SetPageRanges
	//  Set the page ranges.
	SetPageRanges(ranges ICefRangeArray) // procedure
	// GetPageRanges
	//  Retrieve the page ranges.
	GetPageRanges(outRanges *ICefRangeArray) // procedure
	// SetSelectionOnly
	//  Set whether only the selection will be printed.
	SetSelectionOnly(selectionOnly bool) // procedure
	// SetCollate
	//  Set whether pages will be collated.
	SetCollate(collate bool) // procedure
	// SetColorModel
	//  Set the color model.
	SetColorModel(model cefTypes.TCefColorModel) // procedure
	// SetCopies
	//  Set the number of copies.
	SetCopies(copies int32) // procedure
	// SetDuplexMode
	//  Set the duplex mode.
	SetDuplexMode(mode cefTypes.TCefDuplexMode) // procedure
}

// ICefPrintSettingsRef Parent: ICefPrintSettings ICefBaseRefCountedRef
type ICefPrintSettingsRef interface {
	ICefPrintSettings
	ICefBaseRefCountedRef
	AsIntfPrintSettings() uintptr
}

type TCefPrintSettingsRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPrintSettingsRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefPrintSettingsRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPrintSettingsRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPrintSettingsRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPrintSettingsRef) IsLandscape() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPrintSettingsRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPrintSettingsRef) GetDeviceName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefPrintSettingsRefAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCefPrintSettingsRef) GetDpi() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefPrintSettingsRefAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCefPrintSettingsRef) GetPageRangesCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefPrintSettingsRefAPI().SysCallN(6, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefPrintSettingsRef) IsSelectionOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPrintSettingsRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPrintSettingsRef) WillCollate() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPrintSettingsRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPrintSettingsRef) GetColorModel() cefTypes.TCefColorModel {
	if !m.IsValid() {
		return 0
	}
	r := cefPrintSettingsRefAPI().SysCallN(9, m.Instance())
	return cefTypes.TCefColorModel(r)
}

func (m *TCefPrintSettingsRef) GetCopies() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefPrintSettingsRefAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TCefPrintSettingsRef) GetDuplexMode() cefTypes.TCefDuplexMode {
	if !m.IsValid() {
		return 0
	}
	r := cefPrintSettingsRefAPI().SysCallN(11, m.Instance())
	return cefTypes.TCefDuplexMode(r)
}

func (m *TCefPrintSettingsRef) SetOrientation(landscape bool) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(14, m.Instance(), api.PasBool(landscape))
}

func (m *TCefPrintSettingsRef) SetPrinterPrintableArea(physicalSizeDeviceUnits TCefSize, printableAreaDeviceUnits TCefRect, landscapeNeedsFlip bool) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&physicalSizeDeviceUnits)), uintptr(base.UnsafePointer(&printableAreaDeviceUnits)), api.PasBool(landscapeNeedsFlip))
}

func (m *TCefPrintSettingsRef) SetDeviceName(name string) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(16, m.Instance(), api.PasStr(name))
}

func (m *TCefPrintSettingsRef) SetDpi(dpi int32) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(17, m.Instance(), uintptr(dpi))
}

func (m *TCefPrintSettingsRef) SetPageRanges(ranges ICefRangeArray) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(18, m.Instance(), ranges.Instance(), uintptr(int32(ranges.Count())))
}

func (m *TCefPrintSettingsRef) GetPageRanges(outRanges *ICefRangeArray) {
	if !m.IsValid() {
		return
	}
	var rangesPtr uintptr
	var rangesCountPtr uintptr
	cefPrintSettingsRefAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&rangesPtr)), uintptr(base.UnsafePointer(&rangesCountPtr)))
	*outRanges = NewCefRangeArray(int(rangesCountPtr), rangesPtr)
}

func (m *TCefPrintSettingsRef) SetSelectionOnly(selectionOnly bool) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(20, m.Instance(), api.PasBool(selectionOnly))
}

func (m *TCefPrintSettingsRef) SetCollate(collate bool) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(21, m.Instance(), api.PasBool(collate))
}

func (m *TCefPrintSettingsRef) SetColorModel(model cefTypes.TCefColorModel) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(22, m.Instance(), uintptr(model))
}

func (m *TCefPrintSettingsRef) SetCopies(copies int32) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(23, m.Instance(), uintptr(copies))
}

func (m *TCefPrintSettingsRef) SetDuplexMode(mode cefTypes.TCefDuplexMode) {
	if !m.IsValid() {
		return
	}
	cefPrintSettingsRefAPI().SysCallN(24, m.Instance(), uintptr(mode))
}

func (m *TCefPrintSettingsRef) AsIntfPrintSettings() uintptr {
	return m.GetIntfPointer(0)
}

// PrintSettingsRef  is static instance
var PrintSettingsRef _PrintSettingsRefClass

// _PrintSettingsRefClass is class type defined by TCefPrintSettingsRef
type _PrintSettingsRefClass uintptr

func (_PrintSettingsRefClass) New() (result ICefPrintSettings) {
	var resultPtr uintptr
	cefPrintSettingsRefAPI().SysCallN(12, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPrintSettingsRef(resultPtr)
	return
}

func (_PrintSettingsRefClass) UnWrap(data uintptr) (result ICefPrintSettings) {
	var resultPtr uintptr
	cefPrintSettingsRefAPI().SysCallN(13, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPrintSettingsRef(resultPtr)
	return
}

// NewPrintSettingsRef class constructor
func NewPrintSettingsRef(data uintptr) ICefPrintSettingsRef {
	var printSettingsPtr uintptr // ICefPrintSettings
	r := cefPrintSettingsRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&printSettingsPtr)))
	ret := AsCefPrintSettingsRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printSettingsPtr)
	}
	return ret
}

var (
	cefPrintSettingsRefOnce   base.Once
	cefPrintSettingsRefImport *imports.Imports = nil
)

func cefPrintSettingsRefAPI() *imports.Imports {
	cefPrintSettingsRefOnce.Do(func() {
		cefPrintSettingsRefImport = api.NewDefaultImports()
		cefPrintSettingsRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPrintSettingsRef_Create", 0), // constructor NewPrintSettingsRef
			/* 1 */ imports.NewTable("TCefPrintSettingsRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefPrintSettingsRef_IsReadOnly", 0), // function IsReadOnly
			/* 3 */ imports.NewTable("TCefPrintSettingsRef_IsLandscape", 0), // function IsLandscape
			/* 4 */ imports.NewTable("TCefPrintSettingsRef_GetDeviceName", 0), // function GetDeviceName
			/* 5 */ imports.NewTable("TCefPrintSettingsRef_GetDpi", 0), // function GetDpi
			/* 6 */ imports.NewTable("TCefPrintSettingsRef_GetPageRangesCount", 0), // function GetPageRangesCount
			/* 7 */ imports.NewTable("TCefPrintSettingsRef_IsSelectionOnly", 0), // function IsSelectionOnly
			/* 8 */ imports.NewTable("TCefPrintSettingsRef_WillCollate", 0), // function WillCollate
			/* 9 */ imports.NewTable("TCefPrintSettingsRef_GetColorModel", 0), // function GetColorModel
			/* 10 */ imports.NewTable("TCefPrintSettingsRef_GetCopies", 0), // function GetCopies
			/* 11 */ imports.NewTable("TCefPrintSettingsRef_GetDuplexMode", 0), // function GetDuplexMode
			/* 12 */ imports.NewTable("TCefPrintSettingsRef_New", 0), // static function New
			/* 13 */ imports.NewTable("TCefPrintSettingsRef_UnWrap", 0), // static function UnWrap
			/* 14 */ imports.NewTable("TCefPrintSettingsRef_SetOrientation", 0), // procedure SetOrientation
			/* 15 */ imports.NewTable("TCefPrintSettingsRef_SetPrinterPrintableArea", 0), // procedure SetPrinterPrintableArea
			/* 16 */ imports.NewTable("TCefPrintSettingsRef_SetDeviceName", 0), // procedure SetDeviceName
			/* 17 */ imports.NewTable("TCefPrintSettingsRef_SetDpi", 0), // procedure SetDpi
			/* 18 */ imports.NewTable("TCefPrintSettingsRef_SetPageRanges", 0), // procedure SetPageRanges
			/* 19 */ imports.NewTable("TCefPrintSettingsRef_GetPageRanges", 0), // procedure GetPageRanges
			/* 20 */ imports.NewTable("TCefPrintSettingsRef_SetSelectionOnly", 0), // procedure SetSelectionOnly
			/* 21 */ imports.NewTable("TCefPrintSettingsRef_SetCollate", 0), // procedure SetCollate
			/* 22 */ imports.NewTable("TCefPrintSettingsRef_SetColorModel", 0), // procedure SetColorModel
			/* 23 */ imports.NewTable("TCefPrintSettingsRef_SetCopies", 0), // procedure SetCopies
			/* 24 */ imports.NewTable("TCefPrintSettingsRef_SetDuplexMode", 0), // procedure SetDuplexMode
		}
	})
	return cefPrintSettingsRefImport
}
