//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IChromiumOptions Parent: IPersistent
//
//	The TChromiumOptions properties used to fill the TCefBrowserSettings record which is used during the browser creation.
type IChromiumOptions interface {
	IPersistent
	// Javascript
	//  Controls whether JavaScript can be executed. Also configurable using the
	//  "disable-javascript" command-line switch.
	Javascript() TCefState // property
	// SetJavascript Set Javascript
	SetJavascript(AValue TCefState) // property
	// JavascriptCloseWindows
	//  Controls whether JavaScript can be used to close windows that were not
	//  opened via JavaScript. JavaScript can still be used to close windows that
	//  were opened via JavaScript or that have no back/forward history. Also
	//  configurable using the "disable-javascript-close-windows" command-line
	//  switch.
	JavascriptCloseWindows() TCefState // property
	// SetJavascriptCloseWindows Set JavascriptCloseWindows
	SetJavascriptCloseWindows(AValue TCefState) // property
	// JavascriptAccessClipboard
	//  Controls whether JavaScript can access the clipboard. Also configurable
	//  using the "disable-javascript-access-clipboard" command-line switch.
	JavascriptAccessClipboard() TCefState // property
	// SetJavascriptAccessClipboard Set JavascriptAccessClipboard
	SetJavascriptAccessClipboard(AValue TCefState) // property
	// JavascriptDomPaste
	//  Controls whether DOM pasting is supported in the editor via
	//  execCommand("paste"). The |javascript_access_clipboard| setting must also
	//  be enabled. Also configurable using the "disable-javascript-dom-paste"
	//  command-line switch.
	JavascriptDomPaste() TCefState // property
	// SetJavascriptDomPaste Set JavascriptDomPaste
	SetJavascriptDomPaste(AValue TCefState) // property
	// ImageLoading
	//  Controls whether image URLs will be loaded from the network. A cached
	//  image will still be rendered if requested. Also configurable using the
	//  "disable-image-loading" command-line switch.
	ImageLoading() TCefState // property
	// SetImageLoading Set ImageLoading
	SetImageLoading(AValue TCefState) // property
	// ImageShrinkStandaloneToFit
	//  Controls whether standalone images will be shrunk to fit the page. Also
	//  configurable using the "image-shrink-standalone-to-fit" command-line
	//  switch.
	ImageShrinkStandaloneToFit() TCefState // property
	// SetImageShrinkStandaloneToFit Set ImageShrinkStandaloneToFit
	SetImageShrinkStandaloneToFit(AValue TCefState) // property
	// TextAreaResize
	//  Controls whether text areas can be resized. Also configurable using the
	//  "disable-text-area-resize" command-line switch.
	TextAreaResize() TCefState // property
	// SetTextAreaResize Set TextAreaResize
	SetTextAreaResize(AValue TCefState) // property
	// TabToLinks
	//  Controls whether the tab key can advance focus to links. Also configurable
	//  using the "disable-tab-to-links" command-line switch.
	TabToLinks() TCefState // property
	// SetTabToLinks Set TabToLinks
	SetTabToLinks(AValue TCefState) // property
	// LocalStorage
	//  Controls whether local storage can be used. Also configurable using the
	//  "disable-local-storage" command-line switch.
	LocalStorage() TCefState // property
	// SetLocalStorage Set LocalStorage
	SetLocalStorage(AValue TCefState) // property
	// Databases
	//  Controls whether databases can be used. Also configurable using the
	//  "disable-databases" command-line switch.
	Databases() TCefState // property
	// SetDatabases Set Databases
	SetDatabases(AValue TCefState) // property
	// Webgl
	//  Controls whether WebGL can be used. Note that WebGL requires hardware
	//  support and may not work on all systems even when enabled. Also
	//  configurable using the "disable-webgl" command-line switch.
	Webgl() TCefState // property
	// SetWebgl Set Webgl
	SetWebgl(AValue TCefState) // property
	// BackgroundColor
	//  Background color used for the browser before a document is loaded and when
	//  no document color is specified. The alpha component must be either fully
	//  opaque(0xFF) or fully transparent(0x00). If the alpha component is fully
	//  opaque then the RGB components will be used as the background color. If
	//  the alpha component is fully transparent for a windowed browser then the
	//  TCefSettings.background_color value will be used. If the alpha component is
	//  fully transparent for a windowless(off-screen) browser then transparent
	//  painting will be enabled.
	BackgroundColor() TCefColor // property
	// SetBackgroundColor Set BackgroundColor
	SetBackgroundColor(AValue TCefColor) // property
	// WindowlessFrameRate
	//  The maximum rate in frames per second(fps) that ICefRenderHandler.OnPaint
	//  will be called for a windowless browser. The actual fps may be lower if
	//  the browser cannot generate frames at the requested rate. The minimum
	//  value is 1 and the maximum value is 60(default 30). This value can also
	//  be changed dynamically via ICefBrowserHost.SetWindowlessFrameRate.
	//  Use CEF_OSR_SHARED_TEXTURES_FRAMERATE_DEFAULT as default value if the shared textures are enabled.
	//  Use CEF_OSR_FRAMERATE_DEFAULT as default value if the shared textures are disabled.
	WindowlessFrameRate() int32 // property
	// SetWindowlessFrameRate Set WindowlessFrameRate
	SetWindowlessFrameRate(AValue int32) // property
	// ChromeStatusBubble
	//  Controls whether the Chrome status bubble will be used. Only supported
	//  with the Chrome runtime. For details about the status bubble see
	//  https://www.chromium.org/user-experience/status-bubble/
	ChromeStatusBubble() TCefState // property
	// SetChromeStatusBubble Set ChromeStatusBubble
	SetChromeStatusBubble(AValue TCefState) // property
	// ChromeZoomBubble
	//  Controls whether the Chrome zoom bubble will be shown when zooming. Only
	//  supported with the Chrome runtime.
	ChromeZoomBubble() TCefState // property
	// SetChromeZoomBubble Set ChromeZoomBubble
	SetChromeZoomBubble(AValue TCefState) // property
}

