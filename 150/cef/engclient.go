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
)

// IEngClient Parent: ICefClientOwn
type IEngClient interface {
	ICefClientOwn
	SetOnClientProcessMessageReceived(fn TOnClientProcessMessageReceivedEvent) // property event
	SetOnClientGetAudioHandler(fn TOnClientGetAudioHandlerEvent)               // property event
	SetOnClientGetCommandHandler(fn TOnClientGetCommandHandlerEvent)           // property event
	SetOnClientGetContextMenuHandler(fn TOnClientGetContextMenuHandlerEvent)   // property event
	SetOnClientGetDialogHandler(fn TOnClientGetDialogHandlerEvent)             // property event
	SetOnClientGetDisplayHandler(fn TOnClientGetDisplayHandlerEvent)           // property event
	SetOnClientGetDownloadHandler(fn TOnClientGetDownloadHandlerEvent)         // property event
	SetOnClientGetDragHandler(fn TOnClientGetDragHandlerEvent)                 // property event
	SetOnClientGetFindHandler(fn TOnClientGetFindHandlerEvent)                 // property event
	SetOnClientGetFocusHandler(fn TOnClientGetFocusHandlerEvent)               // property event
	SetOnClientGetFrameHandler(fn TOnClientGetFrameHandlerEvent)               // property event
	SetOnClientGetPermissionHandler(fn TOnClientGetPermissionHandlerEvent)     // property event
	SetOnClientGetJsdialogHandler(fn TOnClientGetJsdialogHandlerEvent)         // property event
	SetOnClientGetKeyboardHandler(fn TOnClientGetKeyboardHandlerEvent)         // property event
	SetOnClientGetLifeSpanHandler(fn TOnClientGetLifeSpanHandlerEvent)         // property event
	SetOnClientGetLoadHandler(fn TOnClientGetLoadHandlerEvent)                 // property event
	SetOnClientGetPrintHandler(fn TOnClientGetPrintHandlerEvent)               // property event
	SetOnClientGetRenderHandler(fn TOnClientGetRenderHandlerEvent)             // property event
	SetOnClientGetRequestHandler(fn TOnClientGetRequestHandlerEvent)           // property event
	AsIntfClient() uintptr
}

type TEngClient struct {
	TCefClientOwn
}

