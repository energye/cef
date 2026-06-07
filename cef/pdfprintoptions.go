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

// IPDFPrintOptions Parent: IObject
type IPDFPrintOptions interface {
	IObject
	CopyToSettings(settings *TCefPdfPrintSettings)       // procedure
	Landscape() bool                                     // property Landscape Getter
	SetLandscape(value bool)                             // property Landscape Setter
	PrintBackground() bool                               // property PrintBackground Getter
	SetPrintBackground(value bool)                       // property PrintBackground Setter
	PreferCSSPageSize() bool                             // property PreferCSSPageSize Getter
	SetPreferCSSPageSize(value bool)                     // property PreferCSSPageSize Setter
	PageRanges() string                                  // property PageRanges Getter
	SetPageRanges(value string)                          // property PageRanges Setter
	DisplayHeaderFooter() bool                           // property DisplayHeaderFooter Getter
	SetDisplayHeaderFooter(value bool)                   // property DisplayHeaderFooter Setter
	HeaderTemplate() string                              // property HeaderTemplate Getter
	SetHeaderTemplate(value string)                      // property HeaderTemplate Setter
	FooterTemplate() string                              // property FooterTemplate Getter
	SetFooterTemplate(value string)                      // property FooterTemplate Setter
	Scale() float64                                      // property Scale Getter
	SetScale(value float64)                              // property Scale Setter
	ScalePct() float64                                   // property ScalePct Getter
	SetScalePct(value float64)                           // property ScalePct Setter
	PaperWidthInch() float64                             // property PaperWidthInch Getter
	SetPaperWidthInch(value float64)                     // property PaperWidthInch Setter
	PaperHeightInch() float64                            // property PaperHeightInch Getter
	SetPaperHeightInch(value float64)                    // property PaperHeightInch Setter
	PaperWidthMM() float64                               // property PaperWidthMM Getter
	SetPaperWidthMM(value float64)                       // property PaperWidthMM Setter
	PaperHeightMM() float64                              // property PaperHeightMM Getter
	SetPaperHeightMM(value float64)                      // property PaperHeightMM Setter
	MarginType() cefTypes.TCefPdfPrintMarginType         // property MarginType Getter
	SetMarginType(value cefTypes.TCefPdfPrintMarginType) // property MarginType Setter
	MarginTopInch() float64                              // property MarginTopInch Getter
	SetMarginTopInch(value float64)                      // property MarginTopInch Setter
	MarginRightInch() float64                            // property MarginRightInch Getter
	SetMarginRightInch(value float64)                    // property MarginRightInch Setter
	MarginBottomInch() float64                           // property MarginBottomInch Getter
	SetMarginBottomInch(value float64)                   // property MarginBottomInch Setter
	MarginLeftInch() float64                             // property MarginLeftInch Getter
	SetMarginLeftInch(value float64)                     // property MarginLeftInch Setter
	MarginTopMM() float64                                // property MarginTopMM Getter
	SetMarginTopMM(value float64)                        // property MarginTopMM Setter
	MarginRightMM() float64                              // property MarginRightMM Getter
	SetMarginRightMM(value float64)                      // property MarginRightMM Setter
	MarginBottomMM() float64                             // property MarginBottomMM Getter
	SetMarginBottomMM(value float64)                     // property MarginBottomMM Setter
	MarginLeftMM() float64                               // property MarginLeftMM Getter
	SetMarginLeftMM(value float64)                       // property MarginLeftMM Setter
}

type TPDFPrintOptions struct {
	TObject
}

func (m *TPDFPrintOptions) CopyToSettings(settings *TCefPdfPrintSettings) {
	if !m.IsValid() {
		return
	}
	settingsPtr := settings.ToPas()
	pDFPrintOptionsAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(settingsPtr)))
	*settings = settingsPtr.ToGo()
}

