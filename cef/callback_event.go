//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type callback struct {
	name string
	cb   func(getVal func(i int) uintptr)
}

func getPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}

func makeTConstrainedResizeEvent(cb TConstrainedResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TConstrainedResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);
			sender := lcl.AsObject(getVal(0))
			minWidth := (*types.TConstraintSize)(getPtr(getVal(1)))
			minHeight := (*types.TConstraintSize)(getPtr(getVal(2)))
			maxWidth := (*types.TConstraintSize)(getPtr(getVal(3)))
			maxHeight := (*types.TConstraintSize)(getPtr(getVal(4)))
			cb(sender, minWidth, minHeight, maxWidth, maxHeight)
		},
	}
}

func makeTContextPopupEvent(cb TContextPopupEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TContextPopupEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
			sender := lcl.AsObject(getVal(0))
			mousePos := *(*types.TPoint)(getPtr(getVal(1)))
			handled := (*bool)(getPtr(getVal(2)))
			cb(sender, mousePos, handled)
		},
	}
}

func makeTDragDropEvent(cb TDragDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Source: TObject; X,Y: Integer);
			sender := lcl.AsObject(getVal(0))
			source := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, source, X, Y)
		},
	}
}

func makeTDragOverEvent(cb TDragOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender, Source: TObject; X,Y: Integer; State: TDragState; var Accept: Boolean);
			sender := lcl.AsObject(getVal(0))
			source := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			state := types.TDragState(getVal(4))
			accept := (*bool)(getPtr(getVal(5)))
			cb(sender, source, X, Y, state, accept)
		},
	}
}

func makeTEndDragEvent(cb TEndDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TEndDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Target: TObject; X,Y: Integer);
			sender := lcl.AsObject(getVal(0))
			target := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, target, X, Y)
		},
	}
}

func makeTGetSiteInfoEvent(cb TGetSiteInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetSiteInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
			sender := lcl.AsObject(getVal(0))
			dockClient := lcl.AsControl(getVal(1))
			influenceRect := (*types.TRect)(getPtr(getVal(2)))
			mousePos := *(*types.TPoint)(getPtr(getVal(3)))
			canDock := (*bool)(getPtr(getVal(4)))
			cb(sender, dockClient, influenceRect, mousePos, canDock)
		},
	}
}

func makeTMouseEvent(cb TMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y: Integer);
			sender := lcl.AsObject(getVal(0))
			button := types.TMouseButton(getVal(1))
			shift := types.TShiftState(getVal(2))
			X := int32(getVal(3))
			Y := int32(getVal(4))
			cb(sender, button, shift, X, Y)
		},
	}
}

func makeTMouseMoveEvent(cb TMouseMoveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseMoveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Shift: TShiftState; X, Y: Integer);
			sender := lcl.AsObject(getVal(0))
			shift := types.TShiftState(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, shift, X, Y)
		},
	}
}

func makeTMouseWheelEvent(cb TMouseWheelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseWheelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; Shift: TShiftState; WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
			sender := lcl.AsObject(getVal(0))
			shift := types.TShiftState(getVal(1))
			wheelDelta := int32(getVal(2))
			mousePos := *(*types.TPoint)(getPtr(getVal(3)))
			handled := (*bool)(getPtr(getVal(4)))
			cb(sender, shift, wheelDelta, mousePos, handled)
		},
	}
}

func makeTNotifyEvent(cb TNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTOnAcceleratedPaint(cb TOnAcceleratedPaint) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAcceleratedPaint",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; type_: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; const info: PCefAcceleratedPaintInfo);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			type_ := cefTypes.TCefPaintElementType(getVal(2))
			dirtyRectsCount := cefTypes.NativeUInt(getVal(3))
			dirtyRectsPtr := getVal(4)
			dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
			infoPtr := (*tCefAcceleratedPaintInfo)(getPtr(getVal(5)))
			info := infoPtr.ToGo()
			cb(sender, browser, type_, dirtyRectsCount, dirtyRects, info)
		},
	}
}

func makeTOnAcceleratorEvent(cb TOnAcceleratorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAcceleratorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const window_: ICefWindow; command_id: Integer; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, window, commandId, result)
		},
	}
}

func makeTOnAcceptsFirstMouseEvent(cb TOnAcceptsFirstMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAcceptsFirstMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : TCefState);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*cefTypes.TCefState)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnAccessibilityAccessibilityLocationChangeEvent(cb TOnAccessibilityAccessibilityLocationChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAccessibilityAccessibilityLocationChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const value: ICefValue);
			value := AsCefValueRef(getVal(0))
			cb(value)
		},
	}
}

func makeTOnAccessibilityAccessibilityTreeChangeEvent(cb TOnAccessibilityAccessibilityTreeChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAccessibilityAccessibilityTreeChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const value: ICefValue);
			value := AsCefValueRef(getVal(0))
			cb(value)
		},
	}
}

func makeTOnAccessibilityEvent(cb TOnAccessibilityEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAccessibilityEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const value: ICefValue);
			sender := lcl.AsObject(getVal(0))
			value := AsCefValueRef(getVal(1))
			cb(sender, value)
		},
	}
}

func makeTOnAddressChange(cb TOnAddressChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAddressChange",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const url: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			url := api.GoStr(getVal(3))
			cb(sender, browser, frame, url)
		},
	}
}

func makeTOnAfterCreated(cb TOnAfterCreated) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAfterCreated",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnAfterUserActionEvent(cb TOnAfterUserActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAfterUserActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const textfield: ICefTextfield);
			sender := lcl.AsObject(getVal(0))
			textfield := AsCefTextfieldRef(getVal(1))
			cb(sender, textfield)
		},
	}
}

func makeTOnAllowEvent(cb TOnAllowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAllowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var allow : boolean);
			sender := lcl.AsObject(getVal(0))
			allow := (*bool)(getPtr(getVal(1)))
			cb(sender, allow)
		},
	}
}

func makeTOnAllowMoveForPictureInPictureEvent(cb TOnAllowMoveForPictureInPictureEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAllowMoveForPictureInPictureEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, browserView, result)
		},
	}
}

func makeTOnAllowPictureInPictureWithoutUserActivationEvent(cb TOnAllowPictureInPictureWithoutUserActivationEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAllowPictureInPictureWithoutUserActivationEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, browserView, result)
		},
	}
}

func makeTOnAlreadyRunningAppRelaunchEvent(cb TOnAlreadyRunningAppRelaunchEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAlreadyRunningAppRelaunchEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const commandLine: ICefCommandLine; const current_directory: ustring; var aResult: boolean);
			commandLine := AsCefCommandLineRef(getVal(0))
			currentDirectory := api.GoStr(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(commandLine, currentDirectory, result)
		},
	}
}

func makeTOnAppBeforeCommandLineProcessingEvent(cb TOnAppBeforeCommandLineProcessingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAppBeforeCommandLineProcessingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const processType: ustring; const commandLine: ICefCommandLine);
			processType := api.GoStr(getVal(0))
			commandLine := AsCefCommandLineRef(getVal(1))
			cb(processType, commandLine)
		},
	}
}

func makeTOnAppGetBrowserProcessHandlerEvent(cb TOnAppGetBrowserProcessHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAppGetBrowserProcessHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefBrowserProcessHandler);
			var handler IEngBrowserProcessHandler
			handler = AsEngBrowserProcessHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnAppGetRenderProcessHandlerEvent(cb TOnAppGetRenderProcessHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAppGetRenderProcessHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefRenderProcessHandler);
			var handler IEngRenderProcessHandler
			handler = AsEngRenderProcessHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnAppGetResourceBundleHandlerEvent(cb TOnAppGetResourceBundleHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAppGetResourceBundleHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefResourceBundleHandler);
			var handler IEngResourceBundleHandler
			handler = AsEngResourceBundleHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnAppRegisterCustomSchemesEvent(cb TOnAppRegisterCustomSchemesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAppRegisterCustomSchemesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const registrar: TCefSchemeRegistrarRef);
			registrar := AsCefSchemeRegistrarRef(getVal(0))
			cb(registrar)
		},
	}
}

func makeTOnAudioAudioStreamErrorEvent(cb TOnAudioAudioStreamErrorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioAudioStreamErrorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const message_: ustring);
			browser := AsCefBrowserRef(getVal(0))
			message := api.GoStr(getVal(1))
			cb(browser, message)
		},
	}
}

func makeTOnAudioAudioStreamPacketEvent(cb TOnAudioAudioStreamPacketEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioAudioStreamPacketEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const data: PPSingle; frames: integer; pts: int64);
			browser := AsCefBrowserRef(getVal(0))
			data := PPSingle(getVal(1))
			frames := int32(getVal(2))
			pts := *(*int64)(getPtr(getVal(3)))
			cb(browser, data, frames, pts)
		},
	}
}

func makeTOnAudioAudioStreamStartedEvent(cb TOnAudioAudioStreamStartedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioAudioStreamStartedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const params: TCefAudioParameters; channels: integer);
			browser := AsCefBrowserRef(getVal(0))
			params := *(*TCefAudioParameters)(getPtr(getVal(1)))
			channels := int32(getVal(2))
			cb(browser, params, channels)
		},
	}
}

func makeTOnAudioAudioStreamStoppedEvent(cb TOnAudioAudioStreamStoppedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioAudioStreamStoppedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnAudioGetAudioParametersEvent(cb TOnAudioGetAudioParametersEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioGetAudioParametersEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; var params: TCefAudioParameters; var aResult: boolean);
			browser := AsCefBrowserRef(getVal(0))
			params := (*TCefAudioParameters)(getPtr(getVal(1)))
			result := (*bool)(getPtr(getVal(2)))
			cb(browser, params, result)
			*(*TCefAudioParameters)(getPtr(getVal(1))) = *params
		},
	}
}

func makeTOnAudioStreamErrorEvent(cb TOnAudioStreamErrorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioStreamErrorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			message := api.GoStr(getVal(2))
			cb(sender, browser, message)
		},
	}
}

func makeTOnAudioStreamPacketEvent(cb TOnAudioStreamPacketEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioStreamPacketEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const data : PPSingle; frames: integer; pts: int64);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			data := PPSingle(getVal(2))
			frames := int32(getVal(3))
			pts := *(*int64)(getPtr(getVal(4)))
			cb(sender, browser, data, frames, pts)
		},
	}
}

func makeTOnAudioStreamStartedEvent(cb TOnAudioStreamStartedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioStreamStartedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const params: TCefAudioParameters; channels: integer);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			params := *(*TCefAudioParameters)(getPtr(getVal(2)))
			channels := int32(getVal(3))
			cb(sender, browser, params, channels)
		},
	}
}

func makeTOnAudioStreamStoppedEvent(cb TOnAudioStreamStoppedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAudioStreamStoppedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnAutoResize(cb TOnAutoResize) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAutoResize",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const new_size: PCefSize; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			newSize := *(*TCefSize)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, newSize, result)
		},
	}
}

func makeTOnBeforeBrowse(cb TOnBeforeBrowse) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeBrowse",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; user_gesture, isRedirect: Boolean; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			userGesture := api.GoBool(getVal(4))
			isRedirect := api.GoBool(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, frame, request, userGesture, isRedirect, result)
		},
	}
}

func makeTOnBeforeChildProcessLaunchEvent(cb TOnBeforeChildProcessLaunchEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeChildProcessLaunchEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const commandLine: ICefCommandLine);
			commandLine := AsCefCommandLineRef(getVal(0))
			cb(commandLine)
		},
	}
}

func makeTOnBeforeClose(cb TOnBeforeClose) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeClose",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnBeforeContextMenu(cb TOnBeforeContextMenu) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeContextMenu",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			params := AsCefContextMenuParamsRef(getVal(3))
			model := AsCefMenuModelRef(getVal(4))
			cb(sender, browser, frame, params, model)
		},
	}
}

func makeTOnBeforeDevToolsPopup(cb TOnBeforeDevToolsPopup) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeDevToolsPopup",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var use_default_window: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(2)))
			windowInfo := windowInfoPtr.ToGo()
			var client IEngClient
			client = AsEngClient(*(*uintptr)(getPtr(getVal(3))))
			settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(4)))
			settings := settingsPtr.ToGo()
			extraInfo := ICefDictionaryValue(AsCefDictionaryValueRef(*(*uintptr)(getPtr(getVal(5)))))
			useDefaultWindow := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, &windowInfo, &client, &settings, &extraInfo, useDefaultWindow)
			if r := windowInfo.ToPas(); r != nil {
				*(*tCefWindowInfo)(getPtr(getVal(2))) = *r
			}
			if client != nil {
				*(*uintptr)(getPtr(getVal(3))) = client.Instance()
			}
			if r := settings.ToPas(); r != nil {
				*(*tCefBrowserSettings)(getPtr(getVal(4))) = *r
			}
			if extraInfo != nil {
				*(*uintptr)(getPtr(getVal(5))) = extraInfo.Instance()
			}
		},
	}
}

func makeTOnBeforeDownload(cb TOnBeforeDownload) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeDownload",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const suggestedName: ustring; const callback: ICefBeforeDownloadCallback; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			downloadItem := AsCefDownloadItemRef(getVal(2))
			suggestedName := api.GoStr(getVal(3))
			callback := AsCefBeforeDownloadCallbackRef(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, downloadItem, suggestedName, callback, result)
		},
	}
}

func makeTOnBeforePopup(cb TOnBeforePopup) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforePopup",
		cb: func(getVal func(i int) uintptr) {
			// 15 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; popup_id: Integer; const targetUrl, targetFrameName: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; const popupFeatures: TCefPopupFeatures; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var noJavascriptAccess: Boolean; var Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			popupId := int32(getVal(3))
			targetUrl := api.GoStr(getVal(4))
			targetFrameName := api.GoStr(getVal(5))
			targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(6))
			userGesture := api.GoBool(getVal(7))
			popupFeatures := *(*TCefPopupFeatures)(getPtr(getVal(8)))
			windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(9)))
			windowInfo := windowInfoPtr.ToGo()
			var client IEngClient
			client = AsEngClient(*(*uintptr)(getPtr(getVal(10))))
			settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(11)))
			settings := settingsPtr.ToGo()
			extraInfo := ICefDictionaryValue(AsCefDictionaryValueRef(*(*uintptr)(getPtr(getVal(12)))))
			noJavascriptAccess := (*bool)(getPtr(getVal(13)))
			result := (*bool)(getPtr(getVal(14)))
			cb(sender, browser, frame, popupId, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, &windowInfo, &client, &settings, &extraInfo, noJavascriptAccess, result)
			if r := windowInfo.ToPas(); r != nil {
				*(*tCefWindowInfo)(getPtr(getVal(9))) = *r
			}
			if client != nil {
				*(*uintptr)(getPtr(getVal(10))) = client.Instance()
			}
			if r := settings.ToPas(); r != nil {
				*(*tCefBrowserSettings)(getPtr(getVal(11))) = *r
			}
			if extraInfo != nil {
				*(*uintptr)(getPtr(getVal(12))) = extraInfo.Instance()
			}
		},
	}
}

func makeTOnBeforePopupAborted(cb TOnBeforePopupAborted) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforePopupAborted",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; popup_id: Integer);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			popupId := int32(getVal(2))
			cb(sender, browser, popupId)
		},
	}
}

func makeTOnBeforeResourceLoad(cb TOnBeforeResourceLoad) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeResourceLoad",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const callback: ICefCallback; out Result: TCefReturnValue);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			callback := AsCefCallbackRef(getVal(4))
			result := (*cefTypes.TCefReturnValue)(getPtr(getVal(5)))
			cb(sender, browser, frame, request, callback, result)
		},
	}
}

func makeTOnBeforeUnloadDialog(cb TOnBeforeUnloadDialog) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBeforeUnloadDialog",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const messageText: ustring; isReload: Boolean; const callback: ICefJsDialogCallback; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			messageText := api.GoStr(getVal(2))
			isReload := api.GoBool(getVal(3))
			callback := AsCefJsDialogCallbackRef(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, messageText, isReload, callback, result)
		},
	}
}

func makeTOnBlurEvent(cb TOnBlurEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBlurEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const view: ICefView);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			cb(sender, view)
		},
	}
}

func makeTOnBrowserCreatedEvent(cb TOnBrowserCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const extra_info: ICefDictionaryValue);
			browser := AsCefBrowserRef(getVal(0))
			extraInfo := AsCefDictionaryValueRef(getVal(1))
			cb(browser, extraInfo)
		},
	}
}

func makeTOnBrowserDestroyedEvent(cb TOnBrowserDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnBrowserProcessAlreadyRunningAppRelaunchEvent(cb TOnBrowserProcessAlreadyRunningAppRelaunchEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessAlreadyRunningAppRelaunchEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const commandLine: ICefCommandLine; const current_directory: ustring; var aResult: boolean);
			commandLine := AsCefCommandLineRef(getVal(0))
			currentDirectory := api.GoStr(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(commandLine, currentDirectory, result)
		},
	}
}

func makeTOnBrowserProcessBeforeChildProcessLaunchEvent(cb TOnBrowserProcessBeforeChildProcessLaunchEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessBeforeChildProcessLaunchEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const commandLine: ICefCommandLine);
			commandLine := AsCefCommandLineRef(getVal(0))
			cb(commandLine)
		},
	}
}

func makeTOnBrowserProcessContextInitializedEvent(cb TOnBrowserProcessContextInitializedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessContextInitializedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnBrowserProcessGetDefaultClientEvent(cb TOnBrowserProcessGetDefaultClientEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessGetDefaultClientEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aClient: ICefClient);
			var client IEngClient
			client = AsEngClient(*(*uintptr)(getPtr(getVal(0))))
			cb(&client)
			if client != nil {
				*(*uintptr)(getPtr(getVal(0))) = client.Instance()
			}
		},
	}
}

func makeTOnBrowserProcessGetDefaultRequestContextHandlerEvent(cb TOnBrowserProcessGetDefaultRequestContextHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessGetDefaultRequestContextHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aRequestContextHandler: ICefRequestContextHandler);
			var requestContextHandler IEngRequestContextHandler
			requestContextHandler = AsEngRequestContextHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&requestContextHandler)
			if requestContextHandler != nil {
				*(*uintptr)(getPtr(getVal(0))) = requestContextHandler.Instance()
			}
		},
	}
}

func makeTOnBrowserProcessRegisterCustomPreferencesEvent(cb TOnBrowserProcessRegisterCustomPreferencesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessRegisterCustomPreferencesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(type_: TCefPreferencesType; registrar: PCefPreferenceRegistrar);
			type_ := cefTypes.TCefPreferencesType(getVal(0))
			registrar := *(*TCefPreferenceRegistrar)(getPtr(getVal(1)))
			cb(type_, registrar)
		},
	}
}

func makeTOnBrowserProcessScheduleMessagePumpWorkEvent(cb TOnBrowserProcessScheduleMessagePumpWorkEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessScheduleMessagePumpWorkEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const delayMs: Int64);
			delayMs := *(*int64)(getPtr(getVal(0)))
			cb(delayMs)
		},
	}
}

func makeTOnBrowserViewAllowMoveForPictureInPictureEvent(cb TOnBrowserViewAllowMoveForPictureInPictureEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewAllowMoveForPictureInPictureEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; var aResult: boolean);
			browserView := AsCefBrowserViewRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(browserView, result)
		},
	}
}

func makeTOnBrowserViewAllowPictureInPictureWithoutUserActivationEvent(cb TOnBrowserViewAllowPictureInPictureWithoutUserActivationEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewAllowPictureInPictureWithoutUserActivationEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; var aResult: boolean);
			browserView := AsCefBrowserViewRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(browserView, result)
		},
	}
}

