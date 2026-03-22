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

	cefTypes "github.com/energye/cef/types"
)

// IChromiumFontOptions Parent: IPersistent
type IChromiumFontOptions interface {
	IPersistent
	// StandardFontFamily
	//  Standard font family name.
	StandardFontFamily() string         // property StandardFontFamily Getter
	SetStandardFontFamily(value string) // property StandardFontFamily Setter
	// FixedFontFamily
	//  Fixed font family name.
	FixedFontFamily() string         // property FixedFontFamily Getter
	SetFixedFontFamily(value string) // property FixedFontFamily Setter
	// SerifFontFamily
	//  Serif font family name.
	SerifFontFamily() string         // property SerifFontFamily Getter
	SetSerifFontFamily(value string) // property SerifFontFamily Setter
	// SansSerifFontFamily
	//  SansSerif font family name.
	SansSerifFontFamily() string         // property SansSerifFontFamily Getter
	SetSansSerifFontFamily(value string) // property SansSerifFontFamily Setter
	// CursiveFontFamily
	//  Cursive font family name.
	CursiveFontFamily() string         // property CursiveFontFamily Getter
	SetCursiveFontFamily(value string) // property CursiveFontFamily Setter
	// FantasyFontFamily
	//  Fantasy font family name.
	FantasyFontFamily() string         // property FantasyFontFamily Getter
	SetFantasyFontFamily(value string) // property FantasyFontFamily Setter
	// DefaultFontSize
	//  Default font size.
	DefaultFontSize() int32         // property DefaultFontSize Getter
	SetDefaultFontSize(value int32) // property DefaultFontSize Setter
	// DefaultFixedFontSize
	//  Default fixed font size.
	DefaultFixedFontSize() int32         // property DefaultFixedFontSize Getter
	SetDefaultFixedFontSize(value int32) // property DefaultFixedFontSize Setter
	// MinimumFontSize
	//  Minimum font size.
	MinimumFontSize() int32         // property MinimumFontSize Getter
	SetMinimumFontSize(value int32) // property MinimumFontSize Setter
	// MinimumLogicalFontSize
	//  Minimum logical font size.
	MinimumLogicalFontSize() int32         // property MinimumLogicalFontSize Getter
	SetMinimumLogicalFontSize(value int32) // property MinimumLogicalFontSize Setter
	// RemoteFonts
	//  Controls the loading of fonts from remote sources. Also configurable using
	//  the "disable-remote-fonts" command-line switch.
	RemoteFonts() cefTypes.TCefState         // property RemoteFonts Getter
	SetRemoteFonts(value cefTypes.TCefState) // property RemoteFonts Setter
}

type TChromiumFontOptions struct {
	TPersistent
}

func (m *TChromiumFontOptions) StandardFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetStandardFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) FixedFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetFixedFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) SerifFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetSerifFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) SansSerifFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetSansSerifFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) CursiveFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetCursiveFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) FantasyFontFamily() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumFontOptionsAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumFontOptions) SetFantasyFontFamily(value string) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumFontOptions) DefaultFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumFontOptionsAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumFontOptions) SetDefaultFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumFontOptions) DefaultFixedFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumFontOptionsAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumFontOptions) SetDefaultFixedFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumFontOptions) MinimumFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumFontOptionsAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumFontOptions) SetMinimumFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumFontOptions) MinimumLogicalFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumFontOptionsAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumFontOptions) SetMinimumLogicalFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumFontOptions) RemoteFonts() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumFontOptionsAPI().SysCallN(11, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumFontOptions) SetRemoteFonts(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumFontOptionsAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

// NewChromiumFontOptions class constructor
func NewChromiumFontOptions() IChromiumFontOptions {
	r := chromiumFontOptionsAPI().SysCallN(0)
	return AsChromiumFontOptions(r)
}

var (
	chromiumFontOptionsOnce   base.Once
	chromiumFontOptionsImport *imports.Imports = nil
)

func chromiumFontOptionsAPI() *imports.Imports {
	chromiumFontOptionsOnce.Do(func() {
		chromiumFontOptionsImport = api.NewDefaultImports()
		chromiumFontOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChromiumFontOptions_Create", 0), // constructor NewChromiumFontOptions
			/* 1 */ imports.NewTable("TChromiumFontOptions_StandardFontFamily", 0), // property StandardFontFamily
			/* 2 */ imports.NewTable("TChromiumFontOptions_FixedFontFamily", 0), // property FixedFontFamily
			/* 3 */ imports.NewTable("TChromiumFontOptions_SerifFontFamily", 0), // property SerifFontFamily
			/* 4 */ imports.NewTable("TChromiumFontOptions_SansSerifFontFamily", 0), // property SansSerifFontFamily
			/* 5 */ imports.NewTable("TChromiumFontOptions_CursiveFontFamily", 0), // property CursiveFontFamily
			/* 6 */ imports.NewTable("TChromiumFontOptions_FantasyFontFamily", 0), // property FantasyFontFamily
			/* 7 */ imports.NewTable("TChromiumFontOptions_DefaultFontSize", 0), // property DefaultFontSize
			/* 8 */ imports.NewTable("TChromiumFontOptions_DefaultFixedFontSize", 0), // property DefaultFixedFontSize
			/* 9 */ imports.NewTable("TChromiumFontOptions_MinimumFontSize", 0), // property MinimumFontSize
			/* 10 */ imports.NewTable("TChromiumFontOptions_MinimumLogicalFontSize", 0), // property MinimumLogicalFontSize
			/* 11 */ imports.NewTable("TChromiumFontOptions_RemoteFonts", 0), // property RemoteFonts
		}
	})
	return chromiumFontOptionsImport
}
