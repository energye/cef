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

// ICefDomDocument Parent: ICefBaseRefCounted
type ICefDomDocument interface {
	ICefBaseRefCounted
	GetType() TCefDomDocumentType            // function
	GetDocument() ICefDomNode                // function
	GetBody() ICefDomNode                    // function
	GetHead() ICefDomNode                    // function
	GetTitle() string                        // function
	GetElementById(id string) ICefDomNode    // function
	GetFocusedNode() ICefDomNode             // function
	HasSelection() bool                      // function
	GetSelectionStartOffset() int32          // function
	GetSelectionEndOffset() int32            // function
	GetSelectionAsMarkup() string            // function
	GetSelectionAsText() string              // function
	GetBaseUrl() string                      // function
	GetCompleteUrl(partialURL string) string // function
}

// TCefDomDocument Parent: TCefBaseRefCounted
type TCefDomDocument struct {
	TCefBaseRefCounted
}

// DomDocumentRef -> ICefDomDocument
var DomDocumentRef domDocument

// domDocument TCefDomDocument Ref
type domDocument uintptr

func (m *domDocument) UnWrap(data uintptr) ICefDomDocument {
	var resultCefDomDocument uintptr
	domDocumentImportAPI().SysCallN(14, uintptr(data), uintptr(unsafePointer(&resultCefDomDocument)))
	return AsCefDomDocument(resultCefDomDocument)
}

func (m *TCefDomDocument) GetType() TCefDomDocumentType {
	r1 := domDocumentImportAPI().SysCallN(12, m.Instance())
	return TCefDomDocumentType(r1)
}

func (m *TCefDomDocument) GetDocument() ICefDomNode {
	var resultCefDomNode uintptr
	domDocumentImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetBody() ICefDomNode {
	var resultCefDomNode uintptr
	domDocumentImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetHead() ICefDomNode {
	var resultCefDomNode uintptr
	domDocumentImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetTitle() string {
	r1 := domDocumentImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetElementById(id string) ICefDomNode {
	var resultCefDomNode uintptr
	domDocumentImportAPI().SysCallN(4, m.Instance(), PascalStr(id), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetFocusedNode() ICefDomNode {
	var resultCefDomNode uintptr
	domDocumentImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) HasSelection() bool {
	r1 := domDocumentImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomDocument) GetSelectionStartOffset() int32 {
	r1 := domDocumentImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TCefDomDocument) GetSelectionEndOffset() int32 {
	r1 := domDocumentImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCefDomDocument) GetSelectionAsMarkup() string {
	r1 := domDocumentImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetSelectionAsText() string {
	r1 := domDocumentImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetBaseUrl() string {
	r1 := domDocumentImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetCompleteUrl(partialURL string) string {
	r1 := domDocumentImportAPI().SysCallN(2, m.Instance(), PascalStr(partialURL))
	return GoStr(r1)
}

var (
	domDocumentImport       *imports.Imports = nil
	domDocumentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDomDocument_GetBaseUrl", 0),
		/*1*/ imports.NewTable("CefDomDocument_GetBody", 0),
		/*2*/ imports.NewTable("CefDomDocument_GetCompleteUrl", 0),
		/*3*/ imports.NewTable("CefDomDocument_GetDocument", 0),
		/*4*/ imports.NewTable("CefDomDocument_GetElementById", 0),
		/*5*/ imports.NewTable("CefDomDocument_GetFocusedNode", 0),
		/*6*/ imports.NewTable("CefDomDocument_GetHead", 0),
		/*7*/ imports.NewTable("CefDomDocument_GetSelectionAsMarkup", 0),
		/*8*/ imports.NewTable("CefDomDocument_GetSelectionAsText", 0),
		/*9*/ imports.NewTable("CefDomDocument_GetSelectionEndOffset", 0),
		/*10*/ imports.NewTable("CefDomDocument_GetSelectionStartOffset", 0),
		/*11*/ imports.NewTable("CefDomDocument_GetTitle", 0),
		/*12*/ imports.NewTable("CefDomDocument_GetType", 0),
		/*13*/ imports.NewTable("CefDomDocument_HasSelection", 0),
		/*14*/ imports.NewTable("CefDomDocument_UnWrap", 0),
	}
)

func domDocumentImportAPI() *imports.Imports {
	if domDocumentImport == nil {
		domDocumentImport = NewDefaultImports()
		domDocumentImport.SetImportTable(domDocumentImportTables)
		domDocumentImportTables = nil
	}
	return domDocumentImport
}
