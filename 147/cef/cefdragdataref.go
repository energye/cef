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
	"github.com/energye/lcl/lcl"

	cefTypes "github.com/energye/cef/147/types"
)

// ICefDragData Parent: ICefBaseRefCounted
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
	//  Return the base URL that the fragment came from. This value is used for
	//  resolving relative URLs and may be NULL.
	GetFragmentBaseUrl() string // function
	// GetFileName
	//  Return the name of the file being dragged out of the browser window.
	GetFileName() string // function
	// GetFileContents
	//  Write the contents of the file being dragged out of the web view into
	//  |writer|. Returns the number of bytes sent to |writer|. If |writer| is
	//  NULL this function will return the size of the file contents in bytes.
	//  Call get_file_name() to get a suggested name for the file.
	GetFileContents(writer ICefStreamWriter) cefTypes.NativeUInt // function
	// GetFileNames
	//  Retrieve the list of file names that are being dragged into the browser
	//  window.
	GetFileNames(names *lcl.IStrings) int32 // function
	// GetFilePaths
	//  Retrieve the list of file paths that are being dragged into the browser
	//  window.
	GetFilePaths(paths *lcl.IStrings) int32 // function
	// GetImage
	//  Get the image representation of drag data. May return NULL if no image
	//  representation is available.
	GetImage() ICefImage // function
	// GetImageHotspot
	//  Get the image hotspot (drag start location relative to image dimensions).
	GetImageHotspot() TCefPoint // function
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
	//  Reset the file contents. You should do this before calling
	//  ICefBrowserHost.DragTargetDragEnter as the web view does not allow us
	//  to drag in this kind of data.
	ResetFileContents() // procedure
	// AddFile
	//  Add a file that is being dragged into the webview.
	AddFile(path string, displayName string) // procedure
	// ClearFilenames
	//  Clear list of filenames.
	ClearFilenames() // procedure
}

// ICefDragDataRef Parent: ICefDragData ICefBaseRefCountedRef
type ICefDragDataRef interface {
	ICefDragData
	ICefBaseRefCountedRef
	AsIntfDragData() uintptr
}

type TCefDragDataRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefDragDataRef) Clone() (result ICefDragData) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDragDataRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDragDataRef(resultPtr)
	return
}

func (m *TCefDragDataRef) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDragDataRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDragDataRef) IsLink() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDragDataRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDragDataRef) IsFragment() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDragDataRefAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDragDataRef) IsFile() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDragDataRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDragDataRef) GetLinkUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetLinkTitle() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetLinkMetadata() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetFragmentText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetFragmentHtml() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetFragmentBaseUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetFileName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefDragDataRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefDragDataRef) GetFileContents(writer ICefStreamWriter) cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefDragDataRefAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(writer))
	return cefTypes.NativeUInt(r)
}

func (m *TCefDragDataRef) GetFileNames(names *lcl.IStrings) int32 {
	if !m.IsValid() {
		return 0
	}
	namesPtr := base.GetObjectUintptr(*names)
	r := cefDragDataRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&namesPtr)))
	*names = lcl.AsStrings(namesPtr)
	return int32(r)
}

func (m *TCefDragDataRef) GetFilePaths(paths *lcl.IStrings) int32 {
	if !m.IsValid() {
		return 0
	}
	pathsPtr := base.GetObjectUintptr(*paths)
	r := cefDragDataRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&pathsPtr)))
	*paths = lcl.AsStrings(pathsPtr)
	return int32(r)
}

func (m *TCefDragDataRef) GetImage() (result ICefImage) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefDragDataRefAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefImageRef(resultPtr)
	return
}

func (m *TCefDragDataRef) GetImageHotspot() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefDragDataRef) HasImage() bool {
	if !m.IsValid() {
		return false
	}
	r := cefDragDataRefAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TCefDragDataRef) SetLinkUrl(url string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(21, m.Instance(), api.PasStr(url))
}

func (m *TCefDragDataRef) SetLinkTitle(title string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(22, m.Instance(), api.PasStr(title))
}

func (m *TCefDragDataRef) SetLinkMetadata(data string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(23, m.Instance(), api.PasStr(data))
}

func (m *TCefDragDataRef) SetFragmentText(text string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(24, m.Instance(), api.PasStr(text))
}

func (m *TCefDragDataRef) SetFragmentHtml(html string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(25, m.Instance(), api.PasStr(html))
}

func (m *TCefDragDataRef) SetFragmentBaseUrl(baseUrl string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(26, m.Instance(), api.PasStr(baseUrl))
}

func (m *TCefDragDataRef) ResetFileContents() {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(27, m.Instance())
}

func (m *TCefDragDataRef) AddFile(path string, displayName string) {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(28, m.Instance(), api.PasStr(path), api.PasStr(displayName))
}

func (m *TCefDragDataRef) ClearFilenames() {
	if !m.IsValid() {
		return
	}
	cefDragDataRefAPI().SysCallN(29, m.Instance())
}

