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

// ICefX509CertPrincipal Parent: ICefBaseRefCountedRef
type ICefX509CertPrincipal interface {
	ICefBaseRefCountedRef
	// GetDisplayName
	//  Returns a name that can be used to represent the issuer. It tries in this
	//  order: Common Name (CN), Organization Name (O) and Organizational Unit
	//  Name (OU) and returns the first non-NULL one found.
	GetDisplayName() string // function
	// GetCommonName
	//  Returns the common name.
	GetCommonName() string // function
	// GetLocalityName
	//  Returns the locality name.
	GetLocalityName() string // function
	// GetStateOrProvinceName
	//  Returns the state or province name.
	GetStateOrProvinceName() string // function
	// GetCountryName
	//  Returns the country name.
	GetCountryName() string // function
	// GetOrganizationNames
	//  Retrieve the list of organization names.
	GetOrganizationNames(names lcl.IStrings) // procedure
	// GetOrganizationUnitNames
	//  Retrieve the list of organization unit names.
	GetOrganizationUnitNames(names lcl.IStrings) // procedure
}

// ICefX509CertPrincipalRef Parent: ICefX509CertPrincipal
type ICefX509CertPrincipalRef interface {
	ICefX509CertPrincipal
	AsIntfX509CertPrincipal() uintptr
}

type TCefX509CertPrincipalRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefX509CertPrincipalRef) GetDisplayName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefX509CertPrincipalRefAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCefX509CertPrincipalRef) GetCommonName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefX509CertPrincipalRefAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TCefX509CertPrincipalRef) GetLocalityName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefX509CertPrincipalRefAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCefX509CertPrincipalRef) GetStateOrProvinceName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefX509CertPrincipalRefAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCefX509CertPrincipalRef) GetCountryName() string {
	if !m.IsValid() {
		return ""
	}
	r := cefX509CertPrincipalRefAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCefX509CertPrincipalRef) GetOrganizationNames(names lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(names))
}

func (m *TCefX509CertPrincipalRef) GetOrganizationUnitNames(names lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(names))
}

func (m *TCefX509CertPrincipalRef) AsIntfX509CertPrincipal() uintptr {
	return m.GetIntfPointer(0)
}

// X509CertPrincipalRef  is static instance
var X509CertPrincipalRef _X509CertPrincipalRefClass

// _X509CertPrincipalRefClass is class type defined by TCefX509CertPrincipalRef
type _X509CertPrincipalRefClass uintptr

func (_X509CertPrincipalRefClass) UnWrap(data uintptr) (result ICefX509CertPrincipal) {
	var resultPtr uintptr
	cefX509CertPrincipalRefAPI().SysCallN(6, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefX509CertPrincipalRef(resultPtr)
	return
}

// NewX509CertPrincipalRef class constructor
func NewX509CertPrincipalRef(data uintptr) ICefX509CertPrincipalRef {
	var x509CertPrincipalPtr uintptr // ICefX509CertPrincipal
	r := cefX509CertPrincipalRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&x509CertPrincipalPtr)))
	ret := AsCefX509CertPrincipalRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, x509CertPrincipalPtr)
	}
	return ret
}

var (
	cefX509CertPrincipalRefOnce   base.Once
	cefX509CertPrincipalRefImport *imports.Imports = nil
)

func cefX509CertPrincipalRefAPI() *imports.Imports {
	cefX509CertPrincipalRefOnce.Do(func() {
		cefX509CertPrincipalRefImport = api.NewDefaultImports()
		cefX509CertPrincipalRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefX509CertPrincipalRef_Create", 0), // constructor NewX509CertPrincipalRef
			/* 1 */ imports.NewTable("TCefX509CertPrincipalRef_GetDisplayName", 0), // function GetDisplayName
			/* 2 */ imports.NewTable("TCefX509CertPrincipalRef_GetCommonName", 0), // function GetCommonName
			/* 3 */ imports.NewTable("TCefX509CertPrincipalRef_GetLocalityName", 0), // function GetLocalityName
			/* 4 */ imports.NewTable("TCefX509CertPrincipalRef_GetStateOrProvinceName", 0), // function GetStateOrProvinceName
			/* 5 */ imports.NewTable("TCefX509CertPrincipalRef_GetCountryName", 0), // function GetCountryName
			/* 6 */ imports.NewTable("TCefX509CertPrincipalRef_UnWrap", 0), // static function UnWrap
			/* 7 */ imports.NewTable("TCefX509CertPrincipalRef_GetOrganizationNames", 0), // procedure GetOrganizationNames
			/* 8 */ imports.NewTable("TCefX509CertPrincipalRef_GetOrganizationUnitNames", 0), // procedure GetOrganizationUnitNames
		}
	})
	return cefX509CertPrincipalRefImport
}
