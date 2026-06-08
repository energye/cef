//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/109/types"
	. "github.com/energye/cef/base"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

var getPtr = GetPtr

func makeTConstrainedResizeEvent(cb TConstrainedResizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TConstrainedResizeEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);
		sender := lcl.AsObject(getVal(0))
		minWidth := (*types.TConstraintSize)(getPtr(getVal(1)))
		minHeight := (*types.TConstraintSize)(getPtr(getVal(2)))
		maxWidth := (*types.TConstraintSize)(getPtr(getVal(3)))
		maxHeight := (*types.TConstraintSize)(getPtr(getVal(4)))
		cb(sender, minWidth, minHeight, maxWidth, maxHeight)
	})
}

func makeTContextPopupEvent(cb TContextPopupEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TContextPopupEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
		sender := lcl.AsObject(getVal(0))
		mousePos := *(*types.TPoint)(getPtr(getVal(1)))
		handled := (*bool)(getPtr(getVal(2)))
		cb(sender, mousePos, handled)
	})
}

func makeTDragDropEvent(cb TDragDropEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TDragDropEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender, Source: TObject; X,Y: Integer);
		sender := lcl.AsObject(getVal(0))
		source := lcl.AsObject(getVal(1))
		X := int32(getVal(2))
		Y := int32(getVal(3))
		cb(sender, source, X, Y)
	})
}

func makeTDragOverEvent(cb TDragOverEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TDragOverEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender, Source: TObject; X,Y: Integer; State: TDragState; var Accept: Boolean);
		sender := lcl.AsObject(getVal(0))
		source := lcl.AsObject(getVal(1))
		X := int32(getVal(2))
		Y := int32(getVal(3))
		state := types.TDragState(getVal(4))
		accept := (*bool)(getPtr(getVal(5)))
		cb(sender, source, X, Y, state, accept)
	})
}

func makeTEndDragEvent(cb TEndDragEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TEndDragEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender, Target: TObject; X,Y: Integer);
		sender := lcl.AsObject(getVal(0))
		target := lcl.AsObject(getVal(1))
		X := int32(getVal(2))
		Y := int32(getVal(3))
		cb(sender, target, X, Y)
	})
}

func makeTGetSiteInfoEvent(cb TGetSiteInfoEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TGetSiteInfoEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
		sender := lcl.AsObject(getVal(0))
		dockClient := lcl.AsControl(getVal(1))
		influenceRect := (*types.TRect)(getPtr(getVal(2)))
		mousePos := *(*types.TPoint)(getPtr(getVal(3)))
		canDock := (*bool)(getPtr(getVal(4)))
		cb(sender, dockClient, influenceRect, mousePos, canDock)
	})
}

func makeTMouseEvent(cb TMouseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TMouseEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y: Integer);
		sender := lcl.AsObject(getVal(0))
		button := types.TMouseButton(getVal(1))
		shift := types.TShiftState(getVal(2))
		X := int32(getVal(3))
		Y := int32(getVal(4))
		cb(sender, button, shift, X, Y)
	})
}

func makeTMouseMoveEvent(cb TMouseMoveEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TMouseMoveEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; Shift: TShiftState; X, Y: Integer);
		sender := lcl.AsObject(getVal(0))
		shift := types.TShiftState(getVal(1))
		X := int32(getVal(2))
		Y := int32(getVal(3))
		cb(sender, shift, X, Y)
	})
}

func makeTMouseWheelEvent(cb TMouseWheelEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TMouseWheelEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; Shift: TShiftState; WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
		sender := lcl.AsObject(getVal(0))
		shift := types.TShiftState(getVal(1))
		wheelDelta := int32(getVal(2))
		mousePos := *(*types.TPoint)(getPtr(getVal(3)))
		handled := (*bool)(getPtr(getVal(4)))
		cb(sender, shift, wheelDelta, mousePos, handled)
	})
}

func makeTNotifyEvent(cb TNotifyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TNotifyEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(Sender: TObject);
		sender := lcl.AsObject(getVal(0))
		cb(sender)
	})
}

func makeTOnAcceleratedPaint(cb TOnAcceleratedPaint) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAcceleratedPaint", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; type_: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; shared_handle: Pointer);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		type_ := cefTypes.TCefPaintElementType(getVal(2))
		dirtyRectsCount := cefTypes.NativeUInt(getVal(3))
		dirtyRectsPtr := getVal(4)
		dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
		sharedHandle := uintptr(getVal(5))
		cb(sender, browser, type_, dirtyRectsCount, dirtyRects, sharedHandle)
	})
}

func makeTOnAcceleratorEvent(cb TOnAcceleratorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAcceleratorEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const window_: ICefWindow; command_id: Integer; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		commandId := int32(getVal(2))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, window, commandId, result)
	})
}

func makeTOnAccessibilityAccessibilityLocationChangeEvent(cb TOnAccessibilityAccessibilityLocationChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAccessibilityAccessibilityLocationChangeEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const value: ICefValue);
		value := AsCefValueRef(getVal(0))
		cb(value)
	})
}

func makeTOnAccessibilityAccessibilityTreeChangeEvent(cb TOnAccessibilityAccessibilityTreeChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAccessibilityAccessibilityTreeChangeEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const value: ICefValue);
		value := AsCefValueRef(getVal(0))
		cb(value)
	})
}

func makeTOnAccessibilityEvent(cb TOnAccessibilityEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAccessibilityEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const value: ICefValue);
		sender := lcl.AsObject(getVal(0))
		value := AsCefValueRef(getVal(1))
		cb(sender, value)
	})
}

func makeTOnAddressChange(cb TOnAddressChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAddressChange", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const url: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		url := api.GoStr(getVal(3))
		cb(sender, browser, frame, url)
	})
}

func makeTOnAfterCreated(cb TOnAfterCreated) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAfterCreated", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnAfterUserActionEvent(cb TOnAfterUserActionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAfterUserActionEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const textfield: ICefTextfield);
		sender := lcl.AsObject(getVal(0))
		textfield := AsCefTextfieldRef(getVal(1))
		cb(sender, textfield)
	})
}

func makeTOnAllowEvent(cb TOnAllowEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAllowEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; var allow : boolean);
		sender := lcl.AsObject(getVal(0))
		allow := (*bool)(getPtr(getVal(1)))
		cb(sender, allow)
	})
}

func makeTOnAppBeforeCommandLineProcessingEvent(cb TOnAppBeforeCommandLineProcessingEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAppBeforeCommandLineProcessingEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const processType: ustring; const commandLine: ICefCommandLine);
		processType := api.GoStr(getVal(0))
		commandLine := AsCefCommandLineRef(getVal(1))
		cb(processType, commandLine)
	})
}

func makeTOnAppGetBrowserProcessHandlerEvent(cb TOnAppGetBrowserProcessHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAppGetBrowserProcessHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefBrowserProcessHandler);
		var handler IEngBrowserProcessHandler
		handler = AsEngBrowserProcessHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnAppGetRenderProcessHandlerEvent(cb TOnAppGetRenderProcessHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAppGetRenderProcessHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefRenderProcessHandler);
		var handler IEngRenderProcessHandler
		handler = AsEngRenderProcessHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnAppGetResourceBundleHandlerEvent(cb TOnAppGetResourceBundleHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAppGetResourceBundleHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefResourceBundleHandler);
		var handler IEngResourceBundleHandler
		handler = AsEngResourceBundleHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnAppRegisterCustomSchemesEvent(cb TOnAppRegisterCustomSchemesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAppRegisterCustomSchemesEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const registrar: TCefSchemeRegistrarRef);
		registrar := AsCefSchemeRegistrarRef(getVal(0))
		cb(registrar)
	})
}

func makeTOnAudioAudioStreamErrorEvent(cb TOnAudioAudioStreamErrorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioAudioStreamErrorEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const message_: ustring);
		browser := AsCefBrowserRef(getVal(0))
		message := api.GoStr(getVal(1))
		cb(browser, message)
	})
}

func makeTOnAudioAudioStreamPacketEvent(cb TOnAudioAudioStreamPacketEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioAudioStreamPacketEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const data: PPSingle; frames: integer; pts: int64);
		browser := AsCefBrowserRef(getVal(0))
		data := PPSingle(getVal(1))
		frames := int32(getVal(2))
		pts := *(*int64)(getPtr(getVal(3)))
		cb(browser, data, frames, pts)
	})
}

func makeTOnAudioAudioStreamStartedEvent(cb TOnAudioAudioStreamStartedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioAudioStreamStartedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const params: TCefAudioParameters; channels: integer);
		browser := AsCefBrowserRef(getVal(0))
		params := *(*TCefAudioParameters)(getPtr(getVal(1)))
		channels := int32(getVal(2))
		cb(browser, params, channels)
	})
}

func makeTOnAudioAudioStreamStoppedEvent(cb TOnAudioAudioStreamStoppedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioAudioStreamStoppedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnAudioGetAudioParametersEvent(cb TOnAudioGetAudioParametersEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioGetAudioParametersEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; var params: TCefAudioParameters; var aResult: boolean);
		browser := AsCefBrowserRef(getVal(0))
		params := (*TCefAudioParameters)(getPtr(getVal(1)))
		result := (*bool)(getPtr(getVal(2)))
		cb(browser, params, result)
		*(*TCefAudioParameters)(getPtr(getVal(1))) = *params
	})
}

func makeTOnAudioStreamErrorEvent(cb TOnAudioStreamErrorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioStreamErrorEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		message := api.GoStr(getVal(2))
		cb(sender, browser, message)
	})
}

func makeTOnAudioStreamPacketEvent(cb TOnAudioStreamPacketEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioStreamPacketEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const data : PPSingle; frames: integer; pts: int64);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		data := PPSingle(getVal(2))
		frames := int32(getVal(3))
		pts := *(*int64)(getPtr(getVal(4)))
		cb(sender, browser, data, frames, pts)
	})
}

func makeTOnAudioStreamStartedEvent(cb TOnAudioStreamStartedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioStreamStartedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const params: TCefAudioParameters; channels: integer);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		params := *(*TCefAudioParameters)(getPtr(getVal(2)))
		channels := int32(getVal(3))
		cb(sender, browser, params, channels)
	})
}

func makeTOnAudioStreamStoppedEvent(cb TOnAudioStreamStoppedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAudioStreamStoppedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnAutoResize(cb TOnAutoResize) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnAutoResize", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const new_size: PCefSize; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		newSize := *(*TCefSize)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, newSize, result)
	})
}

func makeTOnBeforeBackgroundBrowserEvent(cb TOnBeforeBackgroundBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeBackgroundBrowserEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const extension: ICefExtension; const url: ustring; var client: ICefClient; var settings: TCefBrowserSettings; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		url := api.GoStr(getVal(2))
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(3))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(4)))
		settings := settingsPtr.ToGo()
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, extension, url, &client, &settings, result)
		if client != nil {
			*(*uintptr)(getPtr(getVal(3))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(4))) = *r
		}
	})
}

func makeTOnBeforeBrowse(cb TOnBeforeBrowse) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeBrowse", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; user_gesture, isRedirect: Boolean; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		userGesture := api.GoBool(getVal(4))
		isRedirect := api.GoBool(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, request, userGesture, isRedirect, result)
	})
}

func makeTOnBeforeBrowserEvent(cb TOnBeforeBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeBrowserEvent", func(getVal func(i int) uintptr) {
		// 11 : procedure(Sender: TObject; const extension: ICefExtension; const browser, active_browser: ICefBrowser; index: Integer; const url: ustring; active: boolean; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		browser := AsCefBrowserRef(getVal(2))
		activeBrowser := AsCefBrowserRef(getVal(3))
		index := int32(getVal(4))
		url := api.GoStr(getVal(5))
		active := api.GoBool(getVal(6))
		windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(7)))
		windowInfo := windowInfoPtr.ToGo()
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(8))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(9)))
		settings := settingsPtr.ToGo()
		result := (*bool)(getPtr(getVal(10)))
		cb(sender, extension, browser, activeBrowser, index, url, active, &windowInfo, &client, &settings, result)
		if r := windowInfo.ToPas(); r != nil {
			*(*tCefWindowInfo)(getPtr(getVal(7))) = *r
		}
		if client != nil {
			*(*uintptr)(getPtr(getVal(8))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(9))) = *r
		}
	})
}

func makeTOnBeforeChildProcessLaunchEvent(cb TOnBeforeChildProcessLaunchEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeChildProcessLaunchEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const commandLine: ICefCommandLine);
		commandLine := AsCefCommandLineRef(getVal(0))
		cb(commandLine)
	})
}

func makeTOnBeforeClose(cb TOnBeforeClose) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeClose", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnBeforeContextMenu(cb TOnBeforeContextMenu) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeContextMenu", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		params := AsCefContextMenuParamsRef(getVal(3))
		model := AsCefMenuModelRef(getVal(4))
		cb(sender, browser, frame, params, model)
	})
}

func makeTOnBeforeDownload(cb TOnBeforeDownload) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeDownload", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const suggestedName: ustring; const callback: ICefBeforeDownloadCallback);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		downloadItem := AsCefDownLoadItemRef(getVal(2))
		suggestedName := api.GoStr(getVal(3))
		callback := AsCefBeforeDownloadCallbackRef(getVal(4))
		cb(sender, browser, downloadItem, suggestedName, callback)
	})
}

func makeTOnBeforePopup(cb TOnBeforePopup) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforePopup", func(getVal func(i int) uintptr) {
		// 14 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const targetUrl, targetFrameName: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; const popupFeatures: TCefPopupFeatures; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var noJavascriptAccess: Boolean; var Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		targetUrl := api.GoStr(getVal(3))
		targetFrameName := api.GoStr(getVal(4))
		targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(5))
		userGesture := api.GoBool(getVal(6))
		popupFeatures := *(*TCefPopupFeatures)(getPtr(getVal(7)))
		windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(8)))
		windowInfo := windowInfoPtr.ToGo()
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(9))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(10)))
		settings := settingsPtr.ToGo()
		extraInfo := ICefDictionaryValue(AsCefDictionaryValueRef(*(*uintptr)(getPtr(getVal(11)))))
		noJavascriptAccess := (*bool)(getPtr(getVal(12)))
		result := (*bool)(getPtr(getVal(13)))
		cb(sender, browser, frame, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, &windowInfo, &client, &settings, &extraInfo, noJavascriptAccess, result)
		if r := windowInfo.ToPas(); r != nil {
			*(*tCefWindowInfo)(getPtr(getVal(8))) = *r
		}
		if client != nil {
			*(*uintptr)(getPtr(getVal(9))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(10))) = *r
		}
		if extraInfo != nil {
			*(*uintptr)(getPtr(getVal(11))) = extraInfo.Instance()
		}
	})
}

func makeTOnBeforeResourceLoad(cb TOnBeforeResourceLoad) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeResourceLoad", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const callback: ICefCallback; out Result: TCefReturnValue);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		callback := AsCefCallbackRef(getVal(4))
		result := (*cefTypes.TCefReturnValue)(getPtr(getVal(5)))
		cb(sender, browser, frame, request, callback, result)
	})
}

func makeTOnBeforeUnloadDialog(cb TOnBeforeUnloadDialog) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBeforeUnloadDialog", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const messageText: ustring; isReload: Boolean; const callback: ICefJsDialogCallback; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		messageText := api.GoStr(getVal(2))
		isReload := api.GoBool(getVal(3))
		callback := AsCefJsDialogCallbackRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, messageText, isReload, callback, result)
	})
}

func makeTOnBlurEvent(cb TOnBlurEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBlurEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const view: ICefView);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		cb(sender, view)
	})
}

func makeTOnBrowserCreatedEvent(cb TOnBrowserCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserCreatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const extra_info: ICefDictionaryValue);
		browser := AsCefBrowserRef(getVal(0))
		extraInfo := AsCefDictionaryValueRef(getVal(1))
		cb(browser, extraInfo)
	})
}

func makeTOnBrowserDestroyedEvent(cb TOnBrowserDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserDestroyedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnBrowserProcessBeforeChildProcessLaunchEvent(cb TOnBrowserProcessBeforeChildProcessLaunchEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserProcessBeforeChildProcessLaunchEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const commandLine: ICefCommandLine);
		commandLine := AsCefCommandLineRef(getVal(0))
		cb(commandLine)
	})
}

func makeTOnBrowserProcessContextInitializedEvent(cb TOnBrowserProcessContextInitializedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserProcessContextInitializedEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnBrowserProcessGetDefaultClientEvent(cb TOnBrowserProcessGetDefaultClientEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserProcessGetDefaultClientEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aClient: ICefClient);
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(0))))
		cb(&client)
		if client != nil {
			*(*uintptr)(getPtr(getVal(0))) = client.Instance()
		}
	})
}

func makeTOnBrowserProcessRegisterCustomPreferencesEvent(cb TOnBrowserProcessRegisterCustomPreferencesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserProcessRegisterCustomPreferencesEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(type_: TCefPreferencesType; registrar: PCefPreferenceRegistrar);
		type_ := cefTypes.TCefPreferencesType(getVal(0))
		registrar := *(*TCefPreferenceRegistrar)(getPtr(getVal(1)))
		cb(type_, registrar)
	})
}

func makeTOnBrowserProcessScheduleMessagePumpWorkEvent(cb TOnBrowserProcessScheduleMessagePumpWorkEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserProcessScheduleMessagePumpWorkEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const delayMs: Int64);
		delayMs := *(*int64)(getPtr(getVal(0)))
		cb(delayMs)
	})
}

func makeTOnBrowserViewBrowserCreatedEvent(cb TOnBrowserViewBrowserCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserViewBrowserCreatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser_view: ICefBrowserView; const browser: ICefBrowser);
		browserView := AsCefBrowserViewRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(browserView, browser)
	})
}

func makeTOnBrowserViewBrowserDestroyedEvent(cb TOnBrowserViewBrowserDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserViewBrowserDestroyedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser_view: ICefBrowserView; const browser: ICefBrowser);
		browserView := AsCefBrowserViewRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(browserView, browser)
	})
}

func makeTOnBrowserViewGetChromeToolbarTypeEvent(cb TOnBrowserViewGetChromeToolbarTypeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserViewGetChromeToolbarTypeEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): TCefChromeToolbarType;
		ret := cb()
		*(*TCefChromeToolbarType)(getPtr(getVal(0))) = ret
	})
}

func makeTOnBrowserViewGetDelegateForPopupBrowserViewEvent(cb TOnBrowserViewGetDelegateForPopupBrowserViewEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserViewGetDelegateForPopupBrowserViewEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser_view: ICefBrowserView; const settings: TCefBrowserSettings; const client: ICefClient; is_devtools: boolean; var aResult: ICefBrowserViewDelegate);
		browserView := AsCefBrowserViewRef(getVal(0))
		settingsPtr := (*tCefBrowserSettings)(getPtr(getVal(1)))
		settings := settingsPtr.ToGo()
		client := AsEngClient(getVal(2))
		isDevtools := api.GoBool(getVal(3))
		var result IEngBrowserViewDelegate
		result = AsEngBrowserViewDelegate(*(*uintptr)(getPtr(getVal(4))))
		cb(browserView, settings, client, isDevtools, &result)
		if result != nil {
			*(*uintptr)(getPtr(getVal(4))) = result.Instance()
		}
	})
}

func makeTOnBrowserViewPopupBrowserViewCreatedEvent(cb TOnBrowserViewPopupBrowserViewCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnBrowserViewPopupBrowserViewCreatedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser_view: ICefBrowserView; const popup_browser_view: ICefBrowserView; is_devtools: boolean; var aResult: boolean);
		browserView := AsCefBrowserViewRef(getVal(0))
		popupBrowserView := AsCefBrowserViewRef(getVal(1))
		isDevtools := api.GoBool(getVal(2))
		result := (*bool)(getPtr(getVal(3)))
		cb(browserView, popupBrowserView, isDevtools, result)
	})
}

func makeTOnButtonButtonPressedEvent(cb TOnButtonButtonPressedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnButtonButtonPressedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const button: ICefButton);
		button := AsCefButtonRef(getVal(0))
		cb(button)
	})
}

func makeTOnButtonButtonStateChangedEvent(cb TOnButtonButtonStateChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnButtonButtonStateChangedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const button: ICefButton);
		button := AsCefButtonRef(getVal(0))
		cb(button)
	})
}

func makeTOnButtonPressedEvent(cb TOnButtonPressedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnButtonPressedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const button: ICefButton);
		sender := lcl.AsObject(getVal(0))
		button := AsCefButtonRef(getVal(1))
		cb(sender, button)
	})
}

func makeTOnButtonStateChangedEvent(cb TOnButtonStateChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnButtonStateChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const button: ICefButton);
		sender := lcl.AsObject(getVal(0))
		button := AsCefButtonRef(getVal(1))
		cb(sender, button)
	})
}

