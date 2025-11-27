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

// ICefNavigationEntry Parent: ICefBaseRefCountedRef
type ICefNavigationEntry interface {
	ICefBaseRefCountedRef
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions
	//  if this function returns false (0).
	IsValid() bool // function
	// GetUrl
	//  Returns the actual URL of the page. For some pages this may be data: URL
	//  or similar. Use get_display_url() to return a display-friendly version.
	GetUrl() string // function
	// GetDisplayUrl
	//  Returns a display-friendly version of the URL.
	GetDisplayUrl() string // function
	// GetOriginalUrl
	//  Returns the original URL that was entered by the user before any
	//  redirects.
	GetOriginalUrl() string // function
	// GetTitle
	//  Returns the title set by the page. This value may be NULL.
	GetTitle() string // function
	// GetTransitionType
	//  Returns the transition type which indicates what the user did to move to
	//  this page from the previous page.
	GetTransitionType() cefTypes.TCefTransitionType // function
	// HasPostData
	//  Returns true (1) if this navigation includes post data.
	HasPostData() bool // function
	// GetCompletionTime
	//  Returns the time for the last known successful navigation completion. A
	//  navigation may be completed more than once if the page is reloaded. May be
	//  0 if the navigation has not yet completed.
	GetCompletionTime() types.TDateTime // function
	// GetHttpStatusCode
	//  Returns the HTTP status code for the last known successful navigation
	//  response. May be 0 if the response has not yet been received or if the
	//  navigation has not yet completed.
	GetHttpStatusCode() int32 // function
	// GetSSLStatus
	//  Returns the SSL information for this navigation entry.
	GetSSLStatus() ICefSSLStatus // function
}

// ICefNavigationEntryRef Parent: ICefNavigationEntry
type ICefNavigationEntryRef interface {
	ICefNavigationEntry
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

func (m *TCefNavigationEntryRef) GetUrl() string {
	if !m.IsValid() {
		return ""
	}
	r := cefNavigationEntryRefAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TCefNavigationEntryRef) GetDisplayUrl() string {
	if !m.IsValid() {
		return ""
	}
	r := cefNavigationEntryRefAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCefNavigationEntryRef) GetOriginalUrl() string {
	if !m.IsValid() {
		return ""
	}
	r := cefNavigationEntryRefAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCefNavigationEntryRef) GetTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := cefNavigationEntryRefAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
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
