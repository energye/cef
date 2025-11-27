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

	cefTypes "github.com/energye/cef/types"
)

// IChromiumCore Parent: lcl.IComponent
type IChromiumCore interface {
	lcl.IComponent
	// CreateClientHandlerWithBool
	//  Used to create the client handler which will also create most of the browser handlers needed for the browser.
	CreateClientHandlerWithBool(isOSR bool) bool // function
	// CreateClientHandlerWithClientBool
	//  Used to create the client handler when a browser requests a new browser in a popup window or tab in the TChromiumCore.OnBeforePopup event.
	CreateClientHandlerWithClientBool(client *IEngClient, isOSR bool) bool // function
	// TryCloseBrowser
	//  Helper for closing a browser. Call this function from the top-level window
	//  close handler (if any). Internally this calls CloseBrowser(false (0)) if
	//  the close has not yet been initiated. This function returns false (0)
	//  while the close is pending and true (1) after the close has completed. See
	//  CloseBrowser() and ICefLifeSpanHandler.DoClose() documentation for
	//  additional usage information. This function must be called on the browser
	//  process UI thread.
	TryCloseBrowser() bool // function
	// SelectBrowser
	//  Select the browser with the aID identifier when TChromiumCore uses the
	//  multi-browser mode.
	SelectBrowser(iD int32) bool // function
	// IndexOfBrowserID
	//  Returns the index in the browsers array of the browser with the aID
	//  identifier when TChromiumCore uses the multi-browser mode.
	IndexOfBrowserID(iD int32) int32 // function
	// ShareRequestContext
	//  Creates a new request context in the aContext parameter that shares
	//  storage with the request context of the current browser and uses an
	//  optional handler.
	ShareRequestContext(context *ICefRequestContext, handler IEngRequestContextHandler) bool // function
	// SetNewBrowserParent
	//  Used to reparent the browser to a different TCEFWindowParent.
	SetNewBrowserParent(newParentHwnd types.HWND) bool // function
	// CreateBrowserWithWindowHandleRectStringRequestContextDictionaryValueBool
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowserWithWindowHandleRectStringRequestContextDictionaryValueBool(parentHandle cefTypes.TCefWindowHandle, parentRect types.TRect, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue, forceAsPopup bool) bool // function
	// CreateBrowserWithStringBrowserViewComponentRequestContextDictionaryValue
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowserWithStringBrowserViewComponentRequestContextDictionaryValue(uRL string, browserViewComp ICEFBrowserViewComponent, context ICefRequestContext, extraInfo ICefDictionaryValue) bool // function
	// ClearCertificateExceptions
	//  Clears all certificate exceptions that were added as part of handling
	//  OnCertificateError. If you call this it is recommended that you also call
	//  CloseAllConnections() or you risk not being prompted again for server
	//  certificates if you reconnect quickly.
	//  If aClearImmediately is false then OnCertificateExceptionsCleared is
	//  triggered when the exceptions are cleared.
	ClearCertificateExceptions(clearImmediately bool) bool // function
	// ClearHttpAuthCredentials
	//  Clears all HTTP authentication credentials that were added as part of
	//  handling GetAuthCredentials. If |callback| is non-NULL it will be executed
	//  on the UI thread after completion.
	//  If aClearImmediately is false then OnHttpAuthCredentialsCleared is triggered
	//  when the credeintials are cleared.
	ClearHttpAuthCredentials(clearImmediately bool) bool // function
	// CloseAllConnections
	//  Clears all active and idle connections that Chromium currently has. This
	//  is only recommended if you have released all other CEF objects but don't
	//  yet want to call cef_shutdown().
	CloseAllConnections(closeImmediately bool) bool // function
	// GetFrameNames
	//  Returns the names of all existing frames.
	GetFrameNames(frameNames *lcl.IStrings) bool // function
	// GetFrameIdentifiers
	//  Returns the identifiers of all existing frames.
	GetFrameIdentifiers(frameIdentifiers *lcl.IStrings) bool // function
	// IsSameBrowser
	//  Used to check if the browser parameter is the same as the selected browser in TChromiumCore.
	IsSameBrowser(browser ICefBrowser) bool // function
	// ExecuteTaskOnCefThread
	//  Calling ExecuteTaskOnCefThread function will trigger the TChromiumCore.OnExecuteTaskOnCefThread event.
	//  <param name="aCefThreadId">Indicates the CEF thread on which TChromiumCore.OnExecuteTaskOnCefThread will be executed.</param>
	//  <param name="aTaskID">Custom ID used to identify the task that triggered the TChromiumCore.OnExecuteTaskOnCefThread event.</param>
	//  <param name="aDelayMs">Optional delay in milliseconds to trigger the TChromiumCore.OnExecuteTaskOnCefThread event.</param>
	ExecuteTaskOnCefThread(cefThreadId cefTypes.TCefThreadId, taskID uint32, delayMs int64) bool // function
	// DeleteCookies
	//  Used to delete cookies immediately or asynchronously. If aDeleteImmediately is false TChromiumCore.DeleteCookies triggers
	//  the TChromiumCore.OnCookiesDeleted event when the cookies are deleted.
	DeleteCookies(url string, cookieName string, deleteImmediately bool) bool // function
	// VisitAllCookies
	//  TChromiumCore.VisitAllCookies triggers the TChromiumCore.OnCookiesVisited event for each cookie
	//  aID is an optional parameter to identify which VisitAllCookies call has triggered the
	//  OnCookiesVisited event.
	//  TChromiumCore.OnCookiesVisited may not be triggered if the cookie store is empty but the
	//  TChromium.OnCookieVisitorDestroyed event will always be triggered to signal when the browser
	//  when the visit is over.
	VisitAllCookies(iD int32) bool // function
	// VisitURLCookies
	//  TChromiumCore.VisitURLCookies triggers the TChromiumCore.OnCookiesVisited event for each cookie
	//  aID is an optional parameter to identify which VisitURLCookies call has triggered the
	//  OnCookiesVisited event.
	//  TChromiumCore.OnCookiesVisited may not be triggered if the cookie store is empty but the
	//  TChromium.OnCookieVisitorDestroyed event will always be triggered to signal when the browser
	//  when the visit is over.
	VisitURLCookies(url string, includeHttpOnly bool, iD int32) bool // function
	// SetCookie
	//  TChromiumCore.SetCookie triggers the TChromiumCore.OnCookieSet event when the cookie has been set
	//  aID is an optional parameter to identify which SetCookie call has triggered the
	//  OnCookieSet event.
	SetCookie(url string, name string, value string, domain string, path string, secure bool, httponly bool, hasExpires bool, creation types.TDateTime, lastAccess types.TDateTime, expires types.TDateTime, sameSite cefTypes.TCefCookieSameSite, priority cefTypes.TCefCookiePriority, setImmediately bool, iD int32) bool // function
	// FlushCookieStore
	//  Flush the backing store (if any) to disk.
	//  <param name="aFlushImmediately">If aFlushImmediately is false the cookies will be flushed on the CEF UI thread and the OnCookiesFlushed event will be triggered.</param>
	//  <returns>Returns false (0) if cookies cannot be accessed.</returns>
	FlushCookieStore(flushImmediately bool) bool // function
	// SendDevToolsMessage
	//  Send a function call message over the DevTools protocol. |message_| must be
	//  a UTF8-encoded JSON dictionary that contains "id" (int), "function"
	//  (string) and "params" (dictionary, optional) values. See the DevTools
	//  protocol documentation at https://chromedevtools.github.io/devtools-
	//  protocol/ for details of supported functions and the expected "params"
	//  dictionary contents. |message_| will be copied if necessary. This function
	//  will return true (1) if called on the UI thread and the message was
	//  successfully submitted for validation, otherwise false (0). Validation
	//  will be applied asynchronously and any messages that fail due to
	//  formatting errors or missing parameters may be discarded without
	//  notification. Prefer ExecuteDevToolsMethod if a more structured approach
	//  to message formatting is desired.
	//  Every valid function call will result in an asynchronous function result
	//  or error message that references the sent message "id". Event messages are
	//  received while notifications are enabled (for example, between function
	//  calls for "Page.enable" and "Page.disable"). All received messages will be
	//  delivered to the observer(s) registered with AddDevToolsMessageObserver.
	//  See ICefDevToolsMessageObserver.OnDevToolsMessage documentation for
	//  details of received message contents.
	//  Usage of the SendDevToolsMessage, ExecuteDevToolsMethod and
	//  AddDevToolsMessageObserver functions does not require an active DevTools
	//  front-end or remote-debugging session. Other active DevTools sessions will
	//  continue to function independently. However, any modification of global
	//  browser state by one session may not be reflected in the UI of other
	//  sessions.
	//  Communication with the DevTools front-end (when displayed) can be logged
	//  for development purposes by passing the `--devtools-protocol-log-
	//  file=<path>` command-line flag.
	SendDevToolsMessage(message string) bool // function
	// ExecuteDevToolsMethod
	//  Execute a function call over the DevTools protocol. This is a more
	//  structured version of SendDevToolsMessage. |message_id| is an incremental
	//  number that uniquely identifies the message (pass 0 to have the next
	//  number assigned automatically based on previous values). |function| is the
	//  function name. |params| are the function parameters, which may be NULL.
	//  See the DevTools protocol documentation (linked above) for details of
	//  supported functions and the expected |params| dictionary contents. This
	//  function will return the assigned message ID if called on the UI thread
	//  and the message was successfully submitted for validation, otherwise 0.
	//  See the SendDevToolsMessage documentation for additional usage
	//  information.
	ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32 // function
	// AddDevToolsMessageObserver
	//  Add an observer for DevTools protocol messages (function results and
	//  events). The observer will remain registered until the returned
	//  Registration object is destroyed. See the SendDevToolsMessage
	//  documentation for additional usage information.
	AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) ICefRegistration // function
	// CanExecuteChromeCommand
	//  Returns true (1) if a Chrome command is supported and enabled. Values for
	//  |command_id| can be found in the cef_command_ids.h file. This function can
	//  only be called on the UI thread. Only used with the Chrome runtime.
	//  <see cref="uCEFConstants">See the IDC_* constants in uCEFConstants.pas for all the |command_id| values.</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/app/chrome_command_ids.h">The command_id values are also available in chrome/app/chrome_command_ids.h</see>
	CanExecuteChromeCommand(commandId int32) bool // function
	// CreateUrlRequestWithRequestUrlrequestClientStringX2
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently (see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  <code>
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//  </code>
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequestWithRequestUrlrequestClientStringX2(request ICefRequest, client IEngUrlrequestClient, frameName string, frameIdentifier string) ICefUrlRequest // function
	// CreateUrlRequestWithRequestUrlrequestClientFrame
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently (see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  <code>
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//  </code>
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequestWithRequestUrlrequestClientFrame(request ICefRequest, client IEngUrlrequestClient, frame ICefFrame) ICefUrlRequest // function
	// AddObserver
	//  Add an observer for MediaRouter events. The observer will remain
	//  registered until the returned Registration object is destroyed.
	AddObserver(observer IEngMediaObserver) ICefRegistration // function
	// GetSource
	//  Returns a MediaSource object for the specified media source URN. Supported
	//  URN schemes include "cast:" and "dial:", and will be already known by the
	//  client application (e.g. "cast:<appId>?clientId=<clientId>").
	GetSource(urn string) ICefMediaSource // function
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
	// CloseBrowser
	//  Request that the browser close. The JavaScript 'onbeforeunload' event will
	//  be fired. If |aForceClose| is false (0) the event handler, if any, will be
	//  allowed to prompt the user and the user can optionally cancel the close.
	//  If |aForceClose| is true (1) the prompt will not be displayed and the
	//  close will proceed. Results in a call to
	//  ICefLifeSpanHandler.DoClose() if the event handler allows the close
	//  or if |aForceClose| is true (1). See ICefLifeSpanHandler.DoClose()
	//  documentation for additional usage information.
	CloseBrowser(forceClose bool) // procedure
	// CloseAllBrowsers
	//  Calls CloseBrowser for all the browsers handled by this TChromiumCore instance.
	CloseAllBrowsers() // procedure
	// InitializeDragAndDrop
	//  Used with browsers in OSR mode to initialize drag and drop in Windows.
	InitializeDragAndDrop(dropTargetWnd types.HWND) // procedure
	// ShutdownDragAndDrop
	//  Used with browsers in OSR mode to shutdown drag and drop in Windows.
	ShutdownDragAndDrop() // procedure
	// LoadURLWithStringX3
	//  Used to navigate to a URL in the specified frame or the main frame.
	LoadURLWithStringX3(uRL string, frameName string, frameIdentifier string) // procedure
	// LoadURLWithStringFrame
	//  Used to navigate to a URL in the specified frame or the main frame.
	LoadURLWithStringFrame(uRL string, frame ICefFrame) // procedure
	// LoadStringWithStringX3
	//  Used to load a DATA URI with the HTML string contents in the specified frame or the main frame.
	LoadStringWithStringX3(hTML string, frameName string, frameIdentifier string) // procedure
	// LoadStringWithStringFrame
	//  Used to load a DATA URI with the HTML string contents in the specified frame or the main frame.
	LoadStringWithStringFrame(hTML string, frame ICefFrame) // procedure
	// LoadResourceWithCustomMemoryStreamStringX4
	//  Used to load a DATA URI with the stream contents in the specified frame or the main frame.
	//  The DATA URI will be configured with the mime type and charset specified in the parameters.
	LoadResourceWithCustomMemoryStreamStringX4(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameName string, frameIdentifier string) // procedure
	// LoadResourceWithCustomMemoryStreamStringX2Frame
	//  Used to load a DATA URI with the stream contents in the specified frame or the main frame.
	//  The DATA URI will be configured with the mime type and charset specified in the parameters.
	LoadResourceWithCustomMemoryStreamStringX2Frame(stream lcl.ICustomMemoryStream, mimeType string, charset string, frame ICefFrame) // procedure
	// LoadRequest
	//  Load the request represented by the aRequest object.
	//  WARNING: This function will fail with bad IPC message reason
	//  INVALID_INITIATOR_ORIGIN (213) unless you first navigate to the request
	//  origin using some other mechanism (LoadURL, link click, etc).
	LoadRequest(request ICefRequest) // procedure
	// GoBack
	//  Navigate backwards.
	GoBack() // procedure
	// GoForward
	//  Navigate forwards.
	GoForward() // procedure
	// Reload
	//  Reload the current page.
	Reload() // procedure
	// ReloadIgnoreCache
	//  Reload the current page ignoring any cached data.
	ReloadIgnoreCache() // procedure
	// StopLoad
	//  Stop loading the page.
	StopLoad() // procedure
	// StartDownload
	//  Starts downloading a file in the specified URL.
	StartDownload(uRL string) // procedure
	// DownloadImage
	//  Starts downloading an image in the specified URL.
	//  Use the TChromiumCore.OnDownloadImageFinished event to receive the image.
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool) // procedure
	// SimulateMouseWheel
	//  Calls ICefBrowserHost.SendMouseWheelEvent to simulate a simple mouse wheel event.
	//  Use TChromiumCore.SendMouseWheelEvent if you need to specify the mouse coordinates or the event flags.
	SimulateMouseWheel(deltaX int32, deltaY int32) // procedure
	// SimulateKeyEvent
	//  Dispatches a key event to the page using the "Input.dispatchKeyEvent"
	//  DevTools method. The browser has to be focused before simulating any
	//  key event.
	//  <param name="type_">Type of the key event.</param>
	//  <param name="modifiers">Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8.(default: 0).</param>
	//  <param name="timestamp">Time at which the event occurred.</param>
	//  <param name="text">Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events.(default: "")</param>
	//  <param name="unmodifiedtext">Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling.(default: "").</param>
	//  <param name="keyIdentifier">Unique key identifier (e.g., 'U+0041').(default: "").</param>
	//  <param name="code">Unique DOM defined string value for each physical key (e.g., 'KeyA').(default: "").</param>
	//  <param name="key">Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr').(default: "").</param>
	//  <param name="windowsVirtualKeyCode">Windows virtual key code.(default: 0).</param>
	//  <param name="nativeVirtualKeyCode">Native virtual key code.(default: 0).</param>
	//  <param name="autoRepeat">Whether the event was generated from auto repeat.(default: false).</param>
	//  <param name="isKeypad">Whether the event was generated from the keypad.(default: false).</param>
	//  <param name="isSystemKey">Whether the event was a system key event.(default: false).</param>
	//  <param name="location">Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right.(default: 0).</param>
	//  <param name="commands">Editing commands to send with the key event (e.g., 'selectAll') (default: []). These are related to but not equal the command names used in document.execCommand and NSStandardKeyBindingResponding. See https://source.chromium.org/chromium/chromium/src/+/main:third_party/blink/renderer/core/editing/commands/editor_command_names.h for valid command names.</param>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	SimulateKeyEvent(type_ cefTypes.TSimulatedCefKeyEventType, modifiers int32, timestamp float32, text string, unmodifiedtext string, keyIdentifier string, code string, key string, windowsVirtualKeyCode int32, nativeVirtualKeyCode int32, autoRepeat bool, isKeypad bool, isSystemKey bool, location cefTypes.TCefKeyLocation, commands cefTypes.TCefEditingCommand) // procedure
	// SimulateMouseEvent
	//  Dispatches a key event to the page using the "Input.dispatchKeyEvent"
	//  DevTools method. The browser has to be focused before simulating any
	//  key event.
	//  <param name="type_">Type of the mouse event.</param>
	//  <param name="x">X coordinate of the event relative to the main frame's viewport in CSS pixels.</param>
	//  <param name="y">Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.</param>
	//  <param name="modifiers">Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0). See the CEF_MOUSETOUCH_EVENT_MODIFIERS_* constants.</param>
	//  <param name="timestamp">Time at which the event occurred.</param>
	//  <param name="button">Mouse button (default: "none").</param>
	//  <param name="buttons">A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.</param>
	//  <param name="clickCount">Number of times the mouse button was clicked (default: 0).</param>
	//  <param name="force">The normalized pressure, which has a range of [0,1] (default: 0). </param>
	//  <param name="tangentialPressure">The normalized tangential pressure, which has a range of [-1,1] (default: 0).</param>
	//  <param name="tiltX">The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0).</param>
	//  <param name="tiltY">The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).</param>
	//  <param name="twist">The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).</param>
	//  <param name="deltaX">X delta in CSS pixels for mouse wheel event (default: 0).</param>
	//  <param name="deltaY">Y delta in CSS pixels for mouse wheel event (default: 0).</param>
	//  <param name="pointerType">Pointer type (default: "mouse").</param>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	SimulateMouseEvent(type_ cefTypes.TCefSimulatedMouseEventType, X float32, Y float32, modifiers int32, timestamp float32, button cefTypes.TCefSimulatedMouseButton, buttons int32, clickCount int32, force float32, tangentialPressure float32, tiltX float32, tiltY float32, twist int32, deltaX float32, deltaY float32, pointerType cefTypes.TCefSimulatedPointerType) // procedure
	// SimulateEditingCommand
	//  Simulate editing commands using the "Input.dispatchKeyEvent" DevTools method.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</see>
	SimulateEditingCommand(command cefTypes.TCefEditingCommand) // procedure
	// RetrieveHTMLWithStringX2
	//  Retrieve all the HTML content from the specified frame or the main frame.
	//  Leave aFrameName empty to get the HTML source from the main frame.
	//  It uses a CefStringVisitor to get the HTML content asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveHTMLWithStringX2(frameName string, frameIdentifier string) // procedure
	// RetrieveHTMLWithFrame
	//  Retrieve all the HTML content from the specified frame or the main frame.
	//  Set aFrame to nil to get the HTML source from the main frame.
	//  It uses a CefStringVisitor to get the HTML content asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveHTMLWithFrame(frame ICefFrame) // procedure
	// RetrieveTextWithStringX2
	//  Retrieve all the text content from the specified frame or the main frame.
	//  Leave aFrameName empty to get the text from the main frame.
	//  It uses a CefStringVisitor to get the text asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveTextWithStringX2(frameName string, frameIdentifier string) // procedure
	// RetrieveTextWithFrame
	//  Retrieve all the text content from the specified frame or the main frame.
	//  Set aFrame to nil to get the text from the main frame.
	//  It uses a CefStringVisitor to get the text asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveTextWithFrame(frame ICefFrame) // procedure
	// GetNavigationEntries
	//  Retrieve a snapshot of current navigation entries asynchronously. The
	//  TChromiumCore.OnNavigationVisitorResultAvailable event will be triggered
	//  for each navigation entry.
	GetNavigationEntries(currentOnly bool) // procedure
	// ExecuteJavaScriptWithStringX4Int
	//  Execute a string of JavaScript code in the specified frame or the main frame.
	//  <param name="aCode">JavaScript code.</param>
	//  <param name="aScriptURL">The URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error.</param>
	//  <param name="aFrameName">Name of the frame where the JavaScript code will be executed. This name is generated automatically by Chromium. See ICefBrowser.GetFrameNames.</param>
	//  <param name="aFrameIdentifier">Identifier of the frame where the JavaScript code will be executed.</param>
	//  <param name="aStartLine">The base line number to use for error reporting.</param>
	ExecuteJavaScriptWithStringX4Int(code string, scriptURL string, frameName string, frameIdentifier string, startLine int32) // procedure
	// ExecuteJavaScriptWithStringX2FrameInt
	//  Execute a string of JavaScript code in the specified frame or the main frame.
	//  <param name="aCode">JavaScript code.</param>
	//  <param name="aScriptURL">The URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error.</param>
	//  <param name="aFrame">Frame where the JavaScript code will be executed.</param>
	//  <param name="aStartLine">The base line number to use for error reporting.</param>
	ExecuteJavaScriptWithStringX2FrameInt(code string, scriptURL string, frame ICefFrame, startLine int32) // procedure
	// UpdatePreferences
	//  Used to update the browser preferences using the TChromiumCore property values asynchronously.
	UpdatePreferences() // procedure
	// SavePreferences
	//  Save the browser preferences as a text file.
	SavePreferences(fileName string) // procedure
	// ResolveHost
	//  Calls CefRequestContext.ResolveHost to resolve the domain in the URL parameter
	//  to a list of IP addresses.
	//  The result will be received in the TChromiumCore.OnResolvedHostAvailable event.
	ResolveHost(uRL string) // procedure
	// SetUserAgentOverride
	//  This procedure calls the Emulation.setUserAgentOverride DevTools method to override the user agent string.
	SetUserAgentOverride(userAgent string, acceptLanguage string, platform string) // procedure
	// ClearDataForOrigin
	//  This procedure calls the Storage.clearDataForOrigin DevTools method to clear the storage data for a given origin.
	ClearDataForOrigin(origin string, storageTypes cefTypes.TCefClearDataStorageTypes) // procedure
	// ClearCache
	//  This procedure calls the Network.clearBrowserCache DevTools method to clear the cache data.
	ClearCache() // procedure
	// ToggleAudioMuted
	//  Enable or disable the browser's audio.
	ToggleAudioMuted() // procedure
	// ShowDevTools
	//  Open developer tools (DevTools) in its own browser. If inspectElementAt has a valid point
	//  with coordinates different than low(integer) then the element at the specified location
	//  will be inspected. If the DevTools browser is already open then it will be focused.
	ShowDevTools(inspectElementAt types.TPoint, windowInfo TCefWindowInfo) // procedure
	// CloseDevTools
	//  close the developer tools.
	CloseDevTools() // procedure
	// CloseDevToolsWithWindowHandle
	//  close the developer tools.
	CloseDevToolsWithWindowHandle(devToolsWnd cefTypes.TCefWindowHandle) // procedure
	// Find
	//  Search for |searchText|. |forward| indicates whether to search forward or
	//  backward within the page. |matchCase| indicates whether the search should
	//  be case-sensitive. |findNext| indicates whether this is the first request
	//  or a follow-up. The search will be restarted if |searchText| or
	//  |matchCase| change. The search will be stopped if |searchText| is NULL.
	//  OnFindResult will be triggered to report find results.
	Find(searchText string, forward bool, matchCase bool, findNext bool) // procedure
	// StopFinding
	//  Cancel all searches that are currently going on.
	StopFinding(clearSelection bool) // procedure
	// Print
	//  Print the current browser contents.
	Print() // procedure
	// PrintToPDF
	//  Print the current browser contents to the PDF file specified by |path| and
	//  execute |callback| on completion. The caller is responsible for deleting
	//  |path| when done. For PDF printing to work on Linux you must implement the
	//  ICefPrintHandler.GetPdfPaperSize function.
	//  The TChromiumCore.OnPdfPrintFinished event will be triggered when the PDF
	//  file is created.
	PrintToPDF(filePath string) // procedure
	// ClipboardCopy
	//  Execute copy on the focused frame.
	ClipboardCopy() // procedure
	// ClipboardPaste
	//  Execute paste on the focused frame.
	ClipboardPaste() // procedure
	// ClipboardCut
	//  Execute cut on the focused frame.
	ClipboardCut() // procedure
	// ClipboardUndo
	//  Execute undo on the focused frame.
	ClipboardUndo() // procedure
	// ClipboardRedo
	//  Execute redo on the focused frame.
	ClipboardRedo() // procedure
	// ClipboardDel
	//  Execute delete on the focused frame.
	ClipboardDel() // procedure
	// SelectAll
	//  Execute select all on the focused frame.
	SelectAll() // procedure
	// IncZoomStep
	//  Increase the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	IncZoomStep() // procedure
	// DecZoomStep
	//  Decrease the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	DecZoomStep() // procedure
	// IncZoomPct
	//  Increase the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	IncZoomPct() // procedure
	// DecZoomPct
	//  Decrease the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	DecZoomPct() // procedure
	// ResetZoomStep
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomStep() // procedure
	// ResetZoomLevel
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomLevel() // procedure
	// ResetZoomPct
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomPct() // procedure
	// ReadZoom
	//  Read the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the zoom value.
	ReadZoom() // procedure
	// IncZoomCommand
	//  Execute a zoom IN command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	IncZoomCommand() // procedure
	// DecZoomCommand
	//  Execute a zoom OUT command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	DecZoomCommand() // procedure
	// ResetZoomCommand
	//  Execute a zoom RESET command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	ResetZoomCommand() // procedure
	// WasResized
	//  Notify the browser that the widget has been resized. The browser will
	//  first call ICefRenderHandler.GetViewRect to get the new size and then
	//  call ICefRenderHandler.OnPaint asynchronously with the updated
	//  regions. This function is only used when window rendering is disabled.
	WasResized() // procedure
	// WasHidden
	//  Notify the browser that it has been hidden or shown. Layouting and
	//  ICefRenderHandler.OnPaint notification will stop when the browser is
	//  hidden. This function is only used when window rendering is disabled.
	WasHidden(hidden bool) // procedure
	// NotifyScreenInfoChanged
	//  Send a notification to the browser that the screen info has changed. The
	//  browser will then call ICefRenderHandler.GetScreenInfo to update the
	//  screen information with the new values. This simulates moving the webview
	//  window from one display to another, or changing the properties of the
	//  current display. This function is only used when window rendering is
	//  disabled.
	NotifyScreenInfoChanged() // procedure
	// NotifyMoveOrResizeStarted
	//  Notify the browser that the window hosting it is about to be moved or
	//  resized. This function is only used on Windows and Linux.
	NotifyMoveOrResizeStarted() // procedure
	// Invalidate
	//  Invalidate the view. The browser will call ICefRenderHandler.OnPaint
	//  asynchronously. This function is only used when window rendering is
	//  disabled.
	Invalidate(type_ cefTypes.TCefPaintElementType) // procedure
	// ExitFullscreen
	//  Requests the renderer to exit browser fullscreen. In most cases exiting
	//  window fullscreen should also exit browser fullscreen. With the Alloy
	//  runtime this function should be called in response to a user action such
	//  as clicking the green traffic light button on MacOS
	//  (ICefWindowDelegate.OnWindowFullscreenTransition callback) or pressing
	//  the "ESC" key (ICefKeyboardHandler.OnPreKeyEvent callback). With the
	//  Chrome runtime these standard exit actions are handled internally but
	//  new/additional user actions can use this function. Set |will_cause_resize|
	//  to true (1) if exiting browser fullscreen will cause a view resize.
	ExitFullscreen(willCauseResize bool) // procedure
	// ExecuteChromeCommand
	//  Execute a Chrome command. Values for |command_id| can be found in the
	//  cef_command_ids.h file. |disposition| provides information about the
	//  intended command target. Only used with the Chrome runtime.
	//  <see cref="uCEFConstants">See the IDC_* constants in uCEFConstants.pas for all the |command_id| values.</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/app/chrome_command_ids.h">The command_id values are also available in chrome/app/chrome_command_ids.h</see>
	ExecuteChromeCommand(commandId int32, disposition cefTypes.TCefWindowOpenDisposition) // procedure
	// SendExternalBeginFrame
	//  Issue a BeginFrame request to Chromium. Only valid when
	//  TCefWindowInfo.external_begin_frame_enabled is set to true (1).
	SendExternalBeginFrame() // procedure
	// SendKeyEvent
	//  Send a key event to the browser.
	SendKeyEvent(event TCefKeyEvent) // procedure
	// SendMouseClickEvent
	//  Send a mouse click event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view.
	SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32) // procedure
	// SendMouseMoveEvent
	//  Send a mouse move event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view.
	SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool) // procedure
	// SendMouseWheelEvent
	//  Send a mouse wheel event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view. The |deltaX| and |deltaY|
	//  values represent the movement delta in the X and Y directions
	//  respectively. In order to scroll inside select popups with window
	//  rendering disabled ICefRenderHandler.GetScreenPoint should be
	//  implemented properly.
	SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32) // procedure
	// SendTouchEvent
	//  Send a touch event to the browser for a windowless browser.
	SendTouchEvent(event TCefTouchEvent) // procedure
	// SendCaptureLostEvent
	//  Send a capture lost event to the browser.
	SendCaptureLostEvent() // procedure
	// SendProcessMessageWithProcessIdProcessMessageStringX2
	//  Send a message to the specified |targetProcess|. Ownership of the message
	//  contents will be transferred and the |ProcMessage| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases (for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
	SendProcessMessageWithProcessIdProcessMessageStringX2(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameName string, frameIdentifier string) // procedure
	// SendProcessMessageWithProcessIdProcessMessageFrame
	//  Send a message to the specified |targetProcess|. Ownership of the message
	//  contents will be transferred and the |ProcMessage| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases (for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
	SendProcessMessageWithProcessIdProcessMessageFrame(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frame ICefFrame) // procedure
	// SetFocus
	//  Set whether the browser is focused.
	SetFocus(focus bool) // procedure
	// SetAccessibilityState
	//  Set accessibility state for all frames. |accessibility_state| may be
	//  default, enabled or disabled. If |accessibility_state| is STATE_DEFAULT
	//  then accessibility will be disabled by default and the state may be
	//  further controlled with the "force-renderer-accessibility" and "disable-
	//  renderer-accessibility" command-line switches. If |accessibility_state| is
	//  STATE_ENABLED then accessibility will be enabled. If |accessibility_state|
	//  is STATE_DISABLED then accessibility will be completely disabled.
	//  For windowed browsers accessibility will be enabled in Complete mode
	//  (which corresponds to kAccessibilityModeComplete in Chromium). In this
	//  mode all platform accessibility objects will be created and managed by
	//  Chromium's internal implementation. The client needs only to detect the
	//  screen reader and call this function appropriately. For example, on macOS
	//  the client can handle the @"AXEnhancedUserStructure" accessibility
	//  attribute to detect VoiceOver state changes and on Windows the client can
	//  handle WM_GETOBJECT with OBJID_CLIENT to detect accessibility readers.
	//  For windowless browsers accessibility will be enabled in TreeOnly mode
	//  (which corresponds to kAccessibilityModeWebContentsOnly in Chromium). In
	//  this mode renderer accessibility is enabled, the full tree is computed,
	//  and events are passed to CefAccessibiltyHandler, but platform
	//  accessibility objects are not created. The client may implement platform
	//  accessibility objects using CefAccessibiltyHandler callbacks if desired.
	SetAccessibilityState(accessibilityState cefTypes.TCefState) // procedure
	// DragTargetDragEnter
	//  Call this function when the user drags the mouse into the web view (before
	//  calling DragTargetDragOver/DragTargetLeave/DragTargetDrop). |drag_data|
	//  should not contain file contents as this type of data is not allowed to be
	//  dragged into the web view. File contents can be removed using
	//  ICefDragData.ResetFileContents (for example, if |drag_data| comes from
	//  ICefRenderHandler.StartDragging). This function is only used when
	//  window rendering is disabled.
	DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) // procedure
	// DragTargetDragOver
	//  Call this function each time the mouse is moved across the web view during
	//  a drag operation (after calling DragTargetDragEnter and before calling
	//  DragTargetDragLeave/DragTargetDrop). This function is only used when
	//  window rendering is disabled.
	DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) // procedure
	// DragTargetDragLeave
	//  Call this function when the user drags the mouse out of the web view
	//  (after calling DragTargetDragEnter). This function is only used when
	//  window rendering is disabled.
	DragTargetDragLeave() // procedure
	// DragTargetDrop
	//  Call this function when the user completes the drag operation by dropping
	//  the object onto the web view (after calling DragTargetDragEnter). The
	//  object being dropped is |drag_data|, given as an argument to the previous
	//  DragTargetDragEnter call. This function is only used when window rendering
	//  is disabled.
	DragTargetDrop(event TCefMouseEvent) // procedure
	// DragSourceEndedAt
	//  Call this function when the drag operation started by a
	//  ICefRenderHandler.StartDragging call has ended either in a drop or by
	//  being cancelled. |x| and |y| are mouse coordinates relative to the upper-
	//  left corner of the view. If the web view is both the drag source and the
	//  drag target then all DragTarget* functions should be called before
	//  DragSource* mthods. This function is only used when window rendering is
	//  disabled.
	DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation) // procedure
	// DragSourceSystemDragEnded
	//  Call this function when the drag operation started by a
	//  ICefRenderHandler.StartDragging call has completed. This function may
	//  be called immediately without first calling DragSourceEndedAt to cancel a
	//  drag operation. If the web view is both the drag source and the drag
	//  target then all DragTarget* functions should be called before DragSource*
	//  mthods. This function is only used when window rendering is disabled.
	DragSourceSystemDragEnded() // procedure
	// IMESetComposition
	//  Begins a new composition or updates the existing composition. Blink has a
	//  special node (a composition node) that allows the input function to change
	//  text without affecting other DOM nodes. |text| is the optional text that
	//  will be inserted into the composition node. |underlines| is an optional
	//  set of ranges that will be underlined in the resulting text.
	//  |replacement_range| is an optional range of the existing text that will be
	//  replaced. |selection_range| is an optional range of the resulting text
	//  that will be selected after insertion or replacement. The
	//  |replacement_range| value is only used on OS X.
	//  This function may be called multiple times as the composition changes.
	//  When the client is done making changes the composition should either be
	//  canceled or completed. To cancel the composition call
	//  ImeCancelComposition. To complete the composition call either
	//  ImeCommitText or ImeFinishComposingText. Completion is usually signaled
	//  when:
	//  <code>
	//  1. The client receives a WM_IME_COMPOSITION message with a GCS_RESULTSTR
	//  flag (on Windows), or;
	//  2. The client receives a "commit" signal of GtkIMContext (on Linux), or;
	//  3. insertText of NSTextInput is called (on Mac).
	//  </code>
	//  This function is only used when window rendering is disabled.
	IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange) // procedure
	// IMECommitText
	//  Completes the existing composition by optionally inserting the specified
	//  |text| into the composition node. |replacement_range| is an optional range
	//  of the existing text that will be replaced. |relative_cursor_pos| is where
	//  the cursor will be positioned relative to the current cursor position. See
	//  comments on ImeSetComposition for usage. The |replacement_range| and
	//  |relative_cursor_pos| values are only used on OS X. This function is only
	//  used when window rendering is disabled.
	IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) // procedure
	// IMEFinishComposingText
	//  Completes the existing composition by applying the current composition
	//  node contents. If |keep_selection| is false (0) the current selection, if
	//  any, will be discarded. See comments on ImeSetComposition for usage. This
	//  function is only used when window rendering is disabled.
	IMEFinishComposingText(keepSelection bool) // procedure
	// IMECancelComposition
	//  Cancels the existing composition and discards the composition node
	//  contents without applying them. See comments on ImeSetComposition for
	//  usage. This function is only used when window rendering is disabled.
	IMECancelComposition() // procedure
	// ReplaceMisspelling
	//  If a misspelled word is currently selected in an editable node calling
	//  this function will replace it with the specified |word|.
	ReplaceMisspelling(word string) // procedure
	// AddWordToDictionary
	//  Add the specified |word| to the spelling dictionary.
	AddWordToDictionary(word string) // procedure
	// UpdateBrowserSize
	//  Used in Linux to resize the browser contents.
	UpdateBrowserSize(left int32, top int32, width int32, height int32) // procedure
	// UpdateXWindowVisibility
	//  Used in Linux to update the browser visibility.
	UpdateXWindowVisibility(visible bool) // procedure
	// NotifyCurrentSinks
	//  Trigger an asynchronous call to ICefMediaObserver.OnSinks on all
	//  registered observers.
	NotifyCurrentSinks() // procedure
	// NotifyCurrentRoutes
	//  Trigger an asynchronous call to ICefMediaObserver.OnRoutes on all
	//  registered observers.
	NotifyCurrentRoutes() // procedure
	// CreateRoute
	//  Create a new route between |source| and |sink|. Source and sink must be
	//  valid, compatible (as reported by ICefMediaSink.IsCompatibleWith), and
	//  a route between them must not already exist. |callback| will be executed
	//  on success or failure. If route creation succeeds it will also trigger an
	//  asynchronous call to ICefMediaObserver.OnRoutes on all registered
	//  observers.
	//  This procedure is asynchronous and the result, ICefMediaRoute and the error
	//  message will be available in the TChromium.OnMediaRouteCreateFinished event.
	CreateRoute(source ICefMediaSource, sink ICefMediaSink) // procedure
	// GetDeviceInfo
	//  Asynchronously retrieves device info.
	//  This procedure will trigger OnMediaSinkDeviceInfo with the device info.
	GetDeviceInfo(mediaSink ICefMediaSink) // procedure
	// SetWebsiteSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass nullptr for |value| to remove the
	//  default value for this content type.
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
	// DefaultUrl
	//  First URL loaded by the browser after its creation.
	DefaultUrl() string         // property DefaultUrl Getter
	SetDefaultUrl(value string) // property DefaultUrl Setter
	// Options
	//  Properties used to fill the TCefBrowserSettings record which is used during the browser creation.
	Options() IChromiumOptions         // property Options Getter
	SetOptions(value IChromiumOptions) // property Options Setter
	// FontOptions
	//  Font properties used to fill the TCefBrowserSettings record which is used during the browser creation.
	FontOptions() IChromiumFontOptions         // property FontOptions Getter
	SetFontOptions(value IChromiumFontOptions) // property FontOptions Setter
	// PDFPrintOptions
	//  Properties used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
	PDFPrintOptions() IPDFPrintOptions         // property PDFPrintOptions Getter
	SetPDFPrintOptions(value IPDFPrintOptions) // property PDFPrintOptions Setter
	// DefaultEncoding
	//  Default encoding for Web content. If empty "ISO-8859-1" will be used. Also
	//  configurable using the "default-encoding" command-line switch. It's used during the browser creation.
	DefaultEncoding() string         // property DefaultEncoding Getter
	SetDefaultEncoding(value string) // property DefaultEncoding Setter
	// BrowserId
	//  Globally unique identifier for the seleted browser.
	BrowserId() int32 // property BrowserId Getter
	// Browser
	//  Returns a ICefBrowser instance of the selected browser.
	Browser() ICefBrowser // property Browser Getter
	// BrowserById
	//  Returns a ICefBrowser instance of the browser with the specified id.
	BrowserById(id int32) ICefBrowser // property BrowserById Getter
	// BrowserCount
	//  Returns the number of browsers in the browser array when the multi-browser mode is enabled.
	BrowserCount() int32 // property BrowserCount Getter
	// BrowserIdByIndex
	//  Returns the identifier of the browser in the specified array position when the multi-browser mode is enabled.
	BrowserIdByIndex(I int32) int32 // property BrowserIdByIndex Getter
	// CefClient
	//  Returns the ICefClient instance used by the selected browser.
	CefClient() IEngClient // property CefClient Getter
	// ReqContextHandler
	//  Returns the ICefRequestContextHandler instance used in this browser.
	ReqContextHandler() IEngRequestContextHandler // property ReqContextHandler Getter
	// ResourceRequestHandler
	//  Returns the ICefResourceRequestHandler instance used in this browser.
	ResourceRequestHandler() IEngResourceRequestHandler // property ResourceRequestHandler Getter
	// CefWindowInfo
	//  Returns the TCefWindowInfo record used when the browser was created.
	CefWindowInfo() TCefWindowInfo // property CefWindowInfo Getter
	// VisibleNavigationEntry
	//  Returns the current visible navigation entry for this browser.
	//  This property can only be used on the CEF UI thread.
	VisibleNavigationEntry() ICefNavigationEntry // property VisibleNavigationEntry Getter
	// RuntimeStyle
	//  Returns the runtime style for this browser (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	//  This property can only be read on the CEF UI thread.
	RuntimeStyle() cefTypes.TCefRuntimeStyle         // property RuntimeStyle Getter
	SetRuntimeStyle(value cefTypes.TCefRuntimeStyle) // property RuntimeStyle Setter
	// RequestContext
	//  Returns a ICefRequestContext instance used by the selected browser.
	RequestContext() ICefRequestContext // property RequestContext Getter
	// MediaRouter
	//  Returns a ICefMediaRouter instance used by the selected browser.
	MediaRouter() ICefMediaRouter // property MediaRouter Getter
	// MediaObserver
	//  Returns a ICefMediaObserver instance used by the selected browser.
	MediaObserver() IEngMediaObserver // property MediaObserver Getter
	// MediaObserverReg
	//  Returns the ICefRegistration instance obtained when the default MediaObserver was added.
	MediaObserverReg() ICefRegistration // property MediaObserverReg Getter
	// DevToolsMsgObserver
	//  Returns a ICefDevToolsMessageObserver instance used by the selected browser.
	DevToolsMsgObserver() IEngDevToolsMessageObserver // property DevToolsMsgObserver Getter
	// DevToolsMsgObserverReg
	//  Returns the ICefRegistration instance obtained when the default DevToolsMessageObserver was added.
	DevToolsMsgObserverReg() ICefRegistration // property DevToolsMsgObserverReg Getter
	// ExtensionHandler
	//  Returns a ICefExtensionHandler instance used by the selected browser.
	ExtensionHandler() IEngExtensionHandler // property ExtensionHandler Getter
	// MultithreadApp
	//  Returns the value of GlobalCEFApp.MultiThreadedMessageLoop.
	MultithreadApp() bool // property MultithreadApp Getter
	// IsLoading
	//  Calls ICefBrowser.IsLoading and returns true if the browser is currently loading.
	IsLoading() bool // property IsLoading Getter
	// HasDocument
	//  Calls ICefBrowser.HasDocument and returns true if a document has been loaded in the browser.
	HasDocument() bool // property HasDocument Getter
	// HasView
	//  Calls ICefBrowserHost.HasView and returns true if this browser is wrapped in a ICefBrowserView.
	HasView() bool // property HasView Getter
	// HasDevTools
	//  Calls ICefBrowserHost.HasDevTools and returns true if this browser currently has an associated DevTools browser.
	HasDevTools() bool // property HasDevTools Getter
	// HasClientHandler
	//  Returns true if ICefClient has a valid value.
	HasClientHandler() bool // property HasClientHandler Getter
	// HasBrowser
	//  Returns true if this component has a valid selected browser.
	HasBrowser() bool // property HasBrowser Getter
	// CanGoBack
	//  Calls ICefBrowser.CanGoBack and returns true if the browser can navigate back.
	CanGoBack() bool // property CanGoBack Getter
	// CanGoForward
	//  Calls ICefBrowser.CanGoForward and returns true if the browser can navigate forward.
	CanGoForward() bool // property CanGoForward Getter
	// IsPopUp
	//  Calls ICefBrowser.IsPopUp and returns true if the window is a popup window.
	IsPopUp() bool // property IsPopUp Getter
	// WindowHandle
	//  Calls ICefBrowserHost.GetWindowHandle and returns the window handle for this browser.
	WindowHandle() cefTypes.TCefWindowHandle // property WindowHandle Getter
	// OpenerWindowHandle
	//  Calls ICefBrowserHost.GetOpenerWindowHandle and returns the window handle of the browser that opened this browser.
	OpenerWindowHandle() cefTypes.TCefWindowHandle // property OpenerWindowHandle Getter
	// BrowserHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	BrowserHandle() types.THandle // property BrowserHandle Getter
	// WidgetHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	WidgetHandle() types.THandle // property WidgetHandle Getter
	// RenderHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	RenderHandle() types.THandle // property RenderHandle Getter
	// FrameIsFocused
	//  Returns true if ICefBrowser.FocusedFrame has a valid value.
	FrameIsFocused() bool // property FrameIsFocused Getter
	// Initialized
	//  Returns true when the browser is fully initialized and it's not being closed.
	Initialized() bool // property Initialized Getter
	// RequestContextCache
	//  Returns the cache value in ICefRequestContext.CachePath.
	RequestContextCache() string // property RequestContextCache Getter
	// RequestContextIsGlobal
	//  Calls ICefRequestContext.IsGlobal to check if the request context is the global context or it's independent.
	RequestContextIsGlobal() bool // property RequestContextIsGlobal Getter
	// ChromeColorSchemeMode
	//  Returns the current Chrome color scheme mode (SYSTEM, LIGHT or DARK). Must
	//  be called on the browser process UI thread.
	ChromeColorSchemeMode() cefTypes.TCefColorVariant // property ChromeColorSchemeMode Getter
	// ChromeColorSchemeColor
	//  Returns the current Chrome color scheme color, or transparent (0) for the
	//  default color. Must be called on the browser process UI thread.
	ChromeColorSchemeColor() cefTypes.TCefColor // property ChromeColorSchemeColor Getter
	// ChromeColorSchemeVariant
	//  Returns the current Chrome color scheme variant. Must be called on the
	//  browser process UI thread.
	ChromeColorSchemeVariant() cefTypes.TCefColorVariant // property ChromeColorSchemeVariant Getter
	// DocumentURL
	//  Returns the URL of the main frame.
	DocumentURL() string // property DocumentURL Getter
	// ZoomLevel
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomLevel() float64         // property ZoomLevel Getter
	SetZoomLevel(value float64) // property ZoomLevel Setter
	// DefaultZoomLevel
	//  Get the default zoom level. This value will be 0.0 by default but can be
	//  configured with the Chrome runtime. This function can only be called on
	//  the CEF UI thread.
	DefaultZoomLevel() float64 // property DefaultZoomLevel Getter
	// CanIncZoom
	//  Returns true (1) if this browser can execute the zoom IN command.
	//  This function can only be called on the CEF UI thread.
	CanIncZoom() bool // property CanIncZoom Getter
	// CanDecZoom
	//  Returns true (1) if this browser can execute the zoom OUT command.
	//  This function can only be called on the CEF UI thread.
	CanDecZoom() bool // property CanDecZoom Getter
	// CanResetZoom
	//  Returns true (1) if this browser can execute the zoom RESET command.
	//  This function can only be called on the CEF UI thread.
	CanResetZoom() bool // property CanResetZoom Getter
	// ZoomPct
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomPct() float64         // property ZoomPct Getter
	SetZoomPct(value float64) // property ZoomPct Setter
	// ZoomStep
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomStep() byte         // property ZoomStep Getter
	SetZoomStep(value byte) // property ZoomStep Setter
	// WindowlessFrameRate
	//  Returns the maximum rate in frames per second (fps) that OnPaint will be called for a browser in OSR mode.
	WindowlessFrameRate() int32         // property WindowlessFrameRate Getter
	SetWindowlessFrameRate(value int32) // property WindowlessFrameRate Setter
	// CustomHeaderName
	//  Custom HTTP header name added to all requests.
	CustomHeaderName() string         // property CustomHeaderName Getter
	SetCustomHeaderName(value string) // property CustomHeaderName Setter
	// CustomHeaderValue
	//  Custom HTTP header value added to all requests.
	CustomHeaderValue() string         // property CustomHeaderValue Getter
	SetCustomHeaderValue(value string) // property CustomHeaderValue Setter
	// DoNotTrack
	//  Set to True if you want to send the DNT header.
	DoNotTrack() bool         // property DoNotTrack Getter
	SetDoNotTrack(value bool) // property DoNotTrack Setter
	// SendReferrer
	//  Set to True if you want to send the referer header.
	SendReferrer() bool         // property SendReferrer Getter
	SetSendReferrer(value bool) // property SendReferrer Setter
	// HyperlinkAuditing
	//  Enable hyperlink auditing.
	HyperlinkAuditing() bool         // property HyperlinkAuditing Getter
	SetHyperlinkAuditing(value bool) // property HyperlinkAuditing Setter
	// AllowOutdatedPlugins
	//  Allow using outdated plugins.
	AllowOutdatedPlugins() bool         // property AllowOutdatedPlugins Getter
	SetAllowOutdatedPlugins(value bool) // property AllowOutdatedPlugins Setter
	// AlwaysAuthorizePlugins
	//  Always authorize plugins.
	AlwaysAuthorizePlugins() bool         // property AlwaysAuthorizePlugins Getter
	SetAlwaysAuthorizePlugins(value bool) // property AlwaysAuthorizePlugins Setter
	// AlwaysOpenPDFExternally
	//  Always open PDF files externally.
	AlwaysOpenPDFExternally() bool         // property AlwaysOpenPDFExternally Getter
	SetAlwaysOpenPDFExternally(value bool) // property AlwaysOpenPDFExternally Setter
	// SpellChecking
	//  Set to True if you want to enable the spell checker.
	SpellChecking() bool         // property SpellChecking Getter
	SetSpellChecking(value bool) // property SpellChecking Setter
	// SpellCheckerDicts
	//  Comma delimited list of language codes used by the spell checker, for example "es-ES,en-US,fr-FR,de-DE,it-IT".
	SpellCheckerDicts() string         // property SpellCheckerDicts Getter
	SetSpellCheckerDicts(value string) // property SpellCheckerDicts Setter
	// HasValidMainFrame
	//  Returns true if the main frame exists and it's valid.
	HasValidMainFrame() bool // property HasValidMainFrame Getter
	// FrameCount
	//  Returns the number of frames that currently exist.
	FrameCount() cefTypes.NativeUInt // property FrameCount Getter
	// DragOperations
	//  Returns the TcefDragOperation value used during drag and drop.
	DragOperations() cefTypes.TCefDragOperations         // property DragOperations Getter
	SetDragOperations(value cefTypes.TCefDragOperations) // property DragOperations Setter
	// AudioMuted
	//  Returns true if the browser's audio is muted.
	AudioMuted() bool         // property AudioMuted Getter
	SetAudioMuted(value bool) // property AudioMuted Setter
	// Fullscreen
	//  Returns true (1) if the renderer is currently in browser fullscreen. This
	//  differs from window fullscreen in that browser fullscreen is entered using
	//  the JavaScript Fullscreen API and modifies CSS attributes such as the
	//  ::backdrop pseudo-element and :fullscreen pseudo-structure. This property
	//  can only be called on the UI thread.
	Fullscreen() bool // property Fullscreen Getter
	// IsRenderProcessUnresponsive
	//  Returns true (1) if the render process associated with this browser is
	//  currently unresponsive as indicated by a lack of input event processing
	//  for at least 15 seconds. To receive associated state change notifications
	//  and optionally handle an unresponsive render process implement
	//  ICefRequestHandler.OnRenderProcessUnresponsive.
	//  This property can only be read on the CEF UI thread.
	IsRenderProcessUnresponsive() bool // property IsRenderProcessUnresponsive Getter
	// SafeSearch
	//  Forces the Google safesearch in the browser preferences.
	SafeSearch() bool         // property SafeSearch Getter
	SetSafeSearch(value bool) // property SafeSearch Setter
	// YouTubeRestrict
	//  Forces the YouTube restrictions in the browser preferences.
	YouTubeRestrict() int32         // property YouTubeRestrict Getter
	SetYouTubeRestrict(value int32) // property YouTubeRestrict Setter
	// PrintingEnabled
	//  Enables printing in the browser preferences.
	PrintingEnabled() bool         // property PrintingEnabled Getter
	SetPrintingEnabled(value bool) // property PrintingEnabled Setter
	// AcceptLanguageList
	//  Set the accept language list in the browser preferences.
	AcceptLanguageList() string         // property AcceptLanguageList Getter
	SetAcceptLanguageList(value string) // property AcceptLanguageList Setter
	// AcceptCookies
	//  Sets the cookies policy value in the browser preferences.
	AcceptCookies() cefTypes.TCefCookiePref         // property AcceptCookies Getter
	SetAcceptCookies(value cefTypes.TCefCookiePref) // property AcceptCookies Setter
	// Block3rdPartyCookies
	//  Blocks third party cookies in the browser preferences.
	Block3rdPartyCookies() bool         // property Block3rdPartyCookies Getter
	SetBlock3rdPartyCookies(value bool) // property Block3rdPartyCookies Setter
	// MultiBrowserMode
	//  Enables the multi-browser mode that allows TChromiumCore to handle several browsers with one component. These browsers are usually the main browser, popup windows and new tabs.
	MultiBrowserMode() bool         // property MultiBrowserMode Getter
	SetMultiBrowserMode(value bool) // property MultiBrowserMode Setter
	// DefaultWindowInfoExStyle
	//  Default ExStyle value used to initialize the browser. A value of WS_EX_NOACTIVATE can be used as a workaround for some focus issues in CEF.
	DefaultWindowInfoExStyle() types.DWORD         // property DefaultWindowInfoExStyle Getter
	SetDefaultWindowInfoExStyle(value types.DWORD) // property DefaultWindowInfoExStyle Setter
	// Offline
	//  Uses the Network.emulateNetworkConditions DevTool method to set the browser in offline mode.
	Offline() bool         // property Offline Getter
	SetOffline(value bool) // property Offline Setter
	// QuicAllowed
	//  Enables the Quic protocol in the browser preferences.
	QuicAllowed() bool         // property QuicAllowed Getter
	SetQuicAllowed(value bool) // property QuicAllowed Setter
	// JavascriptEnabled
	//  Enables JavaScript in the browser preferences.
	JavascriptEnabled() bool         // property JavascriptEnabled Getter
	SetJavascriptEnabled(value bool) // property JavascriptEnabled Setter
	// LoadImagesAutomatically
	//  Enables automatic image loading in the browser preferences.
	LoadImagesAutomatically() bool         // property LoadImagesAutomatically Getter
	SetLoadImagesAutomatically(value bool) // property LoadImagesAutomatically Setter
	// BatterySaverModeState
	//  Battery saver mode state.
	BatterySaverModeState() cefTypes.TCefBatterySaverModeState         // property BatterySaverModeState Getter
	SetBatterySaverModeState(value cefTypes.TCefBatterySaverModeState) // property BatterySaverModeState Setter
	// HighEfficiencyModeState
	//  High efficiency mode state.
	HighEfficiencyModeState() cefTypes.TCefHighEfficiencyModeState         // property HighEfficiencyModeState Getter
	SetHighEfficiencyModeState(value cefTypes.TCefHighEfficiencyModeState) // property HighEfficiencyModeState Setter
	// CanFocus
	//  Indicates whether the browser can receive focus.
	CanFocus() bool // property CanFocus Getter
	// EnableFocusDelayMs
	//  Delay in milliseconds to enable browser focus.
	EnableFocusDelayMs() uint32         // property EnableFocusDelayMs Getter
	SetEnableFocusDelayMs(value uint32) // property EnableFocusDelayMs Setter
	// XDisplay
	//  Gets the Xdisplay pointer in Linux.
	XDisplay() uintptr // property XDisplay Getter
	// WebRTCIPHandlingPolicy
	//  WebRTC handling policy setting in the browser preferences.
	WebRTCIPHandlingPolicy() cefTypes.TCefWebRTCHandlingPolicy         // property WebRTCIPHandlingPolicy Getter
	SetWebRTCIPHandlingPolicy(value cefTypes.TCefWebRTCHandlingPolicy) // property WebRTCIPHandlingPolicy Setter
	// WebRTCMultipleRoutes
	//  WebRTC multiple routes setting in the browser preferences.
	WebRTCMultipleRoutes() cefTypes.TCefState         // property WebRTCMultipleRoutes Getter
	SetWebRTCMultipleRoutes(value cefTypes.TCefState) // property WebRTCMultipleRoutes Setter
	// WebRTCNonproxiedUDP
	//  WebRTC nonproxied UDP setting in the browser preferences.
	WebRTCNonproxiedUDP() cefTypes.TCefState         // property WebRTCNonproxiedUDP Getter
	SetWebRTCNonproxiedUDP(value cefTypes.TCefState) // property WebRTCNonproxiedUDP Setter
	// ProxyType
	//  Proxy type: CEF_PROXYTYPE_DIRECT, CEF_PROXYTYPE_AUTODETECT, CEF_PROXYTYPE_SYSTEM, CEF_PROXYTYPE_FIXED_SERVERS or CEF_PROXYTYPE_PAC_SCRIPT.
	ProxyType() int32         // property ProxyType Getter
	SetProxyType(value int32) // property ProxyType Setter
	// ProxyScheme
	//  Proxy scheme
	ProxyScheme() cefTypes.TCefProxyScheme         // property ProxyScheme Getter
	SetProxyScheme(value cefTypes.TCefProxyScheme) // property ProxyScheme Setter
	// ProxyServer
	//  Proxy server address
	ProxyServer() string         // property ProxyServer Getter
	SetProxyServer(value string) // property ProxyServer Setter
	// ProxyPort
	//  Proxy server port
	ProxyPort() int32         // property ProxyPort Getter
	SetProxyPort(value int32) // property ProxyPort Setter
	// ProxyUsername
	//  Proxy username
	ProxyUsername() string         // property ProxyUsername Getter
	SetProxyUsername(value string) // property ProxyUsername Setter
	// ProxyPassword
	//  Proxy password
	ProxyPassword() string         // property ProxyPassword Getter
	SetProxyPassword(value string) // property ProxyPassword Setter
	// ProxyScriptURL
	//  URL of the PAC script file.
	ProxyScriptURL() string         // property ProxyScriptURL Getter
	SetProxyScriptURL(value string) // property ProxyScriptURL Setter
	// ProxyByPassList
	//  This tells chromium to bypass any specified proxy for the given semi-colon-separated list of hosts.
	ProxyByPassList() string         // property ProxyByPassList Getter
	SetProxyByPassList(value string) // property ProxyByPassList Setter
	// MaxConnectionsPerProxy
	//  Sets the maximum connections per proxy value in the browser preferences (experimental).
	MaxConnectionsPerProxy() int32         // property MaxConnectionsPerProxy Getter
	SetMaxConnectionsPerProxy(value int32) // property MaxConnectionsPerProxy Setter
	// DownloadBubble
	//  Enable the file download bubble when using the Chrome runtime.
	DownloadBubble() cefTypes.TCefState         // property DownloadBubble Getter
	SetDownloadBubble(value cefTypes.TCefState) // property DownloadBubble Setter
	// HTTPSUpgrade
	//  Automatically upgrade to HTTPS connections.
	HTTPSUpgrade() cefTypes.TCefState         // property HTTPSUpgrade Getter
	SetHTTPSUpgrade(value cefTypes.TCefState) // property HTTPSUpgrade Setter
	// HSTSPolicyBypassList
	//  List of comma-delimited single-label hostnames that will skip the check to possibly upgrade from http to https.
	HSTSPolicyBypassList() string                                                      // property HSTSPolicyBypassList Getter
	SetHSTSPolicyBypassList(value string)                                              // property HSTSPolicyBypassList Setter
	SetOnTextResultAvailable(fn TOnTextResultAvailableEvent)                           // property event
	SetOnPdfPrintFinished(fn TOnPdfPrintFinishedEvent)                                 // property event
	SetOnPrefsAvailable(fn TOnPrefsAvailableEvent)                                     // property event
	SetOnPrefsUpdated(fn TNotifyEvent)                                                 // property event
	SetOnCookiesDeleted(fn TOnCookiesDeletedEvent)                                     // property event
	SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailableEvent)                        // property event
	SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailableEvent) // property event
	SetOnDownloadImageFinished(fn TOnDownloadImageFinishedEvent)                       // property event
	SetOnCookiesFlushed(fn TNotifyEvent)                                               // property event
	SetOnCertificateExceptionsCleared(fn TNotifyEvent)                                 // property event
	SetOnHttpAuthCredentialsCleared(fn TNotifyEvent)                                   // property event
	SetOnAllConnectionsClosed(fn TNotifyEvent)                                         // property event
	SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread)                          // property event
	SetOnCookiesVisited(fn TOnCookiesVisited)                                          // property event
	SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed)                          // property event
	SetOnCookieSet(fn TOnCookieSet)                                                    // property event
	SetOnZoomPctAvailable(fn TOnZoomPctAvailable)                                      // property event
	SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinishedEvent)                 // property event
	SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoEvent)                           // property event
	SetOnCanFocus(fn TNotifyEvent)                                                     // property event
	SetOnBrowserCompMsg(fn TOnCompMsgEvent)                                            // property event
	SetOnWidgetCompMsg(fn TOnCompMsgEvent)                                             // property event
	SetOnRenderCompMsg(fn TOnCompMsgEvent)                                             // property event
	SetOnProcessMessageReceived(fn TOnProcessMessageReceived)                          // property event
	SetOnLoadStart(fn TOnLoadStart)                                                    // property event
	SetOnLoadEnd(fn TOnLoadEnd)                                                        // property event
	SetOnLoadError(fn TOnLoadError)                                                    // property event
	SetOnLoadingStateChange(fn TOnLoadingStateChange)                                  // property event
	SetOnTakeFocus(fn TOnTakeFocus)                                                    // property event
	SetOnSetFocus(fn TOnSetFocus)                                                      // property event
	SetOnGotFocus(fn TOnGotFocus)                                                      // property event
	SetOnBeforeContextMenu(fn TOnBeforeContextMenu)                                    // property event
	SetOnRunContextMenu(fn TOnRunContextMenu)                                          // property event
	SetOnContextMenuCommand(fn TOnContextMenuCommand)                                  // property event
	SetOnContextMenuDismissed(fn TOnContextMenuDismissed)                              // property event
	SetOnRunQuickMenu(fn TOnRunQuickMenuEvent)                                         // property event
	SetOnQuickMenuCommand(fn TOnQuickMenuCommandEvent)                                 // property event
	SetOnQuickMenuDismissed(fn TOnQuickMenuDismissedEvent)                             // property event
	SetOnPreKeyEvent(fn TOnPreKeyEvent)                                                // property event
	SetOnKeyEvent(fn TOnKeyEvent)                                                      // property event
	SetOnAddressChange(fn TOnAddressChange)                                            // property event
	SetOnTitleChange(fn TOnTitleChange)                                                // property event
	SetOnFavIconUrlChange(fn TOnFavIconUrlChange)                                      // property event
	SetOnFullScreenModeChange(fn TOnFullScreenModeChange)                              // property event
	SetOnTooltip(fn TOnTooltip)                                                        // property event
	SetOnStatusMessage(fn TOnStatusMessage)                                            // property event
	SetOnConsoleMessage(fn TOnConsoleMessage)                                          // property event
	SetOnAutoResize(fn TOnAutoResize)                                                  // property event
	SetOnLoadingProgressChange(fn TOnLoadingProgressChange)                            // property event
	SetOnCursorChange(fn TOnCursorChange)                                              // property event
	SetOnMediaAccessChange(fn TOnMediaAccessChange)                                    // property event
	SetOnCanDownload(fn TOnCanDownloadEvent)                                           // property event
	SetOnBeforeDownload(fn TOnBeforeDownload)                                          // property event
	SetOnDownloadUpdated(fn TOnDownloadUpdated)                                        // property event
	SetOnJsdialog(fn TOnJsdialog)                                                      // property event
	SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog)                                  // property event
	SetOnResetDialogState(fn TOnResetDialogState)                                      // property event
	SetOnDialogClosed(fn TOnDialogClosed)                                              // property event
	SetOnBeforePopup(fn TOnBeforePopup)                                                // property event
	SetOnBeforeDevToolsPopup(fn TOnBeforeDevToolsPopup)                                // property event
	SetOnAfterCreated(fn TOnAfterCreated)                                              // property event
	SetOnBeforeClose(fn TOnBeforeClose)                                                // property event
	SetOnClose(fn TOnClose)                                                            // property event
	SetOnBeforeBrowse(fn TOnBeforeBrowse)                                              // property event
	SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab)                                          // property event
	SetOnGetAuthCredentials(fn TOnGetAuthCredentials)                                  // property event
	SetOnCertificateError(fn TOnCertificateError)                                      // property event
	SetOnSelectClientCertificate(fn TOnSelectClientCertificate)                        // property event
	SetOnRenderViewReady(fn TOnRenderViewReady)                                        // property event
	SetOnRenderProcessUnresponsive(fn TOnRenderProcessUnresponsive)                    // property event
	SetOnRenderProcessResponsive(fn TOnRenderProcessResponsive)                        // property event
	SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated)                        // property event
	SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler)            // property event
	SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame)              // property event
	SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad)                                  // property event
	SetOnGetResourceHandler(fn TOnGetResourceHandler)                                  // property event
	SetOnResourceRedirect(fn TOnResourceRedirect)                                      // property event
	SetOnResourceResponse(fn TOnResourceResponse)                                      // property event
	SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter)                    // property event
	SetOnResourceLoadComplete(fn TOnResourceLoadComplete)                              // property event
	SetOnProtocolExecution(fn TOnProtocolExecution)                                    // property event
	SetOnCanSendCookie(fn TOnCanSendCookie)                                            // property event
	SetOnCanSaveCookie(fn TOnCanSaveCookie)                                            // property event
	SetOnFileDialog(fn TOnFileDialog)                                                  // property event
	SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler)                        // property event
	SetOnGetRootScreenRect(fn TOnGetRootScreenRect)                                    // property event
	SetOnGetViewRect(fn TOnGetViewRect)                                                // property event
	SetOnGetScreenPoint(fn TOnGetScreenPoint)                                          // property event
	SetOnGetScreenInfo(fn TOnGetScreenInfo)                                            // property event
	SetOnPopupShow(fn TOnPopupShow)                                                    // property event
	SetOnPopupSize(fn TOnPopupSize)                                                    // property event
	SetOnPaint(fn TOnPaint)                                                            // property event
	SetOnAcceleratedPaint(fn TOnAcceleratedPaint)                                      // property event
	SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize)                                  // property event
	SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged)                        // property event
	SetOnStartDragging(fn TOnStartDragging)                                            // property event
	SetOnUpdateDragCursor(fn TOnUpdateDragCursor)                                      // property event
	SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged)                                // property event
	SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged)                  // property event
	SetOnTextSelectionChanged(fn TOnTextSelectionChanged)                              // property event
	SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested)                      // property event
	SetOnDragEnter(fn TOnDragEnter)                                                    // property event
	SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged)                        // property event
	SetOnFindResult(fn TOnFindResult)                                                  // property event
	SetOnRequestContextInitialized(fn TOnRequestContextInitialized)                    // property event
	SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler)         // property event
	SetOnSinks(fn TOnSinksEvent)                                                       // property event
	SetOnRoutes(fn TOnRoutesEvent)                                                     // property event
	SetOnRouteStateChanged(fn TOnRouteStateChangedEvent)                               // property event
	SetOnRouteMessageReceived(fn TOnRouteMessageReceivedEvent)                         // property event
	SetOnGetAudioParameters(fn TOnGetAudioParametersEvent)                             // property event
	SetOnAudioStreamStarted(fn TOnAudioStreamStartedEvent)                             // property event
	SetOnAudioStreamPacket(fn TOnAudioStreamPacketEvent)                               // property event
	SetOnAudioStreamStopped(fn TOnAudioStreamStoppedEvent)                             // property event
	SetOnAudioStreamError(fn TOnAudioStreamErrorEvent)                                 // property event
	SetOnDevToolsMessage(fn TOnDevToolsMessageEvent)                                   // property event
	SetOnDevToolsRawMessage(fn TOnDevToolsRawMessageEvent)                             // property event
	SetOnDevToolsMethodResult(fn TOnDevToolsMethodResultEvent)                         // property event
	SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResultEvent)                   // property event
	SetOnDevToolsEvent(fn TOnDevToolsEventEvent)                                       // property event
	SetOnDevToolsRawEvent(fn TOnDevToolsEventRawEvent)                                 // property event
	SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttachedEvent)                       // property event
	SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetachedEvent)                       // property event
	SetOnExtensionLoadFailed(fn TOnExtensionLoadFailedEvent)                           // property event
	SetOnExtensionLoaded(fn TOnExtensionLoadedEvent)                                   // property event
	SetOnExtensionUnloaded(fn TOnExtensionUnloadedEvent)                               // property event
	SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowserEvent)          // property event
	SetOnExtensionBeforeBrowser(fn TOnBeforeBrowserEvent)                              // property event
	SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowserEvent)                        // property event
	SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowserEvent)                        // property event
	SetOnExtensionGetExtensionResource(fn TOnGetExtensionResourceEvent)                // property event
	SetOnPrintStart(fn TOnPrintStartEvent)                                             // property event
	SetOnPrintSettings(fn TOnPrintSettingsEvent)                                       // property event
	SetOnPrintDialog(fn TOnPrintDialogEvent)                                           // property event
	SetOnPrintJob(fn TOnPrintJobEvent)                                                 // property event
	SetOnPrintReset(fn TOnPrintResetEvent)                                             // property event
	SetOnGetPDFPaperSize(fn TOnGetPDFPaperSizeEvent)                                   // property event
	SetOnFrameCreated(fn TOnFrameCreated)                                              // property event
	SetOnFrameAttached(fn TOnFrameAttached)                                            // property event
	SetOnFrameDetached(fn TOnFrameDetached)                                            // property event
	SetOnMainFrameChanged(fn TOnMainFrameChanged)                                      // property event
	SetOnChromeCommand(fn TOnChromeCommandEvent)                                       // property event
	SetOnIsChromeAppMenuItemVisible(fn TOnIsChromeAppMenuItemVisibleEvent)             // property event
	SetOnIsChromeAppMenuItemEnabled(fn TOnIsChromeAppMenuItemEnabledEvent)             // property event
	SetOnIsChromePageActionIconVisible(fn TOnIsChromePageActionIconVisibleEvent)       // property event
	SetOnIsChromeToolbarButtonVisible(fn TOnIsChromeToolbarButtonVisibleEvent)         // property event
	SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermissionEvent)         // property event
	SetOnShowPermissionPrompt(fn TOnShowPermissionPromptEvent)                         // property event
	SetOnDismissPermissionPrompt(fn TOnDismissPermissionPromptEvent)                   // property event
	AsIntfChromiumEvents() uintptr
}

