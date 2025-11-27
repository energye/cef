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

	cefTypes "github.com/energye/cef/types"
)

// ICEFBrowserViewComponent Parent: ICEFViewComponent
type ICEFBrowserViewComponent interface {
	ICEFViewComponent
	// CreateBrowserView
	//  Create a new ICefBrowserView. The underlying ICefBrowser will not be created
	//  until this view is added to the views hierarchy. The optional |extra_info|
	//  parameter provides an opportunity to specify extra information specific to
	//  the created browser that will be passed to
	//  ICefRenderProcessHandler.OnBrowserCreated in the render process.
	CreateBrowserView(client IEngClient, url string, settings TCefBrowserSettings, extraInfo ICefDictionaryValue, requestContext ICefRequestContext) bool // function
	// GetForBrowser
	//  Updates the internal ICefBrowserView with the ICefBrowserView associated with |browser|.
	GetForBrowser(browser ICefBrowser) bool // function
	// SetPreferAccelerators
	//  Sets whether accelerators registered with ICefWindow.SetAccelerator are
	//  triggered before or after the event is sent to the ICefBrowser. If
	//  |prefer_accelerators| is true (1) then the matching accelerator will be
	//  triggered immediately and the event will not be sent to the ICefBrowser.
	//  If |prefer_accelerators| is false (0) then the matching accelerator will
	//  only be triggered if the event is not handled by web content or by
	//  ICefKeyboardHandler. The default value is false (0).
	SetPreferAccelerators(preferAccelerators bool) // procedure
	// Browser
	//  Returns the ICefBrowser hosted by this BrowserView. Will return NULL if
	//  the browser has not yet been created or has already been destroyed.
	Browser() ICefBrowser // property Browser Getter
	// BrowserView
	//  ICefBrowserView assiciated to this component.
	BrowserView() ICefBrowserView // property BrowserView Getter
	// RuntimeStyle
	//  Returns the runtime style for this BrowserView (ALLOY or CHROME). See
	//  TCefRuntimeStyle documentation for details.
	RuntimeStyle() cefTypes.TCefRuntimeStyle                                                // property RuntimeStyle Getter
	SetOnBrowserCreated(fn TOnBrowserCreatedEvent)                                          // property event
	SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent)                                      // property event
	SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewEvent)          // property event
	SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedEvent)                        // property event
	SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeEvent)                              // property event
	SetOnUseFramelessWindowForPictureInPicture(fn TOnUseFramelessWindowForPictureInPicture) // property event
	SetOnGestureCommand(fn TOnGestureCommandEvent)                                          // property event
	SetOnGetBrowserRuntimeStyle(fn TOnGetBrowserRuntimeStyleEvent)                          // property event
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

func (m *TCEFBrowserViewComponent) RuntimeStyle() cefTypes.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r := cEFBrowserViewComponentAPI().SysCallN(6, m.Instance())
	return cefTypes.TCefRuntimeStyle(r)
}

func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn TOnBrowserCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserCreatedEvent(fn)
	base.SetEvent(m, 7, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn TOnBrowserDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserDestroyedEvent(fn)
	base.SetEvent(m, 8, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetDelegateForPopupBrowserViewEvent(fn)
	base.SetEvent(m, 9, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPopupBrowserViewCreatedEvent(fn)
	base.SetEvent(m, 10, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetChromeToolbarTypeEvent(fn)
	base.SetEvent(m, 11, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnUseFramelessWindowForPictureInPicture(fn TOnUseFramelessWindowForPictureInPicture) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUseFramelessWindowForPictureInPicture(fn)
	base.SetEvent(m, 12, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGestureCommand(fn TOnGestureCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGestureCommandEvent(fn)
	base.SetEvent(m, 13, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCEFBrowserViewComponent) SetOnGetBrowserRuntimeStyle(fn TOnGetBrowserRuntimeStyleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetBrowserRuntimeStyleEvent(fn)
	base.SetEvent(m, 14, cEFBrowserViewComponentAPI(), api.MakeEventDataPtr(cb))
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
			/* 6 */ imports.NewTable("TCEFBrowserViewComponent_RuntimeStyle", 0), // property RuntimeStyle
			/* 7 */ imports.NewTable("TCEFBrowserViewComponent_OnBrowserCreated", 0), // event OnBrowserCreated
			/* 8 */ imports.NewTable("TCEFBrowserViewComponent_OnBrowserDestroyed", 0), // event OnBrowserDestroyed
			/* 9 */ imports.NewTable("TCEFBrowserViewComponent_OnGetDelegateForPopupBrowserView", 0), // event OnGetDelegateForPopupBrowserView
			/* 10 */ imports.NewTable("TCEFBrowserViewComponent_OnPopupBrowserViewCreated", 0), // event OnPopupBrowserViewCreated
			/* 11 */ imports.NewTable("TCEFBrowserViewComponent_OnGetChromeToolbarType", 0), // event OnGetChromeToolbarType
			/* 12 */ imports.NewTable("TCEFBrowserViewComponent_OnUseFramelessWindowForPictureInPicture", 0), // event OnUseFramelessWindowForPictureInPicture
			/* 13 */ imports.NewTable("TCEFBrowserViewComponent_OnGestureCommand", 0), // event OnGestureCommand
			/* 14 */ imports.NewTable("TCEFBrowserViewComponent_OnGetBrowserRuntimeStyle", 0), // event OnGetBrowserRuntimeStyle
		}
	})
	return cEFBrowserViewComponentImport
}
