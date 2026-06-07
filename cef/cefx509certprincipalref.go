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

// ICefX509CertPrincipal Parent: ICefBaseRefCounted
type ICefX509CertPrincipal interface {
	ICefBaseRefCounted
	GetDisplayName() string                      // function
	GetCommonName() string                       // function
	GetLocalityName() string                     // function
	GetStateOrProvinceName() string              // function
	GetCountryName() string                      // function
	GetStreetAddresses(addresses lcl.IStrings)   // procedure
	GetOrganizationNames(names lcl.IStrings)     // procedure
	GetOrganizationUnitNames(names lcl.IStrings) // procedure
	GetDomainComponents(components lcl.IStrings) // procedure
}

// ICefX509CertPrincipalRef Parent: ICefX509CertPrincipal ICefBaseRefCountedRef
type ICefX509CertPrincipalRef interface {
	ICefX509CertPrincipal
	ICefBaseRefCountedRef
	AsIntfX509CertPrincipal() uintptr
}

type TCefX509CertPrincipalRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefX509CertPrincipalRef) GetDisplayName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefX509CertPrincipalRefAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefX509CertPrincipalRef) GetCommonName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefX509CertPrincipalRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefX509CertPrincipalRef) GetLocalityName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefX509CertPrincipalRefAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefX509CertPrincipalRef) GetStateOrProvinceName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefX509CertPrincipalRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefX509CertPrincipalRef) GetCountryName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefX509CertPrincipalRefAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefX509CertPrincipalRef) GetStreetAddresses(addresses lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(addresses))
}

func (m *TCefX509CertPrincipalRef) GetOrganizationNames(names lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(names))
}

func (m *TCefX509CertPrincipalRef) GetOrganizationUnitNames(names lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(names))
}

func (m *TCefX509CertPrincipalRef) GetDomainComponents(components lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	cefX509CertPrincipalRefAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(components))
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
			/* 7 */ imports.NewTable("TCefX509CertPrincipalRef_GetStreetAddresses", 0), // procedure GetStreetAddresses
			/* 8 */ imports.NewTable("TCefX509CertPrincipalRef_GetOrganizationNames", 0), // procedure GetOrganizationNames
			/* 9 */ imports.NewTable("TCefX509CertPrincipalRef_GetOrganizationUnitNames", 0), // procedure GetOrganizationUnitNames
			/* 10 */ imports.NewTable("TCefX509CertPrincipalRef_GetDomainComponents", 0), // procedure GetDomainComponents
		}
	})
	return cefX509CertPrincipalRefImport
}
