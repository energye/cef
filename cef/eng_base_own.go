//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefAccessibilityHandler Parent: ICefBaseRefCountedOwn
type ICefAccessibilityHandler interface {
	ICefBaseRefCountedOwn
}

// TCefAccessibilityHandler Parent: TCefBaseRefCountedOwn
type TCefAccessibilityHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDevToolsMessageObserver Parent: ICefBaseRefCountedOwn
type ICefDevToolsMessageObserver interface {
	ICefBaseRefCountedOwn
}

// TCefDevToolsMessageObserver Parent: TCefBaseRefCountedOwn
type TCefDevToolsMessageObserver struct {
	TCefBaseRefCountedOwn
}

// ICEFServerHandler Parent: ICefBaseRefCountedOwn
type ICEFServerHandler interface {
	ICefBaseRefCountedOwn
}

// TCEFServerHandler Parent: TCefBaseRefCountedOwn
type TCEFServerHandler struct {
	TCefBaseRefCountedOwn
}

// ICefAudioHandler Parent: ICefBaseRefCountedOwn
type ICefAudioHandler interface {
	ICefBaseRefCountedOwn
}

// TCefAudioHandler Parent: TCefBaseRefCountedOwn
type TCefAudioHandler struct {
	TCefBaseRefCountedOwn
}

// ICefBaseRefCountedOwn Parent: IObject
type ICefBaseRefCountedOwn interface {
	IObject
}

// TCefBaseRefCountedOwn Parent: TObject
type TCefBaseRefCountedOwn struct {
	TObject
}

// ICefBrowserViewDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle BrowserView events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_delegate_capi.h">CEF source file: /include/capi/views/cef_browser_view_delegate_capi.h (cef_browser_view_delegate_t)</a>
type ICefBrowserViewDelegate interface {
	ICefViewDelegate
}

// TCefBrowserViewDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle BrowserView events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_delegate_capi.h">CEF source file: /include/capi/views/cef_browser_view_delegate_capi.h (cef_browser_view_delegate_t)</a>
type TCefBrowserViewDelegate struct {
	TCefViewDelegate
}

// ICefButtonDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Button events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_delegate_capi.h">CEF source file: /include/capi/views/cef_button_delegate_capi.h (cef_button_delegate_t)</a>
type ICefButtonDelegate interface {
	ICefViewDelegate
}

// TCefButtonDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Button events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_delegate_capi.h">CEF source file: /include/capi/views/cef_button_delegate_capi.h (cef_button_delegate_t)</a>
type TCefButtonDelegate struct {
	TCefViewDelegate
}

// ICefClient Parent: ICefBaseRefCountedOwn
type ICefClient interface {
	ICefBaseRefCountedOwn
}

// TCefClient Parent: TCefBaseRefCountedOwn
type TCefClient struct {
	TCefBaseRefCountedOwn
}

// ICefCommandHandler Parent: ICefBaseRefCountedOwn
type ICefCommandHandler interface {
	ICefBaseRefCountedOwn
}

// TCefCommandHandler Parent: TCefBaseRefCountedOwn
type TCefCommandHandler struct {
	TCefBaseRefCountedOwn
}

// ICefCompletionCallback Parent: ICefBaseRefCountedOwn
type ICefCompletionCallback interface {
	ICefBaseRefCountedOwn
}

// TCefCompletionCallback Parent: TCefBaseRefCountedOwn
type TCefCompletionCallback struct {
	TCefBaseRefCountedOwn
}

// ICefContextMenuHandler Parent: ICefBaseRefCountedOwn
type ICefContextMenuHandler interface {
	ICefBaseRefCountedOwn
}

// TCefContextMenuHandler Parent: TCefBaseRefCountedOwn
type TCefContextMenuHandler struct {
	TCefBaseRefCountedOwn
}

// ICefCookieVisitor Parent: ICefBaseRefCountedOwn
type ICefCookieVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefCookieVisitor Parent: TCefBaseRefCountedOwn
type TCefCookieVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefDeleteCookiesCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefDeleteCookiesCallback interface {
	ICefBaseRefCountedOwn
}

// TCefDeleteCookiesCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefDeleteCookiesCallback struct {
	TCefBaseRefCountedOwn
}

