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

// ICefResourceReadCallback Parent: ICefBaseRefCounted
type ICefResourceReadCallback interface {
	ICefBaseRefCounted
	Cont(bytesread int64) // procedure
}

// TCefResourceReadCallback Parent: TCefBaseRefCounted
type TCefResourceReadCallback struct {
	TCefBaseRefCounted
}

// ResourceReadCallbackRef -> ICefResourceReadCallback
var ResourceReadCallbackRef resourceReadCallback

// resourceReadCallback TCefResourceReadCallback Ref
type resourceReadCallback uintptr

func (m *resourceReadCallback) UnWrap(data uintptr) ICefResourceReadCallback {
	var resultCefResourceReadCallback uintptr
	resourceReadCallbackImportAPI().SysCallN(1, uintptr(data), uintptr(unsafePointer(&resultCefResourceReadCallback)))
	return AsCefResourceReadCallback(resultCefResourceReadCallback)
}

func (m *TCefResourceReadCallback) Cont(bytesread int64) {
	resourceReadCallbackImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&bytesread)))
}

var (
	resourceReadCallbackImport       *imports.Imports = nil
	resourceReadCallbackImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefResourceReadCallback_Cont", 0),
		/*1*/ imports.NewTable("CefResourceReadCallback_UnWrap", 0),
	}
)

func resourceReadCallbackImportAPI() *imports.Imports {
	if resourceReadCallbackImport == nil {
		resourceReadCallbackImport = NewDefaultImports()
		resourceReadCallbackImport.SetImportTable(resourceReadCallbackImportTables)
		resourceReadCallbackImportTables = nil
	}
	return resourceReadCallbackImport
}
