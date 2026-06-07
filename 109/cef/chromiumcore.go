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
	"github.com/energye/lcl/types"

	cefTypes "github.com/energye/cef/109/types"
)

// IChromiumEvents Parent: IObject
type IChromiumEvents interface {
	IObject
}

// IChromiumCore Parent: IChromiumEvents IComponent
type IChromiumCore interface {
	IChromiumEvents
	IComponent
	CreateClientHandlerWithBool(isOSR bool) bool                                                                                                                                                                            // function
	CreateClientHandlerWithClientBool(client *IEngClient, isOSR bool) bool                                                                                                                                                  // function
	TryCloseBrowser() bool                                                                                                                                                                                                  // function
	SelectBrowser(iD int32) bool                                                                                                                                                                                            // function
	IndexOfBrowserID(iD int32) int32                                                                                                                                                                                        // function
	ShareRequestContext(context *ICefRequestContext, handler IEngRequestContextHandler) bool                                                                                                                                // function
	SetNewBrowserParent(newParentHwnd types.HWND) bool                                                                                                                                                                      // function
	CreateBrowserWithWHandleRectStrRContextDValueBool(parentHandle cefTypes.TCefWindowHandle, parentRect types.TRect, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue, forceAsPopup bool) bool // function
	CreateBrowserWithStrBVComponentRContextDValue(uRL string, browserViewComp ICEFBrowserViewComponent, context ICefRequestContext, extraInfo ICefDictionaryValue) bool                                                     // function
	ClearCertificateExceptions(clearImmediately bool) bool                                                                                                                                                                  // function
	ClearHttpAuthCredentials(clearImmediately bool) bool                                                                                                                                                                    // function
	CloseAllConnections(closeImmediately bool) bool                                                                                                                                                                         // function
	GetFrameNames(frameNames *lcl.IStrings) bool                                                                                                                                                                            // function
	GetFrameIdentifiers(frameCount *cefTypes.NativeUInt, frameIdentifierArray *ICefFrameIdentifierArray) bool                                                                                                               // function
	IsSameBrowser(browser ICefBrowser) bool                                                                                                                                                                                 // function
	ExecuteTaskOnCefThread(cefThreadId cefTypes.TCefThreadId, taskID uint32, delayMs int64) bool                                                                                                                            // function
	DeleteCookies(url string, cookieName string, deleteImmediately bool) bool                                                                                                                                               // function
	VisitAllCookies(iD int32) bool                                                                                                                                                                                          // function
	VisitURLCookies(url string, includeHttpOnly bool, iD int32) bool                                                                                                                                                        // function
	SetCookie(args TChromiumCoreSetCookieArgs) bool                                                                                                                                                                         // function
	FlushCookieStore(flushImmediately bool) bool                                                                                                                                                                            // function
	SendDevToolsMessage(message string) bool                                                                                                                                                                                // function
	ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32                                                                                                                                 // function
	AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) ICefRegistration                                                                                                                                       // function
	CreateUrlRequestWithRequestUClientStr(request ICefRequest, client IEngUrlrequestClient, frameName string) ICefUrlRequest                                                                                                // function
	CreateUrlRequestWithRequestUClientFrame(request ICefRequest, client IEngUrlrequestClient, frame ICefFrame) ICefUrlRequest                                                                                               // function
	CreateUrlRequestWithRequestUClientInt64(request ICefRequest, client IEngUrlrequestClient, frameIdentifier int64) ICefUrlRequest                                                                                         // function
	// AddObserver
	//  ICefMediaRouter methods
	AddObserver(observer IEngMediaObserver) ICefRegistration // function
	GetSource(urn string) ICefMediaSource                    // function
	// LoadExtension
	//  ICefRequestContext methods for extensions
	LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler, requestContext ICefRequestContext) bool // function
	DidLoadExtension(extensionId string) bool                                                                                               // function
	HasExtension(extensionId string) bool                                                                                                   // function
	GetExtensions(extensionIds lcl.IStringList) bool                                                                                        // function
	GetExtension(extensionId string) ICefExtension                                                                                          // function
	CloseBrowser(forceClose bool)                                                                                                           // procedure
	CloseAllBrowsers()                                                                                                                      // procedure
	InitializeDragAndDrop(dropTargetWnd types.HWND)                                                                                         // procedure
	ShutdownDragAndDrop()                                                                                                                   // procedure
	LoadURLWithStrX2(uRL string, frameName string)                                                                                          // procedure
	LoadURLWithStrFrame(uRL string, frame ICefFrame)                                                                                        // procedure
	LoadURLWithStrInt64(uRL string, frameIdentifier int64)                                                                                  // procedure
	LoadStringWithStrX2(hTML string, frameName string)                                                                                      // procedure
	LoadStringWithStrFrame(hTML string, frame ICefFrame)                                                                                    // procedure
	LoadStringWithStrInt64(hTML string, frameIdentifier int64)                                                                              // procedure
	LoadResourceWithCMStreamStrX3(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameName string)                        // procedure
	LoadResourceWithCMStreamStrX2Frame(stream lcl.ICustomMemoryStream, mimeType string, charset string, frame ICefFrame)                    // procedure
	LoadResourceWithCMStreamStrX2Int64(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameIdentifier int64)              // procedure
	LoadRequest(request ICefRequest)                                                                                                        // procedure
	GoBack()                                                                                                                                // procedure
	GoForward()                                                                                                                             // procedure
	Reload()                                                                                                                                // procedure
	ReloadIgnoreCache()                                                                                                                     // procedure
	StopLoad()                                                                                                                              // procedure
	StartDownload(uRL string)                                                                                                               // procedure
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool)                                                   // procedure
	SimulateMouseWheel(deltaX int32, deltaY int32)                                                                                          // procedure
	RetrieveHTMLWithStr(frameName string)                                                                                                   // procedure
	RetrieveHTMLWithFrame(frame ICefFrame)                                                                                                  // procedure
	RetrieveHTMLWithInt64(frameIdentifier int64)                                                                                            // procedure
	RetrieveTextWithStr(frameName string)                                                                                                   // procedure
	RetrieveTextWithFrame(frame ICefFrame)                                                                                                  // procedure
	RetrieveTextWithInt64(frameIdentifier int64)                                                                                            // procedure
	GetNavigationEntries(currentOnly bool)                                                                                                  // procedure
	ExecuteJavaScriptWithStrX3Int(code string, scriptURL string, frameName string, startLine int32)                                         // procedure
	ExecuteJavaScriptWithStrX2FrameInt(code string, scriptURL string, frame ICefFrame, startLine int32)                                     // procedure
	ExecuteJavaScriptWithStrX2Int64Int(code string, scriptURL string, frameIdentifier int64, startLine int32)                               // procedure
	UpdatePreferences()                                                                                                                     // procedure
	SavePreferences(fileName string)                                                                                                        // procedure
	ResolveHost(uRL string)                                                                                                                 // procedure
	SetUserAgentOverride(userAgent string, acceptLanguage string, platform string)                                                          // procedure
	ClearDataForOrigin(origin string, storageTypes cefTypes.TCefClearDataStorageTypes)                                                      // procedure
	ClearCache()                                                                                                                            // procedure
	ToggleAudioMuted()                                                                                                                      // procedure
	ShowDevTools(inspectElementAt types.TPoint, windowInfo TCefWindowInfo)                                                                  // procedure
	CloseDevTools()                                                                                                                         // procedure
	CloseDevToolsWithWindowHandle(devToolsWnd cefTypes.TCefWindowHandle)                                                                    // procedure
	Find(searchText string, forward bool, matchCase bool, findNext bool)                                                                    // procedure
	StopFinding(clearSelection bool)                                                                                                        // procedure
	Print()                                                                                                                                 // procedure
	PrintToPDF(filePath string)                                                                                                             // procedure
	ClipboardCopy()                                                                                                                         // procedure
	ClipboardPaste()                                                                                                                        // procedure
	ClipboardCut()                                                                                                                          // procedure
	ClipboardUndo()                                                                                                                         // procedure
	ClipboardRedo()                                                                                                                         // procedure
	ClipboardDel()                                                                                                                          // procedure
	SelectAll()                                                                                                                             // procedure
	IncZoomStep()                                                                                                                           // procedure
	DecZoomStep()                                                                                                                           // procedure
	IncZoomPct()                                                                                                                            // procedure
	DecZoomPct()                                                                                                                            // procedure
	ResetZoomStep()                                                                                                                         // procedure
	ResetZoomLevel()                                                                                                                        // procedure
	ResetZoomPct()                                                                                                                          // procedure
	ReadZoom()                                                                                                                              // procedure
	WasResized()                                                                                                                            // procedure
	WasHidden(hidden bool)                                                                                                                  // procedure
	NotifyScreenInfoChanged()                                                                                                               // procedure
	NotifyMoveOrResizeStarted()                                                                                                             // procedure
	Invalidate(type_ cefTypes.TCefPaintElementType)                                                                                         // procedure
	SendExternalBeginFrame()                                                                                                                // procedure
	SendKeyEvent(event TCefKeyEvent)                                                                                                        // procedure
	SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32)                           // procedure
	SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool)                                                                               // procedure
	SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32)                                                                   // procedure
	SendTouchEvent(event TCefTouchEvent)                                                                                                    // procedure
	SendCaptureLostEvent()                                                                                                                  // procedure
	SendProcessMessageWithPIdPMessageStr(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameName string)            // procedure
	SendProcessMessageWithPIdPMessageFrame(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frame ICefFrame)           // procedure
	SendProcessMessageWithPIdPMessageInt64(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameIdentifier int64)     // procedure
	SetFocus(focus bool)                                                                                                                    // procedure
	SetAccessibilityState(accessibilityState cefTypes.TCefState)                                                                            // procedure
	DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations)                                // procedure
	DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations)                                                        // procedure
	DragTargetDragLeave()                                                                                                                   // procedure
	DragTargetDrop(event TCefMouseEvent)                                                                                                    // procedure
	DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation)                                                                      // procedure
	DragSourceSystemDragEnded()                                                                                                             // procedure
	IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange)          // procedure
	IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32)                                                         // procedure
	IMEFinishComposingText(keepSelection bool)                                                                                              // procedure
	IMECancelComposition()                                                                                                                  // procedure
	ReplaceMisspelling(word string)                                                                                                         // procedure
	AddWordToDictionary(word string)                                                                                                        // procedure
	UpdateBrowserSize(left int32, top int32, width int32, height int32)                                                                     // procedure
	UpdateXWindowVisibility(visible bool)                                                                                                   // procedure
	NotifyCurrentSinks()                                                                                                                    // procedure
	NotifyCurrentRoutes()                                                                                                                   // procedure
	CreateRoute(source ICefMediaSource, sink ICefMediaSink)                                                                                 // procedure
	GetDeviceInfo(mediaSink ICefMediaSink)                                                                                                  // procedure
	DefaultUrl() string                                                                                                                     // property DefaultUrl Getter
	SetDefaultUrl(value string)                                                                                                             // property DefaultUrl Setter
	Options() IChromiumOptions                                                                                                              // property Options Getter
	SetOptions(value IChromiumOptions)                                                                                                      // property Options Setter
	FontOptions() IChromiumFontOptions                                                                                                      // property FontOptions Getter
	SetFontOptions(value IChromiumFontOptions)                                                                                              // property FontOptions Setter
	PDFPrintOptions() IPDFPrintOptions                                                                                                      // property PDFPrintOptions Getter
	SetPDFPrintOptions(value IPDFPrintOptions)                                                                                              // property PDFPrintOptions Setter
	DefaultEncoding() string                                                                                                                // property DefaultEncoding Getter
	SetDefaultEncoding(value string)                                                                                                        // property DefaultEncoding Setter
	BrowserId() int32                                                                                                                       // property BrowserId Getter
	Browser() ICefBrowser                                                                                                                   // property Browser Getter
	BrowserById(id int32) ICefBrowser                                                                                                       // property BrowserById Getter
	BrowserCount() int32                                                                                                                    // property BrowserCount Getter
	BrowserIdByIndex(I int32) int32                                                                                                         // property BrowserIdByIndex Getter
	CefClient() IEngClient                                                                                                                  // property CefClient Getter
	ReqContextHandler() IEngRequestContextHandler                                                                                           // property ReqContextHandler Getter
	ResourceRequestHandler() IEngResourceRequestHandler                                                                                     // property ResourceRequestHandler Getter
	CefWindowInfo() TCefWindowInfo                                                                                                          // property CefWindowInfo Getter
	VisibleNavigationEntry() ICefNavigationEntry                                                                                            // property VisibleNavigationEntry Getter
	RequestContext() ICefRequestContext                                                                                                     // property RequestContext Getter
	MediaRouter() ICefMediaRouter                                                                                                           // property MediaRouter Getter
	MediaObserver() IEngMediaObserver                                                                                                       // property MediaObserver Getter
	MediaObserverReg() ICefRegistration                                                                                                     // property MediaObserverReg Getter
	DevToolsMsgObserver() IEngDevToolsMessageObserver                                                                                       // property DevToolsMsgObserver Getter
	DevToolsMsgObserverReg() ICefRegistration                                                                                               // property DevToolsMsgObserverReg Getter
	ExtensionHandler() IEngExtensionHandler                                                                                                 // property ExtensionHandler Getter
	MultithreadApp() bool                                                                                                                   // property MultithreadApp Getter
	IsLoading() bool                                                                                                                        // property IsLoading Getter
	HasDocument() bool                                                                                                                      // property HasDocument Getter
	HasView() bool                                                                                                                          // property HasView Getter
	HasDevTools() bool                                                                                                                      // property HasDevTools Getter
	HasClientHandler() bool                                                                                                                 // property HasClientHandler Getter
	HasBrowser() bool                                                                                                                       // property HasBrowser Getter
	CanGoBack() bool                                                                                                                        // property CanGoBack Getter
	CanGoForward() bool                                                                                                                     // property CanGoForward Getter
	IsPopUp() bool                                                                                                                          // property IsPopUp Getter
	WindowHandle() cefTypes.TCefWindowHandle                                                                                                // property WindowHandle Getter
	OpenerWindowHandle() cefTypes.TCefWindowHandle                                                                                          // property OpenerWindowHandle Getter
	BrowserHandle() types.THandle                                                                                                           // property BrowserHandle Getter
	WidgetHandle() types.THandle                                                                                                            // property WidgetHandle Getter
	RenderHandle() types.THandle                                                                                                            // property RenderHandle Getter
	FrameIsFocused() bool                                                                                                                   // property FrameIsFocused Getter
	Initialized() bool                                                                                                                      // property Initialized Getter
	RequestContextCache() string                                                                                                            // property RequestContextCache Getter
	RequestContextIsGlobal() bool                                                                                                           // property RequestContextIsGlobal Getter
	DocumentURL() string                                                                                                                    // property DocumentURL Getter
	ZoomLevel() float64                                                                                                                     // property ZoomLevel Getter
	SetZoomLevel(value float64)                                                                                                             // property ZoomLevel Setter
	ZoomPct() float64                                                                                                                       // property ZoomPct Getter
	SetZoomPct(value float64)                                                                                                               // property ZoomPct Setter
	ZoomStep() byte                                                                                                                         // property ZoomStep Getter
	SetZoomStep(value byte)                                                                                                                 // property ZoomStep Setter
	WindowlessFrameRate() int32                                                                                                             // property WindowlessFrameRate Getter
	SetWindowlessFrameRate(value int32)                                                                                                     // property WindowlessFrameRate Setter
	CustomHeaderName() string                                                                                                               // property CustomHeaderName Getter
	SetCustomHeaderName(value string)                                                                                                       // property CustomHeaderName Setter
	CustomHeaderValue() string                                                                                                              // property CustomHeaderValue Getter
	SetCustomHeaderValue(value string)                                                                                                      // property CustomHeaderValue Setter
	DoNotTrack() bool                                                                                                                       // property DoNotTrack Getter
	SetDoNotTrack(value bool)                                                                                                               // property DoNotTrack Setter
	SendReferrer() bool                                                                                                                     // property SendReferrer Getter
	SetSendReferrer(value bool)                                                                                                             // property SendReferrer Setter
	HyperlinkAuditing() bool                                                                                                                // property HyperlinkAuditing Getter
	SetHyperlinkAuditing(value bool)                                                                                                        // property HyperlinkAuditing Setter
	AllowOutdatedPlugins() bool                                                                                                             // property AllowOutdatedPlugins Getter
	SetAllowOutdatedPlugins(value bool)                                                                                                     // property AllowOutdatedPlugins Setter
	AlwaysAuthorizePlugins() bool                                                                                                           // property AlwaysAuthorizePlugins Getter
	SetAlwaysAuthorizePlugins(value bool)                                                                                                   // property AlwaysAuthorizePlugins Setter
	AlwaysOpenPDFExternally() bool                                                                                                          // property AlwaysOpenPDFExternally Getter
	SetAlwaysOpenPDFExternally(value bool)                                                                                                  // property AlwaysOpenPDFExternally Setter
	SpellChecking() bool                                                                                                                    // property SpellChecking Getter
	SetSpellChecking(value bool)                                                                                                            // property SpellChecking Setter
	SpellCheckerDicts() string                                                                                                              // property SpellCheckerDicts Getter
	SetSpellCheckerDicts(value string)                                                                                                      // property SpellCheckerDicts Setter
	HasValidMainFrame() bool                                                                                                                // property HasValidMainFrame Getter
	FrameCount() cefTypes.NativeUInt                                                                                                        // property FrameCount Getter
	DragOperations() cefTypes.TCefDragOperations                                                                                            // property DragOperations Getter
	SetDragOperations(value cefTypes.TCefDragOperations)                                                                                    // property DragOperations Setter
	AudioMuted() bool                                                                                                                       // property AudioMuted Getter
	SetAudioMuted(value bool)                                                                                                               // property AudioMuted Setter
	SafeSearch() bool                                                                                                                       // property SafeSearch Getter
	SetSafeSearch(value bool)                                                                                                               // property SafeSearch Setter
	YouTubeRestrict() int32                                                                                                                 // property YouTubeRestrict Getter
	SetYouTubeRestrict(value int32)                                                                                                         // property YouTubeRestrict Setter
	PrintingEnabled() bool                                                                                                                  // property PrintingEnabled Getter
	SetPrintingEnabled(value bool)                                                                                                          // property PrintingEnabled Setter
	AcceptLanguageList() string                                                                                                             // property AcceptLanguageList Getter
	SetAcceptLanguageList(value string)                                                                                                     // property AcceptLanguageList Setter
	AcceptCookies() cefTypes.TCefCookiePref                                                                                                 // property AcceptCookies Getter
	SetAcceptCookies(value cefTypes.TCefCookiePref)                                                                                         // property AcceptCookies Setter
	Block3rdPartyCookies() bool                                                                                                             // property Block3rdPartyCookies Getter
	SetBlock3rdPartyCookies(value bool)                                                                                                     // property Block3rdPartyCookies Setter
	MultiBrowserMode() bool                                                                                                                 // property MultiBrowserMode Getter
	SetMultiBrowserMode(value bool)                                                                                                         // property MultiBrowserMode Setter
	DefaultWindowInfoExStyle() types.DWORD                                                                                                  // property DefaultWindowInfoExStyle Getter
	SetDefaultWindowInfoExStyle(value types.DWORD)                                                                                          // property DefaultWindowInfoExStyle Setter
	Offline() bool                                                                                                                          // property Offline Getter
	SetOffline(value bool)                                                                                                                  // property Offline Setter
	QuicAllowed() bool                                                                                                                      // property QuicAllowed Getter
	SetQuicAllowed(value bool)                                                                                                              // property QuicAllowed Setter
	JavascriptEnabled() bool                                                                                                                // property JavascriptEnabled Getter
	SetJavascriptEnabled(value bool)                                                                                                        // property JavascriptEnabled Setter
	LoadImagesAutomatically() bool                                                                                                          // property LoadImagesAutomatically Getter
	SetLoadImagesAutomatically(value bool)                                                                                                  // property LoadImagesAutomatically Setter
	BatterySaverModeState() cefTypes.TCefBatterySaverModeState                                                                              // property BatterySaverModeState Getter
	SetBatterySaverModeState(value cefTypes.TCefBatterySaverModeState)                                                                      // property BatterySaverModeState Setter
	HighEfficiencyMode() cefTypes.TCefState                                                                                                 // property HighEfficiencyMode Getter
	SetHighEfficiencyMode(value cefTypes.TCefState)                                                                                         // property HighEfficiencyMode Setter
	XDisplay() uintptr                                                                                                                      // property XDisplay Getter
	WebRTCIPHandlingPolicy() cefTypes.TCefWebRTCHandlingPolicy                                                                              // property WebRTCIPHandlingPolicy Getter
	SetWebRTCIPHandlingPolicy(value cefTypes.TCefWebRTCHandlingPolicy)                                                                      // property WebRTCIPHandlingPolicy Setter
	WebRTCMultipleRoutes() cefTypes.TCefState                                                                                               // property WebRTCMultipleRoutes Getter
	SetWebRTCMultipleRoutes(value cefTypes.TCefState)                                                                                       // property WebRTCMultipleRoutes Setter
	WebRTCNonproxiedUDP() cefTypes.TCefState                                                                                                // property WebRTCNonproxiedUDP Getter
	SetWebRTCNonproxiedUDP(value cefTypes.TCefState)                                                                                        // property WebRTCNonproxiedUDP Setter
	ProxyType() int32                                                                                                                       // property ProxyType Getter
	SetProxyType(value int32)                                                                                                               // property ProxyType Setter
	ProxyScheme() cefTypes.TCefProxyScheme                                                                                                  // property ProxyScheme Getter
	SetProxyScheme(value cefTypes.TCefProxyScheme)                                                                                          // property ProxyScheme Setter
	ProxyServer() string                                                                                                                    // property ProxyServer Getter
	SetProxyServer(value string)                                                                                                            // property ProxyServer Setter
	ProxyPort() int32                                                                                                                       // property ProxyPort Getter
	SetProxyPort(value int32)                                                                                                               // property ProxyPort Setter
	ProxyUsername() string                                                                                                                  // property ProxyUsername Getter
	SetProxyUsername(value string)                                                                                                          // property ProxyUsername Setter
	ProxyPassword() string                                                                                                                  // property ProxyPassword Getter
	SetProxyPassword(value string)                                                                                                          // property ProxyPassword Setter
	ProxyScriptURL() string                                                                                                                 // property ProxyScriptURL Getter
	SetProxyScriptURL(value string)                                                                                                         // property ProxyScriptURL Setter
	ProxyByPassList() string                                                                                                                // property ProxyByPassList Getter
	SetProxyByPassList(value string)                                                                                                        // property ProxyByPassList Setter
	MaxConnectionsPerProxy() int32                                                                                                          // property MaxConnectionsPerProxy Getter
	SetMaxConnectionsPerProxy(value int32)                                                                                                  // property MaxConnectionsPerProxy Setter
	SetOnTextResultAvailable(fn TOnTextResultAvailableEvent)                                                                                // property event
	SetOnPdfPrintFinished(fn TOnPdfPrintFinishedEvent)                                                                                      // property event
	SetOnPrefsAvailable(fn TOnPrefsAvailableEvent)                                                                                          // property event
	SetOnPrefsUpdated(fn TNotifyEvent)                                                                                                      // property event
	SetOnCookiesDeleted(fn TOnCookiesDeletedEvent)                                                                                          // property event
	SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailableEvent)                                                                             // property event
	SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailableEvent)                                                      // property event
	SetOnDownloadImageFinished(fn TOnDownloadImageFinishedEvent)                                                                            // property event
	SetOnCookiesFlushed(fn TNotifyEvent)                                                                                                    // property event
	SetOnCertificateExceptionsCleared(fn TNotifyEvent)                                                                                      // property event
	SetOnHttpAuthCredentialsCleared(fn TNotifyEvent)                                                                                        // property event
	SetOnAllConnectionsClosed(fn TNotifyEvent)                                                                                              // property event
	SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread)                                                                               // property event
	SetOnCookiesVisited(fn TOnCookiesVisited)                                                                                               // property event
	SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed)                                                                               // property event
	SetOnCookieSet(fn TOnCookieSet)                                                                                                         // property event
	SetOnZoomPctAvailable(fn TOnZoomPctAvailable)                                                                                           // property event
	SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinishedEvent)                                                                      // property event
	SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoEvent)                                                                                // property event
	SetOnBrowserCompMsg(fn TOnCompMsgEvent)                                                                                                 // property event
	SetOnWidgetCompMsg(fn TOnCompMsgEvent)                                                                                                  // property event
	SetOnRenderCompMsg(fn TOnCompMsgEvent)                                                                                                  // property event
	SetOnProcessMessageReceived(fn TOnProcessMessageReceived)                                                                               // property event
	SetOnLoadStart(fn TOnLoadStart)                                                                                                         // property event
	SetOnLoadEnd(fn TOnLoadEnd)                                                                                                             // property event
	SetOnLoadError(fn TOnLoadError)                                                                                                         // property event
	SetOnLoadingStateChange(fn TOnLoadingStateChange)                                                                                       // property event
	SetOnTakeFocus(fn TOnTakeFocus)                                                                                                         // property event
	SetOnSetFocus(fn TOnSetFocus)                                                                                                           // property event
	SetOnGotFocus(fn TOnGotFocus)                                                                                                           // property event
	SetOnBeforeContextMenu(fn TOnBeforeContextMenu)                                                                                         // property event
	SetOnRunContextMenu(fn TOnRunContextMenu)                                                                                               // property event
	SetOnContextMenuCommand(fn TOnContextMenuCommand)                                                                                       // property event
	SetOnContextMenuDismissed(fn TOnContextMenuDismissed)                                                                                   // property event
	SetOnRunQuickMenu(fn TOnRunQuickMenuEvent)                                                                                              // property event
	SetOnQuickMenuCommand(fn TOnQuickMenuCommandEvent)                                                                                      // property event
	SetOnQuickMenuDismissed(fn TOnQuickMenuDismissedEvent)                                                                                  // property event
	SetOnPreKeyEvent(fn TOnPreKeyEvent)                                                                                                     // property event
	SetOnKeyEvent(fn TOnKeyEvent)                                                                                                           // property event
	SetOnAddressChange(fn TOnAddressChange)                                                                                                 // property event
	SetOnTitleChange(fn TOnTitleChange)                                                                                                     // property event
	SetOnFavIconUrlChange(fn TOnFavIconUrlChange)                                                                                           // property event
	SetOnFullScreenModeChange(fn TOnFullScreenModeChange)                                                                                   // property event
	SetOnTooltip(fn TOnTooltip)                                                                                                             // property event
	SetOnStatusMessage(fn TOnStatusMessage)                                                                                                 // property event
	SetOnConsoleMessage(fn TOnConsoleMessage)                                                                                               // property event
	SetOnAutoResize(fn TOnAutoResize)                                                                                                       // property event
	SetOnLoadingProgressChange(fn TOnLoadingProgressChange)                                                                                 // property event
	SetOnCursorChange(fn TOnCursorChange)                                                                                                   // property event
	SetOnMediaAccessChange(fn TOnMediaAccessChange)                                                                                         // property event
	SetOnCanDownload(fn TOnCanDownloadEvent)                                                                                                // property event
	SetOnBeforeDownload(fn TOnBeforeDownload)                                                                                               // property event
	SetOnDownloadUpdated(fn TOnDownloadUpdated)                                                                                             // property event
	SetOnJsdialog(fn TOnJsdialog)                                                                                                           // property event
	SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog)                                                                                       // property event
	SetOnResetDialogState(fn TOnResetDialogState)                                                                                           // property event
	SetOnDialogClosed(fn TOnDialogClosed)                                                                                                   // property event
	SetOnBeforePopup(fn TOnBeforePopup)                                                                                                     // property event
	SetOnAfterCreated(fn TOnAfterCreated)                                                                                                   // property event
	SetOnBeforeClose(fn TOnBeforeClose)                                                                                                     // property event
	SetOnClose(fn TOnClose)                                                                                                                 // property event
	SetOnBeforeBrowse(fn TOnBeforeBrowse)                                                                                                   // property event
	SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab)                                                                                               // property event
	SetOnGetAuthCredentials(fn TOnGetAuthCredentials)                                                                                       // property event
	SetOnCertificateError(fn TOnCertificateError)                                                                                           // property event
	SetOnSelectClientCertificate(fn TOnSelectClientCertificate)                                                                             // property event
	SetOnRenderViewReady(fn TOnRenderViewReady)                                                                                             // property event
	SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated)                                                                             // property event
	SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler)                                                                 // property event
	SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame)                                                                   // property event
	SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad)                                                                                       // property event
	SetOnGetResourceHandler(fn TOnGetResourceHandler)                                                                                       // property event
	SetOnResourceRedirect(fn TOnResourceRedirect)                                                                                           // property event
	SetOnResourceResponse(fn TOnResourceResponse)                                                                                           // property event
	SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter)                                                                         // property event
	SetOnResourceLoadComplete(fn TOnResourceLoadComplete)                                                                                   // property event
	SetOnProtocolExecution(fn TOnProtocolExecution)                                                                                         // property event
	SetOnCanSendCookie(fn TOnCanSendCookie)                                                                                                 // property event
	SetOnCanSaveCookie(fn TOnCanSaveCookie)                                                                                                 // property event
	SetOnFileDialog(fn TOnFileDialog)                                                                                                       // property event
	SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler)                                                                             // property event
	SetOnGetRootScreenRect(fn TOnGetRootScreenRect)                                                                                         // property event
	SetOnGetViewRect(fn TOnGetViewRect)                                                                                                     // property event
	SetOnGetScreenPoint(fn TOnGetScreenPoint)                                                                                               // property event
	SetOnGetScreenInfo(fn TOnGetScreenInfo)                                                                                                 // property event
	SetOnPopupShow(fn TOnPopupShow)                                                                                                         // property event
	SetOnPopupSize(fn TOnPopupSize)                                                                                                         // property event
	SetOnPaint(fn TOnPaint)                                                                                                                 // property event
	SetOnAcceleratedPaint(fn TOnAcceleratedPaint)                                                                                           // property event
	SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize)                                                                                       // property event
	SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged)                                                                             // property event
	SetOnStartDragging(fn TOnStartDragging)                                                                                                 // property event
	SetOnUpdateDragCursor(fn TOnUpdateDragCursor)                                                                                           // property event
	SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged)                                                                                     // property event
	SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged)                                                                       // property event
	SetOnTextSelectionChanged(fn TOnTextSelectionChanged)                                                                                   // property event
	SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested)                                                                           // property event
	SetOnDragEnter(fn TOnDragEnter)                                                                                                         // property event
	SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged)                                                                             // property event
	SetOnFindResult(fn TOnFindResult)                                                                                                       // property event
	SetOnRequestContextInitialized(fn TOnRequestContextInitialized)                                                                         // property event
	SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler)                                                              // property event
	SetOnSinks(fn TOnSinksEvent)                                                                                                            // property event
	SetOnRoutes(fn TOnRoutesEvent)                                                                                                          // property event
	SetOnRouteStateChanged(fn TOnRouteStateChangedEvent)                                                                                    // property event
	SetOnRouteMessageReceived(fn TOnRouteMessageReceivedEvent)                                                                              // property event
	SetOnGetAudioParameters(fn TOnGetAudioParametersEvent)                                                                                  // property event
	SetOnAudioStreamStarted(fn TOnAudioStreamStartedEvent)                                                                                  // property event
	SetOnAudioStreamPacket(fn TOnAudioStreamPacketEvent)                                                                                    // property event
	SetOnAudioStreamStopped(fn TOnAudioStreamStoppedEvent)                                                                                  // property event
	SetOnAudioStreamError(fn TOnAudioStreamErrorEvent)                                                                                      // property event
	SetOnDevToolsMessage(fn TOnDevToolsMessageEvent)                                                                                        // property event
	SetOnDevToolsRawMessage(fn TOnDevToolsRawMessageEvent)                                                                                  // property event
	SetOnDevToolsMethodResult(fn TOnDevToolsMethodResultEvent)                                                                              // property event
	SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResultEvent)                                                                        // property event
	SetOnDevToolsEvent(fn TOnDevToolsEventEvent)                                                                                            // property event
	SetOnDevToolsRawEvent(fn TOnDevToolsEventRawEvent)                                                                                      // property event
	SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttachedEvent)                                                                            // property event
	SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetachedEvent)                                                                            // property event
	SetOnExtensionLoadFailed(fn TOnExtensionLoadFailedEvent)                                                                                // property event
	SetOnExtensionLoaded(fn TOnExtensionLoadedEvent)                                                                                        // property event
	SetOnExtensionUnloaded(fn TOnExtensionUnloadedEvent)                                                                                    // property event
	SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowserEvent)                                                               // property event
	SetOnExtensionBeforeBrowser(fn TOnBeforeBrowserEvent)                                                                                   // property event
	SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowserEvent)                                                                             // property event
	SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowserEvent)                                                                             // property event
	SetOnExtensionGetExtensionResource(fn TOnGetExtensionResourceEvent)                                                                     // property event
	SetOnPrintStart(fn TOnPrintStartEvent)                                                                                                  // property event
	SetOnPrintSettings(fn TOnPrintSettingsEvent)                                                                                            // property event
	SetOnPrintDialog(fn TOnPrintDialogEvent)                                                                                                // property event
	SetOnPrintJob(fn TOnPrintJobEvent)                                                                                                      // property event
	SetOnPrintReset(fn TOnPrintResetEvent)                                                                                                  // property event
	SetOnGetPDFPaperSize(fn TOnGetPDFPaperSizeEvent)                                                                                        // property event
	SetOnFrameCreated(fn TOnFrameCreated)                                                                                                   // property event
	SetOnFrameAttached(fn TOnFrameAttached)                                                                                                 // property event
	SetOnFrameDetached(fn TOnFrameDetached)                                                                                                 // property event
	SetOnMainFrameChanged(fn TOnMainFrameChanged)                                                                                           // property event
	SetOnChromeCommand(fn TOnChromeCommandEvent)                                                                                            // property event
	SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermissionEvent)                                                              // property event
	SetOnShowPermissionPrompt(fn TOnShowPermissionPromptEvent)                                                                              // property event
	SetOnDismissPermissionPrompt(fn TOnDismissPermissionPromptEvent)                                                                        // property event
	AsIntfChromiumEvents() uintptr
}