// ICefDialogHandler Parent: ICefBaseRefCountedOwn
type ICefDialogHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDialogHandler Parent: TCefBaseRefCountedOwn
type TCefDialogHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDisplayHandler Parent: ICefBaseRefCountedOwn
//
//	Event handler related to browser display state.
//	The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h (cef_display_handler_t)</a>
type ICefDisplayHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDisplayHandler Parent: TCefBaseRefCountedOwn
//
//	Event handler related to browser display state.
//	The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h (cef_display_handler_t)</a>
type TCefDisplayHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDomVisitor Parent: ICefBaseRefCountedOwn
type ICefDomVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefDomVisitor Parent: TCefBaseRefCountedOwn
type TCefDomVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadHandler Parent: ICefBaseRefCountedOwn
type ICefDownloadHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDownloadHandler Parent: TCefBaseRefCountedOwn
type TCefDownloadHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadImageCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefDownloadImageCallback interface {
	ICefBaseRefCountedOwn
}

// TCefDownloadImageCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefDownloadImageCallback struct {
	TCefBaseRefCountedOwn
}

// ICefDragHandler Parent: ICefBaseRefCountedOwn
type ICefDragHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDragHandler Parent: TCefBaseRefCountedOwn
type TCefDragHandler struct {
	TCefBaseRefCountedOwn
}

// ICefExtensionHandler Parent: ICefBaseRefCountedOwn
type ICefExtensionHandler interface {
	ICefBaseRefCountedOwn
}

// TCefExtensionHandler Parent: TCefBaseRefCountedOwn
type TCefExtensionHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFindHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefFindHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFindHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefFindHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFocusHandler Parent: ICefBaseRefCountedOwn
type ICefFocusHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFocusHandler Parent: TCefBaseRefCountedOwn
type TCefFocusHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFrameHandler Parent: ICefBaseRefCountedOwn
type ICefFrameHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFrameHandler Parent: TCefBaseRefCountedOwn
type TCefFrameHandler struct {
	TCefBaseRefCountedOwn
}

// ICefJsDialogHandler Parent: ICefBaseRefCountedOwn
type ICefJsDialogHandler interface {
	ICefBaseRefCountedOwn
}

// TCefJsDialogHandler Parent: TCefBaseRefCountedOwn
type TCefJsDialogHandler struct {
	TCefBaseRefCountedOwn
}

// ICefKeyboardHandler Parent: ICefBaseRefCountedOwn
type ICefKeyboardHandler interface {
	ICefBaseRefCountedOwn
}

// TCefKeyboardHandler Parent: TCefBaseRefCountedOwn
type TCefKeyboardHandler struct {
	TCefBaseRefCountedOwn
}

// ICefLifeSpanHandler Parent: ICefBaseRefCountedOwn
type ICefLifeSpanHandler interface {
	ICefBaseRefCountedOwn
}

// TCefLifeSpanHandler Parent: TCefBaseRefCountedOwn
type TCefLifeSpanHandler struct {
	TCefBaseRefCountedOwn
}

// ICefLoadHandler Parent: ICefBaseRefCountedOwn
type ICefLoadHandler interface {
	ICefBaseRefCountedOwn
}

// TCefLoadHandler Parent: TCefBaseRefCountedOwn
type TCefLoadHandler struct {
	TCefBaseRefCountedOwn
}

// ICefMediaObserver Parent: ICefBaseRefCountedOwn
type ICefMediaObserver interface {
	ICefBaseRefCountedOwn
}

// TCefMediaObserver Parent: TCefBaseRefCountedOwn
type TCefMediaObserver struct {
	TCefBaseRefCountedOwn
}

// ICefMediaRouteCreateCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefMediaRouteCreateCallback interface {
	ICefBaseRefCountedOwn
}

// TCefMediaRouteCreateCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefMediaRouteCreateCallback struct {
	TCefBaseRefCountedOwn
}

// ICefMediaSinkDeviceInfoCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefMediaSinkDeviceInfoCallback interface {
	ICefBaseRefCountedOwn
}

// TCefMediaSinkDeviceInfoCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefMediaSinkDeviceInfoCallback struct {
	TCefBaseRefCountedOwn
}

// ICefMenuButtonDelegate Parent: ICefButtonDelegate
//
//	Implement this interface to handle MenuButton events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_delegate_t)</a>
type ICefMenuButtonDelegate interface {
	ICefButtonDelegate
}

// TCefMenuButtonDelegate Parent: TCefButtonDelegate
//
//	Implement this interface to handle MenuButton events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_delegate_t)</a>
type TCefMenuButtonDelegate struct {
	TCefButtonDelegate
}