func makeTOnBrowserViewBrowserCreatedEvent(cb TOnBrowserViewBrowserCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewBrowserCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; const browser: ICefBrowser);
			browserView := AsCefBrowserViewRef(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(browserView, browser)
		},
	}
}

func makeTOnBrowserViewBrowserDestroyedEvent(cb TOnBrowserViewBrowserDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewBrowserDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; const browser: ICefBrowser);
			browserView := AsCefBrowserViewRef(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(browserView, browser)
		},
	}
}

func makeTOnBrowserViewGestureCommandEvent(cb TOnBrowserViewGestureCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewGestureCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser_view: ICefBrowserView; gesture_command: TCefGestureCommand; var aResult: boolean);
			browserView := AsCefBrowserViewRef(getVal(0))
			gestureCommand := cefTypes.TCefGestureCommand(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(browserView, gestureCommand, result)
		},
	}
}

func makeTOnBrowserViewGetBrowserRuntimeStyleEvent(cb TOnBrowserViewGetBrowserRuntimeStyleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewGetBrowserRuntimeStyleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aResult: TCefRuntimeStyle);
			result := (*cefTypes.TCefRuntimeStyle)(getPtr(getVal(0)))
			cb(result)
		},
	}
}

func makeTOnBrowserViewGetChromeToolbarTypeEvent(cb TOnBrowserViewGetChromeToolbarTypeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewGetChromeToolbarTypeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; var aResult: TCefChromeToolbarType);
			browserView := AsCefBrowserViewRef(getVal(0))
			result := (*TCefChromeToolbarType)(getPtr(getVal(1)))
			cb(browserView, result)
		},
	}
}

func makeTOnBrowserViewGetDelegateForPopupBrowserViewEvent(cb TOnBrowserViewGetDelegateForPopupBrowserViewEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewGetDelegateForPopupBrowserViewEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnBrowserViewPopupBrowserViewCreatedEvent(cb TOnBrowserViewPopupBrowserViewCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewPopupBrowserViewCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser_view: ICefBrowserView; const popup_browser_view: ICefBrowserView; is_devtools: boolean; var aResult: boolean);
			browserView := AsCefBrowserViewRef(getVal(0))
			popupBrowserView := AsCefBrowserViewRef(getVal(1))
			isDevtools := api.GoBool(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(browserView, popupBrowserView, isDevtools, result)
		},
	}
}

func makeTOnBrowserViewUseFramelessWindowForPictureInPictureEvent(cb TOnBrowserViewUseFramelessWindowForPictureInPictureEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserViewUseFramelessWindowForPictureInPictureEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser_view: ICefBrowserView; var aResult: boolean);
			browserView := AsCefBrowserViewRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(browserView, result)
		},
	}
}

func makeTOnButtonButtonPressedEvent(cb TOnButtonButtonPressedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnButtonButtonPressedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const button: ICefButton);
			button := AsCefButtonRef(getVal(0))
			cb(button)
		},
	}
}

func makeTOnButtonButtonStateChangedEvent(cb TOnButtonButtonStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnButtonButtonStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const button: ICefButton);
			button := AsCefButtonRef(getVal(0))
			cb(button)
		},
	}
}

func makeTOnButtonPressedEvent(cb TOnButtonPressedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnButtonPressedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const button: ICefButton);
			sender := lcl.AsObject(getVal(0))
			button := AsCefButtonRef(getVal(1))
			cb(sender, button)
		},
	}
}

func makeTOnButtonStateChangedEvent(cb TOnButtonStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnButtonStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const button: ICefButton);
			sender := lcl.AsObject(getVal(0))
			button := AsCefButtonRef(getVal(1))
			cb(sender, button)
		},
	}
}

func makeTOnCanCloseEvent(cb TOnCanCloseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanCloseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnCanDownloadEvent(cb TOnCanDownloadEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanDownloadEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const url, request_method: ustring; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			url := api.GoStr(getVal(2))
			requestMethod := api.GoStr(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, url, requestMethod, result)
		},
	}
}

func makeTOnCanMaximizeEvent(cb TOnCanMaximizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanMaximizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnCanMinimizeEvent(cb TOnCanMinimizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanMinimizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnCanResizeEvent(cb TOnCanResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnCanSaveCookie(cb TOnCanSaveCookie) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanSaveCookie",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnCanSendCookie(cb TOnCanSendCookie) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCanSendCookie",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const cookie: PCefCookie; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			cookiePtr := (*tCefCookie)(getPtr(getVal(4)))
			cookie := cookiePtr.ToGo()
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, frame, request, cookie, result)
		},
	}
}

func makeTOnCertificateError(cb TOnCertificateError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCertificateError",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; certError: TCefErrorcode; const requestUrl: ustring; const sslInfo: ICefSslInfo; const callback: ICefCallback; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			certError := cefTypes.TCefErrorCode(getVal(2))
			requestUrl := api.GoStr(getVal(3))
			sslInfo := AsCefSslInfoRef(getVal(4))
			callback := AsCefCallbackRef(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, certError, requestUrl, sslInfo, callback, result)
		},
	}
}

func makeTOnChildViewChangedEvent(cb TOnChildViewChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnChildViewChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const view: ICefView; added: boolean; const child: ICefView);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			added := api.GoBool(getVal(2))
			child := AsCefViewRef(getVal(3))
			cb(sender, view, added, child)
		},
	}
}

func makeTOnChromeCommandEvent(cb TOnChromeCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnChromeCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; command_id: integer; disposition: TCefWindowOpenDisposition; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			commandId := int32(getVal(2))
			disposition := cefTypes.TCefWindowOpenDisposition(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, commandId, disposition, result)
		},
	}
}

func makeTOnClientConnected(cb TOnClientConnected) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientConnected",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			cb(sender, server, connectionId)
		},
	}
}

func makeTOnClientDisconnected(cb TOnClientDisconnected) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientDisconnected",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			cb(sender, server, connectionId)
		},
	}
}

func makeTOnClientGetAudioHandlerEvent(cb TOnClientGetAudioHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetAudioHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefAudioHandler);
			var handler IEngAudioHandler
			handler = AsEngAudioHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetCommandHandlerEvent(cb TOnClientGetCommandHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetCommandHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefCommandHandler);
			var handler IEngCommandHandler
			handler = AsEngCommandHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetContextMenuHandlerEvent(cb TOnClientGetContextMenuHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetContextMenuHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefContextMenuHandler);
			var handler IEngContextMenuHandler
			handler = AsEngContextMenuHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetDialogHandlerEvent(cb TOnClientGetDialogHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetDialogHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefDialogHandler);
			var handler IEngDialogHandler
			handler = AsEngDialogHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetDisplayHandlerEvent(cb TOnClientGetDisplayHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetDisplayHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefDisplayHandler);
			var handler IEngDisplayHandler
			handler = AsEngDisplayHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetDownloadHandlerEvent(cb TOnClientGetDownloadHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetDownloadHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefDownloadHandler);
			var handler IEngDownloadHandler
			handler = AsEngDownloadHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetDragHandlerEvent(cb TOnClientGetDragHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetDragHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefDragHandler);
			var handler IEngDragHandler
			handler = AsEngDragHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetFindHandlerEvent(cb TOnClientGetFindHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetFindHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefFindHandler);
			var handler IEngFindHandler
			handler = AsEngFindHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetFocusHandlerEvent(cb TOnClientGetFocusHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetFocusHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefFocusHandler);
			var handler IEngFocusHandler
			handler = AsEngFocusHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetFrameHandlerEvent(cb TOnClientGetFrameHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetFrameHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefFrameHandler);
			var handler IEngFrameHandler
			handler = AsEngFrameHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetJsdialogHandlerEvent(cb TOnClientGetJsdialogHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetJsdialogHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefJsdialogHandler);
			var handler IEngJsDialogHandler
			handler = AsEngJsDialogHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetKeyboardHandlerEvent(cb TOnClientGetKeyboardHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetKeyboardHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefKeyboardHandler);
			var handler IEngKeyboardHandler
			handler = AsEngKeyboardHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetLifeSpanHandlerEvent(cb TOnClientGetLifeSpanHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetLifeSpanHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefLifeSpanHandler);
			var handler IEngLifeSpanHandler
			handler = AsEngLifeSpanHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetLoadHandlerEvent(cb TOnClientGetLoadHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetLoadHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefLoadHandler);
			var handler IEngLoadHandler
			handler = AsEngLoadHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetPermissionHandlerEvent(cb TOnClientGetPermissionHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetPermissionHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefPermissionHandler);
			var handler IEngPermissionHandler
			handler = AsEngPermissionHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetPrintHandlerEvent(cb TOnClientGetPrintHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetPrintHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefPrintHandler);
			var handler IEngPrintHandler
			handler = AsEngPrintHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetRenderHandlerEvent(cb TOnClientGetRenderHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetRenderHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefRenderHandler);
			var handler IEngRenderHandler
			handler = AsEngRenderHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientGetRequestHandlerEvent(cb TOnClientGetRequestHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientGetRequestHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aHandler: ICefRequestHandler);
			var handler IEngRequestHandler
			handler = AsEngRequestHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&handler)
			if handler != nil {
				*(*uintptr)(getPtr(getVal(0))) = handler.Instance()
			}
		},
	}
}

func makeTOnClientProcessMessageReceivedEvent(cb TOnClientProcessMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientProcessMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message_: ICefProcessMessage): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			sourceProcess := cefTypes.TCefProcessId(getVal(2))
			message := AsCefProcessMessageRef(getVal(3))
			ret := cb(browser, frame, sourceProcess, message)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnClose(cb TOnClose) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClose",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; var aAction : TCefCloseBrowserAction);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			action := (*cefTypes.TCefCloseBrowserAction)(getPtr(getVal(2)))
			cb(sender, browser, action)
		},
	}
}

func makeTOnCommandChromeCommandEvent(cb TOnCommandChromeCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCommandChromeCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const browser: ICefBrowser; command_id: integer; disposition: TCefWindowOpenDisposition): boolean;
			browser := AsCefBrowserRef(getVal(0))
			commandId := int32(getVal(1))
			disposition := cefTypes.TCefWindowOpenDisposition(getVal(2))
			ret := cb(browser, commandId, disposition)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnCommandIsChromeAppMenuItemEnabledEvent(cb TOnCommandIsChromeAppMenuItemEnabledEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCommandIsChromeAppMenuItemEnabledEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; command_id: integer): boolean;
			browser := AsCefBrowserRef(getVal(0))
			commandId := int32(getVal(1))
			ret := cb(browser, commandId)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnCommandIsChromeAppMenuItemVisibleEvent(cb TOnCommandIsChromeAppMenuItemVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCommandIsChromeAppMenuItemVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; command_id: integer): boolean;
			browser := AsCefBrowserRef(getVal(0))
			commandId := int32(getVal(1))
			ret := cb(browser, commandId)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnCommandIsChromePageActionIconVisibleEvent(cb TOnCommandIsChromePageActionIconVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCommandIsChromePageActionIconVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : function(icon_type: TCefChromePageActionIconType): boolean;
			iconType := cefTypes.TCefChromePageActionIconType(getVal(0))
			ret := cb(iconType)
			*(*bool)(getPtr(getVal(1))) = ret
		},
	}
}

func makeTOnCommandIsChromeToolbarButtonVisibleEvent(cb TOnCommandIsChromeToolbarButtonVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCommandIsChromeToolbarButtonVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : function(button_type: TCefChromeToolbarButtonType): boolean;
			buttonType := cefTypes.TCefChromeToolbarButtonType(getVal(0))
			ret := cb(buttonType)
			*(*bool)(getPtr(getVal(1))) = ret
		},
	}
}

func makeTOnCompMsgEvent(cb TOnCompMsgEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCompMsgEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var aMessage: TMessage; var aHandled: Boolean);
			sender := lcl.AsObject(getVal(0))
			message := (*types.TMessage)(getPtr(getVal(1)))
			handled := (*bool)(getPtr(getVal(2)))
			cb(sender, message, handled)
		},
	}
}

func makeTOnCompletionCallbackCompleteEvent(cb TOnCompletionCallbackCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCompletionCallbackCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnComponentUpdateCompletedEvent(cb TOnComponentUpdateCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnComponentUpdateCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const component_id: ustring; error: TCefComponentUpdateError);
			sender := lcl.AsObject(getVal(0))
			componentId := api.GoStr(getVal(1))
			error_ := cefTypes.TCefComponentUpdateError(getVal(2))
			cb(sender, componentId, error_)
		},
	}
}

func makeTOnConsoleMessage(cb TOnConsoleMessage) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnConsoleMessage",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; level: TCefLogSeverity; const message_, source: ustring; line: Integer; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			level := cefTypes.TCefLogSeverity(getVal(2))
			message := api.GoStr(getVal(3))
			source := api.GoStr(getVal(4))
			line := int32(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, level, message, source, line, result)
		},
	}
}

func makeTOnContentsBoundsChange(cb TOnContentsBoundsChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContentsBoundsChange",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const new_bounds: PCefRect; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			newBounds := *(*TCefRect)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, newBounds, result)
		},
	}
}

func makeTOnContextCreatedEvent(cb TOnContextCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			cb(browser, frame, context)
		},
	}
}

func makeTOnContextInitializedEvent(cb TOnContextInitializedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextInitializedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnContextMenuBeforeContextMenuEvent(cb TOnContextMenuBeforeContextMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuBeforeContextMenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			params := AsCefContextMenuParamsRef(getVal(2))
			model := AsCefMenuModelRef(getVal(3))
			cb(browser, frame, params, model)
		},
	}
}

func makeTOnContextMenuCommand(cb TOnContextMenuCommand) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuCommand",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; commandId: Integer; eventFlags: TCefEventFlags; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			params := AsCefContextMenuParamsRef(getVal(3))
			commandId := int32(getVal(4))
			eventFlags := cefTypes.TCefEventFlags(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, frame, params, commandId, eventFlags, result)
		},
	}
}

func makeTOnContextMenuContextMenuCommandEvent(cb TOnContextMenuContextMenuCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuContextMenuCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; commandId: Integer; eventFlags: TCefEventFlags): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			params := AsCefContextMenuParamsRef(getVal(2))
			commandId := int32(getVal(3))
			eventFlags := cefTypes.TCefEventFlags(getVal(4))
			ret := cb(browser, frame, params, commandId, eventFlags)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnContextMenuContextMenuDismissedEvent(cb TOnContextMenuContextMenuDismissedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuContextMenuDismissedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			cb(browser, frame)
		},
	}
}

func makeTOnContextMenuDismissed(cb TOnContextMenuDismissed) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuDismissed",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			cb(sender, browser, frame)
		},
	}
}

func makeTOnContextMenuQuickMenuCommandEvent(cb TOnContextMenuQuickMenuCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuQuickMenuCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; command_id: integer; event_flags: TCefEventFlags): boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			commandId := int32(getVal(2))
			eventFlags := cefTypes.TCefEventFlags(getVal(3))
			ret := cb(browser, frame, commandId, eventFlags)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnContextMenuQuickMenuDismissedEvent(cb TOnContextMenuQuickMenuDismissedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuQuickMenuDismissedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			cb(browser, frame)
		},
	}
}

func makeTOnContextMenuRunContextMenuEvent(cb TOnContextMenuRunContextMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuRunContextMenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel; const callback: ICefRunContextMenuCallback): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			params := AsCefContextMenuParamsRef(getVal(2))
			model := AsCefMenuModelRef(getVal(3))
			callback := AsCefRunContextMenuCallbackRef(getVal(4))
			ret := cb(browser, frame, params, model, callback)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnContextMenuRunQuickMenuEvent(cb TOnContextMenuRunQuickMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuRunQuickMenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : function(const browser: ICefBrowser; const frame: ICefFrame; location: PCefPoint; size: PCefSize; edit_state_flags: TCefQuickMenuEditStateFlags; const callback: ICefRunQuickMenuCallback): boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			location := *(*TCefPoint)(getPtr(getVal(2)))
			size := *(*TCefSize)(getPtr(getVal(3)))
			editStateFlags := cefTypes.TCefQuickMenuEditStateFlags(getVal(4))
			callback := AsCefRunQuickMenuCallbackRef(getVal(5))
			ret := cb(browser, frame, location, size, editStateFlags, callback)
			*(*bool)(getPtr(getVal(6))) = ret
		},
	}
}

func makeTOnContextReleasedEvent(cb TOnContextReleasedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextReleasedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			cb(browser, frame, context)
		},
	}
}

func makeTOnCookieAccessFilterCanSaveCookieEvent(cb TOnCookieAccessFilterCanSaveCookieEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookieAccessFilterCanSaveCookieEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; const cookie: PCefCookie): boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			response := AsCefResponseRef(getVal(3))
			cookiePtr := (*tCefCookie)(getPtr(getVal(4)))
			cookie := cookiePtr.ToGo()
			ret := cb(browser, frame, request, response, cookie)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnCookieAccessFilterCanSendCookieEvent(cb TOnCookieAccessFilterCanSendCookieEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookieAccessFilterCanSendCookieEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const cookie: PCefCookie): boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			cookiePtr := (*tCefCookie)(getPtr(getVal(3)))
			cookie := cookiePtr.ToGo()
			ret := cb(browser, frame, request, cookie)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnCookieSet(cb TOnCookieSet) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookieSet",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aSuccess : boolean; aID : integer);
			sender := lcl.AsObject(getVal(0))
			success := api.GoBool(getVal(1))
			iD := int32(getVal(2))
			cb(sender, success, iD)
		},
	}
}

func makeTOnCookieVisitorDestroyed(cb TOnCookieVisitorDestroyed) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookieVisitorDestroyed",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aID : integer);
			sender := lcl.AsObject(getVal(0))
			iD := int32(getVal(1))
			cb(sender, iD)
		},
	}
}

func makeTOnCookieVisitorVisitEvent(cb TOnCookieVisitorVisitEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookieVisitorVisitEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnCookiesDeletedEvent(cb TOnCookiesDeletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookiesDeletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; numDeleted : integer);
			sender := lcl.AsObject(getVal(0))
			numDeleted := int32(getVal(1))
			cb(sender, numDeleted)
		},
	}
}

func makeTOnCookiesVisited(cb TOnCookiesVisited) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCookiesVisited",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnCursorChange(cb TOnCursorChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCursorChange",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; cursor_: TCefCursorHandle; cursorType: TCefCursorType; const customCursorInfo: PCefCursorInfo; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cursor := cefTypes.TCefCursorHandle(getVal(2))
			cursorType := cefTypes.TCefCursorType(getVal(3))
			customCursorInfo := *(*TCefCursorInfo)(getPtr(getVal(4)))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, cursor, cursorType, customCursorInfo, result)
		},
	}
}

func makeTOnDeleteCookiesCallbackCompleteEvent(cb TOnDeleteCookiesCallbackCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDeleteCookiesCallbackCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(numDeleted: Integer);
			numDeleted := int32(getVal(0))
			cb(numDeleted)
		},
	}
}

func makeTOnDevToolsAgentAttachedEvent(cb TOnDevToolsAgentAttachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsAgentAttachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnDevToolsAgentDetachedEvent(cb TOnDevToolsAgentDetachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsAgentDetachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnDevToolsEventEvent(cb TOnDevToolsEventEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsEventEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const method: ustring; const params: ICefValue);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			method := api.GoStr(getVal(2))
			params := AsCefValueRef(getVal(3))
			cb(sender, browser, method, params)
		},
	}
}

