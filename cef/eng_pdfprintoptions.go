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

// IPDFPrintOptions Parent: IObject
//
//	The TPDFPrintOptions properties are used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
type IPDFPrintOptions interface {
	IObject
	// Landscape
	//  Set to true for landscape mode or false for portrait mode.
	Landscape() bool // property
	// SetLandscape Set Landscape
	SetLandscape(AValue bool) // property
	// PrintBackground
	//  Set to true to print background graphics.
	PrintBackground() bool // property
	// SetPrintBackground Set PrintBackground
	SetPrintBackground(AValue bool) // property
	// PreferCSSPageSize
	//  Set to true to prefer page size as defined by css. Defaults to false,
	//  in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize() bool // property
	// SetPreferCSSPageSize Set PreferCSSPageSize
	SetPreferCSSPageSize(AValue bool) // property
	// PageRanges
	//  Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed
	//  in the document order, not in the order specified, and no more than once.
	//  Defaults to empty string, which implies the entire document is printed.
	//  The page numbers are quietly capped to actual page count of the document,
	//  and ranges beyond the end of the document are ignored. If this results in
	//  no pages to print, an error is reported. It is an error to specify a range
	//  with start greater than end.
	PageRanges() string // property
	// SetPageRanges Set PageRanges
	SetPageRanges(AValue string) // property
	// DisplayHeaderFooter
	//  Set to true to display the header and/or footer. Modify
	//  HeaderTemplate and/or FooterTemplate to customize the display.
	DisplayHeaderFooter() bool // property
	// SetDisplayHeaderFooter Set DisplayHeaderFooter
	SetDisplayHeaderFooter(AValue bool) // property
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
	HeaderTemplate() string // property
	// SetHeaderTemplate Set HeaderTemplate
	SetHeaderTemplate(AValue string) // property
	// FooterTemplate
	//  HTML template for the print footer. Only displayed if
	//  DisplayHeaderFooter is true. Uses the same format as
	//  HeaderTemplate.
	FooterTemplate() string // property
	// SetFooterTemplate Set FooterTemplate
	SetFooterTemplate(AValue string) // property
	// GenerateTaggedPDF
	//  Set to true to generate tagged(accessible) PDF.
	GenerateTaggedPDF() bool // property
	// SetGenerateTaggedPDF Set GenerateTaggedPDF
	SetGenerateTaggedPDF(AValue bool) // property
	// Scale
	//  The percentage to scale the PDF by before printing(e.g. .5 is 50%).
	//  If this value is less than or equal to zero the default value of 1.0
	//  will be used.
	Scale() (resultFloat64 float64) // property
	// SetScale Set Scale
	SetScale(AValue float64) // property
	// ScalePct
	//  The percentage value to scale the PDF by before printing(e.g. 50 is 50%).
	ScalePct() (resultFloat64 float64) // property
	// SetScalePct Set ScalePct
	SetScalePct(AValue float64) // property
	// PaperWidthInch
	//  Output paper width in inches. If either of these values is less than or
	//  equal to zero then the default paper size(letter, 8.5 x 11 inches) will
	//  be used.
	PaperWidthInch() (resultFloat64 float64) // property
	// SetPaperWidthInch Set PaperWidthInch
	SetPaperWidthInch(AValue float64) // property
	// PaperHeightInch
	//  Output paper height in inches. If either of these values is less than or
	//  equal to zero then the default paper size(letter, 8.5 x 11 inches) will
	//  be used.
	PaperHeightInch() (resultFloat64 float64) // property
	// SetPaperHeightInch Set PaperHeightInch
	SetPaperHeightInch(AValue float64) // property
	// PaperWidthMM
	//  Output paper width in mm.
	PaperWidthMM() (resultFloat64 float64) // property
	// SetPaperWidthMM Set PaperWidthMM
	SetPaperWidthMM(AValue float64) // property
	// PaperHeightMM
	//  Output paper height in mm.
	PaperHeightMM() (resultFloat64 float64) // property
	// SetPaperHeightMM Set PaperHeightMM
	SetPaperHeightMM(AValue float64) // property
	// MarginType
	//  Margin type.
	MarginType() TCefPdfPrintMarginType // property
	// SetMarginType Set MarginType
	SetMarginType(AValue TCefPdfPrintMarginType) // property
	// MarginTopInch
	//  Top margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginTopInch() (resultFloat64 float64) // property
	// SetMarginTopInch Set MarginTopInch
	SetMarginTopInch(AValue float64) // property
	// MarginRightInch
	//  Right margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginRightInch() (resultFloat64 float64) // property
	// SetMarginRightInch Set MarginRightInch
	SetMarginRightInch(AValue float64) // property
	// MarginBottomInch
	//  Bottom margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginBottomInch() (resultFloat64 float64) // property
	// SetMarginBottomInch Set MarginBottomInch
	SetMarginBottomInch(AValue float64) // property
	// MarginLeftInch
	//  Left margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginLeftInch() (resultFloat64 float64) // property
	// SetMarginLeftInch Set MarginLeftInch
	SetMarginLeftInch(AValue float64) // property
	// MarginTopMM
	//  Top margin in mm.
	MarginTopMM() (resultFloat64 float64) // property
	// SetMarginTopMM Set MarginTopMM
	SetMarginTopMM(AValue float64) // property
	// MarginRightMM
	//  Right margin in mm.
	MarginRightMM() (resultFloat64 float64) // property
	// SetMarginRightMM Set MarginRightMM
	SetMarginRightMM(AValue float64) // property
	// MarginBottomMM
	//  Bottom margin in mm.
	MarginBottomMM() (resultFloat64 float64) // property
	// SetMarginBottomMM Set MarginBottomMM
	SetMarginBottomMM(AValue float64) // property
	// MarginLeftMM
	//  Left margin in mm.
	MarginLeftMM() (resultFloat64 float64) // property
	// SetMarginLeftMM Set MarginLeftMM
	SetMarginLeftMM(AValue float64) // property
	// CopyToSettings
	//  Copy the fields of this class to the TCefPdfPrintSettings parameter.
	CopyToSettings(aSettings *TCefPdfPrintSettings) // procedure
}

