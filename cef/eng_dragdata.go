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

// ICefDragData Parent: ICefBaseRefCounted
//
//	Interface used to represent drag data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_data_capi.h">CEF source file: /include/capi/cef_drag_data_capi.h (cef_drag_data_t))</a>
type ICefDragData interface {
	ICefBaseRefCounted
	// Clone
	//  Returns a copy of the current object.
	Clone() ICefDragData // function
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// IsLink
	//  Returns true (1) if the drag data is a link.
	IsLink() bool // function
	// IsFragment
	//  Returns true (1) if the drag data is a text or html fragment.
	IsFragment() bool // function
	// IsFile
	//  Returns true (1) if the drag data is a file.
	IsFile() bool // function
	// GetLinkUrl
	//  Return the link URL that is being dragged.
	GetLinkUrl() string // function
	// GetLinkTitle
	//  Return the title associated with the link being dragged.
	GetLinkTitle() string // function
	// GetLinkMetadata
	//  Return the metadata, if any, associated with the link being dragged.
	GetLinkMetadata() string // function
	// GetFragmentText
	//  Return the plain text fragment that is being dragged.
	GetFragmentText() string // function
	// GetFragmentHtml
	//  Return the text/html fragment that is being dragged.
	GetFragmentHtml() string // function
	// GetFragmentBaseUrl
	//  Return the base URL that the fragment came from. This value is used for resolving relative URLs and may be NULL.
	GetFragmentBaseUrl() string // function
	// GetFileName
	//  Return the name of the file being dragged out of the browser window.
	GetFileName() string // function
	// GetFileContents
	//  Write the contents of the file being dragged out of the web view into |writer|. Returns the number of bytes sent to |writer|. If |writer| is NULL this function will return the size of the file contents in bytes. Call get_file_name() to get a suggested name for the file.
	GetFileContents(writer ICefStreamWriter) NativeUInt // function
	// GetFileNames
	//  Retrieve the list of file names that are being dragged into the browser window.
	GetFileNames(names *IStrings) int32 // function
	// GetFilePaths
	//  Retrieve the list of file paths that are being dragged into the browser window.
	GetFilePaths(paths *IStrings) int32 // function
	// GetImage
	//  Get the image representation of drag data. May return NULL if no image representation is available.
	GetImage() ICefImage // function
	// GetImageHotspot
	//  Get the image hotspot (drag start location relative to image dimensions).
	GetImageHotspot() (resultCefPoint TCefPoint) // function
	// HasImage
	//  Returns true (1) if an image representation of drag data is available.
	HasImage() bool // function
	// SetLinkUrl
	//  Set the link URL that is being dragged.
	SetLinkUrl(url string) // procedure
	// SetLinkTitle
	//  Set the title associated with the link being dragged.
	SetLinkTitle(title string) // procedure
	// SetLinkMetadata
	//  Set the metadata associated with the link being dragged.
	SetLinkMetadata(data string) // procedure
	// SetFragmentText
	//  Set the plain text fragment that is being dragged.
	SetFragmentText(text string) // procedure
	// SetFragmentHtml
	//  Set the text/html fragment that is being dragged.
	SetFragmentHtml(html string) // procedure
	// SetFragmentBaseUrl
	//  Set the base URL that the fragment came from.
	SetFragmentBaseUrl(baseUrl string) // procedure
	// ResetFileContents
	//  Reset the file contents. You should do this before calling ICefBrowserHost.DragTargetDragEnter as the web view does not allow us to drag in this kind of data.
	ResetFileContents() // procedure
	// AddFile
	//  Add a file that is being dragged into the webview.
	AddFile(path, displayName string) // procedure
	// ClearFilenames
	//  Clear list of filenames.
	ClearFilenames() // procedure
}

// TCefDragData Parent: TCefBaseRefCounted
//
//	Interface used to represent drag data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_data_capi.h">CEF source file: /include/capi/cef_drag_data_capi.h (cef_drag_data_t))</a>
type TCefDragData struct {
	TCefBaseRefCounted
}

// DragDataRef -> ICefDragData
var DragDataRef dragData

// dragData TCefDragData Ref
type dragData uintptr

