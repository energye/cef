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

	cefTypes "github.com/energye/cef/127/types"
)

// ICefBrowserHost Parent: ICefBaseRefCounted
type ICefBrowserHost interface {
	ICefBaseRefCounted
	// GetBrowser
	//  Returns the hosted browser object.
	GetBrowser() ICefBrowser // function
	// TryCloseBrowser
	//  Helper for closing a browser. Call this function from the top-level window
	//  close handler (if any). Internally this calls CloseBrowser(false (0)) if
	//  the close has not yet been initiated. This function returns false (0)
	//  while the close is pending and true (1) after the close has completed. See
	//  CloseBrowser() and ICefLifeSpanHandler.DoClose() documentation for
	//  additional usage information. This function must be called on the browser
	//  process UI thread.
	TryCloseBrowser() bool // function
	// GetWindowHandle
	//  Retrieve the window handle (if any) for this browser. If this browser is
	//  wrapped in a ICefBrowserView this function should be called on the
	//  browser process UI thread and it will return the handle for the top-level
	//  native window.
	GetWindowHandle() cefTypes.TCefWindowHandle // function
	// GetOpenerWindowHandle
	//  Retrieve the window handle (if any) of the browser that opened this
	//  browser. Will return NULL for non-popup browsers or if this browser is
	//  wrapped in a ICefBrowserView. This function can be used in combination
	//  with custom handling of modal windows.
	GetOpenerWindowHandle() cefTypes.TCefWindowHandle // function
	// HasView
	//  Returns true (1) if this browser is wrapped in a ICefBrowserView.
	HasView() bool // function
	// GetClient
	//  Returns the client for this browser.
	GetClient() IEngClient // function
	// GetRequestContext
	//  Returns the request context for this browser.
	GetRequestContext() ICefRequestContext // function
	// CanZoom
	//  Returns true (1) if this browser can execute the specified zoom command.
	//  This function can only be called on the UI thread.
	CanZoom(command cefTypes.TCefZoomCommand) bool // function
	// GetDefaultZoomLevel
	//  Get the default zoom level. This value will be 0.0 by default but can be
	//  configured with the Chrome runtime. This function can only be called on
	//  the UI thread.
	GetDefaultZoomLevel() float64 // function
	// GetZoomLevel
	//  Get the current zoom level. This function can only be called on the UI
	//  thread.
	GetZoomLevel() float64 // function
	// HasDevTools
	//  Returns true (1) if this browser currently has an associated DevTools
	//  browser. Must be called on the browser process UI thread.
	HasDevTools() bool // function
	// SendDevToolsMessage
	//  Send a function call message over the DevTools protocol. |message| must be
	//  a UTF8-encoded JSON dictionary that contains "id" (int), "function"
	//  (string) and "params" (dictionary, optional) values. See the DevTools
	//  protocol documentation at https://chromedevtools.github.io/devtools-
	//  protocol/ for details of supported functions and the expected "params"
	//  dictionary contents. |message| will be copied if necessary. This function
	//  will return true (1) if called on the UI thread and the message was
	//  successfully submitted for validation, otherwise false (0). Validation
	//  will be applied asynchronously and any messages that fail due to
	//  formatting errors or missing parameters may be discarded without
	//  notification. Prefer ExecuteDevToolsMethod if a more structured approach
	//  to message formatting is desired.
	//
	//  Every valid function call will result in an asynchronous function result
	//  or error message that references the sent message "id". Event messages are
	//  received while notifications are enabled (for example, between function
	//  calls for "Page.enable" and "Page.disable"). All received messages will be
	//  delivered to the observer(s) registered with AddDevToolsMessageObserver.
	//  See ICefDevToolsMessageObserver.OnDevToolsMessage documentation for
	//  details of received message contents.
	//
	//  Usage of the SendDevToolsMessage, ExecuteDevToolsMethod and
	//  AddDevToolsMessageObserver functions does not require an active DevTools
	//  front-end or remote-debugging session. Other active DevTools sessions will
	//  continue to function independently. However, any modification of global
	//  browser state by one session may not be reflected in the UI of other
	//  sessions.
	//
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
	// IsWindowRenderingDisabled
	//  Returns true (1) if window rendering is disabled.
	IsWindowRenderingDisabled() bool // function
	// GetWindowlessFrameRate
	//  Returns the maximum rate in frames per second (fps) that
	//  ICefRenderHandler.OnPaint will be called for a windowless browser. The
	//  actual fps may be lower if the browser cannot generate frames at the
	//  requested rate. The minimum value is 1 and the maximum value is 60
	//  (default 30). This function can only be called on the UI thread.
	GetWindowlessFrameRate() int32 // function
	// GetVisibleNavigationEntry
	//  Returns the current visible navigation entry for this browser. This
	//  function can only be called on the UI thread.
	GetVisibleNavigationEntry() ICefNavigationEntry // function
	// GetExtension
	//  Returns the extension hosted in this browser or NULL if no extension is
	//  hosted. See ICefRequestContext.LoadExtension for details.
	//  WARNING: This API is deprecated and will be removed in ~M127.
	GetExtension() ICefExtension // function
	// IsBackgroundHost
	//  Returns true (1) if this browser is hosting an extension background
	//  script. Background hosts do not have a window and are not displayable. See
	//  ICefRequestContext.LoadExtension for details.
	//  WARNING: This API is deprecated and will be removed in ~M127.
	IsBackgroundHost() bool // function
	// IsAudioMuted
	//  Returns true (1) if the browser's audio is muted. This function can only
	//  be called on the UI thread.
	IsAudioMuted() bool // function
	// IsFullscreen
	//  Returns true (1) if the renderer is currently in browser fullscreen. This
	//  differs from window fullscreen in that browser fullscreen is entered using
	//  the JavaScript Fullscreen API and modifies CSS attributes such as the
	//  ::backdrop pseudo-element and :fullscreen pseudo-structure. This function
	//  can only be called on the UI thread.
	IsFullscreen() bool // function
	// CanExecuteChromeCommand
	//  Returns true (1) if a Chrome command is supported and enabled. Values for
	//  |command_id| can be found in the cef_command_ids.h file. This function can
	//  only be called on the UI thread. Only used with the Chrome runtime.
	//  <see cref="uCEFConstants">See the IDC_* constants in uCEFConstants.pas for all the |command_id| values.</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/app/chrome_command_ids.h">The command_id values are also available in chrome/app/chrome_command_ids.h</see>
	CanExecuteChromeCommand(commandId int32) bool // function
	// IsRenderProcessUnresponsive
	//  Returns true (1) if the render process associated with this browser is
	//  currently unresponsive as indicated by a lack of input event processing
	//  for at least 15 seconds. To receive associated state change notifications
	//  and optionally handle an unresponsive render process implement
	//  ICefRequestHandler.OnRenderProcessUnresponsive.
	//  This function can only be called on the CEF UI thread.
	IsRenderProcessUnresponsive() bool // function
	// GetRuntimeStyle
	//  Returns the runtime style for this browser (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	//  This function can only be called on the CEF UI thread.
	GetRuntimeStyle() cefTypes.TCefRuntimeStyle // function
	// CloseBrowser
	//  Request that the browser close. The JavaScript 'onbeforeunload' event will
	//  be fired. If |force_close| is false (0) the event handler, if any, will be
	//  allowed to prompt the user and the user can optionally cancel the close.
	//  If |force_close| is true (1) the prompt will not be displayed and the
	//  close will proceed. Results in a call to
	//  ICefLifeSpanHandler.DoClose() if the event handler allows the close
	//  or if |force_close| is true (1). See ICefLifeSpanHandler.DoClose()
	//  documentation for additional usage information.
	CloseBrowser(forceClose bool) // procedure
	// SetFocus
	//  Set whether the browser is focused.
	SetFocus(focus bool) // procedure
	// Zoom
	//  Execute a zoom command in this browser. If called on the UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the UI thread.
	Zoom(command cefTypes.TCefZoomCommand) // procedure
	// SetZoomLevel
	//  Change the zoom level to the specified value. Specify 0.0 to reset the
	//  zoom level to the default. If called on the UI thread the change will be
	//  applied immediately. Otherwise, the change will be applied asynchronously
	//  on the UI thread.
	SetZoomLevel(zoomLevel float64) // procedure
	// RunFileDialog
	//  Call to run a file chooser dialog. Only a single file chooser dialog may
	//  be pending at any given time. |mode| represents the type of dialog to
	//  display. |title| to the title to be used for the dialog and may be NULL to
	//  show the default title ("Open" or "Save" depending on the mode).
	//  |default_file_path| is the path with optional directory and/or file name
	//  component that will be initially selected in the dialog. |accept_filters|
	//  are used to restrict the selectable file types and may any combination of
	//  (a) valid lower-cased MIME types (e.g. "text/*" or "image/*"), (b)
	//  individual file extensions (e.g. ".txt" or ".png"), or (c) combined
	//  description and file extension delimited using "|" and ";" (e.g. "Image
	//  Types|.png;.gif;.jpg"). |callback| will be executed after the dialog is
	//  dismissed or immediately if another dialog is already pending. The dialog
	//  will be initiated asynchronously on the UI thread.
	RunFileDialog(mode cefTypes.TCefFileDialogMode, title string, defaultFilePath string, acceptFilters lcl.IStrings, callback IEngRunFileDialogCallback) // procedure
	// StartDownload
	//  Download the file at |url| using ICefDownloadHandler.
	StartDownload(url string) // procedure
	// DownloadImage
	//  Download |image_url| and execute |callback| on completion with the images
	//  received from the renderer. If |is_favicon| is true (1) then cookies are
	//  not sent and not accepted during download. Images with density independent
	//  pixel (DIP) sizes larger than |max_image_size| are filtered out from the
	//  image results. Versions of the image at different scale factors may be
	//  downloaded up to the maximum scale factor supported by the system. If
	//  there are no image results <= |max_image_size| then the smallest image is
	//  resized to |max_image_size| and is the only result. A |max_image_size| of
	//  0 means unlimited. If |bypass_cache| is true (1) then |image_url| is
	//  requested from the server even if it is present in the browser cache.
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback IEngDownloadImageCallback) // procedure
	// Print
	//  Print the current browser contents.
	Print() // procedure
	// PrintToPdf
	//  Print the current browser contents to the PDF file specified by |path| and
	//  execute |callback| on completion. The caller is responsible for deleting
	//  |path| when done. For PDF printing to work on Linux you must implement the
	//  ICefPrintHandler.GetPdfPaperSize function.
	PrintToPdf(path string, settings TCefPdfPrintSettings, callback IEngPdfPrintCallback) // procedure
	// Find
	//  Search for |searchText|. |forward| indicates whether to search forward or
	//  backward within the page. |matchCase| indicates whether the search should
	//  be case-sensitive. |findNext| indicates whether this is the first request
	//  or a follow-up. The search will be restarted if |searchText| or
	//  |matchCase| change. The search will be stopped if |searchText| is NULL.
	//  The ICefFindHandler instance, if any, returned via
	//  ICefClient.GetFindHandler will be called to report find results.
	Find(searchText string, forward bool, matchCase bool, findNext bool) // procedure
	// StopFinding
	//  Cancel all searches that are currently going on.
	StopFinding(clearSelection bool) // procedure
	// ShowDevTools
	//  Open developer tools (DevTools) in its own browser. The DevTools browser
	//  will remain associated with this browser. If the DevTools browser is
	//  already open then it will be focused, in which case the |windowInfo|,
	//  |client| and |settings| parameters will be ignored. If
	//  |inspectElementAt| is non-NULL then the element at the specified (x,y)
	//  location will be inspected. The |windowInfo| parameter will be ignored if
	//  this browser is wrapped in a ICefBrowserView.
	ShowDevTools(windowInfo TCefWindowInfo, client IEngClient, settings TCefBrowserSettings, inspectElementAt TCefPoint) // procedure
	// CloseDevTools
	//  Explicitly close the associated DevTools browser, if any.
	CloseDevTools() // procedure
	// GetNavigationEntries
	//  Retrieve a snapshot of current navigation entries as values sent to the
	//  specified visitor. If |current_only| is true (1) only the current
	//  navigation entry will be sent, otherwise all navigation entries will be
	//  sent.
	GetNavigationEntries(visitor IEngNavigationEntryVisitor, currentOnly bool) // procedure
	// ReplaceMisspelling
	//  If a misspelled word is currently selected in an editable node calling
	//  this function will replace it with the specified |word|.
	ReplaceMisspelling(word string) // procedure
	// AddWordToDictionary
	//  Add the specified |word| to the spelling dictionary.
	AddWordToDictionary(word string) // procedure
	// WasResized
	//  Notify the browser that the widget has been resized. The browser will
	//  first call ICefRenderHandler.GetViewRect to get the new size and then
	//  call ICefRenderHandler.OnPaint asynchronously with the updated
	//  regions. This function is only used when window rendering is disabled.
	WasResized() // procedure
	// NotifyScreenInfoChanged
	//  Send a notification to the browser that the screen info has changed. The
	//  browser will then call ICefRenderHandler.GetScreenInfo to update the
	//  screen information with the new values. This simulates moving the webview
	//  window from one display to another, or changing the properties of the
	//  current display. This function is only used when window rendering is
	//  disabled.
	NotifyScreenInfoChanged() // procedure
	// WasHidden
	//  Notify the browser that it has been hidden or shown. Layouting and
	//  ICefRenderHandler.OnPaint notification will stop when the browser is
	//  hidden. This function is only used when window rendering is disabled.
	WasHidden(hidden bool) // procedure
	// Invalidate
	//  Invalidate the view. The browser will call ICefRenderHandler.OnPaint
	//  asynchronously. This function is only used when window rendering is
	//  disabled.
	Invalidate(type_ cefTypes.TCefPaintElementType) // procedure
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
	// NotifyMoveOrResizeStarted
	//  Notify the browser that the window hosting it is about to be moved or
	//  resized. This function is only used on Windows and Linux.
	NotifyMoveOrResizeStarted() // procedure
	// SetWindowlessFrameRate
	//  Set the maximum rate in frames per second (fps) that
	//  ICefRenderHandler.OnPaint will be called for a windowless browser.
	//  The actual fps may be lower if the browser cannot generate frames at the
	//  requested rate. The minimum value is 1 and the maximum value is 60
	//  (default 30). Can also be set at browser creation via
	//  TCefBrowserSettings.windowless_frame_rate.
	SetWindowlessFrameRate(frameRate int32) // procedure
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
	//
	//  This function may be called multiple times as the composition changes.
	//  When the client is done making changes the composition should either be
	//  canceled or completed. To cancel the composition call
	//  ImeCancelComposition. To complete the composition call either
	//  ImeCommitText or ImeFinishComposingText. Completion is usually signaled
	//  when:
	//
	//  1. The client receives a WM_IME_COMPOSITION message with a GCS_RESULTSTR
	//  flag (on Windows), or;
	//  2. The client receives a "commit" signal of GtkIMContext (on Linux), or;
	//  3. insertText of NSTextInput is called (on Mac).
	//
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
	// SetAccessibilityState
	//  Set accessibility state for all frames. |accessibility_state| may be
	//  default, enabled or disabled. If |accessibility_state| is STATE_DEFAULT
	//  then accessibility will be disabled by default and the state may be
	//  further controlled with the "force-renderer-accessibility" and "disable-
	//  renderer-accessibility" command-line switches. If |accessibility_state| is
	//  STATE_ENABLED then accessibility will be enabled. If |accessibility_state|
	//  is STATE_DISABLED then accessibility will be completely disabled.
	//
	//  For windowed browsers accessibility will be enabled in Complete mode
	//  (which corresponds to kAccessibilityModeComplete in Chromium). In this
	//  mode all platform accessibility objects will be created and managed by
	//  Chromium's internal implementation. The client needs only to detect the
	//  screen reader and call this function appropriately. For example, on macOS
	//  the client can handle the @"AXEnhancedUserStructure" accessibility
	//  attribute to detect VoiceOver state changes and on Windows the client can
	//  handle WM_GETOBJECT with OBJID_CLIENT to detect accessibility readers.
	//
	//  For windowless browsers accessibility will be enabled in TreeOnly mode
	//  (which corresponds to kAccessibilityModeWebContentsOnly in Chromium). In
	//  this mode renderer accessibility is enabled, the full tree is computed,
	//  and events are passed to CefAccessibiltyHandler, but platform
	//  accessibility objects are not created. The client may implement platform
	//  accessibility objects using CefAccessibiltyHandler callbacks if desired.
	SetAccessibilityState(accessibilityState cefTypes.TCefState) // procedure
	// SetAutoResizeEnabled
	//  Enable notifications of auto resize via
	//  ICefDisplayHandler.OnAutoResize. Notifications are disabled by
	//  default. |min_size| and |max_size| define the range of allowed sizes.
	SetAutoResizeEnabled(enabled bool, minSize TCefSize, maxSize TCefSize) // procedure
	// SetAudioMuted
	//  Set whether the browser's audio is muted.
	SetAudioMuted(mute bool) // procedure
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
}