type TChromiumCore struct {
	lcl.TComponent
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

func (m *TChromiumCore) CreateBrowserWithWindowHandleRectStringRequestContextDictionaryValueBool(parentHandle cefTypes.TCefWindowHandle, parentRect types.TRect, windowName string, context ICefRequestContext, extraInfo ICefDictionaryValue, forceAsPopup bool) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(8, m.Instance(), uintptr(parentHandle), uintptr(base.UnsafePointer(&parentRect)), api.PasStr(windowName), base.GetObjectUintptr(context), base.GetObjectUintptr(extraInfo), api.PasBool(forceAsPopup))
	return api.GoBool(r)
}

func (m *TChromiumCore) CreateBrowserWithStringBrowserViewComponentRequestContextDictionaryValue(uRL string, browserViewComp ICEFBrowserViewComponent, context ICefRequestContext, extraInfo ICefDictionaryValue) bool {
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

func (m *TChromiumCore) GetFrameIdentifiers(frameIdentifiers *lcl.IStrings) bool {
	if !m.IsValid() {
		return false
	}
	frameIdentifiersPtr := base.GetObjectUintptr(*frameIdentifiers)
	r := chromiumCoreAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&frameIdentifiersPtr)))
	*frameIdentifiers = lcl.AsStrings(frameIdentifiersPtr)
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

