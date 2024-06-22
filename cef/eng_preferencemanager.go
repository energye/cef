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

// ICefPreferenceManager Parent: ICefBaseRefCounted
type ICefPreferenceManager interface {
	ICefBaseRefCounted
	HasPreference(name string) bool                                    // function
	GetPreference(name string) ICefValue                               // function
	GetAllPreferences(includeDefaults bool) ICefDictionaryValue        // function
	CanSetPreference(name string) bool                                 // function
	SetPreference(name string, value ICefValue, outError *string) bool // function
}

// TCefPreferenceManager Parent: TCefBaseRefCounted
type TCefPreferenceManager struct {
	TCefBaseRefCounted
}

// PreferenceManagerRef -> ICefPreferenceManager
var PreferenceManagerRef preferenceManager

// preferenceManager TCefPreferenceManager Ref
type preferenceManager uintptr

func (m *preferenceManager) UnWrap(data uintptr) ICefPreferenceManager {
	var resultCefPreferenceManager uintptr
	preferenceManagerImportAPI().SysCallN(6, uintptr(data), uintptr(unsafePointer(&resultCefPreferenceManager)))
	return AsCefPreferenceManager(resultCefPreferenceManager)
}

func (m *preferenceManager) Global() ICefPreferenceManager {
	var resultCefPreferenceManager uintptr
	preferenceManagerImportAPI().SysCallN(3, uintptr(unsafePointer(&resultCefPreferenceManager)))
	return AsCefPreferenceManager(resultCefPreferenceManager)
}

func (m *TCefPreferenceManager) HasPreference(name string) bool {
	r1 := preferenceManagerImportAPI().SysCallN(4, m.Instance(), PascalStr(name))
	return GoBool(r1)
}

func (m *TCefPreferenceManager) GetPreference(name string) ICefValue {
	var resultCefValue uintptr
	preferenceManagerImportAPI().SysCallN(2, m.Instance(), PascalStr(name), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefPreferenceManager) GetAllPreferences(includeDefaults bool) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	preferenceManagerImportAPI().SysCallN(1, m.Instance(), PascalBool(includeDefaults), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefPreferenceManager) CanSetPreference(name string) bool {
	r1 := preferenceManagerImportAPI().SysCallN(0, m.Instance(), PascalStr(name))
	return GoBool(r1)
}

func (m *TCefPreferenceManager) SetPreference(name string, value ICefValue, outError *string) bool {
	var result2 uintptr
	r1 := preferenceManagerImportAPI().SysCallN(5, m.Instance(), PascalStr(name), GetObjectUintptr(value), uintptr(unsafePointer(&result2)))
	*outError = GoStr(result2)
	return GoBool(r1)
}

var (
	preferenceManagerImport       *imports.Imports = nil
	preferenceManagerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CefPreferenceManager_CanSetPreference", 0),
		/*1*/ imports.NewTable("CefPreferenceManager_GetAllPreferences", 0),
		/*2*/ imports.NewTable("CefPreferenceManager_GetPreference", 0),
		/*3*/ imports.NewTable("CefPreferenceManager_Global", 0),
		/*4*/ imports.NewTable("CefPreferenceManager_HasPreference", 0),
		/*5*/ imports.NewTable("CefPreferenceManager_SetPreference", 0),
		/*6*/ imports.NewTable("CefPreferenceManager_UnWrap", 0),
	}
)

func preferenceManagerImportAPI() *imports.Imports {
	if preferenceManagerImport == nil {
		preferenceManagerImport = NewDefaultImports()
		preferenceManagerImport.SetImportTable(preferenceManagerImportTables)
		preferenceManagerImportTables = nil
	}
	return preferenceManagerImport
}
