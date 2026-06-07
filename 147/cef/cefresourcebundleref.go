//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/147/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefResourceBundle Parent: ICefBaseRefCounted
type ICefResourceBundle interface {
	ICefBaseRefCounted
	// GetLocalizedString
	//  Returns the localized string for the specified |string_id| or an NULL
	//  string if the value is not found. Use the cef_id_for_pack_string_name()
	//  function for version-safe mapping of string IDS names from
	//  cef_pack_strings.h to version-specific numerical |string_id| values.
	GetLocalizedString(stringId int32) string // function
	// GetDataResource
	//  Returns a cef_binary_value_t containing the decompressed contents of the
	//  specified scale independent |resource_id| or NULL if not found. Use the
	//  cef_id_for_pack_resource_name() function for version-safe mapping of
	//  resource IDR names from cef_pack_resources.h to version-specific numerical
	//  |resource_id| values.
	GetDataResource(resourceId int32) ICefBinaryValue // function
	// GetDataResourceForScale
	//  Returns a cef_binary_value_t containing the decompressed contents of the
	//  specified |resource_id| nearest the scale factor |scale_factor| or NULL if
	//  not found. Use a |scale_factor| value of SCALE_FACTOR_NONE for scale
	//  independent resources or call GetDataResource instead. Use the
	//  cef_id_for_pack_resource_name() function for version-safe mapping of
	//  resource IDR names from cef_pack_resources.h to version-specific numerical
	//  |resource_id| values.
	GetDataResourceForScale(resourceId int32, scaleFactor cefTypes.TCefScaleFactor) ICefBinaryValue // function
}

// ICefResourceBundleRef Parent: ICefResourceBundle ICefBaseRefCountedRef
type ICefResourceBundleRef interface {
	ICefResourceBundle
	ICefBaseRefCountedRef
	AsIntfResourceBundle() uintptr
}

type TCefResourceBundleRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResourceBundleRef) GetLocalizedString(stringId int32) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResourceBundleRefAPI().SysCallN(1, m.Instance(), uintptr(stringId), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResourceBundleRef) GetDataResource(resourceId int32) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefResourceBundleRefAPI().SysCallN(2, m.Instance(), uintptr(resourceId), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefResourceBundleRef) GetDataResourceForScale(resourceId int32, scaleFactor cefTypes.TCefScaleFactor) (result ICefBinaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefResourceBundleRefAPI().SysCallN(3, m.Instance(), uintptr(resourceId), uintptr(scaleFactor), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBinaryValueRef(resultPtr)
	return
}

func (m *TCefResourceBundleRef) AsIntfResourceBundle() uintptr {
	return m.GetIntfPointer(0)
}

// ResourceBundleRef  is static instance
var ResourceBundleRef _ResourceBundleRefClass

// _ResourceBundleRefClass is class type defined by TCefResourceBundleRef
type _ResourceBundleRefClass uintptr

func (_ResourceBundleRefClass) UnWrap(data uintptr) (result ICefResourceBundle) {
	var resultPtr uintptr
	cefResourceBundleRefAPI().SysCallN(4, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResourceBundleRef(resultPtr)
	return
}

func (_ResourceBundleRefClass) Global() (result ICefResourceBundle) {
	var resultPtr uintptr
	cefResourceBundleRefAPI().SysCallN(5, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResourceBundleRef(resultPtr)
	return
}

// NewResourceBundleRef class constructor
func NewResourceBundleRef(data uintptr) ICefResourceBundleRef {
	var resourceBundlePtr uintptr // ICefResourceBundle
	r := cefResourceBundleRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&resourceBundlePtr)))
	ret := AsCefResourceBundleRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, resourceBundlePtr)
	}
	return ret
}

var (
	cefResourceBundleRefOnce   base.Once
	cefResourceBundleRefImport *imports.Imports = nil
)

func cefResourceBundleRefAPI() *imports.Imports {
	cefResourceBundleRefOnce.Do(func() {
		cefResourceBundleRefImport = api.NewDefaultImports()
		cefResourceBundleRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResourceBundleRef_Create", 0), // constructor NewResourceBundleRef
			/* 1 */ imports.NewTable("TCefResourceBundleRef_GetLocalizedString", 0), // function GetLocalizedString
			/* 2 */ imports.NewTable("TCefResourceBundleRef_GetDataResource", 0), // function GetDataResource
			/* 3 */ imports.NewTable("TCefResourceBundleRef_GetDataResourceForScale", 0), // function GetDataResourceForScale
			/* 4 */ imports.NewTable("TCefResourceBundleRef_UnWrap", 0), // static function UnWrap
			/* 5 */ imports.NewTable("TCefResourceBundleRef_Global", 0), // static function Global
		}
	})
	return cefResourceBundleRefImport
}
