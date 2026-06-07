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

// ICefUrlRequest Parent: ICefBaseRefCounted
type ICefUrlRequest interface {
	ICefBaseRefCounted
	// GetRequest
	//  Returns the request object used to create this URL request. The returned
	//  object is read-only and should not be modified.
	GetRequest() ICefRequest // function
	// GetClient
	//  Returns the client.
	GetClient() IEngUrlrequestClient // function
	// GetRequestStatus
	//  Returns the request status.
	GetRequestStatus() cefTypes.TCefUrlRequestStatus // function
	// GetRequestError
	//  Returns the request error if status is UR_CANCELED or UR_FAILED, or 0
	//  otherwise.
	GetRequestError() int32 // function
	// GetResponse
	//  Returns the response, or NULL if no response information is available.
	//  Response information will only be available after the upload has
	//  completed. The returned object is read-only and should not be modified.
	GetResponse() ICefResponse // function
	// GetResponseWasCached
	//  Returns true (1) if the response body was served from the cache. This
	//  includes responses for which revalidation was required.
	GetResponseWasCached() bool // function
	// Cancel
	//  Cancel the request.
	Cancel() // procedure
}

// ICefUrlRequestRef Parent: ICefUrlRequest ICefBaseRefCountedRef
type ICefUrlRequestRef interface {
	ICefUrlRequest
	ICefBaseRefCountedRef
	AsIntfUrlRequest() uintptr
}

type TCefUrlRequestRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefUrlRequestRef) GetRequest() (result ICefRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefUrlRequestRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestRef(resultPtr)
	return
}

func (m *TCefUrlRequestRef) GetClient() (result IEngUrlrequestClient) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefUrlRequestRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngUrlrequestClient(resultPtr)
	return
}

func (m *TCefUrlRequestRef) GetRequestStatus() cefTypes.TCefUrlRequestStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefUrlRequestRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefUrlRequestStatus(r)
}

func (m *TCefUrlRequestRef) GetRequestError() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefUrlRequestRefAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCefUrlRequestRef) GetResponse() (result ICefResponse) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefUrlRequestRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResponseRef(resultPtr)
	return
}

func (m *TCefUrlRequestRef) GetResponseWasCached() bool {
	if !m.IsValid() {
		return false
	}
	r := cefUrlRequestRefAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCefUrlRequestRef) Cancel() {
	if !m.IsValid() {
		return
	}
	cefUrlRequestRefAPI().SysCallN(9, m.Instance())
}

func (m *TCefUrlRequestRef) AsIntfUrlRequest() uintptr {
	return m.GetIntfPointer(0)
}

// UrlRequestRef  is static instance
var UrlRequestRef _UrlRequestRefClass

// _UrlRequestRefClass is class type defined by TCefUrlRequestRef
type _UrlRequestRefClass uintptr

func (_UrlRequestRefClass) UnWrap(data uintptr) (result ICefUrlRequest) {
	var resultPtr uintptr
	cefUrlRequestRefAPI().SysCallN(7, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (_UrlRequestRefClass) New(request ICefRequest, client IEngUrlrequestClient, requestContext ICefRequestContext) (result ICefUrlRequest) {
	var resultPtr uintptr
	cefUrlRequestRefAPI().SysCallN(8, base.GetObjectUintptr(request), base.GetObjectUintptr(client), base.GetObjectUintptr(requestContext), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

// NewUrlRequestRef class constructor
func NewUrlRequestRef(data uintptr) ICefUrlRequestRef {
	var urlRequestPtr uintptr // ICefUrlRequest
	r := cefUrlRequestRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&urlRequestPtr)))
	ret := AsCefUrlRequestRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, urlRequestPtr)
	}
	return ret
}

var (
	cefUrlRequestRefOnce   base.Once
	cefUrlRequestRefImport *imports.Imports = nil
)

func cefUrlRequestRefAPI() *imports.Imports {
	cefUrlRequestRefOnce.Do(func() {
		cefUrlRequestRefImport = api.NewDefaultImports()
		cefUrlRequestRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefUrlRequestRef_Create", 0), // constructor NewUrlRequestRef
			/* 1 */ imports.NewTable("TCefUrlRequestRef_GetRequest", 0), // function GetRequest
			/* 2 */ imports.NewTable("TCefUrlRequestRef_GetClient", 0), // function GetClient
			/* 3 */ imports.NewTable("TCefUrlRequestRef_GetRequestStatus", 0), // function GetRequestStatus
			/* 4 */ imports.NewTable("TCefUrlRequestRef_GetRequestError", 0), // function GetRequestError
			/* 5 */ imports.NewTable("TCefUrlRequestRef_GetResponse", 0), // function GetResponse
			/* 6 */ imports.NewTable("TCefUrlRequestRef_GetResponseWasCached", 0), // function GetResponseWasCached
			/* 7 */ imports.NewTable("TCefUrlRequestRef_UnWrap", 0), // static function UnWrap
			/* 8 */ imports.NewTable("TCefUrlRequestRef_New", 0), // static function New
			/* 9 */ imports.NewTable("TCefUrlRequestRef_Cancel", 0), // procedure Cancel
		}
	})
	return cefUrlRequestRefImport
}
