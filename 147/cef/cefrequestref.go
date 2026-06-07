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

// ICefRequest Parent: ICefBaseRefCounted
type ICefRequest interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetUrl
	//  Get the fully qualified URL.
	GetUrl() string // function
	// GetMethod
	//  Get the request function type. The value will default to POST if post data
	//  is provided and GET otherwise.
	GetMethod() string // function
	// GetPostData
	//  Get the post data.
	GetPostData() ICefPostData // function
	// GetReferrerUrl
	//  Get the referrer URL.
	GetReferrerUrl() string // function
	// GetReferrerPolicy
	//  Get the referrer policy.
	GetReferrerPolicy() cefTypes.TCefReferrerPolicy // function
	// GetHeaderByName
	//  Returns the first header value for |name| or an NULL string if not found.
	//  Will not return the Referer value if any. Use GetHeaderMap instead if
	//  |name| might have multiple values.
	GetHeaderByName(name string) string // function
	// GetFlags
	//  Get the flags used in combination with ICefUrlRequest. See
	//  TCefUrlRequestFlags for supported values.
	GetFlags() cefTypes.TCefUrlRequestFlags // function
	// GetFirstPartyForCookies
	//  Get the URL to the first party for cookies used in combination with
	//  ICefUrlRequest.
	GetFirstPartyForCookies() string // function
	// GetResourceType
	//  Get the resource type for this request. Only available in the browser
	//  process.
	GetResourceType() cefTypes.TCefResourceType // function
	// GetTransitionType
	//  Get the transition type for this request. Only available in the browser
	//  process and only applies to requests that represent a main frame or sub-
	//  frame navigation.
	GetTransitionType() cefTypes.TCefTransitionType // function
	// GetIdentifier
	//  Returns the globally unique identifier for this request or 0 if not
	//  specified. Can be used by ICefResourceRequestHandler implementations
	//  in the browser process to track a single request across multiple
	//  callbacks.
	GetIdentifier() uint64 // function
	// GetHeaderMap
	//  Get the header values. Will not include the Referer value if any.
	GetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetUrl
	//  Set the fully qualified URL.
	SetUrl(value string) // procedure
	// SetMethod
	//  Set the request function type.
	SetMethod(value string) // procedure
	// SetReferrer
	//  Set the referrer URL and policy. If non-NULL the referrer URL must be
	//  fully qualified with an HTTP or HTTPS scheme component. Any username,
	//  password or ref component will be removed.
	SetReferrer(referrerUrl string, policy cefTypes.TCefReferrerPolicy) // procedure
	// SetPostData
	//  Set the post data.
	SetPostData(value ICefPostData) // procedure
	// SetHeaderMap
	//  Set the header values. If a Referer value exists in the header map it will
	//  be removed and ignored.
	SetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetHeaderByName
	//  Set the header |name| to |value|. If |overwrite| is true (1) any existing
	//  values will be replaced with the new value. If |overwrite| is false (0)
	//  any existing values will not be overwritten. The Referer value cannot be
	//  set using this function.
	SetHeaderByName(name string, value string, overwrite bool) // procedure
	// SetFlags
	//  Set the flags used in combination with ICefUrlRequest. See
	//  TCefUrlRequestFlags for supported values.
	SetFlags(flags cefTypes.TCefUrlRequestFlags) // procedure
	// SetFirstPartyForCookies
	//  Set the URL to the first party for cookies used in combination with
	//  ICefUrlRequest.
	SetFirstPartyForCookies(url string) // procedure
	// Assign
	//  Set all values at one time. This method corresponds to TCefRequest.set_ and cef_request_t.set
	Assign(url string, method string, postData ICefPostData, headerMap ICefStringMultimap) // procedure
}

// ICefRequestRef Parent: ICefRequest ICefBaseRefCountedRef
type ICefRequestRef interface {
	ICefRequest
	ICefBaseRefCountedRef
	AsIntfRequest() uintptr
}

type TCefRequestRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefRequestRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefRequestRef) GetUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestRef) GetMethod() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestRef) GetPostData() (result ICefPostData) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPostDataRef(resultPtr)
	return
}

func (m *TCefRequestRef) GetReferrerUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestRef) GetReferrerPolicy() cefTypes.TCefReferrerPolicy {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestRefAPI().SysCallN(6, m.Instance())
	return cefTypes.TCefReferrerPolicy(r)
}

func (m *TCefRequestRef) GetHeaderByName(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestRefAPI().SysCallN(7, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestRef) GetFlags() cefTypes.TCefUrlRequestFlags {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestRefAPI().SysCallN(8, m.Instance())
	return cefTypes.TCefUrlRequestFlags(r)
}

func (m *TCefRequestRef) GetFirstPartyForCookies() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestRef) GetResourceType() cefTypes.TCefResourceType {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestRefAPI().SysCallN(10, m.Instance())
	return cefTypes.TCefResourceType(r)
}

func (m *TCefRequestRef) GetTransitionType() cefTypes.TCefTransitionType {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestRefAPI().SysCallN(11, m.Instance())
	return cefTypes.TCefTransitionType(r)
}

func (m *TCefRequestRef) GetIdentifier() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefRequestRef) GetHeaderMap(headerMap ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(headerMap))
}

func (m *TCefRequestRef) SetUrl(value string) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(16, m.Instance(), api.PasStr(value))
}