func makeTOnDevToolsEventRawEvent(cb TOnDevToolsEventRawEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsEventRawEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const method: ustring; const params: Pointer; params_size: NativeUInt);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			method := api.GoStr(getVal(2))
			params := uintptr(getVal(3))
			paramsSize := cefTypes.NativeUInt(getVal(4))
			cb(sender, browser, method, params, paramsSize)
		},
	}
}

func makeTOnDevToolsMessageEvent(cb TOnDevToolsMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: ICefValue; var aHandled: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			message := AsCefValueRef(getVal(2))
			handled := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, message, handled)
		},
	}
}

func makeTOnDevToolsMessageObserverDevToolsAgentAttachedEvent(cb TOnDevToolsMessageObserverDevToolsAgentAttachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageObserverDevToolsAgentAttachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnDevToolsMessageObserverDevToolsAgentDetachedEvent(cb TOnDevToolsMessageObserverDevToolsAgentDetachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageObserverDevToolsAgentDetachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnDevToolsMessageObserverDevToolsEvent(cb TOnDevToolsMessageObserverDevToolsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageObserverDevToolsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const method: ustring; const params: Pointer; params_size: NativeUInt);
			browser := AsCefBrowserRef(getVal(0))
			method := api.GoStr(getVal(1))
			params := uintptr(getVal(2))
			paramsSize := cefTypes.NativeUInt(getVal(3))
			cb(browser, method, params, paramsSize)
		},
	}
}

func makeTOnDevToolsMessageObserverDevToolsMessageEvent(cb TOnDevToolsMessageObserverDevToolsMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageObserverDevToolsMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const message_: Pointer; message_size: NativeUInt; var aHandled: boolean);
			browser := AsCefBrowserRef(getVal(0))
			message := uintptr(getVal(1))
			messageSize := cefTypes.NativeUInt(getVal(2))
			handled := (*bool)(getPtr(getVal(3)))
			cb(browser, message, messageSize, handled)
		},
	}
}

func makeTOnDevToolsMessageObserverDevToolsMethodResultEvent(cb TOnDevToolsMessageObserverDevToolsMethodResultEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMessageObserverDevToolsMethodResultEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; message_id: integer; success: boolean; const result: Pointer; result_size: NativeUInt);
			browser := AsCefBrowserRef(getVal(0))
			messageId := int32(getVal(1))
			success := api.GoBool(getVal(2))
			result := uintptr(getVal(3))
			resultSize := cefTypes.NativeUInt(getVal(4))
			cb(browser, messageId, success, result, resultSize)
		},
	}
}

func makeTOnDevToolsMethodRawResultEvent(cb TOnDevToolsMethodRawResultEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMethodRawResultEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; message_id: integer; success: boolean; const result: Pointer; result_size: NativeUInt);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			messageId := int32(getVal(2))
			success := api.GoBool(getVal(3))
			result := uintptr(getVal(4))
			resultSize := cefTypes.NativeUInt(getVal(5))
			cb(sender, browser, messageId, success, result, resultSize)
		},
	}
}

func makeTOnDevToolsMethodResultEvent(cb TOnDevToolsMethodResultEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsMethodResultEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; message_id: integer; success: boolean; const result: ICefValue);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			messageId := int32(getVal(2))
			success := api.GoBool(getVal(3))
			result := AsCefValueRef(getVal(4))
			cb(sender, browser, messageId, success, result)
		},
	}
}

func makeTOnDevToolsRawMessageEvent(cb TOnDevToolsRawMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsRawMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const message_: Pointer; message_size: NativeUInt; var aHandled: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			message := uintptr(getVal(2))
			messageSize := cefTypes.NativeUInt(getVal(3))
			handled := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, message, messageSize, handled)
		},
	}
}

func makeTOnDialogClosed(cb TOnDialogClosed) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDialogClosed",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnDialogFileDialogEvent(cb TOnDialogFileDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDialogFileDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 8 : function(const browser: ICefBrowser; mode: TCefFileDialogMode; const title: ustring; const defaultFilePath: ustring; const acceptFilters: TStrings; const accept_extensions: TStrings; const accept_descriptions: TStrings; const callback: ICefFileDialogCallback): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			mode := cefTypes.TCefFileDialogMode(getVal(1))
			title := api.GoStr(getVal(2))
			defaultFilePath := api.GoStr(getVal(3))
			acceptFilters := lcl.AsStrings(getVal(4))
			acceptExtensions := lcl.AsStrings(getVal(5))
			acceptDescriptions := lcl.AsStrings(getVal(6))
			callback := AsCefFileDialogCallbackRef(getVal(7))
			ret := cb(browser, mode, title, defaultFilePath, acceptFilters, acceptExtensions, acceptDescriptions, callback)
			*(*bool)(getPtr(getVal(8))) = ret
		},
	}
}

func makeTOnDismissPermissionPromptEvent(cb TOnDismissPermissionPromptEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDismissPermissionPromptEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; prompt_id: uint64; result: TCefPermissionRequestResult);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			promptId := *(*uint64)(getPtr(getVal(2)))
			result := cefTypes.TCefPermissionRequestResult(getVal(3))
			cb(sender, browser, promptId, result)
		},
	}
}

func makeTOnDisplayAddressChangeEvent(cb TOnDisplayAddressChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayAddressChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const url: ustring);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			url := api.GoStr(getVal(2))
			cb(browser, frame, url)
		},
	}
}

func makeTOnDisplayAutoResizeEvent(cb TOnDisplayAutoResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayAutoResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; const new_size: PCefSize): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			newSize := *(*TCefSize)(getPtr(getVal(1)))
			ret := cb(browser, newSize)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnDisplayConsoleMessageEvent(cb TOnDisplayConsoleMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayConsoleMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; level: TCefLogSeverity; const message_: ustring; const source: ustring; line: Integer): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			level := cefTypes.TCefLogSeverity(getVal(1))
			message := api.GoStr(getVal(2))
			source := api.GoStr(getVal(3))
			line := int32(getVal(4))
			ret := cb(browser, level, message, source, line)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnDisplayContentsBoundsChangeEvent(cb TOnDisplayContentsBoundsChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayContentsBoundsChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; const new_bounds: PCefRect): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			newBounds := *(*TCefRect)(getPtr(getVal(1)))
			ret := cb(browser, newBounds)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnDisplayCursorChangeEvent(cb TOnDisplayCursorChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayCursorChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; cursor_: TCefCursorHandle; CursorType: TCefCursorType; const customCursorInfo: PCefCursorInfo; var aResult: boolean);
			browser := AsCefBrowserRef(getVal(0))
			cursor := cefTypes.TCefCursorHandle(getVal(1))
			cursorType := cefTypes.TCefCursorType(getVal(2))
			customCursorInfo := *(*TCefCursorInfo)(getPtr(getVal(3)))
			result := (*bool)(getPtr(getVal(4)))
			cb(browser, cursor, cursorType, customCursorInfo, result)
		},
	}
}

func makeTOnDisplayFaviconUrlChangeEvent(cb TOnDisplayFaviconUrlChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayFaviconUrlChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const iconUrls: TStrings);
			browser := AsCefBrowserRef(getVal(0))
			iconUrls := lcl.AsStrings(getVal(1))
			cb(browser, iconUrls)
		},
	}
}

func makeTOnDisplayFullScreenModeChangeEvent(cb TOnDisplayFullScreenModeChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayFullScreenModeChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; fullscreen: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			fullscreen := api.GoBool(getVal(1))
			cb(browser, fullscreen)
		},
	}
}

func makeTOnDisplayGetRootWindowScreenRectEvent(cb TOnDisplayGetRootWindowScreenRectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayGetRootWindowScreenRectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; rect_: PCefRect): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			rect := *(*TCefRect)(getPtr(getVal(1)))
			ret := cb(browser, rect)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnDisplayLoadingProgressChangeEvent(cb TOnDisplayLoadingProgressChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayLoadingProgressChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const progress: double);
			browser := AsCefBrowserRef(getVal(0))
			progress := *(*float64)(getPtr(getVal(1)))
			cb(browser, progress)
		},
	}
}

func makeTOnDisplayMediaAccessChangeEvent(cb TOnDisplayMediaAccessChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayMediaAccessChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; has_video_access: boolean; has_audio_access: boolean);
			browser := AsCefBrowserRef(getVal(0))
			hasVideoAccess := api.GoBool(getVal(1))
			hasAudioAccess := api.GoBool(getVal(2))
			cb(browser, hasVideoAccess, hasAudioAccess)
		},
	}
}

func makeTOnDisplayStatusMessageEvent(cb TOnDisplayStatusMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayStatusMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const value: ustring);
			browser := AsCefBrowserRef(getVal(0))
			value := api.GoStr(getVal(1))
			cb(browser, value)
		},
	}
}

func makeTOnDisplayTitleChangeEvent(cb TOnDisplayTitleChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayTitleChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const title: ustring);
			browser := AsCefBrowserRef(getVal(0))
			title := api.GoStr(getVal(1))
			cb(browser, title)
		},
	}
}

func makeTOnDisplayTooltipEvent(cb TOnDisplayTooltipEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDisplayTooltipEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; var text: ustring): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			text := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			ret := cb(browser, &text)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(text)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnDocumentAvailableInMainFrame(cb TOnDocumentAvailableInMainFrame) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDocumentAvailableInMainFrame",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: Tobject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnDomVisitorVisitEvent(cb TOnDomVisitorVisitEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDomVisitorVisitEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const document: ICefDomDocument);
			document := AsCefDomDocumentRef(getVal(0))
			cb(document)
		},
	}
}

func makeTOnDownloadBeforeDownloadEvent(cb TOnDownloadBeforeDownloadEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadBeforeDownloadEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const suggestedName: ustring; const callback: ICefBeforeDownloadCallback): boolean;
			browser := AsCefBrowserRef(getVal(0))
			downloadItem := AsCefDownloadItemRef(getVal(1))
			suggestedName := api.GoStr(getVal(2))
			callback := AsCefBeforeDownloadCallbackRef(getVal(3))
			ret := cb(browser, downloadItem, suggestedName, callback)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnDownloadCanDownloadEvent(cb TOnDownloadCanDownloadEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadCanDownloadEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const browser: ICefBrowser; const url: ustring; const request_method: ustring): boolean;
			browser := AsCefBrowserRef(getVal(0))
			url := api.GoStr(getVal(1))
			requestMethod := api.GoStr(getVal(2))
			ret := cb(browser, url, requestMethod)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnDownloadData(cb TOnDownloadData) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadData",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; data: Pointer; dataLength: NativeUInt);
			sender := lcl.AsObject(getVal(0))
			request := AsCefUrlRequestRef(getVal(1))
			data := uintptr(getVal(2))
			dataLength := cefTypes.NativeUInt(getVal(3))
			cb(sender, request, data, dataLength)
		},
	}
}

func makeTOnDownloadDownloadUpdatedEvent(cb TOnDownloadDownloadUpdatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadDownloadUpdatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const callback: ICefDownloadItemCallback);
			browser := AsCefBrowserRef(getVal(0))
			downloadItem := AsCefDownloadItemRef(getVal(1))
			callback := AsCefDownloadItemCallbackRef(getVal(2))
			cb(browser, downloadItem, callback)
		},
	}
}

func makeTOnDownloadImageCallbackDownloadImageFinishedEvent(cb TOnDownloadImageCallbackDownloadImageFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadImageCallbackDownloadImageFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const imageUrl: ustring; httpStatusCode: Integer; const image: ICefImage);
			imageUrl := api.GoStr(getVal(0))
			httpStatusCode := int32(getVal(1))
			image := AsCefImageRef(getVal(2))
			cb(imageUrl, httpStatusCode, image)
		},
	}
}

func makeTOnDownloadImageFinishedEvent(cb TOnDownloadImageFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadImageFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const imageUrl: ustring; httpStatusCode: Integer; const image: ICefImage);
			sender := lcl.AsObject(getVal(0))
			imageUrl := api.GoStr(getVal(1))
			httpStatusCode := int32(getVal(2))
			image := AsCefImageRef(getVal(3))
			cb(sender, imageUrl, httpStatusCode, image)
		},
	}
}

func makeTOnDownloadProgress(cb TOnDownloadProgress) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadProgress",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; current, total: Int64);
			sender := lcl.AsObject(getVal(0))
			request := AsCefUrlRequestRef(getVal(1))
			current := *(*int64)(getPtr(getVal(2)))
			total := *(*int64)(getPtr(getVal(3)))
			cb(sender, request, current, total)
		},
	}
}

func makeTOnDownloadUpdated(cb TOnDownloadUpdated) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadUpdated",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const downloadItem: ICefDownloadItem; const callback: ICefDownloadItemCallback);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			downloadItem := AsCefDownloadItemRef(getVal(2))
			callback := AsCefDownloadItemCallbackRef(getVal(3))
			cb(sender, browser, downloadItem, callback)
		},
	}
}

func makeTOnDragDragEnterEvent(cb TOnDragDragEnterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDragDragEnterEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const browser: ICefBrowser; const dragData: ICefDragData; mask: TCefDragOperations): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			dragData := AsCefDragDataRef(getVal(1))
			mask := cefTypes.TCefDragOperations(getVal(2))
			ret := cb(browser, dragData, mask)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnDragDraggableRegionsChangedEvent(cb TOnDragDraggableRegionsChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDragDraggableRegionsChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; regionsCount: NativeUInt; const regions: PCefDraggableRegionArray);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			regionsCount := cefTypes.NativeUInt(getVal(2))
			regionsPtr := getVal(3)
			regions := NewCefDraggableRegionArray(int(regionsCount), regionsPtr)
			cb(browser, frame, regionsCount, regions)
		},
	}
}

func makeTOnDragEnter(cb TOnDragEnter) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDragEnter",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const dragData: ICefDragData; mask: TCefDragOperations; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			dragData := AsCefDragDataRef(getVal(2))
			mask := cefTypes.TCefDragOperations(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, dragData, mask, result)
		},
	}
}

func makeTOnDraggableRegionsChanged(cb TOnDraggableRegionsChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDraggableRegionsChanged",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; regionsCount: NativeUInt; const regions: PCefDraggableRegionArray);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			regionsCount := cefTypes.NativeUInt(getVal(3))
			regionsPtr := getVal(4)
			regions := NewCefDraggableRegionArray(int(regionsCount), regionsPtr)
			cb(sender, browser, frame, regionsCount, regions)
		},
	}
}

func makeTOnEndTracingCallbackEndTracingCompleteEvent(cb TOnEndTracingCallbackEndTracingCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnEndTracingCallbackEndTracingCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const tracingFile: ustring);
			tracingFile := api.GoStr(getVal(0))
			cb(tracingFile)
		},
	}
}

func makeTOnExecuteTaskOnCefThread(cb TOnExecuteTaskOnCefThread) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnExecuteTaskOnCefThread",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aTaskID : cardinal);
			sender := lcl.AsObject(getVal(0))
			taskID := uint32(getVal(1))
			cb(sender, taskID)
		},
	}
}

func makeTOnFavIconUrlChange(cb TOnFavIconUrlChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFavIconUrlChange",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const iconUrls: TStrings);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			iconUrls := lcl.AsStrings(getVal(2))
			cb(sender, browser, iconUrls)
		},
	}
}

func makeTOnFileDialog(cb TOnFileDialog) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFileDialog",
		cb: func(getVal func(i int) uintptr) {
			// 10 : procedure(Sender: TObject; const browser: ICefBrowser; mode: TCefFileDialogMode; const title, defaultFilePath: ustring; const acceptFilters, accept_extensions, accept_descriptions: TStrings; const callback: ICefFileDialogCallback; var Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			mode := cefTypes.TCefFileDialogMode(getVal(2))
			title := api.GoStr(getVal(3))
			defaultFilePath := api.GoStr(getVal(4))
			acceptFilters := lcl.AsStrings(getVal(5))
			acceptExtensions := lcl.AsStrings(getVal(6))
			acceptDescriptions := lcl.AsStrings(getVal(7))
			callback := AsCefFileDialogCallbackRef(getVal(8))
			result := (*bool)(getPtr(getVal(9)))
			cb(sender, browser, mode, title, defaultFilePath, acceptFilters, acceptExtensions, acceptDescriptions, callback, result)
		},
	}
}

func makeTOnFilterEvent(cb TOnFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFilterEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnFindFindResultEvent(cb TOnFindFindResultEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFindFindResultEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(const browser: ICefBrowser; identifier: Integer; count: Integer; const selectionRect: PCefRect; activeMatchOrdinal: Integer; finalUpdate: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			identifier := int32(getVal(1))
			count := int32(getVal(2))
			selectionRect := *(*TCefRect)(getPtr(getVal(3)))
			activeMatchOrdinal := int32(getVal(4))
			finalUpdate := api.GoBool(getVal(5))
			cb(browser, identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
		},
	}
}

func makeTOnFindResult(cb TOnFindResult) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFindResult",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; identifier, count: Integer; const selectionRect: PCefRect; activeMatchOrdinal: Integer; finalUpdate: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			identifier := int32(getVal(2))
			count := int32(getVal(3))
			selectionRect := *(*TCefRect)(getPtr(getVal(4)))
			activeMatchOrdinal := int32(getVal(5))
			finalUpdate := api.GoBool(getVal(6))
			cb(sender, browser, identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
		},
	}
}

func makeTOnFocusEvent(cb TOnFocusEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFocusEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const view: ICefView);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			cb(sender, view)
		},
	}
}

func makeTOnFocusGotFocusEvent(cb TOnFocusGotFocusEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFocusGotFocusEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnFocusSetFocusEvent(cb TOnFocusSetFocusEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFocusSetFocusEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; source: TCefFocusSource): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			source := cefTypes.TCefFocusSource(getVal(1))
			ret := cb(browser, source)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnFocusTakeFocusEvent(cb TOnFocusTakeFocusEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFocusTakeFocusEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; next: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			next := api.GoBool(getVal(1))
			cb(browser, next)
		},
	}
}

func makeTOnFocusedNodeChangedEvent(cb TOnFocusedNodeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFocusedNodeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const node: ICefDomNode);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			node := AsCefDomNodeRef(getVal(2))
			cb(browser, frame, node)
		},
	}
}

func makeTOnFrameAttached(cb TOnFrameAttached) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameAttached",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; reattached: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			reattached := api.GoBool(getVal(3))
			cb(sender, browser, frame, reattached)
		},
	}
}

func makeTOnFrameCreated(cb TOnFrameCreated) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameCreated",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			cb(sender, browser, frame)
		},
	}
}

func makeTOnFrameDestroyed(cb TOnFrameDestroyed) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameDestroyed",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			cb(sender, browser, frame)
		},
	}
}

func makeTOnFrameDetached(cb TOnFrameDetached) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameDetached",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			cb(sender, browser, frame)
		},
	}
}

func makeTOnFrameFrameAttachedEvent(cb TOnFrameFrameAttachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameFrameAttachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; reattached: boolean);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			reattached := api.GoBool(getVal(2))
			cb(browser, frame, reattached)
		},
	}
}

func makeTOnFrameFrameCreatedEvent(cb TOnFrameFrameCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameFrameCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			cb(browser, frame)
		},
	}
}

func makeTOnFrameFrameDestroyedEvent(cb TOnFrameFrameDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameFrameDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			cb(browser, frame)
		},
	}
}

func makeTOnFrameFrameDetachedEvent(cb TOnFrameFrameDetachedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameFrameDetachedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			cb(browser, frame)
		},
	}
}

