//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/lcl/base"

// AsCustomAccessibilityHandler Convert a pointer object to an existing class object
func AsCustomAccessibilityHandler(obj any) ICustomAccessibilityHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomAccessibilityHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCefApp Convert a pointer object to an existing class object
func AsCustomCefApp(obj any) ICustomCefApp {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCefApp)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefApplicationCore Convert a pointer object to an existing class object
func AsCefApplicationCore(obj any) ICefApplicationCore {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefApplicationCore)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefApplication Convert a pointer object to an existing class object
func AsCefApplication(obj any) ICefApplication {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefApplication)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomBrowserProcessHandler Convert a pointer object to an existing class object
func AsCefCustomBrowserProcessHandler(obj any) ICefCustomBrowserProcessHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomBrowserProcessHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChromiumFontOptions Convert a pointer object to an existing class object
func AsChromiumFontOptions(obj any) IChromiumFontOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChromiumFontOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChromiumOptions Convert a pointer object to an existing class object
func AsChromiumOptions(obj any) IChromiumOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChromiumOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFWinControl Convert a pointer object to an existing class object
func AsCEFWinControl(obj any) ICEFWinControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFWinControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFLinkedWinControlBase Convert a pointer object to an existing class object
func AsCEFLinkedWinControlBase(obj any) ICEFLinkedWinControlBase {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFLinkedWinControlBase)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChromiumWindow Convert a pointer object to an existing class object
func AsChromiumWindow(obj any) IChromiumWindow {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChromiumWindow)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomCompletionCallback Convert a pointer object to an existing class object
func AsCefCustomCompletionCallback(obj any) ICefCustomCompletionCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomCompletionCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefEventCompletionCallback Convert a pointer object to an existing class object
func AsCefEventCompletionCallback(obj any) ICefEventCompletionCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefEventCompletionCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomContextMenuHandler Convert a pointer object to an existing class object
func AsCustomContextMenuHandler(obj any) ICustomContextMenuHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomContextMenuHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomCookieVisitor Convert a pointer object to an existing class object
func AsCefCustomCookieVisitor(obj any) ICefCustomCookieVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomCookieVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomDeleteCookiesCallback Convert a pointer object to an existing class object
func AsCefCustomDeleteCookiesCallback(obj any) ICefCustomDeleteCookiesCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomDeleteCookiesCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDialogHandler Convert a pointer object to an existing class object
func AsCustomDialogHandler(obj any) ICustomDialogHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDialogHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDisplayHandler Convert a pointer object to an existing class object
func AsCustomDisplayHandler(obj any) ICustomDisplayHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDisplayHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDownloadHandler Convert a pointer object to an existing class object
func AsCustomDownloadHandler(obj any) ICustomDownloadHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDownloadHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomDownloadImageCallback Convert a pointer object to an existing class object
func AsCefCustomDownloadImageCallback(obj any) ICefCustomDownloadImageCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomDownloadImageCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDragHandler Convert a pointer object to an existing class object
func AsCustomDragHandler(obj any) ICustomDragHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDragHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFindHandler Convert a pointer object to an existing class object
func AsCustomFindHandler(obj any) ICustomFindHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFindHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFocusHandler Convert a pointer object to an existing class object
func AsCustomFocusHandler(obj any) ICustomFocusHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFocusHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomJsDialogHandler Convert a pointer object to an existing class object
func AsCustomJsDialogHandler(obj any) ICustomJsDialogHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomJsDialogHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomKeyboardHandler Convert a pointer object to an existing class object
func AsCustomKeyboardHandler(obj any) ICustomKeyboardHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomKeyboardHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomLifeSpanHandler Convert a pointer object to an existing class object
func AsCustomLifeSpanHandler(obj any) ICustomLifeSpanHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomLifeSpanHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRenderLoadHandler Convert a pointer object to an existing class object
func AsCustomRenderLoadHandler(obj any) ICustomRenderLoadHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRenderLoadHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomLoadHandler Convert a pointer object to an existing class object
func AsCustomLoadHandler(obj any) ICustomLoadHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomLoadHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCefNavigationEntryVisitor Convert a pointer object to an existing class object
func AsCustomCefNavigationEntryVisitor(obj any) ICustomCefNavigationEntryVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCefNavigationEntryVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomPDFPrintCallBack Convert a pointer object to an existing class object
func AsCefCustomPDFPrintCallBack(obj any) ICefCustomPDFPrintCallBack {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomPDFPrintCallBack)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPDFPrintOptions Convert a pointer object to an existing class object
func AsPDFPrintOptions(obj any) IPDFPrintOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPDFPrintOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRenderHandler Convert a pointer object to an existing class object
func AsCustomRenderHandler(obj any) ICustomRenderHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRenderHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomRenderProcessHandler Convert a pointer object to an existing class object
func AsCefCustomRenderProcessHandler(obj any) ICefCustomRenderProcessHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomRenderProcessHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRequestContextHandler Convert a pointer object to an existing class object
func AsCustomRequestContextHandler(obj any) ICustomRequestContextHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRequestContextHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRequestHandler Convert a pointer object to an existing class object
func AsCustomRequestHandler(obj any) ICustomRequestHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRequestHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomResolveCallback Convert a pointer object to an existing class object
func AsCefCustomResolveCallback(obj any) ICefCustomResolveCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomResolveCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomResourceBundleHandler Convert a pointer object to an existing class object
func AsCefCustomResourceBundleHandler(obj any) ICefCustomResourceBundleHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomResourceBundleHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomResponseFilter Convert a pointer object to an existing class object
func AsCustomResponseFilter(obj any) ICustomResponseFilter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomResponseFilter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResourceHandlerRef Convert a pointer object to an existing class object
func AsCefResourceHandlerRef(obj any) ICefResourceHandlerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResourceHandlerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResponseFilterRef Convert a pointer object to an existing class object
func AsCefResponseFilterRef(obj any) ICefResponseFilterRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResponseFilterRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomSetCookieCallback Convert a pointer object to an existing class object
func AsCefCustomSetCookieCallback(obj any) ICefCustomSetCookieCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomSetCookieCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStringListRef Convert a pointer object to an existing class object
func AsCefStringListRef(obj any) ICefStringListRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStringListRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStringMapRef Convert a pointer object to an existing class object
func AsCefStringMapRef(obj any) ICefStringMapRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStringMapRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCefStringVisitor Convert a pointer object to an existing class object
func AsCustomCefStringVisitor(obj any) ICustomCefStringVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCefStringVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefTaskRef Convert a pointer object to an existing class object
func AsCefTaskRef(obj any) ICefTaskRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefTaskRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefManagedTask Convert a pointer object to an existing class object
func AsCefManagedTask(obj any) ICefManagedTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefManagedTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefChromiumTask Convert a pointer object to an existing class object
func AsCefChromiumTask(obj any) ICefChromiumTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefChromiumTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefViewDelegateTask Convert a pointer object to an existing class object
func AsCefViewDelegateTask(obj any) ICefViewDelegateTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefViewDelegateTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefGenericTask Convert a pointer object to an existing class object
func AsCefGenericTask(obj any) ICefGenericTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefGenericTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefURLRequestClientTask Convert a pointer object to an existing class object
func AsCefURLRequestClientTask(obj any) ICefURLRequestClientTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefURLRequestClientTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefUrlrequestClientRef Convert a pointer object to an existing class object
func AsCefUrlrequestClientRef(obj any) ICefUrlrequestClientRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefUrlrequestClientRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCefUrlrequestClient Convert a pointer object to an existing class object
func AsCustomCefUrlrequestClient(obj any) ICustomCefUrlrequestClient {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCefUrlrequestClient)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefv8HandlerRef Convert a pointer object to an existing class object
func AsCefv8HandlerRef(obj any) ICefv8HandlerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefv8HandlerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefv8ArrayBufferReleaseCallbackRef Convert a pointer object to an existing class object
func AsCefv8ArrayBufferReleaseCallbackRef(obj any) ICefv8ArrayBufferReleaseCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefv8ArrayBufferReleaseCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFWindowParent Convert a pointer object to an existing class object
func AsCEFWindowParent(obj any) ICEFWindowParent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFWindowParent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFWorkScheduler Convert a pointer object to an existing class object
func AsCEFWorkScheduler(obj any) ICEFWorkScheduler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFWorkScheduler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBytesWriteHandler Convert a pointer object to an existing class object
func AsCefBytesWriteHandler(obj any) ICefBytesWriteHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBytesWriteHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChromiumCore Convert a pointer object to an existing class object
func AsChromiumCore(obj any) IChromiumCore {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChromiumCore)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChromium Convert a pointer object to an existing class object
func AsChromium(obj any) IChromium {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChromium)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBufferPanel Convert a pointer object to an existing class object
func AsBufferPanel(obj any) IBufferPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBufferPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFServerComponent Convert a pointer object to an existing class object
func AsCEFServerComponent(obj any) ICEFServerComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFServerComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomServerHandler Convert a pointer object to an existing class object
func AsCustomServerHandler(obj any) ICustomServerHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomServerHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFLinkedWindowParent Convert a pointer object to an existing class object
func AsCEFLinkedWindowParent(obj any) ICEFLinkedWindowParent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFLinkedWindowParent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFUrlRequestClientComponent Convert a pointer object to an existing class object
func AsCEFUrlRequestClientComponent(obj any) ICEFUrlRequestClientComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFUrlRequestClientComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFOSRIMEHandler Convert a pointer object to an existing class object
func AsCEFOSRIMEHandler(obj any) ICEFOSRIMEHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFOSRIMEHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCookieAccessFilter Convert a pointer object to an existing class object
func AsCustomCookieAccessFilter(obj any) ICustomCookieAccessFilter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCookieAccessFilter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCookieAccessFilterRef Convert a pointer object to an existing class object
func AsCefCookieAccessFilterRef(obj any) ICefCookieAccessFilterRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCookieAccessFilterRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomResourceRequestHandler Convert a pointer object to an existing class object
func AsCustomResourceRequestHandler(obj any) ICustomResourceRequestHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomResourceRequestHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResourceRequestHandlerRef Convert a pointer object to an existing class object
func AsCefResourceRequestHandlerRef(obj any) ICefResourceRequestHandlerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResourceRequestHandlerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFSentinel Convert a pointer object to an existing class object
func AsCEFSentinel(obj any) ICEFSentinel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFSentinel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFOAuth2Helper Convert a pointer object to an existing class object
func AsCEFOAuth2Helper(obj any) ICEFOAuth2Helper {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFOAuth2Helper)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomMediaObserver Convert a pointer object to an existing class object
func AsCustomMediaObserver(obj any) ICustomMediaObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomMediaObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomMediaRouteCreateCallback Convert a pointer object to an existing class object
func AsCefCustomMediaRouteCreateCallback(obj any) ICefCustomMediaRouteCreateCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomMediaRouteCreateCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomWindowDelegate Convert a pointer object to an existing class object
func AsCustomWindowDelegate(obj any) ICustomWindowDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomWindowDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefViewDelegateRef Convert a pointer object to an existing class object
func AsCefViewDelegateRef(obj any) ICefViewDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefViewDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPanelDelegateRef Convert a pointer object to an existing class object
func AsCefPanelDelegateRef(obj any) ICefPanelDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPanelDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefWindowDelegateRef Convert a pointer object to an existing class object
func AsCefWindowDelegateRef(obj any) ICefWindowDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefWindowDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomMenuButtonDelegate Convert a pointer object to an existing class object
func AsCustomMenuButtonDelegate(obj any) ICustomMenuButtonDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomMenuButtonDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefButtonDelegateRef Convert a pointer object to an existing class object
func AsCefButtonDelegateRef(obj any) ICefButtonDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefButtonDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMenuButtonDelegateRef Convert a pointer object to an existing class object
func AsCefMenuButtonDelegateRef(obj any) ICefMenuButtonDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMenuButtonDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomButtonDelegate Convert a pointer object to an existing class object
func AsCustomButtonDelegate(obj any) ICustomButtonDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomButtonDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomBrowserViewDelegate Convert a pointer object to an existing class object
func AsCustomBrowserViewDelegate(obj any) ICustomBrowserViewDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomBrowserViewDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBrowserViewDelegateRef Convert a pointer object to an existing class object
func AsCefBrowserViewDelegateRef(obj any) ICefBrowserViewDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBrowserViewDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPanelDelegate Convert a pointer object to an existing class object
func AsCustomPanelDelegate(obj any) ICustomPanelDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPanelDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTextfieldDelegate Convert a pointer object to an existing class object
func AsCustomTextfieldDelegate(obj any) ICustomTextfieldDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTextfieldDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefTextfieldDelegateRef Convert a pointer object to an existing class object
func AsCefTextfieldDelegateRef(obj any) ICefTextfieldDelegateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefTextfieldDelegateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomViewDelegate Convert a pointer object to an existing class object
func AsCustomViewDelegate(obj any) ICustomViewDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomViewDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFViewComponent Convert a pointer object to an existing class object
func AsCEFViewComponent(obj any) ICEFViewComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFViewComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFButtonComponent Convert a pointer object to an existing class object
func AsCEFButtonComponent(obj any) ICEFButtonComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFButtonComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFLabelButtonComponent Convert a pointer object to an existing class object
func AsCEFLabelButtonComponent(obj any) ICEFLabelButtonComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFLabelButtonComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFMenuButtonComponent Convert a pointer object to an existing class object
func AsCEFMenuButtonComponent(obj any) ICEFMenuButtonComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFMenuButtonComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFBrowserViewComponent Convert a pointer object to an existing class object
func AsCEFBrowserViewComponent(obj any) ICEFBrowserViewComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFBrowserViewComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFPanelComponent Convert a pointer object to an existing class object
func AsCEFPanelComponent(obj any) ICEFPanelComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFPanelComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFWindowComponent Convert a pointer object to an existing class object
func AsCEFWindowComponent(obj any) ICEFWindowComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFWindowComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFScrollViewComponent Convert a pointer object to an existing class object
func AsCEFScrollViewComponent(obj any) ICEFScrollViewComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFScrollViewComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFTextfieldComponent Convert a pointer object to an existing class object
func AsCEFTextfieldComponent(obj any) ICEFTextfieldComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFTextfieldComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomAudioHandler Convert a pointer object to an existing class object
func AsCustomAudioHandler(obj any) ICustomAudioHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomAudioHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDevToolsMessageObserver Convert a pointer object to an existing class object
func AsCustomDevToolsMessageObserver(obj any) ICustomDevToolsMessageObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDevToolsMessageObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomMediaSinkDeviceInfoCallback Convert a pointer object to an existing class object
func AsCefCustomMediaSinkDeviceInfoCallback(obj any) ICefCustomMediaSinkDeviceInfoCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomMediaSinkDeviceInfoCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFJson Convert a pointer object to an existing class object
func AsCEFJson(obj any) ICEFJson {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFJson)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFBitmapBitBuffer Convert a pointer object to an existing class object
func AsCEFBitmapBitBuffer(obj any) ICEFBitmapBitBuffer {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFBitmapBitBuffer)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPrintHandler Convert a pointer object to an existing class object
func AsCustomPrintHandler(obj any) ICustomPrintHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPrintHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFTimerWorkScheduler Convert a pointer object to an existing class object
func AsCEFTimerWorkScheduler(obj any) ICEFTimerWorkScheduler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFTimerWorkScheduler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFrameHandler Convert a pointer object to an existing class object
func AsCustomFrameHandler(obj any) ICustomFrameHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFrameHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFFileDialogInfo Convert a pointer object to an existing class object
func AsCEFFileDialogInfo(obj any) ICEFFileDialogInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFFileDialogInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCommandHandler Convert a pointer object to an existing class object
func AsCustomCommandHandler(obj any) ICustomCommandHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCommandHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPermissionHandler Convert a pointer object to an existing class object
func AsCustomPermissionHandler(obj any) ICustomPermissionHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPermissionHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFBaseScopedWrapperRef Convert a pointer object to an existing class object
func AsCEFBaseScopedWrapperRef(obj any) ICEFBaseScopedWrapperRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFBaseScopedWrapperRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPreferenceRegistrarRef Convert a pointer object to an existing class object
func AsCefPreferenceRegistrarRef(obj any) ICefPreferenceRegistrarRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPreferenceRegistrarRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSchemeRegistrarRef Convert a pointer object to an existing class object
func AsCefSchemeRegistrarRef(obj any) ICefSchemeRegistrarRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSchemeRegistrarRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomClientHandler Convert a pointer object to an existing class object
func AsCustomClientHandler(obj any) ICustomClientHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomClientHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStringListOwn Convert a pointer object to an existing class object
func AsCefStringListOwn(obj any) ICefStringListOwn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStringListOwn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStringMapOwn Convert a pointer object to an existing class object
func AsCefStringMapOwn(obj any) ICefStringMapOwn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStringMapOwn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFProxySettings Convert a pointer object to an existing class object
func AsCEFProxySettings(obj any) ICEFProxySettings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFProxySettings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngBrowserProcessHandler Convert a pointer object to an existing class object
func AsEngBrowserProcessHandler(obj any) IEngBrowserProcessHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngBrowserProcessHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngCompletionCallback Convert a pointer object to an existing class object
func AsEngCompletionCallback(obj any) IEngCompletionCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngCompletionCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngContextMenuHandler Convert a pointer object to an existing class object
func AsEngContextMenuHandler(obj any) IEngContextMenuHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngContextMenuHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngCookieVisitor Convert a pointer object to an existing class object
func AsEngCookieVisitor(obj any) IEngCookieVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngCookieVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDeleteCookiesCallback Convert a pointer object to an existing class object
func AsEngDeleteCookiesCallback(obj any) IEngDeleteCookiesCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDeleteCookiesCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDialogHandler Convert a pointer object to an existing class object
func AsEngDialogHandler(obj any) IEngDialogHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDialogHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDisplayHandler Convert a pointer object to an existing class object
func AsEngDisplayHandler(obj any) IEngDisplayHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDisplayHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDomVisitor Convert a pointer object to an existing class object
func AsEngDomVisitor(obj any) IEngDomVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDomVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDownloadHandler Convert a pointer object to an existing class object
func AsEngDownloadHandler(obj any) IEngDownloadHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDownloadHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDownloadImageCallback Convert a pointer object to an existing class object
func AsEngDownloadImageCallback(obj any) IEngDownloadImageCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDownloadImageCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDragHandler Convert a pointer object to an existing class object
func AsEngDragHandler(obj any) IEngDragHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDragHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngEndTracingCallback Convert a pointer object to an existing class object
func AsEngEndTracingCallback(obj any) IEngEndTracingCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngEndTracingCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngFindHandler Convert a pointer object to an existing class object
func AsEngFindHandler(obj any) IEngFindHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngFindHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngFocusHandler Convert a pointer object to an existing class object
func AsEngFocusHandler(obj any) IEngFocusHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngFocusHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngJsDialogHandler Convert a pointer object to an existing class object
func AsEngJsDialogHandler(obj any) IEngJsDialogHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngJsDialogHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngKeyboardHandler Convert a pointer object to an existing class object
func AsEngKeyboardHandler(obj any) IEngKeyboardHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngKeyboardHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngLifeSpanHandler Convert a pointer object to an existing class object
func AsEngLifeSpanHandler(obj any) IEngLifeSpanHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngLifeSpanHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngLoadHandler Convert a pointer object to an existing class object
func AsEngLoadHandler(obj any) IEngLoadHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngLoadHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngMenuModelDelegate Convert a pointer object to an existing class object
func AsEngMenuModelDelegate(obj any) IEngMenuModelDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngMenuModelDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngNavigationEntryVisitor Convert a pointer object to an existing class object
func AsEngNavigationEntryVisitor(obj any) IEngNavigationEntryVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngNavigationEntryVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngPdfPrintCallback Convert a pointer object to an existing class object
func AsEngPdfPrintCallback(obj any) IEngPdfPrintCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngPdfPrintCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngRenderHandler Convert a pointer object to an existing class object
func AsEngRenderHandler(obj any) IEngRenderHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngRenderHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngRenderProcessHandler Convert a pointer object to an existing class object
func AsEngRenderProcessHandler(obj any) IEngRenderProcessHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngRenderProcessHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngRequestContextHandler Convert a pointer object to an existing class object
func AsEngRequestContextHandler(obj any) IEngRequestContextHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngRequestContextHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngRequestHandler Convert a pointer object to an existing class object
func AsEngRequestHandler(obj any) IEngRequestHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngRequestHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngResolveCallback Convert a pointer object to an existing class object
func AsEngResolveCallback(obj any) IEngResolveCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngResolveCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngResourceBundleHandler Convert a pointer object to an existing class object
func AsEngResourceBundleHandler(obj any) IEngResourceBundleHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngResourceBundleHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngResourceHandler Convert a pointer object to an existing class object
func AsEngResourceHandler(obj any) IEngResourceHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngResourceHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngResponseFilter Convert a pointer object to an existing class object
func AsEngResponseFilter(obj any) IEngResponseFilter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngResponseFilter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngRunFileDialogCallback Convert a pointer object to an existing class object
func AsEngRunFileDialogCallback(obj any) IEngRunFileDialogCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngRunFileDialogCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngSchemeHandlerFactory Convert a pointer object to an existing class object
func AsEngSchemeHandlerFactory(obj any) IEngSchemeHandlerFactory {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngSchemeHandlerFactory)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngSetCookieCallback Convert a pointer object to an existing class object
func AsEngSetCookieCallback(obj any) IEngSetCookieCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngSetCookieCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngStringVisitor Convert a pointer object to an existing class object
func AsEngStringVisitor(obj any) IEngStringVisitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngStringVisitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngTask Convert a pointer object to an existing class object
func AsEngTask(obj any) IEngTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngUrlrequestClient Convert a pointer object to an existing class object
func AsEngUrlrequestClient(obj any) IEngUrlrequestClient {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngUrlrequestClient)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngV8Accessor Convert a pointer object to an existing class object
func AsEngV8Accessor(obj any) IEngV8Accessor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngV8Accessor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngV8ArrayBufferReleaseCallback Convert a pointer object to an existing class object
func AsEngV8ArrayBufferReleaseCallback(obj any) IEngV8ArrayBufferReleaseCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngV8ArrayBufferReleaseCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngV8Handler Convert a pointer object to an existing class object
func AsEngV8Handler(obj any) IEngV8Handler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngV8Handler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngV8Interceptor Convert a pointer object to an existing class object
func AsEngV8Interceptor(obj any) IEngV8Interceptor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngV8Interceptor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngServerHandler Convert a pointer object to an existing class object
func AsEngServerHandler(obj any) IEngServerHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngServerHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngCookieAccessFilter Convert a pointer object to an existing class object
func AsEngCookieAccessFilter(obj any) IEngCookieAccessFilter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngCookieAccessFilter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngResourceRequestHandler Convert a pointer object to an existing class object
func AsEngResourceRequestHandler(obj any) IEngResourceRequestHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngResourceRequestHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngMediaObserver Convert a pointer object to an existing class object
func AsEngMediaObserver(obj any) IEngMediaObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngMediaObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngMediaRouteCreateCallback Convert a pointer object to an existing class object
func AsEngMediaRouteCreateCallback(obj any) IEngMediaRouteCreateCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngMediaRouteCreateCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngMenuButtonDelegate Convert a pointer object to an existing class object
func AsEngMenuButtonDelegate(obj any) IEngMenuButtonDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngMenuButtonDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngButtonDelegate Convert a pointer object to an existing class object
func AsEngButtonDelegate(obj any) IEngButtonDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngButtonDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngBrowserViewDelegate Convert a pointer object to an existing class object
func AsEngBrowserViewDelegate(obj any) IEngBrowserViewDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngBrowserViewDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngPanelDelegate Convert a pointer object to an existing class object
func AsEngPanelDelegate(obj any) IEngPanelDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngPanelDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngTextfieldDelegate Convert a pointer object to an existing class object
func AsEngTextfieldDelegate(obj any) IEngTextfieldDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngTextfieldDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngViewDelegate Convert a pointer object to an existing class object
func AsEngViewDelegate(obj any) IEngViewDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngViewDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngAudioHandler Convert a pointer object to an existing class object
func AsEngAudioHandler(obj any) IEngAudioHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngAudioHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngDevToolsMessageObserver Convert a pointer object to an existing class object
func AsEngDevToolsMessageObserver(obj any) IEngDevToolsMessageObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngDevToolsMessageObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngMediaSinkDeviceInfoCallback Convert a pointer object to an existing class object
func AsEngMediaSinkDeviceInfoCallback(obj any) IEngMediaSinkDeviceInfoCallback {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngMediaSinkDeviceInfoCallback)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngPrintHandler Convert a pointer object to an existing class object
func AsEngPrintHandler(obj any) IEngPrintHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngPrintHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngFrameHandler Convert a pointer object to an existing class object
func AsEngFrameHandler(obj any) IEngFrameHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngFrameHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngCommandHandler Convert a pointer object to an existing class object
func AsEngCommandHandler(obj any) IEngCommandHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngCommandHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngPermissionHandler Convert a pointer object to an existing class object
func AsEngPermissionHandler(obj any) IEngPermissionHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngPermissionHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngClient Convert a pointer object to an existing class object
func AsEngClient(obj any) IEngClient {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngClient)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngWindowDelegate Convert a pointer object to an existing class object
func AsEngWindowDelegate(obj any) IEngWindowDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngWindowDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngPreferenceObserver Convert a pointer object to an existing class object
func AsEngPreferenceObserver(obj any) IEngPreferenceObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngPreferenceObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngSettingObserver Convert a pointer object to an existing class object
func AsEngSettingObserver(obj any) IEngSettingObserver {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngSettingObserver)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngAccessibilityHandler Convert a pointer object to an existing class object
func AsEngAccessibilityHandler(obj any) IEngAccessibilityHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngAccessibilityHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngWriteHandler Convert a pointer object to an existing class object
func AsEngWriteHandler(obj any) IEngWriteHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngWriteHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngApp Convert a pointer object to an existing class object
func AsEngApp(obj any) IEngApp {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngApp)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefAuthCallbackRef Convert a pointer object to an existing class object
func AsCefAuthCallbackRef(obj any) ICefAuthCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefAuthCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBeforeDownloadCallbackRef Convert a pointer object to an existing class object
func AsCefBeforeDownloadCallbackRef(obj any) ICefBeforeDownloadCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBeforeDownloadCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBinaryValueRef Convert a pointer object to an existing class object
func AsCefBinaryValueRef(obj any) ICefBinaryValueRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBinaryValueRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBrowserRef Convert a pointer object to an existing class object
func AsCefBrowserRef(obj any) ICefBrowserRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBrowserRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBrowserHostRef Convert a pointer object to an existing class object
func AsCefBrowserHostRef(obj any) ICefBrowserHostRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBrowserHostRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCallbackRef Convert a pointer object to an existing class object
func AsCefCallbackRef(obj any) ICefCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCommandLineRef Convert a pointer object to an existing class object
func AsCefCommandLineRef(obj any) ICefCommandLineRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCommandLineRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefContextMenuParamsRef Convert a pointer object to an existing class object
func AsCefContextMenuParamsRef(obj any) ICefContextMenuParamsRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefContextMenuParamsRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCookieManagerRef Convert a pointer object to an existing class object
func AsCefCookieManagerRef(obj any) ICefCookieManagerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCookieManagerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomStreamReader Convert a pointer object to an existing class object
func AsCefCustomStreamReader(obj any) ICefCustomStreamReader {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomStreamReader)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDictionaryValueRef Convert a pointer object to an existing class object
func AsCefDictionaryValueRef(obj any) ICefDictionaryValueRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDictionaryValueRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDomDocumentRef Convert a pointer object to an existing class object
func AsCefDomDocumentRef(obj any) ICefDomDocumentRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDomDocumentRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDomNodeRef Convert a pointer object to an existing class object
func AsCefDomNodeRef(obj any) ICefDomNodeRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDomNodeRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDownloadItemRef Convert a pointer object to an existing class object
func AsCefDownloadItemRef(obj any) ICefDownloadItemRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDownloadItemRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDownloadItemCallbackRef Convert a pointer object to an existing class object
func AsCefDownloadItemCallbackRef(obj any) ICefDownloadItemCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDownloadItemCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDragDataRef Convert a pointer object to an existing class object
func AsCefDragDataRef(obj any) ICefDragDataRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDragDataRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefFileDialogCallbackRef Convert a pointer object to an existing class object
func AsCefFileDialogCallbackRef(obj any) ICefFileDialogCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefFileDialogCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefFrameRef Convert a pointer object to an existing class object
func AsCefFrameRef(obj any) ICefFrameRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefFrameRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefImageRef Convert a pointer object to an existing class object
func AsCefImageRef(obj any) ICefImageRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefImageRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefJsDialogCallbackRef Convert a pointer object to an existing class object
func AsCefJsDialogCallbackRef(obj any) ICefJsDialogCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefJsDialogCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefListValueRef Convert a pointer object to an existing class object
func AsCefListValueRef(obj any) ICefListValueRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefListValueRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMenuModelRef Convert a pointer object to an existing class object
func AsCefMenuModelRef(obj any) ICefMenuModelRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMenuModelRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefNavigationEntryRef Convert a pointer object to an existing class object
func AsCefNavigationEntryRef(obj any) ICefNavigationEntryRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefNavigationEntryRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPostDataRef Convert a pointer object to an existing class object
func AsCefPostDataRef(obj any) ICefPostDataRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPostDataRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPostDataElementRef Convert a pointer object to an existing class object
func AsCefPostDataElementRef(obj any) ICefPostDataElementRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPostDataElementRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPrintSettingsRef Convert a pointer object to an existing class object
func AsCefPrintSettingsRef(obj any) ICefPrintSettingsRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPrintSettingsRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefProcessMessageRef Convert a pointer object to an existing class object
func AsCefProcessMessageRef(obj any) ICefProcessMessageRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefProcessMessageRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefRequestRef Convert a pointer object to an existing class object
func AsCefRequestRef(obj any) ICefRequestRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefRequestRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefRequestContextRef Convert a pointer object to an existing class object
func AsCefRequestContextRef(obj any) ICefRequestContextRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefRequestContextRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResourceBundleRef Convert a pointer object to an existing class object
func AsCefResourceBundleRef(obj any) ICefResourceBundleRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResourceBundleRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResponseRef Convert a pointer object to an existing class object
func AsCefResponseRef(obj any) ICefResponseRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResponseRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefRunContextMenuCallbackRef Convert a pointer object to an existing class object
func AsCefRunContextMenuCallbackRef(obj any) ICefRunContextMenuCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefRunContextMenuCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSelectClientCertificateCallbackRef Convert a pointer object to an existing class object
func AsCefSelectClientCertificateCallbackRef(obj any) ICefSelectClientCertificateCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSelectClientCertificateCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSslInfoRef Convert a pointer object to an existing class object
func AsCefSslInfoRef(obj any) ICefSslInfoRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSslInfoRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSSLStatusRef Convert a pointer object to an existing class object
func AsCefSSLStatusRef(obj any) ICefSSLStatusRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSSLStatusRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStreamReaderRef Convert a pointer object to an existing class object
func AsCefStreamReaderRef(obj any) ICefStreamReaderRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStreamReaderRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStreamWriterRef Convert a pointer object to an existing class object
func AsCefStreamWriterRef(obj any) ICefStreamWriterRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStreamWriterRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefStringMultimapOwn Convert a pointer object to an existing class object
func AsCefStringMultimapOwn(obj any) ICefStringMultimapOwn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefStringMultimapOwn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomStringMultimap Convert a pointer object to an existing class object
func AsCefCustomStringMultimap(obj any) ICefCustomStringMultimap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomStringMultimap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefTaskRunnerRef Convert a pointer object to an existing class object
func AsCefTaskRunnerRef(obj any) ICefTaskRunnerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefTaskRunnerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefThreadRef Convert a pointer object to an existing class object
func AsCefThreadRef(obj any) ICefThreadRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefThreadRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefUrlRequestRef Convert a pointer object to an existing class object
func AsCefUrlRequestRef(obj any) ICefUrlRequestRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefUrlRequestRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefv8ContextRef Convert a pointer object to an existing class object
func AsCefv8ContextRef(obj any) ICefv8ContextRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefv8ContextRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefV8ExceptionRef Convert a pointer object to an existing class object
func AsCefV8ExceptionRef(obj any) ICefV8ExceptionRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefV8ExceptionRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefV8StackFrameRef Convert a pointer object to an existing class object
func AsCefV8StackFrameRef(obj any) ICefV8StackFrameRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefV8StackFrameRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefV8StackTraceRef Convert a pointer object to an existing class object
func AsCefV8StackTraceRef(obj any) ICefV8StackTraceRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefV8StackTraceRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefv8ValueRef Convert a pointer object to an existing class object
func AsCefv8ValueRef(obj any) ICefv8ValueRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefv8ValueRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefValueRef Convert a pointer object to an existing class object
func AsCefValueRef(obj any) ICefValueRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefValueRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefWaitableEventRef Convert a pointer object to an existing class object
func AsCefWaitableEventRef(obj any) ICefWaitableEventRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefWaitableEventRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFX509CertificateRef Convert a pointer object to an existing class object
func AsCEFX509CertificateRef(obj any) ICEFX509CertificateRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFX509CertificateRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefX509CertPrincipalRef Convert a pointer object to an existing class object
func AsCefX509CertPrincipalRef(obj any) ICefX509CertPrincipalRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefX509CertPrincipalRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefXmlReaderRef Convert a pointer object to an existing class object
func AsCefXmlReaderRef(obj any) ICefXmlReaderRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefXmlReaderRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefZipReaderRef Convert a pointer object to an existing class object
func AsCefZipReaderRef(obj any) ICefZipReaderRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefZipReaderRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCEFServerRef Convert a pointer object to an existing class object
func AsCEFServerRef(obj any) ICEFServerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCEFServerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResourceReadCallbackRef Convert a pointer object to an existing class object
func AsCefResourceReadCallbackRef(obj any) ICefResourceReadCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResourceReadCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefResourceSkipCallbackRef Convert a pointer object to an existing class object
func AsCefResourceSkipCallbackRef(obj any) ICefResourceSkipCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefResourceSkipCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMediaRouteRef Convert a pointer object to an existing class object
func AsCefMediaRouteRef(obj any) ICefMediaRouteRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMediaRouteRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMediaRouterRef Convert a pointer object to an existing class object
func AsCefMediaRouterRef(obj any) ICefMediaRouterRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMediaRouterRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMediaSinkRef Convert a pointer object to an existing class object
func AsCefMediaSinkRef(obj any) ICefMediaSinkRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMediaSinkRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMediaSourceRef Convert a pointer object to an existing class object
func AsCefMediaSourceRef(obj any) ICefMediaSourceRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMediaSourceRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefWindowRef Convert a pointer object to an existing class object
func AsCefWindowRef(obj any) ICefWindowRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefWindowRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMenuButtonRef Convert a pointer object to an existing class object
func AsCefMenuButtonRef(obj any) ICefMenuButtonRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMenuButtonRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefLabelButtonRef Convert a pointer object to an existing class object
func AsCefLabelButtonRef(obj any) ICefLabelButtonRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefLabelButtonRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefButtonRef Convert a pointer object to an existing class object
func AsCefButtonRef(obj any) ICefButtonRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefButtonRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBrowserViewRef Convert a pointer object to an existing class object
func AsCefBrowserViewRef(obj any) ICefBrowserViewRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBrowserViewRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPanelRef Convert a pointer object to an existing class object
func AsCefPanelRef(obj any) ICefPanelRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPanelRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefScrollViewRef Convert a pointer object to an existing class object
func AsCefScrollViewRef(obj any) ICefScrollViewRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefScrollViewRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefTextfieldRef Convert a pointer object to an existing class object
func AsCefTextfieldRef(obj any) ICefTextfieldRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefTextfieldRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefViewRef Convert a pointer object to an existing class object
func AsCefViewRef(obj any) ICefViewRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefViewRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBoxLayoutRef Convert a pointer object to an existing class object
func AsCefBoxLayoutRef(obj any) ICefBoxLayoutRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBoxLayoutRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefLayoutRef Convert a pointer object to an existing class object
func AsCefLayoutRef(obj any) ICefLayoutRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefLayoutRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefDisplayRef Convert a pointer object to an existing class object
func AsCefDisplayRef(obj any) ICefDisplayRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefDisplayRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefFillLayoutRef Convert a pointer object to an existing class object
func AsCefFillLayoutRef(obj any) ICefFillLayoutRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefFillLayoutRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPrintDialogCallbackRef Convert a pointer object to an existing class object
func AsCefPrintDialogCallbackRef(obj any) ICefPrintDialogCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPrintDialogCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPrintJobCallbackRef Convert a pointer object to an existing class object
func AsCefPrintJobCallbackRef(obj any) ICefPrintJobCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPrintJobCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefOverlayControllerRef Convert a pointer object to an existing class object
func AsCefOverlayControllerRef(obj any) ICefOverlayControllerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefOverlayControllerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefRunQuickMenuCallbackRef Convert a pointer object to an existing class object
func AsCefRunQuickMenuCallbackRef(obj any) ICefRunQuickMenuCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefRunQuickMenuCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMediaAccessCallbackRef Convert a pointer object to an existing class object
func AsCefMediaAccessCallbackRef(obj any) ICefMediaAccessCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMediaAccessCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPermissionPromptCallbackRef Convert a pointer object to an existing class object
func AsCefPermissionPromptCallbackRef(obj any) ICefPermissionPromptCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPermissionPromptCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSharedProcessMessageBuilderRef Convert a pointer object to an existing class object
func AsCefSharedProcessMessageBuilderRef(obj any) ICefSharedProcessMessageBuilderRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSharedProcessMessageBuilderRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefSharedMemoryRegionRef Convert a pointer object to an existing class object
func AsCefSharedMemoryRegionRef(obj any) ICefSharedMemoryRegionRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefSharedMemoryRegionRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefPreferenceManagerRef Convert a pointer object to an existing class object
func AsCefPreferenceManagerRef(obj any) ICefPreferenceManagerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefPreferenceManagerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefTaskManagerRef Convert a pointer object to an existing class object
func AsCefTaskManagerRef(obj any) ICefTaskManagerRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefTaskManagerRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefRegistrationRef Convert a pointer object to an existing class object
func AsCefRegistrationRef(obj any) ICefRegistrationRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefRegistrationRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomStringMap Convert a pointer object to an existing class object
func AsCefCustomStringMap(obj any) ICefCustomStringMap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomStringMap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomStringList Convert a pointer object to an existing class object
func AsCefCustomStringList(obj any) ICefCustomStringList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomStringList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefUnresponsiveProcessCallbackRef Convert a pointer object to an existing class object
func AsCefUnresponsiveProcessCallbackRef(obj any) ICefUnresponsiveProcessCallbackRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefUnresponsiveProcessCallbackRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefMenuButtonPressedLockRef Convert a pointer object to an existing class object
func AsCefMenuButtonPressedLockRef(obj any) ICefMenuButtonPressedLockRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefMenuButtonPressedLockRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBaseRefCountedRef Convert a pointer object to an existing class object
func AsCefBaseRefCountedRef(obj any) ICefBaseRefCountedRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBaseRefCountedRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefBaseRefCountedOwn Convert a pointer object to an existing class object
func AsCefBaseRefCountedOwn(obj any) ICefBaseRefCountedOwn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefBaseRefCountedOwn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefCustomUserData Convert a pointer object to an existing class object
func AsCefCustomUserData(obj any) ICefCustomUserData {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefCustomUserData)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefComponentRef Convert a pointer object to an existing class object
func AsCefComponentRef(obj any) ICefComponentRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefComponentRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCefv8BackingStoreRef Convert a pointer object to an existing class object
func AsCefv8BackingStoreRef(obj any) ICefv8BackingStoreRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCefv8BackingStoreRef)
	base.SetObjectInstance(result, instance)
	return result
}