// ICefMenuModelDelegate Parent: ICefBaseRefCountedOwn
type ICefMenuModelDelegate interface {
	ICefBaseRefCountedOwn
}

// TCefMenuModelDelegate Parent: TCefBaseRefCountedOwn
type TCefMenuModelDelegate struct {
	TCefBaseRefCountedOwn
}

// ICefNavigationEntryVisitor Parent: ICefBaseRefCountedOwn
type ICefNavigationEntryVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefNavigationEntryVisitor Parent: TCefBaseRefCountedOwn
type TCefNavigationEntryVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefPanelDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Panel events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_delegate_capi.h">CEF source file: /include/capi/views/cef_panel_delegate_capi.h (cef_panel_delegate_t)</a>
type ICefPanelDelegate interface {
	ICefViewDelegate
}

// TCefPanelDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Panel events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_delegate_capi.h">CEF source file: /include/capi/views/cef_panel_delegate_capi.h (cef_panel_delegate_t)</a>
type TCefPanelDelegate struct {
	TCefViewDelegate
}

// ICefPdfPrintCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefPdfPrintCallback interface {
	ICefBaseRefCountedOwn
}

// TCefPdfPrintCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefPdfPrintCallback struct {
	TCefBaseRefCountedOwn
}

// ICefPermissionHandler Parent: ICefBaseRefCountedOwn
type ICefPermissionHandler interface {
	ICefBaseRefCountedOwn
}

// TCefPermissionHandler Parent: TCefBaseRefCountedOwn
type TCefPermissionHandler struct {
	TCefBaseRefCountedOwn
}

// ICefPrintHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefPrintHandler interface {
	ICefBaseRefCountedOwn
}

// TCefPrintHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefPrintHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRenderHandler Parent: ICefBaseRefCountedOwn
type ICefRenderHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRenderHandler Parent: TCefBaseRefCountedOwn
type TCefRenderHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRenderProcessHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefRenderProcessHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRenderProcessHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefRenderProcessHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRequestHandler Parent: ICefBaseRefCountedOwn
type ICefRequestHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRequestHandler Parent: TCefBaseRefCountedOwn
type TCefRequestHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResolveCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefResolveCallback interface {
	ICefBaseRefCountedOwn
}

// TCefResolveCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefResolveCallback struct {
	TCefBaseRefCountedOwn
}

// ICefResourceHandler Parent: ICefBaseRefCountedOwn
type ICefResourceHandler interface {
	ICefBaseRefCountedOwn
}

// TCefResourceHandler Parent: TCefBaseRefCountedOwn
type TCefResourceHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResourceRequestHandler Parent: ICefBaseRefCountedOwn
type ICefResourceRequestHandler interface {
	ICefBaseRefCountedOwn
}

// TCefResourceRequestHandler Parent: TCefBaseRefCountedOwn
type TCefResourceRequestHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResponseFilter Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefResponseFilter interface {
	ICefBaseRefCountedOwn
}

// TCefResponseFilter Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefResponseFilter struct {
	TCefBaseRefCountedOwn
}

// ICefRunFileDialogCallback Parent: ICefBaseRefCountedOwn
type ICefRunFileDialogCallback interface {
	ICefBaseRefCountedOwn
}

// TCefRunFileDialogCallback Parent: TCefBaseRefCountedOwn
type TCefRunFileDialogCallback struct {
	TCefBaseRefCountedOwn
}

// ICefSchemeHandlerFactory Parent: ICefBaseRefCountedOwn
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type ICefSchemeHandlerFactory interface {
	ICefBaseRefCountedOwn
}

// TCefSchemeHandlerFactory Parent: TCefBaseRefCountedOwn
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type TCefSchemeHandlerFactory struct {
	TCefBaseRefCountedOwn
}

// ICefSetCookieCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
type ICefSetCookieCallback interface {
	ICefBaseRefCountedOwn
}

// TCefSetCookieCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
type TCefSetCookieCallback struct {
	TCefBaseRefCountedOwn
}

// ICefStringMultimapOwn Parent: ICefStringMultimap
type ICefStringMultimapOwn interface {
	ICefStringMultimap
}

// TCefStringMultimapOwn Parent: TCefStringMultimap
type TCefStringMultimapOwn struct {
	TCefStringMultimap
}