func makeTOnCanAccessBrowserEvent(cb TOnCanAccessBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanAccessBrowserEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const extension: ICefExtension; const browser: ICefBrowser; include_incognito: boolean; const target_browser: ICefBrowser; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		browser := AsCefBrowserRef(getVal(2))
		includeIncognito := api.GoBool(getVal(3))
		targetBrowser := AsCefBrowserRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, extension, browser, includeIncognito, targetBrowser, result)
	})
}

func makeTOnCanCloseEvent(cb TOnCanCloseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanCloseEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnCanDownloadEvent(cb TOnCanDownloadEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanDownloadEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const url, request_method: ustring; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		url := api.GoStr(getVal(2))
		requestMethod := api.GoStr(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, url, requestMethod, result)
	})
}

func makeTOnCanMaximizeEvent(cb TOnCanMaximizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanMaximizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnCanMinimizeEvent(cb TOnCanMinimizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanMinimizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnCanResizeEvent(cb TOnCanResizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanResizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnCanSaveCookie(cb TOnCanSaveCookie) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanSaveCookie", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; const cookie: PCefCookie; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		response := AsCefResponseRef(getVal(4))
		cookiePtr := (*tCefCookie)(getPtr(getVal(5)))
		cookie := cookiePtr.ToGo()
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, request, response, cookie, result)
	})
}

func makeTOnCanSendCookie(cb TOnCanSendCookie) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCanSendCookie", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const cookie: PCefCookie; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		cookiePtr := (*tCefCookie)(getPtr(getVal(4)))
		cookie := cookiePtr.ToGo()
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, frame, request, cookie, result)
	})
}

func makeTOnCertificateError(cb TOnCertificateError) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCertificateError", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; certError: TCefErrorcode; const requestUrl: ustring; const sslInfo: ICefSslInfo; const callback: ICefCallback; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		certError := cefTypes.TCefErrorCode(getVal(2))
		requestUrl := api.GoStr(getVal(3))
		sslInfo := AsCefSslInfoRef(getVal(4))
		callback := AsCefCallbackRef(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, certError, requestUrl, sslInfo, callback, result)
	})
}

func makeTOnChildViewChangedEvent(cb TOnChildViewChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnChildViewChangedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const view: ICefView; added: boolean; const child: ICefView);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		added := api.GoBool(getVal(2))
		child := AsCefViewRef(getVal(3))
		cb(sender, view, added, child)
	})
}

func makeTOnChromeCommandEvent(cb TOnChromeCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnChromeCommandEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; command_id: integer; disposition: TCefWindowOpenDisposition; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		commandId := int32(getVal(2))
		disposition := cefTypes.TCefWindowOpenDisposition(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, commandId, disposition, result)
	})
}

func makeTOnClientConnected(cb TOnClientConnected) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientConnected", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		cb(sender, server, connectionId)
	})
}

func makeTOnClientDisconnected(cb TOnClientDisconnected) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientDisconnected", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		cb(sender, server, connectionId)
	})
}

func makeTOnClientGetAudioHandlerEvent(cb TOnClientGetAudioHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetAudioHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefAudioHandler);
		var handler IEngAudioHandler
		handler = AsEngAudioHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetCommandHandlerEvent(cb TOnClientGetCommandHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetCommandHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefCommandHandler);
		var handler IEngCommandHandler
		handler = AsEngCommandHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetContextMenuHandlerEvent(cb TOnClientGetContextMenuHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetContextMenuHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefContextMenuHandler);
		var handler IEngContextMenuHandler
		handler = AsEngContextMenuHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetDialogHandlerEvent(cb TOnClientGetDialogHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetDialogHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefDialogHandler);
		var handler IEngDialogHandler
		handler = AsEngDialogHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetDisplayHandlerEvent(cb TOnClientGetDisplayHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetDisplayHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefDisplayHandler);
		var handler IEngDisplayHandler
		handler = AsEngDisplayHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetDownloadHandlerEvent(cb TOnClientGetDownloadHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetDownloadHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefDownloadHandler);
		var handler IEngDownloadHandler
		handler = AsEngDownloadHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetDragHandlerEvent(cb TOnClientGetDragHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetDragHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefDragHandler);
		var handler IEngDragHandler
		handler = AsEngDragHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetFindHandlerEvent(cb TOnClientGetFindHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetFindHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefFindHandler);
		var handler IEngFindHandler
		handler = AsEngFindHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetFocusHandlerEvent(cb TOnClientGetFocusHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetFocusHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefFocusHandler);
		var handler IEngFocusHandler
		handler = AsEngFocusHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetFrameHandlerEvent(cb TOnClientGetFrameHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetFrameHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefFrameHandler);
		var handler IEngFrameHandler
		handler = AsEngFrameHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetJsdialogHandlerEvent(cb TOnClientGetJsdialogHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetJsdialogHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefJsdialogHandler);
		var handler IEngJsDialogHandler
		handler = AsEngJsDialogHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetKeyboardHandlerEvent(cb TOnClientGetKeyboardHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetKeyboardHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefKeyboardHandler);
		var handler IEngKeyboardHandler
		handler = AsEngKeyboardHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetLifeSpanHandlerEvent(cb TOnClientGetLifeSpanHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetLifeSpanHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefLifeSpanHandler);
		var handler IEngLifeSpanHandler
		handler = AsEngLifeSpanHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetLoadHandlerEvent(cb TOnClientGetLoadHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetLoadHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefLoadHandler);
		var handler IEngLoadHandler
		handler = AsEngLoadHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetPermissionHandlerEvent(cb TOnClientGetPermissionHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetPermissionHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefPermissionHandler);
		var handler IEngPermissionHandler
		handler = AsEngPermissionHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetPrintHandlerEvent(cb TOnClientGetPrintHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetPrintHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefPrintHandler);
		var handler IEngPrintHandler
		handler = AsEngPrintHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetRenderHandlerEvent(cb TOnClientGetRenderHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetRenderHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefRenderHandler);
		var handler IEngRenderHandler
		handler = AsEngRenderHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientGetRequestHandlerEvent(cb TOnClientGetRequestHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientGetRequestHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aHandler: ICefRequestHandler);
		var handler IEngRequestHandler
		handler = AsEngRequestHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&handler)
		if handler != nil {
			*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
		}
	})
}

func makeTOnClientProcessMessageReceivedEvent(cb TOnClientProcessMessageReceivedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClientProcessMessageReceivedEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message_: ICefProcessMessage): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		sourceProcess := cefTypes.TCefProcessId(getVal(2))
		message := AsCefProcessMessageRef(getVal(3))
		ret := cb(browser, frame, sourceProcess, message)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnClose(cb TOnClose) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnClose", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; var aAction : TCefCloseBrowserAction);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		action := (*cefTypes.TCefCloseBrowserAction)(getPtr(getVal(2)))
		cb(sender, browser, action)
	})
}

func makeTOnCommandChromeCommandEvent(cb TOnCommandChromeCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCommandChromeCommandEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const browser: ICefBrowser; command_id: integer; disposition: TCefWindowOpenDisposition): boolean;
		browser := AsCefBrowserRef(getVal(0))
		commandId := int32(getVal(1))
		disposition := cefTypes.TCefWindowOpenDisposition(getVal(2))
		ret := cb(browser, commandId, disposition)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnCompMsgEvent(cb TOnCompMsgEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCompMsgEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; var aMessage: TMessage; var aHandled: Boolean);
		sender := lcl.AsObject(getVal(0))
		message := (*types.TMessage)(getPtr(getVal(1)))
		handled := (*bool)(getPtr(getVal(2)))
		cb(sender, message, handled)
	})
}

func makeTOnCompletionCallbackCompleteEvent(cb TOnCompletionCallbackCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCompletionCallbackCompleteEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnConsoleMessage(cb TOnConsoleMessage) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnConsoleMessage", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; level: TCefLogSeverity; const message, source: ustring; line: Integer; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		level := cefTypes.TCefLogSeverity(getVal(2))
		message := api.GoStr(getVal(3))
		source := api.GoStr(getVal(4))
		line := int32(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, level, message, source, line, result)
	})
}

func makeTOnContextCreatedEvent(cb TOnContextCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextCreatedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		cb(browser, frame, context)
	})
}

func makeTOnContextInitializedEvent(cb TOnContextInitializedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextInitializedEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnContextMenuBeforeContextMenuEvent(cb TOnContextMenuBeforeContextMenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuBeforeContextMenuEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		params := AsCefContextMenuParamsRef(getVal(2))
		model := AsCefMenuModelRef(getVal(3))
		cb(browser, frame, params, model)
	})
}

func makeTOnContextMenuCommand(cb TOnContextMenuCommand) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuCommand", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; commandId: Integer; eventFlags: TCefEventFlags; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		params := AsCefContextMenuParamsRef(getVal(3))
		commandId := int32(getVal(4))
		eventFlags := cefTypes.TCefEventFlags(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, params, commandId, eventFlags, result)
	})
}

func makeTOnContextMenuContextMenuCommandEvent(cb TOnContextMenuContextMenuCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuContextMenuCommandEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; commandId: Integer; eventFlags: TCefEventFlags): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		params := AsCefContextMenuParamsRef(getVal(2))
		commandId := int32(getVal(3))
		eventFlags := cefTypes.TCefEventFlags(getVal(4))
		ret := cb(browser, frame, params, commandId, eventFlags)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnContextMenuContextMenuDismissedEvent(cb TOnContextMenuContextMenuDismissedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuContextMenuDismissedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		cb(browser, frame)
	})
}

func makeTOnContextMenuDismissed(cb TOnContextMenuDismissed) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuDismissed", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		cb(sender, browser, frame)
	})
}

func makeTOnContextMenuQuickMenuCommandEvent(cb TOnContextMenuQuickMenuCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuQuickMenuCommandEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; command_id: integer; event_flags: TCefEventFlags): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		commandId := int32(getVal(2))
		eventFlags := cefTypes.TCefEventFlags(getVal(3))
		ret := cb(browser, frame, commandId, eventFlags)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnContextMenuQuickMenuDismissedEvent(cb TOnContextMenuQuickMenuDismissedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuQuickMenuDismissedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		cb(browser, frame)
	})
}

func makeTOnContextMenuRunContextMenuEvent(cb TOnContextMenuRunContextMenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuRunContextMenuEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel; const callback: ICefRunContextMenuCallback): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		params := AsCefContextMenuParamsRef(getVal(2))
		model := AsCefMenuModelRef(getVal(3))
		callback := AsCefRunContextMenuCallbackRef(getVal(4))
		ret := cb(browser, frame, params, model, callback)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnContextMenuRunQuickMenuEvent(cb TOnContextMenuRunQuickMenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextMenuRunQuickMenuEvent", func(getVal func(i int) uintptr) {
		// 6 : function(const browser: ICefBrowser; const frame: ICefFrame; location: PCefPoint; size: PCefSize; edit_state_flags: TCefQuickMenuEditStateFlags; const callback: ICefRunQuickMenuCallback): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		location := *(*TCefPoint)(getPtr(getVal(2)))
		size := *(*TCefSize)(getPtr(getVal(3)))
		editStateFlags := cefTypes.TCefQuickMenuEditStateFlags(getVal(4))
		callback := AsCefRunQuickMenuCallbackRef(getVal(5))
		ret := cb(browser, frame, location, size, editStateFlags, callback)
		*(*bool)(getPtr(getVal(6))) = ret
	})
}

func makeTOnContextReleasedEvent(cb TOnContextReleasedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnContextReleasedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		cb(browser, frame, context)
	})
}

func makeTOnCookieAccessFilterCanSaveCookieEvent(cb TOnCookieAccessFilterCanSaveCookieEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookieAccessFilterCanSaveCookieEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; const cookie: PCefCookie): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		response := AsCefResponseRef(getVal(3))
		cookiePtr := (*tCefCookie)(getPtr(getVal(4)))
		cookie := cookiePtr.ToGo()
		ret := cb(browser, frame, request, response, cookie)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnCookieAccessFilterCanSendCookieEvent(cb TOnCookieAccessFilterCanSendCookieEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookieAccessFilterCanSendCookieEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const cookie: PCefCookie): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		cookiePtr := (*tCefCookie)(getPtr(getVal(3)))
		cookie := cookiePtr.ToGo()
		ret := cb(browser, frame, request, cookie)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnCookieSet(cb TOnCookieSet) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookieSet", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; aSuccess : boolean; aID : integer);
		sender := lcl.AsObject(getVal(0))
		success := api.GoBool(getVal(1))
		iD := int32(getVal(2))
		cb(sender, success, iD)
	})
}

func makeTOnCookieVisitorDestroyed(cb TOnCookieVisitorDestroyed) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookieVisitorDestroyed", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; aID : integer);
		sender := lcl.AsObject(getVal(0))
		iD := int32(getVal(1))
		cb(sender, iD)
	})
}

func makeTOnCookieVisitorVisitEvent(cb TOnCookieVisitorVisitEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookieVisitorVisitEvent", func(getVal func(i int) uintptr) {
		// 15 : function(const name: ustring; const value: ustring; const domain: ustring; const path: ustring; secure: Boolean; httponly: Boolean; hasExpires: Boolean; const creation: TDateTime; const lastAccess: TDateTime; const expires: TDateTime; count: Integer; total: Integer; same_site: TCefCookieSameSite; priority: TCefCookiePriority; out deleteCookie: Boolean): Boolean;
		name := api.GoStr(getVal(0))
		value := api.GoStr(getVal(1))
		domain := api.GoStr(getVal(2))
		path := api.GoStr(getVal(3))
		secure := api.GoBool(getVal(4))
		httponly := api.GoBool(getVal(5))
		hasExpires := api.GoBool(getVal(6))
		creation := types.TDateTime(getVal(7))
		lastAccess := types.TDateTime(getVal(8))
		expires := types.TDateTime(getVal(9))
		count := int32(getVal(10))
		total := int32(getVal(11))
		sameSite := cefTypes.TCefCookieSameSite(getVal(12))
		priority := cefTypes.TCefCookiePriority(getVal(13))
		deleteCookie := (*bool)(getPtr(getVal(14)))
		ret := cb(name, value, domain, path, secure, httponly, hasExpires, creation, lastAccess, expires, count, total, sameSite, priority, deleteCookie)
		*(*bool)(getPtr(getVal(15))) = ret
	})
}

func makeTOnCookiesDeletedEvent(cb TOnCookiesDeletedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookiesDeletedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; numDeleted : integer);
		sender := lcl.AsObject(getVal(0))
		numDeleted := int32(getVal(1))
		cb(sender, numDeleted)
	})
}

func makeTOnCookiesVisited(cb TOnCookiesVisited) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCookiesVisited", func(getVal func(i int) uintptr) {
		// 18 : procedure(Sender: TObject; const name_, value, domain, path: ustring; secure, httponly, hasExpires: Boolean; const creation, lastAccess, expires: TDateTime; count, total, aID : Integer; same_site : TCefCookieSameSite; priority : TCefCookiePriority; var aDeleteCookie, aResult : Boolean);
		sender := lcl.AsObject(getVal(0))
		name := api.GoStr(getVal(1))
		value := api.GoStr(getVal(2))
		domain := api.GoStr(getVal(3))
		path := api.GoStr(getVal(4))
		secure := api.GoBool(getVal(5))
		httponly := api.GoBool(getVal(6))
		hasExpires := api.GoBool(getVal(7))
		creation := types.TDateTime(getVal(8))
		lastAccess := types.TDateTime(getVal(9))
		expires := types.TDateTime(getVal(10))
		count := int32(getVal(11))
		total := int32(getVal(12))
		iD := int32(getVal(13))
		sameSite := cefTypes.TCefCookieSameSite(getVal(14))
		priority := cefTypes.TCefCookiePriority(getVal(15))
		deleteCookie := (*bool)(getPtr(getVal(16)))
		result := (*bool)(getPtr(getVal(17)))
		cb(sender, name, value, domain, path, secure, httponly, hasExpires, creation, lastAccess, expires, count, total, iD, sameSite, priority, deleteCookie, result)
	})
}

func makeTOnCursorChange(cb TOnCursorChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnCursorChange", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; cursor_: TCefCursorHandle; cursorType: TCefCursorType; const customCursorInfo: PCefCursorInfo; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cursor := cefTypes.TCefCursorHandle(getVal(2))
		cursorType := cefTypes.TCefCursorType(getVal(3))
		customCursorInfo := *(*TCefCursorInfo)(getPtr(getVal(4)))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, cursor, cursorType, customCursorInfo, result)
	})
}

func makeTOnDeleteCookiesCallbackCompleteEvent(cb TOnDeleteCookiesCallbackCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDeleteCookiesCallbackCompleteEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(numDeleted: Integer);
		numDeleted := int32(getVal(0))
		cb(numDeleted)
	})
}

func makeTOnDevToolsAgentAttachedEvent(cb TOnDevToolsAgentAttachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsAgentAttachedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnDevToolsAgentDetachedEvent(cb TOnDevToolsAgentDetachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsAgentDetachedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnDevToolsEventEvent(cb TOnDevToolsEventEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsEventEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const method: ustring; const params: ICefValue);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		method := api.GoStr(getVal(2))
		params := AsCefValueRef(getVal(3))
		cb(sender, browser, method, params)
	})
}

func makeTOnDevToolsEventRawEvent(cb TOnDevToolsEventRawEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsEventRawEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const method: ustring; const params: Pointer; params_size: NativeUInt);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		method := api.GoStr(getVal(2))
		params := uintptr(getVal(3))
		paramsSize := cefTypes.NativeUInt(getVal(4))
		cb(sender, browser, method, params, paramsSize)
	})
}

func makeTOnDevToolsMessageEvent(cb TOnDevToolsMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: ICefValue; var aHandled: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		message := AsCefValueRef(getVal(2))
		handled := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, message, handled)
	})
}

func makeTOnDevToolsMessageObserverDevToolsAgentAttachedEvent(cb TOnDevToolsMessageObserverDevToolsAgentAttachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageObserverDevToolsAgentAttachedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnDevToolsMessageObserverDevToolsAgentDetachedEvent(cb TOnDevToolsMessageObserverDevToolsAgentDetachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageObserverDevToolsAgentDetachedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnDevToolsMessageObserverDevToolsEvent(cb TOnDevToolsMessageObserverDevToolsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageObserverDevToolsEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const method: ustring; const params: Pointer; params_size: NativeUInt);
		browser := AsCefBrowserRef(getVal(0))
		method := api.GoStr(getVal(1))
		params := uintptr(getVal(2))
		paramsSize := cefTypes.NativeUInt(getVal(3))
		cb(browser, method, params, paramsSize)
	})
}

func makeTOnDevToolsMessageObserverDevToolsMessageEvent(cb TOnDevToolsMessageObserverDevToolsMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageObserverDevToolsMessageEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const message_: Pointer; message_size: NativeUInt; var aHandled: boolean);
		browser := AsCefBrowserRef(getVal(0))
		message := uintptr(getVal(1))
		messageSize := cefTypes.NativeUInt(getVal(2))
		handled := (*bool)(getPtr(getVal(3)))
		cb(browser, message, messageSize, handled)
	})
}

func makeTOnDevToolsMessageObserverDevToolsMethodResultEvent(cb TOnDevToolsMessageObserverDevToolsMethodResultEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMessageObserverDevToolsMethodResultEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; message_id: integer; success: boolean; const result: Pointer; result_size: NativeUInt);
		browser := AsCefBrowserRef(getVal(0))
		messageId := int32(getVal(1))
		success := api.GoBool(getVal(2))
		result := uintptr(getVal(3))
		resultSize := cefTypes.NativeUInt(getVal(4))
		cb(browser, messageId, success, result, resultSize)
	})
}

func makeTOnDevToolsMethodRawResultEvent(cb TOnDevToolsMethodRawResultEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMethodRawResultEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; message_id: integer; success: boolean; const result: Pointer; result_size: NativeUInt);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		messageId := int32(getVal(2))
		success := api.GoBool(getVal(3))
		result := uintptr(getVal(4))
		resultSize := cefTypes.NativeUInt(getVal(5))
		cb(sender, browser, messageId, success, result, resultSize)
	})
}

func makeTOnDevToolsMethodResultEvent(cb TOnDevToolsMethodResultEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsMethodResultEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; message_id: integer; success: boolean; const result: ICefValue);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		messageId := int32(getVal(2))
		success := api.GoBool(getVal(3))
		result := AsCefValueRef(getVal(4))
		cb(sender, browser, messageId, success, result)
	})
}

func makeTOnDevToolsRawMessageEvent(cb TOnDevToolsRawMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDevToolsRawMessageEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: Pointer; message_size: NativeUInt; var aHandled: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		message := uintptr(getVal(2))
		messageSize := cefTypes.NativeUInt(getVal(3))
		handled := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, message, messageSize, handled)
	})
}

