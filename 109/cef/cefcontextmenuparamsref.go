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

	cefTypes "github.com/energye/cef/109/types"
)

// ICefContextMenuParams Parent: ICefBaseRefCounted
type ICefContextMenuParams interface {
	ICefBaseRefCounted
	GetXCoord() int32                                            // function
	GetYCoord() int32                                            // function
	GetTypeFlags() cefTypes.TCefContextMenuTypeFlags             // function
	GetLinkUrl() string                                          // function
	GetUnfilteredLinkUrl() string                                // function
	GetSourceUrl() string                                        // function
	HasImageContents() bool                                      // function
	GetTitleText() string                                        // function
	GetPageUrl() string                                          // function
	GetFrameUrl() string                                         // function
	GetFrameCharset() string                                     // function
	GetMediaType() cefTypes.TCefContextMenuMediaType             // function
	GetMediaStateFlags() cefTypes.TCefContextMenuMediaStateFlags // function
	GetSelectionText() string                                    // function
	GetMisspelledWord() string                                   // function
	GetDictionarySuggestions(suggestions lcl.IStringList) bool   // function
	IsEditable() bool                                            // function
	IsSpellCheckEnabled() bool                                   // function
	GetEditStateFlags() cefTypes.TCefContextMenuEditStateFlags   // function
	IsCustomMenu() bool                                          // function
}

// ICefContextMenuParamsRef Parent: ICefContextMenuParams ICefBaseRefCountedRef
type ICefContextMenuParamsRef interface {
	ICefContextMenuParams
	ICefBaseRefCountedRef
	AsIntfContextMenuParams() uintptr
}

type TCefContextMenuParamsRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefContextMenuParamsRef) GetXCoord() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TCefContextMenuParamsRef) GetYCoord() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCefContextMenuParamsRef) GetTypeFlags() cefTypes.TCefContextMenuTypeFlags {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(3, m.Instance())
	return cefTypes.TCefContextMenuTypeFlags(r)
}

func (m *TCefContextMenuParamsRef) GetLinkUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetUnfilteredLinkUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetSourceUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) HasImageContents() bool {
	if !m.IsValid() {
		return false
	}
	r := cefContextMenuParamsRefAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCefContextMenuParamsRef) GetTitleText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetPageUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetFrameUrl() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetFrameCharset() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetMediaType() cefTypes.TCefContextMenuMediaType {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(12, m.Instance())
	return cefTypes.TCefContextMenuMediaType(r)
}

func (m *TCefContextMenuParamsRef) GetMediaStateFlags() cefTypes.TCefContextMenuMediaStateFlags {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(13, m.Instance())
	return cefTypes.TCefContextMenuMediaStateFlags(r)
}

func (m *TCefContextMenuParamsRef) GetSelectionText() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetMisspelledWord() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefContextMenuParamsRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefContextMenuParamsRef) GetDictionarySuggestions(suggestions lcl.IStringList) bool {
	if !m.IsValid() {
		return false
	}
	r := cefContextMenuParamsRefAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(suggestions))
	return api.GoBool(r)
}

func (m *TCefContextMenuParamsRef) IsEditable() bool {
	if !m.IsValid() {
		return false
	}
	r := cefContextMenuParamsRefAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TCefContextMenuParamsRef) IsSpellCheckEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := cefContextMenuParamsRefAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TCefContextMenuParamsRef) GetEditStateFlags() cefTypes.TCefContextMenuEditStateFlags {
	if !m.IsValid() {
		return 0
	}
	r := cefContextMenuParamsRefAPI().SysCallN(19, m.Instance())
	return cefTypes.TCefContextMenuEditStateFlags(r)
}

func (m *TCefContextMenuParamsRef) IsCustomMenu() bool {
	if !m.IsValid() {
		return false
	}
	r := cefContextMenuParamsRefAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCefContextMenuParamsRef) AsIntfContextMenuParams() uintptr {
	return m.GetIntfPointer(0)
}

// ContextMenuParamsRef  is static instance
var ContextMenuParamsRef _ContextMenuParamsRefClass

