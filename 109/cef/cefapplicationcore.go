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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/109/types"
)

// ICefApplicationCore Parent: IObject
type ICefApplicationCore interface {
	IObject
	CheckCEFLibrary() bool                                                                                                                     // function
	StartMainProcess() bool                                                                                                                    // function
	StartSubProcess() bool                                                                                                                     // function
	InternalGetLocalizedString(stringId int32, stringVal *string) bool                                                                         // function
	InternalGetDataResource(resourceId int32, data *uintptr, dataSize *cefTypes.NativeUInt) bool                                               // function
	InternalGetDataResourceForScale(resourceId int32, scaleFactor cefTypes.TCefScaleFactor, data *uintptr, dataSize *cefTypes.NativeUInt) bool // function
	AfterConstruction()                                                                                                                        // procedure
	AddCustomCommandLine(commandLine string, value string)                                                                                     // procedure
	DoMessageLoopWork()                                                                                                                        // procedure
	RunMessageLoop()                                                                                                                           // procedure
	QuitMessageLoop()                                                                                                                          // procedure
	UpdateDeviceScaleFactor()                                                                                                                  // procedure
	InitLibLocationFromArgs()                                                                                                                  // procedure
	// InternalOnBeforeCommandLineProcessing
	//  Internal procedures. Only ICefApp, ICefBrowserProcessHandler,
	//  ICefResourceBundleHandler, ICefRenderProcessHandler, ICefRegisterCDMCallback and
	//  ICefLoadHandler should use them.
	InternalOnBeforeCommandLineProcessing(processType string, commandLine ICefCommandLine)                                                                  // procedure
	InternalOnRegisterCustomSchemes(registrar ICefSchemeRegistrarRef)                                                                                       // procedure
	InternalOnRegisterCustomPreferences(type_ cefTypes.TCefPreferencesType, registrar ICefPreferenceRegistrarRef)                                           // procedure
	InternalOnContextInitialized()                                                                                                                          // procedure
	InternalOnBeforeChildProcessLaunch(commandLine ICefCommandLine)                                                                                         // procedure
	InternalOnScheduleMessagePumpWork(delayMs int64)                                                                                                        // procedure
	InternalOnWebKitInitialized()                                                                                                                           // procedure
	InternalOnBrowserCreated(browser ICefBrowser, extraInfo ICefDictionaryValue)                                                                            // procedure
	InternalOnBrowserDestroyed(browser ICefBrowser)                                                                                                         // procedure
	InternalOnContextCreated(browser ICefBrowser, frame ICefFrame, context ICefv8Context)                                                                   // procedure
	InternalOnContextReleased(browser ICefBrowser, frame ICefFrame, context ICefv8Context)                                                                  // procedure
	InternalOnUncaughtException(browser ICefBrowser, frame ICefFrame, context ICefv8Context, exception ICefV8Exception, stackTrace ICefV8StackTrace)        // procedure
	InternalOnFocusedNodeChanged(browser ICefBrowser, frame ICefFrame, node ICefDomNode)                                                                    // procedure
	InternalOnProcessMessageReceived(browser ICefBrowser, frame ICefFrame, sourceProcess cefTypes.TCefProcessId, message ICefProcessMessage, handled *bool) // procedure
	InternalOnLoadingStateChange(browser ICefBrowser, isLoading bool, canGoBack bool, canGoForward bool)                                                    // procedure
	InternalOnLoadStart(browser ICefBrowser, frame ICefFrame, transitionType cefTypes.TCefTransitionType)                                                   // procedure
	InternalOnLoadEnd(browser ICefBrowser, frame ICefFrame, httpStatusCode int32)                                                                           // procedure
	InternalOnLoadError(browser ICefBrowser, frame ICefFrame, errorCode int32, errorText string, failedUrl string)                                          // procedure
	InternalGetDefaultClient(client *IEngClient)                                                                                                            // procedure
	// NoSandbox
	//  Properties used to populate TCefSettings (cef_settings_t)
	NoSandbox() bool                                // property NoSandbox Getter
	SetNoSandbox(value bool)                        // property NoSandbox Setter
	BrowserSubprocessPath() string                  // property BrowserSubprocessPath Getter
	SetBrowserSubprocessPath(value string)          // property BrowserSubprocessPath Setter
	FrameworkDirPath() string                       // property FrameworkDirPath Getter
	SetFrameworkDirPath(value string)               // property FrameworkDirPath Setter
	MainBundlePath() string                         // property MainBundlePath Getter
	SetMainBundlePath(value string)                 // property MainBundlePath Setter
	ChromeRuntime() bool                            // property ChromeRuntime Getter
	SetChromeRuntime(value bool)                    // property ChromeRuntime Setter
	MultiThreadedMessageLoop() bool                 // property MultiThreadedMessageLoop Getter
	SetMultiThreadedMessageLoop(value bool)         // property MultiThreadedMessageLoop Setter
	ExternalMessagePump() bool                      // property ExternalMessagePump Getter
	SetExternalMessagePump(value bool)              // property ExternalMessagePump Setter
	WindowlessRenderingEnabled() bool               // property WindowlessRenderingEnabled Getter
	SetWindowlessRenderingEnabled(value bool)       // property WindowlessRenderingEnabled Setter
	CommandLineArgsDisabled() bool                  // property CommandLineArgsDisabled Getter
	SetCommandLineArgsDisabled(value bool)          // property CommandLineArgsDisabled Setter
	Cache() string                                  // property Cache Getter
	SetCache(value string)                          // property Cache Setter
	RootCache() string                              // property RootCache Getter
	SetRootCache(value string)                      // property RootCache Setter
	UserDataPath() string                           // property UserDataPath Getter
	SetUserDataPath(value string)                   // property UserDataPath Setter
	PersistSessionCookies() bool                    // property PersistSessionCookies Getter
	SetPersistSessionCookies(value bool)            // property PersistSessionCookies Setter
	PersistUserPreferences() bool                   // property PersistUserPreferences Getter
	SetPersistUserPreferences(value bool)           // property PersistUserPreferences Setter
	UserAgent() string                              // property UserAgent Getter
	SetUserAgent(value string)                      // property UserAgent Setter
	UserAgentProduct() string                       // property UserAgentProduct Getter
	SetUserAgentProduct(value string)               // property UserAgentProduct Setter
	Locale() string                                 // property Locale Getter
	SetLocale(value string)                         // property Locale Setter
	LogFile() string                                // property LogFile Getter
	SetLogFile(value string)                        // property LogFile Setter
	LogSeverity() cefTypes.TCefLogSeverity          // property LogSeverity Getter
	SetLogSeverity(value cefTypes.TCefLogSeverity)  // property LogSeverity Setter
	JavaScriptFlags() string                        // property JavaScriptFlags Getter
	SetJavaScriptFlags(value string)                // property JavaScriptFlags Setter
	ResourcesDirPath() string                       // property ResourcesDirPath Getter
	SetResourcesDirPath(value string)               // property ResourcesDirPath Setter
	LocalesDirPath() string                         // property LocalesDirPath Getter
	SetLocalesDirPath(value string)                 // property LocalesDirPath Setter
	PackLoadingDisabled() bool                      // property PackLoadingDisabled Getter
	SetPackLoadingDisabled(value bool)              // property PackLoadingDisabled Setter
	RemoteDebuggingPort() int32                     // property RemoteDebuggingPort Getter
	SetRemoteDebuggingPort(value int32)             // property RemoteDebuggingPort Setter
	UncaughtExceptionStackSize() int32              // property UncaughtExceptionStackSize Getter
	SetUncaughtExceptionStackSize(value int32)      // property UncaughtExceptionStackSize Setter
	IgnoreCertificateErrors() bool                  // property IgnoreCertificateErrors Getter
	SetIgnoreCertificateErrors(value bool)          // property IgnoreCertificateErrors Setter
	BackgroundColor() cefTypes.TCefColor            // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor)    // property BackgroundColor Setter
	AcceptLanguageList() string                     // property AcceptLanguageList Getter
	SetAcceptLanguageList(value string)             // property AcceptLanguageList Setter
	CookieableSchemesList() string                  // property CookieableSchemesList Getter
	SetCookieableSchemesList(value string)          // property CookieableSchemesList Setter
	CookieableSchemesExcludeDefaults() bool         // property CookieableSchemesExcludeDefaults Getter
	SetCookieableSchemesExcludeDefaults(value bool) // property CookieableSchemesExcludeDefaults Setter
	// SingleProcess
	//  Properties used to set command line switches
	SingleProcess() bool                                       // property SingleProcess Getter
	SetSingleProcess(value bool)                               // property SingleProcess Setter
	EnableMediaStream() bool                                   // property EnableMediaStream Getter
	SetEnableMediaStream(value bool)                           // property EnableMediaStream Setter
	EnableSpeechInput() bool                                   // property EnableSpeechInput Getter
	SetEnableSpeechInput(value bool)                           // property EnableSpeechInput Setter
	UseFakeUIForMediaStream() bool                             // property UseFakeUIForMediaStream Getter
	SetUseFakeUIForMediaStream(value bool)                     // property UseFakeUIForMediaStream Setter
	EnableUsermediaScreenCapturing() bool                      // property EnableUsermediaScreenCapturing Getter
	SetEnableUsermediaScreenCapturing(value bool)              // property EnableUsermediaScreenCapturing Setter
	EnableGPU() bool                                           // property EnableGPU Getter
	SetEnableGPU(value bool)                                   // property EnableGPU Setter
	EnableFeatures() string                                    // property EnableFeatures Getter
	SetEnableFeatures(value string)                            // property EnableFeatures Setter
	DisableFeatures() string                                   // property DisableFeatures Getter
	SetDisableFeatures(value string)                           // property DisableFeatures Setter
	EnableBlinkFeatures() string                               // property EnableBlinkFeatures Getter
	SetEnableBlinkFeatures(value string)                       // property EnableBlinkFeatures Setter
	DisableBlinkFeatures() string                              // property DisableBlinkFeatures Getter
	SetDisableBlinkFeatures(value string)                      // property DisableBlinkFeatures Setter
	BlinkSettings() string                                     // property BlinkSettings Getter
	SetBlinkSettings(value string)                             // property BlinkSettings Setter
	ForceFieldTrials() string                                  // property ForceFieldTrials Getter
	SetForceFieldTrials(value string)                          // property ForceFieldTrials Setter
	ForceFieldTrialParams() string                             // property ForceFieldTrialParams Getter
	SetForceFieldTrialParams(value string)                     // property ForceFieldTrialParams Setter
	SmoothScrolling() cefTypes.TCefState                       // property SmoothScrolling Getter
	SetSmoothScrolling(value cefTypes.TCefState)               // property SmoothScrolling Setter
	FastUnload() bool                                          // property FastUnload Getter
	SetFastUnload(value bool)                                  // property FastUnload Setter
	DisableSafeBrowsing() bool                                 // property DisableSafeBrowsing Getter
	SetDisableSafeBrowsing(value bool)                         // property DisableSafeBrowsing Setter
	MuteAudio() bool                                           // property MuteAudio Getter
	SetMuteAudio(value bool)                                   // property MuteAudio Setter
	SitePerProcess() bool                                      // property SitePerProcess Getter
	SetSitePerProcess(value bool)                              // property SitePerProcess Setter
	DisableWebSecurity() bool                                  // property DisableWebSecurity Getter
	SetDisableWebSecurity(value bool)                          // property DisableWebSecurity Setter
	DisablePDFExtension() bool                                 // property DisablePDFExtension Getter
	SetDisablePDFExtension(value bool)                         // property DisablePDFExtension Setter
	DisableSiteIsolationTrials() bool                          // property DisableSiteIsolationTrials Getter
	SetDisableSiteIsolationTrials(value bool)                  // property DisableSiteIsolationTrials Setter
	DisableChromeLoginPrompt() bool                            // property DisableChromeLoginPrompt Getter
	SetDisableChromeLoginPrompt(value bool)                    // property DisableChromeLoginPrompt Setter
	DisableExtensions() bool                                   // property DisableExtensions Getter
	SetDisableExtensions(value bool)                           // property DisableExtensions Setter
	AutoplayPolicy() cefTypes.TCefAutoplayPolicy               // property AutoplayPolicy Getter
	SetAutoplayPolicy(value cefTypes.TCefAutoplayPolicy)       // property AutoplayPolicy Setter
	DisableBackgroundNetworking() bool                         // property DisableBackgroundNetworking Getter
	SetDisableBackgroundNetworking(value bool)                 // property DisableBackgroundNetworking Setter
	MetricsRecordingOnly() bool                                // property MetricsRecordingOnly Getter
	SetMetricsRecordingOnly(value bool)                        // property MetricsRecordingOnly Setter
	AllowFileAccessFromFiles() bool                            // property AllowFileAccessFromFiles Getter
	SetAllowFileAccessFromFiles(value bool)                    // property AllowFileAccessFromFiles Setter
	AllowRunningInsecureContent() bool                         // property AllowRunningInsecureContent Getter
	SetAllowRunningInsecureContent(value bool)                 // property AllowRunningInsecureContent Setter
	EnablePrintPreview() bool                                  // property EnablePrintPreview Getter
	SetEnablePrintPreview(value bool)                          // property EnablePrintPreview Setter
	DefaultEncoding() string                                   // property DefaultEncoding Getter
	SetDefaultEncoding(value string)                           // property DefaultEncoding Setter
	DisableJavascript() bool                                   // property DisableJavascript Getter
	SetDisableJavascript(value bool)                           // property DisableJavascript Setter
	DisableJavascriptCloseWindows() bool                       // property DisableJavascriptCloseWindows Getter
	SetDisableJavascriptCloseWindows(value bool)               // property DisableJavascriptCloseWindows Setter
	DisableJavascriptAccessClipboard() bool                    // property DisableJavascriptAccessClipboard Getter
	SetDisableJavascriptAccessClipboard(value bool)            // property DisableJavascriptAccessClipboard Setter
	DisableJavascriptDomPaste() bool                           // property DisableJavascriptDomPaste Getter
	SetDisableJavascriptDomPaste(value bool)                   // property DisableJavascriptDomPaste Setter
	AllowUniversalAccessFromFileUrls() bool                    // property AllowUniversalAccessFromFileUrls Getter
	SetAllowUniversalAccessFromFileUrls(value bool)            // property AllowUniversalAccessFromFileUrls Setter
	DisableImageLoading() bool                                 // property DisableImageLoading Getter
	SetDisableImageLoading(value bool)                         // property DisableImageLoading Setter
	ImageShrinkStandaloneToFit() bool                          // property ImageShrinkStandaloneToFit Getter
	SetImageShrinkStandaloneToFit(value bool)                  // property ImageShrinkStandaloneToFit Setter
	DisableTextAreaResize() bool                               // property DisableTextAreaResize Getter
	SetDisableTextAreaResize(value bool)                       // property DisableTextAreaResize Setter
	DisableTabToLinks() bool                                   // property DisableTabToLinks Getter
	SetDisableTabToLinks(value bool)                           // property DisableTabToLinks Setter
	EnableProfanityFilter() bool                               // property EnableProfanityFilter Getter
	SetEnableProfanityFilter(value bool)                       // property EnableProfanityFilter Setter
	DisableSpellChecking() bool                                // property DisableSpellChecking Getter
	SetDisableSpellChecking(value bool)                        // property DisableSpellChecking Setter
	OverrideSpellCheckLang() string                            // property OverrideSpellCheckLang Getter
	SetOverrideSpellCheckLang(value string)                    // property OverrideSpellCheckLang Setter
	TouchEvents() cefTypes.TCefState                           // property TouchEvents Getter
	SetTouchEvents(value cefTypes.TCefState)                   // property TouchEvents Setter
	DisableReadingFromCanvas() bool                            // property DisableReadingFromCanvas Getter
	SetDisableReadingFromCanvas(value bool)                    // property DisableReadingFromCanvas Setter
	HyperlinkAuditing() bool                                   // property HyperlinkAuditing Getter
	SetHyperlinkAuditing(value bool)                           // property HyperlinkAuditing Setter
	DisableNewBrowserInfoTimeout() bool                        // property DisableNewBrowserInfoTimeout Getter
	SetDisableNewBrowserInfoTimeout(value bool)                // property DisableNewBrowserInfoTimeout Setter
	DevToolsProtocolLogFile() string                           // property DevToolsProtocolLogFile Getter
	SetDevToolsProtocolLogFile(value string)                   // property DevToolsProtocolLogFile Setter
	ForcedDeviceScaleFactor() float32                          // property ForcedDeviceScaleFactor Getter
	SetForcedDeviceScaleFactor(value float32)                  // property ForcedDeviceScaleFactor Setter
	DisableZygote() bool                                       // property DisableZygote Getter
	SetDisableZygote(value bool)                               // property DisableZygote Setter
	UseMockKeyChain() bool                                     // property UseMockKeyChain Getter
	SetUseMockKeyChain(value bool)                             // property UseMockKeyChain Setter
	DisableRequestHandlingForTesting() bool                    // property DisableRequestHandlingForTesting Getter
	SetDisableRequestHandlingForTesting(value bool)            // property DisableRequestHandlingForTesting Setter
	DisablePopupBlocking() bool                                // property DisablePopupBlocking Getter
	SetDisablePopupBlocking(value bool)                        // property DisablePopupBlocking Setter
	DisableBackForwardCache() bool                             // property DisableBackForwardCache Getter
	SetDisableBackForwardCache(value bool)                     // property DisableBackForwardCache Setter
	DisableComponentUpdate() bool                              // property DisableComponentUpdate Getter
	SetDisableComponentUpdate(value bool)                      // property DisableComponentUpdate Setter
	AllowInsecureLocalhost() bool                              // property AllowInsecureLocalhost Getter
	SetAllowInsecureLocalhost(value bool)                      // property AllowInsecureLocalhost Setter
	KioskPrinting() bool                                       // property KioskPrinting Getter
	SetKioskPrinting(value bool)                               // property KioskPrinting Setter
	TreatInsecureOriginAsSecure() string                       // property TreatInsecureOriginAsSecure Getter
	SetTreatInsecureOriginAsSecure(value string)               // property TreatInsecureOriginAsSecure Setter
	NetLogEnabled() bool                                       // property NetLogEnabled Getter
	SetNetLogEnabled(value bool)                               // property NetLogEnabled Setter
	NetLogFile() string                                        // property NetLogFile Getter
	SetNetLogFile(value string)                                // property NetLogFile Setter
	NetLogCaptureMode() cefTypes.TCefNetLogCaptureMode         // property NetLogCaptureMode Getter
	SetNetLogCaptureMode(value cefTypes.TCefNetLogCaptureMode) // property NetLogCaptureMode Setter
	// WindowsSandboxInfo
	//  Properties used during the CEF initialization
	WindowsSandboxInfo() uintptr         // property WindowsSandboxInfo Getter
	SetWindowsSandboxInfo(value uintptr) // property WindowsSandboxInfo Setter
	EnableHighDPISupport() bool          // property EnableHighDPISupport Getter
	SetEnableHighDPISupport(value bool)  // property EnableHighDPISupport Setter
	ArgcCopy() int32                     // property argcCopy Getter
	ArgvCopy() types.PPAnsiChar          // property argvCopy Getter
	// DeleteCache
	//  Custom properties
	DeleteCache() bool                                                   // property DeleteCache Getter
	SetDeleteCache(value bool)                                           // property DeleteCache Setter
	DeleteCookies() bool                                                 // property DeleteCookies Getter
	SetDeleteCookies(value bool)                                         // property DeleteCookies Setter
	CheckCEFFiles() bool                                                 // property CheckCEFFiles Getter
	SetCheckCEFFiles(value bool)                                         // property CheckCEFFiles Setter
	ShowMessageDlg() bool                                                // property ShowMessageDlg Getter
	SetShowMessageDlg(value bool)                                        // property ShowMessageDlg Setter
	MissingBinariesException() bool                                      // property MissingBinariesException Getter
	SetMissingBinariesException(value bool)                              // property MissingBinariesException Setter
	SetCurrentDir() bool                                                 // property SetCurrentDir Getter
	SetSetCurrentDir(value bool)                                         // property SetCurrentDir Setter
	GlobalContextInitialized() bool                                      // property GlobalContextInitialized Getter
	ChromeMajorVer() uint16                                              // property ChromeMajorVer Getter
	ChromeMinorVer() uint16                                              // property ChromeMinorVer Getter
	ChromeRelease() uint16                                               // property ChromeRelease Getter
	ChromeBuild() uint16                                                 // property ChromeBuild Getter
	ChromeVersion() string                                               // property ChromeVersion Getter
	LibCefVersion() string                                               // property LibCefVersion Getter
	LibCefPath() string                                                  // property LibCefPath Getter
	ChromeElfPath() string                                               // property ChromeElfPath Getter
	LibLoaded() bool                                                     // property LibLoaded Getter
	LogProcessInfo() bool                                                // property LogProcessInfo Getter
	SetLogProcessInfo(value bool)                                        // property LogProcessInfo Setter
	ReRaiseExceptions() bool                                             // property ReRaiseExceptions Getter
	SetReRaiseExceptions(value bool)                                     // property ReRaiseExceptions Setter
	DeviceScaleFactor() float32                                          // property DeviceScaleFactor Getter
	LocalesRequired() string                                             // property LocalesRequired Getter
	SetLocalesRequired(value string)                                     // property LocalesRequired Setter
	ProcessType() cefTypes.TCefProcessType                               // property ProcessType Getter
	MustCreateResourceBundleHandler() bool                               // property MustCreateResourceBundleHandler Getter
	SetMustCreateResourceBundleHandler(value bool)                       // property MustCreateResourceBundleHandler Setter
	MustCreateBrowserProcessHandler() bool                               // property MustCreateBrowserProcessHandler Getter
	SetMustCreateBrowserProcessHandler(value bool)                       // property MustCreateBrowserProcessHandler Setter
	MustCreateRenderProcessHandler() bool                                // property MustCreateRenderProcessHandler Getter
	SetMustCreateRenderProcessHandler(value bool)                        // property MustCreateRenderProcessHandler Setter
	MustCreateLoadHandler() bool                                         // property MustCreateLoadHandler Getter
	SetMustCreateLoadHandler(value bool)                                 // property MustCreateLoadHandler Setter
	SetOsmodalLoop(value bool)                                           // property OsmodalLoop Setter
	Status() cefTypes.TCefAplicationStatus                               // property Status Getter
	MissingLibFiles() string                                             // property MissingLibFiles Getter
	MustFreeLibrary() bool                                               // property MustFreeLibrary Getter
	SetMustFreeLibrary(value bool)                                       // property MustFreeLibrary Setter
	ChildProcessesCount() int32                                          // property ChildProcessesCount Getter
	UsedMemory() uint64                                                  // property UsedMemory Getter
	TotalSystemMemory() uint64                                           // property TotalSystemMemory Getter
	AvailableSystemMemory() uint64                                       // property AvailableSystemMemory Getter
	SystemMemoryLoad() uint32                                            // property SystemMemoryLoad Getter
	ApiHashUniversal() string                                            // property ApiHashUniversal Getter
	ApiHashPlatform() string                                             // property ApiHashPlatform Getter
	ApiHashCommit() string                                               // property ApiHashCommit Getter
	LastErrorMessage() string                                            // property LastErrorMessage Getter
	XDisplay() uintptr                                                   // property XDisplay Getter
	SetOnRegCustomSchemes(fn TOnRegisterCustomSchemesEvent)              // property event
	SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferencesEvent) // property event
	SetOnContextInitialized(fn TOnContextInitializedEvent)               // property event
	SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunchEvent)   // property event
	SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWorkEvent)     // property event
	SetOnGetDefaultClient(fn TOnGetDefaultClientEvent)                   // property event
	SetOnGetLocalizedString(fn TOnGetLocalizedStringEvent)               // property event
	SetOnGetDataResource(fn TOnGetDataResourceEvent)                     // property event
	SetOnGetDataResourceForScale(fn TOnGetDataResourceForScaleEvent)     // property event
	SetOnWebKitInitialized(fn TOnWebKitInitializedEvent)                 // property event
	SetOnBrowserCreated(fn TOnBrowserCreatedEvent)                       // property event
	SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent)                   // property event
	SetOnContextCreated(fn TOnContextCreatedEvent)                       // property event
	SetOnContextReleased(fn TOnContextReleasedEvent)                     // property event
	SetOnUncaughtException(fn TOnUncaughtExceptionEvent)                 // property event
	SetOnFocusedNodeChanged(fn TOnFocusedNodeChangedEvent)               // property event
	SetOnProcessMessageReceived(fn TOnProcessMessageReceivedEvent)       // property event
	SetOnLoadingStateChange(fn TOnRenderLoadingStateChange)              // property event
	SetOnLoadStart(fn TOnRenderLoadStart)                                // property event
	SetOnLoadEnd(fn TOnRenderLoadEnd)                                    // property event
	SetOnLoadError(fn TOnRenderLoadError)                                // property event
}

