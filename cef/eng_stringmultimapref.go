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

// ICefStringMultimapRef Parent: ICefStringMultimap
type ICefStringMultimapRef interface {
	ICefStringMultimap
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringMultimap // procedure
}

// TCefStringMultimapRef Parent: TCefStringMultimap
type TCefStringMultimapRef struct {
	TCefStringMultimap
}

func NewCefStringMultimapRef(aHandle TCefStringMultimapHandle) ICefStringMultimapRef {
	r1 := stringMultimapRefImportAPI().SysCallN(1, uintptr(aHandle))
	return AsCefStringMultimapRef(r1)
}

func (m *TCefStringMultimapRef) AsInterface() ICefStringMultimap {
	var resultCefStringMultimap uintptr
	stringMultimapRefImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringMultimap)))
	return AsCefStringMultimap(resultCefStringMultimap)
}

var (
	stringMultimapRefImport       *imports.Imports = nil
	stringMultimapRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMultimapRef_AsInterface", 0),
		/*1*/ imports.NewTable("CefStringMultimapRef_Create", 0),
	}
)

func stringMultimapRefImportAPI() *imports.Imports {
	if stringMultimapRefImport == nil {
		stringMultimapRefImport = NewDefaultImports()
		stringMultimapRefImport.SetImportTable(stringMultimapRefImportTables)
		stringMultimapRefImportTables = nil
	}
	return stringMultimapRefImport
}