type TChromiumCore struct {
	TComponent
}

func (m *TChromiumCore) CreateClientHandlerWithBool(isOSR bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(1, m.Instance(), api.PasBool(isOSR))
	return api.GoBool(r)
}

func (m *TChromiumCore) CreateClientHandlerWithClientBool(client *IEngClient, isOSR bool) bool {
	if !m.IsValid() {
		return false
	}
	clientPtr := base.GetObjectUintptr(*client)
	r := chromiumCoreAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&clientPtr)), api.PasBool(isOSR))
	*client = AsEngClient(clientPtr)
	return api.GoBool(r)
}

func (m *TChromiumCore) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SelectBrowser(iD int32) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(4, m.Instance(), uintptr(iD))
	return api.GoBool(r)
}

func (m *TChromiumCore) IndexOfBrowserID(iD int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(5, m.Instance(), uintptr(iD))
	return int32(r)
}

func (m *TChromiumCore) ShareRequestContext(context *ICefRequestContext, handler IEngRequestContextHandler) bool {
	if !m.IsValid() {
		return false
	}
	contextPtr := base.GetObjectUintptr(*context)
	r := chromiumCoreAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&contextPtr)), base.GetObjectUintptr(handler))
	*context = AsCefRequestContextRef(contextPtr)
	return api.GoBool(r)
}