type TCefApplicationCore struct {
	TObject
}

func (m *TCefApplicationCore) CheckCEFLibrary() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) StartMainProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) StartSubProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) InternalGetLocalizedString(stringId int32, stringVal *string) bool {
	if !m.IsValid() {
		return false
	}
	stringValPtr := api.PasStr(*stringVal)
	r := cefApplicationCoreAPI().SysCallN(4, m.Instance(), uintptr(stringId), uintptr(base.UnsafePointer(&stringValPtr)))
	*stringVal = api.GoStr(stringValPtr)
	return api.GoBool(r)
}

func (m *TCefApplicationCore) InternalGetDataResource(resourceId int32, data *uintptr, dataSize *cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	dataPtr := uintptr(*data)
	dataSizePtr := uintptr(*dataSize)
	r := cefApplicationCoreAPI().SysCallN(5, m.Instance(), uintptr(resourceId), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(&dataSizePtr)))
	*data = uintptr(dataPtr)
	*dataSize = cefTypes.NativeUInt(dataSizePtr)
	return api.GoBool(r)
}

func (m *TCefApplicationCore) InternalGetDataResourceForScale(resourceId int32, scaleFactor cefTypes.TCefScaleFactor, data *uintptr, dataSize *cefTypes.NativeUInt) bool {
	if !m.IsValid() {
		return false
	}
	dataPtr := uintptr(*data)
	dataSizePtr := uintptr(*dataSize)
	r := cefApplicationCoreAPI().SysCallN(6, m.Instance(), uintptr(resourceId), uintptr(scaleFactor), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(&dataSizePtr)))
	*data = uintptr(dataPtr)
	*dataSize = cefTypes.NativeUInt(dataSizePtr)
	return api.GoBool(r)
}

