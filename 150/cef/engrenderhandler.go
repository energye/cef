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

// IEngRenderHandler Parent: ICefRenderHandlerOwn
type IEngRenderHandler interface {
	ICefRenderHandlerOwn
	SetOnRenderGetRootScreenRect(fn TOnRenderGetRootScreenRectEvent)                   // property event
	SetOnRenderGetScreenPoint(fn TOnRenderGetScreenPointEvent)                         // property event
	SetOnRenderGetScreenInfo(fn TOnRenderGetScreenInfoEvent)                           // property event
	SetOnRenderStartDragging(fn TOnRenderStartDraggingEvent)                           // property event
	SetOnRenderGetAccessibilityHandler(fn TOnRenderGetAccessibilityHandlerEvent)       // property event
	SetOnRenderGetViewRect(fn TOnRenderGetViewRectEvent)                               // property event
	SetOnRenderPopupShow(fn TOnRenderPopupShowEvent)                                   // property event
	SetOnRenderPopupSize(fn TOnRenderPopupSizeEvent)                                   // property event
	SetOnRenderPaint(fn TOnRenderPaintEvent)                                           // property event
	SetOnRenderAcceleratedPaint(fn TOnRenderAcceleratedPaintEvent)                     // property event
	SetOnRenderGetTouchHandleSize(fn TOnRenderGetTouchHandleSizeEvent)                 // property event
	SetOnRenderTouchHandleStateChanged(fn TOnRenderTouchHandleStateChangedEvent)       // property event
	SetOnRenderUpdateDragCursor(fn TOnRenderUpdateDragCursorEvent)                     // property event
	SetOnRenderScrollOffsetChanged(fn TOnRenderScrollOffsetChangedEvent)               // property event
	SetOnRenderIMECompositionRangeChanged(fn TOnRenderIMECompositionRangeChangedEvent) // property event
	SetOnRenderTextSelectionChanged(fn TOnRenderTextSelectionChangedEvent)             // property event
	SetOnRenderVirtualKeyboardRequested(fn TOnRenderVirtualKeyboardRequestedEvent)     // property event
	AsIntfRenderHandler() uintptr
}

type TEngRenderHandler struct {
	TCefRenderHandlerOwn
}

