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

// ICefDomNode Parent: ICefBaseRefCounted
type ICefDomNode interface {
	ICefBaseRefCounted
	GetType() TCefDomNodeType                        // function
	IsText() bool                                    // function
	IsElement() bool                                 // function
	IsEditable() bool                                // function
	IsFormControlElement() bool                      // function
	GetFormControlElementType() string               // function
	IsSame(that ICefDomNode) bool                    // function
	GetName() string                                 // function
	GetValue() string                                // function
	SetValue(value string) bool                      // function
	GetAsMarkup() string                             // function
	GetDocument() ICefDomDocument                    // function
	GetParent() ICefDomNode                          // function
	GetPreviousSibling() ICefDomNode                 // function
	GetNextSibling() ICefDomNode                     // function
	HasChildren() bool                               // function
	GetFirstChild() ICefDomNode                      // function
	GetLastChild() ICefDomNode                       // function
	GetElementTagName() string                       // function
	HasElementAttributes() bool                      // function
	HasElementAttribute(attrName string) bool        // function
	GetElementAttribute(attrName string) string      // function
	SetElementAttribute(attrName, value string) bool // function
	GetElementInnerText() string                     // function
	GetElementBounds() (resultCefRect TCefRect)      // function
	GetElementAttributes(attrMap ICefStringMap)      // procedure
	GetElementAttributes1(attrList *IStrings)        // procedure
}

// TCefDomNode Parent: TCefBaseRefCounted
type TCefDomNode struct {
	TCefBaseRefCounted
}

// DomNodeRef -> ICefDomNode
var DomNodeRef domNode

// domNode TCefDomNode Ref
type domNode uintptr

func (m *domNode) UnWrap(data uintptr) ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(27, uintptr(data), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetType() TCefDomNodeType {
	r1 := domNodeImportAPI().SysCallN(15, m.Instance())
	return TCefDomNodeType(r1)
}

func (m *TCefDomNode) IsText() bool {
	r1 := domNodeImportAPI().SysCallN(24, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsElement() bool {
	r1 := domNodeImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsEditable() bool {
	r1 := domNodeImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsFormControlElement() bool {
	r1 := domNodeImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) GetFormControlElementType() string {
	r1 := domNodeImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) IsSame(that ICefDomNode) bool {
	r1 := domNodeImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDomNode) GetName() string {
	r1 := domNodeImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetValue() string {
	r1 := domNodeImportAPI().SysCallN(16, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) SetValue(value string) bool {
	r1 := domNodeImportAPI().SysCallN(26, m.Instance(), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDomNode) GetAsMarkup() string {
	r1 := domNodeImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetDocument() ICefDomDocument {
	var resultCefDomDocument uintptr
	domNodeImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCefDomDocument)))
	return AsCefDomDocument(resultCefDomDocument)
}

func (m *TCefDomNode) GetParent() ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetPreviousSibling() ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetNextSibling() ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) HasChildren() bool {
	r1 := domNodeImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) GetFirstChild() ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetLastChild() ICefDomNode {
	var resultCefDomNode uintptr
	domNodeImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetElementTagName() string {
	r1 := domNodeImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) HasElementAttributes() bool {
	r1 := domNodeImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) HasElementAttribute(attrName string) bool {
	r1 := domNodeImportAPI().SysCallN(18, m.Instance(), PascalStr(attrName))
	return GoBool(r1)
}

func (m *TCefDomNode) GetElementAttribute(attrName string) string {
	r1 := domNodeImportAPI().SysCallN(2, m.Instance(), PascalStr(attrName))
	return GoStr(r1)
}

func (m *TCefDomNode) SetElementAttribute(attrName, value string) bool {
	r1 := domNodeImportAPI().SysCallN(25, m.Instance(), PascalStr(attrName), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDomNode) GetElementInnerText() string {
	r1 := domNodeImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetElementBounds() (resultCefRect TCefRect) {
	domNodeImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDomNode) GetElementAttributes(attrMap ICefStringMap) {
	domNodeImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(attrMap))
}

func (m *TCefDomNode) GetElementAttributes1(attrList *IStrings) {
	var result0 uintptr
	domNodeImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&result0)))
	*attrList = AsStrings(result0)
}

var (
	domNodeImport       *imports.Imports = nil
	domNodeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDomNode_GetAsMarkup", 0),
		/*1*/ imports.NewTable("CefDomNode_GetDocument", 0),
		/*2*/ imports.NewTable("CefDomNode_GetElementAttribute", 0),
		/*3*/ imports.NewTable("CefDomNode_GetElementAttributes", 0),
		/*4*/ imports.NewTable("CefDomNode_GetElementAttributes1", 0),
		/*5*/ imports.NewTable("CefDomNode_GetElementBounds", 0),
		/*6*/ imports.NewTable("CefDomNode_GetElementInnerText", 0),
		/*7*/ imports.NewTable("CefDomNode_GetElementTagName", 0),
		/*8*/ imports.NewTable("CefDomNode_GetFirstChild", 0),
		/*9*/ imports.NewTable("CefDomNode_GetFormControlElementType", 0),
		/*10*/ imports.NewTable("CefDomNode_GetLastChild", 0),
		/*11*/ imports.NewTable("CefDomNode_GetName", 0),
		/*12*/ imports.NewTable("CefDomNode_GetNextSibling", 0),
		/*13*/ imports.NewTable("CefDomNode_GetParent", 0),
		/*14*/ imports.NewTable("CefDomNode_GetPreviousSibling", 0),
		/*15*/ imports.NewTable("CefDomNode_GetType", 0),
		/*16*/ imports.NewTable("CefDomNode_GetValue", 0),
		/*17*/ imports.NewTable("CefDomNode_HasChildren", 0),
		/*18*/ imports.NewTable("CefDomNode_HasElementAttribute", 0),
		/*19*/ imports.NewTable("CefDomNode_HasElementAttributes", 0),
		/*20*/ imports.NewTable("CefDomNode_IsEditable", 0),
		/*21*/ imports.NewTable("CefDomNode_IsElement", 0),
		/*22*/ imports.NewTable("CefDomNode_IsFormControlElement", 0),
		/*23*/ imports.NewTable("CefDomNode_IsSame", 0),
		/*24*/ imports.NewTable("CefDomNode_IsText", 0),
		/*25*/ imports.NewTable("CefDomNode_SetElementAttribute", 0),
		/*26*/ imports.NewTable("CefDomNode_SetValue", 0),
		/*27*/ imports.NewTable("CefDomNode_UnWrap", 0),
	}
)

func domNodeImportAPI() *imports.Imports {
	if domNodeImport == nil {
		domNodeImport = NewDefaultImports()
		domNodeImport.SetImportTable(domNodeImportTables)
		domNodeImportTables = nil
	}
	return domNodeImport
}
