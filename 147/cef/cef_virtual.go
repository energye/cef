//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefBrowserProcessHandler Parent: ICefBaseRefCounted
type ICefBrowserProcessHandler interface {
	ICefBaseRefCounted
}

type ICefBrowserProcessHandlerOwn interface {
	ICefBrowserProcessHandler
	ICefBaseRefCountedOwn
}
type TCefBrowserProcessHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefCompletionCallback Parent: ICefBaseRefCounted
type ICefCompletionCallback interface {
	ICefBaseRefCounted
}

type ICefCompletionCallbackOwn interface {
	ICefCompletionCallback
	ICefBaseRefCountedOwn
}
type TCefCompletionCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefContextMenuHandler Parent: ICefBaseRefCounted
type ICefContextMenuHandler interface {
	ICefBaseRefCounted
}

type ICefContextMenuHandlerOwn interface {
	ICefContextMenuHandler
	ICefBaseRefCountedOwn
}
type TCefContextMenuHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefCookieVisitor Parent: ICefBaseRefCounted
type ICefCookieVisitor interface {
	ICefBaseRefCounted
}

type ICefCookieVisitorOwn interface {
	ICefCookieVisitor
	ICefBaseRefCountedOwn
}
type TCefCookieVisitorOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDeleteCookiesCallback Parent: ICefBaseRefCounted
type ICefDeleteCookiesCallback interface {
	ICefBaseRefCounted
}

type ICefDeleteCookiesCallbackOwn interface {
	ICefDeleteCookiesCallback
	ICefBaseRefCountedOwn
}
type TCefDeleteCookiesCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDialogHandler Parent: ICefBaseRefCounted
type ICefDialogHandler interface {
	ICefBaseRefCounted
}

type ICefDialogHandlerOwn interface {
	ICefDialogHandler
	ICefBaseRefCountedOwn
}
type TCefDialogHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDisplayHandler Parent: ICefBaseRefCounted
type ICefDisplayHandler interface {
	ICefBaseRefCounted
}

type ICefDisplayHandlerOwn interface {
	ICefDisplayHandler
	ICefBaseRefCountedOwn
}
type TCefDisplayHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDomVisitor Parent: ICefBaseRefCounted
type ICefDomVisitor interface {
	ICefBaseRefCounted
}

type ICefDomVisitorOwn interface {
	ICefDomVisitor
	ICefBaseRefCountedOwn
}
type TCefDomVisitorOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadHandler Parent: ICefBaseRefCounted
type ICefDownloadHandler interface {
	ICefBaseRefCounted
}

type ICefDownloadHandlerOwn interface {
	ICefDownloadHandler
	ICefBaseRefCountedOwn
}
type TCefDownloadHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadImageCallback Parent: ICefBaseRefCounted
type ICefDownloadImageCallback interface {
	ICefBaseRefCounted
}

type ICefDownloadImageCallbackOwn interface {
	ICefDownloadImageCallback
	ICefBaseRefCountedOwn
}
type TCefDownloadImageCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDragHandler Parent: ICefBaseRefCounted
type ICefDragHandler interface {
	ICefBaseRefCounted
}

type ICefDragHandlerOwn interface {
	ICefDragHandler
	ICefBaseRefCountedOwn
}
type TCefDragHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefEndTracingCallback Parent: ICefBaseRefCounted
type ICefEndTracingCallback interface {
	ICefBaseRefCounted
}

type ICefEndTracingCallbackOwn interface {
	ICefEndTracingCallback
	ICefBaseRefCountedOwn
}
type TCefEndTracingCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefFindHandler Parent: ICefBaseRefCounted
type ICefFindHandler interface {
	ICefBaseRefCounted
}

type ICefFindHandlerOwn interface {
	ICefFindHandler
	ICefBaseRefCountedOwn
}
type TCefFindHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefFocusHandler Parent: ICefBaseRefCounted
type ICefFocusHandler interface {
	ICefBaseRefCounted
}

type ICefFocusHandlerOwn interface {
	ICefFocusHandler
	ICefBaseRefCountedOwn
}
type TCefFocusHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefJsDialogHandler Parent: ICefBaseRefCounted
type ICefJsDialogHandler interface {
	ICefBaseRefCounted
}

type ICefJsDialogHandlerOwn interface {
	ICefJsDialogHandler
	ICefBaseRefCountedOwn
}
type TCefJsDialogHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefKeyboardHandler Parent: ICefBaseRefCounted
type ICefKeyboardHandler interface {
	ICefBaseRefCounted
}