func (m *TChromiumCore) SetNewBrowserParent(newParentHwnd types.HWND) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(7, m.Instance(), uintptr(newParentHwnd))
	return api.GoBool(r)
}

func (m *TChromiumCore) CreateBrowserWithWHandleRectStrRContextDValueBool(parentHandle cefTypes.TCefWindowHandle, parentRect types.TRect, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue, forceAsPopup bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(8, m.Instance(), uintptr(parentHandle), uintptr(base.UnsafePointer(&parentRect)), api.PasStr(windowName), base.GetObjectUintptr(context), base.GetObjectUintptr(extraInfo), api.PasBool(forceAsPopup))
	return api.GoBool(r)
}

func (m *TChromiumCore) CreateBrowserWithStrBVComponentRContextDValue(uRL string, browserViewComp ICEFBrowserViewComponent, context ICefRequestContext, extraInfo ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(9, m.Instance(), api.PasStr(uRL), base.GetObjectUintptr(browserViewComp), base.GetObjectUintptr(context), base.GetObjectUintptr(extraInfo))
	return api.GoBool(r)
}

func (m *TChromiumCore) ClearCertificateExceptions(clearImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(10, m.Instance(), api.PasBool(clearImmediately))
	return api.GoBool(r)
}

func (m *TChromiumCore) ClearHttpAuthCredentials(clearImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(11, m.Instance(), api.PasBool(clearImmediately))
	return api.GoBool(r)
}

func (m *TChromiumCore) CloseAllConnections(closeImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(12, m.Instance(), api.PasBool(closeImmediately))
	return api.GoBool(r)
}

func (m *TChromiumCore) GetFrameNames(frameNames *lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	frameNamesPtr := base.GetObjectUintptr(*frameNames)
	r := chromiumCoreAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&frameNamesPtr)))
	*frameNames = lcl.AsStrings(frameNamesPtr)
	return api.GoBool(r)
}

func (m *TChromiumCore) GetFrameIdentifiers(frameCount *cefTypes.NativeUInt, frameIdentifierArray *ICefFrameIdentifierArray) bool {
	if !m.IsValid() {
		return false
	}
	frameCountPtr := uintptr(*frameCount)
	var frameIdentifierArrayPtr uintptr
	var frameIdentifierArrayCountPtr uintptr
	r := chromiumCoreAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&frameCountPtr)), uintptr(base.UnsafePointer(&frameIdentifierArrayPtr)), uintptr(base.UnsafePointer(&frameIdentifierArrayCountPtr)))
	*frameCount = cefTypes.NativeUInt(frameCountPtr)
	*frameIdentifierArray = NewCefFrameIdentifierArray(int(frameIdentifierArrayCountPtr), frameIdentifierArrayPtr)
	return api.GoBool(r)
}

func (m *TChromiumCore) IsSameBrowser(browser ICefBrowser) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(browser))
	return api.GoBool(r)
}

func (m *TChromiumCore) ExecuteTaskOnCefThread(cefThreadId cefTypes.TCefThreadId, taskID uint32, delayMs int64) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(16, m.Instance(), uintptr(cefThreadId), uintptr(taskID), uintptr(base.UnsafePointer(&delayMs)))
	return api.GoBool(r)
}

func (m *TChromiumCore) DeleteCookies(url string, cookieName string, deleteImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(17, m.Instance(), api.PasStr(url), api.PasStr(cookieName), api.PasBool(deleteImmediately))
	return api.GoBool(r)
}

func (m *TChromiumCore) VisitAllCookies(iD int32) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(18, m.Instance(), uintptr(iD))
	return api.GoBool(r)
}

func (m *TChromiumCore) VisitURLCookies(url string, includeHttpOnly bool, iD int32) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(19, m.Instance(), api.PasStr(url), api.PasBool(includeHttpOnly), uintptr(iD))
	return api.GoBool(r)
}

func (m *TChromiumCore) SetCookie(args TChromiumCoreSetCookieArgs) bool {
	if !m.IsValid() {
		return false
	}
	argsPtr := args.ToPas()
	r := chromiumCoreAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(argsPtr)))
	return api.GoBool(r)
}

func (m *TChromiumCore) FlushCookieStore(flushImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(21, m.Instance(), api.PasBool(flushImmediately))
	return api.GoBool(r)
}

func (m *TChromiumCore) SendDevToolsMessage(message string) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(22, m.Instance(), api.PasStr(message))
	return api.GoBool(r)
}

func (m *TChromiumCore) ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(23, m.Instance(), uintptr(messageId), api.PasStr(method), base.GetObjectUintptr(params))
	return int32(r)
}

func (m *TChromiumCore) AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(observer), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) CreateUrlRequestWithRequestUClientStr(request ICefRequest, client IEngUrlrequestClient, frameName string) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), api.PasStr(frameName), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (m *TChromiumCore) CreateUrlRequestWithRequestUClientFrame(request ICefRequest, client IEngUrlrequestClient, frame ICefFrame) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), base.GetObjectUintptr(frame), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (m *TChromiumCore) CreateUrlRequestWithRequestUClientInt64(request ICefRequest, client IEngUrlrequestClient, frameIdentifier int64) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), uintptr(base.UnsafePointer(&frameIdentifier)), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (m *TChromiumCore) AddObserver(observer IEngMediaObserver) (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(28, m.Instance(), base.GetObjectUintptr(observer), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) GetSource(urn string) (result ICefMediaSource) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(29, m.Instance(), api.PasStr(urn), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaSourceRef(resultPtr)
	return
}

func (m *TChromiumCore) LoadExtension(rootDirectory string, manifest ICefDictionaryValue, handler IEngExtensionHandler, requestContext ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(30, m.Instance(), api.PasStr(rootDirectory), base.GetObjectUintptr(manifest), base.GetObjectUintptr(handler), base.GetObjectUintptr(requestContext))
	return api.GoBool(r)
}

func (m *TChromiumCore) DidLoadExtension(extensionId string) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(31, m.Instance(), api.PasStr(extensionId))
	return api.GoBool(r)
}

func (m *TChromiumCore) HasExtension(extensionId string) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(32, m.Instance(), api.PasStr(extensionId))
	return api.GoBool(r)
}

func (m *TChromiumCore) GetExtensions(extensionIds lcl.IStringList) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(extensionIds))
	return api.GoBool(r)
}

func (m *TChromiumCore) GetExtension(extensionId string) (result ICefExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(34, m.Instance(), api.PasStr(extensionId), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefExtensionRef(resultPtr)
	return
}

func (m *TChromiumCore) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(35, m.Instance(), api.PasBool(forceClose))
}

func (m *TChromiumCore) CloseAllBrowsers() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(36, m.Instance())
}

func (m *TChromiumCore) InitializeDragAndDrop(dropTargetWnd types.HWND) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(37, m.Instance(), uintptr(dropTargetWnd))
}

func (m *TChromiumCore) ShutdownDragAndDrop() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(38, m.Instance())
}

func (m *TChromiumCore) LoadURLWithStrX2(uRL string, frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(39, m.Instance(), api.PasStr(uRL), api.PasStr(frameName))
}

func (m *TChromiumCore) LoadURLWithStrFrame(uRL string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(40, m.Instance(), api.PasStr(uRL), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadURLWithStrInt64(uRL string, frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(41, m.Instance(), api.PasStr(uRL), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) LoadStringWithStrX2(hTML string, frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(42, m.Instance(), api.PasStr(hTML), api.PasStr(frameName))
}

func (m *TChromiumCore) LoadStringWithStrFrame(hTML string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(43, m.Instance(), api.PasStr(hTML), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadStringWithStrInt64(hTML string, frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(44, m.Instance(), api.PasStr(hTML), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) LoadResourceWithCMStreamStrX3(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(45, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType), api.PasStr(charset), api.PasStr(frameName))
}

func (m *TChromiumCore) LoadResourceWithCMStreamStrX2Frame(stream lcl.ICustomMemoryStream, mimeType string, charset string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(46, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType), api.PasStr(charset), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadResourceWithCMStreamStrX2Int64(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(47, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType), api.PasStr(charset), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) LoadRequest(request ICefRequest) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(48, m.Instance(), base.GetObjectUintptr(request))
}

func (m *TChromiumCore) GoBack() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(49, m.Instance())
}

func (m *TChromiumCore) GoForward() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(50, m.Instance())
}

func (m *TChromiumCore) Reload() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(51, m.Instance())
}

func (m *TChromiumCore) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(52, m.Instance())
}

func (m *TChromiumCore) StopLoad() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(53, m.Instance())
}

func (m *TChromiumCore) StartDownload(uRL string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(54, m.Instance(), api.PasStr(uRL))
}

func (m *TChromiumCore) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(55, m.Instance(), api.PasStr(imageUrl), api.PasBool(isFavicon), uintptr(maxImageSize), api.PasBool(bypassCache))
}

func (m *TChromiumCore) SimulateMouseWheel(deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(56, m.Instance(), uintptr(deltaX), uintptr(deltaY))
}

func (m *TChromiumCore) RetrieveHTMLWithStr(frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(57, m.Instance(), api.PasStr(frameName))
}

