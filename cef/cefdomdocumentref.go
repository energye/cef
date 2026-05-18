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

// ICefDomDocument Parent: ICefBaseRefCounted
type ICefDomDocument interface {
	ICefBaseRefCounted
	// GetType
	//  Returns the document type.
	GetType() cefTypes.TCefDomDocumentType // function
	// GetDocument
	//  Returns the root document node.
	GetDocument() ICefDomNode // function
	// GetBody
	//  Returns the BODY node of an HTML document.
	GetBody() ICefDomNode // function
	// GetHead
	//  Returns the HEAD node of an HTML document.
	GetHead() ICefDomNode // function
	// GetTitle
	//  Returns the title of an HTML document.
	GetTitle() string // function
	// GetElementById
	//  Returns the document element with the specified ID value.
	GetElementById(id string) ICefDomNode // function
	// GetFocusedNode
	//  Returns the node that currently has keyboard focus.
	GetFocusedNode() ICefDomNode // function
	// HasSelection
	//  Returns true (1) if a portion of the document is selected.
	HasSelection() bool // function
	// GetSelectionStartOffset
	//  Returns the selection offset within the start node.
	GetSelectionStartOffset() int32 // function
	// GetSelectionEndOffset
	//  Returns the selection offset within the end node.
	GetSelectionEndOffset() int32 // function
	// GetSelectionAsMarkup
	//  Returns the contents of this selection as markup.
	GetSelectionAsMarkup() string // function
	// GetSelectionAsText
	//  Returns the contents of this selection as text.
	GetSelectionAsText() string // function
	// GetBaseUrl
	//  Returns the base URL for the document.
	GetBaseUrl() string // function
	// GetCompleteUrl
	//  Returns a complete URL based on the document base URL and the specified
	//  partial URL.
	GetCompleteUrl(partialURL string) string // function
}

// ICefDomDocumentRef Parent: ICefDomDocument ICefBaseRefCountedRef
type ICefDomDocumentRef interface {
	ICefDomDocument
	ICefBaseRefCountedRef
	AsIntfDomDocument() uintptr
}

type TCefDomDocumentRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDomDocumentRef) GetType() cefTypes.TCefDomDocumentType {
	if !m.IsValid() {
		return 0
	}
	r := cefDomDocumentRefAPI().SysCallN(1, m.Instance())
	return cefTypes.TCefDomDocumentType(r)
}

func (m *TCefDomDocumentRef) GetDocument() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomDocumentRef) GetBody() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomDocumentRef) GetHead() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomDocumentRef) GetTitle() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDomDocumentRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDomDocumentRef) GetElementById(id string) (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(6, m.Instance(), api.PasStr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomDocumentRef) GetFocusedNode() (result ICefDomNode) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomNodeRef(resultPtr)
	return
}

func (m *TCefDomDocumentRef) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDomDocumentRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDomDocumentRef) GetSelectionStartOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDomDocumentRefAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TCefDomDocumentRef) GetSelectionEndOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDomDocumentRefAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TCefDomDocumentRef) GetSelectionAsMarkup() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDomDocumentRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDomDocumentRef) GetSelectionAsText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDomDocumentRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDomDocumentRef) GetBaseUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDomDocumentRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDomDocumentRef) GetCompleteUrl(partialURL string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDomDocumentRefAPI().SysCallN(14, m.Instance(), api.PasStr(partialURL), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDomDocumentRef) AsIntfDomDocument() uintptr {
	return m.GetIntfPointer(0)
}

// DomDocumentRef  is static instance
var DomDocumentRef _DomDocumentRefClass

// _DomDocumentRefClass is class type defined by TCefDomDocumentRef
type _DomDocumentRefClass uintptr

func (_DomDocumentRefClass) UnWrap(data uintptr) (result ICefDomDocument) {
	var resultPtr uintptr
	cefDomDocumentRefAPI().SysCallN(15, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDomDocumentRef(resultPtr)
	return
}

// NewDomDocumentRef class constructor
func NewDomDocumentRef(data uintptr) ICefDomDocumentRef {
	var domDocumentPtr uintptr // ICefDomDocument
	r := cefDomDocumentRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&domDocumentPtr)))
	ret := AsCefDomDocumentRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, domDocumentPtr)
	}
	return ret
}

var (
	cefDomDocumentRefOnce   base.Once
	cefDomDocumentRefImport *imports.Imports = nil
)

func cefDomDocumentRefAPI() *imports.Imports {
	cefDomDocumentRefOnce.Do(func() {
		cefDomDocumentRefImport = api.NewDefaultImports()
		cefDomDocumentRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDomDocumentRef_Create", 0), // constructor NewDomDocumentRef
			/* 1 */ imports.NewTable("TCefDomDocumentRef_GetType", 0), // function GetType
			/* 2 */ imports.NewTable("TCefDomDocumentRef_GetDocument", 0), // function GetDocument
			/* 3 */ imports.NewTable("TCefDomDocumentRef_GetBody", 0), // function GetBody
			/* 4 */ imports.NewTable("TCefDomDocumentRef_GetHead", 0), // function GetHead
			/* 5 */ imports.NewTable("TCefDomDocumentRef_GetTitle", 0), // function GetTitle
			/* 6 */ imports.NewTable("TCefDomDocumentRef_GetElementById", 0), // function GetElementById
			/* 7 */ imports.NewTable("TCefDomDocumentRef_GetFocusedNode", 0), // function GetFocusedNode
			/* 8 */ imports.NewTable("TCefDomDocumentRef_HasSelection", 0), // function HasSelection
			/* 9 */ imports.NewTable("TCefDomDocumentRef_GetSelectionStartOffset", 0), // function GetSelectionStartOffset
			/* 10 */ imports.NewTable("TCefDomDocumentRef_GetSelectionEndOffset", 0), // function GetSelectionEndOffset
			/* 11 */ imports.NewTable("TCefDomDocumentRef_GetSelectionAsMarkup", 0), // function GetSelectionAsMarkup
			/* 12 */ imports.NewTable("TCefDomDocumentRef_GetSelectionAsText", 0), // function GetSelectionAsText
			/* 13 */ imports.NewTable("TCefDomDocumentRef_GetBaseUrl", 0), // function GetBaseUrl
			/* 14 */ imports.NewTable("TCefDomDocumentRef_GetCompleteUrl", 0), // function GetCompleteUrl
			/* 15 */ imports.NewTable("TCefDomDocumentRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefDomDocumentRefImport
}