// TPDFPrintOptions Parent: TObject
//
//	The TPDFPrintOptions properties are used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
type TPDFPrintOptions struct {
	TObject
}

func NewPDFPrintOptions() IPDFPrintOptions {
	r1 := pDFPrintOptionsImportAPI().SysCallN(1)
	return AsPDFPrintOptions(r1)
}

func (m *TPDFPrintOptions) Landscape() bool {
	r1 := pDFPrintOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetLandscape(AValue bool) {
	pDFPrintOptionsImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PrintBackground() bool {
	r1 := pDFPrintOptionsImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetPrintBackground(AValue bool) {
	pDFPrintOptionsImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PreferCSSPageSize() bool {
	r1 := pDFPrintOptionsImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetPreferCSSPageSize(AValue bool) {
	pDFPrintOptionsImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PageRanges() string {
	r1 := pDFPrintOptionsImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetPageRanges(AValue string) {
	pDFPrintOptionsImportAPI().SysCallN(16, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) DisplayHeaderFooter() bool {
	r1 := pDFPrintOptionsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetDisplayHeaderFooter(AValue bool) {
	pDFPrintOptionsImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) HeaderTemplate() string {
	r1 := pDFPrintOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetHeaderTemplate(AValue string) {
	pDFPrintOptionsImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) FooterTemplate() string {
	r1 := pDFPrintOptionsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetFooterTemplate(AValue string) {
	pDFPrintOptionsImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) GenerateTaggedPDF() bool {
	r1 := pDFPrintOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetGenerateTaggedPDF(AValue bool) {
	pDFPrintOptionsImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) Scale() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(23, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetScale(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(23, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) ScalePct() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(24, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetScalePct(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(24, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperWidthInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(19, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(19, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperHeightInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(17, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(17, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperWidthMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(20, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(20, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperHeightMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(18, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(18, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginType() TCefPdfPrintMarginType {
	r1 := pDFPrintOptionsImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TCefPdfPrintMarginType(r1)
}

func (m *TPDFPrintOptions) SetMarginType(AValue TCefPdfPrintMarginType) {
	pDFPrintOptionsImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TPDFPrintOptions) MarginTopInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(13, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(13, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginRightInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(11, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(11, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginBottomInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(7, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginLeftInch() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(9, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftInch(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(9, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginTopMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(14, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(14, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginRightMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(12, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(12, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginBottomMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(8, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginLeftMM() (resultFloat64 float64) {
	pDFPrintOptionsImportAPI().SysCallN(10, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftMM(AValue float64) {
	pDFPrintOptionsImportAPI().SysCallN(10, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) CopyToSettings(aSettings *TCefPdfPrintSettings) {
	var result0 uintptr
	pDFPrintOptionsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&result0)))
	*aSettings = *(*TCefPdfPrintSettings)(unsafePointer(result0))
}

var (
	pDFPrintOptionsImport       *imports.Imports = nil
	pDFPrintOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PDFPrintOptions_CopyToSettings", 0),
		/*1*/ imports.NewTable("PDFPrintOptions_Create", 0),
		/*2*/ imports.NewTable("PDFPrintOptions_DisplayHeaderFooter", 0),
		/*3*/ imports.NewTable("PDFPrintOptions_FooterTemplate", 0),
		/*4*/ imports.NewTable("PDFPrintOptions_GenerateTaggedPDF", 0),
		/*5*/ imports.NewTable("PDFPrintOptions_HeaderTemplate", 0),
		/*6*/ imports.NewTable("PDFPrintOptions_Landscape", 0),
		/*7*/ imports.NewTable("PDFPrintOptions_MarginBottomInch", 0),
		/*8*/ imports.NewTable("PDFPrintOptions_MarginBottomMM", 0),
		/*9*/ imports.NewTable("PDFPrintOptions_MarginLeftInch", 0),
		/*10*/ imports.NewTable("PDFPrintOptions_MarginLeftMM", 0),
		/*11*/ imports.NewTable("PDFPrintOptions_MarginRightInch", 0),
		/*12*/ imports.NewTable("PDFPrintOptions_MarginRightMM", 0),
		/*13*/ imports.NewTable("PDFPrintOptions_MarginTopInch", 0),
		/*14*/ imports.NewTable("PDFPrintOptions_MarginTopMM", 0),
		/*15*/ imports.NewTable("PDFPrintOptions_MarginType", 0),
		/*16*/ imports.NewTable("PDFPrintOptions_PageRanges", 0),
		/*17*/ imports.NewTable("PDFPrintOptions_PaperHeightInch", 0),
		/*18*/ imports.NewTable("PDFPrintOptions_PaperHeightMM", 0),
		/*19*/ imports.NewTable("PDFPrintOptions_PaperWidthInch", 0),
		/*20*/ imports.NewTable("PDFPrintOptions_PaperWidthMM", 0),
		/*21*/ imports.NewTable("PDFPrintOptions_PreferCSSPageSize", 0),
		/*22*/ imports.NewTable("PDFPrintOptions_PrintBackground", 0),
		/*23*/ imports.NewTable("PDFPrintOptions_Scale", 0),
		/*24*/ imports.NewTable("PDFPrintOptions_ScalePct", 0),
	}
)

func pDFPrintOptionsImportAPI() *imports.Imports {
	if pDFPrintOptionsImport == nil {
		pDFPrintOptionsImport = NewDefaultImports()
		pDFPrintOptionsImport.SetImportTable(pDFPrintOptionsImportTables)
		pDFPrintOptionsImportTables = nil
	}
	return pDFPrintOptionsImport
}
