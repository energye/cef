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
)

// ICefRequestContext Parent: ICefPreferenceManager
type ICefRequestContext interface {
	ICefPreferenceManager
	IsSame(other ICefRequestContext) bool                                                                     // function
	IsSharingWith(other ICefRequestContext) bool                                                              // function
	IsGlobal() bool                                                                                           // function
	GetHandler() IEngRequestContextHandler                                                                    // function
	GetCachePath() string                                                                                     // function
	GetCookieManager(callback IEngCompletionCallback) ICefCookieManager                                       // function
	RegisterSchemeHandlerFactory(schemeName string, domainName string, factory IEngSchemeHandlerFactory) bool // function
	ClearSchemeHandlerFactories() bool                                                                        // function
	DidLoadExtension(extensionId string) bool                                                                 // function
	HasExtension(extensionId string) bool                                                                     // function
	GetExtensions(extensionIds lcl.IStringList) bool                                                          // function
	GetExtension(extensionId string) ICefExtension                                                            // function
	GetMediaRouter(callback IEngCompletionCallback) ICefMediaRouter                                           // function
	ClearCertificateExceptions(callback IEngCompletionCallback)                                               // procedure
	ClearHttpAuthCredentials(callback IEngCompletionCallback)                                                 // procedure
	CloseAllConnections(callback IEngCompletionCallback)                                                      // procedure
	ResolveHost(origin string, callback IEngResolveCallback)                                                  // procedure
	LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler)           // procedure
}

// ICefRequestContextRef Parent: ICefRequestContext ICefPreferenceManagerRef
type ICefRequestContextRef interface {
	ICefRequestContext
	ICefPreferenceManagerRef
	AsIntfRequestContext() uintptr
	AsIntfPreferenceManager() uintptr
}

type TCefRequestContextRef struct {
	TCefPreferenceManagerRef
}

func (m *TCefRequestContextRef) IsSame(other ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(other))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) IsSharingWith(other ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(other))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) IsGlobal() bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) GetHandler() (result IEngRequestContextHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngRequestContextHandler(resultPtr)
	return
}

func (m *TCefRequestContextRef) GetCachePath() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefRequestContextRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefRequestContextRef) GetCookieManager(callback IEngCompletionCallback) (result ICefCookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(callback), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefCookieManagerRef(resultPtr)
	return
}

func (m *TCefRequestContextRef) RegisterSchemeHandlerFactory(schemeName string, domainName string, factory IEngSchemeHandlerFactory) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(7, m.Instance(), api.PasStr(schemeName), api.PasStr(domainName), base.GetObjectUintptr(factory))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) ClearSchemeHandlerFactories() bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) DidLoadExtension(extensionId string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(9, m.Instance(), api.PasStr(extensionId))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) HasExtension(extensionId string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(10, m.Instance(), api.PasStr(extensionId))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) GetExtensions(extensionIds lcl.IStringList) bool {
	if !m.IsValid() {
		return false
	}
	r := cefRequestContextRefAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(extensionIds))
	return api.GoBool(r)
}

func (m *TCefRequestContextRef) GetExtension(extensionId string) (result ICefExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(12, m.Instance(), api.PasStr(extensionId), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefExtensionRef(resultPtr)
	return
}

func (m *TCefRequestContextRef) GetMediaRouter(callback IEngCompletionCallback) (result ICefMediaRouter) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(callback), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouterRef(resultPtr)
	return
}

func (m *TCefRequestContextRef) ClearCertificateExceptions(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) ClearHttpAuthCredentials(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) CloseAllConnections(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) ResolveHost(origin string, callback IEngResolveCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(22, m.Instance(), api.PasStr(origin), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(23, m.Instance(), api.PasStr(rootDirectory), base.GetObjectUintptr(manifest), base.GetObjectUintptr(handler))
}

func (m *TCefRequestContextRef) AsIntfRequestContext() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCefRequestContextRef) AsIntfPreferenceManager() uintptr {
	return m.GetIntfPointer(1)
}

// RequestContextRef  is static instance
var RequestContextRef _RequestContextRefClass

// _RequestContextRefClass is class type defined by TCefRequestContextRef
type _RequestContextRefClass uintptr

