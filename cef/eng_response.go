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

// ICefResponse Parent: ICefBaseRefCounted
type ICefResponse interface {
	ICefBaseRefCounted
	GetHeaderMap() ICefStringMultimap
	IsReadOnly() bool                                   // function
	GetError() TCefErrorCode                            // function
	GetStatus() int32                                   // function
	GetStatusText() string                              // function
	GetMimeType() string                                // function
	GetCharset() string                                 // function
	GetHeaderByName(name string) string                 // function
	GetURL() string                                     // function
	SetError(error TCefErrorCode)                       // procedure
	SetStatus(status int32)                             // procedure
	SetStatusText(statusText string)                    // procedure
	SetMimeType(mimetype string)                        // procedure
	SetCharset(charset string)                          // procedure
	SetHeaderByName(name, value string, overwrite bool) // procedure
	SetHeaderMap(headerMap ICefStringMultimap)          // procedure
	SetURL(url string)                                  // procedure
}

// TCefResponse Parent: TCefBaseRefCounted
type TCefResponse struct {
	TCefBaseRefCounted
}

// ResponseRef -> ICefResponse
var ResponseRef response

// response TCefResponse Ref
type response uintptr

func (m *response) UnWrap(data uintptr) ICefResponse {
	var resultCefResponse uintptr
	responseImportAPI().SysCallN(17, uintptr(data), uintptr(unsafePointer(&resultCefResponse)))
	return AsCefResponse(resultCefResponse)
}

func (m *response) New() ICefResponse {
	var resultCefResponse uintptr
	responseImportAPI().SysCallN(8, uintptr(unsafePointer(&resultCefResponse)))
	return AsCefResponse(resultCefResponse)
}

func (m *TCefResponse) IsReadOnly() bool {
	r1 := responseImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCefResponse) GetError() TCefErrorCode {
	r1 := responseImportAPI().SysCallN(1, m.Instance())
	return TCefErrorCode(r1)
}

func (m *TCefResponse) GetStatus() int32 {
	r1 := responseImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TCefResponse) GetStatusText() string {
	r1 := responseImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetMimeType() string {
	r1 := responseImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetCharset() string {
	r1 := responseImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) GetHeaderByName(name string) string {
	r1 := responseImportAPI().SysCallN(2, m.Instance(), PascalStr(name))
	return GoStr(r1)
}

func (m *TCefResponse) GetURL() string {
	r1 := responseImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCefResponse) SetError(error TCefErrorCode) {
	responseImportAPI().SysCallN(10, m.Instance(), uintptr(error))
}

func (m *TCefResponse) SetStatus(status int32) {
	responseImportAPI().SysCallN(14, m.Instance(), uintptr(status))
}

func (m *TCefResponse) SetStatusText(statusText string) {
	responseImportAPI().SysCallN(15, m.Instance(), PascalStr(statusText))
}

func (m *TCefResponse) SetMimeType(mimetype string) {
	responseImportAPI().SysCallN(13, m.Instance(), PascalStr(mimetype))
}

func (m *TCefResponse) SetCharset(charset string) {
	responseImportAPI().SysCallN(9, m.Instance(), PascalStr(charset))
}

func (m *TCefResponse) SetHeaderByName(name, value string, overwrite bool) {
	responseImportAPI().SysCallN(11, m.Instance(), PascalStr(name), PascalStr(value), PascalBool(overwrite))
}

func (m *TCefResponse) SetHeaderMap(headerMap ICefStringMultimap) {
	responseImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(headerMap))
}

func (m *TCefResponse) SetURL(url string) {
	responseImportAPI().SysCallN(16, m.Instance(), PascalStr(url))
}

var (
	responseImport       *imports.Imports = nil
	responseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefResponse_GetCharset", 0),
		/*1*/ imports.NewTable("CefResponse_GetError", 0),
		/*2*/ imports.NewTable("CefResponse_GetHeaderByName", 0),
		/*3*/ imports.NewTable("CefResponse_GetMimeType", 0),
		/*4*/ imports.NewTable("CefResponse_GetStatus", 0),
		/*5*/ imports.NewTable("CefResponse_GetStatusText", 0),
		/*6*/ imports.NewTable("CefResponse_GetURL", 0),
		/*7*/ imports.NewTable("CefResponse_IsReadOnly", 0),
		/*8*/ imports.NewTable("CefResponse_New", 0),
		/*9*/ imports.NewTable("CefResponse_SetCharset", 0),
		/*10*/ imports.NewTable("CefResponse_SetError", 0),
		/*11*/ imports.NewTable("CefResponse_SetHeaderByName", 0),
		/*12*/ imports.NewTable("CefResponse_SetHeaderMap", 0),
		/*13*/ imports.NewTable("CefResponse_SetMimeType", 0),
		/*14*/ imports.NewTable("CefResponse_SetStatus", 0),
		/*15*/ imports.NewTable("CefResponse_SetStatusText", 0),
		/*16*/ imports.NewTable("CefResponse_SetURL", 0),
		/*17*/ imports.NewTable("CefResponse_UnWrap", 0),
	}
)

func responseImportAPI() *imports.Imports {
	if responseImport == nil {
		responseImport = NewDefaultImports()
		responseImport.SetImportTable(responseImportTables)
		responseImportTables = nil
	}
	return responseImport
}
