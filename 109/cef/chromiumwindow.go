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
	CreateBrowser() bool               // function
	CloseBrowser(forceClose bool)      // procedure
	LoadURL(uRL string)                // procedure
	NotifyMoveOrResizeStarted()        // procedure
	ChromiumBrowser() IChromium        // property ChromiumBrowser Getter
	Initialized() bool                 // property Initialized Getter
	UseSetFocus() bool                 // property UseSetFocus Getter
	SetUseSetFocus(value bool)         // property UseSetFocus Setter
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

func (m *TChromiumWindow) UseSetFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := chromiumWindowAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TChromiumWindow) SetUseSetFocus(value bool) {
	if !m.IsValid() {
		return
	}
	chromiumWindowAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TChromiumWindow) SetOnClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumWindow) SetOnBeforeClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChromiumWindow) SetOnAfterCreated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, chromiumWindowAPI(), api.MakeEventDataPtr(cb))
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
			/* 7 */ imports.NewTable("TChromiumWindow_UseSetFocus", 0), // property UseSetFocus
			/* 8 */ imports.NewTable("TChromiumWindow_OnClose", 0), // event OnClose
			/* 9 */ imports.NewTable("TChromiumWindow_OnBeforeClose", 0), // event OnBeforeClose
			/* 10 */ imports.NewTable("TChromiumWindow_OnAfterCreated", 0), // event OnAfterCreated
		}
	})
	return chromiumWindowImport
}
