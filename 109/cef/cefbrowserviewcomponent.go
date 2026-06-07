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

// ICefBrowserViewDelegateEvents Parent: ICefViewDelegateEvents
type ICefBrowserViewDelegateEvents interface {
	ICefViewDelegateEvents
}

// ICEFBrowserViewComponent Parent: ICefBrowserViewDelegateEvents ICEFViewComponent
type ICEFBrowserViewComponent interface {
	ICefBrowserViewDelegateEvents
	ICEFViewComponent
	CreateBrowserView(client IEngClient, url string, settings TCefBrowserSettings, extraInfo ICefDictionaryValue, requestContext ICefRequestContext) bool // function
	GetForBrowser(browser ICefBrowser) bool                                                                                                               // function
	SetPreferAccelerators(preferAccelerators bool)                                                                                                        // procedure
	Browser() ICefBrowser                                                                                                                                 // property Browser Getter
	BrowserView() ICefBrowserView                                                                                                                         // property BrowserView Getter
	SetOnBrowserCreated(fn TOnBrowserCreatedEvent)                                                                                                        // property event
	SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent)                                                                                                    // property event
	SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewEvent)                                                                        // property event
	SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedEvent)                                                                                      // property event
	SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeEvent)                                                                                            // property event
	AsIntfBrowserViewDelegateEvents() uintptr
	AsIntfViewDelegateEvents() uintptr
}

type TCEFBrowserViewComponent struct {
	TCEFViewComponent
}

func (m *TCEFBrowserViewComponent) CreateBrowserView(client IEngClient, url string, settings TCefBrowserSettings, extraInfo ICefDictionaryValue, requestContext ICefRequestContext) bool {
	if !m.IsValid() {
		return false
	}
	settingsPtr := settings.ToPas()
	r := cEFBrowserViewComponentAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(client), api.PasStr(url), uintptr(base.UnsafePointer(settingsPtr)), base.GetObjectUintptr(extraInfo), base.GetObjectUintptr(requestContext))
	return api.GoBool(r)
}

func (m *TCEFBrowserViewComponent) GetForBrowser(browser ICefBrowser) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFBrowserViewComponentAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(browser))
	return api.GoBool(r)
}

func (m *TCEFBrowserViewComponent) SetPreferAccelerators(preferAccelerators bool) {
	if !m.IsValid() {
		return
	}
	cEFBrowserViewComponentAPI().SysCallN(3, m.Instance(), api.PasBool(preferAccelerators))
}

func (m *TCEFBrowserViewComponent) Browser() (result ICefBrowser) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFBrowserViewComponentAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserRef(resultPtr)
	return
}

func (m *TCEFBrowserViewComponent) BrowserView() (result ICefBrowserView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cEFBrowserViewComponentAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefBrowserViewRef(resultPtr)
	return
}

func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn TOnBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserCreatedEvent(fn)
	base.SetEvent(m, 6, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserDestroyedEvent(fn)
	base.SetEvent(m, 7, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDelegateForPopupBrowserViewEvent(fn)
	base.SetEvent(m, 8, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupBrowserViewCreatedEvent(fn)
	base.SetEvent(m, 9, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetChromeToolbarTypeEvent(fn)
	base.SetEvent(m, 10, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) AsIntfBrowserViewDelegateEvents() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCEFBrowserViewComponent) AsIntfViewDelegateEvents() uintptr {
	return m.GetIntfPointer(1)
}

// NewBrowserViewComponent class constructor
func NewBrowserViewComponent(owner lcl.IComponent) ICEFBrowserViewComponent {
	var browserViewDelegateEventsPtr uintptr // ICefBrowserViewDelegateEvents
	var viewDelegateEventsPtr uintptr        // ICefViewDelegateEvents
	r := cEFBrowserViewComponentAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&browserViewDelegateEventsPtr)), uintptr(base.UnsafePointer(&viewDelegateEventsPtr)))
	ret := AsCEFBrowserViewComponent(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, browserViewDelegateEventsPtr)
		intf.SetIntfPointer(1, viewDelegateEventsPtr)
	}
	return ret
}

var (
	cEFBrowserViewComponentOnce   base.Once
	cEFBrowserViewComponentImport *imports.Imports = nil
)

func cEFBrowserViewComponentAPI() *imports.Imports {
	cEFBrowserViewComponentOnce.Do(func() {
		cEFBrowserViewComponentImport = api.NewDefaultImports()
		cEFBrowserViewComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFBrowserViewComponent_Create", 0), // constructor NewBrowserViewComponent
			/* 1 */ imports.NewTable("TCEFBrowserViewComponent_CreateBrowserView", 0), // function CreateBrowserView
			/* 2 */ imports.NewTable("TCEFBrowserViewComponent_GetForBrowser", 0), // function GetForBrowser
			/* 3 */ imports.NewTable("TCEFBrowserViewComponent_SetPreferAccelerators", 0), // procedure SetPreferAccelerators
			/* 4 */ imports.NewTable("TCEFBrowserViewComponent_Browser", 0), // property Browser
			/* 5 */ imports.NewTable("TCEFBrowserViewComponent_BrowserView", 0), // property BrowserView
			/* 6 */ imports.NewTable("TCEFBrowserViewComponent_OnBrowserCreated", 0), // event OnBrowserCreated
			/* 7 */ imports.NewTable("TCEFBrowserViewComponent_OnBrowserDestroyed", 0), // event OnBrowserDestroyed
			/* 8 */ imports.NewTable("TCEFBrowserViewComponent_OnGetDelegateForPopupBrowserView", 0), // event OnGetDelegateForPopupBrowserView
			/* 9 */ imports.NewTable("TCEFBrowserViewComponent_OnPopupBrowserViewCreated", 0), // event OnPopupBrowserViewCreated
			/* 10 */ imports.NewTable("TCEFBrowserViewComponent_OnGetChromeToolbarType", 0), // event OnGetChromeToolbarType
		}
	})
	return cEFBrowserViewComponentImport
}
