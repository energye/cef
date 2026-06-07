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

	cefTypes "github.com/energye/cef/types"
)

// ICefSchemeRegistrarRef Parent: ICEFBaseScopedWrapperRef
type ICefSchemeRegistrarRef interface {
	ICEFBaseScopedWrapperRef
	AddCustomScheme(schemeName string, options cefTypes.TCefSchemeOptions) bool // function
}

type TCefSchemeRegistrarRef struct {
	TCEFBaseScopedWrapperRef
}

func (m *TCefSchemeRegistrarRef) AddCustomScheme(schemeName string, options cefTypes.TCefSchemeOptions) bool {
	if !m.IsValid() {
		return false
	}
	r := cefSchemeRegistrarRefAPI().SysCallN(1, m.Instance(), api.PasStr(schemeName), uintptr(options))
	return api.GoBool(r)
}

// NewSchemeRegistrarRef class constructor
func NewSchemeRegistrarRef(data uintptr) ICefSchemeRegistrarRef {
	r := cefSchemeRegistrarRefAPI().SysCallN(0, uintptr(data))
	return AsCefSchemeRegistrarRef(r)
}

var (
	cefSchemeRegistrarRefOnce   base.Once
	cefSchemeRegistrarRefImport *imports.Imports = nil
)

func cefSchemeRegistrarRefAPI() *imports.Imports {
	cefSchemeRegistrarRefOnce.Do(func() {
		cefSchemeRegistrarRefImport = api.NewDefaultImports()
		cefSchemeRegistrarRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefSchemeRegistrarRef_Create", 0), // constructor NewSchemeRegistrarRef
			/* 1 */ imports.NewTable("TCefSchemeRegistrarRef_AddCustomScheme", 0), // function AddCustomScheme
		}
	})
	return cefSchemeRegistrarRefImport
}