// ICefBrowserHostRef Parent: ICefBrowserHost ICefBaseRefCountedRef
type ICefBrowserHostRef interface {
	ICefBrowserHost
	ICefBaseRefCountedRef
	AsIntfBrowserHost() uintptr
}

type TCefBrowserHostRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefBrowserHostRef) GetBrowser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCefBrowserHostRef) GetOpenerWindowHandle() cefTypes.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(4, m.Instance())
	return cefTypes.TCefWindowHandle(r)
}

func (m *TCefBrowserHostRef) HasView() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetClient() (result IEngClient) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsEngClient(resultPtr)
	return
}

func (m *TCefBrowserHostRef) GetRequestContext() (result ICefRequestContext) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRequestContextRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) CanZoom(command cefTypes.TCefZoomCommand) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(8, m.Instance(), uintptr(command))
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetDefaultZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefBrowserHostRef) GetZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCefBrowserHostRef) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) SendDevToolsMessage(message string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(12, m.Instance(), api.PasStr(message))
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) ExecuteDevToolsMethod(messageId int32, method string, params ICefDictionaryValue) int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(13, m.Instance(), uintptr(messageId), api.PasStr(method), base.GetObjectUintptr(params))
	return int32(r)
}

func (m *TCefBrowserHostRef) AddDevToolsMessageObserver(observer IEngDevToolsMessageObserver) (result ICefRegistration) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(observer), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefRegistrationRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) IsWindowRenderingDisabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetWindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(16, m.Instance())
	return int32(r)
}

