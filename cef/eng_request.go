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

// ICefRequest Parent: ICefBaseRefCounted
type ICefRequest interface {
	ICefBaseRefCounted
	GetHeaderMap() ICefStringMultimap
	IsReadOnly() bool                                                               // function
	GetUrl() string                                                                 // function
	GetMethod() string                                                              // function
	GetPostData() ICefPostData                                                      // function
	GetReferrerUrl() string                                                         // function
	GetReferrerPolicy() TCefReferrerPolicy                                          // function
	GetHeaderByName(name string) string                                             // function
	GetFlags() TCefUrlRequestFlags                                                  // function
	GetFirstPartyForCookies() string                                                // function
	GetResourceType() TCefResourceType                                              // function
	GetTransitionType() TCefTransitionType                                          // function
	GetIdentifier() (resultUint64 uint64)                                           // function
	SetUrl(value string)                                                            // procedure
	SetMethod(value string)                                                         // procedure
	SetReferrer(referrerUrl string, policy TCefReferrerPolicy)                      // procedure
	SetPostData(value ICefPostData)                                                 // procedure
	SetHeaderMap(headerMap ICefStringMultimap)                                      // procedure
	SetHeaderByName(name, value string, overwrite bool)                             // procedure
	SetFlags(flags TCefUrlRequestFlags)                                             // procedure
	SetFirstPartyForCookies(url string)                                             // procedure
	Assign(url, method string, postData ICefPostData, headerMap ICefStringMultimap) // procedure
}

// TCefRequest Parent: TCefBaseRefCounted
type TCefRequest struct {
	TCefBaseRefCounted
}

// RequestRef -> ICefRequest
var RequestRef request

// request TCefRequest Ref
type request uintptr

func (m *request) UnWrap(data uintptr) ICefRequest {
	var resultCefRequest uintptr
	requestImportAPI().SysCallN(22, uintptr(data), uintptr(unsafePointer(&resultCefRequest)))
	return AsCefRequest(resultCefRequest)
}

func (m *request) New() ICefRequest {
	var resultCefRequest uintptr
	requestImportAPI().SysCallN(13, uintptr(unsafePointer(&resultCefRequest)))
	return AsCefRequest(resultCefRequest)
}

func (m *TCefRequest) IsReadOnly() bool {
	r1 := requestImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TCefRequest) GetUrl() string {
	r1 := requestImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetMethod() string {
	r1 := requestImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetPostData() ICefPostData {
	var resultCefPostData uintptr
	requestImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefPostData)))
	return AsCefPostData(resultCefPostData)
}

func (m *TCefRequest) GetReferrerUrl() string {
	r1 := requestImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetReferrerPolicy() TCefReferrerPolicy {
	r1 := requestImportAPI().SysCallN(7, m.Instance())
	return TCefReferrerPolicy(r1)
}

func (m *TCefRequest) GetHeaderByName(name string) string {
	r1 := requestImportAPI().SysCallN(3, m.Instance(), PascalStr(name))
	return GoStr(r1)
}

func (m *TCefRequest) GetFlags() TCefUrlRequestFlags {
	r1 := requestImportAPI().SysCallN(2, m.Instance())
	return TCefUrlRequestFlags(r1)
}

func (m *TCefRequest) GetFirstPartyForCookies() string {
	r1 := requestImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequest) GetResourceType() TCefResourceType {
	r1 := requestImportAPI().SysCallN(9, m.Instance())
	return TCefResourceType(r1)
}

func (m *TCefRequest) GetTransitionType() TCefTransitionType {
	r1 := requestImportAPI().SysCallN(10, m.Instance())
	return TCefTransitionType(r1)
}

func (m *TCefRequest) GetIdentifier() (resultUint64 uint64) {
	requestImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCefRequest) SetUrl(value string) {
	requestImportAPI().SysCallN(21, m.Instance(), PascalStr(value))
}

func (m *TCefRequest) SetMethod(value string) {
	requestImportAPI().SysCallN(18, m.Instance(), PascalStr(value))
}

func (m *TCefRequest) SetReferrer(referrerUrl string, policy TCefReferrerPolicy) {
	requestImportAPI().SysCallN(20, m.Instance(), PascalStr(referrerUrl), uintptr(policy))
}

func (m *TCefRequest) SetPostData(value ICefPostData) {
	requestImportAPI().SysCallN(19, m.Instance(), GetObjectUintptr(value))
}

func (m *TCefRequest) SetHeaderMap(headerMap ICefStringMultimap) {
	requestImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefRequest) SetHeaderByName(name, value string, overwrite bool) {
	requestImportAPI().SysCallN(16, m.Instance(), PascalStr(name), PascalStr(value), PascalBool(overwrite))
}

func (m *TCefRequest) SetFlags(flags TCefUrlRequestFlags) {
	requestImportAPI().SysCallN(15, m.Instance(), uintptr(flags))
}

func (m *TCefRequest) SetFirstPartyForCookies(url string) {
	requestImportAPI().SysCallN(14, m.Instance(), PascalStr(url))
}

func (m *TCefRequest) Assign(url, method string, postData ICefPostData, headerMap ICefStringMultimap) {
	requestImportAPI().SysCallN(0, m.Instance(), PascalStr(url), PascalStr(method), GetObjectUintptr(postData), GetObjectUintptr(headerMap))
}

var (
	requestImport       *imports.Imports = nil
	requestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefRequest_Assign", 0),
		/*1*/ imports.NewTable("CefRequest_GetFirstPartyForCookies", 0),
		/*2*/ imports.NewTable("CefRequest_GetFlags", 0),
		/*3*/ imports.NewTable("CefRequest_GetHeaderByName", 0),
		/*4*/ imports.NewTable("CefRequest_GetIdentifier", 0),
		/*5*/ imports.NewTable("CefRequest_GetMethod", 0),
		/*6*/ imports.NewTable("CefRequest_GetPostData", 0),
		/*7*/ imports.NewTable("CefRequest_GetReferrerPolicy", 0),
		/*8*/ imports.NewTable("CefRequest_GetReferrerUrl", 0),
		/*9*/ imports.NewTable("CefRequest_GetResourceType", 0),
		/*10*/ imports.NewTable("CefRequest_GetTransitionType", 0),
		/*11*/ imports.NewTable("CefRequest_GetUrl", 0),
		/*12*/ imports.NewTable("CefRequest_IsReadOnly", 0),
		/*13*/ imports.NewTable("CefRequest_New", 0),
		/*14*/ imports.NewTable("CefRequest_SetFirstPartyForCookies", 0),
		/*15*/ imports.NewTable("CefRequest_SetFlags", 0),
		/*16*/ imports.NewTable("CefRequest_SetHeaderByName", 0),
		/*17*/ imports.NewTable("CefRequest_SetHeaderMap", 0),
		/*18*/ imports.NewTable("CefRequest_SetMethod", 0),
		/*19*/ imports.NewTable("CefRequest_SetPostData", 0),
		/*20*/ imports.NewTable("CefRequest_SetReferrer", 0),
		/*21*/ imports.NewTable("CefRequest_SetUrl", 0),
		/*22*/ imports.NewTable("CefRequest_UnWrap", 0),
	}
)

func requestImportAPI() *imports.Imports {
	if requestImport == nil {
		requestImport = NewDefaultImports()
		requestImport.SetImportTable(requestImportTables)
		requestImportTables = nil
	}
	return requestImport
}