func (m *TEngRenderHandler) SetOnRenderGetRootScreenRect(fn TOnRenderGetRootScreenRectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetRootScreenRectEvent(fn)
	base.SetEvent(m, 1, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderGetScreenPoint(fn TOnRenderGetScreenPointEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetScreenPointEvent(fn)
	base.SetEvent(m, 2, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderGetScreenInfo(fn TOnRenderGetScreenInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetScreenInfoEvent(fn)
	base.SetEvent(m, 3, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderStartDragging(fn TOnRenderStartDraggingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderStartDraggingEvent(fn)
	base.SetEvent(m, 4, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderGetAccessibilityHandler(fn TOnRenderGetAccessibilityHandlerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetAccessibilityHandlerEvent(fn)
	base.SetEvent(m, 5, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderGetViewRect(fn TOnRenderGetViewRectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetViewRectEvent(fn)
	base.SetEvent(m, 6, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderPopupShow(fn TOnRenderPopupShowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderPopupShowEvent(fn)
	base.SetEvent(m, 7, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderPopupSize(fn TOnRenderPopupSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderPopupSizeEvent(fn)
	base.SetEvent(m, 8, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderPaint(fn TOnRenderPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderPaintEvent(fn)
	base.SetEvent(m, 9, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderAcceleratedPaint(fn TOnRenderAcceleratedPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderAcceleratedPaintEvent(fn)
	base.SetEvent(m, 10, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderGetTouchHandleSize(fn TOnRenderGetTouchHandleSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderGetTouchHandleSizeEvent(fn)
	base.SetEvent(m, 11, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderTouchHandleStateChanged(fn TOnRenderTouchHandleStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderTouchHandleStateChangedEvent(fn)
	base.SetEvent(m, 12, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderUpdateDragCursor(fn TOnRenderUpdateDragCursorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderUpdateDragCursorEvent(fn)
	base.SetEvent(m, 13, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderScrollOffsetChanged(fn TOnRenderScrollOffsetChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderScrollOffsetChangedEvent(fn)
	base.SetEvent(m, 14, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderIMECompositionRangeChanged(fn TOnRenderIMECompositionRangeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderIMECompositionRangeChangedEvent(fn)
	base.SetEvent(m, 15, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderTextSelectionChanged(fn TOnRenderTextSelectionChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderTextSelectionChangedEvent(fn)
	base.SetEvent(m, 16, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) SetOnRenderVirtualKeyboardRequested(fn TOnRenderVirtualKeyboardRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRenderVirtualKeyboardRequestedEvent(fn)
	base.SetEvent(m, 17, engRenderHandlerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngRenderHandler) AsIntfRenderHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewEngRenderHandler class constructor
func NewEngRenderHandler() IEngRenderHandler {
	var renderHandlerPtr uintptr // ICefRenderHandler
	r := engRenderHandlerAPI().SysCallN(0, uintptr(base.UnsafePointer(&renderHandlerPtr)))
	ret := AsEngRenderHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, renderHandlerPtr)
	}
	return ret
}

var (
	engRenderHandlerOnce   base.Once
	engRenderHandlerImport *imports.Imports = nil
)

func engRenderHandlerAPI() *imports.Imports {
	engRenderHandlerOnce.Do(func() {
		engRenderHandlerImport = api.NewDefaultImports()
		engRenderHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngRenderHandler_Create", 0), // constructor NewEngRenderHandler
			/* 1 */ imports.NewTable("TEngRenderHandler_OnRenderGetRootScreenRect", 0), // event OnRenderGetRootScreenRect
			/* 2 */ imports.NewTable("TEngRenderHandler_OnRenderGetScreenPoint", 0), // event OnRenderGetScreenPoint
			/* 3 */ imports.NewTable("TEngRenderHandler_OnRenderGetScreenInfo", 0), // event OnRenderGetScreenInfo
			/* 4 */ imports.NewTable("TEngRenderHandler_OnRenderStartDragging", 0), // event OnRenderStartDragging
			/* 5 */ imports.NewTable("TEngRenderHandler_OnRenderGetAccessibilityHandler", 0), // event OnRenderGetAccessibilityHandler
			/* 6 */ imports.NewTable("TEngRenderHandler_OnRenderGetViewRect", 0), // event OnRenderGetViewRect
			/* 7 */ imports.NewTable("TEngRenderHandler_OnRenderPopupShow", 0), // event OnRenderPopupShow
			/* 8 */ imports.NewTable("TEngRenderHandler_OnRenderPopupSize", 0), // event OnRenderPopupSize
			/* 9 */ imports.NewTable("TEngRenderHandler_OnRenderPaint", 0), // event OnRenderPaint
			/* 10 */ imports.NewTable("TEngRenderHandler_OnRenderAcceleratedPaint", 0), // event OnRenderAcceleratedPaint
			/* 11 */ imports.NewTable("TEngRenderHandler_OnRenderGetTouchHandleSize", 0), // event OnRenderGetTouchHandleSize
			/* 12 */ imports.NewTable("TEngRenderHandler_OnRenderTouchHandleStateChanged", 0), // event OnRenderTouchHandleStateChanged
			/* 13 */ imports.NewTable("TEngRenderHandler_OnRenderUpdateDragCursor", 0), // event OnRenderUpdateDragCursor
			/* 14 */ imports.NewTable("TEngRenderHandler_OnRenderScrollOffsetChanged", 0), // event OnRenderScrollOffsetChanged
			/* 15 */ imports.NewTable("TEngRenderHandler_OnRenderIMECompositionRangeChanged", 0), // event OnRenderIMECompositionRangeChanged
			/* 16 */ imports.NewTable("TEngRenderHandler_OnRenderTextSelectionChanged", 0), // event OnRenderTextSelectionChanged
			/* 17 */ imports.NewTable("TEngRenderHandler_OnRenderVirtualKeyboardRequested", 0), // event OnRenderVirtualKeyboardRequested
		}
	})
	return engRenderHandlerImport
}
