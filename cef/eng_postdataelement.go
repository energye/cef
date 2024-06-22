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

// ICefPostDataElement Parent: ICefBaseRefCounted
type ICefPostDataElement interface {
	ICefBaseRefCounted
	IsReadOnly() bool                                   // function
	GetType() TCefPostDataElementType                   // function
	GetFile() string                                    // function
	GetBytesCount() NativeUInt                          // function
	GetBytes(size NativeUInt, bytes uintptr) NativeUInt // function
	SetToEmpty()                                        // procedure
	SetToFile(fileName string)                          // procedure
	SetToBytes(size NativeUInt, bytes uintptr)          // procedure
}

// TCefPostDataElement Parent: TCefBaseRefCounted
type TCefPostDataElement struct {
	TCefBaseRefCounted
}

// PostDataElementRef -> ICefPostDataElement
var PostDataElementRef postDataElement

// postDataElement TCefPostDataElement Ref
type postDataElement uintptr

func (m *postDataElement) UnWrap(data uintptr) ICefPostDataElement {
	var resultCefPostDataElement uintptr
	postDataElementImportAPI().SysCallN(9, uintptr(data), uintptr(unsafePointer(&resultCefPostDataElement)))
	return AsCefPostDataElement(resultCefPostDataElement)
}

func (m *postDataElement) New() ICefPostDataElement {
	var resultCefPostDataElement uintptr
	postDataElementImportAPI().SysCallN(5, uintptr(unsafePointer(&resultCefPostDataElement)))
	return AsCefPostDataElement(resultCefPostDataElement)
}

func (m *TCefPostDataElement) IsReadOnly() bool {
	r1 := postDataElementImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostDataElement) GetType() TCefPostDataElementType {
	r1 := postDataElementImportAPI().SysCallN(3, m.Instance())
	return TCefPostDataElementType(r1)
}

func (m *TCefPostDataElement) GetFile() string {
	r1 := postDataElementImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefPostDataElement) GetBytesCount() NativeUInt {
	r1 := postDataElementImportAPI().SysCallN(1, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPostDataElement) GetBytes(size NativeUInt, bytes uintptr) NativeUInt {
	r1 := postDataElementImportAPI().SysCallN(0, m.Instance(), uintptr(size), uintptr(bytes))
	return NativeUInt(r1)
}

func (m *TCefPostDataElement) SetToEmpty() {
	postDataElementImportAPI().SysCallN(7, m.Instance())
}

func (m *TCefPostDataElement) SetToFile(fileName string) {
	postDataElementImportAPI().SysCallN(8, m.Instance(), PascalStr(fileName))
}

func (m *TCefPostDataElement) SetToBytes(size NativeUInt, bytes uintptr) {
	postDataElementImportAPI().SysCallN(6, m.Instance(), uintptr(size), uintptr(bytes))
}

var (
	postDataElementImport       *imports.Imports = nil
	postDataElementImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPostDataElement_GetBytes", 0),
		/*1*/ imports.NewTable("CefPostDataElement_GetBytesCount", 0),
		/*2*/ imports.NewTable("CefPostDataElement_GetFile", 0),
		/*3*/ imports.NewTable("CefPostDataElement_GetType", 0),
		/*4*/ imports.NewTable("CefPostDataElement_IsReadOnly", 0),
		/*5*/ imports.NewTable("CefPostDataElement_New", 0),
		/*6*/ imports.NewTable("CefPostDataElement_SetToBytes", 0),
		/*7*/ imports.NewTable("CefPostDataElement_SetToEmpty", 0),
		/*8*/ imports.NewTable("CefPostDataElement_SetToFile", 0),
		/*9*/ imports.NewTable("CefPostDataElement_UnWrap", 0),
	}
)

func postDataElementImportAPI() *imports.Imports {
	if postDataElementImport == nil {
		postDataElementImport = NewDefaultImports()
		postDataElementImport.SetImportTable(postDataElementImportTables)
		postDataElementImportTables = nil
	}
	return postDataElementImport
}