type ICefKeyboardHandlerOwn interface {
	ICefKeyboardHandler
	ICefBaseRefCountedOwn
}
type TCefKeyboardHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefLifeSpanHandler Parent: ICefBaseRefCounted
type ICefLifeSpanHandler interface {
	ICefBaseRefCounted
}

type ICefLifeSpanHandlerOwn interface {
	ICefLifeSpanHandler
	ICefBaseRefCountedOwn
}
type TCefLifeSpanHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefLoadHandler Parent: ICefBaseRefCounted
type ICefLoadHandler interface {
	ICefBaseRefCounted
}

type ICefLoadHandlerOwn interface {
	ICefLoadHandler
	ICefBaseRefCountedOwn
}
type TCefLoadHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefMenuModelDelegate Parent: ICefBaseRefCounted
type ICefMenuModelDelegate interface {
	ICefBaseRefCounted
}

type ICefMenuModelDelegateOwn interface {
	ICefMenuModelDelegate
	ICefBaseRefCountedOwn
}
type TCefMenuModelDelegateOwn struct {
	TCefBaseRefCountedOwn
}

// ICefNavigationEntryVisitor Parent: ICefBaseRefCounted
type ICefNavigationEntryVisitor interface {
	ICefBaseRefCounted
}

type ICefNavigationEntryVisitorOwn interface {
	ICefNavigationEntryVisitor
	ICefBaseRefCountedOwn
}
type TCefNavigationEntryVisitorOwn struct {
	TCefBaseRefCountedOwn
}

// ICefPdfPrintCallback Parent: ICefBaseRefCounted
type ICefPdfPrintCallback interface {
	ICefBaseRefCounted
}

type ICefPdfPrintCallbackOwn interface {
	ICefPdfPrintCallback
	ICefBaseRefCountedOwn
}
type TCefPdfPrintCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefRenderHandler Parent: ICefBaseRefCounted
type ICefRenderHandler interface {
	ICefBaseRefCounted
}

type ICefRenderHandlerOwn interface {
	ICefRenderHandler
	ICefBaseRefCountedOwn
}
type TCefRenderHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefRenderProcessHandler Parent: ICefBaseRefCounted
type ICefRenderProcessHandler interface {
	ICefBaseRefCounted
}

type ICefRenderProcessHandlerOwn interface {
	ICefRenderProcessHandler
	ICefBaseRefCountedOwn
}
type TCefRenderProcessHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefRequestContextHandler Parent: ICefBaseRefCounted
type ICefRequestContextHandler interface {
	ICefBaseRefCounted
}

type ICefRequestContextHandlerOwn interface {
	ICefRequestContextHandler
	ICefBaseRefCountedOwn
}
type TCefRequestContextHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefRequestHandler Parent: ICefBaseRefCounted
type ICefRequestHandler interface {
	ICefBaseRefCounted
}

type ICefRequestHandlerOwn interface {
	ICefRequestHandler
	ICefBaseRefCountedOwn
}
type TCefRequestHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefResolveCallback Parent: ICefBaseRefCounted
type ICefResolveCallback interface {
	ICefBaseRefCounted
}

type ICefResolveCallbackOwn interface {
	ICefResolveCallback
	ICefBaseRefCountedOwn
}
type TCefResolveCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefResourceBundleHandler Parent: ICefBaseRefCounted
type ICefResourceBundleHandler interface {
	ICefBaseRefCounted
}

type ICefResourceBundleHandlerOwn interface {
	ICefResourceBundleHandler
	ICefBaseRefCountedOwn
}
type TCefResourceBundleHandlerOwn struct {
	TCefBaseRefCountedOwn
}
type ICefResourceHandlerOwn interface {
	ICefResourceHandler
	ICefBaseRefCountedOwn
}
type TCefResourceHandlerOwn struct {
	TCefBaseRefCountedOwn
}
type ICefResponseFilterOwn interface {
	ICefResponseFilter
	ICefBaseRefCountedOwn
}
type TCefResponseFilterOwn struct {
	TCefBaseRefCountedOwn
}

// ICefRunFileDialogCallback Parent: ICefBaseRefCounted
type ICefRunFileDialogCallback interface {
	ICefBaseRefCounted
}

type ICefRunFileDialogCallbackOwn interface {
	ICefRunFileDialogCallback
	ICefBaseRefCountedOwn
}
type TCefRunFileDialogCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefSchemeHandlerFactory Parent: ICefBaseRefCounted
type ICefSchemeHandlerFactory interface {
	ICefBaseRefCounted
}

type ICefSchemeHandlerFactoryOwn interface {
	ICefSchemeHandlerFactory
	ICefBaseRefCountedOwn
}
type TCefSchemeHandlerFactoryOwn struct {
	TCefBaseRefCountedOwn
}

