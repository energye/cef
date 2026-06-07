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

// IEngWindowDelegate Parent: ICefWindowDelegateOwn
type IEngWindowDelegate interface {
	ICefWindowDelegateOwn
	SetOnWindowWindowCreated(fn TOnWindowWindowCreatedEvent)                           // property event
	SetOnWindowWindowClosing(fn TOnWindowWindowClosingEvent)                           // property event
	SetOnWindowWindowDestroyed(fn TOnWindowWindowDestroyedEvent)                       // property event
	SetOnWindowWindowActivationChanged(fn TOnWindowWindowActivationChangedEvent)       // property event
	SetOnWindowWindowBoundsChanged(fn TOnWindowWindowBoundsChangedEvent)               // property event
	SetOnWindowWindowFullscreenTransition(fn TOnWindowWindowFullscreenTransitionEvent) // property event
	SetOnWindowGetParentWindow(fn TOnWindowGetParentWindowEvent)                       // property event
	SetOnWindowIsWindowModalDialog(fn TOnWindowIsWindowModalDialogEvent)               // property event
	SetOnWindowGetInitialBounds(fn TOnWindowGetInitialBoundsEvent)                     // property event
	SetOnWindowGetInitialShowState(fn TOnWindowGetInitialShowStateEvent)               // property event
	SetOnWindowIsFrameless(fn TOnWindowIsFramelessEvent)                               // property event
	SetOnWindowWithStandardWindowButtons(fn TOnWindowWithStandardWindowButtonsEvent)   // property event
	SetOnWindowGetTitlebarHeight(fn TOnWindowGetTitlebarHeightEvent)                   // property event
	SetOnWindowAcceptsFirstMouse(fn TOnWindowAcceptsFirstMouseEvent)                   // property event
	SetOnWindowCanResize(fn TOnWindowCanResizeEvent)                                   // property event
	SetOnWindowCanMaximize(fn TOnWindowCanMaximizeEvent)                               // property event
	SetOnWindowCanMinimize(fn TOnWindowCanMinimizeEvent)                               // property event
	SetOnWindowCanClose(fn TOnWindowCanCloseEvent)                                     // property event
	SetOnWindowAccelerator(fn TOnWindowAcceleratorEvent)                               // property event
	SetOnWindowKey(fn TOnWindowKeyEvent)                                               // property event
	SetOnWindowThemeColorsChanged(fn TOnWindowThemeColorsChangedEvent)                 // property event
	SetOnWindowGetWindowRuntimeStyle(fn TOnWindowGetWindowRuntimeStyleEvent)           // property event
	SetOnWindowGetLinuxWindowProperties(fn TOnWindowGetLinuxWindowPropertiesEvent)     // property event
	AsIntfWindowDelegate() uintptr
	AsIntfPanelDelegate() uintptr
	AsIntfViewDelegate() uintptr
}

type TEngWindowDelegate struct {
	TCefWindowDelegateOwn
}