func (m *TCefDragDataRef) AsIntfDragData() uintptr {
	return m.GetIntfPointer(0)
}

// DragDataRef  is static instance
var DragDataRef _DragDataRefClass

// _DragDataRefClass is class type defined by TCefDragDataRef
type _DragDataRefClass uintptr

func (_DragDataRefClass) UnWrap(data uintptr) (result ICefDragData) {
	var resultPtr uintptr
	cefDragDataRefAPI().SysCallN(19, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDragDataRef(resultPtr)
	return
}

func (_DragDataRefClass) New() (result ICefDragData) {
	var resultPtr uintptr
	cefDragDataRefAPI().SysCallN(20, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDragDataRef(resultPtr)
	return
}

// NewDragDataRef class constructor
func NewDragDataRef(data uintptr) ICefDragDataRef {
	var dragDataPtr uintptr // ICefDragData
	r := cefDragDataRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&dragDataPtr)))
	ret := AsCefDragDataRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, dragDataPtr)
	}
	return ret
}

var (
	cefDragDataRefOnce   base.Once
	cefDragDataRefImport *imports.Imports = nil
)

func cefDragDataRefAPI() *imports.Imports {
	cefDragDataRefOnce.Do(func() {
		cefDragDataRefImport = api.NewDefaultImports()
		cefDragDataRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefDragDataRef_Create", 0), // constructor NewDragDataRef
			/* 1 */ imports.NewTable("TCefDragDataRef_Clone", 0), // function Clone
			/* 2 */ imports.NewTable("TCefDragDataRef_IsReadOnly", 0), // function IsReadOnly
			/* 3 */ imports.NewTable("TCefDragDataRef_IsLink", 0), // function IsLink
			/* 4 */ imports.NewTable("TCefDragDataRef_IsFragment", 0), // function IsFragment
			/* 5 */ imports.NewTable("TCefDragDataRef_IsFile", 0), // function IsFile
			/* 6 */ imports.NewTable("TCefDragDataRef_GetLinkUrl", 0), // function GetLinkUrl
			/* 7 */ imports.NewTable("TCefDragDataRef_GetLinkTitle", 0), // function GetLinkTitle
			/* 8 */ imports.NewTable("TCefDragDataRef_GetLinkMetadata", 0), // function GetLinkMetadata
			/* 9 */ imports.NewTable("TCefDragDataRef_GetFragmentText", 0), // function GetFragmentText
			/* 10 */ imports.NewTable("TCefDragDataRef_GetFragmentHtml", 0), // function GetFragmentHtml
			/* 11 */ imports.NewTable("TCefDragDataRef_GetFragmentBaseUrl", 0), // function GetFragmentBaseUrl
			/* 12 */ imports.NewTable("TCefDragDataRef_GetFileName", 0), // function GetFileName
			/* 13 */ imports.NewTable("TCefDragDataRef_GetFileContents", 0), // function GetFileContents
			/* 14 */ imports.NewTable("TCefDragDataRef_GetFileNames", 0), // function GetFileNames
			/* 15 */ imports.NewTable("TCefDragDataRef_GetFilePaths", 0), // function GetFilePaths
			/* 16 */ imports.NewTable("TCefDragDataRef_GetImage", 0), // function GetImage
			/* 17 */ imports.NewTable("TCefDragDataRef_GetImageHotspot", 0), // function GetImageHotspot
			/* 18 */ imports.NewTable("TCefDragDataRef_HasImage", 0), // function HasImage
			/* 19 */ imports.NewTable("TCefDragDataRef_UnWrap", 0), // static function UnWrap
			/* 20 */ imports.NewTable("TCefDragDataRef_New", 0), // static function New
			/* 21 */ imports.NewTable("TCefDragDataRef_SetLinkUrl", 0), // procedure SetLinkUrl
			/* 22 */ imports.NewTable("TCefDragDataRef_SetLinkTitle", 0), // procedure SetLinkTitle
			/* 23 */ imports.NewTable("TCefDragDataRef_SetLinkMetadata", 0), // procedure SetLinkMetadata
			/* 24 */ imports.NewTable("TCefDragDataRef_SetFragmentText", 0), // procedure SetFragmentText
			/* 25 */ imports.NewTable("TCefDragDataRef_SetFragmentHtml", 0), // procedure SetFragmentHtml
			/* 26 */ imports.NewTable("TCefDragDataRef_SetFragmentBaseUrl", 0), // procedure SetFragmentBaseUrl
			/* 27 */ imports.NewTable("TCefDragDataRef_ResetFileContents", 0), // procedure ResetFileContents
			/* 28 */ imports.NewTable("TCefDragDataRef_AddFile", 0), // procedure AddFile
			/* 29 */ imports.NewTable("TCefDragDataRef_ClearFilenames", 0), // procedure ClearFilenames
		}
	})
	return cefDragDataRefImport
}