func (m *TCefBrowserHostRef) GetVisibleNavigationEntry() (result ICefNavigationEntry) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefNavigationEntryRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) GetExtension() (result ICefExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefExtensionRef(resultPtr)
	return
}

func (m *TCefBrowserHostRef) IsBackgroundHost() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(19, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) IsAudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) CanExecuteChromeCommand(commandId int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(22, m.Instance(), uintptr(commandId))
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) IsRenderProcessUnresponsive() bool {
	if !m.IsValid() {
		return false
	}
	r := cefBrowserHostRefAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TCefBrowserHostRef) GetRuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := cefBrowserHostRefAPI().SysCallN(24, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TCefBrowserHostRef) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(26, m.Instance(), api.PasBool(forceClose))
}

func (m *TCefBrowserHostRef) SetFocus(focus bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(27, m.Instance(), api.PasBool(focus))
}

func (m *TCefBrowserHostRef) Zoom(command cefTypes.TCefZoomCommand) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(28, m.Instance(), uintptr(command))
}

func (m *TCefBrowserHostRef) SetZoomLevel(zoomLevel float64) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&zoomLevel)))
}

func (m *TCefBrowserHostRef) RunFileDialog(mode cefTypes.TCefFileDialogMode, title string, defaultFilePath string, acceptFilters lcl.IStrings, callback IEngRunFileDialogCallback) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(30, m.Instance(), uintptr(mode), api.PasStr(title), api.PasStr(defaultFilePath), base.GetObjectUintptr(acceptFilters), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) StartDownload(url string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(31, m.Instance(), api.PasStr(url))
}

