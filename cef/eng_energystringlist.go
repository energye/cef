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

// IEnergyStringList Parent: ICefCustomStringList
//
//	CEF string maps are a set of key/value string pairs.
type IEnergyStringList interface {
	ICefCustomStringList
	// AsInterface
	//  Class instance to interface instance
	AsInterface() ICefStringList // procedure
}

// TEnergyStringList Parent: TCefCustomStringList
//
//	CEF string maps are a set of key/value string pairs.
type TEnergyStringList struct {
	TCefCustomStringList
}

func NewEnergyStringList() IEnergyStringList {
	r1 := energyStringListImportAPI().SysCallN(1)
	return AsEnergyStringList(r1)
}

func (m *TEnergyStringList) AsInterface() ICefStringList {
	var resultCefStringList uintptr
	energyStringListImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCefStringList)))
	return AsCefStringList(resultCefStringList)
}

var (
	energyStringListImport       *imports.Imports = nil
	energyStringListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("EnergyStringList_AsInterface", 0),
		/*1*/ imports.NewTable("EnergyStringList_Create", 0),
	}
)

func energyStringListImportAPI() *imports.Imports {
	if energyStringListImport == nil {
		energyStringListImport = NewDefaultImports()
		energyStringListImport.SetImportTable(energyStringListImportTables)
		energyStringListImportTables = nil
	}
	return energyStringListImport
}