func makeTOnDialogClosed(cb TOnDialogClosed) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDialogClosed", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnDialogFileDialogEvent(cb TOnDialogFileDialogEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDialogFileDialogEvent", func(getVal func(i int) uintptr) {
		// 6 : function(const browser: ICefBrowser; mode: TCefFileDialogMode; const title: ustring; const defaultFilePath: ustring; const acceptFilters: TStrings; const callback: ICefFileDialogCallback): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		mode := cefTypes.TCefFileDialogMode(getVal(1))
		title := api.GoStr(getVal(2))
		defaultFilePath := api.GoStr(getVal(3))
		acceptFilters := lcl.AsStrings(getVal(4))
		callback := AsCefFileDialogCallbackRef(getVal(5))
		ret := cb(browser, mode, title, defaultFilePath, acceptFilters, callback)
		*(*bool)(getPtr(getVal(6))) = ret
	})
}

func makeTOnDismissPermissionPromptEvent(cb TOnDismissPermissionPromptEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDismissPermissionPromptEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; prompt_id: uint64; result: TCefPermissionRequestResult);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		promptId := *(*uint64)(getPtr(getVal(2)))
		result := cefTypes.TCefPermissionRequestResult(getVal(3))
		cb(sender, browser, promptId, result)
	})
}

func makeTOnDisplayAddressChangeEvent(cb TOnDisplayAddressChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayAddressChangeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const url: ustring);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		url := api.GoStr(getVal(2))
		cb(browser, frame, url)
	})
}

func makeTOnDisplayAutoResizeEvent(cb TOnDisplayAutoResizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayAutoResizeEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const browser: ICefBrowser; const new_size: PCefSize): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		newSize := *(*TCefSize)(getPtr(getVal(1)))
		ret := cb(browser, newSize)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnDisplayConsoleMessageEvent(cb TOnDisplayConsoleMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayConsoleMessageEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; level: TCefLogSeverity; const message_: ustring; const source: ustring; line: Integer): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		level := cefTypes.TCefLogSeverity(getVal(1))
		message := api.GoStr(getVal(2))
		source := api.GoStr(getVal(3))
		line := int32(getVal(4))
		ret := cb(browser, level, message, source, line)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnDisplayCursorChangeEvent(cb TOnDisplayCursorChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayCursorChangeEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; cursor_: TCefCursorHandle; CursorType: TCefCursorType; const customCursorInfo: PCefCursorInfo; var aResult: boolean);
		browser := AsCefBrowserRef(getVal(0))
		cursor := cefTypes.TCefCursorHandle(getVal(1))
		cursorType := cefTypes.TCefCursorType(getVal(2))
		customCursorInfo := *(*TCefCursorInfo)(getPtr(getVal(3)))
		result := (*bool)(getPtr(getVal(4)))
		cb(browser, cursor, cursorType, customCursorInfo, result)
	})
}

func makeTOnDisplayFaviconUrlChangeEvent(cb TOnDisplayFaviconUrlChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayFaviconUrlChangeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const iconUrls: TStrings);
		browser := AsCefBrowserRef(getVal(0))
		iconUrls := lcl.AsStrings(getVal(1))
		cb(browser, iconUrls)
	})
}

func makeTOnDisplayFullScreenModeChangeEvent(cb TOnDisplayFullScreenModeChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayFullScreenModeChangeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; fullscreen: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		fullscreen := api.GoBool(getVal(1))
		cb(browser, fullscreen)
	})
}

func makeTOnDisplayLoadingProgressChangeEvent(cb TOnDisplayLoadingProgressChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayLoadingProgressChangeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const progress: double);
		browser := AsCefBrowserRef(getVal(0))
		progress := *(*float64)(getPtr(getVal(1)))
		cb(browser, progress)
	})
}

func makeTOnDisplayMediaAccessChangeEvent(cb TOnDisplayMediaAccessChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayMediaAccessChangeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; has_video_access: boolean; has_audio_access: boolean);
		browser := AsCefBrowserRef(getVal(0))
		hasVideoAccess := api.GoBool(getVal(1))
		hasAudioAccess := api.GoBool(getVal(2))
		cb(browser, hasVideoAccess, hasAudioAccess)
	})
}

func makeTOnDisplayStatusMessageEvent(cb TOnDisplayStatusMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayStatusMessageEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const value: ustring);
		browser := AsCefBrowserRef(getVal(0))
		value := api.GoStr(getVal(1))
		cb(browser, value)
	})
}

func makeTOnDisplayTitleChangeEvent(cb TOnDisplayTitleChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayTitleChangeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const title: ustring);
		browser := AsCefBrowserRef(getVal(0))
		title := api.GoStr(getVal(1))
		cb(browser, title)
	})
}

func makeTOnDisplayTooltipEvent(cb TOnDisplayTooltipEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDisplayTooltipEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const browser: ICefBrowser; var text: ustring): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		text := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
		ret := cb(browser, &text)
		*(*uintptr)(getPtr(getVal(1))) = api.PasStr(text)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnDocumentAvailableInMainFrame(cb TOnDocumentAvailableInMainFrame) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDocumentAvailableInMainFrame", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: Tobject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnDomVisitorVisitEvent(cb TOnDomVisitorVisitEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDomVisitorVisitEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const document: ICefDomDocument);
		document := AsCefDomDocumentRef(getVal(0))
		cb(document)
	})
}

func makeTOnDownloadBeforeDownloadEvent(cb TOnDownloadBeforeDownloadEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadBeforeDownloadEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const suggestedName: ustring; const callback: ICefBeforeDownloadCallback);
		browser := AsCefBrowserRef(getVal(0))
		downloadItem := AsCefDownLoadItemRef(getVal(1))
		suggestedName := api.GoStr(getVal(2))
		callback := AsCefBeforeDownloadCallbackRef(getVal(3))
		cb(browser, downloadItem, suggestedName, callback)
	})
}

func makeTOnDownloadCanDownloadEvent(cb TOnDownloadCanDownloadEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadCanDownloadEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const browser: ICefBrowser; const url: ustring; const request_method: ustring): boolean;
		browser := AsCefBrowserRef(getVal(0))
		url := api.GoStr(getVal(1))
		requestMethod := api.GoStr(getVal(2))
		ret := cb(browser, url, requestMethod)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnDownloadData(cb TOnDownloadData) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadData", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; data: Pointer; dataLength: NativeUInt);
		sender := lcl.AsObject(getVal(0))
		request := AsCefUrlRequestRef(getVal(1))
		data := uintptr(getVal(2))
		dataLength := cefTypes.NativeUInt(getVal(3))
		cb(sender, request, data, dataLength)
	})
}

func makeTOnDownloadDownloadUpdatedEvent(cb TOnDownloadDownloadUpdatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadDownloadUpdatedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const callback: ICefDownloadItemCallback);
		browser := AsCefBrowserRef(getVal(0))
		downloadItem := AsCefDownLoadItemRef(getVal(1))
		callback := AsCefDownloadItemCallbackRef(getVal(2))
		cb(browser, downloadItem, callback)
	})
}

func makeTOnDownloadImageCallbackDownloadImageFinishedEvent(cb TOnDownloadImageCallbackDownloadImageFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadImageCallbackDownloadImageFinishedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const imageUrl: ustring; httpStatusCode: Integer; const image: ICefImage);
		imageUrl := api.GoStr(getVal(0))
		httpStatusCode := int32(getVal(1))
		image := AsCefImageRef(getVal(2))
		cb(imageUrl, httpStatusCode, image)
	})
}

func makeTOnDownloadImageFinishedEvent(cb TOnDownloadImageFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadImageFinishedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const imageUrl: ustring; httpStatusCode: Integer; const image: ICefImage);
		sender := lcl.AsObject(getVal(0))
		imageUrl := api.GoStr(getVal(1))
		httpStatusCode := int32(getVal(2))
		image := AsCefImageRef(getVal(3))
		cb(sender, imageUrl, httpStatusCode, image)
	})
}

func makeTOnDownloadProgress(cb TOnDownloadProgress) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadProgress", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; current, total: Int64);
		sender := lcl.AsObject(getVal(0))
		request := AsCefUrlRequestRef(getVal(1))
		current := *(*int64)(getPtr(getVal(2)))
		total := *(*int64)(getPtr(getVal(3)))
		cb(sender, request, current, total)
	})
}

func makeTOnDownloadUpdated(cb TOnDownloadUpdated) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDownloadUpdated", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const callback: ICefDownloadItemCallback);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		downloadItem := AsCefDownLoadItemRef(getVal(2))
		callback := AsCefDownloadItemCallbackRef(getVal(3))
		cb(sender, browser, downloadItem, callback)
	})
}

func makeTOnDragDragEnterEvent(cb TOnDragDragEnterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDragDragEnterEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const browser: ICefBrowser; const dragData: ICefDragData; mask: TCefDragOperations): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		dragData := AsCefDragDataRef(getVal(1))
		mask := cefTypes.TCefDragOperations(getVal(2))
		ret := cb(browser, dragData, mask)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnDragDraggableRegionsChangedEvent(cb TOnDragDraggableRegionsChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDragDraggableRegionsChangedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; regionsCount: NativeUInt; const regions: PCefDraggableRegionArray);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		regionsCount := cefTypes.NativeUInt(getVal(2))
		regionsPtr := getVal(3)
		regions := NewCefDraggableRegionArray(int(regionsCount), regionsPtr)
		cb(browser, frame, regionsCount, regions)
	})
}

func makeTOnDragEnter(cb TOnDragEnter) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDragEnter", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const dragData: ICefDragData; mask: TCefDragOperations; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		dragData := AsCefDragDataRef(getVal(2))
		mask := cefTypes.TCefDragOperations(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, dragData, mask, result)
	})
}

func makeTOnDraggableRegionsChanged(cb TOnDraggableRegionsChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnDraggableRegionsChanged", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; regionsCount: NativeUInt; const regions: PCefDraggableRegionArray);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		regionsCount := cefTypes.NativeUInt(getVal(3))
		regionsPtr := getVal(4)
		regions := NewCefDraggableRegionArray(int(regionsCount), regionsPtr)
		cb(sender, browser, frame, regionsCount, regions)
	})
}

func makeTOnEndTracingCallbackEndTracingCompleteEvent(cb TOnEndTracingCallbackEndTracingCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnEndTracingCallbackEndTracingCompleteEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const tracingFile: ustring);
		tracingFile := api.GoStr(getVal(0))
		cb(tracingFile)
	})
}

func makeTOnExecuteTaskOnCefThread(cb TOnExecuteTaskOnCefThread) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExecuteTaskOnCefThread", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; aTaskID : cardinal);
		sender := lcl.AsObject(getVal(0))
		taskID := uint32(getVal(1))
		cb(sender, taskID)
	})
}

func makeTOnExtensionBeforeBackgroundBrowserEvent(cb TOnExtensionBeforeBackgroundBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionBeforeBackgroundBrowserEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const extension: ICefExtension; const url: ustring; var client: ICefClient; var settings: TCefBrowserSettings): boolean;
		extension := AsCefExtensionRef(getVal(0))
		url := api.GoStr(getVal(1))
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(2))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(3)))
		settings := settingsPtr.ToGo()
		ret := cb(extension, url, &client, &settings)
		if client != nil {
			*(*uintptr)(getPtr(getVal(2))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(3))) = *r
		}
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnExtensionBeforeBrowserEvent(cb TOnExtensionBeforeBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionBeforeBrowserEvent", func(getVal func(i int) uintptr) {
		// 9 : function(const extension: ICefExtension; const browser: ICefBrowser; const active_browser: ICefBrowser; index: Integer; const url: ustring; active: boolean; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings): boolean;
		extension := AsCefExtensionRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		activeBrowser := AsCefBrowserRef(getVal(2))
		index := int32(getVal(3))
		url := api.GoStr(getVal(4))
		active := api.GoBool(getVal(5))
		windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(6)))
		windowInfo := windowInfoPtr.ToGo()
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(7))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(8)))
		settings := settingsPtr.ToGo()
		ret := cb(extension, browser, activeBrowser, index, url, active, &windowInfo, &client, &settings)
		if r := windowInfo.ToPas(); r != nil {
			*(*tCefWindowInfo)(getPtr(getVal(6))) = *r
		}
		if client != nil {
			*(*uintptr)(getPtr(getVal(7))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(8))) = *r
		}
		*(*bool)(getPtr(getVal(9))) = ret
	})
}

func makeTOnExtensionCanAccessBrowserEvent(cb TOnExtensionCanAccessBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionCanAccessBrowserEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const extension: ICefExtension; const browser: ICefBrowser; include_incognito: boolean; const target_browser: ICefBrowser): boolean;
		extension := AsCefExtensionRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		includeIncognito := api.GoBool(getVal(2))
		targetBrowser := AsCefBrowserRef(getVal(3))
		ret := cb(extension, browser, includeIncognito, targetBrowser)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnExtensionExtensionLoadFailedEvent(cb TOnExtensionExtensionLoadFailedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionExtensionLoadFailedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(result: TCefErrorcode);
		result := cefTypes.TCefErrorCode(getVal(0))
		cb(result)
	})
}

func makeTOnExtensionExtensionLoadedEvent(cb TOnExtensionExtensionLoadedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionExtensionLoadedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const extension: ICefExtension);
		extension := AsCefExtensionRef(getVal(0))
		cb(extension)
	})
}

func makeTOnExtensionExtensionUnloadedEvent(cb TOnExtensionExtensionUnloadedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionExtensionUnloadedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const extension: ICefExtension);
		extension := AsCefExtensionRef(getVal(0))
		cb(extension)
	})
}

func makeTOnExtensionGetActiveBrowserEvent(cb TOnExtensionGetActiveBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionGetActiveBrowserEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const extension: ICefExtension; const browser: ICefBrowser; include_incognito: boolean; var aRsltBrowser: ICefBrowser);
		extension := AsCefExtensionRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		includeIncognito := api.GoBool(getVal(2))
		rsltBrowser := ICefBrowser(AsCefBrowserRef(*(*uintptr)(getPtr(getVal(3)))))
		cb(extension, browser, includeIncognito, &rsltBrowser)
		if rsltBrowser != nil {
			*(*uintptr)(getPtr(getVal(3))) = rsltBrowser.Instance()
		}
	})
}

func makeTOnExtensionGetExtensionResourceEvent(cb TOnExtensionGetExtensionResourceEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionGetExtensionResourceEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const extension: ICefExtension; const browser: ICefBrowser; const file_: ustring; const callback: ICefGetExtensionResourceCallback): boolean;
		extension := AsCefExtensionRef(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		file := api.GoStr(getVal(2))
		callback := AsCefGetExtensionResourceCallbackRef(getVal(3))
		ret := cb(extension, browser, file, callback)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnExtensionLoadFailedEvent(cb TOnExtensionLoadFailedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionLoadFailedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; result: TCefErrorcode);
		sender := lcl.AsObject(getVal(0))
		result := cefTypes.TCefErrorCode(getVal(1))
		cb(sender, result)
	})
}

func makeTOnExtensionLoadedEvent(cb TOnExtensionLoadedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionLoadedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const extension: ICefExtension);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		cb(sender, extension)
	})
}

func makeTOnExtensionUnloadedEvent(cb TOnExtensionUnloadedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnExtensionUnloadedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const extension: ICefExtension);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		cb(sender, extension)
	})
}

func makeTOnFavIconUrlChange(cb TOnFavIconUrlChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFavIconUrlChange", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const iconUrls: TStrings);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		iconUrls := lcl.AsStrings(getVal(2))
		cb(sender, browser, iconUrls)
	})
}

func makeTOnFileDialog(cb TOnFileDialog) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFileDialog", func(getVal func(i int) uintptr) {
		// 8 : procedure(Sender: TObject; const browser: ICefBrowser; mode: TCefFileDialogMode; const title, defaultFilePath: ustring; const acceptFilters: TStrings; const callback: ICefFileDialogCallback; var Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		mode := cefTypes.TCefFileDialogMode(getVal(2))
		title := api.GoStr(getVal(3))
		defaultFilePath := api.GoStr(getVal(4))
		acceptFilters := lcl.AsStrings(getVal(5))
		callback := AsCefFileDialogCallbackRef(getVal(6))
		result := (*bool)(getPtr(getVal(7)))
		cb(sender, browser, mode, title, defaultFilePath, acceptFilters, callback, result)
	})
}

func makeTOnFilterEvent(cb TOnFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFilterEvent", func(getVal func(i int) uintptr) {
		// 8 : procedure(Sender: TObject; data_in: Pointer; data_in_size: NativeUInt; var data_in_read: NativeUInt; data_out: Pointer; data_out_size : NativeUInt; var data_out_written: NativeUInt; var aResult : TCefResponseFilterStatus);
		sender := lcl.AsObject(getVal(0))
		dataIn := uintptr(getVal(1))
		dataInSize := cefTypes.NativeUInt(getVal(2))
		dataInRead := (*cefTypes.NativeUInt)(getPtr(getVal(3)))
		dataOut := uintptr(getVal(4))
		dataOutSize := cefTypes.NativeUInt(getVal(5))
		dataOutWritten := (*cefTypes.NativeUInt)(getPtr(getVal(6)))
		result := (*cefTypes.TCefResponseFilterStatus)(getPtr(getVal(7)))
		cb(sender, dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten, result)
	})
}

func makeTOnFindFindResultEvent(cb TOnFindFindResultEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFindFindResultEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(const browser: ICefBrowser; identifier: Integer; count: Integer; const selectionRect: PCefRect; activeMatchOrdinal: Integer; finalUpdate: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		identifier := int32(getVal(1))
		count := int32(getVal(2))
		selectionRect := *(*TCefRect)(getPtr(getVal(3)))
		activeMatchOrdinal := int32(getVal(4))
		finalUpdate := api.GoBool(getVal(5))
		cb(browser, identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
	})
}

func makeTOnFindResult(cb TOnFindResult) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFindResult", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; identifier, count: Integer; const selectionRect: PCefRect; activeMatchOrdinal: Integer; finalUpdate: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		identifier := int32(getVal(2))
		count := int32(getVal(3))
		selectionRect := *(*TCefRect)(getPtr(getVal(4)))
		activeMatchOrdinal := int32(getVal(5))
		finalUpdate := api.GoBool(getVal(6))
		cb(sender, browser, identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
	})
}

func makeTOnFocusEvent(cb TOnFocusEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFocusEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const view: ICefView);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		cb(sender, view)
	})
}

func makeTOnFocusGotFocusEvent(cb TOnFocusGotFocusEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFocusGotFocusEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnFocusSetFocusEvent(cb TOnFocusSetFocusEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFocusSetFocusEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const browser: ICefBrowser; source: TCefFocusSource): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		source := cefTypes.TCefFocusSource(getVal(1))
		ret := cb(browser, source)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnFocusTakeFocusEvent(cb TOnFocusTakeFocusEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFocusTakeFocusEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; next: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		next := api.GoBool(getVal(1))
		cb(browser, next)
	})
}

func makeTOnFocusedNodeChangedEvent(cb TOnFocusedNodeChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFocusedNodeChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const node: ICefDomNode);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		node := AsCefDomNodeRef(getVal(2))
		cb(browser, frame, node)
	})
}

func makeTOnFrameAttached(cb TOnFrameAttached) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameAttached", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; reattached: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		reattached := api.GoBool(getVal(3))
		cb(sender, browser, frame, reattached)
	})
}

func makeTOnFrameCreated(cb TOnFrameCreated) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameCreated", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		cb(sender, browser, frame)
	})
}

func makeTOnFrameDetached(cb TOnFrameDetached) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameDetached", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		cb(sender, browser, frame)
	})
}

func makeTOnFrameFrameAttachedEvent(cb TOnFrameFrameAttachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameFrameAttachedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; reattached: boolean);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		reattached := api.GoBool(getVal(2))
		cb(browser, frame, reattached)
	})
}

func makeTOnFrameFrameCreatedEvent(cb TOnFrameFrameCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameFrameCreatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		cb(browser, frame)
	})
}

func makeTOnFrameFrameDetachedEvent(cb TOnFrameFrameDetachedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameFrameDetachedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		cb(browser, frame)
	})
}

func makeTOnFrameMainFrameChangedEvent(cb TOnFrameMainFrameChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFrameMainFrameChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const old_frame: ICefFrame; const new_frame: ICefFrame);
		browser := AsCefBrowserRef(getVal(0))
		oldFrame := AsCefFrameRef(getVal(1))
		newFrame := AsCefFrameRef(getVal(2))
		cb(browser, oldFrame, newFrame)
	})
}

func makeTOnFullScreenModeChange(cb TOnFullScreenModeChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnFullScreenModeChange", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; fullscreen: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		fullscreen := api.GoBool(getVal(2))
		cb(sender, browser, fullscreen)
	})
}