func (m *TCefBrowserHostRef) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback IEngDownloadImageCallback) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(32, m.Instance(), api.PasStr(imageUrl), api.PasBool(isFavicon), uintptr(maxImageSize), api.PasBool(bypassCache), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) Print() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(33, m.Instance())
}

func (m *TCefBrowserHostRef) PrintToPdf(path string, settings TCefPdfPrintSettings, callback IEngPdfPrintCallback) {
	if !m.IsValid() {
		return
	}
	settingsPtr := settings.ToPas()
	cefBrowserHostRefAPI().SysCallN(34, m.Instance(), api.PasStr(path), uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(callback))
}

func (m *TCefBrowserHostRef) Find(searchText string, forward bool, matchCase bool, findNext bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(35, m.Instance(), api.PasStr(searchText), api.PasBool(forward), api.PasBool(matchCase), api.PasBool(findNext))
}

func (m *TCefBrowserHostRef) StopFinding(clearSelection bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(36, m.Instance(), api.PasBool(clearSelection))
}

func (m *TCefBrowserHostRef) ShowDevTools(windowInfo TCefWindowInfo, client IEngClient, settings TCefBrowserSettings, inspectElementAt TCefPoint) {
	if !m.IsValid() {
		return
	}
	windowInfoPtr := windowInfo.ToPas()
	settingsPtr := settings.ToPas()
	cefBrowserHostRefAPI().SysCallN(37, m.Instance(), uintptr(base.UnsafePointer(windowInfoPtr)), base.GetObjectUintptr(client), uintptr(base.UnsafePointer(settingsPtr)), uintptr(base.UnsafePointer(&inspectElementAt)))
}