func (m *TCefApplicationCore) AfterConstruction() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(7, m.Instance())
}

func (m *TCefApplicationCore) AddCustomCommandLine(commandLine string, value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(8, m.Instance(), api.PasStr(commandLine), api.PasStr(value))
}

func (m *TCefApplicationCore) DoMessageLoopWork() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(9, m.Instance())
}

func (m *TCefApplicationCore) RunMessageLoop() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(10, m.Instance())
}

func (m *TCefApplicationCore) QuitMessageLoop() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(11, m.Instance())
}

func (m *TCefApplicationCore) UpdateDeviceScaleFactor() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(12, m.Instance())
}

func (m *TCefApplicationCore) InitLibLocationFromArgs() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(13, m.Instance())
}

func (m *TCefApplicationCore) InternalOnBeforeCommandLineProcessing(processType string, commandLine ICefCommandLine) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(14, m.Instance(), api.PasStr(processType), base.GetObjectUintptr(commandLine))
}

func (m *TCefApplicationCore) InternalOnRegisterCustomSchemes(registrar ICefSchemeRegistrarRef) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(registrar))
}

func (m *TCefApplicationCore) InternalOnRegisterCustomPreferences(type_ cefTypes.TCefPreferencesType, registrar ICefPreferenceRegistrarRef) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(16, m.Instance(), uintptr(type_), base.GetObjectUintptr(registrar))
}

