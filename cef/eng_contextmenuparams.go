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
type ICefContextMenuParams interface {
	ICefBaseRefCounted
	GetXCoord() int32                                      // function
	GetYCoord() int32                                      // function
	GetTypeFlags() TCefContextMenuTypeFlags                // function
	GetLinkUrl() string                                    // function
	GetUnfilteredLinkUrl() string                          // function
	GetSourceUrl() string                                  // function
	HasImageContents() bool                                // function
	GetTitleText() string                                  // function
	GetPageUrl() string                                    // function
	GetFrameUrl() string                                   // function
	GetFrameCharset() string                               // function
	GetMediaType() TCefContextMenuMediaType                // function
	GetMediaStateFlags() TCefContextMenuMediaStateFlags    // function
	GetSelectionText() string                              // function
	GetMisspelledWord() string                             // function
	GetDictionarySuggestions(suggestions IStringList) bool // function
	IsEditable() bool                                      // function
	IsSpellCheckEnabled() bool                             // function
	GetEditStateFlags() TCefContextMenuEditStateFlags      // function
	IsCustomMenu() bool                                    // function
}

// TCefContextMenuParams Parent: TCefBaseRefCounted
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
