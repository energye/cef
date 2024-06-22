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
type ICefDownloadItem interface {
	ICefBaseRefCounted
	IsValid() bool                                   // function
	IsInProgress() bool                              // function
	IsComplete() bool                                // function
	IsCanceled() bool                                // function
	IsInterrupted() bool                             // function
	GetInterruptReason() TCefDownloadInterruptReason // function
	GetCurrentSpeed() (resultInt64 int64)            // function
	GetPercentComplete() int32                       // function
	GetTotalBytes() (resultInt64 int64)              // function
	GetReceivedBytes() (resultInt64 int64)           // function
	GetStartTime() (resultDateTime TDateTime)        // function
	GetEndTime() (resultDateTime TDateTime)          // function
	GetFullPath() string                             // function
	GetId() uint32                                   // function
	GetUrl() string                                  // function
	GetOriginalUrl() string                          // function
	GetSuggestedFileName() string                    // function
	GetContentDisposition() string                   // function
	GetMimeType() string                             // function
}

// TCefDownloadItem Parent: TCefBaseRefCounted
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
