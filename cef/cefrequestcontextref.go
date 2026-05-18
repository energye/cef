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

	cefTypes "github.com/energye/cef/types"
)

// ICefRequestContext Parent: ICefPreferenceManager
type ICefRequestContext interface {
	ICefPreferenceManager
	// IsSame
	//  Returns true (1) if this object is pointing to the same context as |that|
	//  object.
	IsSame(other ICefRequestContext) bool // function
	// IsSharingWith
	//  Returns true (1) if this object is sharing the same storage as |that|
	//  object.
	IsSharingWith(other ICefRequestContext) bool // function
	// IsGlobal
	//  Returns true (1) if this object is the global context. The global context
	//  is used by default when creating a browser or URL request with a NULL
	//  context argument.
	IsGlobal() bool // function
	// GetHandler
	//  Returns the handler for this context if any.
	GetHandler() IEngRequestContextHandler // function
	// GetCachePath
	//  Returns the cache path for this object. If NULL an "incognito mode" in-
	//  memory cache is being used.
	GetCachePath() string // function
	// GetCookieManager
	//  Returns the cookie manager for this object. If |callback| is non-NULL it
	//  will be executed asnychronously on the UI thread after the manager's
	//  storage has been initialized.
	GetCookieManager(callback IEngCompletionCallback) ICefCookieManager // function
	// RegisterSchemeHandlerFactory
	//  Register a scheme handler factory for the specified |scheme_name| and
	//  optional |domain_name|. An NULL |domain_name| value for a standard scheme
	//  will cause the factory to match all domain names. The |domain_name| value
	//  will be ignored for non-standard schemes. If |scheme_name| is a built-in
	//  scheme and no handler is returned by |factory| then the built-in scheme
	//  handler factory will be called. If |scheme_name| is a custom scheme then
	//  you must also implement the ICefApp.OnRegisterCustomSchemes()
	//  function in all processes. This function may be called multiple times to
	//  change or remove the factory that matches the specified |scheme_name| and
	//  optional |domain_name|. Returns false (0) if an error occurs. This
	//  function may be called on any thread in the browser process.
	RegisterSchemeHandlerFactory(schemeName string, domainName string, factory IEngSchemeHandlerFactory) bool // function
	// ClearSchemeHandlerFactories
	//  Clear all registered scheme handler factories. Returns false (0) on error.
	//  This function may be called on any thread in the browser process.
	ClearSchemeHandlerFactories() bool // function
	// DidLoadExtension
	//  Returns true (1) if this context was used to load the extension identified
	//  by |extension_id|. Other contexts sharing the same storage will also have
	//  access to the extension (see HasExtension). This function must be called
	//  on the browser process UI thread.
	//
	//  WARNING: This function is deprecated and will be removed in ~M127.
	DidLoadExtension(extensionId string) bool // function
	// HasExtension
	//  Returns true (1) if this context has access to the extension identified by
	//  |extension_id|. This may not be the context that was used to load the
	//  extension (see DidLoadExtension). This function must be called on the
	//  browser process UI thread.
	//
	//  WARNING: This function is deprecated and will be removed in ~M127.
	HasExtension(extensionId string) bool // function
	// GetExtensions
	//  Retrieve the list of all extensions that this context has access to (see
	//  HasExtension). |extension_ids| will be populated with the list of
	//  extension ID values. Returns true (1) on success. This function must be
	//  called on the browser process UI thread.
	//
	//  WARNING: This function is deprecated and will be removed in ~M127.
	GetExtensions(extensionIds lcl.IStringList) bool // function
	// GetExtension
	//  Returns the extension matching |extension_id| or NULL if no matching
	//  extension is accessible in this context (see HasExtension). This function
	//  must be called on the browser process UI thread.
	//
	//  WARNING: This function is deprecated and will be removed in ~M127.
	GetExtension(extensionId string) ICefExtension // function
	// GetMediaRouter
	//  Returns the MediaRouter object associated with this context. If
	//  |callback| is non-NULL it will be executed asnychronously on the UI thread
	//  after the manager's context has been initialized.
	GetMediaRouter(callback IEngCompletionCallback) ICefMediaRouter // function
	// GetWebsiteSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns nullptr if no value is configured. Must be called on the browser
	//  process UI thread.
	GetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) ICefValue // function
	// GetContentSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns CEF_CONTENT_SETTING_VALUE_DEFAULT if no value is configured. Must
	//  be called on the browser process UI thread.
	GetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) cefTypes.TCefContentSettingValues // function
	// GetChromeColorSchemeMode
	//  Returns the current Chrome color scheme mode (SYSTEM, LIGHT or DARK). Must
	//  be called on the browser process UI thread.
	GetChromeColorSchemeMode() cefTypes.TCefColorVariant // function
	// GetChromeColorSchemeColor
	//  Returns the current Chrome color scheme color, or transparent (0) for the
	//  default color. Must be called on the browser process UI thread.
	GetChromeColorSchemeColor() cefTypes.TCefColor // function
	// GetChromeColorSchemeVariant
	//  Returns the current Chrome color scheme variant. Must be called on the
	//  browser process UI thread.
	GetChromeColorSchemeVariant() cefTypes.TCefColorVariant // function
	// ClearCertificateExceptions
	//  Clears all certificate exceptions that were added as part of handling
	//  ICefRequestHandler.OnCertificateError(). If you call this it is
	//  recommended that you also call CloseAllConnections() or you risk not
	//  being prompted again for server certificates if you reconnect quickly. If
	//  |callback| is non-NULL it will be executed on the UI thread after
	//  completion.
	ClearCertificateExceptions(callback IEngCompletionCallback) // procedure
	// ClearHttpAuthCredentials
	//  Clears all HTTP authentication credentials that were added as part of
	//  handling GetAuthCredentials. If |callback| is non-NULL it will be executed
	//  on the UI thread after completion.
	ClearHttpAuthCredentials(callback IEngCompletionCallback) // procedure
	// CloseAllConnections
	//  Clears all active and idle connections that Chromium currently has. This
	//  is only recommended if you have released all other CEF objects but don't
	//  yet want to call cef_shutdown(). If |callback| is non-NULL it will be
	//  executed on the UI thread after completion.
	CloseAllConnections(callback IEngCompletionCallback) // procedure
	// ResolveHost
	//  Attempts to resolve |origin| to a list of associated IP addresses.
	//  |callback| will be executed on the UI thread after completion.
	ResolveHost(origin string, callback IEngResolveCallback) // procedure
	// LoadExtension
	//  Load an extension.
	//
	//  If extension resources will be read from disk using the default load
	//  implementation then |root_directory| should be the absolute path to the
	//  extension resources directory and |manifest| should be NULL. If extension
	//  resources will be provided by the client (e.g. via ICefRequestHandler
	//  and/or ICefExtensionHandler) then |root_directory| should be a path
	//  component unique to the extension (if not absolute this will be internally
	//  prefixed with the PK_DIR_RESOURCES path) and |manifest| should contain the
	//  contents that would otherwise be read from the "manifest.json" file on
	//  disk.
	//
	//  The loaded extension will be accessible in all contexts sharing the same
	//  storage (HasExtension returns true (1)). However, only the context on
	//  which this function was called is considered the loader (DidLoadExtension
	//  returns true (1)) and only the loader will receive
	//  ICefRequestContextHandler callbacks for the extension.
	//
	//  ICefExtensionHandler.OnExtensionLoaded will be called on load success
	//  or ICefExtensionHandler.OnExtensionLoadFailed will be called on load
	//  failure.
	//
	//  If the extension specifies a background script via the "background"
	//  manifest key then ICefExtensionHandler.OnBeforeBackgroundBrowser will
	//  be called to create the background browser. See that function for
	//  additional information about background scripts.
	//
	//  For visible extension views the client application should evaluate the
	//  manifest to determine the correct extension URL to load and then pass that
	//  URL to the ICefBrowserHost.CreateBrowser* function after the extension
	//  has loaded. For example, the client can look for the "browser_action"
	//  manifest key as documented at
	//  https://developer.chrome.com/extensions/browserAction. Extension URLs take
	//  the form "chrome-extension://<extension_id>/<path>".
	//
	//  Browsers that host extensions differ from normal browsers as follows:
	//  - Can access chrome.* JavaScript APIs if allowed by the manifest. Visit
	//  chrome://extensions-support for the list of extension APIs currently
	//  supported by CEF.
	//  - Main frame navigation to non-extension content is blocked.
	//  - Pinch-zooming is disabled.
	//  - CefBrowserHost::GetExtension returns the hosted extension.
	//  - CefBrowserHost::IsBackgroundHost returns true for background hosts.
	//
	//  See https://developer.chrome.com/extensions for extension implementation
	//  and usage documentation.
	//
	//  WARNING: This function is deprecated and will be removed in ~M127.
	LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler) // procedure
	// SetWebsiteSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass nullptr for |value| to remove the
	//  default value for this content type.
	//
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value ICefValue) // procedure
	// SetContentSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass CEF_CONTENT_SETTING_VALUE_DEFAULT for
	//  |value| to use the default value for this content type.
	//
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value cefTypes.TCefContentSettingValues) // procedure
	// SetChromeColorScheme
	//  Sets the Chrome color scheme for all browsers that share this request
	//  context. |variant| values of SYSTEM, LIGHT and DARK change the underlying
	//  color mode (e.g. light vs dark). Other |variant| values determine how
	//  |user_color| will be applied in the current color mode. If |user_color| is
	//  transparent (0) the default color will be used.
	SetChromeColorScheme(variant cefTypes.TCefColorVariant, userColor cefTypes.TCefColor) // procedure
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

