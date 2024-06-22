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

// IChromiumFontOptions Parent: IPersistent
//
//	The TChromiumFontOptions properties are used to fill the TCefBrowserSettings record which is used during the browser creation.
type IChromiumFontOptions interface {
	IPersistent
	// StandardFontFamily
	//  Standard font family name.
	StandardFontFamily() string // property
	// SetStandardFontFamily Set StandardFontFamily
	SetStandardFontFamily(AValue string) // property
	// FixedFontFamily
	//  Fixed font family name.
	FixedFontFamily() string // property
	// SetFixedFontFamily Set FixedFontFamily
	SetFixedFontFamily(AValue string) // property
	// SerifFontFamily
	//  Serif font family name.
	SerifFontFamily() string // property
	// SetSerifFontFamily Set SerifFontFamily
	SetSerifFontFamily(AValue string) // property
	// SansSerifFontFamily
	//  SansSerif font family name.
	SansSerifFontFamily() string // property
	// SetSansSerifFontFamily Set SansSerifFontFamily
	SetSansSerifFontFamily(AValue string) // property
	// CursiveFontFamily
	//  Cursive font family name.
	CursiveFontFamily() string // property
	// SetCursiveFontFamily Set CursiveFontFamily
	SetCursiveFontFamily(AValue string) // property
	// FantasyFontFamily
	//  Fantasy font family name.
	FantasyFontFamily() string // property
	// SetFantasyFontFamily Set FantasyFontFamily
	SetFantasyFontFamily(AValue string) // property
	// DefaultFontSize
	//  Default font size.
	DefaultFontSize() int32 // property
	// SetDefaultFontSize Set DefaultFontSize
	SetDefaultFontSize(AValue int32) // property
	// DefaultFixedFontSize
	//  Default fixed font size.
	DefaultFixedFontSize() int32 // property
	// SetDefaultFixedFontSize Set DefaultFixedFontSize
	SetDefaultFixedFontSize(AValue int32) // property
	// MinimumFontSize
	//  Minimum font size.
	MinimumFontSize() int32 // property
	// SetMinimumFontSize Set MinimumFontSize
	SetMinimumFontSize(AValue int32) // property
	// MinimumLogicalFontSize
	//  Minimum logical font size.
	MinimumLogicalFontSize() int32 // property
	// SetMinimumLogicalFontSize Set MinimumLogicalFontSize
	SetMinimumLogicalFontSize(AValue int32) // property
	// RemoteFonts
	//  Controls the loading of fonts from remote sources. Also configurable using
	//  the "disable-remote-fonts" command-line switch.
	RemoteFonts() TCefState // property
	// SetRemoteFonts Set RemoteFonts
	SetRemoteFonts(AValue TCefState) // property
}

// TChromiumFontOptions Parent: TPersistent
//
//	The TChromiumFontOptions properties are used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumFontOptions struct {
	TPersistent
}

func NewChromiumFontOptions() IChromiumFontOptions {
	r1 := chromiumFontOptionsImportAPI().SysCallN(0)
	return AsChromiumFontOptions(r1)
}

func (m *TChromiumFontOptions) StandardFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetStandardFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(11, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) FixedFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetFixedFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) SerifFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetSerifFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) SansSerifFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetSansSerifFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(9, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) CursiveFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetCursiveFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) FantasyFontFamily() string {
	r1 := chromiumFontOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetFantasyFontFamily(AValue string) {
	chromiumFontOptionsImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) DefaultFontSize() int32 {
	r1 := chromiumFontOptionsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetDefaultFontSize(AValue int32) {
	chromiumFontOptionsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) DefaultFixedFontSize() int32 {
	r1 := chromiumFontOptionsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetDefaultFixedFontSize(AValue int32) {
	chromiumFontOptionsImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) MinimumFontSize() int32 {
	r1 := chromiumFontOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetMinimumFontSize(AValue int32) {
	chromiumFontOptionsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) MinimumLogicalFontSize() int32 {
	r1 := chromiumFontOptionsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetMinimumLogicalFontSize(AValue int32) {
	chromiumFontOptionsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) RemoteFonts() TCefState {
	r1 := chromiumFontOptionsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumFontOptions) SetRemoteFonts(AValue TCefState) {
	chromiumFontOptionsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

var (
	chromiumFontOptionsImport       *imports.Imports = nil
	chromiumFontOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ChromiumFontOptions_Create", 0),
		/*1*/ imports.NewTable("ChromiumFontOptions_CursiveFontFamily", 0),
		/*2*/ imports.NewTable("ChromiumFontOptions_DefaultFixedFontSize", 0),
		/*3*/ imports.NewTable("ChromiumFontOptions_DefaultFontSize", 0),
		/*4*/ imports.NewTable("ChromiumFontOptions_FantasyFontFamily", 0),
		/*5*/ imports.NewTable("ChromiumFontOptions_FixedFontFamily", 0),
		/*6*/ imports.NewTable("ChromiumFontOptions_MinimumFontSize", 0),
		/*7*/ imports.NewTable("ChromiumFontOptions_MinimumLogicalFontSize", 0),
		/*8*/ imports.NewTable("ChromiumFontOptions_RemoteFonts", 0),
		/*9*/ imports.NewTable("ChromiumFontOptions_SansSerifFontFamily", 0),
		/*10*/ imports.NewTable("ChromiumFontOptions_SerifFontFamily", 0),
		/*11*/ imports.NewTable("ChromiumFontOptions_StandardFontFamily", 0),
	}
)

func chromiumFontOptionsImportAPI() *imports.Imports {
	if chromiumFontOptionsImport == nil {
		chromiumFontOptionsImport = NewDefaultImports()
		chromiumFontOptionsImport.SetImportTable(chromiumFontOptionsImportTables)
		chromiumFontOptionsImportTables = nil
	}
	return chromiumFontOptionsImport
}