func (m *TCefRequestRef) SetMethod(value string) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(17, m.Instance(), api.PasStr(value))
}

func (m *TCefRequestRef) SetReferrer(referrerUrl string, policy cefTypes.TCefReferrerPolicy) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(18, m.Instance(), api.PasStr(referrerUrl), uintptr(policy))
}

func (m *TCefRequestRef) SetPostData(value ICefPostData) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCefRequestRef) SetHeaderMap(headerMap ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(headerMap))
}

func (m *TCefRequestRef) SetHeaderByName(name string, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(21, m.Instance(), api.PasStr(name), api.PasStr(value), api.PasBool(overwrite))
}

func (m *TCefRequestRef) SetFlags(flags cefTypes.TCefUrlRequestFlags) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(22, m.Instance(), uintptr(flags))
}

func (m *TCefRequestRef) SetFirstPartyForCookies(url string) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(23, m.Instance(), api.PasStr(url))
}

func (m *TCefRequestRef) Assign(url string, method string, postData ICefPostData, headerMap ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cefRequestRefAPI().SysCallN(24, m.Instance(), api.PasStr(url), api.PasStr(method), base.GetObjectUintptr(postData), base.GetObjectUintptr(headerMap))
}

func (m *TCefRequestRef) AsIntfRequest() uintptr {
	return m.GetIntfPointer(0)
}

// RequestRef  is static instance
var RequestRef _RequestRefClass

// _RequestRefClass is class type defined by TCefRequestRef
type _RequestRefClass uintptr

func (_RequestRefClass) UnWrap(data uintptr) (result ICefRequest) {
	var resultPtr uintptr
	cefRequestRefAPI().SysCallN(13, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestRef(resultPtr)
	return
}

func (_RequestRefClass) New() (result ICefRequest) {
	var resultPtr uintptr
	cefRequestRefAPI().SysCallN(14, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestRef(resultPtr)
	return
}

// NewRequestRef class constructor
func NewRequestRef(data uintptr) ICefRequestRef {
	var requestPtr uintptr // ICefRequest
	r := cefRequestRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&requestPtr)))
	ret := AsCefRequestRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, requestPtr)
	}
	return ret
}

var (
	cefRequestRefOnce   base.Once
	cefRequestRefImport *imports.Imports = nil
)

func cefRequestRefAPI() *imports.Imports {
	cefRequestRefOnce.Do(func() {
		cefRequestRefImport = api.NewDefaultImports()
		cefRequestRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefRequestRef_Create", 0), // constructor NewRequestRef
			/* 1 */ imports.NewTable("TCefRequestRef_IsReadOnly", 0), // function IsReadOnly
			/* 2 */ imports.NewTable("TCefRequestRef_GetUrl", 0), // function GetUrl
			/* 3 */ imports.NewTable("TCefRequestRef_GetMethod", 0), // function GetMethod
			/* 4 */ imports.NewTable("TCefRequestRef_GetPostData", 0), // function GetPostData
			/* 5 */ imports.NewTable("TCefRequestRef_GetReferrerUrl", 0), // function GetReferrerUrl
			/* 6 */ imports.NewTable("TCefRequestRef_GetReferrerPolicy", 0), // function GetReferrerPolicy
			/* 7 */ imports.NewTable("TCefRequestRef_GetHeaderByName", 0), // function GetHeaderByName
			/* 8 */ imports.NewTable("TCefRequestRef_GetFlags", 0), // function GetFlags
			/* 9 */ imports.NewTable("TCefRequestRef_GetFirstPartyForCookies", 0), // function GetFirstPartyForCookies
			/* 10 */ imports.NewTable("TCefRequestRef_GetResourceType", 0), // function GetResourceType
			/* 11 */ imports.NewTable("TCefRequestRef_GetTransitionType", 0), // function GetTransitionType
			/* 12 */ imports.NewTable("TCefRequestRef_GetIdentifier", 0), // function GetIdentifier
			/* 13 */ imports.NewTable("TCefRequestRef_UnWrap", 0), // static function UnWrap
			/* 14 */ imports.NewTable("TCefRequestRef_New", 0), // static function New
			/* 15 */ imports.NewTable("TCefRequestRef_GetHeaderMap", 0), // procedure GetHeaderMap
			/* 16 */ imports.NewTable("TCefRequestRef_SetUrl", 0), // procedure SetUrl
			/* 17 */ imports.NewTable("TCefRequestRef_SetMethod", 0), // procedure SetMethod
			/* 18 */ imports.NewTable("TCefRequestRef_SetReferrer", 0), // procedure SetReferrer
			/* 19 */ imports.NewTable("TCefRequestRef_SetPostData", 0), // procedure SetPostData
			/* 20 */ imports.NewTable("TCefRequestRef_SetHeaderMap", 0), // procedure SetHeaderMap
			/* 21 */ imports.NewTable("TCefRequestRef_SetHeaderByName", 0), // procedure SetHeaderByName
			/* 22 */ imports.NewTable("TCefRequestRef_SetFlags", 0), // procedure SetFlags
			/* 23 */ imports.NewTable("TCefRequestRef_SetFirstPartyForCookies", 0), // procedure SetFirstPartyForCookies
			/* 24 */ imports.NewTable("TCefRequestRef_Assign", 0), // procedure Assign
		}
	})
	return cefRequestRefImport
}