func (m *TCefApplicationCore) InternalOnContextInitialized() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(17, m.Instance())
}

func (m *TCefApplicationCore) InternalOnBeforeChildProcessLaunch(commandLine ICefCommandLine) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(18, m.Instance(), base.GetObjectUintptr(commandLine))
}

func (m *TCefApplicationCore) InternalOnScheduleMessagePumpWork(delayMs int64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&delayMs)))
}

func (m *TCefApplicationCore) InternalOnWebKitInitialized() {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(20, m.Instance())
}

func (m *TCefApplicationCore) InternalOnBrowserCreated(browser ICefBrowser, extraInfo ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(extraInfo))
}

func (m *TCefApplicationCore) InternalOnBrowserDestroyed(browser ICefBrowser) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(browser))
}

func (m *TCefApplicationCore) InternalOnContextCreated(browser ICefBrowser, frame ICefFrame, context ICefv8Context) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), base.GetObjectUintptr(context))
}

func (m *TCefApplicationCore) InternalOnContextReleased(browser ICefBrowser, frame ICefFrame, context ICefv8Context) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), base.GetObjectUintptr(context))
}

func (m *TCefApplicationCore) InternalOnUncaughtException(browser ICefBrowser, frame ICefFrame, context ICefv8Context, exception ICefV8Exception, stackTrace ICefV8StackTrace) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), base.GetObjectUintptr(context), base.GetObjectUintptr(exception), base.GetObjectUintptr(stackTrace))
}

func (m *TCefApplicationCore) InternalOnFocusedNodeChanged(browser ICefBrowser, frame ICefFrame, node ICefDomNode) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), base.GetObjectUintptr(node))
}

func (m *TCefApplicationCore) InternalOnProcessMessageReceived(browser ICefBrowser, frame ICefFrame, sourceProcess cefTypes.TCefProcessId, message ICefProcessMessage, handled *bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), uintptr(sourceProcess), base.GetObjectUintptr(message), uintptr(base.UnsafePointer(handled)))
}

func (m *TCefApplicationCore) InternalOnLoadingStateChange(browser ICefBrowser, isLoading bool, canGoBack bool, canGoForward bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(28, m.Instance(), base.GetObjectUintptr(browser), api.PasBool(isLoading), api.PasBool(canGoBack), api.PasBool(canGoForward))
}

func (m *TCefApplicationCore) InternalOnLoadStart(browser ICefBrowser, frame ICefFrame, transitionType cefTypes.TCefTransitionType) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(29, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), uintptr(transitionType))
}

func (m *TCefApplicationCore) InternalOnLoadEnd(browser ICefBrowser, frame ICefFrame, httpStatusCode int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(30, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), uintptr(httpStatusCode))
}

func (m *TCefApplicationCore) InternalOnLoadError(browser ICefBrowser, frame ICefFrame, errorCode int32, errorText string, failedUrl string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(browser), base.GetObjectUintptr(frame), uintptr(errorCode), api.PasStr(errorText), api.PasStr(failedUrl))
}

func (m *TCefApplicationCore) InternalGetDefaultClient(client *IEngClient) {
	if !m.IsValid() {
		return
	}
	clientPtr := base.GetObjectUintptr(*client)
	cefApplicationCoreAPI().SysCallN(32, m.Instance(), uintptr(base.UnsafePointer(&clientPtr)))
	*client = AsEngClient(clientPtr)
}