func (m *TCefRequestContextRef) GetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(14, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TCefRequestContextRef) GetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) cefTypes.TCefContentSettingValues {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestContextRefAPI().SysCallN(15, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType))
	return cefTypes.TCefContentSettingValues(r)
}

func (m *TCefRequestContextRef) GetChromeColorSchemeMode() cefTypes.TCefColorVariant {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestContextRefAPI().SysCallN(16, m.Instance())
	return cefTypes.TCefColorVariant(r)
}

func (m *TCefRequestContextRef) GetChromeColorSchemeColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestContextRefAPI().SysCallN(17, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefRequestContextRef) GetChromeColorSchemeVariant() cefTypes.TCefColorVariant {
	if !m.IsValid() {
		return 0
	}
	r := cefRequestContextRefAPI().SysCallN(18, m.Instance())
	return cefTypes.TCefColorVariant(r)
}

func (m *TCefRequestContextRef) ClearCertificateExceptions(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) ClearHttpAuthCredentials(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) CloseAllConnections(callback IEngCompletionCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) ResolveHost(origin string, callback IEngResolveCallback) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(27, m.Instance(), api.PasStr(origin), base.GetObjectUintptr(callback))
}

func (m *TCefRequestContextRef) LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(28, m.Instance(), api.PasStr(rootDirectory), base.GetObjectUintptr(manifest), base.GetObjectUintptr(handler))
}

