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

// ICefPrintSettings Parent: ICefBaseRefCounted
//
//	Interface representing print settings.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_settings_capi.h">CEF source file: /include/capi/cef_print_settings_capi.h (cef_print_settings_t))</a>
type ICefPrintSettings interface {
	ICefBaseRefCounted
	// SetPageRanges
	//  Set the page ranges.
	SetPageRanges(ranges TRangeArray)
	// GetPageRanges
	//  Retrieve the page ranges.
	GetPageRanges() TRangeArray
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
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
	GetPageRangesCount() NativeUInt // function
	// IsSelectionOnly
	//  Returns true (1) if only the selection will be printed.
	IsSelectionOnly() bool // function
	// WillCollate
	//  Returns true (1) if pages will be collated.
	WillCollate() bool // function
	// GetColorModel
	//  Get the color model.
	GetColorModel() TCefColorModel // function
	// GetCopies
	//  Get the number of copies.
	GetCopies() int32 // function
	// GetDuplexMode
	//  Get the duplex mode.
	GetDuplexMode() TCefDuplexMode // function
	// SetOrientation
	//  Set the page orientation.
	SetOrientation(landscape bool) // procedure
	// SetPrinterPrintableArea
	//  Set the printer printable area in device units. Some platforms already provide flipped area. Set |landscape_needs_flip| to false (0) on those platforms to avoid double flipping.
	SetPrinterPrintableArea(physicalSizeDeviceUnits *TCefSize, printableAreaDeviceUnits *TCefRect, landscapeNeedsFlip bool) // procedure
	// SetDeviceName
	//  Set the device name.
	SetDeviceName(name string) // procedure
	// SetDpi
	//  Set the DPI (dots per inch).
	SetDpi(dpi int32) // procedure
	// SetSelectionOnly
	//  Set whether only the selection will be printed.
	SetSelectionOnly(selectionOnly bool) // procedure
	// SetCollate
	//  Set whether pages will be collated.
	SetCollate(collate bool) // procedure
	// SetColorModel
	//  Set the color model.
	SetColorModel(model TCefColorModel) // procedure
	// SetCopies
	//  Set the number of copies.
	SetCopies(copies int32) // procedure
	// SetDuplexMode
	//  Set the duplex mode.
	SetDuplexMode(mode TCefDuplexMode) // procedure
}

// TCefPrintSettings Parent: TCefBaseRefCounted
//
//	Interface representing print settings.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_settings_capi.h">CEF source file: /include/capi/cef_print_settings_capi.h (cef_print_settings_t))</a>
type TCefPrintSettings struct {
	TCefBaseRefCounted
}

// PrintSettingsRef -> ICefPrintSettings
var PrintSettingsRef printSettings

// printSettings TCefPrintSettings Ref
type printSettings uintptr

func (m *printSettings) New() ICefPrintSettings {
	var resultCefPrintSettings uintptr
	printSettingsImportAPI().SysCallN(10, uintptr(unsafePointer(&resultCefPrintSettings)))
	return AsCefPrintSettings(resultCefPrintSettings)
}

func (m *printSettings) UnWrap(data uintptr) ICefPrintSettings {
	var resultCefPrintSettings uintptr
	printSettingsImportAPI().SysCallN(20, uintptr(data), uintptr(unsafePointer(&resultCefPrintSettings)))
	return AsCefPrintSettings(resultCefPrintSettings)
}

func (m *TCefPrintSettings) IsValid() bool {
	r1 := printSettingsImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) IsReadOnly() bool {
	r1 := printSettingsImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) IsLandscape() bool {
	r1 := printSettingsImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) GetDeviceName() string {
	r1 := printSettingsImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefPrintSettings) GetDpi() int32 {
	r1 := printSettingsImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TCefPrintSettings) GetPageRangesCount() NativeUInt {
	r1 := printSettingsImportAPI().SysCallN(5, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPrintSettings) IsSelectionOnly() bool {
	r1 := printSettingsImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) WillCollate() bool {
	r1 := printSettingsImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) GetColorModel() TCefColorModel {
	r1 := printSettingsImportAPI().SysCallN(0, m.Instance())
	return TCefColorModel(r1)
}

func (m *TCefPrintSettings) GetCopies() int32 {
	r1 := printSettingsImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCefPrintSettings) GetDuplexMode() TCefDuplexMode {
	r1 := printSettingsImportAPI().SysCallN(4, m.Instance())
	return TCefDuplexMode(r1)
}

func (m *TCefPrintSettings) SetOrientation(landscape bool) {
	printSettingsImportAPI().SysCallN(17, m.Instance(), PascalBool(landscape))
}

func (m *TCefPrintSettings) SetPrinterPrintableArea(physicalSizeDeviceUnits *TCefSize, printableAreaDeviceUnits *TCefRect, landscapeNeedsFlip bool) {
	printSettingsImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(physicalSizeDeviceUnits)), uintptr(unsafePointer(printableAreaDeviceUnits)), PascalBool(landscapeNeedsFlip))
}

