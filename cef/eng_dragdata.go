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
type ICefDragData interface {
	ICefBaseRefCounted
	Clone() ICefDragData                                // function
	IsReadOnly() bool                                   // function
	IsLink() bool                                       // function
	IsFragment() bool                                   // function
	IsFile() bool                                       // function
	GetLinkUrl() string                                 // function
	GetLinkTitle() string                               // function
	GetLinkMetadata() string                            // function
	GetFragmentText() string                            // function
	GetFragmentHtml() string                            // function
	GetFragmentBaseUrl() string                         // function
	GetFileName() string                                // function
	GetFileContents(writer ICefStreamWriter) NativeUInt // function
	GetFileNames(names *IStrings) int32                 // function
	GetFilePaths(paths *IStrings) int32                 // function
	GetImage() ICefImage                                // function
	GetImageHotspot() (resultCefPoint TCefPoint)        // function
	HasImage() bool                                     // function
	SetLinkUrl(url string)                              // procedure
	SetLinkTitle(title string)                          // procedure
	SetLinkMetadata(data string)                        // procedure
	SetFragmentText(text string)                        // procedure
	SetFragmentHtml(html string)                        // procedure
	SetFragmentBaseUrl(baseUrl string)                  // procedure
	ResetFileContents()                                 // procedure
	AddFile(path, displayName string)                   // procedure
	ClearFilenames()                                    // procedure
}

// TCefDragData Parent: TCefBaseRefCounted
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