// ICefSetCookieCallback Parent: ICefBaseRefCounted
type ICefSetCookieCallback interface {
	ICefBaseRefCounted
}

type ICefSetCookieCallbackOwn interface {
	ICefSetCookieCallback
	ICefBaseRefCountedOwn
}
type TCefSetCookieCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefStringVisitor Parent: ICefBaseRefCounted
type ICefStringVisitor interface {
	ICefBaseRefCounted
}

type ICefStringVisitorOwn interface {
	ICefStringVisitor
	ICefBaseRefCountedOwn
}
type TCefStringVisitorOwn struct {
	TCefBaseRefCountedOwn
}
type ICefTaskOwn interface {
	ICefTask
	ICefBaseRefCountedOwn
}
type TCefTaskOwn struct {
	TCefBaseRefCountedOwn
}
type ICefUrlrequestClientOwn interface {
	ICefUrlrequestClient
	ICefBaseRefCountedOwn
}
type TCefUrlrequestClientOwn struct {
	TCefBaseRefCountedOwn
}

// ICefV8Accessor Parent: ICefBaseRefCounted
type ICefV8Accessor interface {
	ICefBaseRefCounted
}

type ICefV8AccessorOwn interface {
	ICefV8Accessor
	ICefBaseRefCountedOwn
}
type TCefV8AccessorOwn struct {
	TCefBaseRefCountedOwn
}
type ICefv8ArrayBufferReleaseCallbackOwn interface {
	ICefv8ArrayBufferReleaseCallback
	ICefBaseRefCountedOwn
}
type TCefv8ArrayBufferReleaseCallbackOwn struct {
	TCefBaseRefCountedOwn
}
type ICefv8HandlerOwn interface {
	ICefv8Handler
	ICefBaseRefCountedOwn
}
type TCefv8HandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefV8Interceptor Parent: ICefBaseRefCounted
type ICefV8Interceptor interface {
	ICefBaseRefCounted
}

type ICefV8InterceptorOwn interface {
	ICefV8Interceptor
	ICefBaseRefCountedOwn
}
type TCefV8InterceptorOwn struct {
	TCefBaseRefCountedOwn
}

// ICefServerHandler Parent: ICefBaseRefCounted
type ICefServerHandler interface {
	ICefBaseRefCounted
}

type ICEFServerHandlerOwn interface {
	ICefServerHandler
	ICefBaseRefCountedOwn
}
type TCEFServerHandlerOwn struct {
	TCefBaseRefCountedOwn
}
type ICefCookieAccessFilterOwn interface {
	ICefCookieAccessFilter
	ICefBaseRefCountedOwn
}
type TCefCookieAccessFilterOwn struct {
	TCefBaseRefCountedOwn
}
type ICefResourceRequestHandlerOwn interface {
	ICefResourceRequestHandler
	ICefBaseRefCountedOwn
}
type TCefResourceRequestHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefMediaObserver Parent: ICefBaseRefCounted
type ICefMediaObserver interface {
	ICefBaseRefCounted
}

type ICefMediaObserverOwn interface {
	ICefMediaObserver
	ICefBaseRefCountedOwn
}
type TCefMediaObserverOwn struct {
	TCefBaseRefCountedOwn
}

// ICefMediaRouteCreateCallback Parent: ICefBaseRefCounted
type ICefMediaRouteCreateCallback interface {
	ICefBaseRefCounted
}

type ICefMediaRouteCreateCallbackOwn interface {
	ICefMediaRouteCreateCallback
	ICefBaseRefCountedOwn
}
type TCefMediaRouteCreateCallbackOwn struct {
	TCefBaseRefCountedOwn
}
type ICefMenuButtonDelegateOwn interface {
	ICefMenuButtonDelegate
	ICefButtonDelegateOwn
}
type TCefMenuButtonDelegateOwn struct {
	TCefButtonDelegateOwn
}
type ICefButtonDelegateOwn interface {
	ICefButtonDelegate
	ICefViewDelegateOwn
}
type TCefButtonDelegateOwn struct {
	TCefViewDelegateOwn
}
type ICefBrowserViewDelegateOwn interface {
	ICefBrowserViewDelegate
	ICefViewDelegateOwn
}
type TCefBrowserViewDelegateOwn struct {
	TCefViewDelegateOwn
}
type ICefPanelDelegateOwn interface {
	ICefPanelDelegate
	ICefViewDelegateOwn
}
type TCefPanelDelegateOwn struct {
	TCefViewDelegateOwn
}
type ICefTextfieldDelegateOwn interface {
	ICefTextfieldDelegate
	ICefViewDelegateOwn
}
type TCefTextfieldDelegateOwn struct {
	TCefViewDelegateOwn
}
type ICefViewDelegateOwn interface {
	ICefViewDelegate
	ICefBaseRefCountedOwn
}
type TCefViewDelegateOwn struct {
	TCefBaseRefCountedOwn
}

