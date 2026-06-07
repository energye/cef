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
)

// ICefPreferenceRegistrarRef Parent: ICEFBaseScopedWrapperRef
type ICefPreferenceRegistrarRef interface {
	ICEFBaseScopedWrapperRef
	AddPreference(name string, defaultValue ICefValue) bool // function
}

type TCefPreferenceRegistrarRef struct {
	TCEFBaseScopedWrapperRef
}

func (m *TCefPreferenceRegistrarRef) AddPreference(name string, defaultValue ICefValue) bool {
	if !m.IsValid() {
		return false
	}
	r := cefPreferenceRegistrarRefAPI().SysCallN(1, m.Instance(), api.PasStr(name), base.GetObjectUintptr(defaultValue))
	return api.GoBool(r)
}

// NewPreferenceRegistrarRef class constructor
func NewPreferenceRegistrarRef(data uintptr) ICefPreferenceRegistrarRef {
	r := cefPreferenceRegistrarRefAPI().SysCallN(0, uintptr(data))
	return AsCefPreferenceRegistrarRef(r)
}

var (
	cefPreferenceRegistrarRefOnce   base.Once
	cefPreferenceRegistrarRefImport *imports.Imports = nil
)

func cefPreferenceRegistrarRefAPI() *imports.Imports {
	cefPreferenceRegistrarRefOnce.Do(func() {
		cefPreferenceRegistrarRefImport = api.NewDefaultImports()
		cefPreferenceRegistrarRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPreferenceRegistrarRef_Create", 0), // constructor NewPreferenceRegistrarRef
			/* 1 */ imports.NewTable("TCefPreferenceRegistrarRef_AddPreference", 0), // function AddPreference
		}
	})
	return cefPreferenceRegistrarRefImport
}
