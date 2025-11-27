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

// IChromiumOptions Parent: lcl.IPersistent
type IChromiumOptions interface {
	lcl.IPersistent
	// Javascript
	//  Controls whether JavaScript can be executed. Also configurable using the
	//  "disable-javascript" command-line switch.
	Javascript() cefTypes.TCefState         // property Javascript Getter
	SetJavascript(value cefTypes.TCefState) // property Javascript Setter
	// JavascriptCloseWindows
	//  Controls whether JavaScript can be used to close windows that were not
	//  opened via JavaScript. JavaScript can still be used to close windows that
	//  were opened via JavaScript or that have no back/forward history. Also
	//  configurable using the "disable-javascript-close-windows" command-line
	//  switch.
	JavascriptCloseWindows() cefTypes.TCefState         // property JavascriptCloseWindows Getter
	SetJavascriptCloseWindows(value cefTypes.TCefState) // property JavascriptCloseWindows Setter
	// JavascriptAccessClipboard
	//  Controls whether JavaScript can access the clipboard. Also configurable
	//  using the "disable-javascript-access-clipboard" command-line switch.
	JavascriptAccessClipboard() cefTypes.TCefState         // property JavascriptAccessClipboard Getter
	SetJavascriptAccessClipboard(value cefTypes.TCefState) // property JavascriptAccessClipboard Setter
	// JavascriptDomPaste
	//  Controls whether DOM pasting is supported in the editor via
	//  execCommand("paste"). The |javascript_access_clipboard| setting must also
	//  be enabled. Also configurable using the "disable-javascript-dom-paste"
	//  command-line switch.
	JavascriptDomPaste() cefTypes.TCefState         // property JavascriptDomPaste Getter
	SetJavascriptDomPaste(value cefTypes.TCefState) // property JavascriptDomPaste Setter
	// ImageLoading
	//  Controls whether image URLs will be loaded from the network. A cached
	//  image will still be rendered if requested. Also configurable using the
	//  "disable-image-loading" command-line switch.
	ImageLoading() cefTypes.TCefState         // property ImageLoading Getter
	SetImageLoading(value cefTypes.TCefState) // property ImageLoading Setter
	// ImageShrinkStandaloneToFit
	//  Controls whether standalone images will be shrunk to fit the page. Also
	//  configurable using the "image-shrink-standalone-to-fit" command-line
	//  switch.
	ImageShrinkStandaloneToFit() cefTypes.TCefState         // property ImageShrinkStandaloneToFit Getter
	SetImageShrinkStandaloneToFit(value cefTypes.TCefState) // property ImageShrinkStandaloneToFit Setter
	// TextAreaResize
	//  Controls whether text areas can be resized. Also configurable using the
	//  "disable-text-area-resize" command-line switch.
	TextAreaResize() cefTypes.TCefState         // property TextAreaResize Getter
	SetTextAreaResize(value cefTypes.TCefState) // property TextAreaResize Setter
	// TabToLinks
	//  Controls whether the tab key can advance focus to links. Also configurable
	//  using the "disable-tab-to-links" command-line switch.
	TabToLinks() cefTypes.TCefState         // property TabToLinks Getter
	SetTabToLinks(value cefTypes.TCefState) // property TabToLinks Setter
	// LocalStorage
	//  Controls whether local storage can be used. Also configurable using the
	//  "disable-local-storage" command-line switch.
	LocalStorage() cefTypes.TCefState         // property LocalStorage Getter
	SetLocalStorage(value cefTypes.TCefState) // property LocalStorage Setter
	// Databases
	//  Controls whether databases can be used. Also configurable using the
	//  "disable-databases" command-line switch.
	Databases() cefTypes.TCefState         // property Databases Getter
	SetDatabases(value cefTypes.TCefState) // property Databases Setter
	// Webgl
	//  Controls whether WebGL can be used. Note that WebGL requires hardware
	//  support and may not work on all systems even when enabled. Also
	//  configurable using the "disable-webgl" command-line switch.
	Webgl() cefTypes.TCefState         // property Webgl Getter
	SetWebgl(value cefTypes.TCefState) // property Webgl Setter
	// BackgroundColor
	//  Background color used for the browser before a document is loaded and when
	//  no document color is specified. The alpha component must be either fully
	//  opaque (0xFF) or fully transparent (0x00). If the alpha component is fully
	//  opaque then the RGB components will be used as the background color. If
	//  the alpha component is fully transparent for a windowed browser then the
	//  TCefSettings.background_color value will be used. If the alpha component is
	//  fully transparent for a windowless (off-screen) browser then transparent
	//  painting will be enabled.
	BackgroundColor() cefTypes.TCefColor         // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor) // property BackgroundColor Setter
	// WindowlessFrameRate
	//  The maximum rate in frames per second (fps) that ICefRenderHandler.OnPaint
	//  will be called for a windowless browser. The actual fps may be lower if
	//  the browser cannot generate frames at the requested rate. The minimum
	//  value is 1 and the maximum value is 60 (default 30). This value can also
	//  be changed dynamically via ICefBrowserHost.SetWindowlessFrameRate.
	//  Use CEF_OSR_SHARED_TEXTURES_FRAMERATE_DEFAULT as default value if the shared textures are enabled.
	//  Use CEF_OSR_FRAMERATE_DEFAULT as default value if the shared textures are disabled.
	WindowlessFrameRate() int32         // property WindowlessFrameRate Getter
	SetWindowlessFrameRate(value int32) // property WindowlessFrameRate Setter
	// ChromeStatusBubble
	//  Controls whether the Chrome status bubble will be used. Only supported
	//  with the Chrome runtime. For details about the status bubble see
	//  https://www.chromium.org/user-experience/status-bubble/
	ChromeStatusBubble() cefTypes.TCefState         // property ChromeStatusBubble Getter
	SetChromeStatusBubble(value cefTypes.TCefState) // property ChromeStatusBubble Setter
	// ChromeZoomBubble
	//  Controls whether the Chrome zoom bubble will be shown when zooming. Only
	//  supported with the Chrome runtime.
	ChromeZoomBubble() cefTypes.TCefState         // property ChromeZoomBubble Getter
	SetChromeZoomBubble(value cefTypes.TCefState) // property ChromeZoomBubble Setter
}

