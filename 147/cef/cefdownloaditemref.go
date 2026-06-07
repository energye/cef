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

	cefTypes "github.com/energye/cef/147/types"
)

// ICefDownloadItem Parent: ICefBaseRefCounted
type ICefDownloadItem interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions
	//  if this function returns false (0).
	IsValid() bool // function
	// IsInProgress
	//  Returns true (1) if the download is in progress.
	IsInProgress() bool // function
	// IsComplete
	//  Returns true (1) if the download is complete.
	IsComplete() bool // function
	// IsCanceled
	//  Returns true (1) if the download has been canceled.
	IsCanceled() bool // function
	// IsInterrupted
	//  Returns true (1) if the download has been interrupted.
	IsInterrupted() bool // function
	// GetInterruptReason
	//  Returns the most recent interrupt reason.
	GetInterruptReason() cefTypes.TCefDownloadInterruptReason // function
	// GetCurrentSpeed
	//  Returns a simple speed estimate in bytes/s.
	GetCurrentSpeed() int64 // function
	// GetPercentComplete
	//  Returns the rough percent complete or -1 if the receive total size is
	//  unknown.
	GetPercentComplete() int32 // function
	// GetTotalBytes
	//  Returns the total number of bytes.
	GetTotalBytes() int64 // function
	// GetReceivedBytes
	//  Returns the number of received bytes.
	GetReceivedBytes() int64 // function
	// GetStartTime
	//  Returns the time that the download started.
	GetStartTime() types.TDateTime // function
	// GetEndTime
	//  Returns the time that the download ended.
	GetEndTime() types.TDateTime // function
	// GetFullPath
	//  Returns the full path to the downloaded or downloading file.
	GetFullPath() string // function
	// GetId
	//  Returns the unique identifier for this download.
	GetId() uint32 // function
	// GetUrl
	//  Returns the URL.
	GetUrl() string // function
	// GetOriginalUrl
	//  Returns the original URL before any redirections.
	GetOriginalUrl() string // function
	// GetSuggestedFileName
	//  Returns the suggested file name.
	GetSuggestedFileName() string // function
	// GetContentDisposition
	//  Returns the content disposition.
	GetContentDisposition() string // function
	// GetMimeType
	//  Returns the mime type.
	GetMimeType() string // function
	// IsPaused
	//  Returns true (1) if the download has been paused.
	IsPaused() bool // function
}

// ICefDownloadItemRef Parent: ICefDownloadItem ICefBaseRefCountedRef
type ICefDownloadItemRef interface {
	ICefDownloadItem
	ICefBaseRefCountedRef
	AsIntfDownloadItem() uintptr
}

type TCefDownloadItemRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDownloadItemRef) IsValid() bool {
	if !m.TBase.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) IsInProgress() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) IsComplete() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) IsCanceled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) IsInterrupted() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) GetInterruptReason() cefTypes.TCefDownloadInterruptReason {
	if !m.IsValid() {
		return 0
	}
	r := cefDownloadItemRefAPI().SysCallN(6, m.Instance())
	return cefTypes.TCefDownloadInterruptReason(r)
}

func (m *TCefDownloadItemRef) GetCurrentSpeed() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownloadItemRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownloadItemRef) GetPercentComplete() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDownloadItemRefAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TCefDownloadItemRef) GetTotalBytes() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownloadItemRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownloadItemRef) GetReceivedBytes() (result int64) {
	if !m.IsValid() {
		return
	}
	cefDownloadItemRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownloadItemRef) GetStartTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefDownloadItemRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownloadItemRef) GetEndTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	cefDownloadItemRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDownloadItemRef) GetFullPath() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) GetId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cefDownloadItemRefAPI().SysCallN(14, m.Instance())
	return uint32(r)
}