func makeTOnGetAccessibilityHandler(cb TOnGetAccessibilityHandler) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetAccessibilityHandler", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; var aAccessibilityHandler : ICefAccessibilityHandler);
		sender := lcl.AsObject(getVal(0))
		var accessibilityHandler IEngAccessibilityHandler
		accessibilityHandler = AsEngAccessibilityHandler(*(*uintptr)(getPtr(getVal(1))))
		cb(sender, &accessibilityHandler)
		if accessibilityHandler != nil {
			*(*uintptr)(getPtr(getVal(1))) = accessibilityHandler.Instance()
		}
	})
}

func makeTOnGetActiveBrowserEvent(cb TOnGetActiveBrowserEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetActiveBrowserEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const extension: ICefExtension; const browser: ICefBrowser; include_incognito: boolean; var aRsltBrowser : ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		browser := AsCefBrowserRef(getVal(2))
		includeIncognito := api.GoBool(getVal(3))
		rsltBrowser := ICefBrowser(AsCefBrowserRef(*(*uintptr)(getPtr(getVal(4)))))
		cb(sender, extension, browser, includeIncognito, &rsltBrowser)
		if rsltBrowser != nil {
			*(*uintptr)(getPtr(getVal(4))) = rsltBrowser.Instance()
		}
	})
}

func makeTOnGetAudioParametersEvent(cb TOnGetAudioParametersEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetAudioParametersEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var params: TCefAudioParameters; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		params := (*TCefAudioParameters)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, params, result)
		*(*TCefAudioParameters)(getPtr(getVal(2))) = *params
	})
}

func makeTOnGetAuthCredentials(cb TOnGetAuthCredentials) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetAuthCredentials", func(getVal func(i int) uintptr) {
		// 10 : procedure(Sender: TObject; const browser: ICefBrowser; const originUrl: ustring; isProxy: Boolean; const host: ustring; port: Integer; const realm, scheme: ustring; const callback: ICefAuthCallback; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		originUrl := api.GoStr(getVal(2))
		isProxy := api.GoBool(getVal(3))
		host := api.GoStr(getVal(4))
		port := int32(getVal(5))
		realm := api.GoStr(getVal(6))
		scheme := api.GoStr(getVal(7))
		callback := AsCefAuthCallbackRef(getVal(8))
		result := (*bool)(getPtr(getVal(9)))
		cb(sender, browser, originUrl, isProxy, host, port, realm, scheme, callback, result)
	})
}

func makeTOnGetChromeToolbarTypeEvent(cb TOnGetChromeToolbarTypeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetChromeToolbarTypeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; var aChromeToolbarType: TCefChromeToolbarType);
		sender := lcl.AsObject(getVal(0))
		chromeToolbarType := (*TCefChromeToolbarType)(getPtr(getVal(1)))
		cb(sender, chromeToolbarType)
	})
}

func makeTOnGetDataResourceEvent(cb TOnGetDataResourceEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetDataResourceEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(resourceId: Integer; out data: Pointer; out dataSize: NativeUInt; var aResult : Boolean);
		resourceId := int32(getVal(0))
		data := (*uintptr)(getPtr(getVal(1)))
		dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(resourceId, data, dataSize, result)
	})
}

func makeTOnGetDataResourceForScaleEvent(cb TOnGetDataResourceForScaleEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetDataResourceForScaleEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(resourceId: Integer; scaleFactor: TCefScaleFactor; out data: Pointer; out dataSize: NativeUInt; var aResult : Boolean);
		resourceId := int32(getVal(0))
		scaleFactor := cefTypes.TCefScaleFactor(getVal(1))
		data := (*uintptr)(getPtr(getVal(2)))
		dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(3)))
		result := (*bool)(getPtr(getVal(4)))
		cb(resourceId, scaleFactor, data, dataSize, result)
	})
}

func makeTOnGetDefaultClientEvent(cb TOnGetDefaultClientEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetDefaultClientEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aClient : ICefClient);
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(0))))
		cb(&client)
		if client != nil {
			*(*uintptr)(getPtr(getVal(0))) = client.Instance()
		}
	})
}

func makeTOnGetDelegateForPopupBrowserViewEvent(cb TOnGetDelegateForPopupBrowserViewEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetDelegateForPopupBrowserViewEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; const settings: TCefBrowserSettings; const client: ICefClient; is_devtools: boolean; var aResult : ICefBrowserViewDelegate);
		sender := lcl.AsObject(getVal(0))
		browserView := AsCefBrowserViewRef(getVal(1))
		settingsPtr := (*tCefBrowserSettings)(getPtr(getVal(2)))
		settings := settingsPtr.ToGo()
		client := AsEngClient(getVal(3))
		isDevtools := api.GoBool(getVal(4))
		var result IEngBrowserViewDelegate
		result = AsEngBrowserViewDelegate(*(*uintptr)(getPtr(getVal(5))))
		cb(sender, browserView, settings, client, isDevtools, &result)
		if result != nil {
			*(*uintptr)(getPtr(getVal(5))) = result.Instance()
		}
	})
}

func makeTOnGetExtensionResourceEvent(cb TOnGetExtensionResourceEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetExtensionResourceEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const extension: ICefExtension; const browser: ICefBrowser; const file_: ustring; const callback: ICefGetExtensionResourceCallback; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		extension := AsCefExtensionRef(getVal(1))
		browser := AsCefBrowserRef(getVal(2))
		file := api.GoStr(getVal(3))
		callback := AsCefGetExtensionResourceCallbackRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, extension, browser, file, callback, result)
	})
}

func makeTOnGetHeightForWidthEvent(cb TOnGetHeightForWidthEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetHeightForWidthEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const view: ICefView; width: Integer; var aResult: Integer);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		width := int32(getVal(2))
		result := (*int32)(getPtr(getVal(3)))
		cb(sender, view, width, result)
	})
}

func makeTOnGetInitialBoundsEvent(cb TOnGetInitialBoundsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetInitialBoundsEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : TCefRect);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*TCefRect)(getPtr(getVal(2)))
		cb(sender, window, result)
		*(*TCefRect)(getPtr(getVal(2))) = *result
	})
}

func makeTOnGetInitialShowStateEvent(cb TOnGetInitialShowStateEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetInitialShowStateEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : TCefShowState);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*cefTypes.TCefShowState)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnGetLocalizedStringEvent(cb TOnGetLocalizedStringEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetLocalizedStringEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(stringId: Integer; out stringVal: ustring; var aResult : Boolean);
		stringId := int32(getVal(0))
		var stringVal string
		result := (*bool)(getPtr(getVal(2)))
		cb(stringId, &stringVal, result)
		*(*uintptr)(getPtr(getVal(1))) = api.PasStr(stringVal)
	})
}

func makeTOnGetMaximumSizeEvent(cb TOnGetMaximumSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetMaximumSizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		result := (*TCefSize)(getPtr(getVal(2)))
		cb(sender, view, result)
		*(*TCefSize)(getPtr(getVal(2))) = *result
	})
}

func makeTOnGetMinimumSizeEvent(cb TOnGetMinimumSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetMinimumSizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		result := (*TCefSize)(getPtr(getVal(2)))
		cb(sender, view, result)
		*(*TCefSize)(getPtr(getVal(2))) = *result
	})
}

func makeTOnGetPDFPaperSizeEvent(cb TOnGetPDFPaperSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetPDFPaperSizeEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; deviceUnitsPerInch: Integer; var aResult : TCefSize);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		deviceUnitsPerInch := int32(getVal(2))
		result := (*TCefSize)(getPtr(getVal(3)))
		cb(sender, browser, deviceUnitsPerInch, result)
		*(*TCefSize)(getPtr(getVal(3))) = *result
	})
}

func makeTOnGetParentWindowEvent(cb TOnGetParentWindowEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetParentWindowEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const Sender: TObject; const window_: ICefWindow; var is_menu, can_activate_menu: boolean; var aResult : ICefWindow);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		isMenu := (*bool)(getPtr(getVal(2)))
		canActivateMenu := (*bool)(getPtr(getVal(3)))
		result := ICefWindow(AsCefWindowRef(*(*uintptr)(getPtr(getVal(4)))))
		cb(sender, window, isMenu, canActivateMenu, &result)
		if result != nil {
			*(*uintptr)(getPtr(getVal(4))) = result.Instance()
		}
	})
}

func makeTOnGetPreferredSizeEvent(cb TOnGetPreferredSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetPreferredSizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		result := (*TCefSize)(getPtr(getVal(2)))
		cb(sender, view, result)
		*(*TCefSize)(getPtr(getVal(2))) = *result
	})
}

func makeTOnGetResourceHandler(cb TOnGetResourceHandler) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetResourceHandler", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var aResourceHandler : ICefResourceHandler);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		var resourceHandler IEngResourceHandler
		resourceHandler = AsEngResourceHandler(*(*uintptr)(getPtr(getVal(4))))
		cb(sender, browser, frame, request, &resourceHandler)
		if resourceHandler != nil {
			*(*uintptr)(getPtr(getVal(4))) = resourceHandler.Instance()
		}
	})
}

func makeTOnGetResourceRequestHandler(cb TOnGetResourceRequestHandler) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetResourceRequestHandler", func(getVal func(i int) uintptr) {
		// 9 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; is_navigation, is_download: boolean; const request_initiator: ustring; var disable_default_handling: boolean; var aExternalResourceRequestHandler : ICefResourceRequestHandler);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		isNavigation := api.GoBool(getVal(4))
		isDownload := api.GoBool(getVal(5))
		requestInitiator := api.GoStr(getVal(6))
		disableDefaultHandling := (*bool)(getPtr(getVal(7)))
		var externalResourceRequestHandler IEngResourceRequestHandler
		externalResourceRequestHandler = AsEngResourceRequestHandler(*(*uintptr)(getPtr(getVal(8))))
		cb(sender, browser, frame, request, isNavigation, isDownload, requestInitiator, disableDefaultHandling, &externalResourceRequestHandler)
		if externalResourceRequestHandler != nil {
			*(*uintptr)(getPtr(getVal(8))) = externalResourceRequestHandler.Instance()
		}
	})
}

func makeTOnGetResourceResponseFilter(cb TOnGetResourceResponseFilter) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetResourceResponseFilter", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; out Result: ICefResponseFilter);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		response := AsCefResponseRef(getVal(4))
		var result IEngResponseFilter
		cb(sender, browser, frame, request, response, &result)
		if result != nil {
			*(*uintptr)(getPtr(getVal(5))) = result.Instance()
		}
	})
}

func makeTOnGetRootScreenRect(cb TOnGetRootScreenRect) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetRootScreenRect", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var rect: TCefRect; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		rect := (*TCefRect)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, rect, result)
		*(*TCefRect)(getPtr(getVal(2))) = *rect
	})
}

func makeTOnGetScreenInfo(cb TOnGetScreenInfo) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetScreenInfo", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var screenInfo: TCefScreenInfo; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		screenInfo := (*TCefScreenInfo)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, screenInfo, result)
		*(*TCefScreenInfo)(getPtr(getVal(2))) = *screenInfo
	})
}

func makeTOnGetScreenPoint(cb TOnGetScreenPoint) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetScreenPoint", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; viewX, viewY: Integer; var screenX, screenY: Integer; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		viewX := int32(getVal(2))
		viewY := int32(getVal(3))
		screenX := (*int32)(getPtr(getVal(4)))
		screenY := (*int32)(getPtr(getVal(5)))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, viewX, viewY, screenX, screenY, result)
	})
}

func makeTOnGetTouchHandleSize(cb TOnGetTouchHandleSize) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetTouchHandleSize", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; orientation: TCefHorizontalAlignment; var size: TCefSize);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		orientation := cefTypes.TCefHorizontalAlignment(getVal(2))
		size := (*TCefSize)(getPtr(getVal(3)))
		cb(sender, browser, orientation, size)
		*(*TCefSize)(getPtr(getVal(3))) = *size
	})
}

func makeTOnGetViewRect(cb TOnGetViewRect) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGetViewRect", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; var rect: TCefRect);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		rect := (*TCefRect)(getPtr(getVal(2)))
		cb(sender, browser, rect)
		*(*TCefRect)(getPtr(getVal(2))) = *rect
	})
}

func makeTOnGotFocus(cb TOnGotFocus) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnGotFocus", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnHandledMessageEvent(cb TOnHandledMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnHandledMessageEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; var aMessage: TMessage; var aHandled : boolean);
		sender := lcl.AsObject(getVal(0))
		message := (*types.TMessage)(getPtr(getVal(1)))
		handled := (*bool)(getPtr(getVal(2)))
		cb(sender, message, handled)
	})
}

func makeTOnHttpRequest(cb TOnHttpRequest) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnHttpRequest", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		clientAddress := api.GoStr(getVal(3))
		request := AsCefRequestRef(getVal(4))
		cb(sender, server, connectionId, clientAddress, request)
	})
}

func makeTOnIMECommitTextEvent(cb TOnIMECommitTextEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnIMECommitTextEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const aText : ustring; const replacement_range : PCefRange; relative_cursor_pos : integer);
		sender := lcl.AsObject(getVal(0))
		text := api.GoStr(getVal(1))
		replacementRange := *(*TCefRange)(getPtr(getVal(2)))
		relativeCursorPos := int32(getVal(3))
		cb(sender, text, replacementRange, relativeCursorPos)
	})
}

func makeTOnIMECompositionRangeChanged(cb TOnIMECompositionRangeChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnIMECompositionRangeChanged", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const selected_range: PCefRange; character_boundsCount: NativeUInt; const character_bounds: PCefRect);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		selectedRange := *(*TCefRange)(getPtr(getVal(2)))
		characterBoundsCount := cefTypes.NativeUInt(getVal(3))
		characterBounds := *(*TCefRect)(getPtr(getVal(4)))
		cb(sender, browser, selectedRange, characterBoundsCount, characterBounds)
	})
}

func makeTOnIMESetCompositionEvent(cb TOnIMESetCompositionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnIMESetCompositionEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const aText : ustring; const underlines : TCefCompositionUnderlineDynArray; const replacement_range, selection_range : TCefRange);
		sender := lcl.AsObject(getVal(0))
		text := api.GoStr(getVal(1))
		underlinesPtr := getVal(2)
		underlinesLen := int32(getVal(3))
		underlines := NewCefCompositionUnderlineArray(int(underlinesLen), underlinesPtr)
		replacementRange := *(*TCefRange)(getPtr(getVal(4)))
		selectionRange := *(*TCefRange)(getPtr(getVal(5)))
		cb(sender, text, underlines, replacementRange, selectionRange)
	})
}

func makeTOnInitFilterEvent(cb TOnInitFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnInitFilterEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(sender, result)
	})
}

func makeTOnIsFramelessEvent(cb TOnIsFramelessEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnIsFramelessEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(sender, window, result)
	})
}

func makeTOnJsDialogBeforeUnloadDialogEvent(cb TOnJsDialogBeforeUnloadDialogEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnJsDialogBeforeUnloadDialogEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const messageText: ustring; isReload: Boolean; const callback: ICefJsDialogCallback): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		messageText := api.GoStr(getVal(1))
		isReload := api.GoBool(getVal(2))
		callback := AsCefJsDialogCallbackRef(getVal(3))
		ret := cb(browser, messageText, isReload, callback)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnJsDialogDialogClosedEvent(cb TOnJsDialogDialogClosedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnJsDialogDialogClosedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnJsDialogJsdialogEvent(cb TOnJsDialogJsdialogEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnJsDialogJsdialogEvent", func(getVal func(i int) uintptr) {
		// 7 : function(const browser: ICefBrowser; const originUrl: ustring; dialogType: TCefJsDialogType; const messageText: ustring; const defaultPromptText: ustring; const callback: ICefJsDialogCallback; out suppressMessage: Boolean): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		originUrl := api.GoStr(getVal(1))
		dialogType := cefTypes.TCefJsDialogType(getVal(2))
		messageText := api.GoStr(getVal(3))
		defaultPromptText := api.GoStr(getVal(4))
		callback := AsCefJsDialogCallbackRef(getVal(5))
		suppressMessage := (*bool)(getPtr(getVal(6)))
		ret := cb(browser, originUrl, dialogType, messageText, defaultPromptText, callback, suppressMessage)
		*(*bool)(getPtr(getVal(7))) = ret
	})
}

func makeTOnJsDialogResetDialogStateEvent(cb TOnJsDialogResetDialogStateEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnJsDialogResetDialogStateEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnJsdialog(cb TOnJsdialog) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnJsdialog", func(getVal func(i int) uintptr) {
		// 9 : procedure(Sender: TObject; const browser: ICefBrowser; const originUrl: ustring; dialogType: TCefJsDialogType; const messageText, defaultPromptText: ustring; const callback: ICefJsDialogCallback; out suppressMessage: Boolean; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		originUrl := api.GoStr(getVal(2))
		dialogType := cefTypes.TCefJsDialogType(getVal(3))
		messageText := api.GoStr(getVal(4))
		defaultPromptText := api.GoStr(getVal(5))
		callback := AsCefJsDialogCallbackRef(getVal(6))
		suppressMessage := (*bool)(getPtr(getVal(7)))
		result := (*bool)(getPtr(getVal(8)))
		cb(sender, browser, originUrl, dialogType, messageText, defaultPromptText, callback, suppressMessage, result)
	})
}

func makeTOnKeyEvent(cb TOnKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnKeyEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		event := *(*TCefKeyEvent)(getPtr(getVal(2)))
		osEvent := cefTypes.TCefEventHandle(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, event, osEvent, result)
	})
}

func makeTOnKeyboardKeyEvent(cb TOnKeyboardKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnKeyboardKeyEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		event := *(*TCefKeyEvent)(getPtr(getVal(1)))
		osEvent := cefTypes.TCefEventHandle(getVal(2))
		ret := cb(browser, event, osEvent)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnKeyboardPreKeyEvent(cb TOnKeyboardPreKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnKeyboardPreKeyEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out isKeyboardShortcut: Boolean): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		event := *(*TCefKeyEvent)(getPtr(getVal(1)))
		osEvent := cefTypes.TCefEventHandle(getVal(2))
		isKeyboardShortcut := (*bool)(getPtr(getVal(3)))
		ret := cb(browser, event, osEvent, isKeyboardShortcut)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnLayoutChangedEvent(cb TOnLayoutChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLayoutChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const view: ICefView; new_bounds: TCefRect);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		newBounds := *(*TCefRect)(getPtr(getVal(2)))
		cb(sender, view, newBounds)
	})
}

func makeTOnLifeSpanAfterCreatedEvent(cb TOnLifeSpanAfterCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLifeSpanAfterCreatedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnLifeSpanBeforeCloseEvent(cb TOnLifeSpanBeforeCloseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLifeSpanBeforeCloseEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnLifeSpanBeforePopupEvent(cb TOnLifeSpanBeforePopupEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLifeSpanBeforePopupEvent", func(getVal func(i int) uintptr) {
		// 12 : function(const browser: ICefBrowser; const frame: ICefFrame; const targetUrl: ustring; const targetFrameName: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; const popupFeatures: TCefPopupFeatures; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var noJavascriptAccess: Boolean): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		targetUrl := api.GoStr(getVal(2))
		targetFrameName := api.GoStr(getVal(3))
		targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(4))
		userGesture := api.GoBool(getVal(5))
		popupFeatures := *(*TCefPopupFeatures)(getPtr(getVal(6)))
		windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(7)))
		windowInfo := windowInfoPtr.ToGo()
		var client IEngClient
		client = AsEngClient(*(*uintptr)(getPtr(getVal(8))))
		settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(9)))
		settings := settingsPtr.ToGo()
		extraInfo := ICefDictionaryValue(AsCefDictionaryValueRef(*(*uintptr)(getPtr(getVal(10)))))
		noJavascriptAccess := (*bool)(getPtr(getVal(11)))
		ret := cb(browser, frame, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, &windowInfo, &client, &settings, &extraInfo, noJavascriptAccess)
		if r := windowInfo.ToPas(); r != nil {
			*(*tCefWindowInfo)(getPtr(getVal(7))) = *r
		}
		if client != nil {
			*(*uintptr)(getPtr(getVal(8))) = client.Instance()
		}
		if r := settings.ToPas(); r != nil {
			*(*tCefBrowserSettings)(getPtr(getVal(9))) = *r
		}
		if extraInfo != nil {
			*(*uintptr)(getPtr(getVal(10))) = extraInfo.Instance()
		}
		*(*bool)(getPtr(getVal(12))) = ret
	})
}

func makeTOnLifeSpanDoCloseEvent(cb TOnLifeSpanDoCloseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLifeSpanDoCloseEvent", func(getVal func(i int) uintptr) {
		// 1 : function(const browser: ICefBrowser): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		ret := cb(browser)
		*(*bool)(getPtr(getVal(1))) = ret
	})
}

