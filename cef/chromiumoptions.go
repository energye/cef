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

	cefTypes "github.com/energye/cef/types"
)

// IChromiumOptions Parent: IPersistent
type IChromiumOptions interface {
	IPersistent
	Javascript() cefTypes.TCefState                         // property Javascript Getter
	SetJavascript(value cefTypes.TCefState)                 // property Javascript Setter
	JavascriptCloseWindows() cefTypes.TCefState             // property JavascriptCloseWindows Getter
	SetJavascriptCloseWindows(value cefTypes.TCefState)     // property JavascriptCloseWindows Setter
	JavascriptAccessClipboard() cefTypes.TCefState          // property JavascriptAccessClipboard Getter
	SetJavascriptAccessClipboard(value cefTypes.TCefState)  // property JavascriptAccessClipboard Setter
	JavascriptDomPaste() cefTypes.TCefState                 // property JavascriptDomPaste Getter
	SetJavascriptDomPaste(value cefTypes.TCefState)         // property JavascriptDomPaste Setter
	ImageLoading() cefTypes.TCefState                       // property ImageLoading Getter
	SetImageLoading(value cefTypes.TCefState)               // property ImageLoading Setter
	ImageShrinkStandaloneToFit() cefTypes.TCefState         // property ImageShrinkStandaloneToFit Getter
	SetImageShrinkStandaloneToFit(value cefTypes.TCefState) // property ImageShrinkStandaloneToFit Setter
	TextAreaResize() cefTypes.TCefState                     // property TextAreaResize Getter
	SetTextAreaResize(value cefTypes.TCefState)             // property TextAreaResize Setter
	TabToLinks() cefTypes.TCefState                         // property TabToLinks Getter
	SetTabToLinks(value cefTypes.TCefState)                 // property TabToLinks Setter
	LocalStorage() cefTypes.TCefState                       // property LocalStorage Getter
	SetLocalStorage(value cefTypes.TCefState)               // property LocalStorage Setter
	Databases() cefTypes.TCefState                          // property Databases Getter
	SetDatabases(value cefTypes.TCefState)                  // property Databases Setter
	Webgl() cefTypes.TCefState                              // property Webgl Getter
	SetWebgl(value cefTypes.TCefState)                      // property Webgl Setter
	BackgroundColor() cefTypes.TCefColor                    // property BackgroundColor Getter
	SetBackgroundColor(value cefTypes.TCefColor)            // property BackgroundColor Setter
	AcceptLanguageList() string                             // property AcceptLanguageList Getter
	SetAcceptLanguageList(value string)                     // property AcceptLanguageList Setter
	WindowlessFrameRate() int32                             // property WindowlessFrameRate Getter
	SetWindowlessFrameRate(value int32)                     // property WindowlessFrameRate Setter
	ChromeStatusBubble() cefTypes.TCefState                 // property ChromeStatusBubble Getter
	SetChromeStatusBubble(value cefTypes.TCefState)         // property ChromeStatusBubble Setter
}

type TChromiumOptions struct {
	TPersistent
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

func (m *TChromiumOptions) AcceptLanguageList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	chromiumOptionsAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TChromiumOptions) SetAcceptLanguageList(value string) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TChromiumOptions) WindowlessFrameRate() int32 {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TChromiumOptions) SetWindowlessFrameRate(value int32) {
	if !m.IsValid() {
		return
	}
	chromiumOptionsAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TChromiumOptions) ChromeStatusBubble() cefTypes.TCefState {
	if !m.IsValid() {
		return 0
	}
	r := chromiumOptionsAPI().SysCallN(15, 0, m.Instance())
	return cefTypes.TCefState(r)
}

func (m *TChromiumOptions) SetChromeStatusBubble(value cefTypes.TCefState) {
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
			/* 13 */ imports.NewTable("TChromiumOptions_AcceptLanguageList", 0), // property AcceptLanguageList
			/* 14 */ imports.NewTable("TChromiumOptions_WindowlessFrameRate", 0), // property WindowlessFrameRate
			/* 15 */ imports.NewTable("TChromiumOptions_ChromeStatusBubble", 0), // property ChromeStatusBubble
		}
	})
	return chromiumOptionsImport
}