func (m *TCefDownloadItemRef) GetUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) GetOriginalUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) GetSuggestedFileName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) GetContentDisposition() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) GetMimeType() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDownloadItemRefAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDownloadItemRef) IsPaused() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDownloadItemRefAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDownloadItemRef) AsIntfDownloadItem() uintptr {
	return m.GetIntfPointer(0)
}

// DownloadItemRef  is static instance
var DownloadItemRef _DownloadItemRefClass

// _DownloadItemRefClass is class type defined by TCefDownloadItemRef
type _DownloadItemRefClass uintptr

func (_DownloadItemRefClass) UnWrap(data uintptr) (result ICefDownloadItem) {
	var resultPtr uintptr
	cefDownloadItemRefAPI().SysCallN(21, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDownloadItemRef(resultPtr)
	return
}

// NewDownloadItemRef class constructor
func NewDownloadItemRef(data uintptr) ICefDownloadItemRef {
	var downloadItemPtr uintptr // ICefDownloadItem
	r := cefDownloadItemRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&downloadItemPtr)))
	ret := AsCefDownloadItemRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, downloadItemPtr)
	}
	return ret
}

var (
	cefDownloadItemRefOnce   base.Once
	cefDownloadItemRefImport *imports.Imports = nil
)

func cefDownloadItemRefAPI() *imports.Imports {
	cefDownloadItemRefOnce.Do(func() {
		cefDownloadItemRefImport = api.NewDefaultImports()
		cefDownloadItemRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDownloadItemRef_Create", 0), // constructor NewDownloadItemRef
			/* 1 */ imports.NewTable("TCefDownloadItemRef_IsValid", 0), // function IsValid
			/* 2 */ imports.NewTable("TCefDownloadItemRef_IsInProgress", 0), // function IsInProgress
			/* 3 */ imports.NewTable("TCefDownloadItemRef_IsComplete", 0), // function IsComplete
			/* 4 */ imports.NewTable("TCefDownloadItemRef_IsCanceled", 0), // function IsCanceled
			/* 5 */ imports.NewTable("TCefDownloadItemRef_IsInterrupted", 0), // function IsInterrupted
			/* 6 */ imports.NewTable("TCefDownloadItemRef_GetInterruptReason", 0), // function GetInterruptReason
			/* 7 */ imports.NewTable("TCefDownloadItemRef_GetCurrentSpeed", 0), // function GetCurrentSpeed
			/* 8 */ imports.NewTable("TCefDownloadItemRef_GetPercentComplete", 0), // function GetPercentComplete
			/* 9 */ imports.NewTable("TCefDownloadItemRef_GetTotalBytes", 0), // function GetTotalBytes
			/* 10 */ imports.NewTable("TCefDownloadItemRef_GetReceivedBytes", 0), // function GetReceivedBytes
			/* 11 */ imports.NewTable("TCefDownloadItemRef_GetStartTime", 0), // function GetStartTime
			/* 12 */ imports.NewTable("TCefDownloadItemRef_GetEndTime", 0), // function GetEndTime
			/* 13 */ imports.NewTable("TCefDownloadItemRef_GetFullPath", 0), // function GetFullPath
			/* 14 */ imports.NewTable("TCefDownloadItemRef_GetId", 0), // function GetId
			/* 15 */ imports.NewTable("TCefDownloadItemRef_GetUrl", 0), // function GetUrl
			/* 16 */ imports.NewTable("TCefDownloadItemRef_GetOriginalUrl", 0), // function GetOriginalUrl
			/* 17 */ imports.NewTable("TCefDownloadItemRef_GetSuggestedFileName", 0), // function GetSuggestedFileName
			/* 18 */ imports.NewTable("TCefDownloadItemRef_GetContentDisposition", 0), // function GetContentDisposition
			/* 19 */ imports.NewTable("TCefDownloadItemRef_GetMimeType", 0), // function GetMimeType
			/* 20 */ imports.NewTable("TCefDownloadItemRef_IsPaused", 0), // function IsPaused
			/* 21 */ imports.NewTable("TCefDownloadItemRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefDownloadItemRefImport
}