func makeTOnLoadEnd(cb TOnLoadEnd) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadEnd", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		httpStatusCode := int32(getVal(3))
		cb(sender, browser, frame, httpStatusCode)
	})
}

func makeTOnLoadError(cb TOnLoadError) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadError", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText, failedUrl: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		errorCode := cefTypes.TCefErrorCode(getVal(3))
		errorText := api.GoStr(getVal(4))
		failedUrl := api.GoStr(getVal(5))
		cb(sender, browser, frame, errorCode, errorText, failedUrl)
	})
}

func makeTOnLoadLoadEndEvent(cb TOnLoadLoadEndEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadLoadEndEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		httpStatusCode := int32(getVal(2))
		cb(browser, frame, httpStatusCode)
	})
}

func makeTOnLoadLoadErrorEvent(cb TOnLoadLoadErrorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadLoadErrorEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText: ustring; const failedUrl: ustring);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		errorCode := cefTypes.TCefErrorCode(getVal(2))
		errorText := api.GoStr(getVal(3))
		failedUrl := api.GoStr(getVal(4))
		cb(browser, frame, errorCode, errorText, failedUrl)
	})
}

func makeTOnLoadLoadStartEvent(cb TOnLoadLoadStartEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadLoadStartEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		transitionType := cefTypes.TCefTransitionType(getVal(2))
		cb(browser, frame, transitionType)
	})
}

func makeTOnLoadLoadingStateChangeEvent(cb TOnLoadLoadingStateChangeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadLoadingStateChangeEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; isLoading: Boolean; canGoBack: Boolean; canGoForward: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		isLoading := api.GoBool(getVal(1))
		canGoBack := api.GoBool(getVal(2))
		canGoForward := api.GoBool(getVal(3))
		cb(browser, isLoading, canGoBack, canGoForward)
	})
}

func makeTOnLoadStart(cb TOnLoadStart) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadStart", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		transitionType := cefTypes.TCefTransitionType(getVal(3))
		cb(sender, browser, frame, transitionType)
	})
}

func makeTOnLoadingProgressChange(cb TOnLoadingProgressChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadingProgressChange", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const progress: double);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		progress := *(*float64)(getPtr(getVal(2)))
		cb(sender, browser, progress)
	})
}

func makeTOnLoadingStateChange(cb TOnLoadingStateChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnLoadingStateChange", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; isLoading, canGoBack, canGoForward: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		isLoading := api.GoBool(getVal(2))
		canGoBack := api.GoBool(getVal(3))
		canGoForward := api.GoBool(getVal(4))
		cb(sender, browser, isLoading, canGoBack, canGoForward)
	})
}

func makeTOnMainFrameChanged(cb TOnMainFrameChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMainFrameChanged", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const old_frame, new_frame: ICefFrame);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		oldFrame := AsCefFrameRef(getVal(2))
		newFrame := AsCefFrameRef(getVal(3))
		cb(sender, browser, oldFrame, newFrame)
	})
}

func makeTOnMediaAccessChange(cb TOnMediaAccessChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaAccessChange", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; has_video_access, has_audio_access: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		hasVideoAccess := api.GoBool(getVal(2))
		hasAudioAccess := api.GoBool(getVal(3))
		cb(sender, browser, hasVideoAccess, hasAudioAccess)
	})
}

func makeTOnMediaAccessRequestMediaAccessPermissionEvent(cb TOnMediaAccessRequestMediaAccessPermissionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaAccessRequestMediaAccessPermissionEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const requesting_url: ustring; requested_permissions: TCefMediaAccessPermissionTypes; const callback: ICefMediaAccessCallback): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		requestingUrl := api.GoStr(getVal(2))
		requestedPermissions := cefTypes.TCefMediaAccessPermissionTypes(getVal(3))
		callback := AsCefMediaAccessCallbackRef(getVal(4))
		ret := cb(browser, frame, requestingUrl, requestedPermissions, callback)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnMediaObserverRouteMessageReceivedEvent(cb TOnMediaObserverRouteMessageReceivedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaObserverRouteMessageReceivedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const route: ICefMediaRoute; const message_: ustring);
		route := AsCefMediaRouteRef(getVal(0))
		message := api.GoStr(getVal(1))
		cb(route, message)
	})
}

func makeTOnMediaObserverRouteStateChangedEvent(cb TOnMediaObserverRouteStateChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaObserverRouteStateChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const route: ICefMediaRoute; state: TCefMediaRouteConnectionState);
		route := AsCefMediaRouteRef(getVal(0))
		state := cefTypes.TCefMediaRouteConnectionState(getVal(1))
		cb(route, state)
	})
}

func makeTOnMediaObserverRoutesEvent(cb TOnMediaObserverRoutesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaObserverRoutesEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const routes: TCefMediaRouteArray);
		routesPtr := getVal(0)
		routesLen := int32(getVal(1))
		routes := NewCefMediaRouteArray(int(routesLen), routesPtr)
		cb(routes)
	})
}

func makeTOnMediaObserverSinksEvent(cb TOnMediaObserverSinksEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaObserverSinksEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const sinks: TCefMediaSinkArray);
		sinksPtr := getVal(0)
		sinksLen := int32(getVal(1))
		sinks := NewCefMediaSinkArray(int(sinksLen), sinksPtr)
		cb(sinks)
	})
}

func makeTOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent(cb TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(result: TCefMediaRouterCreateResult; const error: ustring; const route: ICefMediaRoute);
		result := cefTypes.TCefMediaRouterCreateResult(getVal(0))
		error_ := api.GoStr(getVal(1))
		route := AsCefMediaRouteRef(getVal(2))
		cb(result, error_, route)
	})
}

func makeTOnMediaRouteCreateFinishedEvent(cb TOnMediaRouteCreateFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaRouteCreateFinishedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; result: TCefMediaRouterCreateResult; const error: ustring; const route: ICefMediaRoute);
		sender := lcl.AsObject(getVal(0))
		result := cefTypes.TCefMediaRouterCreateResult(getVal(1))
		error_ := api.GoStr(getVal(2))
		route := AsCefMediaRouteRef(getVal(3))
		cb(sender, result, error_, route)
	})
}

func makeTOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent(cb TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const ip_address: ustring; port: integer; const model_name: ustring);
		ipAddress := api.GoStr(getVal(0))
		port := int32(getVal(1))
		modelName := api.GoStr(getVal(2))
		cb(ipAddress, port, modelName)
	})
}

func makeTOnMediaSinkDeviceInfoEvent(cb TOnMediaSinkDeviceInfoEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMediaSinkDeviceInfoEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const ip_address: ustring; port: integer; const model_name: ustring);
		sender := lcl.AsObject(getVal(0))
		ipAddress := api.GoStr(getVal(1))
		port := int32(getVal(2))
		modelName := api.GoStr(getVal(3))
		cb(sender, ipAddress, port, modelName)
	})
}

func makeTOnMenuButtonMenuButtonPressedEvent(cb TOnMenuButtonMenuButtonPressedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuButtonMenuButtonPressedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const menu_button: ICefMenuButton; const screen_point: TCefPoint; const button_pressed_lock: ICefMenuButtonPressedLock);
		menuButton := AsCefMenuButtonRef(getVal(0))
		screenPoint := *(*TCefPoint)(getPtr(getVal(1)))
		buttonPressedLock := AsCefMenuButtonPressedLockRef(getVal(2))
		cb(menuButton, screenPoint, buttonPressedLock)
	})
}

func makeTOnMenuButtonPressedEvent(cb TOnMenuButtonPressedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuButtonPressedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const menu_button: ICefMenuButton; const screen_point: TCefPoint; const button_pressed_lock: ICefMenuButtonPressedLock);
		sender := lcl.AsObject(getVal(0))
		menuButton := AsCefMenuButtonRef(getVal(1))
		screenPoint := *(*TCefPoint)(getPtr(getVal(2)))
		buttonPressedLock := AsCefMenuButtonPressedLockRef(getVal(3))
		cb(sender, menuButton, screenPoint, buttonPressedLock)
	})
}

func makeTOnMenuModelExecuteCommandEvent(cb TOnMenuModelExecuteCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelExecuteCommandEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const menuModel: ICefMenuModel; commandId: Integer; eventFlags: TCefEventFlags);
		menuModel := AsCefMenuModelRef(getVal(0))
		commandId := int32(getVal(1))
		eventFlags := cefTypes.TCefEventFlags(getVal(2))
		cb(menuModel, commandId, eventFlags)
	})
}

func makeTOnMenuModelFormatLabelEvent(cb TOnMenuModelFormatLabelEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelFormatLabelEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const menuModel: ICefMenuModel; var label_: ustring): boolean;
		menuModel := AsCefMenuModelRef(getVal(0))
		label := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
		ret := cb(menuModel, &label)
		*(*uintptr)(getPtr(getVal(1))) = api.PasStr(label)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnMenuModelMenuClosedEvent(cb TOnMenuModelMenuClosedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelMenuClosedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const menuModel: ICefMenuModel);
		menuModel := AsCefMenuModelRef(getVal(0))
		cb(menuModel)
	})
}

func makeTOnMenuModelMenuWillShowEvent(cb TOnMenuModelMenuWillShowEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelMenuWillShowEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const menuModel: ICefMenuModel);
		menuModel := AsCefMenuModelRef(getVal(0))
		cb(menuModel)
	})
}

func makeTOnMenuModelMouseOutsideMenuEvent(cb TOnMenuModelMouseOutsideMenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelMouseOutsideMenuEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const menuModel: ICefMenuModel; const screenPoint: PCefPoint);
		menuModel := AsCefMenuModelRef(getVal(0))
		screenPoint := *(*TCefPoint)(getPtr(getVal(1)))
		cb(menuModel, screenPoint)
	})
}

func makeTOnMenuModelUnhandledCloseSubmenuEvent(cb TOnMenuModelUnhandledCloseSubmenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelUnhandledCloseSubmenuEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const menuModel: ICefMenuModel; isRTL: boolean);
		menuModel := AsCefMenuModelRef(getVal(0))
		isRTL := api.GoBool(getVal(1))
		cb(menuModel, isRTL)
	})
}

func makeTOnMenuModelUnhandledOpenSubmenuEvent(cb TOnMenuModelUnhandledOpenSubmenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnMenuModelUnhandledOpenSubmenuEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const menuModel: ICefMenuModel; isRTL: boolean);
		menuModel := AsCefMenuModelRef(getVal(0))
		isRTL := api.GoBool(getVal(1))
		cb(menuModel, isRTL)
	})
}

func makeTOnNavigationEntryVisitorVisitEvent(cb TOnNavigationEntryVisitorVisitEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnNavigationEntryVisitorVisitEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const entry: ICefNavigationEntry; current: Boolean; index: Integer; total: Integer): Boolean;
		entry := AsCefNavigationEntryRef(getVal(0))
		current := api.GoBool(getVal(1))
		index := int32(getVal(2))
		total := int32(getVal(3))
		ret := cb(entry, current, index, total)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnNavigationVisitorResultAvailableEvent(cb TOnNavigationVisitorResultAvailableEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnNavigationVisitorResultAvailableEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const entry: ICefNavigationEntry; current: Boolean; index, total: Integer; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		entry := AsCefNavigationEntryRef(getVal(1))
		current := api.GoBool(getVal(2))
		index := int32(getVal(3))
		total := int32(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, entry, current, index, total, result)
	})
}

func makeTOnOpenUrlFromTab(cb TOnOpenUrlFromTab) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnOpenUrlFromTab", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const targetUrl: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		targetUrl := api.GoStr(getVal(3))
		targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(4))
		userGesture := api.GoBool(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, targetUrl, targetDisposition, userGesture, result)
	})
}

func makeTOnPaint(cb TOnPaint) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPaint", func(getVal func(i int) uintptr) {
		// 8 : procedure(Sender: TObject; const browser: ICefBrowser; type_: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; const buffer: Pointer; width, height: Integer);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		type_ := cefTypes.TCefPaintElementType(getVal(2))
		dirtyRectsCount := cefTypes.NativeUInt(getVal(3))
		dirtyRectsPtr := getVal(4)
		dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
		buffer := uintptr(getVal(5))
		width := int32(getVal(6))
		height := int32(getVal(7))
		cb(sender, browser, type_, dirtyRectsCount, dirtyRects, buffer, width, height)
	})
}

func makeTOnParentViewChangedEvent(cb TOnParentViewChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnParentViewChangedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const view: ICefView; added: boolean; const parent: ICefView);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		added := api.GoBool(getVal(2))
		parent := AsCefViewRef(getVal(3))
		cb(sender, view, added, parent)
	})
}

func makeTOnPdfPrintCallbackPdfPrintFinishedEvent(cb TOnPdfPrintCallbackPdfPrintFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPdfPrintCallbackPdfPrintFinishedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const path: ustring; ok: Boolean);
		path := api.GoStr(getVal(0))
		ok := api.GoBool(getVal(1))
		cb(path, ok)
	})
}

func makeTOnPdfPrintFinishedEvent(cb TOnPdfPrintFinishedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPdfPrintFinishedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; aResultOK : boolean);
		sender := lcl.AsObject(getVal(0))
		resultOK := api.GoBool(getVal(1))
		cb(sender, resultOK)
	})
}

func makeTOnPermissionDismissPermissionPromptEvent(cb TOnPermissionDismissPermissionPromptEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPermissionDismissPermissionPromptEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; prompt_id: uint64; result: TCefPermissionRequestResult);
		browser := AsCefBrowserRef(getVal(0))
		promptId := *(*uint64)(getPtr(getVal(1)))
		result := cefTypes.TCefPermissionRequestResult(getVal(2))
		cb(browser, promptId, result)
	})
}

func makeTOnPermissionRequestMediaAccessPermissionEvent(cb TOnPermissionRequestMediaAccessPermissionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPermissionRequestMediaAccessPermissionEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefMediaAccessCallback): boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		requestingOrigin := api.GoStr(getVal(2))
		requestedPermissions := uint32(getVal(3))
		callback := AsCefMediaAccessCallbackRef(getVal(4))
		ret := cb(browser, frame, requestingOrigin, requestedPermissions, callback)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnPermissionShowPermissionPromptEvent(cb TOnPermissionShowPermissionPromptEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPermissionShowPermissionPromptEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; prompt_id: uint64; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefPermissionPromptCallback): boolean;
		browser := AsCefBrowserRef(getVal(0))
		promptId := *(*uint64)(getPtr(getVal(1)))
		requestingOrigin := api.GoStr(getVal(2))
		requestedPermissions := uint32(getVal(3))
		callback := AsCefPermissionPromptCallbackRef(getVal(4))
		ret := cb(browser, promptId, requestingOrigin, requestedPermissions, callback)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnPopupBrowserViewCreatedEvent(cb TOnPopupBrowserViewCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPopupBrowserViewCreatedEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const Sender: TObject; const browser_view, popup_browser_view: ICefBrowserView; is_devtools: boolean; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browserView := AsCefBrowserViewRef(getVal(1))
		popupBrowserView := AsCefBrowserViewRef(getVal(2))
		isDevtools := api.GoBool(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browserView, popupBrowserView, isDevtools, result)
	})
}

func makeTOnPopupShow(cb TOnPopupShow) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPopupShow", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; show: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		show := api.GoBool(getVal(2))
		cb(sender, browser, show)
	})
}

func makeTOnPopupSize(cb TOnPopupSize) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPopupSize", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const rect: PCefRect);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		rect := *(*TCefRect)(getPtr(getVal(2)))
		cb(sender, browser, rect)
	})
}

func makeTOnPreKeyEvent(cb TOnPreKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPreKeyEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out isKeyboardShortcut: Boolean; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		event := *(*TCefKeyEvent)(getPtr(getVal(2)))
		osEvent := cefTypes.TCefEventHandle(getVal(3))
		isKeyboardShortcut := (*bool)(getPtr(getVal(4)))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, event, osEvent, isKeyboardShortcut, result)
	})
}

func makeTOnPrefsAvailableEvent(cb TOnPrefsAvailableEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrefsAvailableEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; aResultOK : boolean);
		sender := lcl.AsObject(getVal(0))
		resultOK := api.GoBool(getVal(1))
		cb(sender, resultOK)
	})
}

func makeTOnPrintDialogEvent(cb TOnPrintDialogEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintDialogEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; hasSelection: boolean; const callback: ICefPrintDialogCallback; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		hasSelection := api.GoBool(getVal(2))
		callback := AsCefPrintDialogCallbackRef(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, hasSelection, callback, result)
	})
}

func makeTOnPrintGetPDFPaperSizeEvent(cb TOnPrintGetPDFPaperSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintGetPDFPaperSizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; deviceUnitsPerInch: Integer; var aResult: TCefSize);
		browser := AsCefBrowserRef(getVal(0))
		deviceUnitsPerInch := int32(getVal(1))
		result := (*TCefSize)(getPtr(getVal(2)))
		cb(browser, deviceUnitsPerInch, result)
		*(*TCefSize)(getPtr(getVal(2))) = *result
	})
}

func makeTOnPrintJobEvent(cb TOnPrintJobEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintJobEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const documentName, PDFFilePath: ustring; const callback: ICefPrintJobCallback; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		documentName := api.GoStr(getVal(2))
		pDFFilePath := api.GoStr(getVal(3))
		callback := AsCefPrintJobCallbackRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, documentName, pDFFilePath, callback, result)
	})
}

func makeTOnPrintPrintDialogEvent(cb TOnPrintPrintDialogEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintPrintDialogEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; hasSelection: boolean; const callback: ICefPrintDialogCallback; var aResult: boolean);
		browser := AsCefBrowserRef(getVal(0))
		hasSelection := api.GoBool(getVal(1))
		callback := AsCefPrintDialogCallbackRef(getVal(2))
		result := (*bool)(getPtr(getVal(3)))
		cb(browser, hasSelection, callback, result)
	})
}

func makeTOnPrintPrintJobEvent(cb TOnPrintPrintJobEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintPrintJobEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const documentName: ustring; const PDFFilePath: ustring; const callback: ICefPrintJobCallback; var aResult: boolean);
		browser := AsCefBrowserRef(getVal(0))
		documentName := api.GoStr(getVal(1))
		pDFFilePath := api.GoStr(getVal(2))
		callback := AsCefPrintJobCallbackRef(getVal(3))
		result := (*bool)(getPtr(getVal(4)))
		cb(browser, documentName, pDFFilePath, callback, result)
	})
}

func makeTOnPrintPrintResetEvent(cb TOnPrintPrintResetEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintPrintResetEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnPrintPrintSettingsEvent(cb TOnPrintPrintSettingsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintPrintSettingsEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const settings: ICefPrintSettings; getDefaults: boolean);
		browser := AsCefBrowserRef(getVal(0))
		settings := AsCefPrintSettingsRef(getVal(1))
		getDefaults := api.GoBool(getVal(2))
		cb(browser, settings, getDefaults)
	})
}

func makeTOnPrintPrintStartEvent(cb TOnPrintPrintStartEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintPrintStartEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnPrintResetEvent(cb TOnPrintResetEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintResetEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnPrintSettingsEvent(cb TOnPrintSettingsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintSettingsEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const settings: ICefPrintSettings; getDefaults: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		settings := AsCefPrintSettingsRef(getVal(2))
		getDefaults := api.GoBool(getVal(3))
		cb(sender, browser, settings, getDefaults)
	})
}

func makeTOnPrintStartEvent(cb TOnPrintStartEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnPrintStartEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnProcessMessageReceived(cb TOnProcessMessageReceived) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnProcessMessageReceived", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message: ICefProcessMessage; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		sourceProcess := cefTypes.TCefProcessId(getVal(3))
		message := AsCefProcessMessageRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, frame, sourceProcess, message, result)
	})
}

func makeTOnProcessMessageReceivedEvent(cb TOnProcessMessageReceivedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnProcessMessageReceivedEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message: ICefProcessMessage; var aHandled : boolean);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		sourceProcess := cefTypes.TCefProcessId(getVal(2))
		message := AsCefProcessMessageRef(getVal(3))
		handled := (*bool)(getPtr(getVal(4)))
		cb(browser, frame, sourceProcess, message, handled)
	})
}

func makeTOnProtocolExecution(cb TOnProtocolExecution) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnProtocolExecution", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var allowOsExecution: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		allowOsExecution := (*bool)(getPtr(getVal(4)))
		cb(sender, browser, frame, request, allowOsExecution)
	})
}

func makeTOnQuickMenuCommandEvent(cb TOnQuickMenuCommandEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnQuickMenuCommandEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; command_id: integer; event_flags: TCefEventFlags; var aResult: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		commandId := int32(getVal(3))
		eventFlags := cefTypes.TCefEventFlags(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, frame, commandId, eventFlags, result)
	})
}

func makeTOnQuickMenuDismissedEvent(cb TOnQuickMenuDismissedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnQuickMenuDismissedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		cb(sender, browser, frame)
	})
}