func makeTOnFrameMainFrameChangedEvent(cb TOnFrameMainFrameChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameMainFrameChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const old_frame: ICefFrame; const new_frame: ICefFrame);
			browser := AsCefBrowserRef(getVal(0))
			oldFrame := AsCefFrameRef(getVal(1))
			newFrame := AsCefFrameRef(getVal(2))
			cb(browser, oldFrame, newFrame)
		},
	}
}

func makeTOnFullScreenModeChange(cb TOnFullScreenModeChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFullScreenModeChange",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; fullscreen: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			fullscreen := api.GoBool(getVal(2))
			cb(sender, browser, fullscreen)
		},
	}
}

func makeTOnGestureCommandEvent(cb TOnGestureCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGestureCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; gesture_command: TCefGestureCommand; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			gestureCommand := cefTypes.TCefGestureCommand(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browserView, gestureCommand, result)
		},
	}
}

func makeTOnGetAccessibilityHandler(cb TOnGetAccessibilityHandler) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetAccessibilityHandler",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var aAccessibilityHandler : ICefAccessibilityHandler);
			sender := lcl.AsObject(getVal(0))
			var accessibilityHandler IEngAccessibilityHandler
			accessibilityHandler = AsEngAccessibilityHandler(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &accessibilityHandler)
			if accessibilityHandler != nil {
				*(*uintptr)(getPtr(getVal(1))) = accessibilityHandler.Instance()
			}
		},
	}
}

func makeTOnGetAudioParametersEvent(cb TOnGetAudioParametersEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetAudioParametersEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var params: TCefAudioParameters; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			params := (*TCefAudioParameters)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, params, result)
			*(*TCefAudioParameters)(getPtr(getVal(2))) = *params
		},
	}
}

func makeTOnGetAuthCredentials(cb TOnGetAuthCredentials) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetAuthCredentials",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetBrowserRuntimeStyleEvent(cb TOnGetBrowserRuntimeStyleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetBrowserRuntimeStyleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; var aResult : TCefRuntimeStyle);
			sender := lcl.AsObject(getVal(0))
			result := (*cefTypes.TCefRuntimeStyle)(getPtr(getVal(1)))
			cb(sender, result)
		},
	}
}

func makeTOnGetChromeToolbarTypeEvent(cb TOnGetChromeToolbarTypeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetChromeToolbarTypeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; var aChromeToolbarType: TCefChromeToolbarType);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			chromeToolbarType := (*TCefChromeToolbarType)(getPtr(getVal(2)))
			cb(sender, browserView, chromeToolbarType)
		},
	}
}

func makeTOnGetDataResourceEvent(cb TOnGetDataResourceEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetDataResourceEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(resourceId: Integer; out data: Pointer; out dataSize: NativeUInt; var aResult : Boolean);
			resourceId := int32(getVal(0))
			data := (*uintptr)(getPtr(getVal(1)))
			dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(resourceId, data, dataSize, result)
		},
	}
}

func makeTOnGetDataResourceForScaleEvent(cb TOnGetDataResourceForScaleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetDataResourceForScaleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(resourceId: Integer; scaleFactor: TCefScaleFactor; out data: Pointer; out dataSize: NativeUInt; var aResult : Boolean);
			resourceId := int32(getVal(0))
			scaleFactor := cefTypes.TCefScaleFactor(getVal(1))
			data := (*uintptr)(getPtr(getVal(2)))
			dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(3)))
			result := (*bool)(getPtr(getVal(4)))
			cb(resourceId, scaleFactor, data, dataSize, result)
		},
	}
}

func makeTOnGetDefaultClientEvent(cb TOnGetDefaultClientEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetDefaultClientEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aClient : ICefClient);
			var client IEngClient
			client = AsEngClient(*(*uintptr)(getPtr(getVal(0))))
			cb(&client)
			if client != nil {
				*(*uintptr)(getPtr(getVal(0))) = client.Instance()
			}
		},
	}
}

func makeTOnGetDefaultRequestContextHandlerEvent(cb TOnGetDefaultRequestContextHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetDefaultRequestContextHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aRequestContextHandler : ICefRequestContextHandler);
			var requestContextHandler IEngRequestContextHandler
			requestContextHandler = AsEngRequestContextHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&requestContextHandler)
			if requestContextHandler != nil {
				*(*uintptr)(getPtr(getVal(0))) = requestContextHandler.Instance()
			}
		},
	}
}

func makeTOnGetDelegateForPopupBrowserViewEvent(cb TOnGetDelegateForPopupBrowserViewEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetDelegateForPopupBrowserViewEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetHeightForWidthEvent(cb TOnGetHeightForWidthEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetHeightForWidthEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const view: ICefView; width: Integer; var aResult: Integer);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			width := int32(getVal(2))
			result := (*int32)(getPtr(getVal(3)))
			cb(sender, view, width, result)
		},
	}
}

func makeTOnGetInitialBoundsEvent(cb TOnGetInitialBoundsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetInitialBoundsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : TCefRect);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*TCefRect)(getPtr(getVal(2)))
			cb(sender, window, result)
			*(*TCefRect)(getPtr(getVal(2))) = *result
		},
	}
}

func makeTOnGetInitialShowStateEvent(cb TOnGetInitialShowStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetInitialShowStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : TCefShowState);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*cefTypes.TCefShowState)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnGetLinuxWindowPropertiesEvent(cb TOnGetLinuxWindowPropertiesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetLinuxWindowPropertiesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const window_: ICefWindow; var properties: TLinuxWindowProperties; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			propertiesPtr := *(*tLinuxWindowProperties)(getPtr(getVal(2)))
			properties := propertiesPtr.ToGo()
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, window, &properties, result)
			if r := properties.ToPas(); r != nil {
				*(*tLinuxWindowProperties)(getPtr(getVal(2))) = *r
			}
		},
	}
}

func makeTOnGetLocalizedStringEvent(cb TOnGetLocalizedStringEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetLocalizedStringEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(stringId: Integer; out stringVal: ustring; var aResult : Boolean);
			stringId := int32(getVal(0))
			var stringVal string
			result := (*bool)(getPtr(getVal(2)))
			cb(stringId, &stringVal, result)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(stringVal)
		},
	}
}

func makeTOnGetMaximumSizeEvent(cb TOnGetMaximumSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetMaximumSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			result := (*TCefSize)(getPtr(getVal(2)))
			cb(sender, view, result)
			*(*TCefSize)(getPtr(getVal(2))) = *result
		},
	}
}

func makeTOnGetMinimumSizeEvent(cb TOnGetMinimumSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetMinimumSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			result := (*TCefSize)(getPtr(getVal(2)))
			cb(sender, view, result)
			*(*TCefSize)(getPtr(getVal(2))) = *result
		},
	}
}

func makeTOnGetPDFPaperSizeEvent(cb TOnGetPDFPaperSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetPDFPaperSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; deviceUnitsPerInch: Integer; var aResult : TCefSize);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			deviceUnitsPerInch := int32(getVal(2))
			result := (*TCefSize)(getPtr(getVal(3)))
			cb(sender, browser, deviceUnitsPerInch, result)
			*(*TCefSize)(getPtr(getVal(3))) = *result
		},
	}
}

func makeTOnGetParentWindowEvent(cb TOnGetParentWindowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetParentWindowEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetPreferredSizeEvent(cb TOnGetPreferredSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetPreferredSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const view: ICefView; var aResult : TCefSize);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			result := (*TCefSize)(getPtr(getVal(2)))
			cb(sender, view, result)
			*(*TCefSize)(getPtr(getVal(2))) = *result
		},
	}
}

func makeTOnGetResourceHandler(cb TOnGetResourceHandler) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetResourceHandler",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetResourceRequestHandler(cb TOnGetResourceRequestHandler) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetResourceRequestHandler",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetResourceResponseFilter(cb TOnGetResourceResponseFilter) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetResourceResponseFilter",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnGetRootScreenRect(cb TOnGetRootScreenRect) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetRootScreenRect",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var rect: TCefRect; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			rect := (*TCefRect)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, rect, result)
			*(*TCefRect)(getPtr(getVal(2))) = *rect
		},
	}
}

func makeTOnGetRootWindowScreenRect(cb TOnGetRootWindowScreenRect) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetRootWindowScreenRect",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; rect_: PCefRect; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			rect := *(*TCefRect)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, rect, result)
		},
	}
}

func makeTOnGetScreenInfo(cb TOnGetScreenInfo) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetScreenInfo",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var screenInfo: TCefScreenInfo; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			screenInfo := (*TCefScreenInfo)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, screenInfo, result)
			*(*TCefScreenInfo)(getPtr(getVal(2))) = *screenInfo
		},
	}
}

func makeTOnGetScreenPoint(cb TOnGetScreenPoint) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetScreenPoint",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; viewX, viewY: Integer; var screenX, screenY: Integer; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			viewX := int32(getVal(2))
			viewY := int32(getVal(3))
			screenX := (*int32)(getPtr(getVal(4)))
			screenY := (*int32)(getPtr(getVal(5)))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, viewX, viewY, screenX, screenY, result)
		},
	}
}

func makeTOnGetTitlebarHeightEvent(cb TOnGetTitlebarHeightEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetTitlebarHeightEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const window_: ICefWindow; var titlebar_height: Single; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			titlebarHeight := (*float32)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, window, titlebarHeight, result)
		},
	}
}

func makeTOnGetTouchHandleSize(cb TOnGetTouchHandleSize) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetTouchHandleSize",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; orientation: TCefHorizontalAlignment; var size: TCefSize);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			orientation := cefTypes.TCefHorizontalAlignment(getVal(2))
			size := (*TCefSize)(getPtr(getVal(3)))
			cb(sender, browser, orientation, size)
			*(*TCefSize)(getPtr(getVal(3))) = *size
		},
	}
}

func makeTOnGetViewRect(cb TOnGetViewRect) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetViewRect",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; var rect: TCefRect);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			rect := (*TCefRect)(getPtr(getVal(2)))
			cb(sender, browser, rect)
			*(*TCefRect)(getPtr(getVal(2))) = *rect
		},
	}
}

func makeTOnGetWindowRuntimeStyleEvent(cb TOnGetWindowRuntimeStyleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetWindowRuntimeStyleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; var aResult : TCefRuntimeStyle);
			sender := lcl.AsObject(getVal(0))
			result := (*cefTypes.TCefRuntimeStyle)(getPtr(getVal(1)))
			cb(sender, result)
		},
	}
}

func makeTOnGotFocus(cb TOnGotFocus) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGotFocus",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnHandledMessageEvent(cb TOnHandledMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnHandledMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var aMessage: TMessage; var aHandled : boolean);
			sender := lcl.AsObject(getVal(0))
			message := (*types.TMessage)(getPtr(getVal(1)))
			handled := (*bool)(getPtr(getVal(2)))
			cb(sender, message, handled)
		},
	}
}

func makeTOnHttpRequest(cb TOnHttpRequest) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnHttpRequest",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			clientAddress := api.GoStr(getVal(3))
			request := AsCefRequestRef(getVal(4))
			cb(sender, server, connectionId, clientAddress, request)
		},
	}
}

func makeTOnIMECommitTextEvent(cb TOnIMECommitTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIMECommitTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aText : ustring; const replacement_range : PCefRange; relative_cursor_pos : integer);
			sender := lcl.AsObject(getVal(0))
			text := api.GoStr(getVal(1))
			replacementRange := *(*TCefRange)(getPtr(getVal(2)))
			relativeCursorPos := int32(getVal(3))
			cb(sender, text, replacementRange, relativeCursorPos)
		},
	}
}

func makeTOnIMECompositionRangeChanged(cb TOnIMECompositionRangeChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIMECompositionRangeChanged",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const selected_range: PCefRange; character_boundsCount: NativeUInt; const character_bounds: PCefRect);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			selectedRange := *(*TCefRange)(getPtr(getVal(2)))
			characterBoundsCount := cefTypes.NativeUInt(getVal(3))
			characterBounds := *(*TCefRect)(getPtr(getVal(4)))
			cb(sender, browser, selectedRange, characterBoundsCount, characterBounds)
		},
	}
}

func makeTOnIMESetCompositionEvent(cb TOnIMESetCompositionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIMESetCompositionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const aText : ustring; const underlines : TCefCompositionUnderlineDynArray; const replacement_range, selection_range : TCefRange);
			sender := lcl.AsObject(getVal(0))
			text := api.GoStr(getVal(1))
			underlinesPtr := getVal(2)
			underlinesLen := int32(getVal(3))
			underlines := NewCefCompositionUnderlineArray(int(underlinesLen), underlinesPtr)
			replacementRange := *(*TCefRange)(getPtr(getVal(4)))
			selectionRange := *(*TCefRange)(getPtr(getVal(5)))
			cb(sender, text, underlines, replacementRange, selectionRange)
		},
	}
}

func makeTOnInitFilterEvent(cb TOnInitFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnInitFilterEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(sender, result)
		},
	}
}

func makeTOnIsChromeAppMenuItemEnabledEvent(cb TOnIsChromeAppMenuItemEnabledEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsChromeAppMenuItemEnabledEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; command_id: integer; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, commandId, result)
		},
	}
}

func makeTOnIsChromeAppMenuItemVisibleEvent(cb TOnIsChromeAppMenuItemVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsChromeAppMenuItemVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; command_id: integer; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, commandId, result)
		},
	}
}

func makeTOnIsChromePageActionIconVisibleEvent(cb TOnIsChromePageActionIconVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsChromePageActionIconVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; icon_type: TCefChromePageActionIconType; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			iconType := cefTypes.TCefChromePageActionIconType(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, iconType, result)
		},
	}
}

func makeTOnIsChromeToolbarButtonVisibleEvent(cb TOnIsChromeToolbarButtonVisibleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsChromeToolbarButtonVisibleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; button_type: TCefChromeToolbarButtonType; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			buttonType := cefTypes.TCefChromeToolbarButtonType(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, buttonType, result)
		},
	}
}

func makeTOnIsFramelessEvent(cb TOnIsFramelessEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsFramelessEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnIsWindowModalDialogEvent(cb TOnIsWindowModalDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsWindowModalDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnJsDialogBeforeUnloadDialogEvent(cb TOnJsDialogBeforeUnloadDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnJsDialogBeforeUnloadDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const messageText: ustring; isReload: Boolean; const callback: ICefJsDialogCallback): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			messageText := api.GoStr(getVal(1))
			isReload := api.GoBool(getVal(2))
			callback := AsCefJsDialogCallbackRef(getVal(3))
			ret := cb(browser, messageText, isReload, callback)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnJsDialogDialogClosedEvent(cb TOnJsDialogDialogClosedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnJsDialogDialogClosedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnJsDialogJsdialogEvent(cb TOnJsDialogJsdialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnJsDialogJsdialogEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnJsDialogResetDialogStateEvent(cb TOnJsDialogResetDialogStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnJsDialogResetDialogStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnJsdialog(cb TOnJsdialog) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnJsdialog",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnKeyEvent(cb TOnKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			event := *(*TCefKeyEvent)(getPtr(getVal(2)))
			osEvent := cefTypes.TCefEventHandle(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, event, osEvent, result)
		},
	}
}

func makeTOnKeyboardKeyEvent(cb TOnKeyboardKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnKeyboardKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			event := *(*TCefKeyEvent)(getPtr(getVal(1)))
			osEvent := cefTypes.TCefEventHandle(getVal(2))
			ret := cb(browser, event, osEvent)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnKeyboardPreKeyEvent(cb TOnKeyboardPreKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnKeyboardPreKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out isKeyboardShortcut: Boolean): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			event := *(*TCefKeyEvent)(getPtr(getVal(1)))
			osEvent := cefTypes.TCefEventHandle(getVal(2))
			isKeyboardShortcut := (*bool)(getPtr(getVal(3)))
			ret := cb(browser, event, osEvent, isKeyboardShortcut)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnLayoutChangedEvent(cb TOnLayoutChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLayoutChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const view: ICefView; new_bounds: TCefRect);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			newBounds := *(*TCefRect)(getPtr(getVal(2)))
			cb(sender, view, newBounds)
		},
	}
}

func makeTOnLifeSpanAfterCreatedEvent(cb TOnLifeSpanAfterCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanAfterCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnLifeSpanBeforeCloseEvent(cb TOnLifeSpanBeforeCloseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanBeforeCloseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnLifeSpanBeforeDevToolsPopupEvent(cb TOnLifeSpanBeforeDevToolsPopupEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanBeforeDevToolsPopupEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(const browser: ICefBrowser; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var use_default_window: boolean);
			browser := AsCefBrowserRef(getVal(0))
			windowInfoPtr := *(*tCefWindowInfo)(getPtr(getVal(1)))
			windowInfo := windowInfoPtr.ToGo()
			var client IEngClient
			client = AsEngClient(*(*uintptr)(getPtr(getVal(2))))
			settingsPtr := *(*tCefBrowserSettings)(getPtr(getVal(3)))
			settings := settingsPtr.ToGo()
			extraInfo := ICefDictionaryValue(AsCefDictionaryValueRef(*(*uintptr)(getPtr(getVal(4)))))
			useDefaultWindow := (*bool)(getPtr(getVal(5)))
			cb(browser, &windowInfo, &client, &settings, &extraInfo, useDefaultWindow)
			if r := windowInfo.ToPas(); r != nil {
				*(*tCefWindowInfo)(getPtr(getVal(1))) = *r
			}
			if client != nil {
				*(*uintptr)(getPtr(getVal(2))) = client.Instance()
			}
			if r := settings.ToPas(); r != nil {
				*(*tCefBrowserSettings)(getPtr(getVal(3))) = *r
			}
			if extraInfo != nil {
				*(*uintptr)(getPtr(getVal(4))) = extraInfo.Instance()
			}
		},
	}
}

func makeTOnLifeSpanBeforePopupAbortedEvent(cb TOnLifeSpanBeforePopupAbortedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanBeforePopupAbortedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; popup_id: Integer);
			browser := AsCefBrowserRef(getVal(0))
			popupId := int32(getVal(1))
			cb(browser, popupId)
		},
	}
}

func makeTOnLifeSpanBeforePopupEvent(cb TOnLifeSpanBeforePopupEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanBeforePopupEvent",
		cb: func(getVal func(i int) uintptr) {
			// 13 : function(const browser: ICefBrowser; const frame: ICefFrame; popup_id: Integer; const targetUrl: ustring; const targetFrameName: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; const popupFeatures: TCefPopupFeatures; var windowInfo: TCefWindowInfo; var client: ICefClient; var settings: TCefBrowserSettings; var extra_info: ICefDictionaryValue; var noJavascriptAccess: Boolean): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			popupId := int32(getVal(2))
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
			ret := cb(browser, frame, popupId, targetUrl, targetFrameName, targetDisposition, userGesture, popupFeatures, &windowInfo, &client, &settings, &extraInfo, noJavascriptAccess)
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
			*(*bool)(getPtr(getVal(13))) = ret
		},
	}
}

func makeTOnLifeSpanDoCloseEvent(cb TOnLifeSpanDoCloseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLifeSpanDoCloseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : function(const browser: ICefBrowser): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			ret := cb(browser)
			*(*bool)(getPtr(getVal(1))) = ret
		},
	}
}

func makeTOnLoadEnd(cb TOnLoadEnd) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadEnd",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			httpStatusCode := int32(getVal(3))
			cb(sender, browser, frame, httpStatusCode)
		},
	}
}

func makeTOnLoadError(cb TOnLoadError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadError",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText, failedUrl: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			errorCode := cefTypes.TCefErrorCode(getVal(3))
			errorText := api.GoStr(getVal(4))
			failedUrl := api.GoStr(getVal(5))
			cb(sender, browser, frame, errorCode, errorText, failedUrl)
		},
	}
}