func (m *TChromiumCore) RetrieveHTMLWithFrame(frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(58, m.Instance(), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) RetrieveHTMLWithInt64(frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(59, m.Instance(), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) RetrieveTextWithStr(frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(60, m.Instance(), api.PasStr(frameName))
}

func (m *TChromiumCore) RetrieveTextWithFrame(frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(61, m.Instance(), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) RetrieveTextWithInt64(frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(62, m.Instance(), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) GetNavigationEntries(currentOnly bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(63, m.Instance(), api.PasBool(currentOnly))
}

func (m *TChromiumCore) ExecuteJavaScriptWithStrX3Int(code string, scriptURL string, frameName string, startLine int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(64, m.Instance(), api.PasStr(code), api.PasStr(scriptURL), api.PasStr(frameName), uintptr(startLine))
}

func (m *TChromiumCore) ExecuteJavaScriptWithStrX2FrameInt(code string, scriptURL string, frame ICefFrame, startLine int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(65, m.Instance(), api.PasStr(code), api.PasStr(scriptURL), base.GetObjectUintptr(frame), uintptr(startLine))
}

func (m *TChromiumCore) ExecuteJavaScriptWithStrX2Int64Int(code string, scriptURL string, frameIdentifier int64, startLine int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(66, m.Instance(), api.PasStr(code), api.PasStr(scriptURL), uintptr(base.UnsafePointer(&frameIdentifier)), uintptr(startLine))
}

func (m *TChromiumCore) UpdatePreferences() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(67, m.Instance())
}

func (m *TChromiumCore) SavePreferences(fileName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(68, m.Instance(), api.PasStr(fileName))
}

func (m *TChromiumCore) ResolveHost(uRL string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(69, m.Instance(), api.PasStr(uRL))
}

func (m *TChromiumCore) SetUserAgentOverride(userAgent string, acceptLanguage string, platform string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(70, m.Instance(), api.PasStr(userAgent), api.PasStr(acceptLanguage), api.PasStr(platform))
}

func (m *TChromiumCore) ClearDataForOrigin(origin string, storageTypes cefTypes.TCefClearDataStorageTypes) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(71, m.Instance(), api.PasStr(origin), uintptr(storageTypes))
}

func (m *TChromiumCore) ClearCache() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(72, m.Instance())
}

func (m *TChromiumCore) ToggleAudioMuted() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(73, m.Instance())
}

func (m *TChromiumCore) ShowDevTools(inspectElementAt types.TPoint, windowInfo TCefWindowInfo) {
	if !m.IsValid() {
		return
	}
	windowInfoPtr := windowInfo.ToPas()
	chromiumCoreAPI().SysCallN(74, m.Instance(), uintptr(base.UnsafePointer(&inspectElementAt)), uintptr(base.UnsafePointer(windowInfoPtr)))
}

func (m *TChromiumCore) CloseDevTools() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(75, m.Instance())
}

func (m *TChromiumCore) CloseDevToolsWithWindowHandle(devToolsWnd cefTypes.TCefWindowHandle) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(76, m.Instance(), uintptr(devToolsWnd))
}

func (m *TChromiumCore) Find(searchText string, forward bool, matchCase bool, findNext bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(77, m.Instance(), api.PasStr(searchText), api.PasBool(forward), api.PasBool(matchCase), api.PasBool(findNext))
}

func (m *TChromiumCore) StopFinding(clearSelection bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(78, m.Instance(), api.PasBool(clearSelection))
}

func (m *TChromiumCore) Print() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(79, m.Instance())
}

func (m *TChromiumCore) PrintToPDF(filePath string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(80, m.Instance(), api.PasStr(filePath))
}

func (m *TChromiumCore) ClipboardCopy() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(81, m.Instance())
}

func (m *TChromiumCore) ClipboardPaste() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(82, m.Instance())
}

func (m *TChromiumCore) ClipboardCut() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(83, m.Instance())
}

func (m *TChromiumCore) ClipboardUndo() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(84, m.Instance())
}

func (m *TChromiumCore) ClipboardRedo() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(85, m.Instance())
}

func (m *TChromiumCore) ClipboardDel() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(86, m.Instance())
}

func (m *TChromiumCore) SelectAll() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(87, m.Instance())
}

func (m *TChromiumCore) IncZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(88, m.Instance())
}

func (m *TChromiumCore) DecZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(89, m.Instance())
}

func (m *TChromiumCore) IncZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(90, m.Instance())
}

func (m *TChromiumCore) DecZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(91, m.Instance())
}

func (m *TChromiumCore) ResetZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(92, m.Instance())
}

func (m *TChromiumCore) ResetZoomLevel() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(93, m.Instance())
}

func (m *TChromiumCore) ResetZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(94, m.Instance())
}

func (m *TChromiumCore) ReadZoom() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(95, m.Instance())
}

func (m *TChromiumCore) WasResized() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(96, m.Instance())
}

func (m *TChromiumCore) WasHidden(hidden bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(97, m.Instance(), api.PasBool(hidden))
}

func (m *TChromiumCore) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(98, m.Instance())
}

func (m *TChromiumCore) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(99, m.Instance())
}

func (m *TChromiumCore) Invalidate(type_ cefTypes.TCefPaintElementType) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(100, m.Instance(), uintptr(type_))
}

func (m *TChromiumCore) SendExternalBeginFrame() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(101, m.Instance())
}

func (m *TChromiumCore) SendKeyEvent(event TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(102, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(103, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(type_), api.PasBool(mouseUp), uintptr(clickCount))
}

func (m *TChromiumCore) SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(104, m.Instance(), uintptr(base.UnsafePointer(&event)), api.PasBool(mouseLeave))
}

func (m *TChromiumCore) SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(105, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TChromiumCore) SendTouchEvent(event TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(106, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(107, m.Instance())
}

func (m *TChromiumCore) SendProcessMessageWithPIdPMessageStr(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(108, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(procMessage), api.PasStr(frameName))
}

func (m *TChromiumCore) SendProcessMessageWithPIdPMessageFrame(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(109, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(procMessage), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) SendProcessMessageWithPIdPMessageInt64(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(110, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(procMessage), uintptr(base.UnsafePointer(&frameIdentifier)))
}

func (m *TChromiumCore) SetFocus(focus bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(111, m.Instance(), api.PasBool(focus))
}

func (m *TChromiumCore) SetAccessibilityState(accessibilityState cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(112, m.Instance(), uintptr(accessibilityState))
}

func (m *TChromiumCore) DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(113, m.Instance(), base.GetObjectUintptr(dragData), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(114, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragLeave() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(115, m.Instance())
}

func (m *TChromiumCore) DragTargetDrop(event TCefMouseEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(116, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(117, m.Instance(), uintptr(X), uintptr(Y), uintptr(op))
}

func (m *TChromiumCore) DragSourceSystemDragEnded() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(118, m.Instance())
}

func (m *TChromiumCore) IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(119, m.Instance(), api.PasStr(text), underlines.Instance(), uintptr(int32(underlines.Count())), uintptr(base.UnsafePointer(&replacementRange)), uintptr(base.UnsafePointer(&selectionRange)))
}

func (m *TChromiumCore) IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(120, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&replacementRange)), uintptr(relativeCursorPos))
}

func (m *TChromiumCore) IMEFinishComposingText(keepSelection bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(121, m.Instance(), api.PasBool(keepSelection))
}

func (m *TChromiumCore) IMECancelComposition() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(122, m.Instance())
}

func (m *TChromiumCore) ReplaceMisspelling(word string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(123, m.Instance(), api.PasStr(word))
}

func (m *TChromiumCore) AddWordToDictionary(word string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(124, m.Instance(), api.PasStr(word))
}

func (m *TChromiumCore) UpdateBrowserSize(left int32, top int32, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(125, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

func (m *TChromiumCore) UpdateXWindowVisibility(visible bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(126, m.Instance(), api.PasBool(visible))
}

func (m *TChromiumCore) NotifyCurrentSinks() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(127, m.Instance())
}

func (m *TChromiumCore) NotifyCurrentRoutes() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(128, m.Instance())
}

func (m *TChromiumCore) CreateRoute(source ICefMediaSource, sink ICefMediaSink) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(129, m.Instance(), base.GetObjectUintptr(source), base.GetObjectUintptr(sink))
}

func (m *TChromiumCore) GetDeviceInfo(mediaSink ICefMediaSink) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(130, m.Instance(), base.GetObjectUintptr(mediaSink))
}

