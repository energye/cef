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

// ICefExtension Parent: ICefBaseRefCounted
type ICefExtension interface {
	ICefBaseRefCounted
	GetIdentifier() string                // function
	GetPath() string                      // function
	GetManifest() ICefDictionaryValue     // function
	IsSame(that ICefExtension) bool       // function
	GetHandler() ICefExtensionHandler     // function
	GetLoaderContext() ICefRequestContext // function
	IsLoaded() bool                       // function
	GetBrowserActionPopup() string        // function
	GetBrowserActionIcon() string         // function
	GetPageActionPopup() string           // function
	GetPageActionIcon() string            // function
	GetOptionsPage() string               // function
	GetOptionsUIPage() string             // function
	GetBackgroundPage() string            // function
	GetURL() string                       // function
	Unload()                              // procedure
}

// TCefExtension Parent: TCefBaseRefCounted
type TCefExtension struct {
	TCefBaseRefCounted
}

// ExtensionRef -> ICefExtension
var ExtensionRef extension

// extension TCefExtension Ref
type extension uintptr

func (m *extension) UnWrap(data uintptr) ICefExtension {
	var resultCefExtension uintptr
	extensionImportAPI().SysCallN(15, uintptr(data), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TCefExtension) GetIdentifier() string {
	r1 := extensionImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPath() string {
	r1 := extensionImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetManifest() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	extensionImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefExtension) IsSame(that ICefExtension) bool {
	r1 := extensionImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefExtension) GetHandler() ICefExtensionHandler {
	var resultCefExtensionHandler uintptr
	extensionImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCefExtensionHandler)))
	return AsCefExtensionHandler(resultCefExtensionHandler)
}

func (m *TCefExtension) GetLoaderContext() ICefRequestContext {
	var resultCefRequestContext uintptr
	extensionImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TCefExtension) IsLoaded() bool {
	r1 := extensionImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TCefExtension) GetBrowserActionPopup() string {
	r1 := extensionImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetBrowserActionIcon() string {
	r1 := extensionImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPageActionPopup() string {
	r1 := extensionImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPageActionIcon() string {
	r1 := extensionImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetOptionsPage() string {
	r1 := extensionImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetOptionsUIPage() string {
	r1 := extensionImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetBackgroundPage() string {
	r1 := extensionImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetURL() string {
	r1 := extensionImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) Unload() {
	extensionImportAPI().SysCallN(16, m.Instance())
}

var (
	extensionImport       *imports.Imports = nil
	extensionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefExtension_GetBackgroundPage", 0),
		/*1*/ imports.NewTable("CefExtension_GetBrowserActionIcon", 0),
		/*2*/ imports.NewTable("CefExtension_GetBrowserActionPopup", 0),
		/*3*/ imports.NewTable("CefExtension_GetHandler", 0),
		/*4*/ imports.NewTable("CefExtension_GetIdentifier", 0),
		/*5*/ imports.NewTable("CefExtension_GetLoaderContext", 0),
		/*6*/ imports.NewTable("CefExtension_GetManifest", 0),
		/*7*/ imports.NewTable("CefExtension_GetOptionsPage", 0),
		/*8*/ imports.NewTable("CefExtension_GetOptionsUIPage", 0),
		/*9*/ imports.NewTable("CefExtension_GetPageActionIcon", 0),
		/*10*/ imports.NewTable("CefExtension_GetPageActionPopup", 0),
		/*11*/ imports.NewTable("CefExtension_GetPath", 0),
		/*12*/ imports.NewTable("CefExtension_GetURL", 0),
		/*13*/ imports.NewTable("CefExtension_IsLoaded", 0),
		/*14*/ imports.NewTable("CefExtension_IsSame", 0),
		/*15*/ imports.NewTable("CefExtension_UnWrap", 0),
		/*16*/ imports.NewTable("CefExtension_Unload", 0),
	}
)

func extensionImportAPI() *imports.Imports {
	if extensionImport == nil {
		extensionImport = NewDefaultImports()
		extensionImport.SetImportTable(extensionImportTables)
		extensionImportTables = nil
	}
	return extensionImport
}
