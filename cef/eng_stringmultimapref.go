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
}

// TCefStringMultimapRef Parent: TCefStringMultimap
type TCefStringMultimapRef struct {
	TCefStringMultimap
}

func NewCefStringMultimapRef(aHandle TCefStringMultimapHandle) ICefStringMultimapRef {
	r1 := stringMultimapRefImportAPI().SysCallN(0, uintptr(aHandle))
	return AsCefStringMultimapRef(r1)
}

var (
	stringMultimapRefImport       *imports.Imports = nil
	stringMultimapRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefStringMultimapRef_Create", 0),
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