func (m *TEngClient) SetOnClientProcessMessageReceived(fn TOnClientProcessMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientProcessMessageReceivedEvent(fn)
	base.SetEvent(m, 1, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetAudioHandler(fn TOnClientGetAudioHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetAudioHandlerEvent(fn)
	base.SetEvent(m, 2, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetCommandHandler(fn TOnClientGetCommandHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetCommandHandlerEvent(fn)
	base.SetEvent(m, 3, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetContextMenuHandler(fn TOnClientGetContextMenuHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetContextMenuHandlerEvent(fn)
	base.SetEvent(m, 4, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetDialogHandler(fn TOnClientGetDialogHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetDialogHandlerEvent(fn)
	base.SetEvent(m, 5, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetDisplayHandler(fn TOnClientGetDisplayHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetDisplayHandlerEvent(fn)
	base.SetEvent(m, 6, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetDownloadHandler(fn TOnClientGetDownloadHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetDownloadHandlerEvent(fn)
	base.SetEvent(m, 7, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetDragHandler(fn TOnClientGetDragHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetDragHandlerEvent(fn)
	base.SetEvent(m, 8, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetFindHandler(fn TOnClientGetFindHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetFindHandlerEvent(fn)
	base.SetEvent(m, 9, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetFocusHandler(fn TOnClientGetFocusHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetFocusHandlerEvent(fn)
	base.SetEvent(m, 10, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetFrameHandler(fn TOnClientGetFrameHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetFrameHandlerEvent(fn)
	base.SetEvent(m, 11, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetPermissionHandler(fn TOnClientGetPermissionHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetPermissionHandlerEvent(fn)
	base.SetEvent(m, 12, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetJsdialogHandler(fn TOnClientGetJsdialogHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetJsdialogHandlerEvent(fn)
	base.SetEvent(m, 13, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetKeyboardHandler(fn TOnClientGetKeyboardHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetKeyboardHandlerEvent(fn)
	base.SetEvent(m, 14, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetLifeSpanHandler(fn TOnClientGetLifeSpanHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetLifeSpanHandlerEvent(fn)
	base.SetEvent(m, 15, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetLoadHandler(fn TOnClientGetLoadHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetLoadHandlerEvent(fn)
	base.SetEvent(m, 16, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetPrintHandler(fn TOnClientGetPrintHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetPrintHandlerEvent(fn)
	base.SetEvent(m, 17, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetRenderHandler(fn TOnClientGetRenderHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetRenderHandlerEvent(fn)
	base.SetEvent(m, 18, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) SetOnClientGetRequestHandler(fn TOnClientGetRequestHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientGetRequestHandlerEvent(fn)
	base.SetEvent(m, 19, engClientAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngClient) AsIntfClient() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngClient class constructor
func NewEngClient() IEngClient {
	var clientPtr uintptr // ICefClient
	r := engClientAPI().SysCallN(0, uintptr(base.UnsafePointer(&clientPtr)))
	ret := AsEngClient(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, clientPtr)
	}
	return ret
}

var (
	engClientOnce   base.Once
	engClientImport *imports.Imports = nil
)

func engClientAPI() *imports.Imports {
	engClientOnce.Do(func() {
		engClientImport = api.NewDefaultImports()
		engClientImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngClient_Create", 0), // constructor NewEngClient
			/* 1 */ imports.NewTable("TEngClient_OnClientProcessMessageReceived", 0), // event OnClientProcessMessageReceived
			/* 2 */ imports.NewTable("TEngClient_OnClientGetAudioHandler", 0), // event OnClientGetAudioHandler
			/* 3 */ imports.NewTable("TEngClient_OnClientGetCommandHandler", 0), // event OnClientGetCommandHandler
			/* 4 */ imports.NewTable("TEngClient_OnClientGetContextMenuHandler", 0), // event OnClientGetContextMenuHandler
			/* 5 */ imports.NewTable("TEngClient_OnClientGetDialogHandler", 0), // event OnClientGetDialogHandler
			/* 6 */ imports.NewTable("TEngClient_OnClientGetDisplayHandler", 0), // event OnClientGetDisplayHandler
			/* 7 */ imports.NewTable("TEngClient_OnClientGetDownloadHandler", 0), // event OnClientGetDownloadHandler
			/* 8 */ imports.NewTable("TEngClient_OnClientGetDragHandler", 0), // event OnClientGetDragHandler
			/* 9 */ imports.NewTable("TEngClient_OnClientGetFindHandler", 0), // event OnClientGetFindHandler
			/* 10 */ imports.NewTable("TEngClient_OnClientGetFocusHandler", 0), // event OnClientGetFocusHandler
			/* 11 */ imports.NewTable("TEngClient_OnClientGetFrameHandler", 0), // event OnClientGetFrameHandler
			/* 12 */ imports.NewTable("TEngClient_OnClientGetPermissionHandler", 0), // event OnClientGetPermissionHandler
			/* 13 */ imports.NewTable("TEngClient_OnClientGetJsdialogHandler", 0), // event OnClientGetJsdialogHandler
			/* 14 */ imports.NewTable("TEngClient_OnClientGetKeyboardHandler", 0), // event OnClientGetKeyboardHandler
			/* 15 */ imports.NewTable("TEngClient_OnClientGetLifeSpanHandler", 0), // event OnClientGetLifeSpanHandler
			/* 16 */ imports.NewTable("TEngClient_OnClientGetLoadHandler", 0), // event OnClientGetLoadHandler
			/* 17 */ imports.NewTable("TEngClient_OnClientGetPrintHandler", 0), // event OnClientGetPrintHandler
			/* 18 */ imports.NewTable("TEngClient_OnClientGetRenderHandler", 0), // event OnClientGetRenderHandler
			/* 19 */ imports.NewTable("TEngClient_OnClientGetRequestHandler", 0), // event OnClientGetRequestHandler
		}
	})
	return engClientImport
}