func makeTOnLoadLoadEndEvent(cb TOnLoadLoadEndEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadLoadEndEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			httpStatusCode := int32(getVal(2))
			cb(browser, frame, httpStatusCode)
		},
	}
}

func makeTOnLoadLoadErrorEvent(cb TOnLoadLoadErrorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadLoadErrorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText: ustring; const failedUrl: ustring);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			errorCode := cefTypes.TCefErrorCode(getVal(2))
			errorText := api.GoStr(getVal(3))
			failedUrl := api.GoStr(getVal(4))
			cb(browser, frame, errorCode, errorText, failedUrl)
		},
	}
}

func makeTOnLoadLoadStartEvent(cb TOnLoadLoadStartEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadLoadStartEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			transitionType := cefTypes.TCefTransitionType(getVal(2))
			cb(browser, frame, transitionType)
		},
	}
}

func makeTOnLoadLoadingStateChangeEvent(cb TOnLoadLoadingStateChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadLoadingStateChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; isLoading: Boolean; canGoBack: Boolean; canGoForward: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			isLoading := api.GoBool(getVal(1))
			canGoBack := api.GoBool(getVal(2))
			canGoForward := api.GoBool(getVal(3))
			cb(browser, isLoading, canGoBack, canGoForward)
		},
	}
}

func makeTOnLoadStart(cb TOnLoadStart) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadStart",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			transitionType := cefTypes.TCefTransitionType(getVal(3))
			cb(sender, browser, frame, transitionType)
		},
	}
}

func makeTOnLoadingProgressChange(cb TOnLoadingProgressChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadingProgressChange",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const progress: double);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			progress := *(*float64)(getPtr(getVal(2)))
			cb(sender, browser, progress)
		},
	}
}

func makeTOnLoadingStateChange(cb TOnLoadingStateChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLoadingStateChange",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; isLoading, canGoBack, canGoForward: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			isLoading := api.GoBool(getVal(2))
			canGoBack := api.GoBool(getVal(3))
			canGoForward := api.GoBool(getVal(4))
			cb(sender, browser, isLoading, canGoBack, canGoForward)
		},
	}
}

func makeTOnMainFrameChanged(cb TOnMainFrameChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMainFrameChanged",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const old_frame, new_frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			oldFrame := AsCefFrameRef(getVal(2))
			newFrame := AsCefFrameRef(getVal(3))
			cb(sender, browser, oldFrame, newFrame)
		},
	}
}

func makeTOnMediaAccessChange(cb TOnMediaAccessChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaAccessChange",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; has_video_access, has_audio_access: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			hasVideoAccess := api.GoBool(getVal(2))
			hasAudioAccess := api.GoBool(getVal(3))
			cb(sender, browser, hasVideoAccess, hasAudioAccess)
		},
	}
}

func makeTOnMediaObserverRouteMessageReceivedEvent(cb TOnMediaObserverRouteMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaObserverRouteMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const route: ICefMediaRoute; const message_: ustring);
			route := AsCefMediaRouteRef(getVal(0))
			message := api.GoStr(getVal(1))
			cb(route, message)
		},
	}
}

func makeTOnMediaObserverRouteStateChangedEvent(cb TOnMediaObserverRouteStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaObserverRouteStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const route: ICefMediaRoute; state: TCefMediaRouteConnectionState);
			route := AsCefMediaRouteRef(getVal(0))
			state := cefTypes.TCefMediaRouteConnectionState(getVal(1))
			cb(route, state)
		},
	}
}

func makeTOnMediaObserverRoutesEvent(cb TOnMediaObserverRoutesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaObserverRoutesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const routes: TCefMediaRouteArray);
			routesPtr := getVal(0)
			routesLen := int32(getVal(1))
			routes := NewCefMediaRouteArray(int(routesLen), routesPtr)
			cb(routes)
		},
	}
}

func makeTOnMediaObserverSinksEvent(cb TOnMediaObserverSinksEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaObserverSinksEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const sinks: TCefMediaSinkArray);
			sinksPtr := getVal(0)
			sinksLen := int32(getVal(1))
			sinks := NewCefMediaSinkArray(int(sinksLen), sinksPtr)
			cb(sinks)
		},
	}
}

func makeTOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent(cb TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaRouteCreateCallbackMediaRouteCreateFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(result: TCefMediaRouterCreateResult; const error: ustring; const route: ICefMediaRoute);
			result := cefTypes.TCefMediaRouterCreateResult(getVal(0))
			error_ := api.GoStr(getVal(1))
			route := AsCefMediaRouteRef(getVal(2))
			cb(result, error_, route)
		},
	}
}

func makeTOnMediaRouteCreateFinishedEvent(cb TOnMediaRouteCreateFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaRouteCreateFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; result: TCefMediaRouterCreateResult; const error: ustring; const route: ICefMediaRoute);
			sender := lcl.AsObject(getVal(0))
			result := cefTypes.TCefMediaRouterCreateResult(getVal(1))
			error_ := api.GoStr(getVal(2))
			route := AsCefMediaRouteRef(getVal(3))
			cb(sender, result, error_, route)
		},
	}
}

func makeTOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent(cb TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaSinkDeviceInfoCallbackMediaSinkDeviceInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const ip_address: ustring; port: integer; const model_name: ustring);
			ipAddress := api.GoStr(getVal(0))
			port := int32(getVal(1))
			modelName := api.GoStr(getVal(2))
			cb(ipAddress, port, modelName)
		},
	}
}

func makeTOnMediaSinkDeviceInfoEvent(cb TOnMediaSinkDeviceInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMediaSinkDeviceInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const ip_address: ustring; port: integer; const model_name: ustring);
			sender := lcl.AsObject(getVal(0))
			ipAddress := api.GoStr(getVal(1))
			port := int32(getVal(2))
			modelName := api.GoStr(getVal(3))
			cb(sender, ipAddress, port, modelName)
		},
	}
}

func makeTOnMenuButtonMenuButtonPressedEvent(cb TOnMenuButtonMenuButtonPressedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuButtonMenuButtonPressedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const menu_button: ICefMenuButton; const screen_point: TCefPoint; const button_pressed_lock: ICefMenuButtonPressedLock);
			menuButton := AsCefMenuButtonRef(getVal(0))
			screenPoint := *(*TCefPoint)(getPtr(getVal(1)))
			buttonPressedLock := AsCefMenuButtonPressedLockRef(getVal(2))
			cb(menuButton, screenPoint, buttonPressedLock)
		},
	}
}

func makeTOnMenuButtonPressedEvent(cb TOnMenuButtonPressedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuButtonPressedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const menu_button: ICefMenuButton; const screen_point: TCefPoint; const button_pressed_lock: ICefMenuButtonPressedLock);
			sender := lcl.AsObject(getVal(0))
			menuButton := AsCefMenuButtonRef(getVal(1))
			screenPoint := *(*TCefPoint)(getPtr(getVal(2)))
			buttonPressedLock := AsCefMenuButtonPressedLockRef(getVal(3))
			cb(sender, menuButton, screenPoint, buttonPressedLock)
		},
	}
}

func makeTOnMenuModelExecuteCommandEvent(cb TOnMenuModelExecuteCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelExecuteCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const menuModel: ICefMenuModel; commandId: Integer; eventFlags: TCefEventFlags);
			menuModel := AsCefMenuModelRef(getVal(0))
			commandId := int32(getVal(1))
			eventFlags := cefTypes.TCefEventFlags(getVal(2))
			cb(menuModel, commandId, eventFlags)
		},
	}
}

func makeTOnMenuModelFormatLabelEvent(cb TOnMenuModelFormatLabelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelFormatLabelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const menuModel: ICefMenuModel; var label_: ustring): boolean;
			menuModel := AsCefMenuModelRef(getVal(0))
			label := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			ret := cb(menuModel, &label)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(label)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnMenuModelMenuClosedEvent(cb TOnMenuModelMenuClosedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelMenuClosedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const menuModel: ICefMenuModel);
			menuModel := AsCefMenuModelRef(getVal(0))
			cb(menuModel)
		},
	}
}

func makeTOnMenuModelMenuWillShowEvent(cb TOnMenuModelMenuWillShowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelMenuWillShowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const menuModel: ICefMenuModel);
			menuModel := AsCefMenuModelRef(getVal(0))
			cb(menuModel)
		},
	}
}

func makeTOnMenuModelMouseOutsideMenuEvent(cb TOnMenuModelMouseOutsideMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelMouseOutsideMenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const menuModel: ICefMenuModel; const screenPoint: PCefPoint);
			menuModel := AsCefMenuModelRef(getVal(0))
			screenPoint := *(*TCefPoint)(getPtr(getVal(1)))
			cb(menuModel, screenPoint)
		},
	}
}

func makeTOnMenuModelUnhandledCloseSubmenuEvent(cb TOnMenuModelUnhandledCloseSubmenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelUnhandledCloseSubmenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const menuModel: ICefMenuModel; isRTL: boolean);
			menuModel := AsCefMenuModelRef(getVal(0))
			isRTL := api.GoBool(getVal(1))
			cb(menuModel, isRTL)
		},
	}
}

func makeTOnMenuModelUnhandledOpenSubmenuEvent(cb TOnMenuModelUnhandledOpenSubmenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMenuModelUnhandledOpenSubmenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const menuModel: ICefMenuModel; isRTL: boolean);
			menuModel := AsCefMenuModelRef(getVal(0))
			isRTL := api.GoBool(getVal(1))
			cb(menuModel, isRTL)
		},
	}
}

func makeTOnNavigationEntryVisitorVisitEvent(cb TOnNavigationEntryVisitorVisitEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNavigationEntryVisitorVisitEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const entry: ICefNavigationEntry; current: Boolean; index: Integer; total: Integer): Boolean;
			entry := AsCefNavigationEntryRef(getVal(0))
			current := api.GoBool(getVal(1))
			index := int32(getVal(2))
			total := int32(getVal(3))
			ret := cb(entry, current, index, total)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnNavigationVisitorResultAvailableEvent(cb TOnNavigationVisitorResultAvailableEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNavigationVisitorResultAvailableEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const entry: ICefNavigationEntry; current: Boolean; index, total: Integer; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			entry := AsCefNavigationEntryRef(getVal(1))
			current := api.GoBool(getVal(2))
			index := int32(getVal(3))
			total := int32(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, entry, current, index, total, result)
		},
	}
}

func makeTOnOpenUrlFromTab(cb TOnOpenUrlFromTab) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnOpenUrlFromTab",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const targetUrl: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			targetUrl := api.GoStr(getVal(3))
			targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(4))
			userGesture := api.GoBool(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, frame, targetUrl, targetDisposition, userGesture, result)
		},
	}
}

func makeTOnPaint(cb TOnPaint) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPaint",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnParentViewChangedEvent(cb TOnParentViewChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnParentViewChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const view: ICefView; added: boolean; const parent: ICefView);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			added := api.GoBool(getVal(2))
			parent := AsCefViewRef(getVal(3))
			cb(sender, view, added, parent)
		},
	}
}

func makeTOnPdfPrintCallbackPdfPrintFinishedEvent(cb TOnPdfPrintCallbackPdfPrintFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPdfPrintCallbackPdfPrintFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const path: ustring; ok: Boolean);
			path := api.GoStr(getVal(0))
			ok := api.GoBool(getVal(1))
			cb(path, ok)
		},
	}
}

func makeTOnPdfPrintFinishedEvent(cb TOnPdfPrintFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPdfPrintFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResultOK : boolean);
			sender := lcl.AsObject(getVal(0))
			resultOK := api.GoBool(getVal(1))
			cb(sender, resultOK)
		},
	}
}

func makeTOnPermissionDismissPermissionPromptEvent(cb TOnPermissionDismissPermissionPromptEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPermissionDismissPermissionPromptEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; prompt_id: uint64; result: TCefPermissionRequestResult);
			browser := AsCefBrowserRef(getVal(0))
			promptId := *(*uint64)(getPtr(getVal(1)))
			result := cefTypes.TCefPermissionRequestResult(getVal(2))
			cb(browser, promptId, result)
		},
	}
}

func makeTOnPermissionRequestMediaAccessPermissionEvent(cb TOnPermissionRequestMediaAccessPermissionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPermissionRequestMediaAccessPermissionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefMediaAccessCallback): boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			requestingOrigin := api.GoStr(getVal(2))
			requestedPermissions := uint32(getVal(3))
			callback := AsCefMediaAccessCallbackRef(getVal(4))
			ret := cb(browser, frame, requestingOrigin, requestedPermissions, callback)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnPermissionShowPermissionPromptEvent(cb TOnPermissionShowPermissionPromptEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPermissionShowPermissionPromptEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; prompt_id: uint64; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefPermissionPromptCallback): boolean;
			browser := AsCefBrowserRef(getVal(0))
			promptId := *(*uint64)(getPtr(getVal(1)))
			requestingOrigin := api.GoStr(getVal(2))
			requestedPermissions := uint32(getVal(3))
			callback := AsCefPermissionPromptCallbackRef(getVal(4))
			ret := cb(browser, promptId, requestingOrigin, requestedPermissions, callback)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnPopupBrowserViewCreatedEvent(cb TOnPopupBrowserViewCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPopupBrowserViewCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const Sender: TObject; const browser_view, popup_browser_view: ICefBrowserView; is_devtools: boolean; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			popupBrowserView := AsCefBrowserViewRef(getVal(2))
			isDevtools := api.GoBool(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browserView, popupBrowserView, isDevtools, result)
		},
	}
}

func makeTOnPopupShow(cb TOnPopupShow) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPopupShow",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; show: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			show := api.GoBool(getVal(2))
			cb(sender, browser, show)
		},
	}
}

func makeTOnPopupSize(cb TOnPopupSize) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPopupSize",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const rect: PCefRect);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			rect := *(*TCefRect)(getPtr(getVal(2)))
			cb(sender, browser, rect)
		},
	}
}

func makeTOnPreKeyEvent(cb TOnPreKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPreKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const event: PCefKeyEvent; osEvent: TCefEventHandle; out isKeyboardShortcut: Boolean; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			event := *(*TCefKeyEvent)(getPtr(getVal(2)))
			osEvent := cefTypes.TCefEventHandle(getVal(3))
			isKeyboardShortcut := (*bool)(getPtr(getVal(4)))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, event, osEvent, isKeyboardShortcut, result)
		},
	}
}

func makeTOnPreferenceChangedEvent(cb TOnPreferenceChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPreferenceChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const name_: ustring);
			sender := lcl.AsObject(getVal(0))
			name := api.GoStr(getVal(1))
			cb(sender, name)
		},
	}
}

func makeTOnPreferenceObserverPreferenceChangedEvent(cb TOnPreferenceObserverPreferenceChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPreferenceObserverPreferenceChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const name: ustring);
			name := api.GoStr(getVal(0))
			cb(name)
		},
	}
}

func makeTOnPrefsAvailableEvent(cb TOnPrefsAvailableEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrefsAvailableEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResultOK : boolean);
			sender := lcl.AsObject(getVal(0))
			resultOK := api.GoBool(getVal(1))
			cb(sender, resultOK)
		},
	}
}

func makeTOnPrintDialogEvent(cb TOnPrintDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; hasSelection: boolean; const callback: ICefPrintDialogCallback; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			hasSelection := api.GoBool(getVal(2))
			callback := AsCefPrintDialogCallbackRef(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, hasSelection, callback, result)
		},
	}
}

func makeTOnPrintGetPDFPaperSizeEvent(cb TOnPrintGetPDFPaperSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintGetPDFPaperSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; deviceUnitsPerInch: Integer; var aResult: TCefSize);
			browser := AsCefBrowserRef(getVal(0))
			deviceUnitsPerInch := int32(getVal(1))
			result := (*TCefSize)(getPtr(getVal(2)))
			cb(browser, deviceUnitsPerInch, result)
			*(*TCefSize)(getPtr(getVal(2))) = *result
		},
	}
}

func makeTOnPrintJobEvent(cb TOnPrintJobEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintJobEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const documentName, PDFFilePath: ustring; const callback: ICefPrintJobCallback; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			documentName := api.GoStr(getVal(2))
			pDFFilePath := api.GoStr(getVal(3))
			callback := AsCefPrintJobCallbackRef(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, documentName, pDFFilePath, callback, result)
		},
	}
}

func makeTOnPrintPrintDialogEvent(cb TOnPrintPrintDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintPrintDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; hasSelection: boolean; const callback: ICefPrintDialogCallback; var aResult: boolean);
			browser := AsCefBrowserRef(getVal(0))
			hasSelection := api.GoBool(getVal(1))
			callback := AsCefPrintDialogCallbackRef(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(browser, hasSelection, callback, result)
		},
	}
}

func makeTOnPrintPrintJobEvent(cb TOnPrintPrintJobEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintPrintJobEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const documentName: ustring; const PDFFilePath: ustring; const callback: ICefPrintJobCallback; var aResult: boolean);
			browser := AsCefBrowserRef(getVal(0))
			documentName := api.GoStr(getVal(1))
			pDFFilePath := api.GoStr(getVal(2))
			callback := AsCefPrintJobCallbackRef(getVal(3))
			result := (*bool)(getPtr(getVal(4)))
			cb(browser, documentName, pDFFilePath, callback, result)
		},
	}
}

func makeTOnPrintPrintResetEvent(cb TOnPrintPrintResetEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintPrintResetEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnPrintPrintSettingsEvent(cb TOnPrintPrintSettingsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintPrintSettingsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const settings: ICefPrintSettings; getDefaults: boolean);
			browser := AsCefBrowserRef(getVal(0))
			settings := AsCefPrintSettingsRef(getVal(1))
			getDefaults := api.GoBool(getVal(2))
			cb(browser, settings, getDefaults)
		},
	}
}

func makeTOnPrintPrintStartEvent(cb TOnPrintPrintStartEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintPrintStartEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnPrintResetEvent(cb TOnPrintResetEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintResetEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnPrintSettingsEvent(cb TOnPrintSettingsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintSettingsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const settings: ICefPrintSettings; getDefaults: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			settings := AsCefPrintSettingsRef(getVal(2))
			getDefaults := api.GoBool(getVal(3))
			cb(sender, browser, settings, getDefaults)
		},
	}
}

func makeTOnPrintStartEvent(cb TOnPrintStartEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintStartEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnProcessMessageReceived(cb TOnProcessMessageReceived) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProcessMessageReceived",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message: ICefProcessMessage; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			sourceProcess := cefTypes.TCefProcessId(getVal(3))
			message := AsCefProcessMessageRef(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, frame, sourceProcess, message, result)
		},
	}
}

func makeTOnProcessMessageReceivedEvent(cb TOnProcessMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProcessMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const message: ICefProcessMessage; var aHandled : boolean);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			sourceProcess := cefTypes.TCefProcessId(getVal(2))
			message := AsCefProcessMessageRef(getVal(3))
			handled := (*bool)(getPtr(getVal(4)))
			cb(browser, frame, sourceProcess, message, handled)
		},
	}
}

func makeTOnProtocolExecution(cb TOnProtocolExecution) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProtocolExecution",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var allowOsExecution: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			allowOsExecution := (*bool)(getPtr(getVal(4)))
			cb(sender, browser, frame, request, allowOsExecution)
		},
	}
}

func makeTOnQuickMenuCommandEvent(cb TOnQuickMenuCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnQuickMenuCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; command_id: integer; event_flags: TCefEventFlags; var aResult: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			commandId := int32(getVal(3))
			eventFlags := cefTypes.TCefEventFlags(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, frame, commandId, eventFlags, result)
		},
	}
}