func (m *TCefPrintSettings) SetDeviceName(name string) {
	printSettingsImportAPI().SysCallN(14, m.Instance(), PascalStr(name))
}

func (m *TCefPrintSettings) SetDpi(dpi int32) {
	printSettingsImportAPI().SysCallN(15, m.Instance(), uintptr(dpi))
}

func (m *TCefPrintSettings) SetSelectionOnly(selectionOnly bool) {
	printSettingsImportAPI().SysCallN(19, m.Instance(), PascalBool(selectionOnly))
}

func (m *TCefPrintSettings) SetCollate(collate bool) {
	printSettingsImportAPI().SysCallN(11, m.Instance(), PascalBool(collate))
}

func (m *TCefPrintSettings) SetColorModel(model TCefColorModel) {
	printSettingsImportAPI().SysCallN(12, m.Instance(), uintptr(model))
}

func (m *TCefPrintSettings) SetCopies(copies int32) {
	printSettingsImportAPI().SysCallN(13, m.Instance(), uintptr(copies))
}

func (m *TCefPrintSettings) SetDuplexMode(mode TCefDuplexMode) {
	printSettingsImportAPI().SysCallN(16, m.Instance(), uintptr(mode))
}

var (
	printSettingsImport       *imports.Imports = nil
	printSettingsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPrintSettings_GetColorModel", 0),
		/*1*/ imports.NewTable("CefPrintSettings_GetCopies", 0),
		/*2*/ imports.NewTable("CefPrintSettings_GetDeviceName", 0),
		/*3*/ imports.NewTable("CefPrintSettings_GetDpi", 0),
		/*4*/ imports.NewTable("CefPrintSettings_GetDuplexMode", 0),
		/*5*/ imports.NewTable("CefPrintSettings_GetPageRangesCount", 0),
		/*6*/ imports.NewTable("CefPrintSettings_IsLandscape", 0),
		/*7*/ imports.NewTable("CefPrintSettings_IsReadOnly", 0),
		/*8*/ imports.NewTable("CefPrintSettings_IsSelectionOnly", 0),
		/*9*/ imports.NewTable("CefPrintSettings_IsValid", 0),
		/*10*/ imports.NewTable("CefPrintSettings_New", 0),
		/*11*/ imports.NewTable("CefPrintSettings_SetCollate", 0),
		/*12*/ imports.NewTable("CefPrintSettings_SetColorModel", 0),
		/*13*/ imports.NewTable("CefPrintSettings_SetCopies", 0),
		/*14*/ imports.NewTable("CefPrintSettings_SetDeviceName", 0),
		/*15*/ imports.NewTable("CefPrintSettings_SetDpi", 0),
		/*16*/ imports.NewTable("CefPrintSettings_SetDuplexMode", 0),
		/*17*/ imports.NewTable("CefPrintSettings_SetOrientation", 0),
		/*18*/ imports.NewTable("CefPrintSettings_SetPrinterPrintableArea", 0),
		/*19*/ imports.NewTable("CefPrintSettings_SetSelectionOnly", 0),
		/*20*/ imports.NewTable("CefPrintSettings_UnWrap", 0),
		/*21*/ imports.NewTable("CefPrintSettings_WillCollate", 0),
	}
)

func printSettingsImportAPI() *imports.Imports {
	if printSettingsImport == nil {
		printSettingsImport = NewDefaultImports()
		printSettingsImport.SetImportTable(printSettingsImportTables)
		printSettingsImportTables = nil
	}
	return printSettingsImport
}