func (m *TChromiumCore) DefaultUrl() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(131, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetDefaultUrl(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(131, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) Options() IChromiumOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(132, 0, m.Instance())
	return AsChromiumOptions(r)
}

func (m *TChromiumCore) SetOptions(value IChromiumOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(132, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) FontOptions() IChromiumFontOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(133, 0, m.Instance())
	return AsChromiumFontOptions(r)
}

func (m *TChromiumCore) SetFontOptions(value IChromiumFontOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(133, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) PDFPrintOptions() IPDFPrintOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(134, 0, m.Instance())
	return AsPDFPrintOptions(r)
}

func (m *TChromiumCore) SetPDFPrintOptions(value IPDFPrintOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(134, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) DefaultEncoding() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(135, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(135, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) BrowserId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(136, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) Browser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(137, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TChromiumCore) BrowserById(id int32) (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(138, m.Instance(), uintptr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TChromiumCore) BrowserCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(139, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) BrowserIdByIndex(I int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(140, m.Instance(), uintptr(I))
	return int32(r)
}

func (m *TChromiumCore) CefClient() (result IEngClient) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(141, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngClient(resultPtr)
	return
}

func (m *TChromiumCore) ReqContextHandler() (result IEngRequestContextHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(142, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngRequestContextHandler(resultPtr)
	return
}

func (m *TChromiumCore) ResourceRequestHandler() (result IEngResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(143, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngResourceRequestHandler(resultPtr)
	return
}

func (m *TChromiumCore) CefWindowInfo() (result TCefWindowInfo) {
	if !m.IsValid() {
		return
	}
	resultPtr := result.ToPas()
	chromiumCoreAPI().SysCallN(144, m.Instance(), uintptr(base.UnsafePointer(resultPtr)))
	result = resultPtr.ToGo()
	return
}

func (m *TChromiumCore) VisibleNavigationEntry() (result ICefNavigationEntry) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(145, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefNavigationEntryRef(resultPtr)
	return
}

func (m *TChromiumCore) RequestContext() (result ICefRequestContext) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(146, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (m *TChromiumCore) MediaRouter() (result ICefMediaRouter) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(147, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouterRef(resultPtr)
	return
}

func (m *TChromiumCore) MediaObserver() (result IEngMediaObserver) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(148, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngMediaObserver(resultPtr)
	return
}

func (m *TChromiumCore) MediaObserverReg() (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(149, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) DevToolsMsgObserver() (result IEngDevToolsMessageObserver) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(150, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngDevToolsMessageObserver(resultPtr)
	return
}

func (m *TChromiumCore) DevToolsMsgObserverReg() (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(151, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) ExtensionHandler() (result IEngExtensionHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(152, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngExtensionHandler(resultPtr)
	return
}

func (m *TChromiumCore) MultithreadApp() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(153, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(154, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasDocument() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(155, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasView() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(156, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(157, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasClientHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(158, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(159, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(160, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(161, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) IsPopUp() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(162, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) WindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(163, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TChromiumCore) OpenerWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(164, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TChromiumCore) BrowserHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(165, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) WidgetHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(166, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) RenderHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(167, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) FrameIsFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(168, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(169, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) RequestContextCache() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(170, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) RequestContextIsGlobal() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(171, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) DocumentURL() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(172, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) ZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(173, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TChromiumCore) SetZoomLevel(value float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(173, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TChromiumCore) ZoomPct() (result float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(174, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TChromiumCore) SetZoomPct(value float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(174, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TChromiumCore) ZoomStep() byte {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(175, 0, m.Instance())
	return byte(r)
}

func (m *TChromiumCore) SetZoomStep(value byte) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(175, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(176, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetWindowlessFrameRate(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(176, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) CustomHeaderName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(177, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetCustomHeaderName(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(177, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) CustomHeaderValue() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(178, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetCustomHeaderValue(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(178, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) DoNotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(179, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetDoNotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(179, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SendReferrer() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(180, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSendReferrer(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(180, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) HyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(181, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetHyperlinkAuditing(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(181, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AllowOutdatedPlugins() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(182, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAllowOutdatedPlugins(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(182, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AlwaysAuthorizePlugins() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(183, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAlwaysAuthorizePlugins(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(183, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AlwaysOpenPDFExternally() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(184, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAlwaysOpenPDFExternally(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(184, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SpellChecking() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(185, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSpellChecking(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(185, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SpellCheckerDicts() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(186, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetSpellCheckerDicts(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(186, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) HasValidMainFrame() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(187, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) FrameCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(188, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TChromiumCore) DragOperations() cefTypes.TCefDragOperations {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(189, 0, m.Instance())
	return cefTypes.TCefDragOperations(r)
}

func (m *TChromiumCore) SetDragOperations(value cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(189, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) AudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(190, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAudioMuted(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(190, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SafeSearch() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(191, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSafeSearch(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(191, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) YouTubeRestrict() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(192, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetYouTubeRestrict(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(192, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) PrintingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(193, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetPrintingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(193, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AcceptLanguageList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(194, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetAcceptLanguageList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(194, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) AcceptCookies() cefTypes.TCefCookiePref {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(195, 0, m.Instance())
	return cefTypes.TCefCookiePref(r)
}

func (m *TChromiumCore) SetAcceptCookies(value cefTypes.TCefCookiePref) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(195, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) Block3rdPartyCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(196, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetBlock3rdPartyCookies(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(196, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) MultiBrowserMode() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(197, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetMultiBrowserMode(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(197, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) DefaultWindowInfoExStyle() types.DWORD {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(198, 0, m.Instance())
	return types.DWORD(r)
}

func (m *TChromiumCore) SetDefaultWindowInfoExStyle(value types.DWORD) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(198, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) Offline() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(199, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetOffline(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(199, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) QuicAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(200, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetQuicAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(200, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) JavascriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(201, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetJavascriptEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(201, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) LoadImagesAutomatically() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(202, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetLoadImagesAutomatically(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(202, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) BatterySaverModeState() cefTypes.TCefBatterySaverModeState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(203, 0, m.Instance())
	return cefTypes.TCefBatterySaverModeState(r)
}

func (m *TChromiumCore) SetBatterySaverModeState(value cefTypes.TCefBatterySaverModeState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(203, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) HighEfficiencyMode() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(204, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetHighEfficiencyMode(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(204, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) XDisplay() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(205, m.Instance())
	return uintptr(r)
}

func (m *TChromiumCore) WebRTCIPHandlingPolicy() cefTypes.TCefWebRTCHandlingPolicy {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(206, 0, m.Instance())
	return cefTypes.TCefWebRTCHandlingPolicy(r)
}

func (m *TChromiumCore) SetWebRTCIPHandlingPolicy(value cefTypes.TCefWebRTCHandlingPolicy) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(206, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WebRTCMultipleRoutes() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(207, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetWebRTCMultipleRoutes(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(207, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WebRTCNonproxiedUDP() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(208, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetWebRTCNonproxiedUDP(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(208, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyType() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(209, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetProxyType(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(209, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyScheme() cefTypes.TCefProxyScheme {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(210, 0, m.Instance())
	return cefTypes.TCefProxyScheme(r)
}

func (m *TChromiumCore) SetProxyScheme(value cefTypes.TCefProxyScheme) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(210, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyServer() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(211, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetProxyServer(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(211, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(212, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetProxyPort(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(212, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyUsername() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(213, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetProxyUsername(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(213, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyPassword() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(214, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetProxyPassword(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(214, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyScriptURL() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(215, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetProxyScriptURL(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(215, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyByPassList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumCoreAPI().SysCallN(216, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumCore) SetProxyByPassList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(216, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) MaxConnectionsPerProxy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(217, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetMaxConnectionsPerProxy(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(217, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) SetOnTextResultAvailable(fn TOnTextResultAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextResultAvailableEvent(fn)
	base.SetEvent(m, 218, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPdfPrintFinished(fn TOnPdfPrintFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPdfPrintFinishedEvent(fn)
	base.SetEvent(m, 219, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrefsAvailable(fn TOnPrefsAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrefsAvailableEvent(fn)
	base.SetEvent(m, 220, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrefsUpdated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 221, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesDeleted(fn TOnCookiesDeletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookiesDeletedEvent(fn)
	base.SetEvent(m, 222, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResolvedIPsAvailableEvent(fn)
	base.SetEvent(m, 223, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationVisitorResultAvailableEvent(fn)
	base.SetEvent(m, 224, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDownloadImageFinished(fn TOnDownloadImageFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadImageFinishedEvent(fn)
	base.SetEvent(m, 225, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesFlushed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 226, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCertificateExceptionsCleared(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 227, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnHttpAuthCredentialsCleared(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 228, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAllConnectionsClosed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 229, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExecuteTaskOnCefThread(fn)
	base.SetEvent(m, 230, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesVisited(fn TOnCookiesVisited) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookiesVisited(fn)
	base.SetEvent(m, 231, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieVisitorDestroyed(fn)
	base.SetEvent(m, 232, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookieSet(fn TOnCookieSet) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieSet(fn)
	base.SetEvent(m, 233, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnZoomPctAvailable(fn TOnZoomPctAvailable) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnZoomPctAvailable(fn)
	base.SetEvent(m, 234, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaRouteCreateFinishedEvent(fn)
	base.SetEvent(m, 235, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaSinkDeviceInfoEvent(fn)
	base.SetEvent(m, 236, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBrowserCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 237, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnWidgetCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 238, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 239, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceived) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessMessageReceived(fn)
	base.SetEvent(m, 240, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadStart(fn TOnLoadStart) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadStart(fn)
	base.SetEvent(m, 241, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadEnd(fn TOnLoadEnd) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadEnd(fn)
	base.SetEvent(m, 242, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadError(fn TOnLoadError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadError(fn)
	base.SetEvent(m, 243, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadingStateChange(fn TOnLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadingStateChange(fn)
	base.SetEvent(m, 244, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTakeFocus(fn TOnTakeFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTakeFocus(fn)
	base.SetEvent(m, 245, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSetFocus(fn TOnSetFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSetFocus(fn)
	base.SetEvent(m, 246, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGotFocus(fn TOnGotFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGotFocus(fn)
	base.SetEvent(m, 247, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeContextMenu(fn TOnBeforeContextMenu) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeContextMenu(fn)
	base.SetEvent(m, 248, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRunContextMenu(fn TOnRunContextMenu) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRunContextMenu(fn)
	base.SetEvent(m, 249, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnContextMenuCommand(fn TOnContextMenuCommand) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuCommand(fn)
	base.SetEvent(m, 250, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnContextMenuDismissed(fn TOnContextMenuDismissed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuDismissed(fn)
	base.SetEvent(m, 251, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRunQuickMenu(fn TOnRunQuickMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRunQuickMenuEvent(fn)
	base.SetEvent(m, 252, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnQuickMenuCommand(fn TOnQuickMenuCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnQuickMenuCommandEvent(fn)
	base.SetEvent(m, 253, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnQuickMenuDismissed(fn TOnQuickMenuDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnQuickMenuDismissedEvent(fn)
	base.SetEvent(m, 254, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPreKeyEvent(fn TOnPreKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPreKeyEvent(fn)
	base.SetEvent(m, 255, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnKeyEvent(fn TOnKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnKeyEvent(fn)
	base.SetEvent(m, 256, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAddressChange(fn TOnAddressChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAddressChange(fn)
	base.SetEvent(m, 257, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTitleChange(fn TOnTitleChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTitleChange(fn)
	base.SetEvent(m, 258, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFavIconUrlChange(fn TOnFavIconUrlChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFavIconUrlChange(fn)
	base.SetEvent(m, 259, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFullScreenModeChange(fn TOnFullScreenModeChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFullScreenModeChange(fn)
	base.SetEvent(m, 260, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTooltip(fn TOnTooltip) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTooltip(fn)
	base.SetEvent(m, 261, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnStatusMessage(fn TOnStatusMessage) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStatusMessage(fn)
	base.SetEvent(m, 262, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnConsoleMessage(fn TOnConsoleMessage) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnConsoleMessage(fn)
	base.SetEvent(m, 263, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAutoResize(fn TOnAutoResize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAutoResize(fn)
	base.SetEvent(m, 264, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadingProgressChange(fn TOnLoadingProgressChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadingProgressChange(fn)
	base.SetEvent(m, 265, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCursorChange(fn TOnCursorChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCursorChange(fn)
	base.SetEvent(m, 266, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaAccessChange(fn TOnMediaAccessChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaAccessChange(fn)
	base.SetEvent(m, 267, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanDownload(fn TOnCanDownloadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanDownloadEvent(fn)
	base.SetEvent(m, 268, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeDownload(fn TOnBeforeDownload) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeDownload(fn)
	base.SetEvent(m, 269, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDownloadUpdated(fn TOnDownloadUpdated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadUpdated(fn)
	base.SetEvent(m, 270, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnJsdialog(fn TOnJsdialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsdialog(fn)
	base.SetEvent(m, 271, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeUnloadDialog(fn)
	base.SetEvent(m, 272, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResetDialogState(fn TOnResetDialogState) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResetDialogState(fn)
	base.SetEvent(m, 273, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDialogClosed(fn TOnDialogClosed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDialogClosed(fn)
	base.SetEvent(m, 274, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforePopup(fn TOnBeforePopup) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforePopup(fn)
	base.SetEvent(m, 275, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAfterCreated(fn TOnAfterCreated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAfterCreated(fn)
	base.SetEvent(m, 276, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeClose(fn TOnBeforeClose) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeClose(fn)
	base.SetEvent(m, 277, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnClose(fn TOnClose) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClose(fn)
	base.SetEvent(m, 278, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeBrowse(fn TOnBeforeBrowse) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBrowse(fn)
	base.SetEvent(m, 279, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnOpenUrlFromTab(fn)
	base.SetEvent(m, 280, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAuthCredentials(fn TOnGetAuthCredentials) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAuthCredentials(fn)
	base.SetEvent(m, 281, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCertificateError(fn TOnCertificateError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCertificateError(fn)
	base.SetEvent(m, 282, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSelectClientCertificate(fn TOnSelectClientCertificate) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectClientCertificate(fn)
	base.SetEvent(m, 283, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderViewReady(fn TOnRenderViewReady) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderViewReady(fn)
	base.SetEvent(m, 284, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessTerminated(fn)
	base.SetEvent(m, 285, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceRequestHandler(fn)
	base.SetEvent(m, 286, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDocumentAvailableInMainFrame(fn)
	base.SetEvent(m, 287, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeResourceLoad(fn)
	base.SetEvent(m, 288, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceHandler(fn TOnGetResourceHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceHandler(fn)
	base.SetEvent(m, 289, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceRedirect(fn TOnResourceRedirect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRedirect(fn)
	base.SetEvent(m, 290, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceResponse(fn TOnResourceResponse) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceResponse(fn)
	base.SetEvent(m, 291, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceResponseFilter(fn)
	base.SetEvent(m, 292, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceLoadComplete(fn TOnResourceLoadComplete) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceLoadComplete(fn)
	base.SetEvent(m, 293, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnProtocolExecution(fn TOnProtocolExecution) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProtocolExecution(fn)
	base.SetEvent(m, 294, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanSendCookie(fn TOnCanSendCookie) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanSendCookie(fn)
	base.SetEvent(m, 295, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanSaveCookie(fn TOnCanSaveCookie) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanSaveCookie(fn)
	base.SetEvent(m, 296, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFileDialog(fn TOnFileDialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFileDialog(fn)
	base.SetEvent(m, 297, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAccessibilityHandler(fn)
	base.SetEvent(m, 298, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetRootScreenRect(fn TOnGetRootScreenRect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetRootScreenRect(fn)
	base.SetEvent(m, 299, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetViewRect(fn TOnGetViewRect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetViewRect(fn)
	base.SetEvent(m, 300, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetScreenPoint(fn TOnGetScreenPoint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetScreenPoint(fn)
	base.SetEvent(m, 301, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetScreenInfo(fn TOnGetScreenInfo) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetScreenInfo(fn)
	base.SetEvent(m, 302, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPopupShow(fn TOnPopupShow) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupShow(fn)
	base.SetEvent(m, 303, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPopupSize(fn TOnPopupSize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupSize(fn)
	base.SetEvent(m, 304, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPaint(fn TOnPaint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPaint(fn)
	base.SetEvent(m, 305, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAcceleratedPaint(fn TOnAcceleratedPaint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceleratedPaint(fn)
	base.SetEvent(m, 306, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetTouchHandleSize(fn)
	base.SetEvent(m, 307, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTouchHandleStateChanged(fn)
	base.SetEvent(m, 308, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnStartDragging(fn TOnStartDragging) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStartDragging(fn)
	base.SetEvent(m, 309, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnUpdateDragCursor(fn TOnUpdateDragCursor) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUpdateDragCursor(fn)
	base.SetEvent(m, 310, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScrollOffsetChanged(fn)
	base.SetEvent(m, 311, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIMECompositionRangeChanged(fn)
	base.SetEvent(m, 312, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTextSelectionChanged(fn TOnTextSelectionChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextSelectionChanged(fn)
	base.SetEvent(m, 313, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnVirtualKeyboardRequested(fn)
	base.SetEvent(m, 314, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDragEnter(fn TOnDragEnter) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDragEnter(fn)
	base.SetEvent(m, 315, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDraggableRegionsChanged(fn)
	base.SetEvent(m, 316, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFindResult(fn TOnFindResult) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFindResult(fn)
	base.SetEvent(m, 317, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRequestContextInitialized(fn TOnRequestContextInitialized) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestContextInitialized(fn)
	base.SetEvent(m, 318, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceRequestHandler(fn)
	base.SetEvent(m, 319, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSinks(fn TOnSinksEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSinksEvent(fn)
	base.SetEvent(m, 320, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRoutes(fn TOnRoutesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRoutesEvent(fn)
	base.SetEvent(m, 321, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRouteStateChanged(fn TOnRouteStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRouteStateChangedEvent(fn)
	base.SetEvent(m, 322, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRouteMessageReceived(fn TOnRouteMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRouteMessageReceivedEvent(fn)
	base.SetEvent(m, 323, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAudioParameters(fn TOnGetAudioParametersEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAudioParametersEvent(fn)
	base.SetEvent(m, 324, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamStarted(fn TOnAudioStreamStartedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamStartedEvent(fn)
	base.SetEvent(m, 325, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamPacket(fn TOnAudioStreamPacketEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamPacketEvent(fn)
	base.SetEvent(m, 326, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamStopped(fn TOnAudioStreamStoppedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamStoppedEvent(fn)
	base.SetEvent(m, 327, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamError(fn TOnAudioStreamErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamErrorEvent(fn)
	base.SetEvent(m, 328, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMessage(fn TOnDevToolsMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageEvent(fn)
	base.SetEvent(m, 329, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsRawMessage(fn TOnDevToolsRawMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsRawMessageEvent(fn)
	base.SetEvent(m, 330, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMethodResult(fn TOnDevToolsMethodResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMethodResultEvent(fn)
	base.SetEvent(m, 331, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMethodRawResultEvent(fn)
	base.SetEvent(m, 332, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsEvent(fn TOnDevToolsEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsEventEvent(fn)
	base.SetEvent(m, 333, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsRawEvent(fn TOnDevToolsEventRawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsEventRawEvent(fn)
	base.SetEvent(m, 334, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsAgentAttachedEvent(fn)
	base.SetEvent(m, 335, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsAgentDetachedEvent(fn)
	base.SetEvent(m, 336, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionLoadFailed(fn TOnExtensionLoadFailedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionLoadFailedEvent(fn)
	base.SetEvent(m, 337, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionLoaded(fn TOnExtensionLoadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionLoadedEvent(fn)
	base.SetEvent(m, 338, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionUnloaded(fn TOnExtensionUnloadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionUnloadedEvent(fn)
	base.SetEvent(m, 339, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBackgroundBrowserEvent(fn)
	base.SetEvent(m, 340, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionBeforeBrowser(fn TOnBeforeBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBrowserEvent(fn)
	base.SetEvent(m, 341, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetActiveBrowserEvent(fn)
	base.SetEvent(m, 342, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanAccessBrowserEvent(fn)
	base.SetEvent(m, 343, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionGetExtensionResource(fn TOnGetExtensionResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetExtensionResourceEvent(fn)
	base.SetEvent(m, 344, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintStart(fn TOnPrintStartEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintStartEvent(fn)
	base.SetEvent(m, 345, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintSettings(fn TOnPrintSettingsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintSettingsEvent(fn)
	base.SetEvent(m, 346, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintDialog(fn TOnPrintDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintDialogEvent(fn)
	base.SetEvent(m, 347, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintJob(fn TOnPrintJobEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintJobEvent(fn)
	base.SetEvent(m, 348, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintReset(fn TOnPrintResetEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintResetEvent(fn)
	base.SetEvent(m, 349, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetPDFPaperSize(fn TOnGetPDFPaperSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetPDFPaperSizeEvent(fn)
	base.SetEvent(m, 350, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameCreated(fn TOnFrameCreated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameCreated(fn)
	base.SetEvent(m, 351, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameAttached(fn TOnFrameAttached) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameAttached(fn)
	base.SetEvent(m, 352, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameDetached(fn TOnFrameDetached) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameDetached(fn)
	base.SetEvent(m, 353, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMainFrameChanged(fn TOnMainFrameChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMainFrameChanged(fn)
	base.SetEvent(m, 354, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnChromeCommand(fn TOnChromeCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnChromeCommandEvent(fn)
	base.SetEvent(m, 355, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermissionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestMediaAccessPermissionEvent(fn)
	base.SetEvent(m, 356, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnShowPermissionPrompt(fn TOnShowPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnShowPermissionPromptEvent(fn)
	base.SetEvent(m, 357, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDismissPermissionPrompt(fn TOnDismissPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDismissPermissionPromptEvent(fn)
	base.SetEvent(m, 358, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) AsIntfChromiumEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewChromiumCore class constructor
func NewChromiumCore(owner lcl.IComponent) IChromiumCore {
	var chromiumEventsPtr uintptr // IChromiumEvents
	r := chromiumCoreAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&chromiumEventsPtr)))
	ret := AsChromiumCore(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, chromiumEventsPtr)
	}
	return ret
}

var (
	chromiumCoreOnce   base.Once
	chromiumCoreImport *imports.Imports = nil
)

func chromiumCoreAPI() *imports.Imports {
	chromiumCoreOnce.Do(func() {
		chromiumCoreImport = api.NewDefaultImports()
		chromiumCoreImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChromiumCore_Create", 0), // constructor NewChromiumCore
			/* 1 */ imports.NewTable("TChromiumCore_CreateClientHandlerWithBool", 0), // function CreateClientHandlerWithBool
			/* 2 */ imports.NewTable("TChromiumCore_CreateClientHandlerWithClientBool", 0), // function CreateClientHandlerWithClientBool
			/* 3 */ imports.NewTable("TChromiumCore_TryCloseBrowser", 0), // function TryCloseBrowser
			/* 4 */ imports.NewTable("TChromiumCore_SelectBrowser", 0), // function SelectBrowser
			/* 5 */ imports.NewTable("TChromiumCore_IndexOfBrowserID", 0), // function IndexOfBrowserID
			/* 6 */ imports.NewTable("TChromiumCore_ShareRequestContext", 0), // function ShareRequestContext
			/* 7 */ imports.NewTable("TChromiumCore_SetNewBrowserParent", 0), // function SetNewBrowserParent
			/* 8 */ imports.NewTable("TChromiumCore_CreateBrowserWithWHandleRectStrRContextDValueBool", 0), // function CreateBrowserWithWHandleRectStrRContextDValueBool
			/* 9 */ imports.NewTable("TChromiumCore_CreateBrowserWithStrBVComponentRContextDValue", 0), // function CreateBrowserWithStrBVComponentRContextDValue
			/* 10 */ imports.NewTable("TChromiumCore_ClearCertificateExceptions", 0), // function ClearCertificateExceptions
			/* 11 */ imports.NewTable("TChromiumCore_ClearHttpAuthCredentials", 0), // function ClearHttpAuthCredentials
			/* 12 */ imports.NewTable("TChromiumCore_CloseAllConnections", 0), // function CloseAllConnections
			/* 13 */ imports.NewTable("TChromiumCore_GetFrameNames", 0), // function GetFrameNames
			/* 14 */ imports.NewTable("TChromiumCore_GetFrameIdentifiers", 0), // function GetFrameIdentifiers
			/* 15 */ imports.NewTable("TChromiumCore_IsSameBrowser", 0), // function IsSameBrowser
			/* 16 */ imports.NewTable("TChromiumCore_ExecuteTaskOnCefThread", 0), // function ExecuteTaskOnCefThread
			/* 17 */ imports.NewTable("TChromiumCore_DeleteCookies", 0), // function DeleteCookies
			/* 18 */ imports.NewTable("TChromiumCore_VisitAllCookies", 0), // function VisitAllCookies
			/* 19 */ imports.NewTable("TChromiumCore_VisitURLCookies", 0), // function VisitURLCookies
			/* 20 */ imports.NewTable("TChromiumCore_SetCookie", 0), // function SetCookie
			/* 21 */ imports.NewTable("TChromiumCore_FlushCookieStore", 0), // function FlushCookieStore
			/* 22 */ imports.NewTable("TChromiumCore_SendDevToolsMessage", 0), // function SendDevToolsMessage
			/* 23 */ imports.NewTable("TChromiumCore_ExecuteDevToolsMethod", 0), // function ExecuteDevToolsMethod
			/* 24 */ imports.NewTable("TChromiumCore_AddDevToolsMessageObserver", 0), // function AddDevToolsMessageObserver
			/* 25 */ imports.NewTable("TChromiumCore_CreateUrlRequestWithRequestUClientStr", 0), // function CreateUrlRequestWithRequestUClientStr
			/* 26 */ imports.NewTable("TChromiumCore_CreateUrlRequestWithRequestUClientFrame", 0), // function CreateUrlRequestWithRequestUClientFrame
			/* 27 */ imports.NewTable("TChromiumCore_CreateUrlRequestWithRequestUClientInt64", 0), // function CreateUrlRequestWithRequestUClientInt64
			/* 28 */ imports.NewTable("TChromiumCore_AddObserver", 0), // function AddObserver
			/* 29 */ imports.NewTable("TChromiumCore_GetSource", 0), // function GetSource
			/* 30 */ imports.NewTable("TChromiumCore_LoadExtension", 0), // function LoadExtension
			/* 31 */ imports.NewTable("TChromiumCore_DidLoadExtension", 0), // function DidLoadExtension
			/* 32 */ imports.NewTable("TChromiumCore_HasExtension", 0), // function HasExtension
			/* 33 */ imports.NewTable("TChromiumCore_GetExtensions", 0), // function GetExtensions
			/* 34 */ imports.NewTable("TChromiumCore_GetExtension", 0), // function GetExtension
			/* 35 */ imports.NewTable("TChromiumCore_CloseBrowser", 0), // procedure CloseBrowser
			/* 36 */ imports.NewTable("TChromiumCore_CloseAllBrowsers", 0), // procedure CloseAllBrowsers
			/* 37 */ imports.NewTable("TChromiumCore_InitializeDragAndDrop", 0), // procedure InitializeDragAndDrop
			/* 38 */ imports.NewTable("TChromiumCore_ShutdownDragAndDrop", 0), // procedure ShutdownDragAndDrop
			/* 39 */ imports.NewTable("TChromiumCore_LoadURLWithStrX2", 0), // procedure LoadURLWithStrX2
			/* 40 */ imports.NewTable("TChromiumCore_LoadURLWithStrFrame", 0), // procedure LoadURLWithStrFrame
			/* 41 */ imports.NewTable("TChromiumCore_LoadURLWithStrInt64", 0), // procedure LoadURLWithStrInt64
			/* 42 */ imports.NewTable("TChromiumCore_LoadStringWithStrX2", 0), // procedure LoadStringWithStrX2
			/* 43 */ imports.NewTable("TChromiumCore_LoadStringWithStrFrame", 0), // procedure LoadStringWithStrFrame
			/* 44 */ imports.NewTable("TChromiumCore_LoadStringWithStrInt64", 0), // procedure LoadStringWithStrInt64
			/* 45 */ imports.NewTable("TChromiumCore_LoadResourceWithCMStreamStrX3", 0), // procedure LoadResourceWithCMStreamStrX3
			/* 46 */ imports.NewTable("TChromiumCore_LoadResourceWithCMStreamStrX2Frame", 0), // procedure LoadResourceWithCMStreamStrX2Frame
			/* 47 */ imports.NewTable("TChromiumCore_LoadResourceWithCMStreamStrX2Int64", 0), // procedure LoadResourceWithCMStreamStrX2Int64
			/* 48 */ imports.NewTable("TChromiumCore_LoadRequest", 0), // procedure LoadRequest
			/* 49 */ imports.NewTable("TChromiumCore_GoBack", 0), // procedure GoBack
			/* 50 */ imports.NewTable("TChromiumCore_GoForward", 0), // procedure GoForward
			/* 51 */ imports.NewTable("TChromiumCore_Reload", 0), // procedure Reload
			/* 52 */ imports.NewTable("TChromiumCore_ReloadIgnoreCache", 0), // procedure ReloadIgnoreCache
			/* 53 */ imports.NewTable("TChromiumCore_StopLoad", 0), // procedure StopLoad
			/* 54 */ imports.NewTable("TChromiumCore_StartDownload", 0), // procedure StartDownload
			/* 55 */ imports.NewTable("TChromiumCore_DownloadImage", 0), // procedure DownloadImage
			/* 56 */ imports.NewTable("TChromiumCore_SimulateMouseWheel", 0), // procedure SimulateMouseWheel
			/* 57 */ imports.NewTable("TChromiumCore_RetrieveHTMLWithStr", 0), // procedure RetrieveHTMLWithStr
			/* 58 */ imports.NewTable("TChromiumCore_RetrieveHTMLWithFrame", 0), // procedure RetrieveHTMLWithFrame
			/* 59 */ imports.NewTable("TChromiumCore_RetrieveHTMLWithInt64", 0), // procedure RetrieveHTMLWithInt64
			/* 60 */ imports.NewTable("TChromiumCore_RetrieveTextWithStr", 0), // procedure RetrieveTextWithStr
			/* 61 */ imports.NewTable("TChromiumCore_RetrieveTextWithFrame", 0), // procedure RetrieveTextWithFrame
			/* 62 */ imports.NewTable("TChromiumCore_RetrieveTextWithInt64", 0), // procedure RetrieveTextWithInt64
			/* 63 */ imports.NewTable("TChromiumCore_GetNavigationEntries", 0), // procedure GetNavigationEntries
			/* 64 */ imports.NewTable("TChromiumCore_ExecuteJavaScriptWithStrX3Int", 0), // procedure ExecuteJavaScriptWithStrX3Int
			/* 65 */ imports.NewTable("TChromiumCore_ExecuteJavaScriptWithStrX2FrameInt", 0), // procedure ExecuteJavaScriptWithStrX2FrameInt
			/* 66 */ imports.NewTable("TChromiumCore_ExecuteJavaScriptWithStrX2Int64Int", 0), // procedure ExecuteJavaScriptWithStrX2Int64Int
			/* 67 */ imports.NewTable("TChromiumCore_UpdatePreferences", 0), // procedure UpdatePreferences
			/* 68 */ imports.NewTable("TChromiumCore_SavePreferences", 0), // procedure SavePreferences
			/* 69 */ imports.NewTable("TChromiumCore_ResolveHost", 0), // procedure ResolveHost
			/* 70 */ imports.NewTable("TChromiumCore_SetUserAgentOverride", 0), // procedure SetUserAgentOverride
			/* 71 */ imports.NewTable("TChromiumCore_ClearDataForOrigin", 0), // procedure ClearDataForOrigin
			/* 72 */ imports.NewTable("TChromiumCore_ClearCache", 0), // procedure ClearCache
			/* 73 */ imports.NewTable("TChromiumCore_ToggleAudioMuted", 0), // procedure ToggleAudioMuted
			/* 74 */ imports.NewTable("TChromiumCore_ShowDevTools", 0), // procedure ShowDevTools
			/* 75 */ imports.NewTable("TChromiumCore_CloseDevTools", 0), // procedure CloseDevTools
			/* 76 */ imports.NewTable("TChromiumCore_CloseDevToolsWithWindowHandle", 0), // procedure CloseDevToolsWithWindowHandle
			/* 77 */ imports.NewTable("TChromiumCore_Find", 0), // procedure Find
			/* 78 */ imports.NewTable("TChromiumCore_StopFinding", 0), // procedure StopFinding
			/* 79 */ imports.NewTable("TChromiumCore_Print", 0), // procedure Print
			/* 80 */ imports.NewTable("TChromiumCore_PrintToPDF", 0), // procedure PrintToPDF
			/* 81 */ imports.NewTable("TChromiumCore_ClipboardCopy", 0), // procedure ClipboardCopy
			/* 82 */ imports.NewTable("TChromiumCore_ClipboardPaste", 0), // procedure ClipboardPaste
			/* 83 */ imports.NewTable("TChromiumCore_ClipboardCut", 0), // procedure ClipboardCut
			/* 84 */ imports.NewTable("TChromiumCore_ClipboardUndo", 0), // procedure ClipboardUndo
			/* 85 */ imports.NewTable("TChromiumCore_ClipboardRedo", 0), // procedure ClipboardRedo
			/* 86 */ imports.NewTable("TChromiumCore_ClipboardDel", 0), // procedure ClipboardDel
			/* 87 */ imports.NewTable("TChromiumCore_SelectAll", 0), // procedure SelectAll
			/* 88 */ imports.NewTable("TChromiumCore_IncZoomStep", 0), // procedure IncZoomStep
			/* 89 */ imports.NewTable("TChromiumCore_DecZoomStep", 0), // procedure DecZoomStep
			/* 90 */ imports.NewTable("TChromiumCore_IncZoomPct", 0), // procedure IncZoomPct
			/* 91 */ imports.NewTable("TChromiumCore_DecZoomPct", 0), // procedure DecZoomPct
			/* 92 */ imports.NewTable("TChromiumCore_ResetZoomStep", 0), // procedure ResetZoomStep
			/* 93 */ imports.NewTable("TChromiumCore_ResetZoomLevel", 0), // procedure ResetZoomLevel
			/* 94 */ imports.NewTable("TChromiumCore_ResetZoomPct", 0), // procedure ResetZoomPct
			/* 95 */ imports.NewTable("TChromiumCore_ReadZoom", 0), // procedure ReadZoom
			/* 96 */ imports.NewTable("TChromiumCore_WasResized", 0), // procedure WasResized
			/* 97 */ imports.NewTable("TChromiumCore_WasHidden", 0), // procedure WasHidden
			/* 98 */ imports.NewTable("TChromiumCore_NotifyScreenInfoChanged", 0), // procedure NotifyScreenInfoChanged
			/* 99 */ imports.NewTable("TChromiumCore_NotifyMoveOrResizeStarted", 0), // procedure NotifyMoveOrResizeStarted
			/* 100 */ imports.NewTable("TChromiumCore_Invalidate", 0), // procedure Invalidate
			/* 101 */ imports.NewTable("TChromiumCore_SendExternalBeginFrame", 0), // procedure SendExternalBeginFrame
			/* 102 */ imports.NewTable("TChromiumCore_SendKeyEvent", 0), // procedure SendKeyEvent
			/* 103 */ imports.NewTable("TChromiumCore_SendMouseClickEvent", 0), // procedure SendMouseClickEvent
			/* 104 */ imports.NewTable("TChromiumCore_SendMouseMoveEvent", 0), // procedure SendMouseMoveEvent
			/* 105 */ imports.NewTable("TChromiumCore_SendMouseWheelEvent", 0), // procedure SendMouseWheelEvent
			/* 106 */ imports.NewTable("TChromiumCore_SendTouchEvent", 0), // procedure SendTouchEvent
			/* 107 */ imports.NewTable("TChromiumCore_SendCaptureLostEvent", 0), // procedure SendCaptureLostEvent
			/* 108 */ imports.NewTable("TChromiumCore_SendProcessMessageWithPIdPMessageStr", 0), // procedure SendProcessMessageWithPIdPMessageStr
			/* 109 */ imports.NewTable("TChromiumCore_SendProcessMessageWithPIdPMessageFrame", 0), // procedure SendProcessMessageWithPIdPMessageFrame
			/* 110 */ imports.NewTable("TChromiumCore_SendProcessMessageWithPIdPMessageInt64", 0), // procedure SendProcessMessageWithPIdPMessageInt64
			/* 111 */ imports.NewTable("TChromiumCore_SetFocus", 0), // procedure SetFocus
			/* 112 */ imports.NewTable("TChromiumCore_SetAccessibilityState", 0), // procedure SetAccessibilityState
			/* 113 */ imports.NewTable("TChromiumCore_DragTargetDragEnter", 0), // procedure DragTargetDragEnter
			/* 114 */ imports.NewTable("TChromiumCore_DragTargetDragOver", 0), // procedure DragTargetDragOver
			/* 115 */ imports.NewTable("TChromiumCore_DragTargetDragLeave", 0), // procedure DragTargetDragLeave
			/* 116 */ imports.NewTable("TChromiumCore_DragTargetDrop", 0), // procedure DragTargetDrop
			/* 117 */ imports.NewTable("TChromiumCore_DragSourceEndedAt", 0), // procedure DragSourceEndedAt
			/* 118 */ imports.NewTable("TChromiumCore_DragSourceSystemDragEnded", 0), // procedure DragSourceSystemDragEnded
			/* 119 */ imports.NewTable("TChromiumCore_IMESetComposition", 0), // procedure IMESetComposition
			/* 120 */ imports.NewTable("TChromiumCore_IMECommitText", 0), // procedure IMECommitText
			/* 121 */ imports.NewTable("TChromiumCore_IMEFinishComposingText", 0), // procedure IMEFinishComposingText
			/* 122 */ imports.NewTable("TChromiumCore_IMECancelComposition", 0), // procedure IMECancelComposition
			/* 123 */ imports.NewTable("TChromiumCore_ReplaceMisspelling", 0), // procedure ReplaceMisspelling
			/* 124 */ imports.NewTable("TChromiumCore_AddWordToDictionary", 0), // procedure AddWordToDictionary
			/* 125 */ imports.NewTable("TChromiumCore_UpdateBrowserSize", 0), // procedure UpdateBrowserSize
			/* 126 */ imports.NewTable("TChromiumCore_UpdateXWindowVisibility", 0), // procedure UpdateXWindowVisibility
			/* 127 */ imports.NewTable("TChromiumCore_NotifyCurrentSinks", 0), // procedure NotifyCurrentSinks
			/* 128 */ imports.NewTable("TChromiumCore_NotifyCurrentRoutes", 0), // procedure NotifyCurrentRoutes
			/* 129 */ imports.NewTable("TChromiumCore_CreateRoute", 0), // procedure CreateRoute
			/* 130 */ imports.NewTable("TChromiumCore_GetDeviceInfo", 0), // procedure GetDeviceInfo
			/* 131 */ imports.NewTable("TChromiumCore_DefaultUrl", 0), // property DefaultUrl
			/* 132 */ imports.NewTable("TChromiumCore_Options", 0), // property Options
			/* 133 */ imports.NewTable("TChromiumCore_FontOptions", 0), // property FontOptions
			/* 134 */ imports.NewTable("TChromiumCore_PDFPrintOptions", 0), // property PDFPrintOptions
			/* 135 */ imports.NewTable("TChromiumCore_DefaultEncoding", 0), // property DefaultEncoding
			/* 136 */ imports.NewTable("TChromiumCore_BrowserId", 0), // property BrowserId
			/* 137 */ imports.NewTable("TChromiumCore_Browser", 0), // property Browser
			/* 138 */ imports.NewTable("TChromiumCore_BrowserById", 0), // property BrowserById
			/* 139 */ imports.NewTable("TChromiumCore_BrowserCount", 0), // property BrowserCount
			/* 140 */ imports.NewTable("TChromiumCore_BrowserIdByIndex", 0), // property BrowserIdByIndex
			/* 141 */ imports.NewTable("TChromiumCore_CefClient", 0), // property CefClient
			/* 142 */ imports.NewTable("TChromiumCore_ReqContextHandler", 0), // property ReqContextHandler
			/* 143 */ imports.NewTable("TChromiumCore_ResourceRequestHandler", 0), // property ResourceRequestHandler
			/* 144 */ imports.NewTable("TChromiumCore_CefWindowInfo", 0), // property CefWindowInfo
			/* 145 */ imports.NewTable("TChromiumCore_VisibleNavigationEntry", 0), // property VisibleNavigationEntry
			/* 146 */ imports.NewTable("TChromiumCore_RequestContext", 0), // property RequestContext
			/* 147 */ imports.NewTable("TChromiumCore_MediaRouter", 0), // property MediaRouter
			/* 148 */ imports.NewTable("TChromiumCore_MediaObserver", 0), // property MediaObserver
			/* 149 */ imports.NewTable("TChromiumCore_MediaObserverReg", 0), // property MediaObserverReg
			/* 150 */ imports.NewTable("TChromiumCore_DevToolsMsgObserver", 0), // property DevToolsMsgObserver
			/* 151 */ imports.NewTable("TChromiumCore_DevToolsMsgObserverReg", 0), // property DevToolsMsgObserverReg
			/* 152 */ imports.NewTable("TChromiumCore_ExtensionHandler", 0), // property ExtensionHandler
			/* 153 */ imports.NewTable("TChromiumCore_MultithreadApp", 0), // property MultithreadApp
			/* 154 */ imports.NewTable("TChromiumCore_IsLoading", 0), // property IsLoading
			/* 155 */ imports.NewTable("TChromiumCore_HasDocument", 0), // property HasDocument
			/* 156 */ imports.NewTable("TChromiumCore_HasView", 0), // property HasView
			/* 157 */ imports.NewTable("TChromiumCore_HasDevTools", 0), // property HasDevTools
			/* 158 */ imports.NewTable("TChromiumCore_HasClientHandler", 0), // property HasClientHandler
			/* 159 */ imports.NewTable("TChromiumCore_HasBrowser", 0), // property HasBrowser
			/* 160 */ imports.NewTable("TChromiumCore_CanGoBack", 0), // property CanGoBack
			/* 161 */ imports.NewTable("TChromiumCore_CanGoForward", 0), // property CanGoForward
			/* 162 */ imports.NewTable("TChromiumCore_IsPopUp", 0), // property IsPopUp
			/* 163 */ imports.NewTable("TChromiumCore_WindowHandle", 0), // property WindowHandle
			/* 164 */ imports.NewTable("TChromiumCore_OpenerWindowHandle", 0), // property OpenerWindowHandle
			/* 165 */ imports.NewTable("TChromiumCore_BrowserHandle", 0), // property BrowserHandle
			/* 166 */ imports.NewTable("TChromiumCore_WidgetHandle", 0), // property WidgetHandle
			/* 167 */ imports.NewTable("TChromiumCore_RenderHandle", 0), // property RenderHandle
			/* 168 */ imports.NewTable("TChromiumCore_FrameIsFocused", 0), // property FrameIsFocused
			/* 169 */ imports.NewTable("TChromiumCore_Initialized", 0), // property Initialized
			/* 170 */ imports.NewTable("TChromiumCore_RequestContextCache", 0), // property RequestContextCache
			/* 171 */ imports.NewTable("TChromiumCore_RequestContextIsGlobal", 0), // property RequestContextIsGlobal
			/* 172 */ imports.NewTable("TChromiumCore_DocumentURL", 0), // property DocumentURL
			/* 173 */ imports.NewTable("TChromiumCore_ZoomLevel", 0), // property ZoomLevel
			/* 174 */ imports.NewTable("TChromiumCore_ZoomPct", 0), // property ZoomPct
			/* 175 */ imports.NewTable("TChromiumCore_ZoomStep", 0), // property ZoomStep
			/* 176 */ imports.NewTable("TChromiumCore_WindowlessFrameRate", 0), // property WindowlessFrameRate
			/* 177 */ imports.NewTable("TChromiumCore_CustomHeaderName", 0), // property CustomHeaderName
			/* 178 */ imports.NewTable("TChromiumCore_CustomHeaderValue", 0), // property CustomHeaderValue
			/* 179 */ imports.NewTable("TChromiumCore_DoNotTrack", 0), // property DoNotTrack
			/* 180 */ imports.NewTable("TChromiumCore_SendReferrer", 0), // property SendReferrer
			/* 181 */ imports.NewTable("TChromiumCore_HyperlinkAuditing", 0), // property HyperlinkAuditing
			/* 182 */ imports.NewTable("TChromiumCore_AllowOutdatedPlugins", 0), // property AllowOutdatedPlugins
			/* 183 */ imports.NewTable("TChromiumCore_AlwaysAuthorizePlugins", 0), // property AlwaysAuthorizePlugins
			/* 184 */ imports.NewTable("TChromiumCore_AlwaysOpenPDFExternally", 0), // property AlwaysOpenPDFExternally
			/* 185 */ imports.NewTable("TChromiumCore_SpellChecking", 0), // property SpellChecking
			/* 186 */ imports.NewTable("TChromiumCore_SpellCheckerDicts", 0), // property SpellCheckerDicts
			/* 187 */ imports.NewTable("TChromiumCore_HasValidMainFrame", 0), // property HasValidMainFrame
			/* 188 */ imports.NewTable("TChromiumCore_FrameCount", 0), // property FrameCount
			/* 189 */ imports.NewTable("TChromiumCore_DragOperations", 0), // property DragOperations
			/* 190 */ imports.NewTable("TChromiumCore_AudioMuted", 0), // property AudioMuted
			/* 191 */ imports.NewTable("TChromiumCore_SafeSearch", 0), // property SafeSearch
			/* 192 */ imports.NewTable("TChromiumCore_YouTubeRestrict", 0), // property YouTubeRestrict
			/* 193 */ imports.NewTable("TChromiumCore_PrintingEnabled", 0), // property PrintingEnabled
			/* 194 */ imports.NewTable("TChromiumCore_AcceptLanguageList", 0), // property AcceptLanguageList
			/* 195 */ imports.NewTable("TChromiumCore_AcceptCookies", 0), // property AcceptCookies
			/* 196 */ imports.NewTable("TChromiumCore_Block3rdPartyCookies", 0), // property Block3rdPartyCookies
			/* 197 */ imports.NewTable("TChromiumCore_MultiBrowserMode", 0), // property MultiBrowserMode
			/* 198 */ imports.NewTable("TChromiumCore_DefaultWindowInfoExStyle", 0), // property DefaultWindowInfoExStyle
			/* 199 */ imports.NewTable("TChromiumCore_Offline", 0), // property Offline
			/* 200 */ imports.NewTable("TChromiumCore_QuicAllowed", 0), // property QuicAllowed
			/* 201 */ imports.NewTable("TChromiumCore_JavascriptEnabled", 0), // property JavascriptEnabled
			/* 202 */ imports.NewTable("TChromiumCore_LoadImagesAutomatically", 0), // property LoadImagesAutomatically
			/* 203 */ imports.NewTable("TChromiumCore_BatterySaverModeState", 0), // property BatterySaverModeState
			/* 204 */ imports.NewTable("TChromiumCore_HighEfficiencyMode", 0), // property HighEfficiencyMode
			/* 205 */ imports.NewTable("TChromiumCore_XDisplay", 0), // property XDisplay
			/* 206 */ imports.NewTable("TChromiumCore_WebRTCIPHandlingPolicy", 0), // property WebRTCIPHandlingPolicy
			/* 207 */ imports.NewTable("TChromiumCore_WebRTCMultipleRoutes", 0), // property WebRTCMultipleRoutes
			/* 208 */ imports.NewTable("TChromiumCore_WebRTCNonproxiedUDP", 0), // property WebRTCNonproxiedUDP
			/* 209 */ imports.NewTable("TChromiumCore_ProxyType", 0), // property ProxyType
			/* 210 */ imports.NewTable("TChromiumCore_ProxyScheme", 0), // property ProxyScheme
			/* 211 */ imports.NewTable("TChromiumCore_ProxyServer", 0), // property ProxyServer
			/* 212 */ imports.NewTable("TChromiumCore_ProxyPort", 0), // property ProxyPort
			/* 213 */ imports.NewTable("TChromiumCore_ProxyUsername", 0), // property ProxyUsername
			/* 214 */ imports.NewTable("TChromiumCore_ProxyPassword", 0), // property ProxyPassword
			/* 215 */ imports.NewTable("TChromiumCore_ProxyScriptURL", 0), // property ProxyScriptURL
			/* 216 */ imports.NewTable("TChromiumCore_ProxyByPassList", 0), // property ProxyByPassList
			/* 217 */ imports.NewTable("TChromiumCore_MaxConnectionsPerProxy", 0), // property MaxConnectionsPerProxy
			/* 218 */ imports.NewTable("TChromiumCore_OnTextResultAvailable", 0), // event OnTextResultAvailable
			/* 219 */ imports.NewTable("TChromiumCore_OnPdfPrintFinished", 0), // event OnPdfPrintFinished
			/* 220 */ imports.NewTable("TChromiumCore_OnPrefsAvailable", 0), // event OnPrefsAvailable
			/* 221 */ imports.NewTable("TChromiumCore_OnPrefsUpdated", 0), // event OnPrefsUpdated
			/* 222 */ imports.NewTable("TChromiumCore_OnCookiesDeleted", 0), // event OnCookiesDeleted
			/* 223 */ imports.NewTable("TChromiumCore_OnResolvedHostAvailable", 0), // event OnResolvedHostAvailable
			/* 224 */ imports.NewTable("TChromiumCore_OnNavigationVisitorResultAvailable", 0), // event OnNavigationVisitorResultAvailable
			/* 225 */ imports.NewTable("TChromiumCore_OnDownloadImageFinished", 0), // event OnDownloadImageFinished
			/* 226 */ imports.NewTable("TChromiumCore_OnCookiesFlushed", 0), // event OnCookiesFlushed
			/* 227 */ imports.NewTable("TChromiumCore_OnCertificateExceptionsCleared", 0), // event OnCertificateExceptionsCleared
			/* 228 */ imports.NewTable("TChromiumCore_OnHttpAuthCredentialsCleared", 0), // event OnHttpAuthCredentialsCleared
			/* 229 */ imports.NewTable("TChromiumCore_OnAllConnectionsClosed", 0), // event OnAllConnectionsClosed
			/* 230 */ imports.NewTable("TChromiumCore_OnExecuteTaskOnCefThread", 0), // event OnExecuteTaskOnCefThread
			/* 231 */ imports.NewTable("TChromiumCore_OnCookiesVisited", 0), // event OnCookiesVisited
			/* 232 */ imports.NewTable("TChromiumCore_OnCookieVisitorDestroyed", 0), // event OnCookieVisitorDestroyed
			/* 233 */ imports.NewTable("TChromiumCore_OnCookieSet", 0), // event OnCookieSet
			/* 234 */ imports.NewTable("TChromiumCore_OnZoomPctAvailable", 0), // event OnZoomPctAvailable
			/* 235 */ imports.NewTable("TChromiumCore_OnMediaRouteCreateFinished", 0), // event OnMediaRouteCreateFinished
			/* 236 */ imports.NewTable("TChromiumCore_OnMediaSinkDeviceInfo", 0), // event OnMediaSinkDeviceInfo
			/* 237 */ imports.NewTable("TChromiumCore_OnBrowserCompMsg", 0), // event OnBrowserCompMsg
			/* 238 */ imports.NewTable("TChromiumCore_OnWidgetCompMsg", 0), // event OnWidgetCompMsg
			/* 239 */ imports.NewTable("TChromiumCore_OnRenderCompMsg", 0), // event OnRenderCompMsg
			/* 240 */ imports.NewTable("TChromiumCore_OnProcessMessageReceived", 0), // event OnProcessMessageReceived
			/* 241 */ imports.NewTable("TChromiumCore_OnLoadStart", 0), // event OnLoadStart
			/* 242 */ imports.NewTable("TChromiumCore_OnLoadEnd", 0), // event OnLoadEnd
			/* 243 */ imports.NewTable("TChromiumCore_OnLoadError", 0), // event OnLoadError
			/* 244 */ imports.NewTable("TChromiumCore_OnLoadingStateChange", 0), // event OnLoadingStateChange
			/* 245 */ imports.NewTable("TChromiumCore_OnTakeFocus", 0), // event OnTakeFocus
			/* 246 */ imports.NewTable("TChromiumCore_OnSetFocus", 0), // event OnSetFocus
			/* 247 */ imports.NewTable("TChromiumCore_OnGotFocus", 0), // event OnGotFocus
			/* 248 */ imports.NewTable("TChromiumCore_OnBeforeContextMenu", 0), // event OnBeforeContextMenu
			/* 249 */ imports.NewTable("TChromiumCore_OnRunContextMenu", 0), // event OnRunContextMenu
			/* 250 */ imports.NewTable("TChromiumCore_OnContextMenuCommand", 0), // event OnContextMenuCommand
			/* 251 */ imports.NewTable("TChromiumCore_OnContextMenuDismissed", 0), // event OnContextMenuDismissed
			/* 252 */ imports.NewTable("TChromiumCore_OnRunQuickMenu", 0), // event OnRunQuickMenu
			/* 253 */ imports.NewTable("TChromiumCore_OnQuickMenuCommand", 0), // event OnQuickMenuCommand
			/* 254 */ imports.NewTable("TChromiumCore_OnQuickMenuDismissed", 0), // event OnQuickMenuDismissed
			/* 255 */ imports.NewTable("TChromiumCore_OnPreKeyEvent", 0), // event OnPreKeyEvent
			/* 256 */ imports.NewTable("TChromiumCore_OnKeyEvent", 0), // event OnKeyEvent
			/* 257 */ imports.NewTable("TChromiumCore_OnAddressChange", 0), // event OnAddressChange
			/* 258 */ imports.NewTable("TChromiumCore_OnTitleChange", 0), // event OnTitleChange
			/* 259 */ imports.NewTable("TChromiumCore_OnFavIconUrlChange", 0), // event OnFavIconUrlChange
			/* 260 */ imports.NewTable("TChromiumCore_OnFullScreenModeChange", 0), // event OnFullScreenModeChange
			/* 261 */ imports.NewTable("TChromiumCore_OnTooltip", 0), // event OnTooltip
			/* 262 */ imports.NewTable("TChromiumCore_OnStatusMessage", 0), // event OnStatusMessage
			/* 263 */ imports.NewTable("TChromiumCore_OnConsoleMessage", 0), // event OnConsoleMessage
			/* 264 */ imports.NewTable("TChromiumCore_OnAutoResize", 0), // event OnAutoResize
			/* 265 */ imports.NewTable("TChromiumCore_OnLoadingProgressChange", 0), // event OnLoadingProgressChange
			/* 266 */ imports.NewTable("TChromiumCore_OnCursorChange", 0), // event OnCursorChange
			/* 267 */ imports.NewTable("TChromiumCore_OnMediaAccessChange", 0), // event OnMediaAccessChange
			/* 268 */ imports.NewTable("TChromiumCore_OnCanDownload", 0), // event OnCanDownload
			/* 269 */ imports.NewTable("TChromiumCore_OnBeforeDownload", 0), // event OnBeforeDownload
			/* 270 */ imports.NewTable("TChromiumCore_OnDownloadUpdated", 0), // event OnDownloadUpdated
			/* 271 */ imports.NewTable("TChromiumCore_OnJsdialog", 0), // event OnJsdialog
			/* 272 */ imports.NewTable("TChromiumCore_OnBeforeUnloadDialog", 0), // event OnBeforeUnloadDialog
			/* 273 */ imports.NewTable("TChromiumCore_OnResetDialogState", 0), // event OnResetDialogState
			/* 274 */ imports.NewTable("TChromiumCore_OnDialogClosed", 0), // event OnDialogClosed
			/* 275 */ imports.NewTable("TChromiumCore_OnBeforePopup", 0), // event OnBeforePopup
			/* 276 */ imports.NewTable("TChromiumCore_OnAfterCreated", 0), // event OnAfterCreated
			/* 277 */ imports.NewTable("TChromiumCore_OnBeforeClose", 0), // event OnBeforeClose
			/* 278 */ imports.NewTable("TChromiumCore_OnClose", 0), // event OnClose
			/* 279 */ imports.NewTable("TChromiumCore_OnBeforeBrowse", 0), // event OnBeforeBrowse
			/* 280 */ imports.NewTable("TChromiumCore_OnOpenUrlFromTab", 0), // event OnOpenUrlFromTab
			/* 281 */ imports.NewTable("TChromiumCore_OnGetAuthCredentials", 0), // event OnGetAuthCredentials
			/* 282 */ imports.NewTable("TChromiumCore_OnCertificateError", 0), // event OnCertificateError
			/* 283 */ imports.NewTable("TChromiumCore_OnSelectClientCertificate", 0), // event OnSelectClientCertificate
			/* 284 */ imports.NewTable("TChromiumCore_OnRenderViewReady", 0), // event OnRenderViewReady
			/* 285 */ imports.NewTable("TChromiumCore_OnRenderProcessTerminated", 0), // event OnRenderProcessTerminated
			/* 286 */ imports.NewTable("TChromiumCore_OnGetResourceRequestHandler_ReqHdlr", 0), // event OnGetResourceRequestHandler_ReqHdlr
			/* 287 */ imports.NewTable("TChromiumCore_OnDocumentAvailableInMainFrame", 0), // event OnDocumentAvailableInMainFrame
			/* 288 */ imports.NewTable("TChromiumCore_OnBeforeResourceLoad", 0), // event OnBeforeResourceLoad
			/* 289 */ imports.NewTable("TChromiumCore_OnGetResourceHandler", 0), // event OnGetResourceHandler
			/* 290 */ imports.NewTable("TChromiumCore_OnResourceRedirect", 0), // event OnResourceRedirect
			/* 291 */ imports.NewTable("TChromiumCore_OnResourceResponse", 0), // event OnResourceResponse
			/* 292 */ imports.NewTable("TChromiumCore_OnGetResourceResponseFilter", 0), // event OnGetResourceResponseFilter
			/* 293 */ imports.NewTable("TChromiumCore_OnResourceLoadComplete", 0), // event OnResourceLoadComplete
			/* 294 */ imports.NewTable("TChromiumCore_OnProtocolExecution", 0), // event OnProtocolExecution
			/* 295 */ imports.NewTable("TChromiumCore_OnCanSendCookie", 0), // event OnCanSendCookie
			/* 296 */ imports.NewTable("TChromiumCore_OnCanSaveCookie", 0), // event OnCanSaveCookie
			/* 297 */ imports.NewTable("TChromiumCore_OnFileDialog", 0), // event OnFileDialog
			/* 298 */ imports.NewTable("TChromiumCore_OnGetAccessibilityHandler", 0), // event OnGetAccessibilityHandler
			/* 299 */ imports.NewTable("TChromiumCore_OnGetRootScreenRect", 0), // event OnGetRootScreenRect
			/* 300 */ imports.NewTable("TChromiumCore_OnGetViewRect", 0), // event OnGetViewRect
			/* 301 */ imports.NewTable("TChromiumCore_OnGetScreenPoint", 0), // event OnGetScreenPoint
			/* 302 */ imports.NewTable("TChromiumCore_OnGetScreenInfo", 0), // event OnGetScreenInfo
			/* 303 */ imports.NewTable("TChromiumCore_OnPopupShow", 0), // event OnPopupShow
			/* 304 */ imports.NewTable("TChromiumCore_OnPopupSize", 0), // event OnPopupSize
			/* 305 */ imports.NewTable("TChromiumCore_OnPaint", 0), // event OnPaint
			/* 306 */ imports.NewTable("TChromiumCore_OnAcceleratedPaint", 0), // event OnAcceleratedPaint
			/* 307 */ imports.NewTable("TChromiumCore_OnGetTouchHandleSize", 0), // event OnGetTouchHandleSize
			/* 308 */ imports.NewTable("TChromiumCore_OnTouchHandleStateChanged", 0), // event OnTouchHandleStateChanged
			/* 309 */ imports.NewTable("TChromiumCore_OnStartDragging", 0), // event OnStartDragging
			/* 310 */ imports.NewTable("TChromiumCore_OnUpdateDragCursor", 0), // event OnUpdateDragCursor
			/* 311 */ imports.NewTable("TChromiumCore_OnScrollOffsetChanged", 0), // event OnScrollOffsetChanged
			/* 312 */ imports.NewTable("TChromiumCore_OnIMECompositionRangeChanged", 0), // event OnIMECompositionRangeChanged
			/* 313 */ imports.NewTable("TChromiumCore_OnTextSelectionChanged", 0), // event OnTextSelectionChanged
			/* 314 */ imports.NewTable("TChromiumCore_OnVirtualKeyboardRequested", 0), // event OnVirtualKeyboardRequested
			/* 315 */ imports.NewTable("TChromiumCore_OnDragEnter", 0), // event OnDragEnter
			/* 316 */ imports.NewTable("TChromiumCore_OnDraggableRegionsChanged", 0), // event OnDraggableRegionsChanged
			/* 317 */ imports.NewTable("TChromiumCore_OnFindResult", 0), // event OnFindResult
			/* 318 */ imports.NewTable("TChromiumCore_OnRequestContextInitialized", 0), // event OnRequestContextInitialized
			/* 319 */ imports.NewTable("TChromiumCore_OnGetResourceRequestHandler_ReqCtxHdlr", 0), // event OnGetResourceRequestHandler_ReqCtxHdlr
			/* 320 */ imports.NewTable("TChromiumCore_OnSinks", 0), // event OnSinks
			/* 321 */ imports.NewTable("TChromiumCore_OnRoutes", 0), // event OnRoutes
			/* 322 */ imports.NewTable("TChromiumCore_OnRouteStateChanged", 0), // event OnRouteStateChanged
			/* 323 */ imports.NewTable("TChromiumCore_OnRouteMessageReceived", 0), // event OnRouteMessageReceived
			/* 324 */ imports.NewTable("TChromiumCore_OnGetAudioParameters", 0), // event OnGetAudioParameters
			/* 325 */ imports.NewTable("TChromiumCore_OnAudioStreamStarted", 0), // event OnAudioStreamStarted
			/* 326 */ imports.NewTable("TChromiumCore_OnAudioStreamPacket", 0), // event OnAudioStreamPacket
			/* 327 */ imports.NewTable("TChromiumCore_OnAudioStreamStopped", 0), // event OnAudioStreamStopped
			/* 328 */ imports.NewTable("TChromiumCore_OnAudioStreamError", 0), // event OnAudioStreamError
			/* 329 */ imports.NewTable("TChromiumCore_OnDevToolsMessage", 0), // event OnDevToolsMessage
			/* 330 */ imports.NewTable("TChromiumCore_OnDevToolsRawMessage", 0), // event OnDevToolsRawMessage
			/* 331 */ imports.NewTable("TChromiumCore_OnDevToolsMethodResult", 0), // event OnDevToolsMethodResult
			/* 332 */ imports.NewTable("TChromiumCore_OnDevToolsMethodRawResult", 0), // event OnDevToolsMethodRawResult
			/* 333 */ imports.NewTable("TChromiumCore_OnDevToolsEvent", 0), // event OnDevToolsEvent
			/* 334 */ imports.NewTable("TChromiumCore_OnDevToolsRawEvent", 0), // event OnDevToolsRawEvent
			/* 335 */ imports.NewTable("TChromiumCore_OnDevToolsAgentAttached", 0), // event OnDevToolsAgentAttached
			/* 336 */ imports.NewTable("TChromiumCore_OnDevToolsAgentDetached", 0), // event OnDevToolsAgentDetached
			/* 337 */ imports.NewTable("TChromiumCore_OnExtensionLoadFailed", 0), // event OnExtensionLoadFailed
			/* 338 */ imports.NewTable("TChromiumCore_OnExtensionLoaded", 0), // event OnExtensionLoaded
			/* 339 */ imports.NewTable("TChromiumCore_OnExtensionUnloaded", 0), // event OnExtensionUnloaded
			/* 340 */ imports.NewTable("TChromiumCore_OnExtensionBeforeBackgroundBrowser", 0), // event OnExtensionBeforeBackgroundBrowser
			/* 341 */ imports.NewTable("TChromiumCore_OnExtensionBeforeBrowser", 0), // event OnExtensionBeforeBrowser
			/* 342 */ imports.NewTable("TChromiumCore_OnExtensionGetActiveBrowser", 0), // event OnExtensionGetActiveBrowser
			/* 343 */ imports.NewTable("TChromiumCore_OnExtensionCanAccessBrowser", 0), // event OnExtensionCanAccessBrowser
			/* 344 */ imports.NewTable("TChromiumCore_OnExtensionGetExtensionResource", 0), // event OnExtensionGetExtensionResource
			/* 345 */ imports.NewTable("TChromiumCore_OnPrintStart", 0), // event OnPrintStart
			/* 346 */ imports.NewTable("TChromiumCore_OnPrintSettings", 0), // event OnPrintSettings
			/* 347 */ imports.NewTable("TChromiumCore_OnPrintDialog", 0), // event OnPrintDialog
			/* 348 */ imports.NewTable("TChromiumCore_OnPrintJob", 0), // event OnPrintJob
			/* 349 */ imports.NewTable("TChromiumCore_OnPrintReset", 0), // event OnPrintReset
			/* 350 */ imports.NewTable("TChromiumCore_OnGetPDFPaperSize", 0), // event OnGetPDFPaperSize
			/* 351 */ imports.NewTable("TChromiumCore_OnFrameCreated", 0), // event OnFrameCreated
			/* 352 */ imports.NewTable("TChromiumCore_OnFrameAttached", 0), // event OnFrameAttached
			/* 353 */ imports.NewTable("TChromiumCore_OnFrameDetached", 0), // event OnFrameDetached
			/* 354 */ imports.NewTable("TChromiumCore_OnMainFrameChanged", 0), // event OnMainFrameChanged
			/* 355 */ imports.NewTable("TChromiumCore_OnChromeCommand", 0), // event OnChromeCommand
			/* 356 */ imports.NewTable("TChromiumCore_OnRequestMediaAccessPermission", 0), // event OnRequestMediaAccessPermission
			/* 357 */ imports.NewTable("TChromiumCore_OnShowPermissionPrompt", 0), // event OnShowPermissionPrompt
			/* 358 */ imports.NewTable("TChromiumCore_OnDismissPermissionPrompt", 0), // event OnDismissPermissionPrompt
		}
	})
	return chromiumCoreImport
}
