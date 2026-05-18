//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefResponse Parent: ICefBaseRefCounted
type ICefResponse interface {
	ICefBaseRefCounted
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// GetError
	//  Get the response error code. Returns ERR_NONE if there was no error.
	GetError() cefTypes.TCefErrorCode // function
	// GetStatus
	//  Get the response status code.
	GetStatus() int32 // function
	// GetStatusText
	//  Get the response status text.
	GetStatusText() string // function
	// GetMimeType
	//  Get the response mime type.
	GetMimeType() string // function
	// GetCharset
	//  Get the response charset.
	GetCharset() string // function
	// GetHeaderByName
	//  Get the value for the specified response header field.
	GetHeaderByName(name string) string // function
	// GetURL
	//  Get the resolved URL after redirects or changed as a result of HSTS.
	GetURL() string // function
	// SetError
	//  Set the response error code. This can be used by custom scheme handlers to
	//  return errors during initial request processing.
	SetError(error_ cefTypes.TCefErrorCode) // procedure
	// SetStatus
	//  Set the response status code.
	SetStatus(status int32) // procedure
	// SetStatusText
	//  Set the response status text.
	SetStatusText(statusText string) // procedure
	// SetMimeType
	//  Set the response mime type.
	SetMimeType(mimetype string) // procedure
	// SetCharset
	//  Set the response charset.
	SetCharset(charset string) // procedure
	// SetHeaderByName
	//  Set the header |name| to |value|. If |overwrite| is true (1) any existing
	//  values will be replaced with the new value. If |overwrite| is false (0)
	//  any existing values will not be overwritten.
	SetHeaderByName(name string, value string, overwrite bool) // procedure
	// GetHeaderMap
	//  Get all response header fields.
	GetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetHeaderMap
	//  Set all response header fields.
	SetHeaderMap(headerMap ICefStringMultimap) // procedure
	// SetURL
	//  Set the resolved URL after redirects or changed as a result of HSTS.
	SetURL(url string) // procedure
}

// ICefResponseRef Parent: ICefResponse ICefBaseRefCountedRef
type ICefResponseRef interface {
	ICefResponse
	ICefBaseRefCountedRef
	AsIntfResponse() uintptr
}

type TCefResponseRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefResponseRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefResponseRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefResponseRef) GetError() cefTypes.TCefErrorCode {
	if !m.IsValid() {
		return 0
	}
	r := cefResponseRefAPI().SysCallN(2, m.Instance())
	return cefTypes.TCefErrorCode(r)
}

func (m *TCefResponseRef) GetStatus() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefResponseRefAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TCefResponseRef) GetStatusText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResponseRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResponseRef) GetMimeType() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResponseRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResponseRef) GetCharset() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResponseRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResponseRef) GetHeaderByName(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResponseRefAPI().SysCallN(7, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResponseRef) GetURL() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefResponseRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefResponseRef) SetError(error_ cefTypes.TCefErrorCode) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(11, m.Instance(), uintptr(error_))
}

func (m *TCefResponseRef) SetStatus(status int32) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(12, m.Instance(), uintptr(status))
}

func (m *TCefResponseRef) SetStatusText(statusText string) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(13, m.Instance(), api.PasStr(statusText))
}

func (m *TCefResponseRef) SetMimeType(mimetype string) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(14, m.Instance(), api.PasStr(mimetype))
}

func (m *TCefResponseRef) SetCharset(charset string) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(15, m.Instance(), api.PasStr(charset))
}

func (m *TCefResponseRef) SetHeaderByName(name string, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(16, m.Instance(), api.PasStr(name), api.PasStr(value), api.PasBool(overwrite))
}

func (m *TCefResponseRef) GetHeaderMap(headerMap ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(headerMap))
}

func (m *TCefResponseRef) SetHeaderMap(headerMap ICefStringMultimap) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(18, m.Instance(), base.GetObjectUintptr(headerMap))
}

func (m *TCefResponseRef) SetURL(url string) {
	if !m.IsValid() {
		return
	}
	cefResponseRefAPI().SysCallN(19, m.Instance(), api.PasStr(url))
}

func (m *TCefResponseRef) AsIntfResponse() uintptr {
	return m.GetIntfPointer(0)
}

// ResponseRef  is static instance
var ResponseRef _ResponseRefClass

// _ResponseRefClass is class type defined by TCefResponseRef
type _ResponseRefClass uintptr

func (_ResponseRefClass) UnWrap(data uintptr) (result ICefResponse) {
	var resultPtr uintptr
	cefResponseRefAPI().SysCallN(9, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResponseRef(resultPtr)
	return
}

func (_ResponseRefClass) New() (result ICefResponse) {
	var resultPtr uintptr
	cefResponseRefAPI().SysCallN(10, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefResponseRef(resultPtr)
	return
}

// NewResponseRef class constructor
func NewResponseRef(data uintptr) ICefResponseRef {
	var responsePtr uintptr // ICefResponse
	r := cefResponseRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&responsePtr)))
	ret := AsCefResponseRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, responsePtr)
	}
	return ret
}

var (
	cefResponseRefOnce   base.Once
	cefResponseRefImport *imports.Imports = nil
)

func cefResponseRefAPI() *imports.Imports {
	cefResponseRefOnce.Do(func() {
		cefResponseRefImport = api.NewDefaultImports()
		cefResponseRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefResponseRef_Create", 0), // constructor NewResponseRef
			/* 1 */ imports.NewTable("TCefResponseRef_IsReadOnly", 0), // function IsReadOnly
			/* 2 */ imports.NewTable("TCefResponseRef_GetError", 0), // function GetError
			/* 3 */ imports.NewTable("TCefResponseRef_GetStatus", 0), // function GetStatus
			/* 4 */ imports.NewTable("TCefResponseRef_GetStatusText", 0), // function GetStatusText
			/* 5 */ imports.NewTable("TCefResponseRef_GetMimeType", 0), // function GetMimeType
			/* 6 */ imports.NewTable("TCefResponseRef_GetCharset", 0), // function GetCharset
			/* 7 */ imports.NewTable("TCefResponseRef_GetHeaderByName", 0), // function GetHeaderByName
			/* 8 */ imports.NewTable("TCefResponseRef_GetURL", 0), // function GetURL
			/* 9 */ imports.NewTable("TCefResponseRef_UnWrap", 0), // static function UnWrap
			/* 10 */ imports.NewTable("TCefResponseRef_New", 0), // static function New
			/* 11 */ imports.NewTable("TCefResponseRef_SetError", 0), // procedure SetError
			/* 12 */ imports.NewTable("TCefResponseRef_SetStatus", 0), // procedure SetStatus
			/* 13 */ imports.NewTable("TCefResponseRef_SetStatusText", 0), // procedure SetStatusText
			/* 14 */ imports.NewTable("TCefResponseRef_SetMimeType", 0), // procedure SetMimeType
			/* 15 */ imports.NewTable("TCefResponseRef_SetCharset", 0), // procedure SetCharset
			/* 16 */ imports.NewTable("TCefResponseRef_SetHeaderByName", 0), // procedure SetHeaderByName
			/* 17 */ imports.NewTable("TCefResponseRef_GetHeaderMap", 0), // procedure GetHeaderMap
			/* 18 */ imports.NewTable("TCefResponseRef_SetHeaderMap", 0), // procedure SetHeaderMap
			/* 19 */ imports.NewTable("TCefResponseRef_SetURL", 0), // procedure SetURL
		}
	})
	return cefResponseRefImport
}
