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

// ICefContextMenuParams Parent: ICefBaseRefCounted
//
//	Provides information about the context menu state. The functions of this interface can only be accessed on browser process the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_context_menu_params_t))</a>
type ICefContextMenuParams interface {
	ICefBaseRefCounted
	// GetXCoord
	//  Returns the X coordinate of the mouse where the context menu was invoked. Coords are relative to the associated RenderView's origin.
	GetXCoord() int32 // function
	// GetYCoord
	//  Returns the Y coordinate of the mouse where the context menu was invoked. Coords are relative to the associated RenderView's origin.
	GetYCoord() int32 // function
	// GetTypeFlags
	//  Returns flags representing the type of node that the context menu was invoked on.
	GetTypeFlags() TCefContextMenuTypeFlags // function
	// GetLinkUrl
	//  Returns the URL of the link, if any, that encloses the node that the context menu was invoked on.
	GetLinkUrl() string // function
	// GetUnfilteredLinkUrl
	//  Returns the link URL, if any, to be used ONLY for "copy link address". We don't validate this field in the frontend process.
	GetUnfilteredLinkUrl() string // function
	// GetSourceUrl
	//  Returns the source URL, if any, for the element that the context menu was invoked on. Example of elements with source URLs are img, audio, and video.
	GetSourceUrl() string // function
	// HasImageContents
	//  Returns true (1) if the context menu was invoked on an image which has non-NULL contents.
	HasImageContents() bool // function
	// GetTitleText
	//  Returns the title text or the alt text if the context menu was invoked on an image.
	GetTitleText() string // function
	// GetPageUrl
	//  Returns the URL of the top level page that the context menu was invoked on.
	GetPageUrl() string // function
	// GetFrameUrl
	//  Returns the URL of the subframe that the context menu was invoked on.
	GetFrameUrl() string // function
	// GetFrameCharset
	//  Returns the character encoding of the subframe that the context menu was invoked on.
	GetFrameCharset() string // function
	// GetMediaType
	//  Returns the type of context node that the context menu was invoked on.
	GetMediaType() TCefContextMenuMediaType // function
	// GetMediaStateFlags
	//  Returns flags representing the actions supported by the media element, if any, that the context menu was invoked on.
	GetMediaStateFlags() TCefContextMenuMediaStateFlags // function
	// GetSelectionText
	//  Returns the text of the selection, if any, that the context menu was invoked on.
	GetSelectionText() string // function
	// GetMisspelledWord
	//  Returns the text of the misspelled word, if any, that the context menu was invoked on.
	GetMisspelledWord() string // function
	// GetDictionarySuggestions
	//  Returns true (1) if suggestions exist, false (0) otherwise. Fills in |suggestions| from the spell check service for the misspelled word if there is one.
	GetDictionarySuggestions(suggestions IStringList) bool // function
	// IsEditable
	//  Returns true (1) if the context menu was invoked on an editable node.
	IsEditable() bool // function
	// IsSpellCheckEnabled
	//  Returns true (1) if the context menu was invoked on an editable node where spell-check is enabled.
	IsSpellCheckEnabled() bool // function
	// GetEditStateFlags
	//  Returns flags representing the actions supported by the editable node, if any, that the context menu was invoked on.
	GetEditStateFlags() TCefContextMenuEditStateFlags // function
	// IsCustomMenu
	//  Returns true (1) if the context menu contains items specified by the renderer process.
	IsCustomMenu() bool // function
}

// TCefContextMenuParams Parent: TCefBaseRefCounted
//
//	Provides information about the context menu state. The functions of this interface can only be accessed on browser process the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_context_menu_params_t))</a>
type TCefContextMenuParams struct {
	TCefBaseRefCounted
}

// ContextMenuParamsRef -> ICefContextMenuParams
var ContextMenuParamsRef contextMenuParams

// contextMenuParams TCefContextMenuParams Ref
type contextMenuParams uintptr

func (m *contextMenuParams) UnWrap(data uintptr) ICefContextMenuParams {
	var resultCefContextMenuParams uintptr
	contextMenuParamsImportAPI().SysCallN(20, uintptr(data), uintptr(unsafePointer(&resultCefContextMenuParams)))
	return AsCefContextMenuParams(resultCefContextMenuParams)
}

func (m *TCefContextMenuParams) GetXCoord() int32 {
	r1 := contextMenuParamsImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TCefContextMenuParams) GetYCoord() int32 {
	r1 := contextMenuParamsImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TCefContextMenuParams) GetTypeFlags() TCefContextMenuTypeFlags {
	r1 := contextMenuParamsImportAPI().SysCallN(12, m.Instance())
	return TCefContextMenuTypeFlags(r1)
}