func (m *TCefBrowserHostRef) CloseDevTools() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(38, m.Instance())
}

func (m *TCefBrowserHostRef) GetNavigationEntries(visitor IEngNavigationEntryVisitor, currentOnly bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(visitor), api.PasBool(currentOnly))
}

func (m *TCefBrowserHostRef) ReplaceMisspelling(word string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(40, m.Instance(), api.PasStr(word))
}

func (m *TCefBrowserHostRef) AddWordToDictionary(word string) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(41, m.Instance(), api.PasStr(word))
}

func (m *TCefBrowserHostRef) WasResized() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(42, m.Instance())
}

func (m *TCefBrowserHostRef) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(43, m.Instance())
}

func (m *TCefBrowserHostRef) WasHidden(hidden bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(44, m.Instance(), api.PasBool(hidden))
}

func (m *TCefBrowserHostRef) Invalidate(type_ cefTypes.TCefPaintElementType) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(45, m.Instance(), uintptr(type_))
}

func (m *TCefBrowserHostRef) SendExternalBeginFrame() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(46, m.Instance())
}

func (m *TCefBrowserHostRef) SendKeyEvent(event TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(47, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) SendMouseClickEvent(event TCefMouseEvent, type_ cefTypes.TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(48, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(type_), api.PasBool(mouseUp), uintptr(clickCount))
}

func (m *TCefBrowserHostRef) SendMouseMoveEvent(event TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(49, m.Instance(), uintptr(base.UnsafePointer(&event)), api.PasBool(mouseLeave))
}

func (m *TCefBrowserHostRef) SendMouseWheelEvent(event TCefMouseEvent, deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(50, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCefBrowserHostRef) SendTouchEvent(event TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(51, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(52, m.Instance())
}

func (m *TCefBrowserHostRef) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(53, m.Instance())
}

func (m *TCefBrowserHostRef) SetWindowlessFrameRate(frameRate int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(54, m.Instance(), uintptr(frameRate))
}

func (m *TCefBrowserHostRef) IMESetComposition(text string, underlines ICefCompositionUnderlineArray, replacementRange TCefRange, selectionRange TCefRange) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(55, m.Instance(), api.PasStr(text), underlines.Instance(), uintptr(int32(underlines.Count())), uintptr(base.UnsafePointer(&replacementRange)), uintptr(base.UnsafePointer(&selectionRange)))
}

func (m *TCefBrowserHostRef) IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(56, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&replacementRange)), uintptr(relativeCursorPos))
}