func makeTOnQuickMenuDismissedEvent(cb TOnQuickMenuDismissedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnQuickMenuDismissedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			cb(sender, browser, frame)
		},
	}
}

func makeTOnRegisterCustomPreferencesEvent(cb TOnRegisterCustomPreferencesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRegisterCustomPreferencesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(type_: TCefPreferencesType; const registrar: TCefPreferenceRegistrarRef);
			type_ := cefTypes.TCefPreferencesType(getVal(0))
			registrar := AsCefPreferenceRegistrarRef(getVal(1))
			cb(type_, registrar)
		},
	}
}

func makeTOnRegisterCustomSchemesEvent(cb TOnRegisterCustomSchemesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRegisterCustomSchemesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const registrar: TCefSchemeRegistrarRef);
			registrar := AsCefSchemeRegistrarRef(getVal(0))
			cb(registrar)
		},
	}
}

func makeTOnRenderAcceleratedPaintEvent(cb TOnRenderAcceleratedPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderAcceleratedPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; kind: TCefPaintElementType; dirtyRectsCount: NativeUInt; const dirtyRects: PCefRectArray; const info: PCefAcceleratedPaintInfo);
			browser := AsCefBrowserRef(getVal(0))
			kind := cefTypes.TCefPaintElementType(getVal(1))
			dirtyRectsCount := cefTypes.NativeUInt(getVal(2))
			dirtyRectsPtr := getVal(3)
			dirtyRects := NewCefRectArray(int(dirtyRectsCount), dirtyRectsPtr)
			infoPtr := (*tCefAcceleratedPaintInfo)(getPtr(getVal(4)))
			info := infoPtr.ToGo()
			cb(browser, kind, dirtyRectsCount, dirtyRects, info)
		},
	}
}

func makeTOnRenderGetAccessibilityHandlerEvent(cb TOnRenderGetAccessibilityHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetAccessibilityHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aAccessibilityHandler: ICefAccessibilityHandler);
			var accessibilityHandler IEngAccessibilityHandler
			accessibilityHandler = AsEngAccessibilityHandler(*(*uintptr)(getPtr(getVal(0))))
			cb(&accessibilityHandler)
			if accessibilityHandler != nil {
				*(*uintptr)(getPtr(getVal(0))) = accessibilityHandler.Instance()
			}
		},
	}
}

func makeTOnRenderGetRootScreenRectEvent(cb TOnRenderGetRootScreenRectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetRootScreenRectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; var rect: TCefRect): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			rect := (*TCefRect)(getPtr(getVal(1)))
			ret := cb(browser, rect)
			*(*TCefRect)(getPtr(getVal(1))) = *rect
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnRenderGetScreenInfoEvent(cb TOnRenderGetScreenInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetScreenInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; var screenInfo: TCefScreenInfo): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			screenInfo := (*TCefScreenInfo)(getPtr(getVal(1)))
			ret := cb(browser, screenInfo)
			*(*TCefScreenInfo)(getPtr(getVal(1))) = *screenInfo
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnRenderGetScreenPointEvent(cb TOnRenderGetScreenPointEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetScreenPointEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; viewX: Integer; viewY: Integer; var screenX: Integer; var screenY: Integer): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			viewX := int32(getVal(1))
			viewY := int32(getVal(2))
			screenX := (*int32)(getPtr(getVal(3)))
			screenY := (*int32)(getPtr(getVal(4)))
			ret := cb(browser, viewX, viewY, screenX, screenY)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnRenderGetTouchHandleSizeEvent(cb TOnRenderGetTouchHandleSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetTouchHandleSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; orientation: TCefHorizontalAlignment; var size: TCefSize);
			browser := AsCefBrowserRef(getVal(0))
			orientation := cefTypes.TCefHorizontalAlignment(getVal(1))
			size := (*TCefSize)(getPtr(getVal(2)))
			cb(browser, orientation, size)
			*(*TCefSize)(getPtr(getVal(2))) = *size
		},
	}
}

func makeTOnRenderGetViewRectEvent(cb TOnRenderGetViewRectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderGetViewRectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; var rect: TCefRect);
			browser := AsCefBrowserRef(getVal(0))
			rect := (*TCefRect)(getPtr(getVal(1)))
			cb(browser, rect)
			*(*TCefRect)(getPtr(getVal(1))) = *rect
		},
	}
}

func makeTOnRenderIMECompositionRangeChangedEvent(cb TOnRenderIMECompositionRangeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderIMECompositionRangeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const selected_range: PCefRange; character_boundsCount: NativeUInt; const character_bounds: PCefRect);
			browser := AsCefBrowserRef(getVal(0))
			selectedRange := *(*TCefRange)(getPtr(getVal(1)))
			characterBoundsCount := cefTypes.NativeUInt(getVal(2))
			characterBounds := *(*TCefRect)(getPtr(getVal(3)))
			cb(browser, selectedRange, characterBoundsCount, characterBounds)
		},
	}
}

func makeTOnRenderLoadEnd(cb TOnRenderLoadEnd) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderLoadEnd",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; httpStatusCode: Integer);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			httpStatusCode := int32(getVal(2))
			cb(browser, frame, httpStatusCode)
		},
	}
}

func makeTOnRenderLoadError(cb TOnRenderLoadError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderLoadError",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; errorCode: TCefErrorCode; const errorText, failedUrl: ustring);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			errorCode := cefTypes.TCefErrorCode(getVal(2))
			errorText := api.GoStr(getVal(3))
			failedUrl := api.GoStr(getVal(4))
			cb(browser, frame, errorCode, errorText, failedUrl)
		},
	}
}

func makeTOnRenderLoadStart(cb TOnRenderLoadStart) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderLoadStart",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; transitionType: TCefTransitionType);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			transitionType := cefTypes.TCefTransitionType(getVal(2))
			cb(browser, frame, transitionType)
		},
	}
}

func makeTOnRenderLoadingStateChange(cb TOnRenderLoadingStateChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderLoadingStateChange",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; isLoading, canGoBack, canGoForward: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			isLoading := api.GoBool(getVal(1))
			canGoBack := api.GoBool(getVal(2))
			canGoForward := api.GoBool(getVal(3))
			cb(browser, isLoading, canGoBack, canGoForward)
		},
	}
}

func makeTOnRenderPaintEvent(cb TOnRenderPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderPaintEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnRenderPopupShowEvent(cb TOnRenderPopupShowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderPopupShowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; show: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			show := api.GoBool(getVal(1))
			cb(browser, show)
		},
	}
}

func makeTOnRenderPopupSizeEvent(cb TOnRenderPopupSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderPopupSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const rect: PCefRect);
			browser := AsCefBrowserRef(getVal(0))
			rect := *(*TCefRect)(getPtr(getVal(1)))
			cb(browser, rect)
		},
	}
}

func makeTOnRenderProcessBrowserCreatedEvent(cb TOnRenderProcessBrowserCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessBrowserCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const extra_info: ICefDictionaryValue);
			browser := AsCefBrowserRef(getVal(0))
			extraInfo := AsCefDictionaryValueRef(getVal(1))
			cb(browser, extraInfo)
		},
	}
}

func makeTOnRenderProcessBrowserDestroyedEvent(cb TOnRenderProcessBrowserDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessBrowserDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnRenderProcessContextCreatedEvent(cb TOnRenderProcessContextCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessContextCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			cb(browser, frame, context)
		},
	}
}

func makeTOnRenderProcessContextReleasedEvent(cb TOnRenderProcessContextReleasedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessContextReleasedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			cb(browser, frame, context)
		},
	}
}

func makeTOnRenderProcessFocusedNodeChangedEvent(cb TOnRenderProcessFocusedNodeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessFocusedNodeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const node: ICefDomNode);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			node := AsCefDomNodeRef(getVal(2))
			cb(browser, frame, node)
		},
	}
}

func makeTOnRenderProcessGetLoadHandlerEvent(cb TOnRenderProcessGetLoadHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessGetLoadHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): ICefLoadHandler;
			ret := cb()
			if ret != nil {
				*(*uintptr)(getPtr(getVal(0))) = ret.Instance()
			}
		},
	}
}

func makeTOnRenderProcessProcessMessageReceivedEvent(cb TOnRenderProcessProcessMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessProcessMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; sourceProcess: TCefProcessId; const aMessage: ICefProcessMessage): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			sourceProcess := cefTypes.TCefProcessId(getVal(2))
			message := AsCefProcessMessageRef(getVal(3))
			ret := cb(browser, frame, sourceProcess, message)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnRenderProcessResponsive(cb TOnRenderProcessResponsive) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessResponsive",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: Tobject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnRenderProcessTerminated(cb TOnRenderProcessTerminated) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessTerminated",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const browser: ICefBrowser; status: TCefTerminationStatus; error_code: integer; const error_string: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			status := cefTypes.TCefTerminationStatus(getVal(2))
			errorCode := int32(getVal(3))
			errorString := api.GoStr(getVal(4))
			cb(sender, browser, status, errorCode, errorString)
		},
	}
}

func makeTOnRenderProcessUncaughtExceptionEvent(cb TOnRenderProcessUncaughtExceptionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessUncaughtExceptionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context; const V8Exception: ICefV8Exception; const stackTrace: ICefV8StackTrace);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			v8Exception := AsCefV8ExceptionRef(getVal(3))
			stackTrace := AsCefV8StackTraceRef(getVal(4))
			cb(browser, frame, context, v8Exception, stackTrace)
		},
	}
}

func makeTOnRenderProcessUnresponsive(cb TOnRenderProcessUnresponsive) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessUnresponsive",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: Tobject; const browser: ICefBrowser; const callback: ICefUnresponsiveProcessCallback; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			callback := AsCefUnresponsiveProcessCallbackRef(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, callback, result)
		},
	}
}

func makeTOnRenderProcessWebKitInitializedEvent(cb TOnRenderProcessWebKitInitializedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderProcessWebKitInitializedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnRenderScrollOffsetChangedEvent(cb TOnRenderScrollOffsetChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderScrollOffsetChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; x: Double; y: Double);
			browser := AsCefBrowserRef(getVal(0))
			X := *(*float64)(getPtr(getVal(1)))
			Y := *(*float64)(getPtr(getVal(2)))
			cb(browser, X, Y)
		},
	}
}

func makeTOnRenderStartDraggingEvent(cb TOnRenderStartDraggingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderStartDraggingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const dragData: ICefDragData; allowedOps: TCefDragOperations; x: Integer; y: Integer): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			dragData := AsCefDragDataRef(getVal(1))
			allowedOps := cefTypes.TCefDragOperations(getVal(2))
			X := int32(getVal(3))
			Y := int32(getVal(4))
			ret := cb(browser, dragData, allowedOps, X, Y)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnRenderTextSelectionChangedEvent(cb TOnRenderTextSelectionChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderTextSelectionChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const browser: ICefBrowser; const selected_text: ustring; const selected_range: PCefRange);
			browser := AsCefBrowserRef(getVal(0))
			selectedText := api.GoStr(getVal(1))
			selectedRange := *(*TCefRange)(getPtr(getVal(2)))
			cb(browser, selectedText, selectedRange)
		},
	}
}

func makeTOnRenderTouchHandleStateChangedEvent(cb TOnRenderTouchHandleStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderTouchHandleStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; const state: TCefTouchHandleState);
			browser := AsCefBrowserRef(getVal(0))
			state := *(*TCefTouchHandleState)(getPtr(getVal(1)))
			cb(browser, state)
		},
	}
}

func makeTOnRenderUpdateDragCursorEvent(cb TOnRenderUpdateDragCursorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderUpdateDragCursorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; operation: TCefDragOperation);
			browser := AsCefBrowserRef(getVal(0))
			operation := cefTypes.TCefDragOperation(getVal(1))
			cb(browser, operation)
		},
	}
}

func makeTOnRenderViewReady(cb TOnRenderViewReady) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderViewReady",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: Tobject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnRenderVirtualKeyboardRequestedEvent(cb TOnRenderVirtualKeyboardRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRenderVirtualKeyboardRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const browser: ICefBrowser; input_mode: TCefTextInpuMode);
			browser := AsCefBrowserRef(getVal(0))
			inputMode := cefTypes.TCefTextInpuMode(getVal(1))
			cb(browser, inputMode)
		},
	}
}

func makeTOnRequestBeforeBrowseEvent(cb TOnRequestBeforeBrowseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestBeforeBrowseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; user_gesture: Boolean; isRedirect: Boolean): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			userGesture := api.GoBool(getVal(3))
			isRedirect := api.GoBool(getVal(4))
			ret := cb(browser, frame, request, userGesture, isRedirect)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnRequestCertificateErrorEvent(cb TOnRequestCertificateErrorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestCertificateErrorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; certError: TCefErrorcode; const requestUrl: ustring; const sslInfo: ICefSslInfo; const callback: ICefCallback): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			certError := cefTypes.TCefErrorCode(getVal(1))
			requestUrl := api.GoStr(getVal(2))
			sslInfo := AsCefSslInfoRef(getVal(3))
			callback := AsCefCallbackRef(getVal(4))
			ret := cb(browser, certError, requestUrl, sslInfo, callback)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnRequestComplete(cb TOnRequestComplete) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestComplete",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const request: ICefUrlRequest);
			sender := lcl.AsObject(getVal(0))
			request := AsCefUrlRequestRef(getVal(1))
			cb(sender, request)
		},
	}
}

func makeTOnRequestContextGetResourceRequestHandlerEvent(cb TOnRequestContextGetResourceRequestHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestContextGetResourceRequestHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnRequestContextInitialized(cb TOnRequestContextInitialized) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestContextInitialized",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const request_context: ICefRequestContext);
			sender := lcl.AsObject(getVal(0))
			requestContext := AsCefRequestContextRef(getVal(1))
			cb(sender, requestContext)
		},
	}
}

func makeTOnRequestContextRequestContextInitializedEvent(cb TOnRequestContextRequestContextInitializedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestContextRequestContextInitializedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const request_context: ICefRequestContext);
			requestContext := AsCefRequestContextRef(getVal(0))
			cb(requestContext)
		},
	}
}

func makeTOnRequestDocumentAvailableInMainFrameEvent(cb TOnRequestDocumentAvailableInMainFrameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestDocumentAvailableInMainFrameEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnRequestGetAuthCredentialsEvent(cb TOnRequestGetAuthCredentialsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestGetAuthCredentialsEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnRequestGetResourceRequestHandlerEvent(cb TOnRequestGetResourceRequestHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestGetResourceRequestHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnRequestMediaAccessPermissionEvent(cb TOnRequestMediaAccessPermissionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestMediaAccessPermissionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefMediaAccessCallback; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			requestingOrigin := api.GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := AsCefMediaAccessCallbackRef(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, frame, requestingOrigin, requestedPermissions, callback, result)
		},
	}
}

func makeTOnRequestOpenUrlFromTabEvent(cb TOnRequestOpenUrlFromTabEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestOpenUrlFromTabEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : function(const browser: ICefBrowser; const frame: ICefFrame; const targetUrl: ustring; targetDisposition: TCefWindowOpenDisposition; userGesture: Boolean): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			targetUrl := api.GoStr(getVal(2))
			targetDisposition := cefTypes.TCefWindowOpenDisposition(getVal(3))
			userGesture := api.GoBool(getVal(4))
			ret := cb(browser, frame, targetUrl, targetDisposition, userGesture)
			*(*bool)(getPtr(getVal(5))) = ret
		},
	}
}

func makeTOnRequestRenderProcessResponsiveEvent(cb TOnRequestRenderProcessResponsiveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestRenderProcessResponsiveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnRequestRenderProcessTerminatedEvent(cb TOnRequestRenderProcessTerminatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestRenderProcessTerminatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; status: TCefTerminationStatus; error_code: integer; const error_string: ustring);
			browser := AsCefBrowserRef(getVal(0))
			status := cefTypes.TCefTerminationStatus(getVal(1))
			errorCode := int32(getVal(2))
			errorString := api.GoStr(getVal(3))
			cb(browser, status, errorCode, errorString)
		},
	}
}

func makeTOnRequestRenderProcessUnresponsiveEvent(cb TOnRequestRenderProcessUnresponsiveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestRenderProcessUnresponsiveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const browser: ICefBrowser; const callback: ICefUnresponsiveProcessCallback): boolean;
			browser := AsCefBrowserRef(getVal(0))
			callback := AsCefUnresponsiveProcessCallbackRef(getVal(1))
			ret := cb(browser, callback)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnRequestRenderViewReadyEvent(cb TOnRequestRenderViewReadyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestRenderViewReadyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const browser: ICefBrowser);
			browser := AsCefBrowserRef(getVal(0))
			cb(browser)
		},
	}
}

func makeTOnRequestSelectClientCertificateEvent(cb TOnRequestSelectClientCertificateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRequestSelectClientCertificateEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnResetDialogState(cb TOnResetDialogState) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResetDialogState",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const browser: ICefBrowser);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			cb(sender, browser)
		},
	}
}

func makeTOnResolveCallbackResolveCompletedEvent(cb TOnResolveCallbackResolveCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResolveCallbackResolveCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(result: TCefErrorCode; const resolvedIps: TStrings);
			result := cefTypes.TCefErrorCode(getVal(0))
			resolvedIps := lcl.AsStrings(getVal(1))
			cb(result, resolvedIps)
		},
	}
}

func makeTOnResolvedIPsAvailableEvent(cb TOnResolvedIPsAvailableEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResolvedIPsAvailableEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; result: TCefErrorCode; const resolvedIps: TStrings);
			sender := lcl.AsObject(getVal(0))
			result := cefTypes.TCefErrorCode(getVal(1))
			resolvedIps := lcl.AsStrings(getVal(2))
			cb(sender, result, resolvedIps)
		},
	}
}

func makeTOnResourceBundleGetDataResourceEvent(cb TOnResourceBundleGetDataResourceEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceBundleGetDataResourceEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(resourceId: Integer; var data: Pointer; var dataSize: NativeUInt): Boolean;
			resourceId := int32(getVal(0))
			data := (*uintptr)(getPtr(getVal(1)))
			dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
			ret := cb(resourceId, data, dataSize)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnResourceBundleGetDataResourceForScaleEvent(cb TOnResourceBundleGetDataResourceForScaleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceBundleGetDataResourceForScaleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(resourceId: Integer; scaleFactor: TCefScaleFactor; var data: Pointer; var dataSize: NativeUInt): Boolean;
			resourceId := int32(getVal(0))
			scaleFactor := cefTypes.TCefScaleFactor(getVal(1))
			data := (*uintptr)(getPtr(getVal(2)))
			dataSize := (*cefTypes.NativeUInt)(getPtr(getVal(3)))
			ret := cb(resourceId, scaleFactor, data, dataSize)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnResourceBundleGetLocalizedStringEvent(cb TOnResourceBundleGetLocalizedStringEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceBundleGetLocalizedStringEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(stringid: Integer; var stringVal: ustring): Boolean;
			stringid := int32(getVal(0))
			stringVal := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			ret := cb(stringid, &stringVal)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(stringVal)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnResourceCancelEvent(cb TOnResourceCancelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceCancelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnResourceGetResponseHeadersEvent(cb TOnResourceGetResponseHeadersEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceGetResponseHeadersEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const response: ICefResponse; out responseLength: Int64; out redirectUrl: ustring);
			response := AsCefResponseRef(getVal(0))
			responseLength := (*int64)(getPtr(getVal(1)))
			var redirectUrl string
			cb(response, responseLength, &redirectUrl)
			*(*uintptr)(getPtr(getVal(2))) = api.PasStr(redirectUrl)
		},
	}
}

func makeTOnResourceLoadComplete(cb TOnResourceLoadComplete) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceLoadComplete",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; status: TCefUrlRequestStatus; receivedContentLength: Int64);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			response := AsCefResponseRef(getVal(4))
			status := cefTypes.TCefUrlRequestStatus(getVal(5))
			receivedContentLength := *(*int64)(getPtr(getVal(6)))
			cb(sender, browser, frame, request, response, status, receivedContentLength)
		},
	}
}

