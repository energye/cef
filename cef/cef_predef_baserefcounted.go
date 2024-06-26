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

func (m *TCefBaseRefCounted) FreeAndNil() {
	var data = m.Instance()
	predefBaseRefCountedImportAPI().SysCallN(0, uintptr(unsafePointer(&data)))
	m.SetInstance(nil)
}

var (
	predefBaseRefCountedImport       *imports.Imports = nil
	predefBaseRefCountedImportTables                  = []*imports.Table{
		imports.NewTable("CefBaseRefCounted_FreeAndNil", 0),
	}
)

func predefBaseRefCountedImportAPI() *imports.Imports {
	if predefBaseRefCountedImport == nil {
		predefBaseRefCountedImport = NewDefaultImports()
		predefBaseRefCountedImport.SetImportTable(predefBaseRefCountedImportTables)
		predefBaseRefCountedImportTables = nil
	}
	return predefBaseRefCountedImport
}
