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
)

// IChromiumWindow Parent: ICEFLinkedWinControlBase
type IChromiumWindow interface {
	ICEFLinkedWinControlBase
	// CreateBrowser
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumWindow.OnAfterCreated
	//  event is triggered.
	CreateBrowser() bool // function
	// CloseBrowser
	//  Request that the browser close. The JavaScript 'onbeforeunload' event will
	//  be fired. If |aForceClose| is false (0) the event handler, if any, will be
	//  allowed to prompt the user and the user can optionally cancel the close.
	//  If |aForceClose| is true (1) the prompt will not be displayed and the
	//  close will proceed. Results in a call to
	//  ICefLifeSpanHandler.DoClose() if the event handler allows the close
	//  or if |aForceClose| is true (1). See ICefLifeSpanHandler.DoClose()
	//  documentation for additional usage information.
	CloseBrowser(forceClose bool) // procedure
	// LoadURL
	//  Used to navigate to a URL.
	LoadURL(uRL string) // procedure
	// NotifyMoveOrResizeStarted
	//  Notify the browser that the window hosting it is about to be moved or
	//  resized. This function is only used on Windows and Linux.
	NotifyMoveOrResizeStarted() // procedure
	// ChromiumBrowser
	//  TChromium instance used by this component.
	ChromiumBrowser() IChromium // property ChromiumBrowser Getter
	// Initialized
	//  Returns true when the browser is fully initialized and it's not being closed.
	Initialized() bool                 // property Initialized Getter
	SetOnClose(fn TNotifyEvent)        // property event
	SetOnBeforeClose(fn TNotifyEvent)  // property event
	SetOnAfterCreated(fn TNotifyEvent) // property event
}

type TChromiumWindow struct {
	TCEFLinkedWinControlBase
}

func (m *TChromiumWindow) CreateBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumWindowAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumWindow) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	chromiumWindowAPI().SysCallN(2, m.Instance(), api.PasBool(forceClose))
}

func (m *TChromiumWindow) LoadURL(uRL string) {
	if !m.IsValid() {
		return
	}
	chromiumWindowAPI().SysCallN(3, m.Instance(), api.PasStr(uRL))
}

func (m *TChromiumWindow) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	chromiumWindowAPI().SysCallN(4, m.Instance())
}

func (m *TChromiumWindow) ChromiumBrowser() IChromium {
	if !m.IsValid() {
		return nil
	}
	r := chromiumWindowAPI().SysCallN(5, m.Instance())
	return AsChromium(r)
}

func (m *TChromiumWindow) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumWindowAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumWindow) SetOnClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumWindow) SetOnBeforeClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumWindow) SetOnAfterCreated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
}

// NewChromiumWindow class constructor
func NewChromiumWindow(owner lcl.IComponent) IChromiumWindow {
	r := chromiumWindowAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsChromiumWindow(r)
}

var (
	chromiumWindowOnce   base.Once
	chromiumWindowImport *imports.Imports = nil
)

func chromiumWindowAPI() *imports.Imports {
	chromiumWindowOnce.Do(func() {
		chromiumWindowImport = api.NewDefaultImports()
		chromiumWindowImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChromiumWindow_Create", 0), // constructor NewChromiumWindow
			/* 1 */ imports.NewTable("TChromiumWindow_CreateBrowser", 0), // function CreateBrowser
			/* 2 */ imports.NewTable("TChromiumWindow_CloseBrowser", 0), // procedure CloseBrowser
			/* 3 */ imports.NewTable("TChromiumWindow_LoadURL", 0), // procedure LoadURL
			/* 4 */ imports.NewTable("TChromiumWindow_NotifyMoveOrResizeStarted", 0), // procedure NotifyMoveOrResizeStarted
			/* 5 */ imports.NewTable("TChromiumWindow_ChromiumBrowser", 0), // property ChromiumBrowser
			/* 6 */ imports.NewTable("TChromiumWindow_Initialized", 0), // property Initialized
			/* 7 */ imports.NewTable("TChromiumWindow_OnClose", 0), // event OnClose
			/* 8 */ imports.NewTable("TChromiumWindow_OnBeforeClose", 0), // event OnBeforeClose
			/* 9 */ imports.NewTable("TChromiumWindow_OnAfterCreated", 0), // event OnAfterCreated
		}
	})
	return chromiumWindowImport
}