func (m *TChromiumCore) SetCookie(url string, name string, value string, domain string, path string, secure bool, httponly bool, hasExpires bool, creation types.TDateTime, lastAccess types.TDateTime, expires types.TDateTime, sameSite cefTypes.TCefCookieSameSite, priority cefTypes.TCefCookiePriority, setImmediately bool, iD int32) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(20, m.Instance(), api.PasStr(url), api.PasStr(name), api.PasStr(value), api.PasStr(domain), api.PasStr(path), api.PasBool(secure), api.PasBool(httponly), api.PasBool(hasExpires), uintptr(base.UnsafePointer(&creation)), uintptr(base.UnsafePointer(&lastAccess)), uintptr(base.UnsafePointer(&expires)), uintptr(sameSite), uintptr(priority), api.PasBool(setImmediately), uintptr(iD))
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

func (m *TChromiumCore) CanExecuteChromeCommand(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(25, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TChromiumCore) CreateUrlRequestWithRequestUrlrequestClientStringX2(request ICefRequest, client IEngUrlrequestClient, frameName string, frameIdentifier string) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), api.PasStr(frameName), api.PasStr(frameIdentifier), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefUrlRequestRef(resultPtr)
	return
}

func (m *TChromiumCore) CreateUrlRequestWithRequestUrlrequestClientFrame(request ICefRequest, client IEngUrlrequestClient, frame ICefFrame) (result ICefUrlRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(request), base.GetObjectUintptr(client), base.GetObjectUintptr(frame), uintptr(base.UnsafePointer(&resultPtr)))
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