func (m *TCefRequestContextRef) SetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value ICefValue) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(29, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), base.GetObjectUintptr(value))
}

func (m *TCefRequestContextRef) SetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value cefTypes.TCefContentSettingValues) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(30, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), uintptr(value))
}

func (m *TCefRequestContextRef) SetChromeColorScheme(variant cefTypes.TCefColorVariant, userColor cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefRequestContextRefAPI().SysCallN(31, m.Instance(), uintptr(variant), uintptr(userColor))
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
	cefRequestContextRefAPI().SysCallN(19, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

// GlobalToRequestContext
//
//	Returns the global context object.
func (_RequestContextRefClass) GlobalToRequestContext() (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(20, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

// NewWithPCefRequestContextSettingsRequestContextHandler
//
//	Creates a new context object with the specified |settings| and optional
//	|handler|.
//	<param name="settings">Pointer to TCefRequestContextSettings.</param>
//	<param name="handler">Optional handler for the request context.</param>
func (_RequestContextRefClass) NewWithPCefRequestContextSettingsRequestContextHandler(settings TCefRequestContextSettings, handler IEngRequestContextHandler) (result ICefRequestContext) {
	settingsPtr := settings.ToPas()
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(21, uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

// NewWithStrX3BoolX3RCHandler
//
//	Creates a new context object with the specified settings and optional
//	|handler|.
//	<param name="aCache">The directory where cache data for this request context will be stored on disk. See TCefRequestContextSettings.cache_path for more information.</param>
//	<param name="aAcceptLanguageList">Comma delimited ordered list of language codes without any whitespace that will be used in the "Accept-Language" HTTP header. See TCefRequestContextSettings.accept_language_list for more information.</param>
//	<param name="aCookieableSchemesList">Comma delimited list of schemes supported by the associated ICefCookieManager. See TCefRequestContextSettings.cookieable_schemes_list for more information.</param>
//	<param name="aCookieableSchemesExcludeDefaults">Setting this parameter to true will disable all loading and saving of cookies. See TCefRequestContextSettings.cookieable_schemes_list for more information.</param>
//	<param name="aPersistSessionCookies">To persist session cookies (cookies without an expiry date or validity interval) by default when using the global cookie manager set this value to true. See TCefRequestContextSettings.persist_session_cookies for more information.</param>
//	<param name="aPersistUserPreferences">To persist user preferences as a JSON file in the cache path directory set this value to true. See TCefRequestContextSettings.persist_user_preferences for more information.</param>
//	<param name="handler">Optional handler for the request context.</param>
func (_RequestContextRefClass) NewWithStrX3BoolX3RCHandler(cache string, acceptLanguageList string, cookieableSchemesList string, cookieableSchemesExcludeDefaults bool, persistSessionCookies bool, persistUserPreferences bool, handler IEngRequestContextHandler) (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(22, api.PasStr(cache), api.PasStr(acceptLanguageList), api.PasStr(cookieableSchemesList), api.PasBool(cookieableSchemesExcludeDefaults), api.PasBool(persistSessionCookies), api.PasBool(persistUserPreferences), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

// Shared
//
//	Creates a new context object that shares storage with |other| and uses an
//	optional |handler|.
//	<param name="other">Another ICefRequestContext instance that will share storage with the new ICefRequestContext instance.</param>
//	<param name="handler">Optional handler for the request context.</param>
func (_RequestContextRefClass) Shared(other ICefRequestContext, handler IEngRequestContextHandler) (result ICefRequestContext) {
	var resultPtr uintptr
	cefRequestContextRefAPI().SysCallN(23, base.GetObjectUintptr(other), base.GetObjectUintptr(handler), uintptr(base.UnsafePointer(&resultPtr)))
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
			/* 14 */ imports.NewTable("TCefRequestContextRef_GetWebsiteSetting", 0), // function GetWebsiteSetting
			/* 15 */ imports.NewTable("TCefRequestContextRef_GetContentSetting", 0), // function GetContentSetting
			/* 16 */ imports.NewTable("TCefRequestContextRef_GetChromeColorSchemeMode", 0), // function GetChromeColorSchemeMode
			/* 17 */ imports.NewTable("TCefRequestContextRef_GetChromeColorSchemeColor", 0), // function GetChromeColorSchemeColor
			/* 18 */ imports.NewTable("TCefRequestContextRef_GetChromeColorSchemeVariant", 0), // function GetChromeColorSchemeVariant
			/* 19 */ imports.NewTable("TCefRequestContextRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 20 */ imports.NewTable("TCefRequestContextRef_GlobalToRequestContext", 0), // static function GlobalToRequestContext
			/* 21 */ imports.NewTable("TCefRequestContextRef_NewWithPCefRequestContextSettingsRequestContextHandler", 0), // static function NewWithPCefRequestContextSettingsRequestContextHandler
			/* 22 */ imports.NewTable("TCefRequestContextRef_NewWithStrX3BoolX3RCHandler", 0), // static function NewWithStrX3BoolX3RCHandler
			/* 23 */ imports.NewTable("TCefRequestContextRef_Shared", 0), // static function Shared
			/* 24 */ imports.NewTable("TCefRequestContextRef_ClearCertificateExceptions", 0), // procedure ClearCertificateExceptions
			/* 25 */ imports.NewTable("TCefRequestContextRef_ClearHttpAuthCredentials", 0), // procedure ClearHttpAuthCredentials
			/* 26 */ imports.NewTable("TCefRequestContextRef_CloseAllConnections", 0), // procedure CloseAllConnections
			/* 27 */ imports.NewTable("TCefRequestContextRef_ResolveHost", 0), // procedure ResolveHost
			/* 28 */ imports.NewTable("TCefRequestContextRef_LoadExtension", 0), // procedure LoadExtension
			/* 29 */ imports.NewTable("TCefRequestContextRef_SetWebsiteSetting", 0), // procedure SetWebsiteSetting
			/* 30 */ imports.NewTable("TCefRequestContextRef_SetContentSetting", 0), // procedure SetContentSetting
			/* 31 */ imports.NewTable("TCefRequestContextRef_SetChromeColorScheme", 0), // procedure SetChromeColorScheme
		}
	})
	return cefRequestContextRefImport
}
