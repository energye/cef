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

	cefTypes "github.com/energye/cef/127/types"
)

// IPDFPrintOptions Parent: IObject
type IPDFPrintOptions interface {
	IObject
	CopyToSettings(settings *TCefPdfPrintSettings) // procedure
	// Landscape
	//  Set to true for landscape mode or false for portrait mode.
	Landscape() bool         // property Landscape Getter
	SetLandscape(value bool) // property Landscape Setter
	// PrintBackground
	//  Set to true to print background graphics.
	PrintBackground() bool         // property PrintBackground Getter
	SetPrintBackground(value bool) // property PrintBackground Setter
	// PreferCSSPageSize
	//  Set to true to prefer page size as defined by css. Defaults to false,
	//  in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize() bool         // property PreferCSSPageSize Getter
	SetPreferCSSPageSize(value bool) // property PreferCSSPageSize Setter
	// PageRanges
	//  Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed
	//  in the document order, not in the order specified, and no more than once.
	//  Defaults to empty string, which implies the entire document is printed.
	//  The page numbers are quietly capped to actual page count of the document,
	//  and ranges beyond the end of the document are ignored. If this results in
	//  no pages to print, an error is reported. It is an error to specify a range
	//  with start greater than end.
	PageRanges() string         // property PageRanges Getter
	SetPageRanges(value string) // property PageRanges Setter
	// DisplayHeaderFooter
	//  Set to true to display the header and/or footer. Modify
	//  HeaderTemplate and/or FooterTemplate to customize the display.
	DisplayHeaderFooter() bool         // property DisplayHeaderFooter Getter
	SetDisplayHeaderFooter(value bool) // property DisplayHeaderFooter Setter
	// HeaderTemplate
	//  HTML template for the print header. Only displayed if
	//  DisplayHeaderFooter is true. Should be valid HTML markup with
	//  the following classes used to inject printing values into them:
	//  <code>
	//  - date: formatted print date
	//  - title: document title
	//  - url: document location
	//  - pageNumber: current page number
	//  - totalPages: total pages in the document
	//  </code>
	//  For example, "<span class=title></span>" would generate a span containing
	//  the title.
	HeaderTemplate() string         // property HeaderTemplate Getter
	SetHeaderTemplate(value string) // property HeaderTemplate Setter
	// FooterTemplate
	//  HTML template for the print footer. Only displayed if
	//  DisplayHeaderFooter is true. Uses the same format as
	//  HeaderTemplate.
	FooterTemplate() string         // property FooterTemplate Getter
	SetFooterTemplate(value string) // property FooterTemplate Setter
	// GenerateTaggedPDF
	//  Set to true to generate tagged (accessible) PDF.
	GenerateTaggedPDF() bool         // property GenerateTaggedPDF Getter
	SetGenerateTaggedPDF(value bool) // property GenerateTaggedPDF Setter
	// GenerateDocumentOutline
	//  Set to true to generate a document outline.
	GenerateDocumentOutline() bool         // property GenerateDocumentOutline Getter
	SetGenerateDocumentOutline(value bool) // property GenerateDocumentOutline Setter
	// Scale
	//  The percentage to scale the PDF by before printing (e.g. .5 is 50%).
	//  If this value is less than or equal to zero the default value of 1.0
	//  will be used.
	Scale() float64         // property Scale Getter
	SetScale(value float64) // property Scale Setter
	// ScalePct
	//  The percentage value to scale the PDF by before printing (e.g. 50 is 50%).
	ScalePct() float64         // property ScalePct Getter
	SetScalePct(value float64) // property ScalePct Setter
	// PaperWidthInch
	//  Output paper width in inches. If either of these values is less than or
	//  equal to zero then the default paper size (letter, 8.5 x 11 inches) will
	//  be used.
	PaperWidthInch() float64         // property PaperWidthInch Getter
	SetPaperWidthInch(value float64) // property PaperWidthInch Setter
	// PaperHeightInch
	//  Output paper height in inches. If either of these values is less than or
	//  equal to zero then the default paper size (letter, 8.5 x 11 inches) will
	//  be used.
	PaperHeightInch() float64         // property PaperHeightInch Getter
	SetPaperHeightInch(value float64) // property PaperHeightInch Setter
	// PaperWidthMM
	//  Output paper width in mm.
	PaperWidthMM() float64         // property PaperWidthMM Getter
	SetPaperWidthMM(value float64) // property PaperWidthMM Setter
	// PaperHeightMM
	//  Output paper height in mm.
	PaperHeightMM() float64         // property PaperHeightMM Getter
	SetPaperHeightMM(value float64) // property PaperHeightMM Setter
	// MarginType
	//  Margin type.
	MarginType() cefTypes.TCefPdfPrintMarginType         // property MarginType Getter
	SetMarginType(value cefTypes.TCefPdfPrintMarginType) // property MarginType Setter
	// MarginTopInch
	//  Top margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginTopInch() float64         // property MarginTopInch Getter
	SetMarginTopInch(value float64) // property MarginTopInch Setter
	// MarginRightInch
	//  Right margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginRightInch() float64         // property MarginRightInch Getter
	SetMarginRightInch(value float64) // property MarginRightInch Setter
	// MarginBottomInch
	//  Bottom margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginBottomInch() float64         // property MarginBottomInch Getter
	SetMarginBottomInch(value float64) // property MarginBottomInch Setter
	// MarginLeftInch
	//  Left margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginLeftInch() float64         // property MarginLeftInch Getter
	SetMarginLeftInch(value float64) // property MarginLeftInch Setter
	// MarginTopMM
	//  Top margin in mm.
	MarginTopMM() float64         // property MarginTopMM Getter
	SetMarginTopMM(value float64) // property MarginTopMM Setter
	// MarginRightMM
	//  Right margin in mm.
	MarginRightMM() float64         // property MarginRightMM Getter
	SetMarginRightMM(value float64) // property MarginRightMM Setter
	// MarginBottomMM
	//  Bottom margin in mm.
	MarginBottomMM() float64         // property MarginBottomMM Getter
	SetMarginBottomMM(value float64) // property MarginBottomMM Setter
	// MarginLeftMM
	//  Left margin in mm.
	MarginLeftMM() float64         // property MarginLeftMM Getter
	SetMarginLeftMM(value float64) // property MarginLeftMM Setter
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

func (m *TPDFPrintOptions) GenerateTaggedPDF() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetGenerateTaggedPDF(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) GenerateDocumentOutline() bool {
	if !m.IsValid() {
		return false
	}
	r := pDFPrintOptionsAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPDFPrintOptions) SetGenerateDocumentOutline(value bool) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TPDFPrintOptions) Scale() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetScale(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) ScalePct() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetScalePct(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(12, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperWidthInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(13, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperHeightInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(14, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperWidthMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(15, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(15, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) PaperHeightMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(16, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(16, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginType() cefTypes.TCefPdfPrintMarginType {
	if !m.IsValid() {
		return 0
	}
	r := pDFPrintOptionsAPI().SysCallN(17, 0, m.Instance())
	return cefTypes.TCefPdfPrintMarginType(r)
}

func (m *TPDFPrintOptions) SetMarginType(value cefTypes.TCefPdfPrintMarginType) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TPDFPrintOptions) MarginTopInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(18, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(18, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginRightInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(19, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(19, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginBottomInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(20, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(20, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginLeftInch() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(21, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftInch(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(21, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginTopMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(22, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(22, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginRightMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(23, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(23, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginBottomMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(24, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(24, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPDFPrintOptions) MarginLeftMM() (result float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(25, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftMM(value float64) {
	if !m.IsValid() {
		return
	}
	pDFPrintOptionsAPI().SysCallN(25, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
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
			/* 9 */ imports.NewTable("TPDFPrintOptions_GenerateTaggedPDF", 0), // property GenerateTaggedPDF
			/* 10 */ imports.NewTable("TPDFPrintOptions_GenerateDocumentOutline", 0), // property GenerateDocumentOutline
			/* 11 */ imports.NewTable("TPDFPrintOptions_Scale", 0), // property Scale
			/* 12 */ imports.NewTable("TPDFPrintOptions_ScalePct", 0), // property ScalePct
			/* 13 */ imports.NewTable("TPDFPrintOptions_PaperWidthInch", 0), // property PaperWidthInch
			/* 14 */ imports.NewTable("TPDFPrintOptions_PaperHeightInch", 0), // property PaperHeightInch
			/* 15 */ imports.NewTable("TPDFPrintOptions_PaperWidthMM", 0), // property PaperWidthMM
			/* 16 */ imports.NewTable("TPDFPrintOptions_PaperHeightMM", 0), // property PaperHeightMM
			/* 17 */ imports.NewTable("TPDFPrintOptions_MarginType", 0), // property MarginType
			/* 18 */ imports.NewTable("TPDFPrintOptions_MarginTopInch", 0), // property MarginTopInch
			/* 19 */ imports.NewTable("TPDFPrintOptions_MarginRightInch", 0), // property MarginRightInch
			/* 20 */ imports.NewTable("TPDFPrintOptions_MarginBottomInch", 0), // property MarginBottomInch
			/* 21 */ imports.NewTable("TPDFPrintOptions_MarginLeftInch", 0), // property MarginLeftInch
			/* 22 */ imports.NewTable("TPDFPrintOptions_MarginTopMM", 0), // property MarginTopMM
			/* 23 */ imports.NewTable("TPDFPrintOptions_MarginRightMM", 0), // property MarginRightMM
			/* 24 */ imports.NewTable("TPDFPrintOptions_MarginBottomMM", 0), // property MarginBottomMM
			/* 25 */ imports.NewTable("TPDFPrintOptions_MarginLeftMM", 0), // property MarginLeftMM
		}
	})
	return pDFPrintOptionsImport
}