func (m *TChromiumCore) GetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(30, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TChromiumCore) GetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes) cefTypes.TCefContentSettingValues {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(31, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType))
	return cefTypes.TCefContentSettingValues(r)
}

func (m *TChromiumCore) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(32, m.Instance(), api.PasBool(forceClose))
}

func (m *TChromiumCore) CloseAllBrowsers() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(33, m.Instance())
}

func (m *TChromiumCore) InitializeDragAndDrop(dropTargetWnd types.HWND) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(34, m.Instance(), uintptr(dropTargetWnd))
}

func (m *TChromiumCore) ShutdownDragAndDrop() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(35, m.Instance())
}

func (m *TChromiumCore) LoadURLWithStringX3(uRL string, frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(36, m.Instance(), api.PasStr(uRL), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) LoadURLWithStringFrame(uRL string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(37, m.Instance(), api.PasStr(uRL), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadStringWithStringX3(hTML string, frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(38, m.Instance(), api.PasStr(hTML), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) LoadStringWithStringFrame(hTML string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(39, m.Instance(), api.PasStr(hTML), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadResourceWithCustomMemoryStreamStringX4(stream lcl.ICustomMemoryStream, mimeType string, charset string, frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(40, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType), api.PasStr(charset), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) LoadResourceWithCustomMemoryStreamStringX2Frame(stream lcl.ICustomMemoryStream, mimeType string, charset string, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(41, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType), api.PasStr(charset), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) LoadRequest(request ICefRequest) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(42, m.Instance(), base.GetObjectUintptr(request))
}

func (m *TChromiumCore) GoBack() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(43, m.Instance())
}

func (m *TChromiumCore) GoForward() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(44, m.Instance())
}

func (m *TChromiumCore) Reload() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(45, m.Instance())
}

func (m *TChromiumCore) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(46, m.Instance())
}

func (m *TChromiumCore) StopLoad() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(47, m.Instance())
}

func (m *TChromiumCore) StartDownload(uRL string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(48, m.Instance(), api.PasStr(uRL))
}

func (m *TChromiumCore) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(49, m.Instance(), api.PasStr(imageUrl), api.PasBool(isFavicon), uintptr(maxImageSize), api.PasBool(bypassCache))
}

func (m *TChromiumCore) SimulateMouseWheel(deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(50, m.Instance(), uintptr(deltaX), uintptr(deltaY))
}

func (m *TChromiumCore) SimulateKeyEvent(type_ cefTypes.TSimulatedCefKeyEventType, modifiers int32, timestamp float32, text string, unmodifiedtext string, keyIdentifier string, code string, key string, windowsVirtualKeyCode int32, nativeVirtualKeyCode int32, autoRepeat bool, isKeypad bool, isSystemKey bool, location cefTypes.TCefKeyLocation, commands cefTypes.TCefEditingCommand) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(51, m.Instance(), uintptr(type_), uintptr(modifiers), uintptr(base.UnsafePointer(&timestamp)), api.PasStr(text), api.PasStr(unmodifiedtext), api.PasStr(keyIdentifier), api.PasStr(code), api.PasStr(key), uintptr(windowsVirtualKeyCode), uintptr(nativeVirtualKeyCode), api.PasBool(autoRepeat), api.PasBool(isKeypad), api.PasBool(isSystemKey), uintptr(location), uintptr(commands))
}

func (m *TChromiumCore) SimulateMouseEvent(type_ cefTypes.TCefSimulatedMouseEventType, X float32, Y float32, modifiers int32, timestamp float32, button cefTypes.TCefSimulatedMouseButton, buttons int32, clickCount int32, force float32, tangentialPressure float32, tiltX float32, tiltY float32, twist int32, deltaX float32, deltaY float32, pointerType cefTypes.TCefSimulatedPointerType) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(52, m.Instance(), uintptr(type_), uintptr(base.UnsafePointer(&X)), uintptr(base.UnsafePointer(&Y)), uintptr(modifiers), uintptr(base.UnsafePointer(&timestamp)), uintptr(button), uintptr(buttons), uintptr(clickCount), uintptr(base.UnsafePointer(&force)), uintptr(base.UnsafePointer(&tangentialPressure)), uintptr(base.UnsafePointer(&tiltX)), uintptr(base.UnsafePointer(&tiltY)), uintptr(twist), uintptr(base.UnsafePointer(&deltaX)), uintptr(base.UnsafePointer(&deltaY)), uintptr(pointerType))
}

func (m *TChromiumCore) SimulateEditingCommand(command cefTypes.TCefEditingCommand) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(53, m.Instance(), uintptr(command))
}