func (m *TEngWindowDelegate) SetOnWindowWindowCreated(fn TOnWindowWindowCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowCreatedEvent(fn)
	base.SetEvent(m, 1, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWindowClosing(fn TOnWindowWindowClosingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowClosingEvent(fn)
	base.SetEvent(m, 2, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWindowDestroyed(fn TOnWindowWindowDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowDestroyedEvent(fn)
	base.SetEvent(m, 3, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWindowActivationChanged(fn TOnWindowWindowActivationChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowActivationChangedEvent(fn)
	base.SetEvent(m, 4, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWindowBoundsChanged(fn TOnWindowWindowBoundsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowBoundsChangedEvent(fn)
	base.SetEvent(m, 5, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWindowFullscreenTransition(fn TOnWindowWindowFullscreenTransitionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWindowFullscreenTransitionEvent(fn)
	base.SetEvent(m, 6, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetParentWindow(fn TOnWindowGetParentWindowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetParentWindowEvent(fn)
	base.SetEvent(m, 7, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowIsWindowModalDialog(fn TOnWindowIsWindowModalDialogEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowIsWindowModalDialogEvent(fn)
	base.SetEvent(m, 8, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetInitialBounds(fn TOnWindowGetInitialBoundsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetInitialBoundsEvent(fn)
	base.SetEvent(m, 9, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetInitialShowState(fn TOnWindowGetInitialShowStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetInitialShowStateEvent(fn)
	base.SetEvent(m, 10, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowIsFrameless(fn TOnWindowIsFramelessEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowIsFramelessEvent(fn)
	base.SetEvent(m, 11, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowWithStandardWindowButtons(fn TOnWindowWithStandardWindowButtonsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowWithStandardWindowButtonsEvent(fn)
	base.SetEvent(m, 12, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetTitlebarHeight(fn TOnWindowGetTitlebarHeightEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetTitlebarHeightEvent(fn)
	base.SetEvent(m, 13, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowAcceptsFirstMouse(fn TOnWindowAcceptsFirstMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowAcceptsFirstMouseEvent(fn)
	base.SetEvent(m, 14, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowCanResize(fn TOnWindowCanResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCanResizeEvent(fn)
	base.SetEvent(m, 15, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowCanMaximize(fn TOnWindowCanMaximizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCanMaximizeEvent(fn)
	base.SetEvent(m, 16, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowCanMinimize(fn TOnWindowCanMinimizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCanMinimizeEvent(fn)
	base.SetEvent(m, 17, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowCanClose(fn TOnWindowCanCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowCanCloseEvent(fn)
	base.SetEvent(m, 18, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowAccelerator(fn TOnWindowAcceleratorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowAcceleratorEvent(fn)
	base.SetEvent(m, 19, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowKey(fn TOnWindowKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowKeyEvent(fn)
	base.SetEvent(m, 20, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowThemeColorsChanged(fn TOnWindowThemeColorsChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowThemeColorsChangedEvent(fn)
	base.SetEvent(m, 21, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetWindowRuntimeStyle(fn TOnWindowGetWindowRuntimeStyleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetWindowRuntimeStyleEvent(fn)
	base.SetEvent(m, 22, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) SetOnWindowGetLinuxWindowProperties(fn TOnWindowGetLinuxWindowPropertiesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWindowGetLinuxWindowPropertiesEvent(fn)
	base.SetEvent(m, 23, engWindowDelegateAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEngWindowDelegate) AsIntfWindowDelegate() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TEngWindowDelegate) AsIntfPanelDelegate() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TEngWindowDelegate) AsIntfViewDelegate() uintptr {
	return m.GetIntfPointer(2)
}

// NewEngWindowDelegate class constructor
func NewEngWindowDelegate() IEngWindowDelegate {
	var windowDelegatePtr uintptr // ICefWindowDelegate
	var panelDelegatePtr uintptr  // ICefPanelDelegate
	var viewDelegatePtr uintptr   // ICefViewDelegate
	r := engWindowDelegateAPI().SysCallN(0, uintptr(base.UnsafePointer(&windowDelegatePtr)), uintptr(base.UnsafePointer(&panelDelegatePtr)), uintptr(base.UnsafePointer(&viewDelegatePtr)))
	ret := AsEngWindowDelegate(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(3)
		intf.SetIntfPointer(0, windowDelegatePtr)
		intf.SetIntfPointer(1, panelDelegatePtr)
		intf.SetIntfPointer(2, viewDelegatePtr)
	}
	return ret
}

var (
	engWindowDelegateOnce   base.Once
	engWindowDelegateImport *imports.Imports = nil
)

func engWindowDelegateAPI() *imports.Imports {
	engWindowDelegateOnce.Do(func() {
		engWindowDelegateImport = api.NewDefaultImports()
		engWindowDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngWindowDelegate_Create", 0), // constructor NewEngWindowDelegate
			/* 1 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowCreated", 0), // event OnWindowWindowCreated
			/* 2 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowClosing", 0), // event OnWindowWindowClosing
			/* 3 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowDestroyed", 0), // event OnWindowWindowDestroyed
			/* 4 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowActivationChanged", 0), // event OnWindowWindowActivationChanged
			/* 5 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowBoundsChanged", 0), // event OnWindowWindowBoundsChanged
			/* 6 */ imports.NewTable("TEngWindowDelegate_OnWindowWindowFullscreenTransition", 0), // event OnWindowWindowFullscreenTransition
			/* 7 */ imports.NewTable("TEngWindowDelegate_OnWindowGetParentWindow", 0), // event OnWindowGetParentWindow
			/* 8 */ imports.NewTable("TEngWindowDelegate_OnWindowIsWindowModalDialog", 0), // event OnWindowIsWindowModalDialog
			/* 9 */ imports.NewTable("TEngWindowDelegate_OnWindowGetInitialBounds", 0), // event OnWindowGetInitialBounds
			/* 10 */ imports.NewTable("TEngWindowDelegate_OnWindowGetInitialShowState", 0), // event OnWindowGetInitialShowState
			/* 11 */ imports.NewTable("TEngWindowDelegate_OnWindowIsFrameless", 0), // event OnWindowIsFrameless
			/* 12 */ imports.NewTable("TEngWindowDelegate_OnWindowWithStandardWindowButtons", 0), // event OnWindowWithStandardWindowButtons
			/* 13 */ imports.NewTable("TEngWindowDelegate_OnWindowGetTitlebarHeight", 0), // event OnWindowGetTitlebarHeight
			/* 14 */ imports.NewTable("TEngWindowDelegate_OnWindowAcceptsFirstMouse", 0), // event OnWindowAcceptsFirstMouse
			/* 15 */ imports.NewTable("TEngWindowDelegate_OnWindowCanResize", 0), // event OnWindowCanResize
			/* 16 */ imports.NewTable("TEngWindowDelegate_OnWindowCanMaximize", 0), // event OnWindowCanMaximize
			/* 17 */ imports.NewTable("TEngWindowDelegate_OnWindowCanMinimize", 0), // event OnWindowCanMinimize
			/* 18 */ imports.NewTable("TEngWindowDelegate_OnWindowCanClose", 0), // event OnWindowCanClose
			/* 19 */ imports.NewTable("TEngWindowDelegate_OnWindowAccelerator", 0), // event OnWindowAccelerator
			/* 20 */ imports.NewTable("TEngWindowDelegate_OnWindowKey", 0), // event OnWindowKey
			/* 21 */ imports.NewTable("TEngWindowDelegate_OnWindowThemeColorsChanged", 0), // event OnWindowThemeColorsChanged
			/* 22 */ imports.NewTable("TEngWindowDelegate_OnWindowGetWindowRuntimeStyle", 0), // event OnWindowGetWindowRuntimeStyle
			/* 23 */ imports.NewTable("TEngWindowDelegate_OnWindowGetLinuxWindowProperties", 0), // event OnWindowGetLinuxWindowProperties
		}
	})
	return engWindowDelegateImport
}