func (m *TCefApplicationCore) NoSandbox() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetNoSandbox(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) BrowserSubprocessPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(34, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetBrowserSubprocessPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(34, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) FrameworkDirPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(35, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetFrameworkDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(35, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) MainBundlePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(36, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetMainBundlePath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(36, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ChromeRuntime() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetChromeRuntime(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MultiThreadedMessageLoop() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(38, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMultiThreadedMessageLoop(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(38, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ExternalMessagePump() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(39, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetExternalMessagePump(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(39, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) WindowlessRenderingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(40, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetWindowlessRenderingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(40, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) CommandLineArgsDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(41, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCommandLineArgsDisabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(41, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) Cache() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(42, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetCache(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(42, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) RootCache() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(43, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetRootCache(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(43, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) UserDataPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(44, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetUserDataPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(44, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) PersistSessionCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(45, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPersistSessionCookies(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(45, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) PersistUserPreferences() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(46, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPersistUserPreferences(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(46, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UserAgent() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(47, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetUserAgent(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(47, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) UserAgentProduct() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(48, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetUserAgentProduct(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(48, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) Locale() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(49, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetLocale(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(49, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LogFile() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(50, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(50, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LogSeverity() cefTypes.TCefLogSeverity {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(51, 0, m.Instance())
	return cefTypes.TCefLogSeverity(r)
}

func (m *TCefApplicationCore) SetLogSeverity(value cefTypes.TCefLogSeverity) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) JavaScriptFlags() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(52, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetJavaScriptFlags(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(52, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ResourcesDirPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(53, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetResourcesDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(53, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) LocalesDirPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(54, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetLocalesDirPath(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(54, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) PackLoadingDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetPackLoadingDisabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) RemoteDebuggingPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(56, 0, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) SetRemoteDebuggingPort(value int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(56, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) UncaughtExceptionStackSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(57, 0, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) SetUncaughtExceptionStackSize(value int32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) IgnoreCertificateErrors() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(58, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetIgnoreCertificateErrors(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(58, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) BackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(59, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TCefApplicationCore) SetBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) AcceptLanguageList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(60, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetAcceptLanguageList(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(60, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) CookieableSchemesList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(61, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetCookieableSchemesList(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(61, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) CookieableSchemesExcludeDefaults() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCookieableSchemesExcludeDefaults(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SingleProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(63, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSingleProcess(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(63, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableMediaStream() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableMediaStream(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableSpeechInput() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(65, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableSpeechInput(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(65, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UseFakeUIForMediaStream() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetUseFakeUIForMediaStream(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableUsermediaScreenCapturing() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(67, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableUsermediaScreenCapturing(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(67, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableGPU() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(68, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableGPU(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(68, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableFeatures() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(69, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetEnableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(69, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableFeatures() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(70, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetDisableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(70, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) EnableBlinkFeatures() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(71, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetEnableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(71, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableBlinkFeatures() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(72, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetDisableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(72, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) BlinkSettings() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(73, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetBlinkSettings(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(73, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForceFieldTrials() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(74, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetForceFieldTrials(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(74, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForceFieldTrialParams() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(75, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetForceFieldTrialParams(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(75, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) SmoothScrolling() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(76, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TCefApplicationCore) SetSmoothScrolling(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(76, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) FastUnload() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(77, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetFastUnload(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(77, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableSafeBrowsing() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(78, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableSafeBrowsing(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(78, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MuteAudio() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(79, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMuteAudio(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(79, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SitePerProcess() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(80, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSitePerProcess(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(80, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableWebSecurity() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(81, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableWebSecurity(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(81, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisablePDFExtension() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(82, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisablePDFExtension(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(82, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableSiteIsolationTrials() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(83, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableSiteIsolationTrials(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(83, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableChromeLoginPrompt() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(84, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableChromeLoginPrompt(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(84, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableExtensions() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(85, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableExtensions(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(85, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AutoplayPolicy() cefTypes.TCefAutoplayPolicy {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(86, 0, m.Instance())
	return cefTypes.TCefAutoplayPolicy(r)
}

func (m *TCefApplicationCore) SetAutoplayPolicy(value cefTypes.TCefAutoplayPolicy) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(86, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) DisableBackgroundNetworking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(87, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableBackgroundNetworking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(87, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MetricsRecordingOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(88, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMetricsRecordingOnly(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(88, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowFileAccessFromFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(89, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowFileAccessFromFiles(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(89, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowRunningInsecureContent() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(90, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowRunningInsecureContent(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(90, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnablePrintPreview() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(91, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnablePrintPreview(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(91, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DefaultEncoding() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(92, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(92, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) DisableJavascript() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(93, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascript(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(93, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptCloseWindows() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(94, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptCloseWindows(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(94, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptAccessClipboard() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(95, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptAccessClipboard(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(95, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableJavascriptDomPaste() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(96, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableJavascriptDomPaste(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(96, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowUniversalAccessFromFileUrls() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(97, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowUniversalAccessFromFileUrls(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(97, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableImageLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(98, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableImageLoading(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(98, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ImageShrinkStandaloneToFit() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(99, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetImageShrinkStandaloneToFit(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(99, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableTextAreaResize() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(100, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableTextAreaResize(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(100, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableTabToLinks() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(101, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableTabToLinks(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(101, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) EnableProfanityFilter() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(102, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableProfanityFilter(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(102, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableSpellChecking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(103, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableSpellChecking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(103, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) OverrideSpellCheckLang() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(104, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetOverrideSpellCheckLang(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(104, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) TouchEvents() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(105, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TCefApplicationCore) SetTouchEvents(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(105, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) DisableReadingFromCanvas() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(106, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableReadingFromCanvas(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(106, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) HyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(107, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetHyperlinkAuditing(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(107, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableNewBrowserInfoTimeout() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(108, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableNewBrowserInfoTimeout(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(108, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DevToolsProtocolLogFile() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(109, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetDevToolsProtocolLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(109, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ForcedDeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(110, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) SetForcedDeviceScaleFactor(value float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(110, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCefApplicationCore) DisableZygote() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(111, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableZygote(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(111, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) UseMockKeyChain() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(112, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetUseMockKeyChain(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(112, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableRequestHandlingForTesting() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(113, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableRequestHandlingForTesting(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(113, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisablePopupBlocking() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(114, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisablePopupBlocking(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(114, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableBackForwardCache() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(115, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableBackForwardCache(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(115, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DisableComponentUpdate() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(116, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDisableComponentUpdate(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(116, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) AllowInsecureLocalhost() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(117, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetAllowInsecureLocalhost(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(117, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) KioskPrinting() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(118, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetKioskPrinting(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(118, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) TreatInsecureOriginAsSecure() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(119, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetTreatInsecureOriginAsSecure(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(119, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) NetLogEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(120, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetNetLogEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(120, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) NetLogFile() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(121, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetNetLogFile(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(121, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) NetLogCaptureMode() cefTypes.TCefNetLogCaptureMode {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(122, 0, m.Instance())
	return cefTypes.TCefNetLogCaptureMode(r)
}

func (m *TCefApplicationCore) SetNetLogCaptureMode(value cefTypes.TCefNetLogCaptureMode) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(122, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) WindowsSandboxInfo() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(123, 0, m.Instance())
	return uintptr(r)
}

func (m *TCefApplicationCore) SetWindowsSandboxInfo(value uintptr) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(123, 1, m.Instance(), uintptr(value))
}

func (m *TCefApplicationCore) EnableHighDPISupport() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(124, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetEnableHighDPISupport(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(124, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ArgcCopy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(125, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) ArgvCopy() types.PPAnsiChar {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(126, m.Instance())
	return types.PPAnsiChar(r)
}

func (m *TCefApplicationCore) DeleteCache() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(127, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDeleteCache(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(127, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DeleteCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(128, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetDeleteCookies(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(128, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) CheckCEFFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(129, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetCheckCEFFiles(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(129, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ShowMessageDlg() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(130, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetShowMessageDlg(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(130, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MissingBinariesException() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(131, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMissingBinariesException(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(131, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SetCurrentDir() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(132, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetSetCurrentDir(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(132, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) GlobalContextInitialized() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(133, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) ChromeMajorVer() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(134, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeMinorVer() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(135, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeRelease() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(136, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeBuild() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(137, m.Instance())
	return uint16(r)
}

func (m *TCefApplicationCore) ChromeVersion() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(138, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) LibCefVersion() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(139, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) LibCefPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(140, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) ChromeElfPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(141, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) LibLoaded() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(142, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) LogProcessInfo() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(143, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetLogProcessInfo(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(143, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ReRaiseExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(144, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetReRaiseExceptions(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(144, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) DeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(145, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) LocalesRequired() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(146, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) SetLocalesRequired(value string) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(146, 1, m.Instance(), api.PasStr(value))
}

func (m *TCefApplicationCore) ProcessType() cefTypes.TCefProcessType {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(147, m.Instance())
	return cefTypes.TCefProcessType(r)
}

func (m *TCefApplicationCore) MustCreateResourceBundleHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(148, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateResourceBundleHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(148, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateBrowserProcessHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(149, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateBrowserProcessHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(149, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateRenderProcessHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(150, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateRenderProcessHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(150, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) MustCreateLoadHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(151, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustCreateLoadHandler(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(151, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) SetOsmodalLoop(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(152, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) Status() cefTypes.TCefAplicationStatus {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(153, m.Instance())
	return cefTypes.TCefAplicationStatus(r)
}

func (m *TCefApplicationCore) MissingLibFiles() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(154, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) MustFreeLibrary() bool {
	if !m.IsValid() {
		return false
	}
	r := cefApplicationCoreAPI().SysCallN(155, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCefApplicationCore) SetMustFreeLibrary(value bool) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(155, 1, m.Instance(), api.PasBool(value))
}

func (m *TCefApplicationCore) ChildProcessesCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(156, m.Instance())
	return int32(r)
}

func (m *TCefApplicationCore) UsedMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(157, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) TotalSystemMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(158, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) AvailableSystemMemory() (result uint64) {
	if !m.IsValid() {
		return
	}
	cefApplicationCoreAPI().SysCallN(159, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefApplicationCore) SystemMemoryLoad() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(160, m.Instance())
	return uint32(r)
}

func (m *TCefApplicationCore) ApiHashUniversal() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(161, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) ApiHashPlatform() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(162, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) ApiHashCommit() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(163, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) LastErrorMessage() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefApplicationCoreAPI().SysCallN(164, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefApplicationCore) XDisplay() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cefApplicationCoreAPI().SysCallN(165, m.Instance())
	return uintptr(r)
}

func (m *TCefApplicationCore) SetOnRegCustomSchemes(fn TOnRegisterCustomSchemesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRegisterCustomSchemesEvent(fn)
	base.SetEvent(m, 166, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnRegisterCustomPreferences(fn TOnRegisterCustomPreferencesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRegisterCustomPreferencesEvent(fn)
	base.SetEvent(m, 167, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextInitialized(fn TOnContextInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextInitializedEvent(fn)
	base.SetEvent(m, 168, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBeforeChildProcessLaunch(fn TOnBeforeChildProcessLaunchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeChildProcessLaunchEvent(fn)
	base.SetEvent(m, 169, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnScheduleMessagePumpWork(fn TOnScheduleMessagePumpWorkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScheduleMessagePumpWorkEvent(fn)
	base.SetEvent(m, 170, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDefaultClient(fn TOnGetDefaultClientEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDefaultClientEvent(fn)
	base.SetEvent(m, 171, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetLocalizedString(fn TOnGetLocalizedStringEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetLocalizedStringEvent(fn)
	base.SetEvent(m, 172, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDataResource(fn TOnGetDataResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDataResourceEvent(fn)
	base.SetEvent(m, 173, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnGetDataResourceForScale(fn TOnGetDataResourceForScaleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDataResourceForScaleEvent(fn)
	base.SetEvent(m, 174, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnWebKitInitialized(fn TOnWebKitInitializedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebKitInitializedEvent(fn)
	base.SetEvent(m, 175, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBrowserCreated(fn TOnBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserCreatedEvent(fn)
	base.SetEvent(m, 176, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserDestroyedEvent(fn)
	base.SetEvent(m, 177, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextCreated(fn TOnContextCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextCreatedEvent(fn)
	base.SetEvent(m, 178, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnContextReleased(fn TOnContextReleasedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextReleasedEvent(fn)
	base.SetEvent(m, 179, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnUncaughtException(fn TOnUncaughtExceptionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUncaughtExceptionEvent(fn)
	base.SetEvent(m, 180, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnFocusedNodeChanged(fn TOnFocusedNodeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFocusedNodeChangedEvent(fn)
	base.SetEvent(m, 181, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessMessageReceivedEvent(fn)
	base.SetEvent(m, 182, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadingStateChange(fn TOnRenderLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadingStateChange(fn)
	base.SetEvent(m, 183, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadStart(fn TOnRenderLoadStart) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadStart(fn)
	base.SetEvent(m, 184, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadEnd(fn TOnRenderLoadEnd) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadEnd(fn)
	base.SetEvent(m, 185, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCefApplicationCore) SetOnLoadError(fn TOnRenderLoadError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderLoadError(fn)
	base.SetEvent(m, 186, cefApplicationCoreAPI(), api.MakeEventDataPtr(cb))
}

// NewApplicationCore class constructor
func NewApplicationCore() ICefApplicationCore {
	r := cefApplicationCoreAPI().SysCallN(0)
	return AsCefApplicationCore(r)
}

var (
	cefApplicationCoreOnce   base.Once
	cefApplicationCoreImport *imports.Imports = nil
)

func cefApplicationCoreAPI() *imports.Imports {
	cefApplicationCoreOnce.Do(func() {
		cefApplicationCoreImport = api.NewDefaultImports()
		cefApplicationCoreImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefApplicationCore_Create", 0), // constructor NewApplicationCore
			/* 1 */ imports.NewTable("TCefApplicationCore_CheckCEFLibrary", 0), // function CheckCEFLibrary
			/* 2 */ imports.NewTable("TCefApplicationCore_StartMainProcess", 0), // function StartMainProcess
			/* 3 */ imports.NewTable("TCefApplicationCore_StartSubProcess", 0), // function StartSubProcess
			/* 4 */ imports.NewTable("TCefApplicationCore_Internal_GetLocalizedString", 0), // function InternalGetLocalizedString
			/* 5 */ imports.NewTable("TCefApplicationCore_Internal_GetDataResource", 0), // function InternalGetDataResource
			/* 6 */ imports.NewTable("TCefApplicationCore_Internal_GetDataResourceForScale", 0), // function InternalGetDataResourceForScale
			/* 7 */ imports.NewTable("TCefApplicationCore_AfterConstruction", 0), // procedure AfterConstruction
			/* 8 */ imports.NewTable("TCefApplicationCore_AddCustomCommandLine", 0), // procedure AddCustomCommandLine
			/* 9 */ imports.NewTable("TCefApplicationCore_DoMessageLoopWork", 0), // procedure DoMessageLoopWork
			/* 10 */ imports.NewTable("TCefApplicationCore_RunMessageLoop", 0), // procedure RunMessageLoop
			/* 11 */ imports.NewTable("TCefApplicationCore_QuitMessageLoop", 0), // procedure QuitMessageLoop
			/* 12 */ imports.NewTable("TCefApplicationCore_UpdateDeviceScaleFactor", 0), // procedure UpdateDeviceScaleFactor
			/* 13 */ imports.NewTable("TCefApplicationCore_InitLibLocationFromArgs", 0), // procedure InitLibLocationFromArgs
			/* 14 */ imports.NewTable("TCefApplicationCore_Internal_OnBeforeCommandLineProcessing", 0), // procedure InternalOnBeforeCommandLineProcessing
			/* 15 */ imports.NewTable("TCefApplicationCore_Internal_OnRegisterCustomSchemes", 0), // procedure InternalOnRegisterCustomSchemes
			/* 16 */ imports.NewTable("TCefApplicationCore_Internal_OnRegisterCustomPreferences", 0), // procedure InternalOnRegisterCustomPreferences
			/* 17 */ imports.NewTable("TCefApplicationCore_Internal_OnContextInitialized", 0), // procedure InternalOnContextInitialized
			/* 18 */ imports.NewTable("TCefApplicationCore_Internal_OnBeforeChildProcessLaunch", 0), // procedure InternalOnBeforeChildProcessLaunch
			/* 19 */ imports.NewTable("TCefApplicationCore_Internal_OnScheduleMessagePumpWork", 0), // procedure InternalOnScheduleMessagePumpWork
			/* 20 */ imports.NewTable("TCefApplicationCore_Internal_OnWebKitInitialized", 0), // procedure InternalOnWebKitInitialized
			/* 21 */ imports.NewTable("TCefApplicationCore_Internal_OnBrowserCreated", 0), // procedure InternalOnBrowserCreated
			/* 22 */ imports.NewTable("TCefApplicationCore_Internal_OnBrowserDestroyed", 0), // procedure InternalOnBrowserDestroyed
			/* 23 */ imports.NewTable("TCefApplicationCore_Internal_OnContextCreated", 0), // procedure InternalOnContextCreated
			/* 24 */ imports.NewTable("TCefApplicationCore_Internal_OnContextReleased", 0), // procedure InternalOnContextReleased
			/* 25 */ imports.NewTable("TCefApplicationCore_Internal_OnUncaughtException", 0), // procedure InternalOnUncaughtException
			/* 26 */ imports.NewTable("TCefApplicationCore_Internal_OnFocusedNodeChanged", 0), // procedure InternalOnFocusedNodeChanged
			/* 27 */ imports.NewTable("TCefApplicationCore_Internal_OnProcessMessageReceived", 0), // procedure InternalOnProcessMessageReceived
			/* 28 */ imports.NewTable("TCefApplicationCore_Internal_OnLoadingStateChange", 0), // procedure InternalOnLoadingStateChange
			/* 29 */ imports.NewTable("TCefApplicationCore_Internal_OnLoadStart", 0), // procedure InternalOnLoadStart
			/* 30 */ imports.NewTable("TCefApplicationCore_Internal_OnLoadEnd", 0), // procedure InternalOnLoadEnd
			/* 31 */ imports.NewTable("TCefApplicationCore_Internal_OnLoadError", 0), // procedure InternalOnLoadError
			/* 32 */ imports.NewTable("TCefApplicationCore_Internal_GetDefaultClient", 0), // procedure InternalGetDefaultClient
			/* 33 */ imports.NewTable("TCefApplicationCore_NoSandbox", 0), // property NoSandbox
			/* 34 */ imports.NewTable("TCefApplicationCore_BrowserSubprocessPath", 0), // property BrowserSubprocessPath
			/* 35 */ imports.NewTable("TCefApplicationCore_FrameworkDirPath", 0), // property FrameworkDirPath
			/* 36 */ imports.NewTable("TCefApplicationCore_MainBundlePath", 0), // property MainBundlePath
			/* 37 */ imports.NewTable("TCefApplicationCore_ChromeRuntime", 0), // property ChromeRuntime
			/* 38 */ imports.NewTable("TCefApplicationCore_MultiThreadedMessageLoop", 0), // property MultiThreadedMessageLoop
			/* 39 */ imports.NewTable("TCefApplicationCore_ExternalMessagePump", 0), // property ExternalMessagePump
			/* 40 */ imports.NewTable("TCefApplicationCore_WindowlessRenderingEnabled", 0), // property WindowlessRenderingEnabled
			/* 41 */ imports.NewTable("TCefApplicationCore_CommandLineArgsDisabled", 0), // property CommandLineArgsDisabled
			/* 42 */ imports.NewTable("TCefApplicationCore_Cache", 0), // property Cache
			/* 43 */ imports.NewTable("TCefApplicationCore_RootCache", 0), // property RootCache
			/* 44 */ imports.NewTable("TCefApplicationCore_UserDataPath", 0), // property UserDataPath
			/* 45 */ imports.NewTable("TCefApplicationCore_PersistSessionCookies", 0), // property PersistSessionCookies
			/* 46 */ imports.NewTable("TCefApplicationCore_PersistUserPreferences", 0), // property PersistUserPreferences
			/* 47 */ imports.NewTable("TCefApplicationCore_UserAgent", 0), // property UserAgent
			/* 48 */ imports.NewTable("TCefApplicationCore_UserAgentProduct", 0), // property UserAgentProduct
			/* 49 */ imports.NewTable("TCefApplicationCore_Locale", 0), // property Locale
			/* 50 */ imports.NewTable("TCefApplicationCore_LogFile", 0), // property LogFile
			/* 51 */ imports.NewTable("TCefApplicationCore_LogSeverity", 0), // property LogSeverity
			/* 52 */ imports.NewTable("TCefApplicationCore_JavaScriptFlags", 0), // property JavaScriptFlags
			/* 53 */ imports.NewTable("TCefApplicationCore_ResourcesDirPath", 0), // property ResourcesDirPath
			/* 54 */ imports.NewTable("TCefApplicationCore_LocalesDirPath", 0), // property LocalesDirPath
			/* 55 */ imports.NewTable("TCefApplicationCore_PackLoadingDisabled", 0), // property PackLoadingDisabled
			/* 56 */ imports.NewTable("TCefApplicationCore_RemoteDebuggingPort", 0), // property RemoteDebuggingPort
			/* 57 */ imports.NewTable("TCefApplicationCore_UncaughtExceptionStackSize", 0), // property UncaughtExceptionStackSize
			/* 58 */ imports.NewTable("TCefApplicationCore_IgnoreCertificateErrors", 0), // property IgnoreCertificateErrors
			/* 59 */ imports.NewTable("TCefApplicationCore_BackgroundColor", 0), // property BackgroundColor
			/* 60 */ imports.NewTable("TCefApplicationCore_AcceptLanguageList", 0), // property AcceptLanguageList
			/* 61 */ imports.NewTable("TCefApplicationCore_CookieableSchemesList", 0), // property CookieableSchemesList
			/* 62 */ imports.NewTable("TCefApplicationCore_CookieableSchemesExcludeDefaults", 0), // property CookieableSchemesExcludeDefaults
			/* 63 */ imports.NewTable("TCefApplicationCore_SingleProcess", 0), // property SingleProcess
			/* 64 */ imports.NewTable("TCefApplicationCore_EnableMediaStream", 0), // property EnableMediaStream
			/* 65 */ imports.NewTable("TCefApplicationCore_EnableSpeechInput", 0), // property EnableSpeechInput
			/* 66 */ imports.NewTable("TCefApplicationCore_UseFakeUIForMediaStream", 0), // property UseFakeUIForMediaStream
			/* 67 */ imports.NewTable("TCefApplicationCore_EnableUsermediaScreenCapturing", 0), // property EnableUsermediaScreenCapturing
			/* 68 */ imports.NewTable("TCefApplicationCore_EnableGPU", 0), // property EnableGPU
			/* 69 */ imports.NewTable("TCefApplicationCore_EnableFeatures", 0), // property EnableFeatures
			/* 70 */ imports.NewTable("TCefApplicationCore_DisableFeatures", 0), // property DisableFeatures
			/* 71 */ imports.NewTable("TCefApplicationCore_EnableBlinkFeatures", 0), // property EnableBlinkFeatures
			/* 72 */ imports.NewTable("TCefApplicationCore_DisableBlinkFeatures", 0), // property DisableBlinkFeatures
			/* 73 */ imports.NewTable("TCefApplicationCore_BlinkSettings", 0), // property BlinkSettings
			/* 74 */ imports.NewTable("TCefApplicationCore_ForceFieldTrials", 0), // property ForceFieldTrials
			/* 75 */ imports.NewTable("TCefApplicationCore_ForceFieldTrialParams", 0), // property ForceFieldTrialParams
			/* 76 */ imports.NewTable("TCefApplicationCore_SmoothScrolling", 0), // property SmoothScrolling
			/* 77 */ imports.NewTable("TCefApplicationCore_FastUnload", 0), // property FastUnload
			/* 78 */ imports.NewTable("TCefApplicationCore_DisableSafeBrowsing", 0), // property DisableSafeBrowsing
			/* 79 */ imports.NewTable("TCefApplicationCore_MuteAudio", 0), // property MuteAudio
			/* 80 */ imports.NewTable("TCefApplicationCore_SitePerProcess", 0), // property SitePerProcess
			/* 81 */ imports.NewTable("TCefApplicationCore_DisableWebSecurity", 0), // property DisableWebSecurity
			/* 82 */ imports.NewTable("TCefApplicationCore_DisablePDFExtension", 0), // property DisablePDFExtension
			/* 83 */ imports.NewTable("TCefApplicationCore_DisableSiteIsolationTrials", 0), // property DisableSiteIsolationTrials
			/* 84 */ imports.NewTable("TCefApplicationCore_DisableChromeLoginPrompt", 0), // property DisableChromeLoginPrompt
			/* 85 */ imports.NewTable("TCefApplicationCore_DisableExtensions", 0), // property DisableExtensions
			/* 86 */ imports.NewTable("TCefApplicationCore_AutoplayPolicy", 0), // property AutoplayPolicy
			/* 87 */ imports.NewTable("TCefApplicationCore_DisableBackgroundNetworking", 0), // property DisableBackgroundNetworking
			/* 88 */ imports.NewTable("TCefApplicationCore_MetricsRecordingOnly", 0), // property MetricsRecordingOnly
			/* 89 */ imports.NewTable("TCefApplicationCore_AllowFileAccessFromFiles", 0), // property AllowFileAccessFromFiles
			/* 90 */ imports.NewTable("TCefApplicationCore_AllowRunningInsecureContent", 0), // property AllowRunningInsecureContent
			/* 91 */ imports.NewTable("TCefApplicationCore_EnablePrintPreview", 0), // property EnablePrintPreview
			/* 92 */ imports.NewTable("TCefApplicationCore_DefaultEncoding", 0), // property DefaultEncoding
			/* 93 */ imports.NewTable("TCefApplicationCore_DisableJavascript", 0), // property DisableJavascript
			/* 94 */ imports.NewTable("TCefApplicationCore_DisableJavascriptCloseWindows", 0), // property DisableJavascriptCloseWindows
			/* 95 */ imports.NewTable("TCefApplicationCore_DisableJavascriptAccessClipboard", 0), // property DisableJavascriptAccessClipboard
			/* 96 */ imports.NewTable("TCefApplicationCore_DisableJavascriptDomPaste", 0), // property DisableJavascriptDomPaste
			/* 97 */ imports.NewTable("TCefApplicationCore_AllowUniversalAccessFromFileUrls", 0), // property AllowUniversalAccessFromFileUrls
			/* 98 */ imports.NewTable("TCefApplicationCore_DisableImageLoading", 0), // property DisableImageLoading
			/* 99 */ imports.NewTable("TCefApplicationCore_ImageShrinkStandaloneToFit", 0), // property ImageShrinkStandaloneToFit
			/* 100 */ imports.NewTable("TCefApplicationCore_DisableTextAreaResize", 0), // property DisableTextAreaResize
			/* 101 */ imports.NewTable("TCefApplicationCore_DisableTabToLinks", 0), // property DisableTabToLinks
			/* 102 */ imports.NewTable("TCefApplicationCore_EnableProfanityFilter", 0), // property EnableProfanityFilter
			/* 103 */ imports.NewTable("TCefApplicationCore_DisableSpellChecking", 0), // property DisableSpellChecking
			/* 104 */ imports.NewTable("TCefApplicationCore_OverrideSpellCheckLang", 0), // property OverrideSpellCheckLang
			/* 105 */ imports.NewTable("TCefApplicationCore_TouchEvents", 0), // property TouchEvents
			/* 106 */ imports.NewTable("TCefApplicationCore_DisableReadingFromCanvas", 0), // property DisableReadingFromCanvas
			/* 107 */ imports.NewTable("TCefApplicationCore_HyperlinkAuditing", 0), // property HyperlinkAuditing
			/* 108 */ imports.NewTable("TCefApplicationCore_DisableNewBrowserInfoTimeout", 0), // property DisableNewBrowserInfoTimeout
			/* 109 */ imports.NewTable("TCefApplicationCore_DevToolsProtocolLogFile", 0), // property DevToolsProtocolLogFile
			/* 110 */ imports.NewTable("TCefApplicationCore_ForcedDeviceScaleFactor", 0), // property ForcedDeviceScaleFactor
			/* 111 */ imports.NewTable("TCefApplicationCore_DisableZygote", 0), // property DisableZygote
			/* 112 */ imports.NewTable("TCefApplicationCore_UseMockKeyChain", 0), // property UseMockKeyChain
			/* 113 */ imports.NewTable("TCefApplicationCore_DisableRequestHandlingForTesting", 0), // property DisableRequestHandlingForTesting
			/* 114 */ imports.NewTable("TCefApplicationCore_DisablePopupBlocking", 0), // property DisablePopupBlocking
			/* 115 */ imports.NewTable("TCefApplicationCore_DisableBackForwardCache", 0), // property DisableBackForwardCache
			/* 116 */ imports.NewTable("TCefApplicationCore_DisableComponentUpdate", 0), // property DisableComponentUpdate
			/* 117 */ imports.NewTable("TCefApplicationCore_AllowInsecureLocalhost", 0), // property AllowInsecureLocalhost
			/* 118 */ imports.NewTable("TCefApplicationCore_KioskPrinting", 0), // property KioskPrinting
			/* 119 */ imports.NewTable("TCefApplicationCore_TreatInsecureOriginAsSecure", 0), // property TreatInsecureOriginAsSecure
			/* 120 */ imports.NewTable("TCefApplicationCore_NetLogEnabled", 0), // property NetLogEnabled
			/* 121 */ imports.NewTable("TCefApplicationCore_NetLogFile", 0), // property NetLogFile
			/* 122 */ imports.NewTable("TCefApplicationCore_NetLogCaptureMode", 0), // property NetLogCaptureMode
			/* 123 */ imports.NewTable("TCefApplicationCore_WindowsSandboxInfo", 0), // property WindowsSandboxInfo
			/* 124 */ imports.NewTable("TCefApplicationCore_EnableHighDPISupport", 0), // property EnableHighDPISupport
			/* 125 */ imports.NewTable("TCefApplicationCore_argcCopy", 0), // property ArgcCopy
			/* 126 */ imports.NewTable("TCefApplicationCore_argvCopy", 0), // property ArgvCopy
			/* 127 */ imports.NewTable("TCefApplicationCore_DeleteCache", 0), // property DeleteCache
			/* 128 */ imports.NewTable("TCefApplicationCore_DeleteCookies", 0), // property DeleteCookies
			/* 129 */ imports.NewTable("TCefApplicationCore_CheckCEFFiles", 0), // property CheckCEFFiles
			/* 130 */ imports.NewTable("TCefApplicationCore_ShowMessageDlg", 0), // property ShowMessageDlg
			/* 131 */ imports.NewTable("TCefApplicationCore_MissingBinariesException", 0), // property MissingBinariesException
			/* 132 */ imports.NewTable("TCefApplicationCore_SetCurrentDir", 0), // property SetCurrentDir
			/* 133 */ imports.NewTable("TCefApplicationCore_GlobalContextInitialized", 0), // property GlobalContextInitialized
			/* 134 */ imports.NewTable("TCefApplicationCore_ChromeMajorVer", 0), // property ChromeMajorVer
			/* 135 */ imports.NewTable("TCefApplicationCore_ChromeMinorVer", 0), // property ChromeMinorVer
			/* 136 */ imports.NewTable("TCefApplicationCore_ChromeRelease", 0), // property ChromeRelease
			/* 137 */ imports.NewTable("TCefApplicationCore_ChromeBuild", 0), // property ChromeBuild
			/* 138 */ imports.NewTable("TCefApplicationCore_ChromeVersion", 0), // property ChromeVersion
			/* 139 */ imports.NewTable("TCefApplicationCore_LibCefVersion", 0), // property LibCefVersion
			/* 140 */ imports.NewTable("TCefApplicationCore_LibCefPath", 0), // property LibCefPath
			/* 141 */ imports.NewTable("TCefApplicationCore_ChromeElfPath", 0), // property ChromeElfPath
			/* 142 */ imports.NewTable("TCefApplicationCore_LibLoaded", 0), // property LibLoaded
			/* 143 */ imports.NewTable("TCefApplicationCore_LogProcessInfo", 0), // property LogProcessInfo
			/* 144 */ imports.NewTable("TCefApplicationCore_ReRaiseExceptions", 0), // property ReRaiseExceptions
			/* 145 */ imports.NewTable("TCefApplicationCore_DeviceScaleFactor", 0), // property DeviceScaleFactor
			/* 146 */ imports.NewTable("TCefApplicationCore_LocalesRequired", 0), // property LocalesRequired
			/* 147 */ imports.NewTable("TCefApplicationCore_ProcessType", 0), // property ProcessType
			/* 148 */ imports.NewTable("TCefApplicationCore_MustCreateResourceBundleHandler", 0), // property MustCreateResourceBundleHandler
			/* 149 */ imports.NewTable("TCefApplicationCore_MustCreateBrowserProcessHandler", 0), // property MustCreateBrowserProcessHandler
			/* 150 */ imports.NewTable("TCefApplicationCore_MustCreateRenderProcessHandler", 0), // property MustCreateRenderProcessHandler
			/* 151 */ imports.NewTable("TCefApplicationCore_MustCreateLoadHandler", 0), // property MustCreateLoadHandler
			/* 152 */ imports.NewTable("TCefApplicationCore_OsmodalLoop", 0), // property OsmodalLoop
			/* 153 */ imports.NewTable("TCefApplicationCore_Status", 0), // property Status
			/* 154 */ imports.NewTable("TCefApplicationCore_MissingLibFiles", 0), // property MissingLibFiles
			/* 155 */ imports.NewTable("TCefApplicationCore_MustFreeLibrary", 0), // property MustFreeLibrary
			/* 156 */ imports.NewTable("TCefApplicationCore_ChildProcessesCount", 0), // property ChildProcessesCount
			/* 157 */ imports.NewTable("TCefApplicationCore_UsedMemory", 0), // property UsedMemory
			/* 158 */ imports.NewTable("TCefApplicationCore_TotalSystemMemory", 0), // property TotalSystemMemory
			/* 159 */ imports.NewTable("TCefApplicationCore_AvailableSystemMemory", 0), // property AvailableSystemMemory
			/* 160 */ imports.NewTable("TCefApplicationCore_SystemMemoryLoad", 0), // property SystemMemoryLoad
			/* 161 */ imports.NewTable("TCefApplicationCore_ApiHashUniversal", 0), // property ApiHashUniversal
			/* 162 */ imports.NewTable("TCefApplicationCore_ApiHashPlatform", 0), // property ApiHashPlatform
			/* 163 */ imports.NewTable("TCefApplicationCore_ApiHashCommit", 0), // property ApiHashCommit
			/* 164 */ imports.NewTable("TCefApplicationCore_LastErrorMessage", 0), // property LastErrorMessage
			/* 165 */ imports.NewTable("TCefApplicationCore_XDisplay", 0), // property XDisplay
			/* 166 */ imports.NewTable("TCefApplicationCore_OnRegCustomSchemes", 0), // event OnRegCustomSchemes
			/* 167 */ imports.NewTable("TCefApplicationCore_OnRegisterCustomPreferences", 0), // event OnRegisterCustomPreferences
			/* 168 */ imports.NewTable("TCefApplicationCore_OnContextInitialized", 0), // event OnContextInitialized
			/* 169 */ imports.NewTable("TCefApplicationCore_OnBeforeChildProcessLaunch", 0), // event OnBeforeChildProcessLaunch
			/* 170 */ imports.NewTable("TCefApplicationCore_OnScheduleMessagePumpWork", 0), // event OnScheduleMessagePumpWork
			/* 171 */ imports.NewTable("TCefApplicationCore_OnGetDefaultClient", 0), // event OnGetDefaultClient
			/* 172 */ imports.NewTable("TCefApplicationCore_OnGetLocalizedString", 0), // event OnGetLocalizedString
			/* 173 */ imports.NewTable("TCefApplicationCore_OnGetDataResource", 0), // event OnGetDataResource
			/* 174 */ imports.NewTable("TCefApplicationCore_OnGetDataResourceForScale", 0), // event OnGetDataResourceForScale
			/* 175 */ imports.NewTable("TCefApplicationCore_OnWebKitInitialized", 0), // event OnWebKitInitialized
			/* 176 */ imports.NewTable("TCefApplicationCore_OnBrowserCreated", 0), // event OnBrowserCreated
			/* 177 */ imports.NewTable("TCefApplicationCore_OnBrowserDestroyed", 0), // event OnBrowserDestroyed
			/* 178 */ imports.NewTable("TCefApplicationCore_OnContextCreated", 0), // event OnContextCreated
			/* 179 */ imports.NewTable("TCefApplicationCore_OnContextReleased", 0), // event OnContextReleased
			/* 180 */ imports.NewTable("TCefApplicationCore_OnUncaughtException", 0), // event OnUncaughtException
			/* 181 */ imports.NewTable("TCefApplicationCore_OnFocusedNodeChanged", 0), // event OnFocusedNodeChanged
			/* 182 */ imports.NewTable("TCefApplicationCore_OnProcessMessageReceived", 0), // event OnProcessMessageReceived
			/* 183 */ imports.NewTable("TCefApplicationCore_OnLoadingStateChange", 0), // event OnLoadingStateChange
			/* 184 */ imports.NewTable("TCefApplicationCore_OnLoadStart", 0), // event OnLoadStart
			/* 185 */ imports.NewTable("TCefApplicationCore_OnLoadEnd", 0), // event OnLoadEnd
			/* 186 */ imports.NewTable("TCefApplicationCore_OnLoadError", 0), // event OnLoadError
		}
	})
	return cefApplicationCoreImport
}