func (m *TCefBrowserHostRef) IMEFinishComposingText(keepSelection bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(57, m.Instance(), api.PasBool(keepSelection))
}

func (m *TCefBrowserHostRef) IMECancelComposition() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(58, m.Instance())
}

func (m *TCefBrowserHostRef) DragTargetDragEnter(dragData ICefDragData, event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(59, m.Instance(), base.GetObjectUintptr(dragData), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TCefBrowserHostRef) DragTargetDragOver(event TCefMouseEvent, allowedOps cefTypes.TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(60, m.Instance(), uintptr(base.UnsafePointer(&event)), uintptr(allowedOps))
}

func (m *TCefBrowserHostRef) DragTargetDragLeave() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(61, m.Instance())
}

func (m *TCefBrowserHostRef) DragTargetDrop(event TCefMouseEvent) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(62, m.Instance(), uintptr(base.UnsafePointer(&event)))
}

func (m *TCefBrowserHostRef) DragSourceEndedAt(X int32, Y int32, op cefTypes.TCefDragOperation) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(63, m.Instance(), uintptr(X), uintptr(Y), uintptr(op))
}

func (m *TCefBrowserHostRef) DragSourceSystemDragEnded() {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(64, m.Instance())
}

func (m *TCefBrowserHostRef) SetAccessibilityState(accessibilityState cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(65, m.Instance(), uintptr(accessibilityState))
}

func (m *TCefBrowserHostRef) SetAutoResizeEnabled(enabled bool, minSize TCefSize, maxSize TCefSize) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(66, m.Instance(), api.PasBool(enabled), uintptr(base.UnsafePointer(&minSize)), uintptr(base.UnsafePointer(&maxSize)))
}

func (m *TCefBrowserHostRef) SetAudioMuted(mute bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(67, m.Instance(), api.PasBool(mute))
}

func (m *TCefBrowserHostRef) ExitFullscreen(willCauseResize bool) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(68, m.Instance(), api.PasBool(willCauseResize))
}

func (m *TCefBrowserHostRef) ExecuteChromeCommand(commandId int32, disposition cefTypes.TCefWindowOpenDisposition) {
	if !m.IsValid() {
		return
	}
	cefBrowserHostRefAPI().SysCallN(69, m.Instance(), uintptr(commandId), uintptr(disposition))
}

func (m *TCefBrowserHostRef) AsIntfBrowserHost() uintptr {
	return m.GetIntfPointer(0)
}

// BrowserHostRef  is static instance
var BrowserHostRef _BrowserHostRefClass

// _BrowserHostRefClass is class type defined by TCefBrowserHostRef
type _BrowserHostRefClass uintptr

func (_BrowserHostRefClass) UnWrap(data uintptr) (result ICefBrowserHost) {
	var resultPtr uintptr
	cefBrowserHostRefAPI().SysCallN(25, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserHostRef(resultPtr)
	return
}

// NewBrowserHostRef class constructor
func NewBrowserHostRef(data uintptr) ICefBrowserHostRef {
	var browserHostPtr uintptr // ICefBrowserHost
	r := cefBrowserHostRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&browserHostPtr)))
	ret := AsCefBrowserHostRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserHostPtr)
	}
	return ret
}

var (
	cefBrowserHostRefOnce   base.Once
	cefBrowserHostRefImport *imports.Imports = nil
)

