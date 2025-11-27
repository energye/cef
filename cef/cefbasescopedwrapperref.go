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
)

// ICEFBaseScopedWrapperRef Parent: lcl.IObject
type ICEFBaseScopedWrapperRef interface {
	lcl.IObject
	Wrap() uintptr // function
}

type TCEFBaseScopedWrapperRef struct {
	lcl.TObject
}

func (m *TCEFBaseScopedWrapperRef) Wrap() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := cEFBaseScopedWrapperRefAPI().SysCallN(1, m.Instance())
	return uintptr(r)
}

// NewBaseScopedWrapperRef class constructor
func NewBaseScopedWrapperRef(data uintptr) ICEFBaseScopedWrapperRef {
	r := cEFBaseScopedWrapperRefAPI().SysCallN(0, uintptr(data))
	return AsCEFBaseScopedWrapperRef(r)
}

var (
	cEFBaseScopedWrapperRefOnce   base.Once
	cEFBaseScopedWrapperRefImport *imports.Imports = nil
)

func cEFBaseScopedWrapperRefAPI() *imports.Imports {
	cEFBaseScopedWrapperRefOnce.Do(func() {
		cEFBaseScopedWrapperRefImport = api.NewDefaultImports()
		cEFBaseScopedWrapperRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFBaseScopedWrapperRef_Create", 0), // constructor NewBaseScopedWrapperRef
			/* 1 */ imports.NewTable("TCEFBaseScopedWrapperRef_Wrap", 0), // function Wrap
		}
	})
	return cEFBaseScopedWrapperRefImport
}
