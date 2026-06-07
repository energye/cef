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
)

// ICefDownloadItem Parent: ICefBaseRefCounted
type ICefDownloadItem interface {
	ICefBaseRefCounted
	IsValid() bool                 // function
	IsInProgress() bool            // function
	IsComplete() bool              // function
	IsCanceled() bool              // function
	GetCurrentSpeed() int64        // function
	GetPercentComplete() int32     // function
	GetTotalBytes() int64          // function
	GetReceivedBytes() int64       // function
	GetStartTime() types.TDateTime // function
	GetEndTime() types.TDateTime   // function
	GetFullPath() string           // function
	GetId() uint32                 // function
	GetUrl() string                // function
	GetOriginalUrl() string        // function
	GetSuggestedFileName() string  // function
	GetContentDisposition() string // function
	GetMimeType() string           // function
}

// ICefDownLoadItemRef Parent: ICefDownloadItem ICefBaseRefCountedRef
type ICefDownLoadItemRef interface {
	ICefDownloadItem
	ICefBaseRefCountedRef
	AsIntfDownloadItem() uintptr
}

type TCefDownLoadItemRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDownLoadItemRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefDownLoadItemRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownLoadItemRef) IsInProgress() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownLoadItemRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownLoadItemRef) IsComplete() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownLoadItemRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownLoadItemRef) IsCanceled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownLoadItemRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownLoadItemRef) GetCurrentSpeed() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownLoadItemRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownLoadItemRef) GetPercentComplete() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDownLoadItemRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefDownLoadItemRef) GetTotalBytes() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownLoadItemRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownLoadItemRef) GetReceivedBytes() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownLoadItemRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownLoadItemRef) GetStartTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefDownLoadItemRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownLoadItemRef) GetEndTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefDownLoadItemRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownLoadItemRef) GetFullPath() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) GetId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDownLoadItemRefAPI().SysCallN(12, m.Instance())
	return uint32(r)
}

func (m *TCefDownLoadItemRef) GetUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) GetOriginalUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) GetSuggestedFileName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) GetContentDisposition() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) GetMimeType() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownLoadItemRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownLoadItemRef) AsIntfDownloadItem() uintptr {
	return m.GetIntfPointer(0)
}

// DownLoadItemRef  is static instance
var DownLoadItemRef _DownLoadItemRefClass

// _DownLoadItemRefClass is class type defined by TCefDownLoadItemRef
type _DownLoadItemRefClass uintptr

func (_DownLoadItemRefClass) UnWrap(data uintptr) (result ICefDownloadItem) {
	var resultPtr uintptr
	cefDownLoadItemRefAPI().SysCallN(18, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDownLoadItemRef(resultPtr)
	return
}

// NewDownLoadItemRef class constructor
func NewDownLoadItemRef(data uintptr) ICefDownLoadItemRef {
	var downloadItemPtr uintptr // ICefDownloadItem
	r := cefDownLoadItemRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&downloadItemPtr)))
	ret := AsCefDownLoadItemRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadItemPtr)
	}
	return ret
}

var (
	cefDownLoadItemRefOnce   base.Once
	cefDownLoadItemRefImport *imports.Imports = nil
)

func cefDownLoadItemRefAPI() *imports.Imports {
	cefDownLoadItemRefOnce.Do(func() {
		cefDownLoadItemRefImport = api.NewDefaultImports()
		cefDownLoadItemRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDownLoadItemRef_Create", 0), // constructor NewDownLoadItemRef
			/* 1 */ imports.NewTable("TCefDownLoadItemRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefDownLoadItemRef_IsInProgress", 0), // function IsInProgress
			/* 3 */ imports.NewTable("TCefDownLoadItemRef_IsComplete", 0), // function IsComplete
			/* 4 */ imports.NewTable("TCefDownLoadItemRef_IsCanceled", 0), // function IsCanceled
			/* 5 */ imports.NewTable("TCefDownLoadItemRef_GetCurrentSpeed", 0), // function GetCurrentSpeed
			/* 6 */ imports.NewTable("TCefDownLoadItemRef_GetPercentComplete", 0), // function GetPercentComplete
			/* 7 */ imports.NewTable("TCefDownLoadItemRef_GetTotalBytes", 0), // function GetTotalBytes
			/* 8 */ imports.NewTable("TCefDownLoadItemRef_GetReceivedBytes", 0), // function GetReceivedBytes
			/* 9 */ imports.NewTable("TCefDownLoadItemRef_GetStartTime", 0), // function GetStartTime
			/* 10 */ imports.NewTable("TCefDownLoadItemRef_GetEndTime", 0), // function GetEndTime
			/* 11 */ imports.NewTable("TCefDownLoadItemRef_GetFullPath", 0), // function GetFullPath
			/* 12 */ imports.NewTable("TCefDownLoadItemRef_GetId", 0), // function GetId
			/* 13 */ imports.NewTable("TCefDownLoadItemRef_GetUrl", 0), // function GetUrl
			/* 14 */ imports.NewTable("TCefDownLoadItemRef_GetOriginalUrl", 0), // function GetOriginalUrl
			/* 15 */ imports.NewTable("TCefDownLoadItemRef_GetSuggestedFileName", 0), // function GetSuggestedFileName
			/* 16 */ imports.NewTable("TCefDownLoadItemRef_GetContentDisposition", 0), // function GetContentDisposition
			/* 17 */ imports.NewTable("TCefDownLoadItemRef_GetMimeType", 0), // function GetMimeType
			/* 18 */ imports.NewTable("TCefDownLoadItemRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefDownLoadItemRefImport
}