func (m *TCefContextMenuParams) GetLinkUrl() string {
	r1 := contextMenuParamsImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetUnfilteredLinkUrl() string {
	r1 := contextMenuParamsImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetSourceUrl() string {
	r1 := contextMenuParamsImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) HasImageContents() bool {
	r1 := contextMenuParamsImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TCefContextMenuParams) GetTitleText() string {
	r1 := contextMenuParamsImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetPageUrl() string {
	r1 := contextMenuParamsImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetFrameUrl() string {
	r1 := contextMenuParamsImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetFrameCharset() string {
	r1 := contextMenuParamsImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetMediaType() TCefContextMenuMediaType {
	r1 := contextMenuParamsImportAPI().SysCallN(6, m.Instance())
	return TCefContextMenuMediaType(r1)
}

func (m *TCefContextMenuParams) GetMediaStateFlags() TCefContextMenuMediaStateFlags {
	r1 := contextMenuParamsImportAPI().SysCallN(5, m.Instance())
	return TCefContextMenuMediaStateFlags(r1)
}

func (m *TCefContextMenuParams) GetSelectionText() string {
	r1 := contextMenuParamsImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetMisspelledWord() string {
	r1 := contextMenuParamsImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCefContextMenuParams) GetDictionarySuggestions(suggestions IStringList) bool {
	r1 := contextMenuParamsImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(suggestions))
	return GoBool(r1)
}

func (m *TCefContextMenuParams) IsEditable() bool {
	r1 := contextMenuParamsImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCefContextMenuParams) IsSpellCheckEnabled() bool {
	r1 := contextMenuParamsImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCefContextMenuParams) GetEditStateFlags() TCefContextMenuEditStateFlags {
	r1 := contextMenuParamsImportAPI().SysCallN(1, m.Instance())
	return TCefContextMenuEditStateFlags(r1)
}

func (m *TCefContextMenuParams) IsCustomMenu() bool {
	r1 := contextMenuParamsImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

var (
	contextMenuParamsImport       *imports.Imports = nil
	contextMenuParamsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefContextMenuParams_GetDictionarySuggestions", 0),
		/*1*/ imports.NewTable("CefContextMenuParams_GetEditStateFlags", 0),
		/*2*/ imports.NewTable("CefContextMenuParams_GetFrameCharset", 0),
		/*3*/ imports.NewTable("CefContextMenuParams_GetFrameUrl", 0),
		/*4*/ imports.NewTable("CefContextMenuParams_GetLinkUrl", 0),
		/*5*/ imports.NewTable("CefContextMenuParams_GetMediaStateFlags", 0),
		/*6*/ imports.NewTable("CefContextMenuParams_GetMediaType", 0),
		/*7*/ imports.NewTable("CefContextMenuParams_GetMisspelledWord", 0),
		/*8*/ imports.NewTable("CefContextMenuParams_GetPageUrl", 0),
		/*9*/ imports.NewTable("CefContextMenuParams_GetSelectionText", 0),
		/*10*/ imports.NewTable("CefContextMenuParams_GetSourceUrl", 0),
		/*11*/ imports.NewTable("CefContextMenuParams_GetTitleText", 0),
		/*12*/ imports.NewTable("CefContextMenuParams_GetTypeFlags", 0),
		/*13*/ imports.NewTable("CefContextMenuParams_GetUnfilteredLinkUrl", 0),
		/*14*/ imports.NewTable("CefContextMenuParams_GetXCoord", 0),
		/*15*/ imports.NewTable("CefContextMenuParams_GetYCoord", 0),
		/*16*/ imports.NewTable("CefContextMenuParams_HasImageContents", 0),
		/*17*/ imports.NewTable("CefContextMenuParams_IsCustomMenu", 0),
		/*18*/ imports.NewTable("CefContextMenuParams_IsEditable", 0),
		/*19*/ imports.NewTable("CefContextMenuParams_IsSpellCheckEnabled", 0),
		/*20*/ imports.NewTable("CefContextMenuParams_UnWrap", 0),
	}
)

func contextMenuParamsImportAPI() *imports.Imports {
	if contextMenuParamsImport == nil {
		contextMenuParamsImport = NewDefaultImports()
		contextMenuParamsImport.SetImportTable(contextMenuParamsImportTables)
		contextMenuParamsImportTables = nil
	}
	return contextMenuParamsImport
}
