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

// ICustomResponseFilter Parent: ICefResponseFilter
type ICustomResponseFilter interface {
	ICefResponseFilter
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefResponseFilter  // procedure
	SetOnFilter(fn TOnFilter)         // property event
	SetOnInitFilter(fn TOnInitFilter) // property event
}

// TCustomResponseFilter Parent: TCefResponseFilter
type TCustomResponseFilter struct {
	TCefResponseFilter
	filterPtr     uintptr
	initFilterPtr uintptr
}

func NewCustomResponseFilter() ICustomResponseFilter {
	r1 := customResponseFilterImportAPI().SysCallN(1)
	return AsCustomResponseFilter(r1)
}

func (m *TCustomResponseFilter) AsInterface() ICefResponseFilter {
	var resultCefResponseFilter uintptr
	customResponseFilterImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefResponseFilter)))
	return AsCefResponseFilter(resultCefResponseFilter)
}

func (m *TCustomResponseFilter) SetOnFilter(fn TOnFilter) {
	if m.filterPtr != 0 {
		RemoveEventElement(m.filterPtr)
	}
	m.filterPtr = MakeEventDataPtr(fn)
	customResponseFilterImportAPI().SysCallN(2, m.Instance(), m.filterPtr)
}

func (m *TCustomResponseFilter) SetOnInitFilter(fn TOnInitFilter) {
	if m.initFilterPtr != 0 {
		RemoveEventElement(m.initFilterPtr)
	}
	m.initFilterPtr = MakeEventDataPtr(fn)
	customResponseFilterImportAPI().SysCallN(3, m.Instance(), m.initFilterPtr)
}

var (
	customResponseFilterImport       *imports.Imports = nil
	customResponseFilterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomResponseFilter_AsInterface", 0),
		/*1*/ imports.NewTable("CustomResponseFilter_Create", 0),
		/*2*/ imports.NewTable("CustomResponseFilter_SetOnFilter", 0),
		/*3*/ imports.NewTable("CustomResponseFilter_SetOnInitFilter", 0),
	}
)

func customResponseFilterImportAPI() *imports.Imports {
	if customResponseFilterImport == nil {
		customResponseFilterImport = NewDefaultImports()
		customResponseFilterImport.SetImportTable(customResponseFilterImportTables)
		customResponseFilterImportTables = nil
	}
	return customResponseFilterImport
}