func makeTOnResourceOpenEvent(cb TOnResourceOpenEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceOpenEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const request: ICefRequest; var handle_request: boolean; const callback: ICefCallback): boolean;
			request := AsCefRequestRef(getVal(0))
			handleRequest := (*bool)(getPtr(getVal(1)))
			callback := AsCefCallbackRef(getVal(2))
			ret := cb(request, handleRequest, callback)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnResourceProcessRequestEvent(cb TOnResourceProcessRequestEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceProcessRequestEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const request: ICefRequest; const callback: ICefCallback): Boolean;
			request := AsCefRequestRef(getVal(0))
			callback := AsCefCallbackRef(getVal(1))
			ret := cb(request, callback)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnResourceReadEvent(cb TOnResourceReadEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceReadEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const data_out: Pointer; bytes_to_read: Integer; var bytes_read: Integer; const callback: ICefResourceReadCallback): boolean;
			dataOut := uintptr(getVal(0))
			bytesToRead := int32(getVal(1))
			bytesRead := (*int32)(getPtr(getVal(2)))
			callback := AsCefResourceReadCallbackRef(getVal(3))
			ret := cb(dataOut, bytesToRead, bytesRead, callback)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnResourceReadResponseEvent(cb TOnResourceReadResponseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceReadResponseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const dataOut: Pointer; bytesToRead: Integer; var bytesRead: Integer; const callback: ICefCallback): Boolean;
			dataOut := uintptr(getVal(0))
			bytesToRead := int32(getVal(1))
			bytesRead := (*int32)(getPtr(getVal(2)))
			callback := AsCefCallbackRef(getVal(3))
			ret := cb(dataOut, bytesToRead, bytesRead, callback)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnResourceRedirect(cb TOnResourceRedirect) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRedirect",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; var newUrl: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			response := AsCefResponseRef(getVal(4))
			newUrl := api.GoStr(*(*uintptr)(getPtr(getVal(5))))
			cb(sender, browser, frame, request, response, &newUrl)
			*(*uintptr)(getPtr(getVal(5))) = api.PasStr(newUrl)
		},
	}
}

func makeTOnResourceRequestBeforeResourceLoadEvent(cb TOnResourceRequestBeforeResourceLoadEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestBeforeResourceLoadEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const callback: ICefCallback): TCefReturnValue;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			callback := AsCefCallbackRef(getVal(3))
			ret := cb(browser, frame, request, callback)
			*(*cefTypes.TCefReturnValue)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnResourceRequestGetCookieAccessFilterEvent(cb TOnResourceRequestGetCookieAccessFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestGetCookieAccessFilterEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnResourceRequestGetResourceHandlerEvent(cb TOnResourceRequestGetResourceHandlerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestGetResourceHandlerEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnResourceRequestGetResourceResponseFilterEvent(cb TOnResourceRequestGetResourceResponseFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestGetResourceResponseFilterEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnResourceRequestProtocolExecutionEvent(cb TOnResourceRequestProtocolExecutionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestProtocolExecutionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; var allowOsExecution: Boolean);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			allowOsExecution := (*bool)(getPtr(getVal(3)))
			cb(browser, frame, request, allowOsExecution)
		},
	}
}

func makeTOnResourceRequestResourceLoadCompleteEvent(cb TOnResourceRequestResourceLoadCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestResourceLoadCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; status: TCefUrlRequestStatus; receivedContentLength: Int64);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			response := AsCefResponseRef(getVal(3))
			status := cefTypes.TCefUrlRequestStatus(getVal(4))
			receivedContentLength := *(*int64)(getPtr(getVal(5)))
			cb(browser, frame, request, response, status, receivedContentLength)
		},
	}
}

func makeTOnResourceRequestResourceRedirectEvent(cb TOnResourceRequestResourceRedirectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestResourceRedirectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; var newUrl: ustring);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			response := AsCefResponseRef(getVal(3))
			newUrl := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(browser, frame, request, response, &newUrl)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(newUrl)
		},
	}
}

func makeTOnResourceRequestResourceResponseEvent(cb TOnResourceRequestResourceResponseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceRequestResourceResponseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse): Boolean;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			request := AsCefRequestRef(getVal(2))
			response := AsCefResponseRef(getVal(3))
			ret := cb(browser, frame, request, response)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnResourceResponse(cb TOnResourceResponse) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceResponse",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const request: ICefRequest; const response: ICefResponse; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			request := AsCefRequestRef(getVal(3))
			response := AsCefResponseRef(getVal(4))
			result := (*bool)(getPtr(getVal(5)))
			cb(sender, browser, frame, request, response, result)
		},
	}
}

func makeTOnResourceSkipEvent(cb TOnResourceSkipEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResourceSkipEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(bytes_to_skip: int64; var bytes_skipped: Int64; const callback: ICefResourceSkipCallback): boolean;
			bytesToSkip := *(*int64)(getPtr(getVal(0)))
			bytesSkipped := (*int64)(getPtr(getVal(1)))
			callback := AsCefResourceSkipCallbackRef(getVal(2))
			ret := cb(bytesToSkip, bytesSkipped, callback)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnResponseFilterFilterEvent(cb TOnResponseFilterFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResponseFilterFilterEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : function(data_in: Pointer; data_in_size: NativeUInt; var data_in_read: NativeUInt; data_out: Pointer; data_out_size: NativeUInt; var data_out_written: NativeUInt): TCefResponseFilterStatus;
			dataIn := uintptr(getVal(0))
			dataInSize := cefTypes.NativeUInt(getVal(1))
			dataInRead := (*cefTypes.NativeUInt)(getPtr(getVal(2)))
			dataOut := uintptr(getVal(3))
			dataOutSize := cefTypes.NativeUInt(getVal(4))
			dataOutWritten := (*cefTypes.NativeUInt)(getPtr(getVal(5)))
			ret := cb(dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten)
			*(*cefTypes.TCefResponseFilterStatus)(getPtr(getVal(6))) = ret
		},
	}
}

func makeTOnResponseFilterInitFilterEvent(cb TOnResponseFilterInitFilterEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnResponseFilterInitFilterEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): Boolean;
			ret := cb()
			*(*bool)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTOnRouteMessageReceivedEvent(cb TOnRouteMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRouteMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const route: ICefMediaRoute; const message_: ustring);
			sender := lcl.AsObject(getVal(0))
			route := AsCefMediaRouteRef(getVal(1))
			message := api.GoStr(getVal(2))
			cb(sender, route, message)
		},
	}
}

func makeTOnRouteStateChangedEvent(cb TOnRouteStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRouteStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const route: ICefMediaRoute; state: TCefMediaRouteConnectionState);
			sender := lcl.AsObject(getVal(0))
			route := AsCefMediaRouteRef(getVal(1))
			state := cefTypes.TCefMediaRouteConnectionState(getVal(2))
			cb(sender, route, state)
		},
	}
}

func makeTOnRoutesEvent(cb TOnRoutesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRoutesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const routes: TCefMediaRouteArray);
			sender := lcl.AsObject(getVal(0))
			routesPtr := getVal(1)
			routesLen := int32(getVal(2))
			routes := NewCefMediaRouteArray(int(routesLen), routesPtr)
			cb(sender, routes)
		},
	}
}

func makeTOnRunContextMenu(cb TOnRunContextMenu) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRunContextMenu",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const frame: ICefFrame; const params: ICefContextMenuParams; const model: ICefMenuModel; const callback: ICefRunContextMenuCallback; var aResult : Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			frame := AsCefFrameRef(getVal(2))
			params := AsCefContextMenuParamsRef(getVal(3))
			model := AsCefMenuModelRef(getVal(4))
			callback := AsCefRunContextMenuCallbackRef(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, frame, params, model, callback, result)
		},
	}
}

func makeTOnRunFileDialogCallbackFileDialogDismissedEvent(cb TOnRunFileDialogCallbackFileDialogDismissedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRunFileDialogCallbackFileDialogDismissedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const filePaths: TStrings);
			filePaths := lcl.AsStrings(getVal(0))
			cb(filePaths)
		},
	}
}

func makeTOnRunQuickMenuEvent(cb TOnRunQuickMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRunQuickMenuEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnScheduleMessagePumpWorkEvent(cb TOnScheduleMessagePumpWorkEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnScheduleMessagePumpWorkEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const delayMs: Int64);
			delayMs := *(*int64)(getPtr(getVal(0)))
			cb(delayMs)
		},
	}
}

func makeTOnSchemeFactoryNewEvent(cb TOnSchemeFactoryNewEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSchemeFactoryNewEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const browser: ICefBrowser; const frame: ICefFrame; const schemeName: ustring; const request: ICefRequest): ICefResourceHandler;
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			schemeName := api.GoStr(getVal(2))
			request := AsCefRequestRef(getVal(3))
			ret := cb(browser, frame, schemeName, request)
			if ret != nil {
				*(*uintptr)(getPtr(getVal(4))) = ret.Instance()
			}
		},
	}
}

func makeTOnScrollOffsetChanged(cb TOnScrollOffsetChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnScrollOffsetChanged",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; x, y: Double);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			X := *(*float64)(getPtr(getVal(2)))
			Y := *(*float64)(getPtr(getVal(3)))
			cb(sender, browser, X, Y)
		},
	}
}

func makeTOnSelectClientCertificate(cb TOnSelectClientCertificate) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSelectClientCertificate",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnServerClientConnectedEvent(cb TOnServerClientConnectedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerClientConnectedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const server: ICefServer; connection_id: Integer);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			cb(server, connectionId)
		},
	}
}

func makeTOnServerClientDisconnectedEvent(cb TOnServerClientDisconnectedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerClientDisconnectedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const server: ICefServer; connection_id: Integer);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			cb(server, connectionId)
		},
	}
}

func makeTOnServerCreated(cb TOnServerCreated) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerCreated",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const server: ICefServer);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			cb(sender, server)
		},
	}
}

func makeTOnServerDestroyed(cb TOnServerDestroyed) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerDestroyed",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const server: ICefServer);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			cb(sender, server)
		},
	}
}

func makeTOnServerHttpRequestEvent(cb TOnServerHttpRequestEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerHttpRequestEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			clientAddress := api.GoStr(getVal(2))
			request := AsCefRequestRef(getVal(3))
			cb(server, connectionId, clientAddress, request)
		},
	}
}

func makeTOnServerServerCreatedEvent(cb TOnServerServerCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerServerCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const server: ICefServer);
			server := AsCEFServerRef(getVal(0))
			cb(server)
		},
	}
}

func makeTOnServerServerDestroyedEvent(cb TOnServerServerDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerServerDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const server: ICefServer);
			server := AsCEFServerRef(getVal(0))
			cb(server)
		},
	}
}

func makeTOnServerWebSocketConnectedEvent(cb TOnServerWebSocketConnectedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerWebSocketConnectedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const server: ICefServer; connection_id: Integer);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			cb(server, connectionId)
		},
	}
}

func makeTOnServerWebSocketMessageEvent(cb TOnServerWebSocketMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerWebSocketMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const server: ICefServer; connection_id: Integer; const data: Pointer; data_size: NativeUInt);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			data := uintptr(getVal(2))
			dataSize := cefTypes.NativeUInt(getVal(3))
			cb(server, connectionId, data, dataSize)
		},
	}
}

func makeTOnServerWebSocketRequestEvent(cb TOnServerWebSocketRequestEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerWebSocketRequestEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest; const callback: ICefCallback);
			server := AsCEFServerRef(getVal(0))
			connectionId := int32(getVal(1))
			clientAddress := api.GoStr(getVal(2))
			request := AsCefRequestRef(getVal(3))
			callback := AsCefCallbackRef(getVal(4))
			cb(server, connectionId, clientAddress, request, callback)
		},
	}
}

func makeTOnSetCookieCallbackCompleteEvent(cb TOnSetCookieCallbackCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSetCookieCallbackCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(success: Boolean);
			success := api.GoBool(getVal(0))
			cb(success)
		},
	}
}

func makeTOnSetFocus(cb TOnSetFocus) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSetFocus",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; source: TCefFocusSource; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			source := cefTypes.TCefFocusSource(getVal(2))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, source, result)
		},
	}
}

func makeTOnSettingChangedEvent(cb TOnSettingChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSettingChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const requesting_url, top_level_url : ustring; content_type: TCefContentSettingTypes);
			sender := lcl.AsObject(getVal(0))
			requestingUrl := api.GoStr(getVal(1))
			topLevelUrl := api.GoStr(getVal(2))
			contentType := cefTypes.TCefContentSettingTypes(getVal(3))
			cb(sender, requestingUrl, topLevelUrl, contentType)
		},
	}
}

func makeTOnSettingObserverSettingChangedEvent(cb TOnSettingObserverSettingChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSettingObserverSettingChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const requesting_url: ustring; const top_level_url: ustring; content_type: TCefContentSettingTypes);
			requestingUrl := api.GoStr(getVal(0))
			topLevelUrl := api.GoStr(getVal(1))
			contentType := cefTypes.TCefContentSettingTypes(getVal(2))
			cb(requestingUrl, topLevelUrl, contentType)
		},
	}
}

func makeTOnShowPermissionPromptEvent(cb TOnShowPermissionPromptEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnShowPermissionPromptEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; prompt_id: uint64; const requesting_origin: ustring; requested_permissions: cardinal; const callback: ICefPermissionPromptCallback; var aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			promptId := *(*uint64)(getPtr(getVal(2)))
			requestingOrigin := api.GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := AsCefPermissionPromptCallbackRef(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, promptId, requestingOrigin, requestedPermissions, callback, result)
		},
	}
}

func makeTOnSinksEvent(cb TOnSinksEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSinksEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const sinks: TCefMediaSinkArray);
			sender := lcl.AsObject(getVal(0))
			sinksPtr := getVal(1)
			sinksLen := int32(getVal(2))
			sinks := NewCefMediaSinkArray(int(sinksLen), sinksPtr)
			cb(sender, sinks)
		},
	}
}

func makeTOnStartDragging(cb TOnStartDragging) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnStartDragging",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TObject; const browser: ICefBrowser; const dragData: ICefDragData; allowedOps: TCefDragOperations; x, y: Integer; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			dragData := AsCefDragDataRef(getVal(2))
			allowedOps := cefTypes.TCefDragOperations(getVal(3))
			X := int32(getVal(4))
			Y := int32(getVal(5))
			result := (*bool)(getPtr(getVal(6)))
			cb(sender, browser, dragData, allowedOps, X, Y, result)
		},
	}
}

func makeTOnStatusMessage(cb TOnStatusMessage) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnStatusMessage",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const value: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			value := api.GoStr(getVal(2))
			cb(sender, browser, value)
		},
	}
}

func makeTOnStringVisitorVisitEvent(cb TOnStringVisitorVisitEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnStringVisitorVisitEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const str: ustring);
			str := api.GoStr(getVal(0))
			cb(str)
		},
	}
}

func makeTOnTakeFocus(cb TOnTakeFocus) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTakeFocus",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; next_: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			next := api.GoBool(getVal(2))
			cb(sender, browser, next)
		},
	}
}

func makeTOnTaskExecuteEvent(cb TOnTaskExecuteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTaskExecuteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnTextResultAvailableEvent(cb TOnTextResultAvailableEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTextResultAvailableEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aText : ustring);
			sender := lcl.AsObject(getVal(0))
			text := api.GoStr(getVal(1))
			cb(sender, text)
		},
	}
}

func makeTOnTextSelectionChanged(cb TOnTextSelectionChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTextSelectionChanged",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; const selected_text: ustring; const selected_range: PCefRange);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			selectedText := api.GoStr(getVal(2))
			selectedRange := *(*TCefRange)(getPtr(getVal(3)))
			cb(sender, browser, selectedText, selectedRange)
		},
	}
}

func makeTOnTextfieldAfterUserActionEvent(cb TOnTextfieldAfterUserActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTextfieldAfterUserActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const textfield: ICefTextfield);
			textfield := AsCefTextfieldRef(getVal(0))
			cb(textfield)
		},
	}
}

func makeTOnTextfieldKeyEvent(cb TOnTextfieldKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTextfieldKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const textfield: ICefTextfield; const event: TCefKeyEvent; var aResult: boolean);
			textfield := AsCefTextfieldRef(getVal(0))
			event := *(*TCefKeyEvent)(getPtr(getVal(1)))
			result := (*bool)(getPtr(getVal(2)))
			cb(textfield, event, result)
		},
	}
}

func makeTOnTextfieldKeyEventEvent(cb TOnTextfieldKeyEventEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTextfieldKeyEventEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const textfield: ICefTextfield; const event: TCefKeyEvent; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			textfield := AsCefTextfieldRef(getVal(1))
			event := *(*TCefKeyEvent)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, textfield, event, result)
		},
	}
}

func makeTOnThemeChangedEvent(cb TOnThemeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnThemeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const view: ICefView);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			cb(sender, view)
		},
	}
}

func makeTOnThemeColorsChangedEvent(cb TOnThemeColorsChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnThemeColorsChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; chrome_theme: Integer);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			chromeTheme := int32(getVal(2))
			cb(sender, window, chromeTheme)
		},
	}
}

func makeTOnTitleChange(cb TOnTitleChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTitleChange",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const title: ustring);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			title := api.GoStr(getVal(2))
			cb(sender, browser, title)
		},
	}
}

func makeTOnTooltip(cb TOnTooltip) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTooltip",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const browser: ICefBrowser; var text: ustring; out Result: Boolean);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			text := api.GoStr(*(*uintptr)(getPtr(getVal(2))))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, browser, &text, result)
			*(*uintptr)(getPtr(getVal(2))) = api.PasStr(text)
		},
	}
}

func makeTOnTouchHandleStateChanged(cb TOnTouchHandleStateChanged) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTouchHandleStateChanged",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; const state: TCefTouchHandleState);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			state := *(*TCefTouchHandleState)(getPtr(getVal(2)))
			cb(sender, browser, state)
		},
	}
}

func makeTOnUncaughtExceptionEvent(cb TOnUncaughtExceptionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUncaughtExceptionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(const browser: ICefBrowser; const frame: ICefFrame; const context: ICefv8Context; const exception: ICefV8Exception; const stackTrace: ICefV8StackTrace);
			browser := AsCefBrowserRef(getVal(0))
			frame := AsCefFrameRef(getVal(1))
			context := AsCefv8ContextRef(getVal(2))
			exception := AsCefV8ExceptionRef(getVal(3))
			stackTrace := AsCefV8StackTraceRef(getVal(4))
			cb(browser, frame, context, exception, stackTrace)
		},
	}
}

func makeTOnUpdateDragCursor(cb TOnUpdateDragCursor) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUpdateDragCursor",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; operation: TCefDragOperation);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			operation := cefTypes.TCefDragOperation(getVal(2))
			cb(sender, browser, operation)
		},
	}
}

func makeTOnUploadProgress(cb TOnUploadProgress) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUploadProgress",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const request: ICefUrlRequest; current, total: Int64);
			sender := lcl.AsObject(getVal(0))
			request := AsCefUrlRequestRef(getVal(1))
			current := *(*int64)(getPtr(getVal(2)))
			total := *(*int64)(getPtr(getVal(3)))
			cb(sender, request, current, total)
		},
	}
}

