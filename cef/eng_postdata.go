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

// ICefPostData Parent: ICefBaseRefCounted
type ICefPostData interface {
	ICefBaseRefCounted
	GetElements(elementsCount *NativeUInt, elements *ICefPostDataElementArray)
	IsReadOnly() bool                               // function
	HasExcludedElements() bool                      // function
	GetElementCount() NativeUInt                    // function
	RemoveElement(element ICefPostDataElement) bool // function
	AddElement(element ICefPostDataElement) bool    // function
	RemoveElements()                                // procedure
}

// TCefPostData Parent: TCefBaseRefCounted
type TCefPostData struct {
	TCefBaseRefCounted
}

// PostDataRef -> ICefPostData
var PostDataRef postData

// postData TCefPostData Ref
type postData uintptr

func (m *postData) UnWrap(data uintptr) ICefPostData {
	var resultCefPostData uintptr
	postDataImportAPI().SysCallN(7, uintptr(data), uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *postData) New() ICefPostData {
	var resultCefPostData uintptr
	postDataImportAPI().SysCallN(4, uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *TCefPostData) IsReadOnly() bool {
	r1 := postDataImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostData) HasExcludedElements() bool {
	r1 := postDataImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCefPostData) GetElementCount() NativeUInt {
	r1 := postDataImportAPI().SysCallN(1, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPostData) RemoveElement(element ICefPostDataElement) bool {
	r1 := postDataImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(element))
	return GoBool(r1)
}

func (m *TCefPostData) AddElement(element ICefPostDataElement) bool {
	r1 := postDataImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(element))
	return GoBool(r1)
}

func (m *TCefPostData) RemoveElements() {
	postDataImportAPI().SysCallN(6, m.Instance())
}

var (
	postDataImport       *imports.Imports = nil
	postDataImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPostData_AddElement", 0),
		/*1*/ imports.NewTable("CefPostData_GetElementCount", 0),
		/*2*/ imports.NewTable("CefPostData_HasExcludedElements", 0),
		/*3*/ imports.NewTable("CefPostData_IsReadOnly", 0),
		/*4*/ imports.NewTable("CefPostData_New", 0),
		/*5*/ imports.NewTable("CefPostData_RemoveElement", 0),
		/*6*/ imports.NewTable("CefPostData_RemoveElements", 0),
		/*7*/ imports.NewTable("CefPostData_UnWrap", 0),
	}
)

func postDataImportAPI() *imports.Imports {
	if postDataImport == nil {
		postDataImport = NewDefaultImports()
		postDataImport.SetImportTable(postDataImportTables)
		postDataImportTables = nil
	}
	return postDataImport
}
