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

// ICEFBaseScopedWrapperRef Parent: IObject
type ICEFBaseScopedWrapperRef interface {
	IObject
	Wrap() uintptr // function
}

// TCEFBaseScopedWrapperRef Parent: TObject
type TCEFBaseScopedWrapperRef struct {
	TObject
}

func NewCEFBaseScopedWrapperRef(data uintptr) ICEFBaseScopedWrapperRef {
	r1 := baseScopedWrapperRefImportAPI().SysCallN(0, uintptr(data))
	return AsCEFBaseScopedWrapperRef(r1)
}

func (m *TCEFBaseScopedWrapperRef) Wrap() uintptr {
	r1 := baseScopedWrapperRefImportAPI().SysCallN(1, m.Instance())
	return uintptr(r1)
}

var (
	baseScopedWrapperRefImport       *imports.Imports = nil
	baseScopedWrapperRefImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CEFBaseScopedWrapperRef_Create", 0),
		/*1*/ imports.NewTable("CEFBaseScopedWrapperRef_Wrap", 0),
	}
)

func baseScopedWrapperRefImportAPI() *imports.Imports {
	if baseScopedWrapperRefImport == nil {
		baseScopedWrapperRefImport = NewDefaultImports()
		baseScopedWrapperRefImport.SetImportTable(baseScopedWrapperRefImportTables)
		baseScopedWrapperRefImportTables = nil
	}
	return baseScopedWrapperRefImport
}
