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
//
//	Interface used to represent a DOM node. The functions of this interface should only be called on the render process main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domnode_t))</a>
type ICefDomNode interface {
	ICefBaseRefCounted
	// GetType
	//  Returns the type for this node.
	GetType() TCefDomNodeType // function
	// IsText
	//  Returns true (1) if this is a text node.
	IsText() bool // function
	// IsElement
	//  Returns true (1) if this is an element node.
	IsElement() bool // function
	// IsEditable
	//  Returns true (1) if this is an editable node.
	IsEditable() bool // function
	// IsFormControlElement
	//  Returns true (1) if this is a form control element node.
	IsFormControlElement() bool // function
	// GetFormControlElementType
	//  Returns the type of this form control element node.
	GetFormControlElementType() string // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefDomNode) bool // function
	// GetName
	//  Returns the name of this node.
	GetName() string // function
	// GetValue
	//  Returns the value of this node.
	GetValue() string // function
	// SetValue
	//  Set the value of this node. Returns true (1) on success.
	SetValue(value string) bool // function
	// GetAsMarkup
	//  Returns the contents of this node as markup.
	GetAsMarkup() string // function
	// GetDocument
	//  Returns the document associated with this node.
	GetDocument() ICefDomDocument // function
	// GetParent
	//  Returns the parent node.
	GetParent() ICefDomNode // function
	// GetPreviousSibling
	//  Returns the previous sibling node.
	GetPreviousSibling() ICefDomNode // function
	// GetNextSibling
	//  Returns the next sibling node.
	GetNextSibling() ICefDomNode // function
	// HasChildren
	//  Returns true (1) if this node has child nodes.
	HasChildren() bool // function
	// GetFirstChild
	//  Return the first child node.
	GetFirstChild() ICefDomNode // function
	// GetLastChild
	//  Returns the last child node.
	GetLastChild() ICefDomNode // function
	// GetElementTagName
	//  Returns the tag name of this element.
	GetElementTagName() string // function
	// HasElementAttributes
	//  Returns true (1) if this element has attributes.
	HasElementAttributes() bool // function
	// HasElementAttribute
	//  Returns true (1) if this element has an attribute named |attrName|.
	HasElementAttribute(attrName string) bool // function
	// GetElementAttribute
	//  Returns the element attribute named |attrName|.
	GetElementAttribute(attrName string) string // function
	// SetElementAttribute
	//  Set the value for the element attribute named |attrName|. Returns true (1) on success.
	SetElementAttribute(attrName, value string) bool // function
	// GetElementInnerText
	//  Returns the inner text of the element.
	GetElementInnerText() string // function
	// GetElementBounds
	//  Returns the bounds of the element in device pixels. Use "window.devicePixelRatio" to convert to/from CSS pixels.
	GetElementBounds() (resultCefRect TCefRect) // function
	// GetElementAttributes
	//  Returns a ICefStringMap of all element attributes.
	GetElementAttributes(attrMap ICefStringMap) // procedure
	// GetElementAttributes1
	//  Returns a ICefStringMap of all element attributes.
	GetElementAttributes1(attrList *IStrings) // procedure
}

// TCefDomNode Parent: TCefBaseRefCounted
//
//	Interface used to represent a DOM node. The functions of this interface should only be called on the render process main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domnode_t))</a>
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