func (m *TPDFPrintOptions) Landscape() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetLandscape(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) PrintBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetPrintBackground(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) PreferCSSPageSize() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetPreferCSSPageSize(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) PageRanges() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	pDFPrintOptionsAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TPDFPrintOptions) SetPageRanges(value string) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TPDFPrintOptions) DisplayHeaderFooter() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetDisplayHeaderFooter(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) HeaderTemplate() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	pDFPrintOptionsAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TPDFPrintOptions) SetHeaderTemplate(value string) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TPDFPrintOptions) FooterTemplate() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	pDFPrintOptionsAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TPDFPrintOptions) SetFooterTemplate(value string) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TPDFPrintOptions) Scale() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetScale(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(9, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) ScalePct() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetScalePct(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperWidthInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperHeightInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(12, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperWidthMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(13, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperHeightMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(14, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginType() cefTypes.TCefPdfPrintMarginType {
	if !m.IsValid() {
		return 0
	}
	r := pDFPrintOptionsAPI().SysCallN(15, 0, m.Instance())
	return cefTypes.TCefPdfPrintMarginType(r)
}

func (m *TPDFPrintOptions) SetMarginType(value cefTypes.TCefPdfPrintMarginType) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TPDFPrintOptions) MarginTopInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(16, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(16, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginRightInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(17, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(17, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginBottomInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(18, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(18, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginLeftInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(19, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(19, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginTopMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(20, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(20, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginRightMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(21, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(21, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginBottomMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(22, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(22, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginLeftMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(23, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(23, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewPDFPrintOptions class constructor
func NewPDFPrintOptions() IPDFPrintOptions {
	r := pDFPrintOptionsAPI().SysCallN(0)
	return AsPDFPrintOptions(r)
}

var (
	pDFPrintOptionsOnce   base.Once
	pDFPrintOptionsImport *imports.Imports = nil
)

func pDFPrintOptionsAPI() *imports.Imports {
	pDFPrintOptionsOnce.Do(func() {
		pDFPrintOptionsImport = api.NewDefaultImports()
		pDFPrintOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPDFPrintOptions_Create", 0), // constructor NewPDFPrintOptions
			/* 1 */ imports.NewTable("TPDFPrintOptions_CopyToSettings", 0), // procedure CopyToSettings
			/* 2 */ imports.NewTable("TPDFPrintOptions_Landscape", 0), // property Landscape
			/* 3 */ imports.NewTable("TPDFPrintOptions_PrintBackground", 0), // property PrintBackground
			/* 4 */ imports.NewTable("TPDFPrintOptions_PreferCSSPageSize", 0), // property PreferCSSPageSize
			/* 5 */ imports.NewTable("TPDFPrintOptions_PageRanges", 0), // property PageRanges
			/* 6 */ imports.NewTable("TPDFPrintOptions_DisplayHeaderFooter", 0), // property DisplayHeaderFooter
			/* 7 */ imports.NewTable("TPDFPrintOptions_HeaderTemplate", 0), // property HeaderTemplate
			/* 8 */ imports.NewTable("TPDFPrintOptions_FooterTemplate", 0), // property FooterTemplate
			/* 9 */ imports.NewTable("TPDFPrintOptions_Scale", 0), // property Scale
			/* 10 */ imports.NewTable("TPDFPrintOptions_ScalePct", 0), // property ScalePct
			/* 11 */ imports.NewTable("TPDFPrintOptions_PaperWidthInch", 0), // property PaperWidthInch
			/* 12 */ imports.NewTable("TPDFPrintOptions_PaperHeightInch", 0), // property PaperHeightInch
			/* 13 */ imports.NewTable("TPDFPrintOptions_PaperWidthMM", 0), // property PaperWidthMM
			/* 14 */ imports.NewTable("TPDFPrintOptions_PaperHeightMM", 0), // property PaperHeightMM
			/* 15 */ imports.NewTable("TPDFPrintOptions_MarginType", 0), // property MarginType
			/* 16 */ imports.NewTable("TPDFPrintOptions_MarginTopInch", 0), // property MarginTopInch
			/* 17 */ imports.NewTable("TPDFPrintOptions_MarginRightInch", 0), // property MarginRightInch
			/* 18 */ imports.NewTable("TPDFPrintOptions_MarginBottomInch", 0), // property MarginBottomInch
			/* 19 */ imports.NewTable("TPDFPrintOptions_MarginLeftInch", 0), // property MarginLeftInch
			/* 20 */ imports.NewTable("TPDFPrintOptions_MarginTopMM", 0), // property MarginTopMM
			/* 21 */ imports.NewTable("TPDFPrintOptions_MarginRightMM", 0), // property MarginRightMM
			/* 22 */ imports.NewTable("TPDFPrintOptions_MarginBottomMM", 0), // property MarginBottomMM
			/* 23 */ imports.NewTable("TPDFPrintOptions_MarginLeftMM", 0), // property MarginLeftMM
		}
	})
	return pDFPrintOptionsImport
}