func cefBrowserHostRefAPI() *imports.Imports {
	cefBrowserHostRefOnce.Do(func() {
		cefBrowserHostRefImport = api.NewDefaultImports()
		cefBrowserHostRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefBrowserHostRef_Create", 0), // constructor NewBrowserHostRef
			/* 1 */ imports.NewTable("TCefBrowserHostRef_GetBrowser", 0), // function GetBrowser
			/* 2 */ imports.NewTable("TCefBrowserHostRef_TryCloseBrowser", 0), // function TryCloseBrowser
			/* 3 */ imports.NewTable("TCefBrowserHostRef_GetWindowHandle", 0), // function GetWindowHandle
			/* 4 */ imports.NewTable("TCefBrowserHostRef_GetOpenerWindowHandle", 0), // function GetOpenerWindowHandle
			/* 5 */ imports.NewTable("TCefBrowserHostRef_HasView", 0), // function HasView
			/* 6 */ imports.NewTable("TCefBrowserHostRef_GetClient", 0), // function GetClient
			/* 7 */ imports.NewTable("TCefBrowserHostRef_GetRequestContext", 0), // function GetRequestContext
			/* 8 */ imports.NewTable("TCefBrowserHostRef_CanZoom", 0), // function CanZoom
			/* 9 */ imports.NewTable("TCefBrowserHostRef_GetDefaultZoomLevel", 0), // function GetDefaultZoomLevel
			/* 10 */ imports.NewTable("TCefBrowserHostRef_GetZoomLevel", 0), // function GetZoomLevel
			/* 11 */ imports.NewTable("TCefBrowserHostRef_HasDevTools", 0), // function HasDevTools
			/* 12 */ imports.NewTable("TCefBrowserHostRef_SendDevToolsMessage", 0), // function SendDevToolsMessage
			/* 13 */ imports.NewTable("TCefBrowserHostRef_ExecuteDevToolsMethod", 0), // function ExecuteDevToolsMethod
			/* 14 */ imports.NewTable("TCefBrowserHostRef_AddDevToolsMessageObserver", 0), // function AddDevToolsMessageObserver
			/* 15 */ imports.NewTable("TCefBrowserHostRef_IsWindowRenderingDisabled", 0), // function IsWindowRenderingDisabled
			/* 16 */ imports.NewTable("TCefBrowserHostRef_GetWindowlessFrameRate", 0), // function GetWindowlessFrameRate
			/* 17 */ imports.NewTable("TCefBrowserHostRef_GetVisibleNavigationEntry", 0), // function GetVisibleNavigationEntry
			/* 18 */ imports.NewTable("TCefBrowserHostRef_GetExtension", 0), // function GetExtension
			/* 19 */ imports.NewTable("TCefBrowserHostRef_IsBackgroundHost", 0), // function IsBackgroundHost
			/* 20 */ imports.NewTable("TCefBrowserHostRef_IsAudioMuted", 0), // function IsAudioMuted
			/* 21 */ imports.NewTable("TCefBrowserHostRef_IsFullscreen", 0), // function IsFullscreen
			/* 22 */ imports.NewTable("TCefBrowserHostRef_CanExecuteChromeCommand", 0), // function CanExecuteChromeCommand
			/* 23 */ imports.NewTable("TCefBrowserHostRef_IsRenderProcessUnresponsive", 0), // function IsRenderProcessUnresponsive
			/* 24 */ imports.NewTable("TCefBrowserHostRef_GetRuntimeStyle", 0), // function GetRuntimeStyle
			/* 25 */ imports.NewTable("TCefBrowserHostRef_UnWrap", 0), // static function UnWrap
			/* 26 */ imports.NewTable("TCefBrowserHostRef_CloseBrowser", 0), // procedure CloseBrowser
			/* 27 */ imports.NewTable("TCefBrowserHostRef_SetFocus", 0), // procedure SetFocus
			/* 28 */ imports.NewTable("TCefBrowserHostRef_Zoom", 0), // procedure Zoom
			/* 29 */ imports.NewTable("TCefBrowserHostRef_SetZoomLevel", 0), // procedure SetZoomLevel
			/* 30 */ imports.NewTable("TCefBrowserHostRef_RunFileDialog", 0), // procedure RunFileDialog
			/* 31 */ imports.NewTable("TCefBrowserHostRef_StartDownload", 0), // procedure StartDownload
			/* 32 */ imports.NewTable("TCefBrowserHostRef_DownloadImage", 0), // procedure DownloadImage
			/* 33 */ imports.NewTable("TCefBrowserHostRef_Print", 0), // procedure Print
			/* 34 */ imports.NewTable("TCefBrowserHostRef_PrintToPdf", 0), // procedure PrintToPdf
			/* 35 */ imports.NewTable("TCefBrowserHostRef_Find", 0), // procedure Find
			/* 36 */ imports.NewTable("TCefBrowserHostRef_StopFinding", 0), // procedure StopFinding
			/* 37 */ imports.NewTable("TCefBrowserHostRef_ShowDevTools", 0), // procedure ShowDevTools
			/* 38 */ imports.NewTable("TCefBrowserHostRef_CloseDevTools", 0), // procedure CloseDevTools
			/* 39 */ imports.NewTable("TCefBrowserHostRef_GetNavigationEntries", 0), // procedure GetNavigationEntries
			/* 40 */ imports.NewTable("TCefBrowserHostRef_ReplaceMisspelling", 0), // procedure ReplaceMisspelling
			/* 41 */ imports.NewTable("TCefBrowserHostRef_AddWordToDictionary", 0), // procedure AddWordToDictionary
			/* 42 */ imports.NewTable("TCefBrowserHostRef_WasResized", 0), // procedure WasResized
			/* 43 */ imports.NewTable("TCefBrowserHostRef_NotifyScreenInfoChanged", 0), // procedure NotifyScreenInfoChanged
			/* 44 */ imports.NewTable("TCefBrowserHostRef_WasHidden", 0), // procedure WasHidden
			/* 45 */ imports.NewTable("TCefBrowserHostRef_Invalidate", 0), // procedure Invalidate
			/* 46 */ imports.NewTable("TCefBrowserHostRef_SendExternalBeginFrame", 0), // procedure SendExternalBeginFrame
			/* 47 */ imports.NewTable("TCefBrowserHostRef_SendKeyEvent", 0), // procedure SendKeyEvent
			/* 48 */ imports.NewTable("TCefBrowserHostRef_SendMouseClickEvent", 0), // procedure SendMouseClickEvent
			/* 49 */ imports.NewTable("TCefBrowserHostRef_SendMouseMoveEvent", 0), // procedure SendMouseMoveEvent
			/* 50 */ imports.NewTable("TCefBrowserHostRef_SendMouseWheelEvent", 0), // procedure SendMouseWheelEvent
			/* 51 */ imports.NewTable("TCefBrowserHostRef_SendTouchEvent", 0), // procedure SendTouchEvent
			/* 52 */ imports.NewTable("TCefBrowserHostRef_SendCaptureLostEvent", 0), // procedure SendCaptureLostEvent
			/* 53 */ imports.NewTable("TCefBrowserHostRef_NotifyMoveOrResizeStarted", 0), // procedure NotifyMoveOrResizeStarted
			/* 54 */ imports.NewTable("TCefBrowserHostRef_SetWindowlessFrameRate", 0), // procedure SetWindowlessFrameRate
			/* 55 */ imports.NewTable("TCefBrowserHostRef_IMESetComposition", 0), // procedure IMESetComposition
			/* 56 */ imports.NewTable("TCefBrowserHostRef_IMECommitText", 0), // procedure IMECommitText
			/* 57 */ imports.NewTable("TCefBrowserHostRef_IMEFinishComposingText", 0), // procedure IMEFinishComposingText
			/* 58 */ imports.NewTable("TCefBrowserHostRef_IMECancelComposition", 0), // procedure IMECancelComposition
			/* 59 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragEnter", 0), // procedure DragTargetDragEnter
			/* 60 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragOver", 0), // procedure DragTargetDragOver
			/* 61 */ imports.NewTable("TCefBrowserHostRef_DragTargetDragLeave", 0), // procedure DragTargetDragLeave
			/* 62 */ imports.NewTable("TCefBrowserHostRef_DragTargetDrop", 0), // procedure DragTargetDrop
			/* 63 */ imports.NewTable("TCefBrowserHostRef_DragSourceEndedAt", 0), // procedure DragSourceEndedAt
			/* 64 */ imports.NewTable("TCefBrowserHostRef_DragSourceSystemDragEnded", 0), // procedure DragSourceSystemDragEnded
			/* 65 */ imports.NewTable("TCefBrowserHostRef_SetAccessibilityState", 0), // procedure SetAccessibilityState
			/* 66 */ imports.NewTable("TCefBrowserHostRef_SetAutoResizeEnabled", 0), // procedure SetAutoResizeEnabled
			/* 67 */ imports.NewTable("TCefBrowserHostRef_SetAudioMuted", 0), // procedure SetAudioMuted
			/* 68 */ imports.NewTable("TCefBrowserHostRef_ExitFullscreen", 0), // procedure ExitFullscreen
			/* 69 */ imports.NewTable("TCefBrowserHostRef_ExecuteChromeCommand", 0), // procedure ExecuteChromeCommand
		}
	})
	return cefBrowserHostRefImport
}