func makeTOnUrlrequestClientDownloadDataEvent(cb TOnUrlrequestClientDownloadDataEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUrlrequestClientDownloadDataEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const request: ICefUrlRequest; data: Pointer; dataLength: NativeUInt);
			request := AsCefUrlRequestRef(getVal(0))
			data := uintptr(getVal(1))
			dataLength := cefTypes.NativeUInt(getVal(2))
			cb(request, data, dataLength)
		},
	}
}

func makeTOnUrlrequestClientDownloadProgressEvent(cb TOnUrlrequestClientDownloadProgressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUrlrequestClientDownloadProgressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const request: ICefUrlRequest; current: Int64; total: Int64);
			request := AsCefUrlRequestRef(getVal(0))
			current := *(*int64)(getPtr(getVal(1)))
			total := *(*int64)(getPtr(getVal(2)))
			cb(request, current, total)
		},
	}
}

func makeTOnUrlrequestClientGetAuthCredentialsEvent(cb TOnUrlrequestClientGetAuthCredentialsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUrlrequestClientGetAuthCredentialsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : function(isProxy: Boolean; const host: ustring; port: Integer; const realm: ustring; const scheme: ustring; const callback: ICefAuthCallback): Boolean;
			isProxy := api.GoBool(getVal(0))
			host := api.GoStr(getVal(1))
			port := int32(getVal(2))
			realm := api.GoStr(getVal(3))
			scheme := api.GoStr(getVal(4))
			callback := AsCefAuthCallbackRef(getVal(5))
			ret := cb(isProxy, host, port, realm, scheme, callback)
			*(*bool)(getPtr(getVal(6))) = ret
		},
	}
}

func makeTOnUrlrequestClientRequestCompleteEvent(cb TOnUrlrequestClientRequestCompleteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUrlrequestClientRequestCompleteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const request: ICefUrlRequest);
			request := AsCefUrlRequestRef(getVal(0))
			cb(request)
		},
	}
}

func makeTOnUrlrequestClientUploadProgressEvent(cb TOnUrlrequestClientUploadProgressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUrlrequestClientUploadProgressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const request: ICefUrlRequest; current: Int64; total: Int64);
			request := AsCefUrlRequestRef(getVal(0))
			current := *(*int64)(getPtr(getVal(1)))
			total := *(*int64)(getPtr(getVal(2)))
			cb(request, current, total)
		},
	}
}

func makeTOnUseFramelessWindowForPictureInPicture(cb TOnUseFramelessWindowForPictureInPicture) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUseFramelessWindowForPictureInPicture",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const browser_view: ICefBrowserView; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			browserView := AsCefBrowserViewRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, browserView, result)
		},
	}
}

func makeTOnV8AccessorGetEvent(cb TOnV8AccessorGetEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8AccessorGetEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnV8AccessorSet_Event(cb TOnV8AccessorSet_Event) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8AccessorSet_Event",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const name: ustring; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): Boolean;
			name := api.GoStr(getVal(0))
			object := AsCefv8ValueRef(getVal(1))
			value := AsCefv8ValueRef(getVal(2))
			exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
			ret := cb(name, object, value, &exception)
			*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnV8ArrayBufferReleaseCallbackReleaseBufferEvent(cb TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8ArrayBufferReleaseCallbackReleaseBufferEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(buffer: Pointer);
			buffer := uintptr(getVal(0))
			cb(buffer)
		},
	}
}

func makeTOnV8ExecuteEvent(cb TOnV8ExecuteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8ExecuteEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnV8InterceptorGetByIndexEvent(cb TOnV8InterceptorGetByIndexEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8InterceptorGetByIndexEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnV8InterceptorGetByNameEvent(cb TOnV8InterceptorGetByNameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8InterceptorGetByNameEvent",
		cb: func(getVal func(i int) uintptr) {
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
		},
	}
}

func makeTOnV8InterceptorSetByIndexEvent(cb TOnV8InterceptorSetByIndexEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8InterceptorSetByIndexEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(index: integer; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): boolean;
			index := int32(getVal(0))
			object := AsCefv8ValueRef(getVal(1))
			value := AsCefv8ValueRef(getVal(2))
			exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
			ret := cb(index, object, value, &exception)
			*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnV8InterceptorSetByNameEvent(cb TOnV8InterceptorSetByNameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnV8InterceptorSetByNameEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(const name: ustring; const object_: ICefv8Value; const value: ICefv8Value; var exception: ustring): boolean;
			name := api.GoStr(getVal(0))
			object := AsCefv8ValueRef(getVal(1))
			value := AsCefv8ValueRef(getVal(2))
			exception := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
			ret := cb(name, object, value, &exception)
			*(*uintptr)(getPtr(getVal(3))) = api.PasStr(exception)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTOnViewBlurEvent(cb TOnViewBlurEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewBlurEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const view: ICefView);
			view := AsCefViewRef(getVal(0))
			cb(view)
		},
	}
}

func makeTOnViewChildViewChangedEvent(cb TOnViewChildViewChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewChildViewChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const view: ICefView; added: boolean; const child: ICefView);
			view := AsCefViewRef(getVal(0))
			added := api.GoBool(getVal(1))
			child := AsCefViewRef(getVal(2))
			cb(view, added, child)
		},
	}
}

func makeTOnViewFocusEvent(cb TOnViewFocusEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewFocusEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const view: ICefView);
			view := AsCefViewRef(getVal(0))
			cb(view)
		},
	}
}

func makeTOnViewGetHeightForWidthEvent(cb TOnViewGetHeightForWidthEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewGetHeightForWidthEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const view: ICefView; width: Integer; var aResult: Integer);
			view := AsCefViewRef(getVal(0))
			width := int32(getVal(1))
			result := (*int32)(getPtr(getVal(2)))
			cb(view, width, result)
		},
	}
}

func makeTOnViewGetMaximumSizeEvent(cb TOnViewGetMaximumSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewGetMaximumSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const view: ICefView; var aResult: TCefSize);
			view := AsCefViewRef(getVal(0))
			result := (*TCefSize)(getPtr(getVal(1)))
			cb(view, result)
			*(*TCefSize)(getPtr(getVal(1))) = *result
		},
	}
}

func makeTOnViewGetMinimumSizeEvent(cb TOnViewGetMinimumSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewGetMinimumSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const view: ICefView; var aResult: TCefSize);
			view := AsCefViewRef(getVal(0))
			result := (*TCefSize)(getPtr(getVal(1)))
			cb(view, result)
			*(*TCefSize)(getPtr(getVal(1))) = *result
		},
	}
}

func makeTOnViewGetPreferredSizeEvent(cb TOnViewGetPreferredSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewGetPreferredSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const view: ICefView; var aResult: TCefSize);
			view := AsCefViewRef(getVal(0))
			result := (*TCefSize)(getPtr(getVal(1)))
			cb(view, result)
			*(*TCefSize)(getPtr(getVal(1))) = *result
		},
	}
}

func makeTOnViewLayoutChangedEvent(cb TOnViewLayoutChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewLayoutChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const view: ICefView; new_bounds: TCefRect);
			view := AsCefViewRef(getVal(0))
			newBounds := *(*TCefRect)(getPtr(getVal(1)))
			cb(view, newBounds)
		},
	}
}

func makeTOnViewParentViewChangedEvent(cb TOnViewParentViewChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewParentViewChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const view: ICefView; added: boolean; const parent: ICefView);
			view := AsCefViewRef(getVal(0))
			added := api.GoBool(getVal(1))
			parent := AsCefViewRef(getVal(2))
			cb(view, added, parent)
		},
	}
}

func makeTOnViewThemeChangedEvent(cb TOnViewThemeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewThemeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const view: ICefView);
			view := AsCefViewRef(getVal(0))
			cb(view)
		},
	}
}

func makeTOnViewWindowChangedEvent(cb TOnViewWindowChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnViewWindowChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const view: ICefView; added: boolean);
			view := AsCefViewRef(getVal(0))
			added := api.GoBool(getVal(1))
			cb(view, added)
		},
	}
}

func makeTOnVirtualKeyboardRequested(cb TOnVirtualKeyboardRequested) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnVirtualKeyboardRequested",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const browser: ICefBrowser; input_mode: TCefTextInpuMode);
			sender := lcl.AsObject(getVal(0))
			browser := AsCefBrowserRef(getVal(1))
			inputMode := cefTypes.TCefTextInpuMode(getVal(2))
			cb(sender, browser, inputMode)
		},
	}
}

func makeTOnWebKitInitializedEvent(cb TOnWebKitInitializedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebKitInitializedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTOnWebSocketConnected(cb TOnWebSocketConnected) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebSocketConnected",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			cb(sender, server, connectionId)
		},
	}
}

func makeTOnWebSocketMessage(cb TOnWebSocketMessage) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebSocketMessage",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const data: Pointer; data_size: NativeUInt);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			data := uintptr(getVal(3))
			dataSize := cefTypes.NativeUInt(getVal(4))
			cb(sender, server, connectionId, data, dataSize)
		},
	}
}

func makeTOnWebSocketRequest(cb TOnWebSocketRequest) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebSocketRequest",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const server: ICefServer; connection_id: Integer; const client_address: ustring; const request: ICefRequest; const callback: ICefCallback);
			sender := lcl.AsObject(getVal(0))
			server := AsCEFServerRef(getVal(1))
			connectionId := int32(getVal(2))
			clientAddress := api.GoStr(getVal(3))
			request := AsCefRequestRef(getVal(4))
			callback := AsCefCallbackRef(getVal(5))
			cb(sender, server, connectionId, clientAddress, request, callback)
		},
	}
}

func makeTOnWindowAcceleratorEvent(cb TOnWindowAcceleratorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowAcceleratorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const window_: ICefWindow; command_id: Integer; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			commandId := int32(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(window, commandId, result)
		},
	}
}

func makeTOnWindowAcceptsFirstMouseEvent(cb TOnWindowAcceptsFirstMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowAcceptsFirstMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: TCefState);
			window := AsCefWindowRef(getVal(0))
			result := (*cefTypes.TCefState)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowActivationChangedEvent(cb TOnWindowActivationChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowActivationChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; active: boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			active := api.GoBool(getVal(2))
			cb(sender, window, active)
		},
	}
}

func makeTOnWindowBoundsChangedEvent(cb TOnWindowBoundsChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowBoundsChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; const new_bounds: TCefRect);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			newBounds := *(*TCefRect)(getPtr(getVal(2)))
			cb(sender, window, newBounds)
		},
	}
}

func makeTOnWindowCanCloseEvent(cb TOnWindowCanCloseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowCanCloseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowCanMaximizeEvent(cb TOnWindowCanMaximizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowCanMaximizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowCanMinimizeEvent(cb TOnWindowCanMinimizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowCanMinimizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowCanResizeEvent(cb TOnWindowCanResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowCanResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowChangedEvent(cb TOnWindowChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const view: ICefView; added: boolean);
			sender := lcl.AsObject(getVal(0))
			view := AsCefViewRef(getVal(1))
			added := api.GoBool(getVal(2))
			cb(sender, view, added)
		},
	}
}

func makeTOnWindowClosingEvent(cb TOnWindowClosingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowClosingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			cb(sender, window)
		},
	}
}

func makeTOnWindowCreatedEvent(cb TOnWindowCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			cb(sender, window)
		},
	}
}

func makeTOnWindowDestroyedEvent(cb TOnWindowDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const Sender: TObject; const window_: ICefWindow);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			cb(sender, window)
		},
	}
}

func makeTOnWindowFullscreenTransitionEvent(cb TOnWindowFullscreenTransitionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowFullscreenTransitionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; is_completed: boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			isCompleted := api.GoBool(getVal(2))
			cb(sender, window, isCompleted)
		},
	}
}

func makeTOnWindowGetInitialBoundsEvent(cb TOnWindowGetInitialBoundsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetInitialBoundsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: TCefRect);
			window := AsCefWindowRef(getVal(0))
			result := (*TCefRect)(getPtr(getVal(1)))
			cb(window, result)
			*(*TCefRect)(getPtr(getVal(1))) = *result
		},
	}
}

func makeTOnWindowGetInitialShowStateEvent(cb TOnWindowGetInitialShowStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetInitialShowStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: TCefShowState);
			window := AsCefWindowRef(getVal(0))
			result := (*cefTypes.TCefShowState)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowGetLinuxWindowPropertiesEvent(cb TOnWindowGetLinuxWindowPropertiesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetLinuxWindowPropertiesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const window_: ICefWindow; var properties: TLinuxWindowProperties; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			propertiesPtr := *(*tLinuxWindowProperties)(getPtr(getVal(1)))
			properties := propertiesPtr.ToGo()
			result := (*bool)(getPtr(getVal(2)))
			cb(window, &properties, result)
			if r := properties.ToPas(); r != nil {
				*(*tLinuxWindowProperties)(getPtr(getVal(1))) = *r
			}
		},
	}
}

func makeTOnWindowGetParentWindowEvent(cb TOnWindowGetParentWindowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetParentWindowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const window_: ICefWindow; var is_menu: boolean; var can_activate_menu: boolean; var aResult: ICefWindow);
			window := AsCefWindowRef(getVal(0))
			isMenu := (*bool)(getPtr(getVal(1)))
			canActivateMenu := (*bool)(getPtr(getVal(2)))
			result := ICefWindow(AsCefWindowRef(*(*uintptr)(getPtr(getVal(3)))))
			cb(window, isMenu, canActivateMenu, &result)
			if result != nil {
				*(*uintptr)(getPtr(getVal(3))) = result.Instance()
			}
		},
	}
}

func makeTOnWindowGetTitlebarHeightEvent(cb TOnWindowGetTitlebarHeightEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetTitlebarHeightEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const window_: ICefWindow; var titlebar_height: Single; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			titlebarHeight := (*float32)(getPtr(getVal(1)))
			result := (*bool)(getPtr(getVal(2)))
			cb(window, titlebarHeight, result)
		},
	}
}

func makeTOnWindowGetWindowRuntimeStyleEvent(cb TOnWindowGetWindowRuntimeStyleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowGetWindowRuntimeStyleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var aResult: TCefRuntimeStyle);
			result := (*cefTypes.TCefRuntimeStyle)(getPtr(getVal(0)))
			cb(result)
		},
	}
}

func makeTOnWindowIsFramelessEvent(cb TOnWindowIsFramelessEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowIsFramelessEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowIsWindowModalDialogEvent(cb TOnWindowIsWindowModalDialogEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowIsWindowModalDialogEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWindowKeyEvent(cb TOnWindowKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const window_: ICefWindow; const event: TCefKeyEvent; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			event := *(*TCefKeyEvent)(getPtr(getVal(1)))
			result := (*bool)(getPtr(getVal(2)))
			cb(window, event, result)
		},
	}
}

func makeTOnWindowKeyEventEvent(cb TOnWindowKeyEventEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowKeyEventEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(const Sender: TObject; const window_: ICefWindow; const event: TCefKeyEvent; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			event := *(*TCefKeyEvent)(getPtr(getVal(2)))
			result := (*bool)(getPtr(getVal(3)))
			cb(sender, window, event, result)
		},
	}
}

func makeTOnWindowThemeColorsChangedEvent(cb TOnWindowThemeColorsChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowThemeColorsChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; chrome_theme: Integer);
			window := AsCefWindowRef(getVal(0))
			chromeTheme := int32(getVal(1))
			cb(window, chromeTheme)
		},
	}
}

func makeTOnWindowWindowActivationChangedEvent(cb TOnWindowWindowActivationChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowActivationChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; active: boolean);
			window := AsCefWindowRef(getVal(0))
			active := api.GoBool(getVal(1))
			cb(window, active)
		},
	}
}

func makeTOnWindowWindowBoundsChangedEvent(cb TOnWindowWindowBoundsChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowBoundsChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; const new_bounds: TCefRect);
			window := AsCefWindowRef(getVal(0))
			newBounds := *(*TCefRect)(getPtr(getVal(1)))
			cb(window, newBounds)
		},
	}
}

func makeTOnWindowWindowClosingEvent(cb TOnWindowWindowClosingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowClosingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const window_: ICefWindow);
			window := AsCefWindowRef(getVal(0))
			cb(window)
		},
	}
}

func makeTOnWindowWindowCreatedEvent(cb TOnWindowWindowCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const window_: ICefWindow);
			window := AsCefWindowRef(getVal(0))
			cb(window)
		},
	}
}

func makeTOnWindowWindowDestroyedEvent(cb TOnWindowWindowDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(const window_: ICefWindow);
			window := AsCefWindowRef(getVal(0))
			cb(window)
		},
	}
}

func makeTOnWindowWindowFullscreenTransitionEvent(cb TOnWindowWindowFullscreenTransitionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWindowFullscreenTransitionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; is_completed: boolean);
			window := AsCefWindowRef(getVal(0))
			isCompleted := api.GoBool(getVal(1))
			cb(window, isCompleted)
		},
	}
}

func makeTOnWindowWithStandardWindowButtonsEvent(cb TOnWindowWithStandardWindowButtonsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWindowWithStandardWindowButtonsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const window_: ICefWindow; var aResult: boolean);
			window := AsCefWindowRef(getVal(0))
			result := (*bool)(getPtr(getVal(1)))
			cb(window, result)
		},
	}
}

func makeTOnWithStandardWindowButtonsEvent(cb TOnWithStandardWindowButtonsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWithStandardWindowButtonsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TObject; const window_: ICefWindow; var aResult : boolean);
			sender := lcl.AsObject(getVal(0))
			window := AsCefWindowRef(getVal(1))
			result := (*bool)(getPtr(getVal(2)))
			cb(sender, window, result)
		},
	}
}

func makeTOnWriteEllEvent(cb TOnWriteEllEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWriteEllEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): Int64;
			ret := cb()
			*(*int64)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTOnWriteFlushEvent(cb TOnWriteFlushEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWriteFlushEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): Integer;
			ret := cb()
			*(*int32)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTOnWriteMayBlockEvent(cb TOnWriteMayBlockEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWriteMayBlockEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): Boolean;
			ret := cb()
			*(*bool)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTOnWriteSeekEvent(cb TOnWriteSeekEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWriteSeekEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(offset: Int64; whence: Integer): Integer;
			offset := *(*int64)(getPtr(getVal(0)))
			whence := int32(getVal(1))
			ret := cb(offset, whence)
			*(*int32)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTOnWriteWriteEvent(cb TOnWriteWriteEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWriteWriteEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(const ptr: Pointer; size: NativeUInt; n: NativeUInt): NativeUInt;
			ptr := uintptr(getVal(0))
			size := cefTypes.NativeUInt(getVal(1))
			N := cefTypes.NativeUInt(getVal(2))
			ret := cb(ptr, size, N)
			*(*cefTypes.NativeUInt)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTOnZoomPctAvailable(cb TOnZoomPctAvailable) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnZoomPctAvailable",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aZoomPct : double);
			sender := lcl.AsObject(getVal(0))
			zoomPct := *(*float64)(getPtr(getVal(1)))
			cb(sender, zoomPct)
		},
	}
}

func makeTStartDockEvent(cb TStartDockEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TStartDockEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var DragObject: TDragDockObject);
			sender := lcl.AsObject(getVal(0))
			var dragObject lcl.IDragDockObject
			dragObject = lcl.AsDragDockObject(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &dragObject)
			if dragObject != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
			}
		},
	}
}

func makeTStartDragEvent(cb TStartDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TStartDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var DragObject: TDragObject);
			sender := lcl.AsObject(getVal(0))
			var dragObject lcl.IDragObject
			dragObject = lcl.AsDragObject(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &dragObject)
			if dragObject != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
			}
		},
	}
}