type TChromiumOptions struct {
	lcl.TPersistent
}

func (m *TChromiumOptions) Javascript() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(1, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetJavascript(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) JavascriptCloseWindows() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(2, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetJavascriptCloseWindows(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) JavascriptAccessClipboard() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(3, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetJavascriptAccessClipboard(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) JavascriptDomPaste() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(4, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetJavascriptDomPaste(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) ImageLoading() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(5, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetImageLoading(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) ImageShrinkStandaloneToFit() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(6, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetImageShrinkStandaloneToFit(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) TextAreaResize() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(7, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetTextAreaResize(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) TabToLinks() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(8, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetTabToLinks(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) LocalStorage() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(9, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetLocalStorage(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) Databases() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(10, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetDatabases(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) Webgl() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(11, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetWebgl(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) BackgroundColor() cefTypes.TCefColor {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(12, 0, m.Instance())
	return cefTypes.TCefColor(r)
}

func (m *TChromiumOptions) SetBackgroundColor(value cefTypes.TCefColor) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) WindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumOptions) SetWindowlessFrameRate(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) ChromeStatusBubble() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(14, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetChromeStatusBubble(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) ChromeZoomBubble() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(15, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetChromeZoomBubble(value cefTypes.TCefState) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

// NewChromiumOptions class constructor
func NewChromiumOptions() IChromiumOptions {
	r := chromiumOptionsAPI().SysCallN(0)
	return AsChromiumOptions(r)
}

var (
	chromiumOptionsOnce   base.Once
	chromiumOptionsImport *imports.Imports = nil
)

func chromiumOptionsAPI() *imports.Imports {
	chromiumOptionsOnce.Do(func() {
		chromiumOptionsImport = api.NewDefaultImports()
		chromiumOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChromiumOptions_Create", 0), // constructor NewChromiumOptions
			/* 1 */ imports.NewTable("TChromiumOptions_Javascript", 0), // property Javascript
			/* 2 */ imports.NewTable("TChromiumOptions_JavascriptCloseWindows", 0), // property JavascriptCloseWindows
			/* 3 */ imports.NewTable("TChromiumOptions_JavascriptAccessClipboard", 0), // property JavascriptAccessClipboard
			/* 4 */ imports.NewTable("TChromiumOptions_JavascriptDomPaste", 0), // property JavascriptDomPaste
			/* 5 */ imports.NewTable("TChromiumOptions_ImageLoading", 0), // property ImageLoading
			/* 6 */ imports.NewTable("TChromiumOptions_ImageShrinkStandaloneToFit", 0), // property ImageShrinkStandaloneToFit
			/* 7 */ imports.NewTable("TChromiumOptions_TextAreaResize", 0), // property TextAreaResize
			/* 8 */ imports.NewTable("TChromiumOptions_TabToLinks", 0), // property TabToLinks
			/* 9 */ imports.NewTable("TChromiumOptions_LocalStorage", 0), // property LocalStorage
			/* 10 */ imports.NewTable("TChromiumOptions_Databases", 0), // property Databases
			/* 11 */ imports.NewTable("TChromiumOptions_Webgl", 0), // property Webgl
			/* 12 */ imports.NewTable("TChromiumOptions_BackgroundColor", 0), // property BackgroundColor
			/* 13 */ imports.NewTable("TChromiumOptions_WindowlessFrameRate", 0), // property WindowlessFrameRate
			/* 14 */ imports.NewTable("TChromiumOptions_ChromeStatusBubble", 0), // property ChromeStatusBubble
			/* 15 */ imports.NewTable("TChromiumOptions_ChromeZoomBubble", 0), // property ChromeZoomBubble
		}
	})
	return chromiumOptionsImport
}
