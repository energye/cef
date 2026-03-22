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

// ICefPreferenceManager Parent: ICefBaseRefCounted
type ICefPreferenceManager interface {
	ICefBaseRefCounted
	// HasPreference
	//  Returns true (1) if a preference with the specified |name| exists. This
	//  function must be called on the browser process UI thread.
	HasPreference(name string) bool // function
	// GetPreference
	//  Returns the value for the preference with the specified |name|. Returns
	//  NULL if the preference does not exist. The returned object contains a copy
	//  of the underlying preference value and modifications to the returned
	//  object will not modify the underlying preference value. This function must
	//  be called on the browser process UI thread.
	GetPreference(name string) ICefValue // function
	// GetAllPreferences
	//  Returns all preferences as a dictionary. If |include_defaults| is true (1)
	//  then preferences currently at their default value will be included. The
	//  returned object contains a copy of the underlying preference values and
	//  modifications to the returned object will not modify the underlying
	//  preference values. This function must be called on the browser process UI
	//  thread.
	GetAllPreferences(includeDefaults bool) ICefDictionaryValue // function
	// CanSetPreference
	//  Returns true (1) if the preference with the specified |name| can be
	//  modified using SetPreference. As one example preferences set via the
	//  command-line usually cannot be modified. This function must be called on
	//  the browser process UI thread.
	CanSetPreference(name string) bool // function
	// SetPreference
	//  Set the |value| associated with preference |name|. Returns true (1) if the
	//  value is set successfully and false (0) otherwise. If |value| is NULL the
	//  preference will be restored to its default value. If setting the
	//  preference fails then |error| will be populated with a detailed
	//  description of the problem. This function must be called on the browser
	//  process UI thread.
	SetPreference(name string, value ICefValue, outError_ *string) bool // function
}

// ICefPreferenceManagerRef Parent: ICefPreferenceManager ICefBaseRefCountedRef
type ICefPreferenceManagerRef interface {
	ICefPreferenceManager
	ICefBaseRefCountedRef
	AsIntfPreferenceManager() uintptr
}

type TCefPreferenceManagerRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefPreferenceManagerRef) HasPreference(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefPreferenceManagerRefAPI().SysCallN(1, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCefPreferenceManagerRef) GetPreference(name string) (result ICefValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPreferenceManagerRefAPI().SysCallN(2, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefValueRef(resultPtr)
	return
}

func (m *TCefPreferenceManagerRef) GetAllPreferences(includeDefaults bool) (result ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	cefPreferenceManagerRefAPI().SysCallN(3, m.Instance(), api.PasBool(includeDefaults), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefDictionaryValueRef(resultPtr)
	return
}

func (m *TCefPreferenceManagerRef) CanSetPreference(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefPreferenceManagerRefAPI().SysCallN(4, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCefPreferenceManagerRef) SetPreference(name string, value ICefValue, outError_ *string) bool {
	if !m.IsValid() {
		return false
	}
	var error_Ptr uintptr
	r := cefPreferenceManagerRefAPI().SysCallN(5, m.Instance(), api.PasStr(name), base.GetObjectUintptr(value), uintptr(base.UnsafePointer(&error_Ptr)))
	*outError_ = api.GoStr(error_Ptr)
	return api.GoBool(r)
}

func (m *TCefPreferenceManagerRef) AsIntfPreferenceManager() uintptr {
	return m.GetIntfPointer(0)
}

// PreferenceManagerRef  is static instance
var PreferenceManagerRef _PreferenceManagerRefClass

// _PreferenceManagerRefClass is class type defined by TCefPreferenceManagerRef
type _PreferenceManagerRefClass uintptr

func (_PreferenceManagerRefClass) UnWrapWithPointer(data uintptr) (result ICefPreferenceManager) {
	var resultPtr uintptr
	cefPreferenceManagerRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPreferenceManagerRef(resultPtr)
	return
}

func (_PreferenceManagerRefClass) GlobalToPreferenceManager() (result ICefPreferenceManager) {
	var resultPtr uintptr
	cefPreferenceManagerRefAPI().SysCallN(7, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefPreferenceManagerRef(resultPtr)
	return
}

// NewPreferenceManagerRef class constructor
func NewPreferenceManagerRef(data uintptr) ICefPreferenceManagerRef {
	var preferenceManagerPtr uintptr // ICefPreferenceManager
	r := cefPreferenceManagerRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&preferenceManagerPtr)))
	ret := AsCefPreferenceManagerRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, preferenceManagerPtr)
	}
	return ret
}

var (
	cefPreferenceManagerRefOnce   base.Once
	cefPreferenceManagerRefImport *imports.Imports = nil
)

func cefPreferenceManagerRefAPI() *imports.Imports {
	cefPreferenceManagerRefOnce.Do(func() {
		cefPreferenceManagerRefImport = api.NewDefaultImports()
		cefPreferenceManagerRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefPreferenceManagerRef_Create", 0), // constructor NewPreferenceManagerRef
			/* 1 */ imports.NewTable("TCefPreferenceManagerRef_HasPreference", 0), // function HasPreference
			/* 2 */ imports.NewTable("TCefPreferenceManagerRef_GetPreference", 0), // function GetPreference
			/* 3 */ imports.NewTable("TCefPreferenceManagerRef_GetAllPreferences", 0), // function GetAllPreferences
			/* 4 */ imports.NewTable("TCefPreferenceManagerRef_CanSetPreference", 0), // function CanSetPreference
			/* 5 */ imports.NewTable("TCefPreferenceManagerRef_SetPreference", 0), // function SetPreference
			/* 6 */ imports.NewTable("TCefPreferenceManagerRef_UnWrapWithPointer", 0), // static function UnWrapWithPointer
			/* 7 */ imports.NewTable("TCefPreferenceManagerRef_GlobalToPreferenceManager", 0), // static function GlobalToPreferenceManager
		}
	})
	return cefPreferenceManagerRefImport
}
