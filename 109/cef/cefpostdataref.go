//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefPostData Parent: ICefBaseRefCounted
type ICefPostData interface {
	ICefBaseRefCounted
	IsReadOnly() bool                                                                  // function
	HasExcludedElements() bool                                                         // function
	GetElementCount() cefTypes.NativeUInt                                              // function
	RemoveElement(element ICefPostDataElement) bool                                    // function
	AddElement(element ICefPostDataElement) bool                                       // function
	GetElements(elementsCount cefTypes.NativeUInt, elements *ICefPostDataElementArray) // procedure
	RemoveElements()                                                                   // procedure
}

// ICefPostDataRef Parent: ICefPostData ICefBaseRefCountedRef
type ICefPostDataRef interface {
	ICefPostData
	ICefBaseRefCountedRef
	AsIntfPostData() uintptr
}

type TCefPostDataRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPostDataRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPostDataRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPostDataRef) HasExcludedElements() bool {
	if !m.IsValid() {
		return false
	}
	r := cefPostDataRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefPostDataRef) GetElementCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefPostDataRefAPI().SysCallN(3, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefPostDataRef) RemoveElement(element ICefPostDataElement) bool {
	if !m.IsValid() {
		return false
	}
	r := cefPostDataRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(element))
	return api.GoBool(r)
}

func (m *TCefPostDataRef) AddElement(element ICefPostDataElement) bool {
	if !m.IsValid() {
		return false
	}
	r := cefPostDataRefAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(element))
	return api.GoBool(r)
}

func (m *TCefPostDataRef) GetElements(elementsCount cefTypes.NativeUInt, elements *ICefPostDataElementArray) {
	if !m.IsValid() {
		return
	}
	var elementsPtr uintptr
	cefPostDataRefAPI().SysCallN(8, m.Instance(), uintptr(elementsCount), uintptr(base.UnsafePointer(&elementsPtr)))
	*elements = NewCefPostDataElementArray(int(elementsCount), elementsPtr)
}

func (m *TCefPostDataRef) RemoveElements() {
	if !m.IsValid() {
		return
	}
	cefPostDataRefAPI().SysCallN(9, m.Instance())
}

func (m *TCefPostDataRef) AsIntfPostData() uintptr {
	return m.GetIntfPointer(0)
}

// PostDataRef  is static instance
var PostDataRef _PostDataRefClass

// _PostDataRefClass is class type defined by TCefPostDataRef
type _PostDataRefClass uintptr

func (_PostDataRefClass) UnWrap(data uintptr) (result ICefPostData) {
	var resultPtr uintptr
	cefPostDataRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPostDataRef(resultPtr)
	return
}

func (_PostDataRefClass) New() (result ICefPostData) {
	var resultPtr uintptr
	cefPostDataRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPostDataRef(resultPtr)
	return
}

// NewPostDataRef class constructor
func NewPostDataRef(data uintptr) ICefPostDataRef {
	var postDataPtr uintptr // ICefPostData
	r := cefPostDataRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&postDataPtr)))
	ret := AsCefPostDataRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, postDataPtr)
	}
	return ret
}

var (
	cefPostDataRefOnce   base.Once
	cefPostDataRefImport *imports.Imports = nil
)

func cefPostDataRefAPI() *imports.Imports {
	cefPostDataRefOnce.Do(func() {
		cefPostDataRefImport = api.NewDefaultImports()
		cefPostDataRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPostDataRef_Create", 0), // constructor NewPostDataRef
			/* 1 */ imports.NewTable("TCefPostDataRef_IsReadOnly", 0), // function IsReadOnly
			/* 2 */ imports.NewTable("TCefPostDataRef_HasExcludedElements", 0), // function HasExcludedElements
			/* 3 */ imports.NewTable("TCefPostDataRef_GetElementCount", 0), // function GetElementCount
			/* 4 */ imports.NewTable("TCefPostDataRef_RemoveElement", 0), // function RemoveElement
			/* 5 */ imports.NewTable("TCefPostDataRef_AddElement", 0), // function AddElement
			/* 6 */ imports.NewTable("TCefPostDataRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefPostDataRef_New", 0), // static function New
			/* 8 */ imports.NewTable("TCefPostDataRef_GetElements", 0), // procedure GetElements
			/* 9 */ imports.NewTable("TCefPostDataRef_RemoveElements", 0), // procedure RemoveElements
		}
	})
	return cefPostDataRefImport
}
