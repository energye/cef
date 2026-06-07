//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/types"
)

// ICefNavigationEntry Parent: ICefBaseRefCounted
type ICefNavigationEntry interface {
	ICefBaseRefCounted
	IsValid() bool                                  // function
	GetUrl() string                                 // function
	GetDisplayUrl() string                          // function
	GetOriginalUrl() string                         // function
	GetTitle() string                               // function
	GetTransitionType() cefTypes.TCefTransitionType // function
	HasPostData() bool                              // function
	GetCompletionTime() types.TDateTime             // function
	GetHttpStatusCode() int32                       // function
	GetSSLStatus() ICefSSLStatus                    // function
}

// ICefNavigationEntryRef Parent: ICefNavigationEntry ICefBaseRefCountedRef
type ICefNavigationEntryRef interface {
	ICefNavigationEntry
	ICefBaseRefCountedRef
	AsIntfNavigationEntry() uintptr
}

type TCefNavigationEntryRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefNavigationEntryRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefNavigationEntryRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefNavigationEntryRef) GetUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefNavigationEntryRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefNavigationEntryRef) GetDisplayUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefNavigationEntryRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefNavigationEntryRef) GetOriginalUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefNavigationEntryRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefNavigationEntryRef) GetTitle() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefNavigationEntryRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefNavigationEntryRef) GetTransitionType() cefTypes.TCefTransitionType {
	if !m.IsValid() {
		return 0
	}
	r := cefNavigationEntryRefAPI().SysCallN(6, m.Instance())
	return cefTypes.TCefTransitionType(r)
}

func (m *TCefNavigationEntryRef) HasPostData() bool {
	if !m.IsValid() {
		return false
	}
	r := cefNavigationEntryRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefNavigationEntryRef) GetCompletionTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefNavigationEntryRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefNavigationEntryRef) GetHttpStatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefNavigationEntryRefAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TCefNavigationEntryRef) GetSSLStatus() (result ICefSSLStatus) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefNavigationEntryRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefSSLStatusRef(resultPtr)
	return
}

func (m *TCefNavigationEntryRef) AsIntfNavigationEntry() uintptr {
	return m.GetIntfPointer(0)
}

// NavigationEntryRef  is static instance
var NavigationEntryRef _NavigationEntryRefClass

// _NavigationEntryRefClass is class type defined by TCefNavigationEntryRef
type _NavigationEntryRefClass uintptr

func (_NavigationEntryRefClass) UnWrap(data uintptr) (result ICefNavigationEntry) {
	var resultPtr uintptr
	cefNavigationEntryRefAPI().SysCallN(11, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefNavigationEntryRef(resultPtr)
	return
}

// NewNavigationEntryRef class constructor
func NewNavigationEntryRef(data uintptr) ICefNavigationEntryRef {
	var navigationEntryPtr uintptr // ICefNavigationEntry
	r := cefNavigationEntryRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&navigationEntryPtr)))
	ret := AsCefNavigationEntryRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, navigationEntryPtr)
	}
	return ret
}

var (
	cefNavigationEntryRefOnce   base.Once
	cefNavigationEntryRefImport *imports.Imports = nil
)

func cefNavigationEntryRefAPI() *imports.Imports {
	cefNavigationEntryRefOnce.Do(func() {
		cefNavigationEntryRefImport = api.NewDefaultImports()
		cefNavigationEntryRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefNavigationEntryRef_Create", 0), // constructor NewNavigationEntryRef
			/* 1 */ imports.NewTable("TCefNavigationEntryRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefNavigationEntryRef_GetUrl", 0), // function GetUrl
			/* 3 */ imports.NewTable("TCefNavigationEntryRef_GetDisplayUrl", 0), // function GetDisplayUrl
			/* 4 */ imports.NewTable("TCefNavigationEntryRef_GetOriginalUrl", 0), // function GetOriginalUrl
			/* 5 */ imports.NewTable("TCefNavigationEntryRef_GetTitle", 0), // function GetTitle
			/* 6 */ imports.NewTable("TCefNavigationEntryRef_GetTransitionType", 0), // function GetTransitionType
			/* 7 */ imports.NewTable("TCefNavigationEntryRef_HasPostData", 0), // function HasPostData
			/* 8 */ imports.NewTable("TCefNavigationEntryRef_GetCompletionTime", 0), // function GetCompletionTime
			/* 9 */ imports.NewTable("TCefNavigationEntryRef_GetHttpStatusCode", 0), // function GetHttpStatusCode
			/* 10 */ imports.NewTable("TCefNavigationEntryRef_GetSSLStatus", 0), // function GetSSLStatus
			/* 11 */ imports.NewTable("TCefNavigationEntryRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefNavigationEntryRefImport
}
