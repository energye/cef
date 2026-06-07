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

// IEngDisplayHandler Parent: ICefDisplayHandlerOwn
type IEngDisplayHandler interface {
	ICefDisplayHandlerOwn
	SetOnDisplayTooltip(fn TOnDisplayTooltipEvent)                                 // property event
	SetOnDisplayConsoleMessage(fn TOnDisplayConsoleMessageEvent)                   // property event
	SetOnDisplayAutoResize(fn TOnDisplayAutoResizeEvent)                           // property event
	SetOnDisplayContentsBoundsChange(fn TOnDisplayContentsBoundsChangeEvent)       // property event
	SetOnDisplayGetRootWindowScreenRect(fn TOnDisplayGetRootWindowScreenRectEvent) // property event
	SetOnDisplayAddressChange(fn TOnDisplayAddressChangeEvent)                     // property event
	SetOnDisplayTitleChange(fn TOnDisplayTitleChangeEvent)                         // property event
	SetOnDisplayFaviconUrlChange(fn TOnDisplayFaviconUrlChangeEvent)               // property event
	SetOnDisplayFullScreenModeChange(fn TOnDisplayFullScreenModeChangeEvent)       // property event
	SetOnDisplayStatusMessage(fn TOnDisplayStatusMessageEvent)                     // property event
	SetOnDisplayLoadingProgressChange(fn TOnDisplayLoadingProgressChangeEvent)     // property event
	SetOnDisplayCursorChange(fn TOnDisplayCursorChangeEvent)                       // property event
	SetOnDisplayMediaAccessChange(fn TOnDisplayMediaAccessChangeEvent)             // property event
	AsIntfDisplayHandler() uintptr
}

type TEngDisplayHandler struct {
	TCefDisplayHandlerOwn
}

func (m *TEngDisplayHandler) SetOnDisplayTooltip(fn TOnDisplayTooltipEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayTooltipEvent(fn)
	base.SetEvent(m, 1, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayConsoleMessage(fn TOnDisplayConsoleMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayConsoleMessageEvent(fn)
	base.SetEvent(m, 2, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayAutoResize(fn TOnDisplayAutoResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayAutoResizeEvent(fn)
	base.SetEvent(m, 3, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayContentsBoundsChange(fn TOnDisplayContentsBoundsChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayContentsBoundsChangeEvent(fn)
	base.SetEvent(m, 4, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayGetRootWindowScreenRect(fn TOnDisplayGetRootWindowScreenRectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayGetRootWindowScreenRectEvent(fn)
	base.SetEvent(m, 5, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayAddressChange(fn TOnDisplayAddressChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayAddressChangeEvent(fn)
	base.SetEvent(m, 6, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayTitleChange(fn TOnDisplayTitleChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayTitleChangeEvent(fn)
	base.SetEvent(m, 7, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayFaviconUrlChange(fn TOnDisplayFaviconUrlChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayFaviconUrlChangeEvent(fn)
	base.SetEvent(m, 8, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayFullScreenModeChange(fn TOnDisplayFullScreenModeChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayFullScreenModeChangeEvent(fn)
	base.SetEvent(m, 9, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayStatusMessage(fn TOnDisplayStatusMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayStatusMessageEvent(fn)
	base.SetEvent(m, 10, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayLoadingProgressChange(fn TOnDisplayLoadingProgressChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayLoadingProgressChangeEvent(fn)
	base.SetEvent(m, 11, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayCursorChange(fn TOnDisplayCursorChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayCursorChangeEvent(fn)
	base.SetEvent(m, 12, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) SetOnDisplayMediaAccessChange(fn TOnDisplayMediaAccessChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDisplayMediaAccessChangeEvent(fn)
	base.SetEvent(m, 13, engDisplayHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngDisplayHandler) AsIntfDisplayHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngDisplayHandler class constructor
func NewEngDisplayHandler() IEngDisplayHandler {
	var displayHandlerPtr uintptr // ICefDisplayHandler
	r := engDisplayHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&displayHandlerPtr)))
	ret := AsEngDisplayHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, displayHandlerPtr)
	}
	return ret
}

var (
	engDisplayHandlerOnce   base.Once
	engDisplayHandlerImport *imports.Imports = nil
)

func engDisplayHandlerAPI() *imports.Imports {
	engDisplayHandlerOnce.Do(func() {
		engDisplayHandlerImport = api.NewDefaultImports()
		engDisplayHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngDisplayHandler_Create", 0), // constructor NewEngDisplayHandler
			/* 1 */ imports.NewTable("TEngDisplayHandler_OnDisplayTooltip", 0), // event OnDisplayTooltip
			/* 2 */ imports.NewTable("TEngDisplayHandler_OnDisplayConsoleMessage", 0), // event OnDisplayConsoleMessage
			/* 3 */ imports.NewTable("TEngDisplayHandler_OnDisplayAutoResize", 0), // event OnDisplayAutoResize
			/* 4 */ imports.NewTable("TEngDisplayHandler_OnDisplayContentsBoundsChange", 0), // event OnDisplayContentsBoundsChange
			/* 5 */ imports.NewTable("TEngDisplayHandler_OnDisplayGetRootWindowScreenRect", 0), // event OnDisplayGetRootWindowScreenRect
			/* 6 */ imports.NewTable("TEngDisplayHandler_OnDisplayAddressChange", 0), // event OnDisplayAddressChange
			/* 7 */ imports.NewTable("TEngDisplayHandler_OnDisplayTitleChange", 0), // event OnDisplayTitleChange
			/* 8 */ imports.NewTable("TEngDisplayHandler_OnDisplayFaviconUrlChange", 0), // event OnDisplayFaviconUrlChange
			/* 9 */ imports.NewTable("TEngDisplayHandler_OnDisplayFullScreenModeChange", 0), // event OnDisplayFullScreenModeChange
			/* 10 */ imports.NewTable("TEngDisplayHandler_OnDisplayStatusMessage", 0), // event OnDisplayStatusMessage
			/* 11 */ imports.NewTable("TEngDisplayHandler_OnDisplayLoadingProgressChange", 0), // event OnDisplayLoadingProgressChange
			/* 12 */ imports.NewTable("TEngDisplayHandler_OnDisplayCursorChange", 0), // event OnDisplayCursorChange
			/* 13 */ imports.NewTable("TEngDisplayHandler_OnDisplayMediaAccessChange", 0), // event OnDisplayMediaAccessChange
		}
	})
	return engDisplayHandlerImport
}
