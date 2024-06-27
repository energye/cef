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
)

type IChromiumEvents = IChromium

// IMESetComposition
//
//	Begins a new composition or updates the existing composition. Blink has a
//	special node (a composition node) that allows the input function to change
//	text without affecting other DOM nodes. |text| is the optional text that
//	will be inserted into the composition node. |underlines| is an optional
//	set of ranges that will be underlined in the resulting text.
//	|replacement_range| is an optional range of the existing text that will
//	be replaced. |selection_range| is an optional range of the resulting text
//	that will be selected after insertion or replacement. The |replacement_range| value is only used on OS X.
//	This function may be called multiple times as the composition changes.
//	When the client is done making changes the composition should either be
//	canceled or completed. To cancel the composition call
//	ImeCancelComposition. To complete the composition call either
//	ImeCommitText or ImeFinishComposingText. Completion is usually signaled
//	when:
//	 1. The client receives a WM_IME_COMPOSITION message with a GCS_RESULTSTR flag (on Windows), or
//	 2. The client receives a "commit" signal of GtkIMContext (on Linux), or
//	 3. insertText of NSTextInput is called (on Mac).
//	This function is only used when window rendering is disabled.
func (m *TChromiumCore) IMESetComposition(text string, underlines TCefCompositionUnderlineDynArray, replacementrange, selectionrange *TCefRange) {
	predefImportAPI().SysCallN(4, m.Instance(), api.PascalStr(text), uintptr(unsafePointer(&underlines[0])), uintptr(int32(len(underlines))),
		uintptr(unsafePointer(replacementrange)), uintptr(unsafePointer(selectionrange)))
}

// SetCookie
//
//	TChromiumCore.SetCookie triggers the TChromiumCore.OnCookieSet event when the cookie has been set
//	aID is an optional parameter to identify which SetCookie call has triggered the
//	OnCookieSet event.
func (m *TChromiumCore) SetCookie(url string, aSetImmediately bool, aID int32, cookie TCookie) bool {
	inCookie := cookie.Pointer()
	r1 := predefImportAPI().SysCallN(9, m.Instance(), api.PascalStr(url), api.PascalBool(aSetImmediately), uintptr(aID), uintptr(unsafePointer(inCookie)))
	return api.GoBool(r1)
}

func (m *TChromiumCore) AsChromiumEvents() IChromiumEvents {
	var result uintptr
	predefChromiumImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&result)))
	return AsChromium(result)
}

var (
	predefChromiumImport       *imports.Imports = nil
	predefChromiumImportTables                  = []*imports.Table{
		imports.NewTable("ChromiumCore_AsChromiumEvents", 0),
	}
)

func predefChromiumImportAPI() *imports.Imports {
	if predefChromiumImport == nil {
		predefChromiumImport = api.NewDefaultImports()
		predefChromiumImport.SetImportTable(predefChromiumImportTables)
		predefChromiumImportTables = nil
	}
	return predefChromiumImport
}