// ICefStringVisitor Parent: ICefBaseRefCountedOwn
type ICefStringVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefStringVisitor Parent: TCefBaseRefCountedOwn
type TCefStringVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefTask Parent: ICefBaseRefCountedOwn
type ICefTask interface {
	ICefBaseRefCountedOwn
}

// TCefTask Parent: TCefBaseRefCountedOwn
type TCefTask struct {
	TCefBaseRefCountedOwn
}

// ICefTextfieldDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Textfield events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_delegate_capi.h">CEF source file: /include/capi/views/cef_textfield_delegate_capi.h (cef_textfield_delegate_t)</a>
type ICefTextfieldDelegate interface {
	ICefViewDelegate
}

// TCefTextfieldDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Textfield events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_delegate_capi.h">CEF source file: /include/capi/views/cef_textfield_delegate_capi.h (cef_textfield_delegate_t)</a>
type TCefTextfieldDelegate struct {
	TCefViewDelegate
}

// ICefUrlRequestClient Parent: ICefBaseRefCountedOwn
type ICefUrlRequestClient interface {
	ICefBaseRefCountedOwn
}

// TCefUrlRequestClient Parent: TCefBaseRefCountedOwn
type TCefUrlRequestClient struct {
	TCefBaseRefCountedOwn
}

// ICefV8Accessor Parent: ICefBaseRefCountedOwn
type ICefV8Accessor interface {
	ICefBaseRefCountedOwn
}

// TCefV8Accessor Parent: TCefBaseRefCountedOwn
type TCefV8Accessor struct {
	TCefBaseRefCountedOwn
}

// ICefV8Interceptor Parent: ICefBaseRefCountedOwn
type ICefV8Interceptor interface {
	ICefBaseRefCountedOwn
}

// TCefV8Interceptor Parent: TCefBaseRefCountedOwn
type TCefV8Interceptor struct {
	TCefBaseRefCountedOwn
}

// ICefViewDelegate Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle view events. All size and position values
//	are in density independent pixels (DIP) unless otherwise indicated. The
//	functions of this interface will be called on the browser process UI thread
//	unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_delegate_capi.h">CEF source file: /include/capi/views/cef_view_delegate_capi.h (cef_view_delegate_t)</a>
type ICefViewDelegate interface {
	ICefBaseRefCountedOwn
}

// TCefViewDelegate Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle view events. All size and position values
//	are in density independent pixels (DIP) unless otherwise indicated. The
//	functions of this interface will be called on the browser process UI thread
//	unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_delegate_capi.h">CEF source file: /include/capi/views/cef_view_delegate_capi.h (cef_view_delegate_t)</a>
type TCefViewDelegate struct {
	TCefBaseRefCountedOwn
}

// ICefWindowDelegate Parent: ICefPanelDelegate
//
//	Implement this interface to handle window events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_window_delegate_capi.h">CEF source file: /include/capi/views/cef_window_delegate_capi.h (cef_window_delegate_t)</a>
type ICefWindowDelegate interface {
	ICefPanelDelegate
}

// TCefWindowDelegate Parent: TCefPanelDelegate
//
//	Implement this interface to handle window events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_window_delegate_capi.h">CEF source file: /include/capi/views/cef_window_delegate_capi.h (cef_window_delegate_t)</a>
type TCefWindowDelegate struct {
	TCefPanelDelegate
}

// ICefWriteHandler Parent: ICefBaseRefCountedOwn
type ICefWriteHandler interface {
	ICefBaseRefCountedOwn
}

// TCefWriteHandler Parent: TCefBaseRefCountedOwn
type TCefWriteHandler struct {
	TCefBaseRefCountedOwn
}

// ICefV8ArrayBufferReleaseCallback Parent: ICefBaseRefCountedOwn
type ICefV8ArrayBufferReleaseCallback interface {
	ICefBaseRefCountedOwn
}

// TCefV8ArrayBufferReleaseCallback Parent: TCefBaseRefCountedOwn
type TCefV8ArrayBufferReleaseCallback struct {
	TCefBaseRefCountedOwn
}

// ICefV8Handler Parent: ICefBaseRefCountedOwn
type ICefV8Handler interface {
	ICefBaseRefCountedOwn
}

// TCefV8Handler Parent: TCefBaseRefCountedOwn
type TCefV8Handler struct {
	TCefBaseRefCountedOwn
}