// TChromiumOptions Parent: TPersistent
//
//	The TChromiumOptions properties used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumOptions struct {
	TPersistent
}

func NewChromiumOptions() IChromiumOptions {
	r1 := chromiumOptionsImportAPI().SysCallN(3)
	return AsChromiumOptions(r1)
}

func (m *TChromiumOptions) Javascript() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascript(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptCloseWindows() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptCloseWindows(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptAccessClipboard() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptAccessClipboard(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptDomPaste() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptDomPaste(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ImageLoading() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetImageLoading(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ImageShrinkStandaloneToFit() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetImageShrinkStandaloneToFit(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) TextAreaResize() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetTextAreaResize(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) TabToLinks() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetTabToLinks(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) LocalStorage() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetLocalStorage(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) Databases() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetDatabases(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) Webgl() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetWebgl(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) BackgroundColor() TCefColor {
	r1 := chromiumOptionsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TChromiumOptions) SetBackgroundColor(AValue TCefColor) {
	chromiumOptionsImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) WindowlessFrameRate() int32 {
	r1 := chromiumOptionsImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumOptions) SetWindowlessFrameRate(AValue int32) {
	chromiumOptionsImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ChromeStatusBubble() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetChromeStatusBubble(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ChromeZoomBubble() TCefState {
	r1 := chromiumOptionsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetChromeZoomBubble(AValue TCefState) {
	chromiumOptionsImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

var (
	chromiumOptionsImport       *imports.Imports = nil
	chromiumOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ChromiumOptions_BackgroundColor", 0),
		/*1*/ imports.NewTable("ChromiumOptions_ChromeStatusBubble", 0),
		/*2*/ imports.NewTable("ChromiumOptions_ChromeZoomBubble", 0),
		/*3*/ imports.NewTable("ChromiumOptions_Create", 0),
		/*4*/ imports.NewTable("ChromiumOptions_Databases", 0),
		/*5*/ imports.NewTable("ChromiumOptions_ImageLoading", 0),
		/*6*/ imports.NewTable("ChromiumOptions_ImageShrinkStandaloneToFit", 0),
		/*7*/ imports.NewTable("ChromiumOptions_Javascript", 0),
		/*8*/ imports.NewTable("ChromiumOptions_JavascriptAccessClipboard", 0),
		/*9*/ imports.NewTable("ChromiumOptions_JavascriptCloseWindows", 0),
		/*10*/ imports.NewTable("ChromiumOptions_JavascriptDomPaste", 0),
		/*11*/ imports.NewTable("ChromiumOptions_LocalStorage", 0),
		/*12*/ imports.NewTable("ChromiumOptions_TabToLinks", 0),
		/*13*/ imports.NewTable("ChromiumOptions_TextAreaResize", 0),
		/*14*/ imports.NewTable("ChromiumOptions_Webgl", 0),
		/*15*/ imports.NewTable("ChromiumOptions_WindowlessFrameRate", 0),
	}
)

func chromiumOptionsImportAPI() *imports.Imports {
	if chromiumOptionsImport == nil {
		chromiumOptionsImport = NewDefaultImports()
		chromiumOptionsImport.SetImportTable(chromiumOptionsImportTables)
		chromiumOptionsImportTables = nil
	}
	return chromiumOptionsImport
}
