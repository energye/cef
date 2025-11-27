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

// ICefExtension Parent: ICefBaseRefCountedRef
type ICefExtension interface {
	ICefBaseRefCountedRef
	// GetIdentifier
	//  Returns the unique extension identifier. This is calculated based on the
	//  extension public key, if available, or on the extension path. See
	//  https://developer.chrome.com/extensions/manifest/key for details.
	GetIdentifier() string // function
	// GetPath
	//  Returns the absolute path to the extension directory on disk. This value
	//  will be prefixed with PK_DIR_RESOURCES if a relative path was passed to
	//  ICefRequestContext.LoadExtension.
	GetPath() string // function
	// GetManifest
	//  Returns the extension manifest contents as a ICefDictionaryValue
	//  object. See https://developer.chrome.com/extensions/manifest for details.
	GetManifest() ICefDictionaryValue // function
	// IsSame
	//  Returns true (1) if this object is the same extension as |that| object.
	//  Extensions are considered the same if identifier, path and loader context
	//  match.
	IsSame(that ICefExtension) bool // function
	// GetHandler
	//  Returns the handler for this extension. Will return NULL for internal
	//  extensions or if no handler was passed to
	//  ICefRequestContext.LoadExtension.
	GetHandler() IEngExtensionHandler // function
	// GetLoaderContext
	//  Returns the request context that loaded this extension. Will return NULL
	//  for internal extensions or if the extension has been unloaded. See the
	//  ICefRequestContext.LoadExtension documentation for more information
	//  about loader contexts. Must be called on the browser process UI thread.
	GetLoaderContext() ICefRequestContext // function
	// IsLoaded
	//  Returns true (1) if this extension is currently loaded. Must be called on
	//  the browser process UI thread.
	IsLoaded() bool                // function
	GetBrowserActionPopup() string // function
	GetBrowserActionIcon() string  // function
	GetPageActionPopup() string    // function
	GetPageActionIcon() string     // function
	GetOptionsPage() string        // function
	GetOptionsUIPage() string      // function
	GetBackgroundPage() string     // function
	GetURL() string                // function
	// Unload
	//  Unload this extension if it is not an internal extension and is currently
	//  loaded. Will result in a call to
	//  ICefExtensionHandler.OnExtensionUnloaded on success.
	Unload() // procedure
}

// ICefExtensionRef Parent: ICefExtension
type ICefExtensionRef interface {
	ICefExtension
	AsIntfExtension() uintptr
}

type TCefExtensionRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefExtensionRef) GetIdentifier() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetPath() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
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

func (m *TCefExtensionRef) GetBrowserActionPopup() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetBrowserActionIcon() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetPageActionPopup() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(10, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetPageActionIcon() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(11, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetOptionsPage() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetOptionsUIPage() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(13, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetBackgroundPage() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(14, m.Instance())
	return api.GoStr(r)
}

func (m *TCefExtensionRef) GetURL() string {
	if !m.IsValid() {
		return ""
	}
	r := cefExtensionRefAPI().SysCallN(15, m.Instance())
	return api.GoStr(r)
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
