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

// ICefResourceSkipCallback Parent: ICefBaseRefCounted
type ICefResourceSkipCallback interface {
	ICefBaseRefCounted
	Cont(bytesskipped int64) // procedure
}

// TCefResourceSkipCallback Parent: TCefBaseRefCounted
type TCefResourceSkipCallback struct {
	TCefBaseRefCounted
}

// ResourceSkipCallbackRef -> ICefResourceSkipCallback
var ResourceSkipCallbackRef resourceSkipCallback

// resourceSkipCallback TCefResourceSkipCallback Ref
type resourceSkipCallback uintptr

func (m *resourceSkipCallback) UnWrap(data uintptr) ICefResourceSkipCallback {
	var resultCefResourceSkipCallback uintptr
	resourceSkipCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefResourceSkipCallback)))
	return AsCefResourceSkipCallback(resultCefResourceSkipCallback)
}

func (m *TCefResourceSkipCallback) Cont(bytesskipped int64) {
	resourceSkipCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&bytesskipped)))
}

var (
	resourceSkipCallbackImport       *imports.Imports = nil
	resourceSkipCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefResourceSkipCallback_Cont", 0),
		/*1*/ imports.NewTable("CefResourceSkipCallback_UnWrap", 0),
	}
)

func resourceSkipCallbackImportAPI() *imports.Imports {
	if resourceSkipCallbackImport == nil {
		resourceSkipCallbackImport = NewDefaultImports()
		resourceSkipCallbackImport.SetImportTable(resourceSkipCallbackImportTables)
		resourceSkipCallbackImportTables = nil
	}
	return resourceSkipCallbackImport
}
