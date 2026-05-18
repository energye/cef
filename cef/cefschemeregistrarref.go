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
	// AddCustomScheme
	//  Register a custom scheme. This function should not be called for the
	//  built-in HTTP, HTTPS, FILE, FTP, ABOUT and DATA schemes.
	//  This function may be called on any thread. It should only be called once
	//  per unique |scheme_name| value. If |scheme_name| is already registered or
	//  if an error occurs this function will return false (0).
	//  <see>See the CEF_SCHEME_OPTION_* constants in the uCEFConstants unit for possible values for |options|.</see>
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