func makeTOnRegisterCustomPreferencesEvent(cb TOnRegisterCustomPreferencesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRegisterCustomPreferencesEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(type_: TCefPreferencesType; const registrar: TCefPreferenceRegistrarRef);
		type_ := cefTypes.TCefPreferencesType(getVal(0))
		registrar := AsCefPreferenceRegistrarRef(getVal(1))
		cb(type_, registrar)
	})
}

func makeTOnRegisterCustomSchemesEvent(cb TOnRegisterCustomSchemesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRegisterCustomSchemesEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const registrar: TCefSchemeRegistrarRef);
		registrar := AsCefSchemeRegistrarRef(getVal(0))
		cb(registrar)
	})
}

func makeTOnRenderAcceleratedPaintEvent(cb TOnRenderAcceleratedPaintEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderAcceleratedPaintEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; kind: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; shared_handle: Pointer);
		browser := AsCefBrowserRef(getVal(0))
		kind := cefTypes.TCefPaintElementType(getVal(1))
		dirtyRectsCount := cefTypes.NativeUInt(getVal(2))
		dirtyRectsPtr := getVal(3)
		dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
		sharedHandle := uintptr(getVal(4))
		cb(browser, kind, dirtyRectsCount, dirtyRects, sharedHandle)
	})
}

func makeTOnRenderGetAccessibilityHandlerEvent(cb TOnRenderGetAccessibilityHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetAccessibilityHandlerEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(var aAccessibilityHandler: ICefAccessibilityHandler);
		var accessibilityHandler IEngAccessibilityHandler
		accessibilityHandler = AsEngAccessibilityHandler(*(*uintptr)(getPtr(getVal(0))))
		cb(&accessibilityHandler)
		if accessibilityHandler != nil {
			*(*uintptr)(getPtr(getVal(0))) = accessibilityHandler.Instance()
		}
	})
}

func makeTOnRenderGetRootScreenRectEvent(cb TOnRenderGetRootScreenRectEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetRootScreenRectEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const browser: ICefBrowser; var rect: TCefRect): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		rect := (*TCefRect)(getPtr(getVal(1)))
		ret := cb(browser, rect)
		*(*TCefRect)(getPtr(getVal(1))) = *rect
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnRenderGetScreenInfoEvent(cb TOnRenderGetScreenInfoEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetScreenInfoEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const browser: ICefBrowser; var screenInfo: TCefScreenInfo): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		screenInfo := (*TCefScreenInfo)(getPtr(getVal(1)))
		ret := cb(browser, screenInfo)
		*(*TCefScreenInfo)(getPtr(getVal(1))) = *screenInfo
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnRenderGetScreenPointEvent(cb TOnRenderGetScreenPointEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetScreenPointEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; viewX: Integer; viewY: Integer; var screenX: Integer; var screenY: Integer): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		viewX := int32(getVal(1))
		viewY := int32(getVal(2))
		screenX := (*int32)(getPtr(getVal(3)))
		screenY := (*int32)(getPtr(getVal(4)))
		ret := cb(browser, viewX, viewY, screenX, screenY)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnRenderGetTouchHandleSizeEvent(cb TOnRenderGetTouchHandleSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetTouchHandleSizeEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; orientation: TCefHorizontalAlignment; var size: TCefSize);
		browser := AsCefBrowserRef(getVal(0))
		orientation := cefTypes.TCefHorizontalAlignment(getVal(1))
		size := (*TCefSize)(getPtr(getVal(2)))
		cb(browser, orientation, size)
		*(*TCefSize)(getPtr(getVal(2))) = *size
	})
}

func makeTOnRenderGetViewRectEvent(cb TOnRenderGetViewRectEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderGetViewRectEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; var rect: TCefRect);
		browser := AsCefBrowserRef(getVal(0))
		rect := (*TCefRect)(getPtr(getVal(1)))
		cb(browser, rect)
		*(*TCefRect)(getPtr(getVal(1))) = *rect
	})
}

func makeTOnRenderIMECompositionRangeChangedEvent(cb TOnRenderIMECompositionRangeChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderIMECompositionRangeChangedEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const selected_range: PCefRange; character_boundsCount: NativeUInt; const character_bounds: PCefRect);
		browser := AsCefBrowserRef(getVal(0))
		selectedRange := *(*TCefRange)(getPtr(getVal(1)))
		characterBoundsCount := cefTypes.NativeUInt(getVal(2))
		characterBounds := *(*TCefRect)(getPtr(getVal(3)))
		cb(browser, selectedRange, characterBoundsCount, characterBounds)
	})
}

func makeTOnRenderLoadEnd(cb TOnRenderLoadEnd) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderLoadEnd", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		httpStatusCode := int32(getVal(2))
		cb(browser, frame, httpStatusCode)
	})
}

func makeTOnRenderLoadError(cb TOnRenderLoadError) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderLoadError", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText, failedUrl: ustring);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		errorCode := cefTypes.TCefErrorCode(getVal(2))
		errorText := api.GoStr(getVal(3))
		failedUrl := api.GoStr(getVal(4))
		cb(browser, frame, errorCode, errorText, failedUrl)
	})
}

func makeTOnRenderLoadStart(cb TOnRenderLoadStart) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderLoadStart", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		transitionType := cefTypes.TCefTransitionType(getVal(2))
		cb(browser, frame, transitionType)
	})
}

func makeTOnRenderLoadingStateChange(cb TOnRenderLoadingStateChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderLoadingStateChange", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; isLoading, canGoBack, canGoForward: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		isLoading := api.GoBool(getVal(1))
		canGoBack := api.GoBool(getVal(2))
		canGoForward := api.GoBool(getVal(3))
		cb(browser, isLoading, canGoBack, canGoForward)
	})
}

func makeTOnRenderPaintEvent(cb TOnRenderPaintEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderPaintEvent", func(getVal func(i int) uintptr) {
		// 7 : procedure(const browser: ICefBrowser; kind: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; const buffer: Pointer; width: Integer; height: Integer);
		browser := AsCefBrowserRef(getVal(0))
		kind := cefTypes.TCefPaintElementType(getVal(1))
		dirtyRectsCount := cefTypes.NativeUInt(getVal(2))
		dirtyRectsPtr := getVal(3)
		dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
		buffer := uintptr(getVal(4))
		width := int32(getVal(5))
		height := int32(getVal(6))
		cb(browser, kind, dirtyRectsCount, dirtyRects, buffer, width, height)
	})
}

func makeTOnRenderPopupShowEvent(cb TOnRenderPopupShowEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderPopupShowEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; show: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		show := api.GoBool(getVal(1))
		cb(browser, show)
	})
}

func makeTOnRenderPopupSizeEvent(cb TOnRenderPopupSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderPopupSizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const rect: PCefRect);
		browser := AsCefBrowserRef(getVal(0))
		rect := *(*TCefRect)(getPtr(getVal(1)))
		cb(browser, rect)
	})
}

func makeTOnRenderProcessBrowserCreatedEvent(cb TOnRenderProcessBrowserCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessBrowserCreatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const extra_info: ICefDictionaryValue);
		browser := AsCefBrowserRef(getVal(0))
		extraInfo := AsCefDictionaryValueRef(getVal(1))
		cb(browser, extraInfo)
	})
}

func makeTOnRenderProcessBrowserDestroyedEvent(cb TOnRenderProcessBrowserDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessBrowserDestroyedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnRenderProcessContextCreatedEvent(cb TOnRenderProcessContextCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessContextCreatedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		cb(browser, frame, context)
	})
}

func makeTOnRenderProcessContextReleasedEvent(cb TOnRenderProcessContextReleasedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessContextReleasedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		cb(browser, frame, context)
	})
}

func makeTOnRenderProcessFocusedNodeChangedEvent(cb TOnRenderProcessFocusedNodeChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessFocusedNodeChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const node: ICefDomNode);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		node := AsCefDomNodeRef(getVal(2))
		cb(browser, frame, node)
	})
}

func makeTOnRenderProcessGetLoadHandlerEvent(cb TOnRenderProcessGetLoadHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessGetLoadHandlerEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): ICefLoadHandler;
		ret := cb()
		if ret != nil {
			*(*uintptr)(getPtr(getVal(0))) = ret.Instance()
		}
	})
}

func makeTOnRenderProcessProcessMessageReceivedEvent(cb TOnRenderProcessProcessMessageReceivedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessProcessMessageReceivedEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const aMessage: ICefProcessMessage): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		sourceProcess := cefTypes.TCefProcessId(getVal(2))
		message := AsCefProcessMessageRef(getVal(3))
		ret := cb(browser, frame, sourceProcess, message)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnRenderProcessTerminated(cb TOnRenderProcessTerminated) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessTerminated", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; status: TCefTerminationStatus);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		status := cefTypes.TCefTerminationStatus(getVal(2))
		cb(sender, browser, status)
	})
}

func makeTOnRenderProcessUncaughtExceptionEvent(cb TOnRenderProcessUncaughtExceptionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessUncaughtExceptionEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context; const V8Exception: ICefV8Exception; const stackTrace: ICefV8StackTrace);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		v8Exception := AsCefV8ExceptionRef(getVal(3))
		stackTrace := AsCefV8StackTraceRef(getVal(4))
		cb(browser, frame, context, v8Exception, stackTrace)
	})
}

func makeTOnRenderProcessWebKitInitializedEvent(cb TOnRenderProcessWebKitInitializedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderProcessWebKitInitializedEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnRenderScrollOffsetChangedEvent(cb TOnRenderScrollOffsetChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderScrollOffsetChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; x: Double; y: Double);
		browser := AsCefBrowserRef(getVal(0))
		X := *(*float64)(getPtr(getVal(1)))
		Y := *(*float64)(getPtr(getVal(2)))
		cb(browser, X, Y)
	})
}

func makeTOnRenderStartDraggingEvent(cb TOnRenderStartDraggingEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderStartDraggingEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const dragData: ICefDragData; allowedOps: TCefDragOperations; x: Integer; y: Integer): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		dragData := AsCefDragDataRef(getVal(1))
		allowedOps := cefTypes.TCefDragOperations(getVal(2))
		X := int32(getVal(3))
		Y := int32(getVal(4))
		ret := cb(browser, dragData, allowedOps, X, Y)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnRenderTextSelectionChangedEvent(cb TOnRenderTextSelectionChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderTextSelectionChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const browser: ICefBrowser; const selected_text: ustring; const selected_range: PCefRange);
		browser := AsCefBrowserRef(getVal(0))
		selectedText := api.GoStr(getVal(1))
		selectedRange := *(*TCefRange)(getPtr(getVal(2)))
		cb(browser, selectedText, selectedRange)
	})
}

func makeTOnRenderTouchHandleStateChangedEvent(cb TOnRenderTouchHandleStateChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderTouchHandleStateChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; const state: TCefTouchHandleState);
		browser := AsCefBrowserRef(getVal(0))
		state := *(*TCefTouchHandleState)(getPtr(getVal(1)))
		cb(browser, state)
	})
}

func makeTOnRenderUpdateDragCursorEvent(cb TOnRenderUpdateDragCursorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderUpdateDragCursorEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; operation: TCefDragOperation);
		browser := AsCefBrowserRef(getVal(0))
		operation := cefTypes.TCefDragOperation(getVal(1))
		cb(browser, operation)
	})
}

func makeTOnRenderViewReady(cb TOnRenderViewReady) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderViewReady", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: Tobject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnRenderVirtualKeyboardRequestedEvent(cb TOnRenderVirtualKeyboardRequestedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRenderVirtualKeyboardRequestedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; input_mode: TCefTextInpuMode);
		browser := AsCefBrowserRef(getVal(0))
		inputMode := cefTypes.TCefTextInpuMode(getVal(1))
		cb(browser, inputMode)
	})
}

func makeTOnRequestBeforeBrowseEvent(cb TOnRequestBeforeBrowseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestBeforeBrowseEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; user_gesture: Boolean; isRedirect: Boolean): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		userGesture := api.GoBool(getVal(3))
		isRedirect := api.GoBool(getVal(4))
		ret := cb(browser, frame, request, userGesture, isRedirect)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnRequestCertificateErrorEvent(cb TOnRequestCertificateErrorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestCertificateErrorEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; certError: TCefErrorcode; const requestUrl: ustring; const sslInfo: ICefSslInfo; const callback: ICefCallback): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		certError := cefTypes.TCefErrorCode(getVal(1))
		requestUrl := api.GoStr(getVal(2))
		sslInfo := AsCefSslInfoRef(getVal(3))
		callback := AsCefCallbackRef(getVal(4))
		ret := cb(browser, certError, requestUrl, sslInfo, callback)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnRequestComplete(cb TOnRequestComplete) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestComplete", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const request: ICefUrlRequest);
		sender := lcl.AsObject(getVal(0))
		request := AsCefUrlRequestRef(getVal(1))
		cb(sender, request)
	})
}

func makeTOnRequestContextGetResourceRequestHandlerEvent(cb TOnRequestContextGetResourceRequestHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestContextGetResourceRequestHandlerEvent", func(getVal func(i int) uintptr) {
		// 8 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; is_navigation: boolean; is_download: boolean; const request_initiator: ustring; var disable_default_handling: boolean; var aResourceRequestHandler: ICefResourceRequestHandler);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		isNavigation := api.GoBool(getVal(3))
		isDownload := api.GoBool(getVal(4))
		requestInitiator := api.GoStr(getVal(5))
		disableDefaultHandling := (*bool)(getPtr(getVal(6)))
		var resourceRequestHandler IEngResourceRequestHandler
		resourceRequestHandler = AsEngResourceRequestHandler(*(*uintptr)(getPtr(getVal(7))))
		cb(browser, frame, request, isNavigation, isDownload, requestInitiator, disableDefaultHandling, &resourceRequestHandler)
		if resourceRequestHandler != nil {
			*(*uintptr)(getPtr(getVal(7))) = resourceRequestHandler.Instance()
		}
	})
}

func makeTOnRequestContextInitialized(cb TOnRequestContextInitialized) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestContextInitialized", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const request_context: ICefRequestContext);
		sender := lcl.AsObject(getVal(0))
		requestContext := AsCefRequestContextRef(getVal(1))
		cb(sender, requestContext)
	})
}

func makeTOnRequestContextRequestContextInitializedEvent(cb TOnRequestContextRequestContextInitializedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestContextRequestContextInitializedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const request_context: ICefRequestContext);
		requestContext := AsCefRequestContextRef(getVal(0))
		cb(requestContext)
	})
}

func makeTOnRequestDocumentAvailableInMainFrameEvent(cb TOnRequestDocumentAvailableInMainFrameEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestDocumentAvailableInMainFrameEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnRequestGetAuthCredentialsEvent(cb TOnRequestGetAuthCredentialsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestGetAuthCredentialsEvent", func(getVal func(i int) uintptr) {
		// 8 : function(const browser: ICefBrowser; const originUrl: ustring; isProxy: Boolean; const host: ustring; port: Integer; const realm: ustring; const scheme: ustring; const callback: ICefAuthCallback): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		originUrl := api.GoStr(getVal(1))
		isProxy := api.GoBool(getVal(2))
		host := api.GoStr(getVal(3))
		port := int32(getVal(4))
		realm := api.GoStr(getVal(5))
		scheme := api.GoStr(getVal(6))
		callback := AsCefAuthCallbackRef(getVal(7))
		ret := cb(browser, originUrl, isProxy, host, port, realm, scheme, callback)
		*(*bool)(getPtr(getVal(8))) = ret
	})
}

func makeTOnRequestGetResourceRequestHandlerEvent(cb TOnRequestGetResourceRequestHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestGetResourceRequestHandlerEvent", func(getVal func(i int) uintptr) {
		// 8 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; is_navigation: boolean; is_download: boolean; const request_initiator: ustring; var disable_default_handling: boolean; var aResourceRequestHandler: ICefResourceRequestHandler);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		isNavigation := api.GoBool(getVal(3))
		isDownload := api.GoBool(getVal(4))
		requestInitiator := api.GoStr(getVal(5))
		disableDefaultHandling := (*bool)(getPtr(getVal(6)))
		var resourceRequestHandler IEngResourceRequestHandler
		resourceRequestHandler = AsEngResourceRequestHandler(*(*uintptr)(getPtr(getVal(7))))
		cb(browser, frame, request, isNavigation, isDownload, requestInitiator, disableDefaultHandling, &resourceRequestHandler)
		if resourceRequestHandler != nil {
			*(*uintptr)(getPtr(getVal(7))) = resourceRequestHandler.Instance()
		}
	})
}

func makeTOnRequestMediaAccessPermissionEvent(cb TOnRequestMediaAccessPermissionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestMediaAccessPermissionEvent", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefMediaAccessCallback; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		requestingOrigin := api.GoStr(getVal(3))
		requestedPermissions := uint32(getVal(4))
		callback := AsCefMediaAccessCallbackRef(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, requestingOrigin, requestedPermissions, callback, result)
	})
}

func makeTOnRequestOpenUrlFromTabEvent(cb TOnRequestOpenUrlFromTabEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestOpenUrlFromTabEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const targetUrl: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		targetUrl := api.GoStr(getVal(2))
		targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(3))
		userGesture := api.GoBool(getVal(4))
		ret := cb(browser, frame, targetUrl, targetDisposition, userGesture)
		*(*bool)(getPtr(getVal(5))) = ret
	})
}

func makeTOnRequestRenderProcessTerminatedEvent(cb TOnRequestRenderProcessTerminatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestRenderProcessTerminatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const browser: ICefBrowser; status: TCefTerminationStatus);
		browser := AsCefBrowserRef(getVal(0))
		status := cefTypes.TCefTerminationStatus(getVal(1))
		cb(browser, status)
	})
}

func makeTOnRequestRenderViewReadyEvent(cb TOnRequestRenderViewReadyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestRenderViewReadyEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const browser: ICefBrowser);
		browser := AsCefBrowserRef(getVal(0))
		cb(browser)
	})
}

func makeTOnRequestSelectClientCertificateEvent(cb TOnRequestSelectClientCertificateEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRequestSelectClientCertificateEvent", func(getVal func(i int) uintptr) {
		// 7 : function(const browser: ICefBrowser; isProxy: boolean; const host: ustring; port: integer; certificatesCount: NativeUInt; const certificates: TCefX509CertificateArray; const callback: ICefSelectClientCertificateCallback): boolean;
		browser := AsCefBrowserRef(getVal(0))
		isProxy := api.GoBool(getVal(1))
		host := api.GoStr(getVal(2))
		port := int32(getVal(3))
		certificatesCount := cefTypes.NativeUInt(getVal(4))
		certificatesPtr := getVal(5)
		certificatesLen := int32(getVal(6))
		certificates := NewCefX509CertificateArray(int(certificatesLen), certificatesPtr)
		callback := AsCefSelectClientCertificateCallbackRef(getVal(7))
		ret := cb(browser, isProxy, host, port, certificatesCount, certificates, callback)
		*(*bool)(getPtr(getVal(8))) = ret
	})
}

func makeTOnResetDialogState(cb TOnResetDialogState) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResetDialogState", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		cb(sender, browser)
	})
}

func makeTOnResolveCallbackResolveCompletedEvent(cb TOnResolveCallbackResolveCompletedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResolveCallbackResolveCompletedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(result: TCefErrorCode; const resolvedIps: TStrings);
		result := cefTypes.TCefErrorCode(getVal(0))
		resolvedIps := lcl.AsStrings(getVal(1))
		cb(result, resolvedIps)
	})
}

func makeTOnResolvedIPsAvailableEvent(cb TOnResolvedIPsAvailableEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResolvedIPsAvailableEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; result: TCefErrorCode; const resolvedIps: TStrings);
		sender := lcl.AsObject(getVal(0))
		result := cefTypes.TCefErrorCode(getVal(1))
		resolvedIps := lcl.AsStrings(getVal(2))
		cb(sender, result, resolvedIps)
	})
}