// ICefAudioHandler Parent: ICefBaseRefCounted
type ICefAudioHandler interface {
	ICefBaseRefCounted
}

type ICefAudioHandlerOwn interface {
	ICefAudioHandler
	ICefBaseRefCountedOwn
}
type TCefAudioHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefDevToolsMessageObserver Parent: ICefBaseRefCounted
type ICefDevToolsMessageObserver interface {
	ICefBaseRefCounted
}

type ICEFDevToolsMessageObserverOwn interface {
	ICefDevToolsMessageObserver
	ICefBaseRefCountedOwn
}
type TCEFDevToolsMessageObserverOwn struct {
	TCefBaseRefCountedOwn
}

// ICefMediaSinkDeviceInfoCallback Parent: ICefBaseRefCounted
type ICefMediaSinkDeviceInfoCallback interface {
	ICefBaseRefCounted
}

type ICefMediaSinkDeviceInfoCallbackOwn interface {
	ICefMediaSinkDeviceInfoCallback
	ICefBaseRefCountedOwn
}
type TCefMediaSinkDeviceInfoCallbackOwn struct {
	TCefBaseRefCountedOwn
}

// ICefPrintHandler Parent: ICefBaseRefCounted
type ICefPrintHandler interface {
	ICefBaseRefCounted
}

type ICefPrintHandlerOwn interface {
	ICefPrintHandler
	ICefBaseRefCountedOwn
}
type TCefPrintHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefFrameHandler Parent: ICefBaseRefCounted
type ICefFrameHandler interface {
	ICefBaseRefCounted
}

type ICefFrameHandlerOwn interface {
	ICefFrameHandler
	ICefBaseRefCountedOwn
}
type TCefFrameHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefCommandHandler Parent: ICefBaseRefCounted
type ICefCommandHandler interface {
	ICefBaseRefCounted
}

type ICefCommandHandlerOwn interface {
	ICefCommandHandler
	ICefBaseRefCountedOwn
}
type TCefCommandHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefPermissionHandler Parent: ICefBaseRefCounted
type ICefPermissionHandler interface {
	ICefBaseRefCounted
}

type ICefPermissionHandlerOwn interface {
	ICefPermissionHandler
	ICefBaseRefCountedOwn
}
type TCefPermissionHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefClient Parent: ICefBaseRefCounted
type ICefClient interface {
	ICefBaseRefCounted
}

type ICefClientOwn interface {
	ICefClient
	ICefBaseRefCountedOwn
}
type TCefClientOwn struct {
	TCefBaseRefCountedOwn
}
type ICefWindowDelegateOwn interface {
	ICefWindowDelegate
	ICefPanelDelegateOwn
}
type TCefWindowDelegateOwn struct {
	TCefPanelDelegateOwn
}

// ICefPreferenceObserver Parent: ICefBaseRefCounted
type ICefPreferenceObserver interface {
	ICefBaseRefCounted
}

type ICefPreferenceObserverOwn interface {
	ICefPreferenceObserver
	ICefBaseRefCountedOwn
}
type TCefPreferenceObserverOwn struct {
	TCefBaseRefCountedOwn
}

// ICefSettingObserver Parent: ICefBaseRefCounted
type ICefSettingObserver interface {
	ICefBaseRefCounted
}

type ICefSettingObserverOwn interface {
	ICefSettingObserver
	ICefBaseRefCountedOwn
}
type TCefSettingObserverOwn struct {
	TCefBaseRefCountedOwn
}

// ICefAccessibilityHandler Parent: ICefBaseRefCounted
type ICefAccessibilityHandler interface {
	ICefBaseRefCounted
}

type ICEFAccessibilityHandlerOwn interface {
	ICefAccessibilityHandler
	ICefBaseRefCountedOwn
}
type TCEFAccessibilityHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefWriteHandler Parent: ICefBaseRefCounted
type ICefWriteHandler interface {
	ICefBaseRefCounted
}

type ICefWriteHandlerOwn interface {
	ICefWriteHandler
	ICefBaseRefCountedOwn
}
type TCefWriteHandlerOwn struct {
	TCefBaseRefCountedOwn
}

// ICefApp Parent: ICefBaseRefCounted
type ICefApp interface {
	ICefBaseRefCounted
}

type ICefAppOwn interface {
	ICefApp
	ICefBaseRefCountedOwn
}
type TCefAppOwn struct {
	TCefBaseRefCountedOwn
}