// _ContextMenuParamsRefClass is class type defined by TCefContextMenuParamsRef
type _ContextMenuParamsRefClass uintptr

func (_ContextMenuParamsRefClass) UnWrap(data uintptr) (result ICefContextMenuParams) {
	var resultPtr uintptr
	cefContextMenuParamsRefAPI().SysCallN(21, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefContextMenuParamsRef(resultPtr)
	return
}

// NewContextMenuParamsRef class constructor
func NewContextMenuParamsRef(data uintptr) ICefContextMenuParamsRef {
	var contextMenuParamsPtr uintptr // ICefContextMenuParams
	r := cefContextMenuParamsRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&contextMenuParamsPtr)))
	ret := AsCefContextMenuParamsRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, contextMenuParamsPtr)
	}
	return ret
}

var (
	cefContextMenuParamsRefOnce   base.Once
	cefContextMenuParamsRefImport *imports.Imports = nil
)

func cefContextMenuParamsRefAPI() *imports.Imports {
	cefContextMenuParamsRefOnce.Do(func() {
		cefContextMenuParamsRefImport = api.NewDefaultImports()
		cefContextMenuParamsRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefContextMenuParamsRef_Create", 0), // constructor NewContextMenuParamsRef
			/* 1 */ imports.NewTable("TCefContextMenuParamsRef_GetXCoord", 0), // function GetXCoord
			/* 2 */ imports.NewTable("TCefContextMenuParamsRef_GetYCoord", 0), // function GetYCoord
			/* 3 */ imports.NewTable("TCefContextMenuParamsRef_GetTypeFlags", 0), // function GetTypeFlags
			/* 4 */ imports.NewTable("TCefContextMenuParamsRef_GetLinkUrl", 0), // function GetLinkUrl
			/* 5 */ imports.NewTable("TCefContextMenuParamsRef_GetUnfilteredLinkUrl", 0), // function GetUnfilteredLinkUrl
			/* 6 */ imports.NewTable("TCefContextMenuParamsRef_GetSourceUrl", 0), // function GetSourceUrl
			/* 7 */ imports.NewTable("TCefContextMenuParamsRef_HasImageContents", 0), // function HasImageContents
			/* 8 */ imports.NewTable("TCefContextMenuParamsRef_GetTitleText", 0), // function GetTitleText
			/* 9 */ imports.NewTable("TCefContextMenuParamsRef_GetPageUrl", 0), // function GetPageUrl
			/* 10 */ imports.NewTable("TCefContextMenuParamsRef_GetFrameUrl", 0), // function GetFrameUrl
			/* 11 */ imports.NewTable("TCefContextMenuParamsRef_GetFrameCharset", 0), // function GetFrameCharset
			/* 12 */ imports.NewTable("TCefContextMenuParamsRef_GetMediaType", 0), // function GetMediaType
			/* 13 */ imports.NewTable("TCefContextMenuParamsRef_GetMediaStateFlags", 0), // function GetMediaStateFlags
			/* 14 */ imports.NewTable("TCefContextMenuParamsRef_GetSelectionText", 0), // function GetSelectionText
			/* 15 */ imports.NewTable("TCefContextMenuParamsRef_GetMisspelledWord", 0), // function GetMisspelledWord
			/* 16 */ imports.NewTable("TCefContextMenuParamsRef_GetDictionarySuggestions", 0), // function GetDictionarySuggestions
			/* 17 */ imports.NewTable("TCefContextMenuParamsRef_IsEditable", 0), // function IsEditable
			/* 18 */ imports.NewTable("TCefContextMenuParamsRef_IsSpellCheckEnabled", 0), // function IsSpellCheckEnabled
			/* 19 */ imports.NewTable("TCefContextMenuParamsRef_GetEditStateFlags", 0), // function GetEditStateFlags
			/* 20 */ imports.NewTable("TCefContextMenuParamsRef_IsCustomMenu", 0), // function IsCustomMenu
			/* 21 */ imports.NewTable("TCefContextMenuParamsRef_UnWrap", 0), // static function UnWrap
		}
	})
	return cefContextMenuParamsRefImport
}