func makeTOnResourceBundleGetDataResourceEvent(cb TOnResourceBundleGetDataResourceEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceBundleGetDataResourceEvent", func(getVal func(i int) uintptr) {
		// 3 : function(resourceId: Integer; var data: Pointer; var dataSize: NativeUInt): Boolean;
		resourceId := int32(getVal(0))
		data := (*uintptr)(getPtr(getVal(1)))
		dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
		ret := cb(resourceId, data, dataSize)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnResourceBundleGetDataResourceForScaleEvent(cb TOnResourceBundleGetDataResourceForScaleEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceBundleGetDataResourceForScaleEvent", func(getVal func(i int) uintptr) {
		// 4 : function(resourceId: Integer; scaleFactor: TCefScaleFactor; var data: Pointer; var dataSize: NativeUInt): Boolean;
		resourceId := int32(getVal(0))
		scaleFactor := cefTypes.TCefScaleFactor(getVal(1))
		data := (*uintptr)(getPtr(getVal(2)))
		dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(3)))
		ret := cb(resourceId, scaleFactor, data, dataSize)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnResourceBundleGetLocalizedStringEvent(cb TOnResourceBundleGetLocalizedStringEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceBundleGetLocalizedStringEvent", func(getVal func(i int) uintptr) {
		// 2 : function(stringid: Integer; var stringVal: ustring): Boolean;
		stringid := int32(getVal(0))
		stringVal := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
		ret := cb(stringid, &stringVal)
		*(*uintptr)(getPtr(getVal(1))) = api.PasStr(stringVal)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnResourceCancelEvent(cb TOnResourceCancelEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceCancelEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnResourceGetResponseHeadersEvent(cb TOnResourceGetResponseHeadersEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceGetResponseHeadersEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const response: ICefResponse; out responseLength: Int64; out redirectUrl: ustring);
		response := AsCefResponseRef(getVal(0))
		responseLength := (*int64)(getPtr(getVal(1)))
		var redirectUrl string
		cb(response, responseLength, &redirectUrl)
		*(*uintptr)(getPtr(getVal(2))) = api.PasStr(redirectUrl)
	})
}

func makeTOnResourceLoadComplete(cb TOnResourceLoadComplete) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceLoadComplete", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; status: TCefUrlRequestStatus; receivedContentLength: Int64);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		response := AsCefResponseRef(getVal(4))
		status := cefTypes.TCefUrlRequestStatus(getVal(5))
		receivedContentLength := *(*int64)(getPtr(getVal(6)))
		cb(sender, browser, frame, request, response, status, receivedContentLength)
	})
}

func makeTOnResourceOpenEvent(cb TOnResourceOpenEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceOpenEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const request: ICefRequest; var handle_request: boolean; const callback: ICefCallback): boolean;
		request := AsCefRequestRef(getVal(0))
		handleRequest := (*bool)(getPtr(getVal(1)))
		callback := AsCefCallbackRef(getVal(2))
		ret := cb(request, handleRequest, callback)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnResourceProcessRequestEvent(cb TOnResourceProcessRequestEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceProcessRequestEvent", func(getVal func(i int) uintptr) {
		// 2 : function(const request: ICefRequest; const callback: ICefCallback): Boolean;
		request := AsCefRequestRef(getVal(0))
		callback := AsCefCallbackRef(getVal(1))
		ret := cb(request, callback)
		*(*bool)(getPtr(getVal(2))) = ret
	})
}

func makeTOnResourceReadEvent(cb TOnResourceReadEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceReadEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const data_out: Pointer; bytes_to_read: Integer; var bytes_read: Integer; const callback: ICefResourceReadCallback): boolean;
		dataOut := uintptr(getVal(0))
		bytesToRead := int32(getVal(1))
		bytesRead := (*int32)(getPtr(getVal(2)))
		callback := AsCefResourceReadCallbackRef(getVal(3))
		ret := cb(dataOut, bytesToRead, bytesRead, callback)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnResourceReadResponseEvent(cb TOnResourceReadResponseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceReadResponseEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const dataOut: Pointer; bytesToRead: Integer; var bytesRead: Integer; const callback: ICefCallback): Boolean;
		dataOut := uintptr(getVal(0))
		bytesToRead := int32(getVal(1))
		bytesRead := (*int32)(getPtr(getVal(2)))
		callback := AsCefCallbackRef(getVal(3))
		ret := cb(dataOut, bytesToRead, bytesRead, callback)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnResourceRedirect(cb TOnResourceRedirect) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRedirect", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; var newUrl: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		response := AsCefResponseRef(getVal(4))
		newUrl := api.GoStr(*(*uintptr)(getPtr(getVal(5))))
		cb(sender, browser, frame, request, response, &newUrl)
		*(*uintptr)(getPtr(getVal(5))) = api.PasStr(newUrl)
	})
}

func makeTOnResourceRequestBeforeResourceLoadEvent(cb TOnResourceRequestBeforeResourceLoadEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestBeforeResourceLoadEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const callback: ICefCallback): TCefReturnValue;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		callback := AsCefCallbackRef(getVal(3))
		ret := cb(browser, frame, request, callback)
		*(*cefTypes.TCefReturnValue)(getPtr(getVal(4))) = ret
	})
}

func makeTOnResourceRequestGetCookieAccessFilterEvent(cb TOnResourceRequestGetCookieAccessFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestGetCookieAccessFilterEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var aFilter: ICefCookieAccessFilter);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		var filter IEngCookieAccessFilter
		filter = AsEngCookieAccessFilter(*(*uintptr)(getPtr(getVal(3))))
		cb(browser, frame, request, &filter)
		if filter != nil {
			*(*uintptr)(getPtr(getVal(3))) = filter.Instance()
		}
	})
}

func makeTOnResourceRequestGetResourceHandlerEvent(cb TOnResourceRequestGetResourceHandlerEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestGetResourceHandlerEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var aResourceHandler: ICefResourceHandler);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		var resourceHandler IEngResourceHandler
		resourceHandler = AsEngResourceHandler(*(*uintptr)(getPtr(getVal(3))))
		cb(browser, frame, request, &resourceHandler)
		if resourceHandler != nil {
			*(*uintptr)(getPtr(getVal(3))) = resourceHandler.Instance()
		}
	})
}

func makeTOnResourceRequestGetResourceResponseFilterEvent(cb TOnResourceRequestGetResourceResponseFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestGetResourceResponseFilterEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; var aResponseFilter: ICefResponseFilter);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		response := AsCefResponseRef(getVal(3))
		var responseFilter IEngResponseFilter
		responseFilter = AsEngResponseFilter(*(*uintptr)(getPtr(getVal(4))))
		cb(browser, frame, request, response, &responseFilter)
		if responseFilter != nil {
			*(*uintptr)(getPtr(getVal(4))) = responseFilter.Instance()
		}
	})
}

func makeTOnResourceRequestProtocolExecutionEvent(cb TOnResourceRequestProtocolExecutionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestProtocolExecutionEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var allowOsExecution: Boolean);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		allowOsExecution := (*bool)(getPtr(getVal(3)))
		cb(browser, frame, request, allowOsExecution)
	})
}

func makeTOnResourceRequestResourceLoadCompleteEvent(cb TOnResourceRequestResourceLoadCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestResourceLoadCompleteEvent", func(getVal func(i int) uintptr) {
		// 6 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; status: TCefUrlRequestStatus; receivedContentLength: Int64);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		response := AsCefResponseRef(getVal(3))
		status := cefTypes.TCefUrlRequestStatus(getVal(4))
		receivedContentLength := *(*int64)(getPtr(getVal(5)))
		cb(browser, frame, request, response, status, receivedContentLength)
	})
}

func makeTOnResourceRequestResourceRedirectEvent(cb TOnResourceRequestResourceRedirectEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestResourceRedirectEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; var newUrl: ustring);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		response := AsCefResponseRef(getVal(3))
		newUrl := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
		cb(browser, frame, request, response, &newUrl)
		*(*uintptr)(getPtr(getVal(4))) = api.PasStr(newUrl)
	})
}

func makeTOnResourceRequestResourceResponseEvent(cb TOnResourceRequestResourceResponseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceRequestResourceResponseEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse): Boolean;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		request := AsCefRequestRef(getVal(2))
		response := AsCefResponseRef(getVal(3))
		ret := cb(browser, frame, request, response)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnResourceResponse(cb TOnResourceResponse) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceResponse", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		request := AsCefRequestRef(getVal(3))
		response := AsCefResponseRef(getVal(4))
		result := (*bool)(getPtr(getVal(5)))
		cb(sender, browser, frame, request, response, result)
	})
}

func makeTOnResourceSkipEvent(cb TOnResourceSkipEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResourceSkipEvent", func(getVal func(i int) uintptr) {
		// 3 : function(bytes_to_skip: int64; var bytes_skipped: Int64; const callback: ICefResourceSkipCallback): boolean;
		bytesToSkip := *(*int64)(getPtr(getVal(0)))
		bytesSkipped := (*int64)(getPtr(getVal(1)))
		callback := AsCefResourceSkipCallbackRef(getVal(2))
		ret := cb(bytesToSkip, bytesSkipped, callback)
		*(*bool)(getPtr(getVal(3))) = ret
	})
}

func makeTOnResponseFilterFilterEvent(cb TOnResponseFilterFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResponseFilterFilterEvent", func(getVal func(i int) uintptr) {
		// 6 : function(data_in: Pointer; data_in_size: NativeUInt; var data_in_read: NativeUInt; data_out: Pointer; data_out_size: NativeUInt; var data_out_written: NativeUInt): TCefResponseFilterStatus;
		dataIn := uintptr(getVal(0))
		dataInSize := cefTypes.NativeUInt(getVal(1))
		dataInRead := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
		dataOut := uintptr(getVal(3))
		dataOutSize := cefTypes.NativeUInt(getVal(4))
		dataOutWritten := (*cefTypes.NativeUInt)(getPtr(getVal(5)))
		ret := cb(dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten)
		*(*cefTypes.TCefResponseFilterStatus)(getPtr(getVal(6))) = ret
	})
}

func makeTOnResponseFilterInitFilterEvent(cb TOnResponseFilterInitFilterEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnResponseFilterInitFilterEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): Boolean;
		ret := cb()
		*(*bool)(getPtr(getVal(0))) = ret
	})
}

func makeTOnRouteMessageReceivedEvent(cb TOnRouteMessageReceivedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRouteMessageReceivedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const route: ICefMediaRoute; const message_: ustring);
		sender := lcl.AsObject(getVal(0))
		route := AsCefMediaRouteRef(getVal(1))
		message := api.GoStr(getVal(2))
		cb(sender, route, message)
	})
}

func makeTOnRouteStateChangedEvent(cb TOnRouteStateChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRouteStateChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const route: ICefMediaRoute; state: TCefMediaRouteConnectionState);
		sender := lcl.AsObject(getVal(0))
		route := AsCefMediaRouteRef(getVal(1))
		state := cefTypes.TCefMediaRouteConnectionState(getVal(2))
		cb(sender, route, state)
	})
}

func makeTOnRoutesEvent(cb TOnRoutesEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRoutesEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const routes: TCefMediaRouteArray);
		sender := lcl.AsObject(getVal(0))
		routesPtr := getVal(1)
		routesLen := int32(getVal(2))
		routes := NewCefMediaRouteArray(int(routesLen), routesPtr)
		cb(sender, routes)
	})
}

func makeTOnRunContextMenu(cb TOnRunContextMenu) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRunContextMenu", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel; const callback: ICefRunContextMenuCallback; var aResult : Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		params := AsCefContextMenuParamsRef(getVal(3))
		model := AsCefMenuModelRef(getVal(4))
		callback := AsCefRunContextMenuCallbackRef(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, frame, params, model, callback, result)
	})
}

func makeTOnRunFileDialogCallbackFileDialogDismissedEvent(cb TOnRunFileDialogCallbackFileDialogDismissedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRunFileDialogCallbackFileDialogDismissedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const filePaths: TStrings);
		filePaths := lcl.AsStrings(getVal(0))
		cb(filePaths)
	})
}

func makeTOnRunQuickMenuEvent(cb TOnRunQuickMenuEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnRunQuickMenuEvent", func(getVal func(i int) uintptr) {
		// 8 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; location: TCefPoint; size: TCefSize; edit_state_flags: TCefQuickMenuEditStateFlags; const callback: ICefRunQuickMenuCallback; var aResult : Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		frame := AsCefFrameRef(getVal(2))
		location := *(*TCefPoint)(getPtr(getVal(3)))
		size := *(*TCefSize)(getPtr(getVal(4)))
		editStateFlags := cefTypes.TCefQuickMenuEditStateFlags(getVal(5))
		callback := AsCefRunQuickMenuCallbackRef(getVal(6))
		result := (*bool)(getPtr(getVal(7)))
		cb(sender, browser, frame, location, size, editStateFlags, callback, result)
	})
}

func makeTOnScheduleMessagePumpWorkEvent(cb TOnScheduleMessagePumpWorkEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnScheduleMessagePumpWorkEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const delayMs: Int64);
		delayMs := *(*int64)(getPtr(getVal(0)))
		cb(delayMs)
	})
}

func makeTOnSchemeFactoryNewEvent(cb TOnSchemeFactoryNewEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnSchemeFactoryNewEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const schemeName: ustring; const request: ICefRequest): ICefResourceHandler;
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		schemeName := api.GoStr(getVal(2))
		request := AsCefRequestRef(getVal(3))
		ret := cb(browser, frame, schemeName, request)
		if ret != nil {
			*(*uintptr)(getPtr(getVal(4))) = ret.Instance()
		}
	})
}

func makeTOnScrollOffsetChanged(cb TOnScrollOffsetChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnScrollOffsetChanged", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; x, y: Double);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		X := *(*float64)(getPtr(getVal(2)))
		Y := *(*float64)(getPtr(getVal(3)))
		cb(sender, browser, X, Y)
	})
}

func makeTOnSelectClientCertificate(cb TOnSelectClientCertificate) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnSelectClientCertificate", func(getVal func(i int) uintptr) {
		// 9 : procedure(Sender: TObject; const browser: ICefBrowser; isProxy: boolean; const host: ustring; port: integer; certificatesCount: NativeUInt; const certificates: TCefX509CertificateArray; const callback: ICefSelectClientCertificateCallback; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		isProxy := api.GoBool(getVal(2))
		host := api.GoStr(getVal(3))
		port := int32(getVal(4))
		certificatesCount := cefTypes.NativeUInt(getVal(5))
		certificatesPtr := getVal(6)
		certificates := NewCefX509CertificateArray(int(certificatesCount), certificatesPtr)
		callback := AsCefSelectClientCertificateCallbackRef(getVal(7))
		result := (*bool)(getPtr(getVal(8)))
		cb(sender, browser, isProxy, host, port, certificatesCount, certificates, callback, result)
	})
}

func makeTOnServerClientConnectedEvent(cb TOnServerClientConnectedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerClientConnectedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const server: ICefServer; connection_id: Integer);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		cb(server, connectionId)
	})
}

func makeTOnServerClientDisconnectedEvent(cb TOnServerClientDisconnectedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerClientDisconnectedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const server: ICefServer; connection_id: Integer);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		cb(server, connectionId)
	})
}

func makeTOnServerCreated(cb TOnServerCreated) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerCreated", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const server: ICefServer);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		cb(sender, server)
	})
}

func makeTOnServerDestroyed(cb TOnServerDestroyed) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerDestroyed", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const server: ICefServer);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		cb(sender, server)
	})
}

func makeTOnServerHttpRequestEvent(cb TOnServerHttpRequestEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerHttpRequestEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		clientAddress := api.GoStr(getVal(2))
		request := AsCefRequestRef(getVal(3))
		cb(server, connectionId, clientAddress, request)
	})
}

func makeTOnServerServerCreatedEvent(cb TOnServerServerCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerServerCreatedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const server: ICefServer);
		server := AsCEFServerRef(getVal(0))
		cb(server)
	})
}

func makeTOnServerServerDestroyedEvent(cb TOnServerServerDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerServerDestroyedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const server: ICefServer);
		server := AsCEFServerRef(getVal(0))
		cb(server)
	})
}

func makeTOnServerWebSocketConnectedEvent(cb TOnServerWebSocketConnectedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerWebSocketConnectedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const server: ICefServer; connection_id: Integer);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		cb(server, connectionId)
	})
}

func makeTOnServerWebSocketMessageEvent(cb TOnServerWebSocketMessageEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerWebSocketMessageEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const server: ICefServer; connection_id: Integer; const data: Pointer; data_size: NativeUInt);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		data := uintptr(getVal(2))
		dataSize := cefTypes.NativeUInt(getVal(3))
		cb(server, connectionId, data, dataSize)
	})
}

func makeTOnServerWebSocketRequestEvent(cb TOnServerWebSocketRequestEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnServerWebSocketRequestEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest; const callback: ICefCallback);
		server := AsCEFServerRef(getVal(0))
		connectionId := int32(getVal(1))
		clientAddress := api.GoStr(getVal(2))
		request := AsCefRequestRef(getVal(3))
		callback := AsCefCallbackRef(getVal(4))
		cb(server, connectionId, clientAddress, request, callback)
	})
}

func makeTOnSetCookieCallbackCompleteEvent(cb TOnSetCookieCallbackCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnSetCookieCallbackCompleteEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(success: Boolean);
		success := api.GoBool(getVal(0))
		cb(success)
	})
}

func makeTOnSetFocus(cb TOnSetFocus) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnSetFocus", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; source: TCefFocusSource; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		source := cefTypes.TCefFocusSource(getVal(2))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, source, result)
	})
}

func makeTOnShowPermissionPromptEvent(cb TOnShowPermissionPromptEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnShowPermissionPromptEvent", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; prompt_id: uint64; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefPermissionPromptCallback; var aResult: boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		promptId := *(*uint64)(getPtr(getVal(2)))
		requestingOrigin := api.GoStr(getVal(3))
		requestedPermissions := uint32(getVal(4))
		callback := AsCefPermissionPromptCallbackRef(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, promptId, requestingOrigin, requestedPermissions, callback, result)
	})
}

func makeTOnSinksEvent(cb TOnSinksEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnSinksEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const sinks: TCefMediaSinkArray);
		sender := lcl.AsObject(getVal(0))
		sinksPtr := getVal(1)
		sinksLen := int32(getVal(2))
		sinks := NewCefMediaSinkArray(int(sinksLen), sinksPtr)
		cb(sender, sinks)
	})
}

func makeTOnStartDragging(cb TOnStartDragging) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnStartDragging", func(getVal func(i int) uintptr) {
		// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const dragData: ICefDragData; allowedOps: TCefDragOperations; x, y: Integer; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		dragData := AsCefDragDataRef(getVal(2))
		allowedOps := cefTypes.TCefDragOperations(getVal(3))
		X := int32(getVal(4))
		Y := int32(getVal(5))
		result := (*bool)(getPtr(getVal(6)))
		cb(sender, browser, dragData, allowedOps, X, Y, result)
	})
}

func makeTOnStatusMessage(cb TOnStatusMessage) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnStatusMessage", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const value: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		value := api.GoStr(getVal(2))
		cb(sender, browser, value)
	})
}

func makeTOnStringVisitorVisitEvent(cb TOnStringVisitorVisitEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnStringVisitorVisitEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const str: ustring);
		str := api.GoStr(getVal(0))
		cb(str)
	})
}

func makeTOnTakeFocus(cb TOnTakeFocus) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTakeFocus", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; next: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		next := api.GoBool(getVal(2))
		cb(sender, browser, next)
	})
}

func makeTOnTaskExecuteEvent(cb TOnTaskExecuteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTaskExecuteEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnTextResultAvailableEvent(cb TOnTextResultAvailableEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTextResultAvailableEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const aText : ustring);
		sender := lcl.AsObject(getVal(0))
		text := api.GoStr(getVal(1))
		cb(sender, text)
	})
}

func makeTOnTextSelectionChanged(cb TOnTextSelectionChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTextSelectionChanged", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const selected_text: ustring; const selected_range: PCefRange);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		selectedText := api.GoStr(getVal(2))
		selectedRange := *(*TCefRange)(getPtr(getVal(3)))
		cb(sender, browser, selectedText, selectedRange)
	})
}

func makeTOnTextfieldAfterUserActionEvent(cb TOnTextfieldAfterUserActionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTextfieldAfterUserActionEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const textfield: ICefTextfield);
		textfield := AsCefTextfieldRef(getVal(0))
		cb(textfield)
	})
}

func makeTOnTextfieldKeyEvent(cb TOnTextfieldKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTextfieldKeyEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const textfield: ICefTextfield; const event: TCefKeyEvent; var aResult: boolean);
		textfield := AsCefTextfieldRef(getVal(0))
		event := *(*TCefKeyEvent)(getPtr(getVal(1)))
		result := (*bool)(getPtr(getVal(2)))
		cb(textfield, event, result)
	})
}

func makeTOnTextfieldKeyEventEvent(cb TOnTextfieldKeyEventEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTextfieldKeyEventEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const textfield: ICefTextfield; const event: TCefKeyEvent; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		textfield := AsCefTextfieldRef(getVal(1))
		event := *(*TCefKeyEvent)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, textfield, event, result)
	})
}

func makeTOnTitleChange(cb TOnTitleChange) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTitleChange", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const title: ustring);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		title := api.GoStr(getVal(2))
		cb(sender, browser, title)
	})
}

func makeTOnTooltip(cb TOnTooltip) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTooltip", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var text: ustring; out Result: Boolean);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		text := api.GoStr(*(*uintptr)(getPtr(getVal(2))))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, browser, &text, result)
		*(*uintptr)(getPtr(getVal(2))) = api.PasStr(text)
	})
}