func (m *TChromiumCore) RetrieveHTMLWithStringX2(frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(54, m.Instance(), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) RetrieveHTMLWithFrame(frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(55, m.Instance(), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) RetrieveTextWithStringX2(frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(56, m.Instance(), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) RetrieveTextWithFrame(frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(57, m.Instance(), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) GetNavigationEntries(currentOnly bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(58, m.Instance(), api.PasBool(currentOnly))
}

func (m *TChromiumCore) ExecuteJavaScriptWithStringX4Int(code string, scriptURL string, frameName string, frameIdentifier string, startLine int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(59, m.Instance(), api.PasStr(code), api.PasStr(scriptURL), api.PasStr(frameName), api.PasStr(frameIdentifier), uintptr(startLine))
}

func (m *TChromiumCore) ExecuteJavaScriptWithStringX2FrameInt(code string, scriptURL string, frame ICefFrame, startLine int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(60, m.Instance(), api.PasStr(code), api.PasStr(scriptURL), base.GetObjectUintptr(frame), uintptr(startLine))
}

func (m *TChromiumCore) UpdatePreferences() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(61, m.Instance())
}

func (m *TChromiumCore) SavePreferences(fileName string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(62, m.Instance(), api.PasStr(fileName))
}

func (m *TChromiumCore) ResolveHost(uRL string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(63, m.Instance(), api.PasStr(uRL))
}

func (m *TChromiumCore) SetUserAgentOverride(userAgent string, acceptLanguage string, platform string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(64, m.Instance(), api.PasStr(userAgent), api.PasStr(acceptLanguage), api.PasStr(platform))
}

func (m *TChromiumCore) ClearDataForOrigin(origin string, storageTypes cefTypes.TCefClearDataStorageTypes) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(65, m.Instance(), api.PasStr(origin), uintptr(storageTypes))
}

func (m *TChromiumCore) ClearCache() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(66, m.Instance())
}

func (m *TChromiumCore) ToggleAudioMuted() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(67, m.Instance())
}

func (m *TChromiumCore) ShowDevTools(inspectElementAt types.TPoint, windowInfo TCefWindowInfo) {
	if !m.IsValid() {
		return
	}
	windowInfoPtr := windowInfo.ToPas()
	chromiumCoreAPI().SysCallN(68, m.Instance(), uintptr(base.UnsafePointer(&inspectElementAt)), uintptr(base.UnsafePointer(windowInfoPtr)))
}

func (m *TChromiumCore) CloseDevTools() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(69, m.Instance())
}

func (m *TChromiumCore) CloseDevToolsWithWindowHandle(devToolsWnd cefTypes.TCefWindowHandle) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(70, m.Instance(), uintptr(devToolsWnd))
}

func (m *TChromiumCore) Find(searchText string, forward bool, matchCase bool, findNext bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(71, m.Instance(), api.PasStr(searchText), api.PasBool(forward), api.PasBool(matchCase), api.PasBool(findNext))
}

func (m *TChromiumCore) StopFinding(clearSelection bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(72, m.Instance(), api.PasBool(clearSelection))
}

func (m *TChromiumCore) Print() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(73, m.Instance())
}

func (m *TChromiumCore) PrintToPDF(filePath string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(74, m.Instance(), api.PasStr(filePath))
}

func (m *TChromiumCore) ClipboardCopy() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(75, m.Instance())
}

func (m *TChromiumCore) ClipboardPaste() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(76, m.Instance())
}

func (m *TChromiumCore) ClipboardCut() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(77, m.Instance())
}

func (m *TChromiumCore) ClipboardUndo() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(78, m.Instance())
}

func (m *TChromiumCore) ClipboardRedo() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(79, m.Instance())
}

func (m *TChromiumCore) ClipboardDel() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(80, m.Instance())
}

func (m *TChromiumCore) SelectAll() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(81, m.Instance())
}

func (m *TChromiumCore) IncZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(82, m.Instance())
}

func (m *TChromiumCore) DecZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(83, m.Instance())
}

func (m *TChromiumCore) IncZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(84, m.Instance())
}

func (m *TChromiumCore) DecZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(85, m.Instance())
}

func (m *TChromiumCore) ResetZoomStep() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(86, m.Instance())
}

func (m *TChromiumCore) ResetZoomLevel() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(87, m.Instance())
}

func (m *TChromiumCore) ResetZoomPct() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(88, m.Instance())
}

func (m *TChromiumCore) ReadZoom() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(89, m.Instance())
}

func (m *TChromiumCore) IncZoomCommand() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(90, m.Instance())
}

func (m *TChromiumCore) DecZoomCommand() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(91, m.Instance())
}

func (m *TChromiumCore) ResetZoomCommand() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(92, m.Instance())
}

func (m *TChromiumCore) WasResized() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(93, m.Instance())
}

func (m *TChromiumCore) WasHidden(hidden bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(94, m.Instance(), api.PasBool(hidden))
}

func (m *TChromiumCore) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(95, m.Instance())
}

func (m *TChromiumCore) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(96, m.Instance())
}

func (m *TChromiumCore) Invalidate(type_ cefTypes.TCefPaintElementType) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(97, m.Instance(), uintptr(type_))
}

func (m *TChromiumCore) ExitFullscreen(willCauseResize bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(98, m.Instance(), api.PasBool(willCauseResize))
}

func (m *TChromiumCore) ExecuteChromeCommand(commandId int32, disposition cefTypes.TCefWindowOpenDisposition) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(99, m.Instance(), uintptr(commandId), uintptr(disposition))
}

func (m *TChromiumCore) SendExternalBeginFrame() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(100, m.Instance())
}

func (m *TChromiumCore) SendKeyEvent(event TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(101, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(102, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(type_), api.PasBool(mouseUp), uintptr(clickCount))
}

func (m *TChromiumCore) SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(103, m.Instance(), uintptr(base.UnsafePointer(&event)), api.PasBool(mouseLeave))
}

func (m *TChromiumCore) SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(104, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TChromiumCore) SendTouchEvent(event TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(105, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(106, m.Instance())
}

func (m *TChromiumCore) SendProcessMessageWithProcessIdProcessMessageStringX2(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frameName string, frameIdentifier string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(107, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(procMessage), api.PasStr(frameName), api.PasStr(frameIdentifier))
}

func (m *TChromiumCore) SendProcessMessageWithProcessIdProcessMessageFrame(targetProcess cefTypes.TCefProcessId, procMessage ICefProcessMessage, frame ICefFrame) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(108, m.Instance(), uintptr(targetProcess), base.GetObjectUintptr(procMessage), base.GetObjectUintptr(frame))
}

func (m *TChromiumCore) SetFocus(focus bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(109, m.Instance(), api.PasBool(focus))
}

func (m *TChromiumCore) SetAccessibilityState(accessibilityState cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(110, m.Instance(), uintptr(accessibilityState))
}

func (m *TChromiumCore) DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(111, m.Instance(), base.GetObjectUintptr(dragData), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(112, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragLeave() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(113, m.Instance())
}

func (m *TChromiumCore) DragTargetDrop(event TCefMouseEvent) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(114, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TChromiumCore) DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(115, m.Instance(), uintptr(X), uintptr(Y), uintptr(op))
}

func (m *TChromiumCore) DragSourceSystemDragEnded() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(116, m.Instance())
}

func (m *TChromiumCore) IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(117, m.Instance(), api.PasStr(text), underlines.Instance(), uintptr(int32(underlines.Count())), uintptr(base.UnsafePointer(&replacementRange)), uintptr(base.UnsafePointer(&selectionRange)))
}

func (m *TChromiumCore) IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(118, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&replacementRange)), uintptr(relativeCursorPos))
}

func (m *TChromiumCore) IMEFinishComposingText(keepSelection bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(119, m.Instance(), api.PasBool(keepSelection))
}

func (m *TChromiumCore) IMECancelComposition() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(120, m.Instance())
}

func (m *TChromiumCore) ReplaceMisspelling(word string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(121, m.Instance(), api.PasStr(word))
}

func (m *TChromiumCore) AddWordToDictionary(word string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(122, m.Instance(), api.PasStr(word))
}

func (m *TChromiumCore) UpdateBrowserSize(left int32, top int32, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(123, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

func (m *TChromiumCore) UpdateXWindowVisibility(visible bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(124, m.Instance(), api.PasBool(visible))
}

func (m *TChromiumCore) NotifyCurrentSinks() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(125, m.Instance())
}

func (m *TChromiumCore) NotifyCurrentRoutes() {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(126, m.Instance())
}

func (m *TChromiumCore) CreateRoute(source ICefMediaSource, sink ICefMediaSink) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(127, m.Instance(), base.GetObjectUintptr(source), base.GetObjectUintptr(sink))
}

func (m *TChromiumCore) GetDeviceInfo(mediaSink ICefMediaSink) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(128, m.Instance(), base.GetObjectUintptr(mediaSink))
}

func (m *TChromiumCore) SetWebsiteSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value ICefValue) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(129, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) SetContentSetting(requestingUrl string, topLevelUrl string, contentType cefTypes.TCefContentSettingTypes, value cefTypes.TCefContentSettingValues) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(130, m.Instance(), api.PasStr(requestingUrl), api.PasStr(topLevelUrl), uintptr(contentType), uintptr(value))
}

func (m *TChromiumCore) SetChromeColorScheme(variant cefTypes.TCefColorVariant, userColor cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(131, m.Instance(), uintptr(variant), uintptr(userColor))
}

