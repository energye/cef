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

// ICefMediaSource Parent: ICefBaseRefCounted
type ICefMediaSource interface {
	ICefBaseRefCounted
	GetId() string      // function
	IsCastSource() bool // function
	IsDialSource() bool // function
}

// TCefMediaSource Parent: TCefBaseRefCounted
type TCefMediaSource struct {
	TCefBaseRefCounted
}

// MediaSourceRef -> ICefMediaSource
var MediaSourceRef mediaSource

// mediaSource TCefMediaSource Ref
type mediaSource uintptr

func (m *mediaSource) UnWrap(data uintptr) ICefMediaSource {
	var resultCefMediaSource uintptr
	mediaSourceImportAPI().SysCallN(3, uintptr(data), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TCefMediaSource) GetId() string {
	r1 := mediaSourceImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaSource) IsCastSource() bool {
	r1 := mediaSourceImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCefMediaSource) IsDialSource() bool {
	r1 := mediaSourceImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

var (
	mediaSourceImport       *imports.Imports = nil
	mediaSourceImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefMediaSource_GetId", 0),
		/*1*/ imports.NewTable("CefMediaSource_IsCastSource", 0),
		/*2*/ imports.NewTable("CefMediaSource_IsDialSource", 0),
		/*3*/ imports.NewTable("CefMediaSource_UnWrap", 0),
	}
)

func mediaSourceImportAPI() *imports.Imports {
	if mediaSourceImport == nil {
		mediaSourceImport = NewDefaultImports()
		mediaSourceImport.SetImportTable(mediaSourceImportTables)
		mediaSourceImportTables = nil
	}
	return mediaSourceImport
}