func (m *dragData) UnWrap(data uintptr) ICefDragData {
	var resultCefDragData uintptr
	dragDataImportAPI().SysCallN(28, uintptr(data), uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *dragData) New() ICefDragData {
	var resultCefDragData uintptr
	dragDataImportAPI().SysCallN(20, uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *TCefDragData) Clone() ICefDragData {
	var resultCefDragData uintptr
	dragDataImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *TCefDragData) IsReadOnly() bool {
	r1 := dragDataImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsLink() bool {
	r1 := dragDataImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsFragment() bool {
	r1 := dragDataImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsFile() bool {
	r1 := dragDataImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) GetLinkUrl() string {
	r1 := dragDataImportAPI().SysCallN(14, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetLinkTitle() string {
	r1 := dragDataImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetLinkMetadata() string {
	r1 := dragDataImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentText() string {
	r1 := dragDataImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentHtml() string {
	r1 := dragDataImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentBaseUrl() string {
	r1 := dragDataImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFileName() string {
	r1 := dragDataImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFileContents(writer ICefStreamWriter) NativeUInt {
	r1 := dragDataImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(writer))
	return NativeUInt(r1)
}

func (m *TCefDragData) GetFileNames(names *IStrings) int32 {
	var result0 uintptr
	r1 := dragDataImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&result0)))
	*names = AsStrings(result0)
	return int32(r1)
}

func (m *TCefDragData) GetFilePaths(paths *IStrings) int32 {
	var result0 uintptr
	r1 := dragDataImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&result0)))
	*paths = AsStrings(result0)
	return int32(r1)
}

func (m *TCefDragData) GetImage() ICefImage {
	var resultCefImage uintptr
	dragDataImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefDragData) GetImageHotspot() (resultCefPoint TCefPoint) {
	dragDataImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCefDragData) HasImage() bool {
	r1 := dragDataImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) SetLinkUrl(url string) {
	dragDataImportAPI().SysCallN(27, m.Instance(), PascalStr(url))
}

func (m *TCefDragData) SetLinkTitle(title string) {
	dragDataImportAPI().SysCallN(26, m.Instance(), PascalStr(title))
}

func (m *TCefDragData) SetLinkMetadata(data string) {
	dragDataImportAPI().SysCallN(25, m.Instance(), PascalStr(data))
}

func (m *TCefDragData) SetFragmentText(text string) {
	dragDataImportAPI().SysCallN(24, m.Instance(), PascalStr(text))
}

func (m *TCefDragData) SetFragmentHtml(html string) {
	dragDataImportAPI().SysCallN(23, m.Instance(), PascalStr(html))
}

func (m *TCefDragData) SetFragmentBaseUrl(baseUrl string) {
	dragDataImportAPI().SysCallN(22, m.Instance(), PascalStr(baseUrl))
}

func (m *TCefDragData) ResetFileContents() {
	dragDataImportAPI().SysCallN(21, m.Instance())
}

func (m *TCefDragData) AddFile(path, displayName string) {
	dragDataImportAPI().SysCallN(0, m.Instance(), PascalStr(path), PascalStr(displayName))
}

func (m *TCefDragData) ClearFilenames() {
	dragDataImportAPI().SysCallN(1, m.Instance())
}

var (
	dragDataImport       *imports.Imports = nil
	dragDataImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefDragData_AddFile", 0),
		/*1*/ imports.NewTable("CefDragData_ClearFilenames", 0),
		/*2*/ imports.NewTable("CefDragData_Clone", 0),
		/*3*/ imports.NewTable("CefDragData_GetFileContents", 0),
		/*4*/ imports.NewTable("CefDragData_GetFileName", 0),
		/*5*/ imports.NewTable("CefDragData_GetFileNames", 0),
		/*6*/ imports.NewTable("CefDragData_GetFilePaths", 0),
		/*7*/ imports.NewTable("CefDragData_GetFragmentBaseUrl", 0),
		/*8*/ imports.NewTable("CefDragData_GetFragmentHtml", 0),
		/*9*/ imports.NewTable("CefDragData_GetFragmentText", 0),
		/*10*/ imports.NewTable("CefDragData_GetImage", 0),
		/*11*/ imports.NewTable("CefDragData_GetImageHotspot", 0),
		/*12*/ imports.NewTable("CefDragData_GetLinkMetadata", 0),
		/*13*/ imports.NewTable("CefDragData_GetLinkTitle", 0),
		/*14*/ imports.NewTable("CefDragData_GetLinkUrl", 0),
		/*15*/ imports.NewTable("CefDragData_HasImage", 0),
		/*16*/ imports.NewTable("CefDragData_IsFile", 0),
		/*17*/ imports.NewTable("CefDragData_IsFragment", 0),
		/*18*/ imports.NewTable("CefDragData_IsLink", 0),
		/*19*/ imports.NewTable("CefDragData_IsReadOnly", 0),
		/*20*/ imports.NewTable("CefDragData_New", 0),
		/*21*/ imports.NewTable("CefDragData_ResetFileContents", 0),
		/*22*/ imports.NewTable("CefDragData_SetFragmentBaseUrl", 0),
		/*23*/ imports.NewTable("CefDragData_SetFragmentHtml", 0),
		/*24*/ imports.NewTable("CefDragData_SetFragmentText", 0),
		/*25*/ imports.NewTable("CefDragData_SetLinkMetadata", 0),
		/*26*/ imports.NewTable("CefDragData_SetLinkTitle", 0),
		/*27*/ imports.NewTable("CefDragData_SetLinkUrl", 0),
		/*28*/ imports.NewTable("CefDragData_UnWrap", 0),
	}
)

func dragDataImportAPI() *imports.Imports {
	if dragDataImport == nil {
		dragDataImport = NewDefaultImports()
		dragDataImport.SetImportTable(dragDataImportTables)
		dragDataImportTables = nil
	}
	return dragDataImport
}