func (_RequestContextRefClass) UnWrapWithPointer(data uintptr) (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(14, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (_RequestContextRefClass) GlobalToRequestContext() (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(15, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (_RequestContextRefClass) NewWithPCefRequestContextSettingsRequestContextHandler(settings TCefRequestContextSettings, handler IEngRequestContextHandler) (result ICefRequestContext) {
	settingsPtr := settings.ToPas()
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(16, uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (_RequestContextRefClass) NewWithStrX3BoolX3RCHandler(cache string, acceptLanguageList string, cookieableSchemesList string, cookieableSchemesExcludeDefaults bool, persistSessionCookies bool, persistUserPreferences bool, handler IEngRequestContextHandler) (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(17, api.PasStr(cache), api.PasStr(acceptLanguageList), api.PasStr(cookieableSchemesList), api.PasBool(cookieableSchemesExcludeDefaults), api.PasBool(persistSessionCookies), api.PasBool(persistUserPreferences), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (_RequestContextRefClass) Shared(other ICefRequestContext, handler IEngRequestContextHandler) (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(18, base.GetObjectUintptr(other), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

// NewRequestContextRef class constructor
func NewRequestContextRef(data uintptr) ICefRequestContextRef {
	var requestContextPtr uintptr    // ICefRequestContext
	var preferenceManagerPtr uintptr // ICefPreferenceManager
	r := cefRequestContextRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&requestContextPtr)), uintptr(base.UnsafePointer(&preferenceManagerPtr)))
	ret := AsCefRequestContextRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, requestContextPtr)
		intf.SetIntfPointer(1, preferenceManagerPtr)
	}
	return ret
}

var (
	cefRequestContextRefOnce   base.Once
	cefRequestContextRefImport *imports.Imports = nil
)

func cefRequestContextRefAPI() *imports.Imports {
	cefRequestContextRefOnce.Do(func() {
		cefRequestContextRefImport = api.NewDefaultImports()
		cefRequestContextRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefRequestContextRef_Create", 0), // constructor NewRequestContextRef
			/* 1 */ imports.NewTable("TCefRequestContextRef_IsSame", 0), // function IsSame
			/* 2 */ imports.NewTable("TCefRequestContextRef_IsSharingWith", 0), // function IsSharingWith
			/* 3 */ imports.NewTable("TCefRequestContextRef_IsGlobal", 0), // function IsGlobal
			/* 4 */ imports.NewTable("TCefRequestContextRef_GetHandler", 0), // function GetHandler
			/* 5 */ imports.NewTable("TCefRequestContextRef_GetCachePath", 0), // function GetCachePath
			/* 6 */ imports.NewTable("TCefRequestContextRef_GetCookieManager", 0), // function GetCookieManager
			/* 7 */ imports.NewTable("TCefRequestContextRef_RegisterSchemeHandlerFactory", 0), // function RegisterSchemeHandlerFactory
			/* 8 */ imports.NewTable("TCefRequestContextRef_ClearSchemeHandlerFactories", 0), // function ClearSchemeHandlerFactories
			/* 9 */ imports.NewTable("TCefRequestContextRef_DidLoadExtension", 0), // function DidLoadExtension
			/* 10 */ imports.NewTable("TCefRequestContextRef_HasExtension", 0), // function HasExtension
			/* 11 */ imports.NewTable("TCefRequestContextRef_GetExtensions", 0), // function GetExtensions
			/* 12 */ imports.NewTable("TCefRequestContextRef_GetExtension", 0), // function GetExtension
			/* 13 */ imports.NewTable("TCefRequestContextRef_GetMediaRouter", 0), // function GetMediaRouter
			/* 14 */ imports.NewTable("TCefRequestContextRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 15 */ imports.NewTable("TCefRequestContextRef_GlobalToRequestContext", 0), // static function GlobalToRequestContext
			/* 16 */ imports.NewTable("TCefRequestContextRef_NewWithPCefRequestContextSettingsRequestContextHandler", 0), // static function NewWithPCefRequestContextSettingsRequestContextHandler
			/* 17 */ imports.NewTable("TCefRequestContextRef_NewWithStrX3BoolX3RCHandler", 0), // static function NewWithStrX3BoolX3RCHandler
			/* 18 */ imports.NewTable("TCefRequestContextRef_Shared", 0), // static function Shared
			/* 19 */ imports.NewTable("TCefRequestContextRef_ClearCertificateExceptions", 0), // procedure ClearCertificateExceptions
			/* 20 */ imports.NewTable("TCefRequestContextRef_ClearHttpAuthCredentials", 0), // procedure ClearHttpAuthCredentials
			/* 21 */ imports.NewTable("TCefRequestContextRef_CloseAllConnections", 0), // procedure CloseAllConnections
			/* 22 */ imports.NewTable("TCefRequestContextRef_ResolveHost", 0), // procedure ResolveHost
			/* 23 */ imports.NewTable("TCefRequestContextRef_LoadExtension", 0), // procedure LoadExtension
		}
	})
	return cefRequestContextRefImport
}