func makeTOnTouchHandleStateChanged(cb TOnTouchHandleStateChanged) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnTouchHandleStateChanged", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const state: TCefTouchHandleState);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		state := *(*TCefTouchHandleState)(getPtr(getVal(2)))
		cb(sender, browser, state)
	})
}

func makeTOnUncaughtExceptionEvent(cb TOnUncaughtExceptionEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUncaughtExceptionEvent", func(getVal func(i int) uintptr) {
		// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context; const exception: ICefV8Exception; const stackTrace: ICefV8StackTrace);
		browser := AsCefBrowserRef(getVal(0))
		frame := AsCefFrameRef(getVal(1))
		context := AsCefv8ContextRef(getVal(2))
		exception := AsCefV8ExceptionRef(getVal(3))
		stackTrace := AsCefV8StackTraceRef(getVal(4))
		cb(browser, frame, context, exception, stackTrace)
	})
}

func makeTOnUpdateDragCursor(cb TOnUpdateDragCursor) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUpdateDragCursor", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; operation: TCefDragOperation);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		operation := cefTypes.TCefDragOperation(getVal(2))
		cb(sender, browser, operation)
	})
}

func makeTOnUploadProgress(cb TOnUploadProgress) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUploadProgress", func(getVal func(i int) uintptr) {
		// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; current, total: Int64);
		sender := lcl.AsObject(getVal(0))
		request := AsCefUrlRequestRef(getVal(1))
		current := *(*int64)(getPtr(getVal(2)))
		total := *(*int64)(getPtr(getVal(3)))
		cb(sender, request, current, total)
	})
}

func makeTOnUrlrequestClientDownloadDataEvent(cb TOnUrlrequestClientDownloadDataEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUrlrequestClientDownloadDataEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const request: ICefUrlRequest; data: Pointer; dataLength: NativeUInt);
		request := AsCefUrlRequestRef(getVal(0))
		data := uintptr(getVal(1))
		dataLength := cefTypes.NativeUInt(getVal(2))
		cb(request, data, dataLength)
	})
}

func makeTOnUrlrequestClientDownloadProgressEvent(cb TOnUrlrequestClientDownloadProgressEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUrlrequestClientDownloadProgressEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const request: ICefUrlRequest; current: Int64; total: Int64);
		request := AsCefUrlRequestRef(getVal(0))
		current := *(*int64)(getPtr(getVal(1)))
		total := *(*int64)(getPtr(getVal(2)))
		cb(request, current, total)
	})
}

func makeTOnUrlrequestClientGetAuthCredentialsEvent(cb TOnUrlrequestClientGetAuthCredentialsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUrlrequestClientGetAuthCredentialsEvent", func(getVal func(i int) uintptr) {
		// 6 : function(isProxy: Boolean; const host: ustring; port: Integer; const realm: ustring; const scheme: ustring; const callback: ICefAuthCallback): Boolean;
		isProxy := api.GoBool(getVal(0))
		host := api.GoStr(getVal(1))
		port := int32(getVal(2))
		realm := api.GoStr(getVal(3))
		scheme := api.GoStr(getVal(4))
		callback := AsCefAuthCallbackRef(getVal(5))
		ret := cb(isProxy, host, port, realm, scheme, callback)
		*(*bool)(getPtr(getVal(6))) = ret
	})
}

func makeTOnUrlrequestClientRequestCompleteEvent(cb TOnUrlrequestClientRequestCompleteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUrlrequestClientRequestCompleteEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const request: ICefUrlRequest);
		request := AsCefUrlRequestRef(getVal(0))
		cb(request)
	})
}

func makeTOnUrlrequestClientUploadProgressEvent(cb TOnUrlrequestClientUploadProgressEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnUrlrequestClientUploadProgressEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const request: ICefUrlRequest; current: Int64; total: Int64);
		request := AsCefUrlRequestRef(getVal(0))
		current := *(*int64)(getPtr(getVal(1)))
		total := *(*int64)(getPtr(getVal(2)))
		cb(request, current, total)
	})
}

func makeTOnV8AccessorGetEvent(cb TOnV8AccessorGetEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8AccessorGetEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const name: ustring; const object_: ICefv8Value; var retval: ICefv8Value; var exception: ustring): Boolean;
		name := api.GoStr(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		retval := ICefv8Value(AsCefv8ValueRef(*(*uintptr)(getPtr(getVal(2)))))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(name, object, &retval, &exception)
		if retval != nil {
			*(*uintptr)(getPtr(getVal(2))) = retval.Instance()
		}
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnV8AccessorSet_Event(cb TOnV8AccessorSet_Event) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8AccessorSet_Event", func(getVal func(i int) uintptr) {
		// 4 : function(const name: ustring; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): Boolean;
		name := api.GoStr(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		value := AsCefv8ValueRef(getVal(2))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(name, object, value, &exception)
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnV8ArrayBufferReleaseCallbackReleaseBufferEvent(cb TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(buffer: Pointer);
		buffer := uintptr(getVal(0))
		cb(buffer)
	})
}

func makeTOnV8ExecuteEvent(cb TOnV8ExecuteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8ExecuteEvent", func(getVal func(i int) uintptr) {
		// 5 : function(const name: ustring; const object_: ICefv8Value; const arguments: TCefv8ValueArray; var retval: ICefv8Value; var exception: ustring): Boolean;
		name := api.GoStr(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		argumentsPtr := getVal(2)
		argumentsLen := int32(getVal(3))
		arguments := NewCefv8ValueArray(int(argumentsLen), argumentsPtr)
		retval := ICefv8Value(AsCefv8ValueRef(*(*uintptr)(getPtr(getVal(4)))))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(5))))
		ret := cb(name, object, arguments, &retval, &exception)
		if retval != nil {
			*(*uintptr)(getPtr(getVal(4))) = retval.Instance()
		}
		*(*uintptr)(getPtr(getVal(5))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(6))) = ret
	})
}

func makeTOnV8InterceptorGetByIndexEvent(cb TOnV8InterceptorGetByIndexEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8InterceptorGetByIndexEvent", func(getVal func(i int) uintptr) {
		// 4 : function(index: integer; const object_: ICefv8Value; var retval: ICefv8Value; var exception: ustring): boolean;
		index := int32(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		retval := ICefv8Value(AsCefv8ValueRef(*(*uintptr)(getPtr(getVal(2)))))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(index, object, &retval, &exception)
		if retval != nil {
			*(*uintptr)(getPtr(getVal(2))) = retval.Instance()
		}
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnV8InterceptorGetByNameEvent(cb TOnV8InterceptorGetByNameEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8InterceptorGetByNameEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const name: ustring; const object_: ICefv8Value; var retval: ICefv8Value; var exception: ustring): boolean;
		name := api.GoStr(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		retval := ICefv8Value(AsCefv8ValueRef(*(*uintptr)(getPtr(getVal(2)))))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(name, object, &retval, &exception)
		if retval != nil {
			*(*uintptr)(getPtr(getVal(2))) = retval.Instance()
		}
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnV8InterceptorSetByIndexEvent(cb TOnV8InterceptorSetByIndexEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8InterceptorSetByIndexEvent", func(getVal func(i int) uintptr) {
		// 4 : function(index: integer; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): boolean;
		index := int32(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		value := AsCefv8ValueRef(getVal(2))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(index, object, value, &exception)
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnV8InterceptorSetByNameEvent(cb TOnV8InterceptorSetByNameEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnV8InterceptorSetByNameEvent", func(getVal func(i int) uintptr) {
		// 4 : function(const name: ustring; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): boolean;
		name := api.GoStr(getVal(0))
		object := AsCefv8ValueRef(getVal(1))
		value := AsCefv8ValueRef(getVal(2))
		exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
		ret := cb(name, object, value, &exception)
		*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
		*(*bool)(getPtr(getVal(4))) = ret
	})
}

func makeTOnViewBlurEvent(cb TOnViewBlurEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewBlurEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const view: ICefView);
		view := AsCefViewRef(getVal(0))
		cb(view)
	})
}

func makeTOnViewChildViewChangedEvent(cb TOnViewChildViewChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewChildViewChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const view: ICefView; added: boolean; const child: ICefView);
		view := AsCefViewRef(getVal(0))
		added := api.GoBool(getVal(1))
		child := AsCefViewRef(getVal(2))
		cb(view, added, child)
	})
}

func makeTOnViewFocusEvent(cb TOnViewFocusEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewFocusEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const view: ICefView);
		view := AsCefViewRef(getVal(0))
		cb(view)
	})
}

func makeTOnViewGetHeightForWidthEvent(cb TOnViewGetHeightForWidthEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewGetHeightForWidthEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const view: ICefView; width: Integer; var aResult: Integer);
		view := AsCefViewRef(getVal(0))
		width := int32(getVal(1))
		result := (*int32)(getPtr(getVal(2)))
		cb(view, width, result)
	})
}

func makeTOnViewGetMaximumSizeEvent(cb TOnViewGetMaximumSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewGetMaximumSizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const view: ICefView; var aResult: TCefSize);
		view := AsCefViewRef(getVal(0))
		result := (*TCefSize)(getPtr(getVal(1)))
		cb(view, result)
		*(*TCefSize)(getPtr(getVal(1))) = *result
	})
}

func makeTOnViewGetMinimumSizeEvent(cb TOnViewGetMinimumSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewGetMinimumSizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const view: ICefView; var aResult: TCefSize);
		view := AsCefViewRef(getVal(0))
		result := (*TCefSize)(getPtr(getVal(1)))
		cb(view, result)
		*(*TCefSize)(getPtr(getVal(1))) = *result
	})
}

func makeTOnViewGetPreferredSizeEvent(cb TOnViewGetPreferredSizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewGetPreferredSizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const view: ICefView; var aResult: TCefSize);
		view := AsCefViewRef(getVal(0))
		result := (*TCefSize)(getPtr(getVal(1)))
		cb(view, result)
		*(*TCefSize)(getPtr(getVal(1))) = *result
	})
}

func makeTOnViewLayoutChangedEvent(cb TOnViewLayoutChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewLayoutChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const view: ICefView; new_bounds: TCefRect);
		view := AsCefViewRef(getVal(0))
		newBounds := *(*TCefRect)(getPtr(getVal(1)))
		cb(view, newBounds)
	})
}

func makeTOnViewParentViewChangedEvent(cb TOnViewParentViewChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewParentViewChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const view: ICefView; added: boolean; const parent: ICefView);
		view := AsCefViewRef(getVal(0))
		added := api.GoBool(getVal(1))
		parent := AsCefViewRef(getVal(2))
		cb(view, added, parent)
	})
}

func makeTOnViewWindowChangedEvent(cb TOnViewWindowChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnViewWindowChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const view: ICefView; added: boolean);
		view := AsCefViewRef(getVal(0))
		added := api.GoBool(getVal(1))
		cb(view, added)
	})
}

func makeTOnVirtualKeyboardRequested(cb TOnVirtualKeyboardRequested) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnVirtualKeyboardRequested", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const browser: ICefBrowser; input_mode: TCefTextInpuMode);
		sender := lcl.AsObject(getVal(0))
		browser := AsCefBrowserRef(getVal(1))
		inputMode := cefTypes.TCefTextInpuMode(getVal(2))
		cb(sender, browser, inputMode)
	})
}

func makeTOnWebKitInitializedEvent(cb TOnWebKitInitializedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWebKitInitializedEvent", func(getVal func(i int) uintptr) {
		// 0 : procedure();
		cb()
	})
}

func makeTOnWebSocketConnected(cb TOnWebSocketConnected) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWebSocketConnected", func(getVal func(i int) uintptr) {
		// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		cb(sender, server, connectionId)
	})
}

func makeTOnWebSocketMessage(cb TOnWebSocketMessage) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWebSocketMessage", func(getVal func(i int) uintptr) {
		// 5 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const data: Pointer; data_size: NativeUInt);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		data := uintptr(getVal(3))
		dataSize := cefTypes.NativeUInt(getVal(4))
		cb(sender, server, connectionId, data, dataSize)
	})
}

func makeTOnWebSocketRequest(cb TOnWebSocketRequest) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWebSocketRequest", func(getVal func(i int) uintptr) {
		// 6 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest; const callback: ICefCallback);
		sender := lcl.AsObject(getVal(0))
		server := AsCEFServerRef(getVal(1))
		connectionId := int32(getVal(2))
		clientAddress := api.GoStr(getVal(3))
		request := AsCefRequestRef(getVal(4))
		callback := AsCefCallbackRef(getVal(5))
		cb(sender, server, connectionId, clientAddress, request, callback)
	})
}

func makeTOnWindowAcceleratorEvent(cb TOnWindowAcceleratorEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowAcceleratorEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const window_: ICefWindow; command_id: Integer; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		commandId := int32(getVal(1))
		result := (*bool)(getPtr(getVal(2)))
		cb(window, commandId, result)
	})
}

func makeTOnWindowActivationChangedEvent(cb TOnWindowActivationChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowActivationChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; active: boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		active := api.GoBool(getVal(2))
		cb(sender, window, active)
	})
}

func makeTOnWindowBoundsChangedEvent(cb TOnWindowBoundsChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowBoundsChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const window_: ICefWindow; const new_bounds: TCefRect);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		newBounds := *(*TCefRect)(getPtr(getVal(2)))
		cb(sender, window, newBounds)
	})
}

func makeTOnWindowCanCloseEvent(cb TOnWindowCanCloseEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowCanCloseEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowCanMaximizeEvent(cb TOnWindowCanMaximizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowCanMaximizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowCanMinimizeEvent(cb TOnWindowCanMinimizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowCanMinimizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowCanResizeEvent(cb TOnWindowCanResizeEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowCanResizeEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowChangedEvent(cb TOnWindowChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowChangedEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const Sender: TObject; const view: ICefView; added: boolean);
		sender := lcl.AsObject(getVal(0))
		view := AsCefViewRef(getVal(1))
		added := api.GoBool(getVal(2))
		cb(sender, view, added)
	})
}

func makeTOnWindowClosingEvent(cb TOnWindowClosingEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowClosingEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		cb(sender, window)
	})
}

func makeTOnWindowCreatedEvent(cb TOnWindowCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowCreatedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		cb(sender, window)
	})
}

func makeTOnWindowDestroyedEvent(cb TOnWindowDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowDestroyedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		cb(sender, window)
	})
}

func makeTOnWindowGetInitialBoundsEvent(cb TOnWindowGetInitialBoundsEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowGetInitialBoundsEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: TCefRect);
		window := AsCefWindowRef(getVal(0))
		result := (*TCefRect)(getPtr(getVal(1)))
		cb(window, result)
		*(*TCefRect)(getPtr(getVal(1))) = *result
	})
}

func makeTOnWindowGetInitialShowStateEvent(cb TOnWindowGetInitialShowStateEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowGetInitialShowStateEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: TCefShowState);
		window := AsCefWindowRef(getVal(0))
		result := (*cefTypes.TCefShowState)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowGetParentWindowEvent(cb TOnWindowGetParentWindowEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowGetParentWindowEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const window_: ICefWindow; var is_menu: boolean; var can_activate_menu: boolean; var aResult: ICefWindow);
		window := AsCefWindowRef(getVal(0))
		isMenu := (*bool)(getPtr(getVal(1)))
		canActivateMenu := (*bool)(getPtr(getVal(2)))
		result := ICefWindow(AsCefWindowRef(*(*uintptr)(getPtr(getVal(3)))))
		cb(window, isMenu, canActivateMenu, &result)
		if result != nil {
			*(*uintptr)(getPtr(getVal(3))) = result.Instance()
		}
	})
}

func makeTOnWindowIsFramelessEvent(cb TOnWindowIsFramelessEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowIsFramelessEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		result := (*bool)(getPtr(getVal(1)))
		cb(window, result)
	})
}

func makeTOnWindowKeyEvent(cb TOnWindowKeyEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowKeyEvent", func(getVal func(i int) uintptr) {
		// 3 : procedure(const window_: ICefWindow; const event: TCefKeyEvent; var aResult: boolean);
		window := AsCefWindowRef(getVal(0))
		event := *(*TCefKeyEvent)(getPtr(getVal(1)))
		result := (*bool)(getPtr(getVal(2)))
		cb(window, event, result)
	})
}

func makeTOnWindowKeyEventEvent(cb TOnWindowKeyEventEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowKeyEventEvent", func(getVal func(i int) uintptr) {
		// 4 : procedure(const Sender: TObject; const window_: ICefWindow; const event: TCefKeyEvent; var aResult : boolean);
		sender := lcl.AsObject(getVal(0))
		window := AsCefWindowRef(getVal(1))
		event := *(*TCefKeyEvent)(getPtr(getVal(2)))
		result := (*bool)(getPtr(getVal(3)))
		cb(sender, window, event, result)
	})
}

func makeTOnWindowWindowActivationChangedEvent(cb TOnWindowWindowActivationChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowWindowActivationChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; active: boolean);
		window := AsCefWindowRef(getVal(0))
		active := api.GoBool(getVal(1))
		cb(window, active)
	})
}

func makeTOnWindowWindowBoundsChangedEvent(cb TOnWindowWindowBoundsChangedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowWindowBoundsChangedEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(const window_: ICefWindow; const new_bounds: TCefRect);
		window := AsCefWindowRef(getVal(0))
		newBounds := *(*TCefRect)(getPtr(getVal(1)))
		cb(window, newBounds)
	})
}

func makeTOnWindowWindowClosingEvent(cb TOnWindowWindowClosingEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowWindowClosingEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const window_: ICefWindow);
		window := AsCefWindowRef(getVal(0))
		cb(window)
	})
}

func makeTOnWindowWindowCreatedEvent(cb TOnWindowWindowCreatedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowWindowCreatedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const window_: ICefWindow);
		window := AsCefWindowRef(getVal(0))
		cb(window)
	})
}

func makeTOnWindowWindowDestroyedEvent(cb TOnWindowWindowDestroyedEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWindowWindowDestroyedEvent", func(getVal func(i int) uintptr) {
		// 1 : procedure(const window_: ICefWindow);
		window := AsCefWindowRef(getVal(0))
		cb(window)
	})
}

func makeTOnWriteEllEvent(cb TOnWriteEllEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWriteEllEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): Int64;
		ret := cb()
		*(*int64)(getPtr(getVal(0))) = ret
	})
}

func makeTOnWriteFlushEvent(cb TOnWriteFlushEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWriteFlushEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): Integer;
		ret := cb()
		*(*int32)(getPtr(getVal(0))) = ret
	})
}

func makeTOnWriteMayBlockEvent(cb TOnWriteMayBlockEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWriteMayBlockEvent", func(getVal func(i int) uintptr) {
		// 0 : function(): Boolean;
		ret := cb()
		*(*bool)(getPtr(getVal(0))) = ret
	})
}

func makeTOnWriteSeekEvent(cb TOnWriteSeekEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWriteSeekEvent", func(getVal func(i int) uintptr) {
		// 2 : function(offset: Int64; whence: Integer): Integer;
		offset := *(*int64)(getPtr(getVal(0)))
		whence := int32(getVal(1))
		ret := cb(offset, whence)
		*(*int32)(getPtr(getVal(2))) = ret
	})
}

func makeTOnWriteWriteEvent(cb TOnWriteWriteEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnWriteWriteEvent", func(getVal func(i int) uintptr) {
		// 3 : function(const ptr: Pointer; size: NativeUInt; n: NativeUInt): NativeUInt;
		ptr := uintptr(getVal(0))
		size := cefTypes.NativeUInt(getVal(1))
		N := cefTypes.NativeUInt(getVal(2))
		ret := cb(ptr, size, N)
		*(*cefTypes.NativeUInt)(getPtr(getVal(3))) = ret
	})
}

func makeTOnZoomPctAvailable(cb TOnZoomPctAvailable) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TOnZoomPctAvailable", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; const aZoomPct : double);
		sender := lcl.AsObject(getVal(0))
		zoomPct := *(*float64)(getPtr(getVal(1)))
		cb(sender, zoomPct)
	})
}

func makeTStartDockEvent(cb TStartDockEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TStartDockEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; var DragObject: TDragDockObject);
		sender := lcl.AsObject(getVal(0))
		var dragObject lcl.IDragDockObject
		dragObject = lcl.AsDragDockObject(*(*uintptr)(getPtr(getVal(1))))
		cb(sender, &dragObject)
		if dragObject != nil {
			*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
		}
	})
}

func makeTStartDragEvent(cb TStartDragEvent) *Callback {
	if cb == nil {
		return nil
	}
	return NewCallback("TStartDragEvent", func(getVal func(i int) uintptr) {
		// 2 : procedure(Sender: TObject; var DragObject: TDragObject);
		sender := lcl.AsObject(getVal(0))
		var dragObject lcl.IDragObject
		dragObject = lcl.AsDragObject(*(*uintptr)(getPtr(getVal(1))))
		cb(sender, &dragObject)
		if dragObject != nil {
			*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
		}
	})
}
