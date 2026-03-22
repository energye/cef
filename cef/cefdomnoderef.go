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

// ICefDomNode Parent: ICefBaseRefCounted
type ICefDomNode interface {
	ICefBaseRefCounted
	// GetType
	//  Returns the type for this node.
	GetType() cefTypes.TCefDomNodeType // function
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
	GetFormControlElementType() cefTypes.TCefDomFormControlType // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that|
	//  object.
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
	//  Set the value for the element attribute named |attrName|. Returns true (1)
	//  on success.
	SetElementAttribute(attrName string, value string) bool // function
	// GetElementInnerText
	//  Returns the inner text of the element.
	GetElementInnerText() string // function
	// GetElementBounds
	//  Returns the bounds of the element in device pixels. Use
	//  "window.devicePixelRatio" to convert to/from CSS pixels.
	GetElementBounds() TCefRect // function
	// GetElementAttributes
	//  Returns a TStrings of all element attributes.
	GetElementAttributes(attrMap ICefStringMap) // procedure
}

// ICefDomNodeRef Parent: ICefDomNode ICefBaseRefCountedRef
type ICefDomNodeRef interface {
	ICefDomNode
	ICefBaseRefCountedRef
	AsIntfDomNode() uintptr
}

type TCefDomNodeRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDomNodeRef) GetType() cefTypes.TCefDomNodeType {
	if !m.IsValid() {
		return 0
	}
	r := cefDomNodeRefAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefDomNodeType(r)
}

func (m *TCefDomNodeRef) IsText() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) IsElement() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) IsEditable() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) IsFormControlElement() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetFormControlElementType() cefTypes.TCefDomFormControlType {
	if !m.IsValid() {
		return 0
	}
	r := cefDomNodeRefAPI().SysCallN(6, m.Instance())
	return cefTypes.TCefDomFormControlType(r)
}

func (m *TCefDomNodeRef) IsSame(that ICefDomNode) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) GetValue() string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) SetValue(value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(10, m.Instance(), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetAsMarkup() string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(11, m.Instance())
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) GetDocument() (result ICefDomDocument) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomDocumentRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) GetParent() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) GetPreviousSibling() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) GetNextSibling() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) HasChildren() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetFirstChild() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) GetLastChild() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomNodeRef) GetElementTagName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(19, m.Instance())
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) HasElementAttributes() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) HasElementAttribute(attrName string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(21, m.Instance(), api.PasStr(attrName))
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetElementAttribute(attrName string) string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(22, m.Instance(), api.PasStr(attrName))
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) SetElementAttribute(attrName string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomNodeRefAPI().SysCallN(23, m.Instance(), api.PasStr(attrName), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCefDomNodeRef) GetElementInnerText() string {
	if !m.IsValid() {
		return ""
	}
	r := cefDomNodeRefAPI().SysCallN(24, m.Instance())
	return api.GoStr(r)
}

func (m *TCefDomNodeRef) GetElementBounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	cefDomNodeRefAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDomNodeRef) GetElementAttributes(attrMap ICefStringMap) {
	if !m.IsValid() {
		return
	}
	cefDomNodeRefAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(attrMap))
}

func (m *TCefDomNodeRef) AsIntfDomNode() uintptr {
	return m.GetIntfPointer(0)
}

// DomNodeRef  is static instance
var DomNodeRef _DomNodeRefClass

// _DomNodeRefClass is class type defined by TCefDomNodeRef
type _DomNodeRefClass uintptr

func (_DomNodeRefClass) UnWrap(data uintptr) (result ICefDomNode) {
	var resultPtr uintptr
	cefDomNodeRefAPI().SysCallN(26, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

// NewDomNodeRef class constructor
func NewDomNodeRef(data uintptr) ICefDomNodeRef {
	var domNodePtr uintptr // ICefDomNode
	r := cefDomNodeRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&domNodePtr)))
	ret := AsCefDomNodeRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, domNodePtr)
	}
	return ret
}

var (
	cefDomNodeRefOnce   base.Once
	cefDomNodeRefImport *imports.Imports = nil
)

func cefDomNodeRefAPI() *imports.Imports {
	cefDomNodeRefOnce.Do(func() {
		cefDomNodeRefImport = api.NewDefaultImports()
		cefDomNodeRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDomNodeRef_Create", 0), // constructor NewDomNodeRef
			/* 1 */ imports.NewTable("TCefDomNodeRef_GetType", 0), // function GetType
			/* 2 */ imports.NewTable("TCefDomNodeRef_IsText", 0), // function IsText
			/* 3 */ imports.NewTable("TCefDomNodeRef_IsElement", 0), // function IsElement
			/* 4 */ imports.NewTable("TCefDomNodeRef_IsEditable", 0), // function IsEditable
			/* 5 */ imports.NewTable("TCefDomNodeRef_IsFormControlElement", 0), // function IsFormControlElement
			/* 6 */ imports.NewTable("TCefDomNodeRef_GetFormControlElementType", 0), // function GetFormControlElementType
			/* 7 */ imports.NewTable("TCefDomNodeRef_IsSame", 0), // function IsSame
			/* 8 */ imports.NewTable("TCefDomNodeRef_GetName", 0), // function GetName
			/* 9 */ imports.NewTable("TCefDomNodeRef_GetValue", 0), // function GetValue
			/* 10 */ imports.NewTable("TCefDomNodeRef_SetValue", 0), // function SetValue
			/* 11 */ imports.NewTable("TCefDomNodeRef_GetAsMarkup", 0), // function GetAsMarkup
			/* 12 */ imports.NewTable("TCefDomNodeRef_GetDocument", 0), // function GetDocument
			/* 13 */ imports.NewTable("TCefDomNodeRef_GetParent", 0), // function GetParent
			/* 14 */ imports.NewTable("TCefDomNodeRef_GetPreviousSibling", 0), // function GetPreviousSibling
			/* 15 */ imports.NewTable("TCefDomNodeRef_GetNextSibling", 0), // function GetNextSibling
			/* 16 */ imports.NewTable("TCefDomNodeRef_HasChildren", 0), // function HasChildren
			/* 17 */ imports.NewTable("TCefDomNodeRef_GetFirstChild", 0), // function GetFirstChild
			/* 18 */ imports.NewTable("TCefDomNodeRef_GetLastChild", 0), // function GetLastChild
			/* 19 */ imports.NewTable("TCefDomNodeRef_GetElementTagName", 0), // function GetElementTagName
			/* 20 */ imports.NewTable("TCefDomNodeRef_HasElementAttributes", 0), // function HasElementAttributes
			/* 21 */ imports.NewTable("TCefDomNodeRef_HasElementAttribute", 0), // function HasElementAttribute
			/* 22 */ imports.NewTable("TCefDomNodeRef_GetElementAttribute", 0), // function GetElementAttribute
			/* 23 */ imports.NewTable("TCefDomNodeRef_SetElementAttribute", 0), // function SetElementAttribute
			/* 24 */ imports.NewTable("TCefDomNodeRef_GetElementInnerText", 0), // function GetElementInnerText
			/* 25 */ imports.NewTable("TCefDomNodeRef_GetElementBounds", 0), // function GetElementBounds
			/* 26 */ imports.NewTable("TCefDomNodeRef_UnWrap", 0), // static function UnWrap
			/* 27 */ imports.NewTable("TCefDomNodeRef_GetElementAttributes", 0), // procedure GetElementAttributes
		}
	})
	return cefDomNodeRefImport
}
