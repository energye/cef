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
)

// ICefExtension Parent: ICefBaseRefCounted
type ICefExtension interface {
	ICefBaseRefCounted
	GetIdentifier() string                // function
	GetPath() string                      // function
	GetManifest() ICefDictionaryValue     // function
	IsSame(that ICefExtension) bool       // function
	GetHandler() IEngExtensionHandler     // function
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

// ICefExtensionRef Parent: ICefExtension ICefBaseRefCountedRef
type ICefExtensionRef interface {
	ICefExtension
	ICefBaseRefCountedRef
	AsIntfExtension() uintptr
}

type TCefExtensionRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefExtensionRef) GetIdentifier() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetPath() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetManifest() (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefExtensionRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefExtensionRef) IsSame(that ICefExtension) bool {
	if !m.IsValid() {
		return false
	}
	r := cefExtensionRefAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(that))
	return api.GoBool(r)
}

func (m *TCefExtensionRef) GetHandler() (result IEngExtensionHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefExtensionRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngExtensionHandler(resultPtr)
	return
}

func (m *TCefExtensionRef) GetLoaderContext() (result ICefRequestContext) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefExtensionRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (m *TCefExtensionRef) IsLoaded() bool {
	if !m.IsValid() {
		return false
	}
	r := cefExtensionRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefExtensionRef) GetBrowserActionPopup() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetBrowserActionIcon() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetPageActionPopup() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetPageActionIcon() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetOptionsPage() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetOptionsUIPage() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetBackgroundPage() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) GetURL() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefExtensionRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefExtensionRef) Unload() {
	if !m.IsValid() {
		return
	}
	cefExtensionRefAPI().SysCallN(17, m.Instance())
}

func (m *TCefExtensionRef) AsIntfExtension() uintptr {
	return m.GetIntfPointer(0)
}

// ExtensionRef  is static instance
var ExtensionRef _ExtensionRefClass

// _ExtensionRefClass is class type defined by TCefExtensionRef
type _ExtensionRefClass uintptr

func (_ExtensionRefClass) UnWrap(data uintptr) (result ICefExtension) {
	var resultPtr uintptr
	cefExtensionRefAPI().SysCallN(16, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefExtensionRef(resultPtr)
	return
}

// NewExtensionRef class constructor
func NewExtensionRef(data uintptr) ICefExtensionRef {
	var extensionPtr uintptr // ICefExtension
	r := cefExtensionRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&extensionPtr)))
	ret := AsCefExtensionRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, extensionPtr)
	}
	return ret
}

var (
	cefExtensionRefOnce   base.Once
	cefExtensionRefImport *imports.Imports = nil
)

func cefExtensionRefAPI() *imports.Imports {
	cefExtensionRefOnce.Do(func() {
		cefExtensionRefImport = api.NewDefaultImports()
		cefExtensionRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefExtensionRef_Create", 0), // constructor NewExtensionRef
			/* 1 */ imports.NewTable("TCefExtensionRef_GetIdentifier", 0), // function GetIdentifier
			/* 2 */ imports.NewTable("TCefExtensionRef_GetPath", 0), // function GetPath
			/* 3 */ imports.NewTable("TCefExtensionRef_GetManifest", 0), // function GetManifest
			/* 4 */ imports.NewTable("TCefExtensionRef_IsSame", 0), // function IsSame
			/* 5 */ imports.NewTable("TCefExtensionRef_GetHandler", 0), // function GetHandler
			/* 6 */ imports.NewTable("TCefExtensionRef_GetLoaderContext", 0), // function GetLoaderContext
			/* 7 */ imports.NewTable("TCefExtensionRef_IsLoaded", 0), // function IsLoaded
			/* 8 */ imports.NewTable("TCefExtensionRef_GetBrowserActionPopup", 0), // function GetBrowserActionPopup
			/* 9 */ imports.NewTable("TCefExtensionRef_GetBrowserActionIcon", 0), // function GetBrowserActionIcon
			/* 10 */ imports.NewTable("TCefExtensionRef_GetPageActionPopup", 0), // function GetPageActionPopup
			/* 11 */ imports.NewTable("TCefExtensionRef_GetPageActionIcon", 0), // function GetPageActionIcon
			/* 12 */ imports.NewTable("TCefExtensionRef_GetOptionsPage", 0), // function GetOptionsPage
			/* 13 */ imports.NewTable("TCefExtensionRef_GetOptionsUIPage", 0), // function GetOptionsUIPage
			/* 14 */ imports.NewTable("TCefExtensionRef_GetBackgroundPage", 0), // function GetBackgroundPage
			/* 15 */ imports.NewTable("TCefExtensionRef_GetURL", 0), // function GetURL
			/* 16 */ imports.NewTable("TCefExtensionRef_UnWrap", 0), // static function UnWrap
			/* 17 */ imports.NewTable("TCefExtensionRef_unload", 0), // procedure Unload
		}
	})
	return cefExtensionRefImport
}