func (m *TChromiumCore) DefaultUrl() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(132, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetDefaultUrl(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(132, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) Options() IChromiumOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(133, 0, m.Instance())
	return AsChromiumOptions(r)
}

func (m *TChromiumCore) SetOptions(value IChromiumOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(133, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) FontOptions() IChromiumFontOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(134, 0, m.Instance())
	return AsChromiumFontOptions(r)
}

func (m *TChromiumCore) SetFontOptions(value IChromiumFontOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(134, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) PDFPrintOptions() IPDFPrintOptions {
	if !m.IsValid() {
		return nil
	}
	r := chromiumCoreAPI().SysCallN(135, 0, m.Instance())
	return AsPDFPrintOptions(r)
}

func (m *TChromiumCore) SetPDFPrintOptions(value IPDFPrintOptions) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(135, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChromiumCore) DefaultEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(136, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(136, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) BrowserId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(137, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) Browser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(138, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TChromiumCore) BrowserById(id int32) (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(139, m.Instance(), uintptr(id), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TChromiumCore) BrowserCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(140, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) BrowserIdByIndex(I int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(141, m.Instance(), uintptr(I))
	return int32(r)
}

func (m *TChromiumCore) CefClient() (result IEngClient) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(142, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngClient(resultPtr)
	return
}

func (m *TChromiumCore) ReqContextHandler() (result IEngRequestContextHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(143, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngRequestContextHandler(resultPtr)
	return
}

func (m *TChromiumCore) ResourceRequestHandler() (result IEngResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(144, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngResourceRequestHandler(resultPtr)
	return
}

func (m *TChromiumCore) CefWindowInfo() (result TCefWindowInfo) {
	if !m.IsValid() {
		return
	}
	resultPtr := result.ToPas()
	chromiumCoreAPI().SysCallN(145, m.Instance(), uintptr(base.UnsafePointer(resultPtr)))
	result = resultPtr.ToGo()
	return
}

func (m *TChromiumCore) VisibleNavigationEntry() (result ICefNavigationEntry) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(146, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefNavigationEntryRef(resultPtr)
	return
}

func (m *TChromiumCore) RuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(147, 0, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TChromiumCore) SetRuntimeStyle(value cefTypes.TCefRuntimeStyle) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(147, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) RequestContext() (result ICefRequestContext) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(148, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (m *TChromiumCore) MediaRouter() (result ICefMediaRouter) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(149, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefMediaRouterRef(resultPtr)
	return
}

func (m *TChromiumCore) MediaObserver() (result IEngMediaObserver) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(150, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngMediaObserver(resultPtr)
	return
}

func (m *TChromiumCore) MediaObserverReg() (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(151, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) DevToolsMsgObserver() (result IEngDevToolsMessageObserver) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(152, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngDevToolsMessageObserver(resultPtr)
	return
}

func (m *TChromiumCore) DevToolsMsgObserverReg() (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(153, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TChromiumCore) ExtensionHandler() (result IEngExtensionHandler) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	chromiumCoreAPI().SysCallN(154, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngExtensionHandler(resultPtr)
	return
}

func (m *TChromiumCore) MultithreadApp() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(155, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(156, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasDocument() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(157, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasView() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(158, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(159, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasClientHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(160, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) HasBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(161, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(162, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(163, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) IsPopUp() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(164, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) WindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(165, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TChromiumCore) OpenerWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(166, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TChromiumCore) BrowserHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(167, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) WidgetHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(168, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) RenderHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(169, m.Instance())
	return types.THandle(r)
}

func (m *TChromiumCore) FrameIsFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(170, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(171, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) RequestContextCache() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(172, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) RequestContextIsGlobal() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(173, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) ChromeColorSchemeMode() cefTypes.TCefColorVariant {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(174, m.Instance())
	return cefTypes.TCefColorVariant(r)
}

func (m *TChromiumCore) ChromeColorSchemeColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(175, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TChromiumCore) ChromeColorSchemeVariant() cefTypes.TCefColorVariant {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(176, m.Instance())
	return cefTypes.TCefColorVariant(r)
}

func (m *TChromiumCore) DocumentURL() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(177, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) ZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(178, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TChromiumCore) SetZoomLevel(value float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(178, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TChromiumCore) DefaultZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(179, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TChromiumCore) CanIncZoom() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(180, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanDecZoom() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(181, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) CanResetZoom() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(182, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) ZoomPct() (result float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(183, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TChromiumCore) SetZoomPct(value float64) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(183, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TChromiumCore) ZoomStep() byte {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(184, 0, m.Instance())
	return byte(r)
}

func (m *TChromiumCore) SetZoomStep(value byte) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(184, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(185, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetWindowlessFrameRate(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(185, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) CustomHeaderName() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(186, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetCustomHeaderName(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(186, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) CustomHeaderValue() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(187, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetCustomHeaderValue(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(187, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) DoNotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(188, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetDoNotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(188, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SendReferrer() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(189, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSendReferrer(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(189, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) HyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(190, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetHyperlinkAuditing(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(190, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AllowOutdatedPlugins() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(191, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAllowOutdatedPlugins(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(191, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AlwaysAuthorizePlugins() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(192, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAlwaysAuthorizePlugins(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(192, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AlwaysOpenPDFExternally() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(193, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAlwaysOpenPDFExternally(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(193, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SpellChecking() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(194, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSpellChecking(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(194, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) SpellCheckerDicts() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(195, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetSpellCheckerDicts(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(195, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) HasValidMainFrame() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(196, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) FrameCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(197, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TChromiumCore) DragOperations() cefTypes.TCefDragOperations {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(198, 0, m.Instance())
	return cefTypes.TCefDragOperations(r)
}

func (m *TChromiumCore) SetDragOperations(value cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(198, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) AudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(199, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetAudioMuted(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(199, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) Fullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(200, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) IsRenderProcessUnresponsive() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(201, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SafeSearch() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(202, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetSafeSearch(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(202, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) YouTubeRestrict() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(203, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetYouTubeRestrict(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(203, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) PrintingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(204, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetPrintingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(204, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) AcceptLanguageList() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(205, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetAcceptLanguageList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(205, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) AcceptCookies() cefTypes.TCefCookiePref {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(206, 0, m.Instance())
	return cefTypes.TCefCookiePref(r)
}

func (m *TChromiumCore) SetAcceptCookies(value cefTypes.TCefCookiePref) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(206, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) Block3rdPartyCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(207, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetBlock3rdPartyCookies(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(207, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) MultiBrowserMode() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(208, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetMultiBrowserMode(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(208, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) DefaultWindowInfoExStyle() types.DWORD {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(209, 0, m.Instance())
	return types.DWORD(r)
}

func (m *TChromiumCore) SetDefaultWindowInfoExStyle(value types.DWORD) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(209, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) Offline() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(210, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetOffline(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(210, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) QuicAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(211, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetQuicAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(211, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) JavascriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(212, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetJavascriptEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(212, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) LoadImagesAutomatically() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(213, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) SetLoadImagesAutomatically(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(213, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumCore) BatterySaverModeState() cefTypes.TCefBatterySaverModeState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(214, 0, m.Instance())
	return cefTypes.TCefBatterySaverModeState(r)
}

func (m *TChromiumCore) SetBatterySaverModeState(value cefTypes.TCefBatterySaverModeState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(214, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) HighEfficiencyModeState() cefTypes.TCefHighEfficiencyModeState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(215, 0, m.Instance())
	return cefTypes.TCefHighEfficiencyModeState(r)
}

func (m *TChromiumCore) SetHighEfficiencyModeState(value cefTypes.TCefHighEfficiencyModeState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(215, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) CanFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumCoreAPI().SysCallN(216, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumCore) EnableFocusDelayMs() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(217, 0, m.Instance())
	return uint32(r)
}

func (m *TChromiumCore) SetEnableFocusDelayMs(value uint32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(217, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) XDisplay() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(218, m.Instance())
	return uintptr(r)
}

func (m *TChromiumCore) WebRTCIPHandlingPolicy() cefTypes.TCefWebRTCHandlingPolicy {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(219, 0, m.Instance())
	return cefTypes.TCefWebRTCHandlingPolicy(r)
}

func (m *TChromiumCore) SetWebRTCIPHandlingPolicy(value cefTypes.TCefWebRTCHandlingPolicy) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(219, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WebRTCMultipleRoutes() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(220, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetWebRTCMultipleRoutes(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(220, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) WebRTCNonproxiedUDP() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(221, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetWebRTCNonproxiedUDP(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(221, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyType() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(222, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetProxyType(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(222, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyScheme() cefTypes.TCefProxyScheme {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(223, 0, m.Instance())
	return cefTypes.TCefProxyScheme(r)
}

func (m *TChromiumCore) SetProxyScheme(value cefTypes.TCefProxyScheme) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(223, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyServer() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(224, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetProxyServer(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(224, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(225, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetProxyPort(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(225, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) ProxyUsername() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(226, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetProxyUsername(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(226, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyPassword() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(227, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetProxyPassword(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(227, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyScriptURL() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(228, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetProxyScriptURL(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(228, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) ProxyByPassList() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(229, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetProxyByPassList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(229, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) MaxConnectionsPerProxy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(230, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumCore) SetMaxConnectionsPerProxy(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(230, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) DownloadBubble() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(231, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetDownloadBubble(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(231, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) HTTPSUpgrade() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumCoreAPI().SysCallN(232, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumCore) SetHTTPSUpgrade(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(232, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumCore) HSTSPolicyBypassList() string {
	if !m.IsValid() {
		return ""
	}
	r := chromiumCoreAPI().SysCallN(233, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TChromiumCore) SetHSTSPolicyBypassList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumCoreAPI().SysCallN(233, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumCore) SetOnTextResultAvailable(fn TOnTextResultAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextResultAvailableEvent(fn)
	base.SetEvent(m, 234, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPdfPrintFinished(fn TOnPdfPrintFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPdfPrintFinishedEvent(fn)
	base.SetEvent(m, 235, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrefsAvailable(fn TOnPrefsAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrefsAvailableEvent(fn)
	base.SetEvent(m, 236, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrefsUpdated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 237, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesDeleted(fn TOnCookiesDeletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookiesDeletedEvent(fn)
	base.SetEvent(m, 238, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResolvedIPsAvailableEvent(fn)
	base.SetEvent(m, 239, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationVisitorResultAvailableEvent(fn)
	base.SetEvent(m, 240, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDownloadImageFinished(fn TOnDownloadImageFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadImageFinishedEvent(fn)
	base.SetEvent(m, 241, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesFlushed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 242, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCertificateExceptionsCleared(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 243, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnHttpAuthCredentialsCleared(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 244, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAllConnectionsClosed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 245, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExecuteTaskOnCefThread(fn)
	base.SetEvent(m, 246, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookiesVisited(fn TOnCookiesVisited) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookiesVisited(fn)
	base.SetEvent(m, 247, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieVisitorDestroyed(fn)
	base.SetEvent(m, 248, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCookieSet(fn TOnCookieSet) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCookieSet(fn)
	base.SetEvent(m, 249, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnZoomPctAvailable(fn TOnZoomPctAvailable) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnZoomPctAvailable(fn)
	base.SetEvent(m, 250, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaRouteCreateFinishedEvent(fn)
	base.SetEvent(m, 251, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaSinkDeviceInfoEvent(fn)
	base.SetEvent(m, 252, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanFocus(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 253, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBrowserCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 254, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnWidgetCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 255, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 256, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceived) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessMessageReceived(fn)
	base.SetEvent(m, 257, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadStart(fn TOnLoadStart) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadStart(fn)
	base.SetEvent(m, 258, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadEnd(fn TOnLoadEnd) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadEnd(fn)
	base.SetEvent(m, 259, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadError(fn TOnLoadError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadError(fn)
	base.SetEvent(m, 260, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadingStateChange(fn TOnLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadingStateChange(fn)
	base.SetEvent(m, 261, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTakeFocus(fn TOnTakeFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTakeFocus(fn)
	base.SetEvent(m, 262, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSetFocus(fn TOnSetFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSetFocus(fn)
	base.SetEvent(m, 263, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGotFocus(fn TOnGotFocus) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGotFocus(fn)
	base.SetEvent(m, 264, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeContextMenu(fn TOnBeforeContextMenu) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeContextMenu(fn)
	base.SetEvent(m, 265, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRunContextMenu(fn TOnRunContextMenu) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRunContextMenu(fn)
	base.SetEvent(m, 266, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnContextMenuCommand(fn TOnContextMenuCommand) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuCommand(fn)
	base.SetEvent(m, 267, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnContextMenuDismissed(fn TOnContextMenuDismissed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuDismissed(fn)
	base.SetEvent(m, 268, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRunQuickMenu(fn TOnRunQuickMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRunQuickMenuEvent(fn)
	base.SetEvent(m, 269, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnQuickMenuCommand(fn TOnQuickMenuCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnQuickMenuCommandEvent(fn)
	base.SetEvent(m, 270, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnQuickMenuDismissed(fn TOnQuickMenuDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnQuickMenuDismissedEvent(fn)
	base.SetEvent(m, 271, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPreKeyEvent(fn TOnPreKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPreKeyEvent(fn)
	base.SetEvent(m, 272, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnKeyEvent(fn TOnKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnKeyEvent(fn)
	base.SetEvent(m, 273, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAddressChange(fn TOnAddressChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAddressChange(fn)
	base.SetEvent(m, 274, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTitleChange(fn TOnTitleChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTitleChange(fn)
	base.SetEvent(m, 275, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFavIconUrlChange(fn TOnFavIconUrlChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFavIconUrlChange(fn)
	base.SetEvent(m, 276, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFullScreenModeChange(fn TOnFullScreenModeChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFullScreenModeChange(fn)
	base.SetEvent(m, 277, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTooltip(fn TOnTooltip) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTooltip(fn)
	base.SetEvent(m, 278, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnStatusMessage(fn TOnStatusMessage) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStatusMessage(fn)
	base.SetEvent(m, 279, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnConsoleMessage(fn TOnConsoleMessage) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnConsoleMessage(fn)
	base.SetEvent(m, 280, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAutoResize(fn TOnAutoResize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAutoResize(fn)
	base.SetEvent(m, 281, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnLoadingProgressChange(fn TOnLoadingProgressChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLoadingProgressChange(fn)
	base.SetEvent(m, 282, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCursorChange(fn TOnCursorChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCursorChange(fn)
	base.SetEvent(m, 283, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMediaAccessChange(fn TOnMediaAccessChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMediaAccessChange(fn)
	base.SetEvent(m, 284, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanDownload(fn TOnCanDownloadEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanDownloadEvent(fn)
	base.SetEvent(m, 285, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeDownload(fn TOnBeforeDownload) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeDownload(fn)
	base.SetEvent(m, 286, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDownloadUpdated(fn TOnDownloadUpdated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadUpdated(fn)
	base.SetEvent(m, 287, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnJsdialog(fn TOnJsdialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnJsdialog(fn)
	base.SetEvent(m, 288, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeUnloadDialog(fn)
	base.SetEvent(m, 289, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResetDialogState(fn TOnResetDialogState) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResetDialogState(fn)
	base.SetEvent(m, 290, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDialogClosed(fn TOnDialogClosed) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDialogClosed(fn)
	base.SetEvent(m, 291, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforePopup(fn TOnBeforePopup) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforePopup(fn)
	base.SetEvent(m, 292, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeDevToolsPopup(fn TOnBeforeDevToolsPopup) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeDevToolsPopup(fn)
	base.SetEvent(m, 293, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAfterCreated(fn TOnAfterCreated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAfterCreated(fn)
	base.SetEvent(m, 294, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeClose(fn TOnBeforeClose) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeClose(fn)
	base.SetEvent(m, 295, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnClose(fn TOnClose) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClose(fn)
	base.SetEvent(m, 296, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeBrowse(fn TOnBeforeBrowse) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBrowse(fn)
	base.SetEvent(m, 297, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnOpenUrlFromTab(fn)
	base.SetEvent(m, 298, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAuthCredentials(fn TOnGetAuthCredentials) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAuthCredentials(fn)
	base.SetEvent(m, 299, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCertificateError(fn TOnCertificateError) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCertificateError(fn)
	base.SetEvent(m, 300, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSelectClientCertificate(fn TOnSelectClientCertificate) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectClientCertificate(fn)
	base.SetEvent(m, 301, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderViewReady(fn TOnRenderViewReady) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderViewReady(fn)
	base.SetEvent(m, 302, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderProcessUnresponsive(fn TOnRenderProcessUnresponsive) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessUnresponsive(fn)
	base.SetEvent(m, 303, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderProcessResponsive(fn TOnRenderProcessResponsive) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessResponsive(fn)
	base.SetEvent(m, 304, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderProcessTerminated(fn)
	base.SetEvent(m, 305, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceRequestHandler(fn)
	base.SetEvent(m, 306, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDocumentAvailableInMainFrame(fn)
	base.SetEvent(m, 307, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeResourceLoad(fn)
	base.SetEvent(m, 308, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceHandler(fn TOnGetResourceHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceHandler(fn)
	base.SetEvent(m, 309, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceRedirect(fn TOnResourceRedirect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceRedirect(fn)
	base.SetEvent(m, 310, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceResponse(fn TOnResourceResponse) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceResponse(fn)
	base.SetEvent(m, 311, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceResponseFilter(fn)
	base.SetEvent(m, 312, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnResourceLoadComplete(fn TOnResourceLoadComplete) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnResourceLoadComplete(fn)
	base.SetEvent(m, 313, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnProtocolExecution(fn TOnProtocolExecution) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProtocolExecution(fn)
	base.SetEvent(m, 314, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanSendCookie(fn TOnCanSendCookie) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanSendCookie(fn)
	base.SetEvent(m, 315, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnCanSaveCookie(fn TOnCanSaveCookie) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanSaveCookie(fn)
	base.SetEvent(m, 316, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFileDialog(fn TOnFileDialog) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFileDialog(fn)
	base.SetEvent(m, 317, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAccessibilityHandler(fn)
	base.SetEvent(m, 318, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetRootScreenRect(fn TOnGetRootScreenRect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetRootScreenRect(fn)
	base.SetEvent(m, 319, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetViewRect(fn TOnGetViewRect) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetViewRect(fn)
	base.SetEvent(m, 320, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetScreenPoint(fn TOnGetScreenPoint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetScreenPoint(fn)
	base.SetEvent(m, 321, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetScreenInfo(fn TOnGetScreenInfo) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetScreenInfo(fn)
	base.SetEvent(m, 322, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPopupShow(fn TOnPopupShow) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupShow(fn)
	base.SetEvent(m, 323, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPopupSize(fn TOnPopupSize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupSize(fn)
	base.SetEvent(m, 324, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPaint(fn TOnPaint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPaint(fn)
	base.SetEvent(m, 325, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAcceleratedPaint(fn TOnAcceleratedPaint) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceleratedPaint(fn)
	base.SetEvent(m, 326, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetTouchHandleSize(fn)
	base.SetEvent(m, 327, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTouchHandleStateChanged(fn)
	base.SetEvent(m, 328, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnStartDragging(fn TOnStartDragging) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStartDragging(fn)
	base.SetEvent(m, 329, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnUpdateDragCursor(fn TOnUpdateDragCursor) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUpdateDragCursor(fn)
	base.SetEvent(m, 330, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScrollOffsetChanged(fn)
	base.SetEvent(m, 331, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIMECompositionRangeChanged(fn)
	base.SetEvent(m, 332, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnTextSelectionChanged(fn TOnTextSelectionChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTextSelectionChanged(fn)
	base.SetEvent(m, 333, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnVirtualKeyboardRequested(fn)
	base.SetEvent(m, 334, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDragEnter(fn TOnDragEnter) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDragEnter(fn)
	base.SetEvent(m, 335, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDraggableRegionsChanged(fn)
	base.SetEvent(m, 336, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFindResult(fn TOnFindResult) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFindResult(fn)
	base.SetEvent(m, 337, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRequestContextInitialized(fn TOnRequestContextInitialized) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestContextInitialized(fn)
	base.SetEvent(m, 338, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetResourceRequestHandler(fn)
	base.SetEvent(m, 339, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnSinks(fn TOnSinksEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSinksEvent(fn)
	base.SetEvent(m, 340, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRoutes(fn TOnRoutesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRoutesEvent(fn)
	base.SetEvent(m, 341, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRouteStateChanged(fn TOnRouteStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRouteStateChangedEvent(fn)
	base.SetEvent(m, 342, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRouteMessageReceived(fn TOnRouteMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRouteMessageReceivedEvent(fn)
	base.SetEvent(m, 343, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetAudioParameters(fn TOnGetAudioParametersEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetAudioParametersEvent(fn)
	base.SetEvent(m, 344, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamStarted(fn TOnAudioStreamStartedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamStartedEvent(fn)
	base.SetEvent(m, 345, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamPacket(fn TOnAudioStreamPacketEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamPacketEvent(fn)
	base.SetEvent(m, 346, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamStopped(fn TOnAudioStreamStoppedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamStoppedEvent(fn)
	base.SetEvent(m, 347, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnAudioStreamError(fn TOnAudioStreamErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAudioStreamErrorEvent(fn)
	base.SetEvent(m, 348, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMessage(fn TOnDevToolsMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMessageEvent(fn)
	base.SetEvent(m, 349, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsRawMessage(fn TOnDevToolsRawMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsRawMessageEvent(fn)
	base.SetEvent(m, 350, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMethodResult(fn TOnDevToolsMethodResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMethodResultEvent(fn)
	base.SetEvent(m, 351, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsMethodRawResultEvent(fn)
	base.SetEvent(m, 352, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsEvent(fn TOnDevToolsEventEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsEventEvent(fn)
	base.SetEvent(m, 353, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsRawEvent(fn TOnDevToolsEventRawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsEventRawEvent(fn)
	base.SetEvent(m, 354, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsAgentAttachedEvent(fn)
	base.SetEvent(m, 355, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetachedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsAgentDetachedEvent(fn)
	base.SetEvent(m, 356, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionLoadFailed(fn TOnExtensionLoadFailedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionLoadFailedEvent(fn)
	base.SetEvent(m, 357, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionLoaded(fn TOnExtensionLoadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionLoadedEvent(fn)
	base.SetEvent(m, 358, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionUnloaded(fn TOnExtensionUnloadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExtensionUnloadedEvent(fn)
	base.SetEvent(m, 359, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBackgroundBrowserEvent(fn)
	base.SetEvent(m, 360, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionBeforeBrowser(fn TOnBeforeBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeBrowserEvent(fn)
	base.SetEvent(m, 361, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetActiveBrowserEvent(fn)
	base.SetEvent(m, 362, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowserEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCanAccessBrowserEvent(fn)
	base.SetEvent(m, 363, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnExtensionGetExtensionResource(fn TOnGetExtensionResourceEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetExtensionResourceEvent(fn)
	base.SetEvent(m, 364, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintStart(fn TOnPrintStartEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintStartEvent(fn)
	base.SetEvent(m, 365, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintSettings(fn TOnPrintSettingsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintSettingsEvent(fn)
	base.SetEvent(m, 366, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintDialog(fn TOnPrintDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintDialogEvent(fn)
	base.SetEvent(m, 367, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintJob(fn TOnPrintJobEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintJobEvent(fn)
	base.SetEvent(m, 368, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnPrintReset(fn TOnPrintResetEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintResetEvent(fn)
	base.SetEvent(m, 369, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnGetPDFPaperSize(fn TOnGetPDFPaperSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetPDFPaperSizeEvent(fn)
	base.SetEvent(m, 370, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameCreated(fn TOnFrameCreated) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameCreated(fn)
	base.SetEvent(m, 371, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameAttached(fn TOnFrameAttached) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameAttached(fn)
	base.SetEvent(m, 372, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnFrameDetached(fn TOnFrameDetached) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameDetached(fn)
	base.SetEvent(m, 373, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnMainFrameChanged(fn TOnMainFrameChanged) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMainFrameChanged(fn)
	base.SetEvent(m, 374, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnChromeCommand(fn TOnChromeCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnChromeCommandEvent(fn)
	base.SetEvent(m, 375, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIsChromeAppMenuItemVisible(fn TOnIsChromeAppMenuItemVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsChromeAppMenuItemVisibleEvent(fn)
	base.SetEvent(m, 376, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIsChromeAppMenuItemEnabled(fn TOnIsChromeAppMenuItemEnabledEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsChromeAppMenuItemEnabledEvent(fn)
	base.SetEvent(m, 377, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIsChromePageActionIconVisible(fn TOnIsChromePageActionIconVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsChromePageActionIconVisibleEvent(fn)
	base.SetEvent(m, 378, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnIsChromeToolbarButtonVisible(fn TOnIsChromeToolbarButtonVisibleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsChromeToolbarButtonVisibleEvent(fn)
	base.SetEvent(m, 379, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermissionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRequestMediaAccessPermissionEvent(fn)
	base.SetEvent(m, 380, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnShowPermissionPrompt(fn TOnShowPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnShowPermissionPromptEvent(fn)
	base.SetEvent(m, 381, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumCore) SetOnDismissPermissionPrompt(fn TOnDismissPermissionPromptEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDismissPermissionPromptEvent(fn)
	base.SetEvent(m, 382, chromiumCoreAPI(), api.MakeEventDataPtr(cb))
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
			/* 8 */ imports.NewTable("TChromiumCore_CreateBrowserWithWindowHandleRectStringRequestContextDictionaryValueBool", 0), // function CreateBrowserWithWindowHandleRectStringRequestContextDictionaryValueBool
			/* 9 */ imports.NewTable("TChromiumCore_CreateBrowserWithStringBrowserViewComponentRequestContextDictionaryValue", 0), // function CreateBrowserWithStringBrowserViewComponentRequestContextDictionaryValue
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
			/* 25 */ imports.NewTable("TChromiumCore_CanExecuteChromeCommand", 0), // function CanExecuteChromeCommand
			/* 26 */ imports.NewTable("TChromiumCore_CreateUrlRequestWithRequestUrlrequestClientStringX2", 0), // function CreateUrlRequestWithRequestUrlrequestClientStringX2
			/* 27 */ imports.NewTable("TChromiumCore_CreateUrlRequestWithRequestUrlrequestClientFrame", 0), // function CreateUrlRequestWithRequestUrlrequestClientFrame
			/* 28 */ imports.NewTable("TChromiumCore_AddObserver", 0), // function AddObserver
			/* 29 */ imports.NewTable("TChromiumCore_GetSource", 0), // function GetSource
			/* 30 */ imports.NewTable("TChromiumCore_GetWebsiteSetting", 0), // function GetWebsiteSetting
			/* 31 */ imports.NewTable("TChromiumCore_GetContentSetting", 0), // function GetContentSetting
			/* 32 */ imports.NewTable("TChromiumCore_CloseBrowser", 0), // procedure CloseBrowser
			/* 33 */ imports.NewTable("TChromiumCore_CloseAllBrowsers", 0), // procedure CloseAllBrowsers
			/* 34 */ imports.NewTable("TChromiumCore_InitializeDragAndDrop", 0), // procedure InitializeDragAndDrop
			/* 35 */ imports.NewTable("TChromiumCore_ShutdownDragAndDrop", 0), // procedure ShutdownDragAndDrop
			/* 36 */ imports.NewTable("TChromiumCore_LoadURLWithStringX3", 0), // procedure LoadURLWithStringX3
			/* 37 */ imports.NewTable("TChromiumCore_LoadURLWithStringFrame", 0), // procedure LoadURLWithStringFrame
			/* 38 */ imports.NewTable("TChromiumCore_LoadStringWithStringX3", 0), // procedure LoadStringWithStringX3
			/* 39 */ imports.NewTable("TChromiumCore_LoadStringWithStringFrame", 0), // procedure LoadStringWithStringFrame
			/* 40 */ imports.NewTable("TChromiumCore_LoadResourceWithCustomMemoryStreamStringX4", 0), // procedure LoadResourceWithCustomMemoryStreamStringX4
			/* 41 */ imports.NewTable("TChromiumCore_LoadResourceWithCustomMemoryStreamStringX2Frame", 0), // procedure LoadResourceWithCustomMemoryStreamStringX2Frame
			/* 42 */ imports.NewTable("TChromiumCore_LoadRequest", 0), // procedure LoadRequest
			/* 43 */ imports.NewTable("TChromiumCore_GoBack", 0), // procedure GoBack
			/* 44 */ imports.NewTable("TChromiumCore_GoForward", 0), // procedure GoForward
			/* 45 */ imports.NewTable("TChromiumCore_Reload", 0), // procedure Reload
			/* 46 */ imports.NewTable("TChromiumCore_ReloadIgnoreCache", 0), // procedure ReloadIgnoreCache
			/* 47 */ imports.NewTable("TChromiumCore_StopLoad", 0), // procedure StopLoad
			/* 48 */ imports.NewTable("TChromiumCore_StartDownload", 0), // procedure StartDownload
			/* 49 */ imports.NewTable("TChromiumCore_DownloadImage", 0), // procedure DownloadImage
			/* 50 */ imports.NewTable("TChromiumCore_SimulateMouseWheel", 0), // procedure SimulateMouseWheel
			/* 51 */ imports.NewTable("TChromiumCore_SimulateKeyEvent", 0), // procedure SimulateKeyEvent
			/* 52 */ imports.NewTable("TChromiumCore_SimulateMouseEvent", 0), // procedure SimulateMouseEvent
			/* 53 */ imports.NewTable("TChromiumCore_SimulateEditingCommand", 0), // procedure SimulateEditingCommand
			/* 54 */ imports.NewTable("TChromiumCore_RetrieveHTMLWithStringX2", 0), // procedure RetrieveHTMLWithStringX2
			/* 55 */ imports.NewTable("TChromiumCore_RetrieveHTMLWithFrame", 0), // procedure RetrieveHTMLWithFrame
			/* 56 */ imports.NewTable("TChromiumCore_RetrieveTextWithStringX2", 0), // procedure RetrieveTextWithStringX2
			/* 57 */ imports.NewTable("TChromiumCore_RetrieveTextWithFrame", 0), // procedure RetrieveTextWithFrame
			/* 58 */ imports.NewTable("TChromiumCore_GetNavigationEntries", 0), // procedure GetNavigationEntries
			/* 59 */ imports.NewTable("TChromiumCore_ExecuteJavaScriptWithStringX4Int", 0), // procedure ExecuteJavaScriptWithStringX4Int
			/* 60 */ imports.NewTable("TChromiumCore_ExecuteJavaScriptWithStringX2FrameInt", 0), // procedure ExecuteJavaScriptWithStringX2FrameInt
			/* 61 */ imports.NewTable("TChromiumCore_UpdatePreferences", 0), // procedure UpdatePreferences
			/* 62 */ imports.NewTable("TChromiumCore_SavePreferences", 0), // procedure SavePreferences
			/* 63 */ imports.NewTable("TChromiumCore_ResolveHost", 0), // procedure ResolveHost
			/* 64 */ imports.NewTable("TChromiumCore_SetUserAgentOverride", 0), // procedure SetUserAgentOverride
			/* 65 */ imports.NewTable("TChromiumCore_ClearDataForOrigin", 0), // procedure ClearDataForOrigin
			/* 66 */ imports.NewTable("TChromiumCore_ClearCache", 0), // procedure ClearCache
			/* 67 */ imports.NewTable("TChromiumCore_ToggleAudioMuted", 0), // procedure ToggleAudioMuted
			/* 68 */ imports.NewTable("TChromiumCore_ShowDevTools", 0), // procedure ShowDevTools
			/* 69 */ imports.NewTable("TChromiumCore_CloseDevTools", 0), // procedure CloseDevTools
			/* 70 */ imports.NewTable("TChromiumCore_CloseDevToolsWithWindowHandle", 0), // procedure CloseDevToolsWithWindowHandle
			/* 71 */ imports.NewTable("TChromiumCore_Find", 0), // procedure Find
			/* 72 */ imports.NewTable("TChromiumCore_StopFinding", 0), // procedure StopFinding
			/* 73 */ imports.NewTable("TChromiumCore_Print", 0), // procedure Print
			/* 74 */ imports.NewTable("TChromiumCore_PrintToPDF", 0), // procedure PrintToPDF
			/* 75 */ imports.NewTable("TChromiumCore_ClipboardCopy", 0), // procedure ClipboardCopy
			/* 76 */ imports.NewTable("TChromiumCore_ClipboardPaste", 0), // procedure ClipboardPaste
			/* 77 */ imports.NewTable("TChromiumCore_ClipboardCut", 0), // procedure ClipboardCut
			/* 78 */ imports.NewTable("TChromiumCore_ClipboardUndo", 0), // procedure ClipboardUndo
			/* 79 */ imports.NewTable("TChromiumCore_ClipboardRedo", 0), // procedure ClipboardRedo
			/* 80 */ imports.NewTable("TChromiumCore_ClipboardDel", 0), // procedure ClipboardDel
			/* 81 */ imports.NewTable("TChromiumCore_SelectAll", 0), // procedure SelectAll
			/* 82 */ imports.NewTable("TChromiumCore_IncZoomStep", 0), // procedure IncZoomStep
			/* 83 */ imports.NewTable("TChromiumCore_DecZoomStep", 0), // procedure DecZoomStep
			/* 84 */ imports.NewTable("TChromiumCore_IncZoomPct", 0), // procedure IncZoomPct
			/* 85 */ imports.NewTable("TChromiumCore_DecZoomPct", 0), // procedure DecZoomPct
			/* 86 */ imports.NewTable("TChromiumCore_ResetZoomStep", 0), // procedure ResetZoomStep
			/* 87 */ imports.NewTable("TChromiumCore_ResetZoomLevel", 0), // procedure ResetZoomLevel
			/* 88 */ imports.NewTable("TChromiumCore_ResetZoomPct", 0), // procedure ResetZoomPct
			/* 89 */ imports.NewTable("TChromiumCore_ReadZoom", 0), // procedure ReadZoom
			/* 90 */ imports.NewTable("TChromiumCore_IncZoomCommand", 0), // procedure IncZoomCommand
			/* 91 */ imports.NewTable("TChromiumCore_DecZoomCommand", 0), // procedure DecZoomCommand
			/* 92 */ imports.NewTable("TChromiumCore_ResetZoomCommand", 0), // procedure ResetZoomCommand
			/* 93 */ imports.NewTable("TChromiumCore_WasResized", 0), // procedure WasResized
			/* 94 */ imports.NewTable("TChromiumCore_WasHidden", 0), // procedure WasHidden
			/* 95 */ imports.NewTable("TChromiumCore_NotifyScreenInfoChanged", 0), // procedure NotifyScreenInfoChanged
			/* 96 */ imports.NewTable("TChromiumCore_NotifyMoveOrResizeStarted", 0), // procedure NotifyMoveOrResizeStarted
			/* 97 */ imports.NewTable("TChromiumCore_Invalidate", 0), // procedure Invalidate
			/* 98 */ imports.NewTable("TChromiumCore_ExitFullscreen", 0), // procedure ExitFullscreen
			/* 99 */ imports.NewTable("TChromiumCore_ExecuteChromeCommand", 0), // procedure ExecuteChromeCommand
			/* 100 */ imports.NewTable("TChromiumCore_SendExternalBeginFrame", 0), // procedure SendExternalBeginFrame
			/* 101 */ imports.NewTable("TChromiumCore_SendKeyEvent", 0), // procedure SendKeyEvent
			/* 102 */ imports.NewTable("TChromiumCore_SendMouseClickEvent", 0), // procedure SendMouseClickEvent
			/* 103 */ imports.NewTable("TChromiumCore_SendMouseMoveEvent", 0), // procedure SendMouseMoveEvent
			/* 104 */ imports.NewTable("TChromiumCore_SendMouseWheelEvent", 0), // procedure SendMouseWheelEvent
			/* 105 */ imports.NewTable("TChromiumCore_SendTouchEvent", 0), // procedure SendTouchEvent
			/* 106 */ imports.NewTable("TChromiumCore_SendCaptureLostEvent", 0), // procedure SendCaptureLostEvent
			/* 107 */ imports.NewTable("TChromiumCore_SendProcessMessageWithProcessIdProcessMessageStringX2", 0), // procedure SendProcessMessageWithProcessIdProcessMessageStringX2
			/* 108 */ imports.NewTable("TChromiumCore_SendProcessMessageWithProcessIdProcessMessageFrame", 0), // procedure SendProcessMessageWithProcessIdProcessMessageFrame
			/* 109 */ imports.NewTable("TChromiumCore_SetFocus", 0), // procedure SetFocus
			/* 110 */ imports.NewTable("TChromiumCore_SetAccessibilityState", 0), // procedure SetAccessibilityState
			/* 111 */ imports.NewTable("TChromiumCore_DragTargetDragEnter", 0), // procedure DragTargetDragEnter
			/* 112 */ imports.NewTable("TChromiumCore_DragTargetDragOver", 0), // procedure DragTargetDragOver
			/* 113 */ imports.NewTable("TChromiumCore_DragTargetDragLeave", 0), // procedure DragTargetDragLeave
			/* 114 */ imports.NewTable("TChromiumCore_DragTargetDrop", 0), // procedure DragTargetDrop
			/* 115 */ imports.NewTable("TChromiumCore_DragSourceEndedAt", 0), // procedure DragSourceEndedAt
			/* 116 */ imports.NewTable("TChromiumCore_DragSourceSystemDragEnded", 0), // procedure DragSourceSystemDragEnded
			/* 117 */ imports.NewTable("TChromiumCore_IMESetComposition", 0), // procedure IMESetComposition
			/* 118 */ imports.NewTable("TChromiumCore_IMECommitText", 0), // procedure IMECommitText
			/* 119 */ imports.NewTable("TChromiumCore_IMEFinishComposingText", 0), // procedure IMEFinishComposingText
			/* 120 */ imports.NewTable("TChromiumCore_IMECancelComposition", 0), // procedure IMECancelComposition
			/* 121 */ imports.NewTable("TChromiumCore_ReplaceMisspelling", 0), // procedure ReplaceMisspelling
			/* 122 */ imports.NewTable("TChromiumCore_AddWordToDictionary", 0), // procedure AddWordToDictionary
			/* 123 */ imports.NewTable("TChromiumCore_UpdateBrowserSize", 0), // procedure UpdateBrowserSize
			/* 124 */ imports.NewTable("TChromiumCore_UpdateXWindowVisibility", 0), // procedure UpdateXWindowVisibility
			/* 125 */ imports.NewTable("TChromiumCore_NotifyCurrentSinks", 0), // procedure NotifyCurrentSinks
			/* 126 */ imports.NewTable("TChromiumCore_NotifyCurrentRoutes", 0), // procedure NotifyCurrentRoutes
			/* 127 */ imports.NewTable("TChromiumCore_CreateRoute", 0), // procedure CreateRoute
			/* 128 */ imports.NewTable("TChromiumCore_GetDeviceInfo", 0), // procedure GetDeviceInfo
			/* 129 */ imports.NewTable("TChromiumCore_SetWebsiteSetting", 0), // procedure SetWebsiteSetting
			/* 130 */ imports.NewTable("TChromiumCore_SetContentSetting", 0), // procedure SetContentSetting
			/* 131 */ imports.NewTable("TChromiumCore_SetChromeColorScheme", 0), // procedure SetChromeColorScheme
			/* 132 */ imports.NewTable("TChromiumCore_DefaultUrl", 0), // property DefaultUrl
			/* 133 */ imports.NewTable("TChromiumCore_Options", 0), // property Options
			/* 134 */ imports.NewTable("TChromiumCore_FontOptions", 0), // property FontOptions
			/* 135 */ imports.NewTable("TChromiumCore_PDFPrintOptions", 0), // property PDFPrintOptions
			/* 136 */ imports.NewTable("TChromiumCore_DefaultEncoding", 0), // property DefaultEncoding
			/* 137 */ imports.NewTable("TChromiumCore_BrowserId", 0), // property BrowserId
			/* 138 */ imports.NewTable("TChromiumCore_Browser", 0), // property Browser
			/* 139 */ imports.NewTable("TChromiumCore_BrowserById", 0), // property BrowserById
			/* 140 */ imports.NewTable("TChromiumCore_BrowserCount", 0), // property BrowserCount
			/* 141 */ imports.NewTable("TChromiumCore_BrowserIdByIndex", 0), // property BrowserIdByIndex
			/* 142 */ imports.NewTable("TChromiumCore_CefClient", 0), // property CefClient
			/* 143 */ imports.NewTable("TChromiumCore_ReqContextHandler", 0), // property ReqContextHandler
			/* 144 */ imports.NewTable("TChromiumCore_ResourceRequestHandler", 0), // property ResourceRequestHandler
			/* 145 */ imports.NewTable("TChromiumCore_CefWindowInfo", 0), // property CefWindowInfo
			/* 146 */ imports.NewTable("TChromiumCore_VisibleNavigationEntry", 0), // property VisibleNavigationEntry
			/* 147 */ imports.NewTable("TChromiumCore_RuntimeStyle", 0), // property RuntimeStyle
			/* 148 */ imports.NewTable("TChromiumCore_RequestContext", 0), // property RequestContext
			/* 149 */ imports.NewTable("TChromiumCore_MediaRouter", 0), // property MediaRouter
			/* 150 */ imports.NewTable("TChromiumCore_MediaObserver", 0), // property MediaObserver
			/* 151 */ imports.NewTable("TChromiumCore_MediaObserverReg", 0), // property MediaObserverReg
			/* 152 */ imports.NewTable("TChromiumCore_DevToolsMsgObserver", 0), // property DevToolsMsgObserver
			/* 153 */ imports.NewTable("TChromiumCore_DevToolsMsgObserverReg", 0), // property DevToolsMsgObserverReg
			/* 154 */ imports.NewTable("TChromiumCore_ExtensionHandler", 0), // property ExtensionHandler
			/* 155 */ imports.NewTable("TChromiumCore_MultithreadApp", 0), // property MultithreadApp
			/* 156 */ imports.NewTable("TChromiumCore_IsLoading", 0), // property IsLoading
			/* 157 */ imports.NewTable("TChromiumCore_HasDocument", 0), // property HasDocument
			/* 158 */ imports.NewTable("TChromiumCore_HasView", 0), // property HasView
			/* 159 */ imports.NewTable("TChromiumCore_HasDevTools", 0), // property HasDevTools
			/* 160 */ imports.NewTable("TChromiumCore_HasClientHandler", 0), // property HasClientHandler
			/* 161 */ imports.NewTable("TChromiumCore_HasBrowser", 0), // property HasBrowser
			/* 162 */ imports.NewTable("TChromiumCore_CanGoBack", 0), // property CanGoBack
			/* 163 */ imports.NewTable("TChromiumCore_CanGoForward", 0), // property CanGoForward
			/* 164 */ imports.NewTable("TChromiumCore_IsPopUp", 0), // property IsPopUp
			/* 165 */ imports.NewTable("TChromiumCore_WindowHandle", 0), // property WindowHandle
			/* 166 */ imports.NewTable("TChromiumCore_OpenerWindowHandle", 0), // property OpenerWindowHandle
			/* 167 */ imports.NewTable("TChromiumCore_BrowserHandle", 0), // property BrowserHandle
			/* 168 */ imports.NewTable("TChromiumCore_WidgetHandle", 0), // property WidgetHandle
			/* 169 */ imports.NewTable("TChromiumCore_RenderHandle", 0), // property RenderHandle
			/* 170 */ imports.NewTable("TChromiumCore_FrameIsFocused", 0), // property FrameIsFocused
			/* 171 */ imports.NewTable("TChromiumCore_Initialized", 0), // property Initialized
			/* 172 */ imports.NewTable("TChromiumCore_RequestContextCache", 0), // property RequestContextCache
			/* 173 */ imports.NewTable("TChromiumCore_RequestContextIsGlobal", 0), // property RequestContextIsGlobal
			/* 174 */ imports.NewTable("TChromiumCore_ChromeColorSchemeMode", 0), // property ChromeColorSchemeMode
			/* 175 */ imports.NewTable("TChromiumCore_ChromeColorSchemeColor", 0), // property ChromeColorSchemeColor
			/* 176 */ imports.NewTable("TChromiumCore_ChromeColorSchemeVariant", 0), // property ChromeColorSchemeVariant
			/* 177 */ imports.NewTable("TChromiumCore_DocumentURL", 0), // property DocumentURL
			/* 178 */ imports.NewTable("TChromiumCore_ZoomLevel", 0), // property ZoomLevel
			/* 179 */ imports.NewTable("TChromiumCore_DefaultZoomLevel", 0), // property DefaultZoomLevel
			/* 180 */ imports.NewTable("TChromiumCore_CanIncZoom", 0), // property CanIncZoom
			/* 181 */ imports.NewTable("TChromiumCore_CanDecZoom", 0), // property CanDecZoom
			/* 182 */ imports.NewTable("TChromiumCore_CanResetZoom", 0), // property CanResetZoom
			/* 183 */ imports.NewTable("TChromiumCore_ZoomPct", 0), // property ZoomPct
			/* 184 */ imports.NewTable("TChromiumCore_ZoomStep", 0), // property ZoomStep
			/* 185 */ imports.NewTable("TChromiumCore_WindowlessFrameRate", 0), // property WindowlessFrameRate
			/* 186 */ imports.NewTable("TChromiumCore_CustomHeaderName", 0), // property CustomHeaderName
			/* 187 */ imports.NewTable("TChromiumCore_CustomHeaderValue", 0), // property CustomHeaderValue
			/* 188 */ imports.NewTable("TChromiumCore_DoNotTrack", 0), // property DoNotTrack
			/* 189 */ imports.NewTable("TChromiumCore_SendReferrer", 0), // property SendReferrer
			/* 190 */ imports.NewTable("TChromiumCore_HyperlinkAuditing", 0), // property HyperlinkAuditing
			/* 191 */ imports.NewTable("TChromiumCore_AllowOutdatedPlugins", 0), // property AllowOutdatedPlugins
			/* 192 */ imports.NewTable("TChromiumCore_AlwaysAuthorizePlugins", 0), // property AlwaysAuthorizePlugins
			/* 193 */ imports.NewTable("TChromiumCore_AlwaysOpenPDFExternally", 0), // property AlwaysOpenPDFExternally
			/* 194 */ imports.NewTable("TChromiumCore_SpellChecking", 0), // property SpellChecking
			/* 195 */ imports.NewTable("TChromiumCore_SpellCheckerDicts", 0), // property SpellCheckerDicts
			/* 196 */ imports.NewTable("TChromiumCore_HasValidMainFrame", 0), // property HasValidMainFrame
			/* 197 */ imports.NewTable("TChromiumCore_FrameCount", 0), // property FrameCount
			/* 198 */ imports.NewTable("TChromiumCore_DragOperations", 0), // property DragOperations
			/* 199 */ imports.NewTable("TChromiumCore_AudioMuted", 0), // property AudioMuted
			/* 200 */ imports.NewTable("TChromiumCore_Fullscreen", 0), // property Fullscreen
			/* 201 */ imports.NewTable("TChromiumCore_IsRenderProcessUnresponsive", 0), // property IsRenderProcessUnresponsive
			/* 202 */ imports.NewTable("TChromiumCore_SafeSearch", 0), // property SafeSearch
			/* 203 */ imports.NewTable("TChromiumCore_YouTubeRestrict", 0), // property YouTubeRestrict
			/* 204 */ imports.NewTable("TChromiumCore_PrintingEnabled", 0), // property PrintingEnabled
			/* 205 */ imports.NewTable("TChromiumCore_AcceptLanguageList", 0), // property AcceptLanguageList
			/* 206 */ imports.NewTable("TChromiumCore_AcceptCookies", 0), // property AcceptCookies
			/* 207 */ imports.NewTable("TChromiumCore_Block3rdPartyCookies", 0), // property Block3rdPartyCookies
			/* 208 */ imports.NewTable("TChromiumCore_MultiBrowserMode", 0), // property MultiBrowserMode
			/* 209 */ imports.NewTable("TChromiumCore_DefaultWindowInfoExStyle", 0), // property DefaultWindowInfoExStyle
			/* 210 */ imports.NewTable("TChromiumCore_Offline", 0), // property Offline
			/* 211 */ imports.NewTable("TChromiumCore_QuicAllowed", 0), // property QuicAllowed
			/* 212 */ imports.NewTable("TChromiumCore_JavascriptEnabled", 0), // property JavascriptEnabled
			/* 213 */ imports.NewTable("TChromiumCore_LoadImagesAutomatically", 0), // property LoadImagesAutomatically
			/* 214 */ imports.NewTable("TChromiumCore_BatterySaverModeState", 0), // property BatterySaverModeState
			/* 215 */ imports.NewTable("TChromiumCore_HighEfficiencyModeState", 0), // property HighEfficiencyModeState
			/* 216 */ imports.NewTable("TChromiumCore_CanFocus", 0), // property CanFocus
			/* 217 */ imports.NewTable("TChromiumCore_EnableFocusDelayMs", 0), // property EnableFocusDelayMs
			/* 218 */ imports.NewTable("TChromiumCore_XDisplay", 0), // property XDisplay
			/* 219 */ imports.NewTable("TChromiumCore_WebRTCIPHandlingPolicy", 0), // property WebRTCIPHandlingPolicy
			/* 220 */ imports.NewTable("TChromiumCore_WebRTCMultipleRoutes", 0), // property WebRTCMultipleRoutes
			/* 221 */ imports.NewTable("TChromiumCore_WebRTCNonproxiedUDP", 0), // property WebRTCNonproxiedUDP
			/* 222 */ imports.NewTable("TChromiumCore_ProxyType", 0), // property ProxyType
			/* 223 */ imports.NewTable("TChromiumCore_ProxyScheme", 0), // property ProxyScheme
			/* 224 */ imports.NewTable("TChromiumCore_ProxyServer", 0), // property ProxyServer
			/* 225 */ imports.NewTable("TChromiumCore_ProxyPort", 0), // property ProxyPort
			/* 226 */ imports.NewTable("TChromiumCore_ProxyUsername", 0), // property ProxyUsername
			/* 227 */ imports.NewTable("TChromiumCore_ProxyPassword", 0), // property ProxyPassword
			/* 228 */ imports.NewTable("TChromiumCore_ProxyScriptURL", 0), // property ProxyScriptURL
			/* 229 */ imports.NewTable("TChromiumCore_ProxyByPassList", 0), // property ProxyByPassList
			/* 230 */ imports.NewTable("TChromiumCore_MaxConnectionsPerProxy", 0), // property MaxConnectionsPerProxy
			/* 231 */ imports.NewTable("TChromiumCore_DownloadBubble", 0), // property DownloadBubble
			/* 232 */ imports.NewTable("TChromiumCore_HTTPSUpgrade", 0), // property HTTPSUpgrade
			/* 233 */ imports.NewTable("TChromiumCore_HSTSPolicyBypassList", 0), // property HSTSPolicyBypassList
			/* 234 */ imports.NewTable("TChromiumCore_OnTextResultAvailable", 0), // event OnTextResultAvailable
			/* 235 */ imports.NewTable("TChromiumCore_OnPdfPrintFinished", 0), // event OnPdfPrintFinished
			/* 236 */ imports.NewTable("TChromiumCore_OnPrefsAvailable", 0), // event OnPrefsAvailable
			/* 237 */ imports.NewTable("TChromiumCore_OnPrefsUpdated", 0), // event OnPrefsUpdated
			/* 238 */ imports.NewTable("TChromiumCore_OnCookiesDeleted", 0), // event OnCookiesDeleted
			/* 239 */ imports.NewTable("TChromiumCore_OnResolvedHostAvailable", 0), // event OnResolvedHostAvailable
			/* 240 */ imports.NewTable("TChromiumCore_OnNavigationVisitorResultAvailable", 0), // event OnNavigationVisitorResultAvailable
			/* 241 */ imports.NewTable("TChromiumCore_OnDownloadImageFinished", 0), // event OnDownloadImageFinished
			/* 242 */ imports.NewTable("TChromiumCore_OnCookiesFlushed", 0), // event OnCookiesFlushed
			/* 243 */ imports.NewTable("TChromiumCore_OnCertificateExceptionsCleared", 0), // event OnCertificateExceptionsCleared
			/* 244 */ imports.NewTable("TChromiumCore_OnHttpAuthCredentialsCleared", 0), // event OnHttpAuthCredentialsCleared
			/* 245 */ imports.NewTable("TChromiumCore_OnAllConnectionsClosed", 0), // event OnAllConnectionsClosed
			/* 246 */ imports.NewTable("TChromiumCore_OnExecuteTaskOnCefThread", 0), // event OnExecuteTaskOnCefThread
			/* 247 */ imports.NewTable("TChromiumCore_OnCookiesVisited", 0), // event OnCookiesVisited
			/* 248 */ imports.NewTable("TChromiumCore_OnCookieVisitorDestroyed", 0), // event OnCookieVisitorDestroyed
			/* 249 */ imports.NewTable("TChromiumCore_OnCookieSet", 0), // event OnCookieSet
			/* 250 */ imports.NewTable("TChromiumCore_OnZoomPctAvailable", 0), // event OnZoomPctAvailable
			/* 251 */ imports.NewTable("TChromiumCore_OnMediaRouteCreateFinished", 0), // event OnMediaRouteCreateFinished
			/* 252 */ imports.NewTable("TChromiumCore_OnMediaSinkDeviceInfo", 0), // event OnMediaSinkDeviceInfo
			/* 253 */ imports.NewTable("TChromiumCore_OnCanFocus", 0), // event OnCanFocus
			/* 254 */ imports.NewTable("TChromiumCore_OnBrowserCompMsg", 0), // event OnBrowserCompMsg
			/* 255 */ imports.NewTable("TChromiumCore_OnWidgetCompMsg", 0), // event OnWidgetCompMsg
			/* 256 */ imports.NewTable("TChromiumCore_OnRenderCompMsg", 0), // event OnRenderCompMsg
			/* 257 */ imports.NewTable("TChromiumCore_OnProcessMessageReceived", 0), // event OnProcessMessageReceived
			/* 258 */ imports.NewTable("TChromiumCore_OnLoadStart", 0), // event OnLoadStart
			/* 259 */ imports.NewTable("TChromiumCore_OnLoadEnd", 0), // event OnLoadEnd
			/* 260 */ imports.NewTable("TChromiumCore_OnLoadError", 0), // event OnLoadError
			/* 261 */ imports.NewTable("TChromiumCore_OnLoadingStateChange", 0), // event OnLoadingStateChange
			/* 262 */ imports.NewTable("TChromiumCore_OnTakeFocus", 0), // event OnTakeFocus
			/* 263 */ imports.NewTable("TChromiumCore_OnSetFocus", 0), // event OnSetFocus
			/* 264 */ imports.NewTable("TChromiumCore_OnGotFocus", 0), // event OnGotFocus
			/* 265 */ imports.NewTable("TChromiumCore_OnBeforeContextMenu", 0), // event OnBeforeContextMenu
			/* 266 */ imports.NewTable("TChromiumCore_OnRunContextMenu", 0), // event OnRunContextMenu
			/* 267 */ imports.NewTable("TChromiumCore_OnContextMenuCommand", 0), // event OnContextMenuCommand
			/* 268 */ imports.NewTable("TChromiumCore_OnContextMenuDismissed", 0), // event OnContextMenuDismissed
			/* 269 */ imports.NewTable("TChromiumCore_OnRunQuickMenu", 0), // event OnRunQuickMenu
			/* 270 */ imports.NewTable("TChromiumCore_OnQuickMenuCommand", 0), // event OnQuickMenuCommand
			/* 271 */ imports.NewTable("TChromiumCore_OnQuickMenuDismissed", 0), // event OnQuickMenuDismissed
			/* 272 */ imports.NewTable("TChromiumCore_OnPreKeyEvent", 0), // event OnPreKeyEvent
			/* 273 */ imports.NewTable("TChromiumCore_OnKeyEvent", 0), // event OnKeyEvent
			/* 274 */ imports.NewTable("TChromiumCore_OnAddressChange", 0), // event OnAddressChange
			/* 275 */ imports.NewTable("TChromiumCore_OnTitleChange", 0), // event OnTitleChange
			/* 276 */ imports.NewTable("TChromiumCore_OnFavIconUrlChange", 0), // event OnFavIconUrlChange
			/* 277 */ imports.NewTable("TChromiumCore_OnFullScreenModeChange", 0), // event OnFullScreenModeChange
			/* 278 */ imports.NewTable("TChromiumCore_OnTooltip", 0), // event OnTooltip
			/* 279 */ imports.NewTable("TChromiumCore_OnStatusMessage", 0), // event OnStatusMessage
			/* 280 */ imports.NewTable("TChromiumCore_OnConsoleMessage", 0), // event OnConsoleMessage
			/* 281 */ imports.NewTable("TChromiumCore_OnAutoResize", 0), // event OnAutoResize
			/* 282 */ imports.NewTable("TChromiumCore_OnLoadingProgressChange", 0), // event OnLoadingProgressChange
			/* 283 */ imports.NewTable("TChromiumCore_OnCursorChange", 0), // event OnCursorChange
			/* 284 */ imports.NewTable("TChromiumCore_OnMediaAccessChange", 0), // event OnMediaAccessChange
			/* 285 */ imports.NewTable("TChromiumCore_OnCanDownload", 0), // event OnCanDownload
			/* 286 */ imports.NewTable("TChromiumCore_OnBeforeDownload", 0), // event OnBeforeDownload
			/* 287 */ imports.NewTable("TChromiumCore_OnDownloadUpdated", 0), // event OnDownloadUpdated
			/* 288 */ imports.NewTable("TChromiumCore_OnJsdialog", 0), // event OnJsdialog
			/* 289 */ imports.NewTable("TChromiumCore_OnBeforeUnloadDialog", 0), // event OnBeforeUnloadDialog
			/* 290 */ imports.NewTable("TChromiumCore_OnResetDialogState", 0), // event OnResetDialogState
			/* 291 */ imports.NewTable("TChromiumCore_OnDialogClosed", 0), // event OnDialogClosed
			/* 292 */ imports.NewTable("TChromiumCore_OnBeforePopup", 0), // event OnBeforePopup
			/* 293 */ imports.NewTable("TChromiumCore_OnBeforeDevToolsPopup", 0), // event OnBeforeDevToolsPopup
			/* 294 */ imports.NewTable("TChromiumCore_OnAfterCreated", 0), // event OnAfterCreated
			/* 295 */ imports.NewTable("TChromiumCore_OnBeforeClose", 0), // event OnBeforeClose
			/* 296 */ imports.NewTable("TChromiumCore_OnClose", 0), // event OnClose
			/* 297 */ imports.NewTable("TChromiumCore_OnBeforeBrowse", 0), // event OnBeforeBrowse
			/* 298 */ imports.NewTable("TChromiumCore_OnOpenUrlFromTab", 0), // event OnOpenUrlFromTab
			/* 299 */ imports.NewTable("TChromiumCore_OnGetAuthCredentials", 0), // event OnGetAuthCredentials
			/* 300 */ imports.NewTable("TChromiumCore_OnCertificateError", 0), // event OnCertificateError
			/* 301 */ imports.NewTable("TChromiumCore_OnSelectClientCertificate", 0), // event OnSelectClientCertificate
			/* 302 */ imports.NewTable("TChromiumCore_OnRenderViewReady", 0), // event OnRenderViewReady
			/* 303 */ imports.NewTable("TChromiumCore_OnRenderProcessUnresponsive", 0), // event OnRenderProcessUnresponsive
			/* 304 */ imports.NewTable("TChromiumCore_OnRenderProcessResponsive", 0), // event OnRenderProcessResponsive
			/* 305 */ imports.NewTable("TChromiumCore_OnRenderProcessTerminated", 0), // event OnRenderProcessTerminated
			/* 306 */ imports.NewTable("TChromiumCore_OnGetResourceRequestHandler_ReqHdlr", 0), // event OnGetResourceRequestHandler_ReqHdlr
			/* 307 */ imports.NewTable("TChromiumCore_OnDocumentAvailableInMainFrame", 0), // event OnDocumentAvailableInMainFrame
			/* 308 */ imports.NewTable("TChromiumCore_OnBeforeResourceLoad", 0), // event OnBeforeResourceLoad
			/* 309 */ imports.NewTable("TChromiumCore_OnGetResourceHandler", 0), // event OnGetResourceHandler
			/* 310 */ imports.NewTable("TChromiumCore_OnResourceRedirect", 0), // event OnResourceRedirect
			/* 311 */ imports.NewTable("TChromiumCore_OnResourceResponse", 0), // event OnResourceResponse
			/* 312 */ imports.NewTable("TChromiumCore_OnGetResourceResponseFilter", 0), // event OnGetResourceResponseFilter
			/* 313 */ imports.NewTable("TChromiumCore_OnResourceLoadComplete", 0), // event OnResourceLoadComplete
			/* 314 */ imports.NewTable("TChromiumCore_OnProtocolExecution", 0), // event OnProtocolExecution
			/* 315 */ imports.NewTable("TChromiumCore_OnCanSendCookie", 0), // event OnCanSendCookie
			/* 316 */ imports.NewTable("TChromiumCore_OnCanSaveCookie", 0), // event OnCanSaveCookie
			/* 317 */ imports.NewTable("TChromiumCore_OnFileDialog", 0), // event OnFileDialog
			/* 318 */ imports.NewTable("TChromiumCore_OnGetAccessibilityHandler", 0), // event OnGetAccessibilityHandler
			/* 319 */ imports.NewTable("TChromiumCore_OnGetRootScreenRect", 0), // event OnGetRootScreenRect
			/* 320 */ imports.NewTable("TChromiumCore_OnGetViewRect", 0), // event OnGetViewRect
			/* 321 */ imports.NewTable("TChromiumCore_OnGetScreenPoint", 0), // event OnGetScreenPoint
			/* 322 */ imports.NewTable("TChromiumCore_OnGetScreenInfo", 0), // event OnGetScreenInfo
			/* 323 */ imports.NewTable("TChromiumCore_OnPopupShow", 0), // event OnPopupShow
			/* 324 */ imports.NewTable("TChromiumCore_OnPopupSize", 0), // event OnPopupSize
			/* 325 */ imports.NewTable("TChromiumCore_OnPaint", 0), // event OnPaint
			/* 326 */ imports.NewTable("TChromiumCore_OnAcceleratedPaint", 0), // event OnAcceleratedPaint
			/* 327 */ imports.NewTable("TChromiumCore_OnGetTouchHandleSize", 0), // event OnGetTouchHandleSize
			/* 328 */ imports.NewTable("TChromiumCore_OnTouchHandleStateChanged", 0), // event OnTouchHandleStateChanged
			/* 329 */ imports.NewTable("TChromiumCore_OnStartDragging", 0), // event OnStartDragging
			/* 330 */ imports.NewTable("TChromiumCore_OnUpdateDragCursor", 0), // event OnUpdateDragCursor
			/* 331 */ imports.NewTable("TChromiumCore_OnScrollOffsetChanged", 0), // event OnScrollOffsetChanged
			/* 332 */ imports.NewTable("TChromiumCore_OnIMECompositionRangeChanged", 0), // event OnIMECompositionRangeChanged
			/* 333 */ imports.NewTable("TChromiumCore_OnTextSelectionChanged", 0), // event OnTextSelectionChanged
			/* 334 */ imports.NewTable("TChromiumCore_OnVirtualKeyboardRequested", 0), // event OnVirtualKeyboardRequested
			/* 335 */ imports.NewTable("TChromiumCore_OnDragEnter", 0), // event OnDragEnter
			/* 336 */ imports.NewTable("TChromiumCore_OnDraggableRegionsChanged", 0), // event OnDraggableRegionsChanged
			/* 337 */ imports.NewTable("TChromiumCore_OnFindResult", 0), // event OnFindResult
			/* 338 */ imports.NewTable("TChromiumCore_OnRequestContextInitialized", 0), // event OnRequestContextInitialized
			/* 339 */ imports.NewTable("TChromiumCore_OnGetResourceRequestHandler_ReqCtxHdlr", 0), // event OnGetResourceRequestHandler_ReqCtxHdlr
			/* 340 */ imports.NewTable("TChromiumCore_OnSinks", 0), // event OnSinks
			/* 341 */ imports.NewTable("TChromiumCore_OnRoutes", 0), // event OnRoutes
			/* 342 */ imports.NewTable("TChromiumCore_OnRouteStateChanged", 0), // event OnRouteStateChanged
			/* 343 */ imports.NewTable("TChromiumCore_OnRouteMessageReceived", 0), // event OnRouteMessageReceived
			/* 344 */ imports.NewTable("TChromiumCore_OnGetAudioParameters", 0), // event OnGetAudioParameters
			/* 345 */ imports.NewTable("TChromiumCore_OnAudioStreamStarted", 0), // event OnAudioStreamStarted
			/* 346 */ imports.NewTable("TChromiumCore_OnAudioStreamPacket", 0), // event OnAudioStreamPacket
			/* 347 */ imports.NewTable("TChromiumCore_OnAudioStreamStopped", 0), // event OnAudioStreamStopped
			/* 348 */ imports.NewTable("TChromiumCore_OnAudioStreamError", 0), // event OnAudioStreamError
			/* 349 */ imports.NewTable("TChromiumCore_OnDevToolsMessage", 0), // event OnDevToolsMessage
			/* 350 */ imports.NewTable("TChromiumCore_OnDevToolsRawMessage", 0), // event OnDevToolsRawMessage
			/* 351 */ imports.NewTable("TChromiumCore_OnDevToolsMethodResult", 0), // event OnDevToolsMethodResult
			/* 352 */ imports.NewTable("TChromiumCore_OnDevToolsMethodRawResult", 0), // event OnDevToolsMethodRawResult
			/* 353 */ imports.NewTable("TChromiumCore_OnDevToolsEvent", 0), // event OnDevToolsEvent
			/* 354 */ imports.NewTable("TChromiumCore_OnDevToolsRawEvent", 0), // event OnDevToolsRawEvent
			/* 355 */ imports.NewTable("TChromiumCore_OnDevToolsAgentAttached", 0), // event OnDevToolsAgentAttached
			/* 356 */ imports.NewTable("TChromiumCore_OnDevToolsAgentDetached", 0), // event OnDevToolsAgentDetached
			/* 357 */ imports.NewTable("TChromiumCore_OnExtensionLoadFailed", 0), // event OnExtensionLoadFailed
			/* 358 */ imports.NewTable("TChromiumCore_OnExtensionLoaded", 0), // event OnExtensionLoaded
			/* 359 */ imports.NewTable("TChromiumCore_OnExtensionUnloaded", 0), // event OnExtensionUnloaded
			/* 360 */ imports.NewTable("TChromiumCore_OnExtensionBeforeBackgroundBrowser", 0), // event OnExtensionBeforeBackgroundBrowser
			/* 361 */ imports.NewTable("TChromiumCore_OnExtensionBeforeBrowser", 0), // event OnExtensionBeforeBrowser
			/* 362 */ imports.NewTable("TChromiumCore_OnExtensionGetActiveBrowser", 0), // event OnExtensionGetActiveBrowser
			/* 363 */ imports.NewTable("TChromiumCore_OnExtensionCanAccessBrowser", 0), // event OnExtensionCanAccessBrowser
			/* 364 */ imports.NewTable("TChromiumCore_OnExtensionGetExtensionResource", 0), // event OnExtensionGetExtensionResource
			/* 365 */ imports.NewTable("TChromiumCore_OnPrintStart", 0), // event OnPrintStart
			/* 366 */ imports.NewTable("TChromiumCore_OnPrintSettings", 0), // event OnPrintSettings
			/* 367 */ imports.NewTable("TChromiumCore_OnPrintDialog", 0), // event OnPrintDialog
			/* 368 */ imports.NewTable("TChromiumCore_OnPrintJob", 0), // event OnPrintJob
			/* 369 */ imports.NewTable("TChromiumCore_OnPrintReset", 0), // event OnPrintReset
			/* 370 */ imports.NewTable("TChromiumCore_OnGetPDFPaperSize", 0), // event OnGetPDFPaperSize
			/* 371 */ imports.NewTable("TChromiumCore_OnFrameCreated", 0), // event OnFrameCreated
			/* 372 */ imports.NewTable("TChromiumCore_OnFrameAttached", 0), // event OnFrameAttached
			/* 373 */ imports.NewTable("TChromiumCore_OnFrameDetached", 0), // event OnFrameDetached
			/* 374 */ imports.NewTable("TChromiumCore_OnMainFrameChanged", 0), // event OnMainFrameChanged
			/* 375 */ imports.NewTable("TChromiumCore_OnChromeCommand", 0), // event OnChromeCommand
			/* 376 */ imports.NewTable("TChromiumCore_OnIsChromeAppMenuItemVisible", 0), // event OnIsChromeAppMenuItemVisible
			/* 377 */ imports.NewTable("TChromiumCore_OnIsChromeAppMenuItemEnabled", 0), // event OnIsChromeAppMenuItemEnabled
			/* 378 */ imports.NewTable("TChromiumCore_OnIsChromePageActionIconVisible", 0), // event OnIsChromePageActionIconVisible
			/* 379 */ imports.NewTable("TChromiumCore_OnIsChromeToolbarButtonVisible", 0), // event OnIsChromeToolbarButtonVisible
			/* 380 */ imports.NewTable("TChromiumCore_OnRequestMediaAccessPermission", 0), // event OnRequestMediaAccessPermission
			/* 381 */ imports.NewTable("TChromiumCore_OnShowPermissionPrompt", 0), // event OnShowPermissionPrompt
			/* 382 */ imports.NewTable("TChromiumCore_OnDismissPermissionPrompt", 0), // event OnDismissPermissionPrompt
		}
	})
	return chromiumCoreImport
}
