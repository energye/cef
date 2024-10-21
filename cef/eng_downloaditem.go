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

// ICefDownloadItem Parent: ICefBaseRefCounted
//
//	Interface used to represent a download item.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_item_capi.h">CEF source file: /include/capi/cef_download_item_capi.h (cef_download_item_t))</a>
type ICefDownloadItem interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
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
	GetInterruptReason() TCefDownloadInterruptReason // function
	// GetCurrentSpeed
	//  Returns a simple speed estimate in bytes/s.
	GetCurrentSpeed() (resultInt64 int64) // function
	// GetPercentComplete
	//  Returns the rough percent complete or -1 if the receive total size is unknown.
	GetPercentComplete() int32 // function
	// GetTotalBytes
	//  Returns the total number of bytes.
	GetTotalBytes() (resultInt64 int64) // function
	// GetReceivedBytes
	//  Returns the number of received bytes.
	GetReceivedBytes() (resultInt64 int64) // function
	// GetStartTime
	//  Returns the time that the download started.
	GetStartTime() (resultDateTime TDateTime) // function
	// GetEndTime
	//  Returns the time that the download ended.
	GetEndTime() (resultDateTime TDateTime) // function
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
}

// TCefDownloadItem Parent: TCefBaseRefCounted
//
//	Interface used to represent a download item.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_item_capi.h">CEF source file: /include/capi/cef_download_item_capi.h (cef_download_item_t))</a>
type TCefDownloadItem struct {
	TCefBaseRefCounted
}

// DownloadItemRef -> ICefDownloadItem
var DownloadItemRef downloadItem

// downloadItem TCefDownloadItem Ref
type downloadItem uintptr

func (m *downloadItem) UnWrap(data uintptr) ICefDownloadItem {
	var resultCefDownloadItem uintptr
	downloadItemImportAPI().SysCallN(19, uintptr(data), uintptr(unsafePointer(&resultCefDownloadItem)))
	return AsCefDownloadItem(resultCefDownloadItem)
}

func (m *TCefDownloadItem) IsValid() bool {
	r1 := downloadItemImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsInProgress() bool {
	r1 := downloadItemImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsComplete() bool {
	r1 := downloadItemImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsCanceled() bool {
	r1 := downloadItemImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsInterrupted() bool {
	r1 := downloadItemImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) GetInterruptReason() TCefDownloadInterruptReason {
	r1 := downloadItemImportAPI().SysCallN(5, m.Instance())
	return TCefDownloadInterruptReason(r1)
}

func (m *TCefDownloadItem) GetCurrentSpeed() (resultInt64 int64) {
	downloadItemImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetPercentComplete() int32 {
	r1 := downloadItemImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func (m *TCefDownloadItem) GetTotalBytes() (resultInt64 int64) {
	downloadItemImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetReceivedBytes() (resultInt64 int64) {
	downloadItemImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetStartTime() (resultDateTime TDateTime) {
	downloadItemImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefDownloadItem) GetEndTime() (resultDateTime TDateTime) {
	downloadItemImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefDownloadItem) GetFullPath() string {
	r1 := downloadItemImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetId() uint32 {
	r1 := downloadItemImportAPI().SysCallN(4, m.Instance())
	return uint32(r1)
}

func (m *TCefDownloadItem) GetUrl() string {
	r1 := downloadItemImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetOriginalUrl() string {
	r1 := downloadItemImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetSuggestedFileName() string {
	r1 := downloadItemImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetContentDisposition() string {
	r1 := downloadItemImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetMimeType() string {
	r1 := downloadItemImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

var (
	downloadItemImport       *imports.Imports = nil
	downloadItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDownloadItem_GetContentDisposition", 0),
		/*1*/ imports.NewTable("CefDownloadItem_GetCurrentSpeed", 0),
		/*2*/ imports.NewTable("CefDownloadItem_GetEndTime", 0),
		/*3*/ imports.NewTable("CefDownloadItem_GetFullPath", 0),
		/*4*/ imports.NewTable("CefDownloadItem_GetId", 0),
		/*5*/ imports.NewTable("CefDownloadItem_GetInterruptReason", 0),
		/*6*/ imports.NewTable("CefDownloadItem_GetMimeType", 0),
		/*7*/ imports.NewTable("CefDownloadItem_GetOriginalUrl", 0),
		/*8*/ imports.NewTable("CefDownloadItem_GetPercentComplete", 0),
		/*9*/ imports.NewTable("CefDownloadItem_GetReceivedBytes", 0),
		/*10*/ imports.NewTable("CefDownloadItem_GetStartTime", 0),
		/*11*/ imports.NewTable("CefDownloadItem_GetSuggestedFileName", 0),
		/*12*/ imports.NewTable("CefDownloadItem_GetTotalBytes", 0),
		/*13*/ imports.NewTable("CefDownloadItem_GetUrl", 0),
		/*14*/ imports.NewTable("CefDownloadItem_IsCanceled", 0),
		/*15*/ imports.NewTable("CefDownloadItem_IsComplete", 0),
		/*16*/ imports.NewTable("CefDownloadItem_IsInProgress", 0),
		/*17*/ imports.NewTable("CefDownloadItem_IsInterrupted", 0),
		/*18*/ imports.NewTable("CefDownloadItem_IsValid", 0),
		/*19*/ imports.NewTable("CefDownloadItem_UnWrap", 0),
	}
)

func downloadItemImportAPI() *imports.Imports {
	if downloadItemImport == nil {
		downloadItemImport = NewDefaultImports()
		downloadItemImport.SetImportTable(downloadItemImportTables)
		downloadItemImportTables = nil
	}
	return downloadItemImport
}
